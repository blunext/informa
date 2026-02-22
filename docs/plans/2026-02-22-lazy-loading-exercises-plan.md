# Lazy Loading Exercises + Coaching — Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Replace `exercise get` (returns everything) with 3 lazy-loading commands (question/hints/answer) plus a `coaching` field computed from student progress.

**Architecture:** CLI keeps `ExerciseOut` + `queryExercises` + `scanExercise` internally unchanged. New output types (`QuestionOut`, `HintsOut`, `AnswerOut`) map from `ExerciseOut`. New `buildCoaching()` function queries `progress.db` for student context. `exercise review` and `exercise next` return `QuestionOut` instead of `ExerciseOut`.

**Tech Stack:** Go, SQLite (modernc.org/sqlite), cobra CLI, FSRS-5 spaced repetition

**Design doc:** `docs/plans/2026-02-22-lazy-loading-exercises-design.md`

---

## Task 1: Add new output types to types.go

**Files:**
- Modify: `analiza/cli/types.go:89-105`

**Step 1: Write failing test**

Add to `main_test.go`:

```go
func TestQuestionOutHasNoAnswer(t *testing.T) {
	// QuestionOut must NOT have Odpowiedz or Wskazowki fields
	q := QuestionOut{ID: "1.1", Tresc: "test"}
	data, _ := json.Marshal(q)
	var m map[string]any
	json.Unmarshal(data, &m)
	if _, ok := m["odpowiedz"]; ok {
		t.Error("QuestionOut should not have odpowiedz field")
	}
	if _, ok := m["wskazowki"]; ok {
		t.Error("QuestionOut should not have wskazowki field")
	}
	if _, ok := m["coaching"]; !ok {
		t.Error("QuestionOut must have coaching field")
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd analiza/cli && go test -run TestQuestionOutHasNoAnswer -v`
Expected: FAIL — `QuestionOut` not defined

**Step 3: Add types to types.go**

Add after the `ExerciseOut` block (line ~105), keeping `ExerciseOut` intact:

```go
// QuestionOut is what exercise question returns (no hints, no answer)
type QuestionOut struct {
	ID        string   `json:"id"`
	TypNazwa  string   `json:"typ_nazwa"`
	Kategoria string   `json:"kategoria"`
	Trudnosc  string   `json:"trudnosc"`
	Punkty    int      `json:"punkty"`
	Zrodlo    string   `json:"zrodlo"`
	Tagi      []string `json:"tagi"`
	Tresc     string   `json:"tresc"`
	Coaching  Coaching `json:"coaching"`
}

// Coaching provides student context computed from progress.db
type Coaching struct {
	StudentLevel   string   `json:"student_level"`
	HintDelay      int      `json:"hint_delay"`
	LeechTags      []string `json:"leech_tags"`
	PastMistakes   []string `json:"past_mistakes"`
	PreviousResult string   `json:"previous_result,omitempty"`
}

// HintsOut is what exercise hints returns
type HintsOut struct {
	ID        string   `json:"id"`
	Wskazowki []string `json:"wskazowki"`
	MaxHints  int      `json:"max_hints"`
}

// AnswerOut is what exercise answer returns
type AnswerOut struct {
	ID          string        `json:"id"`
	Odpowiedz  string        `json:"odpowiedz"`
	TypoweBledy []CommonError `json:"typowe_bledy"`
}
```

Update `ReviewOut` (line ~107-113) to use `QuestionOut`:

```go
type ReviewOut struct {
	Exercise       QuestionOut `json:"exercise"`
	Tag            string      `json:"tag"`
	DaysOverdue    int         `json:"days_overdue"`
	Retrievability float64     `json:"retrievability"`
}
```

Update `ExerciseNextOut` (line ~245-256) to use `QuestionOut`:

```go
type ExerciseNextOut struct {
	Mode           string      `json:"mode"`
	Exercise       QuestionOut `json:"exercise"`
	ReviewTag      *string     `json:"review_tag"`
	DaysOverdue    *int        `json:"days_overdue"`
	PoolWarning    *string     `json:"pool_warning"`
	SessionCount   int         `json:"session_count"`
	SessionWeight  int         `json:"session_weight"`
	ResetSuggested bool        `json:"reset_suggested"`
	ChosenTyp      string      `json:"chosen_typ,omitempty"`
}
```

**Step 4: Run test to verify it passes**

Run: `cd analiza/cli && go test -run TestQuestionOutHasNoAnswer -v`
Expected: PASS

**Step 5: Fix compilation errors**

`exercise review` and `exercise next` in commands.go assign `ExerciseOut` to `QuestionOut` fields — won't compile yet. Add a converter function to commands.go temporarily to unblock:

```go
func exerciseToQuestion(ex ExerciseOut) QuestionOut {
	return QuestionOut{
		ID: ex.ID, TypNazwa: ex.TypNazwa, Kategoria: ex.Kategoria,
		Trudnosc: ex.Trudnosc, Punkty: ex.Punkty, Zrodlo: ex.Zrodlo,
		Tagi: ex.Tagi, Tresc: ex.Tresc,
	}
}
```

Update all `out.Exercise = ex` assignments in `exerciseReviewCmd` and `exerciseNextCmd` to use `out.Exercise = exerciseToQuestion(ex)`. Remove `applyMaxHints` calls from review/next (hints are no longer in the output).

Run: `cd analiza/cli && go build -o matura .`
Expected: compiles successfully

**Step 6: Run full test suite**

Run: `cd analiza/cli && go test -v`
Expected: some tests fail (those referencing `ExerciseOut` fields from review/next output). Note which tests fail — they'll be fixed in Task 5.

**Step 7: Commit**

```
feat: add QuestionOut, HintsOut, AnswerOut types; update ReviewOut and ExerciseNextOut
```

---

## Task 2: Add buildCoaching() function

**Files:**
- Modify: `analiza/cli/commands.go`

**Step 1: Write failing test**

```go
func TestBuildCoachingNewStudent(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div"})
	if coaching.StudentLevel != "new" {
		t.Errorf("new student: got level %q, want new", coaching.StudentLevel)
	}
	if coaching.HintDelay != 1 {
		t.Errorf("new student: got hint_delay %d, want 1", coaching.HintDelay)
	}
	if len(coaching.LeechTags) != 0 {
		t.Errorf("new student: got %d leech_tags, want 0", len(coaching.LeechTags))
	}
	if len(coaching.PastMistakes) != 0 {
		t.Errorf("new student: got %d past_mistakes, want 0", len(coaching.PastMistakes))
	}
}

func TestBuildCoachingFamiliarStudent(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Simulate: 5 exercises done, streak 4
	db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'srednie', 4)")
	for i := 0; i < 5; i++ {
		db.Exec("INSERT OR IGNORE INTO progress_zrobione (id, typ, data, wynik) VALUES (?, 'cyfry_liczby', '2026-02-20', 'poprawne_bez_pomocy')",
			fmt.Sprintf("7.%d", i+1))
	}

	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div"})
	if coaching.StudentLevel != "familiar" {
		t.Errorf("familiar student: got level %q, want familiar", coaching.StudentLevel)
	}
	if coaching.HintDelay != 2 {
		t.Errorf("familiar student: got hint_delay %d, want 2", coaching.HintDelay)
	}
}

func TestBuildCoachingLeechTags(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Insert a leech tag: lapses=4, low stability, recent review
	db.Exec(`INSERT INTO progress_tagi (tag, poziom, stability, lapses, reps, state, last_review)
		VALUES ('cyfry-mod-div', 1, 1.0, 4, 6, 2, '2026-02-22')`)

	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div"})
	found := false
	for _, lt := range coaching.LeechTags {
		if lt == "cyfry-mod-div" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected cyfry-mod-div in leech_tags, got %v", coaching.LeechTags)
	}
}

func TestBuildCoachingLeechTagHealed(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Leech tag with high stability = healed (retrievability > 0.85)
	db.Exec(`INSERT INTO progress_tagi (tag, poziom, stability, lapses, reps, state, last_review)
		VALUES ('cyfry-mod-div', 3, 30.0, 4, 10, 2, '2026-02-22')`)

	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div"})
	if len(coaching.LeechTags) != 0 {
		t.Errorf("healed leech: expected 0 leech_tags, got %v", coaching.LeechTags)
	}
}

func TestBuildCoachingPastMistakes(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Record sessions + mistakes matching exercise tags
	db.Exec("INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.1', 'cyfry_liczby', '2026-02-20', 'poprawne_z_pomoca_1')")
	db.Exec("INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.2', 'cyfry_liczby', '2026-02-21', 'poprawne_bez_pomocy')")
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, data)
		VALUES ('7.1', 'cyfry_liczby', 'cyfry-mod-div', 'inicjalizacja iloczynu na 0', '2026-02-20')`)
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, data)
		VALUES ('7.1', 'cyfry_liczby', 'unrelated-tag', 'nieistotny blad', '2026-02-20')`)

	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div"})
	if len(coaching.PastMistakes) != 1 {
		t.Errorf("expected 1 past_mistake, got %d: %v", len(coaching.PastMistakes), coaching.PastMistakes)
	}
	if len(coaching.PastMistakes) > 0 && coaching.PastMistakes[0] != "inicjalizacja iloczynu na 0" {
		t.Errorf("expected 'inicjalizacja iloczynu na 0', got %q", coaching.PastMistakes[0])
	}
}

func TestBuildCoachingMastered(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'trudne', 5)")

	coaching := buildCoaching(db, "cyfry_liczby", []string{})
	if coaching.StudentLevel != "mastered" {
		t.Errorf("mastered student: got level %q, want mastered", coaching.StudentLevel)
	}
	if coaching.HintDelay != 3 {
		t.Errorf("mastered student: got hint_delay %d, want 3", coaching.HintDelay)
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `cd analiza/cli && go test -run TestBuildCoaching -v`
Expected: FAIL — `buildCoaching` not defined

**Step 3: Implement buildCoaching**

Add to `commands.go`:

```go
func buildCoaching(d *sql.DB, typ string, exerciseTags []string) Coaching {
	c := Coaching{
		StudentLevel: "new",
		HintDelay:    1,
		LeechTags:    []string{},
		PastMistakes: []string{},
	}

	// Student level from progress_typy + progress_zrobione
	var level sql.NullString
	var streak int
	d.QueryRow("SELECT poziom_trudnosci, streak FROM progress_typy WHERE typ = ?", typ).Scan(&level, &streak)

	if level.Valid && level.String == "trudne" {
		c.StudentLevel = "mastered"
		c.HintDelay = 3
	} else {
		var done int
		d.QueryRow("SELECT COUNT(*) FROM progress_zrobione WHERE typ = ?", typ).Scan(&done)

		switch {
		case done == 0:
			c.StudentLevel = "new"
			c.HintDelay = 1
		case done <= 3 && streak < 3:
			c.StudentLevel = "learning"
			c.HintDelay = 1
		default:
			c.StudentLevel = "familiar"
			c.HintDelay = 2
		}
	}

	// Leech tags: lapses >= 3 AND retrievability < 0.85
	today := time.Now().Format("2006-01-02")
	fsrsParams := DefaultFSRSParams()
	leechRows, err := d.Query(`SELECT tag, COALESCE(stability, 1.0), COALESCE(last_review, '')
		FROM progress_tagi WHERE lapses >= 3`)
	if err == nil {
		defer leechRows.Close()
		for leechRows.Next() {
			var tag, lastReview string
			var stability float64
			leechRows.Scan(&tag, &stability, &lastReview)
			elapsed := daysBetween(lastReview, today)
			r := fsrsParams.Retrievability(elapsed, stability)
			if r < 0.85 {
				c.LeechTags = append(c.LeechTags, tag)
			}
		}
	}

	// Past mistakes: from last 5 sessions, filtered by exercise tags
	if len(exerciseTags) > 0 {
		placeholders := make([]string, len(exerciseTags))
		params := []any{typ}
		for i, tag := range exerciseTags {
			placeholders[i] = "?"
			params = append(params, tag)
		}
		query := fmt.Sprintf(`SELECT DISTINCT blad_opis FROM progress_bledy
			WHERE typ = ? AND blad_kod IN (%s)
			AND data IN (SELECT DISTINCT data FROM progress_zrobione ORDER BY data DESC LIMIT 5)
			ORDER BY data DESC LIMIT 3`, strings.Join(placeholders, ","))
		rows, err := d.Query(query, params...)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var opis string
				rows.Scan(&opis)
				c.PastMistakes = append(c.PastMistakes, opis)
			}
		}
	}

	return c
}
```

**Step 4: Run tests to verify they pass**

Run: `cd analiza/cli && go test -run TestBuildCoaching -v`
Expected: all 6 tests PASS

**Step 5: Commit**

```
feat: add buildCoaching() — student level, hint delay, leech tags, past mistakes
```

---

## Task 3: Add exercise question command

**Files:**
- Modify: `analiza/cli/commands.go` (add `exerciseQuestionCmd`)
- Modify: `analiza/cli/main.go:63` (register new command)

**Step 1: Write failing test**

```go
func TestExerciseQuestionOutput(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	results, err := queryExercises(db, "cyfry_liczby", "latwe", "")
	if err != nil || len(results) == 0 {
		t.Fatal("no exercises found")
	}

	ex := results[0]
	q := exerciseToQuestion(ex)
	q.Coaching = buildCoaching(db, ex.TypNazwa, ex.Tagi)

	// Must have tresc, no odpowiedz
	if q.Tresc == "" {
		t.Error("QuestionOut.Tresc is empty")
	}
	if q.Coaching.StudentLevel != "new" {
		t.Errorf("fresh DB: got level %q, want new", q.Coaching.StudentLevel)
	}

	data, _ := json.Marshal(q)
	var m map[string]any
	json.Unmarshal(data, &m)
	if _, ok := m["odpowiedz"]; ok {
		t.Error("QuestionOut JSON must not contain odpowiedz")
	}
	if _, ok := m["wskazowki"]; ok {
		t.Error("QuestionOut JSON must not contain wskazowki")
	}
}
```

**Step 2: Run test to verify it passes** (uses existing functions)

Run: `cd analiza/cli && go test -run TestExerciseQuestionOutput -v`
Expected: PASS (this test validates the conversion, not the command itself)

**Step 3: Implement exerciseQuestionCmd**

Add to `commands.go`:

```go
func exerciseQuestionCmd() *cobra.Command {
	var typ, trudnosc, exclude string

	cmd := &cobra.Command{
		Use:   "question",
		Short: "Get exercise question with coaching (no hints/answer)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if typ == "" {
				return fatal("--typ is required")
			}

			d := db(cmd)
			level := getLevel(d, typ)
			if !cmd.Flags().Changed("trudnosc") {
				trudnosc = level
			}

			results, err := queryExercises(d, typ, trudnosc, exclude)
			if err != nil {
				return err
			}
			if len(results) == 0 && !cmd.Flags().Changed("trudnosc") {
				results, err = queryExercises(d, typ, "", exclude)
				if err != nil {
					return err
				}
			}
			if len(results) == 0 {
				return notFound(fmt.Sprintf("no exercises found for typ=%s trudnosc=%s", typ, trudnosc))
			}

			chosen := results[rand.Intn(len(results))]
			q := exerciseToQuestion(chosen)
			q.Coaching = buildCoaching(d, typ, chosen.Tagi)
			addWeight(d, 4)
			jsonOut(q)
			return nil
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Exercise type (e.g. sql_group_by)")
	cmd.Flags().StringVar(&trudnosc, "trudnosc", "", "Difficulty: latwe, srednie, srednie-trudne, trudne")
	cmd.Flags().StringVar(&exclude, "exclude", "", "Comma-separated IDs to exclude")
	return cmd
}
```

**Step 4: Register in main.go**

Change line 63 from:
```go
exerciseCmd.AddCommand(exerciseGetCmd(), exerciseReviewCmd(), exerciseNextCmd())
```
to:
```go
exerciseCmd.AddCommand(exerciseQuestionCmd(), exerciseHintsCmd(), exerciseAnswerCmd(), exerciseReviewCmd(), exerciseNextCmd())
```

(hints and answer commands will be added in Task 4 — for now this won't compile until Task 4 is done. Alternatively, add only `exerciseQuestionCmd()` now and expand in Task 4.)

Pragmatic approach: add only `exerciseQuestionCmd()` now:
```go
exerciseCmd.AddCommand(exerciseGetCmd(), exerciseQuestionCmd(), exerciseReviewCmd(), exerciseNextCmd())
```

**Step 5: Build and test**

Run: `cd analiza/cli && go build -o matura . && ./matura exercise question --typ cyfry_liczby 2>/dev/null | python3 -c "import sys,json; d=json.load(sys.stdin); print('id:', d['id']); print('coaching:', d['coaching']); print('has odpowiedz:', 'odpowiedz' in d)"`
Expected: `has odpowiedz: False`, coaching fields present

**Step 6: Commit**

```
feat: add exercise question command with coaching
```

---

## Task 4: Add exercise hints and exercise answer commands

**Files:**
- Modify: `analiza/cli/commands.go`
- Modify: `analiza/cli/main.go:63`

**Step 1: Write failing tests**

```go
func TestExerciseHintsById(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Get a known exercise ID
	results, _ := queryExercises(db, "cyfry_liczby", "latwe", "")
	if len(results) == 0 {
		t.Fatal("no exercises")
	}
	id := results[0].ID

	hints := getExerciseHints(db, id, "latwe")
	if hints.ID != id {
		t.Errorf("got id %q, want %q", hints.ID, id)
	}
	if len(hints.Wskazowki) == 0 {
		t.Error("expected at least 1 hint")
	}
}

func TestExerciseAnswerById(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	results, _ := queryExercises(db, "cyfry_liczby", "latwe", "")
	if len(results) == 0 {
		t.Fatal("no exercises")
	}
	id := results[0].ID

	answer := getExerciseAnswer(db, id)
	if answer.ID != id {
		t.Errorf("got id %q, want %q", answer.ID, id)
	}
	if answer.Odpowiedz == "" {
		t.Error("expected non-empty odpowiedz")
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `cd analiza/cli && go test -run "TestExerciseHintsById|TestExerciseAnswerById" -v`
Expected: FAIL — functions not defined

**Step 3: Implement helper functions and commands**

Add to `commands.go`:

```go
func getExerciseHints(d *sql.DB, id, level string) HintsOut {
	var wskazowkiJSON string
	d.QueryRow("SELECT wskazowki FROM data.cwiczenia WHERE id = ?", id).Scan(&wskazowkiJSON)
	var wskazowki []string
	json.Unmarshal([]byte(wskazowkiJSON), &wskazowki)
	if wskazowki == nil {
		wskazowki = []string{}
	}
	maxHints := calculateMaxHints(level)
	if maxHints >= 0 && len(wskazowki) > maxHints {
		wskazowki = wskazowki[:maxHints]
	}
	if maxHints < 0 {
		maxHints = len(wskazowki)
	}
	return HintsOut{ID: id, Wskazowki: wskazowki, MaxHints: maxHints}
}

func getExerciseAnswer(d *sql.DB, id string) AnswerOut {
	var odpowiedz, typoweBledyJSON string
	d.QueryRow("SELECT odpowiedz, typowe_bledy FROM data.cwiczenia WHERE id = ?", id).Scan(&odpowiedz, &typoweBledyJSON)
	var bledy []CommonError
	json.Unmarshal([]byte(typoweBledyJSON), &bledy)
	if bledy == nil {
		bledy = []CommonError{}
	}
	return AnswerOut{ID: id, Odpowiedz: odpowiedz, TypoweBledy: bledy}
}

func exerciseHintsCmd() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "hints",
		Short: "Get hints for an exercise by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			if id == "" {
				return fatal("--id is required")
			}
			d := db(cmd)
			// Verify exercise exists
			var exists int
			if err := d.QueryRow("SELECT COUNT(*) FROM data.cwiczenia WHERE id = ?", id).Scan(&exists); err != nil || exists == 0 {
				return notFound(fmt.Sprintf("exercise %s not found", id))
			}
			// Get typ to determine level for max_hints
			var typ string
			d.QueryRow("SELECT typ_nazwa FROM data.cwiczenia WHERE id = ?", id).Scan(&typ)
			level := getLevel(d, typ)
			hints := getExerciseHints(d, id, level)
			addWeight(d, 1)
			jsonOut(hints)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Exercise ID (e.g. 7.1)")
	return cmd
}

func exerciseAnswerCmd() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "answer",
		Short: "Get answer for an exercise by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			if id == "" {
				return fatal("--id is required")
			}
			d := db(cmd)
			var exists int
			if err := d.QueryRow("SELECT COUNT(*) FROM data.cwiczenia WHERE id = ?", id).Scan(&exists); err != nil || exists == 0 {
				return notFound(fmt.Sprintf("exercise %s not found", id))
			}
			answer := getExerciseAnswer(d, id)
			addWeight(d, 2)
			jsonOut(answer)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Exercise ID (e.g. 7.1)")
	return cmd
}
```

**Step 4: Register in main.go**

Change line 63 to:
```go
exerciseCmd.AddCommand(exerciseQuestionCmd(), exerciseHintsCmd(), exerciseAnswerCmd(), exerciseReviewCmd(), exerciseNextCmd())
```

Remove `exerciseGetCmd()` from registration.

**Step 5: Run tests**

Run: `cd analiza/cli && go test -run "TestExerciseHintsById|TestExerciseAnswerById" -v`
Expected: PASS

**Step 6: Manual smoke test**

Run:
```bash
cd analiza/cli && go build -o matura .
./matura exercise question --typ cyfry_liczby 2>/dev/null | python3 -c "import sys,json; d=json.load(sys.stdin); print('ID:', d['id'])"
# capture the ID, e.g. 7.11
./matura exercise hints --id 7.11 2>/dev/null
./matura exercise answer --id 7.11 2>/dev/null
```

**Step 7: Commit**

```
feat: add exercise hints and exercise answer commands
```

---

## Task 5: Update exercise review and exercise next

**Files:**
- Modify: `analiza/cli/commands.go` (lines ~366-445 and ~1642-1767)

**Step 1: Write failing tests**

```go
func TestReviewOutHasCoaching(t *testing.T) {
	// ReviewOut.Exercise should be QuestionOut with coaching
	q := QuestionOut{
		ID: "7.1", Tresc: "test",
		Coaching: Coaching{StudentLevel: "familiar", HintDelay: 2},
	}
	r := ReviewOut{Exercise: q, Tag: "test-tag", DaysOverdue: 3}
	data, _ := json.Marshal(r)
	var m map[string]any
	json.Unmarshal(data, &m)
	ex := m["exercise"].(map[string]any)
	if _, ok := ex["coaching"]; !ok {
		t.Error("ReviewOut.Exercise must have coaching")
	}
	if _, ok := ex["odpowiedz"]; ok {
		t.Error("ReviewOut.Exercise must not have odpowiedz")
	}
}

func TestExerciseNextOutputNoAnswer(t *testing.T) {
	// ExerciseNextOut.Exercise should be QuestionOut
	q := QuestionOut{ID: "7.1", Tresc: "test"}
	out := ExerciseNextOut{Mode: "new", Exercise: q}
	data, _ := json.Marshal(out)
	var m map[string]any
	json.Unmarshal(data, &m)
	ex := m["exercise"].(map[string]any)
	if _, ok := ex["odpowiedz"]; ok {
		t.Error("ExerciseNextOut.Exercise must not have odpowiedz")
	}
}
```

**Step 2: Run tests — should already pass** (types were changed in Task 1)

Run: `cd analiza/cli && go test -run "TestReviewOutHasCoaching|TestExerciseNextOutputNoAnswer" -v`
Expected: PASS

**Step 3: Update exerciseReviewCmd to populate coaching + previous_result**

In `exerciseReviewCmd` (commands.go ~420-437), change the loop body:

Where it currently does:
```go
results = append(results, ReviewOut{
    Exercise:       ex,
    ...
})
```

Change to:
```go
q := exerciseToQuestion(ex)
q.Coaching = buildCoaching(d, ex.TypNazwa, ex.Tagi)
// Add previous_result for reviews
var prevResult sql.NullString
d.QueryRow("SELECT wynik FROM progress_zrobione WHERE id = ?", ex.ID).Scan(&prevResult)
if prevResult.Valid {
    q.Coaching.PreviousResult = prevResult.String
}

results = append(results, ReviewOut{
    Exercise:       q,
    Tag:            rc.tag,
    DaysOverdue:    daysOverdue,
    Retrievability: rc.retrievability,
})
```

Remove the `applyMaxHints` call from this function.

**Step 4: Update exerciseNextCmd similarly**

In `exerciseNextCmd`, every path that sets `out.Exercise = ex` must become:
```go
q := exerciseToQuestion(ex)
q.Coaching = buildCoaching(d, ex.TypNazwa, ex.Tagi)
out.Exercise = q
```

For the review path (Priority 1), also add `previous_result`:
```go
var prevResult sql.NullString
d.QueryRow("SELECT wynik FROM progress_zrobione WHERE id = ?", ex.ID).Scan(&prevResult)
if prevResult.Valid {
    q.Coaching.PreviousResult = prevResult.String
}
```

Remove all `applyHints` calls from `exerciseNextCmd`. Remove the `applyHints` closure entirely. The `maxHints` flag can be removed from this command.

**Step 5: Build and run full tests**

Run: `cd analiza/cli && go build -o matura . && go test -v`
Expected: some existing tests may fail if they check for `Wskazowki` or `Odpowiedz` in review/next output. Fix those in Task 6.

**Step 6: Commit**

```
feat: exercise review and next return QuestionOut with coaching
```

---

## Task 6: Remove exercise get and fix existing tests

**Files:**
- Modify: `analiza/cli/commands.go` (remove `exerciseGetCmd`)
- Modify: `analiza/cli/main_test.go` (update broken tests)

**Step 1: Remove exerciseGetCmd function**

Delete `exerciseGetCmd()` from `commands.go` (lines ~312-362). It's already removed from main.go registration in Task 4.

**Step 2: Fix compilation**

Run: `cd analiza/cli && go build -o matura .`
If `exerciseGetCmd` is still referenced anywhere, remove references.

**Step 3: Update failing tests**

Run `go test -v` and fix each failing test:

- `TestExerciseGetMaxHintsIntegration` (line ~1949): remove or rewrite to test `exercise hints --id X` with max_hints behavior
- `TestExerciseExclude` (line ~347): update to use `exerciseQuestionCmd` pattern
- Any test that accesses `.Exercise.Wskazowki` or `.Exercise.Odpowiedz` from review/next output: remove those assertions
- Any test that references `exerciseGetCmd`: replace with `exerciseQuestionCmd`

**Step 4: Run full test suite**

Run: `cd analiza/cli && go test -v`
Expected: ALL PASS

**Step 5: Commit**

```
refactor: remove exercise get, fix all tests for lazy loading
```

---

## Task 7: Update test_qa.sh

**Files:**
- Modify: `analiza/test_qa.sh`

**Step 1: Update Layer 1 smoke tests**

Replace all `exercise get` references with `exercise question`:

- Line ~163: `test_json_cmd "exercise question --typ cyfry_liczby" "$MATURA" exercise question --typ cyfry_liczby`
- Line ~192: `test_json_cmd "exercise question --typ $typ" "$MATURA" exercise question --typ "$typ"`
- Lines ~288-312: max-hints tests → move to `exercise hints` with `--id`
- Line ~335-336: error test → `exercise question --typ NIEISTNIEJACY`

Add new smoke tests:
```bash
# exercise hints
ex_id=$("$MATURA" exercise question --typ cyfry_liczby 2>/dev/null | jq -r '.id')
test_json_cmd "exercise hints --id $ex_id" "$MATURA" exercise hints --id "$ex_id"

# exercise answer
test_json_cmd "exercise answer --id $ex_id" "$MATURA" exercise answer --id "$ex_id"

# exercise question must NOT have odpowiedz
q_out=$("$MATURA" exercise question --typ cyfry_liczby 2>/dev/null)
if echo "$q_out" | jq -e '.odpowiedz' >/dev/null 2>&1; then
  fail "exercise question should not have odpowiedz"
else
  pass "exercise question has no odpowiedz"
fi

# exercise question must have coaching
if echo "$q_out" | jq -e '.coaching.student_level' >/dev/null 2>&1; then
  pass "exercise question has coaching.student_level"
else
  fail "exercise question missing coaching"
fi
```

**Step 2: Update Layer 6 journey tests**

Lines ~654, ~680, ~782, ~806, ~826: replace `exercise get` with `exercise question`.

**Step 3: Run test_qa.sh**

Run: `cd analiza && ./test_qa.sh`
Expected: all layers pass

**Step 4: Commit**

```
test: update test_qa.sh for exercise question/hints/answer commands
```

---

## Task 8: Update SKILL.md

**Files:**
- Modify: `.claude/skills/matura/SKILL.md`

**Step 1: Read current SKILL.md**

Read the file and identify all references to `exercise get`.

**Step 2: Replace exercise flow**

Replace the exercise retrieval section with the new lazy flow:

```markdown
### Exercise Flow

1. **Fetch question**: `exercise question --typ {typ}` → returns treść + coaching (no hints, no answer)
2. **Teach sokratejsko**: Use own knowledge. `coaching.hint_delay` suggests attempts before hints.
3. **If student stuck** (after hint_delay attempts OR student asks): `exercise hints --id {id}`
4. **When student submits answer**: `exercise answer --id {id}` → compare and grade

Rules:
- `coaching.hint_delay` is a suggestion, not a blocker. If student asks for help → fetch hints immediately.
- `coaching.leech_tags` → these tags are problem areas; watch for related mistakes.
- `coaching.past_mistakes` → student made these errors before in this area.
- `coaching.previous_result` (reviews only) → how student did last time.
- For IMPLEMENTACJA: `tresc` contains `**Oczekiwany wynik**` — compare student output directly.
```

**Step 3: Remove old hint logic**

Remove any SKILL.md rules about interpreting raw streak/zrobione data, max_hints logic, or "don't show answer" rules. CLI now controls this via coaching.

**Step 4: Verify SKILL.md lint passes**

Run: `cd analiza && ./test_qa.sh --layer 3`
Expected: PASS

**Step 5: Commit**

```
docs: update SKILL.md for lazy loading exercise flow
```

---

## Task 9: Final build + full QA

**Files:**
- Run: `analiza/cli/build.sh` (rebuild binaries + reimport matura.db)

**Step 1: Full rebuild**

Run: `cd analiza/cli && bash build.sh`
Expected: builds macOS + Windows binaries, reimports matura.db

**Step 2: Run Go unit tests**

Run: `cd analiza/cli && go test -v`
Expected: ALL PASS

**Step 3: Run full QA suite**

Run: `cd analiza && ./test_qa.sh`
Expected: all 6 layers pass

**Step 4: Manual integration test**

```bash
cd analiza/cli
# Fresh progress DB
rm -f test_progress.db
./matura --db-dir . exercise question --typ cyfry_liczby
# → should show coaching.student_level: "new"

# Get ID from output, then:
./matura --db-dir . exercise hints --id <ID>
./matura --db-dir . exercise answer --id <ID>

# Record a result
./matura --db-dir . progress update --id <ID> --wynik poprawne_bez_pomocy

# Next question should show coaching.student_level: "learning"
./matura --db-dir . exercise question --typ cyfry_liczby
rm -f test_progress.db
```

**Step 5: Commit**

```
chore: rebuild binaries and verify full QA passes
```
