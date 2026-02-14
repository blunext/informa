---
name: generate-exercises
description: Generate new exercises for matura exam preparation. Use when creating practice problems, expanding exercise sets, or generating exercises for a specific topic/type.
argument-hint: [typ_zadania or file_name] [count]
---

# Generowanie cwiczen maturalnych

Skill do generowania nowych cwiczen do bazy JSON. Prowadzi caly workflow: planowanie, generacja, walidacja, weryfikacja.

## Krok 1: Ustal parametry

Zapytaj uzytkownika (jezeli nie podal):
- **Plik docelowy**: ktory z 23 typow? (np. `07_cyfry_liczby`, `20_sql_group_by`)
- **Ile cwiczen**: domyslnie 5
- **Trudnosc**: `latwe` / `srednie` / `srednie-trudne` / `trudne` / mix

## Krok 2: Zaladuj kontekst

Przed generacja ZAWSZE przeczytaj:

1. **Istniejace cwiczenia** z docelowego pliku JSON:
   ```
   analiza/cwiczenia/json/NN_nazwa.json
   ```
2. **Schema i format danych** (kluczowe reguly formatowania):
   ```
   analiza/cwiczenia/json/README.md
   ```
3. **Rejestr tagow** (dozwolone tagi):
   ```
   analiza/cwiczenia/json/tagi_rejestr.json
   ```
4. **Szablony i wzorce** odpowiednie dla kategorii:
   - IMPLEMENTACJA (07-14): `analiza/szablony/cpp_szablony.md`
   - SQL (20-23): `analiza/szablony/sql_szablony.md`
   - ARKUSZ (15-19): `analiza/szablony/arkusz_formuly.md`
   - TEORIA (01-06): `analiza/szablony/pseudokod_wzorce.md`

## Krok 3: Generuj cwiczenia

Dla kazdego cwiczenia generuj pelny obiekt JSON:

```json
{
  "id": "NN.M",
  "trudnosc": "srednie",
  "punkty": 3,
  "zrodlo": "Inspirowane matura 2023 zad. 4",
  "tagi": ["petla", "modulo"],
  "tresc": "...",
  "wskazowki": [
    "**Kierunek**: ...",
    "**Podejscie**: ...",
    "**Kluczowy krok**: ..."
  ],
  "odpowiedz": "...",
  "typowe_bledy": [
    { "opis": "**Opis bledu**: ...", "kara": "-1 pkt" }
  ]
}
```

### Reguly bezwzgledne

- **id**: format `NN.M` — NN = numer pliku, M = kolejny numer (sprawdz ostatni istniejacy!)
- **tagi**: TYLKO tagi z `tagi_rejestr.json`. Jezeli potrzebujesz nowego tagu — najpierw dodaj go do rejestru
- **tagi**: musza byc podzbiorem `tagi_globalne` pliku. Jezeli uzywasz nowego — dodaj tez do `tagi_globalne`
- **wskazowki**: dokladnie 3, kazda >10 znakow, wzorzec: Kierunek / Podejscie / Kluczowy krok
- **tresc**: >50 znakow, musi byc samodzielna (uczen nie potrzebuje dodatkowych informacji)
- **odpowiedz**: >50 znakow, pelne rozwiazanie (nie szkic!)
- **typowe_bledy**: minimum 1, kazdy z `opis` (bold prefix) i `kara` (format: `-N pkt`)
- **punkty_lacznie** w naglowku: MUSI byc zaktualizowane po dodaniu cwiczen (suma punktow)

### Reguly per kategoria

#### IMPLEMENTACJA (pliki 07-14)
- `tresc` MUSI zawierac dane wejsciowe w formacie:
  ```
  **Dane** (`plik.txt`):
  ```
  dane
  ```
  ```
- `tresc` MUSI zawierac oczekiwany wynik:
  ```
  **Oczekiwany wynik**:
  ```
  wynik
  ```
  ```
- `odpowiedz` MUSI zawierac dzialajacy kod C++ w bloku ` ```cpp `
- Wynik kodu MUSI dokladnie odpowiadac oczekiwanemu wynikowi (kolejnosc, formatowanie!)
- Testuj mentalnie: czy ten kod skompiluje sie i da dokladnie ten wynik?

#### SQL (pliki 20-23)
- `tresc` MUSI zawierac definicje tabel: `Tabela **Nazwa**:` + tabela markdown
- `odpowiedz` MUSI zawierac zapytanie w bloku ` ```sql `
- Ostatnia tabela markdown w `odpowiedz` (bez znakow weryfikacyjnych) = oczekiwany wynik
- Testuj mentalnie: czy to zapytanie na tych danych da te wyniki?

#### TEORIA/sledzenie (pliki 01-04)
- `odpowiedz` musi zawierac krokowe sledzenie (tabela, lista krokow, lub drzewo wywolan)
- Dla P/F (plik 04): odpowiedz musi zawierac PRAWDA/FALSZ lub **P**/**F** dla kazdego pytania

#### ARKUSZ (pliki 15-19)
- `odpowiedz` musi zawierac formuly arkuszowe (=FUNKCJA(...) lub =A1+B1)

## Krok 3.5: Weryfikacja merytoryczna (OBOWIAZKOWA dla TEORIA i ARKUSZ)

Dla IMPLEMENTACJA (07-14) i SQL (20-23) `verify_all.py` zweryfikuje automatycznie (kompilacja/exec). Ale dla TEORIA (01-04, 06) i ARKUSZ (15-19) musisz SAM zweryfikowac kazde cwiczenie PRZED wstawieniem do pliku.

### TEORIA — checklista weryfikacji

| Plik | Metoda | Co sprawdzic |
|------|--------|-------------|
| 01_sledzenie_algorytmu | Sledzenie krok po kroku | Tabele stanow zmiennych, wartosci koncowe, liczba wywolan rekurencyjnych |
| 02_projektowanie_algorytmu | Mentalna kompilacja pseudokodu/C++ | Poprawnosc algorytmu na podanych przykladach, edge cases, zgodnosc z ograniczeniami |
| 03_analiza_algorytmu | Weryfikacja rozumowania | O-notacja (policz iteracje!), kontrprzyklady (przelicz!), niezmienniki petli |
| 04_test_prawda_falsz | Sprawdzenie kazdego P/F vs fakty CS | Poprawnosc klasyfikacji, merytorycznosc uzasadnien |
| 06_teoria_bezpieczenstwa | Sprawdzenie poprawnosci faktycznej | Definicje, klasyfikacje atakow/malware, protokoly sieciowe |

**Szczegolowe instrukcje:**

1. **Sledzenie (01)**: Dla KAZDEGO przykladu w `odpowiedz` — recznie przejdz algorytm krok po kroku. Zweryfikuj tabele zmiennych, wartosci koncowe, liczbe wywolan. Jesli odpowiedz zawiera drzewo wywolan — policz wezly.

2. **Projektowanie (02)**: Przetestuj podany pseudokod/C++ mentalnie na KAZDYM przykladzie z `tresc`. Sprawdz edge cases (n=0, n=1, pusta tablica). Zweryfikuj, ze `odpowiedz` zawiera pelne, kompilowalne rozwiazanie.

3. **Analiza (03)**: Zweryfikuj zlozonosc — policz iteracje petli. Przy kontrprzykladach — przelicz je. Przy dowodach indukcyjnych — sprawdz baze i krok. Niezmienniki petli — zweryfikuj na 2-3 iteracjach.

4. **Test P/F (04)**: Dla kazdego zdania — zweryfikuj P/F niezaleznie (nie ufaj wygenerowanej odpowiedzi). Sprawdz uzasadnienia.

5. **Bezpieczenstwo (06)**: Sprawdz definicje vs ogolnie przyjeta wiedza CS. Protokoly — zweryfikuj akronimy i funkcje.

### ARKUSZ — checklista weryfikacji

| Plik | Metoda | Co sprawdzic |
|------|--------|-------------|
| 15_agregacja_warunkowa | Przelicz formuly na danych z tresci | SUMIFS/COUNTIFS — skladnia, zakresy, kryteria, wynik liczbowy |
| 16_symulacja | Sledz wiersz po wierszu | Akumulatory, procent skladany, zaleznosci miedzy komorkami |
| 17_wykres | Sprawdz specyfikacje | Typ wykresu vs dane, osie, ewentualne obliczenia |
| 18_agregacja_podstawowa | Przelicz statystyki | SUM/AVERAGE/MAX/MIN/MEDIAN/COUNT — policz recznie na danych |
| 19_transformacja | Weryfikuj formuly | VLOOKUP argumenty, tabele krzyzowe, odniesienia mieszane ($) |

**Szczegolowe instrukcje:**

1. **Agregacja warunkowa (15)**: Wypisz dane spelniajace kryteria, policz recznie, porownaj z odpowiedzia. Sprawdz skladnie: `=SUMIFS(zakres_sum; zakres_kryterium; kryterium)` — polskie separatory (`;`).

2. **Symulacja (16)**: Zbuduj mentalny "arkusz": wiersze 1-5 minimum. Sledz zaleznosci komorek. Zweryfikuj formule w kontekscie kopiowania ($ vs relative). UWAGA na zaokraglenia (ZAOKR, INT) — liczbowo przelicz krok po kroku.

3. **Wykres (17)**: Sprawdz czy typ wykresu pasuje do danych (kolumnowy vs liniowy vs kolowy). Zweryfikuj obliczenia numeryczne jesli sa.

4. **Agregacja podstawowa (18)**: Policz recznie: SUM (dodaj), AVERAGE (SUM/COUNT), MAX/MIN (znajdz), MEDIAN (posortuj, srodkowy). Porownaj z odpowiedzia.

5. **Transformacja (19)**: Sprawdz VLOOKUP: `=VLOOKUP(szukana; tabela; nr_kolumny; FALSZ)`. Zweryfikuj odniesienia mieszane przy kopiowaniu formuly.

### Uniwersalny checklist (WSZYSTKIE typy)

Dla kazdego wygenerowanego cwiczenia sprawdz:
- `tresc` — jasna, kompletna, samodzielna (uczen nie potrzebuje dodatkowych info)
- `odpowiedz` — merytorycznie poprawna (przeliczona/prześledzona)
- `wskazowki` — 3 hinty, progresja (kierunek -> podejscie -> kluczowy krok)
- `typowe_bledy` — realistyczne, z poprawna kara punktowa
- Brak literowek w kluczowych terminach (nazwy algorytmow, funkcji, protokolow)

## Krok 4: Wstaw do pliku JSON

Dodaj wygenerowane cwiczenia do tablicy `cwiczenia` w docelowym pliku.
Zaktualizuj `punkty_lacznie` = suma wszystkich punktow.
Jezeli uzyles nowych tagow — dodaj je do `tagi_globalne` i do `tagi_rejestr.json`.

## Krok 5: Walidacja (OBOWIAZKOWA)

Po wstawieniu uruchom OBA narzedzia:

```bash
# 1. Walidacja schema JSON (struktura, tagi, punkty):
python3 analiza/cwiczenia/validate_json.py --file NN_nazwa

# 2. Weryfikacja merytoryczna (kompilacja C++, SQL, sanity checks):
python3 analiza/cwiczenia/verify/verify_all.py --file NN_nazwa --verbose
```

### Interpretacja wynikow

- **validate_json.py 0 ERRORS** = struktura OK
- **verify_all.py PASS** = kod/SQL daje poprawny wynik
- **verify_all.py MANUAL_REVIEW** = TEORIA/ARKUSZ — juz zweryfikowane w Kroku 3.5
- **verify_all.py FAIL** = blad merytoryczny — NAPRAW i uruchom ponownie

Jezeli sa bledy — napraw je i uruchom walidacje ponownie. Powtarzaj az:
- validate_json.py: 0 ERRORS
- verify_all.py: 0 FAIL, 0 ERROR

## Krok 6: Podsumowanie

Wyswietl uzytkownikowi:
- Ile cwiczen dodano
- Do jakiego pliku
- Wynik walidacji
- Aktualna liczba cwiczen w pliku (bylo X, jest Y)
