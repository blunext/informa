# 17. Wykresy w arkuszu kalkulacyjnym

Typ zadania: **arkusz_wykres**
Czestotliwosc: 8/11 lat | Laczna punktacja: 25 pkt
Kategoria: ARKUSZ

---

### Cwiczenie 17.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna

Uczniowie klasy pisali test z matematyki. Wyniki pogrupowano w przedzialy:

| | A | B |
|---|---|---|
| 1 | **Przedzial punktow** | **Liczba uczniow** |
| 2 | 0-20% | 2 |
| 3 | 21-40% | 5 |
| 4 | 41-60% | 8 |
| 5 | 61-80% | 7 |
| 6 | 81-100% | 3 |

**Polecenie**: Opisz, jak utworzyc wykres kolumnowy (histogram) przedstawiajacy rozklad wynikow testu. Podaj:
1. Typ wykresu
2. Dane dla osi X i Y
3. Tytul wykresu i osi
4. Dodatkowe formatowanie

<details>
<summary>Odpowiedz</summary>

**Specyfikacja wykresu:**

1. **Typ**: Wykres kolumnowy (slupkowy pionowy)
2. **Dane**:
   - Os X (kategorie): zakres A2:A6 (przedzialy punktow)
   - Os Y (wartosci): zakres B2:B6 (liczba uczniow)
3. **Tytuly**:
   - Tytul wykresu: "Rozklad wynikow testu z matematyki"
   - Os X: "Przedzial punktow"
   - Os Y: "Liczba uczniow"
4. **Formatowanie**:
   - Kolumny w jednym kolorze (jedna seria danych)
   - Etykiety danych nad kazdym slupkiem (2, 5, 8, 7, 3)
   - Linie siatki poziomej dla latwiejszego odczytu
   - Os Y od 0 do 10 (z krokiem co 2)

**Wyjasnienie**: Wykres kolumnowy jest najwlasciwszy, bo os X zawiera kategorie (przedzialy), a os Y wartosci liczbowe. Kazdy slupek reprezentuje jedna kategorie. Na maturze nalezy zadbac o: czytelny tytul, podpisy osi, etykiety danych.
</details>

---

### Cwiczenie 17.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 (Cukier)

Rodzina prowadzi budzet domowy. Wydatki miesieczne wg kategorii:

| | A | B |
|---|---|---|
| 1 | **Kategoria** | **Kwota (zl)** |
| 2 | Zywnosc | 1800 |
| 3 | Mieszkanie | 2200 |
| 4 | Transport | 600 |
| 5 | Rozrywka | 400 |
| 6 | Oszczednosci | 1000 |

**Polecenie**: Opisz, jak utworzyc wykres kolowy z etykietami procentowymi pokazujacy udzial kazdej kategorii w budzecie. Oblicz udzialy procentowe.

<details>
<summary>Odpowiedz</summary>

**Obliczenia:**

Suma wydatkow: 1800 + 2200 + 600 + 400 + 1000 = 6000 zl

| Kategoria | Kwota | Udzial |
|-----------|-------|--------|
| Zywnosc | 1800 | 30.0% |
| Mieszkanie | 2200 | 36.7% |
| Transport | 600 | 10.0% |
| Rozrywka | 400 | 6.7% |
| Oszczednosci | 1000 | 16.7% |

**Specyfikacja wykresu:**

1. **Typ**: Wykres kolowy (pie chart)
2. **Dane**:
   - Etykiety (kategorie): zakres A2:A6
   - Wartosci: zakres B2:B6
3. **Tytul**: "Struktura wydatkow miesiecznych"
4. **Formatowanie**:
   - Etykiety na kazdym wycinku: nazwa kategorii + procent (np. "Zywnosc 30.0%")
   - Rozne kolory dla kazdej kategorii
   - Legenda po prawej stronie wykresu
   - Procenty zaokraglone do 1 miejsca po przecinku

**Wyjasnienie**: Wykres kolowy przedstawia czesci calosci — idealny do pokazania struktury wydatkow. Etykiety procentowe sa obliczane automatycznie przez arkusz (opcja "Pokaz wartosci jako procenty"). Suma wycinkow zawsze = 100%.
</details>

---

### Cwiczenie 17.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2023 (Konfitury)

Firma sprzedaje 3 produkty. Dane miesiecznej sprzedazy (w tys. zl):

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Miesiac** | **Produkt A** | **Produkt B** | **Produkt C** |
| 2 | Styczen | 45 | 30 | 20 |
| 3 | Luty | 38 | 35 | 25 |
| 4 | Marzec | 52 | 28 | 30 |
| 5 | Kwiecien | 41 | 42 | 22 |
| 6 | Maj | 55 | 38 | 35 |
| 7 | Czerwiec | 48 | 45 | 28 |

**Polecenie**: Opisz, jak utworzyc wykres kolumnowy grupowany (3 serie danych obok siebie) z legenda. Podaj pelna specyfikacje wykresu i odpowiedz na pytanie: w ktorym miesiacu laczna sprzedaz wszystkich produktow byla najwyzsza?

<details>
<summary>Odpowiedz</summary>

**Specyfikacja wykresu:**

1. **Typ**: Wykres kolumnowy grupowany (clustered column)
2. **Dane**:
   - Os X (kategorie): zakres A2:A7 (miesiace)
   - Seria 1: B2:B7 (Produkt A) — np. kolor niebieski
   - Seria 2: C2:C7 (Produkt B) — np. kolor pomaranczowy
   - Seria 3: D2:D7 (Produkt C) — np. kolor zielony
3. **Tytuly**:
   - Tytul: "Miesieczna sprzedaz produktow (tys. zl)"
   - Os X: "Miesiac"
   - Os Y: "Sprzedaz (tys. zl)"
4. **Formatowanie**:
   - Legenda u gory lub po prawej: Produkt A, Produkt B, Produkt C
   - 3 slupki obok siebie w kazdym miesiacu
   - Linie siatki poziomej
   - Os Y od 0 do 60

**Laczna sprzedaz:**

| Miesiac | A + B + C | Suma |
|---------|-----------|------|
| Styczen | 45+30+20 | 95 |
| Luty | 38+35+25 | 98 |
| Marzec | 52+28+30 | 110 |
| Kwiecien | 41+42+22 | 105 |
| Maj | 55+38+35 | **128** |
| Czerwiec | 48+45+28 | 121 |

Najwyzsza laczna sprzedaz: **Maj (128 tys. zl)**

**Wyjasnienie**: Wykres grupowany pozwala porownywac wartosci miedzy seriami w ramach jednej kategorii (np. ktory produkt sprzedal sie najlepiej w marcu) oraz miedzy kategoriami (np. trend sprzedazy Produktu A w czasie). Legenda jest niezbedna, gdy jest wiecej niz jedna seria danych.
</details>

---

### Cwiczenie 17.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 (Martianeum)

Zuzycie energii (w GWh) wg zrodel w 4 kwartalach:

| | A | B | C | D | E |
|---|---|---|---|---|---|
| 1 | **Zrodlo** | **Q1** | **Q2** | **Q3** | **Q4** |
| 2 | Wegiel | 120 | 100 | 80 | 130 |
| 3 | Gaz | 60 | 50 | 45 | 65 |
| 4 | OZE | 30 | 55 | 70 | 35 |
| 5 | Atom | 90 | 90 | 90 | 90 |

**Polecenie**:
1. Opisz, jak utworzyc wykres kolumnowy skumulowany (stacked) gdzie kazdy slupek pokazuje calkowite zuzycie energii w kwartale, podzielone na zrodla.
2. Oblicz calkowite zuzycie w kazdym kwartale.
3. W ktorym kwartale udzial OZE byl procentowo najwyzszy?

<details>
<summary>Odpowiedz</summary>

**Specyfikacja wykresu:**

1. **Typ**: Wykres kolumnowy skumulowany (stacked column)
2. **Dane** (uwaga — dane trzeba "transponowac" w ukladzie: kwartaly na osi X, zrodla jako serie):
   - Os X (kategorie): Q1, Q2, Q3, Q4
   - Seria 1 (Wegiel): B2:E2
   - Seria 2 (Gaz): B3:E3
   - Seria 3 (OZE): B4:E4
   - Seria 4 (Atom): B5:E5
3. **Tytuly**:
   - Tytul: "Zuzycie energii wg zrodel (GWh)"
   - Os X: "Kwartal"
   - Os Y: "Zuzycie (GWh)"
4. **Formatowanie**:
   - Kolory: Wegiel=czarny, Gaz=szary, OZE=zielony, Atom=zolty
   - Legenda po prawej stronie
   - Kazdy slupek sklada sie z 4 warstw (jedno zrodlo na warstwe)
   - Calkowita wysokosc slupka = suma wszystkich zrodel

**Obliczenia:**

| Kwartal | Wegiel | Gaz | OZE | Atom | Razem |
|---------|--------|-----|-----|------|-------|
| Q1 | 120 | 60 | 30 | 90 | 300 |
| Q2 | 100 | 50 | 55 | 90 | 295 |
| Q3 | 80 | 45 | 70 | 90 | 285 |
| Q4 | 130 | 65 | 35 | 90 | 320 |

**Udzial OZE:**

| Kwartal | OZE | Razem | Udzial OZE |
|---------|-----|-------|------------|
| Q1 | 30 | 300 | 10.0% |
| Q2 | 55 | 295 | 18.6% |
| Q3 | 70 | 285 | **24.6%** |
| Q4 | 35 | 320 | 10.9% |

Najwyzszy udzial OZE: **Q3 (24.6%)**

**Wyjasnienie**: Wykres skumulowany rozni sie od grupowanego tym, ze serie sa ulozone jedna na drugiej (nie obok siebie). Pozwala to jednoczesnie zobaczyc calkosc (wysokosc slupka) i sklad (czesci). Udzial procentowy = OZE / Razem × 100%. Na maturze trzeba umiec odczytac z wykresu skumulowanego zarowno wartosc calkowita, jak i poszczegolne czesci.
</details>

---

### Cwiczenie 17.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2019 (Pogoda)

Dane pogodowe z 10 dni:

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Dzien** | **Temp. max (°C)** | **Temp. min (°C)** | **Opady (mm)** |
| 2 | 1 | 22 | 12 | 0 |
| 3 | 2 | 25 | 15 | 0 |
| 4 | 3 | 20 | 14 | 8 |
| 5 | 4 | 18 | 11 | 15 |
| 6 | 5 | 16 | 9 | 22 |
| 7 | 6 | 19 | 10 | 5 |
| 8 | 7 | 23 | 13 | 0 |
| 9 | 8 | 27 | 16 | 0 |
| 10 | 9 | 24 | 14 | 3 |
| 11 | 10 | 21 | 12 | 12 |

**Polecenie**: Opisz, jak utworzyc wykres kombinowany (liniowo-kolumnowy) z dwiema osiami Y:
- Os lewa (Y1): temperatura — dwie linie (max i min)
- Os prawa (Y2): opady — slupki
Podaj pelna specyfikacje i odpowiedz: w ktorych dniach opady przekroczyly 10mm, a temperatura maksymalna spadla ponizej 20°C?

<details>
<summary>Odpowiedz</summary>

**Specyfikacja wykresu:**

1. **Typ**: Wykres kombinowany (combo chart) — liniowy + kolumnowy z dwiema osiami Y
2. **Dane**:
   - Os X: A2:A11 (numery dni)
   - Seria 1 (linia, os lewa): B2:B11 — Temp. max, kolor czerwony, linia ciagla z markerami
   - Seria 2 (linia, os lewa): C2:C11 — Temp. min, kolor niebieski, linia ciagla z markerami
   - Seria 3 (slupki, os prawa): D2:D11 — Opady, kolor jasnoniebieski, slupki
3. **Osie**:
   - Os lewa (Y1): "Temperatura (°C)", zakres np. 0-30
   - Os prawa (Y2): "Opady (mm)", zakres np. 0-25
   - Os X: "Dzien"
4. **Formatowanie**:
   - Tytul: "Temperatura i opady w ciagu 10 dni"
   - Legenda u gory: Temp. max, Temp. min, Opady
   - Linie siatki poziomej dla osi lewej
   - Slupki opadow sa "za" liniami temperatury (na drugim planie)

**Tworzenie w arkuszu (krok po kroku):**
1. Zaznacz dane A1:D11
2. Wstaw wykres kolumnowy
3. Zmien serie Temp. max i Temp. min na typ "liniowy"
4. Przesuń serie Opady na os dodatkowa (prawa)
5. Dodaj tytuly osi i legende

**Odpowiedz na pytanie:**

| Dzien | Temp. max | Opady | Opady>10 i Temp<20? |
|-------|-----------|-------|---------------------|
| 4 | 18 | 15 | TAK (18<20 i 15>10) |
| 5 | 16 | 22 | TAK (16<20 i 22>10) |
| 10 | 21 | 12 | NIE (21≥20) |

Dni spelniajace oba warunki: **dzien 4 i dzien 5**

**Wyjasnienie**: Wykres z dwiema osiami Y jest potrzebny, gdy serie maja rozne jednostki lub skale (°C vs mm). Bez osobnej osi slupki opadow (0-22) bylyby nieczytelne na skali temperatury (0-30). Na maturze nalezy precyzyjnie opisac: ktore serie sa na ktorej osi, jaki typ (linia/slupek), kolory i legende.
</details>
