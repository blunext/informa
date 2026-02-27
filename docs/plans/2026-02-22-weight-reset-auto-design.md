# Design: Auto-reset session_context_weight

## Problem

`--weight-reset` flag on `exercise next`/`exercise review` is forgotten by the tutor in 7+ test-tutor reports. The flag is mentioned in 5 places across SKILL.md (sections C1, C2, C3, D, I), creating redundancy that leads to omission.

## Solution

Make `progress status` the **session-start signal** that automatically resets `session_context_weight`. Remove `--weight-reset` flag entirely.

### Key insight

`progress status` is ALWAYS called at the start of every session (SKILL.md enforces this, test-tutor confirms 100% compliance). It's the natural session boundary.

### Isolation

To prevent mid-session weight resets, `progress status` is removed from:
- Proactive detection (every 5 exercises) — replaced by `progress diagnose` (already called there)
- Student `status` command — remapped to `progress diagnose`

This makes `progress status` a session-start-only command.

## Changes

### CLI (commands.go)

1. **`progressStatusCmd`**: Add weight reset
   ```go
   d.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '0')`)
   ```

2. **`progressDiagnoseCmd`**: Add fields from progress status
   - `rekomendacja`, `retencja_szacowana`, `leech_tagi`, `zaleglosci`
   - Logic already exists in progressStatusCmd — extract and reuse

3. **`exerciseNextCmd`**: Remove `--weight-reset` flag and logic (lines 1837-1839, 1933)

4. **`exerciseReviewCmd`**: Remove `--weight-reset` if present

### SKILL.md

1. **Section C** (3 welcome scenarios): Remove 4x `[WYMAGANE] --weight-reset` markers
2. **Section D** (exercise selection): Remove `--weight-reset` from example command
3. **Section F** (proactive detection, line 444-455): Remove `./matura progress status`, keep only `progress diagnose`
4. **Section H** (student commands): `status` → `./matura progress diagnose` (not `progress status`)
5. **Section I** (context reset): Update description — weight resets automatically via `progress status`

### Tests

1. **Go unit tests**: Update weight-reset test to verify `progress status` resets weight. Add test for `progress diagnose` returning new fields.
2. **test_qa.sh L1**: Remove `--weight-reset` from smoke commands
3. **test-tutor**: No longer checks for `--weight-reset` — checks `progress status` at session start (already 100% compliant)

## Trade-offs

- **Eliminated**: 5 redundant SKILL.md mentions, 1 CLI flag, 7+ report failures
- **Accepted**: `progress status` becomes session-start-only (minor semantic restriction)
- **Net**: Simpler SKILL.md, fewer tokens, zero chance of forgotten weight-reset
