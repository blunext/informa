# Cwiczenia JSON — format i weryfikacja

## Struktura katalogow

Kazdy typ cwiczen to katalog z indeksem (`_meta.json`) i osobnymi plikami cwiczen:

```
json/
├── 01_sledzenie_algorytmu/
│   ├── _meta.json           (~2KB — metadane + lista cwiczen)
│   ├── 1.1.json             (~3-5KB — jedno cwiczenie)
│   ├── 1.2.json
│   └── ...1.10.json
├── 02_projektowanie_algorytmu/
│   ├── _meta.json
│   ├── 2.1.json
│   └── ...
├── ...
├── 23_sql_select_where/
│   ├── _meta.json
│   ├── 23.1.json
│   └── ...23.20.json
├── tagi_rejestr.json        (centralny rejestr tagow)
└── README.md                (ten plik)
```

## Schema `_meta.json`

Lekki indeks (~2KB) z metadanymi typu i lista cwiczen:

```jsonc
{
  "typ": "07_cyfry_liczby",           // = nazwa katalogu
  "nazwa": "Implementacja — cyfry i liczby",
  "kategoria": "IMPLEMENTACJA",       // TEORIA | IMPLEMENTACJA | ARKUSZ | SQL
  "czestotliwosc": "8/12 lat",        // format: "N/12 lat"
  "punkty_lacznie": 36,               // suma punktow wszystkich cwiczen
  "tagi_globalne": ["petla", "modulo", "dzielenie-calkowite"],
  "cwiczenia": [                       // indeks cwiczen (bez pelnej tresci!)
    {"id": "7.1", "trudnosc": "latwe", "punkty": 2, "tagi": ["petla", "modulo"]},
    {"id": "7.2", "trudnosc": "latwe", "punkty": 3, "tagi": ["petla", "dzielenie-calkowite"]},
    ...
  ]
}
```

## Schema cwiczenia (`{id}.json`)

Pojedynczy plik (~3-5KB) z pelna trescia:

```jsonc
{
  "id": "7.1",                        // format: "NN.M" (nr katalogu . nr cwiczenia)
  "trudnosc": "latwe",                // latwe | srednie | srednie-trudne | trudne
  "punkty": 2,
  "zrodlo": "Matura 2023 (Gry planszowe)",
  "tagi": ["petla", "modulo"],        // podzbiór tagi_globalne
  "tresc": "...",                      // markdown: polecenie + dane
  "wskazowki": [                       // 3 wskazowki (kierunek, podejscie, kluczowy krok)
    "**Kierunek**: ...",
    "**Podejscie**: ...",
    "**Kluczowy krok**: ..."
  ],
  "odpowiedz": "...",                 // markdown: kod + weryfikacja + tabela wynikowa
  "typowe_bledy": [
    { "opis": "**Opis bledu**: wyjasnienie", "kara": "-1 pkt" }
  ]
}
```

## Jak weryfikator odczytuje cwiczenia

System `verify/verify_all.py` automatycznie testuje cwiczenia z kategorii **IMPLEMENTACJA** (C++), **SQL** i **TEORIA/konwersje** (plik 05). Reszta dostaje status MANUAL_REVIEW.

### Przypisanie weryfikatora wg numeru katalogu

| Katalogi | Kategoria | Weryfikator | Status |
|----------|-----------|-------------|--------|
| 01-04 | TEORIA | manual_sanity | MANUAL_REVIEW / FAIL |
| 05 | TEORIA (konwersje) | numconv | PASS/FAIL |
| 06 | TEORIA | manual_sanity | MANUAL_REVIEW / FAIL |
| **07-14** | **IMPLEMENTACJA** | **cpp** | **PASS/FAIL/ERROR** |
| 15-19 | ARKUSZ | manual_sanity | MANUAL_REVIEW / FAIL |
| **20-23** | **SQL** | **sql** | **PASS/FAIL/ERROR** |

## Format `tresc` — IMPLEMENTACJA (C++, katalogi 07-14)

Weryfikator C++ szuka w `tresc` dwoch rzeczy:

### 1. Pliki wejsciowe

Wzorzec (regex: `\*\*Dane\*\*\s*\(`plik.txt`\)\s*:\s*\n```\n...\n````):

```markdown
**Dane** (`liczby.txt`):
```
12 45 7 89 23
15 3 67 41 8
```
```

Mozna definiowac wiele plikow — kazdy z osobnym blokiem `**Dane**`.

**WAZNE**: Dokladny format naglowka to `**Dane** (\`nazwa.txt\`):` — inne formaty (np. `Liczby pierwsze (\`plik.txt\`):`) nie zostana rozpoznane i kod C++ dostanie puste pliki.

### 2. Oczekiwany wynik

Wzorzec (regex: `\*\*Oczekiwany wynik\*\*\s*:\s*\n```\n...\n````):

```markdown
**Oczekiwany wynik**:
```
a) Suma: 123
b) Srednia: 45.67
```
```

### Jak dziala weryfikacja C++

1. Wyciaga kod C++ z `odpowiedz` (blok ` ```cpp `)
2. Wyciaga pliki wejsciowe i oczekiwany wynik z `tresc`
3. Tworzy katalog tymczasowy, zapisuje pliki wejsciowe
4. Podmienia sciezki plikow w kodzie C++ na bezwzgledne
5. Kompiluje: `g++ -std=c++17 -o prog source.cpp`
6. Uruchamia i przechwytuje stdout
7. Porownuje linia po linii z oczekiwanym wynikiem (tolerancja float: 0.01)

### Typowe przyczyny FAIL w C++

- **Zla kolejnosc wyjscia**: `map` iteruje alfabetycznie, `vector` wg indeksu — oczekiwany wynik musi pasowac do kodu
- **Bledna wartosc oczekiwana**: np. zle policzona suma, zly wynik algorytmu
- **Notatki robocze w wyniku**: oczekiwany wynik powinien zawierac TYLKO czyste linie wyjscia programu
- **Pluralizacja**: jezeli kod drukuje np. `" slow"` ale oczekiwany wynik ma `" slowa"` — trzeba zsynchronizowac

## Format `tresc` — SQL (katalogi 20-23)

### Definicja tabel

Wzorzec (regex: `Tabela \*\*(\w+)\*\*`):

```markdown
Tabela **Uczniowie**:

| id | imie | nazwisko | klasa |
|----|------|----------|-------|
| 1 | Anna | Kowalska | 3A |
| 2 | Jan | Nowak | 2B |
```

Mozna definiowac wiele tabel. Typy kolumn sa inferowane automatycznie (INTEGER/REAL/TEXT).

### Zapytanie SQL w `odpowiedz`

Blok ` ```sql `:

```markdown
**Zapytanie SQL:**
```sql
SELECT imie, nazwisko
FROM Uczniowie
WHERE klasa = '3A';
```
```

### Tabela wynikowa w `odpowiedz`

**Ostatnia** tabela markdown w `odpowiedz` (bez znakow ✓/✗) jest traktowana jako oczekiwany wynik:

```markdown
| imie | nazwisko |
|------|----------|
| Anna | Kowalska |
| Jan | Nowak |
```

Tabele z ✓/✗ (tabele weryfikacyjne) sa automatycznie pomijane.

### Multi-query

Jezeli SQL zawiera kilka SELECT-ow rozdzielonych `;`, a `odpowiedz` zawiera kilka tabel wynikowych — weryfikator dopasowuje je 1:1 w kolejnosci.

### Typowe przyczyny FAIL w SQL

- **Bledna wartosc LENGTH()**: policzyc recznie znaki (wlacznie ze spacjami!)
- **Zla kolejnosc wierszy**: ORDER BY musi byc zgodny z tabela wynikowa
- **Typ kolumny**: weryfikator inferuje typy — `3.0` zostanie porownane jako `3` (int)

## Format `tresc` — konwersje systemow (katalog 05)

Weryfikator szuka w `odpowiedz` wzorcow konwersji:

```
45(10) = 101101(2)      — decimal → binary
101101(2) = 45           — binary → decimal
B6(16) = 182             — hex → decimal
347(8) = 231             — octal → decimal
```

Weryfikacja: Python `int(digits, base)` porownuje wartosci.

## Uruchamianie weryfikacji

```bash
# Calosc (583 cwiczen):
python3 analiza/cwiczenia/verify/verify_all.py

# Jeden katalog:
python3 analiza/cwiczenia/verify/verify_all.py --file 07_cyfry_liczby --verbose

# Jedna kategoria:
python3 analiza/cwiczenia/verify/verify_all.py --category SQL

# Jedno cwiczenie:
python3 analiza/cwiczenia/verify/verify_all.py --id 23.4 --verbose
```

Raporty: `verify/report/verification_report.{md,json}`

Docelowy wynik: **530 PASS, 0 FAIL, 0 ERROR, 407 MANUAL_REVIEW**.

## Dodawanie nowego cwiczenia — checklist

1. Utworz plik `{id}.json` w odpowiednim katalogu (np. `json/07_cyfry_liczby/7.11.json`)
2. Nadaj `id` = `"NN.M"` (kolejny numer — sprawdz `_meta.json` po ostatni istniejacy!)
3. Dodaj wpis do `_meta.json` w tablicy `cwiczenia`: `{"id": "7.11", "trudnosc": "...", "punkty": N, "tagi": [...]}`
4. Zaktualizuj `punkty_lacznie` w `_meta.json` (musi = suma wszystkich punktow)
5. Dla IMPLEMENTACJA (07-14):
   - W `tresc`: dane wejsciowe w formacie `**Dane** (\`plik.txt\`):` + blok kodu
   - W `tresc`: oczekiwany wynik w formacie `**Oczekiwany wynik**:` + blok kodu
   - W `odpowiedz`: dzialajacy kod C++ w bloku ` ```cpp `
   - Upewnij sie ze wynik kodu **dokladnie** pasuje do oczekiwanego (kolejnosc, formatowanie, liczby)
6. Dla SQL (20-23):
   - W `tresc`: tabele w formacie `Tabela **Nazwa**:` + tabela markdown
   - W `odpowiedz`: zapytanie w bloku ` ```sql ` + ostatnia tabela markdown = oczekiwany wynik
   - Tabele weryfikacyjne (z ✓/✗) sa ignorowane
7. Uruchom: `python3 analiza/cwiczenia/verify/verify_all.py --id NN.M --verbose`
8. Popraw do PASS

## Walidacja schema JSON (standalone)

```bash
# Wszystkie katalogi:
python3 analiza/cwiczenia/validate_json.py

# Jeden katalog (prefix match):
python3 analiza/cwiczenia/validate_json.py --file 07_cyfry
```

Standalone walidator (nie wymaga plikow MD). Sprawdza:
- `_meta.json`: wymagane pola naglowka: `typ`, `nazwa`, `kategoria`, `czestotliwosc`
- `punkty_lacznie` = suma `punkty` z cwiczen w `_meta.json`
- `tagi_globalne` — niepuste, kazdy tag istnieje w rejestrze (`tagi_rejestr.json`)
- Spojnosc `_meta.json` z plikami na dysku (kazdy wpis ma plik i odwrotnie)
- Per cwiczenie: format `id`, `trudnosc`, `punkty` (1-10), `zrodlo`, `tresc` (>50 zn.), `wskazowki` (3 elementy), `odpowiedz` (>50 zn.), `typowe_bledy` (>=1 z `opis` i `kara`)
- `tagi` cwiczenia — podzbiór `tagi_globalne`, kazdy w rejestrze
- Zgodnosc pol miedzy `_meta.json` a plikiem cwiczenia (trudnosc, punkty, tagi)
- Reguly per kategoria: IMPLEMENTACJA wymaga ` ```cpp ` i `**Dane**`, SQL wymaga ` ```sql ` i `Tabela **`

## Centralny rejestr tagow

Plik `tagi_rejestr.json` zawiera posortowana liste wszystkich 290 dozwolonych tagow.

Walidator JSON odrzuci tagi spoza rejestru — zapobiega literowkom i niespojnosci.

Aby dodac nowy tag: dopisz go do tablicy `tagi` w `tagi_rejestr.json` (utrzymuj sortowanie alfabetyczne).

## Sanity checki dla MANUAL_REVIEW

Weryfikator `manual_sanity` (katalogi 01-04, 06, 15-19) oproc oznaczenia MANUAL_REVIEW wykonuje sanity checki:

**Uniwersalne** (wszystkie):
- `odpowiedz` >100 znakow
- `odpowiedz` zawiera tresc strukturalna (tabela/kod/P-F/lista/bold)

**Per typ**:
- 01 (sledzenie): tabela, drzewo (blok kodu) lub lista krokow
- 02 (projektowanie): blok kodu
- 03 (analiza): terminologia analityczna (O(), zlozonosc, porownani, krok...)
- 04 (test P/F): liczba odpowiedzi PRAWDA/FALSZ >= liczba pytan
- 06 (bezpieczenstwo): odpowiedz >= 200 znakow
- 15, 16, 18 (arkusz formuly): formula `=FUNKCJA(...)` lub `=komorka+komorka`
- 17 (wykresy): slowo "wykres" w odpowiedzi
- 19 (transformacja): odpowiedz >= 100 znakow

Sanity OK → MANUAL_REVIEW, Sanity FAIL → FAIL.
