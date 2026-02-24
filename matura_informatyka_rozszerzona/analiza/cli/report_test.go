package main

import (
	"bytes"
	"encoding/json"
	"math"
	"os"
	"path/filepath"
	"strings"
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
	s0 := e.Scenarios[0]
	if s0.L1Percent == nil || *s0.L1Percent != 100.0 {
		t.Errorf("scenario 0 l1_percent: got %v, want 100.0", s0.L1Percent)
	}
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

// === Test helpers for Summary tests ===

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
	expectedDelta := 92.3 - 88.0
	if math.Abs(s.Delta-expectedDelta) > 0.01 {
		t.Errorf("delta: got %f, want %f", s.Delta, expectedDelta)
	}
}

func TestComputeSummary_PerScenario(t *testing.T) {
	entries := sampleEntries()
	s := ComputeSummary(entries, 10)

	if len(s.ScenarioStats) != 2 {
		t.Fatalf("scenario stats: got %d, want 2", len(s.ScenarioStats))
	}

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

func TestComputeSummary_Empty(t *testing.T) {
	s := ComputeSummary(nil, 10)
	if s.TotalRuns != 0 {
		t.Errorf("total runs: got %d, want 0", s.TotalRuns)
	}
}

func TestRenderMarkdown_ContainsSections(t *testing.T) {
	entries := sampleEntries()
	s := ComputeSummary(entries, 10)
	md := RenderSummaryMarkdown(s)

	sections := []string{
		"## Dashboard",
		"## Historia",
		"## Analiza per-scenariusz",
		"## Trendy kryteriow L2",
		"92.3",
		"PASS",
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

func TestRenderMarkdown_NoDuplicatesSection(t *testing.T) {
	entries := sampleEntries()
	s := ComputeSummary(entries, 10)
	md := RenderSummaryMarkdown(s)

	if strings.Contains(md, "Szum ewaluatora") {
		t.Error("markdown should not have duplicate section when no duplicates")
	}
}

func TestTestReportSummaryCmd_EmptyFile(t *testing.T) {
	dir := t.TempDir()
	os.WriteFile(filepath.Join(dir, "historia.jsonl"), []byte(""), 0644)

	cmd := testReportSummaryCmd()
	cmd.SetArgs([]string{"--historia", filepath.Join(dir, "historia.jsonl")})
	var buf bytes.Buffer
	cmd.SetOut(&buf)

	if err := cmd.Execute(); err != nil {
		t.Fatalf("unexpected error: %v", err)
	}
	output := buf.String()
	if !strings.Contains(output, "Liczba uruchomien**: 0") {
		t.Errorf("expected 'Liczba uruchomien**: 0' in output, got: %s", output)
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
	var result Summary
	if err := json.Unmarshal(buf.Bytes(), &result); err != nil {
		t.Errorf("output is not valid JSON: %v\noutput: %s", err, buf.String())
	}
}
