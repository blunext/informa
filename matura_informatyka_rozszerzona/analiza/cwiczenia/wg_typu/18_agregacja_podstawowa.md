# 18. Agregacja podstawowa w arkuszu kalkulacyjnym

Typ zadania: **arkusz_agregacja_podstawowa**
Czestotliwosc: 3/12 lat | Laczna punktacja: 9 pkt
Kategoria: ARKUSZ

## Umiejetnosci cwiczone w tym zestawie

`SUM` `AVERAGE` `MAX` `MIN` `COUNT` `COUNTA` `COUNTBLANK` `RANK` `mnozenie-komorek` `zaokraglanie` `ROUND` `MEDIAN` `procent` `rozstep` `ABS` `odchylenie-od-sredniej`

---

### Cwiczenie 18.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `AVERAGE` `MAX` `MIN` `rozstep`

Uczniowie pisali test z 20 pytan. Wyniki (liczba poprawnych odpowiedzi):

| | A | B |
|---|---|---|
| 1 | **Uczen** | **Wynik** |
| 2 | Anna | 16 |
| 3 | Bartek | 12 |
| 4 | Celina | 19 |
| 5 | Dawid | 8 |
| 6 | Ewa | 14 |
| 7 | Filip | 11 |
| 8 | Gosia | 17 |
| 9 | Hubert | 13 |

**Polecenie**: Napisz formuly obliczajace:
1. Srednia wynikow
2. Najwyzszy wynik
3. Najnizszy wynik
4. Rozstep (roznica miedzy max a min)

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Potrzebujesz czterech podstawowych funkcji agregujacych: AVERAGE, MAX, MIN.
2. **Podejscie**: Wszystkie operuja na tym samym zakresie B2:B9.
3. **Kluczowy krok**: Rozstep nie ma dedykowanej funkcji — oblicz jako MAX(...) - MIN(...).

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
Srednia:     =AVERAGE(B2:B9)
Najwyzszy:   =MAX(B2:B9)
Najnizszy:   =MIN(B2:B9)
Rozstep:     =MAX(B2:B9)-MIN(B2:B9)
```

**Weryfikacja:**
- Wartosci: 16, 12, 19, 8, 14, 11, 17, 13
- Srednia: (16+12+19+8+14+11+17+13)/8 = 110/8 = **13.75**
- Max: **19** (Celina)
- Min: **8** (Dawid)
- Rozstep: 19 - 8 = **11**

**Wyjasnienie**: AVERAGE, MAX i MIN to podstawowe funkcje agregujace. Przyjmuja zakres komorek i zwracaja odpowiednio srednia arytmetyczna, wartosc najwieksza i najmniejsza. Rozstep nie ma dedykowanej funkcji — obliczamy go jako roznice MAX - MIN.
</details>

<details>
<summary>Typowe bledy</summary>

- **Zly zakres**: =AVERAGE(B1:B9) wlacza naglowek "Wynik" — spowoduje blad lub zly wynik. CKE: -1 pkt
- **Uzycie SUM zamiast AVERAGE**: =SUM(B2:B9)/8 tez da poprawna srednia, ale jest mniej eleganckie i latwiej o blad (jesli zmieni sie liczba uczniow). CKE: akceptowane

</details>

---

### Cwiczenie 18.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2014 (Przychody)
**Tagi**: `mnozenie-komorek` `SUM`

Sklep internetowy realizuje zamowienie:

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Produkt** | **Cena jedn.** | **Ilosc** | **Wartosc** |
| 2 | Zeszyt | 3.50 | 10 | ? |
| 3 | Dlugopis | 2.80 | 5 | ? |
| 4 | Linijka | 4.20 | 3 | ? |
| 5 | Gumka | 1.50 | 8 | ? |
| 6 | Olowek | 1.90 | 6 | ? |
| 7 | | | **RAZEM:** | ? |

**Polecenie**:
1. Napisz formule w D2 obliczajaca wartosc pozycji (cena × ilosc).
2. Napisz formule w D7 obliczajaca laczna wartosc zamowienia.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wartosc pozycji = cena jednostkowa × ilosc. Laczna wartosc = suma wszystkich pozycji.
2. **Podejscie**: D2 to mnozenie dwoch komorek. D7 to SUM zakresu D2:D6.
3. **Kluczowy krok**: Formula D2 jest kopiowana w dol do D6 — dzieki adresowaniu wzglednemu automatycznie dostosuje numery wierszy.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
D2: =B2*C2        (kopiowana w dol do D6)
D7: =SUM(D2:D6)
```

**Wypelniona tabela:**

| Produkt | Cena | Ilosc | Wartosc |
|---------|------|-------|---------|
| Zeszyt | 3.50 | 10 | 35.00 |
| Dlugopis | 2.80 | 5 | 14.00 |
| Linijka | 4.20 | 3 | 12.60 |
| Gumka | 1.50 | 8 | 12.00 |
| Olowek | 1.90 | 6 | 11.40 |
| | | **RAZEM:** | **85.00** |

**Weryfikacja:**
- D2: 3.50 × 10 = 35.00
- D3: 2.80 × 5 = 14.00
- D4: 4.20 × 3 = 12.60
- D5: 1.50 × 8 = 12.00
- D6: 1.90 × 6 = 11.40
- D7: 35.00 + 14.00 + 12.60 + 12.00 + 11.40 = 85.00

**Wyjasnienie**: Mnozenie komorek (B2*C2) to najprostsza operacja w arkuszu. SUM sumuje caly zakres. Na maturze czesto trzeba najpierw obliczyc wartosci posrednie (kolumna D), a dopiero potem je zagregowac.
</details>

<details>
<summary>Typowe bledy</summary>

- **Wpisanie wartosci zamiast formuly**: D2: =3.50*10 zamiast =B2*C2 — strata elastycznosci, formula nie zaktualizuje sie przy zmianie danych. CKE: -1 pkt
- **SUM na zlym zakresie**: =SUM(D2:D7) wlacza komorke RAZEM do sumy (odwolanie cykliczne). CKE: -1 pkt

</details>

---

### Cwiczenie 18.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `COUNTA` `COUNTBLANK` `COUNT` `procent`

Przeprowadzono ankiete wsrod 10 osob. Kazda osoba mogla udzielic odpowiedzi lub zostawic puste pole:

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Osoba** | **Pyt. 1** | **Pyt. 2** | **Pyt. 3** |
| 2 | O1 | Tak | Nie | Tak |
| 3 | O2 | Tak | | Nie |
| 4 | O3 | | Tak | Tak |
| 5 | O4 | Nie | Nie | |
| 6 | O5 | Tak | Tak | Tak |
| 7 | O6 | | | Nie |
| 8 | O7 | Tak | Nie | Tak |
| 9 | O8 | Nie | Tak | |
| 10 | O9 | Tak | | Tak |
| 11 | O10 | Tak | Tak | Nie |

**Polecenie**: Dla pytania 1 (kolumna B) napisz formuly obliczajace:
1. Ile osob odpowiedzialo (komorki niepuste)?
2. Ile osob nie odpowiedzialo (komorki puste)?
3. Procent wypelnienia (ile odpowiedzialo / ile osob ogolnie × 100%).
4. Ile osob odpowiedzialo "Tak" na pytanie 1?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: COUNTA zlicza niepuste komorki, COUNTBLANK — puste. Roznica miedzy COUNT (tylko liczby) a COUNTA (tekst i liczby) jest kluczowa.
2. **Podejscie**: Procent = niepuste / calkowita liczba osob × 100. Zliczenie "Tak" wymaga COUNTIF.
3. **Kluczowy krok**: Calkowita liczba osob mozna uzyskac z COUNTA kolumny A (imiona) lub wpisac recznie (10).

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
Odpowiedzialo:         =COUNTA(B2:B11)
Nie odpowiedzialo:     =COUNTBLANK(B2:B11)
Procent wypelnienia:   =COUNTA(B2:B11)/COUNTA(A2:A11)*100
   lub prosciej:       =COUNTA(B2:B11)/10*100
Odpowiedzi "Tak":     =COUNTIF(B2:B11;"Tak")
```

**Weryfikacja (kolumna B — Pytanie 1):**

Wartosci w B2:B11: Tak, Tak, (puste), Nie, Tak, (puste), Tak, Nie, Tak, Tak
- Niepuste (COUNTA): Tak, Tak, Nie, Tak, Tak, Nie, Tak, Tak = **8**
- Puste (COUNTBLANK): O3, O6 = **2**
- Procent: 8/10 × 100 = **80%**
- "Tak": O1, O2, O5, O7, O9, O10 = **6**

**Wyjasnienie**:
- `COUNTA` zlicza komorki niepuste (zarowno tekst, jak i liczby).
- `COUNTBLANK` zlicza komorki puste.
- `COUNTIF` zlicza komorki spelniajace podane kryterium.
- Roznica: COUNT zlicza tylko komorki z liczbami, COUNTA zlicza wszystkie niepuste (wlacznie z tekstem).
</details>

<details>
<summary>Typowe bledy</summary>

- **COUNT zamiast COUNTA**: COUNT zlicza tylko LICZBY — dla komorek z tekstem ("Tak", "Nie") zwroci 0. CKE: -2 pkt
- **Pomieszanie COUNTBLANK z COUNTA**: COUNTBLANK zlicza puste, COUNTA — niepuste. Sa komplementarne. CKE: -1 pkt
- **Brak mnozenia razy 100 w procencie**: 8/10 = 0.8, a nie 80%. CKE: -1 pkt (chyba ze komorka jest sformatowana jako procent)

</details>

---

### Cwiczenie 18.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `SUM` `RANK` `zakres-bezwzgledny`

W zawodach sportowych 6 zawodnikow rywalizowalo w 4 konkurencjach. Punkty za kazda konkurencje:

| | A | B | C | D | E | F | G |
|---|---|---|---|---|---|---|---|
| 1 | **Zawodnik** | **K1** | **K2** | **K3** | **K4** | **Suma** | **Ranking** |
| 2 | Adam | 8 | 6 | 9 | 7 | ? | ? |
| 3 | Beata | 9 | 8 | 7 | 9 | ? | ? |
| 4 | Czarek | 6 | 9 | 8 | 5 | ? | ? |
| 5 | Diana | 7 | 7 | 10 | 8 | ? | ? |
| 6 | Emil | 10 | 5 | 6 | 10 | ? | ? |
| 7 | Fiona | 5 | 10 | 9 | 6 | ? | ? |

**Polecenie**:
1. Napisz formule w F2 obliczajaca sume punktow zawodnika.
2. Napisz formule w G2 obliczajaca pozycje w rankingu (1 = najwyzszy wynik).
3. Podaj najwyzsza sume i srednia sume.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: SUM sumuje wiersz (B2:E2). RANK zwraca pozycje wartosci w rankingu.
2. **Podejscie**: RANK(wartosc; zakres; kolejnosc) — zakres musi byc bezwzgledny ($F$2:$F$7), bo formula bedzie kopiowana.
3. **Kluczowy krok**: Trzeci argument RANK: 0 = malejaco (najwieksza = pozycja 1).

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly (wiersz 2, kopiowane w dol):**
```
F2: =SUM(B2:E2)
G2: =RANK(F2;$F$2:$F$7;0)
```

**Wypelniona tabela:**

| Zawodnik | K1 | K2 | K3 | K4 | Suma | Ranking |
|----------|----|----|----|-----|------|---------|
| Adam | 8 | 6 | 9 | 7 | 30 | 4 |
| Beata | 9 | 8 | 7 | 9 | 33 | 1 |
| Czarek | 6 | 9 | 8 | 5 | 28 | 6 |
| Diana | 7 | 7 | 10 | 8 | 32 | 2 |
| Emil | 10 | 5 | 6 | 10 | 31 | 3 |
| Fiona | 5 | 10 | 9 | 6 | 30 | 4 |

**Weryfikacja:**
- Adam: 8+6+9+7 = 30
- Beata: 9+8+7+9 = 33
- Czarek: 6+9+8+5 = 28
- Diana: 7+7+10+8 = 32
- Emil: 10+5+6+10 = 31
- Fiona: 5+10+9+6 = 30

Ranking (malejaco): 33 (1.), 32 (2.), 31 (3.), 30 (4.), 30 (4.), 28 (6.)
- Uwaga: Adam i Fiona maja taki sam wynik (30), wiec obie maja ranking 4. Czarek z 28 ma ranking 6 (nie 5!). Tak dziala RANK — przy remisach pomija kolejne pozycje.

- Najwyzsza suma: =MAX(F2:F7) = **33** (Beata)
- Srednia suma: =AVERAGE(F2:F7) = (30+33+28+32+31+30)/6 = 184/6 = **30.67**

**Wyjasnienie**: RANK(wartosc; zakres; kolejnosc) zwraca pozycje wartosci w rankingu. Trzeci argument: 0 = malejaco (najwieksza = 1), 1 = rosnaco. Zakres `$F$2:$F$7` musi byc bezwzgledny (z $), aby po skopiowaniu formuly zawsze odwolywal sie do calej kolumny sum.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak $ w zakresie RANK**: =RANK(F2;F2:F7;0) — po skopiowaniu do wiersza 3 zakres przesunie sie na F3:F8, dajac bledne wyniki. CKE: -2 pkt
- **Trzeci argument = 1 zamiast 0**: RANK(...;1) daje ranking rosnacy (najnizsza wartosc = 1). CKE: -1 pkt
- **Oczekiwanie rang 1-6 bez przerw**: Przy remisach RANK pomija pozycje (4, 4, 6 zamiast 4, 5, 6). To jest poprawne zachowanie.

</details>

---

### Cwiczenie 18.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `AVERAGE` `ABS` `odchylenie-od-sredniej` `AVERAGEIF` `zakres-bezwzgledny`

Firma ma 10 pracownikow w 3 dzialach. Dane o pensjach:

| | A | B | C |
|---|---|---|---|
| 1 | **Pracownik** | **Dzial** | **Pensja** |
| 2 | P1 | IT | 8500 |
| 3 | P2 | HR | 6200 |
| 4 | P3 | IT | 9200 |
| 5 | P4 | Sprzedaz | 7100 |
| 6 | P5 | HR | 5800 |
| 7 | P6 | IT | 10500 |
| 8 | P7 | Sprzedaz | 6800 |
| 9 | P8 | HR | 6500 |
| 10 | P9 | Sprzedaz | 7500 |
| 11 | P10 | IT | 8800 |

**Polecenie**:
1. Oblicz laczna pensje wszystkich pracownikow.
2. Oblicz srednia pensje w calej firmie.
3. W kolumnie D oblicz odchylenie kazdego pracownika od sredniej firmy (pensja - srednia).
4. W kolumnie E oblicz wartosc bezwzgledna odchylenia (ABS).
5. Oblicz srednia pensje osobno dla kazdego dzialu (IT, HR, Sprzedaz).

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Oblicz srednia w komorce pomocniczej, a potem uzyj jej z adresem bezwzglednym w kolumnie D.
2. **Podejscie**: ABS usuwa znak minus — dzieki temu mozna uśrednić odchylenia (bo suma odchylen od sredniej = 0).
3. **Kluczowy krok**: AVERAGEIF oblicza srednia warunkowo — uzyj go do srednich na dzial.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
Laczna pensja:    =SUM(C2:C11)
Srednia firmy:    =AVERAGE(C2:C11)     (wynik w komorce G1)

D2: =C2-$G$1                           (kopiowana w dol)
E2: =ABS(D2)                           (kopiowana w dol)

Srednia IT:       =AVERAGEIF(B2:B11;"IT";C2:C11)
Srednia HR:       =AVERAGEIF(B2:B11;"HR";C2:C11)
Srednia Sprzedaz: =AVERAGEIF(B2:B11;"Sprzedaz";C2:C11)
```

**Obliczenia:**

Laczna pensja: 8500+6200+9200+7100+5800+10500+6800+6500+7500+8800 = **76900**
Srednia firmy: 76900/10 = **7690.00**

| Pracownik | Dzial | Pensja | Odchylenie (D) | |Odchylenie| (E) |
|-----------|-------|--------|----------------|-------------------|
| P1 | IT | 8500 | +810.00 | 810.00 |
| P2 | HR | 6200 | -1490.00 | 1490.00 |
| P3 | IT | 9200 | +1510.00 | 1510.00 |
| P4 | Sprzedaz | 7100 | -590.00 | 590.00 |
| P5 | HR | 5800 | -1890.00 | 1890.00 |
| P6 | IT | 10500 | +2810.00 | 2810.00 |
| P7 | Sprzedaz | 6800 | -890.00 | 890.00 |
| P8 | HR | 6500 | -1190.00 | 1190.00 |
| P9 | Sprzedaz | 7500 | -190.00 | 190.00 |
| P10 | IT | 8800 | +1110.00 | 1110.00 |

**Srednie na dzial:**
- IT: (8500+9200+10500+8800)/4 = 37000/4 = **9250.00**
- HR: (6200+5800+6500)/3 = 18500/3 = **6166.67**
- Sprzedaz: (7100+6800+7500)/3 = 21400/3 = **7133.33**

**Weryfikacja kontrolna:**
- Suma odchylen: 810+(-1490)+1510+(-590)+(-1890)+2810+(-890)+(-1190)+(-190)+1110 = 0 ✓ (suma odchylen od sredniej zawsze = 0)
- Srednia |odchylen|: (810+1490+1510+590+1890+2810+890+1190+190+1110)/10 = 12480/10 = 1248.00

**Wyjasnienie**: Odchylenie od sredniej to miara rozproszenia — pokazuje, jak daleko dana wartosc jest od sredniej. ABS (wartosc bezwzgledna) usuwa znak, dzieki czemu mozna uśrednić odchylenia (inaczej suma odchylen = 0). AVERAGEIF laczy agregacje ze srednia warunkowa — przydatne do porownywania dzialow. Odwolanie `$G$1` jest bezwzgledne, bo srednia jest jedna dla calej firmy.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak $ w odwolaniu do sredniej**: =C2-G1 po skopiowaniu do wiersza 3 stanie sie =C3-G2, co jest bledem. CKE: -2 pkt
- **Zapomnienie o ABS**: Bez ABS srednia odchylen wyniesie 0, co jest bezuzyteczne. CKE: -1 pkt
- **Pomieszanie AVERAGE z AVERAGEIF**: AVERAGE(C2:C11) to srednia wszystkich, AVERAGEIF filtruje po dziale. CKE: -2 pkt

</details>

---

### Cwiczenie 18.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `SUM` `AVERAGE` `procent`

Wyniki sprzedazy w 5 oddzialach firmy:

| | A | B |
|---|---|---|
| 1 | **Oddzial** | **Sprzedaz (tys. zl)** |
| 2 | Warszawa | 450 |
| 3 | Kraków | 320 |
| 4 | Wrocław | 280 |
| 5 | Gdańsk | 190 |
| 6 | Poznań | 260 |

**Polecenie**:
1. Oblicz laczna sprzedaz wszystkich oddzialow.
2. Oblicz srednia sprzedaz na oddzial.
3. W kolumnie C oblicz udzial procentowy kazdego oddzialu w calkowitej sprzedazy.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: SUM i AVERAGE na kolumnie B. Udzial = sprzedaz oddzialu / suma × 100%.
2. **Podejscie**: W kolumnie C formula odwoluje sie do komorki z suma (bezwzglednie z $).
3. **Kluczowy krok**: Sprawdz, czy suma udzialow = 100%.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
Suma (np. B8):    =SUM(B2:B6)
Srednia (np. B9): =AVERAGE(B2:B6)
C2:              =B2/$B$8*100      (kopiowana w dol do C6)
```

**Obliczenia:**
- Suma: 450+320+280+190+260 = **1500 tys. zl**
- Srednia: 1500/5 = **300 tys. zl**

| Oddzial | Sprzedaz | Udzial (%) |
|---------|----------|------------|
| Warszawa | 450 | 30.0% |
| Kraków | 320 | 21.3% |
| Wrocław | 280 | 18.7% |
| Gdańsk | 190 | 12.7% |
| Poznań | 260 | 17.3% |

Weryfikacja: 30.0+21.3+18.7+12.7+17.3 = 100.0% ✓

**Wyjasnienie**: Udzial procentowy = czesc/calosc × 100. Odwolanie $B$8 jest bezwzgledne, aby po skopiowaniu formuly zawsze odwolywac sie do komorki z suma calkowita.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak $ w odwolaniu do sumy**: =B2/B8*100 po skopiowaniu do C3 stanie sie =B3/B9*100 (B9 jest pusta!). CKE: -2 pkt
- **Brak mnozenia razy 100**: =B2/$B$8 da 0.30 zamiast 30.0%. CKE: -1 pkt (chyba ze format komorki to %)

</details>

---

### Cwiczenie 18.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2016 (Lotniska)
**Tagi**: `ROUND` `zaokraglanie` `AVERAGE` `MAX` `MIN`

Pomiary temperatury w 7 dniach (z dokladnoscia do 0.1°C):

| | A | B | C |
|---|---|---|---|
| 1 | **Dzien** | **Temp. rano** | **Temp. popoludniu** |
| 2 | Pon | 5.37 | 14.82 |
| 3 | Wt | 3.91 | 12.45 |
| 4 | Sr | 6.14 | 16.73 |
| 5 | Czw | 4.58 | 13.29 |
| 6 | Pt | 7.23 | 18.51 |
| 7 | Sob | 8.06 | 20.14 |
| 8 | Ndz | 6.72 | 17.38 |

**Polecenie**:
1. W kolumnie D oblicz srednia temperature dnia (srednia z rana i popoludnia).
2. W kolumnie E zaokraglij srednia do 1 miejsca po przecinku (ROUND).
3. Oblicz najwyzsza i najnizsza srednia temperature w tygodniu.
4. Oblicz srednia tygodniowa temperature (ze srednich dziennych zaokraglonych).

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: AVERAGE z dwoch komorek da srednia dnia. ROUND(wartosc; liczba_miejsc) zaokragla.
2. **Podejscie**: Kolejnosc: D2=AVERAGE, E2=ROUND(D2;1), potem MAX/MIN/AVERAGE na kolumnie E.
3. **Kluczowy krok**: ROUND(wartosc; 1) zaokragla do 1 miejsca po przecinku. ROUND(wartosc; 0) do liczby calkowitej.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
D2: =AVERAGE(B2:C2)          (kopiowana w dol)
E2: =ROUND(D2;1)             (kopiowana w dol)

Najwyzsza srednia: =MAX(E2:E8)
Najnizsza srednia: =MIN(E2:E8)
Srednia tygodniowa: =AVERAGE(E2:E8)
```

**Obliczenia:**

| Dzien | Rano | Popolud. | D: Srednia | E: Zaokr. |
|-------|------|----------|-----------|-----------|
| Pon | 5.37 | 14.82 | 10.095 | 10.1 |
| Wt | 3.91 | 12.45 | 8.180 | 8.2 |
| Sr | 6.14 | 16.73 | 11.435 | 11.4 |
| Czw | 4.58 | 13.29 | 8.935 | 8.9 |
| Pt | 7.23 | 18.51 | 12.870 | 12.9 |
| Sob | 8.06 | 20.14 | 14.100 | 14.1 |
| Ndz | 6.72 | 17.38 | 12.050 | 12.1 |

- Najwyzsza: **14.1** (Sob)
- Najnizsza: **8.2** (Wt)
- Srednia tygodniowa: (10.1+8.2+11.4+8.9+12.9+14.1+12.1)/7 = 77.7/7 = **11.1**

**Wyjasnienie**: ROUND(wartosc; n) zaokragla do n miejsc po przecinku. Dla n=0 zaokragla do calkowitej, n=1 do jednego miejsca itd. Na maturze czesto trzeba zaokraglic wyniki posrednie przed dalszymi obliczeniami.
</details>

<details>
<summary>Typowe bledy</summary>

- **INT zamiast ROUND**: INT(10.095) = 10 (ucinanie), ROUND(10.095;1) = 10.1 (zaokraglanie). CKE: -1 pkt
- **Zaokraglenie do zlej liczby miejsc**: ROUND(D2;2) daje 10.10 zamiast 10.1. CKE: -1 pkt
- **Obliczanie sredniej tygodniowej z niezaokraglonych wartosci**: Jesli polecenie mowi "ze srednich zaokraglonych", nalezy uzywac kolumny E, nie D. CKE: -1 pkt

</details>

---

### Cwiczenie 18.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `MEDIAN` `AVERAGE` `MAX` `MIN` `COUNT` `procent`

Wynagrodzenia 12 pracownikow firmy (w tys. zl):

| | A | B |
|---|---|---|
| 1 | **Pracownik** | **Wynagrodzenie** |
| 2 | P1 | 4.2 |
| 3 | P2 | 5.1 |
| 4 | P3 | 4.8 |
| 5 | P4 | 12.5 |
| 6 | P5 | 5.3 |
| 7 | P6 | 4.5 |
| 8 | P7 | 5.0 |
| 9 | P8 | 4.9 |
| 10 | P9 | 5.5 |
| 11 | P10 | 4.7 |
| 12 | P11 | 5.2 |
| 13 | P12 | 4.6 |

**Polecenie**:
1. Oblicz srednia, mediane, minimum i maximum wynagrodzen.
2. Ktora miara (srednia czy mediana) lepiej opisuje "typowe" wynagrodzenie? Uzasadnij.
3. Ile osob zarabia powyzej sredniej? Jaki to procent?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: MEDIAN zwraca wartosc srodkowa — jest odporniejsza na wartosci odstajace niz AVERAGE.
2. **Podejscie**: Porownaj srednia z mediana. Jesli roznica jest duza, w danych sa wartosci odstajace (tu: P4 = 12.5).
3. **Kluczowy krok**: Uzyj COUNTIF z kryterium ">"&srednia, aby policzyc osoby powyzej sredniej.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
Srednia:   =AVERAGE(B2:B13)
Mediana:   =MEDIAN(B2:B13)
Minimum:   =MIN(B2:B13)
Maximum:   =MAX(B2:B13)

Powyzej sredniej: =COUNTIF(B2:B13;">"&AVERAGE(B2:B13))
Procent:          =COUNTIF(B2:B13;">"&AVERAGE(B2:B13))/COUNT(B2:B13)*100
```

**Obliczenia:**
- Suma: 4.2+5.1+4.8+12.5+5.3+4.5+5.0+4.9+5.5+4.7+5.2+4.6 = 66.3
- Srednia: 66.3/12 = **5.525 tys. zl**
- Posortowane: 4.2, 4.5, 4.6, 4.7, 4.8, 4.9, 5.0, 5.1, 5.2, 5.3, 5.5, 12.5
- Mediana (srednia z 6. i 7. wartosci): (4.9+5.0)/2 = **4.95 tys. zl**
- Min: **4.2** (P1)
- Max: **12.5** (P4)

Powyzej sredniej (>5.525): P4 (12.5) i P9 (5.5) — nie, 5.5 < 5.525... Sprawdzmy:
- P4: 12.5 > 5.525 ✓
- P9: 5.5 < 5.525 ✗
- Wiec tylko **1 osoba** (P4)
- Procent: 1/12 × 100 = **8.33%**

**2. Ktora miara lepsza?**
**Mediana** (4.95) lepiej opisuje typowe wynagrodzenie. Srednia (5.525) jest zawyzona przez jedna wartosc odstajaca (P4 = 12.5 tys. zl). Tylko 1 z 12 osob zarabia powyzej sredniej, co potwierdza, ze srednia jest niereprezentatywna.

**Wyjasnienie**: Mediana jest odporniejsza na wartosci skrajne (outliers). Jesli rozklad jest skosniony (jak tu — jedna bardzo wysoka pensja), mediana lepiej reprezentuje "typowa" wartosc. Na maturze czesto pytaja o interpretacje roznic miedzy srednia a mediana.
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomieszanie MEDIAN z AVERAGE**: Mediana to wartosc srodkowa posortowanego zbioru, nie srednia arytmetyczna. CKE: -2 pkt
- **Bledne obliczenie mediany dla parzystej liczby elementow**: Dla 12 elementow mediana = srednia z 6. i 7. elementu (po posortowaniu). CKE: -1 pkt
- **Brak uzasadnienia**: "Mediana jest lepsza" bez wyjaśnienia dlaczego. CKE: -1 pkt

</details>

---

### Cwiczenie 18.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 (Pogoda)
**Tagi**: `AVERAGE` `MAX` `MIN` `mnozenie-komorek` `ROUND` `procent`

Sklep sprzedaje produkty w 3 kategoriach. Dane za tydzien:

| | A | B | C | D | E |
|---|---|---|---|---|---|
| 1 | **Dzien** | **Elektronika** | **Odzież** | **Żywność** | **Razem** |
| 2 | Pon | 2500 | 1800 | 3200 | ? |
| 3 | Wt | 1900 | 2100 | 2800 | ? |
| 4 | Sr | 3100 | 1500 | 3500 | ? |
| 5 | Czw | 2200 | 1900 | 2600 | ? |
| 6 | Pt | 4500 | 2800 | 4100 | ? |
| 7 | Sob | 5200 | 3500 | 4800 | ? |
| 8 | Ndz | 800 | 900 | 1200 | ? |

**Polecenie**:
1. W kolumnie E oblicz laczna sprzedaz dnia (suma trzech kategorii).
2. Oblicz srednia, najwyzsza i najnizsza dzienna sprzedaz (z kolumny E).
3. W kolumnie F oblicz udzial elektroniki w sprzedazy dnia (B/E × 100), zaokraglony do 1 miejsca.
4. W ktorym dniu udzial elektroniki byl najwyzszy?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: E = SUM trzech kolumn, F = B/E × 100 zaokraglone ROUND.
2. **Podejscie**: Oblicz E, potem F korzysta z E. Na koniec MAX/MIN/AVERAGE na E, oraz analiza F.
3. **Kluczowy krok**: Udzial = ROUND(B2/E2*100;1). Sprawdz, ktory dzien ma najwyzszy F.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
E2: =SUM(B2:D2)                    (kopiowana w dol)
F2: =ROUND(B2/E2*100;1)            (kopiowana w dol)

Srednia dzienna: =AVERAGE(E2:E8)
Najwyzsza:       =MAX(E2:E8)
Najnizsza:       =MIN(E2:E8)
```

**Obliczenia:**

| Dzien | Elek. | Odzież | Żywn. | E: Razem | F: % Elek. |
|-------|-------|--------|-------|----------|------------|
| Pon | 2500 | 1800 | 3200 | 7500 | 33.3 |
| Wt | 1900 | 2100 | 2800 | 6800 | 27.9 |
| Sr | 3100 | 1500 | 3500 | 8100 | 38.3 |
| Czw | 2200 | 1900 | 2600 | 6700 | 32.8 |
| Pt | 4500 | 2800 | 4100 | 11400 | 39.5 |
| Sob | 5200 | 3500 | 4800 | 13500 | 38.5 |
| Ndz | 800 | 900 | 1200 | 2900 | 27.6 |

- Srednia: (7500+6800+8100+6700+11400+13500+2900)/7 = 56900/7 = **8128.57**
- Najwyzsza: **13500** (Sob)
- Najnizsza: **2900** (Ndz)
- Najwyzszy udzial elektroniki: **39.5%** (Piatek)

**Wyjasnienie**: Udzial procentowy pozwala porownac proporcje niezaleznie od calkowitej sprzedazy. Piatek ma najwyzszy udzial elektroniki (39.5%), mimo ze sobota ma wyzsza bezwzgledna sprzedaz elektroniki (5200 vs 4500). To pokazuje roznice miedzy wartoscia bezwzgledna a wzgledna (procentowa).
</details>

<details>
<summary>Typowe bledy</summary>

- **Dzielenie przez B+C+D zamiast E**: =B2/(B2+C2+D2)*100 zadziala, ale lepiej uzyc E2 (unika duplikacji obliczen). CKE: akceptowane
- **Zapomnienie o ROUND**: Bez zaokraglenia wyniki beda mialy wiele miejsc po przecinku (np. 33.333...). CKE: -1 pkt jesli polecenie wymaga zaokraglenia
- **Pomieszanie najwyzszej sprzedazy z najwyzszym udzialem**: Sob ma najwyzsza sprzedaz (13500), ale Pt ma najwyzszy udzial elektroniki (39.5%). CKE: -1 pkt

</details>

---

### Cwiczenie 18.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 (Martianeum)
**Tagi**: `AVERAGE` `MAX` `MIN` `RANK` `ROUND` `odchylenie-od-sredniej` `ABS` `zakres-bezwzgledny`

Uczniowie brali udzial w 5 konkursach przedmiotowych. Punkty (0-100):

| | A | B | C | D | E | F |
|---|---|---|---|---|---|---|
| 1 | **Uczen** | **Mat** | **Fiz** | **Inf** | **Bio** | **Chem** |
| 2 | Alicja | 78 | 65 | 92 | 54 | 71 |
| 3 | Borys | 85 | 72 | 88 | 61 | 68 |
| 4 | Celina | 62 | 91 | 45 | 83 | 77 |
| 5 | Daniel | 95 | 58 | 97 | 42 | 55 |
| 6 | Ewa | 71 | 80 | 76 | 79 | 82 |
| 7 | Filip | 88 | 69 | 84 | 57 | 63 |
| 8 | Gosia | 55 | 87 | 51 | 90 | 85 |

**Polecenie**:
1. W kolumnie G oblicz srednia punktow kazdego ucznia (ze wszystkich 5 konkursow), zaokraglona do 1 miejsca.
2. W kolumnie H oblicz rozstep punktow kazdego ucznia (MAX - MIN z jego wynikow).
3. W kolumnie I wpisz ranking na podstawie sredniej (1 = najlepsza srednia).
4. Oblicz srednia klasy dla kazdego przedmiotu osobno.
5. Ktory uczen ma najbardziej wyrownane wyniki (najmniejszy rozstep)?
6. Ktory uczen ma najwyzsza srednia, ale jednoczesnie duzy rozstep (>30)?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Kazdy uczen to jeden wiersz — AVERAGE, MAX, MIN operuja na B:F danego wiersza.
2. **Podejscie**: Rozstep = MAX(wiersz) - MIN(wiersz). RANK operuje na kolumnie G (srednie).
3. **Kluczowy krok**: Porownaj ranking ze rozstepem — wysoka srednia z duzym rozstepem oznacza nierownomierne wyniki.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
G2: =ROUND(AVERAGE(B2:F2);1)     (kopiowana w dol)
H2: =MAX(B2:F2)-MIN(B2:F2)       (kopiowana w dol)
I2: =RANK(G2;$G$2:$G$8;0)        (kopiowana w dol)

Srednia klasy Mat: =AVERAGE(B2:B8)
Srednia klasy Fiz: =AVERAGE(C2:C8)
(analogicznie dla Inf, Bio, Chem)
```

**Obliczenia:**

| Uczen | Mat | Fiz | Inf | Bio | Chem | G: Sred. | H: Rozstep | I: Rank |
|-------|-----|-----|-----|-----|------|---------|-----------|---------|
| Alicja | 78 | 65 | 92 | 54 | 71 | 72.0 | 38 | 4 |
| Borys | 85 | 72 | 88 | 61 | 68 | 74.8 | 27 | 2 |
| Celina | 62 | 91 | 45 | 83 | 77 | 71.6 | 46 | 5 |
| Daniel | 95 | 58 | 97 | 42 | 55 | 69.4 | 55 | 6 |
| Ewa | 71 | 80 | 76 | 79 | 82 | 77.6 | 11 | 1 |
| Filip | 88 | 69 | 84 | 57 | 63 | 72.2 | 31 | 3 |
| Gosia | 55 | 87 | 51 | 90 | 85 | 73.6 | 39 | — |

Weryfikacja G:
- Alicja: (78+65+92+54+71)/5 = 360/5 = 72.0 ✓
- Borys: (85+72+88+61+68)/5 = 374/5 = 74.8 ✓
- Ewa: (71+80+76+79+82)/5 = 388/5 = 77.6 ✓
- Daniel: (95+58+97+42+55)/5 = 347/5 = 69.4 ✓
- Gosia: (55+87+51+90+85)/5 = 368/5 = 73.6

Ranking: 77.6 (1.), 74.8 (2.), 73.6 (3.), 72.2 (4.), 72.0 (5.), 71.6 (6.), 69.4 (7.)

Korekta rankingu:
| Uczen | Srednia | Ranking |
|-------|---------|---------|
| Ewa | 77.6 | 1 |
| Borys | 74.8 | 2 |
| Gosia | 73.6 | 3 |
| Filip | 72.2 | 4 |
| Alicja | 72.0 | 5 |
| Celina | 71.6 | 6 |
| Daniel | 69.4 | 7 |

**Srednie na przedmiot:**
- Mat: (78+85+62+95+71+88+55)/7 = 534/7 = **76.3**
- Fiz: (65+72+91+58+80+69+87)/7 = 522/7 = **74.6**
- Inf: (92+88+45+97+76+84+51)/7 = 533/7 = **76.1**
- Bio: (54+61+83+42+79+57+90)/7 = 466/7 = **66.6**
- Chem: (71+68+77+55+82+63+85)/7 = 501/7 = **71.6**

**5. Najbardziej wyrownane wyniki:**
**Ewa** — rozstep = 11 (wyniki 71-82, wszystkie zblizone)

**6. Najwyzsza srednia z duzym rozstepem (>30):**
Osoby z rozstepem >30: Alicja (38), Celina (46), Daniel (55), Filip (31), Gosia (39)
Najwyzsza srednia wsrod nich: **Gosia** (73.6, rozstep 39) lub **Borys** (74.8, rozstep 27 — nie spelnia >30)
Odpowiedz: **Gosia** (srednia 73.6, rozstep 39)

**Wyjasnienie**: Rozstep (MAX-MIN) mierzy rozrzut wynikow ucznia. Maly rozstep = wyrownane wyniki (Ewa: 11), duzy = duza zmiennosc (Daniel: 55 — swietny z Mat/Inf, slaby z Bio). Na maturze analiza rozstepu pozwala interpretowac profile uczniow.
</details>

<details>
<summary>Typowe bledy</summary>

- **AVERAGE z kolumny zamiast wiersza**: =AVERAGE(B2:B8) to srednia wszystkich uczniow z matematyki, nie srednia jednego ucznia. CKE: -2 pkt
- **Brak ROUND w sredniej**: Bez zaokraglenia srednia moze miec wiele miejsc po przecinku. CKE: -1 pkt jesli wymagane
- **Pomieszanie rozstepu z odchyleniem standardowym**: Rozstep = MAX-MIN, odchylenie standardowe to inna miara (STDEV). CKE: -1 pkt

</details>

---

## Samoocena

| Poziom | Opis | Cwiczenia |
|--------|------|-----------|
| Podstawowy | Znam SUM, AVERAGE, MAX, MIN i proste mnozenie komorek | 18.1-18.2, 18.6 bez pomocy |
| Dobry | Uzywam COUNT/COUNTA/COUNTBLANK, RANK, ROUND i licze udzialy procentowe | 18.3-18.4, 18.7-18.8 bez pomocy |
| Bardzo dobry | Obliczam odchylenia od sredniej, uzywam ABS i kolumn pomocniczych | 18.5, 18.9 bez pomocy |
| Doskonaly | Lacze wiele agregatow w analizie wielowymiarowej (srednia + rozstep + ranking) | 18.10 bez pomocy |

**Co dalej?**
- Jesli masz trudnosci z podstawowymi: przejrzyj `cheatsheet_arkusz.md` (sekcja Podstawowe funkcje)
- Jesli opanowales srednie: przejdz do cwiczen 15 (Agregacja warunkowa) — rozszerzenie o SUMIF/COUNTIF
- Jesli zrobiles wszystkie 10: sprobuj cwiczen 16 (Symulacja) lub 17 (Wykresy) ze stoperem
