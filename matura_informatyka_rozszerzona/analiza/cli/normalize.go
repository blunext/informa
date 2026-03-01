package main

import (
	"fmt"
	"math"
	"regexp"
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

	// Normalize Unicode arrows to ASCII
	s = strings.ReplaceAll(s, "→", "->")
	s = strings.ReplaceAll(s, "←", "<-")
	s = strings.ReplaceAll(s, "⟶", "->")

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

// multiPartResult extends checkAnswerResult with part counts.
type multiPartResult struct {
	checkAnswerResult
	TrafioneParts int
	TotalParts    int
}

// multiPartRe matches "a)", "b)", "c)", "d)" labels in answers.
var multiPartRe = regexp.MustCompile(`(?i)\b([a-d])\)\s*`)

// splitMultiPart splits an answer into parts by "a)", "b)", "c)", "d)" labels.
// Returns nil if the answer doesn't have multi-part format.
func splitMultiPart(s string) []string {
	re := multiPartRe
	locs := re.FindAllStringIndex(s, -1)
	if len(locs) < 2 {
		return nil
	}

	var parts []string
	for i, loc := range locs {
		start := loc[1]
		var end int
		if i+1 < len(locs) {
			end = locs[i+1][0]
		} else {
			end = len(s)
		}
		part := strings.TrimSpace(s[start:end])
		parts = append(parts, part)
	}
	return parts
}

// checkMultiPartAnswer handles multi-part answers (a/b/c format).
// Falls back to regular checkAnswer for single-part answers.
func checkMultiPartAnswer(correct, student string) multiPartResult {
	correctParts := splitMultiPart(correct)
	studentParts := splitMultiPart(student)

	if correctParts == nil || studentParts == nil {
		r := checkAnswer(correct, student)
		hit := 0
		if r.Poprawne {
			hit = 1
		}
		return multiPartResult{checkAnswerResult: r, TrafioneParts: hit, TotalParts: 1}
	}

	total := len(correctParts)
	hits := 0
	for i := 0; i < total && i < len(studentParts); i++ {
		r := checkAnswer(correctParts[i], studentParts[i])
		if r.Poprawne {
			hits++
		}
	}

	var wynik string
	switch {
	case hits == total:
		wynik = "pelne"
	case hits == 0:
		wynik = "zero"
	default:
		wynik = "czesciowe"
	}

	return multiPartResult{
		checkAnswerResult: checkAnswerResult{
			Poprawne:          hits == total,
			Wynik:             wynik,
			PoprawnaOdpowiedz: correct,
		},
		TrafioneParts: hits,
		TotalParts:    total,
	}
}
