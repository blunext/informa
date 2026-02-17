#!/bin/bash
# =============================================================================
# test_qa.sh — Regression test suite for matura tutoring system
#
# Usage:
#   ./test_qa.sh                    # Run all 5 layers
#   ./test_qa.sh --layer 1          # Run only layer 1 (CLI smoke)
#   ./test_qa.sh --layer 2          # Run only layer 2 (import round-trip)
#   ./test_qa.sh --layer 3          # Run only layer 3 (SKILL lint)
#   ./test_qa.sh --layer 4          # Run only layer 4 (baseline snapshot)
#   ./test_qa.sh --layer 5          # Run only layer 5 (Go unit tests)
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
    *) echo "Unknown layer: $RUN_LAYER"; exit 2 ;;
  esac
else
  run_layer_5
  run_layer_1
  run_layer_2
  run_layer_3
  run_layer_4
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
