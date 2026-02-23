# Update `/generate-exercises` Skill — Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Update the `/generate-exercises` skill with 9 fixes found during audit (3 critical, 3 important, 3 minor).

**Architecture:** Single-file edit of `.claude/skills/generate-exercises/SKILL.md` — reorganize Krok 5 to include re-import+baseline, add missing fields/rules, add "Typowe błędy" section.

**Tech Stack:** Markdown skill file, no code changes.

---

### Task 1: Fix numeracji w Kroku 2

**Files:**
- Modify: `.claude/skills/generate-exercises/SKILL.md:38`

**Step 1: Fix duplicate numbering**

Change line 38 from:
```
4. **Szablony i wzorce** odpowiednie dla kategorii:
```
to:
```
5. **Szablony i wzorce** odpowiednie dla kategorii:
```

**Step 2: Verify** — Read file, confirm items are numbered 1-5.

---

### Task 2: Dodaj `zrodlo` do przykładu JSON + uzupełnij reguły

**Files:**
- Modify: `.claude/skills/generate-exercises/SKILL.md:48-77`

**Step 1: Verify example JSON already has `zrodlo`**

Read lines 48-66. The example JSON should already contain `"zrodlo"`. If it does, skip to Step 2.

If `zrodlo` is missing, add it after `"punkty": 3,`:
```json
  "zrodlo": "Inspirowane matura 2023 zad. 4",
```

**Step 2: Update reguły bezwzględne (after line 76)**

Replace line 76:
```
- **typowe_bledy**: minimum 1, kazdy z `opis` (bold prefix) i `kara` (format: `-N pkt`)
```
with:
```
- **typowe_bledy**: minimum 1, kazdy z `opis` (bold prefix) i `kara` (format: `-N pkt` lub `-N.N pkt`, np. `-0.5 pkt`)
- **punkty**: zakres 1-10
- **zrodlo**: niepusty string opisujacy inspiracje (np. "Wzor: Matura 2023 zad. 4")
```

**Step 3: Verify** — Read the section, confirm all 9 fields documented in reguły.

---

### Task 3: Dodaj regułę konwersji (katalog 05)

**Files:**
- Modify: `.claude/skills/generate-exercises/SKILL.md:106-111`

**Step 1: Add konwersje rule after TEORIA/sledzenie section**

After line 108 (`Dla P/F (plik 04): odpowiedz musi zawierac PRAWDA/FALSZ...`), add:

```markdown

#### Konwersje systemow liczbowych (plik 05)
- `odpowiedz` musi zawierac wzorce konwersji: `45(10) = 101101(2)`, `B6(16) = 182(10)` itp.
- Weryfikator automatycznie sprawdza poprawnosc konwersji (PASS/FAIL, nie MANUAL_REVIEW)
```

**Step 2: Verify** — Read section, confirm 5 category rules: IMPLEMENTACJA, SQL, TEORIA, Konwersje, ARKUSZ.

---

### Task 4: Dodaj warning o field matching w Kroku 4

**Files:**
- Modify: `.claude/skills/generate-exercises/SKILL.md:170-176`

**Step 1: Add field matching warning**

After line 176 (end of Krok 4), add:

```markdown

**UWAGA**: Pola `trudnosc`, `punkty`, `tagi` w `_meta.json` MUSZA byc identyczne z polami w pliku cwiczenia. Walidator sprawdza spojnosc — roznica = ERROR.
```

**Step 2: Verify** — Read Krok 4, confirm warning is present.

---

### Task 5: Reorganizuj Krok 5 (walidacja + re-import + baseline)

**Files:**
- Modify: `.claude/skills/generate-exercises/SKILL.md:178-199`

**Step 1: Replace entire Krok 5 section**

Replace lines 178-199 with:

```markdown
## Krok 5: Walidacja + Re-import (OBOWIAZKOWY)

### 5a. Schema + weryfikacja merytoryczna

```bash
# 1. Walidacja schema JSON (struktura, tagi, punkty):
python3 analiza/cwiczenia/validate_json.py --file NN_nazwa

# 2. Weryfikacja merytoryczna (kompilacja C++, SQL, sanity checks):
python3 analiza/cwiczenia/verify/verify_all.py --file NN_nazwa --verbose
```

### Mapowanie weryfikatorow

| Katalogi | Weryfikator | Wynik |
|----------|-------------|-------|
| 01-04, 06 | manual_sanity | MANUAL_REVIEW (zweryfikowane w Kroku 3.5) |
| 05 | numconv | PASS/FAIL |
| 07-14 | cpp (kompilacja + uruchomienie) | PASS/FAIL/ERROR |
| 15-19 | manual_sanity | MANUAL_REVIEW (zweryfikowane w Kroku 3.5) |
| 20-23 | sql (SQLite exec) | PASS/FAIL/ERROR |

### Interpretacja wynikow

- **validate_json.py 0 ERRORS** = struktura OK
- **verify_all.py PASS** = kod/SQL daje poprawny wynik
- **verify_all.py MANUAL_REVIEW** = TEORIA/ARKUSZ — juz zweryfikowane w Kroku 3.5
- **verify_all.py FAIL** = blad merytoryczny — NAPRAW i uruchom ponownie

Jezeli sa bledy — napraw je i uruchom walidacje ponownie. Powtarzaj az:
- validate_json.py: 0 ERRORS
- verify_all.py: 0 FAIL, 0 ERROR

### 5b. Re-import do CLI

```bash
cd analiza/cli && ./matura data import --source ../
```

Bez re-importu CLI nie zobaczy nowych cwiczen!

### 5c. Aktualizacja baseline

```bash
cd analiza && ./test_qa.sh --update-baseline
```

Bez tego Layer 4 test_qa.sh bedzie failowal przy nastepnym uruchomieniu.
```

**Step 2: Verify** — Read new Krok 5, confirm sub-sections 5a/5b/5c present.

---

### Task 6: Dodaj sekcję "Typowe błędy generacji"

**Files:**
- Modify: `.claude/skills/generate-exercises/SKILL.md` (before Krok 6)

**Step 1: Insert new section before Krok 6**

Before the `## Krok 6: Podsumowanie` line, add:

```markdown
## Typowe bledy generacji (unikaj!)

1. **Zly format `kara`** — zawsze `-N pkt` lub `-N.N pkt` (regex: `^-\d+(\.\d+)? pkt$`). NIE: `"-2 pkt (brak odp.)"`, `""`, `"minus 2"`
2. **Zly format `**Dane**`** — MUSI byc dokladnie `` **Dane** (`plik.txt`): `` (z backtick-ami!). NIE: `Liczby pierwsze (plik.txt):`
3. **Notatki robocze w oczekiwanym wyniku** — blok `**Oczekiwany wynik**` = TYLKO czyste linie stdout. Bez komentarzy, wyjasn, numeracji krokow
4. **SQL: zla tabela wynikowa** — ostatnia tabela markdown w `odpowiedz` (bez ✓/✗) = wynik. Tabele weryfikacyjne z ✓/✗ sa pomijane
5. **Tag spoza rejestru** — walidator ODRZUCI. Dodaj do `tagi_rejestr.json` PRZED uzyciem
6. **Niezgodnosc `_meta.json` ↔ plik** — `trudnosc`, `punkty`, `tagi` musza byc identyczne w obu miejscach
7. **Brak `zrodlo`** — pole wymagane, walidator odrzuci. Min. 1 znak
8. **Zly ID** — format `NN.M` gdzie NN = numer katalogu. Sprawdz ostatni istniejacy w `_meta.json`!

```

**Step 2: Verify** — Read section, confirm 8 items present.

---

### Task 7: Rozszerz Krok 6 podsumowanie

**Files:**
- Modify: `.claude/skills/generate-exercises/SKILL.md` (Krok 6 section)

**Step 1: Add re-import status to summary**

Replace the Krok 6 list:
```markdown
Wyswietl uzytkownikowi:
- Ile cwiczen dodano
- Do jakiego pliku
- Wynik walidacji
- Aktualna liczba cwiczen w pliku (bylo X, jest Y)
```
with:
```markdown
Wyswietl uzytkownikowi:
- Ile cwiczen dodano
- Do jakiego pliku
- Wynik walidacji (validate_json + verify_all)
- Re-import CLI: OK/FAIL
- Baseline: zaktualizowany
- Aktualna liczba cwiczen w pliku (bylo X, jest Y)
```

**Step 2: Verify** — Read Krok 6, confirm 6 bullet points.

---

### Task 8: Final verification

**Step 1: Read entire SKILL.md** — verify flow: Krok 1 → 2 → 3 → 3.5 → 4 → 5(a,b,c) → Typowe błędy → 6

**Step 2: Check file size** — should be ~280-300 lines (was 208, adding ~80 lines)

**Step 3: Sanity check** — verify no broken markdown (unclosed code blocks, missing headers)
