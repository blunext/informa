# 17. Wykresy w arkuszu kalkulacyjnym

Typ zadania: **arkusz_wykres**
Czestotliwosc: 8/11 lat | Laczna punktacja: 25 pkt
Kategoria: ARKUSZ

## Umiejetnosci cwiczone w tym zestawie

`wykres-kolumnowy` `wykres-kolowy` `wykres-liniowy` `wykres-grupowany` `wykres-skumulowany` `wykres-kombinowany` `dwie-osie-Y` `etykiety-danych` `legenda` `tytuly-osi` `odczyt-z-wykresu` `procent-udzialu` `formatowanie-wykresu` `seria-danych`

---

### Cwiczenie 17.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `wykres-kolumnowy` `etykiety-danych` `tytuly-osi`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Wykres kolumnowy (slupkowy pionowy) jest najlepszy dla danych kategorycznych z wartosciami liczbowymi.
2. **Podejscie**: Os X to kategorie (przedzialy), os Y to wartosci (liczba uczniow). Jedna seria danych.
3. **Kluczowy krok**: Pamietaj o: tytule wykresu, podpisach osi, etykietach danych, odpowiedniej skali osi Y.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Brak tytulu wykresu lub podpisow osi**: Na maturze za brakujace elementy opisowe odejmuje sie punkty. CKE: -1 pkt za kazdy brakujacy element
- **Wybor zlego typu wykresu**: Wykres kolowy dla rozkladu wynikow jest niewlasciwy — kolumnowy lepiej pokazuje porownanie miedzy przedzialami. CKE: -1 pkt

</details>

---

### Cwiczenie 17.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 (Cukier)
**Tagi**: `wykres-kolowy` `procent-udzialu` `legenda` `etykiety-danych`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Wykres kolowy najlepiej pokazuje udzial czesci w calosci. Kazdy wycinek = jedna kategoria.
2. **Podejscie**: Oblicz sume wszystkich wydatkow, potem udzial kazdej kategorii (kwota/suma × 100%).
3. **Kluczowy krok**: Etykiety na wykresie: nazwa kategorii + procent. Legenda po prawej stronie.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Brak etykiet procentowych**: Wykres kolowy bez procentow jest nieczytelny — nie wiadomo, jaki udzial ma kazda kategoria. CKE: -1 pkt
- **Zbyt wiele kategorii**: Wykres kolowy jest czytelny dla 3-7 kategorii. Dla wiekszej liczby lepszy jest kolumnowy. CKE: uwaga ogolna

</details>

---

### Cwiczenie 17.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2023 (Konfitury)
**Tagi**: `wykres-grupowany` `legenda` `seria-danych` `odczyt-z-wykresu`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Wykres grupowany — kazdy miesiac ma 3 slupki obok siebie (po jednym na produkt).
2. **Podejscie**: 3 serie danych = 3 kolumny (B, C, D). Os X = miesiace. Legenda rozroznia serie.
3. **Kluczowy krok**: Aby znalezc najwyzsza laczna sprzedaz, zsumuj B+C+D w kazdym wierszu.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Brak legendy**: Przy 3 seriach danych bez legendy nie wiadomo, ktory slupek odpowiada ktoremu produktowi. CKE: -1 pkt
- **Uzycie wykresu skumulowanego zamiast grupowanego**: Skumulowany uklada serie jedna na drugiej — nie pozwala porownac poszczegolnych produktow. CKE: -2 pkt jesli zadanie wymaga porownania

</details>

---

### Cwiczenie 17.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 (Martianeum)
**Tagi**: `wykres-skumulowany` `procent-udzialu` `seria-danych` `odczyt-z-wykresu`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Wykres skumulowany — serie ulozone jedna na drugiej. Wysokosc calego slupka = suma.
2. **Podejscie**: Os X = kwartaly, kazde zrodlo to osobna seria. Udzial OZE = OZE/razem × 100%.
3. **Kluczowy krok**: Uwaga — dane w tabeli sa ulozone "zrodla w wierszach", wiec na wykresie serie to wiersze, nie kolumny.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Pominiecie transpozycji danych**: Jesli dane sa w wierszach (jak tu), trzeba wskazac programowi, ze serie to wiersze, nie kolumny. CKE: -1 pkt
- **Bledne odczytanie wartosci ze skumulowanego**: Wartosc pojedynczej serii to roznica miedzy gorna a dolna krawedzia warstwy, NIE odczyt z osi Y. CKE: -1 pkt

</details>

---

### Cwiczenie 17.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2019 (Pogoda)
**Tagi**: `wykres-kombinowany` `dwie-osie-Y` `wykres-liniowy` `wykres-kolumnowy` `odczyt-z-wykresu`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Wykres kombinowany laczy dwa typy wykresow (liniowy + kolumnowy) z dwiema osiami Y.
2. **Podejscie**: Temperatura i opady maja rozne jednostki i skale — potrzebna os dodatkowa. Linie = temperatura, slupki = opady.
3. **Kluczowy krok**: W arkuszu: wstaw wykres, zmien serie temperatur na liniowe, przesuń opady na os prawa.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Brak drugiej osi Y**: Umieszczenie temperatury i opadow na jednej osi — skale sa rozne, wykres bedzie nieczytelny. CKE: -2 pkt
- **Wszystkie serie jako slupki**: Temperatura powinna byc linia (trend), opady slupkami (pojedyncze zdarzenia). CKE: -1 pkt
- **Bledne przypisanie serii do osi**: Opady na osi lewej, temperatura na prawej — odwrocona logika. CKE: -1 pkt

</details>

---

### Cwiczenie 17.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `wykres-liniowy` `etykiety-danych` `tytuly-osi` `formatowanie-wykresu`

Pomiary temperatury powietrza w ciagu dnia:

| | A | B |
|---|---|---|
| 1 | **Godzina** | **Temperatura (°C)** |
| 2 | 6:00 | 8 |
| 3 | 9:00 | 14 |
| 4 | 12:00 | 21 |
| 5 | 15:00 | 23 |
| 6 | 18:00 | 18 |
| 7 | 21:00 | 12 |

**Polecenie**: Opisz, jak utworzyc wykres liniowy przedstawiajacy zmiane temperatury w ciagu dnia. Podaj:
1. Typ wykresu i uzasadnienie wyboru
2. Dane dla osi X i Y
3. Elementy opisu wykresu
4. W jakiej godzinie temperatura byla najwyzsza?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wykres liniowy jest najlepszy dla danych cioglych w czasie — pokazuje trend.
2. **Podejscie**: Os X = godziny (czas), os Y = temperatura. Jedna seria danych z markerami.
3. **Kluczowy krok**: Odczytaj wartosc maksymalna z tabeli — nie musisz uzywac formuly, wystarczy analiza danych.

</details>

<details>
<summary>Odpowiedz</summary>

**Specyfikacja wykresu:**

1. **Typ**: Wykres liniowy z markerami
   - Uzasadnienie: dane przedstawiaja zmiane jednej wielkosci w czasie — linia najlepiej uwidacznia trend (wzrost rano, szczyt po poludniu, spadek wieczorem)
2. **Dane**:
   - Os X: A2:A7 (godziny)
   - Os Y: B2:B7 (temperatura)
3. **Elementy opisu**:
   - Tytul: "Temperatura powietrza w ciagu dnia"
   - Os X: "Godzina"
   - Os Y: "Temperatura (°C)"
   - Markery w kazdym punkcie pomiarowym
   - Etykiety danych przy markerach (8, 14, 21, 23, 18, 12)
4. Najwyzsza temperatura: **23°C o godzinie 15:00**

**Wyjasnienie**: Wykres liniowy jest preferowany dla danych czasowych, bo linia sugeruje ciagle zmiany miedzy pomiarami. Markery wskazuja dokladne punkty pomiarowe. Na maturze — zawsze uzasadniaj wybor typu wykresu.
</details>

<details>
<summary>Typowe bledy</summary>

- **Wykres kolumnowy zamiast liniowego**: Dla danych czasowych wykres kolumnowy sugeruje niezalezne kategorie, a nie ciagla zmiane. CKE: -1 pkt
- **Brak markerow**: Sama linia bez markerow moze utrudnic dokladny odczyt wartosci. CKE: zwykle akceptowane, ale mniej czytelne

</details>

---

### Cwiczenie 17.7 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2018 (Szkolna biblioteka)
**Tagi**: `wykres-kolumnowy` `wykres-kolowy` `procent-udzialu` `odczyt-z-wykresu`

Wyniki ankiety "Ulubiony sport" wsrod 200 uczniow:

| | A | B |
|---|---|---|
| 1 | **Sport** | **Liczba glosow** |
| 2 | Piłka nożna | 68 |
| 3 | Koszykówka | 42 |
| 4 | Siatkówka | 35 |
| 5 | Tenis | 25 |
| 6 | Pływanie | 18 |
| 7 | Inne | 12 |

**Polecenie**:
1. Opisz, jak utworzyc wykres kolowy z etykietami (nazwa + procent).
2. Oblicz udzial procentowy kazdego sportu.
3. Ktory typ wykresu lepiej nada sie do porownania liczby glosow: kolowy czy kolumnowy? Uzasadnij.
4. Jaki procent uczniow wybralo sport druzynowy (piłka nożna, koszykówka, siatkówka)?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wykres kolowy — struktura (udzialy), kolumnowy — porownanie wartosci bezwzglednych.
2. **Podejscie**: Udzial = glosy/200 × 100%. Sport druzynowy to suma trzech kategorii.
3. **Kluczowy krok**: Oba typy wykresow sa poprawne w roznych kontekstach — uzasadnij wybor na podstawie celu prezentacji.

</details>

<details>
<summary>Odpowiedz</summary>

**1. Specyfikacja wykresu kolowego:**
- Etykiety: A2:A7, Wartosci: B2:B7
- Etykiety na wycinkach: "Piłka nożna 34.0%", "Koszykówka 21.0%" itd.
- Legenda po prawej, tytul: "Ulubiony sport — wyniki ankiety"

**2. Udzialy procentowe:**

| Sport | Glosy | Udzial |
|-------|-------|--------|
| Piłka nożna | 68 | 34.0% |
| Koszykówka | 42 | 21.0% |
| Siatkówka | 35 | 17.5% |
| Tenis | 25 | 12.5% |
| Pływanie | 18 | 9.0% |
| Inne | 12 | 6.0% |

Weryfikacja: 34.0+21.0+17.5+12.5+9.0+6.0 = 100.0% ✓

**3. Porownanie typow wykresow:**
- **Kolowy** — lepszy do pokazania **struktury** (jaki procent calosci stanowi kazdy sport). Czytelny przy 6 kategoriach.
- **Kolumnowy** — lepszy do **porownania wartosci** (latwiej zobaczyc roznice miedzy slupkami niz miedzy wycinkami). Lepsza precyzja odczytu.
- **Rekomendacja**: Jesli celem jest porownanie popularnosci — **kolumnowy**. Jesli celem jest pokazanie udzialow — **kolowy**.

**4. Sport druzynowy:**
Piłka nożna + Koszykówka + Siatkówka = 68 + 42 + 35 = 145
Procent: 145/200 × 100 = **72.5%**

**Wyjasnienie**: Na maturze czesto pojawia sie pytanie o wybor typu wykresu. Klucz: kolowy = czesc z calosci, kolumnowy = porownanie, liniowy = trend w czasie.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak uzasadnienia wyboru typu**: Na maturze nie wystarczy napisac "kolumnowy" — trzeba wyjaśnic DLACZEGO. CKE: -1 pkt
- **Bledne obliczenie procentu**: 68/200 = 0.34, nie 0.34% — pamietaj o mnozeniu razy 100. CKE: -1 pkt
- **Nieuwzglednienie kategorii "Inne"**: Suma procentow musi = 100%. CKE: -1 pkt jesli brakuje

</details>

---

### Cwiczenie 17.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2022 (Wybory)
**Tagi**: `wykres-grupowany` `wykres-skumulowany` `seria-danych` `legenda` `odczyt-z-wykresu`

Wyniki egzaminu w 4 klasach:

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Klasa** | **Celujacy/Bdb** | **Dobry/Dst** | **Dop/Ndst** |
| 2 | 3A | 8 | 14 | 3 |
| 3 | 3B | 5 | 12 | 8 |
| 4 | 3C | 10 | 11 | 4 |
| 5 | 3D | 3 | 9 | 13 |

**Polecenie**:
1. Jaki typ wykresu (grupowany czy skumulowany) lepiej pokaze: (a) liczebnosc klas, (b) rozklad ocen w klasach? Uzasadnij.
2. Opisz specyfikacje wykresu skumulowanego procentowego (100% stacked) dla tych danych.
3. Oblicz rozmiar kazdej klasy i procentowy udzial kazdej grupy ocen.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Skumulowany pokazuje sume (liczebnosc klasy), grupowany pozwala porownac poszczegolne kategorie ocen. Skumulowany 100% normalizuje do procentow.
2. **Podejscie**: Dla wykresu 100% stacked — kazdy slupek ma ta sama wysokosc (100%), a proporcje wewnatrz pokazuja strukture.
3. **Kluczowy krok**: Rozmiar klasy = suma trzech grup. Procent = grupa/rozmiar × 100%.

</details>

<details>
<summary>Odpowiedz</summary>

**1. Wybor typu wykresu:**

(a) **Liczebnosc klas** — wykres **skumulowany** (stacked): calkowita wysokosc slupka = liczba uczniow w klasie. Latwiej porownac rozmiary klas.

(b) **Rozklad ocen** — wykres **skumulowany 100%** (100% stacked): kazdy slupek ma taka sama wysokosc (100%), wiec widac proporcje ocen niezaleznie od wielkosci klasy. Alternatywnie: **grupowany** pozwala dokladniej porownac wartosci bezwzgledne.

**2. Specyfikacja wykresu 100% stacked:**
- Typ: Kolumnowy skumulowany 100%
- Os X: A2:A5 (klasy)
- Serie: Celujacy/Bdb (zielony), Dobry/Dst (zolty), Dop/Ndst (czerwony)
- Kazdy slupek = 100%, podzielony na 3 czesci
- Tytul: "Rozklad ocen w klasach (%)"
- Legenda u gory

**3. Obliczenia:**

| Klasa | Cel/Bdb | Db/Dst | Dop/Ndst | Razem | %Cel/Bdb | %Db/Dst | %Dop/Ndst |
|-------|---------|--------|----------|-------|----------|---------|-----------|
| 3A | 8 | 14 | 3 | 25 | 32.0% | 56.0% | 12.0% |
| 3B | 5 | 12 | 8 | 25 | 20.0% | 48.0% | 32.0% |
| 3C | 10 | 11 | 4 | 25 | 40.0% | 44.0% | 16.0% |
| 3D | 3 | 9 | 13 | 25 | 12.0% | 36.0% | 52.0% |

Kazda klasa ma 25 uczniow. Klasa 3C ma najlepszy rozklad (40% cel/bdb), klasa 3D — najgorszy (52% dop/ndst).

**Wyjasnienie**: Wykres 100% stacked jest idealny do porownywania struktury miedzy grupami o roznych rozmiarach. Na maturze pojawia sie czesto przy analizie ankiet lub rozkladow.
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomieszanie skumulowanego z 100% skumulowanym**: Zwykly skumulowany pokazuje wartosci bezwzgledne, 100% — udzialy procentowe. CKE: -1 pkt
- **Brak uzasadnienia**: "Wybieram kolumnowy" bez wyjaśnienia dlaczego — na maturze wymagane jest uzasadnienie. CKE: -1 pkt

</details>

---

### Cwiczenie 17.9 (trudnosc: srednie-trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)
**Tagi**: `wykres-liniowy` `wykres-kombinowany` `dwie-osie-Y` `odczyt-z-wykresu` `formatowanie-wykresu`

Dane o ruchu na stronie internetowej i sprzedazy online w 8 tygodniach:

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Tydzien** | **Odwiedziny** | **Zamowienia** | **Konwersja (%)** |
| 2 | 1 | 5200 | 156 | 3.0 |
| 3 | 2 | 4800 | 168 | 3.5 |
| 4 | 3 | 6100 | 183 | 3.0 |
| 5 | 4 | 7500 | 300 | 4.0 |
| 6 | 5 | 6800 | 238 | 3.5 |
| 7 | 6 | 8200 | 410 | 5.0 |
| 8 | 7 | 7100 | 284 | 4.0 |
| 9 | 8 | 9000 | 450 | 5.0 |

**Polecenie**:
1. Opisz wykres kombinowany: linia odwiedzin (os lewa) + slupki zamowien (os prawa).
2. Opisz osobny wykres liniowy konwersji z linia trendu.
3. Czy wzrost odwiedzin zawsze przekladal sie na wzrost zamowien? Uzasadnij na podstawie danych.
4. W ktorym tygodniu konwersja byla najwyzsza?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dwa wykresy — kombinowany (odwiedziny + zamowienia) i liniowy (konwersja). Skale sa rozne, wiec potrzebne dwie osie Y.
2. **Podejscie**: Porownaj tygodnie, w ktorych odwiedziny rosna, a zamowienia spadaja (lub odwrotnie) — to dowod na brak idealnej korelacji.
3. **Kluczowy krok**: Linia trendu (regresja liniowa) na wykresie konwersji pokazuje ogolny kierunek zmian.

</details>

<details>
<summary>Odpowiedz</summary>

**1. Wykres kombinowany (odwiedziny + zamowienia):**
- Typ: Combo chart (linia + slupki)
- Os X: A2:A9 (tygodnie)
- Seria 1 (linia, os lewa): B2:B9 — Odwiedziny, kolor niebieski, skala 0-10000
- Seria 2 (slupki, os prawa): C2:C9 — Zamowienia, kolor zielony, skala 0-500
- Tytul: "Odwiedziny i zamowienia wg tygodni"
- Legenda, tytuly osi

**2. Wykres liniowy konwersji:**
- Typ: Liniowy z markerami + linia trendu (regresja liniowa)
- Os X: tygodnie, Os Y: D2:D9 (konwersja %)
- Tytul: "Wskaznik konwersji — trend"
- Linia trendu: rosnaca (konwersja rosnie z 3.0% do 5.0%)

**3. Analiza korelacji odwiedzin vs zamowien:**

| Tydzien | Odwiedziny | Zmiana | Zamowienia | Zmiana | Zgodnosc? |
|---------|------------|--------|------------|--------|-----------|
| 1→2 | 5200→4800 | spadek | 156→168 | wzrost | NIE |
| 2→3 | 4800→6100 | wzrost | 168→183 | wzrost | TAK |
| 3→4 | 6100→7500 | wzrost | 183→300 | wzrost | TAK |
| 4→5 | 7500→6800 | spadek | 300→238 | spadek | TAK |
| 5→6 | 6800→8200 | wzrost | 238→410 | wzrost | TAK |
| 6→7 | 8200→7100 | spadek | 410→284 | spadek | TAK |
| 7→8 | 7100→9000 | wzrost | 284→450 | wzrost | TAK |

Odpowiedz: **Nie zawsze** — w tygodniu 1→2 odwiedziny spadly, a zamowienia wzrosly. Ale w wiekszosci przypadkow (6 z 7) kierunki zmian sa zgodne.

**4. Najwyzsza konwersja**: **5.0% w tygodniach 6 i 8**

**Wyjasnienie**: Wykres kombinowany z dwiema osiami pozwala wizualnie porownac dwie wielkosc o roznych skalach. Linia trendu pomaga dostrzec ogolny kierunek zmian. Na maturze czesto trzeba przeanalizowac, czy dwie wielkosci sa skorelowane.
</details>

<details>
<summary>Typowe bledy</summary>

- **Obie serie na jednej osi**: Odwiedziny (tysiace) i zamowienia (setki) na jednej skali — zamowienia beda niewidoczne. CKE: -2 pkt
- **Stwierdzenie "zawsze rosna razem" bez analizy**: Trzeba sprawdzic kazda pare tygodni, nie zakladac korelacji. CKE: -1 pkt
- **Brak linii trendu**: Jesli polecenie wymaga linii trendu, trzeba ja dodac. CKE: -1 pkt

</details>

---

### Cwiczenie 17.10 (trudnosc: trudne, ~6 pkt)
**Zrodlo inspiracji**: Matura 2021 (Szyfry)
**Tagi**: `wykres-kolumnowy` `wykres-kolowy` `wykres-liniowy` `procent-udzialu` `odczyt-z-wykresu` `formatowanie-wykresu`

Zrodla dochodow gminy w 3 kolejnych latach (w tys. zl):

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Zrodlo** | **2022** | **2023** | **2024** |
| 2 | Podatki lokalne | 1200 | 1350 | 1500 |
| 3 | Dotacje rzadowe | 800 | 750 | 900 |
| 4 | Oplaty komunalne | 400 | 420 | 450 |
| 5 | Sprzedaz majatku | 150 | 80 | 200 |
| 6 | Inne | 100 | 120 | 130 |

**Polecenie**:
1. Dla kazdego roku oblicz laczny dochod i udzial procentowy kazdego zrodla.
2. Opisz 3 wykresy: (a) kolumnowy grupowany — porownanie zrodel w latach, (b) kolowy dla roku 2024, (c) liniowy — trend lacznych dochodow.
3. Ktore zrodlo dochodow roslo najszybciej (najwieksza zmiana procentowa 2022→2024)?
4. Czy struktura dochodow zmienila sie istotnie miedzy 2022 a 2024?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy rozne wykresy — kazdy odpowiada na inne pytanie. Grupowany = porownanie, kolowy = struktura, liniowy = trend.
2. **Podejscie**: Zmiana procentowa = (wartosc_2024 - wartosc_2022) / wartosc_2022 × 100%.
3. **Kluczowy krok**: Porownaj udzialy procentowe 2022 vs 2024, aby ocenic zmiane struktury.

</details>

<details>
<summary>Odpowiedz</summary>

**1. Obliczenia:**

| Zrodlo | 2022 | 2023 | 2024 |
|--------|------|------|------|
| Podatki lokalne | 1200 | 1350 | 1500 |
| Dotacje rzadowe | 800 | 750 | 900 |
| Oplaty komunalne | 400 | 420 | 450 |
| Sprzedaz majatku | 150 | 80 | 200 |
| Inne | 100 | 120 | 130 |
| **RAZEM** | **2650** | **2720** | **3180** |

Udzialy procentowe:

| Zrodlo | 2022 | 2024 |
|--------|------|------|
| Podatki lokalne | 45.3% | 47.2% |
| Dotacje rzadowe | 30.2% | 28.3% |
| Oplaty komunalne | 15.1% | 14.2% |
| Sprzedaz majatku | 5.7% | 6.3% |
| Inne | 3.8% | 4.1% |

**2. Specyfikacje wykresow:**

**(a) Kolumnowy grupowany:**
- Os X: zrodla dochodow (5 kategorii)
- Serie: 2022, 2023, 2024 (3 slupki obok siebie)
- Tytul: "Dochody gminy wg zrodel (tys. zl)"

**(b) Kolowy dla 2024:**
- Etykiety: A2:A6, Wartosci: D2:D6
- Etykiety: nazwa + procent
- Tytul: "Struktura dochodow gminy w 2024"

**(c) Liniowy — trend:**
- Os X: 2022, 2023, 2024
- Os Y: laczne dochody (2650, 2720, 3180)
- Linia z markerami + ewentualna linia trendu
- Tytul: "Trend lacznych dochodow gminy"

**3. Zmiana procentowa 2022→2024:**

| Zrodlo | 2022 | 2024 | Zmiana |
|--------|------|------|--------|
| Podatki lokalne | 1200 | 1500 | +25.0% |
| Dotacje rzadowe | 800 | 900 | +12.5% |
| Oplaty komunalne | 400 | 450 | +12.5% |
| Sprzedaz majatku | 150 | 200 | **+33.3%** |
| Inne | 100 | 130 | +30.0% |

Najszybszy wzrost: **Sprzedaz majatku (+33.3%)**

**4. Analiza zmian struktury:**
Struktura dochodow **nie zmienila sie istotnie** — udzialy procentowe sa zblizone (roznice <3 p.p.). Podatki lokalne dominuja w obu latach (~45-47%), dotacje rzadowe sa drugie (~28-30%). Glowny trend: lagodny wzrost udzialu podatkow kosztem dotacji.

**Wyjasnienie**: Na maturze czesto trzeba utworzyc kilka roznych wykresow z tych samych danych — kazdy odpowiada na inne pytanie. Zmiana procentowa to standardowa miara dynamiki. Porownanie struktur wymaga zestawienia udzialow procentowych, nie wartosci bezwzglednych.
</details>

<details>
<summary>Typowe bledy</summary>

- **Porownanie wartosci bezwzglednych zamiast procentowych**: "Podatki rosly najszybciej bo wzrosly o 300" — ale 300/1200=25%, podczas gdy sprzedaz majatku 50/150=33%. CKE: -2 pkt
- **Brak osobnego wykresu na kazde pytanie**: Jeden wykres nie odpowie na wszystkie pytania jednoczesnie. CKE: -1 pkt
- **Brak wniosku o strukturze**: "Nie zmienila sie" to za malo — trzeba podac DLACZEGO (porownac udzialy). CKE: -1 pkt

</details>

---

## Samoocena

| Poziom | Opis | Cwiczenia |
|--------|------|-----------|
| Podstawowy | Tworze proste wykresy kolumnowe, kolowe i liniowe z opisem | 17.1-17.2, 17.6 bez pomocy |
| Dobry | Tworze wykresy grupowane i skumulowane, obliczam udzialy procentowe | 17.3-17.4, 17.7-17.8 bez pomocy |
| Bardzo dobry | Tworze wykresy kombinowane z dwiema osiami Y, analizuje korelacje | 17.5, 17.9 bez pomocy |
| Doskonaly | Dobiam typ wykresu do pytania, analizuje zmiany struktury i dynamike | 17.10 bez pomocy |

**Co dalej?**
- Jesli masz trudnosci z podstawowymi: przejrzyj `cheatsheet_arkusz.md` (sekcja Wykresy)
- Jesli opanowales srednie: przejdz do cwiczen 15 (Agregacja warunkowa) — czesto laczone z wykresami na maturze
- Jesli zrobiles wszystkie 10: sprobuj cwiczen 16 (Symulacja) ze stoperem (maks. 15 min na wykres + pytania)
