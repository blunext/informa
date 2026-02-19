package main

import (
	"math"
	"time"
)

// FSRS-5 state constants
const (
	StateNew        = 0
	StateLearning   = 1
	StateReview     = 2
	StateRelearning = 3
)

// FSRS-5 grade constants (mapped from wynik via WynikToGrade)
const (
	GradeAgain = 1 // walk_through
	GradeHard  = 2 // poprawne_z_pomoca_2
	GradeGood  = 3 // poprawne_z_pomoca_1
	GradeEasy  = 4 // poprawne_bez_pomocy
)

// FSRS-5 fixed constants
const (
	fsrsDecay  = -0.5
	fsrsFactor = 19.0 / 81.0
)

// FSRSParams holds the 19 weights and desired retention.
type FSRSParams struct {
	W                  [19]float64
	RequestedRetention float64
}

// CardState holds per-tag spaced repetition state.
type CardState struct {
	Stability  float64
	Difficulty float64
	Lapses     int
	Reps       int
	State      int
	LastReview string // "2006-01-02" or ""
}

// DefaultFSRSParams returns FSRS-5 default parameters (from open-spaced-repetition).
func DefaultFSRSParams() FSRSParams {
	return FSRSParams{
		W: [19]float64{
			0.40255, 1.18385, 3.173, 15.69105, // w0-w3: initial stability per grade
			7.1949,                              // w4: D0 base
			0.5345,                              // w5: D0 grade scaling
			1.4604,                              // w6: D' grade sensitivity
			0.0046,                              // w7: D' mean reversion weight
			1.54575,                             // w8: S'_r exp base
			0.1192,                              // w9: S'_r stability decay
			1.01925,                             // w10: S'_r retrievability scaling
			1.9395,                              // w11: S'_f base
			0.11,                                // w12: S'_f difficulty power
			0.29605,                             // w13: S'_f stability power
			2.2698,                              // w14: S'_f retrievability scaling
			0.2315,                              // w15: hard penalty
			2.9898,                              // w16: easy bonus
			0.51655,                             // w17: short-term stability
			0.6621,                              // w18: short-term grade offset
		},
		RequestedRetention: 0.9,
	}
}

func (p FSRSParams) initStability(grade int) float64 {
	g := clampInt(grade, 1, 4)
	return math.Max(p.W[g-1], 0.1)
}

func (p FSRSParams) initDifficulty(grade int) float64 {
	d := p.W[4] - math.Exp(p.W[5]*float64(grade-1)) + 1
	return clamp(d, 1, 10)
}

func (p FSRSParams) nextDifficulty(d float64, grade int) float64 {
	d0 := p.initDifficulty(3) // anchor toward "Good" difficulty
	newD := p.W[7]*d0 + (1-p.W[7])*(d - p.W[6]*float64(grade-3))
	return clamp(newD, 1, 10)
}

// Retrievability computes the probability of recall after elapsedDays with given stability.
func (p FSRSParams) Retrievability(elapsedDays int, stability float64) float64 {
	if stability <= 0 || elapsedDays <= 0 {
		return 1.0
	}
	return math.Pow(1+fsrsFactor*float64(elapsedDays)/stability, fsrsDecay)
}

func (p FSRSParams) nextInterval(stability float64) int {
	r := p.RequestedRetention
	i := stability / fsrsFactor * (math.Pow(r, 1.0/fsrsDecay) - 1)
	return max(1, int(math.Round(i)))
}

func (p FSRSParams) nextStabilitySuccess(d, s, r float64, grade int) float64 {
	hardPenalty := 1.0
	if grade == GradeHard {
		hardPenalty = p.W[15]
	}
	easyBonus := 1.0
	if grade == GradeEasy {
		easyBonus = p.W[16]
	}

	newS := s * (math.Exp(p.W[8]) *
		(11 - d) *
		math.Pow(s, -p.W[9]) *
		(math.Exp(p.W[10]*(1-r)) - 1) *
		hardPenalty * easyBonus + 1)
	return math.Max(newS, 0.1)
}

func (p FSRSParams) nextStabilityFail(d, s, r float64) float64 {
	newS := p.W[11] *
		math.Pow(d, -p.W[12]) *
		(math.Pow(s+1, p.W[13]) - 1) *
		math.Exp(p.W[14]*(1-r))
	return clamp(newS, 0.1, s) // fail never increases stability
}

func (p FSRSParams) shortTermStability(s float64, grade int) float64 {
	newS := s * math.Exp(p.W[17]*(float64(grade)-3+p.W[18]))
	return math.Max(newS, 0.1)
}

// NextState computes the new card state and next interval (days) after a review.
func (p FSRSParams) NextState(card CardState, grade int, today string) (CardState, int) {
	elapsed := daysBetween(card.LastReview, today)
	c := card // copy (value type)
	c.Reps++
	c.LastReview = today

	switch card.State {
	case StateNew:
		c.Stability = p.initStability(grade)
		c.Difficulty = p.initDifficulty(grade)
		if grade == GradeAgain {
			c.State = StateLearning
			c.Lapses++
			return c, 1
		}
		if grade <= GradeHard {
			c.State = StateLearning
			return c, max(1, int(math.Round(c.Stability)))
		}
		c.State = StateReview
		return c, max(1, p.nextInterval(c.Stability))

	case StateLearning, StateRelearning:
		c.Stability = p.shortTermStability(card.Stability, grade)
		c.Difficulty = p.nextDifficulty(card.Difficulty, grade)
		if grade == GradeAgain {
			c.Lapses++
			return c, 1
		}
		if grade <= GradeHard {
			return c, max(1, int(math.Round(c.Stability)))
		}
		c.State = StateReview
		return c, max(1, p.nextInterval(c.Stability))

	case StateReview:
		r := p.Retrievability(elapsed, card.Stability)
		c.Difficulty = p.nextDifficulty(card.Difficulty, grade)
		if grade == GradeAgain {
			c.Lapses++
			c.State = StateRelearning
			c.Stability = p.nextStabilityFail(card.Difficulty, card.Stability, r)
			return c, 1
		}
		c.Stability = p.nextStabilitySuccess(card.Difficulty, card.Stability, r, grade)
		return c, max(1, p.nextInterval(c.Stability))
	}

	return c, 1 // unreachable with valid state
}

// WynikToGrade maps exercise result strings to FSRS grades.
func WynikToGrade(wynik string) int {
	switch wynik {
	case "walk_through":
		return GradeAgain
	case "poprawne_z_pomoca_2":
		return GradeHard
	case "poprawne_z_pomoca_1":
		return GradeGood
	case "poprawne_bez_pomocy":
		return GradeEasy
	default:
		return GradeGood
	}
}

// StabilityToPoziom maps FSRS stability to legacy 0-4 poziom for backward compatibility.
func StabilityToPoziom(stability float64) int {
	switch {
	case stability >= 21:
		return 4
	case stability >= 7:
		return 3
	case stability >= 3:
		return 2
	case stability >= 1:
		return 1
	default:
		return 0
	}
}

func daysBetween(from, to string) int {
	if from == "" {
		return 0
	}
	f, err1 := time.Parse("2006-01-02", from)
	t, err2 := time.Parse("2006-01-02", to)
	if err1 != nil || err2 != nil {
		return 0
	}
	days := int(t.Sub(f).Hours() / 24)
	if days < 0 {
		return 0
	}
	return days
}

func clamp(v, lo, hi float64) float64 {
	return math.Max(lo, math.Min(hi, v))
}

func clampInt(v, lo, hi int) int {
	if v < lo {
		return lo
	}
	if v > hi {
		return hi
	}
	return v
}
