# 15. Agregacja warunkowa w arkuszu kalkulacyjnym

Typ zadania: **arkusz_agregacja_warunkowa**
Czestotliwosc: 9/11 lat | Laczna punktacja: 38 pkt
Kategoria: ARKUSZ

---

### Cwiczenie 15.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)

W arkuszu kalkulacyjnym mamy liste produktow w sklepie:

| | A | B | C |
|---|---|---|---|
| 1 | **Produkt** | **Kategoria** | **Cena** |
| 2 | Mleko | Nabiał | 4.50 |
| 3 | Chleb | Pieczywo | 5.20 |
| 4 | Jogurt | Nabiał | 3.80 |
| 5 | Bułka | Pieczywo | 0.60 |
| 6 | Ser | Nabiał | 8.90 |
| 7 | Masło | Nabiał | 7.40 |
| 8 | Bagietka | Pieczywo | 3.50 |
| 9 | Kefir | Nabiał | 4.20 |
| 10 | Rogal | Pieczywo | 1.80 |
| 11 | Śmietana | Nabiał | 2.90 |

**Polecenie**: Napisz formule, ktora obliczy ile produktow nalezy do kategorii "Nabiał".

<details>
<summary>Odpowiedz</summary>

**Formula:**
```
=COUNTIF(B2:B11;"Nabiał")
```

**Weryfikacja**: Produkty z kategorii "Nabiał": Mleko, Jogurt, Ser, Masło, Kefir, Śmietana = **6**

**Wyjasnienie**: COUNTIF(zakres; kryterium) zlicza komorki w zakresie B2:B11, ktore zawieraja tekst "Nabiał". Srednik oddziela argumenty (konwencja polska/europejska).
</details>

---

### Cwiczenie 15.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 (Demografia)

W arkuszu mamy wyniki uczniow klasy:

| | A | B | C |
|---|---|---|---|
| 1 | **Imie** | **Plec** | **Ocena** |
| 2 | Anna | K | 4 |
| 3 | Bartek | M | 3 |
| 4 | Celina | K | 5 |
| 5 | Dawid | M | 4 |
| 6 | Ewa | K | 3 |
| 7 | Filip | M | 2 |
| 8 | Gosia | K | 5 |
| 9 | Hubert | M | 4 |
| 10 | Iza | K | 4 |
| 11 | Jan | M | 3 |

**Polecenie**: Napisz formuly obliczajace srednia ocene osobno dla chlopcow (M) i dziewczyn (K).

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
Srednia chlopcow: =AVERAGEIF(B2:B11;"M";C2:C11)
Srednia dziewczyn: =AVERAGEIF(B2:B11;"K";C2:C11)
```

**Weryfikacja**:
- Chlopcy (M): Bartek=3, Dawid=4, Filip=2, Hubert=4, Jan=3 → (3+4+2+4+3)/5 = 16/5 = **3.20**
- Dziewczyny (K): Anna=4, Celina=5, Ewa=3, Gosia=5, Iza=4 → (4+5+3+5+4)/5 = 21/5 = **4.20**

**Wyjasnienie**: AVERAGEIF(zakres_kryterium; kryterium; zakres_sredniej) oblicza srednia tylko tych wartosci z C2:C11, dla ktorych odpowiadajaca komorka w B2:B11 spelnia warunek.
</details>

---

### Cwiczenie 15.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)

W arkuszu mamy dane sprzedazy hurtowni:

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Data** | **Region** | **Kwartal** | **Sprzedaz** |
| 2 | 2024-01-15 | Północ | Q1 | 12000 |
| 3 | 2024-02-20 | Południe | Q1 | 8500 |
| 4 | 2024-03-10 | Północ | Q1 | 15000 |
| 5 | 2024-04-05 | Południe | Q2 | 9200 |
| 6 | 2024-05-18 | Północ | Q2 | 11000 |
| 7 | 2024-06-22 | Południe | Q2 | 7800 |
| 8 | 2024-07-14 | Północ | Q3 | 13500 |
| 9 | 2024-08-30 | Południe | Q3 | 10200 |
| 10 | 2024-09-05 | Północ | Q3 | 14000 |
| 11 | 2024-10-12 | Południe | Q4 | 9500 |
| 12 | 2024-11-20 | Północ | Q4 | 16000 |
| 13 | 2024-12-15 | Południe | Q4 | 11000 |

**Polecenie**: Napisz formule SUMIFS, ktora obliczy laczna sprzedaz regionu "Północ" w kwartale "Q3".

<details>
<summary>Odpowiedz</summary>

**Formula:**
```
=SUMIFS(D2:D13;B2:B13;"Północ";C2:C13;"Q3")
```

**Weryfikacja**:
Szukamy wierszy, gdzie Region="Północ" ORAZ Kwartal="Q3":
- Wiersz 8: Region=Północ, Kwartal=Q3, Sprzedaz=13500 ✓
- Wiersz 10: Region=Północ, Kwartal=Q3, Sprzedaz=14000 ✓

Suma: 13500 + 14000 = **27500**

**Wyjasnienie**: SUMIFS(zakres_sumy; zakres_kryt1; kryterium1; zakres_kryt2; kryterium2) sumuje wartosci z D2:D13 tylko dla wierszy spelniajacych OBA warunki jednoczesnie. Roznica wzgledem SUMIF: SUMIFS akceptuje wiele par zakres-kryterium.
</details>

---

### Cwiczenie 15.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 (Pogoda)

W arkuszu mamy dane pogodowe z 12 miesiecy:

| | A | B | C |
|---|---|---|---|
| 1 | **Dzien** | **Miesiac** | **Temperatura** |
| 2 | 1 | Styczeń | -5.2 |
| 3 | 2 | Styczeń | -3.1 |
| 4 | 3 | Styczeń | 0.4 |
| 5 | 4 | Luty | -1.8 |
| 6 | 5 | Luty | 2.3 |
| 7 | 6 | Luty | -0.5 |
| 8 | 7 | Marzec | 5.1 |
| 9 | 8 | Marzec | 8.7 |
| 10 | 9 | Marzec | 3.2 |
| 11 | 10 | Marzec | 7.4 |

Zalozmy, ze normy temperatur sa podane w osobnej tabeli:

| | E | F |
|---|---|---|
| 1 | **Miesiac** | **Norma** |
| 2 | Styczeń | -2.0 |
| 3 | Luty | 0.0 |
| 4 | Marzec | 5.0 |

**Polecenie**:
1. Napisz formule obliczajaca srednia temperature w styczniu (AVERAGEIF).
2. Napisz formule zliczajaca dni w marcu z temperatura powyzej normy (5.0°C) uzywajac COUNTIFS.

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
1. Srednia temperatura w styczniu:
   =AVERAGEIF(B2:B11;"Styczeń";C2:C11)

2. Dni w marcu powyzej normy:
   =COUNTIFS(B2:B11;"Marzec";C2:C11;">"&F4)
```

**Weryfikacja**:

1. Srednia w styczniu:
   - Dni w styczniu: -5.2, -3.1, 0.4
   - Srednia: (-5.2 + (-3.1) + 0.4) / 3 = -7.9/3 = **-2.633...**

2. Dni w marcu powyzej normy (>5.0):
   - Marzec: 5.1 (>5.0 ✓), 8.7 (>5.0 ✓), 3.2 (≤5.0 ✗), 7.4 (>5.0 ✓)
   - Wynik: **3** dni

**Wyjasnienie**:
- AVERAGEIF dziala analogicznie do SUMIF — uśrednia wartosci C tylko tam, gdzie B spelnia kryterium.
- W COUNTIFS kryterium `">"&F4` laczy operator porownania `">"` z wartoscia komorki F4 (5.0), tworzac warunek `">5"`. Operator `&` sluzy do laczenia tekstu (konkatenacji).
</details>

---

### Cwiczenie 15.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 (Martianeum)

Fabryka produkuje czesci metalowe. Kazda czesc ma wymiary i wagi mierzone na 3 stanowiskach kontroli jakosci. Czesc przechodzi kontrole, jezeli:
- srednia waga z 3 pomiarow miesci sie w zakresie [95g, 105g]
- zaden pojedynczy pomiar nie odbiega od sredniej o wiecej niz 5g

| | A | B | C | D | E |
|---|---|---|---|---|---|
| 1 | **ID** | **Typ** | **Pomiar1** | **Pomiar2** | **Pomiar3** |
| 2 | P001 | A | 98 | 101 | 100 |
| 3 | P002 | B | 94 | 96 | 95 |
| 4 | P003 | A | 102 | 108 | 99 |
| 5 | P004 | A | 100 | 100 | 101 |
| 6 | P005 | B | 92 | 88 | 95 |
| 7 | P006 | A | 103 | 104 | 105 |
| 8 | P007 | B | 97 | 99 | 98 |
| 9 | P008 | A | 110 | 112 | 108 |
| 10 | P009 | B | 100 | 101 | 99 |
| 11 | P010 | A | 96 | 97 | 98 |

**Polecenie**:
1. W kolumnie F oblicz srednia wage kazdej czesci (F2: srednia z C2:E2).
2. W kolumnie G sprawdz warunek zakresu: czy srednia F jest miedzy 95 a 105 (TRUE/FALSE).
3. W kolumnie H sprawdz warunek odchylenia: czy kazdy pomiar jest w odleglosci ≤5 od sredniej (TRUE/FALSE).
4. W kolumnie I wyznacz wynik kontroli: "OK" jezeli G=TRUE i H=TRUE, "ODRZUT" w przeciwnym razie.
5. Napisz formule obliczajaca laczna mase (sume srednich wag) czesci typu "A", ktore przeszly kontrole ("OK").

<details>
<summary>Odpowiedz</summary>

**Formuly (w wierszu 2, kopiowane w dol):**
```
F2: =AVERAGE(C2:E2)
G2: =AND(F2>=95;F2<=105)
H2: =AND(ABS(C2-F2)<=5;ABS(D2-F2)<=5;ABS(E2-F2)<=5)
I2: =IF(AND(G2;H2);"OK";"ODRZUT")
```

**Obliczenia krok po kroku:**

| ID | Typ | P1 | P2 | P3 | F: Srednia | G: Zakres? | H: Odchylenie? | I: Wynik |
|----|-----|----|----|----|-----------|-----------|----------------|----------|
| P001 | A | 98 | 101 | 100 | 99.67 | TRUE (95-105) | TRUE (max odch: 1.67) | OK |
| P002 | B | 94 | 96 | 95 | 95.00 | TRUE (95-105) | TRUE (max odch: 1.00) | OK |
| P003 | A | 102 | 108 | 99 | 103.00 | TRUE (95-105) | FALSE (108-103=5, ale 99 odch: 4 → ok... 108-103=5 ≤5 ✓, 102-103=1 ✓, 99-103=4 ✓) → TRUE | OK |
| P004 | A | 100 | 100 | 101 | 100.33 | TRUE | TRUE (max odch: 0.67) | OK |
| P005 | B | 92 | 88 | 95 | 91.67 | FALSE (<95) | FALSE (88-91.67=3.67, ok, ale zakres juz falszuje) | ODRZUT |
| P006 | A | 103 | 104 | 105 | 104.00 | TRUE | TRUE (max odch: 1.00) | OK |
| P007 | B | 97 | 99 | 98 | 98.00 | TRUE | TRUE (max odch: 1.00) | OK |
| P008 | A | 110 | 112 | 108 | 110.00 | FALSE (>105) | TRUE (max odch: 2.00) | ODRZUT |
| P009 | B | 100 | 101 | 99 | 100.00 | TRUE | TRUE (max odch: 1.00) | OK |
| P010 | A | 96 | 97 | 98 | 97.00 | TRUE | TRUE (max odch: 1.00) | OK |

Korekta P003: srednia = (102+108+99)/3 = 309/3 = 103.00. Odchylenia: |102-103|=1 ≤5 ✓, |108-103|=5 ≤5 ✓, |99-103|=4 ≤5 ✓. Wiec H=TRUE, wynik=OK.

**Formula koncowa (lacza masa czesci typu A z wynikiem OK):**
```
=SUMIFS(F2:F11;B2:B11;"A";I2:I11;"OK")
```

Czesci typu A z wynikiem OK: P001 (99.67), P003 (103.00), P004 (100.33), P006 (104.00), P010 (97.00)
Suma: 99.67 + 103.00 + 100.33 + 104.00 + 97.00 = **504.00**

**Wyjasnienie**: Cwiczenie laczy kilka krokow: kolumny pomocnicze z AVERAGE i AND, a nastepnie SUMIFS z wieloma warunkami na kolumnach pomocniczych. Jest to typowy wzorzec na maturze — budowanie rozwiazania etapami.
</details>
