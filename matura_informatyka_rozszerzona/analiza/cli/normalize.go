package main

import (
	"fmt"
	"math"
	"strconv"
	"strings"
)

// autoScorableTypes are types where CLI can check answers deterministically.
var autoScorableTypes = map[string]bool{
	"sledzenie_algorytmu":           true,
	"test_prawda_falsz":             true,
	"konwersja_systemow_liczbowych": true,
}

func isAutoScorable(typ string) bool {
	return autoScorableTypes[typ]
}

// normalizeAnswer normalizes an answer for comparison:
// trim whitespace, lowercase, normalize numbers (remove trailing .0, leading zeros).
func normalizeAnswer(s string) string {
	s = strings.TrimSpace(s)
	s = strings.ToLower(s)

	// Normalize Polish ł -> l for boolean comparison
	s = strings.ReplaceAll(s, "ł", "l")

	// Try numeric normalization
	if f, err := strconv.ParseFloat(s, 64); err == nil {
		if f == math.Trunc(f) && f >= math.MinInt64 && f <= math.MaxInt64 {
			return fmt.Sprintf("%d", int64(f))
		}
		return strconv.FormatFloat(f, 'f', -1, 64)
	}

	return s
}

// checkAnswerResult is the internal result of checking an answer.
type checkAnswerResult struct {
	Poprawne          bool
	Wynik             string
	PoprawnaOdpowiedz string
}

// checkAnswer compares student answer against correct answer with normalization.
func checkAnswer(correct, student string) checkAnswerResult {
	nc := normalizeAnswer(correct)
	ns := normalizeAnswer(student)

	if nc == ns {
		return checkAnswerResult{Poprawne: true, Wynik: "pelne"}
	}
	return checkAnswerResult{
		Poprawne:          false,
		Wynik:             "zero",
		PoprawnaOdpowiedz: correct,
	}
}
