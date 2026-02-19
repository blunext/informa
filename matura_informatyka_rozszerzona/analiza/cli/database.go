package main

import (
	"database/sql"
	"fmt"
	"os"
	"path/filepath"

	_ "modernc.org/sqlite"
)

const currentSchemaVersion = 4

// OpenDB opens progress DB as main, attaches matura.db as "data" read-only.
// Returns the DB, whether matura.db was attached, and any error.
func OpenDB(dbDir string) (*sql.DB, bool, error) {
	progressPath := filepath.Join(dbDir, "matura_progress.db")
	maturaPath := filepath.Join(dbDir, "matura.db")

	db, err := sql.Open("sqlite", progressPath)
	if err != nil {
		return nil, false, fmt.Errorf("open progress db: %w", err)
	}
	db.SetMaxOpenConns(1)

	// Enable WAL mode for better concurrency
	if _, err := db.Exec("PRAGMA journal_mode=WAL"); err != nil {
		db.Close()
		return nil, false, fmt.Errorf("set WAL mode: %w", err)
	}

	// Initialize progress schema if needed
	if err := initProgressSchema(db); err != nil {
		db.Close()
		return nil, false, fmt.Errorf("init progress schema: %w", err)
	}

	// Attach matura.db read-only
	attached := false
	if _, err := os.Stat(maturaPath); err == nil {
		_, err = db.Exec("ATTACH DATABASE ? AS data", maturaPath)
		if err != nil {
			db.Close()
			return nil, false, fmt.Errorf("attach matura.db: %w", err)
		}
		attached = true
	}

	return db, attached, nil
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
	DROP TABLE IF EXISTS benchmarki;

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
		czesc INTEGER,
		pulapki TEXT,
		sciezka_danych TEXT,
		pliki_danych TEXT
	);

	CREATE TABLE cheatsheets (
		kategoria TEXT PRIMARY KEY,
		content TEXT
	);

	CREATE TABLE benchmarki (
		typ TEXT PRIMARY KEY,
		benchmark_sek INTEGER,
		zrodlo TEXT
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
		wynik TEXT,
		czas_sek INTEGER
	);

	CREATE TABLE IF NOT EXISTS progress_bledy (
		id INTEGER PRIMARY KEY AUTOINCREMENT,
		exercise_id TEXT NOT NULL,
		typ TEXT NOT NULL,
		blad_kod TEXT NOT NULL,
		blad_opis TEXT,
		hint_level INTEGER DEFAULT 0,
		data TEXT NOT NULL
	);

	CREATE TABLE IF NOT EXISTS progress_tagi (
		tag TEXT PRIMARY KEY,
		poziom INTEGER DEFAULT 0,
		nastepna_powtorka TEXT,
		stability REAL DEFAULT 0,
		difficulty REAL DEFAULT 5.0,
		lapses INTEGER DEFAULT 0,
		reps INTEGER DEFAULT 0,
		state INTEGER DEFAULT 0,
		last_review TEXT
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

	CREATE TABLE IF NOT EXISTS pulapki_przejrzane (
		id TEXT PRIMARY KEY,
		typ TEXT,
		data TEXT,
		trafienia INTEGER,
		total INTEGER
	);

	CREATE INDEX IF NOT EXISTS idx_tagi_powtorka ON progress_tagi(nastepna_powtorka);
	CREATE INDEX IF NOT EXISTS idx_pulapki_typ ON pulapki_przejrzane(typ);
	CREATE INDEX IF NOT EXISTS idx_bledy_kod ON progress_bledy(blad_kod);
	CREATE INDEX IF NOT EXISTS idx_bledy_typ ON progress_bledy(typ);
	`
	_, err := db.Exec(schema, currentSchemaVersion)
	return err
}

// Migration uses Apply function for idempotent, transactional migrations.
type Migration struct {
	Version int
	Apply   func(tx *sql.Tx) error
}

var migrations = []Migration{
	{Version: 2, Apply: func(tx *sql.Tx) error {
		if _, err := tx.Exec(`
			CREATE TABLE IF NOT EXISTS pulapki_przejrzane (
				id TEXT PRIMARY KEY,
				typ TEXT,
				data TEXT,
				trafienia INTEGER,
				total INTEGER
			)`); err != nil {
			return err
		}
		_, err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_pulapki_typ ON pulapki_przejrzane(typ)")
		return err
	}},
	{Version: 3, Apply: func(tx *sql.Tx) error {
		// Idempotent: check if czas_sek column already exists
		var count int
		if err := tx.QueryRow("SELECT COUNT(*) FROM pragma_table_info('progress_zrobione') WHERE name='czas_sek'").Scan(&count); err != nil {
			return err
		}
		if count == 0 {
			if _, err := tx.Exec("ALTER TABLE progress_zrobione ADD COLUMN czas_sek INTEGER"); err != nil {
				return err
			}
		}
		if _, err := tx.Exec(`CREATE TABLE IF NOT EXISTS progress_bledy (
			id INTEGER PRIMARY KEY AUTOINCREMENT,
			exercise_id TEXT NOT NULL,
			typ TEXT NOT NULL,
			blad_kod TEXT NOT NULL,
			blad_opis TEXT,
			hint_level INTEGER DEFAULT 0,
			data TEXT NOT NULL
		)`); err != nil {
			return err
		}
		if _, err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_bledy_kod ON progress_bledy(blad_kod)"); err != nil {
			return err
		}
		_, err := tx.Exec("CREATE INDEX IF NOT EXISTS idx_bledy_typ ON progress_bledy(typ)")
		return err
	}},
	{Version: 4, Apply: func(tx *sql.Tx) error {
		cols := []struct{ name, def string }{
			{"stability", "REAL DEFAULT 0"},
			{"difficulty", "REAL DEFAULT 5.0"},
			{"lapses", "INTEGER DEFAULT 0"},
			{"reps", "INTEGER DEFAULT 0"},
			{"state", "INTEGER DEFAULT 0"},
			{"last_review", "TEXT"},
		}
		for _, c := range cols {
			var count int
			if err := tx.QueryRow(
				"SELECT COUNT(*) FROM pragma_table_info('progress_tagi') WHERE name=?",
				c.name).Scan(&count); err != nil {
				return fmt.Errorf("check column %s: %w", c.name, err)
			}
			if count == 0 {
				if _, err := tx.Exec(
					"ALTER TABLE progress_tagi ADD COLUMN " + c.name + " " + c.def); err != nil {
					return fmt.Errorf("add column %s: %w", c.name, err)
				}
			}
		}
		// Seed FSRS state from legacy poziom for existing tags
		if _, err := tx.Exec(`UPDATE progress_tagi SET
			stability = CASE poziom
				WHEN 0 THEN 0.4 WHEN 1 THEN 1.2 WHEN 2 THEN 3.2
				WHEN 3 THEN 7.0 WHEN 4 THEN 21.0 ELSE 0.4 END,
			difficulty = 5.0,
			state = CASE WHEN poziom >= 2 THEN 2 ELSE 1 END,
			last_review = nastepna_powtorka
		WHERE stability = 0 OR stability IS NULL`); err != nil {
			return fmt.Errorf("seed FSRS state: %w", err)
		}
		return nil
	}},
}

func migrateProgress(db *sql.DB, fromVersion int) error {
	for _, m := range migrations {
		if m.Version > fromVersion {
			tx, err := db.Begin()
			if err != nil {
				return fmt.Errorf("begin migration v%d: %w", m.Version, err)
			}
			if err := m.Apply(tx); err != nil {
				tx.Rollback()
				return fmt.Errorf("migration to v%d failed: %w", m.Version, err)
			}
			if _, err := tx.Exec("UPDATE schema_version SET version = ?", m.Version); err != nil {
				tx.Rollback()
				return fmt.Errorf("update version to v%d: %w", m.Version, err)
			}
			if err := tx.Commit(); err != nil {
				return fmt.Errorf("commit migration v%d: %w", m.Version, err)
			}
		}
	}
	return nil
}
