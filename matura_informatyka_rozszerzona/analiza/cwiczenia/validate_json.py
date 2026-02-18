#!/usr/bin/env python3
"""Standalone JSON schema validator for exercise files.

Validates all 23 exercise directories (each with _meta.json + per-exercise files)
against the schema, tag registry, and category-specific rules.

Directory structure:
    json/NN_nazwa/
        _meta.json      — metadata + exercise index
        NN.1.json       — individual exercise
        NN.2.json
        ...

Usage:
    python3 validate_json.py                  # validate all
    python3 validate_json.py --file 07_cyfry  # one directory (prefix match)
"""
import re
import json
import glob
import os
import sys
import argparse

BASE_DIR = os.path.dirname(os.path.abspath(__file__))
JSON_DIR = os.path.join(BASE_DIR, "json")
REGISTRY_PATH = os.path.join(JSON_DIR, "tagi_rejestr.json")

VALID_KATEGORIA = {'TEORIA', 'IMPLEMENTACJA', 'ARKUSZ', 'SQL'}
VALID_TRUDNOSC = {'latwe', 'srednie', 'srednie-trudne', 'trudne'}


class ValidationReport:
    def __init__(self):
        self.errors = []
        self.warnings = []
        self.files = 0
        self.exercises = 0

    def error(self, file, context, msg):
        self.errors.append(f"[ERROR] {file} / {context}: {msg}")

    def warn(self, file, context, msg):
        self.warnings.append(f"[WARN]  {file} / {context}: {msg}")


def load_tag_registry(path: str) -> set[str] | None:
    """Load the central tag registry. Returns None if file missing."""
    if not os.path.exists(path):
        return None
    with open(path, 'r', encoding='utf-8') as f:
        data = json.load(f)
    return set(data.get('tagi', []))


def validate_header(meta: dict, basename: str, report: ValidationReport,
                    registry: set[str] | None):
    """Validate _meta.json header fields."""
    # typ must match directory name
    typ = meta.get('typ', '')
    if typ != basename:
        report.error(basename, '_meta', f"typ='{typ}' != dirname '{basename}'")

    # nazwa non-empty
    if not meta.get('nazwa'):
        report.error(basename, '_meta', "empty nazwa")

    # kategoria
    kat = meta.get('kategoria', '')
    if kat not in VALID_KATEGORIA:
        report.error(basename, '_meta', f"invalid kategoria: '{kat}'")

    # czestotliwosc format
    czest = meta.get('czestotliwosc', '')
    if not re.match(r'\d+/\d+ lat', czest):
        report.error(basename, '_meta', f"bad czestotliwosc format: '{czest}'")

    # punkty_lacznie > 0
    punkty_lacznie = meta.get('punkty_lacznie', 0)
    if punkty_lacznie <= 0:
        report.error(basename, '_meta', f"punkty_lacznie={punkty_lacznie} (must be >0)")

    # punkty_lacznie == sum of exercise points (from meta index)
    cwiczenia_index = meta.get('cwiczenia', [])
    suma = sum(ex.get('punkty', 0) for ex in cwiczenia_index)
    if suma != punkty_lacznie:
        report.error(basename, '_meta',
                     f"punkty_lacznie={punkty_lacznie} != sum(cwiczenia.punkty)={suma}")

    # tagi_globalne non-empty
    tagi_globalne = meta.get('tagi_globalne', [])
    if not tagi_globalne:
        report.error(basename, '_meta', "empty tagi_globalne")

    # tagi_globalne in registry
    if registry is not None:
        for tag in tagi_globalne:
            if tag not in registry:
                report.error(basename, '_meta',
                             f"tag '{tag}' in tagi_globalne not in registry")


def validate_exercise(ex: dict, basename: str, file_prefix: int,
                      tagi_globalne: list[str], report: ValidationReport,
                      registry: set[str] | None, seen_ids: set[str]):
    """Validate a single exercise object."""
    eid = ex.get('id', '???')

    # id format NN.M, NN matches file prefix
    m = re.match(r'^(\d+)\.(\d+)$', eid)
    if not m:
        report.error(basename, eid, f"bad id format: '{eid}'")
    else:
        id_prefix = int(m.group(1))
        if id_prefix != file_prefix:
            report.error(basename, eid,
                         f"id prefix {id_prefix} != file prefix {file_prefix}")

    # unique id within directory
    if eid in seen_ids:
        report.error(basename, eid, "duplicate id within directory")
    seen_ids.add(eid)

    # trudnosc
    trudnosc = ex.get('trudnosc', '')
    if trudnosc not in VALID_TRUDNOSC:
        report.error(basename, eid, f"invalid trudnosc: '{trudnosc}'")

    # punkty 1-10
    punkty = ex.get('punkty', 0)
    if not (1 <= punkty <= 10):
        report.error(basename, eid, f"punkty={punkty} (must be 1-10)")

    # zrodlo non-empty
    zrodlo = ex.get('zrodlo', '')
    if not zrodlo:
        report.error(basename, eid, "empty zrodlo")

    # tagi non-empty, subset of tagi_globalne, in registry
    tagi = ex.get('tagi', [])
    if not tagi:
        report.error(basename, eid, "empty tagi")

    tagi_globalne_set = set(tagi_globalne)
    for tag in tagi:
        if tag not in tagi_globalne_set:
            report.error(basename, eid,
                         f"tag '{tag}' not in tagi_globalne")
        if registry is not None and tag not in registry:
            report.error(basename, eid,
                         f"tag '{tag}' not in registry")

    # tresc non-empty, >50 chars
    tresc = ex.get('tresc', '')
    if not tresc:
        report.error(basename, eid, "empty tresc")
    elif len(tresc) < 50:
        report.error(basename, eid, f"tresc too short ({len(tresc)} chars, need >50)")

    # wskazowki: array of 3, each >10 chars
    wskazowki = ex.get('wskazowki', [])
    if len(wskazowki) != 3:
        report.error(basename, eid, f"wskazowki count={len(wskazowki)} (need 3)")
    for i, w in enumerate(wskazowki):
        if len(w) < 10:
            report.error(basename, eid,
                         f"wskazowki[{i}] too short ({len(w)} chars, need >10)")

    # odpowiedz non-empty, >50 chars
    odpowiedz = ex.get('odpowiedz', '')
    if not odpowiedz:
        report.error(basename, eid, "empty odpowiedz")
    elif len(odpowiedz) < 50:
        report.error(basename, eid,
                     f"odpowiedz too short ({len(odpowiedz)} chars, need >50)")

    # typowe_bledy: array >=1, each has opis and kara
    bledy = ex.get('typowe_bledy', [])
    if len(bledy) < 1:
        report.error(basename, eid, "typowe_bledy empty (need >=1)")
    for i, b in enumerate(bledy):
        if not b.get('opis'):
            report.error(basename, eid, f"typowe_bledy[{i}]: empty opis")
        kara = b.get('kara', '')
        if not kara:
            report.error(basename, eid, f"typowe_bledy[{i}]: missing or empty kara")
        elif not re.match(r'^-\d+(\.\d+)? pkt$', kara):
            report.error(basename, eid,
                         f"typowe_bledy[{i}]: invalid kara format '{kara}' (expected '-N pkt')")


def validate_category_rules(ex: dict, basename: str, file_prefix: int,
                            kategoria: str, report: ValidationReport):
    """Category-specific validation rules."""
    eid = ex.get('id', '???')
    tresc = ex.get('tresc', '')
    odpowiedz = ex.get('odpowiedz', '')

    if kategoria == 'IMPLEMENTACJA':
        # odpowiedz must contain ```cpp block
        if '```cpp' not in odpowiedz:
            report.error(basename, eid,
                         "IMPLEMENTACJA: odpowiedz missing ```cpp code block")
        # tresc must contain **Dane** or **Oczekiwany wynik**
        if '**Dane**' not in tresc and '**Oczekiwany wynik**' not in tresc:
            report.error(basename, eid,
                         "IMPLEMENTACJA: tresc missing **Dane** or **Oczekiwany wynik**")

    elif kategoria == 'SQL':
        # odpowiedz must contain ```sql block
        if '```sql' not in odpowiedz:
            report.error(basename, eid,
                         "SQL: odpowiedz missing ```sql code block")
        # tresc must contain table definition
        if 'Tabela **' not in tresc:
            report.error(basename, eid,
                         "SQL: tresc missing 'Tabela **' definition")

    elif file_prefix == 5:
        # TEORIA/konwersje: odpowiedz should contain conversion patterns
        has_conv = bool(re.search(r'\d+\(\d+\)', odpowiedz))
        if not has_conv:
            report.warn(basename, eid,
                        "konwersje (05): odpowiedz has no conversion pattern like '101(2)'")


def validate_meta_consistency(meta: dict, basename: str, dir_path: str,
                              report: ValidationReport):
    """Check that _meta.json index matches actual exercise files on disk."""
    meta_ids = {ex['id'] for ex in meta.get('cwiczenia', [])}

    # Find all exercise JSON files (exclude _meta.json)
    disk_files = glob.glob(os.path.join(dir_path, '[0-9]*.json'))
    disk_ids = set()
    for fp in disk_files:
        fname = os.path.basename(fp).replace('.json', '')
        disk_ids.add(fname)

    # IDs in meta but not on disk
    missing_on_disk = meta_ids - disk_ids
    for mid in sorted(missing_on_disk):
        report.error(basename, mid, f"listed in _meta.json but file {mid}.json not found")

    # Files on disk but not in meta
    extra_on_disk = disk_ids - meta_ids
    for eid in sorted(extra_on_disk):
        report.error(basename, eid, f"file {eid}.json exists but not listed in _meta.json")


def validate_directory(basename: str, report: ValidationReport,
                       registry: set[str] | None):
    """Validate a single exercise directory."""
    dir_path = os.path.join(JSON_DIR, basename)
    meta_path = os.path.join(dir_path, '_meta.json')

    if not os.path.isdir(dir_path):
        report.error(basename, '-', f"directory not found: {dir_path}")
        return

    if not os.path.exists(meta_path):
        report.error(basename, '-', f"_meta.json not found in {dir_path}")
        return

    try:
        with open(meta_path, 'r', encoding='utf-8') as f:
            meta = json.load(f)
    except json.JSONDecodeError as e:
        report.error(basename, '_meta', f"invalid JSON: {e}")
        return

    report.files += 1

    # Extract file prefix
    try:
        file_prefix = int(basename.split('_')[0])
    except ValueError:
        report.error(basename, '-', f"cannot extract numeric prefix from '{basename}'")
        return

    kategoria = meta.get('kategoria', '')
    tagi_globalne = meta.get('tagi_globalne', [])

    # Header validation
    validate_header(meta, basename, report, registry)

    # Meta-disk consistency check
    validate_meta_consistency(meta, basename, dir_path, report)

    # Per-exercise validation
    cwiczenia_index = meta.get('cwiczenia', [])
    if not cwiczenia_index:
        report.error(basename, '-', "no cwiczenia found in _meta.json")
        return

    seen_ids = set()
    for ex_entry in cwiczenia_index:
        eid = ex_entry.get('id', '???')
        ex_path = os.path.join(dir_path, f'{eid}.json')

        if not os.path.exists(ex_path):
            # Already reported in consistency check
            continue

        try:
            with open(ex_path, 'r', encoding='utf-8') as f:
                ex = json.load(f)
        except json.JSONDecodeError as e:
            report.error(basename, eid, f"invalid JSON in {eid}.json: {e}")
            continue

        report.exercises += 1

        # Check that meta index fields match exercise file
        for field in ('trudnosc', 'punkty', 'tagi'):
            if field in ex_entry and ex.get(field) != ex_entry[field]:
                report.error(basename, eid,
                             f"_meta.cwiczenia.{field} != {eid}.json.{field}")

        validate_exercise(ex, basename, file_prefix, tagi_globalne,
                          report, registry, seen_ids)
        validate_category_rules(ex, basename, file_prefix, kategoria, report)


def main():
    parser = argparse.ArgumentParser(
        description='Standalone JSON schema validator for exercise directories')
    parser.add_argument('--file', help='Validate only directories matching this prefix (e.g., 07_cyfry)')
    args = parser.parse_args()

    # Load tag registry
    registry = load_tag_registry(REGISTRY_PATH)
    if registry is None:
        print(f"WARNING: Tag registry not found at {REGISTRY_PATH}")
        print("         Tag validation will be skipped.")
    else:
        print(f"Tag registry loaded: {len(registry)} tags")

    # Find exercise directories
    all_dirs = sorted([
        d for d in os.listdir(JSON_DIR)
        if os.path.isdir(os.path.join(JSON_DIR, d)) and d[0].isdigit()
    ])

    if args.file:
        all_dirs = [d for d in all_dirs if args.file in d]

    print(f"Validating {len(all_dirs)} directories...")
    print()

    report = ValidationReport()

    for dirname in all_dirs:
        validate_directory(dirname, report, registry)

    # --- Report ---
    print(f"Validated: {report.files} directories, {report.exercises} exercises")

    if report.errors:
        print(f"\n{'=' * 60}")
        print(f"ERRORS: {len(report.errors)}")
        print(f"{'=' * 60}")
        for e in report.errors:
            print(e)

    if report.warnings:
        print(f"\n{'=' * 60}")
        print(f"WARNINGS: {len(report.warnings)}")
        print(f"{'=' * 60}")
        for w in report.warnings:
            print(w)

    if not report.errors and not report.warnings:
        print("\nPERFECT — all exercise directories pass schema validation, zero issues.")
    elif not report.errors:
        print(f"\nNo errors. {len(report.warnings)} warnings (see above).")
    else:
        print(f"\n{len(report.errors)} ERRORS found — need fixing!")

    return 1 if report.errors else 0


if __name__ == '__main__':
    sys.exit(main())
