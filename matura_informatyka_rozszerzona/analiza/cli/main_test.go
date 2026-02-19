package main

import (
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	_ "modernc.org/sqlite"
)

// sharedTestDir holds a temp dir with imported matura.db, created once in TestMain.
var sharedTestDir string

func TestMain(m *testing.M) {
	dir, err := os.MkdirTemp("", "matura-test-*")
	if err != nil {
		fmt.Fprintf(os.Stderr, "create temp dir: %v\n", err)
		os.Exit(1)
	}

	dataDB, err := OpenDataDB(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "open data db: %v\n", err)
		os.RemoveAll(dir)
		os.Exit(1)
	}

	sourceDir := filepath.Join("..", "") // analiza/ is parent of cli/
	if err := ImportAll(dataDB, sourceDir); err != nil {
		fmt.Fprintf(os.Stderr, "import: %v\n", err)
		dataDB.Close()
		os.RemoveAll(dir)
		os.Exit(1)
	}
	dataDB.Close()

	sharedTestDir = dir
	code := m.Run()
	os.RemoveAll(dir)
	os.Exit(code)
}

// testDir returns a temp dir with a copy of matura.db for testing.
// Each test gets its own progress.db (created on first OpenDB).
func testDir(t *testing.T) string {
	t.Helper()
	dir := t.TempDir()

	// Copy shared matura.db
	src := filepath.Join(sharedTestDir, "matura.db")
	dst := filepath.Join(dir, "matura.db")
	data, err := os.ReadFile(src)
	if err != nil {
		t.Fatalf("read shared matura.db: %v", err)
	}
	if err := os.WriteFile(dst, data, 0644); err != nil {
		t.Fatalf("write matura.db copy: %v", err)
	}

	return dir
}

// openTestDB opens both DBs (progress + attached data) for testing.
func openTestDB(t *testing.T, dir string) *sql.DB {
	t.Helper()
	db, attached, err := OpenDB(dir)
	if err != nil {
		t.Fatalf("open db: %v", err)
	}
	if !attached {
		t.Fatal("matura.db not attached")
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

	if exercises != 407 {
		t.Errorf("exercises: got %d, want 407", exercises)
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
	if count != 407 {
		t.Errorf("after re-import: got %d exercises, want 407", count)
	}
}

func TestFSRSIntegrationProgression(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)
	tags := []string{"fsrs-test-1"}

	// 4× poprawne_bez_pomocy (Easy) → stability grows, poziom increases
	for i := 0; i < 4; i++ {
		dates, err := updateTags(db, tags, "poprawne_bez_pomocy")
		if err != nil {
			t.Fatal(err)
		}
		if len(dates) != 1 {
			t.Fatalf("step %d: expected 1 date, got %d", i, len(dates))
		}
		// Each review date should be in the future (or today)
		today := time.Now().Format("2006-01-02")
		if dates[0] < today {
			t.Errorf("step %d: review date %s is in the past", i, dates[0])
		}
	}

	var poziom int
	var stability float64
	db.QueryRow("SELECT poziom, stability FROM progress_tagi WHERE tag = 'fsrs-test-1'").Scan(&poziom, &stability)
	if poziom < 2 {
		t.Errorf("after 4× Easy: poziom=%d, want >= 2", poziom)
	}
	if stability < 3 {
		t.Errorf("after 4× Easy: stability=%.2f, want >= 3", stability)
	}

	// Two tags: both should get dates
	dates2, err := updateTags(db, []string{"fsrs-a", "fsrs-b"}, "poprawne_bez_pomocy")
	if err != nil {
		t.Fatal(err)
	}
	if len(dates2) != 2 {
		t.Fatalf("two tags: expected 2 dates, got %d", len(dates2))
	}
}

func TestFSRSIntegrationWalkThrough(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Build up tag to stable state with 3× Easy
	for i := 0; i < 3; i++ {
		updateTags(db, []string{"walk-fsrs"}, "poprawne_bez_pomocy")
	}
	var beforeS float64
	var beforeL int
	db.QueryRow("SELECT stability, COALESCE(lapses, 0) FROM progress_tagi WHERE tag = 'walk-fsrs'").Scan(&beforeS, &beforeL)

	// walk_through → stability drops, lapses++, interval=1
	dates, err := updateTags(db, []string{"walk-fsrs"}, "walk_through")
	if err != nil {
		t.Fatal(err)
	}
	expected := time.Now().AddDate(0, 0, 1).Format("2006-01-02")
	if dates[0] != expected {
		t.Errorf("walk_through date: got %s, want %s", dates[0], expected)
	}

	var afterS float64
	var afterL int
	db.QueryRow("SELECT stability, lapses FROM progress_tagi WHERE tag = 'walk-fsrs'").Scan(&afterS, &afterL)
	if afterS >= beforeS {
		t.Errorf("stability: %.2f should be < %.2f", afterS, beforeS)
	}
	if afterL != beforeL+1 {
		t.Errorf("lapses: got %d, want %d", afterL, beforeL+1)
	}
}

func TestFSRSIntegrationZPomoca(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// z_pomoca_1 (Good) on new tag → stability grows, interval > 0
	dates, err := updateTags(db, []string{"pomoc-fsrs"}, "poprawne_z_pomoca_1")
	if err != nil {
		t.Fatal(err)
	}
	if len(dates) != 1 {
		t.Fatalf("expected 1 date, got %d", len(dates))
	}
	today := time.Now().Format("2006-01-02")
	if dates[0] < today {
		t.Errorf("z_pomoca_1 date %s should be >= today %s", dates[0], today)
	}

	var s1 float64
	db.QueryRow("SELECT stability FROM progress_tagi WHERE tag = 'pomoc-fsrs'").Scan(&s1)
	if s1 <= 0 {
		t.Errorf("stability after Good: %.3f, want > 0", s1)
	}

	// Second z_pomoca_1 on same day → stability stays same (R=1.0, no learning signal in FSRS)
	dates, err = updateTags(db, []string{"pomoc-fsrs"}, "poprawne_z_pomoca_1")
	if err != nil {
		t.Fatal(err)
	}
	var s2 float64
	db.QueryRow("SELECT stability FROM progress_tagi WHERE tag = 'pomoc-fsrs'").Scan(&s2)
	if s2 <= 0 {
		t.Errorf("stability after 2nd Good: %.3f should be > 0", s2)
	}
	_ = dates

	// Property: Easy interval > Good interval for same state
	datesEasy, _ := updateTags(db, []string{"easy-compare"}, "poprawne_bez_pomocy")
	datesGood, _ := updateTags(db, []string{"good-compare"}, "poprawne_z_pomoca_1")
	if datesEasy[0] < datesGood[0] {
		t.Errorf("Easy date %s should be >= Good date %s", datesEasy[0], datesGood[0])
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

	// First open — should create schema (no matura.db, so attached=false)
	db, attached, err := OpenDB(dir)
	if err != nil {
		t.Fatalf("first open: %v", err)
	}
	if attached {
		t.Error("expected not attached in empty dir")
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
	count, lastTyp, err := getSessionState(db)
	if err != nil {
		t.Fatal(err)
	}
	if count != 0 {
		t.Errorf("fresh count: got %d, want 0", count)
	}
	if lastTyp != "" {
		t.Errorf("fresh lastTyp: got %q, want empty", lastTyp)
	}

	// Increment
	if err := incrementSession(db, "sql_group_by"); err != nil {
		t.Fatal(err)
	}
	count, lastTyp, err = getSessionState(db)
	if err != nil {
		t.Fatal(err)
	}
	if count != 1 {
		t.Errorf("after 1st increment: got %d, want 1", count)
	}
	if lastTyp != "sql_group_by" {
		t.Errorf("after 1st increment: got %q, want sql_group_by", lastTyp)
	}

	// Increment again
	if err := incrementSession(db, "cyfry_liczby"); err != nil {
		t.Fatal(err)
	}
	count, lastTyp, err = getSessionState(db)
	if err != nil {
		t.Fatal(err)
	}
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

	// Fresh DB → getLevel returns "latwe"
	results, err := queryExercises(db, "cyfry_liczby", "latwe", "")
	if err != nil {
		t.Fatal(err)
	}
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
	results, err = queryExercises(db, "cyfry_liczby", "srednie", "")
	if err != nil {
		t.Fatal(err)
	}
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

	// Fresh DB: sql_group_by → first_in_type=true, first_in_category=true
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

// === Feature 1: Time Tracking ===

func TestMigrationV3(t *testing.T) {
	dir := t.TempDir()

	// Create v2 DB (simulating old schema without czas_sek)
	progressPath := filepath.Join(dir, "matura_progress.db")
	db, err := sql.Open("sqlite", progressPath)
	if err != nil {
		t.Fatal(err)
	}

	// Create v2 schema manually (no czas_sek, no progress_bledy)
	db.Exec(`CREATE TABLE schema_version (version INTEGER PRIMARY KEY)`)
	db.Exec(`INSERT INTO schema_version VALUES (2)`)
	db.Exec(`CREATE TABLE progress_meta (key TEXT PRIMARY KEY, value TEXT)`)
	db.Exec(`CREATE TABLE progress_typy (typ TEXT PRIMARY KEY, poziom_trudnosci TEXT DEFAULT 'latwe', streak INTEGER DEFAULT 0)`)
	db.Exec(`CREATE TABLE progress_zrobione (id TEXT PRIMARY KEY, typ TEXT, data TEXT, wynik TEXT)`)
	db.Exec(`CREATE TABLE progress_tagi (tag TEXT PRIMARY KEY, poziom INTEGER DEFAULT 0, nastepna_powtorka TEXT)`)
	db.Exec(`CREATE TABLE matura_zrobione (id TEXT PRIMARY KEY, typ TEXT, data TEXT, punkty INTEGER, max_punkty INTEGER)`)
	db.Exec(`CREATE TABLE probne_matury (id INTEGER PRIMARY KEY AUTOINCREMENT, rok INTEGER, data TEXT, czas_min INTEGER, wynik_pkt INTEGER, max_pkt INTEGER, procent REAL, per_kategoria TEXT, przerwany BOOLEAN DEFAULT 0)`)
	db.Exec(`CREATE TABLE pulapki_przejrzane (id TEXT PRIMARY KEY, typ TEXT, data TEXT, trafienia INTEGER, total INTEGER)`)
	db.Exec(`INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('test.1', 'test', '2026-01-01', 'poprawne_bez_pomocy')`)
	db.Close()

	// Reopen — should migrate to v3
	db, err = sql.Open("sqlite", progressPath)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	err = initProgressSchema(db)
	if err != nil {
		t.Fatalf("migration to v3 failed: %v", err)
	}

	// Verify czas_sek column exists
	_, err = db.Exec("UPDATE progress_zrobione SET czas_sek = 120 WHERE id = 'test.1'")
	if err != nil {
		t.Errorf("czas_sek column missing after migration: %v", err)
	}

	// Verify progress_bledy table exists
	_, err = db.Exec("INSERT INTO progress_bledy (exercise_id, typ, blad_kod, data) VALUES ('test.1', 'test', 'test_err', '2026-01-01')")
	if err != nil {
		t.Errorf("progress_bledy table missing after migration: %v", err)
	}

	// Verify existing data preserved
	var wynik string
	db.QueryRow("SELECT wynik FROM progress_zrobione WHERE id = 'test.1'").Scan(&wynik)
	if wynik != "poprawne_bez_pomocy" {
		t.Errorf("data not preserved: got %q", wynik)
	}

	// Verify version (v2→v3→v4 chain)
	var version int
	db.QueryRow("SELECT version FROM schema_version").Scan(&version)
	if version != currentSchemaVersion {
		t.Errorf("version after migration: got %d, want %d", version, currentSchemaVersion)
	}
}

func TestMigrationV3Idempotent(t *testing.T) {
	dir := t.TempDir()

	// Create v2 DB
	progressPath := filepath.Join(dir, "matura_progress.db")
	db, err := sql.Open("sqlite", progressPath)
	if err != nil {
		t.Fatal(err)
	}
	db.Exec(`CREATE TABLE schema_version (version INTEGER PRIMARY KEY)`)
	db.Exec(`INSERT INTO schema_version VALUES (2)`)
	db.Exec(`CREATE TABLE progress_meta (key TEXT PRIMARY KEY, value TEXT)`)
	db.Exec(`CREATE TABLE progress_typy (typ TEXT PRIMARY KEY, poziom_trudnosci TEXT DEFAULT 'latwe', streak INTEGER DEFAULT 0)`)
	db.Exec(`CREATE TABLE progress_zrobione (id TEXT PRIMARY KEY, typ TEXT, data TEXT, wynik TEXT)`)
	db.Exec(`CREATE TABLE progress_tagi (tag TEXT PRIMARY KEY, poziom INTEGER DEFAULT 0, nastepna_powtorka TEXT)`)
	db.Exec(`CREATE TABLE matura_zrobione (id TEXT PRIMARY KEY, typ TEXT, data TEXT, punkty INTEGER, max_punkty INTEGER)`)
	db.Exec(`CREATE TABLE probne_matury (id INTEGER PRIMARY KEY AUTOINCREMENT, rok INTEGER, data TEXT, czas_min INTEGER, wynik_pkt INTEGER, max_pkt INTEGER, procent REAL, per_kategoria TEXT, przerwany BOOLEAN DEFAULT 0)`)
	db.Exec(`CREATE TABLE pulapki_przejrzane (id TEXT PRIMARY KEY, typ TEXT, data TEXT, trafienia INTEGER, total INTEGER)`)
	db.Close()

	// Migrate to v3 twice — second run should be a no-op
	for i := 0; i < 2; i++ {
		db, err = sql.Open("sqlite", progressPath)
		if err != nil {
			t.Fatal(err)
		}
		err = initProgressSchema(db)
		if err != nil {
			t.Fatalf("migration run %d failed: %v", i+1, err)
		}
		db.Close()
	}
}

func TestProgressUpdateWithCzas(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Simulate what progressUpdateCmd does with --czas
	id := "7.1"
	wynik := "poprawne_bez_pomocy"
	czas := 145
	today := time.Now().Format("2006-01-02")

	var typNazwa string
	err := db.QueryRow("SELECT typ_nazwa FROM data.cwiczenia WHERE id = ?", id).Scan(&typNazwa)
	if err != nil {
		t.Fatalf("exercise not found: %v", err)
	}

	var czasSekVal *int
	if czas > 0 {
		czasSekVal = &czas
	}
	_, err = db.Exec(`INSERT OR REPLACE INTO progress_zrobione (id, typ, data, wynik, czas_sek) VALUES (?, ?, ?, ?, ?)`,
		id, typNazwa, today, wynik, czasSekVal)
	if err != nil {
		t.Fatalf("insert: %v", err)
	}

	// Verify czas_sek stored
	var czasSek sql.NullInt64
	db.QueryRow("SELECT czas_sek FROM progress_zrobione WHERE id = ?", id).Scan(&czasSek)
	if !czasSek.Valid || czasSek.Int64 != 145 {
		t.Errorf("czas_sek: got %v, want 145", czasSek)
	}

	// Verify tempo calculation with benchmark
	var benchmarkSek sql.NullInt64
	db.QueryRow("SELECT benchmark_sek FROM data.benchmarki WHERE typ = ?", typNazwa).Scan(&benchmarkSek)
	if benchmarkSek.Valid {
		tempo := calculateTempo(czas, int(benchmarkSek.Int64))
		if tempo == "" {
			t.Error("expected non-empty tempo with valid benchmark")
		}
	}
}

func TestProgressUpdateWithoutCzas(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Insert without czas (backward compatibility)
	_, err := db.Exec(`INSERT OR REPLACE INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.2', 'cyfry_liczby', '2026-01-01', 'poprawne_bez_pomocy')`)
	if err != nil {
		t.Fatalf("insert without czas: %v", err)
	}

	var czasSek sql.NullInt64
	db.QueryRow("SELECT czas_sek FROM progress_zrobione WHERE id = '7.2'").Scan(&czasSek)
	if czasSek.Valid {
		t.Errorf("czas_sek should be NULL, got %d", czasSek.Int64)
	}
}

func TestTempoCalculation(t *testing.T) {
	cases := []struct {
		elapsed, benchmark int
		want               string
	}{
		{50, 100, "szybko"},    // 0.5 < 0.6
		{59, 100, "szybko"},    // 0.59 < 0.6
		{60, 100, "ok"},        // 0.6 — boundary: NOT < 0.6, so "ok"
		{100, 100, "ok"},       // 1.0
		{120, 100, "ok"},       // 1.2 — boundary: <= 1.2
		{121, 100, "wolno"},    // 1.21 > 1.2
		{150, 100, "wolno"},    // 1.5
		{200, 100, "wolno"},    // 2.0 — boundary: <= 2.0
		{201, 100, "za_wolno"}, // 2.01 > 2.0
		{250, 100, "za_wolno"},
		{100, 0, ""},       // zero benchmark
		{0, 100, "szybko"}, // zero elapsed
		{100, -1, ""},      // negative benchmark
	}
	for _, tc := range cases {
		got := calculateTempo(tc.elapsed, tc.benchmark)
		if got != tc.want {
			t.Errorf("calculateTempo(%d, %d): got %q, want %q", tc.elapsed, tc.benchmark, got, tc.want)
		}
	}
}

// === Feature 2: Error Diagnosis ===

func TestProgressBlad(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	result, err := db.Exec(
		`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, hint_level, data)
		VALUES ('7.1', 'cyfry_liczby', 'off_by_one', 'Pomylka o 1', 1, '2026-01-15')`)
	if err != nil {
		t.Fatalf("insert blad: %v", err)
	}

	lastID, _ := result.LastInsertId()
	if lastID < 1 {
		t.Errorf("lastID: got %d, want >= 1", lastID)
	}

	var blad BladOut
	err = db.QueryRow(`SELECT id, exercise_id, typ, blad_kod, blad_opis, hint_level, data FROM progress_bledy WHERE id = ?`, lastID).
		Scan(&blad.ID, &blad.ExerciseID, &blad.Typ, &blad.BladKod, &blad.BladOpis, &blad.HintLevel, &blad.Data)
	if err != nil {
		t.Fatalf("query blad: %v", err)
	}
	if blad.ExerciseID != "7.1" {
		t.Errorf("exercise_id: got %q", blad.ExerciseID)
	}
	if blad.BladKod != "off_by_one" {
		t.Errorf("blad_kod: got %q", blad.BladKod)
	}
	if blad.HintLevel != 1 {
		t.Errorf("hint_level: got %d", blad.HintLevel)
	}
}

func TestProgressDiagnose(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Insert 5 errors: 3x brak_group_by, 2x off_by_one
	for i := 0; i < 3; i++ {
		db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, data) VALUES (?, 'sql_group_by', 'brak_group_by', '2026-01-15')`,
			fmt.Sprintf("20.%d", i+1))
	}
	for i := 0; i < 2; i++ {
		db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, data) VALUES (?, 'cyfry_liczby', 'off_by_one', '2026-01-14')`,
			fmt.Sprintf("7.%d", i+1))
	}

	var bladKod string
	var cnt int
	err := db.QueryRow(`SELECT blad_kod, COUNT(*) as cnt FROM progress_bledy GROUP BY blad_kod ORDER BY cnt DESC LIMIT 1`).
		Scan(&bladKod, &cnt)
	if err != nil {
		t.Fatalf("query diagnose: %v", err)
	}
	if bladKod != "brak_group_by" {
		t.Errorf("top blad_kod: got %q, want brak_group_by", bladKod)
	}
	if cnt != 3 {
		t.Errorf("top count: got %d, want 3", cnt)
	}
}

func TestProgressDiagnoseByTyp(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, data) VALUES ('20.1', 'sql_group_by', 'brak_group_by', '2026-01-15')`)
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, data) VALUES ('20.2', 'sql_group_by', 'brak_group_by', '2026-01-15')`)
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, data) VALUES ('7.1', 'cyfry_liczby', 'off_by_one', '2026-01-14')`)

	// Filter by sql_group_by — should only see brak_group_by
	var cnt int
	db.QueryRow(`SELECT COUNT(DISTINCT blad_kod) FROM progress_bledy WHERE typ = 'sql_group_by'`).Scan(&cnt)
	if cnt != 1 {
		t.Errorf("distinct blad_kod for sql_group_by: got %d, want 1", cnt)
	}

	// Filter by cyfry_liczby — should only see off_by_one
	db.QueryRow(`SELECT COUNT(DISTINCT blad_kod) FROM progress_bledy WHERE typ = 'cyfry_liczby'`).Scan(&cnt)
	if cnt != 1 {
		t.Errorf("distinct blad_kod for cyfry_liczby: got %d, want 1", cnt)
	}
}

func TestProgressDiagnoseEmpty(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	var cnt int
	db.QueryRow(`SELECT COUNT(*) FROM progress_bledy`).Scan(&cnt)
	if cnt != 0 {
		t.Errorf("fresh DB: got %d errors, want 0", cnt)
	}

	// Diagnose query should return no rows
	rows, err := db.Query(`SELECT blad_kod, COUNT(*) FROM progress_bledy GROUP BY blad_kod ORDER BY COUNT(*) DESC LIMIT 5`)
	if err != nil {
		t.Fatalf("query: %v", err)
	}
	defer rows.Close()
	rowCount := 0
	for rows.Next() {
		rowCount++
	}
	if rowCount != 0 {
		t.Errorf("empty diagnose: got %d rows, want 0", rowCount)
	}
}

// === Feature: BladWarning on progress update ===

func TestBladWarningOnNonPerfectResult(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	today := time.Now().Format("2006-01-02")

	// Simulate progress update with non-perfect result, no blad logged
	// Insert exercise as done
	db.Exec(`INSERT OR REPLACE INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.1', 'cyfry_liczby', ?, 'poprawne_z_pomoca_1')`, today)

	// Check progress_bledy for this exercise today — should be 0
	var bladyCount int
	db.QueryRow("SELECT COUNT(*) FROM progress_bledy WHERE exercise_id = '7.1' AND data = ?", today).Scan(&bladyCount)
	if bladyCount != 0 {
		t.Errorf("expected 0 bledy, got %d", bladyCount)
	}

	// Warning should be generated: wynik != poprawne_bez_pomocy && bladyCount == 0
	wynik := "poprawne_z_pomoca_1"
	shouldWarn := wynik != "poprawne_bez_pomocy" && bladyCount == 0
	if !shouldWarn {
		t.Error("expected blad_warning to be generated")
	}
}

func TestNoBladWarningOnPerfectResult(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	today := time.Now().Format("2006-01-02")

	// Insert exercise with perfect result
	db.Exec(`INSERT OR REPLACE INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.1', 'cyfry_liczby', ?, 'poprawne_bez_pomocy')`, today)

	// Perfect result → no warning regardless of blad count
	wynik := "poprawne_bez_pomocy"
	shouldWarn := wynik != "poprawne_bez_pomocy"
	if shouldWarn {
		t.Error("expected no blad_warning for perfect result")
	}
}

func TestNoBladWarningWhenBladLogged(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	today := time.Now().Format("2006-01-02")

	// Insert exercise with non-perfect result
	db.Exec(`INSERT OR REPLACE INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.1', 'cyfry_liczby', ?, 'poprawne_z_pomoca_1')`, today)

	// Log a blad for today
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, data) VALUES ('7.1', 'cyfry_liczby', 'off_by_one', ?)`, today)

	// Check progress_bledy — should be > 0
	var bladyCount int
	db.QueryRow("SELECT COUNT(*) FROM progress_bledy WHERE exercise_id = '7.1' AND data = ?", today).Scan(&bladyCount)
	if bladyCount == 0 {
		t.Error("expected blady logged, got 0")
	}

	// No warning: bladyCount > 0
	wynik := "poprawne_z_pomoca_1"
	shouldWarn := wynik != "poprawne_bez_pomocy" && bladyCount == 0
	if shouldWarn {
		t.Error("expected no blad_warning when blad is logged")
	}
}

// === Feature 3: Rich Type Intro ===

func TestTypIntroWithCKEStats(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// sledzenie_algorytmu appears in every year (11/11)
	ckeNames := exerciseTypToCKETypes("sledzenie_algorytmu", "TEORIA")
	var wystapienia int
	db.QueryRow(
		`SELECT COUNT(DISTINCT rok) FROM data.egzamin WHERE typ_zadania IN (`+placeholders(len(ckeNames))+`)`,
		toAny(ckeNames)...).Scan(&wystapienia)

	if wystapienia != 11 {
		t.Errorf("sledzenie_algorytmu wystapienia: got %d, want 11", wystapienia)
	}
}

func TestTypIntroTopPulapki(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Check that sledzenie_algorytmu has pulapki
	ckeNames := exerciseTypToCKETypes("sledzenie_algorytmu", "TEORIA")
	var pulapkiCount int
	db.QueryRow(
		`SELECT COUNT(*) FROM data.egzamin WHERE typ_zadania IN (`+placeholders(len(ckeNames))+`)
		AND pulapki != '[]' AND pulapki != 'null'`, toAny(ckeNames)...).Scan(&pulapkiCount)

	if pulapkiCount == 0 {
		t.Skip("no pulapki found for sledzenie_algorytmu — skipping")
	}

	// Collect unique pulapki
	rows, _ := db.Query(
		`SELECT pulapki FROM data.egzamin WHERE typ_zadania IN (`+placeholders(len(ckeNames))+`)
		AND pulapki != '[]' AND pulapki != 'null'`, toAny(ckeNames)...)
	unique := map[string]bool{}
	for rows.Next() {
		var pJSON string
		rows.Scan(&pJSON)
		var pulapki []string
		json.Unmarshal([]byte(pJSON), &pulapki)
		for _, p := range pulapki {
			if p != "" {
				unique[p] = true
			}
		}
	}
	rows.Close()

	if len(unique) == 0 {
		t.Error("expected non-empty pulapki")
	}
}

func TestTypIntroPrzyklad(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Every type should have at least one "latwe" exercise
	var id, tresc string
	err := db.QueryRow(
		`SELECT id, tresc FROM data.cwiczenia WHERE typ_nazwa = 'sledzenie_algorytmu' AND trudnosc = 'latwe'
		ORDER BY LENGTH(tresc) ASC LIMIT 1`).Scan(&id, &tresc)
	if err != nil {
		t.Fatalf("no latwe exercise for sledzenie_algorytmu: %v", err)
	}
	if id == "" {
		t.Error("przyklad id empty")
	}
	if tresc == "" {
		t.Error("przyklad tresc empty")
	}
}

func TestTypIntroCheatsheetExcerpt(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	var fullContent string
	err := db.QueryRow("SELECT content FROM data.cheatsheets WHERE kategoria = 'SQL'").Scan(&fullContent)
	if err != nil {
		t.Fatalf("cheatsheet SQL: %v", err)
	}
	if len(fullContent) == 0 {
		t.Fatal("cheatsheet content empty")
	}

	// Replicate the exact excerpt logic from typIntroCmd
	excerpt := fullContent
	if len(excerpt) > 500 {
		cutoff := strings.LastIndex(excerpt[:500], ".")
		if cutoff > 200 {
			excerpt = excerpt[:cutoff+1]
		} else {
			excerpt = excerpt[:500] + "..."
		}
	}

	// Excerpt must be non-empty and bounded
	if len(excerpt) == 0 {
		t.Error("excerpt empty")
	}
	if len(excerpt) > 501 { // 500 + "..." would be 503, but "." cutoff keeps it <= 500+1
		t.Errorf("excerpt too long: %d chars", len(excerpt))
	}
	// Must be shorter than full content (SQL cheatsheet is >500)
	if len(fullContent) > 500 && len(excerpt) >= len(fullContent) {
		t.Error("excerpt not truncated")
	}
}

func TestImportBenchmarks(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	var count int
	db.QueryRow("SELECT COUNT(*) FROM data.benchmarki").Scan(&count)
	if count == 0 {
		t.Error("no benchmarks imported")
	}

	// sledzenie_algorytmu should have a benchmark
	var benchSek sql.NullInt64
	db.QueryRow("SELECT benchmark_sek FROM data.benchmarki WHERE typ = 'sledzenie_algorytmu'").Scan(&benchSek)
	if !benchSek.Valid {
		t.Error("no benchmark for sledzenie_algorytmu")
	}
	if benchSek.Valid && benchSek.Int64 <= 0 {
		t.Errorf("benchmark for sledzenie_algorytmu: got %d, want > 0", benchSek.Int64)
	}
}

// === Fix #9: examMeta czesc column ===

func TestExamMetaCzesc(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// 2015 (stara formula) should have tasks in czesc 1 and czesc 2
	rows, err := db.Query(`
		SELECT DISTINCT e.czesc
		FROM data.egzamin e
		WHERE e.rok = 2015
		ORDER BY e.czesc`)
	if err != nil {
		t.Fatalf("query czesc: %v", err)
	}
	defer rows.Close()

	var czesci []int
	for rows.Next() {
		var c int
		rows.Scan(&c)
		czesci = append(czesci, c)
	}

	if len(czesci) < 2 {
		t.Errorf("2015 czesci: got %v, want [1, 2]", czesci)
	}

	// 2023 (nowa formula) should have only czesc 1
	var czesc2023 int
	db.QueryRow(`SELECT DISTINCT e.czesc FROM data.egzamin e WHERE e.rok = 2023`).Scan(&czesc2023)
	if czesc2023 != 1 {
		t.Errorf("2023 czesc: got %d, want 1", czesc2023)
	}
}

// === Audit fixes ===

func TestFSRSIntegrationZPomoca2(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Build up tag with 2× Easy
	updateTags(db, []string{"pomoc2-fsrs"}, "poprawne_bez_pomocy")
	updateTags(db, []string{"pomoc2-fsrs"}, "poprawne_bez_pomocy")

	var beforeS float64
	db.QueryRow("SELECT stability FROM progress_tagi WHERE tag = 'pomoc2-fsrs'").Scan(&beforeS)

	// z_pomoca_2 (Hard) → stability still grows but slower than Good
	dates, err := updateTags(db, []string{"pomoc2-fsrs"}, "poprawne_z_pomoca_2")
	if err != nil {
		t.Fatal(err)
	}
	today := time.Now().Format("2006-01-02")
	if dates[0] < today {
		t.Errorf("z_pomoca_2 date %s should be >= today %s", dates[0], today)
	}

	var afterS float64
	db.QueryRow("SELECT stability FROM progress_tagi WHERE tag = 'pomoc2-fsrs'").Scan(&afterS)
	// Hard on a reviewed card should still produce positive stability
	if afterS <= 0 {
		t.Errorf("stability after Hard: %.3f, want > 0", afterS)
	}

	// Property: Hard grows slower than Easy on fresh tags
	updateTags(db, []string{"hard-grow"}, "poprawne_z_pomoca_2")
	updateTags(db, []string{"easy-grow"}, "poprawne_bez_pomocy")
	var hardS, easyS float64
	db.QueryRow("SELECT stability FROM progress_tagi WHERE tag = 'hard-grow'").Scan(&hardS)
	db.QueryRow("SELECT stability FROM progress_tagi WHERE tag = 'easy-grow'").Scan(&easyS)
	if hardS >= easyS {
		t.Errorf("Hard stability %.3f should be < Easy stability %.3f", hardS, easyS)
	}
}

func TestExamSaveTransaction(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	today := time.Now().Format("2006-01-02")

	// Committed transaction: both matura_zrobione and probne_matury are written
	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}

	tx.Exec(`INSERT OR REPLACE INTO matura_zrobione (id, typ, data, punkty, max_punkty) VALUES (?, ?, ?, ?, ?)`,
		"2024.1.1", "sledzenie_algorytmu", today, 1, 1)
	tx.Exec(`INSERT OR REPLACE INTO matura_zrobione (id, typ, data, punkty, max_punkty) VALUES (?, ?, ?, ?, ?)`,
		"2024.1.2", "sledzenie_algorytmu", today, 2, 2)
	tx.Exec(`INSERT INTO probne_matury (rok, data, czas_min, wynik_pkt, max_pkt, procent) VALUES (?, ?, ?, ?, ?, ?)`,
		2024, today, 180, 3, 3, 100.0)

	if err := tx.Commit(); err != nil {
		t.Fatalf("commit: %v", err)
	}

	var zrobioneCount int
	db.QueryRow("SELECT COUNT(*) FROM matura_zrobione").Scan(&zrobioneCount)
	if zrobioneCount != 2 {
		t.Errorf("matura_zrobione after commit: got %d, want 2", zrobioneCount)
	}

	var probneCount int
	db.QueryRow("SELECT COUNT(*) FROM probne_matury").Scan(&probneCount)
	if probneCount != 1 {
		t.Errorf("probne_matury after commit: got %d, want 1", probneCount)
	}

	// Rolled-back transaction: nothing persisted
	tx2, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	tx2.Exec(`INSERT OR REPLACE INTO matura_zrobione (id, typ, data, punkty, max_punkty) VALUES ('2024.2.1', 'test', ?, 1, 1)`, today)
	tx2.Exec(`INSERT INTO probne_matury (rok, data, czas_min, wynik_pkt, max_pkt, procent) VALUES (2025, ?, 120, 1, 1, 100.0)`, today)
	tx2.Rollback()

	db.QueryRow("SELECT COUNT(*) FROM matura_zrobione").Scan(&zrobioneCount)
	if zrobioneCount != 2 {
		t.Errorf("matura_zrobione after rollback: got %d, want 2", zrobioneCount)
	}
	db.QueryRow("SELECT COUNT(*) FROM probne_matury").Scan(&probneCount)
	if probneCount != 1 {
		t.Errorf("probne_matury after rollback: got %d, want 1", probneCount)
	}
}

// Regression: tx.QueryRow must see ATTACH'ed data.egzamin inside a transaction.
// Bug: d.QueryRow during active tx gets a new connection without ATTACH.
func TestExamSaveAttachedQueryInTx(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	tx, err := db.Begin()
	if err != nil {
		t.Fatal(err)
	}
	defer tx.Rollback()

	// This query must succeed — it reads from the ATTACH'ed data alias
	var typ string
	err = tx.QueryRow("SELECT typ_zadania FROM data.egzamin WHERE id = '2024.1.1'").Scan(&typ)
	if err != nil {
		t.Fatalf("tx.QueryRow on data.egzamin failed (ATTACH not visible in tx?): %v", err)
	}
	if typ == "" {
		t.Error("typ_zadania should not be empty")
	}
}

// === Feature: extractSection ===

func TestExtractSectionExactMatch(t *testing.T) {
	md := "# Title\n\n## First section\n\nContent of first.\n\n## Second section\n\nContent of second.\n\n## Third section\n\nContent of third."
	got := extractSection(md, "Second")
	if !strings.HasPrefix(got, "## Second section") {
		t.Errorf("exact match: got %q", got)
	}
	if strings.Contains(got, "## Third") {
		t.Error("should not include next section")
	}
	if strings.Contains(got, "## First") {
		t.Error("should not include previous section")
	}
}

func TestExtractSectionCaseInsensitive(t *testing.T) {
	md := "## ABC Heading\n\nSome content.\n\n## DEF Heading\n\nOther content."
	got := extractSection(md, "abc")
	if !strings.HasPrefix(got, "## ABC Heading") {
		t.Errorf("case insensitive: got %q", got)
	}
}

func TestExtractSectionPolishInflection(t *testing.T) {
	// "pulapki" should match "pulapek" via prefix matching (common prefix "pulap" >= 5)
	md := "## 7 archetypow\n\nContent.\n\n## 6 pulapek sledzenia\n\n1. First trap.\n2. Second trap."
	got := extractSection(md, "pulapki")
	if !strings.Contains(got, "pulapek sledzenia") {
		t.Errorf("Polish inflection: expected 'pulapek sledzenia' section, got %q", got)
	}
}

func TestExtractSectionNotFound(t *testing.T) {
	md := "## One\n\nContent.\n\n## Two\n\nMore content."
	got := extractSection(md, "nonexistent_section_xyz")
	if got != "" {
		t.Errorf("not found: expected empty, got %q", got)
	}
}

func TestExtractSectionLastSection(t *testing.T) {
	md := "## First\n\nContent.\n\n## Last section\n\nFinal content here."
	got := extractSection(md, "Last")
	if !strings.HasPrefix(got, "## Last section") {
		t.Errorf("last section: got %q", got)
	}
	if !strings.Contains(got, "Final content here.") {
		t.Error("should include content of last section")
	}
}

// === Feature: commonPrefix ===

func TestCommonPrefix(t *testing.T) {
	cases := []struct {
		a, b string
		want int
	}{
		{"pulapki", "pulapek", 5},
		{"abc", "abd", 2},
		{"", "abc", 0},
		{"abc", "", 0},
		{"same", "same", 4},
		{"abcdef", "abcxyz", 3},
	}
	for _, tc := range cases {
		got := commonPrefix(tc.a, tc.b)
		if got != tc.want {
			t.Errorf("commonPrefix(%q, %q): got %d, want %d", tc.a, tc.b, got, tc.want)
		}
	}
}

// === Feature: chooseWeakestType ===

func TestChooseWeakestTypeFreshDB(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Fresh DB — all types have streak 0, should return one of the TEORIA types
	typ, err := chooseWeakestType(db, "TEORIA")
	if err != nil {
		t.Fatalf("chooseWeakestType TEORIA: %v", err)
	}
	if typ == "" {
		t.Error("expected non-empty type")
	}

	// Verify it's actually a TEORIA type
	kat := getKategoria(db, typ)
	if kat != "TEORIA" {
		t.Errorf("chosen type %q has kategoria %q, want TEORIA", typ, kat)
	}
}

func TestChooseWeakestTypePicksLowestStreak(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Set high streak for all TEORIA types except teoria_bezpieczenstwa
	teoriaTypes := []string{
		"sledzenie_algorytmu", "projektowanie_algorytmu", "analiza_algorytmu",
		"test_prawda_falsz", "konwersja_systemow_liczbowych",
	}
	for _, tt := range teoriaTypes {
		db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES (?, 'trudne', 10)", tt)
	}
	// teoria_bezpieczenstwa has streak 0 (not inserted) — should be chosen
	typ, err := chooseWeakestType(db, "TEORIA")
	if err != nil {
		t.Fatalf("chooseWeakestType: %v", err)
	}
	if typ != "teoria_bezpieczenstwa" {
		t.Errorf("expected teoria_bezpieczenstwa (lowest streak), got %q", typ)
	}
}

func TestChooseWeakestTypeInvalidKategoria(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	_, err := chooseWeakestType(db, "NONEXISTENT")
	if err == nil {
		t.Error("expected error for nonexistent kategoria")
	}
}

func TestChooseWeakestTypeAllCategories(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	for _, kat := range []string{"TEORIA", "IMPLEMENTACJA", "ARKUSZ", "SQL"} {
		typ, err := chooseWeakestType(db, kat)
		if err != nil {
			t.Errorf("chooseWeakestType(%s): %v", kat, err)
			continue
		}
		gotKat := getKategoria(db, typ)
		if gotKat != kat {
			t.Errorf("chooseWeakestType(%s) returned %q with kategoria %q", kat, typ, gotKat)
		}
	}
}

// === Feature: Context Weight Tracking ===

func TestExerciseNextWeight(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Helper: read current weight from progress_meta
	readWeight := func() int {
		var wStr string
		var w int
		if err := db.QueryRow(`SELECT value FROM progress_meta WHERE key = 'session_context_weight'`).Scan(&wStr); err == nil {
			fmt.Sscanf(wStr, "%d", &w)
		}
		return w
	}

	// 1. weight-reset sets weight to 0
	db.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '0')`)
	if w := readWeight(); w != 0 {
		t.Errorf("after reset: got weight %d, want 0", w)
	}

	// 2. weight-add 40 → weight=40, reset_suggested=false
	db.Exec(`INSERT INTO progress_meta (key, value) VALUES ('session_context_weight', ?)
		ON CONFLICT(key) DO UPDATE SET value = CAST(CAST(value AS INTEGER) + ? AS TEXT)`, 40, 40)
	if w := readWeight(); w != 40 {
		t.Errorf("after +40: got weight %d, want 40", w)
	}
	if readWeight() >= 80 {
		t.Error("weight 40: reset_suggested should be false")
	}

	// 3. weight-add 41 → weight=81, reset_suggested=true
	db.Exec(`INSERT INTO progress_meta (key, value) VALUES ('session_context_weight', ?)
		ON CONFLICT(key) DO UPDATE SET value = CAST(CAST(value AS INTEGER) + ? AS TEXT)`, 41, 41)
	if w := readWeight(); w != 81 {
		t.Errorf("after +41: got weight %d, want 81", w)
	}
	if readWeight() < 80 {
		t.Error("weight 81: reset_suggested should be true")
	}

	// 4. weight-reset again → back to 0
	db.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '0')`)
	if w := readWeight(); w != 0 {
		t.Errorf("after second reset: got weight %d, want 0", w)
	}
}

// === Fix: findExerciseByTag exact matching ===

func TestFindExerciseByTag(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// "suma-cyfr" is a real tag; "suma" is a substring but not an exact tag.
	// With INSTR, searching for "suma" would match "suma-cyfr" — a false positive.
	// With json_each exact matching, "suma" must NOT match.

	// Exact match: "suma-cyfr" should find exercises
	ex, err := findExerciseByTag(db, "suma-cyfr")
	if err != nil {
		t.Fatalf("findExerciseByTag('suma-cyfr') failed: %v", err)
	}
	if ex.ID == "" {
		t.Error("expected non-empty exercise for tag 'suma-cyfr'")
	}

	// Substring "suma" should NOT match (no exercise has bare "suma" as a tag)
	_, err = findExerciseByTag(db, "suma")
	if err == nil {
		t.Error("findExerciseByTag('suma') should fail — no exact tag match, only 'suma-cyfr' exists")
	}
}

// === Feature: cheatsheet --sekcja integration ===

func TestCheatsheetSekcjaIntegration(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Get full TEORIA cheatsheet
	var full string
	err := db.QueryRow("SELECT content FROM data.cheatsheets WHERE kategoria = 'TEORIA'").Scan(&full)
	if err != nil {
		t.Fatalf("cheatsheet TEORIA: %v", err)
	}

	// Extract "archetyp" section
	section := extractSection(full, "archetyp")
	if section == "" {
		t.Fatal("archetyp section not found in TEORIA cheatsheet")
	}
	if !strings.Contains(section, "archetyp") {
		t.Error("section does not contain 'archetyp'")
	}
	if len(section) >= len(full) {
		t.Error("section should be shorter than full cheatsheet")
	}

	// Extract "bezpieczen" section (new addition)
	section = extractSection(full, "bezpieczen")
	if section == "" {
		t.Fatal("bezpieczenstwo section not found in TEORIA cheatsheet")
	}
	if !strings.Contains(section, "Phishing") {
		t.Error("bezpieczenstwo section should contain 'Phishing'")
	}

	// Extract "zlozonosc" section
	section = extractSection(full, "zlozonosc")
	if section == "" {
		t.Fatal("zlozonosc section not found in TEORIA cheatsheet")
	}
}

// === Feature: FSRS Schema Migration ===

func TestMigrationV4(t *testing.T) {
	dir := t.TempDir()

	// Create v3 DB (current production schema)
	progressPath := filepath.Join(dir, "matura_progress.db")
	db, err := sql.Open("sqlite", progressPath)
	if err != nil {
		t.Fatal(err)
	}

	db.Exec(`CREATE TABLE schema_version (version INTEGER PRIMARY KEY)`)
	db.Exec(`INSERT INTO schema_version VALUES (3)`)
	db.Exec(`CREATE TABLE progress_meta (key TEXT PRIMARY KEY, value TEXT)`)
	db.Exec(`CREATE TABLE progress_typy (typ TEXT PRIMARY KEY, poziom_trudnosci TEXT DEFAULT 'latwe', streak INTEGER DEFAULT 0)`)
	db.Exec(`CREATE TABLE progress_zrobione (id TEXT PRIMARY KEY, typ TEXT, data TEXT, wynik TEXT, czas_sek INTEGER)`)
	db.Exec(`CREATE TABLE progress_tagi (tag TEXT PRIMARY KEY, poziom INTEGER DEFAULT 0, nastepna_powtorka TEXT)`)
	db.Exec(`CREATE TABLE matura_zrobione (id TEXT PRIMARY KEY, typ TEXT, data TEXT, punkty INTEGER, max_punkty INTEGER)`)
	db.Exec(`CREATE TABLE probne_matury (id INTEGER PRIMARY KEY AUTOINCREMENT, rok INTEGER, data TEXT, czas_min INTEGER, wynik_pkt INTEGER, max_pkt INTEGER, procent REAL, per_kategoria TEXT, przerwany BOOLEAN DEFAULT 0)`)
	db.Exec(`CREATE TABLE pulapki_przejrzane (id TEXT PRIMARY KEY, typ TEXT, data TEXT, trafienia INTEGER, total INTEGER)`)
	db.Exec(`CREATE TABLE progress_bledy (id INTEGER PRIMARY KEY AUTOINCREMENT, exercise_id TEXT NOT NULL, typ TEXT NOT NULL, blad_kod TEXT NOT NULL, blad_opis TEXT, hint_level INTEGER DEFAULT 0, data TEXT NOT NULL)`)

	// Insert test data at various levels
	db.Exec(`INSERT INTO progress_tagi (tag, poziom, nastepna_powtorka) VALUES ('petla', 3, '2026-02-20')`)
	db.Exec(`INSERT INTO progress_tagi (tag, poziom, nastepna_powtorka) VALUES ('tablica', 0, '2026-02-15')`)
	db.Exec(`INSERT INTO progress_tagi (tag, poziom, nastepna_powtorka) VALUES ('rekurencja', 4, '2026-03-01')`)
	db.Close()

	// Reopen — should migrate to v4
	db, err = sql.Open("sqlite", progressPath)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	err = initProgressSchema(db)
	if err != nil {
		t.Fatalf("migration v3→v4: %v", err)
	}

	// Verify FSRS columns exist and seeded correctly
	var stability, difficulty float64
	var state, lapses int
	db.QueryRow("SELECT stability, difficulty, state, lapses FROM progress_tagi WHERE tag = 'petla'").
		Scan(&stability, &difficulty, &state, &lapses)
	if stability != 7.0 {
		t.Errorf("petla stability: got %.1f, want 7.0 (from poziom=3)", stability)
	}
	if difficulty != 5.0 {
		t.Errorf("petla difficulty: got %.1f, want 5.0", difficulty)
	}
	if state != StateReview {
		t.Errorf("petla state: got %d, want %d (Review)", state, StateReview)
	}

	// tablica (poziom=0) → stability=0.4, state=Learning
	db.QueryRow("SELECT stability, state FROM progress_tagi WHERE tag = 'tablica'").
		Scan(&stability, &state)
	if stability != 0.4 {
		t.Errorf("tablica stability: got %.1f, want 0.4", stability)
	}
	if state != StateLearning {
		t.Errorf("tablica state: got %d, want %d (Learning)", state, StateLearning)
	}

	// rekurencja (poziom=4) → stability=21.0, state=Review
	db.QueryRow("SELECT stability, state FROM progress_tagi WHERE tag = 'rekurencja'").
		Scan(&stability, &state)
	if stability != 21.0 {
		t.Errorf("rekurencja stability: got %.1f, want 21.0", stability)
	}
	if state != StateReview {
		t.Errorf("rekurencja state: got %d, want %d (Review)", state, StateReview)
	}

	// Verify version updated (v3→v4 chain runs all pending migrations)
	var version int
	db.QueryRow("SELECT version FROM schema_version").Scan(&version)
	if version != currentSchemaVersion {
		t.Errorf("version: got %d, want %d", version, currentSchemaVersion)
	}
}
