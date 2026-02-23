# Layer 7: Exercise Consistency Check — Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Add `matura data verify` CLI command + Layer 7 in test_qa.sh that checks JSON↔DB sync, ID continuity, and difficulty distribution.

**Architecture:** New Go function `VerifyExercises()` reads JSON files + DB, compares per-exercise. Layer 7 in bash calls it + does ID/difficulty checks via `python3 -c` on `_meta.json`.

**Tech Stack:** Go (CLI command), bash (test layer), python3 one-liners (JSON parsing in bash)

---

### Task 1: Add `VerifyExercises` function in Go

**Files:**
- Modify: `analiza/cli/importer.go` (add VerifyExercises after ImportAll)
- Modify: `analiza/cli/types.go` (add VerifyOut type)

**Step 1: Add output type to `types.go`**

After `DataStatsOut` (line ~280), add:

```go
// VerifyOut is what data verify returns
type VerifyOut struct {
	TotalDisk     int      `json:"total_disk"`
	TotalDB       int      `json:"total_db"`
	Matched       int      `json:"matched"`
	Mismatched    []string `json:"mismatched"`
	MissingInDB   []string `json:"missing_in_db"`
	MissingOnDisk []string `json:"missing_on_disk"`
}
```

**Step 2: Add `VerifyExercises` function to `importer.go`**

After `ImportExercises` function (~line 138), add:

```go
// VerifyExercises compares JSON files on disk with exercises in matura.db.
// Returns a report of mismatches. sourceDir is the analiza/ directory.
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
```

**Step 3: Verify it compiles**

Run: `cd /Users/blt1wz/priv/informa/matura_informatyka_rozszerzona/analiza/cli && go build -o matura .`
Expected: exits 0, no errors

---

### Task 2: Add `data verify` command

**Files:**
- Modify: `analiza/cli/commands.go` (add dataVerifyCmd after dataStatsCmd ~line 2695)
- Modify: `analiza/cli/main.go:91` (register command)

**Step 1: Add command function to `commands.go`**

After `dataStatsCmd()` (line ~2695), add:

```go
// === data verify ===

func dataVerifyCmd() *cobra.Command {
	var source string

	cmd := &cobra.Command{
		Use:   "verify",
		Short: "Verify JSON files match matura.db contents",
		RunE: func(cmd *cobra.Command, args []string) error {
			dataDB, err := OpenDataDB(dbDir)
			if err != nil {
				return fatal(fmt.Sprintf("open data DB: %v", err))
			}
			defer dataDB.Close()

			result, err := VerifyExercises(dataDB, source)
			if err != nil {
				return fatal(fmt.Sprintf("verify failed: %v", err))
			}

			jsonOut(result)

			if len(result.Mismatched) > 0 || len(result.MissingInDB) > 0 || len(result.MissingOnDisk) > 0 {
				return fatal(fmt.Sprintf("%d mismatched, %d missing in DB, %d missing on disk",
					len(result.Mismatched), len(result.MissingInDB), len(result.MissingOnDisk)))
			}
			return nil
		},
	}

	cmd.Flags().StringVar(&source, "source", "", "Path to analiza/ directory")
	cmd.MarkFlagRequired("source")
	return cmd
}
```

**Step 2: Register in main.go**

Change line 91 in `main.go`:
```go
dataCmd.AddCommand(dataImportCmd(), dataStatsCmd())
```
to:
```go
dataCmd.AddCommand(dataImportCmd(), dataStatsCmd(), dataVerifyCmd())
```

**Step 3: Build and test**

Run: `cd /Users/blt1wz/priv/informa/matura_informatyka_rozszerzona/analiza/cli && go build -o matura . && ./matura data verify --source ../`
Expected: JSON output with `matched: 407`, `mismatched: []`, `missing_in_db: []`, `missing_on_disk: []`, exit 0

---

### Task 3: Add Go unit test for data verify

**Files:**
- Modify: `analiza/cli/main_test.go` (add TestDataVerify)

**Step 1: Add test**

After `TestDataImportCreatesDBFile` (~line 697), add:

```go
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

	if result.TotalDisk != 407 {
		t.Errorf("total_disk: got %d, want 407", result.TotalDisk)
	}
	if result.TotalDB != 407 {
		t.Errorf("total_db: got %d, want 407", result.TotalDB)
	}
	if result.Matched != 407 {
		t.Errorf("matched: got %d, want 407", result.Matched)
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
```

**Step 2: Run test**

Run: `cd /Users/blt1wz/priv/informa/matura_informatyka_rozszerzona/analiza/cli && go test -run TestDataVerify -v`
Expected: PASS

---

### Task 4: Add Layer 7 to test_qa.sh

**Files:**
- Modify: `analiza/test_qa.sh` (add run_layer_7 function + dispatch)

**Step 1: Add `run_layer_7` function**

Before the dispatch section (line ~880), add:

```bash
# =============================================================================
# Layer 7: Exercise Consistency
# =============================================================================

run_layer_7() {
  header 7 "Exercise Consistency"

  # --- 7a: Data verify (JSON ↔ DB sync) ---
  echo "  -- 7a: JSON ↔ DB sync --"
  verify_out=$("$MATURA" data verify --source "$SCRIPT_DIR" 2>/dev/null)
  if [ $? -eq 0 ]; then
    matched=$(echo "$verify_out" | python3 -c "import sys,json; print(json.load(sys.stdin)['matched'])")
    pass "data verify: $matched exercises matched"
  else
    fail "data verify: mismatch detected"
    echo "$verify_out" | python3 -c "
import sys, json
d = json.load(sys.stdin)
for m in d.get('mismatched', []): print(f'    MISMATCH: {m}')
for m in d.get('missing_in_db', []): print(f'    MISSING IN DB: {m}')
for m in d.get('missing_on_disk', []): print(f'    MISSING ON DISK: {m}')
" 2>/dev/null
  fi

  # --- 7b: ID continuity ---
  echo "  -- 7b: ID continuity --"
  local json_dir="$SCRIPT_DIR/cwiczenia/json"
  local id_ok=true
  for meta in "$json_dir"/[0-9][0-9]_*/_meta.json; do
    local dir_name=$(basename "$(dirname "$meta")")
    local result
    result=$(python3 -c "
import json, sys
with open('$meta') as f:
    data = json.load(f)
typ = data['typ']
prefix = int(typ.split('_')[0])
ids = sorted([int(e['id'].split('.')[1]) for e in data['cwiczenia']])
expected = list(range(1, len(ids) + 1))
if ids != expected:
    gaps = [i for i in expected if i not in ids]
    dups = [i for i in ids if ids.count(i) > 1]
    print(f'FAIL:{typ}:gaps={gaps},dups={list(set(dups))}')
else:
    print(f'OK:{typ}:{len(ids)}')
")
    if [[ "$result" == OK:* ]]; then
      local typ=$(echo "$result" | cut -d: -f2)
      local count=$(echo "$result" | cut -d: -f3)
      pass "$dir_name: IDs 1-$count sequential"
    else
      local detail=$(echo "$result" | cut -d: -f2-)
      fail "$dir_name: $detail"
      id_ok=false
    fi
  done

  # --- 7c: Difficulty distribution ---
  echo "  -- 7c: Difficulty distribution --"
  for meta in "$json_dir"/[0-9][0-9]_*/_meta.json; do
    local dir_name=$(basename "$(dirname "$meta")")
    local result
    result=$(python3 -c "
import json
from collections import Counter
with open('$meta') as f:
    data = json.load(f)
dist = Counter(e['trudnosc'] for e in data['cwiczenia'])
levels = len(dist)
parts = ' '.join(f'{k}={v}' for k, v in sorted(dist.items()))
if levels < 2:
    print(f'WARN:{parts} (only {levels} level)')
else:
    print(f'OK:{parts}')
")
    if [[ "$result" == OK:* ]]; then
      pass "$dir_name: ${result#OK:}"
    else
      warn "$dir_name: ${result#WARN:}"
    fi
  done
}
```

**Step 2: Add dispatch**

In the dispatch section, add `7)` to the case statement (after line 892):
```bash
    7) run_layer_7 ;;
```

Add `run_layer_7` to the default run (after line 901, before `fi`):
```bash
  run_layer_7
```

**Step 3: Update header comment**

Change line 2 usage comments to mention layer 7:
```bash
#   ./test_qa.sh --layer 7          # Run only layer 7 (exercise consistency)
```

**Step 4: Test Layer 7 alone**

Run: `cd /Users/blt1wz/priv/informa/matura_informatyka_rozszerzona/analiza && ./test_qa.sh --layer 7`
Expected: All PASS (407 matched, 23 dirs sequential IDs, difficulty distribution)

---

### Task 5: Full regression test + rebuild

**Step 1: Rebuild binary + re-import**

Run: `cd /Users/blt1wz/priv/informa/matura_informatyka_rozszerzona/analiza/cli && ./build.sh`
Expected: "Build OK" message

**Step 2: Run Go tests**

Run: `cd /Users/blt1wz/priv/informa/matura_informatyka_rozszerzona/analiza/cli && go test -v -count=1 ./...`
Expected: All PASS (107+ tests including new TestDataVerify)

**Step 3: Run full test suite**

Run: `cd /Users/blt1wz/priv/informa/matura_informatyka_rozszerzona/analiza && ./test_qa.sh`
Expected: All 7 layers PASS, total 150+ tests, 0 FAIL

---

### Task 6: Update CLAUDE.md and MEMORY.md

**Files:**
- Modify: `CLAUDE.md` — update test_qa.sh description (add Layer 7)
- No need to update MEMORY.md (auto-updated)

**Step 1: Update CLAUDE.md**

In the `test_qa.sh` section, add Layer 7 reference. In the expected results, update test count.

**Step 2: Update /generate-exercises skill**

Add `data verify` to Krok 5b in `.claude/skills/generate-exercises/SKILL.md`:

After the re-import command, add:
```bash
# Verify JSON ↔ DB sync:
./matura data verify --source ../
```
