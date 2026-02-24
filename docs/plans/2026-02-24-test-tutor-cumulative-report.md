# Test-Tutor Cumulative Report Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Automated cumulative report after each test-tutor run — Go CLI computes deterministic stats from JSONL history, LLM interprets trends.

**Architecture:** New file `report.go` in CLI handles JSONL parsing + summary generation. Test-tutor SKILL.md appends structured entry to `historia.jsonl` after each run, calls `matura test-report summary`, then writes LLM interpretation into `RAPORT_ZBIORCZY.md`. One-time migration converts 24 existing markdown reports to JSONL.

**Tech Stack:** Go (cobra CLI), JSONL (append-only structured data), test-tutor SKILL.md (LLM layer)

**Design doc:** `docs/plans/2026-02-24-test-tutor-cumulative-report-design.md`

---

### Task 1: JSONL Types + Parser

**Files:**
- Create: `analiza/cli/report.go`
- Create: `analiza/cli/report_test.go`

**Step 1: Write the failing test for JSONL parsing**

```go
// report_test.go
package main

import (
	"os"
	"path/filepath"
	"testing"
)

func TestParseHistoria_Empty(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "historia.jsonl")
	os.WriteFile(path, []byte(""), 0644)

	entries, err := ParseHistoria(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("got %d entries, want 0", len(entries))
	}
}

func TestParseHistoria_SingleEntry(t *testing.T) {
	dir := t.TempDir()
	path := filepath.Join(dir, "historia.jsonl")
	line := `{"date":"2026-02-24","commit":"abc1234","mode":"full","overall_score":92.3,"pass":true,"scenario_count":2,"scenarios":[{"persona":"beginner","scenario":"first_session","score":95.0,"l1_percent":100.0,"l1_total":10,"l1_passed":10,"l2":{"socratic":4,"tone":5},"issues":["issue1"]},{"persona":"intermediate","scenario":"difficulty_climb","score":85.0,"l1_percent":null,"l1_total":null,"l1_passed":null,"l2":{"socratic":3,"tone":5},"issues":[]}]}` + "\n"
	os.WriteFile(path, []byte(line), 0644)

	entries, err := ParseHistoria(path)
	if err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	if len(entries) != 1 {
		t.Errorf("got %d entries, want 1", len(entries))
	}
	e := entries[0]
	if e.Commit != "abc1234" {
		t.Errorf("commit: got %q, want %q", e.Commit, "abc1234")
	}
	if e.OverallScore != 92.3 {
		t.Errorf("overall_score: got %f, want 92.3", e.OverallScore)
	}
	if len(e.Scenarios) != 2 {
		t.Errorf("scenarios: got %d, want 2", len(e.Scenarios))
	}
	// First scenario has L1
	s0 := e.Scenarios[0]
	if s0.L1Percent == nil || *s0.L1Percent != 100.0 {
		t.Errorf("scenario 0 l1_percent: got %v, want 100.0", s0.L1Percent)
	}
	// Second scenario has null L1
	s1 := e.Scenarios[1]
	if s1.L1Percent != nil {
		t.Errorf("scenario 1 l1_percent: got %v, want nil", s1.L1Percent)
	}
}

func TestParseHistoria_FileNotFound(t *testing.T) {
	entries, err := ParseHistoria("/nonexistent/historia.jsonl")
	if err != nil {
		t.Fatalf("missing file should return empty, got error: %v", err)
	}
	if len(entries) != 0 {
		t.Errorf("got %d entries, want 0", len(entries))
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd analiza/cli && go test -run TestParseHistoria -v`
Expected: FAIL — `ParseHistoria` not defined

**Step 3: Write types + parser**

```go
// report.go
package main

import (
	"bufio"
	"encoding/json"
	"errors"
	"fmt"
	"math"
	"os"
	"sort"
	"strings"
)

// === JSONL types ===

type ReportEntry struct {
	Date          string           `json:"date"`
	Commit        string           `json:"commit"`
	Mode          string           `json:"mode"`
	OverallScore  float64          `json:"overall_score"`
	Pass          bool             `json:"pass"`
	ScenarioCount int              `json:"scenario_count"`
	Scenarios     []ScenarioEntry  `json:"scenarios"`
}

type ScenarioEntry struct {
	Persona   string             `json:"persona"`
	Scenario  string             `json:"scenario"`
	Score     float64            `json:"score"`
	L1Percent *float64           `json:"l1_percent"`
	L1Total   *int               `json:"l1_total"`
	L1Passed  *int               `json:"l1_passed"`
	L2        map[string]float64 `json:"l2"`
	Issues    []string           `json:"issues"`
}

// ParseHistoria reads a JSONL file and returns all entries.
// Returns empty slice (no error) if file doesn't exist.
func ParseHistoria(path string) ([]ReportEntry, error) {
	f, err := os.Open(path)
	if err != nil {
		if errors.Is(err, os.ErrNotExist) {
			return nil, nil
		}
		return nil, fmt.Errorf("open historia: %w", err)
	}
	defer f.Close()

	var entries []ReportEntry
	scanner := bufio.NewScanner(f)
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024) // 1MB max line
	lineNum := 0
	for scanner.Scan() {
		lineNum++
		line := strings.TrimSpace(scanner.Text())
		if line == "" {
			continue
		}
		var e ReportEntry
		if err := json.Unmarshal([]byte(line), &e); err != nil {
			return nil, fmt.Errorf("line %d: %w", lineNum, err)
		}
		entries = append(entries, e)
	}
	if err := scanner.Err(); err != nil {
		return nil, fmt.Errorf("scan historia: %w", err)
	}
	return entries, nil
}
```

**Step 4: Run test to verify it passes**

Run: `cd analiza/cli && go test -run TestParseHistoria -v`
Expected: PASS (all 3 tests)

**Step 5: Commit**

```bash
git add analiza/cli/report.go analiza/cli/report_test.go
git commit -m "feat(cli): add JSONL types and parser for test-tutor report history"
```

---

### Task 2: Summary Computation

**Files:**
- Modify: `analiza/cli/report.go`
- Modify: `analiza/cli/report_test.go`

**Step 1: Write the failing test for summary computation**

Add to `report_test.go`:

```go
func sampleEntries() []ReportEntry {
	return []ReportEntry{
		{
			Date: "2026-02-20", Commit: "aaa1111", Mode: "full",
			OverallScore: 85.0, Pass: true, ScenarioCount: 2,
			Scenarios: []ScenarioEntry{
				{Persona: "beginner", Scenario: "first_session", Score: 80.0,
					L2: map[string]float64{"socratic": 4, "tone": 5}, Issues: []string{"issue A"}},
				{Persona: "intermediate", Scenario: "difficulty_climb", Score: 90.0,
					L2: map[string]float64{"socratic": 5, "tone": 5}, Issues: []string{}},
			},
		},
		{
			Date: "2026-02-22", Commit: "bbb2222", Mode: "full",
			OverallScore: 88.0, Pass: true, ScenarioCount: 2,
			Scenarios: []ScenarioEntry{
				{Persona: "beginner", Scenario: "first_session", Score: 90.0,
					L1Percent: ptr(95.0), L1Total: intPtr(10), L1Passed: intPtr(9),
					L2: map[string]float64{"socratic": 4, "tone": 5}, Issues: []string{}},
				{Persona: "intermediate", Scenario: "difficulty_climb", Score: 86.0,
					L1Percent: ptr(100.0), L1Total: intPtr(8), L1Passed: intPtr(8),
					L2: map[string]float64{"socratic": 3, "tone": 5}, Issues: []string{"issue B"}},
			},
		},
		{
			Date: "2026-02-24", Commit: "ccc3333", Mode: "full",
			OverallScore: 92.3, Pass: true, ScenarioCount: 2,
			Scenarios: []ScenarioEntry{
				{Persona: "beginner", Scenario: "first_session", Score: 95.0,
					L1Percent: ptr(100.0), L1Total: intPtr(10), L1Passed: intPtr(10),
					L2: map[string]float64{"socratic": 4, "tone": 5}, Issues: []string{}},
				{Persona: "intermediate", Scenario: "difficulty_climb", Score: 85.0,
					L1Percent: ptr(90.0), L1Total: intPtr(10), L1Passed: intPtr(9),
					L2: map[string]float64{"socratic": 3, "tone": 5}, Issues: []string{"issue C"}},
			},
		},
	}
}

func ptr(f float64) *float64 { return &f }
func intPtr(i int) *int      { return &i }

func TestComputeSummary_Dashboard(t *testing.T) {
	entries := sampleEntries()
	s := ComputeSummary(entries, 10)

	if s.LatestScore != 92.3 {
		t.Errorf("latest score: got %f, want 92.3", s.LatestScore)
	}
	if !s.LatestPass {
		t.Error("latest pass: got false, want true")
	}
	if s.PassStreak != 3 {
		t.Errorf("pass streak: got %d, want 3", s.PassStreak)
	}
	if s.TotalRuns != 3 {
		t.Errorf("total runs: got %d, want 3", s.TotalRuns)
	}
	// Delta vs previous
	expectedDelta := 92.3 - 88.0
	if math.Abs(s.Delta-expectedDelta) > 0.01 {
		t.Errorf("delta: got %f, want %f", s.Delta, expectedDelta)
	}
}

func TestComputeSummary_PerScenario(t *testing.T) {
	entries := sampleEntries()
	s := ComputeSummary(entries, 10)

	// Should have 2 scenarios
	if len(s.ScenarioStats) != 2 {
		t.Fatalf("scenario stats: got %d, want 2", len(s.ScenarioStats))
	}

	// Find first_session
	var fs *ScenarioStat
	for i := range s.ScenarioStats {
		if s.ScenarioStats[i].Scenario == "first_session" {
			fs = &s.ScenarioStats[i]
			break
		}
	}
	if fs == nil {
		t.Fatal("first_session not found")
	}
	// Avg of 80, 90, 95 = 88.33
	expectedAvg := (80.0 + 90.0 + 95.0) / 3.0
	if math.Abs(fs.Avg-expectedAvg) > 0.01 {
		t.Errorf("first_session avg: got %f, want %f", fs.Avg, expectedAvg)
	}
	if fs.Current != 95.0 {
		t.Errorf("first_session current: got %f, want 95.0", fs.Current)
	}
}

func TestComputeSummary_DuplicateCommits(t *testing.T) {
	entries := []ReportEntry{
		{Date: "2026-02-20", Commit: "aaa1111", Mode: "full", OverallScore: 89.0, Pass: true, Scenarios: []ScenarioEntry{}},
		{Date: "2026-02-20", Commit: "aaa1111", Mode: "full", OverallScore: 91.0, Pass: true, Scenarios: []ScenarioEntry{}},
	}
	s := ComputeSummary(entries, 10)

	if len(s.DuplicateCommits) != 1 {
		t.Fatalf("duplicate commits: got %d, want 1", len(s.DuplicateCommits))
	}
	dc := s.DuplicateCommits[0]
	if dc.Commit != "aaa1111" {
		t.Errorf("commit: got %q, want %q", dc.Commit, "aaa1111")
	}
	if len(dc.Scores) != 2 {
		t.Errorf("scores count: got %d, want 2", len(dc.Scores))
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd analiza/cli && go test -run TestComputeSummary -v`
Expected: FAIL — `ComputeSummary` not defined

**Step 3: Write summary computation**

Add to `report.go`:

```go
// === Summary types ===

type Summary struct {
	// Dashboard
	LatestScore float64
	LatestPass  bool
	Delta       float64 // vs previous run
	PassStreak  int
	TotalRuns   int
	WindowSize  int
	WindowAvg   float64
	WindowMin   float64
	WindowMax   float64

	// History (all runs, newest first)
	History []HistoryRow

	// Per-scenario (windowed)
	ScenarioStats []ScenarioStat

	// Per-criterion L2 (windowed)
	CriterionStats []CriterionStat

	// Duplicate commits
	DuplicateCommits []DuplicateCommit
}

type HistoryRow struct {
	Index    int
	Date     string
	Commit   string
	Mode     string
	Score    float64
	Pass     bool
	Best     string // "scenario score"
	Worst    string // "scenario score"
}

type ScenarioStat struct {
	Scenario  string
	Avg       float64
	Current   float64
	Trend     string // ↑ ↓ →
	StdDev    float64
	Min       float64
	Max       float64
	Regressed bool // current < avg - stddev
}

type CriterionStat struct {
	Name    string
	Avg     float64
	Current float64
	Trend   string
}

type DuplicateCommit struct {
	Commit string
	Scores []float64
	StdDev float64
}

// ComputeSummary computes deterministic summary from report entries.
// window controls how many recent entries are used for per-scenario/criterion analysis.
func ComputeSummary(entries []ReportEntry, window int) Summary {
	if len(entries) == 0 {
		return Summary{}
	}

	// Sort by date ascending (stable for same date)
	sorted := make([]ReportEntry, len(entries))
	copy(sorted, entries)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Date < sorted[j].Date
	})

	latest := sorted[len(sorted)-1]

	s := Summary{
		LatestScore: latest.OverallScore,
		LatestPass:  latest.Pass,
		TotalRuns:   len(sorted),
		WindowSize:  window,
	}

	// Delta
	if len(sorted) >= 2 {
		s.Delta = latest.OverallScore - sorted[len(sorted)-2].OverallScore
	}

	// Pass streak (count from end)
	for i := len(sorted) - 1; i >= 0; i-- {
		if sorted[i].Pass {
			s.PassStreak++
		} else {
			break
		}
	}

	// History rows (newest first)
	for i := len(sorted) - 1; i >= 0; i-- {
		e := sorted[i]
		row := HistoryRow{
			Index:  i + 1,
			Date:   e.Date,
			Commit: e.Commit,
			Mode:   e.Mode,
			Score:  e.OverallScore,
			Pass:   e.Pass,
		}
		// Best/worst scenario
		if len(e.Scenarios) > 0 {
			best, worst := e.Scenarios[0], e.Scenarios[0]
			for _, sc := range e.Scenarios[1:] {
				if sc.Score > best.Score {
					best = sc
				}
				if sc.Score < worst.Score {
					worst = sc
				}
			}
			row.Best = fmt.Sprintf("%s %.0f", best.Scenario, best.Score)
			row.Worst = fmt.Sprintf("%s %.0f", worst.Scenario, worst.Score)
		}
		s.History = append(s.History, row)
	}

	// Windowed entries
	winStart := len(sorted) - window
	if winStart < 0 {
		winStart = 0
	}
	windowed := sorted[winStart:]

	// Window stats
	var winSum float64
	s.WindowMin = windowed[0].OverallScore
	s.WindowMax = windowed[0].OverallScore
	for _, e := range windowed {
		winSum += e.OverallScore
		if e.OverallScore < s.WindowMin {
			s.WindowMin = e.OverallScore
		}
		if e.OverallScore > s.WindowMax {
			s.WindowMax = e.OverallScore
		}
	}
	s.WindowAvg = winSum / float64(len(windowed))

	// Per-scenario stats (windowed)
	scenarioScores := map[string][]float64{}
	scenarioCurrent := map[string]float64{}
	for i, e := range windowed {
		for _, sc := range e.Scenarios {
			scenarioScores[sc.Scenario] = append(scenarioScores[sc.Scenario], sc.Score)
			if i == len(windowed)-1 {
				scenarioCurrent[sc.Scenario] = sc.Score
			}
		}
	}
	for name, scores := range scenarioScores {
		avg := mean(scores)
		sd := stddev(scores)
		current := scenarioCurrent[name]
		trend := trendArrow(scores, current)
		s.ScenarioStats = append(s.ScenarioStats, ScenarioStat{
			Scenario:  name,
			Avg:       avg,
			Current:   current,
			Trend:     trend,
			StdDev:    sd,
			Min:       minFloat(scores),
			Max:       maxFloat(scores),
			Regressed: current < avg-sd,
		})
	}
	sort.Slice(s.ScenarioStats, func(i, j int) bool {
		return s.ScenarioStats[i].Scenario < s.ScenarioStats[j].Scenario
	})

	// Per-criterion L2 stats (windowed)
	criterionScores := map[string][]float64{}
	criterionCurrent := map[string]float64{}
	for i, e := range windowed {
		for _, sc := range e.Scenarios {
			for k, v := range sc.L2 {
				criterionScores[k] = append(criterionScores[k], v)
				if i == len(windowed)-1 {
					criterionCurrent[k] = v // last scenario wins for current
				}
			}
		}
	}
	for name, scores := range criterionScores {
		avg := mean(scores)
		current := criterionCurrent[name]
		trend := trendArrow(scores, current)
		s.CriterionStats = append(s.CriterionStats, CriterionStat{
			Name:    name,
			Avg:     avg,
			Current: current,
			Trend:   trend,
		})
	}
	sort.Slice(s.CriterionStats, func(i, j int) bool {
		return s.CriterionStats[i].Name < s.CriterionStats[j].Name
	})

	// Duplicate commits
	commitRuns := map[string][]float64{}
	for _, e := range sorted {
		commitRuns[e.Commit] = append(commitRuns[e.Commit], e.OverallScore)
	}
	for commit, scores := range commitRuns {
		if len(scores) > 1 {
			s.DuplicateCommits = append(s.DuplicateCommits, DuplicateCommit{
				Commit: commit,
				Scores: scores,
				StdDev: stddev(scores),
			})
		}
	}
	sort.Slice(s.DuplicateCommits, func(i, j int) bool {
		return s.DuplicateCommits[i].Commit < s.DuplicateCommits[j].Commit
	})

	return s
}

// === math helpers ===

func mean(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}
	var sum float64
	for _, v := range vals {
		sum += v
	}
	return sum / float64(len(vals))
}

func stddev(vals []float64) float64 {
	if len(vals) < 2 {
		return 0
	}
	m := mean(vals)
	var sumSq float64
	for _, v := range vals {
		d := v - m
		sumSq += d * d
	}
	return math.Sqrt(sumSq / float64(len(vals)))
}

func minFloat(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v < m {
			m = v
		}
	}
	return m
}

func maxFloat(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}
	m := vals[0]
	for _, v := range vals[1:] {
		if v > m {
			m = v
		}
	}
	return m
}

func trendArrow(history []float64, current float64) string {
	if len(history) < 2 {
		return "→"
	}
	avg := mean(history)
	diff := current - avg
	threshold := 1.0
	switch {
	case diff > threshold:
		return "↑"
	case diff < -threshold:
		return "↓"
	default:
		return "→"
	}
}
```

**Step 4: Run test to verify it passes**

Run: `cd analiza/cli && go test -run TestComputeSummary -v`
Expected: PASS (all 3 tests)

**Step 5: Commit**

```bash
git add analiza/cli/report.go analiza/cli/report_test.go
git commit -m "feat(cli): add summary computation for test-tutor report history"
```

---

### Task 3: Markdown Renderer

**Files:**
- Modify: `analiza/cli/report.go`
- Modify: `analiza/cli/report_test.go`

**Step 1: Write the failing test for markdown output**

Add to `report_test.go`:

```go
func TestRenderMarkdown_ContainsSections(t *testing.T) {
	entries := sampleEntries()
	s := ComputeSummary(entries, 10)
	md := RenderSummaryMarkdown(s)

	sections := []string{
		"## Dashboard",
		"## Historia",
		"## Analiza per-scenariusz",
		"## Trendy kryteriow L2",
		"92.3",   // latest score
		"PASS",   // pass status
		"first_session",
		"difficulty_climb",
	}
	for _, want := range sections {
		if !strings.Contains(md, want) {
			t.Errorf("markdown missing %q", want)
		}
	}
}

func TestRenderMarkdown_DuplicateCommitsSection(t *testing.T) {
	entries := []ReportEntry{
		{Date: "2026-02-20", Commit: "aaa1111", Mode: "full", OverallScore: 89.0, Pass: true, Scenarios: []ScenarioEntry{}},
		{Date: "2026-02-20", Commit: "aaa1111", Mode: "full", OverallScore: 91.0, Pass: true, Scenarios: []ScenarioEntry{}},
	}
	s := ComputeSummary(entries, 10)
	md := RenderSummaryMarkdown(s)

	if !strings.Contains(md, "Szum ewaluatora") {
		t.Error("markdown missing duplicate commits section")
	}
	if !strings.Contains(md, "aaa1111") {
		t.Error("markdown missing commit hash in duplicate section")
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd analiza/cli && go test -run TestRenderMarkdown -v`
Expected: FAIL — `RenderSummaryMarkdown` not defined

**Step 3: Write markdown renderer**

Add to `report.go`:

```go
// RenderSummaryMarkdown renders the summary as a markdown report.
func RenderSummaryMarkdown(s Summary) string {
	var b strings.Builder

	b.WriteString("# Test-Tutor Report Summary\n\n")

	// Dashboard
	b.WriteString("## Dashboard\n\n")
	passStr := "PASS"
	if !s.LatestPass {
		passStr = "FAIL"
	}
	deltaStr := ""
	if s.TotalRuns >= 2 {
		sign := "+"
		if s.Delta < 0 {
			sign = ""
		}
		deltaStr = fmt.Sprintf(" (%s%.1f vs poprzedni)", sign, s.Delta)
	}
	fmt.Fprintf(&b, "- **Score**: %.1f/100 (%s)%s\n", s.LatestScore, passStr, deltaStr)
	fmt.Fprintf(&b, "- **Seria PASS**: %d consecutive\n", s.PassStreak)
	fmt.Fprintf(&b, "- **Runs**: %d total\n", s.TotalRuns)
	if s.TotalRuns > 1 {
		fmt.Fprintf(&b, "- **Okno** (last %d): %.1f — %.1f (avg %.1f)\n",
			s.WindowSize, s.WindowMin, s.WindowMax, s.WindowAvg)
	}
	b.WriteString("\n")

	// History table
	b.WriteString("## Historia\n\n")
	b.WriteString("| # | Data | Commit | Mode | Score | Pass | Najlepszy | Najslabszy |\n")
	b.WriteString("|---|------|--------|------|-------|------|-----------|------------|\n")
	for _, h := range s.History {
		passCell := "PASS"
		if !h.Pass {
			passCell = "FAIL"
		}
		fmt.Fprintf(&b, "| %d | %s | %s | %s | %.1f | %s | %s | %s |\n",
			h.Index, h.Date, h.Commit, h.Mode, h.Score, passCell, h.Best, h.Worst)
	}
	b.WriteString("\n")

	// Per-scenario
	if len(s.ScenarioStats) > 0 {
		b.WriteString("## Analiza per-scenariusz\n\n")
		b.WriteString("| Scenariusz | Avg | Current | Trend | σ | Min | Max | Regresja |\n")
		b.WriteString("|------------|-----|---------|-------|---|-----|-----|----------|\n")
		for _, sc := range s.ScenarioStats {
			regFlag := ""
			if sc.Regressed {
				regFlag = "⚠"
			}
			fmt.Fprintf(&b, "| %s | %.1f | %.1f | %s | %.1f | %.1f | %.1f | %s |\n",
				sc.Scenario, sc.Avg, sc.Current, sc.Trend, sc.StdDev, sc.Min, sc.Max, regFlag)
		}
		b.WriteString("\n")
	}

	// Per-criterion L2
	if len(s.CriterionStats) > 0 {
		b.WriteString("## Trendy kryteriow L2\n\n")
		b.WriteString("| Kryterium | Avg | Current | Trend |\n")
		b.WriteString("|-----------|-----|---------|-------|\n")
		for _, cr := range s.CriterionStats {
			fmt.Fprintf(&b, "| %s | %.2f | %.1f | %s |\n",
				cr.Name, cr.Avg, cr.Current, cr.Trend)
		}
		b.WriteString("\n")
	}

	// Duplicate commits
	if len(s.DuplicateCommits) > 0 {
		b.WriteString("## Szum ewaluatora (duplikaty commitow)\n\n")
		for _, dc := range s.DuplicateCommits {
			scores := make([]string, len(dc.Scores))
			for i, sc := range dc.Scores {
				scores[i] = fmt.Sprintf("%.1f", sc)
			}
			fmt.Fprintf(&b, "- `%s`: %s (σ=%.1f)\n", dc.Commit, strings.Join(scores, ", "), dc.StdDev)
		}
		b.WriteString("\n")
	}

	return b.String()
}
```

**Step 4: Run test to verify it passes**

Run: `cd analiza/cli && go test -run TestRenderMarkdown -v`
Expected: PASS

**Step 5: Commit**

```bash
git add analiza/cli/report.go analiza/cli/report_test.go
git commit -m "feat(cli): add markdown renderer for test-tutor summary report"
```

---

### Task 4: CLI Command Wiring

**Files:**
- Modify: `analiza/cli/commands.go` (add `testReportSummaryCmd()`)
- Modify: `analiza/cli/main.go` (add `test-report` parent command, skip DB for this command)

**Step 1: Write the failing test**

Add to `report_test.go`:

```go
func TestTestReportSummaryCmd_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	// Create empty historia.jsonl
	os.WriteFile(filepath.Join(dir, "historia.jsonl"), []byte(""), 0644)

	cmd := testReportSummaryCmd()
	cmd.SetArgs([]string{"--historia", filepath.Join(dir, "historia.jsonl")})
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Should produce markdown with 0 runs
	output := buf.String()
	if !strings.Contains(output, "0 total") {
		t.Errorf("expected '0 total' in output, got: %s", output)
	}
}

func TestTestReportSummaryCmd_WithData(t *testing.T) {
	dir := t.TempDir()
	line := `{"date":"2026-02-24","commit":"abc1234","mode":"full","overall_score":92.3,"pass":true,"scenario_count":1,"scenarios":[{"persona":"beginner","scenario":"first_session","score":92.3,"l1_percent":100.0,"l1_total":10,"l1_passed":10,"l2":{"socratic":4,"tone":5},"issues":[]}]}` + "\n"
	os.WriteFile(filepath.Join(dir, "historia.jsonl"), []byte(line), 0644)

	cmd := testReportSummaryCmd()
	cmd.SetArgs([]string{"--historia", filepath.Join(dir, "historia.jsonl")})
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := buf.String()
	if !strings.Contains(output, "92.3") {
		t.Errorf("expected '92.3' in output, got: %s", output)
	}
}

func TestTestReportSummaryCmd_JSONFormat(t *testing.T) {
	dir := t.TempDir()
	line := `{"date":"2026-02-24","commit":"abc1234","mode":"full","overall_score":92.3,"pass":true,"scenario_count":1,"scenarios":[{"persona":"beginner","scenario":"first_session","score":92.3,"l1_percent":100.0,"l1_total":10,"l1_passed":10,"l2":{"socratic":4,"tone":5},"issues":[]}]}` + "\n"
	os.WriteFile(filepath.Join(dir, "historia.jsonl"), []byte(line), 0644)

	cmd := testReportSummaryCmd()
	cmd.SetArgs([]string{"--historia", filepath.Join(dir, "historia.jsonl"), "--format", "json"})
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	// Should be valid JSON
	var result Summary
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Errorf("output is not valid JSON: %v", err)
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd analiza/cli && go test -run TestTestReportSummaryCmd -v`
Expected: FAIL — `testReportSummaryCmd` not defined

**Step 3: Write the CLI command**

Add at the end of `commands.go`:

```go
func testReportSummaryCmd() *cobra.Command {
	var historiaPath string
	var window int
	var format string

	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Generate summary report from test-tutor history",
		RunE: func(cmd *cobra.Command, args []string) error {
			// Auto-detect historia path if not given
			if historiaPath == "" {
				exe, err := os.Executable()
				if err == nil {
					historiaPath = filepath.Join(filepath.Dir(exe),
						"..", "test_pedagogical", "reports", "historia.jsonl")
				}
			}

			entries, err := ParseHistoria(historiaPath)
			if err != nil {
				return fmt.Errorf("parse historia: %w", err)
			}

			summary := ComputeSummary(entries, window)

			switch format {
			case "json":
				enc := json.NewEncoder(cmd.OutOrStdout())
				enc.SetIndent("", "  ")
				return enc.Encode(summary)
			default:
				md := RenderSummaryMarkdown(summary)
				fmt.Fprint(cmd.OutOrStdout(), md)
				return nil
			}
		},
	}

	cmd.Flags().StringVar(&historiaPath, "historia", "", "Path to historia.jsonl (default: auto-detect)")
	cmd.Flags().IntVar(&window, "window", 10, "Analysis window size")
	cmd.Flags().StringVar(&format, "format", "md", "Output format: md or json")

	return cmd
}
```

Add to `main.go` inside `func main()`, after the `dataCmd` block and before `rootCmd.AddCommand(...)`:

```go
	// === test-report ===
	testReportCmd := &cobra.Command{Use: "test-report", Short: "Test-tutor report analysis"}
	testReportCmd.AddCommand(testReportSummaryCmd())
```

Update the `rootCmd.AddCommand` call to include `testReportCmd`:

```go
	rootCmd.AddCommand(exerciseCmd, progressCmd, ckeCmd, examCmd, typCmd, trapCmd, cheatsheetCmd, dataCmd, testReportCmd)
```

Also update the `PersistentPreRunE` skip check to also skip DB for `test-report` commands (they read files, not DB):

```go
			// Skip DB open for commands that don't need it
			if cmd.Name() == "import" && cmd.Parent() != nil && cmd.Parent().Name() == "data" {
				return nil
			}
			if cmd.Parent() != nil && cmd.Parent().Name() == "test-report" {
				return nil
			}
```

**Step 4: Run test to verify it passes**

Run: `cd analiza/cli && go test -run TestTestReportSummaryCmd -v`
Expected: PASS (all 3 tests)

**Step 5: Run ALL existing tests to check for regressions**

Run: `cd analiza/cli && go test -v ./...`
Expected: All tests PASS (existing 13 + new ~9 = ~22 total)

**Step 6: Build and smoke test**

Run: `cd analiza/cli && ./build.sh`
Run: `cd analiza/cli && ./matura test-report summary --help`
Expected: Shows usage with --historia, --window, --format flags

**Step 7: Commit**

```bash
git add analiza/cli/commands.go analiza/cli/main.go analiza/cli/report.go analiza/cli/report_test.go
git commit -m "feat(cli): add test-report summary command"
```

---

### Task 5: Migration — Convert Old Reports to JSONL

**Files:**
- Create: `analiza/test_pedagogical/reports/historia.jsonl`

**Step 1: Migrate old reports using LLM agent**

This is a one-time operation. Use a Task agent to:
1. Read all 24 markdown reports from `analiza/test_pedagogical/reports/`
2. Extract from each: date, commit, mode, overall_score, pass, per-scenario data (persona, scenario, score, L1 if available, L2 criteria, issues)
3. Handle two formats:
   - Old (Feb 17-22): no L1, 7-8 L2 criteria
   - New (Feb 24+): L1 binary checkpoints + 2 L2 criteria
4. Write one JSONL line per report, chronologically ordered
5. Save to `analiza/test_pedagogical/reports/historia.jsonl`

**Important:** Agent must read actual report files to extract data. Do not guess values.

**Step 2: Verify migration**

Run: `cd analiza/cli && go build -o matura . && ./matura test-report summary --historia ../test_pedagogical/reports/historia.jsonl`
Expected: Produces a markdown summary with ~24 rows in the history table, reasonable scores

Run: `wc -l analiza/test_pedagogical/reports/historia.jsonl`
Expected: ~24 lines (one per report)

**Step 3: Commit**

```bash
git add analiza/test_pedagogical/reports/historia.jsonl
git commit -m "data: migrate 24 test-tutor reports to historia.jsonl"
```

---

### Task 6: Update test-tutor SKILL.md

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md` (sections 7-10)

**Step 1: Understand current sections 7-10**

Re-read sections 7-10 of SKILL.md (report generation, format, saving, cleanup).

**Step 2: Add JSONL append step after report generation**

In section 9 (Zapis raportu), after saving the individual report, add:

```markdown
5. Append JSONL entry to `{REPORT_DIR}/historia.jsonl`:
   - Parse the per-scenario results from step 3 into JSON matching this schema:
     ```json
     {"date":"{REPORT_DATE}","commit":"{COMMIT_HASH}","mode":"{MODE}","overall_score":{SCORE},"pass":{PASS},"scenario_count":{N},"scenarios":[{"persona":"...","scenario":"...","score":N,"l1_percent":N,"l1_total":N,"l1_passed":N,"l2":{"socratic":N,"tone":N},"issues":["..."]}]}
     ```
   - Use Bash: `echo '{JSON_LINE}' >> {REPORT_DIR}/historia.jsonl`
   - One line, no pretty-printing — this is JSONL (one JSON object per line)
```

**Step 3: Add cumulative report generation step**

After the JSONL append step, add:

```markdown
6. Generate cumulative report:
   a. Run: `{CLI_PATH} test-report summary --historia {REPORT_DIR}/historia.jsonl --format md`
   b. Capture the markdown output (this is the deterministic Go-computed section)
   c. Write interpretation section — read the Go output (numbers, trends, regressions) and add:
      - **Co sie zmienilo?** — explain score changes vs previous run
      - **Top 3 do naprawienia** — most impactful issues to fix next
      - **Uporczywe problemy** — issues appearing in 3+ consecutive runs (check issues arrays in historia.jsonl)
      - **Co dziala dobrze** — stable/improving metrics
   d. Combine: Go markdown + "## Interpretacja" section
   e. Save to `{REPORT_DIR}/RAPORT_ZBIORCZY.md` (Write tool — overwrites each time)
```

**Step 4: Verify SKILL.md is syntactically correct**

Run: `cd analiza && bash test_qa.sh --layer 3`
Expected: PASS (SKILL lint layer)

**Step 5: Commit**

```bash
git add .claude/skills/test-tutor/SKILL.md
git commit -m "feat(test-tutor): add cumulative report generation (JSONL + Go summary + LLM interpretation)"
```

---

### Task 7: Update build.sh and test_qa.sh

**Files:**
- Modify: `analiza/cli/build.sh` (no changes needed — build already covers all Go files)
- Modify: `analiza/test_qa.sh` (add Layer for test-report command smoke test)

**Step 1: Check if test_qa.sh needs a new smoke test**

Read `analiza/test_qa.sh` Layer 1 (CLI smoke) to understand the pattern.

**Step 2: Add smoke test for `test-report summary`**

Add to Layer 1 (CLI smoke tests):

```bash
run_test "test-report summary --historia with empty file" \
  "$CLI test-report summary --historia /dev/null" 0
```

**Step 3: Run full QA**

Run: `cd analiza && bash test_qa.sh`
Expected: All layers PASS, total test count increases by 1

**Step 4: Commit**

```bash
git add analiza/test_qa.sh
git commit -m "test: add test-report summary smoke test to QA suite"
```

---

### Task 8: End-to-End Verification

**Step 1: Build everything**

Run: `cd analiza/cli && ./build.sh`
Expected: Build OK

**Step 2: Run Go tests**

Run: `cd analiza/cli && go test -v ./...`
Expected: All PASS

**Step 3: Run full QA suite**

Run: `cd analiza && bash test_qa.sh`
Expected: All layers PASS

**Step 4: Manual smoke test with real data**

Run: `cd analiza/cli && ./matura test-report summary --historia ../test_pedagogical/reports/historia.jsonl`
Expected: Readable markdown with dashboard, history table, per-scenario analysis

**Step 5: Verify JSON format**

Run: `cd analiza/cli && ./matura test-report summary --historia ../test_pedagogical/reports/historia.jsonl --format json | python3 -m json.tool > /dev/null`
Expected: Valid JSON (exit 0)
