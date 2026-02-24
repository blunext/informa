# Design: Test-Tutor Cumulative Report

**Date:** 2026-02-24
**Status:** Approved

## Problem

Test-tutor generates individual run reports (`YYYY-MM-DD_COMMITHASH.md`) but has no automated cumulative analysis. A manually-curated `ANALIZA_ZBIORCZA.md` existed but was high-maintenance, non-deterministic, and has been removed.

## Solution: Hybrid Go + LLM

Two-layer approach:
- **Go CLI** (deterministic): parses structured data, computes stats, generates tables/trends
- **LLM** (test-tutor agent): interprets numbers, identifies root causes, recommends next actions

Neither does the other's job.

## Architecture

```
test-tutor run
    │
    ├─► generates individual report (.md)    [existing]
    │
    ├─► appends entry to historia.jsonl      [NEW - structured data]
    │
    ├─► calls: matura test-report summary    [NEW - Go CLI command]
    │         reads historia.jsonl
    │         outputs deterministic markdown (tables, trends, stats)
    │
    └─► LLM writes interpretation section    [NEW - appended to cumulative report]
              reads Go output (numbers only)
              writes: why, what to fix, stubborn problems
              saves: RAPORT_ZBIORCZY.md
```

## Component 1: JSONL Schema

**File:** `analiza/test_pedagogical/reports/historia.jsonl`
**Format:** JSON Lines (one line = one run, append-only)

```json
{
  "date": "2026-02-24",
  "commit": "6baa479",
  "mode": "full",
  "overall_score": 92.3,
  "pass": true,
  "scenario_count": 7,
  "scenarios": [
    {
      "persona": "beginner",
      "scenario": "first_session",
      "score": 95.0,
      "l1_percent": 100.0,
      "l1_total": 10,
      "l1_passed": 10,
      "l2": {
        "socratic": 4,
        "tone": 5
      },
      "issues": ["Tutor used exercise question instead of exercise next"]
    }
  ]
}
```

### Schema rules:
- `l1_percent`, `l1_total`, `l1_passed`: nullable (old reports pre-Feb 24 lack L1 data)
- `l2`: flexible map (old format has 7-8 criteria, new has 2)
- `issues`: per-scenario list of problem strings
- Duplicate commits allowed (same commit, different runs) — Go CLI handles averaging

## Component 2: Go CLI Command

**Command:** `matura test-report summary [--window N] [--format md|json]`

Reads `historia.jsonl`, outputs deterministic analysis.

### Output sections:

#### 2a. Dashboard
```
## Dashboard
Score: 92.3/100 (PASS) ↑ +3.0 vs previous
Streak: 21 consecutive PASS
Window (last 10): 85.3 — 92.3 (avg 89.1)
```

#### 2b. History table (all runs, newest first)
```
| # | Date       | Commit  | Mode | Score | Pass | Best               | Worst              |
|---|------------|---------|------|-------|------|--------------------|---------------------|
| 22| 2026-02-24 | 6baa479 | full | 92.3  | PASS | first_session 95   | difficulty_climb 85 |
```

#### 2c. Per-scenario analysis (windowed)
```
| Scenario         | Avg  | Current | Trend | σ   | Min  | Max  |
|------------------|------|---------|-------|-----|------|------|
| first_session    | 87.2 | 95.0    | ↑     | 5.1 | 67.0 | 95.0 |
```
Regressions flagged: current < avg - σ

#### 2d. Per-criterion L2 trends (windowed)
```
| Criterion  | Avg  | Current | Trend |
|------------|------|---------|-------|
| socratic   | 4.33 | 4.0     | ↓     |
| tone       | 4.86 | 4.9     | →     |
```

#### 2e. Evaluator noise (duplicate commits)
```
Commit abc: 89, 90, 91 → σ=1.0
```

### Parameters:
- `--window N`: analysis window (default: 10)
- `--format md|json`: output format (default: md)
- `--historia PATH`: path to historia.jsonl (default: auto-detect relative to reports dir)

## Component 3: LLM Interpretation (test-tutor SKILL.md changes)

After generating the individual report, test-tutor:

1. **Appends JSONL entry** — parses current run results into schema, appends to `historia.jsonl`
2. **Calls Go CLI** — `matura test-report summary --format md` → gets deterministic tables
3. **Writes interpretation** — reads Go output (numbers only), adds:
   - What changed since last run?
   - Top 3 issues to fix
   - Stubborn problems (appearing in N+ consecutive runs)
   - What's working well
4. **Saves** `RAPORT_ZBIORCZY.md` (Go tables + LLM interpretation combined)

### Key principle:
LLM does NOT compute — it takes numbers from Go output.
LLM DOES interpret — explains why, recommends what next.

## Component 4: Migration (one-time)

Convert 22 existing markdown reports → `historia.jsonl` entries.

### Approach:
- LLM-assisted (agent reads all reports, extracts structured data)
- Two format paths:
  - **Old format (Feb 17-22):** `l1_percent = null`, `l2 = {all 7-8 criteria}`
  - **New format (Feb 24+):** full L1 + L2 data
- Output: chronologically ordered JSONL
- Go CLI handles null L1 gracefully (uses `overall_score` for trends)

## File locations

| What | Path |
|------|------|
| JSONL data | `analiza/test_pedagogical/reports/historia.jsonl` |
| Cumulative report | `analiza/test_pedagogical/reports/RAPORT_ZBIORCZY.md` |
| Individual reports | `analiza/test_pedagogical/reports/YYYY-MM-DD_COMMIT.md` (unchanged) |
| Go CLI source | `analiza/cli/` (new command in commands.go) |
| Test-tutor skill | `.claude/skills/test-tutor/SKILL.md` (modified) |

## Out of scope
- Web UI / dashboard
- Alerting / notifications
- Automatic SKILL.md fixes based on report
