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

func TestGetLevel(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Fresh DB → "latwe"
	if got := getLevel(db, "cyfry_liczby"); got != "latwe" {
		t.Errorf("fresh DB: got %q, want latwe", got)
	}

	// Insert level
	db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'srednie', 3)")
	if got := getLevel(db, "cyfry_liczby"); got != "srednie" {
		t.Errorf("after insert: got %q, want srednie", got)
	}

	// Unknown type → "latwe"
	if got := getLevel(db, "nonexistent_type"); got != "latwe" {
		t.Errorf("unknown type: got %q, want latwe", got)
	}
}

func TestGetKategoria(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	cases := map[string]string{
		"sql_group_by":        "SQL",
		"cyfry_liczby":        "IMPLEMENTACJA",
		"sledzenie_algorytmu": "TEORIA",
		"agregacja_warunkowa": "ARKUSZ",
	}
	for typ, want := range cases {
		if got := getKategoria(db, typ); got != want {
			t.Errorf("getKategoria(%s): got %q, want %q", typ, got, want)
		}
	}

	if got := getKategoria(db, "nonexistent"); got != "" {
		t.Errorf("unknown type: got %q, want empty", got)
	}
}

func TestSessionTracking(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Fresh → count=0
	count, lastTyp := getSessionState(db)
	if count != 0 {
		t.Errorf("fresh count: got %d, want 0", count)
	}
	if lastTyp != "" {
		t.Errorf("fresh lastTyp: got %q, want empty", lastTyp)
	}

	// Increment
	incrementSession(db, "sql_group_by")
	count, lastTyp = getSessionState(db)
	if count != 1 {
		t.Errorf("after 1st increment: got %d, want 1", count)
	}
	if lastTyp != "sql_group_by" {
		t.Errorf("after 1st increment: got %q, want sql_group_by", lastTyp)
	}

	// Increment again
	incrementSession(db, "cyfry_liczby")
	count, lastTyp = getSessionState(db)
	if count != 2 {
		t.Errorf("after 2nd increment: got %d, want 2", count)
	}
	if lastTyp != "cyfry_liczby" {
		t.Errorf("after 2nd increment: got %q, want cyfry_liczby", lastTyp)
	}
}

func TestAutodifficulty(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Set mainDB for queryExercises
	mainDB = db

	// Fresh DB → getLevel returns "latwe"
	results := queryExercises("cyfry_liczby", "latwe", "")
	if len(results) == 0 {
		t.Fatal("no latwe exercises for cyfry_liczby")
	}
	for _, r := range results {
		if r.Trudnosc != "latwe" {
			t.Errorf("got trudnosc %q, want latwe", r.Trudnosc)
		}
	}

	// Set level to srednie
	db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'srednie', 3)")
	if got := getLevel(db, "cyfry_liczby"); got != "srednie" {
		t.Errorf("level after insert: got %q, want srednie", got)
	}

	// Query srednie exercises
	results = queryExercises("cyfry_liczby", "srednie", "")
	for _, r := range results {
		if r.Trudnosc != "srednie" {
			t.Errorf("got trudnosc %q, want srednie", r.Trudnosc)
		}
	}
}

func TestPoolWarning(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// sql_group_by has many exercises → no warning
	if w := poolWarning(db, "sql_group_by", "latwe"); w != nil {
		t.Errorf("sql_group_by latwe: expected nil warning, got %q", *w)
	}

	// Mark all but 2 teoria_bezpieczenstwa/latwe as done
	rows, _ := db.Query(`SELECT id FROM data.cwiczenia WHERE typ_nazwa = 'teoria_bezpieczenstwa' AND trudnosc = 'latwe'`)
	var ids []string
	for rows.Next() {
		var id string
		rows.Scan(&id)
		ids = append(ids, id)
	}
	rows.Close()

	// Mark all but last 2 as done
	for i := 0; i < len(ids)-2; i++ {
		db.Exec("INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES (?, 'teoria_bezpieczenstwa', '2026-01-01', 'poprawne_bez_pomocy')", ids[i])
	}

	if w := poolWarning(db, "teoria_bezpieczenstwa", "latwe"); w == nil {
		t.Error("teoria_bezpieczenstwa latwe: expected warning, got nil")
	}
}

func TestTypIntro(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	mainDB = db

	// Fresh DB: sql_group_by → first_in_type=true, first_in_category=true
	// We need to simulate the typ intro logic directly
	kat := getKategoria(db, "sql_group_by")
	if kat != "SQL" {
		t.Fatalf("sql_group_by kategoria: got %q, want SQL", kat)
	}

	var doneInType int
	db.QueryRow(`SELECT COUNT(*) FROM progress_zrobione z
		JOIN data.cwiczenia c ON c.id = z.id
		WHERE c.typ_nazwa = 'sql_group_by'`).Scan(&doneInType)
	if doneInType != 0 {
		t.Errorf("fresh DB doneInType: got %d, want 0", doneInType)
	}

	// Mark 1 exercise from sql_group_by as done
	db.Exec("INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('20.1', 'sql_group_by', '2026-01-01', 'poprawne_bez_pomocy')")

	// Now check sql_join: first_in_type=true, first_in_category=false (sql_group_by done)
	var doneJoin int
	db.QueryRow(`SELECT COUNT(*) FROM progress_zrobione z
		JOIN data.cwiczenia c ON c.id = z.id
		WHERE c.typ_nazwa = 'sql_join'`).Scan(&doneJoin)
	if doneJoin != 0 {
		t.Errorf("sql_join doneInType: got %d, want 0", doneJoin)
	}

	otherRows, _ := db.Query(`SELECT DISTINCT c.typ_nazwa FROM progress_zrobione z
		JOIN data.cwiczenia c ON c.id = z.id
		WHERE c.kategoria = 'SQL' AND c.typ_nazwa != 'sql_join'`)
	var otherTypes []string
	for otherRows.Next() {
		var tt string
		otherRows.Scan(&tt)
		otherTypes = append(otherTypes, tt)
	}
	otherRows.Close()

	if len(otherTypes) == 0 {
		t.Error("expected other types done in SQL category")
	}
	found := false
	for _, tt := range otherTypes {
		if tt == "sql_group_by" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected sql_group_by in other types, got %v", otherTypes)
	}
}

func TestCKEUnlockCheck(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Fresh → not "trudne"
	level := getLevel(db, "sledzenie_algorytmu")
	if level == "trudne" {
		t.Error("fresh DB: expected non-trudne level")
	}

	// Set to trudne
	db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('sledzenie_algorytmu', 'trudne', 8)")
	level = getLevel(db, "sledzenie_algorytmu")
	if level != "trudne" {
		t.Errorf("after insert: got %q, want trudne", level)
	}
}

func TestExamList(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Check all 11 years are in data.egzamin
	var yearCount int
	db.QueryRow("SELECT COUNT(DISTINCT rok) FROM data.egzamin").Scan(&yearCount)
	if yearCount != 11 {
		t.Errorf("years: got %d, want 11", yearCount)
	}

	// No mock exams → none Done
	var doneCount int
	db.QueryRow("SELECT COUNT(*) FROM probne_matury WHERE przerwany = 0").Scan(&doneCount)
	if doneCount != 0 {
		t.Errorf("fresh done count: got %d, want 0", doneCount)
	}

	// Insert a mock exam for 2024
	db.Exec("INSERT INTO probne_matury (rok, data, czas_min, wynik_pkt, max_pkt, procent, przerwany) VALUES (2024, '2026-01-15', 180, 35, 50, 70.0, 0)")

	var procent float64
	err := db.QueryRow("SELECT procent FROM probne_matury WHERE rok = 2024 AND przerwany = 0 ORDER BY procent DESC LIMIT 1").Scan(&procent)
	if err != nil {
		t.Fatalf("query procent: %v", err)
	}
	if procent != 70.0 {
		t.Errorf("procent: got %.1f, want 70.0", procent)
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
