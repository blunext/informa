package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"path/filepath"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

func jsonOut(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	if err := enc.Encode(v); err != nil {
		fmt.Fprintf(os.Stderr, "jsonOut encode error: %v\n", err)
		os.Exit(2)
	}
}

// Error types for proper exit codes (1=not found, 2=fatal)
type errNotFound struct{ msg string }

func (e errNotFound) Error() string { return e.msg }

type errFatal struct{ msg string }

func (e errFatal) Error() string { return e.msg }

func notFound(msg string) error { return errNotFound{msg} }
func fatal(msg string) error    { return errFatal{msg} }

// Context key for DB access
type ctxKey struct{}

func db(cmd *cobra.Command) *sql.DB {
	return cmd.Context().Value(ctxKey{}).(*sql.DB)
}

// dbExecer is satisfied by both *sql.DB and *sql.Tx
type dbExecer interface {
	Exec(string, ...any) (sql.Result, error)
	QueryRow(string, ...any) *sql.Row
}

// === helpers ===

func calculateTempo(elapsed, benchmark int) string {
	if benchmark <= 0 {
		return ""
	}
	ratio := float64(elapsed) / float64(benchmark)
	switch {
	case ratio < 0.6:
		return "szybko"
	case ratio <= 1.2:
		return "ok"
	case ratio <= 2.0:
		return "wolno"
	default:
		return "za_wolno"
	}
}

func feedbackCzasowy(tempo string, elapsed, benchmark int) string {
	switch tempo {
	case "szybko":
		return fmt.Sprintf("Świetne tempo! (Ty: %ds, CKE benchmark: ~%ds)", elapsed, benchmark)
	case "wolno":
		return fmt.Sprintf("Na egzaminie miałbyś na to ~%ds — Ty: %ds. Spróbuj szybciej.", benchmark, elapsed)
	case "za_wolno":
		return fmt.Sprintf("To zajęło %ds, CKE benchmark to %ds. Potrenuj szybkość.", elapsed, benchmark)
	default:
		return ""
	}
}

func exerciseTypToCKETypes(typ, kat string) []string {
	switch kat {
	case "ARKUSZ":
		return []string{"arkusz_" + typ}
	case "IMPLEMENTACJA":
		if typ == "geometryczne" {
			return []string{"obliczenia_" + typ}
		}
		return []string{"przetwarzanie_" + typ}
	default:
		return []string{typ} // SQL i TEORIA — bez prefiksu
	}
}

func placeholders(n int) string {
	ph := make([]string, n)
	for i := range ph {
		ph[i] = "?"
	}
	return strings.Join(ph, ",")
}

func toAny(ss []string) []any {
	a := make([]any, len(ss))
	for i, s := range ss {
		a[i] = s
	}
	return a
}

const exerciseColumns = `c.id, c.typ_nazwa, c.kategoria, c.trudnosc, c.punkty, c.zrodlo, c.tagi, c.tresc, c.wskazowki, c.odpowiedz, c.typowe_bledy`

func scanExercise(scanner interface{ Scan(dest ...any) error }) (ExerciseOut, error) {
	var ex ExerciseOut
	var tagiJSON, wskazowkiJSON, typoweBledyJSON string
	err := scanner.Scan(&ex.ID, &ex.TypNazwa, &ex.Kategoria, &ex.Trudnosc, &ex.Punkty, &ex.Zrodlo,
		&tagiJSON, &ex.Tresc, &wskazowkiJSON, &ex.Odpowiedz, &typoweBledyJSON)
	if err != nil {
		return ex, err
	}
	if err := json.Unmarshal([]byte(tagiJSON), &ex.Tagi); err != nil {
		fmt.Fprintf(os.Stderr, "warning: unmarshal tagi for %s: %v\n", ex.ID, err)
	}
	if err := json.Unmarshal([]byte(wskazowkiJSON), &ex.Wskazowki); err != nil {
		fmt.Fprintf(os.Stderr, "warning: unmarshal wskazowki for %s: %v\n", ex.ID, err)
	}
	if err := json.Unmarshal([]byte(typoweBledyJSON), &ex.TypoweBledy); err != nil {
		fmt.Fprintf(os.Stderr, "warning: unmarshal typowe_bledy for %s: %v\n", ex.ID, err)
	}
	if ex.Tagi == nil {
		ex.Tagi = []string{}
	}
	if ex.Wskazowki == nil {
		ex.Wskazowki = []string{}
	}
	if ex.TypoweBledy == nil {
		ex.TypoweBledy = []CommonError{}
	}
	return ex, nil
}

// exerciseToQuestion converts internal ExerciseOut to public QuestionOut (no hints, no answer).
func exerciseToQuestion(ex ExerciseOut) QuestionOut {
	return QuestionOut{
		ID: ex.ID, TypNazwa: ex.TypNazwa, Kategoria: ex.Kategoria,
		Trudnosc: ex.Trudnosc, Punkty: ex.Punkty, Zrodlo: ex.Zrodlo,
		Tagi: ex.Tagi, Tresc: ex.Tresc,
	}
}

// buildCoaching computes student context from progress.db for the given exercise type and tags.
func buildCoaching(d *sql.DB, typ string, exerciseTags []string) Coaching {
	c := Coaching{
		StudentLevel: "new",
		HintDelay:    1,
		LeechTags:    []string{},
		PastMistakes: []string{},
	}

	// Student level from progress_typy + progress_zrobione
	var level sql.NullString
	var streak int
	d.QueryRow("SELECT poziom_trudnosci, streak FROM progress_typy WHERE typ = ?", typ).Scan(&level, &streak)

	if level.Valid && level.String == "trudne" {
		c.StudentLevel = "mastered"
		c.HintDelay = 3
	} else {
		var done int
		d.QueryRow("SELECT COUNT(*) FROM progress_zrobione WHERE typ = ?", typ).Scan(&done)

		switch {
		case done == 0:
			c.StudentLevel = "new"
			c.HintDelay = 1
		case done <= 3 && streak < 3:
			c.StudentLevel = "learning"
			c.HintDelay = 1
		default:
			c.StudentLevel = "familiar"
			c.HintDelay = 2
		}
	}

	// Leech tags: lapses >= 3 AND retrievability < 0.85
	today := time.Now().Format("2006-01-02")
	fsrsParams := DefaultFSRSParams()
	leechRows, err := d.Query(`SELECT tag, COALESCE(stability, 1.0), COALESCE(last_review, '')
		FROM progress_tagi WHERE lapses >= 3`)
	if err == nil {
		defer leechRows.Close()
		for leechRows.Next() {
			var tag, lastReview string
			var stability float64
			leechRows.Scan(&tag, &stability, &lastReview)
			elapsed := daysBetween(lastReview, today)
			r := fsrsParams.Retrievability(elapsed, stability)
			if r < 0.85 {
				c.LeechTags = append(c.LeechTags, tag)
			}
		}
	}

	// Past mistakes: from last 5 sessions, filtered by exercise tags
	if len(exerciseTags) > 0 {
		placeholders := make([]string, len(exerciseTags))
		params := []any{typ}
		for i, tag := range exerciseTags {
			placeholders[i] = "?"
			params = append(params, tag)
		}
		query := fmt.Sprintf(`SELECT blad_opis FROM progress_bledy
			WHERE typ = ? AND blad_kod IN (%s)
			AND data IN (SELECT DISTINCT data FROM progress_zrobione ORDER BY data DESC LIMIT 5)
			GROUP BY blad_opis ORDER BY MAX(data) DESC LIMIT 3`, strings.Join(placeholders, ","))
		rows, err := d.Query(query, params...)
		if err == nil {
			defer rows.Close()
			for rows.Next() {
				var opis string
				rows.Scan(&opis)
				c.PastMistakes = append(c.PastMistakes, opis)
			}
		}
	}

	// Generate coaching_actions
	var actions []string
	for _, tag := range c.LeechTags {
		actions = append(actions, fmt.Sprintf("WARN_LEECH: Tag '%s' sprawia Ci trudnosc — zwroc uwage", tag))
	}
	for _, m := range c.PastMistakes {
		actions = append(actions, fmt.Sprintf("MENTION_PAST: Ostatnio mialeS problem z '%s'", m))
	}
	if c.HintDelay >= 2 {
		actions = append(actions, fmt.Sprintf("HINT_DELAY: %d (Od teraz mniej podpowiedzi — rozwijasz samodzielnosc)", c.HintDelay))
	}
	if actions == nil {
		actions = []string{}
	}
	c.Actions = actions

	return c
}

// getExerciseHints returns hints for an exercise by ID, respecting max_hints for the student's level.
func getExerciseHints(d *sql.DB, id, level string) HintsOut {
	var wskazowkiJSON string
	d.QueryRow("SELECT wskazowki FROM data.cwiczenia WHERE id = ?", id).Scan(&wskazowkiJSON)
	var wskazowki []string
	json.Unmarshal([]byte(wskazowkiJSON), &wskazowki)
	if wskazowki == nil {
		wskazowki = []string{}
	}
	maxHints := calculateMaxHints(level)
	if maxHints >= 0 && len(wskazowki) > maxHints {
		wskazowki = wskazowki[:maxHints]
	}
	if maxHints < 0 {
		maxHints = len(wskazowki)
	}
	out := HintsOut{ID: id, Wskazowki: wskazowki, MaxHints: maxHints}

	// Auto-attach cheatsheet excerpt when returning 2+ hints (student sees L2 content)
	if len(wskazowki) >= 2 {
		out.CheatsheetExcerpt = findCheatsheetExcerptByTags(d, id)
	}

	return out
}

// getExerciseAnswer returns the full answer for an exercise by ID.
func getExerciseAnswer(d *sql.DB, id string) AnswerOut {
	var odpowiedz, typoweBledyJSON string
	d.QueryRow("SELECT odpowiedz, typowe_bledy FROM data.cwiczenia WHERE id = ?", id).Scan(&odpowiedz, &typoweBledyJSON)
	var bledy []CommonError
	json.Unmarshal([]byte(typoweBledyJSON), &bledy)
	if bledy == nil {
		bledy = []CommonError{}
	}
	return AnswerOut{ID: id, Odpowiedz: odpowiedz, TypoweBledy: bledy}
}

func queryExercises(d *sql.DB, typ, trudnosc, exclude string) ([]ExerciseOut, error) {
	query := `SELECT ` + exerciseColumns + ` FROM data.cwiczenia c WHERE c.typ_nazwa = ?`
	params := []any{typ}
	if trudnosc != "" {
		query += " AND c.trudnosc = ?"
		params = append(params, trudnosc)
	}
	query += " AND c.id NOT IN (SELECT id FROM progress_zrobione)"
	if exclude != "" {
		for _, id := range strings.Split(exclude, ",") {
			query += " AND c.id != ?"
			params = append(params, strings.TrimSpace(id))
		}
	}
	rows, err := d.Query(query, params...)
	if err != nil {
		return nil, fatal(fmt.Sprintf("query error: %v", err))
	}
	defer rows.Close()
	var results []ExerciseOut
	for rows.Next() {
		ex, err := scanExercise(rows)
		if err != nil {
			return nil, fatal(fmt.Sprintf("scan error: %v", err))
		}
		results = append(results, ex)
	}
	if err := rows.Err(); err != nil {
		return nil, fatal(fmt.Sprintf("rows error: %v", err))
	}
	return results, nil
}

func getLevel(d *sql.DB, typ string) string {
	var level sql.NullString
	d.QueryRow("SELECT poziom_trudnosci FROM progress_typy WHERE typ = ?", typ).Scan(&level)
	if level.Valid {
		return level.String
	}
	return "latwe"
}

func getStreak(d *sql.DB, typ string) int {
	var streak int
	d.QueryRow("SELECT COALESCE(streak, 0) FROM progress_typy WHERE typ = ?", typ).Scan(&streak)
	return streak
}

func getKategoria(d *sql.DB, typ string) string {
	var kat string
	d.QueryRow("SELECT DISTINCT kategoria FROM data.cwiczenia WHERE typ_nazwa = ? LIMIT 1", typ).Scan(&kat)
	return kat
}

func poolWarning(d *sql.DB, typ, trudnosc string) *string {
	var available int
	d.QueryRow(`SELECT COUNT(*) FROM data.cwiczenia
		WHERE typ_nazwa = ? AND trudnosc = ?
		AND id NOT IN (SELECT id FROM progress_zrobione)`, typ, trudnosc).Scan(&available)
	if available <= 2 {
		msg := fmt.Sprintf("Pozostaly %d cwiczenia typu %s na poziomie %s. Dogeneruj: /generate-exercises %s 5",
			available, typ, trudnosc, typ)
		return &msg
	}
	return nil
}

func getSessionState(d *sql.DB) (count int, lastTyp string, err error) {
	today := time.Now().Format("2006-01-02")
	var sessionDate sql.NullString
	d.QueryRow("SELECT value FROM progress_meta WHERE key = 'session_date'").Scan(&sessionDate)
	if !sessionDate.Valid || sessionDate.String != today {
		if _, err = d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_date', ?)", today); err != nil {
			return 0, "", fmt.Errorf("reset session date: %w", err)
		}
		if _, err = d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_exercise_count', '0')"); err != nil {
			return 0, "", fmt.Errorf("reset session count: %w", err)
		}
		if _, err = d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_last_typ', '')"); err != nil {
			return 0, "", fmt.Errorf("reset session last typ: %w", err)
		}
		return 0, "", nil
	}
	var countStr, lastTypStr sql.NullString
	d.QueryRow("SELECT value FROM progress_meta WHERE key = 'session_exercise_count'").Scan(&countStr)
	d.QueryRow("SELECT value FROM progress_meta WHERE key = 'session_last_typ'").Scan(&lastTypStr)
	if countStr.Valid {
		fmt.Sscanf(countStr.String, "%d", &count)
	}
	if lastTypStr.Valid {
		lastTyp = lastTypStr.String
	}
	return count, lastTyp, nil
}

func incrementSession(d dbExecer, typ string) error {
	today := time.Now().Format("2006-01-02")
	var sessionDate sql.NullString
	d.QueryRow("SELECT value FROM progress_meta WHERE key = 'session_date'").Scan(&sessionDate)
	if !sessionDate.Valid || sessionDate.String != today {
		if _, err := d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_date', ?)", today); err != nil {
			return fmt.Errorf("set session date: %w", err)
		}
		if _, err := d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_exercise_count', '1')"); err != nil {
			return fmt.Errorf("set session count: %w", err)
		}
	} else {
		if _, err := d.Exec(`UPDATE progress_meta SET value = CAST(CAST(value AS INTEGER) + 1 AS TEXT) WHERE key = 'session_exercise_count'`); err != nil {
			return fmt.Errorf("increment session count: %w", err)
		}
	}
	if _, err := d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_last_typ', ?)", typ); err != nil {
		return fmt.Errorf("set session last typ: %w", err)
	}
	return nil
}

// addWeight increments session_context_weight by amount.
// Silently ignores errors (weight tracking is non-critical).
func addWeight(d dbExecer, amount int) {
	if amount <= 0 {
		return
	}
	d.Exec(`INSERT INTO progress_meta (key, value) VALUES ('session_context_weight', ?)
		ON CONFLICT(key) DO UPDATE SET value = CAST(CAST(value AS INTEGER) + ? AS TEXT)`,
		amount, amount)
}

func findExerciseByTag(d *sql.DB, tag string) (ExerciseOut, error) {
	row := d.QueryRow(`SELECT `+exerciseColumns+` FROM data.cwiczenia c
		WHERE EXISTS (SELECT 1 FROM json_each(c.tagi) WHERE value = ?)
		AND c.id NOT IN (SELECT id FROM progress_zrobione)
		ORDER BY RANDOM() LIMIT 1`, tag)
	ex, err := scanExercise(row)
	if err == nil {
		return ex, nil
	}
	row = d.QueryRow(`SELECT `+exerciseColumns+` FROM data.cwiczenia c
		WHERE EXISTS (SELECT 1 FROM json_each(c.tagi) WHERE value = ?)
		ORDER BY RANDOM() LIMIT 1`, tag)
	return scanExercise(row)
}

func findExerciseByTypAndLevel(d *sql.DB, typ, level string) (ExerciseOut, error) {
	query := `SELECT ` + exerciseColumns + ` FROM data.cwiczenia c
		WHERE c.typ_nazwa = ? AND c.id NOT IN (SELECT id FROM progress_zrobione)`
	params := []any{typ}
	if level != "" {
		query += " AND c.trudnosc = ?"
		params = append(params, level)
	}
	query += " ORDER BY RANDOM() LIMIT 1"
	return scanExercise(d.QueryRow(query, params...))
}

func findInterleaveExercise(d *sql.DB, excludeTyp string) (ExerciseOut, error) {
	var altTyp, altLevel string
	err := d.QueryRow(`
		SELECT c.typ_nazwa, COALESCE(p.poziom_trudnosci, 'latwe')
		FROM data.cwiczenia c
		LEFT JOIN progress_typy p ON p.typ = c.typ_nazwa
		WHERE c.typ_nazwa != ?
		AND c.id NOT IN (SELECT id FROM progress_zrobione)
		GROUP BY c.typ_nazwa
		ORDER BY COALESCE(p.streak, 0) ASC, RANDOM()
		LIMIT 1`, excludeTyp).Scan(&altTyp, &altLevel)
	if err != nil {
		return ExerciseOut{}, err
	}
	return findExerciseByTypAndLevel(d, altTyp, altLevel)
}

// === exercise question ===

func exerciseQuestionCmd() *cobra.Command {
	var typ, trudnosc, exclude string

	cmd := &cobra.Command{
		Use:   "question",
		Short: "Get exercise question with coaching (no hints/answer)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if typ == "" {
				return fatal("--typ is required")
			}

			d := db(cmd)

			level := getLevel(d, typ)
			if !cmd.Flags().Changed("trudnosc") {
				trudnosc = level
			}

			results, err := queryExercises(d, typ, trudnosc, exclude)
			if err != nil {
				return err
			}
			if len(results) == 0 && !cmd.Flags().Changed("trudnosc") {
				results, err = queryExercises(d, typ, "", exclude)
				if err != nil {
					return err
				}
			}
			if len(results) == 0 {
				return notFound(fmt.Sprintf("no exercises found for typ=%s trudnosc=%s", typ, trudnosc))
			}

			chosen := results[rand.Intn(len(results))]
			q := exerciseToQuestion(chosen)
			q.Coaching = buildCoaching(d, typ, chosen.Tagi)
			// Register fetch for guardrails
			if err := registerFetch(d, q.ID, q.TypNazwa, q.Coaching.HintDelay); err != nil {
				fmt.Fprintf(os.Stderr, "warning: registerFetch: %v\n", err)
			}
			addWeight(d, 4)
			jsonOut(q)
			return nil
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Exercise type (e.g. sql_group_by, cyfry_liczby)")
	cmd.Flags().StringVar(&trudnosc, "trudnosc", "", "Difficulty: latwe, srednie, srednie-trudne, trudne")
	cmd.Flags().StringVar(&exclude, "exclude", "", "Comma-separated IDs to exclude")
	return cmd
}

// === exercise hints ===

func exerciseHintsCmd() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "hints",
		Short: "Get hints for an exercise by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			if id == "" {
				return fatal("--id is required")
			}
			d := db(cmd)
			var exists int
			if err := d.QueryRow("SELECT COUNT(*) FROM data.cwiczenia WHERE id = ?", id).Scan(&exists); err != nil || exists == 0 {
				return notFound(fmt.Sprintf("exercise %s not found", id))
			}
			// Guardrail: block hints before hint_delay
			canH, attempts, delay, err := checkCanFetchHints(d, id)
			if err != nil {
				return fmt.Errorf("checkCanFetchHints: %w", err)
			}
			if !canH {
				out := map[string]interface{}{
					"status":      "HINT_LOCKED",
					"exercise_id": id,
					"attempt":     attempts,
					"hint_delay":  delay,
					"action":      "Zadaj pytanie sokratejskie BEZ hintow",
				}
				jsonOut(out)
				return nil
			}
			// Mark hints as fetched
			d.Exec(`UPDATE active_exercises SET hints_fetched = 1 WHERE exercise_id = ?`, id)
			var typ string
			d.QueryRow("SELECT typ_nazwa FROM data.cwiczenia WHERE id = ?", id).Scan(&typ)
			level := getLevel(d, typ)
			hints := getExerciseHints(d, id, level)
			if hints.CheatsheetExcerpt != "" {
				addWeight(d, 3)
			} else {
				addWeight(d, 1)
			}
			jsonOut(hints)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Exercise ID (e.g. 7.1)")
	return cmd
}

// === exercise answer ===

func exerciseAnswerCmd() *cobra.Command {
	var id string
	cmd := &cobra.Command{
		Use:   "answer",
		Short: "Get answer for an exercise by ID",
		RunE: func(cmd *cobra.Command, args []string) error {
			if id == "" {
				return fatal("--id is required")
			}
			d := db(cmd)
			var exists int
			if err := d.QueryRow("SELECT COUNT(*) FROM data.cwiczenia WHERE id = ?", id).Scan(&exists); err != nil || exists == 0 {
				return notFound(fmt.Sprintf("exercise %s not found", id))
			}
			// Guardrail: block answer before attempt
			can, attempts, err := checkCanFetchAnswer(d, id)
			if err != nil {
				return fmt.Errorf("checkCanFetchAnswer: %w", err)
			}
			if !can {
				out := map[string]interface{}{
					"status":        "LAZY_LOADING_BLOCKED",
					"exercise_id":   id,
					"attempt_count": attempts,
					"action":        "Student hasn't attempted yet. Record attempt first via 'progress blad' or 'progress update'.",
				}
				jsonOut(out)
				return nil
			}
			// Mark answer as fetched
			d.Exec(`UPDATE active_exercises SET answer_fetched = 1 WHERE exercise_id = ?`, id)
			answer := getExerciseAnswer(d, id)
			addWeight(d, 2)
			jsonOut(answer)
			return nil
		},
	}
	cmd.Flags().StringVar(&id, "id", "", "Exercise ID (e.g. 7.1)")
	return cmd
}

// === exercise review ===

func exerciseReviewCmd() *cobra.Command {
	var limit int

	cmd := &cobra.Command{
		Use:   "review",
		Short: "Get exercises due for spaced repetition review",
		RunE: func(cmd *cobra.Command, args []string) error {
			d := db(cmd)
			today := time.Now().Format("2006-01-02")
			fsrsParams := DefaultFSRSParams()

			rows, err := d.Query(`
				SELECT t.tag, t.nastepna_powtorka,
					COALESCE(t.stability, 1.0), COALESCE(t.last_review, '')
				FROM progress_tagi t
				WHERE t.nastepna_powtorka <= ?
				ORDER BY t.nastepna_powtorka ASC`, today)
			if err != nil {
				return fatal(fmt.Sprintf("query error: %v", err))
			}

			type reviewCandidate struct {
				tag, powtorka, lastReview string
				stability, retrievability float64
			}
			var candidates []reviewCandidate
			for rows.Next() {
				var rc reviewCandidate
				if err := rows.Scan(&rc.tag, &rc.powtorka, &rc.stability, &rc.lastReview); err != nil {
					rows.Close()
					return fatal(fmt.Sprintf("scan error: %v", err))
				}
				elapsed := daysBetween(rc.lastReview, today)
				rc.retrievability = fsrsParams.Retrievability(elapsed, rc.stability)
				candidates = append(candidates, rc)
			}
			if err := rows.Err(); err != nil {
				rows.Close()
				return fatal(fmt.Sprintf("rows error: %v", err))
			}
			rows.Close()

			// Sort by retrievability ASC (lowest = most urgent)
			sort.Slice(candidates, func(i, j int) bool {
				return candidates[i].retrievability < candidates[j].retrievability
			})

			// Limit after sorting
			if len(candidates) > limit {
				candidates = candidates[:limit]
			}

			var results []ReviewOut
			for _, rc := range candidates {
				powtorkaDate, _ := time.Parse("2006-01-02", rc.powtorka)
				daysOverdue := int(time.Since(powtorkaDate).Hours() / 24)

				ex, err := findExerciseByTag(d, rc.tag)
				if err != nil {
					continue
				}

				q := exerciseToQuestion(ex)
				q.Coaching = buildCoaching(d, ex.TypNazwa, ex.Tagi)
				// Add previous_result for reviews
				var prevResult sql.NullString
				d.QueryRow("SELECT wynik FROM progress_zrobione WHERE id = ?", ex.ID).Scan(&prevResult)
				if prevResult.Valid {
					q.Coaching.PreviousResult = prevResult.String
				}

				// Register fetch for guardrails
				if err := registerFetch(d, q.ID, q.TypNazwa, q.Coaching.HintDelay); err != nil {
					fmt.Fprintf(os.Stderr, "warning: registerFetch: %v\n", err)
				}

				results = append(results, ReviewOut{
					Exercise:       q,
					Tag:            rc.tag,
					DaysOverdue:    daysOverdue,
					Retrievability: rc.retrievability,
				})
			}

			if len(results) == 0 {
				return notFound("no reviews due")
			}
			addWeight(d, 4*len(results))
			jsonOut(results)
			return nil
		},
	}

	cmd.Flags().IntVar(&limit, "limit", 3, "Max number of reviews to return")
	return cmd
}

// === progress update ===

func progressUpdateCmd() *cobra.Command {
	var id, wynik string
	var czas int

	cmd := &cobra.Command{
		Use:   "update",
		Short: "Record exercise result and update spaced repetition",
		RunE: func(cmd *cobra.Command, args []string) error {
			if id == "" || wynik == "" {
				return fatal("--id and --wynik are required")
			}

			validWynik := map[string]bool{
				"poprawne_bez_pomocy": true,
				"poprawne_z_pomoca_1": true,
				"poprawne_z_pomoca_2": true,
				"walk_through":        true,
			}
			if !validWynik[wynik] {
				return fatal("--wynik must be one of: poprawne_bez_pomocy, poprawne_z_pomoca_1, poprawne_z_pomoca_2, walk_through")
			}

			d := db(cmd)
			today := time.Now().Format("2006-01-02")

			// Read-only: get exercise info from data DB
			var typNazwa, tagiJSON string
			err := d.QueryRow("SELECT typ_nazwa, tagi FROM data.cwiczenia WHERE id = ?", id).Scan(&typNazwa, &tagiJSON)
			if err != nil {
				return fatal(fmt.Sprintf("exercise %s not found: %v", id, err))
			}

			var tagi []string
			if err := json.Unmarshal([]byte(tagiJSON), &tagi); err != nil {
				return fatal(fmt.Sprintf("invalid tagi JSON for %s: %v", id, err))
			}

			// Read-only: benchmark (for output enrichment)
			var benchmarkSek sql.NullInt64
			if czas > 0 {
				d.QueryRow("SELECT benchmark_sek FROM data.benchmarki WHERE typ = ?", typNazwa).Scan(&benchmarkSek)
			}

			// Begin transaction for all writes
			tx, err := d.Begin()
			if err != nil {
				return fatal(fmt.Sprintf("begin tx: %v", err))
			}
			defer tx.Rollback()

			// Record done exercise
			var czasSekVal *int
			if czas > 0 {
				czasSekVal = &czas
			}
			if _, err = tx.Exec(`INSERT OR REPLACE INTO progress_zrobione (id, typ, data, wynik, czas_sek) VALUES (?, ?, ?, ?, ?)`,
				id, typNazwa, today, wynik, czasSekVal); err != nil {
				return fatal(fmt.Sprintf("save progress: %v", err))
			}

			// Update streak and level
			if _, err = tx.Exec(`INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES (?, 'latwe', 1)
				ON CONFLICT(typ) DO UPDATE SET streak = streak + 1`, typNazwa); err != nil {
				return fatal(fmt.Sprintf("update typ: %v", err))
			}

			// Level progression based on streak
			var streak int
			tx.QueryRow("SELECT streak FROM progress_typy WHERE typ = ?", typNazwa).Scan(&streak)

			newLevel := calculateLevel(streak, wynik)
			oldLevel := calculateLevel(streak-1, wynik)
			bumped := newLevel != oldLevel
			if _, err = tx.Exec("UPDATE progress_typy SET poziom_trudnosci = ? WHERE typ = ?", newLevel, typNazwa); err != nil {
				return fatal(fmt.Sprintf("update level: %v", err))
			}

			// Spaced repetition for tags
			nextReviewDates, err := updateTags(tx, tagi, wynik)
			if err != nil {
				return err
			}

			// Update session metadata
			if _, err = tx.Exec(`INSERT INTO progress_meta (key, value) VALUES ('sesje', '1')
				ON CONFLICT(key) DO UPDATE SET value = CAST(CAST(value AS INTEGER) + 1 AS TEXT)
				WHERE key = 'sesje' AND NOT EXISTS (
					SELECT 1 FROM progress_meta WHERE key = 'ostatnia_sesja' AND value = ?
				)`, today); err != nil {
				return fatal(fmt.Sprintf("update sesje: %v", err))
			}
			if _, err = tx.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('ostatnia_sesja', ?)`, today); err != nil {
				return fatal(fmt.Sprintf("update ostatnia_sesja: %v", err))
			}
			if _, err = tx.Exec(`INSERT INTO progress_meta (key, value) VALUES ('cwiczenia_lacznie', '1')
				ON CONFLICT(key) DO UPDATE SET value = CAST(CAST(value AS INTEGER) + 1 AS TEXT)`); err != nil {
				return fatal(fmt.Sprintf("update cwiczenia_lacznie: %v", err))
			}

			// Session tracking
			if err := incrementSession(tx, typNazwa); err != nil {
				return fatal(fmt.Sprintf("session tracking: %v", err))
			}

			if err := tx.Commit(); err != nil {
				return fatal(fmt.Sprintf("commit: %v", err))
			}

			// Clear exercise from active tracking
			if err := clearExercise(d, id); err != nil {
				fmt.Fprintf(os.Stderr, "warning: clearExercise: %v\n", err)
			}

			out := ProgressUpdateOut{
				ID:               id,
				NewLevel:         newLevel,
				Streak:           streak,
				DifficultyBumped: bumped,
				TagsUpdated:      tagi,
				NextReviewDates:  nextReviewDates,
			}

			// FSRS stats from updated tags
			var totalS float64
			var maxLapses int
			for _, tag := range tagi {
				var s float64
				var l int
				d.QueryRow("SELECT COALESCE(stability, 0), COALESCE(lapses, 0) FROM progress_tagi WHERE tag = ?", tag).Scan(&s, &l)
				totalS += s
				if l > maxLapses {
					maxLapses = l
				}
			}
			if len(tagi) > 0 {
				avgS := totalS / float64(len(tagi))
				out.Stability = &avgS
			}
			if maxLapses > 0 {
				out.Lapses = &maxLapses
			}

			// Time tracking
			if czas > 0 {
				out.CzasSek = &czas
				if benchmarkSek.Valid {
					b := int(benchmarkSek.Int64)
					out.BenchmarkSek = &b
					tempo := calculateTempo(czas, b)
					if tempo != "" {
						out.Tempo = &tempo
						fb := feedbackCzasowy(tempo, czas, b)
						if fb != "" {
							out.FeedbackCzasowy = &fb
						}
					}
				}
			}

			// Warn if non-perfect result but no errors logged
			if wynik != "poprawne_bez_pomocy" {
				var bladyCount int
				d.QueryRow("SELECT COUNT(*) FROM progress_bledy WHERE exercise_id = ? AND data = ?",
					id, today).Scan(&bladyCount)
				if bladyCount == 0 {
					w := "UWAGA: wynik " + wynik + " ale brak progress blad dla tego cwiczenia"
					out.BladWarning = &w
				}
			}

			addWeight(d, 1)
			if wynik == "walk_through" {
				addWeight(d, 5)
			}

			// Auto-diagnose every 5 exercises
			var totalDoneStr string
			d.QueryRow(`SELECT COALESCE(value, '0') FROM progress_meta WHERE key='cwiczenia_lacznie'`).Scan(&totalDoneStr)
			var totalDone int
			fmt.Sscanf(totalDoneStr, "%d", &totalDone)
			if totalDone > 0 && totalDone%5 == 0 {
				diag, err := runDiagnose(d, typNazwa, 3)
				if err == nil {
					out.AutoDiagnose = diag
				}
			}

			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Exercise ID (e.g. 7.3)")
	cmd.Flags().StringVar(&wynik, "wynik", "", "Result: poprawne_bez_pomocy, poprawne_z_pomoca_1, poprawne_z_pomoca_2, walk_through")
	cmd.Flags().IntVar(&czas, "czas", 0, "Time spent in seconds")
	return cmd
}

func calculateLevel(streak int, wynik string) string {
	if wynik == "walk_through" {
		return "latwe"
	}
	switch {
	case streak >= 8:
		return "trudne"
	case streak >= 5:
		return "srednie-trudne"
	case streak >= 3:
		return "srednie"
	default:
		return "latwe"
	}
}

// calculateMaxHints returns the max number of hints based on difficulty level.
// -1 means no limit (all hints). Based on level (which resets on walk_through).
func calculateMaxHints(level string) int {
	switch level {
	case "srednie-trudne":
		return 1
	case "trudne":
		return 0
	default: // latwe, srednie
		return -1
	}
}

// applyMaxHints truncates wskazowki and sets MaxHints on the exercise.
// maxHintsFlag: -1 = auto (use level), >= 0 = explicit override.
func applyMaxHints(ex *ExerciseOut, level string, maxHintsFlag int) {
	effective := maxHintsFlag
	if effective < 0 {
		effective = calculateMaxHints(level)
	}
	ex.MaxHints = effective
	if effective >= 0 && len(ex.Wskazowki) > effective {
		ex.Wskazowki = ex.Wskazowki[:effective]
	}
	if effective < 0 {
		ex.MaxHints = len(ex.Wskazowki) // report actual count when unlimited
	}
}

func updateTags(d dbExecer, tags []string, wynik string) ([]string, error) {
	grade := WynikToGrade(wynik)
	params := DefaultFSRSParams()
	today := time.Now().Format("2006-01-02")
	var dates []string

	for _, tag := range tags {
		var card CardState
		err := d.QueryRow(`SELECT COALESCE(stability, 0), COALESCE(difficulty, 5.0),
			COALESCE(lapses, 0), COALESCE(reps, 0), COALESCE(state, 0),
			COALESCE(last_review, '')
			FROM progress_tagi WHERE tag = ?`, tag).
			Scan(&card.Stability, &card.Difficulty, &card.Lapses,
				&card.Reps, &card.State, &card.LastReview)
		if err != nil {
			card = CardState{Difficulty: 5.0} // new tag
		}

		newCard, interval := params.NextState(card, grade, today)
		newPoziom := StabilityToPoziom(newCard.Stability)
		nextReview := time.Now().AddDate(0, 0, interval).Format("2006-01-02")

		if _, err := d.Exec(`INSERT INTO progress_tagi
			(tag, poziom, nastepna_powtorka, stability, difficulty, lapses, reps, state, last_review)
			VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?)
			ON CONFLICT(tag) DO UPDATE SET
				poziom=?, nastepna_powtorka=?, stability=?, difficulty=?,
				lapses=?, reps=?, state=?, last_review=?`,
			tag, newPoziom, nextReview, newCard.Stability, newCard.Difficulty,
			newCard.Lapses, newCard.Reps, newCard.State, today,
			newPoziom, nextReview, newCard.Stability, newCard.Difficulty,
			newCard.Lapses, newCard.Reps, newCard.State, today); err != nil {
			return nil, fatal(fmt.Sprintf("update tag %s: %v", tag, err))
		}
		dates = append(dates, nextReview)
	}
	return dates, nil
}

// === progress status ===

func progressStatusCmd() *cobra.Command {
	var typ string

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Show progress dashboard",
		RunE: func(cmd *cobra.Command, args []string) error {
			d := db(cmd)
			out := ProgressStatusOut{}

			// Auto-reset session context weight — progress status is the session-start signal
			d.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '0')`)

			// Meta
			var sesjeStr, ostatnia, cwiczeniaStr sql.NullString
			d.QueryRow("SELECT value FROM progress_meta WHERE key = 'sesje'").Scan(&sesjeStr)
			d.QueryRow("SELECT value FROM progress_meta WHERE key = 'ostatnia_sesja'").Scan(&ostatnia)
			d.QueryRow("SELECT value FROM progress_meta WHERE key = 'cwiczenia_lacznie'").Scan(&cwiczeniaStr)

			if sesjeStr.Valid {
				fmt.Sscanf(sesjeStr.String, "%d", &out.Sesje)
			}
			if ostatnia.Valid {
				out.OstatniaSesja = ostatnia.String
			}
			if cwiczeniaStr.Valid {
				fmt.Sscanf(cwiczeniaStr.String, "%d", &out.CwiczeniaLacznie)
			}

			// Per-typ stats
			query := `
				SELECT
					c.typ_nazwa,
					COALESCE(p.poziom_trudnosci, 'latwe') as poziom,
					COALESCE(p.streak, 0) as streak,
					COUNT(DISTINCT d.id) as zrobione,
					COUNT(DISTINCT c.id) as dostepne
				FROM data.cwiczenia c
				LEFT JOIN progress_typy p ON p.typ = c.typ_nazwa
				LEFT JOIN progress_zrobione d ON d.id = c.id`

			if typ != "" {
				query += " WHERE c.typ_nazwa = ?"
			}
			query += " GROUP BY c.typ_nazwa ORDER BY c.typ_nazwa"

			var rows *sql.Rows
			var err error
			if typ != "" {
				rows, err = d.Query(query, typ)
			} else {
				rows, err = d.Query(query)
			}
			if err != nil {
				return fatal(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			for rows.Next() {
				var ts TypStatusOut
				if err := rows.Scan(&ts.Typ, &ts.PoziomTrudnosci, &ts.Streak, &ts.Zrobione, &ts.Dostepne); err != nil {
					return fatal(fmt.Sprintf("scan error: %v", err))
				}
				out.PerTyp = append(out.PerTyp, ts)
			}
			if err := rows.Err(); err != nil {
				return fatal(fmt.Sprintf("rows error: %v", err))
			}
			if out.PerTyp == nil {
				out.PerTyp = []TypStatusOut{}
			}

			// Enrich per-typ with avg_czas_sek and benchmark
			avgCzasMap := map[string]int{}
			avgRows, err := d.Query(`SELECT typ, CAST(AVG(czas_sek) AS INTEGER) FROM progress_zrobione WHERE czas_sek IS NOT NULL GROUP BY typ`)
			if err == nil {
				for avgRows.Next() {
					var t string
					var avg int
					avgRows.Scan(&t, &avg)
					avgCzasMap[t] = avg
				}
				if rowErr := avgRows.Err(); rowErr != nil {
					fmt.Fprintf(os.Stderr, "warning: avg czas rows: %v\n", rowErr)
				}
				avgRows.Close()
			}

			for i := range out.PerTyp {
				if avg, ok := avgCzasMap[out.PerTyp[i].Typ]; ok {
					out.PerTyp[i].AvgCzasSek = &avg
				}
				var benchSek sql.NullInt64
				d.QueryRow("SELECT benchmark_sek FROM data.benchmarki WHERE typ = ?", out.PerTyp[i].Typ).Scan(&benchSek)
				if benchSek.Valid {
					b := int(benchSek.Int64)
					out.PerTyp[i].BenchmarkSek = &b
				}
			}

			// Per-kategoria aggregation
			typKatMap := map[string]string{}
			katRows, err := d.Query("SELECT DISTINCT typ_nazwa, kategoria FROM data.cwiczenia")
			if err == nil {
				for katRows.Next() {
					var t, k string
					katRows.Scan(&t, &k)
					typKatMap[t] = k
				}
				if rowErr := katRows.Err(); rowErr != nil {
					fmt.Fprintf(os.Stderr, "warning: kategoria rows: %v\n", rowErr)
				}
				katRows.Close()
			}

			katAgg := map[string]*KategoriaStatusOut{}
			for _, ts := range out.PerTyp {
				kat := typKatMap[ts.Typ]
				if kat == "" {
					continue
				}
				if katAgg[kat] == nil {
					katAgg[kat] = &KategoriaStatusOut{Kategoria: kat}
				}
				katAgg[kat].TypyTotal++
				if ts.Zrobione > 0 {
					katAgg[kat].TypyRuszane++
				}
				katAgg[kat].Zrobione += ts.Zrobione
				katAgg[kat].Dostepne += ts.Dostepne
				katAgg[kat].AvgStreak += float64(ts.Streak)
			}
			katOrder := []string{"TEORIA", "IMPLEMENTACJA", "ARKUSZ", "SQL"}
			for _, k := range katOrder {
				if ks, ok := katAgg[k]; ok {
					if ks.TypyTotal > 0 {
						ks.AvgStreak /= float64(ks.TypyTotal)
					}
					out.PerKategoria = append(out.PerKategoria, *ks)
				}
			}
			if out.PerKategoria == nil {
				out.PerKategoria = []KategoriaStatusOut{}
			}

			// Overdue reviews
			today := time.Now().Format("2006-01-02")
			d.QueryRow("SELECT COUNT(*) FROM progress_tagi WHERE nastepna_powtorka <= ?", today).Scan(&out.Zaleglosci)

			// Mastered tags (poziom >= 3)
			tagRows, err := d.Query("SELECT tag FROM progress_tagi WHERE poziom >= 3 ORDER BY tag")
			if err == nil {
				for tagRows.Next() {
					var tag string
					tagRows.Scan(&tag)
					out.TagiOpanowane = append(out.TagiOpanowane, tag)
				}
				if rowErr := tagRows.Err(); rowErr != nil {
					fmt.Fprintf(os.Stderr, "warning: tagi opanowane rows: %v\n", rowErr)
				}
				tagRows.Close()
			}
			if out.TagiOpanowane == nil {
				out.TagiOpanowane = []string{}
			}

			// Problematic tags (poziom <= 1 AND has reviews)
			probRows, err := d.Query("SELECT tag FROM progress_tagi WHERE poziom <= 1 AND nastepna_powtorka IS NOT NULL ORDER BY tag")
			if err == nil {
				for probRows.Next() {
					var tag string
					probRows.Scan(&tag)
					out.TagiProblematyczne = append(out.TagiProblematyczne, tag)
				}
				if rowErr := probRows.Err(); rowErr != nil {
					fmt.Fprintf(os.Stderr, "warning: tagi problematyczne rows: %v\n", rowErr)
				}
				probRows.Close()
			}
			if out.TagiProblematyczne == nil {
				out.TagiProblematyczne = []string{}
			}

			// Leech tags (3+ lapses)
			leechRows, leechErr := d.Query(`SELECT tag, COALESCE(lapses, 0), COALESCE(stability, 0)
				FROM progress_tagi WHERE lapses >= 3 ORDER BY lapses DESC`)
			if leechErr == nil {
				for leechRows.Next() {
					var lt LeechTagOut
					leechRows.Scan(&lt.Tag, &lt.Lapses, &lt.Stability)
					out.LeechTagi = append(out.LeechTagi, lt)
				}
				if rowErr := leechRows.Err(); rowErr != nil {
					fmt.Fprintf(os.Stderr, "warning: leech rows: %v\n", rowErr)
				}
				leechRows.Close()
			}
			if out.LeechTagi == nil {
				out.LeechTagi = []LeechTagOut{}
			}

			// Global retencja_szacowana — average retrievability across all tracked tags
			fsrsParams := DefaultFSRSParams()
			retRows, retErr := d.Query(`SELECT tag, COALESCE(stability, 1.0), COALESCE(last_review, '')
				FROM progress_tagi WHERE nastepna_powtorka IS NOT NULL`)
			if retErr == nil {
				var totalR float64
				var retCount int
				for retRows.Next() {
					var tag, lastReview string
					var stability float64
					retRows.Scan(&tag, &stability, &lastReview)
					elapsed := daysBetween(lastReview, today)
					r := fsrsParams.Retrievability(elapsed, stability)
					totalR += r
					retCount++
				}
				if rowErr := retRows.Err(); rowErr != nil {
					fmt.Fprintf(os.Stderr, "warning: retencja rows: %v\n", rowErr)
				}
				retRows.Close()
				if retCount > 0 {
					avgR := totalR / float64(retCount)
					out.RetencjaSzacowana = &avgR
				}
			}

			// Per-kategoria retencja — tag → exercise → category mapping
			katRetRows, katRetErr := d.Query(`
				SELECT c.kategoria, t.tag, COALESCE(t.stability, 1.0), COALESCE(t.last_review, '')
				FROM progress_tagi t
				JOIN data.cwiczenia c ON EXISTS (SELECT 1 FROM json_each(c.tagi) WHERE value = t.tag)
				WHERE t.nastepna_powtorka IS NOT NULL
				GROUP BY c.kategoria, t.tag`)
			if katRetErr == nil {
				type katRetAcc struct {
					totalR float64
					count  int
				}
				katRetMap := map[string]*katRetAcc{}
				for katRetRows.Next() {
					var kat, tag, lastReview string
					var stability float64
					katRetRows.Scan(&kat, &tag, &stability, &lastReview)
					elapsed := daysBetween(lastReview, today)
					r := fsrsParams.Retrievability(elapsed, stability)
					if katRetMap[kat] == nil {
						katRetMap[kat] = &katRetAcc{}
					}
					katRetMap[kat].totalR += r
					katRetMap[kat].count++
				}
				if rowErr := katRetRows.Err(); rowErr != nil {
					fmt.Fprintf(os.Stderr, "warning: kat retencja rows: %v\n", rowErr)
				}
				katRetRows.Close()
				for i := range out.PerKategoria {
					if acc, ok := katRetMap[out.PerKategoria[i].Kategoria]; ok && acc.count > 0 {
						avgR := acc.totalR / float64(acc.count)
						out.PerKategoria[i].Retencja = &avgR
					}
				}
			}

			// Rekomendacja — priority-based single recommendation
			out.Rekomendacja = computeRekomendacja(d, out)

			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Filter by type (e.g. sql_group_by)")
	return cmd
}

// tier1Types are types that appear in 100% of exams — critical for every student.
var tier1Types = []string{
	"sledzenie_algorytmu", "cyfry_liczby", "napisy",
	"agregacja_warunkowa", "symulacja",
	"sql_group_by", "sql_join",
}

// computeRekomendacja returns a single priority recommendation based on progress status.
func computeRekomendacja(_ *sql.DB, out ProgressStatusOut) string {
	// (a) Leech tag — highest priority
	if len(out.LeechTagi) > 0 {
		lt := out.LeechTagi[0]
		return fmt.Sprintf("Powtórz: %s (lapses: %d, stability: %.1fd)", lt.Tag, lt.Lapses, lt.Stability)
	}

	// (b) Category with retencja < 0.80
	for _, ks := range out.PerKategoria {
		if ks.Retencja != nil && *ks.Retencja < 0.80 {
			return fmt.Sprintf("Kategoria %s ma retencję %.0f%% — wymaga powtórek", ks.Kategoria, *ks.Retencja*100)
		}
	}

	// (c) Untouched TIER 1 type
	touchedTypes := map[string]bool{}
	for _, ts := range out.PerTyp {
		if ts.Zrobione > 0 {
			touchedTypes[ts.Typ] = true
		}
	}
	for _, t1 := range tier1Types {
		if !touchedTypes[t1] {
			return fmt.Sprintf("Nie ćwiczyłeś: %s — pojawia się co roku na maturze", t1)
		}
	}

	// (d) Streak imbalance: max avg_streak / min avg_streak >= 3
	if len(out.PerKategoria) >= 2 {
		var maxStreak, minStreak float64
		var maxKat, minKat string
		first := true
		for _, ks := range out.PerKategoria {
			if ks.TypyRuszane == 0 {
				continue
			}
			if first {
				maxStreak = ks.AvgStreak
				minStreak = ks.AvgStreak
				maxKat = ks.Kategoria
				minKat = ks.Kategoria
				first = false
			} else {
				if ks.AvgStreak > maxStreak {
					maxStreak = ks.AvgStreak
					maxKat = ks.Kategoria
				}
				if ks.AvgStreak < minStreak {
					minStreak = ks.AvgStreak
					minKat = ks.Kategoria
				}
			}
		}
		if !first && minStreak > 0 && maxStreak/minStreak >= 3 {
			return fmt.Sprintf("%s avg streak %.0f, %s avg streak %.0f — wyrównaj", maxKat, maxStreak, minKat, minStreak)
		}
	}

	return ""
}

// === progress blad ===

func progressBladCmd() *cobra.Command {
	var exerciseID, typ, kod, opis string
	var hint int

	cmd := &cobra.Command{
		Use:   "blad",
		Short: "Record an error/mistake for diagnosis",
		RunE: func(cmd *cobra.Command, args []string) error {
			if exerciseID == "" || typ == "" || kod == "" {
				return fatal("--exercise-id, --typ, and --kod are required")
			}

			// Require --hint
			if hint < 0 {
				return fatal("--hint wymagany. Uzyj --hint 0 (przed hintem) lub --hint 1/2/3 (po hincie)")
			}

			// Validate error code against whitelist
			valid, _ := validateErrorCode(typ, kod)
			if !valid {
				suggestions := suggestClosestCodes(typ, kod)
				_, allowedCodes := validateErrorCode(typ, "___force_invalid___")
				type codeEntry struct {
					Kod  string `json:"kod"`
					Opis string `json:"opis"`
				}
				validList := make([]codeEntry, len(allowedCodes))
				for i, c := range allowedCodes {
					validList[i] = codeEntry{c, errorCodeDescriptions[c]}
				}
				jsonOut(map[string]any{
					"error":       fmt.Sprintf("kod '%s' niedostepny dla typ '%s'", kod, typ),
					"valid_codes": validList,
					"suggestions": suggestions,
				})
				return fatal(fmt.Sprintf("kod '%s' niedostepny", kod))
			}

			d := db(cmd)

			// Validate exercise exists
			var exists int
			d.QueryRow("SELECT COUNT(*) FROM data.cwiczenia WHERE id = ?", exerciseID).Scan(&exists)
			if exists == 0 {
				return fatal(fmt.Sprintf("exercise %s not found", exerciseID))
			}

			today := time.Now().Format("2006-01-02")

			result, err := d.Exec(
				`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, hint_level, data) VALUES (?, ?, ?, ?, ?, ?)`,
				exerciseID, typ, kod, opis, hint, today)
			if err != nil {
				return fatal(fmt.Sprintf("save error: %v", err))
			}

			lastID, _ := result.LastInsertId()

			// Track attempt for guardrails
			if err := incrementAttempt(d, exerciseID); err != nil {
				fmt.Fprintf(os.Stderr, "warning: incrementAttempt: %v\n", err)
			}

			addWeight(d, 1)
			if hint == 3 {
				addWeight(d, 2)
			}

			jsonOut(BladOut{
				ID:         int(lastID),
				ExerciseID: exerciseID,
				Typ:        typ,
				BladKod:    kod,
				BladOpis:   opis,
				HintLevel:  hint,
				Data:       today,
			})
			return nil
		},
	}

	cmd.Flags().StringVar(&exerciseID, "exercise-id", "", "Exercise ID (e.g. 7.3)")
	cmd.Flags().StringVar(&typ, "typ", "", "Exercise type")
	cmd.Flags().StringVar(&kod, "kod", "", "Short error code (e.g. brak_group_by)")
	cmd.Flags().StringVar(&opis, "opis", "", "Full error description")
	cmd.Flags().IntVar(&hint, "hint", -1, "Hint level (0=before any hint, 1/2/3=after hint) — REQUIRED")
	return cmd
}

// === progress diagnose ===

func progressDiagnoseCmd() *cobra.Command {
	var typ string
	var limit int

	cmd := &cobra.Command{
		Use:   "diagnose",
		Short: "Analyze recurring errors",
		RunE: func(cmd *cobra.Command, args []string) error {
			d := db(cmd)

			// Get total count first
			totalQuery := "SELECT COUNT(*) FROM progress_bledy"
			var totalParams []any
			if typ != "" {
				totalQuery += " WHERE typ = ?"
				totalParams = append(totalParams, typ)
			}
			out := DiagnoseOut{}
			d.QueryRow(totalQuery, totalParams...).Scan(&out.Total)

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
				return fatal(fmt.Sprintf("query error: %v", err))
			}

			for rows.Next() {
				var entry DiagnoseEntry
				var typyStr string
				if err := rows.Scan(&entry.BladKod, &entry.Count, &typyStr, &entry.Ostatnio); err != nil {
					rows.Close()
					return fatal(fmt.Sprintf("scan error: %v", err))
				}
				entry.Typy = strings.Split(typyStr, ",")
				out.TopBledy = append(out.TopBledy, entry)
			}
			if err := rows.Err(); err != nil {
				rows.Close()
				return fatal(fmt.Sprintf("rows error: %v", err))
			}
			rows.Close()

			if out.TopBledy == nil {
				out.TopBledy = []DiagnoseEntry{}
			}

			// --- Status fields (mirror progress status without weight reset) ---

			// (a) Zaleglosci — overdue review count
			today := time.Now().Format("2006-01-02")
			d.QueryRow("SELECT COUNT(*) FROM progress_tagi WHERE nastepna_powtorka <= ?", today).Scan(&out.Zaleglosci)

			// (b) LeechTagi — tags with 3+ lapses
			leechRows, leechErr := d.Query(`SELECT tag, COALESCE(lapses, 0), COALESCE(stability, 0)
				FROM progress_tagi WHERE lapses >= 3 ORDER BY lapses DESC`)
			if leechErr == nil {
				for leechRows.Next() {
					var lt LeechTagOut
					leechRows.Scan(&lt.Tag, &lt.Lapses, &lt.Stability)
					out.LeechTagi = append(out.LeechTagi, lt)
				}
				leechRows.Close()
			}
			if out.LeechTagi == nil {
				out.LeechTagi = []LeechTagOut{}
			}

			// (c) RetencjaSzacowana — average retrievability
			fsrsParams := DefaultFSRSParams()
			retRows, retErr := d.Query(`SELECT tag, COALESCE(stability, 1.0), COALESCE(last_review, '')
				FROM progress_tagi WHERE nastepna_powtorka IS NOT NULL`)
			if retErr == nil {
				var totalR float64
				var retCount int
				for retRows.Next() {
					var tag, lastReview string
					var stability float64
					retRows.Scan(&tag, &stability, &lastReview)
					elapsed := daysBetween(lastReview, today)
					r := fsrsParams.Retrievability(elapsed, stability)
					totalR += r
					retCount++
				}
				retRows.Close()
				if retCount > 0 {
					avgR := totalR / float64(retCount)
					out.RetencjaSzacowana = &avgR
				}
			}

			// (d) Rekomendacja — build minimal ProgressStatusOut for computeRekomendacja
			statusOut := ProgressStatusOut{
				Zaleglosci: out.Zaleglosci,
				LeechTagi:  out.LeechTagi,
			}
			if out.RetencjaSzacowana != nil {
				statusOut.RetencjaSzacowana = out.RetencjaSzacowana
			}

			// Build typ→kategoria map first (separate query, must complete before typRows)
			typKatMap := map[string]string{}
			katQueryRows, katQueryErr := d.Query("SELECT DISTINCT typ_nazwa, kategoria FROM data.cwiczenia")
			if katQueryErr == nil {
				for katQueryRows.Next() {
					var t, k string
					katQueryRows.Scan(&t, &k)
					typKatMap[t] = k
				}
				katQueryRows.Close()
			}

			typRows, typErr := d.Query(`
				SELECT c.typ_nazwa,
					COALESCE(p.poziom_trudnosci, 'latwe') as poziom,
					COALESCE(p.streak, 0) as streak,
					COUNT(DISTINCT z.id) as zrobione,
					COUNT(DISTINCT c.id) as dostepne
				FROM data.cwiczenia c
				LEFT JOIN progress_typy p ON p.typ = c.typ_nazwa
				LEFT JOIN progress_zrobione z ON z.id = c.id
				GROUP BY c.typ_nazwa ORDER BY c.typ_nazwa`)
			if typErr == nil {
				katAgg := map[string]*KategoriaStatusOut{}
				for typRows.Next() {
					var ts TypStatusOut
					typRows.Scan(&ts.Typ, &ts.PoziomTrudnosci, &ts.Streak, &ts.Zrobione, &ts.Dostepne)
					statusOut.PerTyp = append(statusOut.PerTyp, ts)
					kat := typKatMap[ts.Typ]
					if kat == "" {
						continue
					}
					if katAgg[kat] == nil {
						katAgg[kat] = &KategoriaStatusOut{Kategoria: kat}
					}
					katAgg[kat].TypyTotal++
					if ts.Zrobione > 0 {
						katAgg[kat].TypyRuszane++
					}
					katAgg[kat].Zrobione += ts.Zrobione
					katAgg[kat].Dostepne += ts.Dostepne
					katAgg[kat].AvgStreak += float64(ts.Streak)
				}
				typRows.Close()
				for _, k := range []string{"TEORIA", "IMPLEMENTACJA", "ARKUSZ", "SQL"} {
					if ks, ok := katAgg[k]; ok {
						if ks.TypyTotal > 0 {
							ks.AvgStreak /= float64(ks.TypyTotal)
						}
						statusOut.PerKategoria = append(statusOut.PerKategoria, *ks)
					}
				}
			}
			out.Rekomendacja = computeRekomendacja(d, statusOut)

			addWeight(d, 2)
			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Filter by type")
	cmd.Flags().IntVar(&limit, "limit", 5, "Max number of error codes to return")
	return cmd
}

// === cke get ===

func ckeGetCmd() *cobra.Command {
	var typ, exclude string
	var force bool

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a random CKE subtask",
		RunE: func(cmd *cobra.Command, args []string) error {
			if typ == "" {
				return fatal("--typ is required")
			}

			d := db(cmd)

			// Unlock check: require level "trudne" unless --force
			if !force {
				level := getLevel(d, typ)
				if level != "trudne" {
					return notFound(fmt.Sprintf("Sprawdzian typu %s wymaga poziomu trudne. Twoj poziom: %s. Uzyj --force aby pominac.", typ, level))
				}
			}

			query := `SELECT e.id, e.rok, e.numer_zadania, e.tytul, e.kontekst, e.typ_zadania, e.kategoria, e.punkty, e.tresc, e.odpowiedz, e.zasady_oceniania, e.pulapki, e.sciezka_danych, e.pliki_danych
				FROM data.egzamin e
				WHERE e.typ_zadania = ?
				AND e.id NOT IN (SELECT id FROM matura_zrobione)
				AND e.id NOT IN (SELECT id FROM worked_examples_shown)`
			params := []any{typ}

			if exclude != "" {
				ids := strings.Split(exclude, ",")
				ph := make([]string, len(ids))
				for i, id := range ids {
					ph[i] = "?"
					params = append(params, strings.TrimSpace(id))
				}
				query += " AND e.id NOT IN (" + strings.Join(ph, ",") + ")"
			}

			query += " ORDER BY RANDOM() LIMIT 1"

			var out CKEOut
			var pulapkiJSON, plikiJSON string
			err := d.QueryRow(query, params...).Scan(&out.ID, &out.Rok, &out.NumerZadania, &out.Tytul, &out.Kontekst,
				&out.TypZadania, &out.Kategoria, &out.Punkty, &out.Tresc, &out.Odpowiedz,
				&out.ZasadyOceniania, &pulapkiJSON, &out.SciezkaDanych, &plikiJSON)
			if err == sql.ErrNoRows {
				return notFound(fmt.Sprintf("no CKE subtasks found for typ=%s", typ))
			}
			if err != nil {
				return fatal(fmt.Sprintf("query error: %v", err))
			}

			if err := json.Unmarshal([]byte(pulapkiJSON), &out.Pulapki); err != nil {
				fmt.Fprintf(os.Stderr, "warning: unmarshal pulapki: %v\n", err)
			}
			if err := json.Unmarshal([]byte(plikiJSON), &out.PlikiDanych); err != nil {
				fmt.Fprintf(os.Stderr, "warning: unmarshal pliki_danych: %v\n", err)
			}
			if out.Pulapki == nil {
				out.Pulapki = []string{}
			}
			if out.PlikiDanych == nil {
				out.PlikiDanych = []string{}
			}

			addWeight(d, 5)
			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Task type (e.g. sledzenie_algorytmu)")
	cmd.Flags().StringVar(&exclude, "exclude", "", "Comma-separated IDs to exclude")
	cmd.Flags().BoolVar(&force, "force", false, "Bypass unlock check")
	return cmd
}

// === cke worked-example ===

func ckeWorkedExampleCmd() *cobra.Command {
	var typ string
	var force bool

	cmd := &cobra.Command{
		Use:   "worked-example",
		Short: "Get a solved CKE example for study (shows answer + traps)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if typ == "" {
				return fatal("--typ is required")
			}

			d := db(cmd)

			// Gate: require at least srednie-trudne unless --force
			if !force {
				level := getLevel(d, typ)
				if level != "srednie-trudne" && level != "trudne" {
					return notFound(fmt.Sprintf("Przyklad rozwiazany wymaga poziomu srednie-trudne. Twoj poziom: %s. Uzyj --force aby pominac.", level))
				}
			}

			query := `SELECT e.id, e.rok, e.numer_zadania, e.tytul, e.kontekst, e.typ_zadania, e.kategoria, e.punkty, e.tresc, e.odpowiedz, e.zasady_oceniania, e.pulapki, e.sciezka_danych, e.pliki_danych
				FROM data.egzamin e
				WHERE e.typ_zadania = ?
				AND e.id NOT IN (SELECT id FROM matura_zrobione)
				AND e.id NOT IN (SELECT id FROM worked_examples_shown WHERE typ = ?)
				ORDER BY RANDOM() LIMIT 1`

			var out CKEOut
			var pulapkiJSON, plikiJSON string
			err := d.QueryRow(query, typ, typ).Scan(&out.ID, &out.Rok, &out.NumerZadania, &out.Tytul, &out.Kontekst,
				&out.TypZadania, &out.Kategoria, &out.Punkty, &out.Tresc, &out.Odpowiedz,
				&out.ZasadyOceniania, &pulapkiJSON, &out.SciezkaDanych, &plikiJSON)
			if err == sql.ErrNoRows {
				return notFound(fmt.Sprintf("no CKE worked examples available for typ=%s", typ))
			}
			if err != nil {
				return fatal(fmt.Sprintf("query error: %v", err))
			}

			if err := json.Unmarshal([]byte(pulapkiJSON), &out.Pulapki); err != nil {
				fmt.Fprintf(os.Stderr, "warning: unmarshal pulapki: %v\n", err)
			}
			if err := json.Unmarshal([]byte(plikiJSON), &out.PlikiDanych); err != nil {
				fmt.Fprintf(os.Stderr, "warning: unmarshal pliki_danych: %v\n", err)
			}
			if out.Pulapki == nil {
				out.Pulapki = []string{}
			}
			if out.PlikiDanych == nil {
				out.PlikiDanych = []string{}
			}

			// Record that this example was shown (so we don't repeat it)
			today := time.Now().Format("2006-01-02")
			d.Exec("INSERT OR IGNORE INTO worked_examples_shown (id, typ, data) VALUES (?, ?, ?)",
				out.ID, typ, today)

			addWeight(d, 2)
			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Task type (e.g. sledzenie_algorytmu)")
	cmd.Flags().BoolVar(&force, "force", false, "Bypass level check")
	return cmd
}

// === cke save ===

func ckeSaveCmd() *cobra.Command {
	var id string
	var punkty, maxPunkty int

	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save CKE subtask result",
		RunE: func(cmd *cobra.Command, args []string) error {
			if id == "" {
				return fatal("--id is required")
			}

			d := db(cmd)

			var typ string
			err := d.QueryRow("SELECT typ_zadania FROM data.egzamin WHERE id = ?", id).Scan(&typ)
			if err != nil {
				return fatal(fmt.Sprintf("CKE subtask %s not found", id))
			}

			today := time.Now().Format("2006-01-02")
			_, err = d.Exec(`INSERT OR REPLACE INTO matura_zrobione (id, typ, data, punkty, max_punkty) VALUES (?, ?, ?, ?, ?)`,
				id, typ, today, punkty, maxPunkty)
			if err != nil {
				return fatal(fmt.Sprintf("save error: %v", err))
			}

			jsonOut(map[string]any{
				"id":     id,
				"punkty": punkty,
				"max":    maxPunkty,
				"saved":  true,
			})
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "CKE subtask ID (e.g. 2025.1.1)")
	cmd.Flags().IntVar(&punkty, "punkty", 0, "Points scored")
	cmd.Flags().IntVar(&maxPunkty, "max", 0, "Maximum points")
	return cmd
}

// === exam meta ===

func examMetaCmd() *cobra.Command {
	var rok int

	cmd := &cobra.Command{
		Use:   "meta",
		Short: "Get exam metadata (tasks without content)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if rok == 0 {
				return fatal("--rok is required")
			}

			d := db(cmd)

			rows, err := d.Query(`
				SELECT e.numer_zadania, e.tytul, SUM(e.punkty) as total_pkt, MIN(e.czesc) as czesc
				FROM data.egzamin e
				WHERE e.rok = ?
				GROUP BY e.numer_zadania, e.tytul
				ORDER BY e.numer_zadania`, rok)
			if err != nil {
				return fatal(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			var zadania []ExamTaskBrief
			for rows.Next() {
				var tb ExamTaskBrief
				if err := rows.Scan(&tb.Numer, &tb.Tytul, &tb.Punkty, &tb.Czesc); err != nil {
					return fatal(fmt.Sprintf("scan error: %v", err))
				}
				zadania = append(zadania, tb)
			}
			if err := rows.Err(); err != nil {
				return fatal(fmt.Sprintf("rows error: %v", err))
			}

			if len(zadania) == 0 {
				return notFound(fmt.Sprintf("no exam data for rok=%d", rok))
			}

			totalPunkty := 0
			for _, z := range zadania {
				totalPunkty += z.Punkty
			}

			formula := "2023"
			// czas_minuty = 210 for ALL formulas: 2014 (90+120), 2015 (60+150), 2023 (210).
			czasMinuty := 210
			if rok == 2014 {
				formula = "2014"
			} else if rok <= 2022 {
				formula = "2015"
			}

			out := ExamMetaOut{
				Rok:         rok,
				Formula:     formula,
				CzasMinuty:  czasMinuty,
				TotalPunkty: totalPunkty,
				LiczbaZadan: len(zadania),
				Zadania:     zadania,
			}

			switch formula {
			case "2014":
				out.Czesci = []ExamPart{
					{Czesc: 1, CzasMinuty: 90, Punkty: 20, Opis: "Czesc I"},
					{Czesc: 2, CzasMinuty: 120, Punkty: 30, Opis: "Czesc II"},
				}
			case "2015":
				out.Czesci = []ExamPart{
					{Czesc: 1, CzasMinuty: 60, Punkty: 15, Opis: "Czesc I"},
					{Czesc: 2, CzasMinuty: 150, Punkty: 35, Opis: "Czesc II"},
				}
			default:
				out.Czesci = []ExamPart{
					{Czesc: 1, CzasMinuty: 210, Punkty: totalPunkty, Opis: "Jedna czesc (nowa formula)"},
				}
			}

			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().IntVar(&rok, "rok", 0, "Exam year (e.g. 2024)")
	return cmd
}

// === exam task ===

func examTaskCmd() *cobra.Command {
	var rok, zadanie int

	cmd := &cobra.Command{
		Use:   "task",
		Short: "Get exam task with subtasks (full content)",
		RunE: func(cmd *cobra.Command, args []string) error {
			if rok == 0 || zadanie == 0 {
				return fatal("--rok and --zadanie are required")
			}

			d := db(cmd)

			rows, err := d.Query(`
				SELECT e.id, e.numer_podzadania, e.tytul, e.kontekst, e.typ_zadania, e.kategoria, e.punkty, e.tresc, e.odpowiedz, e.zasady_oceniania, e.pulapki, e.sciezka_danych, e.pliki_danych
				FROM data.egzamin e
				WHERE e.rok = ? AND e.numer_zadania = ?
				ORDER BY e.id`, rok, zadanie)
			if err != nil {
				return fatal(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			var out ExamTaskOut
			first := true
			for rows.Next() {
				var sub ExamSubtaskOut
				var tytul, kontekst, sciezkaDanych, plikiJSON, pulapkiJSON string
				err := rows.Scan(&sub.ID, &sub.Numer, &tytul, &kontekst,
					&sub.TypZadania, &sub.Kategoria, &sub.Punkty,
					&sub.Tresc, &sub.Odpowiedz, &sub.ZasadyOceniania,
					&pulapkiJSON, &sciezkaDanych, &plikiJSON)
				if err != nil {
					return fatal(fmt.Sprintf("scan error: %v", err))
				}

				if err := json.Unmarshal([]byte(pulapkiJSON), &sub.Pulapki); err != nil {
					fmt.Fprintf(os.Stderr, "warning: unmarshal pulapki for %s: %v\n", sub.ID, err)
				}
				if sub.Pulapki == nil {
					sub.Pulapki = []string{}
				}

				if first {
					out.Tytul = tytul
					out.Numer = zadanie
					out.Kontekst = kontekst
					out.SciezkaDanych = sciezkaDanych
					if err := json.Unmarshal([]byte(plikiJSON), &out.PlikiDanych); err != nil {
						fmt.Fprintf(os.Stderr, "warning: unmarshal pliki_danych: %v\n", err)
					}
					if out.PlikiDanych == nil {
						out.PlikiDanych = []string{}
					}
					first = false
				}

				out.Podzadania = append(out.Podzadania, sub)
				out.Punkty += sub.Punkty
			}
			if err := rows.Err(); err != nil {
				return fatal(fmt.Sprintf("rows error: %v", err))
			}

			if first {
				return notFound(fmt.Sprintf("no task found for rok=%d zadanie=%d", rok, zadanie))
			}

			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().IntVar(&rok, "rok", 0, "Exam year")
	cmd.Flags().IntVar(&zadanie, "zadanie", 0, "Task number")
	return cmd
}

// === exam save ===

func examSaveCmd() *cobra.Command {
	var rok int
	var resultsJSON string
	var czasMin int

	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save mock exam results",
		RunE: func(cmd *cobra.Command, args []string) error {
			if rok == 0 || resultsJSON == "" {
				return fatal("--rok and --results are required")
			}

			type ResultEntry struct {
				ID  string `json:"id"`
				Pkt int    `json:"pkt"`
				Max int    `json:"max"`
			}

			var results []ResultEntry
			if err := json.Unmarshal([]byte(resultsJSON), &results); err != nil {
				return fatal(fmt.Sprintf("invalid --results JSON: %v", err))
			}

			d := db(cmd)
			today := time.Now().Format("2006-01-02")
			wynikPkt := 0
			maxPkt := 0

			// Begin transaction for all writes (atomicity: all results + summary or nothing)
			tx, err := d.Begin()
			if err != nil {
				return fatal(fmt.Sprintf("begin tx: %v", err))
			}
			defer tx.Rollback()

			for _, r := range results {
				wynikPkt += r.Pkt
				maxPkt += r.Max

				// Use tx.QueryRow (not d.QueryRow) — d.QueryRow may get a new
				// connection from the pool that lacks the ATTACH data alias.
				var typ string
				if err := tx.QueryRow("SELECT typ_zadania FROM data.egzamin WHERE id = ?", r.ID).Scan(&typ); err != nil {
					return fatal(fmt.Sprintf("CKE subtask %s not found: %v", r.ID, err))
				}
				if _, err := tx.Exec(`INSERT OR REPLACE INTO matura_zrobione (id, typ, data, punkty, max_punkty) VALUES (?, ?, ?, ?, ?)`,
					r.ID, typ, today, r.Pkt, r.Max); err != nil {
					return fatal(fmt.Sprintf("save result %s: %v", r.ID, err))
				}
			}

			procent := 0.0
			if maxPkt > 0 {
				procent = float64(wynikPkt) / float64(maxPkt) * 100
			}

			if _, err := tx.Exec(`INSERT INTO probne_matury (rok, data, czas_min, wynik_pkt, max_pkt, procent) VALUES (?, ?, ?, ?, ?, ?)`,
				rok, today, czasMin, wynikPkt, maxPkt, procent); err != nil {
				return fatal(fmt.Sprintf("save error: %v", err))
			}

			if err := tx.Commit(); err != nil {
				return fatal(fmt.Sprintf("commit: %v", err))
			}

			jsonOut(map[string]any{
				"rok":     rok,
				"wynik":   wynikPkt,
				"max":     maxPkt,
				"procent": fmt.Sprintf("%.1f%%", procent),
				"saved":   true,
			})
			return nil
		},
	}

	cmd.Flags().IntVar(&rok, "rok", 0, "Exam year")
	cmd.Flags().StringVar(&resultsJSON, "results", "", "JSON array of results")
	cmd.Flags().IntVar(&czasMin, "czas", 0, "Time spent in minutes")
	return cmd
}

// === exercise next ===

func exerciseNextCmd() *cobra.Command {
	var typ, kategoria string

	cmd := &cobra.Command{
		Use:   "next",
		Short: "Get next exercise with smart priority (review > interleave > new)",
		RunE: func(cmd *cobra.Command, args []string) error {
			d := db(cmd)

			if typ == "" && kategoria != "" {
				chosen, err := chooseWeakestType(d, kategoria)
				if err != nil {
					return err
				}
				typ = chosen
			}

			if typ == "" {
				return fatal("--typ or --kategoria is required")
			}

			sessionCount, _, err := getSessionState(d)
			if err != nil {
				return fatal(fmt.Sprintf("session state: %v", err))
			}
			out := ExerciseNextOut{
				SessionCount: sessionCount,
			}

			finalizeWeight := func() {
				addWeight(d, 4)
				var wStr string
				if err := d.QueryRow(`SELECT value FROM progress_meta WHERE key = 'session_context_weight'`).Scan(&wStr); err == nil {
					fmt.Sscanf(wStr, "%d", &out.SessionWeight)
				}
				out.ResetSuggested = out.SessionWeight >= 80
			}

			if cmd.Flags().Changed("kategoria") {
				out.ChosenTyp = typ
			}

			today := time.Now().Format("2006-01-02")

			// Priority 1: overdue review
			var tag, powtorka string
			err = d.QueryRow(`
				SELECT t.tag, t.nastepna_powtorka
				FROM progress_tagi t
				WHERE t.nastepna_powtorka <= ?
				ORDER BY t.nastepna_powtorka ASC
				LIMIT 1`, today).Scan(&tag, &powtorka)
			if err == nil {
				ex, findErr := findExerciseByTag(d, tag)
				if findErr == nil {
					powtorkaDate, _ := time.Parse("2006-01-02", powtorka)
					daysOverdue := int(time.Since(powtorkaDate).Hours() / 24)
					q := exerciseToQuestion(ex)
					q.Coaching = buildCoaching(d, ex.TypNazwa, ex.Tagi)
					var prevResult sql.NullString
					d.QueryRow("SELECT wynik FROM progress_zrobione WHERE id = ?", ex.ID).Scan(&prevResult)
					if prevResult.Valid {
						q.Coaching.PreviousResult = prevResult.String
					}
					out.Mode = "review"
					out.Exercise = q
					out.ReviewTag = &tag
					out.DaysOverdue = &daysOverdue
					out.PoolWarning = poolWarning(d, ex.TypNazwa, ex.Trudnosc)
					if err := registerFetch(d, q.ID, q.TypNazwa, q.Coaching.HintDelay); err != nil {
						fmt.Fprintf(os.Stderr, "warning: registerFetch: %v\n", err)
					}
					finalizeWeight()
					jsonOut(out)
					return nil
				}
			}

			// Priority 2: interleaving (every 3rd exercise from a different type)
			// Skip interleave when streak=3 (difficulty climb threshold) to avoid stealing the auto-difficulty slot
			if sessionCount > 0 && sessionCount%3 == 0 && getStreak(d, typ) < 3 {
				ex, err := findInterleaveExercise(d, typ)
				if err == nil {
					q := exerciseToQuestion(ex)
					q.Coaching = buildCoaching(d, ex.TypNazwa, ex.Tagi)
					out.Mode = "interleave"
					out.Exercise = q
					out.PoolWarning = poolWarning(d, ex.TypNazwa, ex.Trudnosc)
					if err := registerFetch(d, q.ID, q.TypNazwa, q.Coaching.HintDelay); err != nil {
						fmt.Fprintf(os.Stderr, "warning: registerFetch: %v\n", err)
					}
					finalizeWeight()
					jsonOut(out)
					return nil
				}
			}

			// Priority 3: new exercise with auto-difficulty
			level := getLevel(d, typ)
			ex, err := findExerciseByTypAndLevel(d, typ, level)
			if err != nil {
				ex, err = findExerciseByTypAndLevel(d, typ, "")
				if err != nil {
					return notFound(fmt.Sprintf("no exercises available for typ=%s", typ))
				}
			}

			q := exerciseToQuestion(ex)
			q.Coaching = buildCoaching(d, typ, ex.Tagi)
			out.Mode = "new"
			out.Exercise = q
			out.PoolWarning = poolWarning(d, typ, level)
			if err := registerFetch(d, q.ID, q.TypNazwa, q.Coaching.HintDelay); err != nil {
				fmt.Fprintf(os.Stderr, "warning: registerFetch: %v\n", err)
			}
			finalizeWeight()
			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Exercise type")
	cmd.Flags().StringVar(&kategoria, "kategoria", "", "Category (TEORIA/IMPLEMENTACJA/ARKUSZ/SQL) — auto-selects weakest type")
	return cmd
}

// chooseWeakestType picks the type with lowest streak (or oldest review) within a category.
func chooseWeakestType(d *sql.DB, kategoria string) (string, error) {
	rows, err := d.Query(`
		SELECT DISTINCT c.typ_nazwa, COALESCE(p.streak, 0) as streak,
			COALESCE(MAX(z.data), '1970-01-01') as last_done
		FROM data.cwiczenia c
		LEFT JOIN progress_typy p ON p.typ = c.typ_nazwa
		LEFT JOIN progress_zrobione z ON z.typ = c.typ_nazwa
		WHERE c.kategoria = ?
		GROUP BY c.typ_nazwa
		ORDER BY streak ASC, last_done ASC
		LIMIT 1`, kategoria)
	if err != nil {
		return "", fatal(fmt.Sprintf("query types for kategoria=%s: %v", kategoria, err))
	}
	defer rows.Close()

	if rows.Next() {
		var typ string
		var streak int
		var lastDone string
		if err := rows.Scan(&typ, &streak, &lastDone); err != nil {
			return "", fatal(fmt.Sprintf("scan type: %v", err))
		}
		return typ, nil
	}
	if err := rows.Err(); err != nil {
		return "", fatal(fmt.Sprintf("rows error: %v", err))
	}
	return "", notFound(fmt.Sprintf("no types found for kategoria=%s", kategoria))
}

// === typ intro ===

func typIntroCmd() *cobra.Command {
	var typ string

	cmd := &cobra.Command{
		Use:   "intro",
		Short: "Get type introduction context",
		RunE: func(cmd *cobra.Command, args []string) error {
			if typ == "" {
				return fatal("--typ is required")
			}

			d := db(cmd)

			kat := getKategoria(d, typ)
			if kat == "" {
				return notFound(fmt.Sprintf("unknown type: %s", typ))
			}

			out := TypIntroOut{
				Typ:       typ,
				Kategoria: kat,
			}

			// Check if first in type
			var doneInType int
			d.QueryRow(`SELECT COUNT(*) FROM progress_zrobione z
				JOIN data.cwiczenia c ON c.id = z.id
				WHERE c.typ_nazwa = ?`, typ).Scan(&doneInType)
			out.FirstInType = doneInType == 0
			out.Done = doneInType

			// Check if OTHER types in this category have been done
			otherRows, err := d.Query(`SELECT DISTINCT c.typ_nazwa FROM progress_zrobione z
				JOIN data.cwiczenia c ON c.id = z.id
				WHERE c.kategoria = ? AND c.typ_nazwa != ?`, kat, typ)
			var otherTypes []string
			if err == nil {
				for otherRows.Next() {
					var t string
					otherRows.Scan(&t)
					otherTypes = append(otherTypes, t)
				}
				if rowErr := otherRows.Err(); rowErr != nil {
					fmt.Fprintf(os.Stderr, "warning: other types rows: %v\n", rowErr)
				}
				otherRows.Close()
			}
			out.FirstInCategory = len(otherTypes) == 0
			if otherTypes == nil {
				otherTypes = []string{}
			}
			out.OtherTypesDoneInCategory = otherTypes

			// Level and streak
			out.Level = getLevel(d, typ)
			var streak int
			d.QueryRow("SELECT COALESCE(streak, 0) FROM progress_typy WHERE typ = ?", typ).Scan(&streak)
			out.Streak = streak

			// Available exercises
			var available int
			d.QueryRow(`SELECT COUNT(*) FROM data.cwiczenia WHERE typ_nazwa = ?
				AND id NOT IN (SELECT id FROM progress_zrobione)`, typ).Scan(&available)
			out.Available = available

			out.SprawdzianUnlocked = out.Level == "trudne"

			// CKE stats
			ckeNames := exerciseTypToCKETypes(typ, kat)
			var ckeStats CKETypStats
			ckeStats.LatTotal = 11
			err = d.QueryRow(
				`SELECT COUNT(DISTINCT rok), COALESCE(AVG(punkty), 0), COALESCE(SUM(punkty), 0)
				FROM data.egzamin WHERE typ_zadania IN (`+placeholders(len(ckeNames))+`)`,
				toAny(ckeNames)...).
				Scan(&ckeStats.Wystapienia, &ckeStats.AvgPunkty, &ckeStats.TotalPunkty)
			if err == nil && ckeStats.Wystapienia > 0 {
				out.CKEStats = &ckeStats
			}

			// Top pulapki — single pass: count + preserve first-seen casing
			trapRows, err := d.Query(
				`SELECT pulapki FROM data.egzamin
				WHERE typ_zadania IN (`+placeholders(len(ckeNames))+`)
				AND pulapki != '[]' AND pulapki != 'null'`, toAny(ckeNames)...)
			if err == nil {
				type pulapkaEntry struct {
					Original string
					Count    int
				}
				pulapkiMap := map[string]*pulapkaEntry{} // lowercase -> entry

				for trapRows.Next() {
					var pJSON string
					trapRows.Scan(&pJSON)
					var pulapki []string
					if err := json.Unmarshal([]byte(pJSON), &pulapki); err != nil {
						fmt.Fprintf(os.Stderr, "warning: unmarshal pulapki: %v\n", err)
						continue
					}
					for _, p := range pulapki {
						p = strings.TrimSpace(p)
						if p == "" {
							continue
						}
						lp := strings.ToLower(p)
						if e, ok := pulapkiMap[lp]; ok {
							e.Count++
						} else {
							pulapkiMap[lp] = &pulapkaEntry{Original: p, Count: 1}
						}
					}
				}
				if rowErr := trapRows.Err(); rowErr != nil {
					fmt.Fprintf(os.Stderr, "warning: trap rows: %v\n", rowErr)
				}
				trapRows.Close()

				// Sort by frequency descending, then alphabetically for stability
				sorted := make([]*pulapkaEntry, 0, len(pulapkiMap))
				for _, e := range pulapkiMap {
					sorted = append(sorted, e)
				}
				sort.Slice(sorted, func(i, j int) bool {
					if sorted[i].Count != sorted[j].Count {
						return sorted[i].Count > sorted[j].Count
					}
					return sorted[i].Original < sorted[j].Original
				})

				topN := 3
				if len(sorted) < topN {
					topN = len(sorted)
				}
				for i := 0; i < topN; i++ {
					out.TopPulapki = append(out.TopPulapki, sorted[i].Original)
				}
			}

			// Przyklad (shortest easy exercise)
			var brief ExerciseBrief
			err = d.QueryRow(
				`SELECT id, tresc, odpowiedz FROM data.cwiczenia
				WHERE typ_nazwa = ? AND trudnosc = 'latwe'
				ORDER BY LENGTH(tresc) ASC LIMIT 1`, typ).
				Scan(&brief.ID, &brief.Tresc, &brief.Odpowiedz)
			if err == nil {
				out.Przyklad = &brief
			}

			// Cheatsheet excerpt
			var fullContent string
			d.QueryRow("SELECT content FROM data.cheatsheets WHERE kategoria = ?", kat).
				Scan(&fullContent)
			if len(fullContent) > 500 {
				cutoff := strings.LastIndex(fullContent[:500], ".")
				if cutoff > 200 {
					fullContent = fullContent[:cutoff+1]
				} else {
					fullContent = fullContent[:500] + "..."
				}
			}
			out.CheatsheetExcerpt = fullContent

			if out.FirstInCategory {
				addWeight(d, 10)
			} else {
				addWeight(d, 4)
			}
			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Exercise type")
	return cmd
}

// === cke status ===

func ckeStatusCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "status",
		Short: "Show CKE unlock status for all types",
		RunE: func(cmd *cobra.Command, args []string) error {
			d := db(cmd)

			// Get all types with CKE subtasks
			rows, err := d.Query(`
				SELECT e.typ_zadania, e.kategoria, COUNT(*) as total
				FROM data.egzamin e
				GROUP BY e.typ_zadania, e.kategoria
				ORDER BY e.kategoria, e.typ_zadania`)
			if err != nil {
				return fatal(fmt.Sprintf("query error: %v", err))
			}

			var results []CKEStatusOut
			for rows.Next() {
				var out CKEStatusOut
				var total int
				if err := rows.Scan(&out.Typ, &out.Kategoria, &total); err != nil {
					rows.Close()
					return fatal(fmt.Sprintf("scan error: %v", err))
				}
				out.CKEAvailable = total
				results = append(results, out)
			}
			if err := rows.Err(); err != nil {
				rows.Close()
				return fatal(fmt.Sprintf("rows error: %v", err))
			}
			rows.Close()

			// Enrich after rows closed to avoid deadlock with MaxOpenConns(1)
			for i := range results {
				results[i].Level = getLevel(d, results[i].Typ)
				results[i].Unlocked = results[i].Level == "trudne"

				var done int
				d.QueryRow("SELECT COUNT(*) FROM matura_zrobione WHERE typ = ?", results[i].Typ).Scan(&done)
				results[i].CKEDone = done

				var worked int
				d.QueryRow("SELECT COUNT(*) FROM worked_examples_shown WHERE typ = ?", results[i].Typ).Scan(&worked)
				results[i].CKEWorkedExamples = worked
				results[i].CKEAvailable = results[i].CKEAvailable - done - worked
			}

			if results == nil {
				results = []CKEStatusOut{}
			}
			jsonOut(results)
			return nil
		},
	}
}

// === exam list ===

func examListCmd() *cobra.Command {
	var formula string
	var random bool

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List available exam years with status",
		RunE: func(cmd *cobra.Command, args []string) error {
			d := db(cmd)

			rows, err := d.Query(`
				SELECT rok, SUM(punkty) as total
				FROM data.egzamin
				GROUP BY rok
				ORDER BY rok`)
			if err != nil {
				return fatal(fmt.Sprintf("query error: %v", err))
			}

			var rawEntries []ExamListEntry
			for rows.Next() {
				var e ExamListEntry
				if err := rows.Scan(&e.Rok, &e.TotalPkt); err != nil {
					rows.Close()
					return fatal(fmt.Sprintf("scan error: %v", err))
				}

				if e.Rok == 2014 {
					e.Formula = "2014"
					e.CzasMin = 210 // 90+120 total
				} else if e.Rok <= 2022 {
					e.Formula = "2015"
					e.CzasMin = 210 // 60+150 total
				} else {
					e.Formula = "2023"
					e.CzasMin = 210
				}

				rawEntries = append(rawEntries, e)
			}
			if err := rows.Err(); err != nil {
				rows.Close()
				return fatal(fmt.Sprintf("rows error: %v", err))
			}
			rows.Close()

			// Enrich after rows closed to avoid deadlock with MaxOpenConns(1)
			var entries []ExamListEntry
			for _, e := range rawEntries {
				// Check if done (has a non-interrupted mock exam)
				var procent sql.NullFloat64
				d.QueryRow("SELECT procent FROM probne_matury WHERE rok = ? AND przerwany = 0 ORDER BY procent DESC LIMIT 1", e.Rok).Scan(&procent)
				if procent.Valid {
					e.Done = true
					p := procent.Float64
					e.WynikProcent = &p
				}

				// Apply formula filter
				if formula != "" {
					if formula == "nowa" && e.Formula != "2023" {
						continue
					}
					if formula == "stara" && e.Formula == "2023" {
						continue
					}
				}

				entries = append(entries, e)
			}

			if len(entries) == 0 {
				return notFound("no exams found")
			}

			out := ExamListOut{Available: entries}

			// Suggestion logic: prefer undone with nowa formula
			for _, e := range entries {
				if !e.Done && e.Formula == "2023" {
					out.Suggested = &ExamSuggestion{
						Rok:     e.Rok,
						Formula: e.Formula,
						Reason:  "Nowa formula, jeszcze nie zrobiona",
					}
					break
				}
			}
			if out.Suggested == nil {
				for _, e := range entries {
					if !e.Done {
						out.Suggested = &ExamSuggestion{
							Rok:     e.Rok,
							Formula: e.Formula,
							Reason:  "Jeszcze nie zrobiona",
						}
						break
					}
				}
			}

			if random && len(entries) > 0 {
				// Filter to undone first
				var undone []ExamListEntry
				for _, e := range entries {
					if !e.Done {
						undone = append(undone, e)
					}
				}
				if len(undone) > 0 {
					chosen := undone[rand.Intn(len(undone))]
					out.Available = []ExamListEntry{chosen}
					out.Suggested = &ExamSuggestion{
						Rok:     chosen.Rok,
						Formula: chosen.Formula,
						Reason:  "Losowy wybor z niezrobionych",
					}
				} else {
					chosen := entries[rand.Intn(len(entries))]
					out.Available = []ExamListEntry{chosen}
					out.Suggested = &ExamSuggestion{
						Rok:     chosen.Rok,
						Formula: chosen.Formula,
						Reason:  "Losowy wybor (wszystkie zrobione)",
					}
				}
			}

			jsonOut(out)
			return nil
		},
	}

	cmd.Flags().StringVar(&formula, "formula", "", "Filter: nowa, stara, or empty for all")
	cmd.Flags().BoolVar(&random, "random", false, "Pick one random year")
	return cmd
}

// === trap list ===

func trapListCmd() *cobra.Command {
	var typ, kategoria string

	cmd := &cobra.Command{
		Use:   "list",
		Short: "List traps/gotchas from CKE exams",
		RunE: func(cmd *cobra.Command, args []string) error {
			d := db(cmd)

			query := `SELECT e.id, e.rok, e.tytul, e.pulapki FROM data.egzamin e WHERE e.pulapki != '[]' AND e.pulapki != 'null'`
			var params []any

			if typ != "" {
				query += " AND e.typ_zadania = ?"
				params = append(params, typ)
			}
			if kategoria != "" {
				query += " AND e.kategoria = ?"
				params = append(params, kategoria)
			}

			query += " ORDER BY e.rok DESC, e.id"

			rows, err := d.Query(query, params...)
			if err != nil {
				return fatal(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			var results []TrapOut
			for rows.Next() {
				var t TrapOut
				var pulapkiJSON string
				if err := rows.Scan(&t.SourceID, &t.Rok, &t.Tytul, &pulapkiJSON); err != nil {
					return fatal(fmt.Sprintf("scan error: %v", err))
				}
				if err := json.Unmarshal([]byte(pulapkiJSON), &t.Pulapki); err != nil {
					fmt.Fprintf(os.Stderr, "warning: unmarshal pulapki for %s: %v\n", t.SourceID, err)
					continue
				}
				if len(t.Pulapki) > 0 {
					results = append(results, t)
				}
			}
			if err := rows.Err(); err != nil {
				return fatal(fmt.Sprintf("rows error: %v", err))
			}

			if len(results) == 0 {
				return notFound("no traps found")
			}

			jsonOut(results)
			return nil
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Filter by task type")
	cmd.Flags().StringVar(&kategoria, "kategoria", "", "Filter by category (TEORIA/IMPLEMENTACJA/ARKUSZ/SQL)")
	return cmd
}

// === trap save ===

func trapSaveCmd() *cobra.Command {
	var id, typ string
	var trafienia, total int

	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save trap quiz result",
		RunE: func(cmd *cobra.Command, args []string) error {
			if id == "" || typ == "" {
				return fatal("--id and --typ are required")
			}
			if total <= 0 {
				return fatal("--total must be > 0")
			}

			d := db(cmd)

			today := time.Now().Format("2006-01-02")
			_, err := d.Exec(
				`INSERT OR REPLACE INTO pulapki_przejrzane (id, typ, data, trafienia, total) VALUES (?, ?, ?, ?, ?)`,
				id, typ, today, trafienia, total)
			if err != nil {
				return fatal(fmt.Sprintf("save trap result: %v", err))
			}

			jsonOut(TrapSaveOut{ID: id, Typ: typ, Trafienia: trafienia, Total: total})
			return nil
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "CKE subtask ID (e.g. 2024.1.1)")
	cmd.Flags().StringVar(&typ, "typ", "", "Task type")
	cmd.Flags().IntVar(&trafienia, "trafienia", 0, "Number of traps correctly identified")
	cmd.Flags().IntVar(&total, "total", 0, "Total number of traps in the task")
	return cmd
}

// === cheatsheet get ===

func cheatsheetGetCmd() *cobra.Command {
	var kategoria, sekcja string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get cheatsheet content",
		RunE: func(cmd *cobra.Command, args []string) error {
			if kategoria == "" {
				return fatal("--kategoria is required")
			}

			d := db(cmd)

			var content string
			err := d.QueryRow("SELECT content FROM data.cheatsheets WHERE kategoria = ?", kategoria).Scan(&content)
			if err != nil {
				return notFound(fmt.Sprintf("cheatsheet for kategoria=%s not found", kategoria))
			}

			// If --sekcja given, extract only the matching section
			if sekcja != "" {
				if extracted := extractSection(content, sekcja); extracted != "" {
					content = extracted
				}
				// fallback: if not found, return full content
			}

			if sekcja != "" {
				addWeight(d, 2)
			} else {
				addWeight(d, 8)
			}

			// Raw markdown output (not jsonOut) — intentional.
			// SKILL.md expects plain text for direct citation by the AI tutor.
			fmt.Print(content)
			return nil
		},
	}

	cmd.Flags().StringVar(&kategoria, "kategoria", "", "Category: TEORIA, IMPLEMENTACJA, ARKUSZ, SQL")
	cmd.Flags().StringVar(&sekcja, "sekcja", "", "Section header (partial match, case-insensitive)")
	return cmd
}

// findCheatsheetExcerptByTags searches the cheatsheet for the exercise's
// category, matching against the exercise's tags. Returns the first matching
// section, or "" if no match.
func findCheatsheetExcerptByTags(d *sql.DB, exerciseID string) string {
	var kategoria, tagiJSON string
	err := d.QueryRow(`SELECT kategoria, tagi FROM data.cwiczenia WHERE id = ?`, exerciseID).Scan(&kategoria, &tagiJSON)
	if err != nil {
		return ""
	}
	var tagi []string
	json.Unmarshal([]byte(tagiJSON), &tagi)
	if len(tagi) == 0 {
		return ""
	}

	var content string
	err = d.QueryRow(`SELECT content FROM data.cheatsheets WHERE kategoria = ?`, kategoria).Scan(&content)
	if err != nil {
		return ""
	}

	// Try each tag against cheatsheet sections
	for _, tag := range tagi {
		if section := extractSection(content, tag); section != "" {
			return section
		}
	}
	return ""
}

// extractSection finds a ## section whose title contains query (case-insensitive)
// and returns its content up to the next ## or EOF.
// Uses prefix matching (min 5 chars) to handle Polish inflection (e.g. "pulapki" matches "pulapek").
func extractSection(markdown, query string) string {
	queryLower := strings.ToLower(query)
	sections := strings.Split("\n"+markdown, "\n## ")
	for i, sec := range sections {
		if i == 0 {
			continue // skip content before first ##
		}
		newlineIdx := strings.Index(sec, "\n")
		if newlineIdx < 0 {
			newlineIdx = len(sec)
		}
		heading := strings.ToLower(sec[:newlineIdx])
		if strings.Contains(heading, queryLower) {
			return "## " + sec
		}
		// Prefix matching: check if any word in heading shares a common prefix (>=5 chars) with query
		if len(queryLower) >= 5 {
			for _, word := range strings.Fields(heading) {
				cp := commonPrefix(word, queryLower)
				if cp >= 5 {
					return "## " + sec
				}
			}
		}
	}
	return ""
}

func commonPrefix(a, b string) int {
	n := len(a)
	if len(b) < n {
		n = len(b)
	}
	for i := 0; i < n; i++ {
		if a[i] != b[i] {
			return i
		}
	}
	return n
}

// === data import ===

func dataImportCmd() *cobra.Command {
	var source string

	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import JSON data into matura.db",
		RunE: func(cmd *cobra.Command, args []string) error {
			dataDB, err := OpenDataDB(dbDir)
			if err != nil {
				return fatal(fmt.Sprintf("open data DB: %v", err))
			}
			defer dataDB.Close()

			if err := ImportAll(dataDB, source); err != nil {
				return fatal(fmt.Sprintf("import failed: %v", err))
			}

			fmt.Fprintln(os.Stderr, "Import complete.")
			return nil
		},
	}

	cmd.Flags().StringVar(&source, "source", "", "Path to analiza/ directory")
	cmd.MarkFlagRequired("source")
	return cmd
}

// === data stats ===

func dataStatsCmd() *cobra.Command {
	return &cobra.Command{
		Use:   "stats",
		Short: "Show data counts",
		RunE: func(cmd *cobra.Command, args []string) error {
			d := db(cmd)
			var out DataStatsOut
			d.QueryRow("SELECT COUNT(*) FROM data.cwiczenia").Scan(&out.Cwiczenia)
			d.QueryRow("SELECT COUNT(*) FROM data.egzamin").Scan(&out.Podzadania)
			d.QueryRow("SELECT COUNT(*) FROM data.cheatsheets").Scan(&out.Cheatsheets)
			jsonOut(out)
			return nil
		},
	}
}

// === data verify ===

func dataVerifyCmd() *cobra.Command {
	var source string

	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Verify JSON files match matura.db contents",
		RunE: func(cmd *cobra.Command, args []string) error {
			dataDB, err := OpenDataDB(dbDir)
			if err != nil {
				return fatal(fmt.Sprintf("open data DB: %v", err))
			}
			defer dataDB.Close()

			result, err := VerifyExercises(dataDB, source)
			if err != nil {
				return fatal(fmt.Sprintf("verify failed: %v", err))
			}

			jsonOut(result)

			if len(result.Mismatched) > 0 || len(result.MissingInDB) > 0 || len(result.MissingOnDisk) > 0 {
				return fatal(fmt.Sprintf("%d mismatched, %d missing in DB, %d missing on disk",
					len(result.Mismatched), len(result.MissingInDB), len(result.MissingOnDisk)))
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&source, "source", "", "Path to analiza/ directory")
	cmd.MarkFlagRequired("source")
	return cmd
}

// === test-report summary ===

func testReportSummaryCmd() *cobra.Command {
	var historiaPath string
	var window int
	var format string

	cmd := &cobra.Command{
		Use:   "summary",
		Short: "Generate summary report from test-tutor history",
		RunE: func(cmd *cobra.Command, args []string) error {
			if historiaPath == "" {
				exe, err := os.Executable()
				if err == nil {
					historiaPath = filepath.Join(filepath.Dir(exe),
						"..", "test_pedagogical", "reports", "historia.json")
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
				enc.SetEscapeHTML(false)
				return enc.Encode(summary)
			default:
				md := RenderSummaryMarkdown(summary)
				fmt.Fprint(cmd.OutOrStdout(), md)
				return nil
			}
		},
	}

	cmd.Flags().StringVar(&historiaPath, "historia", "", "Path to historia.json (default: auto-detect)")
	cmd.Flags().IntVar(&window, "window", 10, "Analysis window size")
	cmd.Flags().StringVar(&format, "format", "md", "Output format: md or json")

	return cmd
}

// === Exercise Rubric ===

var typKategoriaMap = map[string]string{
	"sledzenie_algorytmu": "TEORIA", "projektowanie_algorytmu": "TEORIA",
	"analiza_algorytmu": "TEORIA", "test_prawda_falsz": "TEORIA",
	"konwersja_systemow_liczbowych": "TEORIA", "teoria_bezpieczenstwa": "TEORIA",
	"cyfry_liczby": "IMPLEMENTACJA", "napisy": "IMPLEMENTACJA",
	"zlozone": "IMPLEMENTACJA", "zliczanie": "IMPLEMENTACJA",
	"minmax": "IMPLEMENTACJA", "sekwencje": "IMPLEMENTACJA",
	"obrazy_2D": "IMPLEMENTACJA", "geometryczne": "IMPLEMENTACJA",
	"agregacja_warunkowa": "ARKUSZ", "symulacja": "ARKUSZ",
	"wykres": "ARKUSZ", "agregacja_podstawowa": "ARKUSZ",
	"transformacja": "ARKUSZ",
	"sql_group_by": "SQL", "sql_podzapytania": "SQL",
	"sql_join": "SQL", "sql_select_where": "SQL",
}

func implRubric() RubricDetail {
	return RubricDetail{
		Full:  RubricLevel{"Program kompiluje sie, daje poprawne wyniki dla danych przykladowych i pelnych", 100},
		Half:  RubricLevel{"Poprawny algorytm, ale bledy: off-by-one, brak obslugi brzegowych, bledne I/O", 50},
		Zero:  RubricLevel{"Zly algorytm lub program sie nie kompiluje", 0},
		Notes: "Czesciowe punkty: poprawne wczytanie (25%), poprawny algorytm z drobnymi bledami (50%), poprawne wypisanie (75%).",
	}
}

func arkuszRubric() RubricDetail {
	return RubricDetail{
		Full:  RubricLevel{"Poprawna formula dajaca prawidlowe wyniki + poprawne adresowanie", 100},
		Half:  RubricLevel{"Poprawna idea formuly, ale blad adresowania ($) lub zly zakres", 50},
		Zero:  RubricLevel{"Zla formula lub brak rozwiazania", 0},
		Notes: "Blad $ (brak dolara przy stale) = najczestszy blad. Poprawna formula ze zlym zakresem = 50%.",
	}
}

func sqlRubric() RubricDetail {
	return RubricDetail{
		Full:  RubricLevel{"Poprawne zapytanie SQL dajace prawidlowy wynik", 100},
		Half:  RubricLevel{"Poprawna struktura (tabele, JOIN, GROUP BY), ale blad w warunku lub agregacji", 50},
		Zero:  RubricLevel{"Zla struktura zapytania lub brak rozwiazania", 0},
		Notes: "Kolejnosc kolumn w wyniku nie ma znaczenia. Aliasy opcjonalne. Brak GROUP BY przy agregacji = 0%.",
	}
}

var rubricData = map[string]RubricDetail{
	// TEORIA — unique per type
	"sledzenie_algorytmu": {
		Full:  RubricLevel{"Tabela poprawna + wynik koncowy poprawny", 100},
		Half:  RubricLevel{"Poprawny tok rozumowania, 1-2 bledy rachunkowe w wierszach", 50},
		Zero:  RubricLevel{"Zly algorytm lub brak tabeli", 0},
		Notes: "Kazdy wiersz tabeli jest wart punkty. Bledny wynik przy poprawnym toku = 50%.",
	},
	"projektowanie_algorytmu": {
		Full:  RubricLevel{"Poprawny pseudokod/C++ rozwiazujacy problem", 100},
		Half:  RubricLevel{"Poprawna idea, bledy skladniowe lub drobne luki", 50},
		Zero:  RubricLevel{"Zly algorytm lub brak rozwiazania", 0},
		Notes: "Liczy sie poprawnosc algorytmu, nie skladnia. Brak obslugi przypadkow brzegowych = -25%.",
	},
	"analiza_algorytmu": {
		Full:  RubricLevel{"Poprawna klasa zlozonosci O() + uzasadnienie", 100},
		Half:  RubricLevel{"Poprawna klasa zlozonosci bez uzasadnienia", 50},
		Zero:  RubricLevel{"Zla klasa zlozonosci", 0},
		Notes: "Uzasadnienie musi odwolywac sie do struktury algorytmu (petla, rekurencja).",
	},
	"test_prawda_falsz": {
		Full:  RubricLevel{"Poprawne P/F + poprawne uzasadnienie", 100},
		Half:  RubricLevel{"Poprawne P/F bez uzasadnienia lub z blednym uzasadnieniem", 50},
		Zero:  RubricLevel{"Bledne P/F", 0},
		Notes: "Brak uzasadnienia = ZAWSZE max 50%. CKE wymaga uzasadnienia nawet przy poprawnej odpowiedzi.",
	},
	"konwersja_systemow_liczbowych": {
		Full:  RubricLevel{"Poprawny wynik + zapis posredni obliczen", 100},
		Half:  RubricLevel{"Poprawny wynik bez zapisu posredniego", 50},
		Zero:  RubricLevel{"Bledny wynik", 0},
		Notes: "Zapis posredni: kolumna dzielenia z resztami, grupowanie bitow, itd.",
	},
	"teoria_bezpieczenstwa": {
		Full:  RubricLevel{"Poprawne dopasowanie + definicja/wyjasnienie", 100},
		Half:  RubricLevel{"Poprawne dopasowanie bez definicji", 50},
		Zero:  RubricLevel{"Bledne dopasowanie", 0},
		Notes: "Przy pytaniach otwartych: wymagana precyzyjna definicja, nie ogolniki.",
	},
	// IMPLEMENTACJA — shared rubric
	"cyfry_liczby":  implRubric(),
	"napisy":        implRubric(),
	"zlozone":       implRubric(),
	"zliczanie":     implRubric(),
	"minmax":        implRubric(),
	"sekwencje":     implRubric(),
	"obrazy_2D":     implRubric(),
	"geometryczne":  implRubric(),
	// ARKUSZ — shared rubric
	"agregacja_warunkowa":  arkuszRubric(),
	"symulacja":            arkuszRubric(),
	"wykres":               arkuszRubric(),
	"agregacja_podstawowa": arkuszRubric(),
	"transformacja":        arkuszRubric(),
	// SQL — shared rubric
	"sql_group_by":      sqlRubric(),
	"sql_podzapytania":  sqlRubric(),
	"sql_join":          sqlRubric(),
	"sql_select_where":  sqlRubric(),
}

func getRubric(typ string) (RubricOut, error) {
	rubric, ok := rubricData[typ]
	if !ok {
		return RubricOut{}, fmt.Errorf("unknown exercise type: %s", typ)
	}
	kat := typKategoriaMap[typ]
	return RubricOut{Typ: typ, Kategoria: kat, Rubric: rubric}, nil
}

func exerciseRubricCmd() *cobra.Command {
	var typ string
	cmd := &cobra.Command{
		Use:   "rubric",
		Short: "Get grading rubric for an exercise type",
		RunE: func(cmd *cobra.Command, args []string) error {
			if typ == "" {
				return fatal("--typ is required")
			}
			out, err := getRubric(typ)
			if err != nil {
				return fatal(err.Error())
			}
			d := db(cmd)
			addWeight(d, 1)
			jsonOut(out)
			return nil
		},
	}
	cmd.Flags().StringVar(&typ, "typ", "", "Exercise type (e.g. sledzenie_algorytmu)")
	return cmd
}
