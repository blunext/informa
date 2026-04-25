# Ranking algorytmow w zadaniach maturalnych CKE 2014-2025

_Wygenerowano automatycznie z `analiza/scripts/generate_ranking.py`_
_Zrodla: 30 plikow `matura_*.json` (641 podzadan) + `algorytmy_rejestr.json` (65 algorytmow)_

## Streszczenie

- **641 podzadan** sklasyfikowanych w 30 sesjach CKE (2014-2025), 1500 punktow lacznie.
- **604/641 (94.2%)** podzadan ma przynajmniej 1 tag algorytmu.
- **1456** lacznie wystapien tagow (srednio 2.41 na podzadanie z tagami).
- **61/65** algorytmow z rejestru rzeczywiscie pojawia sie w zadaniach.

**TOP 5 algorytmow** (wg liczby wystapien):

| # | Algorytm | Kategoria | Wystapien | Punkty | Lat |
|---|---|---|---:|---:|---:|
| 1 | `iteracja-po-pliku` | wzorce | 150 | 435 | 12/12 |
| 2 | `SQL-JOIN` | wzorce | 85 | 191 | 12/12 |
| 3 | `SQL-aggregacja` | wzorce | 85 | 177 | 11/12 |
| 4 | `akumulator-licznik` | wzorce | 74 | 212 | 12/12 |
| 5 | `sledzenie-pseudokod` | wzorce | 74 | 136 | 12/12 |

**Glowne wnioski edukacyjne**:

1. **Praca z plikiem dominuje**: `iteracja-po-pliku` jest w prawie kazdej sesji CKE — to umiejetnosc niezbedna.
2. **SQL ma TIER 1** — `SQL-JOIN`, `SQL-aggregacja`, `SQL-GROUP-BY`, `SQL-WHERE` to filary zadania bazodanowego.
3. **Sledzenie pseudokodu** rownie czeste co programowanie — wymagana zarowno teoria, jak i praktyka.
4. **Konwersje systemow liczbowych** + **Horner** wystepuje regularnie — fundamenty teorii.
5. **Programowanie dynamiczne** pojawilo sie po raz pierwszy w 2024 — nowy obszar wymagajacy uwagi.

---

## Ranking glowny

Lista wszystkich algorytmow uzytych w zadaniach CKE 2014-2025, posortowana wg liczby wystapien.

| # | Algorytm | Kategoria | Wystapien | Punkty | Lat | Podstawa | Przyklady |
|---|---|---|---:|---:|---:|---|---|
| 1 | `iteracja-po-pliku` | wzorce | 150 | 435 | 12/12 | — | `2014M.5a`, `2014M.5b`, `2014M.5c` |
| 2 | `SQL-JOIN` | wzorce | 85 | 191 | 12/12 | — | `2014M.6c`, `2014M.6d`, `2014P.6.1` |
| 3 | `SQL-aggregacja` | wzorce | 85 | 177 | 11/12 | — | `2014M.6c`, `2014M.6d`, `2014P.6.2` |
| 4 | `akumulator-licznik` | wzorce | 74 | 212 | 12/12 | — | `2014M.5a`, `2014P.3.2`, `2014P.5.1` |
| 5 | `sledzenie-pseudokod` | wzorce | 74 | 136 | 12/12 | — | `2014M.1a`, `2014M.2a`, `2014M.3a` |
| 6 | `arkusz-agregacja-warunkowa` | wzorce | 72 | 192 | 12/12 | — | `2014M.4a`, `2015C.4.1`, `2015M.5.1` |
| 7 | `SQL-GROUP-BY` | wzorce | 72 | 156 | 11/12 | — | `2014M.6c`, `2014M.6d`, `2014P.6.2` |
| 8 | `current-max` | wzorce | 68 | 197 | 12/12 | — | `2014M.4a`, `2014M.4c`, `2014M.5c` |
| 9 | `SQL-WHERE` | wzorce | 64 | 132 | 12/12 | — | `2014M.6a`, `2014M.6b`, `2014M.6c` |
| 10 | `iteracja-po-cyfrach` | wzorce | 59 | 171 | 11/12 | — | `2014M.1c`, `2014P.1.1`, `2014P.1.2` |
| 11 | `konwersja-systemow` | klasyczne | 56 | 108 | 11/12 | I.2.a, I+II.2.b | `2014M.1a`, `2014M.1c`, `2014M.3c` |
| 12 | `arkusz-symulacja-iteracyjna` | wzorce | 53 | 155 | 11/12 | — | `2014M.4c`, `2014P.4.1`, `2014P.4.2` |
| 13 | `rekurencja` | techniki | 48 | 102 | 12/12 | I.zr.1, I+II.3.b | `2014M.1a`, `2014M.1b`, `2015M.3.1` |
| 14 | `SQL-ORDER-BY` | wzorce | 43 | 94 | 11/12 | — | `2014M.6a`, `2014M.6c`, `2014M.6d` |
| 15 | `analiza-zlozonosci` | techniki | 41 | 82 | 12/12 | I.zr.5 | `2014M.1b`, `2014M.2b`, `2015C.2.2` |
| 16 | `znajdz-i-policz` | wzorce | 32 | 93 | 11/12 | — | `2014P.4.1`, `2015C.5.1`, `2015X.4.2` |
| 17 | `porownywanie-tekstow` | klasyczne | 29 | 84 | 11/12 | I.2.b | `2014M.5b`, `2015C.6.3`, `2015M.4.3` |
| 18 | `akumulator-suma` | wzorce | 24 | 67 | 9/12 | — | `2014M.5b`, `2014P.1.3`, `2014P.4.3` |
| 19 | `wyszukiwanie-liniowe` | techniki | 20 | 60 | 10/12 | I+II.3.a | `2014P.3.2`, `2014P.3.3`, `2015C.2.1` |
| 20 | `current-min` | wzorce | 19 | 63 | 10/12 | — | `2015C.5.4`, `2015M.4.3`, `2016C.6.5` |
| 21 | `SQL-podzapytanie-niezalezne` | wzorce | 19 | 45 | 9/12 | — | `2015M.6.1`, `2015M.6.2`, `2015X.6.4` |
| 22 | `SQL-LIKE` | wzorce | 17 | 37 | 9/12 | — | `2014M.6b`, `2014P.6.1`, `2014P.6.5` |
| 23 | `tablica-2D` | struktury | 16 | 45 | 6/12 | — | `2017M.6.1`, `2017M.6.2`, `2017M.6.3` |
| 24 | `Horner` | klasyczne | 16 | 41 | 9/12 | I+II.1.h | `2014M.3a`, `2014M.3b`, `2016C.6.4` |
| 25 | `SQL-HAVING` | wzorce | 16 | 36 | 8/12 | — | `2015M.6.3`, `2015X.6.4`, `2016M.5.5` |
| 26 | `SQL-LEFT-JOIN-NULL` | wzorce | 15 | 36 | 10/12 | — | `2014P.6.4`, `2015M.6.2`, `2016M.5.4` |
| 27 | `najdluzszy-podciag-niemalejacy` | klasyczne | 12 | 35 | 8/12 | I+II.2.c | `2014P.5.1`, `2014P.5.3`, `2015X.3.1` |
| 28 | `prefix-sum` | wzorce | 12 | 31 | 6/12 | — | `2014M.4c`, `2018M.5.4`, `2020C.1.2` |
| 29 | `przeszukiwanie-binarne` | klasyczne | 12 | 29 | 5/12 | I+II.1.b, I+II.3.a | `2015C.2.2`, `2017C.1.3`, `2018C.1.1` |
| 30 | `drzewo` | struktury | 11 | 23 | 4/12 | I+II.3.h | `2019M.2.1`, `2021M.2.1`, `2021M.2.2` |
| 31 | `SQL-funkcje-tekstowe` | wzorce | 10 | 25 | 6/12 | — | `2014M.6b`, `2016M.5.1`, `2016M.5.3` |
| 32 | `bisekcja` | klasyczne | 10 | 21 | 4/12 | I+II.1.f, I+II.3.a | `2014M.2a`, `2014M.2b`, `2014M.2c` |
| 33 | `jednoczesne-min-max` | klasyczne | 9 | 25 | 6/12 | I+II.1.d, I+II.3.c | `2014M.5c`, `2014P.5.2`, `2017M.1.1` |
| 34 | `NWD-Euklidesa` | klasyczne | 9 | 23 | 6/12 | I.2.a, I+II.1.a | `2014M.3d`, `2015C.3.1`, `2015M.3.1` |
| 35 | `arkusz-agregacja-podstawowa` | wzorce | 9 | 21 | 6/12 | — | `2014M.4b`, `2015C.4.3`, `2017C.5.1` |
| 36 | `akumulator-warunkowy` | wzorce | 8 | 19 | 5/12 | — | `2015C.5.2`, `2020M.5.4`, `2020M.5.5` |
| 37 | `zachlanne` | techniki | 7 | 22 | 4/12 | I.zr.1, I+II.3.d | `2015M.1.1`, `2015M.1.2`, `2020C.2.1` |
| 38 | `test-pierwszosci` | klasyczne | 7 | 20 | 5/12 | I.2.a | `2017C.4.1`, `2019C.4.1`, `2019C.4.2` |
| 39 | `wyszukiwanie-wzorca-naiwne` | klasyczne | 6 | 18 | 4/12 | I.2.b | `2015C.5.3`, `2021M.4.4`, `2024C.3.1` |
| 40 | `SQL-aggregacja-warunkowa` | wzorce | 6 | 17 | 5/12 | — | `2014P.6.5`, `2016M.5.3`, `2017C.6.5` |
| 41 | `pierwiastek-kwadratowy` | klasyczne | 6 | 14 | 3/12 | I+II.1.g | `2015X.2.1`, `2015X.2.2`, `2015X.2.3` |
| 42 | `sito-Eratostenesa` | klasyczne | 6 | 14 | 3/12 | I+II.1.c | `2017C.1.1`, `2017C.1.2`, `2020C.1.1` |
| 43 | `SQL-podzapytanie-skorelowane` | wzorce | 5 | 15 | 4/12 | — | `2015M.6.3`, `2019M.6.2`, `2021M.6.5` |
| 44 | `sliding-window` | wzorce | 5 | 13 | 3/12 | — | `2023M.3.1`, `2023M.3.2`, `2023M.3.3` |
| 45 | `sortowanie-szybkie` | klasyczne | 5 | 12 | 2/12 | I+II.3.c | `2016M.2.1`, `2016M.2.2`, `2016M.2.3` |
| 46 | `szyfr-Cezara` | klasyczne | 4 | 15 | 2/12 | I.2.b | `2016M.6.1`, `2016M.6.2`, `2016M.6.3` |
| 47 | `faktoryzacja` | klasyczne | 4 | 11 | 3/12 | I+II.2.a | `2014M.3d`, `2022M.4.2`, `2024M.4.3` |
| 48 | `szybkie-potegowanie` | klasyczne | 4 | 10 | 1/12 | I+II.1.i | `2023C.1.1`, `2023C.1.3`, `2023X.3.2` |
| 49 | `szyfrowanie-klucz-publiczny` | klasyczne | 4 | 4 | 4/12 | I+II.3.f | `2017M.3.3`, `2022P.7.1`, `2023M.4.1` |
| 50 | `sortowanie-babelkowe` | klasyczne | 3 | 10 | 3/12 | I.2.c | `2018M.2.2`, `2023C.2.3`, `2024P.3.3` |
| 51 | `ciag-Fibonacciego` | klasyczne | 3 | 5 | 1/12 | I.2.d | `2018C.2.1`, `2018C.2.2`, `2018C.2.3` |
| 52 | `graf` | struktury | 2 | 7 | 2/12 | I+II.3.h | `2017M.6.4`, `2024C.4.3` |
| 53 | `programowanie-dynamiczne` | techniki | 2 | 4 | 1/12 | I+II.3.e | `2024M.1.1`, `2024M.1.2` |
| 54 | `kolejka` | struktury | 1 | 4 | 1/12 | I+II.3.g | `2020C.5.5` |
| 55 | `lista-dynamiczna` | struktury | 1 | 4 | 1/12 | I+II.3.g | `2021M.4.4` |
| 56 | `sortowanie-przez-scalanie` | klasyczne | 1 | 4 | 1/12 | I+II.1.e, I+II.3.c | `2018C.4.4` |
| 57 | `sortowanie-przez-wstawianie` | klasyczne | 1 | 4 | 1/12 | I.2.c | `2018M.2.2` |
| 58 | `dziel-i-zwyciezaj` | techniki | 1 | 2 | 1/12 | I+II.3.c | `2018C.2.3` |
| 59 | `NWW` | klasyczne | 1 | 1 | 1/12 | I.2.a | `2014M.3d` |
| 60 | `ONP` | klasyczne | 1 | 1 | 1/12 | I+II.2.d, I+II.3.g | `2018C.3.2` |
| 61 | `stos` | struktury | 1 | 1 | 1/12 | I+II.3.g | `2018C.3.2` |

---

## Rozbicie per kategoria

### Klasyczne (23 algorytmow, 209 wystapien, 509 pkt)

| Algorytm | Wystapien | Punkty | Lat |
|---|---:|---:|---:|
| `konwersja-systemow` | 56 | 108 | 11/12 |
| `porownywanie-tekstow` | 29 | 84 | 11/12 |
| `Horner` | 16 | 41 | 9/12 |
| `najdluzszy-podciag-niemalejacy` | 12 | 35 | 8/12 |
| `przeszukiwanie-binarne` | 12 | 29 | 5/12 |
| `bisekcja` | 10 | 21 | 4/12 |
| `jednoczesne-min-max` | 9 | 25 | 6/12 |
| `NWD-Euklidesa` | 9 | 23 | 6/12 |
| `test-pierwszosci` | 7 | 20 | 5/12 |
| `wyszukiwanie-wzorca-naiwne` | 6 | 18 | 4/12 |
| `pierwiastek-kwadratowy` | 6 | 14 | 3/12 |
| `sito-Eratostenesa` | 6 | 14 | 3/12 |
| `sortowanie-szybkie` | 5 | 12 | 2/12 |
| `szyfr-Cezara` | 4 | 15 | 2/12 |
| `faktoryzacja` | 4 | 11 | 3/12 |
| `szybkie-potegowanie` | 4 | 10 | 1/12 |
| `szyfrowanie-klucz-publiczny` | 4 | 4 | 4/12 |
| `sortowanie-babelkowe` | 3 | 10 | 3/12 |
| `ciag-Fibonacciego` | 3 | 5 | 1/12 |
| `sortowanie-przez-scalanie` | 1 | 4 | 1/12 |
| `sortowanie-przez-wstawianie` | 1 | 4 | 1/12 |
| `NWW` | 1 | 1 | 1/12 |
| `ONP` | 1 | 1 | 1/12 |

### Techniki (6 algorytmow, 119 wystapien, 272 pkt)

| Algorytm | Wystapien | Punkty | Lat |
|---|---:|---:|---:|
| `rekurencja` | 48 | 102 | 12/12 |
| `analiza-zlozonosci` | 41 | 82 | 12/12 |
| `wyszukiwanie-liniowe` | 20 | 60 | 10/12 |
| `zachlanne` | 7 | 22 | 4/12 |
| `programowanie-dynamiczne` | 2 | 4 | 1/12 |
| `dziel-i-zwyciezaj` | 1 | 2 | 1/12 |

### Struktury (6 algorytmow, 32 wystapien, 84 pkt)

| Algorytm | Wystapien | Punkty | Lat |
|---|---:|---:|---:|
| `tablica-2D` | 16 | 45 | 6/12 |
| `drzewo` | 11 | 23 | 4/12 |
| `graf` | 2 | 7 | 2/12 |
| `kolejka` | 1 | 4 | 1/12 |
| `lista-dynamiczna` | 1 | 4 | 1/12 |
| `stos` | 1 | 1 | 1/12 |

### Wzorce (26 algorytmow, 1096 wystapien, 2766 pkt)

| Algorytm | Wystapien | Punkty | Lat |
|---|---:|---:|---:|
| `iteracja-po-pliku` | 150 | 435 | 12/12 |
| `SQL-JOIN` | 85 | 191 | 12/12 |
| `SQL-aggregacja` | 85 | 177 | 11/12 |
| `akumulator-licznik` | 74 | 212 | 12/12 |
| `sledzenie-pseudokod` | 74 | 136 | 12/12 |
| `arkusz-agregacja-warunkowa` | 72 | 192 | 12/12 |
| `SQL-GROUP-BY` | 72 | 156 | 11/12 |
| `current-max` | 68 | 197 | 12/12 |
| `SQL-WHERE` | 64 | 132 | 12/12 |
| `iteracja-po-cyfrach` | 59 | 171 | 11/12 |
| `arkusz-symulacja-iteracyjna` | 53 | 155 | 11/12 |
| `SQL-ORDER-BY` | 43 | 94 | 11/12 |
| `znajdz-i-policz` | 32 | 93 | 11/12 |
| `akumulator-suma` | 24 | 67 | 9/12 |
| `current-min` | 19 | 63 | 10/12 |
| `SQL-podzapytanie-niezalezne` | 19 | 45 | 9/12 |
| `SQL-LIKE` | 17 | 37 | 9/12 |
| `SQL-HAVING` | 16 | 36 | 8/12 |
| `SQL-LEFT-JOIN-NULL` | 15 | 36 | 10/12 |
| `prefix-sum` | 12 | 31 | 6/12 |
| `SQL-funkcje-tekstowe` | 10 | 25 | 6/12 |
| `arkusz-agregacja-podstawowa` | 9 | 21 | 6/12 |
| `akumulator-warunkowy` | 8 | 19 | 5/12 |
| `SQL-aggregacja-warunkowa` | 6 | 17 | 5/12 |
| `SQL-podzapytanie-skorelowane` | 5 | 15 | 4/12 |
| `sliding-window` | 5 | 13 | 3/12 |

---

## Heatmapa: rok x algorytm (TOP 20)

Liczba wystapien algorytmu w danym roku. `·` = brak.

| algorytm | 2014 | 2015 | 2016 | 2017 | 2018 | 2019 | 2020 | 2021 | 2022 | 2023 | 2024 | 2025 |
|---|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|:-:|
| iteracja-po-pliku | 6 | 12 | 14 | 6 | 7 | 6 | 14 | 11 | 17 | 22 | 16 | 19 |
| SQL-JOIN | 6 | 6 | 2 | 6 | 8 | 10 | 8 | 11 | 8 | 9 | 9 | 2 |
| SQL-aggregacja | 4 | 5 | 4 | 5 | 9 | 6 | 9 | 15 | 10 | 9 | 9 | · |
| sledzenie-pseudokod | 6 | 7 | 6 | 6 | 4 | 5 | 9 | 9 | 6 | 7 | 7 | 2 |
| akumulator-licznik | 3 | 10 | 6 | 5 | 5 | 2 | 4 | 5 | 11 | 11 | 5 | 7 |
| arkusz-agregacja-warunkowa | 1 | 8 | 7 | 5 | 7 | 5 | 8 | 8 | 4 | 7 | 7 | 5 |
| SQL-GROUP-BY | 5 | 5 | 2 | 7 | 6 | 6 | 4 | 13 | 9 | 8 | 7 | · |
| current-max | 4 | 6 | 5 | 2 | 3 | 2 | 7 | 4 | 10 | 12 | 5 | 8 |
| SQL-WHERE | 6 | 4 | 1 | 5 | 3 | 6 | 10 | 8 | 6 | 7 | 6 | 2 |
| iteracja-po-cyfrach | 4 | 7 | 4 | 1 | · | 3 | 3 | 4 | 9 | 8 | 9 | 7 |
| konwersja-systemow | 6 | 4 | 4 | 3 | · | 2 | 4 | 6 | 5 | 9 | 9 | 4 |
| arkusz-symulacja-iteracyjna | 4 | 5 | · | 2 | 3 | 5 | 7 | 6 | 8 | 5 | 5 | 3 |
| rekurencja | 2 | 2 | 4 | 6 | 3 | 5 | 6 | 3 | 5 | 4 | 4 | 4 |
| SQL-ORDER-BY | 5 | 4 | 2 | 4 | 6 | 5 | 4 | 4 | 2 | 3 | 4 | · |
| analiza-zlozonosci | 2 | 4 | 3 | 3 | 4 | 4 | 4 | 3 | 5 | 3 | 3 | 3 |
| znajdz-i-policz | 1 | 2 | 2 | · | 1 | 2 | 3 | 1 | 3 | 6 | 5 | 6 |
| porownywanie-tekstow | 1 | 3 | 2 | 3 | 3 | · | 4 | 2 | 2 | 5 | 1 | 3 |
| akumulator-suma | 3 | 4 | 3 | 3 | · | · | 2 | 1 | 3 | 3 | · | 2 |
| wyszukiwanie-liniowe | 2 | 3 | 4 | 2 | 1 | 1 | 3 | · | 2 | 1 | 1 | · |
| current-min | · | 2 | 1 | 1 | 1 | · | 3 | 2 | 1 | 3 | 2 | 3 |

---

## Rekomendacje kolejnosci nauki (TIER 1/2/3)

### TIER 1 — Must Have (16 algorytmow, ≥30 wystapien)

**Znajomosc TIER 1 pozwala podejsc do 575/641 podzadan (89.7%) i 1398/1500 pkt (93.2%).**
Te algorytmy musisz znac na 100%.

- `iteracja-po-pliku` (wzorce) — 150x, 12/12 lat
- `SQL-JOIN` (wzorce) — 85x, 12/12 lat
- `SQL-aggregacja` (wzorce) — 85x, 11/12 lat
- `akumulator-licznik` (wzorce) — 74x, 12/12 lat
- `sledzenie-pseudokod` (wzorce) — 74x, 12/12 lat
- `arkusz-agregacja-warunkowa` (wzorce) — 72x, 12/12 lat
- `SQL-GROUP-BY` (wzorce) — 72x, 11/12 lat
- `current-max` (wzorce) — 68x, 12/12 lat
- `SQL-WHERE` (wzorce) — 64x, 12/12 lat
- `iteracja-po-cyfrach` (wzorce) — 59x, 11/12 lat
- `konwersja-systemow` (klasyczne) — 56x, 11/12 lat
- `arkusz-symulacja-iteracyjna` (wzorce) — 53x, 11/12 lat
- `rekurencja` (techniki) — 48x, 12/12 lat
- `SQL-ORDER-BY` (wzorce) — 43x, 11/12 lat
- `analiza-zlozonosci` (techniki) — 41x, 12/12 lat
- `znajdz-i-policz` (wzorce) — 32x, 11/12 lat

### TIER 2 — Powinno sie znac (16 algorytmow, 10-29 wystapien)

Razem TIER 1+2 pozwala podejsc do 590/641 podzadan (92.0%) i 1435/1500 pkt (95.7%).

- `porownywanie-tekstow` (klasyczne) — 29x, 11/12 lat
- `akumulator-suma` (wzorce) — 24x, 9/12 lat
- `wyszukiwanie-liniowe` (techniki) — 20x, 10/12 lat
- `current-min` (wzorce) — 19x, 10/12 lat
- `SQL-podzapytanie-niezalezne` (wzorce) — 19x, 9/12 lat
- `SQL-LIKE` (wzorce) — 17x, 9/12 lat
- `tablica-2D` (struktury) — 16x, 6/12 lat
- `Horner` (klasyczne) — 16x, 9/12 lat
- `SQL-HAVING` (wzorce) — 16x, 8/12 lat
- `SQL-LEFT-JOIN-NULL` (wzorce) — 15x, 10/12 lat
- `najdluzszy-podciag-niemalejacy` (klasyczne) — 12x, 8/12 lat
- `prefix-sum` (wzorce) — 12x, 6/12 lat
- `przeszukiwanie-binarne` (klasyczne) — 12x, 5/12 lat
- `drzewo` (struktury) — 11x, 4/12 lat
- `SQL-funkcje-tekstowe` (wzorce) — 10x, 6/12 lat
- `bisekcja` (klasyczne) — 10x, 4/12 lat

### TIER 3 — Nice to have (29 algorytmow, 1-9 wystapien)

Rzadziej spotykane, ale mozna na nie trafic.

- `jednoczesne-min-max` (klasyczne) — 9x, 25 pkt
- `NWD-Euklidesa` (klasyczne) — 9x, 23 pkt
- `arkusz-agregacja-podstawowa` (wzorce) — 9x, 21 pkt
- `akumulator-warunkowy` (wzorce) — 8x, 19 pkt
- `zachlanne` (techniki) — 7x, 22 pkt
- `test-pierwszosci` (klasyczne) — 7x, 20 pkt
- `wyszukiwanie-wzorca-naiwne` (klasyczne) — 6x, 18 pkt
- `SQL-aggregacja-warunkowa` (wzorce) — 6x, 17 pkt
- `pierwiastek-kwadratowy` (klasyczne) — 6x, 14 pkt
- `sito-Eratostenesa` (klasyczne) — 6x, 14 pkt
- `SQL-podzapytanie-skorelowane` (wzorce) — 5x, 15 pkt
- `sliding-window` (wzorce) — 5x, 13 pkt
- `sortowanie-szybkie` (klasyczne) — 5x, 12 pkt
- `szyfr-Cezara` (klasyczne) — 4x, 15 pkt
- `faktoryzacja` (klasyczne) — 4x, 11 pkt
- `szybkie-potegowanie` (klasyczne) — 4x, 10 pkt
- `szyfrowanie-klucz-publiczny` (klasyczne) — 4x, 4 pkt
- `sortowanie-babelkowe` (klasyczne) — 3x, 10 pkt
- `ciag-Fibonacciego` (klasyczne) — 3x, 5 pkt
- `graf` (struktury) — 2x, 7 pkt
- `programowanie-dynamiczne` (techniki) — 2x, 4 pkt
- `kolejka` (struktury) — 1x, 4 pkt
- `lista-dynamiczna` (struktury) — 1x, 4 pkt
- `sortowanie-przez-scalanie` (klasyczne) — 1x, 4 pkt
- `sortowanie-przez-wstawianie` (klasyczne) — 1x, 4 pkt
- `dziel-i-zwyciezaj` (techniki) — 1x, 2 pkt
- `NWW` (klasyczne) — 1x, 1 pkt
- `ONP` (klasyczne) — 1x, 1 pkt
- `stos` (struktury) — 1x, 1 pkt

### Algorytmy z rejestru NIE testowane przez CKE (2014-2025)

Algorytmy z podstawy programowej ktorych CKE nigdy nie pyta — niski priorytet nauki.

- `fraktale-rekurencyjne` (podstawa: I+II.1.j) — Rekurencyjne tworzenie fraktali (drzewo binarne, plotek Kocha, dywan Sierpinskiego itp.).
- `metoda-wstepujaca-zstepujaca` (podstawa: I.zr.3) — Top-down (dekompozycja od ogolu) vs bottom-up (od szczegolow do calosci).
- `najdluzszy-wspolny-podciag` (podstawa: I+II.3.e) — Longest Common Subsequence (LCS) za pomoca programowania dynamicznego, O(n*m).
- `podciag-najwieksza-suma` (podstawa: I+II.2.c) — Najwiekszy podciag spojny o najwiekszej sumie (algorytm Kadane'a, O(n)).

---

## TOP 10 kombinacji 2-algorytmowych

Pary tagow ktore najczesciej wystepuja razem w jednym podzadaniu.

| # | Algorytm A | Algorytm B | Wspolnie |
|---|---|---|---:|
| 1 | `SQL-GROUP-BY` | `SQL-aggregacja` | 64 |
| 2 | `SQL-JOIN` | `SQL-aggregacja` | 61 |
| 3 | `akumulator-licznik` | `iteracja-po-pliku` | 56 |
| 4 | `SQL-GROUP-BY` | `SQL-JOIN` | 55 |
| 5 | `current-max` | `iteracja-po-pliku` | 43 |
| 6 | `SQL-JOIN` | `SQL-WHERE` | 40 |
| 7 | `SQL-WHERE` | `SQL-aggregacja` | 36 |
| 8 | `SQL-JOIN` | `SQL-ORDER-BY` | 31 |
| 9 | `iteracja-po-cyfrach` | `iteracja-po-pliku` | 30 |
| 10 | `iteracja-po-pliku` | `znajdz-i-policz` | 27 |

---

## Statystyki kategorii

Liczba wystapien tagu = ile razy algorytm pojawia sie w klasyfikacji (1 podzadanie moze miec wiele tagow).

| Kategoria | Algorytmow w rejestrze | Uzywanych | Wystapien | % wszystkich tagow |
|---|---:|---:|---:|---:|
| klasyczne | 26 | 23 | 209 | 14.4% |
| techniki | 7 | 6 | 119 | 8.2% |
| struktury | 6 | 6 | 32 | 2.2% |
| wzorce | 26 | 26 | 1096 | 75.3% |

