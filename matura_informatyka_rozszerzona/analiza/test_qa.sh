#!/bin/bash
# =============================================================================
# test_qa.sh — Regression test suite for matura tutoring system
#
# Usage:
#   ./test_qa.sh                    # Run all 6 layers
#   ./test_qa.sh --layer 1          # Run only layer 1 (CLI smoke)
#   ./test_qa.sh --layer 2          # Run only layer 2 (import round-trip)
#   ./test_qa.sh --layer 3          # Run only layer 3 (SKILL lint)
#   ./test_qa.sh --layer 4          # Run only layer 4 (baseline snapshot)
#   ./test_qa.sh --layer 5          # Run only layer 5 (Go unit tests)
#   ./test_qa.sh --layer 6          # Run only layer 6 (pedagogical journey)
#   ./test_qa.sh --update-baseline  # Update baseline (runs only layer 4)
# =============================================================================

# Disable set -e: test functions handle errors themselves; set -e would kill
# the script on expected non-zero exits (e.g., validate_json.py with errors).
set -uo pipefail

SCRIPT_DIR="$(cd "$(dirname "$0")" && pwd)"
CLI_DIR="$SCRIPT_DIR/cli"
MATURA="$CLI_DIR/matura"
SKILL_DIR="$(cd "$SCRIPT_DIR/../.." && pwd)/.claude/skills"
BASELINE_FILE="$SCRIPT_DIR/cwiczenia/verify/baseline.json"

# Colors
RED='\033[0;31m'
GREEN='\033[0;32m'
YELLOW='\033[0;33m'
BLUE='\033[0;34m'
BOLD='\033[1m'
NC='\033[0m'

PASS_COUNT=0
FAIL_COUNT=0
WARN_COUNT=0
FAILURES=()
CLEANUP_DIRS=()

# --- Helpers ----------------------------------------------------------------

pass() {
  PASS_COUNT=$((PASS_COUNT + 1))
  printf "  ${GREEN}✓${NC} %s\n" "$1"
}

fail() {
  FAIL_COUNT=$((FAIL_COUNT + 1))
  FAILURES+=("$1")
  printf "  ${RED}✗${NC} %s\n" "$1"
}

warn() {
  WARN_COUNT=$((WARN_COUNT + 1))
  printf "  ${YELLOW}⚠${NC} %s\n" "$1"
}

header() {
  echo ""
  printf "${BOLD}${BLUE}=== Layer %s: %s ===${NC}\n" "$1" "$2"
}

cleanup() {
  for d in "${CLEANUP_DIRS[@]}"; do
    rm -rf "$d" 2>/dev/null || true
  done
}
trap cleanup EXIT

# Test: command exits 0 and produces non-empty output
test_cmd() {
  local label="$1"
  shift
  local output
  if output=$("$@" 2>&1); then
    if [ -n "$output" ]; then
      pass "$label"
    else
      fail "$label (empty output)"
    fi
  else
    fail "$label (exit code $?)"
  fi
}

# Test: command exits 0 and output is valid JSON
test_json_cmd() {
  local label="$1"
  shift
  local output
  if output=$("$@" 2>&1); then
    if echo "$output" | python3 -c "import sys,json; json.load(sys.stdin)" 2>/dev/null; then
      pass "$label"
    else
      fail "$label (invalid JSON)"
    fi
  else
    fail "$label (exit code $?)"
  fi
}

# Test: command exits with expected non-zero code (e.g., "not found" = 1)
test_cmd_exitcode() {
  local label="$1"
  local expected_exit="$2"
  shift 2
  local actual_exit=0
  "$@" >/dev/null 2>&1 || actual_exit=$?
  if [ "$actual_exit" -eq "$expected_exit" ]; then
    pass "$label (expected exit $expected_exit)"
  else
    fail "$label (expected exit $expected_exit, got $actual_exit)"
  fi
}

# --- Parse args -------------------------------------------------------------

RUN_LAYER=""
UPDATE_BASELINE=false

while [[ $# -gt 0 ]]; do
  case "$1" in
    --layer) RUN_LAYER="$2"; shift 2 ;;
    --update-baseline) UPDATE_BASELINE=true; shift ;;
    -h|--help)
      echo "Usage: $0 [--layer N] [--update-baseline]"
      exit 0
      ;;
    *) echo "Unknown option: $1"; exit 2 ;;
  esac
done

# --- Preflight --------------------------------------------------------------

if [ ! -x "$MATURA" ]; then
  printf "${RED}ERROR: CLI binary not found at %s${NC}\n" "$MATURA"
  exit 2
fi

# Create temp dir for write tests (isolated progress.db)
TMPDIR_QA=$(mktemp -d)
CLEANUP_DIRS+=("$TMPDIR_QA")
cp "$CLI_DIR/matura.db" "$TMPDIR_QA/matura.db"

# Helper: run CLI with temp DB (for write tests that should not touch real DB)
matura_tmp() {
  "$MATURA" --db-dir "$TMPDIR_QA" "$@"
}

printf "${BOLD}Matura QA Test Suite${NC}\n"
printf "CLI: %s\n" "$MATURA"
printf "Temp DB: %s\n" "$TMPDIR_QA"

# =============================================================================
# Layer 1: CLI Smoke Test
# =============================================================================

run_layer_1() {
  header 1 "CLI Smoke Test"

  echo "  -- Read-only commands --"
  test_json_cmd "data stats" "$MATURA" data stats
  test_json_cmd "exercise get --typ cyfry_liczby" "$MATURA" exercise get --typ cyfry_liczby
  test_json_cmd "exercise next --typ napisy" matura_tmp exercise next --typ napisy
  # exercise review on fresh DB = "no reviews due" (exit 1) — expected
  test_cmd_exitcode "exercise review --limit 1 (fresh DB, expect 1)" 1 \
    matura_tmp exercise review --limit 1
  test_json_cmd "typ intro --typ cyfry_liczby" matura_tmp typ intro --typ cyfry_liczby
  test_json_cmd "progress status" matura_tmp progress status
  test_json_cmd "progress diagnose" matura_tmp progress diagnose
  test_json_cmd "cke get --typ sledzenie_algorytmu --force" "$MATURA" cke get --typ sledzenie_algorytmu --force
  test_json_cmd "cke status" matura_tmp cke status
  test_json_cmd "exam list" "$MATURA" exam list
  test_json_cmd "exam meta --rok 2024" "$MATURA" exam meta --rok 2024
  test_json_cmd "exam task --rok 2024 --zadanie 1" "$MATURA" exam task --rok 2024 --zadanie 1
  test_json_cmd "trap list --typ sledzenie_algorytmu" "$MATURA" trap list --typ sledzenie_algorytmu
  test_json_cmd "trap list --kategoria TEORIA" "$MATURA" trap list --kategoria TEORIA
  test_cmd "cheatsheet get --kategoria TEORIA" "$MATURA" cheatsheet get --kategoria TEORIA
  test_cmd "cheatsheet get --kategoria SQL --sekcja join" "$MATURA" cheatsheet get --kategoria SQL --sekcja "join"

  echo "  -- All 4 categories --"
  for kat in TEORIA IMPLEMENTACJA ARKUSZ SQL; do
    test_cmd "cheatsheet get --kategoria $kat" "$MATURA" cheatsheet get --kategoria "$kat"
  done

  echo "  -- All 23 exercise types --"
  for typ in sledzenie_algorytmu projektowanie_algorytmu analiza_algorytmu \
             test_prawda_falsz konwersja_systemow_liczbowych teoria_bezpieczenstwa \
             cyfry_liczby napisy zlozone zliczanie minmax sekwencje obrazy_2D geometryczne \
             agregacja_warunkowa symulacja wykres agregacja_podstawowa transformacja \
             sql_group_by sql_podzapytania sql_join sql_select_where; do
    test_json_cmd "exercise get --typ $typ" "$MATURA" exercise get --typ "$typ"
  done

  echo "  -- All 11 exam years --"
  for rok in 2014 2015 2016 2017 2018 2019 2021 2022 2023 2024 2025; do
    test_json_cmd "exam meta --rok $rok" "$MATURA" exam meta --rok "$rok"
  done

  echo "  -- Write commands (temp DB) --"
  # Get an exercise ID for write tests
  local ex_id
  ex_id=$("$MATURA" exercise get --typ cyfry_liczby 2>/dev/null \
    | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])" 2>/dev/null) || ex_id="7.1"

  test_json_cmd "progress update --wynik poprawne_bez_pomocy" \
    matura_tmp progress update --id "$ex_id" --wynik poprawne_bez_pomocy --czas 60

  # progress update with non-perfect result (no prior blad) → blad_warning present
  local ex_id2
  ex_id2=$("$MATURA" exercise get --typ napisy 2>/dev/null \
    | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])" 2>/dev/null) || ex_id2="8.1"
  local update_out2
  update_out2=$(matura_tmp progress update --id "$ex_id2" --wynik poprawne_z_pomoca_1 --czas 45 2>&1)
  local has_warning
  has_warning=$(echo "$update_out2" | python3 -c "import sys,json; d=json.loads(sys.stdin.read()); print('yes' if d.get('blad_warning') else 'no')" 2>/dev/null) || has_warning=""
  if [ "$has_warning" = "yes" ]; then
    pass "progress update non-perfect → blad_warning present"
  else
    fail "progress update non-perfect → blad_warning (expected yes, got: $has_warning)"
  fi

  test_json_cmd "progress blad" \
    matura_tmp progress blad --exercise-id "$ex_id" --typ cyfry_liczby --kod brak_inicjalizacji

  # cke save needs a valid CKE id
  local cke_id
  cke_id=$("$MATURA" cke get --typ sledzenie_algorytmu --force 2>/dev/null \
    | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])" 2>/dev/null) || cke_id=""
  if [ -n "$cke_id" ]; then
    test_json_cmd "cke save" matura_tmp cke save --id "$cke_id" --punkty 2 --max 3
  else
    warn "cke save (skipped — could not get CKE id)"
  fi

  # trap save needs a valid trap id — use sledzenie_algorytmu (has traps)
  local trap_id
  trap_id=$("$MATURA" trap list --typ sledzenie_algorytmu 2>/dev/null \
    | python3 -c "import sys,json; d=json.loads(sys.stdin.read()); print(d[0]['source_id'] if d else '')" 2>/dev/null) || trap_id=""
  if [ -n "$trap_id" ]; then
    test_json_cmd "trap save" matura_tmp trap save --id "$trap_id" --typ sledzenie_algorytmu --trafienia 2 --total 3
  else
    warn "trap save (skipped — no traps found)"
  fi

  # exam save — use 2024, minimal results
  local exam_task_id
  exam_task_id=$("$MATURA" exam task --rok 2024 --zadanie 1 2>/dev/null \
    | python3 -c "
import sys,json
d=json.loads(sys.stdin.read())
sub=d['podzadania'][0]
print(sub['id'])
" 2>/dev/null) || exam_task_id=""
  if [ -n "$exam_task_id" ]; then
    local max_pts
    max_pts=$("$MATURA" exam task --rok 2024 --zadanie 1 2>/dev/null \
      | python3 -c "
import sys,json
d=json.loads(sys.stdin.read())
print(d['podzadania'][0]['punkty'])
" 2>/dev/null) || max_pts="3"
    test_json_cmd "exam save" matura_tmp exam save --rok 2024 \
      --results "[{\"id\":\"$exam_task_id\",\"pkt\":2,\"max\":$max_pts}]" --czas 30
  else
    warn "exam save (skipped — could not get exam task id)"
  fi

  echo "  -- Context weight tracking --"
  test_json_cmd "exercise next --weight-reset" \
    matura_tmp exercise next --typ sql_group_by --weight-reset

  # --weight-add 30 → reset_suggested should be true
  local weight_out weight_reset
  weight_out=$(matura_tmp exercise next --typ sql_group_by --weight-add 30 2>&1)
  weight_reset=$(echo "$weight_out" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['reset_suggested'])" 2>/dev/null) || weight_reset=""
  if [ "$weight_reset" = "True" ]; then
    pass "exercise next --weight-add 30 → reset_suggested=true"
  else
    fail "exercise next --weight-add 30 → reset_suggested (expected True, got: $weight_reset)"
  fi

  echo "  -- Error handling --"
  test_cmd_exitcode "exercise get --typ NIEISTNIEJACY (expect error)" 1 \
    "$MATURA" exercise get --typ NIEISTNIEJACY
  test_cmd_exitcode "exam meta --rok 2020 (expect not found)" 1 \
    "$MATURA" exam meta --rok 2020
}

# =============================================================================
# Layer 2: Import Round-Trip
# =============================================================================

run_layer_2() {
  header 2 "Import Round-Trip"

  local import_dir
  import_dir=$(mktemp -d)
  CLEANUP_DIRS+=("$import_dir")

  echo "  -- Running data import --"
  local import_out
  if import_out=$("$MATURA" --db-dir "$import_dir" data import --source "$SCRIPT_DIR/" 2>&1); then
    pass "data import exits 0"
  else
    fail "data import (exit code $?): $import_out"
    return
  fi

  echo "  -- Verifying counts --"
  local stats
  if stats=$("$MATURA" --db-dir "$import_dir" data stats 2>&1); then
    :  # success
  else
    fail "data stats after import (exit code $?)"
    return
  fi

  local cwiczenia podzadania cheatsheets
  cwiczenia=$(echo "$stats" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['cwiczenia'])" 2>/dev/null) || cwiczenia=0
  podzadania=$(echo "$stats" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['podzadania'])" 2>/dev/null) || podzadania=0
  cheatsheets=$(echo "$stats" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['cheatsheets'])" 2>/dev/null) || cheatsheets=0

  [ "$cwiczenia" -ge 407 ] && pass "exercises: $cwiczenia (>= 407)" || fail "exercises: $cwiczenia (expected >= 407)"
  [ "$podzadania" -ge 230 ] && pass "subtasks: $podzadania (>= 230)" || fail "subtasks: $podzadania (expected >= 230)"
  [ "$cheatsheets" -ge 4 ] && pass "cheatsheets: $cheatsheets (>= 4)" || fail "cheatsheets: $cheatsheets (expected >= 4)"
}

# =============================================================================
# Layer 3: SKILL.md Lint
# =============================================================================

lint_skill_file() {
  local skill_file="$1"
  local skill_name="$2"

  if [ ! -f "$skill_file" ]; then
    fail "$skill_name: SKILL.md not found at $skill_file"
    return
  fi

  pass "$skill_name: SKILL.md exists"

  # Extract unique CLI command patterns from SKILL.md
  # Only match patterns that reference known CLI command groups
  local commands
  commands=$(python3 -c "
import re, sys

with open('$skill_file') as f:
    text = f.read()

# Known top-level CLI command groups
KNOWN_GROUPS = {'exercise', 'progress', 'cke', 'exam', 'trap', 'cheatsheet', 'data', 'typ'}

# Match: ./matura <word> <word> or \$MATURA <word> <word>
patterns = re.findall(r'(?:\./matura|\\\$MATURA)\s+(\w+)\s+(\w+)', text)

# Also match backtick-wrapped: \`matura <word> <word>\`
patterns += re.findall(r'\x60matura\s+(\w+)\s+(\w+)', text)

# Deduplicate, only keep known CLI groups
seen = set()
for cmd, sub in patterns:
    key = f'{cmd} {sub}'
    if key not in seen and cmd in KNOWN_GROUPS:
        seen.add(key)
        print(key)
" 2>/dev/null | sort -u)

  local cmd_count=0
  while IFS= read -r line; do
    [ -z "$line" ] && continue
    cmd_count=$((cmd_count + 1))
    local group sub
    group=$(echo "$line" | awk '{print $1}')
    sub=$(echo "$line" | awk '{print $2}')

    # Check if command exists (--help should work)
    if "$MATURA" "$group" "$sub" --help >/dev/null 2>&1; then
      pass "$skill_name: matura $group $sub"
    else
      fail "$skill_name: matura $group $sub (command not found in CLI)"
    fi
  done <<< "$commands"

  if [ "$cmd_count" -eq 0 ]; then
    warn "$skill_name: no CLI commands extracted"
  else
    pass "$skill_name: checked $cmd_count unique commands"
  fi
}

run_layer_3() {
  header 3 "SKILL.md Lint"

  # Lint all skill files that reference the CLI
  for skill_dir in "$SKILL_DIR"/*/; do
    local skill_name
    skill_name=$(basename "$skill_dir")
    local skill_file="$skill_dir/SKILL.md"
    if [ -f "$skill_file" ] && grep -q "matura" "$skill_file" 2>/dev/null; then
      echo "  -- Skill: $skill_name --"
      lint_skill_file "$skill_file" "$skill_name"
    fi
  done

  # Check category names in main skill
  local main_skill="$SKILL_DIR/matura/SKILL.md"
  if [ -f "$main_skill" ]; then
    echo "  -- Checking category references --"
    for kat in TEORIA IMPLEMENTACJA ARKUSZ SQL; do
      if grep -q "$kat" "$main_skill" 2>/dev/null; then
        pass "Category $kat referenced in SKILL.md"
      else
        fail "Category $kat NOT found in SKILL.md"
      fi
    done
  fi
}

# =============================================================================
# Layer 4: Baseline Snapshot
# =============================================================================

run_layer_4() {
  header 4 "Baseline Snapshot"

  echo "  -- Running schema validation --"
  local validate_out validate_exit
  validate_out=$(python3 "$SCRIPT_DIR/cwiczenia/validate_json.py" 2>&1) && validate_exit=0 || validate_exit=$?

  if [ $validate_exit -eq 0 ]; then
    pass "validate_json.py exits 0"
  else
    fail "validate_json.py exits $validate_exit"
  fi

  # Extract counts from "Validated: 23 directories, 407 exercises" line
  local val_exercises val_dirs
  val_exercises=$(echo "$validate_out" | grep "Validated:" | grep -oE '[0-9]+ exercises' | grep -oE '[0-9]+') || val_exercises=0
  val_dirs=$(echo "$validate_out" | grep "Validated:" | grep -oE '[0-9]+ directories' | grep -oE '[0-9]+') || val_dirs=0
  echo "    Schema: $val_dirs dirs, $val_exercises exercises"

  echo "  -- Running content verification --"
  local verify_out verify_exit
  verify_out=$(python3 "$SCRIPT_DIR/cwiczenia/verify/verify_all.py" 2>&1) && verify_exit=0 || verify_exit=$?

  # Parse results from SUMMARY block
  local v_pass v_fail v_error v_manual
  v_pass=$(echo "$verify_out" | grep -E '^\s*PASS' | grep -oE '[0-9]+' | head -1) || v_pass=0
  v_fail=$(echo "$verify_out" | grep -E '^\s*FAIL\b' | grep -oE '[0-9]+' | head -1) || v_fail=0
  v_error=$(echo "$verify_out" | grep -E '^\s*ERROR' | grep -oE '[0-9]+' | head -1) || v_error=0
  v_manual=$(echo "$verify_out" | grep -E '^\s*MANUAL_REVIEW' | grep -oE '[0-9]+' | head -1) || v_manual=0

  echo "    Verify: PASS=$v_pass FAIL=$v_fail ERROR=$v_error MANUAL=$v_manual"

  [ "$v_fail" -eq 0 ] && pass "verify_all.py: 0 FAIL" || fail "verify_all.py: $v_fail FAIL"
  [ "$v_error" -eq 0 ] && pass "verify_all.py: 0 ERROR" || fail "verify_all.py: $v_error ERROR"

  # Current results as JSON
  local current_json
  current_json=$(python3 -c "
import json
print(json.dumps({
    'validate_exercises': $val_exercises,
    'validate_dirs': $val_dirs,
    'verify_pass': $v_pass,
    'verify_fail': $v_fail,
    'verify_error': $v_error,
    'verify_manual': $v_manual
}, indent=2))
" 2>/dev/null)

  if $UPDATE_BASELINE; then
    echo "$current_json" > "$BASELINE_FILE"
    pass "Baseline saved to $BASELINE_FILE"
    return
  fi

  # Compare with baseline
  if [ ! -f "$BASELINE_FILE" ]; then
    warn "No baseline found. Run with --update-baseline to create one."
    echo "$current_json" > "$BASELINE_FILE"
    pass "Initial baseline created at $BASELINE_FILE"
    return
  fi

  echo "  -- Comparing with baseline --"
  local cmp_exit
  python3 -c "
import json, sys

with open('$BASELINE_FILE') as f:
    baseline = json.load(f)

current = json.loads('''$current_json''')

ok = True
for key in sorted(baseline):
    b = baseline[key]
    c = current.get(key, 0)
    if key in ('verify_fail', 'verify_error'):
        # These should not increase
        if c > b:
            print(f'  \033[0;31m✗\033[0m {key}: {b} -> {c} (REGRESSION!)')
            ok = False
        elif c < b:
            print(f'  \033[0;32m✓\033[0m {key}: {b} -> {c} (improved)')
        else:
            print(f'  \033[0;32m✓\033[0m {key}: {c} (unchanged)')
    else:
        # These should not decrease
        if c < b:
            print(f'  \033[0;31m✗\033[0m {key}: {b} -> {c} (REGRESSION!)')
            ok = False
        elif c > b:
            print(f'  \033[0;32m✓\033[0m {key}: {b} -> {c} (improved)')
        else:
            print(f'  \033[0;32m✓\033[0m {key}: {c} (unchanged)')

sys.exit(0 if ok else 1)
" 2>/dev/null && cmp_exit=0 || cmp_exit=$?
  if [ $cmp_exit -eq 0 ]; then
    pass "No regressions vs baseline"
  else
    fail "Regressions detected vs baseline!"
  fi
}

# =============================================================================
# Layer 5: Go Unit Tests
# =============================================================================

run_layer_5() {
  header 5 "Go Unit Tests"
  local test_out test_exit
  test_out=$(cd "$CLI_DIR" && go test -v -count=1 ./... 2>&1) && test_exit=0 || test_exit=$?
  if [ $test_exit -eq 0 ]; then
    local test_count
    test_count=$(echo "$test_out" | grep -c -- "--- PASS:" 2>/dev/null) || test_count="?"
    pass "go test: $test_count tests pass"
  else
    fail "go test: failures detected (exit $test_exit)"
    echo "$test_out" | grep -E "FAIL|--- FAIL" | head -20
  fi
}

# =============================================================================
# Layer 6: Pedagogical Journey (deterministic CLI flow tests)
# =============================================================================

run_layer_6() {
  header 6 "Pedagogical Journey"

  # Each test uses an isolated temp DB to simulate student journeys

  # Shared fresh DB for read-only L6 tests (separate from global matura_tmp)
  local l6_fresh
  l6_fresh=$(mktemp -d)
  CLEANUP_DIRS+=("$l6_fresh")
  cp "$CLI_DIR/matura.db" "$l6_fresh/matura.db"

  # --- 6a: Fresh DB has cwiczenia_lacznie == 0 ---
  echo "  -- 6a: Fresh DB progress status --"
  local status_out
  status_out=$("$MATURA" --db-dir "$l6_fresh" progress status 2>&1)
  local cwicz_total
  cwicz_total=$(echo "$status_out" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['cwiczenia_lacznie'])" 2>/dev/null) || cwicz_total=""
  if [ "$cwicz_total" = "0" ]; then
    pass "6a: fresh DB cwiczenia_lacznie == 0"
  else
    fail "6a: fresh DB cwiczenia_lacznie == 0 (got: $cwicz_total)"
  fi

  # --- 6b: typ intro returns first_in_type == true on fresh DB ---
  echo "  -- 6b: typ intro first_in_type --"
  local intro_out first_in
  intro_out=$("$MATURA" --db-dir "$l6_fresh" typ intro --typ cyfry_liczby 2>&1)
  first_in=$(echo "$intro_out" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['first_in_type'])" 2>/dev/null) || first_in=""
  if [ "$first_in" = "True" ]; then
    pass "6b: typ intro --typ cyfry_liczby → first_in_type=true"
  else
    fail "6b: typ intro first_in_type (got: $first_in)"
  fi

  # --- 6c: 3x poprawne_bez_pomocy → poziom srednie ---
  echo "  -- 6c: Streak 3 → poziom srednie --"
  # Create a second temp dir for this multi-step journey
  local journey_dir
  journey_dir=$(mktemp -d)
  CLEANUP_DIRS+=("$journey_dir")
  cp "$CLI_DIR/matura.db" "$journey_dir/matura.db"

  # Get 3 distinct exercises of the same type
  local ex_ids=()
  local ex_json
  for i in 1 2 3; do
    local exclude_arg=""
    if [ ${#ex_ids[@]} -gt 0 ]; then
      exclude_arg=$(IFS=,; echo "${ex_ids[*]}")
    fi
    ex_json=$("$MATURA" --db-dir "$journey_dir" exercise get --typ cyfry_liczby --trudnosc latwe ${exclude_arg:+--exclude "$exclude_arg"} 2>/dev/null)
    local eid
    eid=$(echo "$ex_json" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])" 2>/dev/null) || eid=""
    if [ -n "$eid" ]; then
      ex_ids+=("$eid")
      "$MATURA" --db-dir "$journey_dir" progress update --id "$eid" --wynik poprawne_bez_pomocy --czas 60 >/dev/null 2>&1
    fi
  done

  local new_level
  new_level=$("$MATURA" --db-dir "$journey_dir" progress status --typ cyfry_liczby 2>/dev/null \
    | python3 -c "import sys,json; d=json.loads(sys.stdin.read()); print(d['per_typ'][0]['poziom_trudnosci'])" 2>/dev/null) || new_level=""
  if [ "$new_level" = "srednie" ]; then
    pass "6c: 3x poprawne_bez_pomocy → poziom srednie"
  else
    fail "6c: 3x poprawne_bez_pomocy → poziom srednie (got: $new_level)"
  fi

  # --- 6d: 3x progress blad with same code → diagnose returns that code ---
  echo "  -- 6d: Error diagnosis --"
  local diag_dir
  diag_dir=$(mktemp -d)
  CLEANUP_DIRS+=("$diag_dir")
  cp "$CLI_DIR/matura.db" "$diag_dir/matura.db"

  local blad_eid
  blad_eid=$("$MATURA" exercise get --typ cyfry_liczby 2>/dev/null \
    | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])" 2>/dev/null) || blad_eid="7.1"
  for i in 1 2 3; do
    "$MATURA" --db-dir "$diag_dir" progress blad --exercise-id "$blad_eid" --typ cyfry_liczby --kod mylenie_div_mod >/dev/null 2>&1
  done

  local top_blad
  top_blad=$("$MATURA" --db-dir "$diag_dir" progress diagnose --typ cyfry_liczby --limit 1 2>/dev/null \
    | python3 -c "import sys,json; d=json.loads(sys.stdin.read()); print(d['top_bledy'][0]['blad_kod'] if d['top_bledy'] else '')" 2>/dev/null) || top_blad=""
  if [ "$top_blad" = "mylenie_div_mod" ]; then
    pass "6d: 3x blad mylenie_div_mod → diagnose returns it"
  else
    fail "6d: diagnose top blad (expected mylenie_div_mod, got: $top_blad)"
  fi

  # --- 6e: cke get without --force at latwe level → exit 1 ---
  echo "  -- 6e: CKE unlock gate --"
  local cke_exit=0
  "$MATURA" --db-dir "$l6_fresh" cke get --typ sledzenie_algorytmu >/dev/null 2>&1 || cke_exit=$?
  if [ "$cke_exit" -eq 1 ]; then
    pass "6e: cke get without --force at latwe → exit 1"
  else
    fail "6e: cke get without --force (expected exit 1, got $cke_exit)"
  fi

  # --- 6f: cke get --force → exit 0, valid JSON ---
  echo "  -- 6f: CKE force bypass --"
  test_json_cmd "6f: cke get --force → valid JSON" \
    "$MATURA" --db-dir "$l6_fresh" cke get --typ sledzenie_algorytmu --force

  # --- 6g: exercise next on fresh DB → mode == new ---
  echo "  -- 6g: exercise next fresh → mode new --"
  local next_out next_mode
  next_out=$("$MATURA" --db-dir "$l6_fresh" exercise next --typ cyfry_liczby 2>&1)
  next_mode=$(echo "$next_out" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['mode'])" 2>/dev/null) || next_mode=""
  if [ "$next_mode" = "new" ]; then
    pass "6g: exercise next fresh DB → mode=new"
  else
    fail "6g: exercise next mode (expected new, got: $next_mode)"
  fi

  # --- 6h: exam meta + exam task → correct structure ---
  echo "  -- 6h: Exam meta + task structure --"
  local meta_out meta_ok=true
  meta_out=$("$MATURA" exam meta --rok 2024 2>&1)
  local meta_rok meta_formula meta_zadania
  meta_rok=$(echo "$meta_out" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['rok'])" 2>/dev/null) || meta_rok=""
  meta_formula=$(echo "$meta_out" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['formula'])" 2>/dev/null) || meta_formula=""
  meta_zadania=$(echo "$meta_out" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['liczba_zadan'])" 2>/dev/null) || meta_zadania=""

  [ "$meta_rok" = "2024" ] || meta_ok=false
  [ "$meta_formula" = "2023" ] || meta_ok=false
  [ -n "$meta_zadania" ] && [ "$meta_zadania" -gt 0 ] 2>/dev/null || meta_ok=false

  local task_out task_podzadania
  task_out=$("$MATURA" exam task --rok 2024 --zadanie 1 2>&1)
  task_podzadania=$(echo "$task_out" | python3 -c "import sys,json; print(len(json.loads(sys.stdin.read())['podzadania']))" 2>/dev/null) || task_podzadania=""
  [ -n "$task_podzadania" ] && [ "$task_podzadania" -gt 0 ] 2>/dev/null || meta_ok=false

  if $meta_ok; then
    pass "6h: exam meta 2024 + exam task 2024.1 → valid structure"
  else
    fail "6h: exam meta/task structure (rok=$meta_rok formula=$meta_formula zadania=$meta_zadania podzadania=$task_podzadania)"
  fi

  # --- 6i: exam save → succeeds ---
  echo "  -- 6i: exam save --"
  local save_dir
  save_dir=$(mktemp -d)
  CLEANUP_DIRS+=("$save_dir")
  cp "$CLI_DIR/matura.db" "$save_dir/matura.db"

  local save_task_id save_max
  save_task_id=$("$MATURA" exam task --rok 2024 --zadanie 1 2>/dev/null \
    | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['podzadania'][0]['id'])" 2>/dev/null) || save_task_id=""
  save_max=$("$MATURA" exam task --rok 2024 --zadanie 1 2>/dev/null \
    | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['podzadania'][0]['punkty'])" 2>/dev/null) || save_max="3"

  if [ -n "$save_task_id" ]; then
    local save_out save_ok
    save_out=$("$MATURA" --db-dir "$save_dir" exam save --rok 2024 \
      --results "[{\"id\":\"$save_task_id\",\"pkt\":2,\"max\":$save_max}]" --czas 30 2>&1)
    save_ok=$(echo "$save_out" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['saved'])" 2>/dev/null) || save_ok=""
    if [ "$save_ok" = "True" ]; then
      pass "6i: exam save → saved=true"
    else
      fail "6i: exam save (saved=$save_ok)"
    fi
  else
    warn "6i: exam save (skipped — could not get task id)"
  fi

  # --- 6j: exercise next after 2 exercises → session tracking works ---
  echo "  -- 6j: exercise next session tracking --"
  local sess_dir
  sess_dir=$(mktemp -d)
  CLEANUP_DIRS+=("$sess_dir")
  cp "$CLI_DIR/matura.db" "$sess_dir/matura.db"

  # Do 2 exercises to build session_count
  for i in 1 2; do
    local sess_eid
    sess_eid=$("$MATURA" --db-dir "$sess_dir" exercise get --typ napisy --trudnosc latwe 2>/dev/null \
      | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])" 2>/dev/null) || sess_eid=""
    if [ -n "$sess_eid" ]; then
      "$MATURA" --db-dir "$sess_dir" progress update --id "$sess_eid" --wynik poprawne_bez_pomocy --czas 45 >/dev/null 2>&1
    fi
  done

  local sess_next sess_count
  sess_next=$("$MATURA" --db-dir "$sess_dir" exercise next --typ napisy 2>&1)
  sess_count=$(echo "$sess_next" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['session_count'])" 2>/dev/null) || sess_count=""
  if [ -n "$sess_count" ] && [ "$sess_count" -ge 2 ] 2>/dev/null; then
    pass "6j: exercise next after 2 exercises → session_count=$sess_count"
  else
    fail "6j: exercise next session_count (expected >= 2, got: $sess_count)"
  fi

  # --- 6k: progress update with non-perfect result and no blad → blad_warning present ---
  echo "  -- 6k: blad_warning on non-perfect result --"
  local warn_dir
  warn_dir=$(mktemp -d)
  CLEANUP_DIRS+=("$warn_dir")
  cp "$CLI_DIR/matura.db" "$warn_dir/matura.db"

  local warn_eid
  warn_eid=$("$MATURA" --db-dir "$warn_dir" exercise get --typ cyfry_liczby 2>/dev/null \
    | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])" 2>/dev/null) || warn_eid="7.1"

  local warn_out warn_field
  warn_out=$("$MATURA" --db-dir "$warn_dir" progress update --id "$warn_eid" --wynik poprawne_z_pomoca_1 --czas 60 2>&1)
  warn_field=$(echo "$warn_out" | python3 -c "import sys,json; d=json.loads(sys.stdin.read()); print('yes' if d.get('blad_warning') else 'no')" 2>/dev/null) || warn_field=""
  if [ "$warn_field" = "yes" ]; then
    pass "6k: non-perfect result without blad → blad_warning present"
  else
    fail "6k: blad_warning (expected yes, got: $warn_field)"
  fi

  # --- 6l: progress update with blad logged first → no blad_warning ---
  echo "  -- 6l: no blad_warning when blad logged --"
  local nowarn_dir
  nowarn_dir=$(mktemp -d)
  CLEANUP_DIRS+=("$nowarn_dir")
  cp "$CLI_DIR/matura.db" "$nowarn_dir/matura.db"

  local nowarn_eid
  nowarn_eid=$("$MATURA" --db-dir "$nowarn_dir" exercise get --typ napisy 2>/dev/null \
    | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])" 2>/dev/null) || nowarn_eid="8.1"

  # Log blad first
  "$MATURA" --db-dir "$nowarn_dir" progress blad --exercise-id "$nowarn_eid" --typ napisy --kod off_by_one --hint 1 >/dev/null 2>&1

  # Then update with non-perfect result
  local nowarn_out nowarn_field
  nowarn_out=$("$MATURA" --db-dir "$nowarn_dir" progress update --id "$nowarn_eid" --wynik poprawne_z_pomoca_1 --czas 60 2>&1)
  nowarn_field=$(echo "$nowarn_out" | python3 -c "import sys,json; d=json.loads(sys.stdin.read()); print('yes' if d.get('blad_warning') else 'no')" 2>/dev/null) || nowarn_field=""
  if [ "$nowarn_field" = "no" ]; then
    pass "6l: blad logged before update → no blad_warning"
  else
    fail "6l: no blad_warning (expected no, got: $nowarn_field)"
  fi
}

# =============================================================================
# Main
# =============================================================================

if $UPDATE_BASELINE; then
  # --update-baseline only needs layer 4
  run_layer_4
elif [ -n "$RUN_LAYER" ]; then
  case "$RUN_LAYER" in
    1) run_layer_1 ;;
    2) run_layer_2 ;;
    3) run_layer_3 ;;
    4) run_layer_4 ;;
    5) run_layer_5 ;;
    6) run_layer_6 ;;
    *) echo "Unknown layer: $RUN_LAYER"; exit 2 ;;
  esac
else
  run_layer_5
  run_layer_1
  run_layer_2
  run_layer_3
  run_layer_4
  run_layer_6
fi

# --- Summary ----------------------------------------------------------------

echo ""
printf "${BOLD}=== SUMMARY ===${NC}\n"
printf "  ${GREEN}PASS${NC}: %d\n" "$PASS_COUNT"
[ "$WARN_COUNT" -gt 0 ] && printf "  ${YELLOW}WARN${NC}: %d\n" "$WARN_COUNT"
if [ "$FAIL_COUNT" -gt 0 ]; then
  printf "  ${RED}FAIL${NC}: %d\n" "$FAIL_COUNT"
  echo ""
  printf "${RED}Failed tests:${NC}\n"
  for f in "${FAILURES[@]}"; do
    printf "  ${RED}✗${NC} %s\n" "$f"
  done
  exit 1
else
  printf "\n${GREEN}${BOLD}All tests passed!${NC}\n"
  exit 0
fi
