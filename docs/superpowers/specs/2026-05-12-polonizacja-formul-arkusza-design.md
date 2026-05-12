# Polonizacja formuł arkusza (MS Excel PL)

**Data:** 2026-05-12
**Status:** Design — gotowy do planu implementacji
**Autor:** Tomek + Claude

## Cel

Wszystkie formuły arkusza kalkulacyjnego w materiałach dydaktycznych mają używać oficjalnych polskich nazw funkcji MS Excel (zgodnie z konwencją CKE), zamiast angielskich. Powód: uczeń przygotowujący się do matury widzi w naszych materiałach `VLOOKUP`, a w polskim Excelu (i na egzaminie CKE) działa wyłącznie `WYSZUKAJ.PIONOWO`. Niespójność dezorientuje i wymusza dodatkowe tłumaczenie w głowie.

Decyzje wstępne (wybrane przez użytkownika):
- **Konwencja**: MS Excel PL. Większość nazw zgodna z LibreOffice Calc PL; realne rozbieżności (np. Excel `MOD` vs LibreOffice `MOD.CZ`) rozstrzygamy na korzyść Excel.
- **Zakres**: tylko pliki dydaktyczne arkusza, NIE matura_*.json (cytaty CKE), NIE SQL/C++
- **Tagi**: rename + migracja schema v9 (zachowanie historii FSRS uczniów)
- **Pliki "mieszane"**: wszystkie 8 zawiera tłumaczenie (zweryfikowane: każde wystąpienie jednoznacznie arkuszowe)
- **Separator argumentów**: normalizujemy `,` → `;` (polski locale)

## Mapa tłumaczeń (46 funkcji, zweryfikowana przez 2 niezależne agenty research)

Pełna mapa EN → PL (MS Excel) zweryfikowana przy MS Support PL.

| Angielska | Polska MS Excel | Uwagi |
|---|---|---|
| VLOOKUP | WYSZUKAJ.PIONOWO | |
| HLOOKUP | WYSZUKAJ.POZIOMO | |
| IF | JEŻELI | |
| AND | ORAZ | |
| OR | LUB | |
| NOT | NIE | |
| IFERROR | JEŻELI.BŁĄD | |
| SUM | SUMA | |
| SUMIF | SUMA.JEŻELI | |
| SUMIFS | SUMA.WARUNKÓW | NIE `SUMY.WARUNKÓW` |
| SUMPRODUCT | SUMA.ILOCZYNÓW | |
| AVERAGE | ŚREDNIA | |
| AVERAGEIF | ŚREDNIA.JEŻELI | |
| AVERAGEIFS | ŚREDNIA.WARUNKÓW | NIE `ŚREDNIE.WARUNKÓW` |
| COUNT | ILE.LICZB | |
| COUNTA | ILE.NIEPUSTYCH | |
| COUNTIF | LICZ.JEŻELI | |
| COUNTIFS | LICZ.WARUNKI | |
| MAXIFS | MAKS.WARUNKÓW | |
| MINIFS | MIN.WARUNKÓW | |
| ROUND | ZAOKR | |
| ROUNDUP | ZAOKR.GÓRA | NIE `ZAOKR.W.GÓRĘ` (= CEILING) |
| ROUNDDOWN | ZAOKR.DÓŁ | NIE `ZAOKR.W.DÓŁ` (= FLOOR) |
| INT | ZAOKR.DO.CAŁK | |
| ABS | MODUŁ.LICZBY | |
| LEFT | LEWY | |
| RIGHT | PRAWY | |
| MID | FRAGMENT.TEKSTU | |
| LEN | DŁ | |
| TRIM | USUŃ.ZBĘDNE.ODSTĘPY | |
| UPPER | LITERY.WIELKIE | |
| LOWER | LITERY.MAŁE | |
| CONCATENATE | ZŁĄCZ.TEKSTY | Legacy. Modern `CONCAT` = `ZŁĄCZ.TEKST` (bez Y) |
| FIND | ZNAJDŹ | |
| SEARCH | SZUKAJ.TEKST | |
| SUBSTITUTE | PODSTAW | |
| REPLACE | ZASTĄP | |
| INDEX | INDEKS | |
| MATCH | PODAJ.POZYCJĘ | |
| RANK | POZYCJA | Legacy. Modern: `POZYCJA.NAJW` (RANK.EQ) |
| YEAR | ROK | |
| MONTH | MIESIĄC | |
| DAY | DZIEŃ | |
| TODAY | DZIŚ | |
| NOW | TERAZ | |
| DATE | DATA | |
| MAX | MAX | bez zmian |
| MIN | MIN | bez zmian |
| MOD | MOD | bez zmian |

## Zakres plików

### Warstwa 1: Pliki MD dedykowane arkuszowi (3)

Pełne globalne tłumaczenie — cały plik jest o arkuszu, zero ryzyka.

- `analiza/szablony/arkusz_formuly.md` (588 linii)
- `analiza/cheatsheets/cheatsheet_arkusz.md` (126 linii)
- `analiza/rozwiazania_wzorcowe/arkusz_kalkulacyjny.md` (351 linii)

### Warstwa 2: Pliki MD mieszane (6) — DETERMINISTYCZNA LISTA ZAMIAN

Audit potwierdził: WSZYSTKIE wystąpienia w plikach mieszanych to **funkcje spreadsheet-specific** (`SUMIF`, `SUMIFS`, `COUNTIF`, `COUNTIFS`, `AVERAGEIF`, `AVERAGEIFS`, `COUNTA`) — żadne nie zawiera generycznych nazw (`SUM`, `COUNT`, `IF`, `AND`, `OR`, `MAX`, `MIN`), które mogłyby kolidować z SQL/C++.

**Strategia**: dla plików mieszanych skrypt stosuje **TYLKO wyliczoną listę poniżej** (nie ogólny replace). Każda linia ma jawnie podaną starą i nową treść — skrypt walidacyjnie sprawdza że stary fragment istnieje w danej linii zanim zamieni.

Listę można też przekleić do `polonizacja_warstwa2.json` jako wkład skryptu.

---

**Plik 1: `analiza/cheatsheets/debug_checklist.md`** (3 zmiany)

```
LINIA 92:
  STARY:   - [ ] SUMIF: `=SUMIF(zakres_warunku; warunek; zakres_sumy)` — 3 argumenty!
  NOWY:    - [ ] SUMA.JEŻELI: `=SUMA.JEŻELI(zakres_warunku; warunek; zakres_sumy)` — 3 argumenty!

LINIA 93:
  STARY:   - [ ] SUMIFS: `=SUMIFS(zakres_sumy; zakres_war1; war1; zakres_war2; war2)` — suma PIERWSZA!
  NOWY:    - [ ] SUMA.WARUNKÓW: `=SUMA.WARUNKÓW(zakres_sumy; zakres_war1; war1; zakres_war2; war2)` — suma PIERWSZA!

LINIA 94:
  STARY:   - [ ] COUNTIF: `=COUNTIF(zakres; warunek)` — 2 argumenty
  NOWY:    - [ ] LICZ.JEŻELI: `=LICZ.JEŻELI(zakres; warunek)` — 2 argumenty
```

---

**Plik 2: `analiza/cheatsheets/przed_egzaminem.md`** (2 zmiany)

```
LINIA 55:
  STARY:   - [ ] `=SUMIF(zakres_war; warunek; zakres_sum)` / SUMIFS
  NOWY:    - [ ] `=SUMA.JEŻELI(zakres_war; warunek; zakres_sum)` / SUMA.WARUNKÓW

LINIA 56:
  STARY:   - [ ] `=COUNTIF(zakres; warunek)`
  NOWY:    - [ ] `=LICZ.JEŻELI(zakres; warunek)`
```

---

**Plik 3: `analiza/szablony/wzorce_2015.md`** (6 zmian + normalizacja separatora w 5 z nich)

```
LINIA 159:
  STARY:   =SUMIF(A:A,"A",B:B)
  NOWY:    =SUMA.JEŻELI(A:A;"A";B:B)              # + separator , → ;

LINIA 162:
  STARY:   =SUMIF(A:A,"B",B:B)
  NOWY:    =SUMA.JEŻELI(A:A;"B";B:B)              # + separator

LINIA 168:
  STARY:   =COUNTA(A2:A20)
  NOWY:    =ILE.NIEPUSTYCH(A2:A20)

LINIA 171:
  STARY:   =COUNTIF(A:A,"A")
  NOWY:    =LICZ.JEŻELI(A:A;"A")                   # + separator

LINIA 183:
  STARY:   =COUNTIF(E:E, ">0")
  NOWY:    =LICZ.JEŻELI(E:E;">0")                  # + separator

LINIA 269 (nagłówek pułapki):
  STARY:   ### Pułapka 8: SUMIF w Excel
  NOWY:    ### Pułapka 8: SUMA.JEŻELI w Excel

LINIA 272:
  STARY:   **Błąd:** Użycie SUM zamiast SUMIF dla warunkowego sumowania
  NOWY:    **Błąd:** Użycie SUMA zamiast SUMA.JEŻELI dla warunkowego sumowania

LINIA 273:
  STARY:   **Poprawnie:** =SUMIF(zakres_kryterium, kryterium, zakres_do_zsumowania)
  NOWY:    **Poprawnie:** =SUMA.JEŻELI(zakres_kryterium; kryterium; zakres_do_zsumowania)   # + separator
```

UWAGA: Linia 272 zawiera GENERYCZNE `SUM` (bez nawiasu) — w tym konkretnym kontekście jednoznacznie nazwa funkcji Excela (pisana wielkimi literami obok `SUMIF`). Tłumaczymy ręcznie, nie regexem.

---

**Plik 4: `analiza/PRZEWODNIK_UCZNIA.md`** (1 zmiana)

```
LINIA 185:
  STARY:   8. `15_agregacja_warunkowa.md` — arkusz: SUMIFS/COUNTIFS (38 pkt)
  NOWY:    8. `15_agregacja_warunkowa.md` — arkusz: SUMA.WARUNKÓW/LICZ.WARUNKI (38 pkt)
```

---

**Plik 5: `analiza/strategia_egzaminacyjna.md`** (3 zmiany)

```
LINIA 613:
  STARY:   2. SUMIF(zakres_kryt, kryterium, zakres_sum)
  NOWY:    2. SUMA.JEŻELI(zakres_kryt; kryterium; zakres_sum)        # + separator

LINIA 614:
  STARY:   3. Dla wielu warunkow: SUMIFS (warunki w parach zakres+kryterium)
  NOWY:    3. Dla wielu warunkow: SUMA.WARUNKÓW (warunki w parach zakres+kryterium)

LINIA 617:
  STARY:   - Pomylenie SUMIF z SUMIFS (kolejnosc argumentow!)
  NOWY:    - Pomylenie SUMA.JEŻELI z SUMA.WARUNKÓW (kolejnosc argumentow!)
```

---

**Plik 6: `analiza/drzewo_decyzyjne.md`** (3 zmiany)

```
LINIA 146:
  STARY:   | "suma/liczba wg warunku" | SUMIF / COUNTIF | `=SUMIFS(D:D; B:B; "X"; C:C; ">100")` |
  NOWY:    | "suma/liczba wg warunku" | SUMA.JEŻELI / LICZ.JEŻELI | `=SUMA.WARUNKÓW(D:D; B:B; "X"; C:C; ">100")` |

LINIA 147:
  STARY:   | "srednia wg warunku" | AVERAGEIF(S) | `=AVERAGEIFS(D:D; B:B; "X")` |
  NOWY:    | "srednia wg warunku" | ŚREDNIA.JEŻELI(.WARUNKÓW) | `=ŚREDNIA.WARUNKÓW(D:D; B:B; "X")` |

LINIA 148:
  STARY:   | "zlicz unikalne" | COUNTIF + pomocnicza | `=1/COUNTIF(zakres; wartosc)` -> SUMA |
  NOWY:    | "zlicz unikalne" | LICZ.JEŻELI + pomocnicza | `=1/LICZ.JEŻELI(zakres; wartosc)` -> SUMA |
```

UWAGA: na końcu linii 148 jest już słowo "SUMA" (po polsku) — pozostaje bez zmian.

---

**Walidacja warstwy 2 (skrypt)**:

Dla każdego pliku w warstwie 2 skrypt:
1. **Pre-check**: dla każdej zaplanowanej zmiany sprawdza że `STARY` fragment dokładnie pojawia się w `LINII X` pliku. Jeśli linia się przesunęła lub treść zmieniła — błąd, nie zamienia.
2. **Apply**: stosuje zamianę dokładnie `STARY` → `NOWY` (Python `str.replace` z asercją `count == 1` w obrębie linii).
3. **Post-check**: po wszystkich zamianach grep'em weryfikuje że w pliku **nie ma** już starych nazw spreadsheet-specific (`SUMIF\(`, `SUMIFS\(`, `COUNTIF\(`, `COUNTIFS\(`, `AVERAGEIF\(`, `AVERAGEIFS\(`, `COUNTA\(`).
4. **Numery linii** są wskazówką — pre-check używa pełnego fragmentu `STARY`, więc skrypt zadziała nawet jeśli linia przesunie się o ±5.

Linie generyczne (`SUM`/`IF`/`COUNT`/`AND`/`OR`) w plikach mieszanych zostają **nietknięte** poza explicite wymienionym przypadkiem w `wzorce_2015.md:272` (`SUM` jako nazwa funkcji w jednoznacznym kontekście).

### Warstwa 3: JSON ćwiczeń arkuszowych (4 katalogi)

Pełne tłumaczenie pól `tresc`, `odpowiedz`, `wskazowki[].opis`, `typowe_bledy[].opis`, `weryfikacja_szczegolowa`, `tagi`, oraz `tagi_globalne` w `_meta.json`.

- `analiza/cwiczenia/json/15_agregacja_warunkowa/` (40 + `_meta.json`)
- `analiza/cwiczenia/json/16_symulacja/` (~34 + `_meta.json`)
- `analiza/cwiczenia/json/18_agregacja_podstawowa/` (~40 + `_meta.json`)
- `analiza/cwiczenia/json/19_transformacja/` (~40 + `_meta.json`)

### Warstwa 4: Rejestry (2)

- `analiza/cwiczenia/json/tagi_rejestr.json` — rename dokładnie 7 wpisów (`VLOOKUP`, `SUMIF`, `SUMIFS`, `COUNTIF`, `COUNTIFS`, `AVERAGEIF`, `AVERAGEIFS` → polskie). Reszta pliku nietknięta.
- `analiza/json/algorytmy_rejestr.json` — 1 linia tekstu opisowego (pole `definicja` z `SUMIFS, COUNTIFS, AVERAGEIFS`)

### Warstwa 5: CLI — schema migration v9

- `analiza/cli/database.go` — dodać migrację dla tabeli `progress_tagi` (7 rename, INSERT OR REPLACE pattern dla bezpieczeństwa kolizji)

### NIE TYKAMY (jawnie wykluczone)

- Wszystkie `analiza/json/matura_*.json` (30 plików) — cytaty CKE, dodatkowo zawierają SQL keywords (SUM/COUNT/AVG)
- Wszystkie pliki w `analiza/szablony/{sql,cpp}_*.md`
- Wszystkie pliki w `analiza/rozwiazania_wzorcowe/{sql_zapytania,implementacja_cpp,teoria_algorytmy}.md`
- Wszystkie katalogi `analiza/cwiczenia/json/01_*` do `14_*` (C++) oraz `20_*` do `23_*` (SQL)
- `analiza/cwiczenia/json/17_wykres/` (sprawdzone: zero wystąpień)
- Pliki PDF (`arkusz.pdf`, `odpowiedzi.pdf`)
- `analiza/cli/` poza `database.go` (importer.go, types.go itd. nietknięte)
- `analiza/cwiczenia/wg_typu/*.md` (generowane z JSON — odświeży się przy regeneracji)

## Strategia podmian

### Skrypt `analiza/scripts/polonizuj_formuly.py`

Idempotentny skrypt z dwoma trybami:

```
python3 analiza/scripts/polonizuj_formuly.py --dry-run    # pokazuje diff, nic nie zmienia
python3 analiza/scripts/polonizuj_formuly.py --apply      # robi zmiany
python3 analiza/scripts/polonizuj_formuly.py --apply --files-only "warstwa_1,warstwa_2"  # selektywnie
```

**Reguły podmian** (w kolejności):

1. **Długie nazwy najpierw** — `SUMIFS` zanim `SUMIF`, `COUNTIFS` zanim `COUNTIF`, `AVERAGEIFS` zanim `AVERAGEIF` (uniknięcie błędnej podmiany prefiksu)
2. **Granica słowa** — regex `\b` lub `(?<![A-Z_])NAZWA(?=\s*\()` żeby nie złapać `MYSUM(` itp.
3. **Tylko z otwierającym `(`** — `SUM(` tak, ale `SUMA jest największa` zostaje (no paren = no formula)
4. **Case-sensitive** — w C++ `sum()` lowercase nie zostanie zmienione

**Whitelist plików** zaszyta w skrypcie (lista z warstw 1+2+3+4). Skrypt **odmawia** edycji pliku spoza listy.

**Per-plik logika** dla warstw 1+3 (czysto arkuszowe): podmiana globalna stosująca pełną mapę 46 funkcji. Bezpieczna bo cały plik/katalog jest dedykowany arkuszowi.

**Per-plik logika dla warstwy 2** (mieszane MD): **NIE używamy heurystyki ani globalnego replace**. Skrypt stosuje **deterministyczną listę zamian** z sekcji "Warstwa 2: Pliki MD mieszane" powyżej. Dla każdej zaplanowanej zmiany:
- Pre-check: stary fragment dokładnie istnieje w pliku (assertion `count == 1` w odpowiedniej linii)
- Apply: `str.replace(STARY, NOWY, 1)`
- Post-check: po wszystkich zamianach w pliku, grep'em weryfikuje że żadna ze spreadsheet-specific nazw (`SUMIF`, `SUMIFS`, `COUNTIF`, `COUNTIFS`, `AVERAGEIF`, `AVERAGEIFS`, `COUNTA`) już nie występuje

Jeśli pre-check się nie powiedzie (ktoś zmienił plik między audytem a uruchomieniem skryptu) — skrypt przerywa z błędem i listuje brakujące fragmenty do ręcznej weryfikacji.

Generyczne nazwy (`SUM`, `IF`, `COUNT`, `AND`, `OR`, `MAX`, `MIN`, `ROUND` itd.) w plikach warstwy 2 zostają **nietknięte** — z jednym wyjątkiem (`wzorce_2015.md:272`), gdzie `SUM` jest jawnie wpisane do listy zamian jako tłumaczenie nazwy funkcji w jednoznacznym kontekście.

### Separator argumentów `,` → `;`

W obrębie formuły arkusza (`=NAZWA(...)` lub `NAZWA(...)` w kontekście arkusza) — po podmianie nazwy normalizujemy `,` na `;`.

UWAGA: separator dziesiętny w Polsce to `,` (np. `1,5` zamiast `1.5`). Skrypt musi rozróżnić:
- `=SUMIF(A:A, "X", B:B)` — przecinki to separatory argumentów → `;`
- `=ROUND(3,14; 2)` — `3,14` to liczba dziesiętna → zostaje

Heurystyka: przecinek jest separatorem argumentów jeśli **NIE jest otoczony cyframi z obu stron**. Konserwatywnie — jeśli wątpliwe, log do review file i nie zmieniaj.

### Migracja schema v9 (CLI)

W `analiza/cli/database.go`, w sekcji migracji schema (już istnieje pattern dla v8):

```go
// v9: Polish formula tag rename (Polonizacja formuł arkusza, 2026-05-12)
if currentVersion < 9 {
    renames := []struct{ old, new string }{
        {"VLOOKUP",    "WYSZUKAJ.PIONOWO"},
        {"SUMIF",      "SUMA.JEŻELI"},
        {"SUMIFS",     "SUMA.WARUNKÓW"},
        {"COUNTIF",    "LICZ.JEŻELI"},
        {"COUNTIFS",   "LICZ.WARUNKI"},
        {"AVERAGEIF",  "ŚREDNIA.JEŻELI"},
        {"AVERAGEIFS", "ŚREDNIA.WARUNKÓW"},
    }
    for _, r := range renames {
        // Jeśli nowa nazwa już istnieje (kolizja), preferuj merge zachowując max(stability, reps)
        if _, err := tx.Exec(`
            INSERT INTO progress_tagi(tag, stability, difficulty, reps, lapses, state, last_review, nastepna_powtorka, poziom)
            SELECT ?, stability, difficulty, reps, lapses, state, last_review, nastepna_powtorka, poziom
            FROM progress_tagi WHERE tag = ?
            ON CONFLICT(tag) DO UPDATE SET
                stability = max(progress_tagi.stability, excluded.stability),
                reps = max(progress_tagi.reps, excluded.reps)
        `, r.new, r.old); err != nil {
            return err
        }
        if _, err := tx.Exec(`DELETE FROM progress_tagi WHERE tag = ?`, r.old); err != nil {
            return err
        }
    }
    // Update wersji schema do 9
}
```

Plus sprawdzić tabelę `progress_bledy` — czy `blad_kod` może zawierać stare nazwy tagów. Jeśli tak, dodać analogiczną migrację.

## Walidacja (6 gates)

Po każdej fazie implementacji uruchamiamy:

1. **Schema lint**: `python3 analiza/cwiczenia/validate_json.py` — żaden tag w `tagi_globalne` ani `tagi` nie używa starej angielskiej nazwy. Expected: 0 ERRORS.
2. **Content verification**: `python3 analiza/cwiczenia/verify/verify_all.py` — nadal 530 PASS / 0 FAIL / 407 MANUAL.
3. **QA suite**: `./test_qa.sh` — 180 testów (L1-L7) pass.
4. **Grep guard**: `grep -rE '\b(VLOOKUP|HLOOKUP|SUMIF|SUMIFS|COUNTIF|COUNTIFS|AVERAGEIF|AVERAGEIFS|CONCATENATE|IFERROR|SUMPRODUCT)\s*\(' <whitelist>` — zero hits w plikach z whitelisty.
5. **CLI smoke**: po rebuildzie (`cd analiza/cli && ./build.sh`):
   - `./matura exercise get --typ agregacja_warunkowa` — zwraca ćwiczenie z polskimi formułami
   - `./matura cke get --typ agregacja_warunkowa` — zwraca cytat CKE (BEZ zmiany, sprawdzenie że matura_*.json nietknięte)
   - `./matura progress status` — działa po migracji v9
6. **Review file**: przejrzeć `_polonizacja_review.txt` — wątpliwe przypadki (oczekiwane: <10). Każdy ręcznie rozstrzygnąć.

## Bonus — pułapki polskich nazw (cheatsheet_arkusz.md)

Dodać nową sekcję na końcu `cheatsheet_arkusz.md`:

```markdown
## Pułapki polskich nazw funkcji — uważaj!

| Funkcja | Polska nazwa | Pułapka |
|---|---|---|
| `ROUNDUP` | `ZAOKR.GÓRA` | NIE `ZAOKR.W.GÓRĘ` (to **CEILING** — zaokrąglanie do wielokrotności) |
| `ROUNDDOWN` | `ZAOKR.DÓŁ` | NIE `ZAOKR.W.DÓŁ` (to **FLOOR**) |
| `CONCATENATE` | `ZŁĄCZ.TEKSTY` (z "Y") | `ZŁĄCZ.TEKST` (bez Y) to nowsza `CONCAT` |
| `SUBSTITUTE` vs `REPLACE` | `PODSTAW` vs `ZASTĄP` | `PODSTAW` = po tekście, `ZASTĄP` = po pozycji znaków |
| `MAX` / `MIN` / `MOD` | bez zmian | Excel PL używa tych samych nazw — NIE tłumaczyć |
| Separator argumentów | `;` | Polski locale używa średnika, NIE przecinka |
| Separator dziesiętny | `,` | Polski locale używa przecinka (np. `3,14` zamiast `3.14`) |
```

## Plan implementacji (fazy)

Po zatwierdzeniu spec — przejście do writing-plans skill, który rozbije to na konkretne kroki. Wstępny zarys:

1. **Faza 0**: Skrypt `polonizuj_formuly.py` w trybie dry-run + mapa w osobnym pliku JSON
2. **Faza 1**: Warstwa 1 (3 pliki MD dedykowane) + walidacja
3. **Faza 2**: Warstwa 3 (89 plików JSON ćwiczeń) + Warstwa 4 (rejestry) + walidacja
4. **Faza 3**: Warstwa 2 (6 plików MD mieszanych) — manual review wątpliwych przypadków
5. **Faza 4**: Warstwa 5 (CLI schema v9) + rebuild + reimport
6. **Faza 5**: Bonus — sekcja pułapek w cheatsheet_arkusz.md
7. **Faza 6**: Full validation (6 gates) + commit

## Wpływ na progres ucznia

Po migracji v9:
- Uczniowie z istniejącym `progress.db` — bezszwowy upgrade (FSRS state zachowany przez SQL UPDATE)
- Komunikaty CLI typu `WARN_LEECH: Tag 'X'` — pokażą polskie nazwy (`Tag 'LICZ.JEŻELI'`)
- Filtrowanie ćwiczeń po tagu (`exercise get --tag SUMIF` przestanie działać; trzeba `--tag SUMA.JEŻELI`)
- ⚠️ Skrypty/aliasy używające starych nazw (jeśli istnieją w skill `matura` lub w dokumentach typu `JAK_UZYWAC_KOREPETYTORA.md`) muszą być zaktualizowane

## Ryzyka i mitygacje

| Ryzyko | Prawdopodobieństwo | Mitygacja |
|---|---|---|
| False positive w pliku mieszanym | Niskie | Whitelist + per-plik review file |
| Pomyłka separatora `,` (dziesiętny vs argumentowy) | Średnie | Konserwatywna heurystyka + review file |
| Kolizja tagów w migracji v9 | Bardzo niskie | INSERT OR REPLACE + max() merge |
| Złamanie validate_json.py | Średnie | Aktualizacja `tagi_rejestr.json` jako pierwsza |
| Złamanie matura_*.json (nie powinno się stać) | Bardzo niskie | Skrypt jawnie odmawia edycji plików spoza whitelisty |
| Reimport CLI z nieaktualnymi tagami | Średnie | Test L1 (CLI smoke) wykryje natychmiast |

## Definition of Done

- [ ] Skrypt `polonizuj_formuly.py` istnieje i przechodzi `--dry-run` bez błędów
- [ ] Wszystkie pliki w whitelist mają polskie nazwy funkcji
- [ ] `validate_json.py` zwraca 0 ERRORS
- [ ] `verify_all.py` zwraca 530 PASS / 0 FAIL / 407 MANUAL
- [ ] `test_qa.sh` przechodzi (180 testów)
- [ ] Grep guard zwraca 0 hits dla angielskich nazw w whitelist
- [ ] `tagi_rejestr.json` zawiera polskie tagi (7 zmienionych, 283 nietknięte)
- [ ] CLI po `build.sh` reimportuje matura.db bez błędów
- [ ] `progress.db` po migracji v9 zachowuje FSRS state (test: ręcznie wstawić wpis `'COUNTIF'`, uruchomić migrację, sprawdzić że jest jako `'LICZ.JEŻELI'`)
- [ ] Nowa sekcja "Pułapki polskich nazw" w `cheatsheet_arkusz.md`
- [ ] Wszystkie zmiany committed w jednym (lub maks. 2-3) PR
- [ ] `MEMORY.md` zaktualizowane o decyzję polonizacji i jej zakres
