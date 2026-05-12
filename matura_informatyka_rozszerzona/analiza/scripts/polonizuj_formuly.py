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


def polonizuj_md_warstwa2(zamiany: list) -> int:
    """Stosuje deterministyczna liste zamian dla plikow MD mieszanych.

    Dla kazdej zamiany:
    - Pre-check: STARY istnieje w pliku (count >= 1)
    - Apply: str.replace(STARY, NOWY, 1)
    - Bledy: zbierane do listy i raportowane na koncu

    Zwraca: liczba pomyslnych zamian.
    """
    total = 0
    errors = []
    # Grupuj po pliku
    by_file = {}
    for z in zamiany:
        by_file.setdefault(z["plik"], []).append(z)

    for plik, zlist in by_file.items():
        path = Path(plik)
        if not path.exists():
            errors.append(f"BRAK PLIKU: {plik}")
            continue
        content = path.read_text()
        for z in zlist:
            if z["stary"] not in content:
                errors.append(f"NIE ZNALEZIONO STARY w {plik}:{z['linia']}: {z['stary'][:60]}...")
                continue
            if content.count(z["stary"]) > 1:
                errors.append(f"WIELOKROTNE STARY w {plik}:{z['linia']}: {z['stary'][:60]}...")
                continue
            content = content.replace(z["stary"], z["nowy"], 1)
            total += 1
        path.write_text(content)

    if errors:
        print("BLEDY warstwy 2:", file=sys.stderr)
        for e in errors:
            print(f"  {e}", file=sys.stderr)
        raise SystemExit(2)
    return total


def _polonizuj_string(text: str, mapa: dict, require_paren: bool = True) -> tuple[str, int]:
    """Wewnetrzna helper: podmiana w stringu, sortuje keys po dlugosci desc.

    require_paren=True: tylko z otwierajacym '(' (jak w warstwie 1, dla tresc/odpowiedz).
    require_paren=False: tez bez nawiasu (dla wskazowek/bledow, gdzie nazwa funkcji
    moze byc cytowana w prozie). Stosuje sie tylko do dluzszych nazw (>=5 znakow),
    aby uniknac kolizji z angielskimi krotszymi slowami (IF, SUM, OR, AND, NOT).
    """
    keys = sorted(mapa["mapowanie"].keys(), key=len, reverse=True)
    total = 0
    for en in keys:
        pl = mapa["mapowanie"][en]
        if require_paren:
            pattern = r"\b" + re.escape(en) + r"(\s*\()"
            text, count = re.subn(pattern, pl + r"\1", text)
        else:
            # Najpierw forma z nawiasem (zawsze bezpieczna)
            pattern1 = r"\b" + re.escape(en) + r"(\s*\()"
            text, c1 = re.subn(pattern1, pl + r"\1", text)
            count = c1
            # Forma bez nawiasu — tylko dla dluzszych unikalnych nazw (>=5 znakow)
            if len(en) >= 5:
                pattern2 = r"\b" + re.escape(en) + r"\b"
                text, c2 = re.subn(pattern2, pl, text)
                count += c2
        total += count
    return text, total


def _rename_tag(tag: str, tagi_map: dict) -> str:
    return tagi_map.get(tag, tag)


def _sort_pl(items):
    """Sortuje po polsku: litery z diakrytykami tuz po bazowych literach.

    Klucz: dla kazdego znaku zwraca pare (baza_lower, oryginalny_ord),
    co daje porzadek typu: S, SUMA, ŚREDNIA, T, ... oraz wielkie litery
    przed malymi tej samej bazy."""
    import unicodedata
    def key(s):
        out = []
        for ch in s:
            decomp = unicodedata.normalize("NFKD", ch)
            base = decomp[0]
            out.append((base.lower(), ord(ch)))
        return out
    return sorted(items, key=key)


def polonizuj_json_cwiczenie(path: Path, mapa: dict) -> int:
    """Polonizuje pola tekstowe oraz `tagi` w pliku cwiczenia JSON."""
    data = json.loads(path.read_text())
    total = 0

    # Pola tekstowe — tresc/odpowiedz: wymagaja '(' (formuly w kodzie)
    for field in ("tresc", "odpowiedz"):
        if field in data and isinstance(data[field], str):
            new, c = _polonizuj_string(data[field], mapa, require_paren=True)
            data[field] = new
            total += c

    # Wskazowki/typowe_bledy: dopuszczamy bare-name (proza wspomina nazwe funkcji)
    for field in ("wskazowki", "typowe_bledy"):
        if field in data and isinstance(data[field], list):
            for item in data[field]:
                for k in ("tekst", "opis"):
                    if k in item and isinstance(item[k], str):
                        new, c = _polonizuj_string(item[k], mapa, require_paren=False)
                        item[k] = new
                        total += c

    if "weryfikacja_szczegolowa" in data and isinstance(data["weryfikacja_szczegolowa"], str):
        new, c = _polonizuj_string(data["weryfikacja_szczegolowa"], mapa, require_paren=True)
        data["weryfikacja_szczegolowa"] = new
        total += c

    # Tagi
    if "tagi" in data and isinstance(data["tagi"], list):
        tagi_map = mapa["tagi_rejestr_rename"]
        new_tagi = [_rename_tag(t, tagi_map) for t in data["tagi"]]
        if new_tagi != data["tagi"]:
            total += sum(1 for a, b in zip(data["tagi"], new_tagi) if a != b)
            data["tagi"] = new_tagi

    path.write_text(json.dumps(data, ensure_ascii=False, indent=2) + "\n")
    return total


def polonizuj_json_meta(path: Path, mapa: dict) -> int:
    """Polonizuje pole `tagi_globalne` w _meta.json. Zachowuje sortowanie polskie."""
    data = json.loads(path.read_text())
    if "tagi_globalne" not in data:
        return 0
    tagi_map = mapa["tagi_rejestr_rename"]
    new = _sort_pl(set(_rename_tag(t, tagi_map) for t in data["tagi_globalne"]))
    if new != data["tagi_globalne"]:
        data["tagi_globalne"] = new
        path.write_text(json.dumps(data, ensure_ascii=False, indent=2) + "\n")
        return 1
    return 0


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
