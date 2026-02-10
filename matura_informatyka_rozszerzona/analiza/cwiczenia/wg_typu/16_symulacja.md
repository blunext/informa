# 16. Symulacja w arkuszu kalkulacyjnym

Typ zadania: **arkusz_symulacja**
Czestotliwosc: 9/11 lat | Laczna punktacja: 37 pkt
Kategoria: ARKUSZ

---

### Cwiczenie 16.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna (akumulator prosty)

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

---

### Cwiczenie 16.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 (Demografia)

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
- B3: 50000 × 1.025 = 51250.00
- B4: 51250 × 1.025 = 52531.25 ≈ 52531
- B5: 52531.25 × 1.025 = 53844.53 ≈ 53845
- B6: 53844.53 × 1.025 = 55190.64 ≈ 55191
- B7: 55190.64 × 1.025 = 56570.41 ≈ 56571

**Wyjasnienie**: Znak `$` przed D i 1 (`$D$1`) sprawia, ze odwolanie jest bezwzgledne — nie zmienia sie przy kopiowaniu formuly. Bez `$` kopiowanie z B3 do B4 zmieniloby D1 na D2, co daloby bledny wynik.
</details>

---

### Cwiczenie 16.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2023 (Konfitury)

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
- D3: 200 + 0 - 80 = 120 (≥50, brak alarmu)
- D4: 120 + 50 - 60 = 110 (≥50, brak alarmu)
- D5: 110 + 0 - 70 = 40 (<50, NISKI STAN)
- D6: 40 + 100 - 30 = 110 (≥50, brak alarmu)
- D7: 110 + 0 - 90 = 20 (<50, NISKI STAN)
- D8: 20 + 20 - 50 = -10 (<50, NISKI STAN)
- D9: -10 + 0 - 40 = -50 (<50, NISKI STAN)

**Wyjasnienie**: Kolumna D to akumulator z dwoma zrodlami zmian (przyjecia i wydania). Kolumna E uzywa IF do warunkowego wyswietlania tekstu. Ujemny zapas oznacza, ze magazyn jest "na minusie" (zamowienia przekraczaja stan).
</details>

---

### Cwiczenie 16.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)

Sklep internetowy stosuje rabaty kumulacyjne — rabat zalezy od lacznej kwoty wszystkich dotychczasowych zakupow klienta:
- laczna kwota < 500 zl → rabat 0%
- laczna kwota 500-999 zl → rabat 5%
- laczna kwota 1000-1999 zl → rabat 10%
- laczna kwota >= 2000 zl → rabat 15%

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
3. W E2 oblicz cene po rabacie = kwota × (1 - rabat).

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
- C2: SUM(B2:B2) = 350 → prog 0% → E2: 350×1.00 = 350.00
- C3: SUM(B2:B3) = 350+280 = 630 → prog 5% → E3: 280×0.95 = 266.00
- C4: SUM(B2:B4) = 630+420 = 1050 → prog 10% → E4: 420×0.90 = 378.00
- C5: SUM(B2:B5) = 1050+150 = 1200 → prog 10% → E5: 150×0.90 = 135.00
- C6: SUM(B2:B6) = 1200+600 = 1800 → prog 10% → E6: 600×0.90 = 540.00
- C7: SUM(B2:B7) = 1800+300 = 2100 → prog 15% → E7: 300×0.85 = 255.00

**Wyjasnienie**: Kluczowy jest zakres `$B$2:B2` — poczatek jest zakotwiczony ($B$2), ale koniec (B2) jest wzgledny. Po skopiowaniu do wiersza 5 formula staje sie `=SUM($B$2:B5)`, co daje sume od poczatku do biezacego wiersza. Zagniezdzony IF sprawdza progi od najwiekszego, bo pierwszy speliony warunek konczy ewaluacje.
</details>

---

### Cwiczenie 16.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 (Martianeum)

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
- Dzien 1: C=0+35=35, <100 → NIE, E=35
- Dzien 2: C=35+42=77, <100 → NIE, E=77
- Dzien 3: C=77+28=105, ≥100 → TAK, E=105-100=5
- Dzien 4: C=5+50=55, <100 → NIE, E=55
- Dzien 5: C=55+15=70, <100 → NIE, E=70
- Dzien 6: C=70+45=115, ≥100 → TAK, E=115-100=15
- Dzien 7: C=15+38=53, <100 → NIE, E=53
- Dzien 8: C=53+55=108, ≥100 → TAK, E=108-100=8
- Dzien 9: C=8+22=30, <100 → NIE, E=30
- Dzien 10: C=30+48=78, <100 → NIE, E=78

**Wyjasnienie**: To typowa symulacja z akumulatorem i warunkiem resetowania. Kolumna E "przenosi" stan do nastepnego dnia — odwolanie E2 w formule C3 tworzy lancuch zaleznosci. Kazdy dzien zalezy od wyniku poprzedniego dnia. COUNTIF na koncu zlicza ile razy wyslano transport. Na maturze tego typu zadania wymagaja dokladnego sledzenia stanu wiersz po wierszu.
</details>
