#!/usr/bin/env python3
"""
polonizuj_formuly.py — polonizacja angielskich nazw formul Excel w materialach.

Tryby:
  --dry-run                  Pokaz diff, nic nie zmieniaj
  --apply                    Wykonaj zamiany
  --warstwa {1,2,3,4,all}    Wybierz warstwe (1=dedykowane MD, 2=mieszane MD,
                             3=JSON cwiczen, 4=rejestry, all=wszystkie)
"""
import argparse
import json
import re
import sys
from pathlib import Path

REPO_ROOT = Path(__file__).resolve().parents[3]
SCRIPTS_DIR = Path(__file__).resolve().parent

# WHITELIST: jedyne pliki ktore skrypt moze edytowac
WHITELIST_WARSTWA_1 = [
    "matura_informatyka_rozszerzona/analiza/szablony/arkusz_formuly.md",
    "matura_informatyka_rozszerzona/analiza/cheatsheets/cheatsheet_arkusz.md",
    "matura_informatyka_rozszerzona/analiza/rozwiazania_wzorcowe/arkusz_kalkulacyjny.md",
]
WHITELIST_WARSTWA_3_DIRS = [
    "matura_informatyka_rozszerzona/analiza/cwiczenia/json/15_agregacja_warunkowa",
    "matura_informatyka_rozszerzona/analiza/cwiczenia/json/16_symulacja",
    "matura_informatyka_rozszerzona/analiza/cwiczenia/json/18_agregacja_podstawowa",
    "matura_informatyka_rozszerzona/analiza/cwiczenia/json/19_transformacja",
]
WHITELIST_WARSTWA_4 = [
    "matura_informatyka_rozszerzona/analiza/cwiczenia/json/tagi_rejestr.json",
    "matura_informatyka_rozszerzona/analiza/json/algorytmy_rejestr.json",
]


def load_mapa():
    with open(SCRIPTS_DIR / "polonizacja_mapa.json") as f:
        return json.load(f)


def load_warstwa2():
    with open(SCRIPTS_DIR / "polonizacja_warstwa2.json") as f:
        return json.load(f)


def is_whitelisted(path: str, warstwa: str) -> bool:
    """Zwraca True jesli sciezka jest na liscie pozwolen dla danej warstwy."""
    p = Path(path)
    rel = str(p.relative_to(REPO_ROOT)) if p.is_absolute() else path
    if warstwa in ("1", "all") and rel in WHITELIST_WARSTWA_1:
        return True
    if warstwa in ("2", "all"):
        warstwa2 = load_warstwa2()
        if rel in {z["plik"] for z in warstwa2["zamiany"]}:
            return True
    if warstwa in ("3", "all"):
        for d in WHITELIST_WARSTWA_3_DIRS:
            if rel.startswith(d + "/"):
                return True
    if warstwa in ("4", "all") and rel in WHITELIST_WARSTWA_4:
        return True
    return False


def polonizuj_md_warstwa1(path: Path, mapa: dict) -> int:
    """Globalna podmiana 46 funkcji EN->PL w pliku MD warstwy 1.

    Reguly:
    - Dluzsze nazwy najpierw (SUMIFS przed SUMIF)
    - Tylko z otwierajacym '(' (\\b NAZWA\\s*\\()
    - Case-sensitive (wielkie litery)

    Zwraca: liczba zamian w pliku.
    """
    content = path.read_text()
    original = content
    # Sortuj klucze od najdluzszych — uniknij SUMIFS -> SUMA + IFS
    keys = sorted(mapa["mapowanie"].keys(), key=len, reverse=True)
    total_changes = 0
    for en in keys:
        pl = mapa["mapowanie"][en]
        # \b NAZWA \s* \( — pozwala na biale znaki przed (
        pattern = r"\b" + re.escape(en) + r"(\s*\()"
        new_content, count = re.subn(pattern, pl + r"\1", content)
        if count > 0:
            total_changes += count
            content = new_content
    if content != original:
        path.write_text(content)
    return total_changes


def main():
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    g = parser.add_mutually_exclusive_group(required=True)
    g.add_argument("--dry-run", action="store_true", help="Pokaz diff, nic nie zmieniaj")
    g.add_argument("--apply", action="store_true", help="Wykonaj zamiany")
    parser.add_argument("--warstwa", choices=["1", "2", "3", "4", "all"], default="all")
    args = parser.parse_args()

    mapa = load_mapa()
    print(f"Mapa: {len(mapa['mapowanie'])} funkcji EN->PL")
    print(f"Warstwa: {args.warstwa}, Tryb: {'DRY-RUN' if args.dry_run else 'APPLY'}")
    # Faza implementacji per warstwa - dodane w kolejnych zadaniach
    print("TODO: implementuj logike per warstwa")


if __name__ == "__main__":
    main()
