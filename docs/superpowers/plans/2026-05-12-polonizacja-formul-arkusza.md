# Polonizacja formuł arkusza — Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Zamienić wszystkie angielskie nazwy formuł arkusza w materiałach dydaktycznych (3 pliki MD dedykowane + 6 mieszanych + 164 plików JSON ćwiczeń + 2 rejestry) na polskie nazwy MS Excel, zachowując progres uczniów przez migrację schema v9.

**Architecture:** Idempotentny skrypt `polonizuj_formuly.py` z mapą EN→PL i deterministyczną listą zamian per linia dla plików mieszanych. Whitelist plików zapobiega edycji SQL/C++. CLI dostaje schema v9 migrację z `INSERT OR REPLACE` dla zachowania FSRS state.

**Tech Stack:** Python 3 (skrypt), Go (CLI migration), SQLite (progress.db), JSON Schema (walidacja).

**Spec:** `docs/superpowers/specs/2026-05-12-polonizacja-formul-arkusza-design.md`

---

## File Structure

### Pliki tworzone (NEW)

| Plik | Cel |
|---|---|
| `analiza/scripts/polonizuj_formuly.py` | Główny skrypt podmian (dry-run + apply, whitelist guard) |
| `analiza/scripts/polonizacja_mapa.json` | Mapa 46 funkcji EN→PL (dane, nie kod) |
| `analiza/scripts/polonizacja_warstwa2.json` | Deterministyczna lista 18 zamian dla 6 plików mieszanych |
| `analiza/scripts/test_polonizuj_formuly.py` | Testy pytest dla skryptu |
| `analiza/scripts/tests/fixtures/` | Fixtures (kopie fragmentów plików do testów) |

### Pliki modyfikowane (MODIFY)

| Plik | Zakres |
|---|---|
| `analiza/szablony/arkusz_formuly.md` | Globalna podmiana (warstwa 1) |
| `analiza/cheatsheets/cheatsheet_arkusz.md` | Globalna podmiana + nowa sekcja "Pułapki" |
| `analiza/rozwiazania_wzorcowe/arkusz_kalkulacyjny.md` | Globalna podmiana (warstwa 1) |
| `analiza/cheatsheets/debug_checklist.md` | 3 zamiany z listy (warstwa 2) |
| `analiza/cheatsheets/przed_egzaminem.md` | 2 zamiany z listy (warstwa 2) |
| `analiza/szablony/wzorce_2015.md` | 8 zamian z listy + separator (warstwa 2) |
| `analiza/PRZEWODNIK_UCZNIA.md` | 1 zamiana z listy (warstwa 2) |
| `analiza/strategia_egzaminacyjna.md` | 3 zamiany z listy (warstwa 2) |
| `analiza/drzewo_decyzyjne.md` | 3 zamiany z listy (warstwa 2) |
| `analiza/cwiczenia/json/15_agregacja_warunkowa/*.json` | 41 plików (40 + _meta), pełna podmiana (warstwa 3) |
| `analiza/cwiczenia/json/16_symulacja/*.json` | 41 plików, pełna podmiana |
| `analiza/cwiczenia/json/18_agregacja_podstawowa/*.json` | 41 plików, pełna podmiana |
| `analiza/cwiczenia/json/19_transformacja/*.json` | 41 plików, pełna podmiana |
| `analiza/cwiczenia/json/tagi_rejestr.json` | Rename 7 tagów |
| `analiza/json/algorytmy_rejestr.json` | 1 linia tekstu opisowego |
| `analiza/cli/database.go` | Schema v9 migration (currentSchemaVersion: 8→9) |
| `analiza/cli/main_test.go` | Test migracji v8→v9 |
| `MEMORY.md` | Notatka o zakończeniu polonizacji |

---

## Faza 0 — Mapy danych (3 zadania)

### Task 1: Utworzenie mapy EN→PL jako JSON

**Files:**
- Create: `analiza/scripts/polonizacja_mapa.json`

- [ ] **Step 1: Utwórz plik z mapą 46 funkcji**

```json
{
  "_metadata": {
    "version": "1.0",
    "data": "2026-05-12",
    "zrodlo": "Microsoft Support PL, zweryfikowane przez 2 niezalezne agenty",
    "konwencja": "MS Excel PL (NIE LibreOffice)"
  },
  "mapowanie": {
    "VLOOKUP": "WYSZUKAJ.PIONOWO",
    "HLOOKUP": "WYSZUKAJ.POZIOMO",
    "IF": "JEŻELI",
    "AND": "ORAZ",
    "OR": "LUB",
    "NOT": "NIE",
    "IFERROR": "JEŻELI.BŁĄD",
    "SUM": "SUMA",
    "SUMIF": "SUMA.JEŻELI",
    "SUMIFS": "SUMA.WARUNKÓW",
    "SUMPRODUCT": "SUMA.ILOCZYNÓW",
    "AVERAGE": "ŚREDNIA",
    "AVERAGEIF": "ŚREDNIA.JEŻELI",
    "AVERAGEIFS": "ŚREDNIA.WARUNKÓW",
    "COUNT": "ILE.LICZB",
    "COUNTA": "ILE.NIEPUSTYCH",
    "COUNTIF": "LICZ.JEŻELI",
    "COUNTIFS": "LICZ.WARUNKI",
    "MAXIFS": "MAKS.WARUNKÓW",
    "MINIFS": "MIN.WARUNKÓW",
    "ROUND": "ZAOKR",
    "ROUNDUP": "ZAOKR.GÓRA",
    "ROUNDDOWN": "ZAOKR.DÓŁ",
    "INT": "ZAOKR.DO.CAŁK",
    "ABS": "MODUŁ.LICZBY",
    "LEFT": "LEWY",
    "RIGHT": "PRAWY",
    "MID": "FRAGMENT.TEKSTU",
    "LEN": "DŁ",
    "TRIM": "USUŃ.ZBĘDNE.ODSTĘPY",
    "UPPER": "LITERY.WIELKIE",
    "LOWER": "LITERY.MAŁE",
    "CONCATENATE": "ZŁĄCZ.TEKSTY",
    "FIND": "ZNAJDŹ",
    "SEARCH": "SZUKAJ.TEKST",
    "SUBSTITUTE": "PODSTAW",
    "REPLACE": "ZASTĄP",
    "INDEX": "INDEKS",
    "MATCH": "PODAJ.POZYCJĘ",
    "RANK": "POZYCJA",
    "YEAR": "ROK",
    "MONTH": "MIESIĄC",
    "DAY": "DZIEŃ",
    "TODAY": "DZIŚ",
    "NOW": "TERAZ",
    "DATE": "DATA"
  },
  "bez_zmian": ["MAX", "MIN", "MOD"],
  "tagi_rejestr_rename": {
    "VLOOKUP": "WYSZUKAJ.PIONOWO",
    "SUMIF": "SUMA.JEŻELI",
    "SUMIFS": "SUMA.WARUNKÓW",
    "COUNTIF": "LICZ.JEŻELI",
    "COUNTIFS": "LICZ.WARUNKI",
    "AVERAGEIF": "ŚREDNIA.JEŻELI",
    "AVERAGEIFS": "ŚREDNIA.WARUNKÓW"
  }
}
```

- [ ] **Step 2: Weryfikuj struktura JSON**

Run: `python3 -c "import json; json.load(open('matura_informatyka_rozszerzona/analiza/scripts/polonizacja_mapa.json'))"`
Expected: brak błędu (poprawny JSON)

- [ ] **Step 3: Verify mapa ma 46 wpisów + 3 bez zmian + 7 tagów**

Run: `python3 -c "import json; m=json.load(open('matura_informatyka_rozszerzona/analiza/scripts/polonizacja_mapa.json')); print(len(m['mapowanie']), len(m['bez_zmian']), len(m['tagi_rejestr_rename']))"`
Expected: `46 3 7`

- [ ] **Step 4: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/scripts/polonizacja_mapa.json
git commit -m "Dodaj mape EN->PL formul Excel (46 funkcji)"
```

---

### Task 2: Deterministyczna lista zamian dla warstwy 2

**Files:**
- Create: `analiza/scripts/polonizacja_warstwa2.json`

- [ ] **Step 1: Utwórz plik z listą 18 zamian**

```json
{
  "_metadata": {
    "data": "2026-05-12",
    "opis": "Deterministyczna lista zamian dla 6 plikow MD mieszanych. Skrypt sprawdza ze STARY istnieje w LINII +-5 zanim zamieni."
  },
  "zamiany": [
    {
      "plik": "matura_informatyka_rozszerzona/analiza/cheatsheets/debug_checklist.md",
      "linia": 92,
      "stary": "- [ ] SUMIF: `=SUMIF(zakres_warunku; warunek; zakres_sumy)` — 3 argumenty!",
      "nowy": "- [ ] SUMA.JEŻELI: `=SUMA.JEŻELI(zakres_warunku; warunek; zakres_sumy)` — 3 argumenty!"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/cheatsheets/debug_checklist.md",
      "linia": 93,
      "stary": "- [ ] SUMIFS: `=SUMIFS(zakres_sumy; zakres_war1; war1; zakres_war2; war2)` — suma PIERWSZA!",
      "nowy": "- [ ] SUMA.WARUNKÓW: `=SUMA.WARUNKÓW(zakres_sumy; zakres_war1; war1; zakres_war2; war2)` — suma PIERWSZA!"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/cheatsheets/debug_checklist.md",
      "linia": 94,
      "stary": "- [ ] COUNTIF: `=COUNTIF(zakres; warunek)` — 2 argumenty",
      "nowy": "- [ ] LICZ.JEŻELI: `=LICZ.JEŻELI(zakres; warunek)` — 2 argumenty"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/cheatsheets/przed_egzaminem.md",
      "linia": 55,
      "stary": "- [ ] `=SUMIF(zakres_war; warunek; zakres_sum)` / SUMIFS",
      "nowy": "- [ ] `=SUMA.JEŻELI(zakres_war; warunek; zakres_sum)` / SUMA.WARUNKÓW"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/cheatsheets/przed_egzaminem.md",
      "linia": 56,
      "stary": "- [ ] `=COUNTIF(zakres; warunek)`",
      "nowy": "- [ ] `=LICZ.JEŻELI(zakres; warunek)`"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md",
      "linia": 159,
      "stary": "=SUMIF(A:A,\"A\",B:B)",
      "nowy": "=SUMA.JEŻELI(A:A;\"A\";B:B)"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md",
      "linia": 162,
      "stary": "=SUMIF(A:A,\"B\",B:B)",
      "nowy": "=SUMA.JEŻELI(A:A;\"B\";B:B)"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md",
      "linia": 168,
      "stary": "=COUNTA(A2:A20)",
      "nowy": "=ILE.NIEPUSTYCH(A2:A20)"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md",
      "linia": 171,
      "stary": "=COUNTIF(A:A,\"A\")",
      "nowy": "=LICZ.JEŻELI(A:A;\"A\")"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md",
      "linia": 183,
      "stary": "=COUNTIF(E:E, \">0\")",
      "nowy": "=LICZ.JEŻELI(E:E;\">0\")"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md",
      "linia": 269,
      "stary": "### Pułapka 8: SUMIF w Excel",
      "nowy": "### Pułapka 8: SUMA.JEŻELI w Excel"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md",
      "linia": 272,
      "stary": "**Błąd:** Użycie SUM zamiast SUMIF dla warunkowego sumowania",
      "nowy": "**Błąd:** Użycie SUMA zamiast SUMA.JEŻELI dla warunkowego sumowania"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md",
      "linia": 273,
      "stary": "**Poprawnie:** =SUMIF(zakres_kryterium, kryterium, zakres_do_zsumowania)",
      "nowy": "**Poprawnie:** =SUMA.JEŻELI(zakres_kryterium; kryterium; zakres_do_zsumowania)"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/PRZEWODNIK_UCZNIA.md",
      "linia": 185,
      "stary": "8. `15_agregacja_warunkowa.md` — arkusz: SUMIFS/COUNTIFS (38 pkt)",
      "nowy": "8. `15_agregacja_warunkowa.md` — arkusz: SUMA.WARUNKÓW/LICZ.WARUNKI (38 pkt)"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/strategia_egzaminacyjna.md",
      "linia": 613,
      "stary": "2. SUMIF(zakres_kryt, kryterium, zakres_sum)",
      "nowy": "2. SUMA.JEŻELI(zakres_kryt; kryterium; zakres_sum)"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/strategia_egzaminacyjna.md",
      "linia": 614,
      "stary": "3. Dla wielu warunkow: SUMIFS (warunki w parach zakres+kryterium)",
      "nowy": "3. Dla wielu warunkow: SUMA.WARUNKÓW (warunki w parach zakres+kryterium)"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/strategia_egzaminacyjna.md",
      "linia": 617,
      "stary": "- Pomylenie SUMIF z SUMIFS (kolejnosc argumentow!)",
      "nowy": "- Pomylenie SUMA.JEŻELI z SUMA.WARUNKÓW (kolejnosc argumentow!)"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/drzewo_decyzyjne.md",
      "linia": 146,
      "stary": "| \"suma/liczba wg warunku\" | SUMIF / COUNTIF | `=SUMIFS(D:D; B:B; \"X\"; C:C; \">100\")` |",
      "nowy": "| \"suma/liczba wg warunku\" | SUMA.JEŻELI / LICZ.JEŻELI | `=SUMA.WARUNKÓW(D:D; B:B; \"X\"; C:C; \">100\")` |"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/drzewo_decyzyjne.md",
      "linia": 147,
      "stary": "| \"srednia wg warunku\" | AVERAGEIF(S) | `=AVERAGEIFS(D:D; B:B; \"X\")` |",
      "nowy": "| \"srednia wg warunku\" | ŚREDNIA.JEŻELI(.WARUNKÓW) | `=ŚREDNIA.WARUNKÓW(D:D; B:B; \"X\")` |"
    },
    {
      "plik": "matura_informatyka_rozszerzona/analiza/drzewo_decyzyjne.md",
      "linia": 148,
      "stary": "| \"zlicz unikalne\" | COUNTIF + pomocnicza | `=1/COUNTIF(zakres; wartosc)` -> SUMA |",
      "nowy": "| \"zlicz unikalne\" | LICZ.JEŻELI + pomocnicza | `=1/LICZ.JEŻELI(zakres; wartosc)` -> SUMA |"
    }
  ]
}
```

- [ ] **Step 2: Weryfikuj struktura JSON i 20 zamian**

Run: `python3 -c "import json; w=json.load(open('matura_informatyka_rozszerzona/analiza/scripts/polonizacja_warstwa2.json')); print(len(w['zamiany']))"`
Expected: `20` (3+2+8+1+3+3)

- [ ] **Step 3: Weryfikuj że każdy STARY rzeczywiście istnieje w odpowiednim pliku**

```bash
python3 << 'EOF'
import json
m = json.load(open('matura_informatyka_rozszerzona/analiza/scripts/polonizacja_warstwa2.json'))
errors = []
for z in m['zamiany']:
    with open(z['plik']) as f:
        content = f.read()
    if z['stary'] not in content:
        errors.append(f"MISSING: {z['plik']}:{z['linia']}: {z['stary'][:60]}...")
if errors:
    for e in errors: print(e)
    exit(1)
print(f"OK: wszystkie 20 fragmentow STARY znalezione w plikach")
EOF
```

Expected: `OK: wszystkie 20 fragmentow STARY znalezione w plikach`

- [ ] **Step 4: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/scripts/polonizacja_warstwa2.json
git commit -m "Dodaj deterministyczna liste 20 zamian dla 6 plikow mieszanych"
```

---

### Task 3: Fixtures dla testów skryptu

**Files:**
- Create: `analiza/scripts/tests/fixtures/fixture_warstwa1.md`
- Create: `analiza/scripts/tests/fixtures/fixture_warstwa2.md`
- Create: `analiza/scripts/tests/fixtures/fixture_cwiczenie.json`
- Create: `analiza/scripts/tests/fixtures/fixture_sql_NOT_TOUCH.md`

- [ ] **Step 1: Fixture dla warstwy 1 (czysty arkusz)**

Create `analiza/scripts/tests/fixtures/fixture_warstwa1.md`:

```markdown
# Test arkusz

## Formuly

=SUM(A1:A10) + IF(B1>0, COUNT(B:B), 0)
=VLOOKUP(C1, $D$1:$E$10; 2; 0)
=SUMIFS(D:D, A:A, "X")

## Tabela

| Funkcja | Opis |
|---|---|
| ROUND(x;2) | Zaokraglenie do 2 miejsc |
| YEAR(data) | Rok z daty |
```

- [ ] **Step 2: Fixture dla warstwy 2 (mieszany)**

Create `analiza/scripts/tests/fixtures/fixture_warstwa2.md`:

```markdown
# Test mieszany

## SQL (NIE TYKAJ)

```sql
SELECT COUNT(*), SUM(price), MAX(date) FROM tab WHERE x IF NULL THEN 0;
```

## C++ (NIE TYKAJ)

```cpp
int sum = 0;
if (x > 0) sum += abs(x);
```

## Arkusz (TYLKO WYBRANE LINIE)

- [ ] `=SUMIF(A:A; "x"; B:B)` — to ma byc zmienione
- [ ] `=COUNTIF(C:C; ">0")` — to ma byc zmienione
```

- [ ] **Step 3: Fixture JSON ćwiczenia**

Create `analiza/scripts/tests/fixtures/fixture_cwiczenie.json`:

```json
{
  "id": "TEST.1",
  "typ_nazwa": "agregacja_warunkowa",
  "punkty": 2,
  "trudnosc": "latwe",
  "tagi": ["SUMIF", "warunek-tekstowy"],
  "tresc": "Uzyj formuly =SUMIF(B2:B11;\"X\";C2:C11)",
  "odpowiedz": "Wynik: =SUMIF(B:B,\"X\",C:C)",
  "wskazowki": [
    {"poziom": 1, "tekst": "SUMIF ma 3 argumenty"}
  ],
  "typowe_bledy": [
    {"kod": "TST.001", "opis": "Pomylenie SUMIF z COUNTIF"}
  ]
}
```

- [ ] **Step 4: Fixture SQL (NIGDY nie powinien być zmieniony)**

Create `analiza/scripts/tests/fixtures/fixture_sql_NOT_TOUCH.md`:

```markdown
# SQL Patterns

```sql
SELECT COUNT(*) FROM users WHERE active = 1;
SELECT SUM(amount), MAX(date), MIN(date) FROM transactions;
SELECT IF(x > 0, 'positive', 'negative') FROM nums;
```
```

- [ ] **Step 5: Commit fixtures**

```bash
git add matura_informatyka_rozszerzona/analiza/scripts/tests/fixtures/
git commit -m "Dodaj fixtures testowe dla skryptu polonizuj_formuly"
```

---

## Faza 1 — Skrypt: szkielet, testy, podmiany (5 zadań)

### Task 4: Szkielet skryptu z whitelist guard

**Files:**
- Create: `analiza/scripts/polonizuj_formuly.py`

- [ ] **Step 1: Utwórz szkielet skryptu z parsowaniem argumentów**

```python
#!/usr/bin/env python3
"""
polonizuj_formuly.py — polonizacja angielskich nazw formul Excel w materialach.

Tryby:
  --dry-run                  Pokaz diff, nic nie zmieniaj
  --apply                    Wykonaj zamiany
  --warstwa {1,2,3,4,5,all}  Wybierz warstwe (1=dedykowane MD, 2=mieszane MD,
                             3=JSON cwiczen, 4=rejestry, 5=cli/database.go, all=wszystkie)
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
    rel = str(Path(path).relative_to(REPO_ROOT)) if Path(path).is_absolute() else path
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


def main():
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    g = parser.add_mutually_exclusive_group(required=True)
    g.add_argument("--dry-run", action="store_true", help="Pokaz diff, nic nie zmieniaj")
    g.add_argument("--apply", action="store_true", help="Wykonaj zamiany")
    parser.add_argument("--warstwa", choices=["1", "2", "3", "4", "5", "all"], default="all")
    args = parser.parse_args()

    mapa = load_mapa()
    print(f"Mapa: {len(mapa['mapowanie'])} funkcji EN->PL")
    print(f"Warstwa: {args.warstwa}, Tryb: {'DRY-RUN' if args.dry_run else 'APPLY'}")
    # Faza implementacji per warstwa - dodane w kolejnych zadaniach
    print("TODO: implementuj logike per warstwa")


if __name__ == "__main__":
    main()
```

- [ ] **Step 2: Spraw że skrypt jest wykonywalny i uruchom**

```bash
chmod +x matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py
matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py --dry-run --warstwa 1
```

Expected output:
```
Mapa: 46 funkcji EN->PL
Warstwa: 1, Tryb: DRY-RUN
TODO: implementuj logike per warstwa
```

- [ ] **Step 3: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py
git commit -m "Szkielet polonizuj_formuly.py: argparser + whitelist"
```

---

### Task 5: Logika podmian warstwy 1 (cały plik) + testy

**Files:**
- Modify: `analiza/scripts/polonizuj_formuly.py`
- Create: `analiza/scripts/test_polonizuj_formuly.py`

- [ ] **Step 1: Napisz test (czerwony) dla zamiany w warstwie 1**

Create `analiza/scripts/test_polonizuj_formuly.py`:

```python
"""Testy dla polonizuj_formuly.py."""
import json
import shutil
import sys
from pathlib import Path

import pytest

SCRIPTS_DIR = Path(__file__).resolve().parent
sys.path.insert(0, str(SCRIPTS_DIR))
import polonizuj_formuly as pf

FIXTURES = SCRIPTS_DIR / "tests" / "fixtures"


def test_warstwa1_polonizuje_caly_plik_md(tmp_path):
    """Warstwa 1: globalna podmiana wszystkich formul w pliku."""
    src = FIXTURES / "fixture_warstwa1.md"
    dst = tmp_path / "fixture.md"
    shutil.copy(src, dst)

    mapa = pf.load_mapa()
    changes = pf.polonizuj_md_warstwa1(dst, mapa)

    content = dst.read_text()
    # Sprawdz konkretne zamiany
    assert "=SUMA(A1:A10)" in content
    assert "JEŻELI(B1>0" in content
    assert "ILE.LICZB(B:B)" in content
    assert "WYSZUKAJ.PIONOWO(C1" in content
    assert "SUMA.WARUNKÓW(D:D" in content
    assert "ZAOKR(x" in content
    assert "ROK(data)" in content
    # Nie zostawia angielskich nazw
    assert "SUM(" not in content
    assert "IF(" not in content
    assert "COUNT(" not in content
    assert "VLOOKUP(" not in content
    assert "SUMIFS(" not in content
    assert "ROUND(" not in content
    assert "YEAR(" not in content
    assert changes > 0
```

- [ ] **Step 2: Uruchom test żeby zobaczyć że failuje**

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: FAIL (`polonizuj_md_warstwa1` nie istnieje)

- [ ] **Step 3: Zaimplementuj `polonizuj_md_warstwa1` w skrypcie**

Dodaj do `polonizuj_formuly.py` przed funkcją `main()`:

```python
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
```

- [ ] **Step 4: Uruchom test żeby zobaczyć że przechodzi**

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: PASS

- [ ] **Step 5: Dodaj test "Tylko nazwy bez nawiasów NIE są zmieniane"**

Dodaj do `test_polonizuj_formuly.py`:

```python
def test_warstwa1_nie_rusza_nazw_bez_nawiasow(tmp_path):
    """Funkcja SUMA jest największa — bez nawiasu zostaje."""
    f = tmp_path / "test.md"
    f.write_text("Tekst o funkcji SUM jest dluzszy. Ale =SUM(A:A) tak.")
    mapa = pf.load_mapa()
    pf.polonizuj_md_warstwa1(f, mapa)
    content = f.read_text()
    assert "funkcji SUM jest" in content  # bez nawiasu — bez zmian
    assert "=SUMA(A:A)" in content        # z nawiasem — zmiana
```

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: 2 PASS

- [ ] **Step 6: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py matura_informatyka_rozszerzona/analiza/scripts/test_polonizuj_formuly.py
git commit -m "Dodaj polonizuj_md_warstwa1 + testy (globalna podmiana 46 funkcji)"
```

---

### Task 6: Logika podmian warstwy 2 (deterministyczna lista) + testy

**Files:**
- Modify: `analiza/scripts/polonizuj_formuly.py`
- Modify: `analiza/scripts/test_polonizuj_formuly.py`

- [ ] **Step 1: Napisz test (czerwony) dla deterministycznej zamiany**

Dodaj do `test_polonizuj_formuly.py`:

```python
def test_warstwa2_stosuje_tylko_liste_deterministyczna(tmp_path):
    """Warstwa 2: zamienia tylko fragmenty z polonizacja_warstwa2.json."""
    src = FIXTURES / "fixture_warstwa2.md"
    dst = tmp_path / "fixture.md"
    shutil.copy(src, dst)

    # Zbuduj fake liste zamian wskazujaca na nasz fixture
    zamiany = [
        {
            "plik": str(dst),
            "linia": 16,
            "stary": "- [ ] `=SUMIF(A:A; \"x\"; B:B)` — to ma byc zmienione",
            "nowy": "- [ ] `=SUMA.JEŻELI(A:A; \"x\"; B:B)` — to ma byc zmienione"
        },
        {
            "plik": str(dst),
            "linia": 17,
            "stary": "- [ ] `=COUNTIF(C:C; \">0\")` — to ma byc zmienione",
            "nowy": "- [ ] `=LICZ.JEŻELI(C:C; \">0\")` — to ma byc zmienione"
        }
    ]
    changes = pf.polonizuj_md_warstwa2(zamiany)

    content = dst.read_text()
    # SQL i C++ bloki NIETKNIETE
    assert "SELECT COUNT(*), SUM(price), MAX(date)" in content
    assert "if (x > 0) sum += abs(x);" in content
    # Tylko 2 wyliczone linie zmienione
    assert "=SUMA.JEŻELI(A:A;" in content
    assert "=LICZ.JEŻELI(C:C;" in content
    assert "=SUMIF(" not in content
    assert "=COUNTIF(" not in content
    assert changes == 2
```

- [ ] **Step 2: Uruchom test (powinien fail)**

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py::test_warstwa2_stosuje_tylko_liste_deterministyczna -v`
Expected: FAIL

- [ ] **Step 3: Zaimplementuj `polonizuj_md_warstwa2`**

Dodaj do `polonizuj_formuly.py`:

```python
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
```

- [ ] **Step 4: Uruchom test (powinien przejść)**

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: 3 PASS

- [ ] **Step 5: Test odporności — gdy stary fragment nie istnieje, błąd**

Dodaj do `test_polonizuj_formuly.py`:

```python
def test_warstwa2_przerwa_z_bledem_gdy_stary_brak(tmp_path):
    """Jesli STARY fragment nie istnieje w pliku — exit code 2."""
    f = tmp_path / "test.md"
    f.write_text("Inna tresc bez tego fragmentu")
    zamiany = [{"plik": str(f), "linia": 1, "stary": "NIEISTNIEJACE", "nowy": "X"}]
    with pytest.raises(SystemExit) as exc:
        pf.polonizuj_md_warstwa2(zamiany)
    assert exc.value.code == 2
```

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: 4 PASS

- [ ] **Step 6: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py matura_informatyka_rozszerzona/analiza/scripts/test_polonizuj_formuly.py
git commit -m "Dodaj polonizuj_md_warstwa2 + testy (deterministyczna lista + pre-check)"
```

---

### Task 7: Logika podmian warstwy 3 (JSON ćwiczeń) + testy

**Files:**
- Modify: `analiza/scripts/polonizuj_formuly.py`
- Modify: `analiza/scripts/test_polonizuj_formuly.py`

- [ ] **Step 1: Napisz test JSON ćwiczenia**

Dodaj do `test_polonizuj_formuly.py`:

```python
def test_warstwa3_polonizuje_json_cwiczenia(tmp_path):
    """Warstwa 3: zamienia formuly w tresc/odpowiedz/wskazowki/typowe_bledy + tagi."""
    src = FIXTURES / "fixture_cwiczenie.json"
    dst = tmp_path / "test.json"
    shutil.copy(src, dst)
    mapa = pf.load_mapa()
    changes = pf.polonizuj_json_cwiczenie(dst, mapa)

    data = json.loads(dst.read_text())
    assert "SUMA.JEŻELI(B2:B11" in data["tresc"]
    assert "SUMA.JEŻELI(B:B" in data["odpowiedz"]
    assert "SUMA.JEŻELI ma 3 argumenty" in data["wskazowki"][0]["tekst"]
    assert "Pomylenie SUMA.JEŻELI z LICZ.JEŻELI" in data["typowe_bledy"][0]["opis"]
    assert "SUMA.JEŻELI" in data["tagi"]
    assert "SUMIF" not in data["tagi"]
    assert changes > 0


def test_warstwa3_polonizuje_meta_json(tmp_path):
    """_meta.json: tagi_globalne tez rename'owane."""
    f = tmp_path / "_meta.json"
    f.write_text(json.dumps({
        "tagi_globalne": ["SUMIF", "SUMIFS", "AVERAGEIFS", "warunek-liczbowy"]
    }))
    mapa = pf.load_mapa()
    pf.polonizuj_json_meta(f, mapa)
    data = json.loads(f.read_text())
    assert data["tagi_globalne"] == ["SUMA.JEŻELI", "SUMA.WARUNKÓW", "ŚREDNIA.WARUNKÓW", "warunek-liczbowy"]
```

- [ ] **Step 2: Uruchom (powinno fail)**

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: 2 FAIL

- [ ] **Step 3: Zaimplementuj `polonizuj_json_cwiczenie` i `polonizuj_json_meta`**

Dodaj do `polonizuj_formuly.py`:

```python
def _polonizuj_string(text: str, mapa: dict) -> tuple[str, int]:
    """Wewnetrzna helper: podmiana w stringu, sortuje keys po dlugosci desc."""
    keys = sorted(mapa["mapowanie"].keys(), key=len, reverse=True)
    total = 0
    for en in keys:
        pl = mapa["mapowanie"][en]
        pattern = r"\b" + re.escape(en) + r"(\s*\()"
        text, count = re.subn(pattern, pl + r"\1", text)
        total += count
    return text, total


def _rename_tag(tag: str, tagi_map: dict) -> str:
    return tagi_map.get(tag, tag)


def polonizuj_json_cwiczenie(path: Path, mapa: dict) -> int:
    """Polonizuje pola tekstowe oraz `tagi` w pliku cwiczenia JSON."""
    data = json.loads(path.read_text())
    total = 0

    # Pola tekstowe (rekurencyjnie po listach obiektow)
    for field in ("tresc", "odpowiedz"):
        if field in data and isinstance(data[field], str):
            new, c = _polonizuj_string(data[field], mapa)
            data[field] = new
            total += c

    for field in ("wskazowki", "typowe_bledy"):
        if field in data and isinstance(data[field], list):
            for item in data[field]:
                for k in ("tekst", "opis"):
                    if k in item and isinstance(item[k], str):
                        new, c = _polonizuj_string(item[k], mapa)
                        item[k] = new
                        total += c

    if "weryfikacja_szczegolowa" in data and isinstance(data["weryfikacja_szczegolowa"], str):
        new, c = _polonizuj_string(data["weryfikacja_szczegolowa"], mapa)
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
    """Polonizuje pole `tagi_globalne` w _meta.json. Zachowuje sortowanie."""
    data = json.loads(path.read_text())
    if "tagi_globalne" not in data:
        return 0
    tagi_map = mapa["tagi_rejestr_rename"]
    new = sorted({_rename_tag(t, tagi_map) for t in data["tagi_globalne"]})
    if new != data["tagi_globalne"]:
        data["tagi_globalne"] = new
        path.write_text(json.dumps(data, ensure_ascii=False, indent=2) + "\n")
        return 1
    return 0
```

- [ ] **Step 4: Uruchom testy**

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: 6 PASS

- [ ] **Step 5: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py matura_informatyka_rozszerzona/analiza/scripts/test_polonizuj_formuly.py
git commit -m "Dodaj polonizuj_json_cwiczenie + polonizuj_json_meta + testy"
```

---

### Task 8: Logika rejestrów (warstwa 4) + normalizacja separatora

**Files:**
- Modify: `analiza/scripts/polonizuj_formuly.py`
- Modify: `analiza/scripts/test_polonizuj_formuly.py`

- [ ] **Step 1: Test rejestru tagów**

Dodaj do `test_polonizuj_formuly.py`:

```python
def test_warstwa4_rename_w_tagi_rejestr(tmp_path):
    """tagi_rejestr.json: rename 7 specyficznych wpisow."""
    f = tmp_path / "tagi_rejestr.json"
    f.write_text(json.dumps({
        "_meta": "central tag registry",
        "tagi": ["AVERAGEIF", "AVERAGEIFS", "COUNTIF", "COUNTIFS", "SUMIF", "SUMIFS", "VLOOKUP",
                 "warunek-tekstowy", "warunek-liczbowy", "JOIN", "GROUP_BY"]
    }))
    mapa = pf.load_mapa()
    changes = pf.polonizuj_tagi_rejestr(f, mapa)
    data = json.loads(f.read_text())
    # 7 starych tagow nie istnieje
    for stary in ["AVERAGEIF", "AVERAGEIFS", "COUNTIF", "COUNTIFS", "SUMIF", "SUMIFS", "VLOOKUP"]:
        assert stary not in data["tagi"]
    # 7 nowych tagow istnieje
    for nowy in ["ŚREDNIA.JEŻELI", "ŚREDNIA.WARUNKÓW", "LICZ.JEŻELI", "LICZ.WARUNKI",
                 "SUMA.JEŻELI", "SUMA.WARUNKÓW", "WYSZUKAJ.PIONOWO"]:
        assert nowy in data["tagi"]
    # Nietkniete: warunek-*, JOIN, GROUP_BY
    assert "warunek-tekstowy" in data["tagi"]
    assert "JOIN" in data["tagi"]
    assert "GROUP_BY" in data["tagi"]
    assert changes == 7


def test_separator_normalizuje_przecinek_w_formule(tmp_path):
    """Separator argumentow ',' -> ';' w formulach arkusza."""
    f = tmp_path / "test.md"
    # 3,14 to liczba dziesiętna (zostaje); reszta to separatory
    f.write_text("Test: =SUMA(A:A, \"X\", B:B) oraz =ZAOKR(3,14; 2)")
    pf.normalizuj_separator(f)
    content = f.read_text()
    assert "=SUMA(A:A; \"X\"; B:B)" in content
    assert "=ZAOKR(3,14; 2)" in content  # 3,14 zachowane
```

- [ ] **Step 2: Uruchom (fail)**

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: 2 FAIL

- [ ] **Step 3: Zaimplementuj `polonizuj_tagi_rejestr` i `normalizuj_separator`**

Dodaj do `polonizuj_formuly.py`:

```python
def polonizuj_tagi_rejestr(path: Path, mapa: dict) -> int:
    """Rename 7 wpisow w tagi_rejestr.json (centralna lista tagow)."""
    data = json.loads(path.read_text())
    if "tagi" not in data:
        return 0
    tagi_map = mapa["tagi_rejestr_rename"]
    new = sorted({_rename_tag(t, tagi_map) for t in data["tagi"]})
    changes = sum(1 for old in data["tagi"] if old in tagi_map)
    data["tagi"] = new
    path.write_text(json.dumps(data, ensure_ascii=False, indent=2) + "\n")
    return changes


def polonizuj_algorytmy_rejestr(path: Path) -> int:
    """1 linia tekstu opisowego w algorytmy_rejestr.json."""
    content = path.read_text()
    stary = "Arkusz: SUMIFS, COUNTIFS, AVERAGEIFS, MAXIFS itp. - agregacja z warunkami."
    nowy = "Arkusz: SUMA.WARUNKÓW, LICZ.WARUNKI, ŚREDNIA.WARUNKÓW, MAKS.WARUNKÓW itp. - agregacja z warunkami."
    if stary not in content:
        return 0
    content = content.replace(stary, nowy, 1)
    path.write_text(content)
    return 1


def normalizuj_separator(path: Path) -> int:
    """Zamienia ',' -> ';' tylko jako separator argumentow formul arkusza.

    Heurystyka: w obrebie formuly arkusza (between `=NAZWA(` i matching `)`),
    przecinek jest separatorem argumentow jesli NIE jest otoczony cyframi
    z obu stron (np. 3,14 = liczba dziesietna, zostaje).
    """
    content = path.read_text()
    original = content

    def replace_in_formula(match):
        body = match.group(1)
        # Zamien przecinek na srednik jesli nie jest otoczony cyframi
        # (?<!\d)  — przed przecinkiem nie cyfra
        # (?!\d)   — po przecinku nie cyfra
        new_body = re.sub(r"(?<!\d),(?!\d)", ";", body)
        # Cleanup: ";<space>" -> "; "
        return "=" + match.group(2) + "(" + new_body + ")"

    # Wzorzec: =NAZWA(zawartosc)  — nieliterackie zagniezdzanie nawiasow!
    # Konserwatywnie: tylko pojedynczy poziom nawiasow.
    # Lapie =NAZWA( ... ) gdzie ... nie ma '('
    pattern = r"=([A-ZŁŻŚĆŃÓĄĘŹ.]+)\(([^()]*)\)"
    content = re.sub(pattern, replace_in_formula, content)

    if content != original:
        path.write_text(content)
        return 1
    return 0
```

- [ ] **Step 4: Uruchom testy**

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: 8 PASS

- [ ] **Step 5: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py matura_informatyka_rozszerzona/analiza/scripts/test_polonizuj_formuly.py
git commit -m "Dodaj polonizuj_tagi_rejestr + algorytmy_rejestr + normalizuj_separator"
```

---

### Task 9: Whitelist guard test + dispatcher main()

**Files:**
- Modify: `analiza/scripts/polonizuj_formuly.py`
- Modify: `analiza/scripts/test_polonizuj_formuly.py`

- [ ] **Step 1: Test whitelist guard**

Dodaj do `test_polonizuj_formuly.py`:

```python
def test_whitelist_guard_odmawia_edycji_pliku_sql(tmp_path):
    """Skrypt nie moze edytowac plikow spoza whitelisty."""
    src = FIXTURES / "fixture_sql_NOT_TOUCH.md"
    dst = tmp_path / "fixture_sql.md"
    shutil.copy(src, dst)
    original = dst.read_text()
    # Skrypt powinien sprawdzic czy plik jest na whitelist
    assert not pf.is_whitelisted(str(dst), "all")
```

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: 9 PASS

- [ ] **Step 2: Zaktualizuj `main()` w skrypcie żeby uruchamiał odpowiednią warstwę**

Zastąp `main()` w `polonizuj_formuly.py`:

```python
def run_warstwa1(mapa, apply: bool):
    total = 0
    for rel in WHITELIST_WARSTWA_1:
        path = REPO_ROOT / rel
        if not path.exists():
            print(f"SKIP: {rel} (nie istnieje)")
            continue
        if not apply:
            # Dry-run: podlicz potencjalne zmiany
            count = sum(len(re.findall(r"\b" + re.escape(en) + r"\s*\(", path.read_text()))
                        for en in mapa["mapowanie"])
            print(f"DRY-RUN {rel}: {count} potencjalnych zamian")
            total += count
        else:
            count = polonizuj_md_warstwa1(path, mapa)
            normalizuj_separator(path)
            print(f"APPLIED {rel}: {count} zamian")
            total += count
    print(f"Warstwa 1: {total} zmian")
    return total


def run_warstwa2(apply: bool):
    w2 = load_warstwa2()
    if not apply:
        print(f"DRY-RUN warstwa 2: {len(w2['zamiany'])} planowanych zamian")
        return len(w2["zamiany"])
    total = polonizuj_md_warstwa2(w2["zamiany"])
    # Po zamianach: normalizacja separatora tylko w plikach warstwy 2
    for plik in {z["plik"] for z in w2["zamiany"]}:
        normalizuj_separator(Path(plik))
    print(f"Warstwa 2: {total} zamian")
    return total


def run_warstwa3(mapa, apply: bool):
    total = 0
    for d in WHITELIST_WARSTWA_3_DIRS:
        dpath = REPO_ROOT / d
        if not dpath.exists():
            print(f"SKIP: {d}")
            continue
        for json_file in sorted(dpath.glob("*.json")):
            if json_file.name == "_meta.json":
                if apply:
                    total += polonizuj_json_meta(json_file, mapa)
            else:
                if apply:
                    total += polonizuj_json_cwiczenie(json_file, mapa)
    print(f"Warstwa 3: {total} zmian")
    return total


def run_warstwa4(mapa, apply: bool):
    total = 0
    for rel in WHITELIST_WARSTWA_4:
        path = REPO_ROOT / rel
        if not path.exists():
            continue
        if not apply:
            print(f"DRY-RUN {rel}")
            continue
        if "tagi_rejestr" in rel:
            total += polonizuj_tagi_rejestr(path, mapa)
        elif "algorytmy_rejestr" in rel:
            total += polonizuj_algorytmy_rejestr(path)
    print(f"Warstwa 4: {total} zmian")
    return total


def main():
    parser = argparse.ArgumentParser(description=__doc__, formatter_class=argparse.RawDescriptionHelpFormatter)
    g = parser.add_mutually_exclusive_group(required=True)
    g.add_argument("--dry-run", action="store_true")
    g.add_argument("--apply", action="store_true")
    parser.add_argument("--warstwa", choices=["1", "2", "3", "4", "all"], default="all")
    args = parser.parse_args()

    mapa = load_mapa()
    print(f"Mapa: {len(mapa['mapowanie'])} funkcji EN->PL")
    print(f"Warstwa: {args.warstwa}, Tryb: {'DRY-RUN' if args.dry_run else 'APPLY'}")

    apply = args.apply
    grand_total = 0
    if args.warstwa in ("1", "all"):
        grand_total += run_warstwa1(mapa, apply)
    if args.warstwa in ("2", "all"):
        grand_total += run_warstwa2(apply)
    if args.warstwa in ("3", "all"):
        grand_total += run_warstwa3(mapa, apply)
    if args.warstwa in ("4", "all"):
        grand_total += run_warstwa4(mapa, apply)

    print(f"\nTOTAL: {grand_total} zmian")
```

- [ ] **Step 3: Dry-run całość — pokaż liczby**

Run: `matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py --dry-run --warstwa all`
Expected: wypisuje liczby per warstwa, nie modyfikuje plików.

- [ ] **Step 4: Verify że nic się nie zmieniło**

Run: `git diff --stat matura_informatyka_rozszerzona/analiza/`
Expected: brak zmian w plikach analiza/ (poza tworzonymi przez nas w scripts/)

- [ ] **Step 5: Uruchom wszystkie testy**

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: 9 PASS

- [ ] **Step 6: Commit**

```bash
git add matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py matura_informatyka_rozszerzona/analiza/scripts/test_polonizuj_formuly.py
git commit -m "Dispatcher main() per warstwa + dry-run/apply tryby"
```

---

## Faza 2 — Aplikacja podmian (4 zadania, każde z osobnym commitem)

### Task 10: Apply warstwa 1 (3 pliki MD dedykowane)

**Files:**
- Modify: `analiza/szablony/arkusz_formuly.md`
- Modify: `analiza/cheatsheets/cheatsheet_arkusz.md`
- Modify: `analiza/rozwiazania_wzorcowe/arkusz_kalkulacyjny.md`

- [ ] **Step 1: Dry-run warstwy 1**

Run: `matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py --dry-run --warstwa 1`
Expected: wypisuje ~140 potencjalnych zamian (15 SUMIFS + 6 SUMIF + 4 COUNTIFS + 4 COUNTIF + 14 SUM + 15 IF + ... ze szczegółowego audytu)

- [ ] **Step 2: Apply warstwa 1**

Run: `matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py --apply --warstwa 1`
Expected: ~140 zmian, "Warstwa 1: 140 zmian"

- [ ] **Step 3: Grep guard — żadnych angielskich nazw spreadsheet-specific**

Run:
```bash
grep -E '\b(VLOOKUP|HLOOKUP|SUMIF|SUMIFS|COUNTIF|COUNTIFS|AVERAGEIF|AVERAGEIFS|COUNTA|CONCATENATE|IFERROR|SUMPRODUCT|MAXIFS|MINIFS)\s*\(' \
  matura_informatyka_rozszerzona/analiza/szablony/arkusz_formuly.md \
  matura_informatyka_rozszerzona/analiza/cheatsheets/cheatsheet_arkusz.md \
  matura_informatyka_rozszerzona/analiza/rozwiazania_wzorcowe/arkusz_kalkulacyjny.md
```
Expected: brak wyników (exit 1 z grep = OK)

- [ ] **Step 4: Spot-check ręczny — otwórz i przejrzyj 5 linii każdego pliku**

Run: `head -50 matura_informatyka_rozszerzona/analiza/cheatsheets/cheatsheet_arkusz.md | grep -E '(=|JEŻELI|SUMA|LICZ|WYSZUKAJ)'`
Expected: widać polskie nazwy

- [ ] **Step 5: Commit warstwa 1**

```bash
git add matura_informatyka_rozszerzona/analiza/szablony/arkusz_formuly.md \
        matura_informatyka_rozszerzona/analiza/cheatsheets/cheatsheet_arkusz.md \
        matura_informatyka_rozszerzona/analiza/rozwiazania_wzorcowe/arkusz_kalkulacyjny.md
git commit -m "Polonizacja warstwa 1: 3 pliki MD dedykowane arkuszowi"
```

---

### Task 11: Apply warstwa 3 (164 plików JSON ćwiczeń arkuszowych)

**Files:**
- Modify: 4 katalogi po 41 plików (40 ćwiczeń + _meta.json)

- [ ] **Step 1: Apply warstwa 3**

Run: `matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py --apply --warstwa 3`
Expected: ~200+ zmian (głównie w tresc/odpowiedz/tagi)

- [ ] **Step 2: Walidator schema (krytyczne — czy tagi są nadal w rejestrze?)**

Run: `python3 matura_informatyka_rozszerzona/analiza/cwiczenia/validate_json.py`
Expected: WYSTAPIA BLEDY — bo zmieniliśmy tagi w ćwiczeniach, ale `tagi_rejestr.json` jeszcze nie zaktualizowany. To OK — naprawimy w Task 12.

- [ ] **Step 3: Sprawdź ile błędów (powinno być ograniczone do 7 tagów)**

Run: `python3 matura_informatyka_rozszerzona/analiza/cwiczenia/validate_json.py 2>&1 | grep -E 'SUMA|LICZ|WYSZUKAJ|ŚREDNIA' | wc -l`
Expected: liczba błędów == liczba użyć 7 nowych polskich tagów (kilkadziesiąt)

- [ ] **Step 4: Spot-check — otwórz losowe ćwiczenie**

Run: `python3 -c "import json; print(json.dumps(json.load(open('matura_informatyka_rozszerzona/analiza/cwiczenia/json/15_agregacja_warunkowa/15.1.json')), ensure_ascii=False, indent=2))" | head -40`
Expected: widać polskie nazwy formuł i polskie tagi

- [ ] **Step 5: Commit warstwa 3 (z notą że validate failuje przejściowo)**

```bash
git add matura_informatyka_rozszerzona/analiza/cwiczenia/json/15_agregacja_warunkowa/ \
        matura_informatyka_rozszerzona/analiza/cwiczenia/json/16_symulacja/ \
        matura_informatyka_rozszerzona/analiza/cwiczenia/json/18_agregacja_podstawowa/ \
        matura_informatyka_rozszerzona/analiza/cwiczenia/json/19_transformacja/
git commit -m "Polonizacja warstwa 3: 164 JSON cwiczen + 4 _meta (formuly + tagi)

UWAGA: validate_json.py failuje przejsciowo — naprawione w nastepnym
commicie (warstwa 4 = rename w tagi_rejestr.json)."
```

---

### Task 12: Apply warstwa 4 (rejestry)

**Files:**
- Modify: `analiza/cwiczenia/json/tagi_rejestr.json`
- Modify: `analiza/json/algorytmy_rejestr.json`

- [ ] **Step 1: Apply warstwa 4**

Run: `matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py --apply --warstwa 4`
Expected: 8 zmian (7 w tagi_rejestr + 1 w algorytmy_rejestr)

- [ ] **Step 2: Walidator schema TERAZ powinien przejść**

Run: `python3 matura_informatyka_rozszerzona/analiza/cwiczenia/validate_json.py`
Expected: 0 ERRORS

- [ ] **Step 3: Verify rejestru tagów**

Run: `python3 -c "import json; t=json.load(open('matura_informatyka_rozszerzona/analiza/cwiczenia/json/tagi_rejestr.json')); old=['VLOOKUP','SUMIF','SUMIFS','COUNTIF','COUNTIFS','AVERAGEIF','AVERAGEIFS']; new=['WYSZUKAJ.PIONOWO','SUMA.JEŻELI','SUMA.WARUNKÓW','LICZ.JEŻELI','LICZ.WARUNKI','ŚREDNIA.JEŻELI','ŚREDNIA.WARUNKÓW']; print('STARE w rejestrze:', sum(1 for x in old if x in t['tagi'])); print('NOWE w rejestrze:', sum(1 for x in new if x in t['tagi']))"`
Expected: `STARE w rejestrze: 0` and `NOWE w rejestrze: 7`

- [ ] **Step 4: Commit warstwa 4**

```bash
git add matura_informatyka_rozszerzona/analiza/cwiczenia/json/tagi_rejestr.json \
        matura_informatyka_rozszerzona/analiza/json/algorytmy_rejestr.json
git commit -m "Polonizacja warstwa 4: rename 7 tagow + opis w algorytmy_rejestr"
```

---

### Task 13: Apply warstwa 2 (6 plików MD mieszanych — deterministyczna lista)

**Files:**
- Modify: 6 plików MD wymienionych w polonizacja_warstwa2.json

- [ ] **Step 1: Pre-check — sprawdź że wszystkie STARY istnieją**

Run:
```bash
python3 << 'EOF'
import json
m = json.load(open('matura_informatyka_rozszerzona/analiza/scripts/polonizacja_warstwa2.json'))
errors = []
for z in m['zamiany']:
    with open(z['plik']) as f:
        if z['stary'] not in f.read():
            errors.append(f"{z['plik']}:{z['linia']}")
print(f"Sprawdzono {len(m['zamiany'])} zamian. Bledy: {len(errors)}")
for e in errors: print(f"  BRAK: {e}")
EOF
```
Expected: `Sprawdzono 20 zamian. Bledy: 0`

- [ ] **Step 2: Apply warstwa 2**

Run: `matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py --apply --warstwa 2`
Expected: `Warstwa 2: 20 zamian`

- [ ] **Step 3: Grep guard dla 6 plików mieszanych**

Run:
```bash
grep -E '\b(VLOOKUP|HLOOKUP|SUMIF|SUMIFS|COUNTIF|COUNTIFS|AVERAGEIF|AVERAGEIFS|COUNTA|CONCATENATE|IFERROR|SUMPRODUCT|MAXIFS|MINIFS)\s*\(' \
  matura_informatyka_rozszerzona/analiza/cheatsheets/debug_checklist.md \
  matura_informatyka_rozszerzona/analiza/cheatsheets/przed_egzaminem.md \
  matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md \
  matura_informatyka_rozszerzona/analiza/PRZEWODNIK_UCZNIA.md \
  matura_informatyka_rozszerzona/analiza/strategia_egzaminacyjna.md \
  matura_informatyka_rozszerzona/analiza/drzewo_decyzyjne.md
```
Expected: brak wyników

- [ ] **Step 4: Spot check — SQL/C++ bloki NIETKNIETE**

Run: `grep -E 'SELECT|JOIN|GROUP BY|#include|ifstream' matura_informatyka_rozszerzona/analiza/strategia_egzaminacyjna.md | head -5`
Expected: bloki SQL/C++ widać niezmienione

- [ ] **Step 5: Commit warstwa 2**

```bash
git add matura_informatyka_rozszerzona/analiza/cheatsheets/debug_checklist.md \
        matura_informatyka_rozszerzona/analiza/cheatsheets/przed_egzaminem.md \
        matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md \
        matura_informatyka_rozszerzona/analiza/PRZEWODNIK_UCZNIA.md \
        matura_informatyka_rozszerzona/analiza/strategia_egzaminacyjna.md \
        matura_informatyka_rozszerzona/analiza/drzewo_decyzyjne.md
git commit -m "Polonizacja warstwa 2: 20 deterministycznych zamian w 6 plikach mieszanych"
```

---

## Faza 3 — CLI Migration v9 (3 zadania)

### Task 14: Test migracji v8→v9 w main_test.go

**Files:**
- Modify: `analiza/cli/main_test.go`

- [ ] **Step 1: Sprawdź gdzie są testy migracji**

Run: `grep -nE 'TestMigration|TestSchema|migrate' matura_informatyka_rozszerzona/analiza/cli/main_test.go | head -10`

- [ ] **Step 2: Dodaj test migracji v9 (czerwony — bo migracja jeszcze nie zaimplementowana)**

Append do `matura_informatyka_rozszerzona/analiza/cli/main_test.go`:

```go
func TestMigrationV9RenameTagi(t *testing.T) {
    tmpDir := t.TempDir()
    dbPath := filepath.Join(tmpDir, "matura_progress.db")

    // 1. Stworz DB w v8 ze starymi tagami
    db, err := sql.Open("sqlite", dbPath)
    if err != nil { t.Fatal(err) }
    schema := `CREATE TABLE schema_version (version INTEGER PRIMARY KEY);
               INSERT INTO schema_version VALUES (8);
               CREATE TABLE progress_tagi (
                 tag TEXT PRIMARY KEY, poziom INTEGER DEFAULT 0,
                 nastepna_powtorka TEXT, stability REAL DEFAULT 0,
                 difficulty REAL DEFAULT 5.0, lapses INTEGER DEFAULT 0,
                 reps INTEGER DEFAULT 0, state INTEGER DEFAULT 0, last_review TEXT
               );
               INSERT INTO progress_tagi(tag, stability, reps) VALUES
                 ('VLOOKUP', 5.0, 10),
                 ('SUMIF', 3.0, 5),
                 ('JOIN', 2.0, 3);`
    if _, err := db.Exec(schema); err != nil { t.Fatal(err) }
    db.Close()

    // 2. Otworz przez OpenDB — powinno trigger migrate v8 -> v9
    db, _, err = OpenDB(tmpDir)
    if err != nil { t.Fatal(err) }
    defer db.Close()

    // 3. Verify: VLOOKUP -> WYSZUKAJ.PIONOWO, SUMIF -> SUMA.JEŻELI, JOIN nietkniete
    var stability float64
    var reps int
    err = db.QueryRow("SELECT stability, reps FROM progress_tagi WHERE tag = ?", "WYSZUKAJ.PIONOWO").Scan(&stability, &reps)
    if err != nil { t.Fatalf("WYSZUKAJ.PIONOWO not found: %v", err) }
    if stability != 5.0 || reps != 10 { t.Errorf("FSRS state lost: got stability=%v reps=%v want 5.0/10", stability, reps) }

    err = db.QueryRow("SELECT stability, reps FROM progress_tagi WHERE tag = ?", "SUMA.JEŻELI").Scan(&stability, &reps)
    if err != nil { t.Fatalf("SUMA.JEŻELI not found: %v", err) }
    if stability != 3.0 || reps != 5 { t.Errorf("FSRS state lost: got stability=%v reps=%v want 3.0/5", stability, reps) }

    // Stary VLOOKUP nie istnieje
    var count int
    db.QueryRow("SELECT COUNT(*) FROM progress_tagi WHERE tag IN ('VLOOKUP','SUMIF')").Scan(&count)
    if count != 0 { t.Errorf("Stare tagi nadal w bazie: %d", count) }

    // JOIN nietkniete
    db.QueryRow("SELECT COUNT(*) FROM progress_tagi WHERE tag = 'JOIN'").Scan(&count)
    if count != 1 { t.Errorf("Tag JOIN powinien zostac: %d", count) }
}
```

- [ ] **Step 3: Uruchom test — powinien fail (migracja v9 nie istnieje)**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestMigrationV9RenameTagi -v`
Expected: FAIL (`progress DB v8 niekompatybilny z v8` lub `WYSZUKAJ.PIONOWO not found`)

- [ ] **Step 4: Commit test (failujący)**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/main_test.go
git commit -m "Test (failujacy): migracja v9 rename 7 polskich tagow"
```

---

### Task 15: Implementacja migracji v9 w database.go

**Files:**
- Modify: `analiza/cli/database.go`

- [ ] **Step 1: Zwiększ `currentSchemaVersion` 8 → 9**

W `database.go` linia 12:

```go
const currentSchemaVersion = 9
```

- [ ] **Step 2: Dodaj migrację v9 do listy `migrations`**

Po `{Version: 8, Apply: func(tx *sql.Tx) error { ... }},` (~linia 391), dodaj:

```go
{Version: 9, Apply: func(tx *sql.Tx) error {
    // Polonizacja formul arkusza — rename 7 tagow z zachowaniem FSRS state
    renames := []struct{ old, new string }{
        {"VLOOKUP", "WYSZUKAJ.PIONOWO"},
        {"SUMIF", "SUMA.JEŻELI"},
        {"SUMIFS", "SUMA.WARUNKÓW"},
        {"COUNTIF", "LICZ.JEŻELI"},
        {"COUNTIFS", "LICZ.WARUNKI"},
        {"AVERAGEIF", "ŚREDNIA.JEŻELI"},
        {"AVERAGEIFS", "ŚREDNIA.WARUNKÓW"},
    }
    for _, r := range renames {
        // Jesli NOWY juz istnieje (np. uzytkownik przeszedl na nowy CLI ale stary tag w starych skryptach)
        // — merge zachowujac max(stability, reps)
        _, err := tx.Exec(`
            INSERT INTO progress_tagi(tag, poziom, nastepna_powtorka, stability, difficulty, lapses, reps, state, last_review)
            SELECT ?, poziom, nastepna_powtorka, stability, difficulty, lapses, reps, state, last_review
            FROM progress_tagi WHERE tag = ?
            ON CONFLICT(tag) DO UPDATE SET
                stability = max(progress_tagi.stability, excluded.stability),
                reps = max(progress_tagi.reps, excluded.reps),
                last_review = CASE WHEN excluded.last_review > progress_tagi.last_review
                                   THEN excluded.last_review ELSE progress_tagi.last_review END
        `, r.new, r.old)
        if err != nil {
            return fmt.Errorf("v9 rename %s -> %s (insert): %w", r.old, r.new, err)
        }
        if _, err := tx.Exec("DELETE FROM progress_tagi WHERE tag = ?", r.old); err != nil {
            return fmt.Errorf("v9 rename %s -> %s (delete): %w", r.old, r.new, err)
        }
    }
    return nil
}},
```

- [ ] **Step 3: Uruchom test — powinien przejść**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestMigrationV9RenameTagi -v`
Expected: PASS

- [ ] **Step 4: Uruchom WSZYSTKIE testy żeby sprawdzić regresje**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test ./...`
Expected: ALL PASS (poprzednie 117 testów + 1 nowy)

- [ ] **Step 5: Commit migracji**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/database.go
git commit -m "Migracja schema v9: rename 7 polskich tagow w progress_tagi"
```

---

### Task 16: Rebuild CLI + reimport matura.db

**Files:**
- Modify: `analiza/cli/matura`, `matura.exe`, `matura.db` (regenerowane)

- [ ] **Step 1: Rebuild CLI**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && ./build.sh`
Expected: budują się matura (macOS) + matura.exe (Windows), reimport matura.db

- [ ] **Step 2: Smoke test CLI — exercise get z arkusza zwraca polskie formuły**

Run: `matura_informatyka_rozszerzona/analiza/cli/matura exercise get --typ agregacja_warunkowa 2>&1 | head -30`
Expected: w polach tresc/odpowiedz polskie nazwy (SUMA.JEŻELI, LICZ.JEŻELI itp.)

- [ ] **Step 3: Smoke test — CKE matura_*.json NIETKNIETE**

Run: `matura_informatyka_rozszerzona/analiza/cli/matura cke get --typ agregacja_warunkowa 2>&1 | head -20`
Expected: tresc z matura_*.json jest BEZ ZMIAN (np. angielskie nazwy lub oryginalna treść CKE)

- [ ] **Step 4: data stats — sanity check**

Run: `matura_informatyka_rozszerzona/analiza/cli/matura data stats`
Expected: `937 cwiczenia, 641 CKE subtasks, 4 cheatsheets`

- [ ] **Step 5: Commit rebuild artifacts**

```bash
git add matura_informatyka_rozszerzona/analiza/cli/matura \
        matura_informatyka_rozszerzona/analiza/cli/matura.exe \
        matura_informatyka_rozszerzona/analiza/cli/matura.db
git commit -m "Rebuild CLI + reimport matura.db po polonizacji"
```

---

## Faza 4 — Bonus + Final Validation (3 zadania)

### Task 17: Sekcja "Pułapki polskich nazw" w cheatsheet_arkusz.md

**Files:**
- Modify: `analiza/cheatsheets/cheatsheet_arkusz.md`

- [ ] **Step 1: Dodaj sekcję na końcu pliku**

Append do `matura_informatyka_rozszerzona/analiza/cheatsheets/cheatsheet_arkusz.md`:

```markdown

---

## Pułapki polskich nazw funkcji — uważaj!

| Funkcja | Polska nazwa | Pułapka |
|---|---|---|
| `ROUNDUP` | `ZAOKR.GÓRA` | NIE `ZAOKR.W.GÓRĘ` (to **CEILING** — zaokrąglanie do wielokrotności) |
| `ROUNDDOWN` | `ZAOKR.DÓŁ` | NIE `ZAOKR.W.DÓŁ` (to **FLOOR**) |
| `CONCATENATE` | `ZŁĄCZ.TEKSTY` (z "Y") | `ZŁĄCZ.TEKST` (bez Y) to nowsza `CONCAT` |
| `SUBSTITUTE` vs `REPLACE` | `PODSTAW` vs `ZASTĄP` | `PODSTAW` = po tekście, `ZASTĄP` = po pozycji znaków |
| `MAX` / `MIN` / `MOD` | bez zmian | Excel PL używa tych samych nazw — NIE tłumaczyć |
| Separator argumentów | `;` (średnik) | Polski locale używa średnika, NIE przecinka |
| Separator dziesiętny | `,` (przecinek) | Polski locale używa przecinka (np. `3,14` zamiast `3.14`) |
```

- [ ] **Step 2: Verify zawartość**

Run: `tail -15 matura_informatyka_rozszerzona/analiza/cheatsheets/cheatsheet_arkusz.md`
Expected: widać sekcję z 7-wierszową tabelą

- [ ] **Step 3: Commit bonusu**

```bash
git add matura_informatyka_rozszerzona/analiza/cheatsheets/cheatsheet_arkusz.md
git commit -m "Dodaj sekcje Pulapki polskich nazw funkcji w cheatsheet_arkusz"
```

---

### Task 18: Full validation suite

**Files:**
- (validation only — no file changes)

- [ ] **Step 1: validate_json.py — schema lint**

Run: `python3 matura_informatyka_rozszerzona/analiza/cwiczenia/validate_json.py`
Expected: `0 ERRORS`

- [ ] **Step 2: verify_all.py — content verification**

Run: `python3 matura_informatyka_rozszerzona/analiza/cwiczenia/verify/verify_all.py 2>&1 | tail -5`
Expected: `530 PASS, 0 FAIL, 0 ERROR, 407 MANUAL_REVIEW`

- [ ] **Step 3: test_qa.sh — pełna suita 180 testów**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh 2>&1 | tail -20`
Expected: wszystkie 7 warstw OK, 180/180 testów PASS

- [ ] **Step 4: Grep guard — żadnych angielskich nazw spreadsheet-specific w whitelisty**

Run:
```bash
echo "=== Warstwa 1+2 ==="
grep -rE '\b(VLOOKUP|HLOOKUP|SUMIF|SUMIFS|COUNTIF|COUNTIFS|AVERAGEIF|AVERAGEIFS|COUNTA|CONCATENATE|IFERROR|SUMPRODUCT|MAXIFS|MINIFS)\s*\(' \
  matura_informatyka_rozszerzona/analiza/szablony/arkusz_formuly.md \
  matura_informatyka_rozszerzona/analiza/cheatsheets/cheatsheet_arkusz.md \
  matura_informatyka_rozszerzona/analiza/rozwiazania_wzorcowe/arkusz_kalkulacyjny.md \
  matura_informatyka_rozszerzona/analiza/cheatsheets/debug_checklist.md \
  matura_informatyka_rozszerzona/analiza/cheatsheets/przed_egzaminem.md \
  matura_informatyka_rozszerzona/analiza/szablony/wzorce_2015.md \
  matura_informatyka_rozszerzona/analiza/PRZEWODNIK_UCZNIA.md \
  matura_informatyka_rozszerzona/analiza/strategia_egzaminacyjna.md \
  matura_informatyka_rozszerzona/analiza/drzewo_decyzyjne.md
echo "=== Warstwa 3 (cwiczenia arkuszowe) ==="
grep -rE '\b(VLOOKUP|HLOOKUP|SUMIF|SUMIFS|COUNTIF|COUNTIFS|AVERAGEIF|AVERAGEIFS|COUNTA|CONCATENATE|IFERROR|SUMPRODUCT|MAXIFS|MINIFS)\s*\(' \
  matura_informatyka_rozszerzona/analiza/cwiczenia/json/15_agregacja_warunkowa/ \
  matura_informatyka_rozszerzona/analiza/cwiczenia/json/16_symulacja/ \
  matura_informatyka_rozszerzona/analiza/cwiczenia/json/18_agregacja_podstawowa/ \
  matura_informatyka_rozszerzona/analiza/cwiczenia/json/19_transformacja/
echo "=== Warstwa 4 (rejestry) ==="
grep -E '"(VLOOKUP|SUMIF|SUMIFS|COUNTIF|COUNTIFS|AVERAGEIF|AVERAGEIFS)"' \
  matura_informatyka_rozszerzona/analiza/cwiczenia/json/tagi_rejestr.json
```
Expected: brak wyników wszędzie

- [ ] **Step 5: Sprawdź że SQL/C++ pliki NIE zostały zmienione (sanity)**

Run: `git log --oneline -20 -- matura_informatyka_rozszerzona/analiza/szablony/sql_szablony.md matura_informatyka_rozszerzona/analiza/szablony/cpp_szablony.md matura_informatyka_rozszerzona/analiza/cwiczenia/json/20_sql_group_by/ matura_informatyka_rozszerzona/analiza/json/matura_2024M.json`
Expected: ostatnie commity dla tych plików są SPRZED naszej polonizacji (cały dzień 2026-05-12 ich nie tyka)

- [ ] **Step 6: Pytest skryptu**

Run: `cd matura_informatyka_rozszerzona/analiza/scripts && python3 -m pytest test_polonizuj_formuly.py -v`
Expected: 9 PASS

---

### Task 19: Update MEMORY.md + final commit

**Files:**
- Modify: `/Users/blt1wz/.claude/projects/-Users-blt1wz-priv-informa/memory/MEMORY.md`
- Create: `/Users/blt1wz/.claude/projects/-Users-blt1wz-priv-informa/memory/polonizacja_formul_2026.md`

- [ ] **Step 1: Utwórz plik memory dla polonizacji**

Create `/Users/blt1wz/.claude/projects/-Users-blt1wz-priv-informa/memory/polonizacja_formul_2026.md`:

```markdown
---
name: polonizacja-formul-arkusza
description: Ukonczono 2026-05-12. Wszystkie materialy dydaktyczne arkusza uzywaja polskich nazw MS Excel (zgodnie z konwencja CKE). matura_*.json NIETKNIETE.
metadata:
  type: project
---

Polonizacja formul arkusza ukonczona 2026-05-12.

**Why:** CKE używa polskich nazw funkcji (np. WYSZUKAJ.PIONOWO, JEŻELI, SUMA.WARUNKÓW). Uczeń widzący w naszych materiałach VLOOKUP/IF/SUMIFS musi tłumaczyć w głowie. Niespójność dezorientowała.

**How to apply:** Przy dodawaniu nowych ćwiczeń arkuszowych / materiałów dydaktycznych — używaj polskich nazw MS Excel. Pełna mapa: `matura_informatyka_rozszerzona/analiza/scripts/polonizacja_mapa.json` (46 funkcji EN→PL).

**Zakres zmian:**
- 3 pliki MD dedykowane arkuszowi (warstwa 1)
- 6 plików MD mieszanych — tylko 20 wyliczonych linii (warstwa 2)
- 164 plików JSON ćwiczeń arkuszowych w katalogach 15/16/18/19 (warstwa 3)
- 2 rejestry: tagi_rejestr.json (rename 7 tagów), algorytmy_rejestr.json (1 linia)
- CLI: schema v9 migration w database.go (rename tagów w progress_tagi)
- cheatsheet_arkusz.md: dodana sekcja "Pułapki polskich nazw"

**NIE TYKANO:**
- matura_*.json (30 plików, 641 podzadań) — cytaty CKE
- Wszystkie pliki SQL i C++ (szablony, rozwiązania, ćwiczenia 01-14 + 20-23)
- 17_wykres (0 wystąpień formuł)

**Pułapki nazw (na przyszłość):**
- `SUMIFS` → `SUMA.WARUNKÓW` (NIE `SUMY.WARUNKÓW`)
- `AVERAGEIFS` → `ŚREDNIA.WARUNKÓW` (NIE `ŚREDNIE.WARUNKÓW`)
- `ROUNDUP` → `ZAOKR.GÓRA` (NIE `ZAOKR.W.GÓRĘ` = CEILING)
- `ROUNDDOWN` → `ZAOKR.DÓŁ` (NIE `ZAOKR.W.DÓŁ` = FLOOR)
- `MAX`, `MIN`, `MOD` — bez zmian (Excel PL używa tych samych)

**Reproducibility:** skrypt `analiza/scripts/polonizuj_formuly.py --dry-run --warstwa all` powinien zwrócić 0 zmian (sprawdzenie idempotentności po commit).

**Spec & Plan:**
- `docs/superpowers/specs/2026-05-12-polonizacja-formul-arkusza-design.md`
- `docs/superpowers/plans/2026-05-12-polonizacja-formul-arkusza.md`
```

- [ ] **Step 2: Dodaj wpis do MEMORY.md index**

Edit `/Users/blt1wz/.claude/projects/-Users-blt1wz-priv-informa/memory/MEMORY.md` — w sekcji "Completed Projects" dodaj jako drugi punkt:

```markdown
- [Polonizacja formuł arkusza (MS Excel PL)](polonizacja_formul_2026.md) — ✅ UKOŃCZONA 2026-05-12. Mapa 46 funkcji EN→PL. matura_*.json NIETKNIETE.
```

- [ ] **Step 3: Final commit całej polonizacji w MEMORY**

```bash
git add /Users/blt1wz/.claude/projects/-Users-blt1wz-priv-informa/memory/polonizacja_formul_2026.md \
        /Users/blt1wz/.claude/projects/-Users-blt1wz-priv-informa/memory/MEMORY.md
# UWAGA: te pliki sa POZA repo informa — moga byc w innym repo lub poza VCS.
# Jesli grep nie znajdzie git repo: po prostu zapisz i pomin commit.
git commit -m "MEMORY: dodaj wpis o ukonczonej polonizacji formul arkusza" 2>/dev/null || echo "Memory poza repo — zapisane bez commit"
```

- [ ] **Step 4: Re-run idempotentności skryptu**

Run: `matura_informatyka_rozszerzona/analiza/scripts/polonizuj_formuly.py --dry-run --warstwa all`
Expected: `TOTAL: 0 zmian` lub bardzo małe liczby (tylko opisowe wzmianki bez nawiasów)

- [ ] **Step 5: Final smoke test SKILL.md (czy używa nadal polskich tagów)**

Run: `grep -E 'SUMIF|VLOOKUP|COUNTIF' matura_informatyka_rozszerzona/analiza/SKILL.md 2>/dev/null`
Expected: brak wyników (lub jeśli są — dodaj jako 21 task)

---

## Self-Review

- [x] **Spec coverage**: każda sekcja spec'a mapowana na task: Mapa (T1), Warstwa 2 lista (T2), Skrypt (T4-9), Apply (T10-13), Migracja CLI (T14-16), Bonus pułapki (T17), Walidacja (T18), Memory (T19) ✓
- [x] **Placeholder scan**: brak TBD/TODO/implement_later w stepach
- [x] **Type consistency**: nazwy funkcji `polonizuj_md_warstwa1`, `polonizuj_md_warstwa2`, `polonizuj_json_cwiczenie`, `polonizuj_json_meta`, `polonizuj_tagi_rejestr`, `polonizuj_algorytmy_rejestr`, `normalizuj_separator`, `_polonizuj_string`, `_rename_tag`, `load_mapa`, `load_warstwa2`, `is_whitelisted`, `run_warstwa1/2/3/4` — wszystkie spójne, używane w testach i main()

## Execution

Plan zapisany w `docs/superpowers/plans/2026-05-12-polonizacja-formul-arkusza.md`. Po zatwierdzeniu przez użytkownika — wybór trybu wykonania.
