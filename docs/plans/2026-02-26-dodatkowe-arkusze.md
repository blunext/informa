# Dodatkowe Arkusze CKE — Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Migrate CKE exam IDs from `YYYY.Z.S` to `YYYYM.Z.S` format, extend CLI to support multiple exam sessions (maj/czerwiec/próbna/przykładowy), and prepare the system for ingesting ~5 new exam papers from nowa formuła (2023+).

**Architecture:** Three-phase approach: (1) JSON file migration — rename files, re-ID 230 subtasks, add `sesja` field; (2) Go CLI changes — schema, importer, commands, progress migration; (3) New exam paper creation (separate effort). Phases 1-2 are this plan. Phase 3 is future work.

**Tech Stack:** Python 3 (migration script), Go (CLI), SQLite (matura.db + progress.db)

**Design doc:** `docs/plans/2026-02-26-dodatkowe-arkusze-design.md`

**Key context:**
- Current schema version: 6 (next migration = v7)
- 11 JSON files: `matura_YYYY.json` with 230 subtasks total
- ID formats: 2014 uses `2014.1a` (letter suffix), 2015+ uses `YYYY.Z.S` (numeric)
- Migration rule for all years: insert session letter after 4-digit year (`2025.1.1` → `2025M.1.1`, `2014.1a` → `2014M.1a`)
- `matura_indeks.json` references `matura_YYYY.json` in `plik` field
- `matura.db` is DROP+CREATE on import, so schema change is automatic
- `progress.db` needs schema migration v7 (re-ID entries, add `sesja` column)
- Hardcoded IDs in: `main_test.go` (lines 154, 1355, 1357, 1410, 2480-2546), `commands.go` (lines 1797, 2695 — help text)
- User preference: NEVER run `git commit` — only write/edit code
- Build command: `cd analiza/cli && ./build.sh`
- Test commands: `go test ./...`, `./test_qa.sh`

---

## Phase 1: JSON Migration

### Task 1: Create and run JSON migration script

**Files:**
- Create: `analiza/json/migrate_ids.py` (temporary migration script)
- Modify: `analiza/json/matura_2014.json` through `matura_2025.json` (rename + content)
- Modify: `analiza/json/matura_indeks.json`

**Step 1: Write the migration script**

```python
#!/usr/bin/env python3
"""Migrate matura JSON files: rename YYYY→YYYYM, add sesja field, re-ID subtasks."""

import json
import os
import sys

JSON_DIR = os.path.dirname(os.path.abspath(__file__))

YEARS = [2014, 2015, 2016, 2017, 2018, 2019, 2021, 2022, 2023, 2024, 2025]

def migrate_id(old_id: str) -> str:
    """Insert 'M' after 4-digit year prefix. Works for both formats:
    2025.1.1 → 2025M.1.1, 2014.1a → 2014M.1a"""
    return old_id[:4] + "M" + old_id[4:]

def migrate_exam_file(year: int) -> None:
    old_path = os.path.join(JSON_DIR, f"matura_{year}.json")
    new_path = os.path.join(JSON_DIR, f"matura_{year}M.json")

    if not os.path.exists(old_path):
        print(f"SKIP: {old_path} not found")
        return

    with open(old_path, "r", encoding="utf-8") as f:
        data = json.load(f)

    # Add sesja field
    data["sesja"] = "maj"

    # Re-ID all podzadania
    for zadanie in data["zadania"]:
        for pod in zadanie["podzadania"]:
            pod["id"] = migrate_id(pod["id"])

    with open(new_path, "w", encoding="utf-8") as f:
        json.dump(data, f, ensure_ascii=False, indent=2)

    os.remove(old_path)
    print(f"OK: {os.path.basename(old_path)} → {os.path.basename(new_path)}")

def migrate_indeks() -> None:
    path = os.path.join(JSON_DIR, "matura_indeks.json")
    with open(path, "r", encoding="utf-8") as f:
        data = json.load(f)

    # Update per_rok keys (stay as strings of years, no change needed)
    # Update lata (stay as year strings, no change needed)

    # Re-ID all podzadania entries
    for entry in data["wszystkie_podzadania"]:
        entry["id"] = migrate_id(entry["id"])
        # Update plik reference: matura_YYYY.json → matura_YYYYM.json
        old_plik = entry["plik"]
        year_str = old_plik.replace("matura_", "").replace(".json", "")
        entry["plik"] = f"matura_{year_str}M.json"

    with open(path, "w", encoding="utf-8") as f:
        json.dump(data, f, ensure_ascii=False, indent=2)

    print(f"OK: matura_indeks.json updated ({len(data['wszystkie_podzadania'])} entries)")

if __name__ == "__main__":
    for year in YEARS:
        migrate_exam_file(year)
    migrate_indeks()
    print("\nDone. Verify with: ls analiza/json/matura_*M.json | wc -l  (expect 11)")
```

**Step 2: Run the migration script**

Run: `cd /Users/blt1wz/priv/informa/matura_informatyka_rozszerzona && python3 analiza/json/migrate_ids.py`

Expected: 11 "OK" lines + indeks update. Old files removed, new `*M.json` files created.

**Step 3: Verify migration**

Run: `ls analiza/json/matura_*M.json | wc -l`
Expected: `11`

Run: `python3 -c "import json; d=json.load(open('analiza/json/matura_2025M.json')); print(d['sesja'], d['zadania'][0]['podzadania'][0]['id'])"`
Expected: `maj 2025M.1.1`

Run: `python3 -c "import json; d=json.load(open('analiza/json/matura_2014M.json')); print(d['sesja'], d['zadania'][0]['podzadania'][0]['id'])"`
Expected: `maj 2014M.1a`

Run: `python3 -c "import json; d=json.load(open('analiza/json/matura_indeks.json')); e=d['wszystkie_podzadania'][0]; print(e['id'], e['plik'])"`
Expected: `2014M.1a matura_2014M.json`

**Step 4: Remove migration script**

Run: `rm analiza/json/migrate_ids.py`

---

## Phase 2: Go CLI Changes

### Task 2: Update Go types

**Files:**
- Modify: `analiza/cli/types.go`

**Step 1: Add Sesja field to MaturaExam struct**

Find `MaturaExam` struct in types.go. Add `Sesja` field:

```go
type MaturaExam struct {
    Rok        int           `json:"rok"`
    Sesja      string        `json:"sesja"`    // NEW: "maj", "czerwiec", "probna", "przykladowy"
    Formula    string        `json:"formula"`
    CzasMinuty int           `json:"czas_minuty"`
    // ... rest unchanged
}
```

**Step 2: Add Sesja to ExamListEntry**

Find `ExamListEntry` struct. Add `Sesja` field:

```go
type ExamListEntry struct {
    Rok          int      `json:"rok"`
    Sesja        string   `json:"sesja"`    // NEW
    Formula      string   `json:"formula"`
    // ... rest unchanged
}
```

**Step 3: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success (no errors)

---

### Task 3: Update DB schema (matura.db — egzamin table)

**Files:**
- Modify: `analiza/cli/database.go` (schema definition, ~line 92)

**Step 1: Add sesja column to egzamin CREATE TABLE**

In `database.go`, find the `CREATE TABLE egzamin` statement (~line 92). Add `sesja TEXT DEFAULT 'maj'` column after `rok`:

```sql
CREATE TABLE egzamin (
    id TEXT PRIMARY KEY,
    rok INTEGER,
    sesja TEXT DEFAULT 'maj',     -- NEW
    numer_zadania INTEGER,
    ...
);
```

**Step 2: Add index for (rok, sesja)**

Find the `CREATE INDEX idx_egzamin_rok` line (~line 124). Change it to:

```sql
CREATE INDEX idx_egzamin_rok ON egzamin(rok, sesja, numer_zadania);
```

**Step 3: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success

---

### Task 4: Update importer to read and store sesja

**Files:**
- Modify: `analiza/cli/importer.go` (~line 244, ImportExams function)

**Step 1: Add sesja to INSERT statement**

In `ImportExams()`, find the `stmt.Prepare` with `INSERT INTO egzamin`. Add `sesja` column:

```go
stmt, err := tx.Prepare(`INSERT INTO egzamin (id, rok, sesja, numer_zadania, numer_podzadania, tytul, kontekst, tresc, odpowiedz, zasady_oceniania, typ_zadania, kategoria, punkty, czesc, pulapki, sciezka_danych, pliki_danych)
    VALUES (?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?, ?)`)
```

**Step 2: Add exam.Sesja to Exec call**

In the loop where `stmt.Exec` is called, add `exam.Sesja` after `exam.Rok`:

```go
_, err = stmt.Exec(
    sub.ID, exam.Rok, exam.Sesja, task.Numer, sub.Numer,
    // ... rest unchanged
)
```

**Step 3: Handle missing sesja (backward compatibility)**

After unmarshaling the JSON, default sesja if empty:

```go
if exam.Sesja == "" {
    exam.Sesja = "maj"
}
```

**Step 4: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success

---

### Task 5: Update exam list command

**Files:**
- Modify: `analiza/cli/commands.go` (~line 2461, examListCmd)

**Step 1: Update query to group by (rok, sesja)**

Find the query in `examListCmd` (~line 2473). Change to:

```sql
SELECT rok, MIN(sesja) as sesja, SUM(punkty) as total
FROM data.egzamin
GROUP BY rok, sesja
ORDER BY rok, sesja
```

Note: `MIN(sesja)` is used because all rows in a (rok, sesja) group have the same sesja value.

**Step 2: Scan sesja from result**

In the `rows.Next()` loop, scan the new sesja column:

```go
if err := rows.Scan(&e.Rok, &e.Sesja, &e.TotalPkt); err != nil {
```

**Step 3: Update "done" check to include sesja**

Find the `probne_matury` query (~line 2510). Add sesja filter:

```go
d.QueryRow("SELECT procent FROM probne_matury WHERE rok = ? AND sesja = ? AND przerwany = 0 ORDER BY procent DESC LIMIT 1", e.Rok, e.Sesja).Scan(&procent)
```

**Step 4: Add --sesja filter flag**

Find where `--formula` flag is defined (~line 2594). Add `--sesja` flag:

```go
cmd.Flags().StringVar(&sesja, "sesja", "", "Filter by session: maj, czerwiec, probna, przykladowy")
```

Add filtering logic before returning results — if sesja is specified, filter entries.

**Step 5: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success

---

### Task 6: Update exam task command

**Files:**
- Modify: `analiza/cli/commands.go` (~line 1896, examTaskCmd)

**Step 1: Add --sesja flag with default "maj"**

```go
var rok, zadanie int
var sesja string  // NEW
// ...
cmd.Flags().StringVar(&sesja, "sesja", "maj", "Session: maj, czerwiec, probna, przykladowy")
```

**Step 2: Update query to filter by sesja**

In the query (~line 1909), add `AND e.sesja = ?`:

```sql
SELECT e.id, e.numer_podzadania, ...
FROM data.egzamin e
WHERE e.rok = ? AND e.sesja = ? AND e.numer_zadania = ?
ORDER BY e.id
```

Add `sesja` to params.

**Step 3: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success

---

### Task 7: Update exam meta command

**Files:**
- Modify: `analiza/cli/commands.go` (examMetaCmd — similar to exam task)

**Step 1: Add --sesja flag with default "maj"**

Same pattern as Task 6. Find `examMetaCmd` and add `--sesja` flag.

**Step 2: Update query to filter by sesja**

Add `AND e.sesja = ?` to the WHERE clause.

**Step 3: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success

---

### Task 8: Update exam save command

**Files:**
- Modify: `analiza/cli/commands.go` (~line 1976, examSaveCmd)

**Step 1: Add --sesja flag with default "maj"**

```go
var rok int
var sesja string    // NEW
var resultsJSON string
var czasMin int
// ...
cmd.Flags().StringVar(&sesja, "sesja", "maj", "Session: maj, czerwiec, probna, przykladowy")
```

**Step 2: Add sesja to probne_matury INSERT**

Find the `INSERT INTO probne_matury` statement (~line 2033). Add sesja:

```sql
INSERT INTO probne_matury (rok, sesja, data, czas_min, wynik_pkt, max_pkt, procent)
VALUES (?, ?, ?, ?, ?, ?, ?)
```

Add `sesja` to Exec args.

**Step 3: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success

---

### Task 9: Update cke get command

**Files:**
- Modify: `analiza/cli/commands.go` (~line 1608, ckeGetCmd)

**Step 1: Add optional --sesja flag**

```go
var typ, exclude, sesja string  // ADD sesja
// ...
cmd.Flags().StringVar(&sesja, "sesja", "", "Filter by session (empty = all)")
```

**Step 2: Add sesja filter to query**

If sesja is not empty, append `AND e.sesja = ?` to the query and add to params:

```go
if sesja != "" {
    query += " AND e.sesja = ?"
    params = append(params, sesja)
}
```

**Step 3: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success

---

### Task 10: Update data stats command

**Files:**
- Modify: `analiza/cli/commands.go` (~line 2854, dataStatsCmd)
- Modify: `analiza/cli/types.go` (DataStatsOut struct)

**Step 1: Add per-session stats to DataStatsOut**

In types.go, find `DataStatsOut` and add:

```go
type DataStatsOut struct {
    Cwiczenia   int            `json:"cwiczenia"`
    Podzadania  int            `json:"podzadania"`
    Cheatsheets int            `json:"cheatsheets"`
    PerSesja    map[string]int `json:"per_sesja,omitempty"`  // NEW
}
```

**Step 2: Query per-session counts**

In `dataStatsCmd`, after existing counts, add:

```go
out.PerSesja = make(map[string]int)
rows, _ := d.Query("SELECT COALESCE(sesja, 'maj'), COUNT(*) FROM data.egzamin GROUP BY sesja")
if rows != nil {
    defer rows.Close()
    for rows.Next() {
        var s string
        var c int
        rows.Scan(&s, &c)
        out.PerSesja[s] = c
    }
}
```

**Step 3: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success

---

### Task 11: Update cke save and help texts

**Files:**
- Modify: `analiza/cli/commands.go` (ckeSaveCmd ~line 1760, help texts ~lines 1797, 2695)

**Step 1: Update help text for --id flag in cke save**

Find line 1797: `"CKE subtask ID (e.g. 2025.1.1)"` → `"CKE subtask ID (e.g. 2025M.1.1)"`

**Step 2: Update any other help text referencing old ID format**

Find line 2695 or similar — update to new format.

**Step 3: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success

---

### Task 12: Add progress.db migration v7

**Files:**
- Modify: `analiza/cli/database.go` (~line 12 for version, ~line 360 for migrations)

**Step 1: Bump schema version**

Change `const currentSchemaVersion = 6` → `const currentSchemaVersion = 7`

**Step 2: Add migration v7**

After the v6 migration entry (the `active_exercises` one at ~line 360), add:

```go
{version: 7, fn: func(tx *sql.Tx) error {
    // Re-ID matura_zrobione entries: 2025.1.1 → 2025M.1.1
    // Rule: insert 'M' after first 4 chars (the year)
    if _, err := tx.Exec(`UPDATE matura_zrobione
        SET id = substr(id, 1, 4) || 'M' || substr(id, 5)
        WHERE id NOT LIKE '____M%' AND id NOT LIKE '____C%' AND id NOT LIKE '____P%' AND id NOT LIKE '____X%'`); err != nil {
        return err
    }
    // Add sesja column to probne_matury
    if _, err := tx.Exec(`ALTER TABLE probne_matury ADD COLUMN sesja TEXT DEFAULT 'maj'`); err != nil {
        return err
    }
    return nil
}},
```

**Step 3: Verify compilation**

Run: `cd analiza/cli && go build ./...`
Expected: success

---

### Task 13: Update Go unit tests

**Files:**
- Modify: `analiza/cli/main_test.go`

**Step 1: Update hardcoded IDs**

Replace all hardcoded old-format IDs with new format:

| Line | Old | New |
|------|-----|-----|
| ~154 | `2025.1.1` | `2025M.1.1` |
| ~156 | `2025.1.1` | `2025M.1.1` |
| ~1355 | `2024.1.1` | `2024M.1.1` |
| ~1357 | `2024.1.2` | `2024M.1.2` |
| ~1410 | `2024.1.1` | `2024M.1.1` |
| ~2480-2546 | `2024.1.1` | `2024M.1.1` |

Search: `grep -n '20[12][0-9]\.[0-9]' main_test.go` to find all instances.

**Step 2: Add migration v7 test**

Add a test `TestMigrationV7` following the pattern of existing migration tests (e.g., TestMigrationV5V6 at ~line 2564):

```go
func TestMigrationV7(t *testing.T) {
    // Create progress.db at v6
    // Insert matura_zrobione with old IDs
    // Insert probne_matury
    // Run migration
    // Verify IDs have M suffix
    // Verify probne_matury has sesja column
}
```

**Step 3: Run tests**

Run: `cd analiza/cli && go test ./... -v -run TestMigration`
Expected: all migration tests pass

Run: `cd analiza/cli && go test ./... -v`
Expected: all 13+ tests pass

---

### Task 14: Update test_qa.sh

**Files:**
- Modify: `analiza/test_qa.sh`

**Step 1: Check for hardcoded old-format IDs**

Run: `grep -n '20[12][0-9]\.[0-9]' analiza/test_qa.sh`

The test_qa.sh dynamically extracts IDs from CLI output, so it should work automatically with new format. But verify and fix any hardcoded references.

**Step 2: Verify exam list output parsing**

If test_qa.sh parses `exam list` output, ensure it handles the new `sesja` field.

---

### Task 15: Build, reimport, and run full test suite

**Step 1: Build binaries and reimport DB**

Run: `cd analiza/cli && ./build.sh`
Expected: macOS + Windows binaries built, matura.db reimported with new IDs

**Step 2: Verify data stats**

Run: `./matura data stats`
Expected: `{"cwiczenia": 937, "podzadania": 230, "cheatsheets": 4, "per_sesja": {"maj": 230}}`

**Step 3: Verify exam list**

Run: `./matura exam list`
Expected: all 11 years listed with `"sesja": "maj"`

**Step 4: Verify a specific exam task**

Run: `./matura exam task --rok 2025 --zadanie 1`
Expected: subtask IDs start with `2025M.`

**Step 5: Run full test suite**

Run: `cd analiza/cli && go test ./... -v`
Expected: all tests pass

Run: `cd analiza && ./test_qa.sh`
Expected: all 112 tests pass

**Step 6: Update baseline if needed**

If verify_all.py counts changed (unlikely — only CKE IDs changed, not exercises):

Run: `./test_qa.sh --update-baseline`

---

### Task 16: Update documentation

**Files:**
- Modify: `CLAUDE.md` — update ID format examples, file names
- Modify: `analiza/json/matura_indeks.json` — already done in Task 1, verify `opis` field text

**Step 1: Update CLAUDE.md**

Update any references to `matura_YYYY.json` → `matura_YYYYM.json` and ID format `YYYY.Z.S` → `YYYYM.Z.S`.

**Step 2: Verify no stale references remain**

Run: `grep -rn 'matura_20[12][0-9]\.json' analiza/ --include='*.md'`

Should only match design docs (which are historical) and README references that may need updating.

---

## Summary

| Task | Description | Key files |
|------|-------------|-----------|
| 1 | JSON migration (rename + re-ID + add sesja) | `analiza/json/matura_*.json` |
| 2 | Go types (add Sesja fields) | `types.go` |
| 3 | DB schema (egzamin + sesja column) | `database.go` |
| 4 | Importer (read + store sesja) | `importer.go` |
| 5 | exam list (group by rok+sesja) | `commands.go` |
| 6 | exam task (add --sesja flag) | `commands.go` |
| 7 | exam meta (add --sesja flag) | `commands.go` |
| 8 | exam save (add --sesja to probne_matury) | `commands.go` |
| 9 | cke get (add --sesja filter) | `commands.go` |
| 10 | data stats (per-session counts) | `commands.go`, `types.go` |
| 11 | cke save + help texts | `commands.go` |
| 12 | progress.db migration v7 | `database.go` |
| 13 | Unit tests (new IDs + migration test) | `main_test.go` |
| 14 | test_qa.sh updates | `test_qa.sh` |
| 15 | Build + full integration test | `build.sh` |
| 16 | Documentation | `CLAUDE.md` |

**Quality gates:**
- After Task 4: `go build ./...` succeeds
- After Task 13: `go test ./...` passes
- After Task 15: `test_qa.sh` passes (all 112 tests)
