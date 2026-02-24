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

// === JSONL types for test-tutor report history ===

type ReportEntry struct {
	Date          string          `json:"date"`
	Commit        string          `json:"commit"`
	Mode          string          `json:"mode"`
	OverallScore  float64         `json:"overall_score"`
	Pass          bool            `json:"pass"`
	ScenarioCount int             `json:"scenario_count"`
	Scenarios     []ScenarioEntry `json:"scenarios"`
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
	scanner.Buffer(make([]byte, 0, 64*1024), 1024*1024)
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

// === Summary types ===

type Summary struct {
	LatestScore float64 `json:"latest_score"`
	LatestPass  bool    `json:"latest_pass"`
	Delta       float64 `json:"delta"`
	PassStreak  int     `json:"pass_streak"`
	TotalRuns   int     `json:"total_runs"`
	WindowSize  int     `json:"window_size"`
	WindowAvg   float64 `json:"window_avg"`
	WindowMin   float64 `json:"window_min"`
	WindowMax   float64 `json:"window_max"`

	History          []HistoryRow      `json:"history"`
	ScenarioStats    []ScenarioStat    `json:"scenario_stats"`
	CriterionStats   []CriterionStat   `json:"criterion_stats"`
	DuplicateCommits []DuplicateCommit `json:"duplicate_commits"`
}

type HistoryRow struct {
	Index  int     `json:"index"`
	Date   string  `json:"date"`
	Commit string  `json:"commit"`
	Mode   string  `json:"mode"`
	Score  float64 `json:"score"`
	Pass   bool    `json:"pass"`
	Best   string  `json:"best"`
	Worst  string  `json:"worst"`
}

type ScenarioStat struct {
	Scenario  string  `json:"scenario"`
	Avg       float64 `json:"avg"`
	Current   float64 `json:"current"`
	Trend     string  `json:"trend"`
	StdDev    float64 `json:"std_dev"`
	Min       float64 `json:"min"`
	Max       float64 `json:"max"`
	Regressed bool    `json:"regressed"`
}

type CriterionStat struct {
	Name    string  `json:"name"`
	Avg     float64 `json:"avg"`
	Current float64 `json:"current"`
	Trend   string  `json:"trend"`
}

type DuplicateCommit struct {
	Commit string    `json:"commit"`
	Scores []float64 `json:"scores"`
	StdDev float64   `json:"std_dev"`
}

// === Math helpers ===

func mean(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}
	sum := 0.0
	for _, v := range vals {
		sum += v
	}
	return sum / float64(len(vals))
}

func stddev(vals []float64) float64 {
	if len(vals) == 0 {
		return 0
	}
	m := mean(vals)
	sum := 0.0
	for _, v := range vals {
		d := v - m
		sum += d * d
	}
	return math.Sqrt(sum / float64(len(vals)))
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
	avg := mean(history)
	if current > avg+1 {
		return "↑"
	}
	if current < avg-1 {
		return "↓"
	}
	return "→"
}

// === ComputeSummary ===

func ComputeSummary(entries []ReportEntry, windowSize int) Summary {
	s := Summary{}
	if len(entries) == 0 {
		return s
	}

	// 1. Sort by date ascending (stable)
	sorted := make([]ReportEntry, len(entries))
	copy(sorted, entries)
	sort.SliceStable(sorted, func(i, j int) bool {
		return sorted[i].Date < sorted[j].Date
	})

	n := len(sorted)
	s.TotalRuns = n

	// 2. Dashboard: latest score/pass, delta, pass streak
	latest := sorted[n-1]
	s.LatestScore = latest.OverallScore
	s.LatestPass = latest.Pass
	if n >= 2 {
		s.Delta = latest.OverallScore - sorted[n-2].OverallScore
	}
	// Pass streak from end
	for i := n - 1; i >= 0; i-- {
		if sorted[i].Pass {
			s.PassStreak++
		} else {
			break
		}
	}

	// 3. Window stats
	winStart := 0
	if n > windowSize {
		winStart = n - windowSize
	}
	window := sorted[winStart:]
	s.WindowSize = len(window)

	windowScores := make([]float64, len(window))
	for i, e := range window {
		windowScores[i] = e.OverallScore
	}
	s.WindowAvg = mean(windowScores)
	s.WindowMin = minFloat(windowScores)
	s.WindowMax = maxFloat(windowScores)

	// 4. History rows (newest first)
	s.History = make([]HistoryRow, n)
	for i := 0; i < n; i++ {
		e := sorted[n-1-i]
		best, worst := "", ""
		if len(e.Scenarios) > 0 {
			bestScore := e.Scenarios[0].Score
			worstScore := e.Scenarios[0].Score
			best = e.Scenarios[0].Scenario
			worst = e.Scenarios[0].Scenario
			for _, sc := range e.Scenarios[1:] {
				if sc.Score > bestScore {
					bestScore = sc.Score
					best = sc.Scenario
				}
				if sc.Score < worstScore {
					worstScore = sc.Score
					worst = sc.Scenario
				}
			}
		}
		s.History[i] = HistoryRow{
			Index:  i + 1,
			Date:   e.Date,
			Commit: e.Commit,
			Mode:   e.Mode,
			Score:  e.OverallScore,
			Pass:   e.Pass,
			Best:   best,
			Worst:  worst,
		}
	}

	// 5. Per-scenario stats (windowed)
	scenarioScores := map[string][]float64{}
	scenarioCurrent := map[string]float64{}
	scenarioOrder := []string{}
	for _, e := range window {
		for _, sc := range e.Scenarios {
			if _, exists := scenarioScores[sc.Scenario]; !exists {
				scenarioOrder = append(scenarioOrder, sc.Scenario)
			}
			scenarioScores[sc.Scenario] = append(scenarioScores[sc.Scenario], sc.Score)
		}
	}
	// Current = from latest run
	for _, sc := range latest.Scenarios {
		scenarioCurrent[sc.Scenario] = sc.Score
	}

	s.ScenarioStats = make([]ScenarioStat, 0, len(scenarioOrder))
	for _, name := range scenarioOrder {
		scores := scenarioScores[name]
		avg := mean(scores)
		sd := stddev(scores)
		cur := scenarioCurrent[name]
		s.ScenarioStats = append(s.ScenarioStats, ScenarioStat{
			Scenario:  name,
			Avg:       avg,
			Current:   cur,
			Trend:     trendArrow(scores, cur),
			StdDev:    sd,
			Min:       minFloat(scores),
			Max:       maxFloat(scores),
			Regressed: cur < avg-sd,
		})
	}

	// 6. Per-criterion L2 stats (windowed)
	criterionAll := map[string][]float64{}
	criterionCurrent := map[string]float64{}
	criterionOrder := []string{}
	for _, e := range window {
		for _, sc := range e.Scenarios {
			for k, v := range sc.L2 {
				if _, exists := criterionAll[k]; !exists {
					criterionOrder = append(criterionOrder, k)
				}
				criterionAll[k] = append(criterionAll[k], v)
			}
		}
	}
	// Current = from latest run's scenarios
	for _, sc := range latest.Scenarios {
		for k, v := range sc.L2 {
			// Use the last value seen (overwrite) — last scenario wins
			criterionCurrent[k] = v
		}
	}
	// Actually, collect average of current run's L2 values per criterion
	criterionCurrentAll := map[string][]float64{}
	for _, sc := range latest.Scenarios {
		for k, v := range sc.L2 {
			criterionCurrentAll[k] = append(criterionCurrentAll[k], v)
		}
	}
	for k, vals := range criterionCurrentAll {
		criterionCurrent[k] = mean(vals)
	}

	sort.Strings(criterionOrder)
	s.CriterionStats = make([]CriterionStat, 0, len(criterionOrder))
	for _, name := range criterionOrder {
		// Only include criteria present in the latest run
		if _, inLatest := criterionCurrent[name]; !inLatest {
			continue
		}
		vals := criterionAll[name]
		cur := criterionCurrent[name]
		s.CriterionStats = append(s.CriterionStats, CriterionStat{
			Name:    name,
			Avg:     mean(vals),
			Current: cur,
			Trend:   trendArrow(vals, cur),
		})
	}

	// 7. Duplicate commits
	commitScores := map[string][]float64{}
	commitOrder := []string{}
	for _, e := range sorted {
		if _, exists := commitScores[e.Commit]; !exists {
			commitOrder = append(commitOrder, e.Commit)
		}
		commitScores[e.Commit] = append(commitScores[e.Commit], e.OverallScore)
	}
	s.DuplicateCommits = nil
	for _, c := range commitOrder {
		scores := commitScores[c]
		if len(scores) >= 2 {
			s.DuplicateCommits = append(s.DuplicateCommits, DuplicateCommit{
				Commit: c,
				Scores: scores,
				StdDev: stddev(scores),
			})
		}
	}

	return s
}

// === Markdown Renderer ===

func RenderSummaryMarkdown(s Summary) string {
	var b strings.Builder

	// Dashboard
	passLabel := "FAIL"
	if s.LatestPass {
		passLabel = "PASS"
	}
	deltaSign := ""
	if s.Delta > 0 {
		deltaSign = "+"
	}
	b.WriteString("## Dashboard\n\n")
	b.WriteString(fmt.Sprintf("- **Wynik**: %.1f  %s\n", s.LatestScore, passLabel))
	b.WriteString(fmt.Sprintf("- **Delta**: %s%.1f\n", deltaSign, s.Delta))
	b.WriteString(fmt.Sprintf("- **Seria PASS**: %d\n", s.PassStreak))
	b.WriteString(fmt.Sprintf("- **Liczba uruchomien**: %d\n", s.TotalRuns))
	b.WriteString(fmt.Sprintf("- **Okno (%d)**: avg=%.1f  min=%.1f  max=%.1f\n\n",
		s.WindowSize, s.WindowAvg, s.WindowMin, s.WindowMax))

	// Historia
	b.WriteString("## Historia\n\n")
	b.WriteString("| # | Data | Commit | Mode | Wynik | Pass | Best | Worst |\n")
	b.WriteString("|---|------|--------|------|-------|------|------|-------|\n")
	for _, h := range s.History {
		passStr := "PASS"
		if !h.Pass {
			passStr = "FAIL"
		}
		b.WriteString(fmt.Sprintf("| %d | %s | %s | %s | %.1f | %s | %s | %s |\n",
			h.Index, h.Date, h.Commit, h.Mode, h.Score, passStr, h.Best, h.Worst))
	}
	b.WriteString("\n")

	// Analiza per-scenariusz
	b.WriteString("## Analiza per-scenariusz\n\n")
	b.WriteString("| Scenariusz | Avg | Current | Trend | StdDev | Min | Max | Regressed |\n")
	b.WriteString("|------------|-----|---------|-------|--------|-----|-----|-----------|\n")
	for _, sc := range s.ScenarioStats {
		regFlag := ""
		if sc.Regressed {
			regFlag = "⚠"
		}
		b.WriteString(fmt.Sprintf("| %s | %.1f | %.1f | %s | %.1f | %.1f | %.1f | %s |\n",
			sc.Scenario, sc.Avg, sc.Current, sc.Trend, sc.StdDev, sc.Min, sc.Max, regFlag))
	}
	b.WriteString("\n")

	// Trendy kryteriow L2
	b.WriteString("## Trendy kryteriow L2\n\n")
	b.WriteString("| Kryterium | Avg | Current | Trend |\n")
	b.WriteString("|-----------|-----|---------|-------|\n")
	for _, cr := range s.CriterionStats {
		b.WriteString(fmt.Sprintf("| %s | %.1f | %.1f | %s |\n",
			cr.Name, cr.Avg, cr.Current, cr.Trend))
	}
	b.WriteString("\n")

	// Szum ewaluatora (only if duplicates exist)
	if len(s.DuplicateCommits) > 0 {
		b.WriteString("## Szum ewaluatora (duplikaty commitow)\n\n")
		b.WriteString("| Commit | Scores | StdDev |\n")
		b.WriteString("|--------|--------|--------|\n")
		for _, dc := range s.DuplicateCommits {
			scoreStrs := make([]string, len(dc.Scores))
			for i, sc := range dc.Scores {
				scoreStrs[i] = fmt.Sprintf("%.1f", sc)
			}
			b.WriteString(fmt.Sprintf("| %s | %s | %.1f |\n",
				dc.Commit, strings.Join(scoreStrs, ", "), dc.StdDev))
		}
		b.WriteString("\n")
	}

	return b.String()
}
