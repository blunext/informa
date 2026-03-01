# Design: Test-Tutor Findings R2 — CLI + SKILL.md Fixes

**Date**: 2026-03-01
**Trigger**: test-tutor run on commit 98ef87d scored 83.1/100 (2 FAILs: hint_progression 67.7, difficulty_climb 59.5)

## Problem Space

| # | Problem | Domain | Root Cause |
|---|---------|--------|------------|
| 1 | `check-answer` false negatives on arrows | CLI bug | `normalizeAnswer()` doesn't normalize Unicode → to -> |
| 2 | `check-answer` can't handle multi-part answers (a/b/c) | CLI limitation | Whole-string comparison, no part splitting |
| 3 | MENTION_PAST never triggers | CLI bug | Query joins `blad_kod` (error codes) with `exerciseTags` (skill tags) — different namespaces |
| 4 | coaching_actions_v2 ordering unclear | SKILL.md gap | No rule for WHEN actions happen relative to exercise content |
| 5 | 3-attempt GATE conflicts with L1→L2→L3 hint progression | SKILL.md design tension | GATE fires before L2/L3 can be delivered |

## Fixes

### Fix 1: Arrow normalization in `normalize.go`

Add to `normalizeAnswer()`:
```go
s = strings.ReplaceAll(s, "→", "->")
s = strings.ReplaceAll(s, "←", "<-")
s = strings.ReplaceAll(s, "⟶", "->")
```

### Fix 2: Multi-part answer support

Add `checkMultiPartAnswer()` that:
1. Detects multi-part format: `a)...b)...c)...` or newline-separated with labels
2. Splits both correct and student answers into parts
3. Normalizes and compares each part
4. Returns proportional score: all correct → "pelne", some → "czesciowe", none → "zero"
5. Falls back to single `checkAnswer()` if no parts detected

### Fix 3: MENTION_PAST query fix in `commands.go`

Current (broken):
```sql
SELECT blad_opis FROM progress_bledy
WHERE typ = ? AND blad_kod IN ({exerciseTags})
```

Fixed (remove tag filter, match by type only):
```sql
SELECT blad_opis FROM progress_bledy
WHERE typ = ?
AND data IN (SELECT DISTINCT data FROM progress_zrobione ORDER BY data DESC LIMIT 5)
GROUP BY blad_opis ORDER BY MAX(data) DESC LIMIT 3
```

### Fix 4: SKILL.md section E2 — coaching_actions_v2 ordering

Add explicit ordering rule after E2:
- WARN_LEECH → say BEFORE exercise content
- HINT_DELAY → inform BEFORE exercise content
- MENTION_PAST → reference AFTER student's first error
- difficulty_bumped=true → ask deepening question BEFORE exercise content

### Fix 5: SKILL.md — Walk-through GATE compatibility with hint progression

Clarify the 3-attempt GATE:
- Count only distinct student wrong attempts (not Socratic questions without response)
- Progression: attempt 1 (no hint) → attempt 2 (L1) → attempt 3 (L2 + cheatsheet) → walk_through (includes L3)
- This makes L1→L2→L3 compatible with 3-attempt GATE

## Testing

- Unit tests for normalize.go: arrow normalization, multi-part answers
- Unit test for buildCoaching: MENTION_PAST triggers with matching type errors
- `test_qa.sh` full pass after changes
- test-tutor re-run to verify score improvement (target: hint_progression > 80, difficulty_climb > 85)

## Files Changed

| File | Type | Est. Lines |
|------|------|-----------|
| `analiza/cli/normalize.go` | edit | +40 |
| `analiza/cli/commands.go` | edit | +5 |
| `analiza/cli/main_test.go` | edit | +60 |
| `.claude/skills/matura/SKILL.md` | edit | +15 |
