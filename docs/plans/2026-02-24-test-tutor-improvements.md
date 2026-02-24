# Test-Tutor Improvements Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Fix 3 issues from test-tutor analysis to raise pedagogical score from 92.3 to ~95-97/100.

**Architecture:** Three independent changes applied as separate commits with targeted verification: (1) CLI interleave streak guard in Go, (2) test-tutor cke_unlock scenario fix, (3) Socratic hint delivery in SKILL.md.

**Tech Stack:** Go (CLI), Markdown (SKILL.md files), Bash (test_qa.sh, build.sh)

**Design doc:** `docs/plans/2026-02-24-test-tutor-improvements-design.md`

---

### Task 1: Write failing test for interleave streak guard

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/main_test.go`

**Step 1: Write the failing test**

Add this test after the existing `TestExerciseNextWeight` function (around line 1603):

```go
func TestInterleaveSkippedAtStreakThreshold(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Set up a student with streak=3 on cyfry_liczby (at difficulty climb threshold)
	db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'latwe', 3)")

	// sessionCount=3 would normally trigger interleave (3%3==0)
	// But streak=3 means student is at difficulty climb threshold — skip interleave
	streak := getStreak(db, "cyfry_liczby")
	if streak != 3 {
		t.Fatalf("setup: streak=%d, want 3", streak)
	}

	// Verify: interleave should be skipped when streak >= 3
	// The interleave condition should be: sessionCount%3==0 && streak < 3
	shouldInterleave := (3 > 0 && 3%3 == 0 && streak < 3)
	if shouldInterleave {
		t.Error("interleave should be skipped at streak=3 (difficulty climb threshold)")
	}

	// Verify: interleave is allowed when streak < 3
	db.Exec("UPDATE progress_typy SET streak = 2 WHERE typ = 'cyfry_liczby'")
	streak = getStreak(db, "cyfry_liczby")
	shouldInterleave = (3 > 0 && 3%3 == 0 && streak < 3)
	if !shouldInterleave {
		t.Error("interleave should be allowed when streak < 3")
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestInterleaveSkippedAtStreakThreshold -v`
Expected: FAIL — `getStreak` function does not exist yet.

---

### Task 2: Implement interleave streak guard

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/commands.go:312` (add `getStreak` helper near `getLevel`)
- Modify: `matura_informatyka_rozszerzona/analiza/cli/commands.go:2099-2100` (add guard)

**Step 1: Add `getStreak` helper**

Add this function right after the existing `getLevel` function (after line 319 in commands.go):

```go
func getStreak(d *sql.DB, typ string) int {
	var streak int
	d.QueryRow("SELECT COALESCE(streak, 0) FROM progress_typy WHERE typ = ?", typ).Scan(&streak)
	return streak
}
```

**Step 2: Add streak guard to interleave condition**

In `commands.go`, change line 2100 from:

```go
if sessionCount > 0 && sessionCount%3 == 0 {
```

to:

```go
if sessionCount > 0 && sessionCount%3 == 0 && getStreak(d, typ) < 3 {
```

**Step 3: Run the new test**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestInterleaveSkippedAtStreakThreshold -v`
Expected: PASS

**Step 4: Run all existing tests**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v`
Expected: All tests PASS (no regressions)

**Step 5: Build binaries + reimport**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && ./build.sh`
Expected: macOS + Windows binaries built, matura.db reimported

**Step 6: Run test_qa.sh**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh`
Expected: All 6 layers PASS

---

### Task 3: Fix cke_unlock scenario in test-tutor SKILL.md

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md` (section 5.5, lines 275-301)
- Modify: `.claude/skills/test-tutor/SKILL.md` (section 2, pre-fetch — add cke_unlock setup)

**Step 1: Update pre-fetch section**

In section 2 (pre-fetch danych), after the coaching_aware setup block (around line 84), add a cke_unlock setup block:

```bash
# cke_unlock: zasymuluj studenta tuz przed progiem trudne (streak=7, srednie-trudne)
sqlite3 /tmp/test-tutor-$$/matura_progress.db "
INSERT OR REPLACE INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('sledzenie_algorytmu', 'srednie-trudne', 7);
"
EX_CKE_PRE=$($MATURA --db-dir /tmp/test-tutor-$$ exercise question --typ sledzenie_algorytmu --trudnosc srednie-trudne)
EX_CKE_PRE_ID=$(echo "$EX_CKE_PRE" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])")
ANSWER_CKE_PRE=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $EX_CKE_PRE_ID)
```

**Step 2: Update section 5.5 fixed student script**

Replace the current section 5.5 fixed student script (lines 277-282) with:

```
wymiana_0_uczen: "[poprawna odpowiedz na cwiczenie srednie-trudne — streak rośnie do 8]"
wymiana_1_uczen: "sprawdzian sledzenie_algorytmu"
wymiana_2_uczen: "[poprawna odpowiedz na worked-example pytanie o pulapki]"
wymiana_3_uczen: "[czesciowo poprawna odpowiedz na sprawdzianie — 70% punktow]"
```

**Step 3: Update section 5.5 binary checkpoints**

Replace the current checkpoints (lines 284-301) with:

```
**Binary checkpoints:**
```
CLI compliance:
[ ] exercise next --typ sledzenie_algorytmu (dla wymiana_0)
[ ] progress update --wynik poprawne_bez_pomocy (po wymiana_0, streak→8)
[ ] cke worked-example --typ X PRZED sprawdzianem
[ ] cke get --typ X --exclude (wyklucz wczesniej robione)
[ ] START_TS i ELAPSED
[ ] cke save --id X --punkty N --max M

Coaching reaction:
[ ] coaching_actions zrealizowane (jesli obecne)

Scenario-specific:
[ ] Ogloszenie odblokowania w formacie "*** ODBLOKOWANO ***" (po progress update gdy streak=8→trudne)
[ ] Pytanie o pulapki po worked-example ("Co zapamiętasz z tych pułapek?")
[ ] Brak hintow na sprawdzianie ("To sprawdzian — na egzaminie tez nie bedzie hintow")
[ ] Ocena czesciowa wg zasady_oceniania
[ ] Ogloszenie formatu "=== SPRAWDZIAN TYPU ==="
```
```

**Step 4: Update agent prompt section**

In section 6 (orchestracja agentow), update the pre-fetched data block for the cke_unlock scenario to include:
```
- Question (CKE_PRE): {EX_CKE_PRE} (only for cke_unlock scenario — exercise before unlock)
- Answer (CKE_PRE): {ANSWER_CKE_PRE}
```

**Step 5: Verify SKILL.md lint passes**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh --layer 3`
Expected: Layer 3 (SKILL lint) PASS

---

### Task 4: Add Socratic diagnostic questions to hint delivery

**Files:**
- Modify: `.claude/skills/matura/SKILL.md` (section F, point 4, lines 311-318)

**Step 1: Edit hint delivery section**

In `.claude/skills/matura/SKILL.md`, replace lines 311-318 (the hint levels inside point 4):

FROM:
```
   - Jesli hinty dostepne → podaj nastepna wskazowke z `wskazowki[]`:
     * **Poziom 1**: `wskazowki[0]` + pytanie sokratejskie
     * **Poziom 2**: `wskazowki[1]` + cytat z cheatsheet:
       `./matura cheatsheet get --kategoria {kat} --sekcja "{temat}"`
       Mapowanie: mod/div→"archetyp", rekurencja→"rekurencj", zlozonosc→"zlozonosc",
       JOIN→"join", GROUP BY→"group", sortowanie→"sort", adresowanie→"adresow",
       szyfrowanie→"bezpieczen", P/F→"prawda", konwersja→"konwersj"
     * **Poziom 3**: `wskazowki[2]` (kluczowy krok) + rozpisz krok po kroku, ostatni krok zostaw uczniowi
```

TO:
```
   - Jesli hinty dostepne → podaj nastepna wskazowke z `wskazowki[]`:
     * **Poziom 1**: NAJPIERW zapytaj: "Gdzie wedlug Ciebie jest blad?" (czekaj na odpowiedz).
       POTEM: `wskazowki[0]` + pytanie sokratejskie
     * **Poziom 2**: NAJPIERW zapytaj: "Co juz wiesz o [temat hintu]?" (czekaj na odpowiedz).
       POTEM: `wskazowki[1]` + cytat z cheatsheet:
       `./matura cheatsheet get --kategoria {kat} --sekcja "{temat}"`
       Mapowanie: mod/div→"archetyp", rekurencja→"rekurencj", zlozonosc→"zlozonosc",
       JOIN→"join", GROUP BY→"group", sortowanie→"sort", adresowanie→"adresow",
       szyfrowanie→"bezpieczen", P/F→"prawda", konwersja→"konwersj"
     * **Poziom 3**: `wskazowki[2]` (kluczowy krok) + rozpisz krok po kroku, ostatni krok zostaw uczniowi
```

**Step 2: Verify SKILL.md lint passes**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh --layer 3`
Expected: Layer 3 (SKILL lint) PASS

---

### Task 5: Full verification

**Step 1: Run full test_qa.sh**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh`
Expected: All 6 layers PASS

**Step 2: Verify build is clean**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go vet ./...`
Expected: No issues

---

## Summary of all changes

| Task | File | Lines Changed | Nature |
|------|------|--------------|--------|
| 1-2 | `cli/commands.go` | +7 (getStreak func) +1 (guard condition) | Go code |
| 1-2 | `cli/main_test.go` | +30 (new test) | Go test |
| 3 | `.claude/skills/test-tutor/SKILL.md` | ~25 lines in section 2 + 5.5 | Markdown |
| 4 | `.claude/skills/matura/SKILL.md` | 2 lines added in section F point 4 | Markdown |
