#!/usr/bin/env python3
"""
Walidator klasyfikacji algorytmow w 641 podzadaniach CKE.

Sprawdza:
- Kazdy tag w polu `algorytmy` istnieje w rejestrze
- Rejestr ma poprawna strukture (kategorie, definicje, podstawa)
- Brak duplikatow nazw w rejestrze
- Pokrycie: ile podzadan ma tagi vs puste []

Uzycie:
    python3 analiza/scripts/validate_algorytmy.py
Exit code 0 = OK, 1 = blad walidacji.
"""
from __future__ import annotations
import json
import sys
from pathlib import Path
from collections import Counter

ROOT = Path(__file__).resolve().parent.parent
JSON_DIR = ROOT / "json"
REJESTR_PATH = JSON_DIR / "algorytmy_rejestr.json"

VALID_KATEGORIE = {"klasyczne", "techniki", "struktury", "wzorce"}


def load_rejestr() -> dict:
    with REJESTR_PATH.open(encoding="utf-8") as f:
        return json.load(f)


def validate_rejestr(rejestr: dict) -> list[str]:
    errors: list[str] = []
    if "algorytmy" not in rejestr:
        errors.append("Rejestr: brak klucza 'algorytmy'")
        return errors

    algos = rejestr["algorytmy"]
    if not isinstance(algos, dict):
        errors.append("Rejestr: 'algorytmy' powinno byc dict")
        return errors

    for nazwa, info in algos.items():
        if not isinstance(info, dict):
            errors.append(f"Rejestr [{nazwa}]: info powinno byc dict")
            continue
        kat = info.get("kategoria")
        if kat not in VALID_KATEGORIE:
            errors.append(f"Rejestr [{nazwa}]: nieprawidlowa kategoria '{kat}'")
        if not info.get("definicja"):
            errors.append(f"Rejestr [{nazwa}]: brak definicji")

    # Duplikaty (case-insensitive)
    lower_names = [n.lower() for n in algos.keys()]
    dup_counter = Counter(lower_names)
    for name, count in dup_counter.items():
        if count > 1:
            errors.append(f"Rejestr: duplikat (case-insensitive) '{name}' wystapil {count}x")

    return errors


def validate_subtasks(rejestr_keys: set[str]) -> tuple[list[str], dict]:
    errors: list[str] = []
    stats = {"total": 0, "tagged": 0, "empty": 0, "tag_counts": Counter(), "missing_field": 0}

    matura_files = sorted(JSON_DIR.glob("matura_*.json"))
    if not matura_files:
        errors.append(f"Brak plikow matura_*.json w {JSON_DIR}")
        return errors, stats

    for fpath in matura_files:
        if "indeks" in fpath.name:
            continue
        with fpath.open(encoding="utf-8") as f:
            data = json.load(f)
        for z in data.get("zadania", []):
            for p in z.get("podzadania", []):
                pid = p.get("id", "?")
                stats["total"] += 1
                if "algorytmy" not in p:
                    errors.append(f"{pid}: brak pola 'algorytmy'")
                    stats["missing_field"] += 1
                    continue
                tags = p["algorytmy"]
                if not isinstance(tags, list):
                    errors.append(f"{pid}: 'algorytmy' powinno byc list")
                    continue
                if not tags:
                    stats["empty"] += 1
                else:
                    stats["tagged"] += 1
                for t in tags:
                    if t not in rejestr_keys:
                        errors.append(f"{pid}: tag '{t}' nie istnieje w rejestrze")
                    stats["tag_counts"][t] += 1

    return errors, stats


def main() -> int:
    print(f"Walidator klasyfikacji algorytmow")
    print(f"  Rejestr: {REJESTR_PATH}")
    print(f"  JSON dir: {JSON_DIR}")
    print()

    rejestr = load_rejestr()
    rejestr_errors = validate_rejestr(rejestr)
    if rejestr_errors:
        print("BLEDY REJESTRU:")
        for e in rejestr_errors:
            print(f"  - {e}")
        return 1

    rejestr_keys = set(rejestr["algorytmy"].keys())
    print(f"Rejestr: {len(rejestr_keys)} algorytmow w 4 kategoriach. OK.")
    print()

    errors, stats = validate_subtasks(rejestr_keys)
    if errors:
        print(f"BLEDY KLASYFIKACJI ({len(errors)}):")
        for e in errors[:20]:
            print(f"  - {e}")
        if len(errors) > 20:
            print(f"  ... i {len(errors) - 20} wiecej")
        return 1

    print(f"Klasyfikacja: {stats['total']} podzadan")
    print(f"  z tagami:    {stats['tagged']} ({100*stats['tagged']/stats['total']:.1f}%)")
    print(f"  pustych []:  {stats['empty']} ({100*stats['empty']/stats['total']:.1f}%)")
    print(f"  brak pola:   {stats['missing_field']}")
    print(f"  unikalne tagi w uzyciu: {len(stats['tag_counts'])}/{len(rejestr_keys)}")
    print()
    print(f"Top 10 tagow:")
    for t, n in stats["tag_counts"].most_common(10):
        print(f"  {n:3}x {t}")

    nieuzywane = rejestr_keys - set(stats["tag_counts"].keys())
    if nieuzywane:
        print()
        print(f"Algorytmy w rejestrze NIEUZYWANE ({len(nieuzywane)}):")
        for n in sorted(nieuzywane):
            print(f"  - {n}")

    print()
    print("OK - walidacja zakonczona pomyslnie.")
    return 0


if __name__ == "__main__":
    sys.exit(main())
