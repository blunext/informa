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
- **Data folders** - Named variably (dane_PR, Dane_NOWA, DANE_XXXX, Dane-NF-XXXX) containing text files with input data for computational problems
- **zalaczniki.zip** - Original compressed archive of attachments (kept for reference)

### Data File Characteristics

Data files in the attachment folders are typically:
- Plain text files (.txt)
- Contain structured data (CSV-like, space-separated, or custom formats)
- Include example files (e.g., `liczby_przyklad.txt`) and full datasets (e.g., `liczby.txt`)
- Used for programming tasks requiring file I/O, data processing, and algorithm implementation

## Important Context

### Exam Format Changes

- **Formula 2015** (stara formuła): Used for exams 2014-2022
- **Formula 2023** (nowa formuła): New exam format starting from 2023
- The 2023 format change represents a significant shift in exam structure and expectations

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
