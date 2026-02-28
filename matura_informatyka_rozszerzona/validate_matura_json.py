#!/usr/bin/env python3
"""
Comprehensive validator for matura exam JSON files.

Checks:
 1. JSON is valid and parseable
 2. All required top-level fields exist
 3. Each zadanie has required fields
 4. Each podzadanie has required fields
 5. ID format matches YYYYS.Z.S or YYYYS.Z or YYYYS.Za pattern
 6. Session letter in ID matches sesja field
 7. Points sum correctly at all levels
 8. Part I + Part II points match czesci definition (for stara formula)
 9. typ_zadania values are from the known 23 types
10. kategoria is one of: TEORIA, IMPLEMENTACJA, ARKUSZ, SQL
11. No duplicate IDs across ALL files
12. sciezka_danych directory actually exists on disk
13. pliki_danych files actually exist in the sciezka_danych directory
"""

import json
import os
import sys
import re
from pathlib import Path
from collections import Counter

# ── Constants ──────────────────────────────────────────────────────────────

REPO_ROOT = Path(__file__).resolve().parent  # matura_informatyka_rozszerzona/
JSON_DIR = REPO_ROOT / "analiza" / "json"

VALID_TYP_ZADANIA = {
    # TEORIA (6)
    "sledzenie_algorytmu",
    "projektowanie_algorytmu",
    "analiza_algorytmu",
    "test_prawda_falsz",
    "konwersja_systemow_liczbowych",
    "teoria_bezpieczenstwa",
    # IMPLEMENTACJA (8)
    "przetwarzanie_cyfry_liczby",
    "przetwarzanie_napisy",
    "przetwarzanie_zlozone",
    "przetwarzanie_zliczanie",
    "przetwarzanie_minmax",
    "przetwarzanie_sekwencje",
    "przetwarzanie_obrazy_2D",
    "przetwarzanie_geometryczne",
    # ARKUSZ (5)
    "arkusz_agregacja_warunkowa",
    "arkusz_symulacja",
    "arkusz_wykres",
    "arkusz_agregacja_podstawowa",
    "arkusz_transformacja",
    # SQL (4)
    "sql_group_by",
    "sql_podzapytania",
    "sql_join",
    "sql_select_where",
}

VALID_KATEGORIA = {"TEORIA", "IMPLEMENTACJA", "ARKUSZ", "SQL"}

REQUIRED_TOP_LEVEL = {"rok", "sesja", "formula", "czas_minuty", "total_punkty",
                      "liczba_zadan", "czesci", "zadania"}

REQUIRED_ZADANIE = {"numer", "tytul", "punkty", "czesc", "kontekst", "podzadania"}

REQUIRED_PODZADANIE = {"id", "numer", "punkty", "typ_zadania", "kategoria",
                       "tresc", "odpowiedz", "zasady_oceniania", "pulapki"}

# sesja value -> expected letter in ID
SESJA_TO_LETTER = {
    "maj": "M",
    "czerwiec": "C",
    "probna": "P",
    "probna_grudzien": "P",
    "probna_kwiecien": "P",
    "probna_marzec": "P",
    "przykladowy": "X",
}

# Accepted ID patterns:
#   YYYYS.Z.S  (e.g. 2025M.1.1) — standard
#   YYYYS.Z    (e.g. 2025M.4)   — single-subtask tasks (no subtask number)
#   YYYYS.Za   (e.g. 2014M.1a)  — letter-based subtask (legacy 2014M format)
ID_PATTERN_FULL = re.compile(r"^(\d{4})([MCPX])\.(\d+)\.(\d+)$")
ID_PATTERN_SINGLE = re.compile(r"^(\d{4})([MCPX])\.(\d+)$")
ID_PATTERN_LETTER = re.compile(r"^(\d{4})([MCPX])\.(\d+)([a-z])$")


def resolve_sciezka(sciezka: str, rok: int = None, sesja: str = None):
    """Try to resolve sciezka_danych to an actual directory on disk.

    sciezka_danych values come in several formats:
    - Full relative from parent: "matura_informatyka_rozszerzona/2025_maj/dane maj 23/"
    - Bare directory name: "DANE", "dane_PR/"

    Returns (resolved_path, is_inconsistent):
    - resolved_path: Path or None if not found anywhere
    - is_inconsistent: True if path doesn't follow the standard format but was
      found via fallback (year-dir relative resolution)
    """
    s = sciezka.rstrip("/")

    # Try from parent of repo root (handles "matura_informatyka_rozszerzona/..." prefix)
    candidate = REPO_ROOT.parent / s
    if candidate.is_dir():
        return candidate, False

    # Try from repo root itself
    candidate = REPO_ROOT / s
    if candidate.is_dir():
        return candidate, False

    # Fallback: try inside each year directory that matches rok/sesja
    # This handles bare paths like "DANE" or "dane_PR/" that are relative to the year dir
    if rok is not None:
        sesja_map = {
            "maj": "maj", "czerwiec": "czerwiec",
            "probna": "probna", "probna_grudzien": "probna_grudzien",
            "probna_kwiecien": "probna_kwiecien", "probna_marzec": "probna_marzec",
            "przykladowy": "przykladowy",
        }
        sesja_dir = sesja_map.get(sesja, sesja) if sesja else None
        if sesja_dir:
            year_dir = REPO_ROOT / f"{rok}_{sesja_dir}"
            candidate = year_dir / s
            if candidate.is_dir():
                return candidate, True  # Found but via inconsistent path

    return None, False


# ── Validation ─────────────────────────────────────────────────────────────

def validate_file(filepath: Path, all_ids: dict, errors: list, warnings: list):
    """Validate a single matura JSON file. Appends issues to errors/warnings."""
    fname = filepath.name
    prefix = f"[{fname}]"

    # 1. JSON is valid and parseable
    try:
        with open(filepath, "r", encoding="utf-8") as f:
            data = json.load(f)
    except json.JSONDecodeError as e:
        errors.append(f"{prefix} PARSE ERROR: {e}")
        return
    except Exception as e:
        errors.append(f"{prefix} READ ERROR: {e}")
        return

    # 2. All required top-level fields exist
    missing_top = REQUIRED_TOP_LEVEL - set(data.keys())
    for field in sorted(missing_top):
        errors.append(f"{prefix} Missing top-level field: '{field}'")

    if missing_top & {"zadania", "czesci", "total_punkty", "sesja", "rok"}:
        return

    rok = data.get("rok")
    sesja = data.get("sesja")
    formula = data.get("formula")
    total_punkty = data.get("total_punkty")
    czesci = data.get("czesci", [])
    zadania = data.get("zadania", [])
    liczba_zadan = data.get("liczba_zadan")

    # Cross-check filename vs content
    filename_match = re.match(r"matura_(\d{4})([MCPX])\.json", fname)
    if filename_match:
        file_rok = int(filename_match.group(1))
        file_letter = filename_match.group(2)
        if file_rok != rok:
            errors.append(f"{prefix} rok={rok} does not match filename year {file_rok}")
        expected_letter = SESJA_TO_LETTER.get(sesja)
        if expected_letter and expected_letter != file_letter:
            errors.append(f"{prefix} sesja='{sesja}' -> letter '{expected_letter}' does not match filename letter '{file_letter}'")

    # Check sesja value is known
    if sesja not in SESJA_TO_LETTER:
        errors.append(f"{prefix} Unknown sesja value: '{sesja}'")

    # Check liczba_zadan matches actual count
    if liczba_zadan is not None and liczba_zadan != len(zadania):
        errors.append(f"{prefix} liczba_zadan={liczba_zadan} but found {len(zadania)} zadania")

    # 8. czesci total must match total_punkty
    czesci_points = {}
    total_czesci_points = 0
    for c in czesci:
        cz = c.get("czesc")
        pts = c.get("punkty", 0)
        czesci_points[cz] = pts
        total_czesci_points += pts

    if total_czesci_points != total_punkty:
        errors.append(f"{prefix} Sum of czesci punkty ({total_czesci_points}) != total_punkty ({total_punkty})")

    # Process each zadanie
    grand_total = 0
    czesc_sums = Counter()

    for zad in zadania:
        zad_num = zad.get("numer", "?")
        zad_prefix = f"{prefix} Zad.{zad_num}"

        # 3. Each zadanie has required fields
        missing_zad = REQUIRED_ZADANIE - set(zad.keys())
        for field in sorted(missing_zad):
            errors.append(f"{zad_prefix} Missing zadanie field: '{field}'")

        zad_punkty = zad.get("punkty", 0)
        zad_czesc = zad.get("czesc")
        podzadania = zad.get("podzadania", [])

        grand_total += zad_punkty
        if zad_czesc is not None:
            czesc_sums[zad_czesc] += zad_punkty

        # 12 & 13. sciezka_danych and pliki_danych (per-zadanie)
        sciezka = zad.get("sciezka_danych")
        pliki = zad.get("pliki_danych", [])

        if sciezka is not None:
            resolved, is_inconsistent = resolve_sciezka(sciezka, rok, sesja)
            if resolved is None:
                errors.append(f"{zad_prefix} sciezka_danych directory does not exist: '{sciezka}'")
            else:
                if is_inconsistent:
                    warnings.append(
                        f"{zad_prefix} sciezka_danych '{sciezka}' uses bare dir name "
                        f"(should be 'matura_informatyka_rozszerzona/{rok}_.../{sciezka.rstrip('/')}/')"
                    )
                # 13. Check each file exists
                for plik in pliki:
                    plik_path = resolved / plik
                    if not plik_path.is_file():
                        errors.append(f"{zad_prefix} pliki_danych file missing: '{plik}' in '{sciezka}'")
        elif pliki:
            warnings.append(f"{zad_prefix} Has pliki_danych but sciezka_danych is null")

        # 7. Points sum: sum of podzadania == zadanie punkty
        podz_sum = sum(p.get("punkty", 0) for p in podzadania)
        if podz_sum != zad_punkty:
            errors.append(f"{zad_prefix} Sum of podzadania punkty ({podz_sum}) != zadanie punkty ({zad_punkty})")

        # Process each podzadanie
        for podz in podzadania:
            podz_id = podz.get("id", "???")
            podz_prefix = f"{prefix} [{podz_id}]"

            # 4. Each podzadanie has required fields
            missing_podz = REQUIRED_PODZADANIE - set(podz.keys())
            for field in sorted(missing_podz):
                errors.append(f"{podz_prefix} Missing podzadanie field: '{field}'")

            # 5. ID format validation (three accepted patterns)
            m_full = ID_PATTERN_FULL.match(podz_id)
            m_single = ID_PATTERN_SINGLE.match(podz_id)
            m_letter = ID_PATTERN_LETTER.match(podz_id)

            id_rok = None
            id_letter = None
            id_zad = None

            if m_full:
                id_rok = int(m_full.group(1))
                id_letter = m_full.group(2)
                id_zad = int(m_full.group(3))
            elif m_single:
                id_rok = int(m_single.group(1))
                id_letter = m_single.group(2)
                id_zad = int(m_single.group(3))
                # Single-subtask: warn if task has multiple podzadania
                if len(podzadania) > 1:
                    errors.append(f"{podz_prefix} ID has no subtask number but task has {len(podzadania)} podzadania")
            elif m_letter:
                id_rok = int(m_letter.group(1))
                id_letter = m_letter.group(2)
                id_zad = int(m_letter.group(3))
            else:
                errors.append(f"{podz_prefix} ID does not match any valid pattern (YYYYS.Z.S / YYYYS.Z / YYYYS.Za)")

            if id_rok is not None:
                # Check rok in ID matches file rok
                if id_rok != rok:
                    errors.append(f"{podz_prefix} ID year {id_rok} != file rok {rok}")

                # 6. Session letter in ID matches sesja field
                expected = SESJA_TO_LETTER.get(sesja)
                if expected and id_letter != expected:
                    errors.append(f"{podz_prefix} ID letter '{id_letter}' does not match sesja '{sesja}' (expected '{expected}')")

                # Check zadanie number in ID matches parent
                if id_zad != zad_num:
                    errors.append(f"{podz_prefix} ID zadanie number {id_zad} != parent numer {zad_num}")

            # 9. typ_zadania is from the known 23 types
            typ = podz.get("typ_zadania")
            if typ and typ not in VALID_TYP_ZADANIA:
                errors.append(f"{podz_prefix} Unknown typ_zadania: '{typ}'")

            # 10. kategoria is valid
            kat = podz.get("kategoria")
            if kat and kat not in VALID_KATEGORIA:
                errors.append(f"{podz_prefix} Unknown kategoria: '{kat}'")

            # Cross-check: kategoria should match typ_zadania prefix
            if typ and kat:
                expected_kat = None
                if typ.startswith("przetwarzanie_"):
                    expected_kat = "IMPLEMENTACJA"
                elif typ.startswith("arkusz_"):
                    expected_kat = "ARKUSZ"
                elif typ.startswith("sql_"):
                    expected_kat = "SQL"
                elif typ in {"sledzenie_algorytmu", "projektowanie_algorytmu",
                             "analiza_algorytmu", "test_prawda_falsz",
                             "konwersja_systemow_liczbowych", "teoria_bezpieczenstwa"}:
                    expected_kat = "TEORIA"

                if expected_kat and kat != expected_kat:
                    errors.append(f"{podz_prefix} typ_zadania '{typ}' implies kategoria '{expected_kat}' but got '{kat}'")

            # 11. Duplicate ID check (across all files)
            if podz_id in all_ids:
                errors.append(f"{podz_prefix} DUPLICATE ID! Also in {all_ids[podz_id]}")
            else:
                all_ids[podz_id] = fname

    # 7. Grand total check
    if grand_total != total_punkty:
        errors.append(f"{prefix} Sum of all zadania punkty ({grand_total}) != total_punkty ({total_punkty})")

    # 8. Per-czesc totals match czesci definition (for stara formula with 2 parts)
    if len(czesci) >= 2:
        for cz_num, expected_pts in czesci_points.items():
            actual_pts = czesc_sums.get(cz_num, 0)
            if actual_pts != expected_pts:
                errors.append(f"{prefix} Czesc {cz_num}: zadania sum={actual_pts} != czesci definition={expected_pts}")


def main():
    print("=" * 70)
    print("MATURA JSON COMPREHENSIVE VALIDATOR")
    print("=" * 70)

    json_files = sorted(JSON_DIR.glob("matura_*[MCPX].json"))

    if not json_files:
        print(f"ERROR: No matura JSON files found in {JSON_DIR}")
        sys.exit(1)

    print(f"\nFound {len(json_files)} exam JSON files to validate.\n")

    all_ids = {}
    errors = []
    warnings = []

    for fp in json_files:
        validate_file(fp, all_ids, errors, warnings)

    # ── Summary by file ───────────────────────────────────────────────

    # Count errors per file for a nice summary
    file_error_counts = Counter()
    for e in errors:
        m = re.match(r"\[([^\]]+)\]", e)
        if m:
            file_error_counts[m.group(1)] += 1

    print("-" * 70)
    print(f"TOTAL IDs processed: {len(all_ids)}")
    print(f"TOTAL FILES: {len(json_files)}")
    print("-" * 70)

    # Per-file summary
    print(f"\nPER-FILE SUMMARY:")
    for fp in json_files:
        ec = file_error_counts.get(fp.name, 0)
        status = "PASS" if ec == 0 else f"FAIL ({ec} errors)"
        print(f"  {fp.name:30s} {status}")

    if warnings:
        print(f"\n{'='*70}")
        print(f"WARNINGS ({len(warnings)}):")
        print(f"{'='*70}")
        for w in warnings:
            print(f"  WARN: {w}")

    if errors:
        print(f"\n{'='*70}")
        print(f"ERRORS ({len(errors)}):")
        print(f"{'='*70}")
        for e in errors:
            print(f"  ERR:  {e}")

        # Categorize errors
        categories = Counter()
        for e in errors:
            if "Missing top-level" in e:
                categories["missing_top_level"] += 1
            elif "Missing zadanie field" in e:
                categories["missing_zadanie_field"] += 1
            elif "Missing podzadanie field" in e:
                categories["missing_podzadanie_field"] += 1
            elif "ID does not match" in e or "ID has no subtask" in e:
                categories["id_format"] += 1
            elif "ID letter" in e or "ID year" in e or "ID zadanie number" in e:
                categories["id_mismatch"] += 1
            elif "DUPLICATE ID" in e:
                categories["duplicate_id"] += 1
            elif "sciezka_danych directory" in e:
                categories["missing_directory"] += 1
            elif "pliki_danych file" in e:
                categories["missing_file"] += 1
            elif "Unknown typ_zadania" in e:
                categories["unknown_typ"] += 1
            elif "Unknown kategoria" in e:
                categories["unknown_kategoria"] += 1
            elif "implies kategoria" in e:
                categories["typ_kat_mismatch"] += 1
            elif "Sum of" in e or "Czesc" in e or "punkty" in e:
                categories["points_mismatch"] += 1
            else:
                categories["other"] += 1

        print(f"\nERROR BREAKDOWN:")
        for cat, count in categories.most_common():
            print(f"  {cat:30s} {count}")

        print(f"\n*** VALIDATION FAILED: {len(errors)} error(s), {len(warnings)} warning(s) ***")
        sys.exit(1)
    else:
        print(f"\n*** ALL {len(json_files)} FILES PASSED ({len(all_ids)} subtasks, {len(warnings)} warnings) ***")
        sys.exit(0)


if __name__ == "__main__":
    main()
