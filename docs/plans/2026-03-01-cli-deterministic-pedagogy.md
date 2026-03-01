# CLI Deterministic Pedagogy Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Move pedagogical decisions from AI to deterministic CLI where possible — enum scoring, auto error detection, answer validation, hint guardrails.

**Architecture:** 6 tasks in TDD order (simplest first). All changes in `matura_informatyka_rozszerzona/analiza/cli/`. Each task: write failing test → implement → green → commit.

**Tech Stack:** Go 1.21+, SQLite (modernc.org/sqlite), cobra CLI framework.

**Base path:** `matura_informatyka_rozszerzona/analiza/cli/`

---

### Task 1: `progress update --wynik` enum guard (5 levels)

**Files:**
- Modify: `commands.go:719-726` (validWynik map)
- Modify: `commands.go:900` (flag help text)
- Test: `main_test.go` (new test + update existing)

**Context:** Currently `--wynik` accepts 4 values: `poprawne_bez_pomocy`, `poprawne_z_pomoca_1`, `poprawne_z_pomoca_2`, `walk_through`. These map to FSRS grades. We need to keep FSRS working while adding a 5-level semantic enum.

**Design decision:** The 4 current values ARE the FSRS grades and drive spaced repetition (e.g., `walk_through` resets streak, `poprawne_z_pomoca_2` = Again). Changing them to `pelne/czesciowe/zero` would break FSRS semantics. Instead: **keep current `--wynik` values AND add `--punktacja` as the 5-level scoring enum**.

The 5-level `--punktacja` is for human-readable scoring. The existing `--wynik` stays for FSRS. Both are required on `progress update`.

**Step 1: Write failing tests**

In `main_test.go`, add:

```go
func TestProgressUpdatePunktacja(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Register exercise first (guardrail requirement)
	registerFetch(db, "1.1", "sledzenie_algorytmu", 1)
	incrementAttempt(db, "1.1")

	// Valid punktacja values
	validValues := []string{"pelne", "prawie_pelne", "czesciowe", "minimalne", "zero"}
	for _, v := range validValues {
		if !isValidPunktacja(v) {
			t.Errorf("expected %q to be valid punktacja", v)
		}
	}

	// Invalid values
	invalidValues := []string{"3", "75%", "4/5", "full", ""}
	for _, v := range invalidValues {
		if isValidPunktacja(v) {
			t.Errorf("expected %q to be invalid punktacja", v)
		}
	}
}

func TestPunktacjaToPercent(t *testing.T) {
	cases := []struct {
		input   string
		percent int
	}{
		{"pelne", 100},
		{"prawie_pelne", 75},
		{"czesciowe", 50},
		{"minimalne", 25},
		{"zero", 0},
	}
	for _, c := range cases {
		got := punktacjaToPercent(c.input)
		if got != c.percent {
			t.Errorf("punktacjaToPercent(%q) = %d, want %d", c.input, got, c.percent)
		}
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestProgressUpdatePunktacja -v`
Expected: FAIL — `isValidPunktacja` and `punktacjaToPercent` undefined

**Step 3: Implement minimal code**

In `commands.go`, add helper functions (near `calculateLevel`, around line 905):

```go
var validPunktacja = map[string]int{
	"pelne":         100,
	"prawie_pelne":  75,
	"czesciowe":     50,
	"minimalne":     25,
	"zero":          0,
}

func isValidPunktacja(p string) bool {
	_, ok := validPunktacja[p]
	return ok
}

func punktacjaToPercent(p string) int {
	return validPunktacja[p]
}
```

Add `--punktacja` flag to `progressUpdateCmd` (around line 707):

```go
func progressUpdateCmd() *cobra.Command {
	var id, wynik, punktacja string
	var czas int
	// ... in RunE, after wynik validation:
	if punktacja != "" && !isValidPunktacja(punktacja) {
		return fatal("--punktacja must be one of: pelne, prawie_pelne, czesciowe, minimalne, zero")
	}
```

Add `Punktacja` field to `ProgressUpdateOut` in `types.go`:

```go
type ProgressUpdateOut struct {
	// ... existing fields ...
	Punktacja       *string  `json:"punktacja,omitempty"`
	PunktacjaPct    *int     `json:"punktacja_pct,omitempty"`
}
```

Populate in output (around line 820):

```go
if punktacja != "" {
	pct := punktacjaToPercent(punktacja)
	out.Punktacja = &punktacja
	out.PunktacjaPct = &pct
}
```

Add flag (around line 900):

```go
cmd.Flags().StringVar(&punktacja, "punktacja", "", "Scoring: pelne, prawie_pelne, czesciowe, minimalne, zero")
```

**Step 4: Run tests to verify they pass**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run "TestProgressUpdatePunktacja|TestPunktacjaToPercent" -v`
Expected: PASS

**Step 5: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test ./... -v -count=1`
Expected: All existing tests PASS (backwards compatible — `--punktacja` is optional)

**Step 6: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/commands.go matura_informatyka_rozszerzona/analiza/cli/types.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "feat(cli): add --punktacja enum (5 levels) to progress update"
```

---

### Task 2: `exercise hints` guard — no attempt since last hint

**Files:**
- Modify: `guardrails.go` (new column tracking + check function)
- Modify: `database.go` (migration v8 for new column)
- Modify: `commands.go:514-561` (exerciseHintsCmd — add guard)
- Test: `main_test.go` (new test)

**Step 1: Write failing test**

```go
func TestHintBlockedNoAttemptSinceLastHint(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Register exercise
	registerFetch(db, "1.1", "sledzenie_algorytmu", 1)

	// First attempt (to pass hint_delay)
	incrementAttempt(db, "1.1")

	// First hints fetch should succeed
	canH, _, err := checkCanFetchHintsSinceAttempt(db, "1.1")
	if err != nil {
		t.Fatalf("checkCanFetchHintsSinceAttempt: %v", err)
	}
	if !canH {
		t.Error("first hint fetch should be allowed")
	}

	// Mark hints fetched
	markHintsFetched(db, "1.1")

	// Second hints fetch without new attempt should FAIL
	canH2, _, err := checkCanFetchHintsSinceAttempt(db, "1.1")
	if err != nil {
		t.Fatalf("checkCanFetchHintsSinceAttempt: %v", err)
	}
	if canH2 {
		t.Error("second hint fetch without attempt should be blocked")
	}

	// New attempt should unblock
	incrementAttempt(db, "1.1")
	canH3, _, err := checkCanFetchHintsSinceAttempt(db, "1.1")
	if err != nil {
		t.Fatalf("checkCanFetchHintsSinceAttempt: %v", err)
	}
	if !canH3 {
		t.Error("hint fetch after new attempt should be allowed")
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestHintBlockedNoAttemptSinceLastHint -v`
Expected: FAIL — `checkCanFetchHintsSinceAttempt` and `markHintsFetched` undefined

**Step 3: Implement**

In `database.go`, add migration v8 (add `hints_given` column to `active_exercises`):

```go
{Version: 8, Apply: func(tx *sql.Tx) error {
    _, err := tx.Exec(`ALTER TABLE active_exercises ADD COLUMN hints_given INTEGER NOT NULL DEFAULT 0`)
    return err
}},
```

Update `currentSchemaVersion = 8`.

Also update the CREATE TABLE in `initProgressSchema` to include `hints_given INTEGER NOT NULL DEFAULT 0`.

In `guardrails.go`, add:

```go
// checkCanFetchHintsSinceAttempt checks if student attempted since last hint.
// Returns (allowed, attempts_since_last_hint, error).
func checkCanFetchHintsSinceAttempt(d *sql.DB, exerciseID string) (bool, int, error) {
	var attemptCount, hintsGiven int
	err := d.QueryRow(`
		SELECT attempt_count, hints_given FROM active_exercises WHERE exercise_id = ?
	`, exerciseID).Scan(&attemptCount, &hintsGiven)
	if err == sql.ErrNoRows {
		return true, 0, nil // Not tracked = allow (backwards compat)
	}
	if err != nil {
		return false, 0, err
	}
	// First hint is always allowed (hintsGiven == 0)
	if hintsGiven == 0 {
		return true, attemptCount, nil
	}
	// After first hint, require at least one new attempt
	return attemptCount > hintsGiven, attemptCount - hintsGiven, nil
}

// markHintsFetched records that hints were given at current attempt count.
func markHintsFetched(d *sql.DB, exerciseID string) error {
	_, err := d.Exec(`UPDATE active_exercises SET hints_given = attempt_count WHERE exercise_id = ?`, exerciseID)
	return err
}
```

In `commands.go:exerciseHintsCmd` (around line 543), after the existing `canH` check, add:

```go
// Additional guard: block hints if no attempt since last hint
canSince, attSince, err := checkCanFetchHintsSinceAttempt(d, id)
if err != nil {
    return fmt.Errorf("checkCanFetchHintsSinceAttempt: %w", err)
}
if !canSince {
    out := map[string]interface{}{
        "status":                     "HINT_BLOCKED_NO_ATTEMPT",
        "exercise_id":                id,
        "attempts_since_last_hint":   attSince,
        "action":                     "Uczeń nie podjął próby od ostatniego hinta. Zarejestruj próbę (progress blad) przed kolejnym hintem.",
    }
    jsonOut(out)
    return nil
}
```

After the hints are returned (around line 545), call `markHintsFetched`:

```go
markHintsFetched(d, id)
```

**Step 4: Run test to verify it passes**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestHintBlockedNoAttemptSinceLastHint -v`
Expected: PASS

**Step 5: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test ./... -v -count=1`
Expected: All tests PASS

**Step 6: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/guardrails.go matura_informatyka_rozszerzona/analiza/cli/database.go matura_informatyka_rozszerzona/analiza/cli/commands.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "feat(cli): block hints until student attempts since last hint"
```

---

### Task 3: Coaching actions with pre-formatted text

**Files:**
- Modify: `types.go:122-130` (change `Actions []string` to structured type)
- Modify: `commands.go:230-244` (buildCoaching — generate structured actions)
- Test: `main_test.go` (new test)

**Step 1: Write failing test**

```go
func TestCoachingActionsStructured(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Insert a leech tag
	db.Exec(`INSERT INTO progress_tagi (tag, lapses, stability, last_review, nastepna_powtorka)
		VALUES ('off_by_one', 4, 1.0, '2026-02-28', '2026-03-01')`)

	// Insert a past mistake
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, hint_level, data)
		VALUES ('7.1', 'cyfry_liczby', 'off_by_one', 'Pomylka o 1', 1, '2026-02-28')`)

	coaching := buildCoaching(db, "cyfry_liczby", []string{"off_by_one"})

	if len(coaching.StructuredActions) == 0 {
		t.Fatal("expected structured coaching actions")
	}

	// Check WARN_LEECH action exists with required fields
	found := false
	for _, a := range coaching.StructuredActions {
		if a.Typ == "WARN_LEECH" {
			found = true
			if a.Tekst == "" {
				t.Error("WARN_LEECH action should have tekst")
			}
			if a.Priorytet == "" {
				t.Error("WARN_LEECH action should have priorytet")
			}
			if a.Tag == "" {
				t.Error("WARN_LEECH action should have tag")
			}
		}
	}
	if !found {
		t.Error("expected WARN_LEECH action for leech tag")
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestCoachingActionsStructured -v`
Expected: FAIL — `StructuredActions` field doesn't exist

**Step 3: Implement**

In `types.go`, add new type and field:

```go
// CoachingAction is a structured coaching action with pre-formatted text
type CoachingAction struct {
	Typ       string `json:"typ"`                 // WARN_LEECH, MENTION_PAST, HINT_DELAY, SUGGEST_SLOWER
	Tag       string `json:"tag,omitempty"`        // relevant tag (for WARN_LEECH)
	Tekst     string `json:"tekst"`                // pre-formatted Polish sentence
	Priorytet string `json:"priorytet"`            // wysoki, niski
}
```

Add field to `Coaching` struct:

```go
type Coaching struct {
	StudentLevel      string           `json:"student_level"`
	HintDelay         int              `json:"hint_delay"`
	LeechTags         []string         `json:"leech_tags"`
	PastMistakes      []string         `json:"past_mistakes"`
	PreviousResult    string           `json:"previous_result,omitempty"`
	Actions           []string         `json:"coaching_actions"`
	StructuredActions []CoachingAction `json:"coaching_actions_v2"`
}
```

In `commands.go`, update `buildCoaching` (around line 230-244):

```go
// Generate coaching_actions (legacy string format + new structured)
var actions []string
var structured []CoachingAction

for _, tag := range c.LeechTags {
    actions = append(actions, fmt.Sprintf("WARN_LEECH: Tag '%s' sprawia Ci trudnosc — zwroc uwage", tag))
    structured = append(structured, CoachingAction{
        Typ:       "WARN_LEECH",
        Tag:       tag,
        Tekst:     fmt.Sprintf("Błędy z tagiem '%s' powtarzają się — to Twój leech tag. Zwróć szczególną uwagę na ten wzorzec.", tag),
        Priorytet: "wysoki",
    })
}
for _, m := range c.PastMistakes {
    actions = append(actions, fmt.Sprintf("MENTION_PAST: Ostatnio mialeS problem z '%s'", m))
    structured = append(structured, CoachingAction{
        Typ:       "MENTION_PAST",
        Tekst:     fmt.Sprintf("Ostatnio miałeś problem z: %s. Sprawdź czy tym razem jest lepiej.", m),
        Priorytet: "niski",
    })
}
if c.HintDelay >= 2 {
    actions = append(actions, fmt.Sprintf("HINT_DELAY: %d (Od teraz mniej podpowiedzi — rozwijasz samodzielnosc)", c.HintDelay))
    structured = append(structured, CoachingAction{
        Typ:       "HINT_DELAY",
        Tekst:     fmt.Sprintf("Od teraz mniej podpowiedzi (opóźnienie: %d próby) — rozwijasz samodzielność.", c.HintDelay),
        Priorytet: "niski",
    })
}

if actions == nil {
    actions = []string{}
}
if structured == nil {
    structured = []CoachingAction{}
}
c.Actions = actions
c.StructuredActions = structured
```

**Step 4: Run test to verify it passes**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestCoachingActionsStructured -v`
Expected: PASS

**Step 5: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test ./... -v -count=1`
Expected: All tests PASS (legacy `coaching_actions` still works)

**Step 6: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/types.go matura_informatyka_rozszerzona/analiza/cli/commands.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "feat(cli): add structured coaching_actions_v2 with pre-formatted text"
```

---

### Task 4: `exercise suggest-error` command

**Files:**
- Modify: `commands.go` (new command `exerciseSuggestErrorCmd`)
- Modify: `types.go` (new `SuggestErrorOut` type)
- Modify: `error_codes.go` (add keyword lists per error code)
- Modify: `main.go` (register command under `exercise`)
- Test: `main_test.go` (new tests)

**Context:** Two paths: (A) auto-detect from `--student-answer` comparison for exercises with clear answers, (B) return typed code list when no student answer given.

**Step 1: Write failing tests**

```go
func TestSuggestErrorAutoDetectOffByOne(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Exercise 1.1 is sledzenie_algorytmu with a known answer
	var odpowiedz string
	err := db.QueryRow("SELECT odpowiedz FROM data.cwiczenia WHERE id = '1.1'").Scan(&odpowiedz)
	if err != nil {
		t.Fatalf("query answer: %v", err)
	}

	// Test the detection logic directly
	result := detectErrorPattern(odpowiedz, "wrong_answer_placeholder")
	// We just need the function to exist and return a struct
	if result == nil {
		// nil is acceptable when no pattern detected
	}
	_ = db // used for context
}

func TestSuggestErrorCodeList(t *testing.T) {
	// When no --student-answer, should return full code list for type
	codes := getCodesForType("cyfry_liczby")
	if len(codes) == 0 {
		t.Fatal("expected codes for cyfry_liczby")
	}

	// Each code should have kod + opis
	for _, c := range codes {
		if c.Kod == "" {
			t.Error("code should have kod")
		}
		if c.Opis == "" {
			t.Errorf("code %q should have opis", c.Kod)
		}
	}
}

func TestDetectErrorPatternOffByOne(t *testing.T) {
	cases := []struct {
		correct, student string
		expectedCode     string
	}{
		{"6", "5", "off_by_one"},
		{"6", "7", "off_by_one"},
		{"100", "99", "off_by_one"},
		{"PRAWDA", "FALSZ", "odwrocona_logika"},
		{"prawda", "fałsz", "odwrocona_logika"},
		{"42", "0", "brak_algorytmu"},
		{"42", "", "brak_algorytmu"},
		{"42", "24", ""}, // no pattern detected
	}
	for _, c := range cases {
		result := detectErrorPattern(c.correct, c.student)
		got := ""
		if result != nil {
			got = result.Kod
		}
		if got != c.expectedCode {
			t.Errorf("detectErrorPattern(%q, %q) = %q, want %q", c.correct, c.student, got, c.expectedCode)
		}
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run "TestSuggestError|TestDetectErrorPattern" -v`
Expected: FAIL — functions undefined

**Step 3: Implement**

In `types.go`, add:

```go
// SuggestErrorOut is what exercise suggest-error returns
type SuggestErrorOut struct {
	AutoDetected bool             `json:"auto_detected"`
	Rekomendowany *CodeSuggestion `json:"rekomendowany,omitempty"`
	Powod         string          `json:"powod,omitempty"`
	KodyDlaTypu   []CodeSuggestion `json:"kody_dla_typu"`
}
```

In `error_codes.go`, add:

```go
// getCodesForType returns all error codes with descriptions for a given type.
func getCodesForType(typ string) []CodeSuggestion {
	codes, ok := errorCodeWhitelist[typ]
	if !ok {
		return nil
	}
	result := make([]CodeSuggestion, len(codes))
	for i, c := range codes {
		result[i] = CodeSuggestion{Kod: c, Opis: errorCodeDescriptions[c]}
	}
	return result
}

// detectErrorPattern compares correct and student answers, returns detected error code or nil.
func detectErrorPattern(correct, student string) *CodeSuggestion {
	correct = strings.TrimSpace(correct)
	student = strings.TrimSpace(student)

	if student == "" || student == "0" {
		if correct != "0" && correct != "" {
			return &CodeSuggestion{Kod: "brak_algorytmu", Opis: "Brak odpowiedzi lub wynik zerowy"}
		}
	}

	// Boolean inversion: PRAWDA <-> FALSZ
	cLow := strings.ToLower(correct)
	sLow := strings.ToLower(student)
	boolTrue := map[string]bool{"prawda": true, "p": true, "true": true, "tak": true}
	boolFalse := map[string]bool{"falsz": true, "fałsz": true, "f": true, "false": true, "nie": true}
	if (boolTrue[cLow] && boolFalse[sLow]) || (boolFalse[cLow] && boolTrue[sLow]) {
		return &CodeSuggestion{Kod: "odwrocona_logika", Opis: "Odwrócona wartość logiczna (PRAWDA↔FAŁSZ)"}
	}

	// Numeric off-by-one
	cNum, cErr := strconv.ParseFloat(correct, 64)
	sNum, sErr := strconv.ParseFloat(student, 64)
	if cErr == nil && sErr == nil {
		diff := math.Abs(cNum - sNum)
		if diff == 1 {
			return &CodeSuggestion{Kod: "off_by_one", Opis: fmt.Sprintf("Wynik o 1 za %s (student=%s, poprawna=%s)",
				map[bool]string{true: "mały", false: "duży"}[sNum < cNum], student, correct)}
		}
	}

	return nil
}
```

Add import `"strconv"` and `"math"` to `error_codes.go`.

In `commands.go`, add new command:

```go
func exerciseSuggestErrorCmd() *cobra.Command {
	var id, studentAnswer string

	cmd := &cobra.Command{
		Use:   "suggest-error",
		Short: "Suggest error code based on exercise and student answer",
		RunE: func(cmd *cobra.Command, args []string) error {
			if id == "" {
				return fatal("--id is required")
			}
			d := db(cmd)

			// Get exercise info
			var typNazwa, odpowiedz string
			err := d.QueryRow("SELECT typ_nazwa, odpowiedz FROM data.cwiczenia WHERE id = ?", id).
				Scan(&typNazwa, &odpowiedz)
			if err != nil {
				return notFound(fmt.Sprintf("exercise %s not found", id))
			}

			out := SuggestErrorOut{
				KodyDlaTypu: getCodesForType(typNazwa),
			}

			// Auto-detect if student answer provided
			if studentAnswer != "" {
				pattern := detectErrorPattern(odpowiedz, studentAnswer)
				if pattern != nil {
					out.AutoDetected = true
					out.Rekomendowany = pattern
					out.Powod = pattern.Opis
				}
			}

			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Exercise ID (e.g. 1.1)")
	cmd.Flags().StringVar(&studentAnswer, "student-answer", "", "Student's answer for auto-detection")
	return cmd
}
```

In `main.go`, register under `exercise` parent command (find where other exercise subcommands are added):

```go
exerciseCmd.AddCommand(exerciseSuggestErrorCmd())
```

**Step 4: Run tests to verify they pass**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run "TestSuggestError|TestDetectErrorPattern" -v`
Expected: PASS

**Step 5: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test ./... -v -count=1`
Expected: All tests PASS

**Step 6: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/error_codes.go matura_informatyka_rozszerzona/analiza/cli/commands.go matura_informatyka_rozszerzona/analiza/cli/types.go matura_informatyka_rozszerzona/analiza/cli/main.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "feat(cli): add exercise suggest-error with auto-detection"
```

---

### Task 5: `exercise check-answer` — auto-scoring for TEORIA

**Files:**
- Modify: `commands.go` (new `exerciseCheckAnswerCmd`)
- Modify: `types.go` (new `CheckAnswerOut`)
- Create: `normalize.go` (answer normalization logic — separate file for clarity)
- Modify: `main.go` (register command)
- Test: `main_test.go` (new tests)

**Context:** Auto-scoring works for types with deterministic answers: `sledzenie_algorytmu`, `test_prawda_falsz`, `konwersja_systemow_liczbowych`. For `analiza_algorytmu` — only when answer is a value (not explanation).

**Step 1: Write failing tests**

```go
func TestNormalizeAnswer(t *testing.T) {
	cases := []struct {
		input, expected string
	}{
		{"  42  ", "42"},
		{"42.0", "42"},
		{"PRAWDA", "prawda"},
		{" Fałsz ", "falsz"},
		{"13.00", "13"},
		{"  hello world  ", "hello world"},
		{"0013", "13"},
	}
	for _, c := range cases {
		got := normalizeAnswer(c.input)
		if got != c.expected {
			t.Errorf("normalizeAnswer(%q) = %q, want %q", c.input, got, c.expected)
		}
	}
}

func TestCheckAnswerAutoScorable(t *testing.T) {
	// These types should be auto-scorable
	scorable := []string{"sledzenie_algorytmu", "test_prawda_falsz", "konwersja_systemow_liczbowych"}
	for _, typ := range scorable {
		if !isAutoScorable(typ) {
			t.Errorf("expected %q to be auto-scorable", typ)
		}
	}
	// These should NOT be auto-scorable
	notScorable := []string{"projektowanie_algorytmu", "cyfry_liczby", "sql_group_by", "symulacja"}
	for _, typ := range notScorable {
		if isAutoScorable(typ) {
			t.Errorf("expected %q to NOT be auto-scorable", typ)
		}
	}
}

func TestCheckAnswerCorrect(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Find a sledzenie_algorytmu exercise and get its answer
	var id, odpowiedz string
	err := db.QueryRow(`SELECT id, odpowiedz FROM data.cwiczenia
		WHERE typ_nazwa = 'sledzenie_algorytmu' LIMIT 1`).Scan(&id, &odpowiedz)
	if err != nil {
		t.Skipf("no sledzenie exercise: %v", err)
	}

	result := checkAnswer(odpowiedz, odpowiedz)
	if !result.Poprawne {
		t.Errorf("expected correct answer to be poprawne for %s", id)
	}
	if result.Wynik != "pelne" {
		t.Errorf("expected wynik=pelne, got %q", result.Wynik)
	}
}

func TestCheckAnswerIncorrect(t *testing.T) {
	result := checkAnswer("42", "43")
	if result.Poprawne {
		t.Error("expected incorrect answer")
	}
	if result.Wynik != "zero" {
		t.Errorf("expected wynik=zero, got %q", result.Wynik)
	}
	if result.PoprawnaOdpowiedz == "" {
		t.Error("expected poprawna_odpowiedz to be set")
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run "TestNormalizeAnswer|TestCheckAnswer" -v`
Expected: FAIL — functions undefined

**Step 3: Implement**

Create `normalize.go`:

```go
package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// autoScorableTypes are types where CLI can check answers deterministically.
var autoScorableTypes = map[string]bool{
	"sledzenie_algorytmu":          true,
	"test_prawda_falsz":            true,
	"konwersja_systemow_liczbowych": true,
}

func isAutoScorable(typ string) bool {
	return autoScorableTypes[typ]
}

// normalizeAnswer normalizes an answer for comparison:
// trim whitespace, lowercase, normalize numbers (remove trailing .0, leading zeros).
func normalizeAnswer(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	// Replace Polish ł with l for falsz comparison
	s = strings.ReplaceAll(s, "ł", "l")

	// Try numeric normalization
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		if f == math.Trunc(f) {
			return fmt.Sprintf("%d", int(f))
		}
		return strconv.FormatFloat(f, 'f', -1, 64)
	}

	return s
}

// CheckAnswerResult is the internal result of checking an answer.
type CheckAnswerResult struct {
	Poprawne         bool   `json:"poprawne"`
	Wynik            string `json:"wynik"`
	PoprawnaOdpowiedz string `json:"poprawna_odpowiedz,omitempty"`
}

// checkAnswer compares student answer against correct answer with normalization.
func checkAnswer(correct, student string) CheckAnswerResult {
	nc := normalizeAnswer(correct)
	ns := normalizeAnswer(student)

	if nc == ns {
		return CheckAnswerResult{Poprawne: true, Wynik: "pelne"}
	}
	return CheckAnswerResult{
		Poprawne:          false,
		Wynik:             "zero",
		PoprawnaOdpowiedz: correct,
	}
}
```

In `types.go`, add:

```go
// CheckAnswerOut is what exercise check-answer returns
type CheckAnswerOut struct {
	ID                string `json:"id"`
	Poprawne          bool   `json:"poprawne"`
	Wynik             string `json:"wynik"`
	AutoScored        bool   `json:"auto_scored"`
	PoprawnaOdpowiedz string `json:"poprawna_odpowiedz,omitempty"`
}
```

In `commands.go`, add:

```go
func exerciseCheckAnswerCmd() *cobra.Command {
	var id, answer string

	cmd := &cobra.Command{
		Use:   "check-answer",
		Short: "Auto-score exercise answer (TEORIA types only)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if id == "" || answer == "" {
				return fatal("--id and --answer are required")
			}
			d := db(cmd)

			var typNazwa, odpowiedz string
			err := d.QueryRow("SELECT typ_nazwa, odpowiedz FROM data.cwiczenia WHERE id = ?", id).
				Scan(&typNazwa, &odpowiedz)
			if err != nil {
				return notFound(fmt.Sprintf("exercise %s not found", id))
			}

			if !isAutoScorable(typNazwa) {
				return fatal(fmt.Sprintf("typ %q is not auto-scorable. Use manual scoring.", typNazwa))
			}

			result := checkAnswer(odpowiedz, answer)
			jsonOut(CheckAnswerOut{
				ID:                 id,
				Poprawne:           result.Poprawne,
				Wynik:              result.Wynik,
				AutoScored:         true,
				PoprawnaOdpowiedz: result.PoprawnaOdpowiedz,
			})
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Exercise ID (e.g. 1.1)")
	cmd.Flags().StringVar(&answer, "answer", "", "Student's answer")
	return cmd
}
```

In `main.go`, register: `exerciseCmd.AddCommand(exerciseCheckAnswerCmd())`

**Step 4: Run tests to verify they pass**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run "TestNormalizeAnswer|TestCheckAnswer" -v`
Expected: PASS

**Step 5: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test ./... -v -count=1`
Expected: All tests PASS

**Step 6: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/normalize.go matura_informatyka_rozszerzona/analiza/cli/commands.go matura_informatyka_rozszerzona/analiza/cli/types.go matura_informatyka_rozszerzona/analiza/cli/main.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "feat(cli): add exercise check-answer with auto-scoring for TEORIA types"
```

---

### Task 6: SKILL.md HARD GATES + build + QA

**Files:**
- Modify: SKILL.md (matura skill — add 4 new HARD GATES)
- Run: `build.sh` (rebuild binaries + reimport matura.db)
- Run: `test_qa.sh` (full QA suite)

**Step 1: Rebuild CLI**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && ./build.sh`
Expected: macOS + Windows binaries built, matura.db reimported

**Step 2: Run full Go test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test ./... -v -count=1`
Expected: All tests PASS

**Step 3: Run QA suite**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh`
Expected: All 7 layers pass. If Layer 1 (CLI smoke) fails on new commands, add smoke tests for them.

**Step 4: Update SKILL.md with new HARD GATES**

Find the SKILL.md for the matura skill and add these gates in the appropriate section:

```markdown
### HARD GATE: Scoring
- For auto-scorable types (sledzenie_algorytmu, test_prawda_falsz, konwersja_systemow_liczbowych):
  MUST use `exercise check-answer --id X --answer Y` instead of manual scoring.
- For all types: `progress update --punktacja` MUST be one of: pelne, prawie_pelne, czesciowe, minimalne, zero.

### HARD GATE: Error Codes
- Before `progress blad --kod X`: MUST call `exercise suggest-error --id X [--student-answer Y]` first.
- If suggest-error returns auto_detected=true with rekomendowany, use that code by default.
- Override only with explicit reasoning.

### HARD GATE: Hint Sequencing
- CLI enforces: no second hint without student attempt (HINT_BLOCKED_NO_ATTEMPT).
- If blocked: ask student Socratic question, wait for answer, register via `progress blad`, then retry hints.
```

**Step 5: Commit**

```bash
git add -A
git commit -m "feat: SKILL.md hard gates for deterministic pedagogy + build"
```
