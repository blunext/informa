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
- **verify_all.py MANUAL_REVIEW** = OK dla TEORIA/ARKUSZ (brak auto-weryfikacji)
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
