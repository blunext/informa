# 15. Agregacja warunkowa w arkuszu kalkulacyjnym

Typ zadania: **arkusz_agregacja_warunkowa**
Czestotliwosc: 11/12 lat | Laczna punktacja: 38 pkt
Kategoria: ARKUSZ

## Umiejetnosci cwiczone w tym zestawie

`COUNTIF` `COUNTIFS` `SUMIF` `SUMIFS` `AVERAGEIF` `AVERAGEIFS` `warunek-tekstowy` `warunek-liczbowy` `operator-konkatenacji` `zakres-bezwzgledny` `kolumna-pomocnicza` `IF-AND-OR` `MAXIFS-MINIFS` `zagniezdzone-funkcje`

---

### Cwiczenie 15.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)
**Tagi**: `COUNTIF` `warunek-tekstowy`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Potrzebujesz funkcji zliczajacej komorki spelniajace warunek tekstowy.
2. **Podejscie**: Funkcja COUNTIF przyjmuje zakres i kryterium — przeszukaj kolumne kategorii.
3. **Kluczowy krok**: =COUNTIF(B2:B11;"Nabiał") — zakres to kolumna B, kryterium to dokladny tekst.

</details>

<details>
<summary>Odpowiedz</summary>

**Formula:**
```
=COUNTIF(B2:B11;"Nabiał")
```

**Weryfikacja**: Produkty z kategorii "Nabiał": Mleko, Jogurt, Ser, Masło, Kefir, Śmietana = **6**

**Wyjasnienie**: COUNTIF(zakres; kryterium) zlicza komorki w zakresie B2:B11, ktore zawieraja tekst "Nabiał". Srednik oddziela argumenty (konwencja polska/europejska).
</details>

<details>
<summary>Typowe bledy</summary>

- **Przecinek zamiast srednika**: =COUNTIF(B2:B11,"Nabiał") — w polskiej wersji Excela separatorem jest srednik. CKE: -1 pkt
- **Brak cudzyslowow wokol tekstu**: =COUNTIF(B2:B11;Nabiał) — kryterium tekstowe musi byc w cudzyslowach. CKE: -1 pkt

</details>

---

### Cwiczenie 15.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 (Demografia)
**Tagi**: `AVERAGEIF` `warunek-tekstowy`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Potrzebujesz funkcji obliczajacej srednia z warunkowym filtrowaniem — nie zwyklego AVERAGE.
2. **Podejscie**: AVERAGEIF ma 3 argumenty: zakres kryterium, kryterium, zakres sredniej.
3. **Kluczowy krok**: Zakres kryterium to kolumna B (plec), zakres sredniej to kolumna C (oceny).

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Zamiana kolejnosci argumentow**: =AVERAGEIF(C2:C11;"M";B2:B11) — pierwszy zakres to zakres kryterium (B), nie zakres sredniej (C). CKE: -1 pkt
- **Uzycie AVERAGE zamiast AVERAGEIF**: =AVERAGE(C2:C11) oblicza srednia WSZYSTKICH uczniow, nie tylko jednej plci. CKE: -2 pkt

</details>

---

### Cwiczenie 15.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)
**Tagi**: `SUMIFS` `warunek-tekstowy` `wiele-warunkow`

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
<summary>Wskazowki</summary>

1. **Kierunek**: SUMIFS pozwala sumowac z wieloma warunkami jednoczesnie (w przeciwienstwie do SUMIF z jednym warunkiem).
2. **Podejscie**: Pierwszy argument to zakres sumowania, potem pary: zakres kryterium + kryterium.
3. **Kluczowy krok**: =SUMIFS(D2:D13;B2:B13;"Północ";C2:C13;"Q3") — dwa warunki: region i kwartal.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Uzycie SUMIF zamiast SUMIFS**: SUMIF obsluguje tylko jeden warunek. Aby filtrowac po regionie I kwartale, trzeba SUMIFS. CKE: -2 pkt
- **Zly porzadek argumentow w SUMIFS**: W SUMIFS zakres sumy jest PIERWSZY (inaczej niz w SUMIF, gdzie jest trzeci). CKE: -1 pkt
- **Niezgodne rozmiary zakresow**: Wszystkie zakresy w SUMIFS musza miec te sama liczbe wierszy. CKE: -1 pkt

</details>

---

### Cwiczenie 15.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 (Pogoda)
**Tagi**: `AVERAGEIF` `COUNTIFS` `operator-konkatenacji` `warunek-liczbowy`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Punkt 1 to standardowy AVERAGEIF. Punkt 2 wymaga dwoch warunkow — COUNTIFS.
2. **Podejscie**: W punkcie 2 kryterium liczbowe musi byc polaczone z operatorem porownania za pomoca `&`.
3. **Kluczowy krok**: Kryterium `">"&F4` laczy tekst ">" z wartoscia komorki F4 (5.0), dajac ">5".

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Brak cudzyslowow wokol operatora**: =">"&F4 jest poprawne, ale >F4 bez cudzyslowow nie zadziala. CKE: -1 pkt
- **Uzycie COUNTIF zamiast COUNTIFS**: COUNTIF obsluguje tylko jeden warunek (np. tylko ">5"), ale nie filtruje jednoczesnie po miesiacu. CKE: -2 pkt
- **Wpisanie wartosci zamiast odwolania**: =COUNTIFS(B2:B11;"Marzec";C2:C11;">5") zadziala, ale traci elastycznosc — lepiej uzywac odwolania do komorki z norma. CKE: zwykle akceptowane, ale mniej eleganckie.

</details>

---

### Cwiczenie 15.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 (Martianeum)
**Tagi**: `SUMIFS` `kolumna-pomocnicza` `IF-AND-OR` `AVERAGEIF` `zagniezdzone-funkcje`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Zadanie wymaga budowania rozwiazania etapami — kazda kolumna pomocnicza odpowiada jednemu krokowi logicznemu.
2. **Podejscie**: AND laczy wiele warunkow logicznych. ABS oblicza wartosc bezwzgledna. IF zamienia TRUE/FALSE na teksty.
3. **Kluczowy krok**: Koncowa formula SUMIFS filtruje po kolumnach pomocniczych — typ="A" i wynik="OK".

</details>

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
| P003 | A | 102 | 108 | 99 | 103.00 | TRUE (95-105) | TRUE (max odch: 5.00) | OK |
| P004 | A | 100 | 100 | 101 | 100.33 | TRUE | TRUE (max odch: 0.67) | OK |
| P005 | B | 92 | 88 | 95 | 91.67 | FALSE (<95) | TRUE (max odch: 3.67) | ODRZUT |
| P006 | A | 103 | 104 | 105 | 104.00 | TRUE | TRUE (max odch: 1.00) | OK |
| P007 | B | 97 | 99 | 98 | 98.00 | TRUE | TRUE (max odch: 1.00) | OK |
| P008 | A | 110 | 112 | 108 | 110.00 | FALSE (>105) | TRUE (max odch: 2.00) | ODRZUT |
| P009 | B | 100 | 101 | 99 | 100.00 | TRUE | TRUE (max odch: 1.00) | OK |
| P010 | A | 96 | 97 | 98 | 97.00 | TRUE | TRUE (max odch: 1.00) | OK |

Korekta P003: srednia = (102+108+99)/3 = 309/3 = 103.00. Odchylenia: |102-103|=1 ≤5 ✓, |108-103|=5 ≤5 ✓, |99-103|=4 ≤5 ✓. Wiec H=TRUE, wynik=OK.

**Formula koncowa (laczna masa czesci typu A z wynikiem OK):**
```
=SUMIFS(F2:F11;B2:B11;"A";I2:I11;"OK")
```

Czesci typu A z wynikiem OK: P001 (99.67), P003 (103.00), P004 (100.33), P006 (104.00), P010 (97.00)
Suma: 99.67 + 103.00 + 100.33 + 104.00 + 97.00 = **504.00**

**Wyjasnienie**: Cwiczenie laczy kilka krokow: kolumny pomocnicze z AVERAGE i AND, a nastepnie SUMIFS z wieloma warunkami na kolumnach pomocniczych. Jest to typowy wzorzec na maturze — budowanie rozwiazania etapami.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak ABS w warunku odchylenia**: Bez ABS roznica moze byc ujemna i warunek ≤5 bedzie falszywiy dla pomiarow nizszych od sredniej. CKE: -1 pkt
- **Uzycie OR zamiast AND w kolumnie I**: Wynik "OK" wymaga spelnienia OBU warunkow (zakres I odchylenie), nie jednego z nich. CKE: -2 pkt
- **Zapomnienie o kolumnach pomocniczych**: Proba napisania jednej ogromnej formuly SUMIFS bez etapow posrednich — trudne do debugowania i czesto bledne. CKE: akceptowane, ale ryzykowne.

</details>

---

### Cwiczenie 15.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `SUMIF` `warunek-tekstowy`

Sklep prowadzi ewidencje sprzedazy owocow w ciagu tygodnia:

| | A | B | C |
|---|---|---|---|
| 1 | **Dzien** | **Owoc** | **Sprzedaz (kg)** |
| 2 | Pon | Jabłka | 25 |
| 3 | Pon | Banany | 15 |
| 4 | Wt | Jabłka | 30 |
| 5 | Wt | Banany | 20 |
| 6 | Sr | Jabłka | 18 |
| 7 | Sr | Banany | 22 |
| 8 | Czw | Jabłka | 35 |
| 9 | Czw | Banany | 12 |
| 10 | Pt | Jabłka | 40 |
| 11 | Pt | Banany | 28 |

**Polecenie**: Napisz formuly obliczajace:
1. Laczna sprzedaz jablek w calym tygodniu.
2. Laczna sprzedaz bananow w calym tygodniu.
3. Ktory owoc mial wieksza laczna sprzedaz?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Potrzebujesz funkcji sumujacej z warunkiem — SUMIF.
2. **Podejscie**: SUMIF(zakres_kryterium; kryterium; zakres_sumy) — filtruj po kolumnie B, sumuj kolumne C.
3. **Kluczowy krok**: Porownaj wyniki obu SUMIF, aby odpowiedziec na punkt 3.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
Jabłka: =SUMIF(B2:B11;"Jabłka";C2:C11)
Banany: =SUMIF(B2:B11;"Banany";C2:C11)
```

**Weryfikacja:**
- Jabłka: 25 + 30 + 18 + 35 + 40 = **148 kg**
- Banany: 15 + 20 + 22 + 12 + 28 = **97 kg**
- Wieksza sprzedaz: **Jabłka (148 kg)**

**Wyjasnienie**: SUMIF(zakres_kryterium; kryterium; zakres_sumy) sumuje wartosci z kolumny C tylko dla wierszy, gdzie kolumna B zawiera podany tekst. Roznica SUMIF vs SUMIFS: SUMIF obsluguje jeden warunek, SUMIFS — wiele.
</details>

<details>
<summary>Typowe bledy</summary>

- **Zamiana zakresow**: =SUMIF(C2:C11;"Jabłka";B2:B11) — zakres kryterium (B) i zakres sumy (C) sa zamienione. CKE: -1 pkt
- **Uzycie COUNTIF zamiast SUMIF**: COUNTIF zlicza ILOSC wierszy, a nie sumuje wartosci sprzedazy. CKE: -2 pkt

</details>

---

### Cwiczenie 15.7 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2017 (Cukier)
**Tagi**: `COUNTIFS` `warunek-liczbowy` `operator-konkatenacji` `wiele-warunkow`

Nauczyciel ma wyniki egzaminu z 12 uczniow:

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Imie** | **Klasa** | **Punkty** | **Plec** |
| 2 | Adam | 3A | 72 | M |
| 3 | Beata | 3B | 85 | K |
| 4 | Czarek | 3A | 58 | M |
| 5 | Diana | 3B | 91 | K |
| 6 | Emil | 3A | 45 | M |
| 7 | Fiona | 3B | 78 | K |
| 8 | Grzegorz | 3A | 83 | M |
| 9 | Hanna | 3B | 67 | K |
| 10 | Igor | 3A | 39 | M |
| 11 | Justyna | 3B | 74 | K |
| 12 | Kamil | 3A | 90 | M |
| 13 | Laura | 3B | 55 | K |

**Polecenie**:
1. Ile dziewczyn z klasy 3B uzyskalo wiecej niz 70 punktow?
2. Ile osob z klasy 3A nie zdalo (ponizej 50 punktow)?
3. Jaki procent uczniow klasy 3A zdalo egzamin (>=50 punktow)?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Punkt 1 i 2 wymagaja COUNTIFS z wieloma warunkami. Punkt 3 to iloraz dwoch COUNTIFS.
2. **Podejscie**: Kazdy warunek to osobna para (zakres; kryterium) w COUNTIFS. Warunki liczbowe wymagaja operatora w cudzyslowach.
3. **Kluczowy krok**: Procent = zdali / wszyscy z 3A × 100. Zdali = COUNTIFS z warunkiem ">=50".

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
1. =COUNTIFS(B2:B13;"3B";D2:D13;"K";C2:C13;">70")
2. =COUNTIFS(B2:B13;"3A";C2:C13;"<50")
3. =COUNTIFS(B2:B13;"3A";C2:C13;">=50")/COUNTIF(B2:B13;"3A")*100
```

**Weryfikacja:**

1. Dziewczyny z 3B z wynikiem >70:
   - Beata (85 ✓), Diana (91 ✓), Fiona (78 ✓), Hanna (67 ✗), Justyna (74 ✓), Laura (55 ✗)
   - Wynik: **4**

2. Osoby z 3A ponizej 50:
   - Adam (72 ✗), Czarek (58 ✗), Emil (45 ✓), Grzegorz (83 ✗), Igor (39 ✓), Kamil (90 ✗)
   - Wynik: **2**

3. Procent zdanych z 3A:
   - Zdali (≥50): Adam (72), Czarek (58), Grzegorz (83), Kamil (90) = 4 osoby
   - Wszyscy z 3A: 6 osob
   - Procent: 4/6 × 100 = **66.67%**

**Wyjasnienie**: COUNTIFS obsluguje do 127 par warunek-zakres. Warunki liczbowe zapisujemy w cudzyslowach: ">70", "<50", ">=50". Procent to iloraz — licznik (COUNTIFS ze warunkami) / mianownik (COUNTIF z jednym warunkiem).
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak cudzyslowow wokol operatora liczbowego**: =COUNTIFS(B2:B13;"3B";C2:C13;>70) — operator musi byc w cudzyslowach ">70". CKE: -1 pkt
- **Pominiecie warunku plci w punkcie 1**: Bez warunku D2:D13;"K" policzymy wszystkich z 3B, nie tylko dziewczyny. CKE: -2 pkt
- **Pomieszanie >= z >**: ">=50" to zdali (50 tez zdal), ">50" wyklucza osoby z dokladnie 50 punktami. CKE: -1 pkt jesli zmienia wynik

</details>

---

### Cwiczenie 15.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2023 (Konfitury)
**Tagi**: `SUMIFS` `AVERAGEIFS` `warunek-liczbowy` `wiele-warunkow`

Hurtownia prowadzi rejestr dostaw towaru:

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Dostawca** | **Kategoria** | **Waga (kg)** | **Ocena jakosci** |
| 2 | Firma A | Owoce | 500 | 8 |
| 3 | Firma B | Warzywa | 300 | 6 |
| 4 | Firma A | Warzywa | 450 | 9 |
| 5 | Firma C | Owoce | 200 | 5 |
| 6 | Firma B | Owoce | 350 | 7 |
| 7 | Firma A | Owoce | 600 | 8 |
| 8 | Firma C | Warzywa | 250 | 4 |
| 9 | Firma B | Warzywa | 400 | 7 |
| 10 | Firma A | Warzywa | 550 | 9 |
| 11 | Firma C | Owoce | 180 | 6 |

**Polecenie**:
1. Oblicz laczna wage dostaw owocow od Firmy A.
2. Oblicz srednia ocene jakosci dostaw warzyw o wadze powyzej 300 kg.
3. Oblicz laczna wage dostaw o ocenie jakosci >= 7.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Punkt 1 to SUMIFS z dwoma warunkami tekstowymi. Punkt 2 to AVERAGEIFS. Punkt 3 to SUMIFS z warunkiem liczbowym.
2. **Podejscie**: W AVERAGEIFS — pierwszy argument to zakres sredniej (D), potem pary warunek. W warunku liczbowym: ">300".
3. **Kluczowy krok**: Pamietaj o kolejnosci argumentow: SUMIFS(suma; zakr1; kryt1; ...) i AVERAGEIFS(srednia; zakr1; kryt1; ...).

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
1. =SUMIFS(C2:C11;A2:A11;"Firma A";B2:B11;"Owoce")
2. =AVERAGEIFS(D2:D11;B2:B11;"Warzywa";C2:C11;">300")
3. =SUMIFS(C2:C11;D2:D11;">="&7)
   lub: =SUMIFS(C2:C11;D2:D11;">=7")
```

**Weryfikacja:**

1. Owoce od Firmy A:
   - Wiersz 2: Firma A, Owoce, 500 ✓
   - Wiersz 7: Firma A, Owoce, 600 ✓
   - Suma: 500 + 600 = **1100 kg**

2. Warzywa o wadze >300 — srednia ocena:
   - Wiersz 4: Warzywa, 450>300, ocena=9 ✓
   - Wiersz 9: Warzywa, 400>300, ocena=7 ✓
   - Wiersz 10: Warzywa, 550>300, ocena=9 ✓
   - (Wiersz 3: Warzywa, 300 — nie spelnia >300)
   - (Wiersz 8: Warzywa, 250 — nie spelnia >300)
   - Srednia: (9+7+9)/3 = **8.33**

3. Dostawy z ocena >=7:
   - Wiersz 2: ocena 8 ✓, waga 500
   - Wiersz 4: ocena 9 ✓, waga 450
   - Wiersz 6: ocena 7 ✓, waga 350
   - Wiersz 7: ocena 8 ✓, waga 600
   - Wiersz 9: ocena 7 ✓, waga 400
   - Wiersz 10: ocena 9 ✓, waga 550
   - Suma: 500+450+350+600+400+550 = **2850 kg**

**Wyjasnienie**: AVERAGEIFS ma taka sama skladnie jak SUMIFS — pierwszy argument to zakres sredniej/sumy, potem pary warunek. Warunki moga byc tekstowe ("Owoce") lub liczbowe (">=7"). Mozna mieszac warunki tekstowe i liczbowe w jednej formule.
</details>

<details>
<summary>Typowe bledy</summary>

- **Pominiecie warunku wagi w punkcie 2**: Bez ">300" policzylibysmy srednia dla WSZYSTKICH warzyw. CKE: -2 pkt
- **Uzycie > zamiast >= w punkcie 3**: ">7" wyklucza dostawy z ocena dokladnie 7 (Firma B Owoce i Firma B Warzywa). CKE: -1 pkt
- **Pomieszanie AVERAGEIF z AVERAGEIFS**: AVERAGEIF obsluguje 1 warunek, AVERAGEIFS — wiele. Dla 2 warunkow trzeba AVERAGEIFS. CKE: -1 pkt

</details>

---

### Cwiczenie 15.9 (trudnosc: srednie-trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2021 (Szyfry)
**Tagi**: `COUNTIFS` `SUMIFS` `kolumna-pomocnicza` `IF-AND-OR` `zakres-bezwzgledny`

W arkuszu mamy dane pracownikow firmy:

| | A | B | C | D | E |
|---|---|---|---|---|---|
| 1 | **Pracownik** | **Dzial** | **Staz (lata)** | **Pensja** | **Wyksztalcenie** |
| 2 | Kowalski | IT | 8 | 9500 | Wyzsze |
| 3 | Nowak | HR | 3 | 5200 | Srednie |
| 4 | Wisniewski | IT | 12 | 11000 | Wyzsze |
| 5 | Wojcik | Sprzedaz | 5 | 6800 | Wyzsze |
| 6 | Kaminska | HR | 7 | 6100 | Wyzsze |
| 7 | Lewandowski | IT | 2 | 7200 | Wyzsze |
| 8 | Zielinska | Sprzedaz | 10 | 8500 | Srednie |
| 9 | Szymanski | IT | 15 | 13000 | Wyzsze |
| 10 | Dabrowska | HR | 4 | 5800 | Srednie |
| 11 | Mazur | Sprzedaz | 6 | 7000 | Wyzsze |

**Polecenie**:
1. W kolumnie F wpisz "Senior" jesli staz >= 10, "Mid" jesli staz >= 5, "Junior" jesli staz < 5.
2. Oblicz srednia pensje seniorow z dzialu IT.
3. Ile osob z wyzszym wyksztalceniem zarabia powyzej sredniej firmy?
4. Oblicz laczna pensje juniorow (staz < 5).

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Punkt 1 wymaga zagniezdzanego IF. Punkt 3 wymaga odwolania do sredniej firmy jako progu — uzyj komorki pomocniczej.
2. **Podejscie**: Zagniezdzone IF: =IF(C2>=10;"Senior";IF(C2>=5;"Mid";"Junior")). COUNTIFS w punkcie 3 z operatorem `">"&komorka_sredniej`.
3. **Kluczowy krok**: Srednia firmy w komorce pomocniczej (np. H1), potem COUNTIFS z odwolaniem do niej.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
F2: =IF(C2>=10;"Senior";IF(C2>=5;"Mid";"Junior"))

Srednia pensja seniorow IT:
=AVERAGEIFS(D2:D11;F2:F11;"Senior";B2:B11;"IT")

Srednia firmy (w H1): =AVERAGE(D2:D11)

Osoby wyzsze powyzej sredniej:
=COUNTIFS(E2:E11;"Wyzsze";D2:D11;">"&H1)

Laczna pensja juniorow:
=SUMIF(F2:F11;"Junior";D2:D11)
```

**Weryfikacja:**

Kolumna F (poziom stazu):
| Pracownik | Staz | Poziom |
|-----------|------|--------|
| Kowalski | 8 | Mid |
| Nowak | 3 | Junior |
| Wisniewski | 12 | Senior |
| Wojcik | 5 | Mid |
| Kaminska | 7 | Mid |
| Lewandowski | 2 | Junior |
| Zielinska | 10 | Senior |
| Szymanski | 15 | Senior |
| Dabrowska | 4 | Junior |
| Mazur | 6 | Mid |

Srednia pensja seniorow IT:
- Seniorzy z IT: Wisniewski (11000), Szymanski (13000)
- Srednia: (11000+13000)/2 = **12000**

Srednia firmy: (9500+5200+11000+6800+6100+7200+8500+13000+5800+7000)/10 = 80100/10 = **8010**

Osoby z wyzszym wyksztalceniem zarabiajace >8010:
- Kowalski (9500, Wyzsze) ✓
- Wisniewski (11000, Wyzsze) ✓
- Szymanski (13000, Wyzsze) ✓
- (Wojcik 6800 ✗, Kaminska 6100 ✗, Lewandowski 7200 ✗, Mazur 7000 ✗)
- Wynik: **3**

Laczna pensja juniorow:
- Nowak (5200) + Lewandowski (7200) + Dabrowska (5800) = **18200**

**Wyjasnienie**: Zagniezdzone IF tworzy warunki kaskadowe — sprawdza od najwyzszego progu w dol. AVERAGEIFS filtruje po dwoch kolumnach jednoczesnie (F i B). Komorka pomocnicza ze srednia (H1) pozwala uzyc jej dynamicznie w COUNTIFS.
</details>

<details>
<summary>Typowe bledy</summary>

- **Zla kolejnosc warunkow w IF**: =IF(C2>=5;"Mid";IF(C2>=10;"Senior";"Junior")) — warunek >=5 jest sprawdzany pierwszy i „pochłania" tez seniorow. Sprawdzaj od NAJWYZSZEGO progu. CKE: -2 pkt
- **Brak komorki pomocniczej ze srednia**: Wpisanie =COUNTIFS(E2:E11;"Wyzsze";D2:D11;">"&AVERAGE(D2:D11)) tez zadziala, ale jest mniej czytelne i trudniejsze do debugowania.
- **Zapomnienie o $ w odwolaniu do H1**: Jesli formula jest kopiowana, $H$1 zabezpiecza odwolanie. CKE: -1 pkt jesli kopiowanie jest wymagane.

</details>

---

### Cwiczenie 15.10 (trudnosc: trudne, ~6 pkt)
**Zrodlo inspiracji**: Matura 2022 (Wybory)
**Tagi**: `SUMIFS` `COUNTIFS` `MAXIFS-MINIFS` `kolumna-pomocnicza` `zagniezdzone-funkcje` `zakres-bezwzgledny`

Siec sklepow prowadzi miesieczna analize sprzedazy. Dane za styczen:

| | A | B | C | D | E |
|---|---|---|---|---|---|
| 1 | **Sklep** | **Miasto** | **Kategoria** | **Przychod** | **Koszty** |
| 2 | S1 | Warszawa | Elektronika | 85000 | 62000 |
| 3 | S2 | Kraków | Odzież | 45000 | 31000 |
| 4 | S3 | Warszawa | Odzież | 52000 | 38000 |
| 5 | S4 | Gdańsk | Elektronika | 71000 | 55000 |
| 6 | S5 | Kraków | Elektronika | 63000 | 48000 |
| 7 | S6 | Warszawa | Spożywcze | 38000 | 28000 |
| 8 | S7 | Gdańsk | Odzież | 41000 | 30000 |
| 9 | S8 | Kraków | Spożywcze | 35000 | 26000 |
| 10 | S9 | Warszawa | Elektronika | 92000 | 68000 |
| 11 | S10 | Gdańsk | Spożywcze | 29000 | 22000 |

**Polecenie**:
1. W kolumnie F oblicz zysk kazdego sklepu (przychod - koszty).
2. W kolumnie G oblicz marze procentowa (zysk/przychod × 100).
3. Znajdz najwyzszy zysk wsrod sklepow w Warszawie (uzyj MAXIFS).
4. Oblicz srednia marze sklepow z kategorii "Elektronika".
5. Ile sklepow ma marze powyzej 20%?
6. Oblicz laczny zysk sklepow, ktore maja przychod powyzej 50000 i marze powyzej 15%.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Zbuduj kolumny pomocnicze F (zysk) i G (marza), a potem uzyj ich w agregacjach warunkowych.
2. **Podejscie**: MAXIFS(zakres_max; zakres_kryt; kryterium) dziala jak SUMIFS, ale zwraca maximum. Punkt 6 wymaga SUMIFS z wieloma warunkami na kolumnach pomocniczych.
3. **Kluczowy krok**: Marza = F2/D2*100. MAXIFS i AVERAGEIFS operuja na kolumnach F i G.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
F2: =D2-E2                         (zysk, kopiowana w dol)
G2: =F2/D2*100                     (marza %, kopiowana w dol)

Najwyzszy zysk w Warszawie:
=MAXIFS(F2:F11;B2:B11;"Warszawa")

Srednia marza Elektroniki:
=AVERAGEIFS(G2:G11;C2:C11;"Elektronika")

Sklepy z marza >20%:
=COUNTIF(G2:G11;">20")

Laczny zysk (przychod>50000 i marza>15%):
=SUMIFS(F2:F11;D2:D11;">50000";G2:G11;">15")
```

**Obliczenia:**

| Sklep | Miasto | Kat. | Przychod | Koszty | F: Zysk | G: Marza (%) |
|-------|--------|------|----------|--------|---------|-------------|
| S1 | Warszawa | Elek. | 85000 | 62000 | 23000 | 27.06 |
| S2 | Kraków | Odzież | 45000 | 31000 | 14000 | 31.11 |
| S3 | Warszawa | Odzież | 52000 | 38000 | 14000 | 26.92 |
| S4 | Gdańsk | Elek. | 71000 | 55000 | 16000 | 22.54 |
| S5 | Kraków | Elek. | 63000 | 48000 | 15000 | 23.81 |
| S6 | Warszawa | Spoż. | 38000 | 28000 | 10000 | 26.32 |
| S7 | Gdańsk | Odzież | 41000 | 30000 | 11000 | 26.83 |
| S8 | Kraków | Spoż. | 35000 | 26000 | 9000 | 25.71 |
| S9 | Warszawa | Elek. | 92000 | 68000 | 24000 | 26.09 |
| S10 | Gdańsk | Spoż. | 29000 | 22000 | 7000 | 24.14 |

3. Najwyzszy zysk w Warszawie: max(23000, 14000, 10000, 24000) = **24000** (S9)

4. Srednia marza Elektroniki: (27.06+22.54+23.81+26.09)/4 = 99.50/4 = **24.87%**

5. Sklepy z marza >20%: Wszystkie 10 sklepow maja marze >20%, wiec wynik = **10**

6. Przychod>50000 i marza>15%:
   - S1: 85000>50000 ✓, 27.06>15 ✓, zysk=23000
   - S3: 52000>50000 ✓, 26.92>15 ✓, zysk=14000
   - S4: 71000>50000 ✓, 22.54>15 ✓, zysk=16000
   - S5: 63000>50000 ✓, 23.81>15 ✓, zysk=15000
   - S9: 92000>50000 ✓, 26.09>15 ✓, zysk=24000
   - Suma: 23000+14000+16000+15000+24000 = **92000**

**Wyjasnienie**: MAXIFS i MINIFS to nowsze funkcje (od Excel 2019 / LibreCalc 7.0). Dzialaja analogicznie do SUMIFS — pierwszy argument to zakres, potem pary warunek. Kolumny pomocnicze (zysk, marza) upraszczaja formuly koncowe. Na maturze to czesty wzorzec: oblicz wartosci posrednie → uzyj ich w agregacji warunkowej.
</details>

<details>
<summary>Typowe bledy</summary>

- **Marza jako zysk/koszty**: Marza procentowa to zysk/PRZYCHOD × 100, nie zysk/koszty. CKE: -2 pkt
- **Brak MAXIFS — uzycie MAX z IF tablicowym**: =MAX(IF(B2:B11="Warszawa";F2:F11)) wymaga Ctrl+Shift+Enter (formula tablicowa). MAXIFS jest prostsza. CKE: akceptowane oba podejscia.
- **Zapomnienie o kolumnie marzy w punkcie 6**: Bez kolumny G nie mozna filtrowac po marzy w SUMIFS. CKE: -2 pkt

</details>

---

## Samoocena

| Poziom | Opis | Cwiczenia |
|--------|------|-----------|
| Podstawowy | Znam COUNTIF, SUMIF, AVERAGEIF z jednym warunkiem | 15.1-15.2, 15.6 bez pomocy |
| Dobry | Stosuje COUNTIFS, SUMIFS z wieloma warunkami i operatorami liczbowymi | 15.3-15.4, 15.7-15.8 bez pomocy |
| Bardzo dobry | Tworze kolumny pomocnicze i lacze zagniezdzone IF z agregacjami | 15.5, 15.9 bez pomocy |
| Doskonaly | Uzywam MAXIFS/MINIFS, wieloetapowych analiz i formul z odwolaniami dynamicznymi | 15.10 bez pomocy |

**Co dalej?**
- Jesli masz trudnosci z podstawowymi: przejrzyj `cheatsheet_arkusz.md` (sekcja COUNTIF/SUMIF/AVERAGEIF)
- Jesli opanowales srednie: przejdz do cwiczen 16 (Symulacja) i 17 (Wykresy)
- Jesli zrobiles wszystkie 10: sprobuj cwiczen 18 (Agregacja podstawowa) dla uzupelnienia lub wrocic do 15.5/15.10 ze stoperem (maks. 10 min)
