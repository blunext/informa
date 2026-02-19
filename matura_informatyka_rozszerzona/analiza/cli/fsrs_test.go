package main

import (
	"fmt"
	"testing"
	"time"
)

func TestFSRSInitStability(t *testing.T) {
	p := DefaultFSRSParams()
	prev := 0.0
	for _, g := range []int{GradeAgain, GradeHard, GradeGood, GradeEasy} {
		s := p.initStability(g)
		if s <= prev {
			t.Errorf("grade %d: stability %.3f not > prev %.3f", g, s, prev)
		}
		prev = s
	}
}

func TestFSRSInitDifficulty(t *testing.T) {
	p := DefaultFSRSParams()
	prev := 11.0
	for _, g := range []int{GradeAgain, GradeHard, GradeGood, GradeEasy} {
		d := p.initDifficulty(g)
		if d >= prev {
			t.Errorf("grade %d: D %.3f not < prev %.3f", g, d, prev)
		}
		if d < 1 || d > 10 {
			t.Errorf("grade %d: D %.3f out of [1,10]", g, d)
		}
		prev = d
	}
}

func TestFSRSRetrievability(t *testing.T) {
	p := DefaultFSRSParams()
	cases := []struct {
		elapsed   int
		stability float64
		wantMin   float64
		wantMax   float64
	}{
		{0, 10.0, 1.0, 1.0},    // t=0 -> R=1.0
		{10, 10.0, 0.85, 0.95}, // t=S -> R~0.9
		{100, 10.0, 0.0, 0.6},  // t>>S -> R drops significantly
	}
	for _, tc := range cases {
		r := p.Retrievability(tc.elapsed, tc.stability)
		if r < tc.wantMin || r > tc.wantMax {
			t.Errorf("R(%d, %.1f) = %.4f, want [%.2f, %.2f]", tc.elapsed, tc.stability, r, tc.wantMin, tc.wantMax)
		}
	}
}

func TestFSRSNextInterval(t *testing.T) {
	p := DefaultFSRSParams()
	for _, s := range []float64{1, 3, 7, 21, 100} {
		i := p.nextInterval(s)
		ratio := float64(i) / s
		if ratio < 0.8 || ratio > 1.2 {
			t.Errorf("nextInterval(%.1f) = %d, ratio %.2f not near 1.0", s, i, ratio)
		}
	}
}

func TestFSRSNextStateNew(t *testing.T) {
	p := DefaultFSRSParams()
	card := CardState{Difficulty: 5.0}
	c, i := p.NextState(card, GradeEasy, "2026-02-19")
	if c.State != StateReview {
		t.Errorf("state: got %d, want %d", c.State, StateReview)
	}
	if c.Stability <= 0 {
		t.Error("stability should be > 0")
	}
	if i < 1 {
		t.Errorf("interval: got %d, want >= 1", i)
	}
	if c.Lapses != 0 {
		t.Errorf("lapses: got %d, want 0", c.Lapses)
	}
}

func TestFSRSNextStateAgainLapse(t *testing.T) {
	p := DefaultFSRSParams()
	card := CardState{Stability: 10, Difficulty: 5, State: StateReview, LastReview: "2026-02-10"}
	c, i := p.NextState(card, GradeAgain, "2026-02-19")
	if c.State != StateRelearning {
		t.Errorf("state: got %d, want %d", c.State, StateRelearning)
	}
	if c.Lapses != 1 {
		t.Errorf("lapses: got %d, want 1", c.Lapses)
	}
	if i != 1 {
		t.Errorf("interval after lapse: got %d, want 1", i)
	}
	if c.Stability >= card.Stability {
		t.Error("stability should decrease after lapse")
	}
}

func TestFSRSLeechDetection(t *testing.T) {
	p := DefaultFSRSParams()
	card := CardState{Stability: 5, Difficulty: 5, State: StateReview, LastReview: "2026-02-01"}
	for i := 0; i < 3; i++ {
		card, _ = p.NextState(card, GradeAgain, fmt.Sprintf("2026-02-%02d", 10+i))
		if card.State == StateRelearning {
			card, _ = p.NextState(card, GradeGood, fmt.Sprintf("2026-02-%02d", 11+i))
		}
	}
	if card.Lapses < 3 {
		t.Errorf("lapses after 3 Agains: got %d, want >= 3", card.Lapses)
	}
}

func TestFSRSStabilityGrowsMonotonically(t *testing.T) {
	p := DefaultFSRSParams()
	card := CardState{Difficulty: 5.0}
	prevS := 0.0
	date := "2026-01-01"
	for i := 0; i < 10; i++ {
		var interval int
		card, interval = p.NextState(card, GradeGood, date)
		if card.Stability <= prevS {
			t.Errorf("step %d: S=%.3f not > prev %.3f", i, card.Stability, prevS)
		}
		prevS = card.Stability
		d, _ := time.Parse("2006-01-02", date)
		date = d.AddDate(0, 0, interval).Format("2006-01-02")
	}
}

func TestStabilityToPoziom(t *testing.T) {
	cases := []struct {
		s    float64
		want int
	}{
		{0.0, 0}, {0.5, 0}, {1.0, 1}, {2.5, 1},
		{3.0, 2}, {6.0, 2}, {7.0, 3}, {20.0, 3},
		{21.0, 4}, {100.0, 4},
	}
	for _, tc := range cases {
		got := StabilityToPoziom(tc.s)
		if got != tc.want {
			t.Errorf("StabilityToPoziom(%.1f) = %d, want %d", tc.s, got, tc.want)
		}
	}
}

func TestWynikToGrade(t *testing.T) {
	cases := map[string]int{
		"walk_through":        GradeAgain,
		"poprawne_z_pomoca_2": GradeHard,
		"poprawne_z_pomoca_1": GradeGood,
		"poprawne_bez_pomocy": GradeEasy,
		"unknown":             GradeGood,
	}
	for w, want := range cases {
		if got := WynikToGrade(w); got != want {
			t.Errorf("WynikToGrade(%q) = %d, want %d", w, got, want)
		}
	}
}

func TestDaysBetween(t *testing.T) {
	cases := []struct {
		from, to string
		want     int
	}{
		{"2026-02-01", "2026-02-10", 9},
		{"", "2026-02-10", 0},
		{"2026-02-10", "2026-02-01", 0},
		{"2026-02-01", "2026-02-01", 0},
	}
	for _, tc := range cases {
		if got := daysBetween(tc.from, tc.to); got != tc.want {
			t.Errorf("daysBetween(%q, %q) = %d, want %d", tc.from, tc.to, got, tc.want)
		}
	}
}
