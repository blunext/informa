package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"regexp"
	"sort"
	"strconv"
	"strings"
)

// ImportAll runs all importers: exercises, exams, cheatsheets.
func ImportAll(db *sql.DB, sourceDir string) error {
	if err := CreateDataSchema(db); err != nil {
		return fmt.Errorf("create schema: %w", err)
	}

	nEx, err := ImportExercises(db, sourceDir)
	if err != nil {
		return fmt.Errorf("import exercises: %w", err)
	}
	fmt.Fprintf(os.Stderr, "Imported %d exercises\n", nEx)

	nCKE, err := ImportExams(db, sourceDir)
	if err != nil {
		return fmt.Errorf("import exams: %w", err)
	}
	fmt.Fprintf(os.Stderr, "Imported %d CKE subtasks\n", nCKE)

	nCS, err := ImportCheatsheets(db, sourceDir)
	if err != nil {
		return fmt.Errorf("import cheatsheets: %w", err)
	}
	fmt.Fprintf(os.Stderr, "Imported %d cheatsheets\n", nCS)

	nBench, err := ImportBenchmarks(db, sourceDir)
	if err != nil {
		return fmt.Errorf("import benchmarks: %w", err)
	}
	fmt.Fprintf(os.Stderr, "Imported %d benchmarks\n", nBench)

	return nil
}

// ImportExercises reads cwiczenia/json/NN_*/ directories and inserts into cwiczenia table.
func ImportExercises(db *sql.DB, sourceDir string) (int, error) {
	jsonDir := filepath.Join(sourceDir, "cwiczenia", "json")

	entries, err := os.ReadDir(jsonDir)
	if err != nil {
		return 0, fmt.Errorf("read exercise dir: %w", err)
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO cwiczenia (id, typ_nr, typ_nazwa, kategoria, trudnosc, punkty, zrodlo, tresc, odpowiedz, wskazowki, typowe_bledy, tagi)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	count := 0
	dirPattern := regexp.MustCompile(`^(\d{2})_(.+)$`)

	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		matches := dirPattern.FindStringSubmatch(entry.Name())
		if matches == nil {
			continue
		}

		typNr, _ := strconv.Atoi(matches[1])
		typNazwa := matches[2]

		dirPath := filepath.Join(jsonDir, entry.Name())

		// Read _meta.json for kategoria
		metaPath := filepath.Join(dirPath, "_meta.json")
		metaData, err := os.ReadFile(metaPath)
		if err != nil {
			return 0, fmt.Errorf("read %s: %w", metaPath, err)
		}
		var meta ExerciseMeta
		if err := json.Unmarshal(metaData, &meta); err != nil {
			return 0, fmt.Errorf("parse %s: %w", metaPath, err)
		}

		// Read each exercise file
		files, err := os.ReadDir(dirPath)
		if err != nil {
			return 0, err
		}

		for _, f := range files {
			if f.Name() == "_meta.json" || f.IsDir() || !strings.HasSuffix(f.Name(), ".json") {
				continue
			}

			exPath := filepath.Join(dirPath, f.Name())
			exData, err := os.ReadFile(exPath)
			if err != nil {
				return 0, fmt.Errorf("read %s: %w", exPath, err)
			}

			var ex ExerciseJSON
			if err := json.Unmarshal(exData, &ex); err != nil {
				return 0, fmt.Errorf("parse %s: %w", exPath, err)
			}

			wskazowkiJSON, _ := json.Marshal(ex.Wskazowki)
			typoweBledyJSON, _ := json.Marshal(ex.TypoweBledy)
			tagiJSON, _ := json.Marshal(ex.Tagi)

			_, err = stmt.Exec(
				ex.ID, typNr, typNazwa, meta.Kategoria,
				ex.Trudnosc, ex.Punkty, ex.Zrodlo,
				ex.Tresc, ex.Odpowiedz,
				string(wskazowkiJSON), string(typoweBledyJSON), string(tagiJSON),
			)
			if err != nil {
				return 0, fmt.Errorf("insert exercise %s: %w", ex.ID, err)
			}
			count++
		}
	}

	return count, tx.Commit()
}

// VerifyExercises compares JSON files on disk with exercises in matura.db.
func VerifyExercises(db *sql.DB, sourceDir string) (*VerifyOut, error) {
	jsonDir := filepath.Join(sourceDir, "cwiczenia", "json")
	out := &VerifyOut{}

	// 1. Read all exercises from disk
	diskExercises := map[string]*ExerciseJSON{}
	entries, err := os.ReadDir(jsonDir)
	if err != nil {
		return nil, fmt.Errorf("read json dir: %w", err)
	}

	dirPattern := regexp.MustCompile(`^(\d{2})_(.+)$`)
	for _, entry := range entries {
		if !entry.IsDir() {
			continue
		}
		if dirPattern.FindStringSubmatch(entry.Name()) == nil {
			continue
		}
		dirPath := filepath.Join(jsonDir, entry.Name())
		files, err := os.ReadDir(dirPath)
		if err != nil {
			return nil, err
		}
		for _, f := range files {
			if f.Name() == "_meta.json" || f.IsDir() || !strings.HasSuffix(f.Name(), ".json") {
				continue
			}
			data, err := os.ReadFile(filepath.Join(dirPath, f.Name()))
			if err != nil {
				return nil, err
			}
			var ex ExerciseJSON
			if err := json.Unmarshal(data, &ex); err != nil {
				return nil, fmt.Errorf("parse %s/%s: %w", entry.Name(), f.Name(), err)
			}
			diskExercises[ex.ID] = &ex
		}
	}
	out.TotalDisk = len(diskExercises)

	// 2. Read all exercises from DB
	rows, err := db.Query("SELECT id, trudnosc, punkty, tresc, odpowiedz FROM cwiczenia")
	if err != nil {
		return nil, fmt.Errorf("query db: %w", err)
	}
	defer rows.Close()

	dbIDs := map[string]bool{}
	for rows.Next() {
		var id, trudnosc, tresc, odpowiedz string
		var punkty int
		if err := rows.Scan(&id, &trudnosc, &punkty, &tresc, &odpowiedz); err != nil {
			return nil, err
		}
		dbIDs[id] = true

		disk, ok := diskExercises[id]
		if !ok {
			out.MissingOnDisk = append(out.MissingOnDisk, id)
			continue
		}

		// Compare key fields
		mismatches := []string{}
		if disk.Trudnosc != trudnosc {
			mismatches = append(mismatches, fmt.Sprintf("trudnosc: disk=%s db=%s", disk.Trudnosc, trudnosc))
		}
		if disk.Punkty != punkty {
			mismatches = append(mismatches, fmt.Sprintf("punkty: disk=%d db=%d", disk.Punkty, punkty))
		}
		if len(disk.Tresc) != len(tresc) {
			mismatches = append(mismatches, fmt.Sprintf("tresc length: disk=%d db=%d", len(disk.Tresc), len(tresc)))
		}
		if len(disk.Odpowiedz) != len(odpowiedz) {
			mismatches = append(mismatches, fmt.Sprintf("odpowiedz length: disk=%d db=%d", len(disk.Odpowiedz), len(odpowiedz)))
		}
		if len(mismatches) > 0 {
			out.Mismatched = append(out.Mismatched, id+": "+strings.Join(mismatches, ", "))
		} else {
			out.Matched++
		}
	}
	out.TotalDB = len(dbIDs)

	// 3. Find exercises on disk but missing in DB
	for id := range diskExercises {
		if !dbIDs[id] {
			out.MissingInDB = append(out.MissingInDB, id)
		}
	}

	sort.Strings(out.Mismatched)
	sort.Strings(out.MissingInDB)
	sort.Strings(out.MissingOnDisk)

	return out, nil
}

// ImportExams reads json/matura_YYYY.json files and inserts flattened subtasks into egzamin table.
func ImportExams(db *sql.DB, sourceDir string) (int, error) {
	jsonDir := filepath.Join(sourceDir, "json")

	pattern := filepath.Join(jsonDir, "matura_*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return 0, err
	}

	// Filter out matura_indeks.json
	var examFiles []string
	for _, f := range files {
		base := filepath.Base(f)
		if base == "matura_indeks.json" {
			continue
		}
		examFiles = append(examFiles, f)
	}

	sort.Strings(examFiles)

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	stmt, err := tx.Prepare(`INSERT INTO egzamin (id, rok, numer_zadania, numer_podzadania, tytul, kontekst, tresc, odpowiedz, zasady_oceniania, typ_zadania, kategoria, punkty, czesc, pulapki, sciezka_danych, pliki_danych)
		VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
	if err != nil {
		return 0, err
	}
	defer stmt.Close()

	count := 0
	for _, f := range examFiles {
		data, err := os.ReadFile(f)
		if err != nil {
			return 0, fmt.Errorf("read %s: %w", f, err)
		}

		var exam MaturaExam
		if err := json.Unmarshal(data, &exam); err != nil {
			return 0, fmt.Errorf("parse %s: %w", f, err)
		}

		for _, task := range exam.Zadania {
			kontekst := ""
			if task.Kontekst != nil {
				kontekst = *task.Kontekst
			}
			sciezkaDanych := ""
			if task.SciezkaDanych != nil {
				sciezkaDanych = *task.SciezkaDanych
			}
			plikiDanychJSON, _ := json.Marshal(task.PlikiDanych)

			for _, sub := range task.Podzadania {
				pulapkiJSON, _ := json.Marshal(sub.Pulapki)

				_, err = stmt.Exec(
					sub.ID, exam.Rok, task.Numer, sub.Numer,
					task.Tytul, kontekst,
					sub.Tresc, sub.Odpowiedz, sub.ZasadyOceniania,
					sub.TypZadania, sub.Kategoria, sub.Punkty,
					task.Czesc,
					string(pulapkiJSON), sciezkaDanych, string(plikiDanychJSON),
				)
				if err != nil {
					return 0, fmt.Errorf("insert subtask %s: %w", sub.ID, err)
				}
				count++
			}
		}
	}

	return count, tx.Commit()
}

// ckeTypToExerciseTyp converts CKE typ_zadania names to exercise typ_nazwa.
func ckeTypToExerciseTyp(ckeTyp string) string {
	for _, p := range []string{"arkusz_", "przetwarzanie_", "obliczenia_"} {
		if strings.HasPrefix(ckeTyp, p) {
			return strings.TrimPrefix(ckeTyp, p)
		}
	}
	return ckeTyp
}

// ImportBenchmarks reads matura_*.json files, extracts wymiary.czasowy_min,
// distributes time proportionally to subtasks by punkty, and aggregates per typ.
func ImportBenchmarks(db *sql.DB, sourceDir string) (int, error) {
	jsonDir := filepath.Join(sourceDir, "json")

	pattern := filepath.Join(jsonDir, "matura_*.json")
	files, err := filepath.Glob(pattern)
	if err != nil {
		return 0, err
	}

	// Collect per-typ seconds
	benchmarks := map[string][]int{} // exercise typ_nazwa -> []seconds

	for _, f := range files {
		base := filepath.Base(f)
		if base == "matura_indeks.json" {
			continue
		}

		data, err := os.ReadFile(f)
		if err != nil {
			return 0, fmt.Errorf("read %s: %w", f, err)
		}

		var exam MaturaExam
		if err := json.Unmarshal(data, &exam); err != nil {
			return 0, fmt.Errorf("parse %s: %w", f, err)
		}

		for _, task := range exam.Zadania {
			if task.Wymiary == nil || task.Wymiary.CzasowyMin == 0 {
				continue
			}

			taskTotalPkt := 0
			for _, sub := range task.Podzadania {
				taskTotalPkt += sub.Punkty
			}
			if taskTotalPkt == 0 {
				continue
			}

			for _, sub := range task.Podzadania {
				subSek := task.Wymiary.CzasowyMin * 60 * sub.Punkty / taskTotalPkt
				exTyp := ckeTypToExerciseTyp(sub.TypZadania)
				benchmarks[exTyp] = append(benchmarks[exTyp], subSek)
			}
		}
	}

	// Insert averages in a transaction
	tx, err := db.Begin()
	if err != nil {
		return 0, fmt.Errorf("begin benchmark tx: %w", err)
	}
	defer tx.Rollback()

	count := 0
	for typ, secs := range benchmarks {
		total := 0
		for _, s := range secs {
			total += s
		}
		avg := total / len(secs)

		_, err := tx.Exec("INSERT OR REPLACE INTO benchmarki (typ, benchmark_sek, zrodlo) VALUES (?, ?, 'cke_avg')",
			typ, avg)
		if err != nil {
			return 0, fmt.Errorf("insert benchmark %s: %w", typ, err)
		}
		count++
	}

	return count, tx.Commit()
}

// ImportCheatsheets reads cheatsheet MD files into cheatsheets table.
func ImportCheatsheets(db *sql.DB, sourceDir string) (int, error) {
	cheatsheetDir := filepath.Join(sourceDir, "cheatsheets")

	// Map filename suffix → kategoria
	mapping := map[string]string{
		"cheatsheet_cpp.md":    "IMPLEMENTACJA",
		"cheatsheet_sql.md":    "SQL",
		"cheatsheet_teoria.md": "TEORIA",
		"cheatsheet_arkusz.md": "ARKUSZ",
	}

	tx, err := db.Begin()
	if err != nil {
		return 0, err
	}
	defer tx.Rollback()

	count := 0
	for filename, kategoria := range mapping {
		path := filepath.Join(cheatsheetDir, filename)
		content, err := os.ReadFile(path)
		if err != nil {
			return 0, fmt.Errorf("read %s: %w", path, err)
		}

		_, err = tx.Exec("INSERT INTO cheatsheets (kategoria, content) VALUES (?, ?)", kategoria, string(content))
		if err != nil {
			return 0, fmt.Errorf("insert cheatsheet %s: %w", kategoria, err)
		}
		count++
	}

	return count, tx.Commit()
}
