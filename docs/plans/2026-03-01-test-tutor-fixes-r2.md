# Test-Tutor Findings R2: CLI + SKILL.md Fixes

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Fix 5 issues found by test-tutor run on commit 98ef87d (score 83.1/100, 2 FAILs) — arrow normalization, multi-part answer scoring, MENTION_PAST query bug, SKILL.md coaching ordering, hint GATE flexibility.

**Architecture:** Three layers of fixes: (1) `normalize.go` — expand answer comparison to handle Unicode arrows and multi-part `a)/b)/c)` answers; (2) `commands.go` — fix SQL query in `buildCoaching` so MENTION_PAST actually triggers; (3) `SKILL.md` — clarify coaching_actions ordering and walk-through GATE compatibility with hint progression.

**Tech Stack:** Go (normalize.go, commands.go, main_test.go), Markdown (SKILL.md)

---

### Task 1: Arrow Normalization in `normalizeAnswer`

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/normalize.go:23-39`
- Test: `matura_informatyka_rozszerzona/analiza/cli/main_test.go:3660-3678`

**Step 1: Write the failing test**

Add arrow normalization cases to `TestNormalizeAnswer` at `main_test.go:3671` (after the `"0013"` case):

```go
		{"2 → 5 → 8", "2 -> 5 -> 8"},
		{"A ← B", "a <- b"},
		{"1⟶2⟶3", "1->2->3"},
		{"tak → nie → tak", "tak -> nie -> tak"},
```

**Step 2: Run test to verify it fails**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestNormalizeAnswer -v`
Expected: FAIL — "2 → 5 → 8" normalizes to "2 → 5 → 8" (arrows kept), not "2 -> 5 -> 8"

**Step 3: Write minimal implementation**

In `normalize.go`, add arrow replacement BEFORE the numeric normalization (after line 28, the `ł` replacement):

```go
	// Normalize Unicode arrows to ASCII
	s = strings.ReplaceAll(s, "→", "->")
	s = strings.ReplaceAll(s, "←", "<-")
	s = strings.ReplaceAll(s, "⟶", "->")
```

**Step 4: Run test to verify it passes**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestNormalizeAnswer -v`
Expected: PASS

**Step 5: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/normalize.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "fix: normalize Unicode arrows in check-answer (→ to ->)"
```

---

### Task 2: Multi-Part Answer Support

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/normalize.go` (add `checkMultiPartAnswer` function)
- Modify: `matura_informatyka_rozszerzona/analiza/cli/commands.go` (call `checkMultiPartAnswer` from check-answer command handler)
- Test: `matura_informatyka_rozszerzona/analiza/cli/main_test.go`

**Step 1: Write the failing tests**

Add after `TestCheckAnswerNonAutoScorable` (after line 3755):

```go
func TestCheckMultiPartAnswerAllCorrect(t *testing.T) {
	result := checkMultiPartAnswer("a) 5  b) 13  c) PRAWDA", "a) 5 b) 13 c) prawda")
	if !result.Poprawne {
		t.Error("all parts correct should be poprawne")
	}
	if result.Wynik != "pelne" {
		t.Errorf("expected wynik=pelne, got %q", result.Wynik)
	}
}

func TestCheckMultiPartAnswerPartial(t *testing.T) {
	result := checkMultiPartAnswer("a) 5  b) 13  c) PRAWDA", "a) 5 b) 99 c) prawda")
	if result.Poprawne {
		t.Error("partial answer should not be poprawne")
	}
	if result.Wynik != "czesciowe" {
		t.Errorf("expected wynik=czesciowe, got %q", result.Wynik)
	}
	if result.TrafioneParts != 2 || result.TotalParts != 3 {
		t.Errorf("expected 2/3 parts, got %d/%d", result.TrafioneParts, result.TotalParts)
	}
}

func TestCheckMultiPartAnswerAllWrong(t *testing.T) {
	result := checkMultiPartAnswer("a) 5  b) 13  c) PRAWDA", "a) 1 b) 2 c) falsz")
	if result.Poprawne {
		t.Error("all wrong should not be poprawne")
	}
	if result.Wynik != "zero" {
		t.Errorf("expected wynik=zero, got %q", result.Wynik)
	}
}

func TestCheckMultiPartAnswerSinglePart(t *testing.T) {
	// Single-part answers should fall through to regular checkAnswer
	result := checkMultiPartAnswer("42", "42")
	if !result.Poprawne {
		t.Error("single-part correct should be poprawne")
	}
}

func TestCheckMultiPartAnswerNewlineSeparated(t *testing.T) {
	result := checkMultiPartAnswer("a) 5\nb) 13\nc) PRAWDA", "a) 5\nb) 13\nc) prawda")
	if !result.Poprawne {
		t.Error("newline-separated all correct should be poprawne")
	}
}
```

**Step 2: Run tests to verify they fail**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestCheckMultiPart -v`
Expected: FAIL — `checkMultiPartAnswer` undefined

**Step 3: Write the implementation**

Add to `normalize.go` after `checkAnswer` (after line 61):

```go
// multiPartResult extends checkAnswerResult with part counts.
type multiPartResult struct {
	checkAnswerResult
	TrafioneParts int
	TotalParts    int
}

// splitMultiPart splits an answer into parts by "a)", "b)", "c)", "d)" labels.
// Returns nil if the answer doesn't have multi-part format.
func splitMultiPart(s string) []string {
	// Match patterns like "a)" "b)" "c)" "d)" at start or after whitespace/newline
	re := regexp.MustCompile(`(?i)\b([a-d])\)\s*`)
	locs := re.FindAllStringIndex(s, -1)
	if len(locs) < 2 {
		return nil // not multi-part
	}

	var parts []string
	for i, loc := range locs {
		start := loc[1] // after the "X) " match
		var end int
		if i+1 < len(locs) {
			end = locs[i+1][0]
		} else {
			end = len(s)
		}
		part := strings.TrimSpace(s[start:end])
		parts = append(parts, part)
	}
	return parts
}

// checkMultiPartAnswer handles multi-part answers (a/b/c format).
// Falls back to regular checkAnswer for single-part answers.
func checkMultiPartAnswer(correct, student string) multiPartResult {
	correctParts := splitMultiPart(correct)
	studentParts := splitMultiPart(student)

	// If either is not multi-part, fall back to simple comparison
	if correctParts == nil || studentParts == nil {
		r := checkAnswer(correct, student)
		total := 1
		hit := 0
		if r.Poprawne {
			hit = 1
		}
		return multiPartResult{checkAnswerResult: r, TrafioneParts: hit, TotalParts: total}
	}

	// Compare part by part (use minimum of both lengths)
	total := len(correctParts)
	hits := 0
	for i := 0; i < total && i < len(studentParts); i++ {
		r := checkAnswer(correctParts[i], studentParts[i])
		if r.Poprawne {
			hits++
		}
	}

	var wynik string
	switch {
	case hits == total:
		wynik = "pelne"
	case hits == 0:
		wynik = "zero"
	default:
		wynik = "czesciowe"
	}

	return multiPartResult{
		checkAnswerResult: checkAnswerResult{
			Poprawne:          hits == total,
			Wynik:             wynik,
			PoprawnaOdpowiedz: correct,
		},
		TrafioneParts: hits,
		TotalParts:    total,
	}
}
```

Also add `"regexp"` to the imports at the top of `normalize.go`.

**Step 4: Run tests to verify they pass**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestCheckMultiPart -v`
Expected: all 5 PASS

**Step 5: Wire up in commands.go**

Find the check-answer command handler in `commands.go`. Currently it calls `checkAnswer(correct, student)`. Replace with `checkMultiPartAnswer(correct, student)` and extract `TrafioneParts`/`TotalParts` into the JSON output.

Search for `checkAnswer(` in commands.go (should be in the check-answer cobra command handler) and replace the call. Add `trafione_parts` and `total_parts` to the JSON output for partial scores.

**Step 6: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v -count=1`
Expected: all tests PASS (117+ existing + 5 new)

**Step 7: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/normalize.go matura_informatyka_rozszerzona/analiza/cli/commands.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "feat: multi-part answer scoring (a/b/c splitting with partial credit)"
```

---

### Task 3: Fix MENTION_PAST Query Bug

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/commands.go:208-228`
- Test: `matura_informatyka_rozszerzona/analiza/cli/main_test.go:2760-2777`

**Context:** The `buildCoaching` function at `commands.go:216-219` queries past mistakes by matching `blad_kod IN (exerciseTags)`. But `blad_kod` values are error codes (e.g., `mylenie_div_mod`) while `exerciseTags` are skill tags (e.g., `cyfry-mod-div`). These are different namespaces and almost never match. The existing test at line 2766 uses `blad_kod = 'cyfry-mod-div'` which happens to also be a tag — masking the bug.

**Step 1: Write the failing test**

Add after `TestBuildCoachingPastMistakes` (after line 2777):

```go
func TestBuildCoachingPastMistakesRealErrorCodes(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Insert error with a REAL error code (not a tag name)
	db.Exec("INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.1', 'cyfry_liczby', '2026-02-20', 'poprawne_z_pomoca_1')")
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, data)
		VALUES ('7.1', 'cyfry_liczby', 'mylenie_div_mod', 'Pomylenie div z mod', '2026-02-20')`)

	// Exercise tags are skill tags, NOT error codes
	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div", "cyfry-dzielniki"})
	if len(coaching.PastMistakes) != 1 {
		t.Errorf("expected 1 past_mistake with real error code, got %d: %v", len(coaching.PastMistakes), coaching.PastMistakes)
	}
	if len(coaching.PastMistakes) > 0 && coaching.PastMistakes[0] != "Pomylenie div z mod" {
		t.Errorf("expected 'Pomylenie div z mod', got %q", coaching.PastMistakes[0])
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestBuildCoachingPastMistakesRealErrorCodes -v`
Expected: FAIL — `expected 1 past_mistake with real error code, got 0: []`

**Step 3: Fix the query**

In `commands.go`, replace lines 208-219:

**Old (broken):**
```go
	// Past mistakes: from last 5 sessions, filtered by exercise tags
	if len(exerciseTags) > 0 {
		placeholders := make([]string, len(exerciseTags))
		params := []any{typ}
		for i, tag := range exerciseTags {
			placeholders[i] = "?"
			params = append(params, tag)
		}
		query := fmt.Sprintf(`SELECT blad_opis FROM progress_bledy
			WHERE typ = ? AND blad_kod IN (%s)
			AND data IN (SELECT DISTINCT data FROM progress_zrobione ORDER BY data DESC LIMIT 5)
			GROUP BY blad_opis ORDER BY MAX(data) DESC LIMIT 3`, strings.Join(placeholders, ","))
		rows, err := d.Query(query, params...)
```

**New (fixed):**
```go
	// Past mistakes: from last 5 sessions, filtered by type
	{
		rows, err := d.Query(`SELECT blad_opis FROM progress_bledy
			WHERE typ = ?
			AND data IN (SELECT DISTINCT data FROM progress_zrobione ORDER BY data DESC LIMIT 5)
			GROUP BY blad_opis ORDER BY MAX(data) DESC LIMIT 3`, typ)
```

Note: the closing `if err == nil { ... }` block stays the same. The key change:
1. Remove `exerciseTags` guard (`if len(exerciseTags) > 0`)
2. Remove `blad_kod IN (...)` filter
3. Keep `typ = ?` filter (match errors for same exercise type)
4. Keep the `data IN (...)` recency filter

**Step 4: Run test to verify it passes**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestBuildCoachingPastMistakes -v`
Expected: both `TestBuildCoachingPastMistakes` and `TestBuildCoachingPastMistakesRealErrorCodes` PASS

**Step 5: Update old test for correctness**

The old `TestBuildCoachingPastMistakes` at line 2766 inserts `blad_kod = 'cyfry-mod-div'` and `blad_kod = 'unrelated-tag'`. With the fixed query (no tag filtering), BOTH errors will now match. Update the test to expect 2 past mistakes, or change the second insert to a different `typ` to keep testing type isolation:

Change line 2768 from:
```go
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, data)
		VALUES ('7.1', 'cyfry_liczby', 'unrelated-tag', 'nieistotny blad', '2026-02-20')`)
```
to:
```go
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, data)
		VALUES ('8.1', 'sql_group_by', 'unrelated-tag', 'nieistotny blad', '2026-02-20')`)
```

This tests that the type filter correctly excludes errors from other types.

**Step 6: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v -count=1`
Expected: all tests PASS

**Step 7: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/commands.go matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "fix: MENTION_PAST query now matches by type, not by tag/error-code join"
```

---

### Task 4: SKILL.md — coaching_actions_v2 Ordering Rule

**Files:**
- Modify: `.claude/skills/matura/SKILL.md:254-269` (section E2)

**Step 1: Read current E2 section**

Current text at SKILL.md:262:
```
**Przeczytaj `coaching_actions_v2` (preferowane) lub `coaching_actions` (legacy) i wlacz naturalnie w dialog PRZED podaniem tresci cwiczenia.**
```

**Step 2: Replace with explicit ordering**

Replace lines 262-269 with:

```markdown
**Realizuj `coaching_actions_v2` (preferowane) w tej kolejnosci:**

1. **PRZED trescia cwiczenia** (priorytet: wysoki → niski):
   - `WARN_LEECH` → powiedz uczniowi o leech tagu ("Ten temat sprawia Ci trudnosc...")
   - `HINT_DELAY` → poinformuj ("Od teraz czekam {N} prob zanim dam podpowiedz")
2. **PO pierwszym bledzie ucznia**:
   - `MENTION_PAST` → nawiaz do historii bledow ("Ostatnio miales problem z...")

Jesli `coaching_actions_v2` puste → pomin, przejdz do tresci.

**Jesli `difficulty_bumped=true`** w odpowiedzi `exercise next`:
- PRZED trescia cwiczenia powiedz: "Wlasnie awansowales na poziom {nowy}! Gotowy na wyzszy poziom?"
- Czekaj na reakcje ucznia, krotki feedback, potem tresc cwiczenia.

`coaching_actions_v2` zwraca gotowe zdania — mozesz je parafrazowac, ale zachowaj kluczowy przekaz:
- `typ: "WARN_LEECH"` (priorytet: wysoki) → tekst o leech tagu, MUSI byc wlaczony
- `typ: "MENTION_PAST"` (priorytet: niski) → tekst o poprzednich bledach
- `typ: "HINT_DELAY"` (priorytet: niski) → tekst o zmniejszonej liczbie podpowiedzi
```

**Step 3: Verify no other references broken**

Search SKILL.md for "coaching_actions" — should only appear in E2 and in guardrails section. No other changes needed.

**Step 4: Commit**

```bash
git add .claude/skills/matura/SKILL.md
git commit -m "docs: clarify coaching_actions_v2 ordering (WARN_LEECH before, MENTION_PAST after error)"
```

---

### Task 5: SKILL.md — Walk-Through GATE Compatibility with Hint Progression

**Files:**
- Modify: `.claude/skills/matura/SKILL.md:362` (section F, step 3 GATE)
- Modify: `.claude/skills/matura/SKILL.md:384` (section F, step 5)

**Step 1: Read current GATE text**

Current at line 362:
```
   - **[GATE]** Jesli to 3. bledna proba LUB uczen mowi "poddaje sie" → POMIN krok 4, przejdz BEZPOSREDNIO do kroku 5 (walk_through). NIE probuj kolejnego hintu.
```

**Step 2: Replace with clarified GATE**

Replace line 362 with:

```markdown
   - **[GATE]** Jesli to 3. bledna **proba ucznia** (kazda odpowiedz ucznia po hincie = osobna proba;
     pytanie sokratejskie bez odpowiedzi NIE liczy sie jako proba) LUB uczen mowi "poddaje sie"
     → POMIN krok 4, przejdz BEZPOSREDNIO do kroku 5 (walk_through). NIE probuj kolejnego hintu.
     Progresja: proba 1 (bez hinta) → proba 2 (po L1) → proba 3 (po L2 + cheatsheet) → walk_through (z L3).
```

**Step 3: Verify step 5 is compatible**

Current step 5 at line 384: `**5. Po 3 probach / "poddaje sie"**` — this is consistent. No change needed.

**Step 4: Commit**

```bash
git add .claude/skills/matura/SKILL.md
git commit -m "docs: clarify walk-through GATE counts student attempts only (L1→L2→L3 compatible)"
```

---

### Task 6: Build + Full Test Suite

**Files:**
- Build: `matura_informatyka_rozszerzona/analiza/cli/build.sh`

**Step 1: Build binaries + reimport**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && ./build.sh`
Expected: macOS + Windows binaries built, matura.db reimported

**Step 2: Run Go unit tests**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v -count=1`
Expected: all tests PASS (122+ tests)

**Step 3: Run test_qa.sh (if available)**

Run: `matura_informatyka_rozszerzona/analiza/test_qa.sh`
Expected: all layers PASS

**Step 4: Final commit with built binaries**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/matura matura_informatyka_rozszerzona/analiza/cli/matura.exe
git commit -m "build: rebuild binaries with arrow norm + multi-part + MENTION_PAST fix"
```
