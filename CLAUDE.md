# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is an archive of Polish extended-level computer science (informatyka rozszerzona) matura examination materials from the Central Examination Commission (CKE). The repository contains official exam papers, answer keys, and data files from 2014-2025 (including 2020, when the main exam was held in June instead of May due to COVID-19).

## Repository Structure

The repository is organized by year, with each year in a separate folder:

```
matura_informatyka_rozszerzona/
├── 2014_maj/
│   ├── arkusz.pdf          # Exam paper
│   ├── odpowiedzi.pdf      # Answer key and scoring guidelines
│   ├── dane_PR/            # Data files for exam problems
│   └── zalaczniki.zip      # Original attachments archive
├── 2015_maj/
├── ... (similar structure for each year)
└── 2025_maj/
```

### File Naming Conventions

- **arkusz.pdf** - The main examination paper containing all problems
- **odpowiedzi.pdf** - Official answer key with scoring criteria ("zasady oceniania")
- **Data folders** - Named differently each year (see table below), containing text files with input data for computational problems

### Data Folder Names by Year

| Year | Folder name |
|------|-------------|
| 2014 | `dane_PR/` |
| 2015 | `Dane_PR/` |
| 2016 | `Dane_NOWA/` |
| 2017-2019 | `Dane_PR/` |
| 2021 | `DANE_2105/` |
| 2022 | `Dane_2205/` |
| 2023 | `Dane_2305/` |
| 2024 | `Dane-NF-2405/` |
| 2025 | `dane maj 23/` (space in name — quote the path!) |
- **zalaczniki.zip** - Original compressed archive of attachments (kept for reference)

### Data File Characteristics

Data files in the attachment folders are typically:
- Plain text files (.txt)
- Contain structured data (CSV-like, space-separated, or custom formats)
- Include example files (e.g., `liczby_przyklad.txt`) and full datasets (e.g., `liczby.txt`)
- Used for programming tasks requiring file I/O, data processing, and algorithm implementation

## Analysis and Study Materials (`analiza/`)

The `analiza/` directory contains a comprehensive study system built on analysis of all exam years (2014-2025):

```
analiza/
├── PODSUMOWANIE_FINAL.md          # Year-by-year breakdown, topic/task type rankings
├── PRZEWODNIK_UCZNIA.md           # Student guide (study plan, self-assessment)
├── strategia_egzaminacyjna.md     # Main strategy: TOP 14 algorithms, C++ code, SQL, task types
├── drzewo_decyzyjne.md            # Decision tree: "see X → use algorithm Y"
├── podsumowanie_szybkie_wszystkie_lata.md  # Quick summary of all years
│
├── json/                          # Machine-readable analysis data
│   ├── matura_YYYYS.json          # **30 files — COMPLETE exam database** (641 subtasks, self-contained)
│   │                              # Full task text, answers, scoring criteria, typ_zadania, traps
│   │                              # No PDF needed — each subtask is standalone and solvable
│   ├── matura_indeks.json         # **Cross-reference index** for all 641 subtasks across 30 exams
│   │                              # Filter by typ_zadania, kategoria, rok — instant access
│   ├── algorytmy_rejestr.json     # **65 algorytmów w 4 kategoriach** — kanoniczny słownik dla
│   │                              # pola `algorytmy` w matura_*.json (faza 1, 2026-04).
│   │                              # 26 klasyczne + 7 techniki + 6 struktury + 26 wzorce.
│   │                              # 38 z linkiem do podstawy programowej 2024 (Dz.U. 2024 poz. 1019),
│   │                              # 27 wzorców CKE-pragmatycznych (iteracja-po-pliku, akumulator-*,
│   │                              # SQL-*, arkusz-*) bez direct linku do podstawy.
│   ├── ranking_tematow.csv        # Topic frequency matrix: 21 topics × 11 years
│   └── ranking_typow_zadan.csv    # Task type frequency: 23 types × 11 years + total points
│
├── cheatsheets/                   # Quick reference cards (7 files)
│   ├── cheatsheet_{cpp,sql,teoria,arkusz}.md  # 4 topic cheatsheets
│   ├── debug_checklist.md         # Error diagnosis by category (C++, SQL, spreadsheet)
│   ├── podczas_egzaminu.md        # Exam-day time strategy (phases + time allocation)
│   └── przed_egzaminem.md         # Pre-exam checklist (tools, algorithms, mindset)
│
├── szablony/                      # Full templates and patterns (6 files, ~97 KB total)
│   ├── cpp_szablony.md            # C++ templates: file I/O, digits, sort, GCD, binary search (27 KB)
│   ├── sql_szablony.md            # SQL: JOINs, GROUP BY, subqueries, advanced patterns (19 KB)
│   ├── arkusz_formuly.md          # Spreadsheet: $ addressing, SUMIFS, VLOOKUP, simulation (18 KB)
│   ├── pseudokod_wzorce.md        # CKE pseudocode → C++ mapping, 7 archetypes (15 KB)
│   ├── wzorce_2014.md             # Year-specific patterns: recursion, bisection (8.5 KB)
│   └── wzorce_2015.md             # Year-specific: greedy, extended Euclidean (9.1 KB)
│
├── cwiczenia/                     # Training exercises
│   ├── cwiczenia_sledzenie.md     # 24 algorithm-tracing exercises (7 archetypes + bonus)
│   ├── validate_json.py           # **Standalone JSON schema validator** (no MD needed)
│   ├── wg_typu/                   # 937 ćwiczeń w formie czytelnej markdown (referencja dla ucznia)
│   │   ├── README.md              # Index with difficulty levels and self-assessment
│   │   ├── 01_sledzenie_algorytmu.md ... 23_sql_select_where.md
│   │   └── (each has: skill tags, 3-level hints, full answer, common CKE mistakes)
│   ├── json/                      # 23 directories — 937 exercises, machine-readable
│   │   ├── README.md              # **Schema, format danych, checklist dodawania cwiczen**
│   │   ├── tagi_rejestr.json      # **Central tag registry (290 tags, enforced by validator)**
│   │   └── NN_nazwa_typu/         # 23 directories, each with:
│   │       ├── _meta.json         #   Lightweight index (~2KB): metadata + exercise list
│   │       └── {id}.json          #   Individual exercise (~3-5KB each)
│   └── verify/                    # Automated verification system
│       ├── verify_all.py          # Main runner (--file, --id, --category, --verbose)
│       └── verifiers/             # cpp, sql, numconv, manual_sanity verifiers
│
├── cli/                          # **Go CLI binary + SQLite backend**
│   ├── matura                    # macOS/Linux binary (pure Go, zero deps)
│   ├── matura.exe                # Windows binary
│   ├── matura.db                 # SQLite DB: 937 exercises + 641 CKE + 4 cheatsheets
│   ├── main.go, commands.go, database.go, importer.go, types.go
│   ├── main_test.go              # 117 tests
│   ├── build.sh                  # build macOS + Windows + import
│   ├── go.mod, go.sum
│   └── matura_progress.db        # Student progress (auto-created, .gitignore)
│
├── rozwiazania_wzorcowe/          # Model solutions for real past exam problems
│   ├── implementacja_cpp.md       # C++ implementations (e.g., 2024 Zad.3)
│   ├── sql_zapytania.md           # SQL queries (e.g., 2023 Zad.7)
│   ├── teoria_algorytmy.md        # Theory/tracing (e.g., 2025 Zad.1)
│   └── arkusz_kalkulacyjny.md     # Spreadsheet (e.g., 2023 Zad.6)
│
├── RANKING_ALGORYTMOW.md          # **Raport pedagogiczny algorytmów** (faza 1, 2026-04, 336 linii).
│                                  # TIER 1/2/3 algorytmów, heatmapa rok×algorytm, top kombinacje,
│                                  # algorytmy z podstawy programowej NIE testowane przez CKE.
│                                  # Generowany skryptem — reproducible po dodaniu nowych roczników.
│
└── scripts/
    ├── validate_algorytmy.py      # Walidator klasyfikacji algorytmów (sprawdza pole algorytmy
    │                              # we wszystkich matura_*.json względem rejestru, statystyki użycia)
    └── generate_ranking.py        # Generator RANKING_ALGORYTMOW.md z 30 plików JSON + rejestru
```

### Klasyfikacja algorytmów (faza 1 ukończona 2026-04-25)

Wszystkie 641 podzadań w `matura_*.json` ma pole `algorytmy: [...]` z tagami z zamkniętego rejestru `algorytmy_rejestr.json`.

- **Pokrycie**: 604/641 podzadań (94.2%) z tagami, 1456 wystąpień łącznie, średnio 2.41 tagu/podzadanie.
- **TIER 1 (16 algorytmów ≥30 wystąpień) pokrywa 93.2% punktów na maturze** — patrz `analiza/RANKING_ALGORYTMOW.md`. Top 5: `iteracja-po-pliku`, `SQL-JOIN`, `SQL-aggregacja`, `akumulator-licznik`, `sledzenie-pseudokod`.
- **4 algorytmy z podstawy programowej NIE testowane przez CKE 2014-2025**: `fraktale-rekurencyjne`, `najdluzszy-wspolny-podciag` (LCS), `podciag-najwieksza-suma` (Kadane), `metoda-wstepujaca-zstepujaca`.
- **Walidacja**: `python3 analiza/scripts/validate_algorytmy.py` — sprawdza wszystkie tagi w polu `algorytmy` względem rejestru (exit 0 = OK).
- **Regeneracja raportu**: `python3 analiza/scripts/generate_ranking.py` po dodaniu nowych roczników.
- **Faza 2 (NIE zaimplementowana, na później)**: import pola `algorytmy` do `matura.db`, CLI commands typu `cke get --algorytm`, integracja ze `SKILL.md` (filtrowanie zadań po algorytmie), tagowanie 937 ćwiczeń treningowych.

### Key Analysis Data

- **Complete exam database**: `matura_YYYYS.json` — 30 files (12 maj + 8 czerwiec + 3 probna + 2 przykladowy + 5 other), 641 subtasks, 1500 points total. Each subtask has full text (`tresc`), answer (`odpowiedz`), scoring (`zasady_oceniania`), type (`typ_zadania`), traps (`pulapki`), and **`algorytmy: [...]`** (faza 1, 2026-04 — tagi algorytmów z `algorytmy_rejestr.json`). Self-contained: no PDF needed to solve any task. ID format: `YYYYS.Z.S` (e.g. `2025M.1.1`, `2024C.3.2`, `2024P.1.1`), where S=session letter (M=maj, C=czerwiec, P=probna, X=przykladowy).
- **Cross-reference index**: `matura_indeks.json` — all 641 subtasks indexed by typ_zadania, kategoria, rok. Supports filtering for cross-year practice (e.g., "all SQL tasks" or "all 2025 tasks").
- **23 task types** in 4 categories: TEORIA (6), IMPLEMENTACJA (8), ARKUSZ (5), SQL (4). All types use canonical prefixed names (e.g., `przetwarzanie_napisy`, `arkusz_symulacja`, `sql_group_by`).
- **Topic frequency tiers**: TIER 1 (100%): SQL, number ops, file processing, spreadsheet; TIER 2 (73-82%): number systems, recursion, sorting, GCD
- **No code files** exist in the year directories — the repo is documentation and reference material only

### Go CLI (`matura`)

Binary in `analiza/cli/` — provides fast SQLite-backed access to all data. JSON files remain the source of truth; the CLI imports from them.

```bash
# Build + import (macOS + Windows + matura.db):
analiza/cli/build.sh
```

Key commands: `exercise question`, `exercise hints`, `exercise answer`, `progress update`, `cke get`, `exam meta/task`, `trap list`, `cheatsheet get`, `data stats`. See `./matura --help` for full reference.

## Important Context

### Exam Format Changes

- **2014**: Part I 90 min / 20 pts + Part II 120 min / 30 pts (unique to this year)
- **2015-2022** (stara formuła): Part I 60 min / 15 pts + Part II 150 min / 35 pts (6 tasks)
- **2023-2025** (nowa formuła): Single paper 210 min / 50 pts, 7-8 tasks

### 2020 Note

- **2020**: The main exam was held in June 2020 (not May) due to COVID-19. The resit exam was held in July 2020. A diagnostic/practice exam (próbna) was also held in April 2020. All three sessions are in the repository.

### Typical Exam Content

Extended-level computer science matura exams typically cover:
- Algorithms and data structures
- File processing and text manipulation
- Number theory and mathematical computations
- Database queries (SQL)
- Sorting, searching, and optimization problems
- Complexity analysis
- Programming logic and problem-solving

## Working with This Repository

### Accessing Exam Materials

To view a specific year's exam:
```bash
cd matura_informatyka_rozszerzona/YYYY_maj/
open arkusz.pdf          # View the exam paper
open odpowiedzi.pdf      # View answer key
ls -la Dane*/            # List available data files
```

### Data File Analysis

When working with data files:
1. Check for example files first (files with `_przyklad` suffix) to understand format
2. Data files may contain Polish characters (UTF-8 encoding)
3. File formats vary by year and problem - always read the exam paper for specifications

### Downloading Additional Years

The materials were downloaded from:
- Primary source: https://arkusze.pl/informatyka-matura-poziom-rozszerzony/
- URL pattern: `https://arkusze.pl/maturalne/informatyka-YYYY-maj-matura-rozszerzona.pdf`
- Attachments: `https://arkusze.pl/maturalne/informatyka-YYYY-maj-matura-rozszerzona-zalaczniki.zip`

To download additional materials or update existing ones:
```bash
curl -L -o arkusz.pdf "https://arkusze.pl/maturalne/informatyka-YYYY-maj-matura-rozszerzona.pdf"
curl -L -o odpowiedzi.pdf "https://arkusze.pl/maturalne/informatyka-YYYY-maj-matura-rozszerzona-odpowiedzi.pdf"
curl -L -o zalaczniki.zip "https://arkusze.pl/maturalne/informatyka-YYYY-maj-matura-rozszerzona-zalaczniki.zip"
unzip -q zalaczniki.zip
```

## Working with Exercises (`analiza/cwiczenia/json/`)

Full documentation: **`analiza/cwiczenia/json/README.md`** — JSON schema, required data format patterns, verification commands, and a step-by-step checklist for adding new exercises.

Key points:
- Exercises live in `json/NN_nazwa/` directories (23 dirs, 937 exercises total)
- **C++ exercises (07-14)**: input data in `tresc` must use `**Dane** (\`plik.txt\`):` format or verifier won't find it
- **SQL exercises (20-23)**: tables in `tresc` via `Tabela **Name**:`, last non-verification markdown table in `odpowiedz` = expected result

### Quality gates (MUST run after any exercise edit)

```bash
# 1. Schema validation (structure, tags, points):
python3 analiza/cwiczenia/validate_json.py              # all directories
python3 analiza/cwiczenia/validate_json.py --file NN     # one directory

# 2. Content verification (C++ compile+run, SQL exec, sanity checks):
python3 analiza/cwiczenia/verify/verify_all.py           # all directories
python3 analiza/cwiczenia/verify/verify_all.py --file NN_nazwa --verbose  # one directory
```

**Expected results:**
- `validate_json.py`: **0 ERRORS**
- `verify_all.py`: **530 PASS, 0 FAIL, 0 ERROR, 407 MANUAL_REVIEW**

### Tag registry (`json/tagi_rejestr.json`)

Central list of 290 allowed tags. Validator rejects any tag not in this file.
To add a new tag: add it to `tagi_rejestr.json` (keep alphabetical), then also to `tagi_globalne` of the relevant JSON file.

### Generating new exercises

Use the `/generate-exercises` skill — it guides the full workflow: load context, generate, insert, validate, fix.

## Common Tasks

### Adding Solutions

When creating solution files for exam problems:
- Create a solutions directory: `YYYY_maj/solutions/`
- Name solution files clearly: `zadanie_1.py`, `zad_4_1.cpp`, etc.
- Include comments referencing the specific problem number
- Test solutions against the provided example data files first

### Analyzing Trends

To analyze patterns across years:
- Compare data file sizes and formats
- Track changes in problem types between Formula 2015 and Formula 2023
- Examine scoring criteria evolution in `odpowiedzi.pdf` files

### Generating Statistics

The repository currently contains 12 years of exam materials spanning 2014-2025, with 30 exam sessions (maj, czerwiec, próbna, przykładowy) totaling 641 subtasks and 1500 points.
