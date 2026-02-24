# CLI Guardrails + SKILL.md + Test-Tutor Redesign — Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Move validation logic from SKILL.md (LLM must remember) into Go CLI (binary enforces), then simplify SKILL.md and redesign test-tutor for higher scores and lower variance.

**Architecture:** CLI gains guardrail layer — active_exercises tracking, lazy loading enforcement, hint delay enforcement, error code validation, coaching_actions generation. SKILL.md simplified (checklist 9→6 steps). Test-tutor switches to 2-layer scoring (60% binary checkpoints + 40% holistic) with deterministic student scripts.

**Tech Stack:** Go 1.21+, SQLite (modernc.org/sqlite), Cobra CLI, Markdown skills

**Design doc:** `docs/plans/2026-02-23-cli-guardrails-redesign-design.md`

---

## Phase 1: CLI Guardrails (Go code)

### Task 1: Schema migration v5→v6 (active_exercises table)

**Files:**
- Modify: `analiza/cli/database.go:12` (version bump)
- Modify: `analiza/cli/database.go` (add migration + table creation)

**Step 1: Add active_exercises table to schema**

In `database.go`, bump `currentSchemaVersion` from 5 to 6.

Add to `ensureSchema()` (in the CREATE TABLE IF NOT EXISTS block):

```go
const currentSchemaVersion = 6
```

```sql
CREATE TABLE IF NOT EXISTS active_exercises (
    exercise_id TEXT PRIMARY KEY,
    typ TEXT NOT NULL,
    fetched_at TEXT NOT NULL DEFAULT (datetime('now')),
    attempt_count INTEGER NOT NULL DEFAULT 0,
    hint_delay INTEGER NOT NULL DEFAULT 1,
    hints_fetched INTEGER NOT NULL DEFAULT 0,
    answer_fetched INTEGER NOT NULL DEFAULT 0
);
```

Add migration v5→v6 in the migration switch:

```go
case 5:
    _, err = tx.Exec(`CREATE TABLE IF NOT EXISTS active_exercises (
        exercise_id TEXT PRIMARY KEY,
        typ TEXT NOT NULL,
        fetched_at TEXT NOT NULL DEFAULT (datetime('now')),
        attempt_count INTEGER NOT NULL DEFAULT 0,
        hint_delay INTEGER NOT NULL DEFAULT 1,
        hints_fetched INTEGER NOT NULL DEFAULT 0,
        answer_fetched INTEGER NOT NULL DEFAULT 0
    )`)
    if err != nil {
        return fmt.Errorf("migrate v5→v6: %w", err)
    }
    fallthrough
```

**Step 2: Write test for migration**

In `main_test.go`, add:

```go
func TestMigrationV5ToV6(t *testing.T) {
    dir := testDir(t)
    d := openTestDB(t, dir)
    defer d.Close()

    // Table should exist after migration
    var count int
    err := d.QueryRow("SELECT COUNT(*) FROM active_exercises").Scan(&count)
    if err != nil {
        t.Fatalf("active_exercises table should exist: %v", err)
    }
    if count != 0 {
        t.Fatalf("expected 0 rows, got %d", count)
    }
}
```

**Step 3: Run test**

```bash
cd analiza/cli && go test -run TestMigrationV5ToV6 -v
```

Expected: PASS

**Step 4: Commit**

```
feat(cli): add active_exercises table (schema v6)
```

---

### Task 2: Register exercise fetch in active_exercises

**Files:**
- Modify: `analiza/cli/commands.go` (exerciseNextCmd, exerciseQuestionCmd, exerciseReviewCmd)
- Create helper: `analiza/cli/guardrails.go`

**Step 1: Create guardrails.go with registerFetch**

```go
package main

import "database/sql"

// registerFetch records an exercise fetch in active_exercises.
// Replaces any previous entry for the same exercise_id.
func registerFetch(d *sql.DB, exerciseID, typ string, hintDelay int) error {
    _, err := d.Exec(`
        INSERT OR REPLACE INTO active_exercises
            (exercise_id, typ, fetched_at, attempt_count, hint_delay, hints_fetched, answer_fetched)
        VALUES (?, ?, datetime('now'), 0, ?, 0, 0)
    `, exerciseID, typ, hintDelay)
    return err
}

// incrementAttempt increments attempt_count for an exercise.
// Called by progress blad.
func incrementAttempt(d *sql.DB, exerciseID string) error {
    _, err := d.Exec(`
        UPDATE active_exercises SET attempt_count = attempt_count + 1
        WHERE exercise_id = ?
    `, exerciseID)
    return err
}

// clearExercise removes an exercise from active tracking.
// Called by progress update (exercise completed).
func clearExercise(d *sql.DB, exerciseID string) error {
    _, err := d.Exec(`DELETE FROM active_exercises WHERE exercise_id = ?`, exerciseID)
    return err
}

// checkCanFetchAnswer checks if answer can be fetched (attempt_count >= 1).
func checkCanFetchAnswer(d *sql.DB, exerciseID string) (bool, int, error) {
    var attemptCount int
    err := d.QueryRow(`
        SELECT attempt_count FROM active_exercises WHERE exercise_id = ?
    `, exerciseID).Scan(&attemptCount)
    if err == sql.ErrNoRows {
        // Not tracked = allow (backwards compat)
        return true, 0, nil
    }
    if err != nil {
        return false, 0, err
    }
    return attemptCount >= 1, attemptCount, nil
}

// checkCanFetchHints checks if hints can be fetched (attempt_count >= hint_delay).
func checkCanFetchHints(d *sql.DB, exerciseID string) (bool, int, int, error) {
    var attemptCount, hintDelay int
    err := d.QueryRow(`
        SELECT attempt_count, hint_delay FROM active_exercises WHERE exercise_id = ?
    `, exerciseID).Scan(&attemptCount, &hintDelay)
    if err == sql.ErrNoRows {
        return true, 0, 1, nil
    }
    if err != nil {
        return false, 0, 0, err
    }
    return attemptCount >= hintDelay, attemptCount, hintDelay, nil
}
```

**Step 2: Call registerFetch in exerciseNextCmd**

In `commands.go`, after the exercise is resolved in `exerciseNextCmd` (around line ~2030, after building `QuestionOut`), add:

```go
// Register fetch for guardrails
if err := registerFetch(d, out.Exercise.ID, out.Exercise.TypNazwa, out.Exercise.Coaching.HintDelay); err != nil {
    // Non-fatal: guardrail tracking, don't break main flow
    fmt.Fprintf(os.Stderr, "warning: registerFetch: %v\n", err)
}
```

Do the same in `exerciseQuestionCmd` (around line ~470) and `exerciseReviewCmd` (for each returned exercise).

**Step 3: Call incrementAttempt in progressBladCmd**

In `progressBladCmd` (around line ~1280), after the INSERT:

```go
if err := incrementAttempt(d, exerciseID); err != nil {
    fmt.Fprintf(os.Stderr, "warning: incrementAttempt: %v\n", err)
}
```

**Step 4: Call clearExercise in progressUpdateCmd**

In `progressUpdateCmd` (around line ~795), after the transaction commits:

```go
if err := clearExercise(d, id); err != nil {
    fmt.Fprintf(os.Stderr, "warning: clearExercise: %v\n", err)
}
```

**Step 5: Write tests**

```go
func TestRegisterFetchAndClear(t *testing.T) {
    dir := testDir(t)
    d := openTestDB(t, dir)
    defer d.Close()

    // Register
    if err := registerFetch(d, "7.1", "cyfry_liczby", 2); err != nil {
        t.Fatal(err)
    }

    // Check answer blocked
    can, attempts, _ := checkCanFetchAnswer(d, "7.1")
    if can {
        t.Fatal("answer should be blocked before any attempt")
    }
    if attempts != 0 {
        t.Fatalf("expected 0 attempts, got %d", attempts)
    }

    // Increment attempt
    incrementAttempt(d, "7.1")
    can, attempts, _ = checkCanFetchAnswer(d, "7.1")
    if !can {
        t.Fatal("answer should be allowed after 1 attempt")
    }

    // Check hints blocked (delay=2, attempt=1)
    canH, att, delay, _ := checkCanFetchHints(d, "7.1")
    if canH {
        t.Fatalf("hints should be blocked (attempt=%d < delay=%d)", att, delay)
    }

    // Second attempt
    incrementAttempt(d, "7.1")
    canH, _, _, _ = checkCanFetchHints(d, "7.1")
    if !canH {
        t.Fatal("hints should be allowed after 2 attempts (delay=2)")
    }

    // Clear
    clearExercise(d, "7.1")
    can, _, _ = checkCanFetchAnswer(d, "7.1")
    if !can {
        t.Fatal("after clear, should be allowed (no tracking)")
    }
}
```

**Step 6: Run tests**

```bash
cd analiza/cli && go test -run TestRegisterFetch -v
```

**Step 7: Commit**

```
feat(cli): track exercise fetch + attempt count in active_exercises
```

---

### Task 3: Lazy loading enforcement (answer + hints)

**Files:**
- Modify: `analiza/cli/commands.go` (exerciseAnswerCmd, exerciseHintsCmd)

**Step 1: Block exercise answer before attempt**

In `exerciseAnswerCmd` (line ~515), before the query, add:

```go
can, attempts, err := checkCanFetchAnswer(d, id)
if err != nil {
    return fmt.Errorf("checkCanFetchAnswer: %w", err)
}
if !can {
    // Return structured error JSON instead of crashing
    out := map[string]interface{}{
        "status":        "LAZY_LOADING_BLOCKED",
        "exercise_id":   id,
        "attempt_count": attempts,
        "action":        "Student hasn't attempted yet. Record attempt first via 'progress blad' or 'progress update'.",
    }
    return printJSON(out)
}
```

**Step 2: Block exercise hints before hint_delay**

In `exerciseHintsCmd` (line ~490), before the query, add:

```go
canH, attempts, delay, err := checkCanFetchHints(d, id)
if err != nil {
    return fmt.Errorf("checkCanFetchHints: %w", err)
}
if !canH {
    out := map[string]interface{}{
        "status":     "HINT_LOCKED",
        "exercise_id": id,
        "attempt":    attempts,
        "hint_delay": delay,
        "action":     "Zadaj pytanie sokratejskie BEZ hintow",
    }
    return printJSON(out)
}
```

Also mark hints as fetched:

```go
d.Exec(`UPDATE active_exercises SET hints_fetched = 1 WHERE exercise_id = ?`, id)
```

**Step 3: Write integration test**

```go
func TestLazyLoadingEnforcement(t *testing.T) {
    dir := testDir(t)
    d := openTestDB(t, dir)
    defer d.Close()

    // Fetch exercise
    registerFetch(d, "1.1", "sledzenie_algorytmu", 1)

    // Answer should be blocked
    can, _, _ := checkCanFetchAnswer(d, "1.1")
    if can {
        t.Fatal("answer should be blocked before attempt")
    }

    // Record error → attempt_count = 1
    incrementAttempt(d, "1.1")

    // Now answer should be allowed
    can, _, _ = checkCanFetchAnswer(d, "1.1")
    if !can {
        t.Fatal("answer should be allowed after attempt")
    }
}
```

**Step 4: Run tests**

```bash
cd analiza/cli && go test -run TestLazyLoading -v
```

**Step 5: Commit**

```
feat(cli): enforce lazy loading — block answer/hints before attempt
```

---

### Task 4: Error code validation whitelist

**Files:**
- Create: `analiza/cli/error_codes.go`
- Modify: `analiza/cli/commands.go` (progressBladCmd)

**Step 1: Create error_codes.go with whitelist**

```go
package main

// errorCodeWhitelist maps exercise type → allowed error codes.
// CLI rejects any code not in the whitelist for the given type.
var errorCodeWhitelist = map[string][]string{
    "sledzenie_algorytmu": {
        "mylenie_div_mod", "zla_kolejnosc_sledzenia", "pominiecie_bazy_rekurencji",
        "zly_mnoznik", "brak_tabeli_sledzenia", "zla_parzystosc_cyfry", "bledne_wciecia_blok",
    },
    "projektowanie_algorytmu": {
        "zly_algorytm", "brak_warunku_stopu", "bledna_skladnia_pseudokod",
        "niepoprawna_petla", "brak_inicjalizacji",
    },
    "analiza_algorytmu": {
        "zla_zlozonosc_klasa", "brak_uzasadnienia_zlozonosc", "mylenie_avg_worst",
        "zly_kontrprzyklad", "brak_wzoru",
    },
    "test_prawda_falsz": {
        "brak_uzasadnienia_pf", "mylenie_avg_worst_pf", "nieprecyzyjne_uzasadnienie",
        "pomylenie_stabilnosci_sortowania",
    },
    "konwersja_systemow_liczbowych": {
        "zla_baza_konwersji", "zla_kolejnosc_reszt", "brak_zapisu_posredniego",
        "zle_grupowanie_bitow", "blad_uzupelnienia_do_2",
    },
    "teoria_bezpieczenstwa": {
        "mylenie_typow_malware", "mylenie_szyfrowania_sym_asym",
        "mylenie_protokolow", "brak_rozroznienia_klucz_pub_pryw",
    },
    // SQL types share same codes
    "sql_group_by":       {"brak_group_by", "zly_join_warunek", "brak_having", "zla_agregacja", "null_zamiast_is_null", "count_star_vs_kolumna", "zla_kolejnosc_klauzul"},
    "sql_podzapytania":   {"brak_group_by", "zly_join_warunek", "brak_having", "zla_agregacja", "null_zamiast_is_null", "count_star_vs_kolumna", "zla_kolejnosc_klauzul"},
    "sql_join":           {"brak_group_by", "zly_join_warunek", "brak_having", "zla_agregacja", "null_zamiast_is_null", "count_star_vs_kolumna", "zla_kolejnosc_klauzul"},
    "sql_select_where":   {"brak_group_by", "zly_join_warunek", "brak_having", "zla_agregacja", "null_zamiast_is_null", "count_star_vs_kolumna", "zla_kolejnosc_klauzul"},
    // IMPLEMENTACJA types share same codes
    "cyfry_liczby":       {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji", "mylenie_div_mod"},
    "napisy":             {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
    "zlozone":            {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
    "zliczanie":          {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
    "minmax":             {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
    "sekwencje":          {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
    "obrazy_2D":          {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
    "geometryczne":       {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
    // ARKUSZ types share same codes
    "agregacja_warunkowa":  {"zle_adresowanie", "brak_dolara", "zla_formula_warunkowa", "stala_zamiast_odwolania", "brak_kolumny_pomocniczej"},
    "symulacja":            {"zle_adresowanie", "brak_dolara", "zla_formula_warunkowa", "stala_zamiast_odwolania", "brak_kolumny_pomocniczej"},
    "wykres":               {"zle_adresowanie", "brak_dolara", "zla_formula_warunkowa", "stala_zamiast_odwolania", "brak_kolumny_pomocniczej"},
    "agregacja_podstawowa": {"zle_adresowanie", "brak_dolara", "zla_formula_warunkowa", "stala_zamiast_odwolania", "brak_kolumny_pomocniczej"},
    "transformacja":        {"zle_adresowanie", "brak_dolara", "zla_formula_warunkowa", "stala_zamiast_odwolania", "brak_kolumny_pomocniczej"},
}

// validateErrorCode checks if kod is allowed for the given typ.
// Returns allowed list if invalid.
func validateErrorCode(typ, kod string) (valid bool, allowed []string) {
    codes, ok := errorCodeWhitelist[typ]
    if !ok {
        // Unknown type — allow any code (backwards compat)
        return true, nil
    }
    for _, c := range codes {
        if c == kod {
            return true, nil
        }
    }
    return false, codes
}
```

**Step 2: Add validation to progressBladCmd**

In `commands.go`, `progressBladCmd` (line ~1260), after parsing flags:

```go
valid, allowed := validateErrorCode(typ, kod)
if !valid {
    return fmt.Errorf("ERROR: kod '%s' niedostepny dla typ '%s'. Dozwolone: %v", kod, typ, allowed)
}
```

**Step 3: Make --hint required**

Change `--hint` from optional to required in progressBladCmd flag definition:

```go
bladCmd.Flags().IntVar(&hintLevel, "hint", -1, "Hint level (0=before any hint, 1/2/3=after hint)")
bladCmd.MarkFlagRequired("hint")
```

Wait — this would break backwards compat. Instead, validate:

```go
if hintLevel < 0 {
    return fmt.Errorf("ERROR: --hint wymagany. Uzyj --hint 0 (przed hintem) lub --hint 1/2/3 (po hincie)")
}
```

**Step 4: Write tests**

```go
func TestValidateErrorCode(t *testing.T) {
    // Valid
    ok, _ := validateErrorCode("sql_group_by", "brak_having")
    if !ok {
        t.Fatal("brak_having should be valid for sql_group_by")
    }

    // Invalid
    ok, allowed := validateErrorCode("sql_group_by", "brak_inicjalizacji")
    if ok {
        t.Fatal("brak_inicjalizacji should be invalid for sql_group_by")
    }
    if len(allowed) == 0 {
        t.Fatal("should return allowed list")
    }

    // Unknown type
    ok, _ = validateErrorCode("unknown_type", "anything")
    if !ok {
        t.Fatal("unknown type should allow any code")
    }

    // cyfry_liczby has mylenie_div_mod (cross-category)
    ok, _ = validateErrorCode("cyfry_liczby", "mylenie_div_mod")
    if !ok {
        t.Fatal("mylenie_div_mod should be valid for cyfry_liczby")
    }
}
```

**Step 5: Run tests**

```bash
cd analiza/cli && go test -run TestValidateErrorCode -v
```

**Step 6: Commit**

```
feat(cli): validate error codes against per-type whitelist + require --hint
```

---

### Task 5: coaching_actions generation

**Files:**
- Modify: `analiza/cli/types.go` (add CoachingActions to QuestionOut/ExerciseNextOut)
- Modify: `analiza/cli/commands.go` (buildCoaching or new function)

**Step 1: Add CoachingActions to output structs**

In `types.go`, add to `Coaching` struct:

```go
type Coaching struct {
    StudentLevel   string   `json:"student_level"`
    HintDelay      int      `json:"hint_delay"`
    LeechTags      []string `json:"leech_tags"`
    PastMistakes   []string `json:"past_mistakes"`
    PreviousResult string   `json:"previous_result,omitempty"`
    Actions        []string `json:"coaching_actions"` // NEW
}
```

**Step 2: Generate actions in buildCoaching**

At the end of `buildCoaching()` in `commands.go` (around line ~225), before return:

```go
// Generate coaching_actions
var actions []string
if len(c.LeechTags) > 0 {
    for _, tag := range c.LeechTags {
        actions = append(actions, fmt.Sprintf("WARN_LEECH: Tag '%s' sprawia Ci trudnosc — zwroc uwage", tag))
    }
}
if len(c.PastMistakes) > 0 {
    for _, m := range c.PastMistakes {
        actions = append(actions, fmt.Sprintf("MENTION_PAST: Ostatnio mialeS problem z '%s'", m))
    }
}
if c.HintDelay >= 2 {
    actions = append(actions, fmt.Sprintf("HINT_DELAY: %d (Od teraz mniej podpowiedzi — rozwijasz samodzielnosc)", c.HintDelay))
}
c.Actions = actions
```

**Step 3: Write test**

```go
func TestCoachingActions(t *testing.T) {
    dir := testDir(t)
    d := openTestDB(t, dir)
    defer d.Close()

    // Insert leech tag
    d.Exec(`INSERT INTO progress_tagi (tag, lapses, stability, last_review, poziom, nastepna_powtorka, difficulty, reps, state)
        VALUES ('cyfry-mod-div', 4, 1.0, '2026-01-01', 1, '2026-01-01', 5.0, 4, 2)`)
    d.Exec(`INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'srednie', 4)`)

    c := buildCoaching(d, "cyfry_liczby", []string{"cyfry-mod-div"})

    if len(c.Actions) == 0 {
        t.Fatal("expected coaching_actions")
    }
    foundLeech := false
    foundDelay := false
    for _, a := range c.Actions {
        if strings.Contains(a, "WARN_LEECH") {
            foundLeech = true
        }
        if strings.Contains(a, "HINT_DELAY") {
            foundDelay = true
        }
    }
    if !foundLeech {
        t.Fatal("expected WARN_LEECH action")
    }
    if !foundDelay {
        t.Fatal("expected HINT_DELAY action (familiar = delay 2)")
    }
}
```

**Step 4: Run tests**

```bash
cd analiza/cli && go test -run TestCoachingActions -v
```

**Step 5: Commit**

```
feat(cli): generate coaching_actions in exercise responses
```

---

### Task 6: Auto-diagnose in progress update

**Files:**
- Modify: `analiza/cli/commands.go` (progressUpdateCmd)
- Modify: `analiza/cli/types.go` (ProgressUpdateOut)

**Step 1: Add Diagnose field to ProgressUpdateOut**

In `types.go`, find `ProgressUpdateOut` and add:

```go
AutoDiagnose *DiagnoseOut `json:"auto_diagnose,omitempty"`
```

**Step 2: Trigger diagnose when count % 5 == 0**

In `progressUpdateCmd`, after the transaction commits (around line ~795), add:

```go
// Auto-diagnose every 5 exercises
var totalDone int
d.QueryRow(`SELECT COALESCE(value, '0') FROM progress_meta WHERE key='cwiczenia_lacznie'`).Scan(&totalDone)
if totalDone > 0 && totalDone%5 == 0 {
    diag, err := runDiagnose(d, typ, 3) // top 3 errors
    if err == nil {
        result.AutoDiagnose = diag
    }
}
```

Extract diagnose logic from `progressDiagnoseCmd` into a reusable `runDiagnose(d, typ, limit)` function.

**Step 3: Write test**

```go
func TestAutoDiagnose(t *testing.T) {
    // Insert 4 completed exercises, verify no auto_diagnose
    // Insert 5th, verify auto_diagnose present in response
}
```

**Step 4: Run tests + commit**

```bash
cd analiza/cli && go test -run TestAutoDiagnose -v
```

```
feat(cli): auto-diagnose in progress update every 5 exercises
```

---

### Task 7: Build, reimport, run existing tests

**Step 1: Build**

```bash
cd analiza/cli && go build -o matura .
```

Expected: no errors

**Step 2: Run all existing tests**

```bash
cd analiza/cli && go test -v
```

Expected: all tests pass (existing + new guardrail tests)

**Step 3: Reimport data**

```bash
cd analiza/cli && ./matura data import --source ../
```

**Step 4: Run test_qa.sh Layers 1 + 5**

```bash
cd analiza && ./test_qa.sh --layer 1
cd analiza && ./test_qa.sh --layer 5
```

Expected: all pass. Layer 1 smoke tests exercise CLI commands — some may need updating if `progress blad` now requires `--hint`.

**Step 5: Fix any Layer 1 failures**

Update `test_qa.sh` Layer 1 commands that use `progress blad` to include `--hint 0`.

**Step 6: Commit**

```
fix(cli): update test_qa.sh for --hint requirement in progress blad
```

---

## Phase 2: SKILL.md Rewrite

### Task 8: Rewrite SKILL.md

**Files:**
- Modify: `.claude/skills/matura/SKILL.md`

**Step 1: Read current SKILL.md fully**

```bash
cat .claude/skills/matura/SKILL.md
```

**Step 2: Rewrite with new structure**

Key changes:
1. **Section B (CLI Reference)**: Remove `exercise question` from table. Add guardrail responses (LAZY_LOADING_BLOCKED, HINT_LOCKED) to output docs.
2. **Section C**: Add first_session checklist with explicit lazy loading steps.
3. **Section D**: Replace "exercise question or exercise next" with "ONLY exercise next".
4. **Section E**: Replace coaching JSON parsing instructions with "Read coaching_actions from response and include each naturally in dialog."
5. **Section F**: Simplify checklist from 9→6 steps. Remove steps that CLI now enforces (lazy loading check, hint_delay check, diagnose trigger).
6. **Remove error codes section** (moved to CLI validation — mention "CLI validates error codes; use any plausible code and CLI will suggest corrections").

**Step 3: Run test_qa.sh Layer 3 (SKILL lint)**

```bash
cd analiza && ./test_qa.sh --layer 3
```

Fix any lint failures from changed sections.

**Step 4: Commit**

```
refactor(skill): simplify SKILL.md — CLI guardrails enforce validation
```

---

## Phase 3: Test-Tutor Rewrite

### Task 9: Rewrite test-tutor SKILL.md

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md`

**Step 1: Read current test-tutor SKILL.md**

**Step 2: Rewrite with new design**

Key changes:
1. **Section 4 (Rubric)**: Replace 8 criteria with 5 (CLI compliance, metoda sokratejska, ton i jezyk, coaching reaction, scenario-specific). Add 2-layer scoring description.
2. **Section 5 (Scenarios)**: Add fixed student scripts for each scenario (deterministic responses).
3. **Section 6 (Agent prompt)**: Add binary checkpoint lists per scenario. Add anchor examples for holistic criteria (metoda sokratejska 5/3/1, ton 5/3/1). Change scoring formula to `0.6 * L1 + 0.4 * (L2/5 * 100)`.
4. **Section 3 (Personas)**: Keep accuracy/behavior descriptions but note that RESPONSES are scripted (persona affects script content, not randomness).

**Step 3: Run test_qa.sh Layer 3**

```bash
cd analiza && ./test_qa.sh --layer 3
```

**Step 4: Commit**

```
refactor(test-tutor): 2-layer scoring, 5 criteria, deterministic scripts
```

---

## Phase 4: Validation

### Task 10: Full test suite

**Step 1: Build CLI**

```bash
cd analiza/cli && ./build.sh
```

**Step 2: Run full test_qa.sh**

```bash
cd analiza && ./test_qa.sh
```

Expected: all layers pass

**Step 3: Fix any failures**

---

### Task 11: Validation test-tutor run

**Step 1: Run quick test-tutor**

```bash
# In Claude Code session:
/test-tutor quick
```

Compare score vs previous runs. Expected: higher score due to CLI guardrails reducing compliance errors.

**Step 2: Run full test-tutor**

```bash
/test-tutor
```

Expected:
- Overall > 92/100
- Per-scenario variance < 10 pkt (vs current 23 pkt)
- No lazy loading violations (CLI blocks them)

**Step 3: Compare with baseline**

Read previous reports in `analiza/test_pedagogical/reports/` and compute delta.

---

## Summary of deliverables

| # | Deliverable | Files changed |
|---|-------------|---------------|
| 1 | Schema v6 (active_exercises) | database.go |
| 2 | registerFetch + incrementAttempt + clearExercise | guardrails.go (new), commands.go |
| 3 | Lazy loading enforcement | commands.go (answer + hints) |
| 4 | Error code whitelist + --hint required | error_codes.go (new), commands.go |
| 5 | coaching_actions generation | types.go, commands.go |
| 6 | Auto-diagnose in progress update | commands.go, types.go |
| 7 | Build + test existing suite | build.sh, test_qa.sh |
| 8 | SKILL.md rewrite | .claude/skills/matura/SKILL.md |
| 9 | Test-tutor rewrite | .claude/skills/test-tutor/SKILL.md |
| 10 | Full test suite validation | test_qa.sh |
| 11 | Test-tutor validation runs | test_pedagogical/reports/ |
