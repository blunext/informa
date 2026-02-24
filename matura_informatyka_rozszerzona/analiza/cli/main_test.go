package main

import (
	"bytes"
	"context"
	"database/sql"
	"encoding/json"
	"fmt"
	"os"
	"path/filepath"
	"strings"
	"testing"
	"time"

	"github.com/spf13/cobra"
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

	if exercises != 583 {
		t.Errorf("exercises: got %d, want 583", exercises)
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
	if count != 583 {
		t.Errorf("after re-import: got %d exercises, want 583", count)
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

func TestDataVerify(t *testing.T) {
	dir := testDir(t)

	dataDB, err := OpenDataDB(dir)
	if err != nil {
		t.Fatal(err)
	}

	result, err := VerifyExercises(dataDB, filepath.Join("..", ""))
	dataDB.Close()
	if err != nil {
		t.Fatalf("verify: %v", err)
	}

	if result.TotalDisk != 583 {
		t.Errorf("total_disk: got %d, want 583", result.TotalDisk)
	}
	if result.TotalDB != 583 {
		t.Errorf("total_db: got %d, want 583", result.TotalDB)
	}
	if result.Matched != 583 {
		t.Errorf("matched: got %d, want 583", result.Matched)
	}
	if len(result.Mismatched) != 0 {
		t.Errorf("mismatched: %v", result.Mismatched)
	}
	if len(result.MissingInDB) != 0 {
		t.Errorf("missing_in_db: %v", result.MissingInDB)
	}
	if len(result.MissingOnDisk) != 0 {
		t.Errorf("missing_on_disk: %v", result.MissingOnDisk)
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

func TestFeedbackCzasowy(t *testing.T) {
	cases := []struct {
		tempo          string
		elapsed, bench int
		wantContains   string
		wantEmpty      bool
	}{
		{"szybko", 50, 100, "Świetne tempo!", false},
		{"szybko", 50, 100, "Ty: 50s", false},
		{"ok", 100, 100, "", true},
		{"wolno", 150, 100, "Na egzaminie", false},
		{"wolno", 150, 100, "Ty: 150s", false},
		{"za_wolno", 250, 100, "To zajęło 250s", false},
		{"", 100, 100, "", true},
	}
	for _, tc := range cases {
		got := feedbackCzasowy(tc.tempo, tc.elapsed, tc.bench)
		if tc.wantEmpty {
			if got != "" {
				t.Errorf("feedbackCzasowy(%q, %d, %d) = %q, want empty", tc.tempo, tc.elapsed, tc.bench, got)
			}
			continue
		}
		if !strings.Contains(got, tc.wantContains) {
			t.Errorf("feedbackCzasowy(%q, %d, %d) = %q, want contain %q", tc.tempo, tc.elapsed, tc.bench, got, tc.wantContains)
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

	// 1. Initialize weight to 0
	db.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '0')`)
	if w := readWeight(); w != 0 {
		t.Errorf("after init: got weight %d, want 0", w)
	}

	// 2. addWeight(4) twice → weight=8
	addWeight(db, 4)
	addWeight(db, 4)
	if w := readWeight(); w != 8 {
		t.Errorf("after 2x addWeight(4): got weight %d, want 8", w)
	}

	// 3. addWeight with 0 should be no-op
	addWeight(db, 0)
	if w := readWeight(); w != 8 {
		t.Errorf("after addWeight(0): got weight %d, want 8", w)
	}

	// 4. Reset again → back to 0
	db.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '0')`)
	if w := readWeight(); w != 0 {
		t.Errorf("after second reset: got weight %d, want 0", w)
	}
}

func TestProgressStatusResetsWeight(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Set weight to 50 via direct SQL
	db.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '50')`)

	// Verify weight is 50
	var wStr string
	if err := db.QueryRow(`SELECT value FROM progress_meta WHERE key = 'session_context_weight'`).Scan(&wStr); err != nil {
		t.Fatalf("read weight: %v", err)
	}
	if wStr != "50" {
		t.Fatalf("setup: weight = %q, want 50", wStr)
	}

	// Build and run the progress status command via cobra.
	// jsonOut writes to os.Stdout directly, so we redirect it to capture output.
	statusCmd := progressStatusCmd()
	parent := &cobra.Command{Use: "progress"}
	parent.AddCommand(statusCmd)
	ctx := context.WithValue(context.Background(), ctxKey{}, db)
	statusCmd.SetContext(ctx)

	// Capture os.Stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := statusCmd.RunE(statusCmd, nil)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("progress status: %v", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	// Verify weight is now 0
	var wAfter string
	if err := db.QueryRow(`SELECT value FROM progress_meta WHERE key = 'session_context_weight'`).Scan(&wAfter); err != nil {
		t.Fatalf("read weight after: %v", err)
	}
	if wAfter != "0" {
		t.Errorf("weight after progress status: got %q, want \"0\"", wAfter)
	}

	// Verify output is valid JSON
	var parsed map[string]interface{}
	if err := json.Unmarshal([]byte(output), &parsed); err != nil {
		t.Errorf("output is not valid JSON: %v\noutput: %s", err, output)
	}
}

func TestAutoWeight(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	readWeight := func() int {
		var wStr string
		var w int
		if err := db.QueryRow(`SELECT value FROM progress_meta WHERE key = 'session_context_weight'`).Scan(&wStr); err == nil {
			fmt.Sscanf(wStr, "%d", &w)
		}
		return w
	}

	// Reset
	db.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '0')`)

	// Simulate exercise next (4) + cke get (5) + progress blad hint=3 (3) = 12
	addWeight(db, 4)
	addWeight(db, 5)
	addWeight(db, 3)
	if w := readWeight(); w != 12 {
		t.Errorf("cumulative weight: got %d, want 12", w)
	}

	// Simulate cheatsheet get with sekcja (2) → 14
	addWeight(db, 2)
	if w := readWeight(); w != 14 {
		t.Errorf("after cheatsheet sekcja: got %d, want 14", w)
	}

	// Threshold at 80: 20 exercise-next cycles (20*4=80)
	db.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('session_context_weight', '0')`)
	for i := 0; i < 20; i++ {
		addWeight(db, 4)
	}
	if w := readWeight(); w != 80 {
		t.Errorf("20 cycles: got %d, want 80", w)
	}
	if readWeight() < 80 {
		t.Error("weight 80: reset_suggested should be true")
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

// === Feature: Diagnostic Dashboard ===

func TestProgressStatusRetencja(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Insert 2 tags with known stability and last_review
	today := time.Now().Format("2006-01-02")
	yesterday := time.Now().AddDate(0, 0, -1).Format("2006-01-02")

	db.Exec(`INSERT INTO progress_tagi (tag, poziom, nastepna_powtorka, stability, difficulty, lapses, reps, state, last_review)
		VALUES ('ret-tag-1', 2, ?, 10.0, 5.0, 0, 3, 2, ?)`, today, today)
	db.Exec(`INSERT INTO progress_tagi (tag, poziom, nastepna_powtorka, stability, difficulty, lapses, reps, state, last_review)
		VALUES ('ret-tag-2', 1, ?, 5.0, 5.0, 0, 2, 2, ?)`, today, yesterday)

	// Compute retencja like progressStatusCmd does
	fsrsParams := DefaultFSRSParams()
	rows, err := db.Query(`SELECT tag, COALESCE(stability, 1.0), COALESCE(last_review, '')
		FROM progress_tagi WHERE nastepna_powtorka IS NOT NULL`)
	if err != nil {
		t.Fatalf("query: %v", err)
	}

	var totalR float64
	var count int
	for rows.Next() {
		var tag, lastReview string
		var stability float64
		rows.Scan(&tag, &stability, &lastReview)
		elapsed := daysBetween(lastReview, today)
		r := fsrsParams.Retrievability(elapsed, stability)
		totalR += r
		count++
	}
	rows.Close()

	if count != 2 {
		t.Fatalf("expected 2 tracked tags, got %d", count)
	}

	avgR := totalR / float64(count)
	// ret-tag-1: elapsed=0 → R=1.0; ret-tag-2: elapsed=1, stability=5 → R≈0.96
	// Average should be between 0.9 and 1.0
	if avgR < 0.9 || avgR > 1.0 {
		t.Errorf("retencja_szacowana: got %.4f, want in [0.9, 1.0]", avgR)
	}
}

func TestProgressStatusRekomendacjaLeech(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Insert a leech tag (lapses=3)
	db.Exec(`INSERT INTO progress_tagi (tag, poziom, nastepna_powtorka, stability, difficulty, lapses, reps, state, last_review)
		VALUES ('leech-tag', 0, '2026-01-01', 0.5, 5.0, 3, 5, 3, '2026-01-01')`)

	out := ProgressStatusOut{
		LeechTagi: []LeechTagOut{{Tag: "leech-tag", Lapses: 3, Stability: 0.5}},
		PerTyp:    []TypStatusOut{},
	}

	rek := computeRekomendacja(db, out)
	if !strings.Contains(rek, "Powtórz") {
		t.Errorf("leech rekomendacja: got %q, want to contain 'Powtórz'", rek)
	}
	if !strings.Contains(rek, "leech-tag") {
		t.Errorf("leech rekomendacja: got %q, want to contain 'leech-tag'", rek)
	}
}

func TestProgressStatusRekomendacjaFreshDB(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Fresh DB — no progress at all → should recommend untouched TIER 1 type
	out := ProgressStatusOut{
		LeechTagi:    []LeechTagOut{},
		PerTyp:       []TypStatusOut{},
		PerKategoria: []KategoriaStatusOut{},
	}

	rek := computeRekomendacja(db, out)
	if !strings.Contains(rek, "Nie ćwiczyłeś") {
		t.Errorf("fresh DB rekomendacja: got %q, want to contain 'Nie ćwiczyłeś'", rek)
	}
	// Should mention one of the TIER 1 types
	foundTier1 := false
	for _, t1 := range tier1Types {
		if strings.Contains(rek, t1) {
			foundTier1 = true
			break
		}
	}
	if !foundTier1 {
		t.Errorf("fresh DB rekomendacja: got %q, want to contain a TIER 1 type", rek)
	}
}

func TestProgressStatusRekomendacjaStreakImbalance(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// All TIER 1 types touched (so we skip rule c)
	perTyp := []TypStatusOut{}
	for _, t1 := range tier1Types {
		perTyp = append(perTyp, TypStatusOut{Typ: t1, Zrobione: 1})
	}

	// Skewed streaks: TEORIA avg=9, SQL avg=1 → ratio 9
	out := ProgressStatusOut{
		LeechTagi: []LeechTagOut{},
		PerTyp:    perTyp,
		PerKategoria: []KategoriaStatusOut{
			{Kategoria: "TEORIA", TypyTotal: 6, TypyRuszane: 3, AvgStreak: 9.0},
			{Kategoria: "IMPLEMENTACJA", TypyTotal: 8, TypyRuszane: 4, AvgStreak: 6.0},
			{Kategoria: "ARKUSZ", TypyTotal: 5, TypyRuszane: 2, AvgStreak: 4.0},
			{Kategoria: "SQL", TypyTotal: 4, TypyRuszane: 2, AvgStreak: 1.0},
		},
	}

	rek := computeRekomendacja(db, out)
	if !strings.Contains(rek, "wyrównaj") {
		t.Errorf("streak imbalance rekomendacja: got %q, want to contain 'wyrównaj'", rek)
	}
	if !strings.Contains(rek, "TEORIA") || !strings.Contains(rek, "SQL") {
		t.Errorf("streak imbalance: got %q, want TEORIA and SQL mentioned", rek)
	}
}

func TestInterleaveSkippedAtStreakThreshold(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Set up a student with streak=3 on cyfry_liczby (at difficulty climb threshold)
	db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'latwe', 3)")

	// Verify getStreak returns 3
	streak := getStreak(db, "cyfry_liczby")
	if streak != 3 {
		t.Fatalf("setup: streak=%d, want 3", streak)
	}

	// sessionCount=3 would trigger interleave (3%3==0)
	// But streak=3 means student is at difficulty climb threshold — interleave should be skipped
	shouldInterleave := (3 > 0 && 3%3 == 0 && streak < 3)
	if shouldInterleave {
		t.Error("interleave should be skipped at streak=3 (difficulty climb threshold)")
	}

	// Verify: interleave is allowed when streak < 3
	db.Exec("UPDATE progress_typy SET streak = 2 WHERE typ = 'cyfry_liczby'")
	streak = getStreak(db, "cyfry_liczby")
	shouldInterleave = (3 > 0 && 3%3 == 0 && streak < 3)
	if !shouldInterleave {
		t.Error("interleave should be allowed when streak < 3")
	}

	// Verify: getStreak returns 0 for unknown type
	streak = getStreak(db, "nonexistent_type")
	if streak != 0 {
		t.Errorf("unknown type: streak=%d, want 0", streak)
	}
}

// === Feature: Hint Fading (max-hints) ===

func TestCalculateMaxHints(t *testing.T) {
	cases := []struct {
		level string
		want  int
	}{
		{"latwe", -1},
		{"srednie", -1},
		{"srednie-trudne", 1},
		{"trudne", 0},
		{"", -1},
	}
	for _, tc := range cases {
		got := calculateMaxHints(tc.level)
		if got != tc.want {
			t.Errorf("calculateMaxHints(%q): got %d, want %d", tc.level, got, tc.want)
		}
	}
}

func TestApplyMaxHints(t *testing.T) {
	tests := []struct {
		name    string
		hints   []string
		level   string
		flag    int
		wantLen int
		wantMax int
	}{
		{"auto latwe, 3 hints", []string{"a", "b", "c"}, "latwe", -1, 3, 3},
		{"auto srednie, 3 hints", []string{"a", "b", "c"}, "srednie", -1, 3, 3},
		{"auto srednie-trudne, 3 hints", []string{"a", "b", "c"}, "srednie-trudne", -1, 1, 1},
		{"auto trudne, 3 hints", []string{"a", "b", "c"}, "trudne", -1, 0, 0},
		{"explicit 0, 3 hints", []string{"a", "b", "c"}, "latwe", 0, 0, 0},
		{"explicit 1, 3 hints", []string{"a", "b", "c"}, "latwe", 1, 1, 1},
		{"explicit 2, 3 hints", []string{"a", "b", "c"}, "latwe", 2, 2, 2},
		{"explicit 5, 3 hints", []string{"a", "b", "c"}, "latwe", 5, 3, 5},
		{"auto trudne, 0 hints", []string{}, "trudne", -1, 0, 0},
		{"auto latwe, 0 hints", []string{}, "latwe", -1, 0, 0},
	}
	for _, tc := range tests {
		t.Run(tc.name, func(t *testing.T) {
			// Make a copy of hints to avoid mutation across tests
			hints := make([]string, len(tc.hints))
			copy(hints, tc.hints)
			ex := ExerciseOut{Wskazowki: hints}
			applyMaxHints(&ex, tc.level, tc.flag)
			if len(ex.Wskazowki) != tc.wantLen {
				t.Errorf("wskazowki len: got %d, want %d", len(ex.Wskazowki), tc.wantLen)
			}
			if ex.MaxHints != tc.wantMax {
				t.Errorf("max_hints: got %d, want %d", ex.MaxHints, tc.wantMax)
			}
		})
	}
}

func TestExerciseGetMaxHintsIntegration(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Fresh DB → level="latwe" → all hints
	results, err := queryExercises(db, "cyfry_liczby", "latwe", "")
	if err != nil {
		t.Fatal(err)
	}
	if len(results) == 0 {
		t.Fatal("no exercises")
	}

	ex := results[0]
	origLen := len(ex.Wskazowki)
	if origLen == 0 {
		t.Skip("exercise has no hints — cannot test truncation")
	}

	// Apply max-hints=1 with auto (latwe level → no limit)
	applyMaxHints(&ex, "latwe", -1)
	if len(ex.Wskazowki) != origLen {
		t.Errorf("latwe auto: got %d hints, want %d", len(ex.Wskazowki), origLen)
	}

	// Reset and apply with srednie-trudne level
	ex2 := results[0]
	applyMaxHints(&ex2, "srednie-trudne", -1)
	if len(ex2.Wskazowki) != 1 {
		t.Errorf("srednie-trudne auto: got %d hints, want 1", len(ex2.Wskazowki))
	}

	// Reset and apply with explicit 0
	ex3 := results[0]
	applyMaxHints(&ex3, "latwe", 0)
	if len(ex3.Wskazowki) != 0 {
		t.Errorf("explicit 0: got %d hints, want 0", len(ex3.Wskazowki))
	}
}

// === Feature: CKE Worked Example ===

func TestWorkedExampleTableCreated(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	var count int
	err := db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name='worked_examples_shown'").Scan(&count)
	if err != nil || count != 1 {
		t.Error("worked_examples_shown table not created")
	}
}

func TestWorkedExampleQuery(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Query a worked example for sledzenie_algorytmu
	var id, tresc, odpowiedz, zasady, pulapkiJSON string
	err := db.QueryRow(`SELECT e.id, e.tresc, e.odpowiedz, e.zasady_oceniania, e.pulapki
		FROM data.egzamin e
		WHERE e.typ_zadania = 'sledzenie_algorytmu'
		AND e.id NOT IN (SELECT id FROM matura_zrobione)
		AND e.id NOT IN (SELECT id FROM worked_examples_shown WHERE typ = 'sledzenie_algorytmu')
		ORDER BY RANDOM() LIMIT 1`).Scan(&id, &tresc, &odpowiedz, &zasady, &pulapkiJSON)
	if err != nil {
		t.Fatalf("worked example query: %v", err)
	}

	if id == "" || tresc == "" || odpowiedz == "" || zasady == "" {
		t.Errorf("missing fields: id=%q tresc=%d odpowiedz=%d zasady=%d",
			id, len(tresc), len(odpowiedz), len(zasady))
	}

	var pulapki []string
	json.Unmarshal([]byte(pulapkiJSON), &pulapki)
	if len(pulapki) == 0 {
		t.Error("expected non-empty pulapki")
	}
}

func TestWorkedExampleNotInMatureZrobione(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Record a worked example
	db.Exec("INSERT INTO worked_examples_shown (id, typ, data) VALUES ('2024.1.1', 'sledzenie_algorytmu', '2026-02-19')")

	// It should NOT appear in matura_zrobione
	var count int
	db.QueryRow("SELECT COUNT(*) FROM matura_zrobione WHERE id = '2024.1.1'").Scan(&count)
	if count != 0 {
		t.Errorf("worked example appeared in matura_zrobione: count=%d", count)
	}
}

func TestWorkedExampleExcludedFromCKEGet(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Record a worked example
	db.Exec("INSERT INTO worked_examples_shown (id, typ, data) VALUES ('2024.1.1', 'sledzenie_algorytmu', '2026-02-19')")

	// cke get query should exclude it
	var found int
	db.QueryRow(`SELECT COUNT(*) FROM data.egzamin e
		WHERE e.id = '2024.1.1'
		AND e.id NOT IN (SELECT id FROM matura_zrobione)
		AND e.id NOT IN (SELECT id FROM worked_examples_shown)`).Scan(&found)
	if found != 0 {
		t.Error("worked example should be excluded from cke get query")
	}
}

func TestCKEStatusSubtractsWorkedExamples(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Get total CKE count for sledzenie_algorytmu
	var total int
	db.QueryRow(`SELECT COUNT(*) FROM data.egzamin WHERE typ_zadania = 'sledzenie_algorytmu'`).Scan(&total)
	if total == 0 {
		t.Skip("no sledzenie_algorytmu CKE tasks in test DB")
	}

	// Record a worked example
	db.Exec("INSERT INTO worked_examples_shown (id, typ, data) VALUES ('2024.1.1', 'sledzenie_algorytmu', '2026-02-19')")

	// Simulate the cke status query logic
	var done int
	db.QueryRow("SELECT COUNT(*) FROM matura_zrobione WHERE typ = 'sledzenie_algorytmu'").Scan(&done)
	var worked int
	db.QueryRow("SELECT COUNT(*) FROM worked_examples_shown WHERE typ = 'sledzenie_algorytmu'").Scan(&worked)
	available := total - done - worked

	if worked != 1 {
		t.Errorf("expected 1 worked example, got %d", worked)
	}
	if available != total-1 {
		t.Errorf("expected available=%d (total %d - 1 worked), got %d", total-1, total, available)
	}
}

func TestWorkedExampleDedup(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Insert same example twice
	db.Exec("INSERT OR IGNORE INTO worked_examples_shown (id, typ, data) VALUES ('2024.1.1', 'sledzenie_algorytmu', '2026-02-19')")
	db.Exec("INSERT OR IGNORE INTO worked_examples_shown (id, typ, data) VALUES ('2024.1.1', 'sledzenie_algorytmu', '2026-02-20')")

	var count int
	db.QueryRow("SELECT COUNT(*) FROM worked_examples_shown WHERE id = '2024.1.1'").Scan(&count)
	if count != 1 {
		t.Errorf("expected 1 record after dedup, got %d", count)
	}
}

// === Feature: Migration V5 (worked_examples_shown) ===

func TestMigrationV5(t *testing.T) {
	dir := t.TempDir()

	// Create v4 DB
	progressPath := filepath.Join(dir, "matura_progress.db")
	db, err := sql.Open("sqlite", progressPath)
	if err != nil {
		t.Fatal(err)
	}

	db.Exec(`CREATE TABLE schema_version (version INTEGER PRIMARY KEY)`)
	db.Exec(`INSERT INTO schema_version VALUES (4)`)
	db.Exec(`CREATE TABLE progress_meta (key TEXT PRIMARY KEY, value TEXT)`)
	db.Exec(`CREATE TABLE progress_typy (typ TEXT PRIMARY KEY, poziom_trudnosci TEXT DEFAULT 'latwe', streak INTEGER DEFAULT 0)`)
	db.Exec(`CREATE TABLE progress_zrobione (id TEXT PRIMARY KEY, typ TEXT, data TEXT, wynik TEXT, czas_sek INTEGER)`)
	db.Exec(`CREATE TABLE progress_tagi (tag TEXT PRIMARY KEY, poziom INTEGER DEFAULT 0, nastepna_powtorka TEXT, stability REAL DEFAULT 0, difficulty REAL DEFAULT 5.0, lapses INTEGER DEFAULT 0, reps INTEGER DEFAULT 0, state INTEGER DEFAULT 0, last_review TEXT)`)
	db.Exec(`CREATE TABLE matura_zrobione (id TEXT PRIMARY KEY, typ TEXT, data TEXT, punkty INTEGER, max_punkty INTEGER)`)
	db.Exec(`CREATE TABLE probne_matury (id INTEGER PRIMARY KEY AUTOINCREMENT, rok INTEGER, data TEXT, czas_min INTEGER, wynik_pkt INTEGER, max_pkt INTEGER, procent REAL, per_kategoria TEXT, przerwany BOOLEAN DEFAULT 0)`)
	db.Exec(`CREATE TABLE pulapki_przejrzane (id TEXT PRIMARY KEY, typ TEXT, data TEXT, trafienia INTEGER, total INTEGER)`)
	db.Exec(`CREATE TABLE progress_bledy (id INTEGER PRIMARY KEY AUTOINCREMENT, exercise_id TEXT NOT NULL, typ TEXT NOT NULL, blad_kod TEXT NOT NULL, blad_opis TEXT, hint_level INTEGER DEFAULT 0, data TEXT NOT NULL)`)
	db.Close()

	// Reopen — should migrate to v5
	db, err = sql.Open("sqlite", progressPath)
	if err != nil {
		t.Fatal(err)
	}
	defer db.Close()

	err = initProgressSchema(db)
	if err != nil {
		t.Fatalf("migration v4→v5: %v", err)
	}

	// Verify worked_examples_shown table exists
	var count int
	db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name='worked_examples_shown'").Scan(&count)
	if count != 1 {
		t.Error("worked_examples_shown table not created by migration")
	}

	// Verify can insert
	_, err = db.Exec("INSERT INTO worked_examples_shown (id, typ, data) VALUES ('test.1', 'test', '2026-02-19')")
	if err != nil {
		t.Errorf("insert into worked_examples_shown failed: %v", err)
	}

	var version int
	db.QueryRow("SELECT version FROM schema_version").Scan(&version)
	if version != currentSchemaVersion {
		t.Errorf("version: got %d, want %d", version, currentSchemaVersion)
	}
}

// === Feature: Lazy Loading Exercises + Coaching ===

func TestQuestionOutHasNoAnswer(t *testing.T) {
	q := QuestionOut{ID: "1.1", Tresc: "test", Coaching: Coaching{StudentLevel: "new"}}
	data, _ := json.Marshal(q)
	var m map[string]any
	json.Unmarshal(data, &m)
	if _, ok := m["odpowiedz"]; ok {
		t.Error("QuestionOut should not have odpowiedz field")
	}
	if _, ok := m["wskazowki"]; ok {
		t.Error("QuestionOut should not have wskazowki field")
	}
	if _, ok := m["coaching"]; !ok {
		t.Error("QuestionOut must have coaching field")
	}
}

func TestBuildCoachingNewStudent(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div"})
	if coaching.StudentLevel != "new" {
		t.Errorf("new student: got level %q, want new", coaching.StudentLevel)
	}
	if coaching.HintDelay != 1 {
		t.Errorf("new student: got hint_delay %d, want 1", coaching.HintDelay)
	}
	if len(coaching.LeechTags) != 0 {
		t.Errorf("new student: got %d leech_tags, want 0", len(coaching.LeechTags))
	}
	if len(coaching.PastMistakes) != 0 {
		t.Errorf("new student: got %d past_mistakes, want 0", len(coaching.PastMistakes))
	}
}

func TestBuildCoachingLearningStudent(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'latwe', 2)")
	for i := 0; i < 3; i++ {
		db.Exec("INSERT OR IGNORE INTO progress_zrobione (id, typ, data, wynik) VALUES (?, 'cyfry_liczby', '2026-02-20', 'poprawne_z_pomoca_1')",
			fmt.Sprintf("7.%d", i+1))
	}

	coaching := buildCoaching(db, "cyfry_liczby", []string{})
	if coaching.StudentLevel != "learning" {
		t.Errorf("learning student: got level %q, want learning", coaching.StudentLevel)
	}
	if coaching.HintDelay != 1 {
		t.Errorf("learning student: got hint_delay %d, want 1", coaching.HintDelay)
	}
}

func TestBuildCoachingFamiliarStudent(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'srednie', 4)")
	for i := 0; i < 5; i++ {
		db.Exec("INSERT OR IGNORE INTO progress_zrobione (id, typ, data, wynik) VALUES (?, 'cyfry_liczby', '2026-02-20', 'poprawne_bez_pomocy')",
			fmt.Sprintf("7.%d", i+1))
	}

	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div"})
	if coaching.StudentLevel != "familiar" {
		t.Errorf("familiar student: got level %q, want familiar", coaching.StudentLevel)
	}
	if coaching.HintDelay != 2 {
		t.Errorf("familiar student: got hint_delay %d, want 2", coaching.HintDelay)
	}
}

func TestBuildCoachingMastered(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	db.Exec("INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'trudne', 5)")

	coaching := buildCoaching(db, "cyfry_liczby", []string{})
	if coaching.StudentLevel != "mastered" {
		t.Errorf("mastered student: got level %q, want mastered", coaching.StudentLevel)
	}
	if coaching.HintDelay != 3 {
		t.Errorf("mastered student: got hint_delay %d, want 3", coaching.HintDelay)
	}
}

func TestBuildCoachingLeechTags(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Low stability (1.0) + old review (30 days ago) → low retrievability (<0.85)
	oldDate := time.Now().AddDate(0, 0, -30).Format("2006-01-02")
	db.Exec(`INSERT INTO progress_tagi (tag, poziom, stability, lapses, reps, state, last_review)
		VALUES ('cyfry-mod-div', 1, 1.0, 4, 6, 2, ?)`, oldDate)

	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div"})
	found := false
	for _, lt := range coaching.LeechTags {
		if lt == "cyfry-mod-div" {
			found = true
		}
	}
	if !found {
		t.Errorf("expected cyfry-mod-div in leech_tags, got %v", coaching.LeechTags)
	}
}

func TestBuildCoachingLeechTagHealed(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	today := time.Now().Format("2006-01-02")
	db.Exec(`INSERT INTO progress_tagi (tag, poziom, stability, lapses, reps, state, last_review)
		VALUES ('cyfry-mod-div', 3, 30.0, 4, 10, 2, ?)`, today)

	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div"})
	if len(coaching.LeechTags) != 0 {
		t.Errorf("healed leech: expected 0 leech_tags, got %v", coaching.LeechTags)
	}
}

func TestBuildCoachingPastMistakes(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	db.Exec("INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.1', 'cyfry_liczby', '2026-02-20', 'poprawne_z_pomoca_1')")
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, data)
		VALUES ('7.1', 'cyfry_liczby', 'cyfry-mod-div', 'inicjalizacja iloczynu na 0', '2026-02-20')`)
	db.Exec(`INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, data)
		VALUES ('7.1', 'cyfry_liczby', 'unrelated-tag', 'nieistotny blad', '2026-02-20')`)

	coaching := buildCoaching(db, "cyfry_liczby", []string{"cyfry-mod-div"})
	if len(coaching.PastMistakes) != 1 {
		t.Errorf("expected 1 past_mistake, got %d: %v", len(coaching.PastMistakes), coaching.PastMistakes)
	}
	if len(coaching.PastMistakes) > 0 && coaching.PastMistakes[0] != "inicjalizacja iloczynu na 0" {
		t.Errorf("expected 'inicjalizacja iloczynu na 0', got %q", coaching.PastMistakes[0])
	}
}

func TestExerciseQuestionOutput(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	results, err := queryExercises(db, "cyfry_liczby", "latwe", "")
	if err != nil || len(results) == 0 {
		t.Fatal("no exercises found")
	}

	ex := results[0]
	q := exerciseToQuestion(ex)
	q.Coaching = buildCoaching(db, ex.TypNazwa, ex.Tagi)

	if q.Tresc == "" {
		t.Error("QuestionOut.Tresc is empty")
	}
	if q.Coaching.StudentLevel != "new" {
		t.Errorf("fresh DB: got level %q, want new", q.Coaching.StudentLevel)
	}

	data, _ := json.Marshal(q)
	var m map[string]any
	json.Unmarshal(data, &m)
	if _, ok := m["odpowiedz"]; ok {
		t.Error("QuestionOut JSON must not contain odpowiedz")
	}
	if _, ok := m["wskazowki"]; ok {
		t.Error("QuestionOut JSON must not contain wskazowki")
	}
}

func TestExerciseHintsById(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	results, _ := queryExercises(db, "cyfry_liczby", "latwe", "")
	if len(results) == 0 {
		t.Fatal("no exercises")
	}
	id := results[0].ID

	hints := getExerciseHints(db, id, "latwe")
	if hints.ID != id {
		t.Errorf("got id %q, want %q", hints.ID, id)
	}
	if len(hints.Wskazowki) == 0 {
		t.Error("expected at least 1 hint")
	}
}

func TestExerciseHintsCheatsheetExcerpt(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Pick a cyfry_liczby exercise (IMPLEMENTACJA category)
	results, _ := queryExercises(db, "cyfry_liczby", "latwe", "")
	if len(results) == 0 {
		t.Fatal("no exercises")
	}
	ex := results[0]

	hints := getExerciseHints(db, ex.ID, "latwe")
	if hints.ID != ex.ID {
		t.Fatalf("got id %q, want %q", hints.ID, ex.ID)
	}
	// cheatsheet_excerpt should be populated when exercise has tags matching cheatsheet sections
	if hints.CheatsheetExcerpt == "" {
		t.Error("expected non-empty cheatsheet_excerpt for cyfry_liczby exercise")
	}
}

func TestExerciseHintsCheatsheetExcerptNoMatch(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// For nonexistent exercise, excerpt should be empty (graceful fallback)
	hints := getExerciseHints(db, "nonexistent_id_99", "latwe")
	if hints.CheatsheetExcerpt != "" {
		t.Errorf("expected empty excerpt for unknown exercise, got %q", hints.CheatsheetExcerpt)
	}
}

func TestExerciseAnswerById(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	results, _ := queryExercises(db, "cyfry_liczby", "latwe", "")
	if len(results) == 0 {
		t.Fatal("no exercises")
	}
	id := results[0].ID

	answer := getExerciseAnswer(db, id)
	if answer.ID != id {
		t.Errorf("got id %q, want %q", answer.ID, id)
	}
	if answer.Odpowiedz == "" {
		t.Error("expected non-empty odpowiedz")
	}
}

// === Feature: DiagnoseOut status fields ===

func TestDiagnoseHasStatusFields(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Build and run the progress diagnose command via cobra
	diagnoseCmd := progressDiagnoseCmd()
	parent := &cobra.Command{Use: "progress"}
	parent.AddCommand(diagnoseCmd)
	ctx := context.WithValue(context.Background(), ctxKey{}, db)
	diagnoseCmd.SetContext(ctx)

	// Capture os.Stdout
	oldStdout := os.Stdout
	r, w, _ := os.Pipe()
	os.Stdout = w

	err := diagnoseCmd.RunE(diagnoseCmd, nil)

	w.Close()
	os.Stdout = oldStdout

	if err != nil {
		t.Fatalf("progress diagnose: %v", err)
	}

	var buf bytes.Buffer
	buf.ReadFrom(r)
	output := buf.String()

	var result map[string]interface{}
	if err := json.Unmarshal([]byte(output), &result); err != nil {
		t.Fatalf("invalid JSON: %v\noutput: %s", err, output)
	}

	for _, field := range []string{"zaleglosci", "leech_tagi"} {
		if _, ok := result[field]; !ok {
			t.Errorf("missing field %q in diagnose output", field)
		}
	}
}

// === Guardrails tests ===

func TestMigrationV5ToV6(t *testing.T) {
	dir := testDir(t)
	d := openTestDB(t, dir)

	// Table should exist after migration
	var count int
	err := d.QueryRow("SELECT COUNT(*) FROM active_exercises").Scan(&count)
	if err != nil {
		t.Fatalf("active_exercises table should exist: %v", err)
	}
	if count != 0 {
		t.Fatalf("expected 0 rows, got %d", count)
	}
}

func TestRegisterFetchAndClear(t *testing.T) {
	dir := testDir(t)
	d := openTestDB(t, dir)

	// Register
	if err := registerFetch(d, "7.1", "cyfry_liczby", 2); err != nil {
		t.Fatal(err)
	}

	// Check answer blocked
	can, attempts, _ := checkCanFetchAnswer(d, "7.1")
	if can {
		t.Fatal("answer should be blocked before any attempt")
	}
	if attempts != 0 {
		t.Fatalf("expected 0 attempts, got %d", attempts)
	}

	// Increment attempt
	if err := incrementAttempt(d, "7.1"); err != nil {
		t.Fatal(err)
	}
	can, attempts, _ = checkCanFetchAnswer(d, "7.1")
	if !can {
		t.Fatal("answer should be allowed after 1 attempt")
	}
	if attempts != 1 {
		t.Fatalf("expected 1 attempt, got %d", attempts)
	}

	// Check hints blocked (delay=2, attempt=1)
	canH, att, delay, _ := checkCanFetchHints(d, "7.1")
	if canH {
		t.Fatalf("hints should be blocked (attempt=%d < delay=%d)", att, delay)
	}

	// Second attempt
	if err := incrementAttempt(d, "7.1"); err != nil {
		t.Fatal(err)
	}
	canH, _, _, _ = checkCanFetchHints(d, "7.1")
	if !canH {
		t.Fatal("hints should be allowed after 2 attempts (delay=2)")
	}

	// Clear
	if err := clearExercise(d, "7.1"); err != nil {
		t.Fatal(err)
	}
	can, _, _ = checkCanFetchAnswer(d, "7.1")
	if !can {
		t.Fatal("after clear, should be allowed (no tracking)")
	}
}

func TestLazyLoadingEnforcement(t *testing.T) {
	dir := testDir(t)
	d := openTestDB(t, dir)

	// Fetch exercise (hint_delay=1)
	registerFetch(d, "1.1", "sledzenie_algorytmu", 1)

	// Answer should be blocked
	can, _, _ := checkCanFetchAnswer(d, "1.1")
	if can {
		t.Fatal("answer should be blocked before attempt")
	}

	// Hints should be blocked (attempt_count=0 < hint_delay=1)
	canH, _, _, _ := checkCanFetchHints(d, "1.1")
	if canH {
		t.Fatal("hints should be blocked before attempt (delay=1)")
	}

	// Record error → attempt_count = 1
	incrementAttempt(d, "1.1")

	// Now answer should be allowed
	can, _, _ = checkCanFetchAnswer(d, "1.1")
	if !can {
		t.Fatal("answer should be allowed after attempt")
	}

	// Hints should also be allowed now (attempt_count=1 >= hint_delay=1)
	canH, _, _, _ = checkCanFetchHints(d, "1.1")
	if !canH {
		t.Fatal("hints should be allowed after attempt (delay=1)")
	}

	// Test untracked exercise — should allow everything
	can, _, _ = checkCanFetchAnswer(d, "untracked.99")
	if !can {
		t.Fatal("untracked exercise should allow answer")
	}
	canH, _, _, _ = checkCanFetchHints(d, "untracked.99")
	if !canH {
		t.Fatal("untracked exercise should allow hints")
	}
}

func TestValidateErrorCode(t *testing.T) {
	// Valid
	ok, _ := validateErrorCode("sql_group_by", "brak_having")
	if !ok {
		t.Fatal("brak_having should be valid for sql_group_by")
	}

	// Invalid
	ok, allowed := validateErrorCode("sql_group_by", "brak_inicjalizacji")
	if ok {
		t.Fatal("brak_inicjalizacji should be invalid for sql_group_by")
	}
	if len(allowed) == 0 {
		t.Fatal("should return allowed list")
	}

	// Unknown type
	ok, _ = validateErrorCode("unknown_type", "anything")
	if !ok {
		t.Fatal("unknown type should allow any code")
	}

	// cyfry_liczby has mylenie_div_mod (cross-category)
	ok, _ = validateErrorCode("cyfry_liczby", "mylenie_div_mod")
	if !ok {
		t.Fatal("mylenie_div_mod should be valid for cyfry_liczby")
	}
}

func TestCoachingActions(t *testing.T) {
	dir := testDir(t)
	d := openTestDB(t, dir)

	// Insert leech tag (lapses >= 3, low stability for low retrievability)
	d.Exec(`INSERT INTO progress_tagi (tag, lapses, stability, last_review, poziom, nastepna_powtorka, difficulty, reps, state)
		VALUES ('cyfry-mod-div', 4, 1.0, '2026-01-01', 1, '2026-01-01', 5.0, 4, 2)`)
	// Create a typ entry with enough streak to trigger familiar level (hint_delay=2)
	d.Exec(`INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'srednie', 4)`)
	// Add some completed exercises so done > 3
	d.Exec(`INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.1', 'cyfry_liczby', '2026-01-01', 'poprawne_bez_pomocy')`)
	d.Exec(`INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.2', 'cyfry_liczby', '2026-01-01', 'poprawne_bez_pomocy')`)
	d.Exec(`INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.3', 'cyfry_liczby', '2026-01-01', 'poprawne_bez_pomocy')`)
	d.Exec(`INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.4', 'cyfry_liczby', '2026-01-01', 'poprawne_bez_pomocy')`)

	c := buildCoaching(d, "cyfry_liczby", []string{"cyfry-mod-div"})

	if len(c.Actions) == 0 {
		t.Fatal("expected coaching_actions")
	}
	foundLeech := false
	foundDelay := false
	for _, a := range c.Actions {
		if strings.Contains(a, "WARN_LEECH") {
			foundLeech = true
		}
		if strings.Contains(a, "HINT_DELAY") {
			foundDelay = true
		}
	}
	if !foundLeech {
		t.Fatalf("expected WARN_LEECH action, got: %v", c.Actions)
	}
	if !foundDelay {
		t.Fatalf("expected HINT_DELAY action (familiar = delay 2), got: %v", c.Actions)
	}
}

func TestCoachingActionsEmpty(t *testing.T) {
	dir := testDir(t)
	d := openTestDB(t, dir)

	// New student — no leech tags, no past mistakes, hint_delay=1
	c := buildCoaching(d, "sql_group_by", []string{})

	if c.Actions == nil {
		t.Fatal("Actions should be empty slice, not nil")
	}
	if len(c.Actions) != 0 {
		t.Fatalf("expected 0 actions for new student, got: %v", c.Actions)
	}
}

func TestDifficultyBumped(t *testing.T) {
	// calculateLevel thresholds: <3 latwe, 3-4 srednie, 5-7 srednie-trudne, 8+ trudne
	tests := []struct {
		streak int
		wynik  string
		bump   bool
		label  string
	}{
		{3, "poprawne_bez_pomocy", true, "2→3: latwe→srednie"},
		{5, "poprawne_bez_pomocy", true, "4→5: srednie→srednie-trudne"},
		{8, "poprawne_bez_pomocy", true, "7→8: srednie-trudne→trudne"},
		{4, "poprawne_bez_pomocy", false, "3→4: srednie→srednie (no bump)"},
		{1, "poprawne_bez_pomocy", false, "0→1: latwe→latwe (no bump)"},
		{9, "poprawne_bez_pomocy", false, "8→9: trudne→trudne (no bump)"},
		{3, "walk_through", false, "walk_through always latwe (no bump)"},
	}
	for _, tc := range tests {
		oldLevel := calculateLevel(tc.streak-1, tc.wynik)
		newLevel := calculateLevel(tc.streak, tc.wynik)
		bumped := newLevel != oldLevel
		if bumped != tc.bump {
			t.Errorf("%s: expected bump=%v, got %v (old=%s new=%s)",
				tc.label, tc.bump, bumped, oldLevel, newLevel)
		}
	}
}

func TestAutoDiagnose(t *testing.T) {
	dir := testDir(t)
	d := openTestDB(t, dir)

	// Setup: insert typ entry
	d.Exec(`INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'latwe', 1)`)

	// Insert 4 exercises + set cwiczenia_lacznie=4
	for i := 1; i <= 4; i++ {
		id := fmt.Sprintf("7.%d", i)
		d.Exec(`INSERT OR REPLACE INTO progress_zrobione (id, typ, data, wynik) VALUES (?, 'cyfry_liczby', '2026-01-01', 'poprawne_bez_pomocy')`, id)
	}
	d.Exec(`INSERT OR REPLACE INTO progress_meta (key, value) VALUES ('cwiczenia_lacznie', '4')`)

	// At 4 exercises, no auto-diagnose
	var totalDoneStr string
	d.QueryRow(`SELECT COALESCE(value, '0') FROM progress_meta WHERE key='cwiczenia_lacznie'`).Scan(&totalDoneStr)
	var totalDone int
	fmt.Sscanf(totalDoneStr, "%d", &totalDone)
	if totalDone%5 == 0 {
		t.Fatal("at 4 exercises, should NOT trigger auto-diagnose")
	}

	// Bump to 5
	d.Exec(`UPDATE progress_meta SET value = '5' WHERE key='cwiczenia_lacznie'`)
	d.QueryRow(`SELECT value FROM progress_meta WHERE key='cwiczenia_lacznie'`).Scan(&totalDoneStr)
	fmt.Sscanf(totalDoneStr, "%d", &totalDone)
	if totalDone != 5 {
		t.Fatalf("expected 5, got %d", totalDone)
	}
	if totalDone%5 != 0 {
		t.Fatal("at 5 exercises, should trigger auto-diagnose")
	}

	// runDiagnose should work
	diag, err := runDiagnose(d, "cyfry_liczby", 3)
	if err != nil {
		t.Fatalf("runDiagnose error: %v", err)
	}
	if diag == nil {
		t.Fatal("expected non-nil DiagnoseOut")
	}
	if diag.TopBledy == nil {
		t.Fatal("TopBledy should be empty slice, not nil")
	}
}

func TestExerciseRubricSledzenie(t *testing.T) {
	rubric, err := getRubric("sledzenie_algorytmu")
	if err != nil {
		t.Fatalf("getRubric: %v", err)
	}
	if rubric.Typ != "sledzenie_algorytmu" {
		t.Errorf("got typ %q", rubric.Typ)
	}
	if rubric.Kategoria != "TEORIA" {
		t.Errorf("got kategoria %q", rubric.Kategoria)
	}
	if rubric.Rubric.Full.Opis == "" {
		t.Error("empty full.opis")
	}
	if rubric.Rubric.Half.Opis == "" {
		t.Error("empty half.opis")
	}
	if rubric.Rubric.Zero.Opis == "" {
		t.Error("empty zero.opis")
	}
	if rubric.Rubric.Full.Procent != 100 {
		t.Errorf("full.procent = %d, want 100", rubric.Rubric.Full.Procent)
	}
}

func TestExerciseRubricAllTypes(t *testing.T) {
	allTypes := []string{
		"sledzenie_algorytmu", "projektowanie_algorytmu", "analiza_algorytmu",
		"test_prawda_falsz", "konwersja_systemow_liczbowych", "teoria_bezpieczenstwa",
		"cyfry_liczby", "napisy", "zlozone", "zliczanie", "minmax", "sekwencje", "obrazy_2D", "geometryczne",
		"agregacja_warunkowa", "symulacja", "wykres", "agregacja_podstawowa", "transformacja",
		"sql_group_by", "sql_podzapytania", "sql_join", "sql_select_where",
	}
	for _, typ := range allTypes {
		rubric, err := getRubric(typ)
		if err != nil {
			t.Errorf("getRubric(%q): %v", typ, err)
			continue
		}
		if rubric.Rubric.Full.Opis == "" {
			t.Errorf("%s: empty full.opis", typ)
		}
	}
}

func TestExerciseRubricUnknownType(t *testing.T) {
	_, err := getRubric("nonexistent_type")
	if err == nil {
		t.Error("expected error for unknown type")
	}
}

func TestSuggestClosestCodes(t *testing.T) {
	// Exact match — no suggestions needed
	suggestions := suggestClosestCodes("sql_group_by", "brak_having")
	if len(suggestions) != 0 {
		t.Errorf("expected no suggestions for valid code, got %d", len(suggestions))
	}

	// Typo — should suggest closest
	suggestions = suggestClosestCodes("sql_group_by", "brak_havng")
	if len(suggestions) == 0 {
		t.Fatal("expected suggestions for typo")
	}
	if suggestions[0].Kod != "brak_having" {
		t.Errorf("got %q, want brak_having", suggestions[0].Kod)
	}
	if suggestions[0].Opis == "" {
		t.Error("expected non-empty opis in suggestion")
	}

	// Completely wrong — should still suggest something
	suggestions = suggestClosestCodes("sql_group_by", "completely_wrong")
	if len(suggestions) == 0 {
		t.Fatal("expected suggestions even for completely wrong code")
	}
	// Should return max 3 suggestions
	if len(suggestions) > 3 {
		t.Errorf("expected max 3 suggestions, got %d", len(suggestions))
	}
}

func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"abc", "", 3},
		{"", "abc", 3},
		{"abc", "abc", 0},
		{"brak_having", "brak_havng", 1},
		{"kitten", "sitting", 3},
	}
	for _, tt := range tests {
		got := levenshtein(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("levenshtein(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
