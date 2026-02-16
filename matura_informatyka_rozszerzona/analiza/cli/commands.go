package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"math/rand"
	"os"
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

			query := `SELECT c.id, c.typ_nazwa, c.kategoria, c.trudnosc, c.punkty, c.zrodlo, c.tagi, c.tresc, c.wskazowki, c.odpowiedz, c.typowe_bledy
				FROM data.cwiczenia c
				WHERE c.typ_nazwa = ?`
			params := []any{typ}

			if trudnosc != "" {
				query += " AND c.trudnosc = ?"
				params = append(params, trudnosc)
			}

			// Exclude already done exercises
			query += ` AND c.id NOT IN (SELECT id FROM progress_zrobione)`

			if exclude != "" {
				excludeIDs := strings.Split(exclude, ",")
				placeholders := make([]string, len(excludeIDs))
				for i, id := range excludeIDs {
					placeholders[i] = "?"
					params = append(params, strings.TrimSpace(id))
				}
				query += " AND c.id NOT IN (" + strings.Join(placeholders, ",") + ")"
			}

			rows, err := db().Query(query, params...)
			if err != nil {
				exitError(fmt.Sprintf("query error: %v", err))
			}
			defer rows.Close()

			var results []ExerciseOut
			for rows.Next() {
				var ex ExerciseOut
				var tagiJSON, wskazowkiJSON, typoweBledyJSON string
				err := rows.Scan(&ex.ID, &ex.TypNazwa, &ex.Kategoria, &ex.Trudnosc, &ex.Punkty, &ex.Zrodlo,
					&tagiJSON, &ex.Tresc, &wskazowkiJSON, &ex.Odpowiedz, &typoweBledyJSON)
				if err != nil {
					exitError(fmt.Sprintf("scan error: %v", err))
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
				results = append(results, ex)
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

				// Find an exercise that has this tag and isn't done
				var ex ExerciseOut
				var tagiJSON, wskazowkiJSON, typoweBledyJSON string
				err := db().QueryRow(`
					SELECT c.id, c.typ_nazwa, c.kategoria, c.trudnosc, c.punkty, c.zrodlo, c.tagi, c.tresc, c.wskazowki, c.odpowiedz, c.typowe_bledy
					FROM data.cwiczenia c
					WHERE c.tagi LIKE ?
					AND c.id NOT IN (SELECT id FROM progress_zrobione)
					ORDER BY RANDOM() LIMIT 1`,
					"%"+tag+"%",
				).Scan(&ex.ID, &ex.TypNazwa, &ex.Kategoria, &ex.Trudnosc, &ex.Punkty, &ex.Zrodlo,
					&tagiJSON, &ex.Tresc, &wskazowkiJSON, &ex.Odpowiedz, &typoweBledyJSON)

				if err == sql.ErrNoRows {
					// Try with done exercises too
					err = db().QueryRow(`
						SELECT c.id, c.typ_nazwa, c.kategoria, c.trudnosc, c.punkty, c.zrodlo, c.tagi, c.tresc, c.wskazowki, c.odpowiedz, c.typowe_bledy
						FROM data.cwiczenia c
						WHERE c.tagi LIKE ?
						ORDER BY RANDOM() LIMIT 1`,
						"%"+tag+"%",
					).Scan(&ex.ID, &ex.TypNazwa, &ex.Kategoria, &ex.Trudnosc, &ex.Punkty, &ex.Zrodlo,
						&tagiJSON, &ex.Tresc, &wskazowkiJSON, &ex.Odpowiedz, &typoweBledyJSON)
				}

				if err != nil {
					continue
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
			_, err = db().Exec(`INSERT OR REPLACE INTO progress_zrobione (id, typ, data, wynik) VALUES (?, ?, ?, ?)`,
				id, typNazwa, today, wynik)
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

			out := ProgressUpdateOut{
				ID:              id,
				NewLevel:        newLevel,
				Streak:          streak,
				TagsUpdated:     tagi,
				NextReviewDates: nextReviewDates,
			}
			jsonOut(out)
		},
	}

	cmd.Flags().StringVar(&id, "id", "", "Exercise ID (e.g. 7.3)")
	cmd.Flags().StringVar(&wynik, "wynik", "", "Result: poprawne_bez_pomocy, poprawne_z_pomoca_1, poprawne_z_pomoca_2, walk_through")
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

// === cke get ===

func ckeGetCmd() *cobra.Command {
	var typ, exclude string

	cmd := &cobra.Command{
		Use:   "get",
		Short: "Get a random CKE subtask",
		Run: func(cmd *cobra.Command, args []string) {
			if typ == "" {
				exitError("--typ is required")
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
