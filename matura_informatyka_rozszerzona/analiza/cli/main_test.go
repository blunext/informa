package main

import (
	"database/sql"
	"os"
	"path/filepath"
	"testing"
	"time"

	_ "modernc.org/sqlite"
)

// testDir returns a temp dir with imported data for testing.
func testDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()

	// Build data DB
	dataDB, err := OpenDataDB(dir)
	if err != nil {
		t.Fatalf("open data db: %v", err)
	}

	sourceDir := filepath.Join("..", "") // analiza/ is parent of cli/
	if err := ImportAll(dataDB, sourceDir); err != nil {
		t.Fatalf("import: %v", err)
	}
	dataDB.Close()

	return dir
}

// openTestDB opens both DBs (progress + attached data) for testing.
func openTestDB(t *testing.T, dir string) *sql.DB {
	t.Helper()
	db, err := OpenDB(dir)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	t.Cleanup(func() { db.Close() })
	return db
}

func TestImportCounts(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	var exercises, exams, cheatsheets int
	db.QueryRow("SELECT COUNT(*) FROM data.cwiczenia").Scan(&exercises)
	db.QueryRow("SELECT COUNT(*) FROM data.egzamin").Scan(&exams)
	db.QueryRow("SELECT COUNT(*) FROM data.cheatsheets").Scan(&cheatsheets)

	if exercises != 310 {
		t.Errorf("exercises: got %d, want 310", exercises)
	}
	if exams != 230 {
		t.Errorf("exams: got %d, want 230", exams)
	}
	if cheatsheets != 4 {
		t.Errorf("cheatsheets: got %d, want 4", cheatsheets)
	}
}

func TestImportExerciseFields(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	var id, typNazwa, kategoria, trudnosc, zrodlo, tresc, odpowiedz, tagi string
	var punkty int
	err := db.QueryRow(`SELECT id, typ_nazwa, kategoria, trudnosc, punkty, zrodlo, tresc, odpowiedz, tagi
		FROM data.cwiczenia WHERE id = '7.1'`).Scan(&id, &typNazwa, &kategoria, &trudnosc, &punkty, &zrodlo, &tresc, &odpowiedz, &tagi)
	if err != nil {
		t.Fatalf("query 7.1: %v", err)
	}

	if id != "7.1" {
		t.Errorf("id: got %q", id)
	}
	if typNazwa != "cyfry_liczby" {
		t.Errorf("typ_nazwa: got %q", typNazwa)
	}
	if kategoria != "IMPLEMENTACJA" {
		t.Errorf("kategoria: got %q", kategoria)
	}
	if trudnosc != "latwe" {
		t.Errorf("trudnosc: got %q", trudnosc)
	}
	if punkty != 2 {
		t.Errorf("punkty: got %d", punkty)
	}
	if zrodlo == "" {
		t.Error("zrodlo empty")
	}
	if tresc == "" {
		t.Error("tresc empty")
	}
	if odpowiedz == "" {
		t.Error("odpowiedz empty")
	}
	if tagi == "" || tagi == "[]" {
		t.Error("tagi empty")
	}
}

func TestImportExamFields(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	var id, typZadania, kategoria, tresc, odpowiedz string
	var rok, numerZadania, punkty int
	err := db.QueryRow(`SELECT id, rok, numer_zadania, typ_zadania, kategoria, punkty, tresc, odpowiedz
		FROM data.egzamin WHERE id = '2025.1.1'`).Scan(&id, &rok, &numerZadania, &typZadania, &kategoria, &punkty, &tresc, &odpowiedz)
	if err != nil {
		t.Fatalf("query 2025.1.1: %v", err)
	}

	if rok != 2025 {
		t.Errorf("rok: got %d", rok)
	}
	if numerZadania != 1 {
		t.Errorf("numer_zadania: got %d", numerZadania)
	}
	if typZadania != "sledzenie_algorytmu" {
		t.Errorf("typ_zadania: got %q", typZadania)
	}
	if kategoria != "TEORIA" {
		t.Errorf("kategoria: got %q", kategoria)
	}
	if punkty != 3 {
		t.Errorf("punkty: got %d", punkty)
	}
	if tresc == "" {
		t.Error("tresc empty")
	}
}

func TestImportIdempotent(t *testing.T) {
	dir := testDir(t)

	// Import again
	dataDB, err := OpenDataDB(dir)
	if err != nil {
		t.Fatal(err)
	}
	sourceDir := filepath.Join("..", "")
	if err := ImportAll(dataDB, sourceDir); err != nil {
		t.Fatalf("second import: %v", err)
	}
	dataDB.Close()

	db := openTestDB(t, dir)
	var count int
	db.QueryRow("SELECT COUNT(*) FROM data.cwiczenia").Scan(&count)
	if count != 310 {
		t.Errorf("after re-import: got %d exercises, want 310", count)
	}
}

func TestSpacedRepetitionLevels(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	tags := []string{"test-tag-1", "test-tag-2"}

	// Level 0 → poprawne_bez_pomocy → level 1, interval = 1 day
	dates := updateTags(db, tags, "poprawne_bez_pomocy")
	if len(dates) != 2 {
		t.Fatalf("expected 2 dates, got %d", len(dates))
	}

	expected := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	if dates[0] != expected {
		t.Errorf("level 0→1: got %s, want %s", dates[0], expected)
	}

	// Check DB state
	var poziom int
	db.QueryRow("SELECT poziom FROM progress_tagi WHERE tag = 'test-tag-1'").Scan(&poziom)
	if poziom != 1 {
		t.Errorf("poziom after first update: got %d, want 1", poziom)
	}

	// Level 1 → poprawne_bez_pomocy → level 2, interval = 3 days
	dates = updateTags(db, []string{"test-tag-1"}, "poprawne_bez_pomocy")
	expected = time.Now().AddDate(0, 0, 3).Format("2006-01-02")
	if dates[0] != expected {
		t.Errorf("level 1→2: got %s, want %s", dates[0], expected)
	}
	db.QueryRow("SELECT poziom FROM progress_tagi WHERE tag = 'test-tag-1'").Scan(&poziom)
	if poziom != 2 {
		t.Errorf("poziom: got %d, want 2", poziom)
	}

	// Level 2 → poprawne_bez_pomocy → level 3, interval = 7 days
	dates = updateTags(db, []string{"test-tag-1"}, "poprawne_bez_pomocy")
	expected = time.Now().AddDate(0, 0, 7).Format("2006-01-02")
	if dates[0] != expected {
		t.Errorf("level 2→3: got %s, want %s", dates[0], expected)
	}

	// Level 3 → poprawne_bez_pomocy → level 4, interval = 21 days
	dates = updateTags(db, []string{"test-tag-1"}, "poprawne_bez_pomocy")
	expected = time.Now().AddDate(0, 0, 21).Format("2006-01-02")
	if dates[0] != expected {
		t.Errorf("level 3→4: got %s, want %s", dates[0], expected)
	}

	// Level 4 → poprawne_bez_pomocy → stays level 4 (capped), interval = 21
	dates = updateTags(db, []string{"test-tag-1"}, "poprawne_bez_pomocy")
	if dates[0] != expected {
		t.Errorf("level 4→4 (capped): got %s, want %s", dates[0], expected)
	}
	db.QueryRow("SELECT poziom FROM progress_tagi WHERE tag = 'test-tag-1'").Scan(&poziom)
	if poziom != 4 {
		t.Errorf("capped poziom: got %d, want 4", poziom)
	}
}

func TestSpacedRepetitionWalkThrough(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Set up tag at level 3
	db.Exec("INSERT INTO progress_tagi (tag, poziom, nastepna_powtorka) VALUES ('walk-test', 3, '2026-01-01')")

	// walk_through → level drops to 2, interval = 1 day
	dates := updateTags(db, []string{"walk-test"}, "walk_through")
	expected := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	if dates[0] != expected {
		t.Errorf("walk_through interval: got %s, want %s", dates[0], expected)
	}

	var poziom int
	db.QueryRow("SELECT poziom FROM progress_tagi WHERE tag = 'walk-test'").Scan(&poziom)
	if poziom != 2 {
		t.Errorf("walk_through poziom: got %d, want 2", poziom)
	}
}

func TestSpacedRepetitionZPomoca(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Level 0 → z_pomoca_1 → level 1, interval from level 0 = 0 days
	dates := updateTags(db, []string{"pomoc-test"}, "poprawne_z_pomoca_1")
	expected := time.Now().Format("2006-01-02") // interval[0] = 0
	if dates[0] != expected {
		t.Errorf("z_pomoca_1 from 0: got %s, want %s", dates[0], expected)
	}

	// Level 1 → z_pomoca_1 → level 2, interval from level 1 = 1 day
	dates = updateTags(db, []string{"pomoc-test"}, "poprawne_z_pomoca_1")
	expected = time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	if dates[0] != expected {
		t.Errorf("z_pomoca_1 from 1: got %s, want %s", dates[0], expected)
	}
}

func TestProgressMigration(t *testing.T) {
	dir := t.TempDir()

	// Create v1 DB
	progressPath := filepath.Join(dir, "matura_progress.db")
	db, err := sql.Open("sqlite", progressPath)
	if err != nil {
		t.Fatal(err)
	}
	createProgressSchema(db)
	db.Exec("INSERT INTO progress_meta (key, value) VALUES ('test_key', 'test_value')")
	db.Close()

	// Reopen — should detect v1 and pass
	db, err = sql.Open("sqlite", progressPath)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	err = initProgressSchema(db)
	if err != nil {
		t.Fatalf("initProgressSchema: %v", err)
	}

	// Data preserved
	var val string
	db.QueryRow("SELECT value FROM progress_meta WHERE key = 'test_key'").Scan(&val)
	if val != "test_value" {
		t.Errorf("data not preserved after migration: got %q", val)
	}
}

func TestExerciseExclude(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Mark exercise 20.1 as done
	db.Exec("INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('20.1', 'sql_group_by', '2026-01-01', 'poprawne_bez_pomocy')")

	// Query should not return 20.1
	var count int
	db.QueryRow(`SELECT COUNT(*) FROM data.cwiczenia
		WHERE typ_nazwa = 'sql_group_by'
		AND id NOT IN (SELECT id FROM progress_zrobione)`).Scan(&count)

	if count != 29 { // 30 total - 1 done = 29
		t.Errorf("expected 29 available, got %d", count)
	}
}

func TestExamYearCoverage(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	years := []int{2014, 2015, 2016, 2017, 2018, 2019, 2021, 2022, 2023, 2024, 2025}
	expectedCounts := map[int]int{
		2014: 22, 2015: 19, 2016: 20, 2017: 21, 2018: 20,
		2019: 20, 2021: 22, 2022: 20, 2023: 23, 2024: 22, 2025: 21,
	}

	for _, year := range years {
		var count int
		db.QueryRow("SELECT COUNT(*) FROM data.egzamin WHERE rok = ?", year).Scan(&count)
		if count != expectedCounts[year] {
			t.Errorf("year %d: got %d subtasks, want %d", year, count, expectedCounts[year])
		}
	}
}

func TestCheatsheetCategories(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	categories := []string{"TEORIA", "IMPLEMENTACJA", "ARKUSZ", "SQL"}
	for _, cat := range categories {
		var content string
		err := db.QueryRow("SELECT content FROM data.cheatsheets WHERE kategoria = ?", cat).Scan(&content)
		if err != nil {
			t.Errorf("cheatsheet %s: %v", cat, err)
		}
		if len(content) < 100 {
			t.Errorf("cheatsheet %s: too short (%d bytes)", cat, len(content))
		}
	}
}

func TestProgressSchemaCreation(t *testing.T) {
	dir := t.TempDir()

	// First open — should create schema
	db, err := OpenDB(dir)
	if err != nil {
		// Expected: no matura.db, but progress schema should be created
		// Actually OpenDB doesn't fail if matura.db is missing — it just doesn't attach.
		t.Fatalf("first open: %v", err)
	}
	defer db.Close()

	// Verify tables exist
	tables := []string{"schema_version", "progress_meta", "progress_typy", "progress_zrobione", "progress_tagi", "matura_zrobione", "probne_matury"}
	for _, table := range tables {
		var count int
		err := db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name=?", table).Scan(&count)
		if err != nil || count != 1 {
			t.Errorf("table %s not created", table)
		}
	}
}

func TestDataImportCreatesDBFile(t *testing.T) {
	dir := testDir(t)

	dbPath := filepath.Join(dir, "matura.db")
	info, err := os.Stat(dbPath)
	if err != nil {
		t.Fatalf("matura.db not created: %v", err)
	}
	if info.Size() < 1000 {
		t.Errorf("matura.db too small: %d bytes", info.Size())
	}
}
