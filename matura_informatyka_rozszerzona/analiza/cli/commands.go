package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
	"sort"
	"strings"
	"time"

	"github.com/spf13/cobra"
)

// Spaced repetition intervals (days) by poziom 0-4
var intervals = [5]int{0, 1, 3, 7, 21}

func jsonOut(v any) {
	enc := json.NewEncoder(os.Stdout)
	enc.SetIndent("", "  ")
	enc.SetEscapeHTML(false)
	enc.Encode(v)
}

func exitNotFound(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(1)
}

func exitError(msg string) {
	fmt.Fprintln(os.Stderr, msg)
	os.Exit(2)
}

func db() *sql.DB { return mainDB }

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
	json.Unmarshal([]byte(tagiJSON), &ex.Tagi)
	json.Unmarshal([]byte(wskazowkiJSON), &ex.Wskazowki)
	json.Unmarshal([]byte(typoweBledyJSON), &ex.TypoweBledy)
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

func queryExercises(typ, trudnosc, exclude string) []ExerciseOut {
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
	rows, err := db().Query(query, params...)
	if err != nil {
		exitError(fmt.Sprintf("query error: %v", err))
	}
	defer rows.Close()
	var results []ExerciseOut
	for rows.Next() {
		ex, err := scanExercise(rows)
		if err != nil {
			exitError(fmt.Sprintf("scan error: %v", err))
		}
		results = append(results, ex)
	}
	return results
}

func getLevel(d *sql.DB, typ string) string {
	var level sql.NullString
	d.QueryRow("SELECT poziom_trudnosci FROM progress_typy WHERE typ = ?", typ).Scan(&level)
	if level.Valid {
		return level.String
	}
	return "latwe"
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

func getSessionState(d *sql.DB) (count int, lastTyp string) {
	today := time.Now().Format("2006-01-02")
	var sessionDate sql.NullString
	d.QueryRow("SELECT value FROM progress_meta WHERE key = 'session_date'").Scan(&sessionDate)
	if !sessionDate.Valid || sessionDate.String != today {
		d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_date', ?)", today)
		d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_exercise_count', '0')")
		d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_last_typ', '')")
		return 0, ""
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
	return
}

func incrementSession(d *sql.DB, typ string) {
	today := time.Now().Format("2006-01-02")
	var sessionDate sql.NullString
	d.QueryRow("SELECT value FROM progress_meta WHERE key = 'session_date'").Scan(&sessionDate)
	if !sessionDate.Valid || sessionDate.String != today {
		d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_date', ?)", today)
		d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_exercise_count', '1')")
	} else {
		d.Exec(`UPDATE progress_meta SET value = CAST(CAST(value AS INTEGER) + 1 AS TEXT) WHERE key = 'session_exercise_count'`)
	}
	d.Exec("INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_last_typ', ?)", typ)
}

func findExerciseByTag(d *sql.DB, tag string) (ExerciseOut, error) {
	row := d.QueryRow(`SELECT `+exerciseColumns+` FROM data.cwiczenia c
		WHERE c.tagi LIKE ? AND c.id NOT IN (SELECT id FROM progress_zrobione)
		ORDER BY RANDOM() LIMIT 1`, "%"+tag+"%")
	ex, err := scanExercise(row)
	if err == nil {
		return ex, nil
	}
	row = d.QueryRow(`SELECT `+exerciseColumns+` FROM data.cwiczenia c
		WHERE c.tagi LIKE ? ORDER BY RANDOM() LIMIT 1`, "%"+tag+"%")
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

// === exercise get ===

func exerciseGetCmd() *cobra.Command {
	var typ, trudnosc, exclude string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a random exercise by type",
		Run: func(cmd *cobra.Command, args []string) {
			if typ == "" {
				exitError("--typ is required")
			}

			// Auto-difficulty: if --trudnosc not explicitly set, use current level
			if !cmd.Flags().Changed("trudnosc") {
				trudnosc = getLevel(db(), typ)
			}

			results := queryExercises(typ, trudnosc, exclude)

			// Fallback: if auto-difficulty yielded 0 results, try without difficulty filter
			if len(results) == 0 && !cmd.Flags().Changed("trudnosc") {
				results = queryExercises(typ, "", exclude)
			}

			if len(results) == 0 {
				exitNotFound(fmt.Sprintf("no exercises found for typ=%s trudnosc=%s", typ, trudnosc))
			}

			chosen := results[rand.Intn(len(results))]
			jsonOut(chosen)
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Exercise type (e.g. sql_group_by, cyfry_liczby)")
	cmd.Flags().StringVar(&trudnosc, "trudnosc", "", "Difficulty: latwe, srednie, srednie-trudne, trudne")
	cmd.Flags().StringVar(&exclude, "exclude", "", "Comma-separated IDs to exclude")
	return cmd
}

// === exercise review ===

func exerciseReviewCmd() *cobra.Command {
	var limit int

	cmd := &cobra.Command{
		Use:   "review",
		Short: "Get exercises due for spaced repetition review",
		Run: func(cmd *cobra.Command, args []string) {
			today := time.Now().Format("2006-01-02")

			rows, err := db().Query(`
				SELECT t.tag, t.nastepna_powtorka
				FROM progress_tagi t
				WHERE t.nastepna_powtorka <= ?
				ORDER BY t.nastepna_powtorka ASC
				LIMIT ?`, today, limit)
			if err != nil {
				exitError(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			var results []ReviewOut
			for rows.Next() {
				var tag, powtorka string
				if err := rows.Scan(&tag, &powtorka); err != nil {
					exitError(fmt.Sprintf("scan error: %v", err))
				}

				powtorkaDate, _ := time.Parse("2006-01-02", powtorka)
				daysOverdue := int(time.Since(powtorkaDate).Hours() / 24)

				ex, err := findExerciseByTag(db(), tag)
				if err != nil {
					continue
				}

				results = append(results, ReviewOut{
					Exercise:    ex,
					Tag:         tag,
					DaysOverdue: daysOverdue,
				})
			}

			if len(results) == 0 {
				exitNotFound("no reviews due")
			}
			jsonOut(results)
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
		Run: func(cmd *cobra.Command, args []string) {
			if id == "" || wynik == "" {
				exitError("--id and --wynik are required")
			}

			validWynik := map[string]bool{
				"poprawne_bez_pomocy": true,
				"poprawne_z_pomoca_1": true,
				"poprawne_z_pomoca_2": true,
				"walk_through":        true,
			}
			if !validWynik[wynik] {
				exitError("--wynik must be one of: poprawne_bez_pomocy, poprawne_z_pomoca_1, poprawne_z_pomoca_2, walk_through")
			}

			today := time.Now().Format("2006-01-02")

			// Get exercise info from data DB
			var typNazwa, tagiJSON string
			err := db().QueryRow("SELECT typ_nazwa, tagi FROM data.cwiczenia WHERE id = ?", id).Scan(&typNazwa, &tagiJSON)
			if err != nil {
				exitError(fmt.Sprintf("exercise %s not found: %v", id, err))
			}

			var tagi []string
			json.Unmarshal([]byte(tagiJSON), &tagi)

			// Record done exercise
			var czasSekVal *int
			if czas > 0 {
				czasSekVal = &czas
			}
			_, err = db().Exec(`INSERT OR REPLACE INTO progress_zrobione (id, typ, data, wynik, czas_sek) VALUES (?, ?, ?, ?, ?)`,
				id, typNazwa, today, wynik, czasSekVal)
			if err != nil {
				exitError(fmt.Sprintf("save progress: %v", err))
			}

			// Update streak and level
			_, err = db().Exec(`INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES (?, 'latwe', 1)
				ON CONFLICT(typ) DO UPDATE SET streak = streak + 1`, typNazwa)
			if err != nil {
				exitError(fmt.Sprintf("update typ: %v", err))
			}

			// Level progression based on streak
			var streak int
			db().QueryRow("SELECT streak FROM progress_typy WHERE typ = ?", typNazwa).Scan(&streak)

			newLevel := calculateLevel(streak, wynik)
			db().Exec("UPDATE progress_typy SET poziom_trudnosci = ? WHERE typ = ?", newLevel, typNazwa)

			// Spaced repetition for tags
			nextReviewDates := updateTags(db(), tagi, wynik)

			// Update session metadata
			db().Exec(`INSERT INTO progress_meta (key, value) VALUES ('sesje', '1')
				ON CONFLICT(key) DO UPDATE SET value = CAST(CAST(value AS INTEGER) + 1 AS TEXT)
				WHERE key = 'sesje' AND NOT EXISTS (
					SELECT 1 FROM progress_meta WHERE key = 'ostatnia_sesja' AND value = ?
				)`, today)
			db().Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('ostatnia_sesja', ?)`, today)
			db().Exec(`INSERT INTO progress_meta (key, value) VALUES ('cwiczenia_lacznie', '1')
				ON CONFLICT(key) DO UPDATE SET value = CAST(CAST(value AS INTEGER) + 1 AS TEXT)`)

			// Session tracking
			incrementSession(db(), typNazwa)

			out := ProgressUpdateOut{
				ID:              id,
				NewLevel:        newLevel,
				Streak:          streak,
				TagsUpdated:     tagi,
				NextReviewDates: nextReviewDates,
			}

			// Time tracking
			if czas > 0 {
				out.CzasSek = &czas
				var benchmarkSek sql.NullInt64
				db().QueryRow("SELECT benchmark_sek FROM data.benchmarki WHERE typ = ?", typNazwa).Scan(&benchmarkSek)
				if benchmarkSek.Valid {
					b := int(benchmarkSek.Int64)
					out.BenchmarkSek = &b
					tempo := calculateTempo(czas, b)
					if tempo != "" {
						out.Tempo = &tempo
					}
				}
			}

			jsonOut(out)
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

func updateTags(d *sql.DB, tags []string, wynik string) []string {
	var dates []string
	for _, tag := range tags {
		var poziom int
		err := d.QueryRow("SELECT poziom FROM progress_tagi WHERE tag = ?", tag).Scan(&poziom)
		if err != nil {
			poziom = 0
		}

		var newPoziom int
		var interval int

		switch wynik {
		case "poprawne_bez_pomocy":
			newPoziom = min(poziom+1, 4)
			interval = intervals[newPoziom]
		case "poprawne_z_pomoca_1":
			newPoziom = min(poziom+1, 4)
			idx := max(newPoziom-1, 0)
			interval = intervals[idx]
		case "poprawne_z_pomoca_2":
			newPoziom = poziom
			interval = intervals[max(newPoziom-1, 0)]
		case "walk_through":
			newPoziom = max(poziom-1, 0)
			interval = 1
		}

		nextReview := time.Now().AddDate(0, 0, interval).Format("2006-01-02")
		d.Exec(`INSERT INTO progress_tagi (tag, poziom, nastepna_powtorka) VALUES (?, ?, ?)
			ON CONFLICT(tag) DO UPDATE SET poziom = ?, nastepna_powtorka = ?`,
			tag, newPoziom, nextReview, newPoziom, nextReview)
		dates = append(dates, nextReview)
	}
	return dates
}

// === progress status ===

func progressStatusCmd() *cobra.Command {
	var typ string

	cmd := &cobra.Command{
		Use:   "status",
		Short: "Show progress dashboard",
		Run: func(cmd *cobra.Command, args []string) {
			out := ProgressStatusOut{}

			// Meta
			var sesjeStr, ostatnia, cwiczeniaStr sql.NullString
			db().QueryRow("SELECT value FROM progress_meta WHERE key = 'sesje'").Scan(&sesjeStr)
			db().QueryRow("SELECT value FROM progress_meta WHERE key = 'ostatnia_sesja'").Scan(&ostatnia)
			db().QueryRow("SELECT value FROM progress_meta WHERE key = 'cwiczenia_lacznie'").Scan(&cwiczeniaStr)

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
				rows, err = db().Query(query, typ)
			} else {
				rows, err = db().Query(query)
			}
			if err != nil {
				exitError(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			for rows.Next() {
				var ts TypStatusOut
				rows.Scan(&ts.Typ, &ts.PoziomTrudnosci, &ts.Streak, &ts.Zrobione, &ts.Dostepne)
				out.PerTyp = append(out.PerTyp, ts)
			}
			if out.PerTyp == nil {
				out.PerTyp = []TypStatusOut{}
			}

			// Enrich per-typ with avg_czas_sek and benchmark
			avgCzasMap := map[string]int{}
			avgRows, _ := db().Query(`SELECT typ, CAST(AVG(czas_sek) AS INTEGER) FROM progress_zrobione WHERE czas_sek IS NOT NULL GROUP BY typ`)
			if avgRows != nil {
				for avgRows.Next() {
					var t string
					var avg int
					avgRows.Scan(&t, &avg)
					avgCzasMap[t] = avg
				}
				avgRows.Close()
			}

			for i := range out.PerTyp {
				if avg, ok := avgCzasMap[out.PerTyp[i].Typ]; ok {
					out.PerTyp[i].AvgCzasSek = &avg
				}
				var benchSek sql.NullInt64
				db().QueryRow("SELECT benchmark_sek FROM data.benchmarki WHERE typ = ?", out.PerTyp[i].Typ).Scan(&benchSek)
				if benchSek.Valid {
					b := int(benchSek.Int64)
					out.PerTyp[i].BenchmarkSek = &b
				}
			}

			// Per-kategoria aggregation
			typKatMap := map[string]string{}
			katRows, _ := db().Query("SELECT DISTINCT typ_nazwa, kategoria FROM data.cwiczenia")
			if katRows != nil {
				for katRows.Next() {
					var t, k string
					katRows.Scan(&t, &k)
					typKatMap[t] = k
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
			db().QueryRow("SELECT COUNT(*) FROM progress_tagi WHERE nastepna_powtorka <= ?", today).Scan(&out.Zaleglosci)

			// Mastered tags (poziom >= 3)
			tagRows, _ := db().Query("SELECT tag FROM progress_tagi WHERE poziom >= 3 ORDER BY tag")
			if tagRows != nil {
				defer tagRows.Close()
				for tagRows.Next() {
					var tag string
					tagRows.Scan(&tag)
					out.TagiOpanowane = append(out.TagiOpanowane, tag)
				}
			}
			if out.TagiOpanowane == nil {
				out.TagiOpanowane = []string{}
			}

			// Problematic tags (poziom <= 1 AND has reviews)
			probRows, _ := db().Query("SELECT tag FROM progress_tagi WHERE poziom <= 1 AND nastepna_powtorka IS NOT NULL ORDER BY tag")
			if probRows != nil {
				defer probRows.Close()
				for probRows.Next() {
					var tag string
					probRows.Scan(&tag)
					out.TagiProblematyczne = append(out.TagiProblematyczne, tag)
				}
			}
			if out.TagiProblematyczne == nil {
				out.TagiProblematyczne = []string{}
			}

			jsonOut(out)
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Filter by type (e.g. sql_group_by)")
	return cmd
}

// === progress blad ===

func progressBladCmd() *cobra.Command {
	var exerciseID, typ, kod, opis string
	var hint int

	cmd := &cobra.Command{
		Use:   "blad",
		Short: "Record an error/mistake for diagnosis",
		Run: func(cmd *cobra.Command, args []string) {
			if exerciseID == "" || typ == "" || kod == "" {
				exitError("--exercise-id, --typ, and --kod are required")
			}

			// Validate exercise exists
			var exists int
			db().QueryRow("SELECT COUNT(*) FROM data.cwiczenia WHERE id = ?", exerciseID).Scan(&exists)
			if exists == 0 {
				exitError(fmt.Sprintf("exercise %s not found", exerciseID))
			}

			today := time.Now().Format("2006-01-02")

			result, err := db().Exec(
				`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, hint_level, data) VALUES (?, ?, ?, ?, ?, ?)`,
				exerciseID, typ, kod, opis, hint, today)
			if err != nil {
				exitError(fmt.Sprintf("save error: %v", err))
			}

			lastID, _ := result.LastInsertId()
			jsonOut(BladOut{
				ID:         int(lastID),
				ExerciseID: exerciseID,
				Typ:        typ,
				BladKod:    kod,
				BladOpis:   opis,
				HintLevel:  hint,
				Data:       today,
			})
		},
	}

	cmd.Flags().StringVar(&exerciseID, "exercise-id", "", "Exercise ID (e.g. 7.3)")
	cmd.Flags().StringVar(&typ, "typ", "", "Exercise type")
	cmd.Flags().StringVar(&kod, "kod", "", "Short error code (e.g. brak_group_by)")
	cmd.Flags().StringVar(&opis, "opis", "", "Full error description")
	cmd.Flags().IntVar(&hint, "hint", 0, "Hint level at which error was identified")
	return cmd
}

// === progress diagnose ===

func progressDiagnoseCmd() *cobra.Command {
	var typ string
	var limit int

	cmd := &cobra.Command{
		Use:   "diagnose",
		Short: "Analyze recurring errors",
		Run: func(cmd *cobra.Command, args []string) {
			// Get total count first
			totalQuery := "SELECT COUNT(*) FROM progress_bledy"
			var totalParams []any
			if typ != "" {
				totalQuery += " WHERE typ = ?"
				totalParams = append(totalParams, typ)
			}
			out := DiagnoseOut{}
			db().QueryRow(totalQuery, totalParams...).Scan(&out.Total)

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

			rows, err := db().Query(query, params...)
			if err != nil {
				exitError(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			for rows.Next() {
				var entry DiagnoseEntry
				var typyStr string
				rows.Scan(&entry.BladKod, &entry.Count, &typyStr, &entry.Ostatnio)
				entry.Typy = strings.Split(typyStr, ",")
				out.TopBledy = append(out.TopBledy, entry)
			}

			if out.TopBledy == nil {
				out.TopBledy = []DiagnoseEntry{}
			}

			jsonOut(out)
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
		Run: func(cmd *cobra.Command, args []string) {
			if typ == "" {
				exitError("--typ is required")
			}

			// Unlock check: require level "trudne" unless --force
			if !force {
				level := getLevel(db(), typ)
				if level != "trudne" {
					exitNotFound(fmt.Sprintf("Sprawdzian typu %s wymaga poziomu trudne. Twoj poziom: %s. Uzyj --force aby pominac.", typ, level))
				}
			}

			query := `SELECT e.id, e.rok, e.numer_zadania, e.tytul, e.kontekst, e.typ_zadania, e.kategoria, e.punkty, e.tresc, e.odpowiedz, e.zasady_oceniania, e.pulapki, e.sciezka_danych, e.pliki_danych
				FROM data.egzamin e
				WHERE e.typ_zadania = ?
				AND e.id NOT IN (SELECT id FROM matura_zrobione)`
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

			rows, err := db().Query(query, params...)
			if err != nil {
				exitError(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			var results []CKEOut
			for rows.Next() {
				var out CKEOut
				var pulapkiJSON, plikiJSON string
				err := rows.Scan(&out.ID, &out.Rok, &out.NumerZadania, &out.Tytul, &out.Kontekst,
					&out.TypZadania, &out.Kategoria, &out.Punkty, &out.Tresc, &out.Odpowiedz,
					&out.ZasadyOceniania, &pulapkiJSON, &out.SciezkaDanych, &plikiJSON)
				if err != nil {
					exitError(fmt.Sprintf("scan error: %v", err))
				}
				json.Unmarshal([]byte(pulapkiJSON), &out.Pulapki)
				json.Unmarshal([]byte(plikiJSON), &out.PlikiDanych)
				if out.Pulapki == nil {
					out.Pulapki = []string{}
				}
				if out.PlikiDanych == nil {
					out.PlikiDanych = []string{}
				}
				results = append(results, out)
			}

			if len(results) == 0 {
				exitNotFound(fmt.Sprintf("no CKE subtasks found for typ=%s", typ))
			}

			chosen := results[rand.Intn(len(results))]
			jsonOut(chosen)
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Task type (e.g. sledzenie_algorytmu)")
	cmd.Flags().StringVar(&exclude, "exclude", "", "Comma-separated IDs to exclude")
	cmd.Flags().BoolVar(&force, "force", false, "Bypass unlock check")
	return cmd
}

// === cke save ===

func ckeSaveCmd() *cobra.Command {
	var id string
	var punkty, maxPunkty int

	cmd := &cobra.Command{
		Use:   "save",
		Short: "Save CKE subtask result",
		Run: func(cmd *cobra.Command, args []string) {
			if id == "" {
				exitError("--id is required")
			}

			var typ string
			err := db().QueryRow("SELECT typ_zadania FROM data.egzamin WHERE id = ?", id).Scan(&typ)
			if err != nil {
				exitError(fmt.Sprintf("CKE subtask %s not found", id))
			}

			today := time.Now().Format("2006-01-02")
			_, err = db().Exec(`INSERT OR REPLACE INTO matura_zrobione (id, typ, data, punkty, max_punkty) VALUES (?, ?, ?, ?, ?)`,
				id, typ, today, punkty, maxPunkty)
			if err != nil {
				exitError(fmt.Sprintf("save error: %v", err))
			}

			jsonOut(map[string]any{
				"id":     id,
				"punkty": punkty,
				"max":    maxPunkty,
				"saved":  true,
			})
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
		Run: func(cmd *cobra.Command, args []string) {
			if rok == 0 {
				exitError("--rok is required")
			}

			rows, err := db().Query(`
				SELECT e.numer_zadania, e.tytul, SUM(e.punkty) as total_pkt
				FROM data.egzamin e
				WHERE e.rok = ?
				GROUP BY e.numer_zadania, e.tytul
				ORDER BY e.numer_zadania`, rok)
			if err != nil {
				exitError(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			var zadania []ExamTaskBrief
			for rows.Next() {
				var tb ExamTaskBrief
				rows.Scan(&tb.Numer, &tb.Tytul, &tb.Punkty)
				tb.Czesc = 1
				zadania = append(zadania, tb)
			}

			if len(zadania) == 0 {
				exitNotFound(fmt.Sprintf("no exam data for rok=%d", rok))
			}

			totalPunkty := 0
			for _, z := range zadania {
				totalPunkty += z.Punkty
			}

			formula := "2023"
			// czas_minuty = 210 for ALL formulas: 2014 (90+120), 2015 (60+150), 2023 (210).
			// Per-part times are in Czesci below.
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
		Run: func(cmd *cobra.Command, args []string) {
			if rok == 0 || zadanie == 0 {
				exitError("--rok and --zadanie are required")
			}

			rows, err := db().Query(`
				SELECT e.id, e.numer_podzadania, e.tytul, e.kontekst, e.typ_zadania, e.kategoria, e.punkty, e.tresc, e.odpowiedz, e.zasady_oceniania, e.pulapki, e.sciezka_danych, e.pliki_danych
				FROM data.egzamin e
				WHERE e.rok = ? AND e.numer_zadania = ?
				ORDER BY e.id`, rok, zadanie)
			if err != nil {
				exitError(fmt.Sprintf("query error: %v", err))
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
					exitError(fmt.Sprintf("scan error: %v", err))
				}

				json.Unmarshal([]byte(pulapkiJSON), &sub.Pulapki)
				if sub.Pulapki == nil {
					sub.Pulapki = []string{}
				}

				if first {
					out.Tytul = tytul
					out.Numer = zadanie
					out.Kontekst = kontekst
					out.SciezkaDanych = sciezkaDanych
					json.Unmarshal([]byte(plikiJSON), &out.PlikiDanych)
					if out.PlikiDanych == nil {
						out.PlikiDanych = []string{}
					}
					first = false
				}

				out.Podzadania = append(out.Podzadania, sub)
				out.Punkty += sub.Punkty
			}

			if first {
				exitNotFound(fmt.Sprintf("no task found for rok=%d zadanie=%d", rok, zadanie))
			}

			jsonOut(out)
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
		Run: func(cmd *cobra.Command, args []string) {
			if rok == 0 || resultsJSON == "" {
				exitError("--rok and --results are required")
			}

			type ResultEntry struct {
				ID  string `json:"id"`
				Pkt int    `json:"pkt"`
				Max int    `json:"max"`
			}

			var results []ResultEntry
			if err := json.Unmarshal([]byte(resultsJSON), &results); err != nil {
				exitError(fmt.Sprintf("invalid --results JSON: %v", err))
			}

			today := time.Now().Format("2006-01-02")
			wynikPkt := 0
			maxPkt := 0

			for _, r := range results {
				wynikPkt += r.Pkt
				maxPkt += r.Max

				var typ string
				db().QueryRow("SELECT typ_zadania FROM data.egzamin WHERE id = ?", r.ID).Scan(&typ)
				db().Exec(`INSERT OR REPLACE INTO matura_zrobione (id, typ, data, punkty, max_punkty) VALUES (?, ?, ?, ?, ?)`,
					r.ID, typ, today, r.Pkt, r.Max)
			}

			procent := 0.0
			if maxPkt > 0 {
				procent = float64(wynikPkt) / float64(maxPkt) * 100
			}

			_, err := db().Exec(`INSERT INTO probne_matury (rok, data, czas_min, wynik_pkt, max_pkt, procent) VALUES (?, ?, ?, ?, ?, ?)`,
				rok, today, czasMin, wynikPkt, maxPkt, procent)
			if err != nil {
				exitError(fmt.Sprintf("save error: %v", err))
			}

			jsonOut(map[string]any{
				"rok":     rok,
				"wynik":   wynikPkt,
				"max":     maxPkt,
				"procent": fmt.Sprintf("%.1f%%", procent),
				"saved":   true,
			})
		},
	}

	cmd.Flags().IntVar(&rok, "rok", 0, "Exam year")
	cmd.Flags().StringVar(&resultsJSON, "results", "", "JSON array of results")
	cmd.Flags().IntVar(&czasMin, "czas", 0, "Time spent in minutes")
	return cmd
}

// === exercise next ===

func exerciseNextCmd() *cobra.Command {
	var typ string

	cmd := &cobra.Command{
		Use:   "next",
		Short: "Get next exercise with smart priority (review > interleave > new)",
		Run: func(cmd *cobra.Command, args []string) {
			if typ == "" {
				exitError("--typ is required")
			}

			sessionCount, _ := getSessionState(db())
			out := ExerciseNextOut{
				SessionCount:   sessionCount,
				ResetSuggested: sessionCount > 0 && sessionCount%16 == 0,
			}

			today := time.Now().Format("2006-01-02")

			// Priority 1: overdue review
			var tag, powtorka string
			err := db().QueryRow(`
				SELECT t.tag, t.nastepna_powtorka
				FROM progress_tagi t
				WHERE t.nastepna_powtorka <= ?
				ORDER BY t.nastepna_powtorka ASC
				LIMIT 1`, today).Scan(&tag, &powtorka)
			if err == nil {
				ex, findErr := findExerciseByTag(db(), tag)
				if findErr == nil {
					powtorkaDate, _ := time.Parse("2006-01-02", powtorka)
					daysOverdue := int(time.Since(powtorkaDate).Hours() / 24)
					out.Mode = "review"
					out.Exercise = ex
					out.ReviewTag = &tag
					out.DaysOverdue = &daysOverdue
					out.PoolWarning = poolWarning(db(), ex.TypNazwa, ex.Trudnosc)
					jsonOut(out)
					return
				}
			}

			// Priority 2: interleaving (every 3rd exercise from a different type)
			if sessionCount > 0 && sessionCount%3 == 0 {
				ex, err := findInterleaveExercise(db(), typ)
				if err == nil {
					out.Mode = "interleave"
					out.Exercise = ex
					out.PoolWarning = poolWarning(db(), ex.TypNazwa, ex.Trudnosc)
					jsonOut(out)
					return
				}
			}

			// Priority 3: new exercise with auto-difficulty
			level := getLevel(db(), typ)
			ex, err := findExerciseByTypAndLevel(db(), typ, level)
			if err != nil {
				// Fallback: any difficulty
				ex, err = findExerciseByTypAndLevel(db(), typ, "")
				if err != nil {
					exitNotFound(fmt.Sprintf("no exercises available for typ=%s", typ))
				}
			}

			out.Mode = "new"
			out.Exercise = ex
			out.PoolWarning = poolWarning(db(), typ, level)
			jsonOut(out)
		},
	}

	cmd.Flags().StringVar(&typ, "typ", "", "Exercise type")
	return cmd
}

// === typ intro ===

func typIntroCmd() *cobra.Command {
	var typ string

	cmd := &cobra.Command{
		Use:   "intro",
		Short: "Get type introduction context",
		Run: func(cmd *cobra.Command, args []string) {
			if typ == "" {
				exitError("--typ is required")
			}

			kat := getKategoria(db(), typ)
			if kat == "" {
				exitNotFound(fmt.Sprintf("unknown type: %s", typ))
			}

			out := TypIntroOut{
				Typ:       typ,
				Kategoria: kat,
			}

			// Check if first in type
			var doneInType int
			db().QueryRow(`SELECT COUNT(*) FROM progress_zrobione z
				JOIN data.cwiczenia c ON c.id = z.id
				WHERE c.typ_nazwa = ?`, typ).Scan(&doneInType)
			out.FirstInType = doneInType == 0
			out.Done = doneInType

			// Check if OTHER types in this category have been done
			otherRows, _ := db().Query(`SELECT DISTINCT c.typ_nazwa FROM progress_zrobione z
				JOIN data.cwiczenia c ON c.id = z.id
				WHERE c.kategoria = ? AND c.typ_nazwa != ?`, kat, typ)
			var otherTypes []string
			if otherRows != nil {
				for otherRows.Next() {
					var t string
					otherRows.Scan(&t)
					otherTypes = append(otherTypes, t)
				}
				otherRows.Close()
			}
			out.FirstInCategory = len(otherTypes) == 0
			if otherTypes == nil {
				otherTypes = []string{}
			}
			out.OtherTypesDoneInCategory = otherTypes

			// Level and streak
			out.Level = getLevel(db(), typ)
			var streak int
			db().QueryRow("SELECT COALESCE(streak, 0) FROM progress_typy WHERE typ = ?", typ).Scan(&streak)
			out.Streak = streak

			// Available exercises
			var available int
			db().QueryRow(`SELECT COUNT(*) FROM data.cwiczenia WHERE typ_nazwa = ?
				AND id NOT IN (SELECT id FROM progress_zrobione)`, typ).Scan(&available)
			out.Available = available

			out.SprawdzianUnlocked = out.Level == "trudne"

			// CKE stats
			ckeNames := exerciseTypToCKETypes(typ, kat)
			var ckeStats CKETypStats
			ckeStats.LatTotal = 11
			err := db().QueryRow(
				`SELECT COUNT(DISTINCT rok), COALESCE(AVG(punkty), 0), COALESCE(SUM(punkty), 0)
				FROM data.egzamin WHERE typ_zadania IN (`+placeholders(len(ckeNames))+`)`,
				toAny(ckeNames)...).
				Scan(&ckeStats.Wystapienia, &ckeStats.AvgPunkty, &ckeStats.TotalPunkty)
			if err == nil && ckeStats.Wystapienia > 0 {
				out.CKEStats = &ckeStats
			}

			// Top pulapki — single pass: count + preserve first-seen casing
			trapRows, _ := db().Query(
				`SELECT pulapki FROM data.egzamin
				WHERE typ_zadania IN (`+placeholders(len(ckeNames))+`)
				AND pulapki != '[]' AND pulapki != 'null'`, toAny(ckeNames)...)
			if trapRows != nil {
				type pulapkaEntry struct {
					Original string
					Count    int
				}
				pulapkiMap := map[string]*pulapkaEntry{} // lowercase -> entry

				for trapRows.Next() {
					var pJSON string
					trapRows.Scan(&pJSON)
					var pulapki []string
					json.Unmarshal([]byte(pJSON), &pulapki)
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
			err = db().QueryRow(
				`SELECT id, tresc, odpowiedz FROM data.cwiczenia
				WHERE typ_nazwa = ? AND trudnosc = 'latwe'
				ORDER BY LENGTH(tresc) ASC LIMIT 1`, typ).
				Scan(&brief.ID, &brief.Tresc, &brief.Odpowiedz)
			if err == nil {
				out.Przyklad = &brief
			}

			// Cheatsheet excerpt
			var fullContent string
			db().QueryRow("SELECT content FROM data.cheatsheets WHERE kategoria = ?", kat).
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

			jsonOut(out)
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
		Run: func(cmd *cobra.Command, args []string) {
			// Get all types with CKE subtasks
			rows, err := db().Query(`
				SELECT e.typ_zadania, e.kategoria, COUNT(*) as total
				FROM data.egzamin e
				GROUP BY e.typ_zadania, e.kategoria
				ORDER BY e.kategoria, e.typ_zadania`)
			if err != nil {
				exitError(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			var results []CKEStatusOut
			for rows.Next() {
				var out CKEStatusOut
				var total int
				rows.Scan(&out.Typ, &out.Kategoria, &total)
				out.CKEAvailable = total

				out.Level = getLevel(db(), out.Typ)
				out.Unlocked = out.Level == "trudne"

				var done int
				db().QueryRow("SELECT COUNT(*) FROM matura_zrobione WHERE typ = ?", out.Typ).Scan(&done)
				out.CKEDone = done
				out.CKEAvailable = total - done

				results = append(results, out)
			}

			if results == nil {
				results = []CKEStatusOut{}
			}
			jsonOut(results)
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
		Run: func(cmd *cobra.Command, args []string) {
			rows, err := db().Query(`
				SELECT rok, SUM(punkty) as total
				FROM data.egzamin
				GROUP BY rok
				ORDER BY rok`)
			if err != nil {
				exitError(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			var entries []ExamListEntry
			for rows.Next() {
				var e ExamListEntry
				rows.Scan(&e.Rok, &e.TotalPkt)

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

				// Check if done (has a non-interrupted mock exam)
				var procent sql.NullFloat64
				db().QueryRow("SELECT procent FROM probne_matury WHERE rok = ? AND przerwany = 0 ORDER BY procent DESC LIMIT 1", e.Rok).Scan(&procent)
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
				exitNotFound("no exams found")
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
		Run: func(cmd *cobra.Command, args []string) {
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

			rows, err := db().Query(query, params...)
			if err != nil {
				exitError(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			var results []TrapOut
			for rows.Next() {
				var t TrapOut
				var pulapkiJSON string
				rows.Scan(&t.SourceID, &t.Rok, &t.Tytul, &pulapkiJSON)
				json.Unmarshal([]byte(pulapkiJSON), &t.Pulapki)
				if len(t.Pulapki) > 0 {
					results = append(results, t)
				}
			}

			if len(results) == 0 {
				exitNotFound("no traps found")
			}

			jsonOut(results)
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
		Run: func(cmd *cobra.Command, args []string) {
			if id == "" || typ == "" {
				exitError("--id and --typ are required")
			}
			if total <= 0 {
				exitError("--total must be > 0")
			}

			today := time.Now().Format("2006-01-02")
			_, err := db().Exec(
				`INSERT OR REPLACE INTO pulapki_przejrzane (id, typ, data, trafienia, total) VALUES (?, ?, ?, ?, ?)`,
				id, typ, today, trafienia, total)
			if err != nil {
				exitError(fmt.Sprintf("save trap result: %v", err))
			}

			jsonOut(TrapSaveOut{ID: id, Typ: typ, Trafienia: trafienia, Total: total})
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
	var kategoria string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get cheatsheet content",
		Run: func(cmd *cobra.Command, args []string) {
			if kategoria == "" {
				exitError("--kategoria is required")
			}

			var content string
			err := db().QueryRow("SELECT content FROM data.cheatsheets WHERE kategoria = ?", kategoria).Scan(&content)
			if err != nil {
				exitNotFound(fmt.Sprintf("cheatsheet for kategoria=%s not found", kategoria))
			}

			fmt.Print(content)
		},
	}

	cmd.Flags().StringVar(&kategoria, "kategoria", "", "Category: TEORIA, IMPLEMENTACJA, ARKUSZ, SQL")
	return cmd
}

// === data import ===

func dataImportCmd() *cobra.Command {
	var source string

	cmd := &cobra.Command{
		Use:   "import",
		Short: "Import JSON data into matura.db",
		Run: func(cmd *cobra.Command, args []string) {
			dataDB, err := OpenDataDB(dbDir)
			if err != nil {
				exitError(fmt.Sprintf("open data DB: %v", err))
			}
			defer dataDB.Close()

			if err := ImportAll(dataDB, source); err != nil {
				exitError(fmt.Sprintf("import failed: %v", err))
			}

			fmt.Fprintln(os.Stderr, "Import complete.")
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
		Run: func(cmd *cobra.Command, args []string) {
			var out DataStatsOut
			db().QueryRow("SELECT COUNT(*) FROM data.cwiczenia").Scan(&out.Cwiczenia)
			db().QueryRow("SELECT COUNT(*) FROM data.egzamin").Scan(&out.Podzadania)
			db().QueryRow("SELECT COUNT(*) FROM data.cheatsheets").Scan(&out.Cheatsheets)
			jsonOut(out)
		},
	}
}
