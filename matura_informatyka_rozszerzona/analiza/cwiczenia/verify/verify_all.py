#!/usr/bin/env python3
"""Orchestrator for verifying all 230 exercises."""
import argparse
import json
import os
import sys
import glob as globmod
from datetime import datetime

# Add parent to path for imports
sys.path.insert(0, os.path.dirname(__file__))

from verifiers.base import VerificationResult
from verifiers.sql_verifier import verify_sql
from verifiers.cpp_verifier import verify_cpp
from verifiers.numconv_verifier import verify_numconv
from verifiers.manual_tagger import tag_manual

BASE_DIR = os.path.dirname(os.path.abspath(__file__))
JSON_DIR = os.path.join(os.path.dirname(BASE_DIR), 'json')
REPORT_DIR = os.path.join(BASE_DIR, 'report')
TMP_DIR = os.path.join(BASE_DIR, 'tmp')

# Mapping: file prefix -> verifier category
# 01-04: TEORIA (manual)
# 05: TEORIA numconv (auto)
# 06: TEORIA security (manual)
# 07-14: IMPLEMENTACJA (cpp)
# 15-19: ARKUSZ (manual)
# 20-23: SQL (auto)

SQL_RANGE = range(20, 24)          # 20, 21, 22, 23
CPP_RANGE = range(7, 15)           # 07, 08, 09, 10, 11, 12, 13, 14
NUMCONV_IDS = {5}                  # 05
MANUAL_TEORIA = {1, 2, 3, 4, 6}   # 01, 02, 03, 04, 06
MANUAL_ARKUSZ = range(15, 20)      # 15, 16, 17, 18, 19


def get_file_prefix(file_typ: str) -> int:
    """Extract numeric prefix from file_typ like '20_sql_group_by' -> 20."""
    m = file_typ.split('_')[0]
    return int(m)


def choose_verifier(file_typ: str):
    """Return the appropriate verifier function for a file type."""
    prefix = get_file_prefix(file_typ)
    if prefix in SQL_RANGE:
        return 'sql'
    elif prefix in CPP_RANGE:
        return 'cpp'
    elif prefix in NUMCONV_IDS:
        return 'numconv'
    else:
        return 'manual'


def run_verification(exercise: dict, file_typ: str, verifier_type: str) -> VerificationResult:
    """Run the appropriate verifier on an exercise."""
    if verifier_type == 'sql':
        return verify_sql(exercise, file_typ)
    elif verifier_type == 'cpp':
        return verify_cpp(exercise, file_typ, TMP_DIR)
    elif verifier_type == 'numconv':
        return verify_numconv(exercise, file_typ)
    else:
        return tag_manual(exercise, file_typ)


def load_all_exercises(json_dir: str) -> list[tuple[str, dict, dict]]:
    """Load all exercises from directory structure.

    Each exercise type is a directory with _meta.json + individual exercise files.
    Returns list of (file_typ, meta_data, exercise) tuples.
    """
    exercises = []
    # Find exercise directories (start with digit)
    all_dirs = sorted([
        d for d in os.listdir(json_dir)
        if os.path.isdir(os.path.join(json_dir, d)) and d[0].isdigit()
    ])
    for dirname in all_dirs:
        dir_path = os.path.join(json_dir, dirname)
        meta_path = os.path.join(dir_path, '_meta.json')
        if not os.path.exists(meta_path):
            continue
        with open(meta_path, 'r', encoding='utf-8') as f:
            meta = json.load(f)
        for ex_entry in meta.get('cwiczenia', []):
            eid = ex_entry['id']
            ex_path = os.path.join(dir_path, f'{eid}.json')
            if not os.path.exists(ex_path):
                continue
            with open(ex_path, 'r', encoding='utf-8') as f:
                ex = json.load(f)
            exercises.append((dirname, meta, ex))
    return exercises


def generate_report_md(results: list[VerificationResult]) -> str:
    """Generate markdown report."""
    lines = []
    lines.append(f"# Verification Report")
    lines.append(f"Generated: {datetime.now().strftime('%Y-%m-%d %H:%M:%S')}")
    lines.append(f"Total exercises: {len(results)}")
    lines.append("")

    # Summary counts
    counts = {}
    for r in results:
        counts[r.status] = counts.get(r.status, 0) + 1

    lines.append("## Summary")
    lines.append("")
    for status in ['PASS', 'FAIL', 'ERROR', 'MANUAL_REVIEW']:
        c = counts.get(status, 0)
        pct = f"{c/len(results)*100:.1f}%" if results else "0%"
        emoji = {'PASS': 'OK', 'FAIL': '!!', 'ERROR': '??', 'MANUAL_REVIEW': '--'}[status]
        lines.append(f"- [{emoji}] **{status}**: {c} ({pct})")
    lines.append("")

    # By category
    cat_counts = {}
    for r in results:
        key = (r.kategoria, r.status)
        cat_counts[key] = cat_counts.get(key, 0) + 1

    categories = sorted(set(r.kategoria for r in results))
    lines.append("## By Category")
    lines.append("")
    lines.append("| Kategoria | PASS | FAIL | ERROR | MANUAL |")
    lines.append("|-----------|------|------|-------|--------|")
    for cat in categories:
        p = cat_counts.get((cat, 'PASS'), 0)
        f = cat_counts.get((cat, 'FAIL'), 0)
        e = cat_counts.get((cat, 'ERROR'), 0)
        m = cat_counts.get((cat, 'MANUAL_REVIEW'), 0)
        lines.append(f"| {cat} | {p} | {f} | {e} | {m} |")
    lines.append("")

    # Failures
    failures = [r for r in results if r.status == 'FAIL']
    if failures:
        lines.append("## Failures")
        lines.append("")
        for r in failures:
            lines.append(f"### {r.exercise_id} ({r.file_typ})")
            lines.append(f"**Method**: {r.method}")
            lines.append(f"**Details**: {r.details}")
            if r.expected:
                lines.append(f"**Expected**:")
                lines.append(f"```\n{r.expected}\n```")
            if r.actual:
                lines.append(f"**Actual**:")
                lines.append(f"```\n{r.actual}\n```")
            lines.append("")

    # Errors
    errors = [r for r in results if r.status == 'ERROR']
    if errors:
        lines.append("## Errors")
        lines.append("")
        for r in errors:
            lines.append(f"### {r.exercise_id} ({r.file_typ})")
            lines.append(f"**Method**: {r.method}")
            lines.append(f"**Details**: {r.details}")
            if r.error:
                lines.append(f"**Error**: `{r.error[:200]}`")
            lines.append("")

    # Manual review (grouped by type)
    manual = [r for r in results if r.status == 'MANUAL_REVIEW']
    if manual:
        lines.append("## Manual Review Required")
        lines.append("")
        by_file = {}
        for r in manual:
            by_file.setdefault(r.file_typ, []).append(r)
        for ftyp in sorted(by_file.keys()):
            items = by_file[ftyp]
            lines.append(f"### {ftyp} ({len(items)} exercises)")
            lines.append(f"Hint: {items[0].details}")
            lines.append(f"IDs: {', '.join(r.exercise_id for r in items)}")
            lines.append("")

    return '\n'.join(lines)


def generate_report_json(results: list[VerificationResult]) -> dict:
    """Generate JSON report."""
    return {
        'generated': datetime.now().isoformat(),
        'total': len(results),
        'summary': {
            status: sum(1 for r in results if r.status == status)
            for status in ['PASS', 'FAIL', 'ERROR', 'MANUAL_REVIEW']
        },
        'results': [
            {
                'exercise_id': r.exercise_id,
                'file_typ': r.file_typ,
                'kategoria': r.kategoria,
                'status': r.status,
                'method': r.method,
                'details': r.details,
                'expected': r.expected,
                'actual': r.actual,
                'error': r.error,
            }
            for r in results
        ]
    }


def main():
    parser = argparse.ArgumentParser(description='Verify exercise correctness')
    parser.add_argument('--category', choices=['SQL', 'IMPLEMENTACJA', 'TEORIA', 'ARKUSZ'],
                        help='Only verify this category')
    parser.add_argument('--file', help='Only verify this file type (e.g., 20_sql_group_by)')
    parser.add_argument('--id', help='Only verify this exercise ID (e.g., 20.1)')
    parser.add_argument('--verbose', '-v', action='store_true', help='Print details for each exercise')
    args = parser.parse_args()

    # Load exercises
    all_exercises = load_all_exercises(JSON_DIR)
    print(f"Loaded {len(all_exercises)} exercises from {JSON_DIR}")

    # Filter
    exercises = all_exercises
    if args.file:
        exercises = [(ft, fd, ex) for ft, fd, ex in exercises if ft == args.file]
    if args.id:
        exercises = [(ft, fd, ex) for ft, fd, ex in exercises if ex['id'] == args.id]
    if args.category:
        cat_map = {
            'SQL': SQL_RANGE,
            'IMPLEMENTACJA': CPP_RANGE,
        }
        if args.category == 'SQL':
            exercises = [(ft, fd, ex) for ft, fd, ex in exercises if get_file_prefix(ft) in SQL_RANGE]
        elif args.category == 'IMPLEMENTACJA':
            exercises = [(ft, fd, ex) for ft, fd, ex in exercises if get_file_prefix(ft) in CPP_RANGE]
        elif args.category == 'TEORIA':
            exercises = [(ft, fd, ex) for ft, fd, ex in exercises if get_file_prefix(ft) in NUMCONV_IDS or get_file_prefix(ft) in MANUAL_TEORIA]
        elif args.category == 'ARKUSZ':
            exercises = [(ft, fd, ex) for ft, fd, ex in exercises if get_file_prefix(ft) in MANUAL_ARKUSZ]

    print(f"Verifying {len(exercises)} exercises...")
    print()

    # Ensure tmp dir exists
    os.makedirs(TMP_DIR, exist_ok=True)
    os.makedirs(REPORT_DIR, exist_ok=True)

    # Run verification
    results = []
    for file_typ, file_data, exercise in exercises:
        verifier_type = choose_verifier(file_typ)
        result = run_verification(exercise, file_typ, verifier_type)
        results.append(result)

        if args.verbose:
            status_mark = {
                'PASS': '[OK]', 'FAIL': '[!!]', 'ERROR': '[??]', 'MANUAL_REVIEW': '[--]'
            }[result.status]
            print(f"  {status_mark} {result.exercise_id:8s} ({file_typ}) — {result.details[:80]}")
        else:
            # Progress indicator
            mark = {'PASS': '.', 'FAIL': 'F', 'ERROR': 'E', 'MANUAL_REVIEW': '-'}[result.status]
            print(mark, end='', flush=True)

    if not args.verbose:
        print()  # newline after progress dots

    print()

    # Summary
    counts = {}
    for r in results:
        counts[r.status] = counts.get(r.status, 0) + 1

    print("=" * 50)
    print("SUMMARY")
    print("=" * 50)
    for status in ['PASS', 'FAIL', 'ERROR', 'MANUAL_REVIEW']:
        c = counts.get(status, 0)
        print(f"  {status:15s}: {c:3d}")
    print(f"  {'TOTAL':15s}: {len(results):3d}")
    print()

    # Show failures and errors briefly
    failures = [r for r in results if r.status == 'FAIL']
    errors = [r for r in results if r.status == 'ERROR']

    if failures:
        print(f"FAILURES ({len(failures)}):")
        for r in failures:
            print(f"  {r.exercise_id} ({r.file_typ}): {r.details[:100]}")
        print()

    if errors:
        print(f"ERRORS ({len(errors)}):")
        for r in errors:
            detail = r.error[:80] if r.error else r.details[:80]
            print(f"  {r.exercise_id} ({r.file_typ}): {detail}")
        print()

    # Generate reports
    report_md = generate_report_md(results)
    report_json = generate_report_json(results)

    md_path = os.path.join(REPORT_DIR, 'verification_report.md')
    json_path = os.path.join(REPORT_DIR, 'verification_report.json')

    with open(md_path, 'w', encoding='utf-8') as f:
        f.write(report_md)
    with open(json_path, 'w', encoding='utf-8') as f:
        json.dump(report_json, f, indent=2, ensure_ascii=False)

    print(f"Reports saved:")
    print(f"  {md_path}")
    print(f"  {json_path}")

    # Exit code: 0 if no failures, 1 if failures/errors
    return 1 if (failures or errors) else 0


if __name__ == '__main__':
    sys.exit(main())
