package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

const currentSchemaVersion = 1

// OpenDB opens progress DB as main, attaches matura.db as "data" read-only.
func OpenDB(dbDir string) (*sql.DB, error) {
	progressPath := filepath.Join(dbDir, "matura_progress.db")
	maturaPath := filepath.Join(dbDir, "matura.db")

	db, err := sql.Open("sqlite", progressPath)
	if err != nil {
		return nil, fmt.Errorf("open progress db: %w", err)
	}

	// Enable WAL mode for better concurrency
	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		db.Close()
		return nil, fmt.Errorf("set WAL mode: %w", err)
	}

	// Initialize progress schema if needed
	if err := initProgressSchema(db); err != nil {
		db.Close()
		return nil, fmt.Errorf("init progress schema: %w", err)
	}

	// Attach matura.db read-only
	if _, err := os.Stat(maturaPath); err == nil {
		_, err = db.Exec("ATTACH DATABASE ? AS data", maturaPath)
		if err != nil {
			db.Close()
			return nil, fmt.Errorf("attach matura.db: %w", err)
		}
	}

	return db, nil
}

// OpenDataDB opens matura.db directly for import (write mode).
func OpenDataDB(dbDir string) (*sql.DB, error) {
	maturaPath := filepath.Join(dbDir, "matura.db")

	db, err := sql.Open("sqlite", maturaPath)
	if err != nil {
		return nil, fmt.Errorf("open matura.db: %w", err)
	}

	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		db.Close()
		return nil, fmt.Errorf("set WAL mode: %w", err)
	}

	return db, nil
}

// CreateDataSchema creates tables in matura.db (for import).
func CreateDataSchema(db *sql.DB) error {
	schema := `
	DROP TABLE IF EXISTS cwiczenia;
	DROP TABLE IF EXISTS egzamin;
	DROP TABLE IF EXISTS cheatsheets;

	CREATE TABLE cwiczenia (
		id TEXT PRIMARY KEY,
		typ_nr INTEGER,
		typ_nazwa TEXT,
		kategoria TEXT,
		trudnosc TEXT,
		punkty INTEGER,
		zrodlo TEXT,
		tresc TEXT,
		odpowiedz TEXT,
		wskazowki TEXT,
		typowe_bledy TEXT,
		tagi TEXT
	);

	CREATE TABLE egzamin (
		id TEXT PRIMARY KEY,
		rok INTEGER,
		numer_zadania INTEGER,
		numer_podzadania TEXT,
		tytul TEXT,
		kontekst TEXT,
		tresc TEXT,
		odpowiedz TEXT,
		zasady_oceniania TEXT,
		typ_zadania TEXT,
		kategoria TEXT,
		punkty INTEGER,
		pulapki TEXT,
		sciezka_danych TEXT,
		pliki_danych TEXT
	);

	CREATE TABLE cheatsheets (
		kategoria TEXT PRIMARY KEY,
		content TEXT
	);

	CREATE INDEX idx_cwiczenia_typ ON cwiczenia(typ_nazwa, trudnosc);
	CREATE INDEX idx_egzamin_typ ON egzamin(typ_zadania);
	CREATE INDEX idx_egzamin_rok ON egzamin(rok, numer_zadania);
	`
	_, err := db.Exec(schema)
	return err
}

func initProgressSchema(db *sql.DB) error {
	// Check if schema_version table exists
	var count int
	err := db.QueryRow("SELECT count(*) FROM sqlite_master WHERE type='table' AND name='schema_version'").Scan(&count)
	if err != nil {
		return err
	}

	if count == 0 {
		// Fresh database — create all tables
		return createProgressSchema(db)
	}

	// Check version and migrate if needed
	var version int
	err = db.QueryRow("SELECT version FROM schema_version").Scan(&version)
	if err != nil {
		return err
	}

	if version == currentSchemaVersion {
		return nil
	}

	if version < currentSchemaVersion {
		return migrateProgress(db, version)
	}

	// Future version — incompatible
	return fmt.Errorf("progress DB v%d niekompatybilny z v%d", version, currentSchemaVersion)
}

func createProgressSchema(db *sql.DB) error {
	schema := `
	CREATE TABLE IF NOT EXISTS schema_version (version INTEGER PRIMARY KEY);
	INSERT OR REPLACE INTO schema_version VALUES (?);

	CREATE TABLE IF NOT EXISTS progress_meta (
		key TEXT PRIMARY KEY,
		value TEXT
	);

	CREATE TABLE IF NOT EXISTS progress_typy (
		typ TEXT PRIMARY KEY,
		poziom_trudnosci TEXT DEFAULT 'latwe',
		streak INTEGER DEFAULT 0
	);

	CREATE TABLE IF NOT EXISTS progress_zrobione (
		id TEXT PRIMARY KEY,
		typ TEXT,
		data TEXT,
		wynik TEXT
	);

	CREATE TABLE IF NOT EXISTS progress_tagi (
		tag TEXT PRIMARY KEY,
		poziom INTEGER DEFAULT 0,
		nastepna_powtorka TEXT
	);

	CREATE TABLE IF NOT EXISTS matura_zrobione (
		id TEXT PRIMARY KEY,
		typ TEXT,
		data TEXT,
		punkty INTEGER,
		max_punkty INTEGER
	);

	CREATE TABLE IF NOT EXISTS probne_matury (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		rok INTEGER,
		data TEXT,
		czas_min INTEGER,
		wynik_pkt INTEGER,
		max_pkt INTEGER,
		procent REAL,
		per_kategoria TEXT,
		przerwany BOOLEAN DEFAULT 0
	);

	CREATE INDEX IF NOT EXISTS idx_tagi_powtorka ON progress_tagi(nastepna_powtorka);
	`
	_, err := db.Exec(schema, currentSchemaVersion)
	return err
}

// Migration table for future schema changes
type Migration struct {
	Version int
	SQL     string
}

var migrations = []Migration{
	// Future migrations go here:
	// {Version: 2, SQL: "ALTER TABLE ..."},
}

func migrateProgress(db *sql.DB, fromVersion int) error {
	for _, m := range migrations {
		if m.Version > fromVersion {
			if _, err := db.Exec(m.SQL); err != nil {
				return fmt.Errorf("migration to v%d failed: %w", m.Version, err)
			}
			if _, err := db.Exec("UPDATE schema_version SET version = ?", m.Version); err != nil {
				return err
			}
		}
	}
	return nil
}
