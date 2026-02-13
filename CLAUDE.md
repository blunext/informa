# CLAUDE.md

This file provides guidance to Claude Code (claude.ai/code) when working with code in this repository.

## Repository Overview

This is an archive of Polish extended-level computer science (informatyka rozszerzona) matura examination materials from the Central Examination Commission (CKE). The repository contains official exam papers, answer keys, and data files from 2014-2025 (excluding 2020, when the exam was not held due to COVID-19).

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

The `analiza/` directory contains a comprehensive study system built on analysis of all 11 exam years:

```
analiza/
├── PODSUMOWANIE_FINAL.md          # Year-by-year breakdown, topic/task type rankings
├── PRZEWODNIK_UCZNIA.md           # Student guide (study plan, self-assessment)
├── strategia_egzaminacyjna.md     # Main strategy: TOP 14 algorithms, C++ code, SQL, task types
├── drzewo_decyzyjne.md            # Decision tree: "see X → use algorithm Y"
├── podsumowanie_szybkie_wszystkie_lata.md  # Quick summary of all years
│
├── json/                          # Machine-readable analysis data
│   ├── analiza_YYYY.json          # 11 files, per-year detailed analysis (each subtask has typ_zadania)
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
│   ├── wg_typu/                   # 23 files × 10 exercises = 230 total (markdown)
│   │   ├── README.md              # Index with difficulty levels and self-assessment
│   │   ├── 01_sledzenie_algorytmu.md ... 23_sql_select_where.md
│   │   └── (each has: skill tags, 3-level hints, full answer, common CKE mistakes)
│   ├── json/                      # 23 JSON files — same 230 exercises, machine-readable
│   │   ├── README.md              # **Schema, format danych, checklist dodawania cwiczen**
│   │   └── 01_sledzenie_algorytmu.json ... 23_sql_select_where.json
│   └── verify/                    # Automated verification system
│       ├── verify_all.py          # Main runner (--file, --id, --category, --verbose)
│       └── verifiers/             # cpp, sql, numconv, manual verifiers
│
└── rozwiazania_wzorcowe/          # Model solutions for real past exam problems
    ├── implementacja_cpp.md       # C++ implementations (e.g., 2024 Zad.3)
    ├── sql_zapytania.md           # SQL queries (e.g., 2023 Zad.7)
    ├── teoria_algorytmy.md        # Theory/tracing (e.g., 2025 Zad.1)
    └── arkusz_kalkulacyjny.md     # Spreadsheet (e.g., 2023 Zad.6)
```

### Key Analysis Data

- **23 task types** in 4 categories: TEORIA (6), IMPLEMENTACJA (8), ARKUSZ (5), SQL (4)
- **Topic frequency tiers**: TIER 1 (100%): SQL, number ops, file processing, spreadsheet; TIER 2 (73-82%): number systems, recursion, sorting, GCD
- **Known gap**: `analiza_2018.json` only has Part I data (Part II was never analyzed)
- **No code files** exist in the year directories — the repo is documentation and reference material only

## Important Context

### Exam Format Changes

- **2014**: Part I 90 min / 20 pts + Part II 120 min / 30 pts (unique to this year)
- **2015-2022** (stara formuła): Part I 60 min / 15 pts + Part II 150 min / 35 pts (6 tasks)
- **2023-2025** (nowa formuła): Single paper 210 min / 50 pts, 7-8 tasks

### Missing Year

- **2020**: No exam materials available - the extended-level computer science matura exam did not take place in May 2020 due to the COVID-19 pandemic

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
- Exercises live in `json/NN_nazwa.json` files (23 files, 230 exercises total)
- Automated verification: `python3 analiza/cwiczenia/verify/verify_all.py` (target: 130 PASS, 0 FAIL)
- **C++ exercises (07-14)**: input data in `tresc` must use `**Dane** (\`plik.txt\`):` format or verifier won't find it
- **SQL exercises (20-23)**: tables in `tresc` via `Tabela **Name**:`, last non-verification markdown table in `odpowiedz` = expected result
- Always verify after edits: `--file <name> --verbose` or `--id NN.M --verbose`

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

The repository currently contains 11 years of exam materials spanning 2014-2025 (excluding 2020), totaling approximately 17 MB of data.
