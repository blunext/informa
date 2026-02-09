# Podsumowanie Analizy Matur z Informatyki Rozszerzonej 2014-2025

## Status Wykonania

### Ukonczone Analizy (11/11 lat):
- 2014 - Pelna analiza (analiza_2014.json)
- 2015 - Pelna analiza (analiza_2015.json)
- 2016 - Pelna analiza (analiza_2016.json)
- 2017 - Pelna analiza (analiza_2017.json)
- 2018 - Czesciowa analiza: tylko Czesc I (analiza_2018.json)
- 2019 - Pelna analiza (analiza_2019.json)
- 2021 - Pelna analiza (analiza_2021.json)
- 2022 - Pelna analiza (analiza_2022.json)
- 2023 - Pelna analiza (analiza_2023.json)
- 2024 - Pelna analiza (analiza_2024.json)
- 2025 - Pelna analiza (analiza_2025.json)

Brak: 2020 (egzamin odwolany - COVID-19)

---

## Formuly Egzaminacyjne

### Formula 2015 (stara) - lata 2014-2022
**2014**: Czesc I: 90 min / 20 pkt, Czesc II: 120 min / 30 pkt = 210 min / 50 pkt
**2015-2022**: Czesc I: 60 min / 15 pkt, Czesc II: 150 min / 35 pkt = 210 min / 50 pkt
- 6 zadan (3 teoria + 3 praktyka)
- Czesc I: algorytmy, analiza, test P/F
- Czesc II: programowanie + arkusz kalkulacyjny + bazy danych SQL

### Formula 2023 (nowa) - lata 2023-2025
**2023-2025**: Jeden arkusz, 210 min, 50 pkt
- 7-8 zadan roznego typu w jednym arkuszu
- Krotkie pytania teoretyczne za 1-2 pkt (quick wins)
- Zadania programistyczne, arkuszowe i SQL polaczone w jednym bloku
- Wieksza roznorodnosc typow zadan

### Kluczowe roznice:
| Cecha | Formula 2015 | Formula 2023 |
|-------|-------------|-------------|
| Czas | 210 min (60+150) | 210 min (jeden blok) |
| Punkty | 50 (15+35) | 50 |
| Czesci | 2 oddzielne | 1 arkusz |

| Liczba zadan | 6 | 7-8 |
| Quick wins | 3-5 pkt (test P/F) | 2-3 pkt (krotkie pytania) |
| Teoria | Oddzielna czesc | Wpleciona w zadania |

---

## Ranking Tematow (Na podstawie 11 lat)

### TIER 1 - Pewne (100%):
1. **SQL / Bazy danych** - 11/11 = 100% - ZAWSZE na egzaminie
2. **Operacje na liczbach/cyfrach** - 10/10 = 100% - mod/div, cyfry, podzielnosc
3. **Przetwarzanie plikow** - 10/10 = 100% - ifstream, parsowanie danych
4. **Arkusz kalkulacyjny** - 10/10 = 100% - SUM, IF, COUNTIF, wykresy

### TIER 2 - Bardzo czeste (70-85%):
5. **Systemy liczbowe** - 9/11 = 82% - bin/oct/hex, konwersje
6. **Rekurencja** - 8/11 = 73% - sledzenie, konwersja na iteracje
7. **Sortowanie** - 8/11 = 73% - sort, partition, klucze
8. **Teoria liczb / NWD** - 8/11 = 73% - Euklides, dzielniki, l. pierwsze

### TIER 3 - Czeste (55-65%):
9. **Zlozonosc algorytmow** - 7/11 = 64% - O(n), O(n^2), O(log n)
10. **Operacje na stringach** - 6/10 = 60% - ASCII, manipulacja tekstem

### TIER 4 - Sporadyczne (25-40%):
11. **Kryptografia / bezpieczenstwo** - 4/11 = 36%
12. **Sieci komputerowe** - 3/11 = 27%
13. **Przeszukiwanie binarne** - 3/11 = 27%
14. **Geometria / matematyka** - 3/11 = 27%

### TIER 5 - Rzadkie (<20%):
15. Struktury danych (kopiec, BST) - 2/11 = 18%
16. Programowanie dynamiczne - 1/11 = 9%
17. Algorytmy zachlanne - 1/11 = 9%
18. DFS/BFS - 1/11 = 9%
19. Operacje bitowe (XOR) - 1/11 = 9%
20. Modele barw - 1/11 = 9%
21. Technologie web - 1/11 = 9%

---

## Rozklad Punktow wg Typu Zadania (24 typy w 4 kategoriach)

Pelna macierz: `ranking_typow_zadan.csv`

### TEORIA I ANALIZA (6 typow, lacznie 164 pkt):
| Typ zadania | Lat | Laczne pkt | Opis |
|-------------|-----|-----------|------|
| sledzenie_algorytmu | 11/11 | 45 | Przesledzic algorytm krok po kroku |
| projektowanie_algorytmu | 11/11 | 43 | Napisac algorytm/pseudokod |
| analiza_algorytmu | 10/11 | 37 | Zlozonosc, wlasciwosci, dowody |
| test_prawda_falsz | 10/11 | 25 | Ocenic prawdziwosc zdan (P/F) |
| konwersja_systemow_liczbowych | 9/11 | 12 | Konwersje miedzy bazami |
| teoria_bezpieczenstwa | 2/11 | 2 | Szyfrowanie, protokoly |

### IMPLEMENTACJA (8 typow, lacznie 147 pkt):
| Typ zadania | Lat | Laczne pkt | Opis |
|-------------|-----|-----------|------|
| przetwarzanie_cyfry_liczby | 6/11 | 36 | Cyfry, NWD, potegi, faktoryzacja |
| przetwarzanie_napisy | 4/11 | 25 | Palindromy, szyfry, ASCII |
| przetwarzanie_zlozone | 4/11 | 24 | Wieloetapowy algorytm na danych |
| przetwarzanie_zliczanie | 5/11 | 17 | Zliczanie/filtrowanie danych |
| przetwarzanie_minmax | 5/11 | 17 | Min/max, sortowanie, rozklad |
| przetwarzanie_sekwencje | 3/11 | 13 | Najdluzszy podciag, bloki |
| przetwarzanie_obrazy_2D | 2/11 | 11 | Piksele, siatki, DFS/BFS |
| obliczenia_geometryczne | 1/11 | 4 | Odleglosci, srodki, pola |

### ARKUSZ KALKULACYJNY (5 typow, lacznie 112 pkt):
| Typ zadania | Lat | Laczne pkt | Opis |
|-------------|-----|-----------|------|
| arkusz_agregacja_warunkowa | 9/11 | 38 | SUMIF, COUNTIF, AVERAGEIF |
| arkusz_symulacja | 9/11 | 37 | Symulacje krokowe, formuly dynamiczne |
| arkusz_wykres | 8/11 | 25 | Kolumnowy, kolowy, liniowy |
| arkusz_agregacja_podstawowa | 3/11 | 9 | SUM, COUNT, AVERAGE, MAX/MIN |
| arkusz_transformacja | 2/11 | 3 | Grupowanie, pivoty |

### SQL (4 typy, lacznie 92 pkt):
| Typ zadania | Lat | Laczne pkt | Opis |
|-------------|-----|-----------|------|
| sql_group_by | 8/11 | 36 | GROUP BY z COUNT/SUM/AVG |
| sql_podzapytania | 7/11 | 25 | Podzapytania, NOT IN, EXISTS |
| sql_join | 8/11 | 21 | Laczenie 2-3 tabel JOIN |
| sql_select_where | 4/11 | 10 | Prosty SELECT z WHERE |

### Podsumowanie kategorii:
| Kategoria | Typow | Laczne pkt | Srednia/rok |
|-----------|-------|-----------|-------------|
| TEORIA | 6 | 164 | ~15 |
| IMPLEMENTACJA | 8 | 147 | ~13 |
| ARKUSZ | 5 | 112 | ~10 |
| SQL | 4 | 92 | ~8 |
| **RAZEM** | **23** | **515** | **~47** |

---

## Analiza Rok po Roku

### 2014 (Formula 2015 - stara, 90+120 min)
- Zad 1: **Korale** (8 pkt) - rekurencja, konwersja na iteracje, rep. binarna
- Zad 2: **Bisekcja** (6 pkt) - algorytm numeryczny, przeszukiwanie binarne
- Zad 3: **Test mieszany** (6 pkt) - systemy liczbowe, NWD, kombinatoryka
- Zad 4: **Arkusz - przychody/koszty** (9 pkt) - SUM, IF, wykresy
- Zad 5: **Programowanie - napisy** (10 pkt) - pliki, stringi, ASCII, grupowanie
- Zad 6: **SQL - przedszkola** (11 pkt) - JOIN, GROUP BY, AVG, LIKE

### 2015 (Formula 2015, 60+120 min)
- Zad 1: **Problem telewidza** (5 pkt) - algorytm zachlanny (Activity Selection)
- Zad 2: **Test P/F** (5 pkt) - systemy liczbowe, kompresja, Excel
- Zad 3: **Rozszerzony Euklides** (5 pkt) - NWD, rekurencja
- Zad 4: **Programowanie - liczby** (12 pkt) - cyfry, podzielnosc, min/max
- Zad 5: **Arkusz - demografia** (13 pkt) - SUMIF, wykresy, prognozy
- Zad 6: **SQL - Formula 1** (10 pkt) - JOIN, LEFT JOIN, GROUP BY

### 2016 (Formula 2015, 60+120 min)
- Zad 1: **Liczby skojarzone** (5 pkt) - dzielniki, optymalizacja O(sqrt(n))
- Zad 2: **Przestawienia** (6 pkt) - partycjonowanie (quicksort partition)
- Zad 3: **Test P/F** (4 pkt) - DNS/SMTP, rekurencja, arytm. binarna, SO
- Zad 4: **Arkusz - Monte Carlo pi** (11 pkt) - symulacja, COUNTIF, wykresy
- Zad 5: **SQL - uniwersytet** (12 pkt) - JOIN, LEFT JOIN, HAVING, LENGTH
- Zad 6: **Programowanie - szyfr Cezara** (12 pkt) - kryptografia, pliki, ASCII

### 2017 (Formula 2015, 60+120 min)
- Zad 1: **Prostokat** (6 pkt) - optymalizacja O(n), podzielnosc
- Zad 2: **Rekurencja licz(x)** (6 pkt) - sledzenie, rep. binarna
- Zad 3: **Test P/F** (3 pkt) - SQL ORDER BY, GROUP BY/HAVING, kryptografia
- Zad 4: **Arkusz - cukier** (13 pkt) - rabaty progresywne, SUMIF, wykresy
- Zad 5: **SQL - pilka reczna** (11 pkt) - JOIN, LEFT JOIN, CASE WHEN
- Zad 6: **Programowanie - obraz rastrowy** (11 pkt) - piksele, DFS/BFS, connected components

### 2018 (Nowa Formula 2015, 60+150 min) - CZESCIOWA ANALIZA
- Zad 1: **Pierwiastek szescienny** (6 pkt) - binary search, zlozonosc
- Zad 2: **Krajobraz** (6 pkt) - geometria, sortowanie punktow
- Zad 3: **Test P/F** (3 pkt) - PHP/JS, RGB/CMYK, SQL
- Czesc II: NIE PRZEANALIZOWANA

### 2019 (Formula 2015, 60+150 min)
- Zad 1: **Ulubione liczby** (6 pkt) - wyszukiwanie binarne, zlozonosc
- Zad 2: **Analiza algorytmu** (6 pkt) - rekurencja, drzewo wywolan
- Zad 3: **Test P/F** (3 pkt) - SQL, systemy liczbowe, DNS
- Zad 4: **Programowanie - liczby** (12 pkt) - potegi 3, silnia cyfr, NWD
- Zad 5: **Arkusz - pogoda** (11 pkt) - filtrowanie, srednie, wykresy
- Zad 6: **SQL - perfumy** (12 pkt) - 3 tabele, JOIN, LIKE, podzapytania

### 2021 (Formula 2015, 60+150 min) - pierwszy po COVID
- Zad 1: **Cyfrowe dopelnienie** (6 pkt) - operacje mod/div na cyfrach
- Zad 2: **Kopiec binarny** (6 pkt) - struktura danych, sift-up
- Zad 3: **Test P/F** (3 pkt) - rekurencja, systemy liczbowe, SQL
- Zad 4: **Programowanie - napisy** (11 pkt) - DOPISZ/USUN/ZMIEN/PRZESUN
- Zad 5: **Arkusz - wodociagi** (12 pkt) - agregacja, prognoza, wykresy
- Zad 6: **SQL - gra strategiczna** (12 pkt) - COUNT DISTINCT, LIKE, JOIN

### 2022 (Formula 2015, 60+150 min) - ostatni rok starej formuly
- Zad 1: **n-permutacja** (6 pkt) - tablica zliczajaca, counting array
- Zad 2: **ab-slowo** (6 pkt) - prefix/suffix sum, analiza algorytmu
- Zad 3: **Test P/F** (3 pkt) - zlozonosc O(n^2), systemy liczbowe, SQL
- Zad 4: **Programowanie - liczby** (12 pkt) - cyfry, czynniki pierwsze, trojki
- Zad 5: **Arkusz - soki** (12 pkt) - zamowienia, wykresy kolowe, daty
- Zad 6: **SQL - ewidencja uczniow** (11 pkt) - JOIN, daty, czas pobytu

### 2023 (Formula 2023 - NOWA, 210 min)
- Zad 1: **Biblioteczka Adama** (7 pkt) - BST, rekurencja, preorder
- Zad 2: **Liczby binarne** (11 pkt) - bloki, XOR, konwersje, kod Graya
- Zad 3: **Liczba Pi** (10 pkt) - przetwarzanie ciagow, zliczanie, rosnaco-malejace
- Zad 4: **Szyfrowanie asymetryczne** (1 pkt) - quick win, P/F
- Zad 5: **Systemy pozycyjne** (1 pkt) - quick win, porownywanie
- Zad 6: **Konfitury owocowe** (10 pkt) - arkusz, symulacja produkcji
- Zad 7: **Gry planszowe** (10 pkt) - baza danych, SQL, AVG, CASE WHEN

### 2024 (Formula 2023, 210 min)
- Zad 1: **Plansza** (5 pkt) - programowanie dynamiczne, tablica 2D
- Zad 2: **Cyfry** (3 pkt) - sledzenie algorytmu, mod/div
- Zad 3: **Nieparzysty skrot** (10 pkt) - pseudokod, pliki, NWD
- Zad 4: **Liczby** (10 pkt) - l. pierwsze, faktoryzacja, sliding window
- Zad 5: **Protokoly** (1 pkt) - quick win, HTTP/FTP/DHCP
- Zad 6: **Systemy liczbowe** (2 pkt) - dodawanie/odejmowanie w syst. 3 i 9
- Zad 7: **Hurtownia** (10 pkt) - arkusz, sprzedaz, rabaty kumulacyjne
- Zad 8: **Rejestr wykroczen** (9 pkt) - SQL, JOIN, LEFT JOIN, NOT IN

### 2025 (Formula 2023, 210 min)
- Zad 1: **Funkcja rekurencyjna** (9 pkt) - rekurencja, zamiana par cyfr
- Zad 2: **Zapis symboliczny** (11 pkt) - palindromy, syst. trojkowy, wzorce 2D
- Zad 3: **Dron** (6 pkt) - NWD, geometria, srodek odcinka
- Zad 4: **Keylogger** (1 pkt) - quick win, bezpieczenstwo
- Zad 5: **Dodawanie binarne** (2 pkt) - arytmetyka binarna
- Zad 6: **Martianeum** (11 pkt) - arkusz, symulacja, wykresy skumulowane
- Zad 7: **Woda na Marsie** (10 pkt) - baza danych, 4 tabele, SQL

---

## Kluczowe Wzorce Kodu

### 1. Operacje na cyfrach liczby (100%):
```cpp
while (n > 0) {
    int cyfra = n % 10;
    // przetworz cyfre
    n = n / 10;
}
```

### 2. NWD - algorytm Euklidesa (73%):
```cpp
int nwd(int a, int b) {
    while (b != 0) {
        int temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}
```

### 3. Przeszukiwanie binarne (27%):
```cpp
int left = 0, right = n - 1;
while (left < right) {
    int mid = (left + right) / 2;
    if (condition(mid)) right = mid;
    else left = mid + 1;
}
```

### 4. SQL JOIN + agregacja (100%):
```sql
SELECT t1.col, COUNT(*), SUM(t2.value)
FROM table1 t1
INNER JOIN table2 t2 ON t1.id = t2.fk
WHERE condition
GROUP BY t1.col
HAVING COUNT(*) > threshold
ORDER BY SUM(t2.value) DESC;
```

### 5. SQL LEFT JOIN + IS NULL (czeste):
```sql
SELECT t1.name FROM table1 t1
LEFT JOIN table2 t2 ON t1.id = t2.fk
WHERE t2.id IS NULL;
```

### 6. Sortowanie z kluczem:
```cpp
sort(arr, arr+n, [](auto& a, auto& b) {
    return key(a) < key(b);
});
```

---

## Strategia Punktowa

### Quick wins (3-5 pkt, 100% accuracy):
- Pytania P/F (test)
- Krotkie pytania za 1-2 pkt (nowa formula)
- Proste sledzenie algorytmu

### Zadania standardowe (25-30 pkt, 90% accuracy):
- SQL (JOIN, GROUP BY, agregacje)
- Arkusz kalkulacyjny (formuly, wykresy)
- Proste programowanie (pliki, cyfry, zliczanie)

### Zadania trudne (15-20 pkt, 70% accuracy):
- Optymalizacja algorytmow
- Zlozone programowanie (DP, symulacje)
- Trudne SQL (podzapytania, LEFT JOIN + NULL)

### Kolejnosc rozwiazywania:
1. Quick wins (5-10 min)
2. SQL - jezeli znasz, szybkie punkty (30-40 min)
3. Arkusz kalkulacyjny (40-50 min)
4. Programowanie - proste podzadania (30-40 min)
5. Teoria/algorytmy - jezeli zostal czas (20-30 min)
6. Trudne podzadania - na koniec (reszta czasu)

---

## Pliki w Projekcie

### Analizy szczegolow (JSON):
- `analiza_2014.json` - `analiza_2025.json` (11 plikow)

### Dokumenty strategiczne:
- `strategia_egzaminacyjna.md` - GLOWNY DOKUMENT (TOP 14 algorytmow + implementacje + SQL + przewodnik typow zadan)
- `ranking_tematow.csv` - Macierz temat x rok (21 tematow x 11 lat)
- `ranking_typow_zadan.csv` - Macierz typ zadania x rok (23 typy x 11 lat, z punktami)
- `PODSUMOWANIE_FINAL.md` - Ten dokument
- `podsumowanie_szybkie_wszystkie_lata.md` - Przeglad chronologiczny

### Wzorce kodu:
- `cpp_szablony.md` - **Samowystarczalna sciagawka C++** (wczytywanie plikow CKE, algorytmy, pulapki)
- `pseudokod_wzorce.md` - **Sciagawka pseudokodu CKE** (slowniczek, 7 archetypow, tabelki sledzenia, pulapki)
- `wzorce_2014.md` - Wzorce i pulapki z 2014
- `wzorce_2015.md` - Wzorce i pulapki z 2015

---

## TODO: Materialy do Stworzenia

### Priorytet WYSOKI:
- [x] C++ Templates — `cpp_szablony.md` (samowystarczalna sciagawka: wczytywanie plikow CKE, cyfry/liczby, napisy, zliczanie, sortowanie, DP, BFS/DFS, systemy liczbowe, pulapki)
- [x] SQL Templates — `sql_szablony.md` (samowystarczalna sciagawka: kolejnosc klauzul, SELECT/WHERE, JOIN 2-4 tabel, LEFT JOIN, GROUP BY + agregacje, podzapytania, CASE WHEN, daty, schematy baz CKE, 10 pulapek)
- [ ] Checklisty (przed_egzaminem.md, podczas_egzaminu.md, debug_checklist.md)
- [x] Arkusz kalkulacyjny templates (`arkusz_formuly.md` — SUMIF/SUMIFS, COUNTIF, odniesienia $, wykresy, symulacje, VLOOKUP, pulapki)
- [x] Wzorce pseudokodu CKE (`pseudokod_wzorce.md` — slowniczek CKE, 7 archetypow z tabelkami sledzenia, struktury danych, ograniczenia, pulapki)
- [ ] Cwiczenia ze sledzenia algorytmow (`cwiczenia_sledzenie.md` — tabelki krok-po-kroku)
- [x] Schemat decyzyjny / drzewko rozwiazywania (`drzewo_decyzyjne.md` — "widzisz X w zadaniu → uzyj algorytmu Y", szybkie rozpoznawanie typu)
- [x] Wzorce wczytywania plikow — wlaczone w `cpp_szablony.md` sekcja 1 (8 szablonow: spacje, linia-po-linii, CSV srednik, TSV tab, pary/trojki, wiele plikow, przecinek europejski, siatka 2D)

### Priorytet SREDNI:
- [ ] Wzorce dla lat 2018-2025
- [ ] Zestaw cwiczen dla TOP 10 algorytmow
- [ ] Probny egzamin w nowej formule 2023+ (`probny_egzamin.md` — 210 min / 50 pkt / 7 zadan)
- [ ] Karta bledow z odpowiedzi CKE (`typowe_bledy_cke.md` — za co traca punkty)
- [ ] Zadania treningowe wg 23 typow zadan (`cwiczenia_wg_typu/` — po 5 zadan na typ, 115 lacznie)
  - [x] Etap 1: TEORIA (6 typow, 30 cwiczen) — pliki 01-06
  - [x] Etap 2: IMPLEMENTACJA (8 typow, 40 cwiczen) — pliki 07-14
  - [x] Etap 3: ARKUSZ (5 typow, 25 cwiczen) — pliki 15-19
  - [x] Etap 4: SQL (4 typy, 20 cwiczen) — pliki 20-23
- [ ] Rozwiazania wzorcowe z komentarzami (`rozwiazania_wzorcowe/` — 2-3 rozwiazane przyklady na najczestszy typ zadania)
- [ ] Fiszki / active recall (`fiszki/` — pytanie-odpowiedz: zlozonosci, wzorce, SQL, formuly arkuszowe)

### Priorytet NISKI:
- [ ] Karta szybkiego dostepu / cheat sheet (`cheat_sheet.md` — 1-2 strony do wydruku)
- [ ] Plan nauki z harmonogramem (`plan_nauki.md` — rozklad tygodniowy wg priorytetow TIER 1-5)

---

*Ostatnia aktualizacja: 2026-02-09*
*Przeanalizowane lata: 11 (2014-2019, 2021-2025)*
*Zidentyfikowane tematy: 21, Zidentyfikowane typy zadan: 23*
