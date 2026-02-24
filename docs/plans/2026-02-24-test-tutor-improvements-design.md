# Test-Tutor Improvements Design

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement the companion plan task-by-task.

**Goal:** Fix 3 issues identified by test-tutor analysis to raise score from 92.3 to ~95-97/100.

**Context:** 23 test-tutor runs over 8 days showed the system plateaued at 85-92. The latest run (92.3/100, commit 6baa479) with the new L1+L2 rubric identified 3 actionable issues with clear fixes.

## Change 1: CLI Interleave Streak Guard

**Problem:** `commands.go:2100` — `sessionCount%3==0` forces interleave mode on every 3rd exercise. This collides with auto-difficulty progression which triggers at streak=3. Interleave has higher priority (Priority 2 vs 3), so the student never gets a harder exercise — interleave "steals" the slot.

**Fix:** Add a streak guard to the interleave condition. Don't interleave when the student's current type has streak >= 3 and hasn't yet received a difficulty upgrade.

```go
// BEFORE:
if sessionCount > 0 && sessionCount%3 == 0 {
    ex, err := findInterleaveExercise(d, typ)
    ...
}

// AFTER:
currentStreak := getCurrentStreak(d, typ)  // new helper
if sessionCount > 0 && sessionCount%3 == 0 && currentStreak < 3 {
    ex, err := findInterleaveExercise(d, typ)
    ...
}
```

`getCurrentStreak(d, typ)` queries `progress_typy` for the current type's streak. If streak >= 3, skip interleave and fall through to Priority 3 (auto-difficulty) which will handle the difficulty upgrade.

**Files:** `matura_informatyka_rozszerzona/analiza/cli/commands.go`

**Verification:** `test-tutor intermediate --scenario difficulty_climb` — checkpoint `nastepne_cwiczenie_po_streak_3_trudnosc_gte_srednie` should PASS.

## Change 2: cke_unlock Scenario Fix

**Problem:** Test-tutor scenario `cke_unlock` has a checkpoint `*** ODBLOKOWANO ***` that expects the tutor to announce unlocking `trudne` difficulty. But the fixed student script starts with the student already at `trudne` level — there's no progression trigger in the scenario.

**Fix:** Modify `test-tutor/SKILL.md` section 5.5:
- Change pre-fetched data: student starts at `srednie-trudne` with `streak=7` (one below threshold)
- Add `wymiana_0`: student answers correctly → streak=8 → CLI triggers difficulty upgrade → tutor announces `*** ODBLOKOWANO ***`
- Shift existing exchanges by 1 (wymiana_1 becomes wymiana_2, etc.)
- Add corresponding binary checkpoint for the progression trigger

**Updated script:**
```
wymiana_0_uczen: "[poprawna odpowiedz — cwiczenie srednie-trudne, streak→8]"
  → tutor announces: *** ODBLOKOWANO typ trudne ***
wymiana_1_uczen: "sprawdzian sledzenie_algorytmu"
wymiana_2_uczen: "[poprawna odpowiedz na worked-example pytanie o pulapki]"
wymiana_3_uczen: "[czesciowo poprawna odpowiedz na sprawdzianie — 70% punktow]"
```

**Files:** `.claude/skills/test-tutor/SKILL.md` (section 5.5 only)

**Verification:** `test-tutor advanced --scenario cke_unlock` — checkpoint `*** ODBLOKOWANO ***` should PASS.

## Change 3: Socratic Hint Delivery

**Problem:** Layer 2 `metoda_sokratejska` = 4/5 across all 7 scenarios consistently. Tutor delivers hints declaratively ("Wskazowka: mod 10 daje ostatnia cyfre") instead of Socratically ("Co daje mod 10?"). The SKILL.md doesn't explicitly require a diagnostic question BEFORE each hint.

**Fix:** Add 1-2 lines per hint level in `matura/SKILL.md` section F, point 4:

```
Jesli hinty dostepne → podaj nastepna wskazowke z wskazowki[]:
  * Poziom 1:
    NAJPIERW zapytaj: "Gdzie wedlug Ciebie jest blad?"
    POTEM: wskazowki[0] + pytanie sokratejskie
  * Poziom 2:
    NAJPIERW zapytaj: "Co juz wiesz o [temat hintu]?"
    POTEM: wskazowki[1] + cytat z cheatsheet
  * Poziom 3: wskazowki[2] + rozpisz krok po kroku,
    ostatni krok zostaw uczniowi
```

This adds a self-diagnosis step before L1 and L2 hints. L3 stays unchanged (student is already stuck, more questions would frustrate).

**Files:** `.claude/skills/matura/SKILL.md` (section F, point 4 only)

**Verification:** `test-tutor` full run — L2 metoda_sokratejska expected to improve from 4.0/5 to 4.5-5.0/5.

## Execution Strategy

Three separate commits, each with targeted verification:

| # | Change | File | Test | Expected Impact |
|---|--------|------|------|-----------------|
| 1 | Interleave streak guard | `commands.go` | `test-tutor intermediate --scenario difficulty_climb` | difficulty_climb 85→89+ |
| 2 | cke_unlock wymiana_0 | `test-tutor/SKILL.md` | `test-tutor advanced --scenario cke_unlock` | cke_unlock 89→92+ |
| 3 | Socratic "NAJPIERW zapytaj" | `matura/SKILL.md` | `test-tutor` full run | L2 sok 4→4.5-5 |

**Combined expected impact:** 92.3 → 95-97/100

## Risks

- **Evaluator variance:** ±5 points per scenario. Single-scenario partial runs reduce noise. Full run after change 3 provides final measurement.
- **Streak guard query cost:** One additional DB query per `exercise next`. Negligible — SQLite query on indexed table.
- **Socratic frustration:** Adding diagnostic questions could slow hint delivery. Mitigated by keeping it at L1/L2 only, not L3.

## Success Criteria

- All 3 changes pass their targeted verification
- Final full test-tutor run: overall >= 95/100
- No regressions in previously passing checkpoints
- test_qa.sh passes (all 6 layers)
