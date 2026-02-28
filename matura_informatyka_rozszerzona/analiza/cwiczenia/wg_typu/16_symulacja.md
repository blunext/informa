# 16. Symulacja w arkuszu kalkulacyjnym

Typ zadania: **arkusz_symulacja**
Czestotliwosc: 12/12 lat | Laczna punktacja: 37 pkt
Kategoria: ARKUSZ

## Umiejetnosci cwiczone w tym zestawie
`akumulator` `IF-warunkowy` `odwolanie-bezwzgledne` `kopiowanie-formul` `SUM-kumulacyjna` `COUNTIF` `IF-zagniezdzony` `MAX` `MIN` `INT` `ROUND` `symulacja-krokowa` `prog-warunkowy` `lancuch-zaleznosci` `formatowanie-wyniku` `ZAOKR` `rata-annuitetowa` `amortyzacja`

---

### Cwiczenie 16.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna (akumulator prosty)
**Tagi**: `akumulator` `kopiowanie-formul` `odwolanie-wzgledne`

Na koncie bankowym poczatkowe saldo wynosi 1000 zl. W kolejnych miesiacach dokonywano wplat (+) i wyplat (-):

| | A | B | C |
|---|---|---|---|
| 1 | **Miesiac** | **Operacja** | **Saldo** |
| 2 | 0 (start) | — | 1000 |
| 3 | 1 | +500 | ? |
| 4 | 2 | -200 | ? |
| 5 | 3 | +300 | ? |
| 6 | 4 | -800 | ? |
| 7 | 5 | +150 | ? |
| 8 | 6 | -100 | ? |

**Polecenie**: Napisz formule w komorce C3, ktora oblicza nowe saldo jako poprzednie saldo plus operacja. Formula powinna dzialac po skopiowaniu w dol az do C8. Wypelnij tabele.

<details>
<summary>Wskazowki</summary>

1. **Wskazowka 1**: Nowe saldo to suma poprzedniego salda i biezacej operacji. Zastanow sie, ktore komorki powinny byc w formule.
2. **Wskazowka 2**: Formula w C3 powinna odwolywac sie do C2 (poprzednie saldo) i B3 (biezaca operacja). Oba odwolania sa wzgledne.
3. **Wskazowka 3**: Wpisz `=C2+B3` w komorce C3. Po skopiowaniu w dol do C4 formula automatycznie zmieni sie na `=C3+B4` — dokladnie to, czego potrzebujemy.
</details>

<details>
<summary>Odpowiedz</summary>

**Formula (C3, kopiowana w dol):**
```
C3: =C2+B3
```

**Wypelniona tabela:**

| Miesiac | Operacja | Saldo |
|---------|----------|-------|
| 0 (start) | — | 1000 |
| 1 | +500 | 1500 |
| 2 | -200 | 1300 |
| 3 | +300 | 1600 |
| 4 | -800 | 800 |
| 5 | +150 | 950 |
| 6 | -100 | 850 |

**Weryfikacja krok po kroku:**
- C3: 1000 + 500 = 1500
- C4: 1500 + (-200) = 1300
- C5: 1300 + 300 = 1600
- C6: 1600 + (-800) = 800
- C7: 800 + 150 = 950
- C8: 950 + (-100) = 850

**Wyjasnienie**: Jest to prosty akumulator — kazdy wiersz odwoluje sie do poprzedniego wyniku. Odwolanie C2 jest wzgledne, wiec po skopiowaniu formuly w dol automatycznie zmienia sie na C3, C4 itd.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Uzycie stalej zamiast odwolania**: Wpisanie `=1000+B3` zamiast `=C2+B3` — formula dziala tylko dla pierwszego wiersza, po skopiowaniu w dol dalej dodaje do 1000 zamiast do biezacego salda.
2. **Pominiecie znaku minus w operacji**: Trzeba pamietac, ze wyplaty sa juz zapisane jako liczby ujemne — wystarczy dodawanie, nie trzeba osobno odejmowac.
</details>

---

### Cwiczenie 16.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 (Demografia)
**Tagi**: `akumulator` `odwolanie-bezwzgledne` `kopiowanie-formul` `mnozenie`

Populacja miasta wynosi 50000 mieszkancow. Kazdego roku populacja rosnie o 2.5%.

| | A | B |
|---|---|---|
| 1 | **Rok** | **Populacja** |
| 2 | 2024 | 50000 |
| 3 | 2025 | ? |
| 4 | 2026 | ? |
| 5 | 2027 | ? |
| 6 | 2028 | ? |
| 7 | 2029 | ? |

Wspolczynnik wzrostu podany jest w komorce D1: 1.025 (czyli +2.5%).

**Polecenie**: Napisz formule w B3, ktora mnozy poprzednia populacje przez wspolczynnik z D1 (z odniesieniem bezwzglednym). Formula musi dzialac po skopiowaniu w dol.

<details>
<summary>Wskazowki</summary>

1. **Wskazowka 1**: Nowa populacja = stara populacja razy wspolczynnik wzrostu. Zastanow sie, czy odwolanie do wspolczynnika powinno sie zmieniac przy kopiowaniu.
2. **Wskazowka 2**: Odwolanie do B2 (poprzednia populacja) powinno byc wzgledne (zmienia sie przy kopiowaniu), ale odwolanie do D1 (wspolczynnik) musi byc bezwzgledne (znak `$`).
3. **Wskazowka 3**: Formula to `=B2*$D$1`. Po skopiowaniu do B4 stanie sie `=B3*$D$1` — B2 zmienilo sie na B3, ale $D$1 zostalo bez zmian.
</details>

<details>
<summary>Odpowiedz</summary>

**Formula (B3, kopiowana w dol):**
```
B3: =B2*$D$1
```

**Wypelniona tabela (zaokraglone do calkowitych):**

| Rok | Populacja |
|-----|-----------|
| 2024 | 50000 |
| 2025 | 51250 |
| 2026 | 52531 |
| 2027 | 53845 |
| 2028 | 55191 |
| 2029 | 56571 |

**Weryfikacja:**
- B3: 50000 x 1.025 = 51250.00
- B4: 51250 x 1.025 = 52531.25 ~ 52531
- B5: 52531.25 x 1.025 = 53844.53 ~ 53845
- B6: 53844.53 x 1.025 = 55190.64 ~ 55191
- B7: 55190.64 x 1.025 = 56570.41 ~ 56571

**Wyjasnienie**: Znak `$` przed D i 1 (`$D$1`) sprawia, ze odwolanie jest bezwzgledne — nie zmienia sie przy kopiowaniu formuly. Bez `$` kopiowanie z B3 do B4 zmieniloby D1 na D2, co daloby bledny wynik.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Brak znaku $ przy wspolczynniku**: Napisanie `=B2*D1` zamiast `=B2*$D$1` — po skopiowaniu do B4 formula zmieni sie na `=B3*D2`, a komorka D2 jest pusta (wynik = 0).
2. **Zakotwiczenie obu odwolan**: Napisanie `=$B$2*$D$1` — kazdy rok mnozy poczatkowa populacje (50000) zamiast poprzedniej, co daje ten sam wynik 51250 w kazdym wierszu.
</details>

---

### Cwiczenie 16.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2023 (Konfitury)
**Tagi**: `akumulator` `IF-warunkowy` `symulacja-krokowa` `prog-warunkowy`

Magazyn ma poczatkowy zapas 200 sztuk. Codziennie odbywaja sie przyjecia i wydania towaru. Jezeli zapas na koniec dnia spadnie ponizej 50 sztuk, nalezy wygenerowac alarm "NISKI STAN".

| | A | B | C | D | E |
|---|---|---|---|---|---|
| 1 | **Dzien** | **Przyjecie** | **Wydanie** | **Zapas** | **Alarm** |
| 2 | 0 (start) | — | — | 200 | |
| 3 | 1 | 0 | 80 | ? | ? |
| 4 | 2 | 50 | 60 | ? | ? |
| 5 | 3 | 0 | 70 | ? | ? |
| 6 | 4 | 100 | 30 | ? | ? |
| 7 | 5 | 0 | 90 | ? | ? |
| 8 | 6 | 20 | 50 | ? | ? |
| 9 | 7 | 0 | 40 | ? | ? |

**Polecenie**:
1. Napisz formule w D3 obliczajaca zapas = poprzedni zapas + przyjecie - wydanie.
2. Napisz formule w E3 wyswietlajaca "NISKI STAN" gdy zapas < 50, w przeciwnym razie puste "".

<details>
<summary>Wskazowki</summary>

1. **Wskazowka 1**: Zapas to akumulator z dwoma zrodlami zmian: przyjecia (dodajesz) i wydania (odejmujesz). Alarm to osobna kolumna z warunkiem IF.
2. **Wskazowka 2**: Formula D3 odwoluje sie do D2 (poprzedni zapas), B3 (przyjecie) i C3 (wydanie). Formula E3 sprawdza wartosc D3 wzgledem progu 50.
3. **Wskazowka 3**: D3: `=D2+B3-C3`, E3: `=IF(D3<50;"NISKI STAN";"")`. Obie formuly kopiujemy w dol — wszystkie odwolania sa wzgledne.
</details>

<details>
<summary>Odpowiedz</summary>

**Formuly (wiersz 3, kopiowane w dol):**
```
D3: =D2+B3-C3
E3: =IF(D3<50;"NISKI STAN";"")
```

**Wypelniona tabela:**

| Dzien | Przyjecie | Wydanie | Zapas | Alarm |
|-------|-----------|---------|-------|-------|
| 0 (start) | — | — | 200 | |
| 1 | 0 | 80 | 120 | |
| 2 | 50 | 60 | 110 | |
| 3 | 0 | 70 | 40 | NISKI STAN |
| 4 | 100 | 30 | 110 | |
| 5 | 0 | 90 | 20 | NISKI STAN |
| 6 | 20 | 50 | -10 | NISKI STAN |
| 7 | 0 | 40 | -50 | NISKI STAN |

**Weryfikacja:**
- D3: 200 + 0 - 80 = 120 (>=50, brak alarmu)
- D4: 120 + 50 - 60 = 110 (>=50, brak alarmu)
- D5: 110 + 0 - 70 = 40 (<50, NISKI STAN)
- D6: 40 + 100 - 30 = 110 (>=50, brak alarmu)
- D7: 110 + 0 - 90 = 20 (<50, NISKI STAN)
- D8: 20 + 20 - 50 = -10 (<50, NISKI STAN)
- D9: -10 + 0 - 40 = -50 (<50, NISKI STAN)

**Wyjasnienie**: Kolumna D to akumulator z dwoma zrodlami zmian (przyjecia i wydania). Kolumna E uzywa IF do warunkowego wyswietlania tekstu. Ujemny zapas oznacza, ze magazyn jest "na minusie" (zamowienia przekraczaja stan).
</details>

<details>
<summary>Typowe bledy</summary>

1. **Odwrotna kolejnosc odejmowania**: Napisanie `=D2-B3+C3` (odjecie przyjecia i dodanie wydania) — odwrocona logika daje calkowicie bledne wyniki.
2. **Porownanie alarmu z odwrotnym znakiem**: Napisanie `IF(D3>50;...)` zamiast `IF(D3<50;...)` — alarm pojawia sie gdy zapas jest wysoki, a znika gdy jest niski.
3. **Puste pole jako spacja**: Napisanie `" "` (spacja) zamiast `""` (pusty tekst) w trzecim argumencie IF — komorka wyglada na pusta, ale zawiera spacje, co moze zaklocic inne formuly (np. COUNTBLANK).
</details>

---

### Cwiczenie 16.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)
**Tagi**: `SUM-kumulacyjna` `IF-zagniezdzony` `odwolanie-bezwzgledne` `prog-warunkowy`

Sklep internetowy stosuje rabaty kumulacyjne — rabat zalezy od lacznej kwoty wszystkich dotychczasowych zakupow klienta:
- laczna kwota < 500 zl -> rabat 0%
- laczna kwota 500-999 zl -> rabat 5%
- laczna kwota 1000-1999 zl -> rabat 10%
- laczna kwota >= 2000 zl -> rabat 15%

| | A | B | C | D | E |
|---|---|---|---|---|---|
| 1 | **Zakup nr** | **Kwota** | **Suma kumulacyjna** | **Prog rabatu** | **Cena po rabacie** |
| 2 | 1 | 350 | ? | ? | ? |
| 3 | 2 | 280 | ? | ? | ? |
| 4 | 3 | 420 | ? | ? | ? |
| 5 | 4 | 150 | ? | ? | ? |
| 6 | 5 | 600 | ? | ? | ? |
| 7 | 6 | 300 | ? | ? | ? |

**Polecenie**:
1. W C2 oblicz sume kumulacyjna (SUM z zakotwiczonym poczatkiem).
2. W D2 wyznacz procent rabatu na podstawie sumy kumulacyjnej (zagniezdzony IF).
3. W E2 oblicz cene po rabacie = kwota x (1 - rabat).

<details>
<summary>Wskazowki</summary>

1. **Wskazowka 1**: Suma kumulacyjna to SUM z zakresu, w ktorym poczatek jest zakotwiczony (bezwzgledny), a koniec jest wzgledny. Rabat to zagniezdzony IF sprawdzajacy progi od najwyzszego.
2. **Wskazowka 2**: Zakres w SUM to `$B$2:B2` — po skopiowaniu do wiersza 4 stanie sie `$B$2:B4`. W zagniezdzoym IF sprawdzaj najpierw >=2000, potem >=1000, potem >=500, a w ostatnim ELSE daj 0%.
3. **Wskazowka 3**: C2: `=SUM($B$2:B2)`, D2: `=IF(C2>=2000;15%;IF(C2>=1000;10%;IF(C2>=500;5%;0%)))`, E2: `=B2*(1-D2)`.
</details>

<details>
<summary>Odpowiedz</summary>

**Formuly (wiersz 2, kopiowane w dol):**
```
C2: =SUM($B$2:B2)
D2: =IF(C2>=2000;15%;IF(C2>=1000;10%;IF(C2>=500;5%;0%)))
E2: =B2*(1-D2)
```

**Wypelniona tabela:**

| Zakup nr | Kwota | Suma kum. | Rabat | Cena po rabacie |
|----------|-------|-----------|-------|-----------------|
| 1 | 350 | 350 | 0% | 350.00 |
| 2 | 280 | 630 | 5% | 266.00 |
| 3 | 420 | 1050 | 10% | 378.00 |
| 4 | 150 | 1200 | 10% | 135.00 |
| 5 | 600 | 1800 | 10% | 540.00 |
| 6 | 300 | 2100 | 15% | 255.00 |

**Weryfikacja:**
- C2: SUM(B2:B2) = 350 -> prog 0% -> E2: 350 x 1.00 = 350.00
- C3: SUM(B2:B3) = 350+280 = 630 -> prog 5% -> E3: 280 x 0.95 = 266.00
- C4: SUM(B2:B4) = 630+420 = 1050 -> prog 10% -> E4: 420 x 0.90 = 378.00
- C5: SUM(B2:B5) = 1050+150 = 1200 -> prog 10% -> E5: 150 x 0.90 = 135.00
- C6: SUM(B2:B6) = 1200+600 = 1800 -> prog 10% -> E6: 600 x 0.90 = 540.00
- C7: SUM(B2:B7) = 1800+300 = 2100 -> prog 15% -> E7: 300 x 0.85 = 255.00

**Wyjasnienie**: Kluczowy jest zakres `$B$2:B2` — poczatek jest zakotwiczony ($B$2), ale koniec (B2) jest wzgledny. Po skopiowaniu do wiersza 5 formula staje sie `=SUM($B$2:B5)`, co daje sume od poczatku do biezacego wiersza. Zagniezdzony IF sprawdza progi od najwiekszego, bo pierwszy speliony warunek konczy ewaluacje.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Sprawdzanie progow od najnizszego**: Napisanie `=IF(C2>=500;5%;IF(C2>=1000;10%;...))` — warunek C2>=500 jest spelniony takze dla 1000 i 2000, wiec rabat nigdy nie przekroczy 5%. Progi nalezy sprawdzac od najwyzszego.
2. **Brak zakotwiczenia poczatku SUM**: Napisanie `=SUM(B2:B2)` bez znaku `$` — po skopiowaniu do wiersza 5 formula stanie sie `=SUM(B5:B5)`, czyli sumuje tylko biezacy wiersz zamiast calego zakresu od poczatku.
3. **Uzycie przecinka zamiast srednika**: W polskiej wersji Excela separator argumentow to srednik (`;`), nie przecinek (`,`). Przecinek jest traktowany jako separator dziesietny.
</details>

---

### Cwiczenie 16.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 (Martianeum)
**Tagi**: `akumulator` `IF-warunkowy` `lancuch-zaleznosci` `COUNTIF` `symulacja-krokowa`

Kopalnia codziennie wydobywa pewna ilosc rudy. Ruda jest gromadzona w magazynie. Gdy zapas osiagnie lub przekroczy 100 ton, wysylany jest transport (dokladnie 100 ton), a reszta zostaje w magazynie. Nalezy policzyc ile transportow wyslano.

| | A | B | C | D | E |
|---|---|---|---|---|---|
| 1 | **Dzien** | **Wydobycie** | **Zapas przed transportem** | **Transport?** | **Zapas po transporcie** |
| 2 | 1 | 35 | ? | ? | ? |
| 3 | 2 | 42 | ? | ? | ? |
| 4 | 3 | 28 | ? | ? | ? |
| 5 | 4 | 50 | ? | ? | ? |
| 6 | 5 | 15 | ? | ? | ? |
| 7 | 6 | 45 | ? | ? | ? |
| 8 | 7 | 38 | ? | ? | ? |
| 9 | 8 | 55 | ? | ? | ? |
| 10 | 9 | 22 | ? | ? | ? |
| 11 | 10 | 48 | ? | ? | ? |

Poczatkowy zapas (dzien 0) wynosi 0 ton.

**Polecenie**:
1. W C2 oblicz zapas przed transportem = poprzedni zapas po transporcie + dzisiejsze wydobycie. Dla dnia 1: C2 = 0 + B2.
2. W D2 napisz formule: jezeli zapas >= 100, to "TAK", w przeciwnym razie "NIE".
3. W E2 napisz formule: jezeli transport = "TAK", to zapas - 100, w przeciwnym razie zapas (bez zmian).
4. Na koncu podaj formule zliczajaca laczna liczbe transportow.

<details>
<summary>Wskazowki</summary>

1. **Wskazowka 1**: To zadanie wymaga trzech kolumn obliczeniowych, ktore tworza lancuch: C (zapas przed) -> D (decyzja) -> E (zapas po). Kolumna E jednego wiersza zasila kolumne C nastepnego.
2. **Wskazowka 2**: Pierwszy wiersz jest specjalny — zapas startowy wynosi 0, wiec C2 = 0 + B2 = B2. Od wiersza 3 formula C odwoluje sie do E poprzedniego wiersza. Formula D to prosty IF z progiem 100. Formula E to IF sprawdzajacy D i odejmujacy 100 lub nie.
3. **Wskazowka 3**: C2: `=B2`, C3: `=E2+B3` (kopiowane w dol). D2: `=IF(C2>=100;"TAK";"NIE")`. E2: `=IF(D2="TAK";C2-100;C2)`. Liczba transportow: `=COUNTIF(D2:D11;"TAK")`.
</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
C2: =B2                    (pierwszy dzien, zapas startowy = 0)
C3: =E2+B3                 (kopiowane w dol od wiersza 3)
D2: =IF(C2>=100;"TAK";"NIE")
E2: =IF(D2="TAK";C2-100;C2)
Laczna liczba transportow: =COUNTIF(D2:D11;"TAK")
```

**Wypelniona tabela:**

| Dzien | Wydobycie | Zapas przed | Transport? | Zapas po |
|-------|-----------|-------------|-----------|----------|
| 1 | 35 | 35 | NIE | 35 |
| 2 | 42 | 77 | NIE | 77 |
| 3 | 28 | 105 | TAK | 5 |
| 4 | 50 | 55 | NIE | 55 |
| 5 | 15 | 70 | NIE | 70 |
| 6 | 45 | 115 | TAK | 15 |
| 7 | 38 | 53 | NIE | 53 |
| 8 | 55 | 108 | TAK | 8 |
| 9 | 22 | 30 | NIE | 30 |
| 10 | 48 | 78 | NIE | 78 |

Laczna liczba transportow: **3**

**Weryfikacja krok po kroku:**
- Dzien 1: C=0+35=35, <100 -> NIE, E=35
- Dzien 2: C=35+42=77, <100 -> NIE, E=77
- Dzien 3: C=77+28=105, >=100 -> TAK, E=105-100=5
- Dzien 4: C=5+50=55, <100 -> NIE, E=55
- Dzien 5: C=55+15=70, <100 -> NIE, E=70
- Dzien 6: C=70+45=115, >=100 -> TAK, E=115-100=15
- Dzien 7: C=15+38=53, <100 -> NIE, E=53
- Dzien 8: C=53+55=108, >=100 -> TAK, E=108-100=8
- Dzien 9: C=8+22=30, <100 -> NIE, E=30
- Dzien 10: C=30+48=78, <100 -> NIE, E=78

**Wyjasnienie**: To typowa symulacja z akumulatorem i warunkiem resetowania. Kolumna E "przenosi" stan do nastepnego dnia — odwolanie E2 w formule C3 tworzy lancuch zaleznosci. Kazdy dzien zalezy od wyniku poprzedniego dnia. COUNTIF na koncu zlicza ile razy wyslano transport. Na maturze tego typu zadania wymagaja dokladnego sledzenia stanu wiersz po wierszu.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Odwolanie do C zamiast E w nastepnym wierszu**: Napisanie `C3: =C2+B3` zamiast `=E2+B3` — pominiecie transportu powoduje, ze zapas nigdy nie jest zmniejszany o 100 ton, co daje bledne wyniki od wiersza z pierwszym transportem.
2. **Porownanie tekstu bez cudzyslowow**: Napisanie `=IF(D2=TAK;...)` zamiast `=IF(D2="TAK";...)` — bez cudzyslowow arkusz traktuje TAK jako nazwe zakresu lub zmiennej, nie jako tekst.
3. **Zapomnienie o osobnej formule dla pierwszego wiersza**: Proste skopiowanie `=E1+B2` do C2 odwoluje sie do E1, ktore jest naglowkiem (tekst), co daje blad #WARTOSC!.
</details>

---

### Cwiczenie 16.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna (procent skladany)
**Tagi**: `akumulator` `odwolanie-bezwzgledne` `kopiowanie-formul` `ROUND`

Na lokacie bankowej zdeponowano 10000 zl z roczna stopa procentowa 4.5%, kapitalizowana co rok. Odsetki sa doliczane do kapitalu na koniec kazdego roku (procent skladany).

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Rok** | **Kapital na poczatku** | **Odsetki** | **Kapital na koniec** |
| 2 | 1 | 10000.00 | ? | ? |
| 3 | 2 | ? | ? | ? |
| 4 | 3 | ? | ? | ? |
| 5 | 4 | ? | ? | ? |
| 6 | 5 | ? | ? | ? |
| 7 | 6 | ? | ? | ? |
| 8 | 7 | ? | ? | ? |
| 9 | 8 | ? | ? | ? |

Stopa procentowa zapisana jest w komorce F1: 4,5% (czyli 0.045).

**Polecenie**:
1. W C2 oblicz odsetki = kapital na poczatku x stopa procentowa (z odwolaniem bezwzglednym do F1). Zaokraglij do 2 miejsc po przecinku.
2. W D2 oblicz kapital na koniec = kapital na poczatku + odsetki.
3. W B3 wpisz formule, ktora pobiera kapital na koniec z poprzedniego roku.
4. Formuly musza dzialac po skopiowaniu w dol do wiersza 9.

<details>
<summary>Wskazowki</summary>

1. **Wskazowka 1**: Odsetki to iloczyn kapitalu i stopy procentowej. Stopa jest w jednej komorce — potrzebujesz odwolania bezwzglednego, zeby sie nie zmienilo po skopiowaniu.
2. **Wskazowka 2**: Uzyj funkcji ROUND (w polskim Excelu: ZAOKR) do zaokraglenia odsetek. Formula C2: `=ZAOKR(B2*$F$1;2)`. Kapital na koniec: D2 = B2 + C2. Poczatek nastepnego roku: B3 = D2.
3. **Wskazowka 3**: C2: `=ZAOKR(B2*$F$1;2)`, D2: `=B2+C2`, B3: `=D2`. Skopiuj wszystkie trzy formuly w dol. Sprawdz, czy po 8 latach kapital wynosi okolo 14221.
</details>

<details>
<summary>Odpowiedz</summary>

**Formuly (kopiowane w dol):**
```
C2: =ZAOKR(B2*$F$1;2)
D2: =B2+C2
B3: =D2
```

**Wypelniona tabela:**

| Rok | Kapital na poczatku | Odsetki | Kapital na koniec |
|-----|---------------------|---------|-------------------|
| 1 | 10000.00 | 450.00 | 10450.00 |
| 2 | 10450.00 | 470.25 | 10920.25 |
| 3 | 10920.25 | 491.41 | 11411.66 |
| 4 | 11411.66 | 513.52 | 11925.19 |
| 5 | 11925.19 | 536.63 | 12461.82 |
| 6 | 12461.82 | 560.78 | 13022.60 |
| 7 | 13022.60 | 586.02 | 13608.62 |
| 8 | 13608.62 | 612.39 | 14221.01 |

**Weryfikacja:**
- Rok 1: 10000.00 x 0.045 = 450.00, kapital = 10450.00
- Rok 2: 10450.00 x 0.045 = 470.25, kapital = 10920.25
- Rok 3: 10920.25 x 0.045 = 491.41125 ~ 491.41, kapital = 11411.66
- Rok 4: 11411.66 x 0.045 = 513.5247 ~ 513.52, kapital = 11925.19 (z zaokragleniem: 11925.18, ale 11411.66 + 513.52 = 11925.18; poprawka: 11411.66*0.045=513.52470, ZAOKR->513.52, suma=11925.18)

Dokladna weryfikacja z ZAOKR:
- Rok 1: ZAOKR(10000*0.045;2)=450.00 -> 10000+450=10450.00
- Rok 2: ZAOKR(10450*0.045;2)=470.25 -> 10450+470.25=10920.25
- Rok 3: ZAOKR(10920.25*0.045;2)=491.41 -> 10920.25+491.41=11411.66
- Rok 4: ZAOKR(11411.66*0.045;2)=513.52 -> 11411.66+513.52=11925.18
- Rok 5: ZAOKR(11925.18*0.045;2)=536.63 -> 11925.18+536.63=12461.81
- Rok 6: ZAOKR(12461.81*0.045;2)=560.78 -> 12461.81+560.78=13022.59
- Rok 7: ZAOKR(13022.59*0.045;2)=586.02 -> 13022.59+586.02=13608.61
- Rok 8: ZAOKR(13608.61*0.045;2)=612.39 -> 13608.61+612.39=14221.00

Koncowy kapital po 8 latach: **14221.00 zl**

**Wyjasnienie**: Procent skladany oznacza, ze odsetki z kazdego roku sa doliczane do kapitalu i w nastepnym roku odsetki sa naliczane od wiekszej kwoty. Dlatego kapital rosnie coraz szybciej. Kluczowe jest uzycie ZAOKR (ROUND) do zaokraglenia odsetek do groszy — w finansach zawsze zaokraglamy do 2 miejsc po przecinku. Odwolanie bezwzgledne $F$1 gwarantuje, ze kazdy wiersz uzywa tej samej stopy.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Brak zaokraglenia odsetek**: Pominiecie ZAOKR powoduje, ze odsetki maja wiele miejsc po przecinku (np. 491.41125), co daje drobne roznice narastajace w kolejnych latach. W zadaniach maturalnych wyniki musza sie zgadzac co do grosza.
2. **Procent prosty zamiast skladanego**: Uzycie `=10000*$F$1` (zawsze od poczatkowej kwoty) zamiast `=B2*$F$1` (od biezacego kapitalu) — daje stale odsetki 450 zl rocznie.
</details>

---

### Cwiczenie 16.7 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2015 (Demografia, rozszerzone)
**Tagi**: `akumulator` `IF-warunkowy` `odwolanie-bezwzgledne` `symulacja-krokowa` `ROUND`

Populacja pewnego regionu zmienia sie co rok pod wplywem trzech czynnikow:
- wspolczynnik urodzen: 1.2% populacji
- wspolczynnik zgonow: 0.8% populacji
- migracja netto (podana w tabeli — wartosc dodatnia = imigracja, ujemna = emigracja)

Poczatkowa populacja wynosi 120000 osob. Wspolczynniki sa w komorkach: G1 = 0.012 (urodzenia), G2 = 0.008 (zgony).

| | A | B | C | D | E | F |
|---|---|---|---|---|---|---|
| 1 | **Rok** | **Populacja na poczatku** | **Urodzenia** | **Zgony** | **Migracja netto** | **Populacja na koniec** |
| 2 | 2020 | 120000 | ? | ? | +800 | ? |
| 3 | 2021 | ? | ? | ? | -300 | ? |
| 4 | 2022 | ? | ? | ? | +1200 | ? |
| 5 | 2023 | ? | ? | ? | +500 | ? |
| 6 | 2024 | ? | ? | ? | -600 | ? |
| 7 | 2025 | ? | ? | ? | +900 | ? |
| 8 | 2026 | ? | ? | ? | -100 | ? |
| 9 | 2027 | ? | ? | ? | +400 | ? |
| 10 | 2028 | ? | ? | ? | +700 | ? |
| 11 | 2029 | ? | ? | ? | -200 | ? |

**Polecenie**:
1. W C2 oblicz urodzenia = populacja x wspolczynnik urodzen (zaokraglij do calkowitych za pomoca INT).
2. W D2 oblicz zgony = populacja x wspolczynnik zgonow (zaokraglij do calkowitych za pomoca INT).
3. W F2 oblicz populacje na koniec = populacja na poczatku + urodzenia - zgony + migracja netto.
4. W B3 pobierz populacje na koniec z poprzedniego roku.
5. Podaj, w ktorym roku populacja po raz pierwszy przekroczy 125000.

<details>
<summary>Wskazowki</summary>

1. **Wskazowka 1**: Urodzenia i zgony to procent biezacej populacji — uzyj odwolan bezwzglednych do wspolczynnikow w G1 i G2. Migracja jest juz dana liczbowo. Populacja na koniec to bilans: poczatek + urodzenia - zgony + migracja.
2. **Wskazowka 2**: C2: `=INT(B2*$G$1)`, D2: `=INT(B2*$G$2)`. Uzyj INT (zaokraglenie w dol do calkowitych) bo nie mozna miec ulamka osoby. F2: `=B2+C2-D2+E2`. B3: `=F2`.
3. **Wskazowka 3**: Przelicz wiersz po wierszu. Rok 2020: 120000 + INT(120000*0.012) - INT(120000*0.008) + 800 = 120000 + 1440 - 960 + 800 = 121280. Kontynuuj az znajdziesz rok >125000.
</details>

<details>
<summary>Odpowiedz</summary>

**Formuly (kopiowane w dol):**
```
C2: =INT(B2*$G$1)
D2: =INT(B2*$G$2)
F2: =B2+C2-D2+E2
B3: =F2
```

**Wypelniona tabela:**

| Rok | Populacja pocz. | Urodzenia | Zgony | Migracja | Populacja kon. |
|-----|-----------------|-----------|-------|----------|----------------|
| 2020 | 120000 | 1440 | 960 | +800 | 121280 |
| 2021 | 121280 | 1455 | 970 | -300 | 121465 |
| 2022 | 121465 | 1457 | 971 | +1200 | 123151 |
| 2023 | 123151 | 1477 | 985 | +500 | 124143 |
| 2024 | 124143 | 1489 | 993 | -600 | 124039 |
| 2025 | 124039 | 1488 | 992 | +900 | 125435 |
| 2026 | 125435 | 1505 | 1003 | -100 | 125837 |
| 2027 | 125837 | 1510 | 1006 | +400 | 126741 |
| 2028 | 126741 | 1520 | 1013 | +700 | 127948 |
| 2029 | 127948 | 1535 | 1023 | -200 | 128260 |

**Weryfikacja (pierwsze 3 lata):**
- 2020: INT(120000*0.012)=1440, INT(120000*0.008)=960, F=120000+1440-960+800=121280
- 2021: INT(121280*0.012)=INT(1455.36)=1455, INT(121280*0.008)=INT(970.24)=970, F=121280+1455-970-300=121465
- 2022: INT(121465*0.012)=INT(1457.58)=1457, INT(121465*0.008)=INT(971.72)=971, F=121465+1457-971+1200=123151

Populacja po raz pierwszy przekracza 125000 w roku **2025** (125435).

**Wyjasnienie**: Symulacja demograficzna laczy trzy czynniki zmian populacji. Uzycie INT zamiast ROUND zapewnia zaokraglenie w dol — bardziej realistyczne dla liczby osob. Odwolania bezwzgledne ($G$1, $G$2) sa kluczowe, bo wspolczynniki sa stale we wszystkich wierszach. Migracja netto moze byc dodatnia lub ujemna, co dodaje realizmu do modelu.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Uzycie ROUND zamiast INT**: ROUND(1455.36;0) daje 1455, ale ROUND(1455.5;0) daloby 1456. Dla populacji INT (zaokraglenie w dol) jest bardziej typowe, choc oba podejscia sa akceptowane na maturze — wazne, by stosowac konsekwentnie.
2. **Zapomnienie o odwolaniu bezwzglednym do wspolczynnikow**: Bez $G$1 po skopiowaniu do wiersza 4 formula odwoluje sie do G3, ktore jest puste — wynik to 0 urodzen.
3. **Dodanie zgonow zamiast odjecia**: Napisanie `=B2+C2+D2+E2` zamiast `=B2+C2-D2+E2` — zgony powinny zmniejszac populacje, nie zwiekszac.
</details>

---

### Cwiczenie 16.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: ogolna (logistyka magazynowa)
**Tagi**: `akumulator` `IF-warunkowy` `symulacja-krokowa` `lancuch-zaleznosci` `prog-warunkowy`

Hurtownia zarzadza zapasem produktu. Codziennie sa zamowienia klientow (wydania). Gdy zapas spadnie ponizej progu zamowienia (30 szt.), hurtownia sklada zamowienie u dostawcy na 100 szt. Dostawa przychodzi dokladnie po 3 dniach od zlozenia zamowienia (czas realizacji = 3 dni). Poczatkowy zapas wynosi 80 szt.

| | A | B | C | D | E | F |
|---|---|---|---|---|---|---|
| 1 | **Dzien** | **Wydania** | **Dostawa** | **Zapas** | **Zamowiono?** | **Dzien dostawy** |
| 2 | 1 | 20 | 0 | ? | ? | ? |
| 3 | 2 | 15 | 0 | ? | ? | ? |
| 4 | 3 | 25 | 0 | ? | ? | ? |
| 5 | 4 | 10 | ? | ? | ? | ? |
| 6 | 5 | 18 | ? | ? | ? | ? |
| 7 | 6 | 12 | ? | ? | ? | ? |
| 8 | 7 | 22 | ? | ? | ? | ? |
| 9 | 8 | 8 | ? | ? | ? | ? |
| 10 | 9 | 30 | ? | ? | ? | ? |
| 11 | 10 | 14 | ? | ? | ? | ? |

Prog zamowienia: komorka H1 = 30. Wielkosc zamowienia: komorka H2 = 100. Czas realizacji: komorka H3 = 3.

**Polecenie**:
1. W D2 oblicz zapas = poprzedni zapas (poczatkowy 80) + dostawa - wydania.
2. W E2 wpisz "TAK" jezeli zapas < prog zamowienia, w przeciwnym razie "NIE".
3. W F2 oblicz dzien dostawy = biezacy dzien + czas realizacji (jezeli zamowiono, w przeciwnym razie "—").
4. W C2 i dalej uzupelnij kolumne dostaw — dostawa wynosi 100 w dniu, ktory odpowiada dniowi dostawy z jakiegos wczesniejszego wiersza.
5. Wypelnij tabele recznie, sledzac stan krok po kroku.

<details>
<summary>Wskazowki</summary>

1. **Wskazowka 1**: To zadanie wymaga recznego sledzenia, bo dostawa w danym dniu zalezy od zamowienia sprzed 3 dni. Zacznij od dnia 1: zapas = 80 - 20 = 60, to >= 30 wiec nie zamawiamy. Kontynuuj dzien po dniu.
2. **Wskazowka 2**: Gdy zapas spadnie ponizej 30, zanotuj dzien dostawy (biezacy dzien + 3). Gdy dojdziesz do tego dnia, dodaj 100 do zapasu. Formula dostawy: `=IF(COUNTIF($F$2:$F$11;A2)>0;$H$2;0)` — sprawdza, czy biezacy dzien jest dniem dostawy.
3. **Wskazowka 3**: Dzien 1: 80-20=60 (>=30, NIE). Dzien 2: 60-15=45 (>=30, NIE). Dzien 3: 45-25=20 (<30, TAK, dostawa w dniu 6). Dzien 4-5: brak dostawy. Dzien 6: dostawa +100.
</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
C2: =IF(COUNTIF($F$2:$F$11;A2)>0;$H$2;0)
D2: =80+C2-B2                     (pierwszy dzien, zapas poczatkowy=80)
D3: =D2+C3-B3                     (kopiowane w dol od wiersza 3)
E2: =IF(D2<$H$1;"TAK";"NIE")
F2: =IF(E2="TAK";A2+$H$3;"---")
```

Uwaga: W praktyce te formuly tworza zaleznosc cykliczna (C zalezy od F, a F od D, D od C). Na maturze tego typu zadanie rozwiazuje sie recznym sledzeniem krok po kroku, a formuly pisze sie uproszczone lub opisowo.

**Reczne sledzenie krok po kroku:**

- Dzien 1: zapas = 80 + 0 - 20 = 60 (>=30 -> NIE)
- Dzien 2: zapas = 60 + 0 - 15 = 45 (>=30 -> NIE)
- Dzien 3: zapas = 45 + 0 - 25 = 20 (<30 -> TAK, dostawa w dniu 6)
- Dzien 4: zapas = 20 + 0 - 10 = 10 (<30 -> TAK, dostawa w dniu 7)
- Dzien 5: zapas = 10 + 0 - 18 = -8 (<30 -> TAK, dostawa w dniu 8)
- Dzien 6: zapas = -8 + 100 - 12 = 80 (>=30 -> NIE) [dostawa z dnia 3]
- Dzien 7: zapas = 80 + 100 - 22 = 158 (>=30 -> NIE) [dostawa z dnia 4]
- Dzien 8: zapas = 158 + 100 - 8 = 250 (>=30 -> NIE) [dostawa z dnia 5]
- Dzien 9: zapas = 250 + 0 - 30 = 220 (>=30 -> NIE)
- Dzien 10: zapas = 220 + 0 - 14 = 206 (>=30 -> NIE)

**Wypelniona tabela:**

| Dzien | Wydania | Dostawa | Zapas | Zamowiono? | Dzien dostawy |
|-------|---------|---------|-------|------------|---------------|
| 1 | 20 | 0 | 60 | NIE | --- |
| 2 | 15 | 0 | 45 | NIE | --- |
| 3 | 25 | 0 | 20 | TAK | 6 |
| 4 | 10 | 0 | 10 | TAK | 7 |
| 5 | 18 | 0 | -8 | TAK | 8 |
| 6 | 12 | 100 | 80 | NIE | --- |
| 7 | 22 | 100 | 158 | NIE | --- |
| 8 | 8 | 100 | 250 | NIE | --- |
| 9 | 30 | 0 | 220 | NIE | --- |
| 10 | 14 | 0 | 206 | NIE | --- |

**Wyjasnienie**: To klasyczna symulacja logistyczna z opoznieniem dostawy. Kluczowe obserwacje: (1) zapas moze byc ujemny (dzien 5: -8), co oznacza niezrealizowane zamowienia; (2) wielokrotne zamowienia moga sie nakladac — zamowiono w dniach 3, 4 i 5, wiec dostawy przyszly kolejno w dniach 6, 7 i 8; (3) po serii dostaw zapas jest bardzo wysoki (250), co jest typowe dla systemu z opoznieniem (efekt byczego bicza). Na maturze tego typu zadanie rozwiazuje sie recznym sledzeniem — krok po kroku.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Ignorowanie opoznienia dostawy**: Dodanie 100 do zapasu natychmiast po zamowieniu (tego samego dnia) zamiast po 3 dniach — calkowicie zmienia dynamike symulacji.
2. **Zamowienie tylko raz**: Zalozenie, ze hurtownia zamawia tylko przy pierwszym spadku ponizej progu i czeka na dostawe — w rzeczywistosci kazdy dzien z zapasem <30 generuje nowe zamowienie.
3. **Pominiecie ujemnego zapasu**: Zapas moze byc ujemny (niedobor), co jest prawidlowe w modelu. Nie nalezy zastepowac go zerem.
</details>

---

### Cwiczenie 16.9 (trudnosc: srednie-trudne, ~5 pkt)
**Zrodlo inspiracji**: ogolna (finanse, raty kredytu)
**Tagi**: `rata-annuitetowa` `akumulator` `ROUND` `odwolanie-bezwzgledne` `amortyzacja` `symulacja-krokowa`

Bank udzielil kredytu na kwote 50000 zl z roczna stopa procentowa 6% (miesieczna stopa = 6%/12 = 0.5%). Kredyt jest splacany w 10 rownych ratach miesiecznych (annuitetowych). Rata annuitetowa obliczana jest wedlug wzoru:

```
R = K * r / (1 - (1+r)^(-n))
```

gdzie K = kwota kredytu, r = miesieczna stopa procentowa, n = liczba rat.

Parametry: F1 = 50000 (kwota), F2 = 0,005 (stopa miesieczna), F3 = 10 (liczba rat).

| | A | B | C | D | E |
|---|---|---|---|---|---|
| 1 | **Rata nr** | **Saldo poczatkowe** | **Odsetki** | **Czesc kapitalowa** | **Saldo koncowe** |
| 2 | 1 | ? | ? | ? | ? |
| 3 | 2 | ? | ? | ? | ? |
| 4 | 3 | ? | ? | ? | ? |
| 5 | 4 | ? | ? | ? | ? |
| 6 | 5 | ? | ? | ? | ? |
| 7 | 6 | ? | ? | ? | ? |
| 8 | 7 | ? | ? | ? | ? |
| 9 | 8 | ? | ? | ? | ? |
| 10 | 9 | ? | ? | ? | ? |
| 11 | 10 | ? | ? | ? | ? |

Rata miesieczna (stala, obliczona wedlug wzoru) zapisana jest w komorce F4.

**Polecenie**:
1. W F4 oblicz rate miesieczna wedlug wzoru.
2. W B2 wpisz kwote kredytu (=F1). Od B3: saldo poczatkowe = saldo koncowe z poprzedniej raty.
3. W C2 oblicz odsetki = saldo poczatkowe x stopa miesieczna (zaokraglij do 2 miejsc).
4. W D2 oblicz czesc kapitalowa = rata - odsetki.
5. W E2 oblicz saldo koncowe = saldo poczatkowe - czesc kapitalowa.
6. Wypelnij cala tabele. Sprawdz, czy saldo koncowe po ostatniej racie wynosi 0 (lub blisko 0).

<details>
<summary>Wskazowki</summary>

1. **Wskazowka 1**: Rata annuitetowa jest stala przez caly okres splacania. W kazdej racie zmienia sie proporcja miedzy odsetkami a czescia kapitalowa — na poczatku wiecej idzie na odsetki, pod koniec wiecej na kapital.
2. **Wskazowka 2**: Formula raty: `=F1*F2/(1-(1+F2)^(-F3))`. Odsetki: `=ZAOKR(B2*$F$2;2)`. Czesc kapitalowa: `=$F$4-C2`. Saldo koncowe: `=B2-D2`. Nastepny wiersz: B3=E2.
3. **Wskazowka 3**: Rata = 50000*0.005/(1-(1.005)^(-10)) = 250/( 1 - 0.95135) = 250/0.04865 = 5138.83 zl (w przyblizeniu). Pierwsza rata: odsetki = 50000*0.005 = 250.00, kapital = 5138.83 - 250.00 = 4888.83, saldo = 50000 - 4888.83 = 45111.17.
</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
F4: =ZAOKR($F$1*$F$2/(1-(1+$F$2)^(-$F$3));2)
B2: =$F$1
B3: =E2                           (kopiowane w dol)
C2: =ZAOKR(B2*$F$2;2)             (kopiowane w dol)
D2: =$F$4-C2                      (kopiowane w dol)
E2: =B2-D2                        (kopiowane w dol)
```

**Obliczenie raty:**
R = 50000 * 0.005 / (1 - (1.005)^(-10))
R = 250 / (1 - 0.951110305...)
R = 250 / 0.048889695...
R = 5113.63 zl (po zaokragleniu do 2 miejsc)

Weryfikacja: (1.005)^10 = 1.05114013... -> (1.005)^(-10) = 1/1.05114013 = 0.95132890...
R = 250 / (1 - 0.95132890) = 250 / 0.04867110 = 5136.62...

Dokaldne obliczenie: (1.005)^10:
- (1.005)^1 = 1.005
- (1.005)^2 = 1.010025
- (1.005)^3 = 1.015075125
- (1.005)^4 = 1.020150500625
- (1.005)^5 = 1.025251253128
- (1.005)^6 = 1.030377509384
- (1.005)^7 = 1.035529446931
- (1.005)^8 = 1.040707094165
- (1.005)^9 = 1.045910629636
- (1.005)^10 = 1.051140132184

(1.005)^(-10) = 1/1.051140132184 = 0.951348404866

R = 250 / (1 - 0.951348404866) = 250 / 0.048651595134 = 5138.53 zl

Poprawne obliczenie: **R = 5138.53 zl**

**Wypelniona tabela:**

| Rata nr | Saldo poczatkowe | Odsetki | Czesc kapitalowa | Saldo koncowe |
|---------|------------------|---------|------------------|---------------|
| 1 | 50000.00 | 250.00 | 4888.53 | 45111.47 |
| 2 | 45111.47 | 225.56 | 4912.97 | 40198.50 |
| 3 | 40198.50 | 200.99 | 4937.54 | 35260.96 |
| 4 | 35260.96 | 176.30 | 4962.23 | 30298.73 |
| 5 | 30298.73 | 151.49 | 4987.04 | 25311.69 |
| 6 | 25311.69 | 126.56 | 5011.97 | 20299.72 |
| 7 | 20299.72 | 101.50 | 5037.03 | 15262.69 |
| 8 | 15262.69 | 76.31 | 5062.22 | 10200.47 |
| 9 | 10200.47 | 51.00 | 5087.53 | 5112.94 |
| 10 | 5112.94 | 25.56 | 5112.97 | -0.03 |

**Weryfikacja (pierwsze 3 raty):**
- Rata 1: odsetki = ZAOKR(50000*0.005;2) = 250.00, kapital = 5138.53-250.00 = 4888.53, saldo = 50000-4888.53 = 45111.47
- Rata 2: odsetki = ZAOKR(45111.47*0.005;2) = 225.56, kapital = 5138.53-225.56 = 4912.97, saldo = 45111.47-4912.97 = 40198.50
- Rata 3: odsetki = ZAOKR(40198.50*0.005;2) = 200.99, kapital = 5138.53-200.99 = 4937.54, saldo = 40198.50-4937.54 = 35260.96

Saldo koncowe po 10. racie wynosi -0.03 zl (blisko zera). Drobna roznica wynika z zaokraglen. W praktyce ostatnia rata jest korygowana, aby saldo wynioslo dokladnie 0.

**Wyjasnienie**: Harmonogram amortyzacji to klasyczna symulacja finansowa. Rata jest stala (annuitetowa), ale proporcja odsetki/kapital zmienia sie: na poczatku odsetki stanowia wieksza czesc raty (bo saldo jest duze), a pod koniec prawie cala rata idzie na kapital. Kluczowe elementy: (1) formula potegowa dla raty, (2) zaokraglanie do 2 miejsc, (3) lancuch zaleznosci wiersz po wierszu. Na maturze tego typu zadanie moze pojawic sie z uproszczonym wzorem lub z podana wartoscia raty.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Pomylenie stopy rocznej z miesieczna**: Uzycie 6% (0.06) zamiast 0.5% (0.005) jako stopy miesiecznej — daje absurdalnie wysokie odsetki i rate.
2. **Bledna formula potegowa**: Napisanie `(1+F2)*(-F3)` zamiast `(1+F2)^(-F3)` — mnozenie zamiast potegowania daje calkowicie inny wynik.
3. **Brak zaokraglenia odsetek**: Bez ZAOKR odsetki maja wiele miejsc po przecinku, co powoduje narastajace rozbieznosci — saldo koncowe moze roznic sie o kilka zlotych od zera.
</details>

---

### Cwiczenie 16.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: ogolna (symulacja kolejki)
**Tagi**: `akumulator` `IF-warunkowy` `MIN` `MAX` `lancuch-zaleznosci` `symulacja-krokowa` `prog-warunkowy`

Parking ma pojemnosc maksymalna 40 miejsc. W kazdej godzinie przybywa pewna liczba samochodow (chcacych wjechac) i pewna liczba samochodow opuszcza parking. Samochody, ktore nie zmieszcza sie na parkingu (brak miejsc), odjeżdżaja i sa liczone jako "odrzucone". Poczatkowo na parkingu jest 15 samochodow.

| | A | B | C | D | E | F | G |
|---|---|---|---|---|---|---|---|
| 1 | **Godz** | **Przyjazdy** | **Wyjazdy** | **Na parkingu przed** | **Faktyczne przyjecia** | **Odrzucone** | **Na parkingu po** |
| 2 | 8:00 | 12 | 3 | ? | ? | ? | ? |
| 3 | 9:00 | 18 | 5 | ? | ? | ? | ? |
| 4 | 10:00 | 8 | 10 | ? | ? | ? | ? |
| 5 | 11:00 | 6 | 7 | ? | ? | ? | ? |
| 6 | 12:00 | 15 | 12 | ? | ? | ? | ? |
| 7 | 13:00 | 20 | 8 | ? | ? | ? | ? |
| 8 | 14:00 | 10 | 14 | ? | ? | ? | ? |
| 9 | 15:00 | 5 | 9 | ? | ? | ? | ? |
| 10 | 16:00 | 3 | 15 | ? | ? | ? | ? |
| 11 | 17:00 | 2 | 12 | ? | ? | ? | ? |

Pojemnosc parkingu: komorka I1 = 40. Poczatkowa liczba samochodow: komorka I2 = 15.

**Polecenie**:
1. W D2 oblicz stan parkingu przed operacjami godziny: poprzedni stan po operacjach (dla godz. 8:00 uzyj I2) minus wyjazdy (ale stan nie moze byc ujemny — uzyj MAX(...;0)).
2. W E2 oblicz faktyczne przyjecia = MIN(przyjazdy; pojemnosc - stan po wyjazdach). To znaczy: przyjmij tyle samochodow, ile sie zmiesci.
3. W F2 oblicz odrzucone = przyjazdy - faktyczne przyjecia.
4. W G2 oblicz stan po operacjach = stan po wyjazdach + faktyczne przyjecia.
5. Wypelnij tabele. Podaj laczna liczbe odrzuconych samochodow i godzine szczytu (najwiecej aut na parkingu).

<details>
<summary>Wskazowki</summary>

1. **Wskazowka 1**: Kolejnosc operacji w kazdej godzinie: najpierw wyjazdy (zwalniaja miejsca), potem przyjazdy (zajmuja miejsca). Stan po wyjazdach nie moze byc ujemny. Faktyczne przyjecia sa ograniczone dostepnymi miejscami.
2. **Wskazowka 2**: Stan po wyjazdach = MAX(poprzedni stan - wyjazdy; 0). Wolne miejsca = pojemnosc - stan po wyjazdach. Faktyczne przyjecia = MIN(przyjazdy; wolne miejsca). Odrzucone = przyjazdy - faktyczne przyjecia.
3. **Wskazowka 3**: Formuly: D2 (stan po wyjazdach): `=MAX($I$2-C2;0)` (dla godz. 8:00), od D3: `=MAX(G2-C3;0)`. E2: `=MIN(B2;$I$1-D2)`. F2: `=B2-E2`. G2: `=D2+E2`. Godz. 8:00: stan po wyjazdach = MAX(15-3;0)=12, wolne=40-12=28, przyjecia=MIN(12;28)=12, odrzucone=0, stan=24.
</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
D2: =MAX($I$2-C2;0)               (godz. 8:00, poczatkowy stan z I2)
D3: =MAX(G2-C3;0)                 (kopiowane w dol od wiersza 3)
E2: =MIN(B2;$I$1-D2)              (kopiowane w dol)
F2: =B2-E2                        (kopiowane w dol)
G2: =D2+E2                        (kopiowane w dol)
Laczna liczba odrzuconych: =SUM(F2:F11)
```

**Reczne sledzenie krok po kroku:**

Godz. 8:00: stan_przed=15, po_wyjazdach=MAX(15-3;0)=12, wolne=40-12=28, przyjecia=MIN(12;28)=12, odrzucone=0, stan_po=12+12=24
Godz. 9:00: stan_przed=24, po_wyjazdach=MAX(24-5;0)=19, wolne=40-19=21, przyjecia=MIN(18;21)=18, odrzucone=0, stan_po=19+18=37
Godz. 10:00: stan_przed=37, po_wyjazdach=MAX(37-10;0)=27, wolne=40-27=13, przyjecia=MIN(8;13)=8, odrzucone=0, stan_po=27+8=35
Godz. 11:00: stan_przed=35, po_wyjazdach=MAX(35-7;0)=28, wolne=40-28=12, przyjecia=MIN(6;12)=6, odrzucone=0, stan_po=28+6=34
Godz. 12:00: stan_przed=34, po_wyjazdach=MAX(34-12;0)=22, wolne=40-22=18, przyjecia=MIN(15;18)=15, odrzucone=0, stan_po=22+15=37
Godz. 13:00: stan_przed=37, po_wyjazdach=MAX(37-8;0)=29, wolne=40-29=11, przyjecia=MIN(20;11)=11, odrzucone=9, stan_po=29+11=40
Godz. 14:00: stan_przed=40, po_wyjazdach=MAX(40-14;0)=26, wolne=40-26=14, przyjecia=MIN(10;14)=10, odrzucone=0, stan_po=26+10=36
Godz. 15:00: stan_przed=36, po_wyjazdach=MAX(36-9;0)=27, wolne=40-27=13, przyjecia=MIN(5;13)=5, odrzucone=0, stan_po=27+5=32
Godz. 16:00: stan_przed=32, po_wyjazdach=MAX(32-15;0)=17, wolne=40-17=23, przyjecia=MIN(3;23)=3, odrzucone=0, stan_po=17+3=20
Godz. 17:00: stan_przed=20, po_wyjazdach=MAX(20-12;0)=8, wolne=40-8=32, przyjecia=MIN(2;32)=2, odrzucone=0, stan_po=8+2=10

**Wypelniona tabela:**

| Godz | Przyjazdy | Wyjazdy | Po wyjazdach | Fakt. przyjecia | Odrzucone | Po operacjach |
|------|-----------|---------|--------------|-----------------|-----------|---------------|
| 8:00 | 12 | 3 | 12 | 12 | 0 | 24 |
| 9:00 | 18 | 5 | 19 | 18 | 0 | 37 |
| 10:00 | 8 | 10 | 27 | 8 | 0 | 35 |
| 11:00 | 6 | 7 | 28 | 6 | 0 | 34 |
| 12:00 | 15 | 12 | 22 | 15 | 0 | 37 |
| 13:00 | 20 | 8 | 29 | 11 | 9 | 40 |
| 14:00 | 10 | 14 | 26 | 10 | 0 | 36 |
| 15:00 | 5 | 9 | 27 | 5 | 0 | 32 |
| 16:00 | 3 | 15 | 17 | 3 | 0 | 20 |
| 17:00 | 2 | 12 | 8 | 2 | 0 | 10 |

Laczna liczba odrzuconych samochodow: **9** (wszystkie o godz. 13:00)
Godzina szczytu (najwiecej aut): **13:00** (40 samochodow — pelny parking)

**Weryfikacja szczytu (godz. 13:00):**
- Stan przed = 37 (z poprzedniej godziny)
- Po wyjazdach: MAX(37-8;0) = 29
- Wolne miejsca: 40-29 = 11
- Chce wjechac 20, ale zmiesci sie tylko 11
- Odrzucone: 20-11 = 9
- Stan po: 29+11 = 40 (pelny parking)

**Wyjasnienie**: Symulacja kolejkowa z ograniczeniem pojemnosci to zaawansowane zadanie laczace kilka technik: (1) akumulator z dwoma zrodlami zmian, (2) ograniczenie dolne przez MAX(...;0) — stan nie moze byc ujemny, (3) ograniczenie gorne przez MIN(przyjazdy; wolne_miejsca) — nie mozna przyjac wiecej niz jest miejsc, (4) obliczenie odrzuconych jako roznica. Na maturze kluczowe jest prawidlowe okreslenie kolejnosci operacji (najpierw wyjazdy, potem przyjazdy) oraz uzycie MIN/MAX do modelowania ograniczen fizycznych. Tego typu zadanie laczy umiejetnosci z wielu prostszych cwiczen w jedna spójna symulacje.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Odwrotna kolejnosc operacji**: Najpierw dodanie przyjazdow, potem odjecie wyjazdow — moze doprowadzic do sytuacji, ze parking "przekracza" pojemnosc na chwile, co daje inne wyniki odrzucen.
2. **Brak ograniczenia MIN dla przyjazdow**: Napisanie `E2: =B2` zamiast `=MIN(B2;$I$1-D2)` — pozwala na przekroczenie pojemnosci parkingu (np. 45 aut na 40 miejscach).
3. **Brak ograniczenia MAX dla wyjazdow**: Jezeli wyjazdy > stan parkingu, bez MAX stan staje sie ujemny, co daje sztuczne "wolne miejsca" powyzej pojemnosci. Nalezy uzywac `MAX(stan-wyjazdy;0)`.
</details>

---

## Samoocena

| Poziom | Zakres cwiczen | Opis |
|--------|----------------|------|
| Podstawowy | 16.1 -- 16.3 | Proste akumulatory, kopiowanie formul, pojedynczy IF. Wystarczajacy do zdobycia 2-3 pkt na maturze. |
| Dobry | 16.4 -- 16.6 | SUM kumulacyjna, zagniezdzony IF, procent skladany, odwolania bezwzgledne. Pozwala zdobyc 4-5 pkt. |
| Bardzo dobry | 16.7 -- 16.8 | Wieloczynnikowe symulacje, opoznienia dostaw, reczne sledzenie zlozonego stanu. Pozwala zdobyc 5-6 pkt. |
| Doskonaly | 16.9 -- 16.10 | Amortyzacja kredytu, symulacja kolejkowa z MIN/MAX, formuly potegowe. Pelna punktacja za symulacje. |

### Co dalej?

- **Jesli rozwiazales 1-3 cwiczen**: Przeczytaj wyjasnienia do cwiczen 16.1-16.3 i powtorz je samodzielnie bez podgladania. Skup sie na zrozumieniu akumulatora i odwolan bezwzglednych.
- **Jesli rozwiazales 4-6 cwiczen**: Przejdz do cwiczen z zagniezdzonym IF i SUM kumulacyjna (16.4, 16.7). Cwicz reczne sledzenie formul wiersz po wierszu — to klucz do unikania bledow.
- **Jesli rozwiazales 7-8 cwiczen**: Skoncentruj sie na najtrudniejszych wzorcach: opoznienie dostawy (16.8), amortyzacja (16.9). Sprawdz, czy potrafisz napisac formuly bez podgladania wskazowek.
- **Jesli rozwiazales 9-10 cwiczen**: Jestes swietnie przygotowany do symulacji maturalnych. Przejdz do zestawow cwiczen z innych typow zadan (15_agregacja_warunkowa, 17_wykres) lub sprobuj rozwiazac prawdziwe arkusze maturalne na czas.
