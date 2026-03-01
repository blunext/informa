package main

import (
	"database/sql"
	"fmt"
	"strings"
	"time"
)

// registerFetch records an exercise fetch in active_exercises.
// Replaces any previous entry for the same exercise_id.
func registerFetch(d *sql.DB, exerciseID, typ string, hintDelay int) error {
	_, err := d.Exec(`
		INSERT OR REPLACE INTO active_exercises
			(exercise_id, typ, fetched_at, attempt_count, hint_delay, hints_fetched, answer_fetched, hints_given)
		VALUES (?, ?, datetime('now'), 0, ?, 0, 0, 0)
	`, exerciseID, typ, hintDelay)
	return err
}

// incrementAttempt increments attempt_count for an exercise.
// Called by progress blad.
func incrementAttempt(d *sql.DB, exerciseID string) error {
	_, err := d.Exec(`
		UPDATE active_exercises SET attempt_count = attempt_count + 1
		WHERE exercise_id = ?
	`, exerciseID)
	return err
}

// clearExercise removes an exercise from active tracking.
// Called by progress update (exercise completed).
func clearExercise(d *sql.DB, exerciseID string) error {
	_, err := d.Exec(`DELETE FROM active_exercises WHERE exercise_id = ?`, exerciseID)
	return err
}

// checkCanFetchAnswer checks if answer can be fetched (attempt_count >= 1).
func checkCanFetchAnswer(d *sql.DB, exerciseID string) (bool, int, error) {
	var attemptCount int
	err := d.QueryRow(`
		SELECT attempt_count FROM active_exercises WHERE exercise_id = ?
	`, exerciseID).Scan(&attemptCount)
	if err == sql.ErrNoRows {
		// Not tracked = allow (backwards compat)
		return true, 0, nil
	}
	if err != nil {
		return false, 0, err
	}
	return attemptCount >= 1, attemptCount, nil
}

// checkCanFetchHints checks if hints can be fetched (attempt_count >= hint_delay).
func checkCanFetchHints(d *sql.DB, exerciseID string) (bool, int, int, error) {
	var attemptCount, hintDelay int
	err := d.QueryRow(`
		SELECT attempt_count, hint_delay FROM active_exercises WHERE exercise_id = ?
	`, exerciseID).Scan(&attemptCount, &hintDelay)
	if err == sql.ErrNoRows {
		return true, 0, 1, nil
	}
	if err != nil {
		return false, 0, 0, err
	}
	return attemptCount >= hintDelay, attemptCount, hintDelay, nil
}

// checkCanFetchHintsSinceAttempt checks if student attempted since last hint.
// First hint is always allowed. After that, requires new attempt.
// Returns (allowed, attempts_since_last_hint, error).
func checkCanFetchHintsSinceAttempt(d *sql.DB, exerciseID string) (bool, int, error) {
	var attemptCount, hintsGiven int
	err := d.QueryRow(`
		SELECT attempt_count, hints_given FROM active_exercises WHERE exercise_id = ?
	`, exerciseID).Scan(&attemptCount, &hintsGiven)
	if err == sql.ErrNoRows {
		return true, 0, nil // Not tracked = allow (backwards compat)
	}
	if err != nil {
		return false, 0, err
	}
	if hintsGiven == 0 {
		return true, attemptCount, nil
	}
	return attemptCount > hintsGiven, attemptCount - hintsGiven, nil
}

// markHintsFetched records that hints were given at current attempt count.
func markHintsFetched(d *sql.DB, exerciseID string) error {
	_, err := d.Exec(`UPDATE active_exercises SET hints_given = attempt_count WHERE exercise_id = ?`, exerciseID)
	return err
}

// runDiagnose returns a lightweight DiagnoseOut (top errors + leech tags).
// Used by auto-diagnose in progressUpdateCmd.
func runDiagnose(d *sql.DB, typ string, limit int) (*DiagnoseOut, error) {
	out := &DiagnoseOut{}

	// Total errors
	totalQuery := "SELECT COUNT(*) FROM progress_bledy"
	var totalParams []any
	if typ != "" {
		totalQuery += " WHERE typ = ?"
		totalParams = append(totalParams, typ)
	}
	d.QueryRow(totalQuery, totalParams...).Scan(&out.Total)

	// Top errors
	query := `SELECT blad_kod, COUNT(*) as cnt,
		GROUP_CONCAT(DISTINCT typ) as typy,
		MAX(data) as ostatnio
		FROM progress_bledy`
	var params []any
	if typ != "" {
		query += " WHERE typ = ?"
		params = append(params, typ)
	}
	query += " GROUP BY blad_kod ORDER BY cnt DESC LIMIT ?"
	params = append(params, limit)

	rows, err := d.Query(query, params...)
	if err != nil {
		return nil, err
	}
	for rows.Next() {
		var entry DiagnoseEntry
		var typyStr string
		if err := rows.Scan(&entry.BladKod, &entry.Count, &typyStr, &entry.Ostatnio); err != nil {
			rows.Close()
			return nil, err
		}
		entry.Typy = strings.Split(typyStr, ",")
		out.TopBledy = append(out.TopBledy, entry)
	}
	rows.Close()
	if err := rows.Err(); err != nil {
		return nil, err
	}
	if out.TopBledy == nil {
		out.TopBledy = []DiagnoseEntry{}
	}

	// Zaleglosci
	today := time.Now().Format("2006-01-02")
	d.QueryRow("SELECT COUNT(*) FROM progress_tagi WHERE nastepna_powtorka <= ?", today).Scan(&out.Zaleglosci)

	// Leech tags
	leechRows, leechErr := d.Query(`SELECT tag, COALESCE(lapses, 0), COALESCE(stability, 0)
		FROM progress_tagi WHERE lapses >= 3 ORDER BY lapses DESC`)
	if leechErr == nil {
		for leechRows.Next() {
			var lt LeechTagOut
			leechRows.Scan(&lt.Tag, &lt.Lapses, &lt.Stability)
			out.LeechTagi = append(out.LeechTagi, lt)
		}
		leechRows.Close()
		if err := leechRows.Err(); err != nil {
			return nil, err
		}
	}
	if out.LeechTagi == nil {
		out.LeechTagi = []LeechTagOut{}
	}

	// Rekomendacja (lightweight — based on available data)
	if len(out.LeechTagi) > 0 {
		lt := out.LeechTagi[0]
		out.Rekomendacja = fmt.Sprintf("Powtórz: %s (lapses: %d, stability: %.1fd)", lt.Tag, lt.Lapses, lt.Stability)
	} else if out.Zaleglosci > 0 {
		out.Rekomendacja = fmt.Sprintf("Masz %d zaległości — użyj 'exercise review'", out.Zaleglosci)
	} else if len(out.TopBledy) > 0 && out.TopBledy[0].Count >= 3 {
		out.Rekomendacja = fmt.Sprintf("Powtarzający się błąd: %s (%dx)", out.TopBledy[0].BladKod, out.TopBledy[0].Count)
	}

	return out, nil
}
