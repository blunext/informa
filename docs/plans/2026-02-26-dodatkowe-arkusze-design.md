# Rozszerzenie systemu o dodatkowe arkusze CKE

**Data:** 2026-02-26
**Status:** Zatwierdzony

## Kontekst

Obecnie system zawiera 11 arkuszy majowych (2014-2025, bez 2020) z 230 subtaskami CKE. Na arkusze.pl dostępnych jest ~36 dodatkowych arkuszy: dodatkowe (czerwiec), próbne, przykładowe, stara formuła, pre-2014.

## Zakres (faza 1)

Nowa formuła (2023+) — 5 dodatkowych arkuszy:

| Plik | Sesja | Źródło |
|------|-------|--------|
| `matura_2023C.json` | czerwiec 2023 | arkusze.pl |
| `matura_2023X.json` | przykładowy 2023 | arkusze.pl |
| `matura_2024C.json` | czerwiec 2024 | arkusze.pl |
| `matura_2024P.json` | próbna grudzień 2024 | arkusze.pl |
| `matura_2025C.json` | czerwiec 2025 | arkusze.pl |

Szacunkowo ~100 nowych subtasków CKE (230 → ~330).

## Schemat identyfikatorów

**Format:** `ROKS.ZADANIE.PODZADANIE` gdzie S = litera sesji

| Sesja | Prefix | Przykład |
|-------|--------|----------|
| Maj (główna) | `M` | `2025M.1.1` |
| Czerwiec (dodatkowa) | `C` | `2024C.1.1` |
| Próbna | `P` | `2024P.1.1` |
| Przykładowy | `X` | `2023X.1.1` |

**Migracja:** Obecne `2025.1.1` → `2025M.1.1` (230 subtasków).

## Zmiany w plikach JSON

### Pliki egzaminów (`analiza/json/`)

- Rename: `matura_YYYY.json` → `matura_YYYYM.json` (11 plików)
- Nowe pole: `"sesja": "maj"` | `"czerwiec"` | `"probna"` | `"przykladowy"`
- Prze-ID-owanie wszystkich podzadań: `YYYY.Z.P` → `YYYYM.Z.P`
- Nowe pliki: 5 arkuszy nowej formuły (jak w tabeli powyżej)

### Indeks (`analiza/json/matura_indeks.json`)

- Prze-ID-owanie 230 istniejących wpisów
- Dodanie nowych wpisów z nowych arkuszy

## Zmiany w SQLite

### matura.db (tabela `egzamin`)

```sql
-- Nowa kolumna:
ALTER TABLE egzamin ADD COLUMN sesja TEXT DEFAULT 'maj';
CREATE INDEX idx_egzamin_sesja ON egzamin(rok, sesja);
```

Import jest DROP+CREATE, więc wystarczy zmienić schemat w Go i reimportować.

### progress.db (migracja schema version N+1)

```sql
-- Migracja ID w matura_zrobione:
UPDATE matura_zrobione
SET id = substr(id, 1, 4) || 'M.' || substr(id, 6)
WHERE id NOT LIKE '%M.%' AND id NOT LIKE '%C.%' AND id NOT LIKE '%P.%' AND id NOT LIKE '%X.%';

-- Nowa kolumna w probne_matury:
ALTER TABLE probne_matury ADD COLUMN sesja TEXT DEFAULT 'maj';
```

## Zmiany w CLI

### Nowe flagi

- `--sesja {maj|czerwiec|probna|przykladowy}` na komendach: `exam task`, `exam list`, `cke get`

### Zmienione komendy

- `exam list`: wyświetla sesje, np. `2024 maj ✓ | 2024 czerwiec | 2024 próbna`
- `exam task --rok 2024 --sesja czerwiec`: nowy flag (default: `maj`)
- `cke get`: losuje z pełnej puli (wszystkie sesje); `--sesja` do filtrowania
- `data stats`: raportuje count per sesja
- `exam save`: przyjmuje `--sesja`

## Strategia implementacji (3 fazy z bramkami)

### Faza 1: Migracja ID (bez nowych danych)

1. Rename `matura_YYYY.json` → `matura_YYYYM.json`
2. Dodaj `"sesja": "maj"` do każdego pliku
3. Prze-ID-uj 230 subtasków we wszystkich plikach
4. Prze-ID-uj `matura_indeks.json`
5. Zaktualizuj Go: schemat, importer, komendy
6. Dodaj migrację progress.db (schema version N+1)
7. **Bramka:** `test_qa.sh` musi przejść na 100%

### Faza 2: Rozszerzenie CLI o sesje

1. Dodaj `--sesja` flag do `exam`, `cke` komend
2. Rozszerz `exam list` o wyświetlanie sesji
3. Rozszerz importer o parsowanie sesji z nazwy pliku (`matura_2024C.json` → sesja=czerwiec)
4. Zaktualizuj `data stats`
5. **Bramka:** `go test ./...` + `test_qa.sh`

### Faza 3: Dodanie nowych arkuszy

1. Pobranie PDF-ów z arkusze.pl (5 arkuszy nowej formuły)
2. Tworzenie JSON-ów z treścią zadań (Claude czyta PDF → JSON)
3. Ręczna weryfikacja treści i odpowiedzi
4. Import do matura.db
5. **Bramka:** `validate_json.py` + `verify_all.py` + `test_qa.sh --update-baseline`

## Ryzyka i mitigacja

| Ryzyko | Mitigacja |
|--------|-----------|
| Utrata progressu (progress.db) | Backup przed migracją; testowany na kopii |
| CLI się psuje | 112 testów QA; każda faza = osobny commit |
| Nowe JSON-y z błędami | validate_json.py + verify_all.py |
| SKILL.md/indeks rozjeżdżenie | Lint w test_qa.sh Layer 3 |

## Przyszłe rozszerzenia (poza zakresem)

- Dodatkowe arkusze starej formuły (2015-2022)
- Pre-2014 egzaminy (2005-2013)
- Stara formuła równoległa (2015-2025)
