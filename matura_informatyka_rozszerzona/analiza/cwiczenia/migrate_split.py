#!/usr/bin/env python3
"""Migrate exercise JSON files from flat structure to directory-per-type structure.

Before:
  json/01_sledzenie_algorytmu.json  (35KB — 10 exercises in one file)

After:
  json/01_sledzenie_algorytmu/
    _meta.json    (~2KB — metadata + exercise index)
    1.1.json      (~3-5KB — single exercise)
    1.2.json
    ...

Usage:
    python3 migrate_split.py           # migrate all
    python3 migrate_split.py --dry-run # preview without changes
"""
import json
import os
import sys
import glob
import argparse
import shutil

BASE_DIR = os.path.dirname(os.path.abspath(__file__))
JSON_DIR = os.path.join(BASE_DIR, "json")


def migrate_file(json_path: str, dry_run: bool = False) -> tuple[int, list[str]]:
    """Migrate a single JSON file to directory structure.

    Returns (exercise_count, list_of_errors).
    """
    errors = []
    basename = os.path.basename(json_path).replace('.json', '')
    target_dir = os.path.join(JSON_DIR, basename)

    # Read source
    with open(json_path, 'r', encoding='utf-8') as f:
        data = json.load(f)

    cwiczenia = data.get('cwiczenia', [])
    if not cwiczenia:
        errors.append(f"{basename}: no cwiczenia found")
        return 0, errors

    # Build _meta.json
    meta = {
        "typ": data["typ"],
        "nazwa": data["nazwa"],
        "kategoria": data["kategoria"],
        "czestotliwosc": data["czestotliwosc"],
        "punkty_lacznie": data["punkty_lacznie"],
        "tagi_globalne": data["tagi_globalne"],
        "cwiczenia": []
    }

    exercises_to_write = []
    for ex in cwiczenia:
        eid = ex["id"]
        meta["cwiczenia"].append({
            "id": eid,
            "trudnosc": ex["trudnosc"],
            "punkty": ex["punkty"],
            "tagi": ex["tagi"]
        })
        exercises_to_write.append((eid, ex))

    if dry_run:
        print(f"  [DRY] {basename}: {len(cwiczenia)} exercises -> {target_dir}/")
        return len(cwiczenia), errors

    # Create directory
    os.makedirs(target_dir, exist_ok=True)

    # Write _meta.json
    meta_path = os.path.join(target_dir, '_meta.json')
    with open(meta_path, 'w', encoding='utf-8') as f:
        json.dump(meta, f, indent=2, ensure_ascii=False)
        f.write('\n')

    # Write individual exercise files
    for eid, ex in exercises_to_write:
        ex_path = os.path.join(target_dir, f'{eid}.json')
        with open(ex_path, 'w', encoding='utf-8') as f:
            json.dump(ex, f, indent=2, ensure_ascii=False)
            f.write('\n')

    # Verify: re-read and compare
    with open(meta_path, 'r', encoding='utf-8') as f:
        meta_check = json.load(f)
    if len(meta_check['cwiczenia']) != len(cwiczenia):
        errors.append(f"{basename}: meta has {len(meta_check['cwiczenia'])} entries, expected {len(cwiczenia)}")

    for eid, ex in exercises_to_write:
        ex_path = os.path.join(target_dir, f'{eid}.json')
        with open(ex_path, 'r', encoding='utf-8') as f:
            ex_check = json.load(f)
        if ex_check.get('tresc') != ex.get('tresc'):
            errors.append(f"{basename}/{eid}: tresc mismatch after write")
        if ex_check.get('odpowiedz') != ex.get('odpowiedz'):
            errors.append(f"{basename}/{eid}: odpowiedz mismatch after write")

    # Remove source file only if no errors
    if not errors:
        os.remove(json_path)

    return len(cwiczenia), errors


def main():
    parser = argparse.ArgumentParser(description='Migrate exercise JSONs to directory structure')
    parser.add_argument('--dry-run', action='store_true', help='Preview without changes')
    parser.add_argument('--file', help='Migrate only this file (prefix match)')
    args = parser.parse_args()

    # Find JSON files (exclude tagi_rejestr.json and any non-exercise files)
    all_files = sorted(glob.glob(os.path.join(JSON_DIR, '[0-9]*.json')))

    if args.file:
        all_files = [f for f in all_files if args.file in os.path.basename(f)]

    if not all_files:
        print("No files to migrate.")
        return 0

    print(f"Migrating {len(all_files)} files...")
    if args.dry_run:
        print("(DRY RUN — no changes will be made)\n")
    print()

    total_exercises = 0
    total_errors = []

    for json_path in all_files:
        basename = os.path.basename(json_path).replace('.json', '')
        count, errors = migrate_file(json_path, args.dry_run)
        total_exercises += count
        total_errors.extend(errors)

        status = "OK" if not errors else f"ERRORS: {len(errors)}"
        print(f"  {basename}: {count} exercises [{status}]")

    print()
    print(f"Total: {len(all_files)} files, {total_exercises} exercises")

    if total_errors:
        print(f"\nERRORS ({len(total_errors)}):")
        for e in total_errors:
            print(f"  {e}")
        return 1

    if not args.dry_run:
        print("\nMigration complete. Old JSON files removed.")
    return 0


if __name__ == '__main__':
    sys.exit(main())
