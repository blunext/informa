# Auto-Reset Weight Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Eliminate `--weight-reset` flag by auto-resetting `session_context_weight` inside `progress status`, making weight reset impossible to forget.

**Architecture:** `progress status` becomes the session-start signal that resets weight to 0. Mid-session callers (`status` student command, proactive detection) are remapped to `progress diagnose`, which gains new fields (`rekomendacja`, `retencja_szacowana`, `leech_tagi`, `zaleglosci`). The `--weight-reset` flag is removed from `exerciseNextCmd`.

**Tech Stack:** Go (commands.go, types.go, main_test.go), SKILL.md, test_qa.sh

**Design doc:** `docs/plans/2026-02-22-weight-reset-auto-design.md`

---

### Task 1: Add weight reset to `progressStatusCmd`

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/commands.go:897-899`
- Test: `matura_informatyka_rozszerzona/analiza/cli/main_test.go`

**Step 1: Write the failing test**

In `main_test.go`, add a new test after `TestExerciseNextWeight` (line ~1566):

```go
func TestProgressStatusResetsWeight(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Set weight to 50
	db.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '50')`)

	// Run progress status command
	out := runCmd(t, dir, "progress", "status")

	// Verify weight was reset to 0
	var wStr string
	var w int
	if err := db.QueryRow(`SELECT value FROM progress_meta WHERE key = 'session_context_weight'`).Scan(&wStr); err != nil {
		t.Fatalf("could not read weight: %v", err)
	}
	fmt.Sscanf(wStr, "%d", &w)
	if w != 0 {
		t.Errorf("after progress status: got weight %d, want 0", w)
	}

	// Verify output is valid JSON with expected fields
	if !json.Valid([]byte(out)) {
		t.Errorf("progress status output is not valid JSON")
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestProgressStatusResetsWeight -v`
Expected: FAIL — weight stays at 50 because `progressStatusCmd` doesn't reset it yet.

**Step 3: Write minimal implementation**

In `commands.go`, inside `progressStatusCmd()`, add weight reset right after `out := ProgressStatusOut{}` (line 899):

```go
out := ProgressStatusOut{}

// Auto-reset session context weight — progress status is the session-start signal
d.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '0')`)
```

**Step 4: Run test to verify it passes**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestProgressStatusResetsWeight -v`
Expected: PASS

**Step 5: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/commands.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "feat(cli): auto-reset session_context_weight in progress status"
```

---

### Task 2: Extend `DiagnoseOut` with status fields

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/types.go:344-347`
- Modify: `matura_informatyka_rozszerzona/analiza/cli/commands.go:1294-1358`
- Test: `matura_informatyka_rozszerzona/analiza/cli/main_test.go`

**Step 1: Write the failing test**

In `main_test.go`, add:

```go
func TestDiagnoseHasStatusFields(t *testing.T) {
	dir := testDir(t)

	out := runCmd(t, dir, "progress", "diagnose")

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(out), &result); err != nil {
		t.Fatalf("invalid JSON: %v", err)
	}

	// Check new fields exist
	for _, field := range []string{"zaleglosci", "leech_tagi"} {
		if _, ok := result[field]; !ok {
			t.Errorf("missing field %q in diagnose output", field)
		}
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestDiagnoseHasStatusFields -v`
Expected: FAIL — `zaleglosci` and `leech_tagi` fields missing from `DiagnoseOut`.

**Step 3: Update `DiagnoseOut` type**

In `types.go`, change `DiagnoseOut` (line 344-347):

```go
type DiagnoseOut struct {
	TopBledy          []DiagnoseEntry `json:"top_bledy"`
	Total             int             `json:"total"`
	Zaleglosci        int             `json:"zaleglosci"`
	LeechTagi         []LeechTagOut   `json:"leech_tagi"`
	RetencjaSzacowana *float64        `json:"retencja_szacowana,omitempty"`
	Rekomendacja      string          `json:"rekomendacja"`
}
```

**Step 4: Add logic to `progressDiagnoseCmd`**

In `commands.go`, inside `progressDiagnoseCmd()`, add after the existing `out.TopBledy` loop (before `jsonOut(out)`), at approximately line 1349:

```go
// --- Status fields (reused from progressStatusCmd) ---

// Overdue reviews
today := time.Now().Format("2006-01-02")
d.QueryRow("SELECT COUNT(*) FROM progress_tagi WHERE nastepna_powtorka <= ?", today).Scan(&out.Zaleglosci)

// Leech tags (3+ lapses)
leechRows, leechErr := d.Query(`SELECT tag, COALESCE(lapses, 0), COALESCE(stability, 0)
	FROM progress_tagi WHERE lapses >= 3 ORDER BY lapses DESC`)
if leechErr == nil {
	for leechRows.Next() {
		var lt LeechTagOut
		leechRows.Scan(&lt.Tag, &lt.Lapses, &lt.Stability)
		out.LeechTagi = append(out.LeechTagi, lt)
	}
	leechRows.Close()
}
if out.LeechTagi == nil {
	out.LeechTagi = []LeechTagOut{}
}

// Retencja szacowana
fsrsParams := DefaultFSRSParams()
retRows, retErr := d.Query(`SELECT tag, COALESCE(stability, 1.0), COALESCE(last_review, '')
	FROM progress_tagi WHERE nastepna_powtorka IS NOT NULL`)
if retErr == nil {
	var totalR float64
	var retCount int
	for retRows.Next() {
		var tag, lastReview string
		var stability float64
		retRows.Scan(&tag, &stability, &lastReview)
		elapsed := daysBetween(lastReview, today)
		r := fsrsParams.Retrievability(elapsed, stability)
		totalR += r
		retCount++
	}
	retRows.Close()
	if retCount > 0 {
		avgR := totalR / float64(retCount)
		out.RetencjaSzacowana = &avgR
	}
}

// Rekomendacja — build a minimal ProgressStatusOut to reuse computeRekomendacja
statusOut := ProgressStatusOut{
	Zaleglosci: out.Zaleglosci,
	LeechTagi:  out.LeechTagi,
}
if out.RetencjaSzacowana != nil {
	statusOut.RetencjaSzacowana = out.RetencjaSzacowana
}
// Populate per-typ and per-kategoria for rekomendacja
typRows, typErr := d.Query(`
	SELECT c.typ_nazwa,
		COALESCE(p.poziom_trudnosci, 'latwe') as poziom,
		COALESCE(p.streak, 0) as streak,
		COUNT(DISTINCT z.id) as zrobione,
		COUNT(DISTINCT c.id) as dostepne
	FROM data.cwiczenia c
	LEFT JOIN progress_typy p ON p.typ = c.typ_nazwa
	LEFT JOIN progress_zrobione z ON z.id = c.id
	GROUP BY c.typ_nazwa ORDER BY c.typ_nazwa`)
if typErr == nil {
	typKatMap := map[string]string{}
	katQueryRows, _ := d.Query("SELECT DISTINCT typ_nazwa, kategoria FROM data.cwiczenia")
	if katQueryRows != nil {
		for katQueryRows.Next() {
			var t, k string
			katQueryRows.Scan(&t, &k)
			typKatMap[t] = k
		}
		katQueryRows.Close()
	}

	katAgg := map[string]*KategoriaStatusOut{}
	for typRows.Next() {
		var ts TypStatusOut
		typRows.Scan(&ts.Typ, &ts.PoziomTrudnosci, &ts.Streak, &ts.Zrobione, &ts.Dostepne)
		statusOut.PerTyp = append(statusOut.PerTyp, ts)
		kat := typKatMap[ts.Typ]
		if kat == "" {
			continue
		}
		if katAgg[kat] == nil {
			katAgg[kat] = &KategoriaStatusOut{Kategoria: kat}
		}
		katAgg[kat].TypyTotal++
		if ts.Zrobione > 0 {
			katAgg[kat].TypyRuszane++
		}
		katAgg[kat].Zrobione += ts.Zrobione
		katAgg[kat].Dostepne += ts.Dostepne
		katAgg[kat].AvgStreak += float64(ts.Streak)
	}
	typRows.Close()
	for _, k := range []string{"TEORIA", "IMPLEMENTACJA", "ARKUSZ", "SQL"} {
		if ks, ok := katAgg[k]; ok {
			if ks.TypyTotal > 0 {
				ks.AvgStreak /= float64(ks.TypyTotal)
			}
			statusOut.PerKategoria = append(statusOut.PerKategoria, *ks)
		}
	}
}
out.Rekomendacja = computeRekomendacja(d, statusOut)
```

**Step 5: Run test to verify it passes**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestDiagnoseHasStatusFields -v`
Expected: PASS

**Step 6: Run all tests**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v`
Expected: All tests pass

**Step 7: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/types.go matura_informatyka_rozszerzona/analiza/cli/commands.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "feat(cli): add zaleglosci/leech_tagi/retencja/rekomendacja to progress diagnose"
```

---

### Task 3: Remove `--weight-reset` flag from `exerciseNextCmd`

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/commands.go:1817,1837-1839,1933`
- Modify: `matura_informatyka_rozszerzona/analiza/cli/main_test.go`

**Step 1: Update existing test**

In `main_test.go`, rename `TestExerciseNextWeight` and remove the weight-reset parts. Keep only the `addWeight` logic tests (the function still exists):

```go
func TestExerciseNextWeight(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	readWeight := func() int {
		var wStr string
		var w int
		if err := db.QueryRow(`SELECT value FROM progress_meta WHERE key = 'session_context_weight'`).Scan(&wStr); err == nil {
			fmt.Sscanf(wStr, "%d", &w)
		}
		return w
	}

	// 1. Start from 0
	db.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '0')`)
	if w := readWeight(); w != 0 {
		t.Errorf("initial: got weight %d, want 0", w)
	}

	// 2. addWeight(4) twice → weight=8
	addWeight(db, 4)
	addWeight(db, 4)
	if w := readWeight(); w != 8 {
		t.Errorf("after 2x addWeight(4): got weight %d, want 8", w)
	}

	// 3. addWeight with 0 should be no-op
	addWeight(db, 0)
	if w := readWeight(); w != 8 {
		t.Errorf("after addWeight(0): got weight %d, want 8", w)
	}
}
```

**Step 2: Remove flag from `exerciseNextCmd`**

In `commands.go`:

1. Remove line 1817: `var weightReset bool` → remove the variable
2. Remove lines 1837-1839: the `if weightReset { ... }` block
3. Remove line 1933: `cmd.Flags().BoolVar(&weightReset, "weight-reset", false, "Reset session context weight to 0")`

**Step 3: Run tests**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v`
Expected: All tests pass (no test references `--weight-reset` flag anymore)

**Step 4: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/commands.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "refactor(cli): remove --weight-reset flag from exercise next"
```

---

### Task 4: Update `test_qa.sh` smoke tests

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/test_qa.sh:269-283`

**Step 1: Remove `--weight-reset` from smoke test**

Replace lines 269-271:
```bash
  echo "  -- Context weight tracking --"
  test_json_cmd "exercise next --weight-reset" \
    matura_tmp exercise next --typ sql_group_by --weight-reset
```

With:
```bash
  echo "  -- Context weight tracking --"
  test_json_cmd "exercise next (weight tracking)" \
    matura_tmp exercise next --typ sql_group_by
```

**Step 2: Add `progress status` weight-reset smoke test**

After the existing weight tracking section (after line ~283), add:
```bash
  # progress status resets weight
  sqlite3 "$TMPDIR_QA/matura_progress.db" \
    "INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '99')"
  matura_tmp progress status > /dev/null 2>&1
  local weight_after_status
  weight_after_status=$(sqlite3 "$TMPDIR_QA/matura_progress.db" \
    "SELECT value FROM progress_meta WHERE key = 'session_context_weight'")
  if [ "$weight_after_status" = "0" ]; then
    pass "progress status resets weight to 0"
  else
    fail "progress status should reset weight (expected 0, got: $weight_after_status)"
  fi
```

**Step 3: Run L1 smoke tests**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh --layer 1`
Expected: All smoke tests pass

**Step 4: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/test_qa.sh
git commit -m "test: update test_qa.sh for auto weight reset via progress status"
```

---

### Task 5: Build binaries and re-import

**Step 1: Run build.sh**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && ./build.sh`
Expected: Builds macOS + Windows binaries, re-imports matura.db

**Step 2: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh`
Expected: All 6 layers pass (112 tests, 0 failures)

**Step 3: Commit binaries**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/matura matura_informatyka_rozszerzona/analiza/cli/matura.exe matura_informatyka_rozszerzona/analiza/cli/matura.db
git commit -m "build: rebuild binaries after weight-reset removal"
```

---

### Task 6: Update SKILL.md — remove `--weight-reset` mentions

**Files:**
- Modify: `.claude/skills/matura/SKILL.md`

**Step 1: Remove `--weight-reset` from Section C1 (line 158)**

Replace:
```
**[WYMAGANE]** Pierwsze `exercise review` w sesji MUSI miec `--weight-reset` (tak samo jak `exercise next`).
```
With: (delete entirely — the line is removed)

**Step 2: Remove `--weight-reset` from Section C1 (line 170)**

Replace:
```
**[WYMAGANE]** Pierwsze `exercise next` w sesji MUSI miec `--weight-reset`.
```
With: (delete entirely)

**Step 3: Remove `--weight-reset` from Section C2 (line 183)**

Replace:
```
**[WYMAGANE]** Pierwsze `exercise next` w sesji MUSI miec `--weight-reset`.
```
With: (delete entirely)

**Step 4: Remove `--weight-reset` from Section C3 (line 194)**

Replace:
```
**[WYMAGANE]** Pierwsze `exercise next` w sesji MUSI miec `--weight-reset`.
```
With: (delete entirely)

**Step 5: Remove `--weight-reset` from Section D (lines 200-201)**

Replace:
```
# Pierwsze wywolanie w sesji (zeruje wage kontekstu):
./matura exercise next --typ {typ} --weight-reset

# Kazde kolejne:
./matura exercise next --typ {typ}
```
With:
```
./matura exercise next --typ {typ}
```

**Step 6: Remove `progress status` from proactive detection (line 447)**

In Section F (proactive detection, lines 444-455), remove line 447:
```
./matura progress status
```

Keep only `./matura progress diagnose --typ {aktualny_typ} --limit 1`.

Update the condition checks that follow — they now reference `diagnose` output fields:

Replace lines 450-455:
```
Jesli `top_bledy[0].count >= 3`:
  "Zauwazam powtarzajacy sie blad: {blad_kod}. Chcesz, zebym wyjasnil to zagadnienie?"

Jesli `rekomendacja` niepuste → wyswietl: "Dashboard: {rekomendacja}"
Jesli `retencja_szacowana < 0.80` → "Uwaga: ogolna retencja {retencja_szacowana*100:.0f}% — rozważ powtorki"
Uzyj `rekomendacja` do zasugerowania nastepnego typu (zamiast kontynuacji biezacego).
```

No change needed in the text — the field names (`rekomendacja`, `retencja_szacowana`) are now in `diagnose` output.

**Step 7: Remap student `status` command (line 471)**

Replace:
```
| `status` | `./matura progress status` |
```
With:
```
| `status` | `./matura progress diagnose` — dashboard z rekomendacja, retencja, zaleglosci |
```

**Step 8: Update Section I (lines 550-554)**

Replace:
```
## I. Reset kontekstu

CLI automatycznie liczy wage kontekstu. Kazda komenda dodaje swoja wage do sesji.

**Na starcie sesji** (po `progress status`): `--weight-reset` przy pierwszym `exercise next`.
```
With:
```
## I. Reset kontekstu

CLI automatycznie liczy wage kontekstu. Kazda komenda dodaje swoja wage do sesji.
`progress status` automatycznie resetuje wage (wywolywany na starcie sesji — sekcja C).
```

**Step 9: Run SKILL lint**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh --layer 3`
Expected: SKILL lint passes (no mentions of `--weight-reset` remain)

**Step 10: Commit**

```bash
git add .claude/skills/matura/SKILL.md
git commit -m "refactor(SKILL.md): remove --weight-reset, auto-reset via progress status"
```

---

### Task 7: Update probna.md

**Files:**
- Modify: `.claude/skills/matura/probna.md:14`

**Step 1: Update the weight-reset note**

Line 14 currently says:
```
0. **[WYMAGANE]** Sprawdz status: `./matura progress status` (SKILL.md C wymaga ZAWSZE)
```

This is already correct — `progress status` is called and it auto-resets weight. No change needed.

**Step 2: Verify no other `--weight-reset` mentions exist in skill files**

Run: `grep -r "weight-reset" .claude/skills/`
Expected: No matches

**Step 3: Commit (skip if no changes)**

---

### Task 8: Final verification

**Step 1: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh`
Expected: All layers pass

**Step 2: Run Go unit tests individually**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v -run "TestProgressStatusResetsWeight|TestDiagnoseHasStatusFields|TestExerciseNextWeight"`
Expected: 3 tests PASS

**Step 3: Verify no `--weight-reset` remains in codebase**

Run: `grep -rn "weight-reset\|weight_reset" matura_informatyka_rozszerzona/analiza/cli/ .claude/skills/matura/ --include="*.go" --include="*.md" --include="*.sh"`
Expected: Only `session_context_weight` references remain (in `addWeight`, `getSessionState`, progress_meta queries). No `--weight-reset` flag references.

**Step 4: Manual smoke test**

```bash
cd matura_informatyka_rozszerzona/analiza/cli

# Set weight high
sqlite3 matura_progress.db "INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '99')"

# Run progress status — should reset to 0
./matura progress status | python3 -c "import sys,json; print('OK')"

# Verify weight is 0
sqlite3 matura_progress.db "SELECT value FROM progress_meta WHERE key = 'session_context_weight'"
# Expected: 0

# Run progress diagnose — should have new fields
./matura progress diagnose | python3 -c "import sys,json; d=json.loads(sys.stdin.read()); print('zaleglosci:', d['zaleglosci'], 'leech_tagi:', d['leech_tagi'])"
# Expected: zaleglosci: 0 leech_tagi: []

# Verify exercise next no longer accepts --weight-reset
./matura exercise next --typ sql_group_by --weight-reset 2>&1 | head -1
# Expected: error about unknown flag
```
