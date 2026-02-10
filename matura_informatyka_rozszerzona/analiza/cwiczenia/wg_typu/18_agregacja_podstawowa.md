# 18. Agregacja podstawowa w arkuszu kalkulacyjnym

Typ zadania: **arkusz_agregacja_podstawowa**
Czestotliwosc: 3/11 lat | Laczna punktacja: 9 pkt
Kategoria: ARKUSZ

---

### Cwiczenie 18.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna

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

---

### Cwiczenie 18.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2014 (Przychody)

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

---

### Cwiczenie 18.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: ogolna

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
<summary>Odpowiedz</summary>

**Formuly:**
```
Odpowiedzialo:         =COUNTA(B2:B11)
Nie odpowiedzialo:     =COUNTBLANK(B2:B11)
Procent wypelnienia:   =COUNTA(B2:B11)/COUNT(A2:A11)*100
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

---

### Cwiczenie 18.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: ogolna

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
| Czarek | 6 | 9 | 8 | 5 | 28 | 5 |
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

Ranking (malejaco): 33 (1.), 32 (2.), 31 (3.), 30 (4.), 30 (4.), 28 (5.)
- Uwaga: Adam i Fiona maja taki sam wynik (30), wiec obie maja ranking 4. Czarek z 28 ma ranking 5 (nie 6!). Tak dziala RANK — pomija pozycje przy remisach, ale kolejna pozycja jest o tyle nizsza ile bylo remisow plus 1 mniej. W tym wypadku: po pozycji 4 (dwa razy) nastepna jest 6, nie 5.

Korekta: RANK z dwiema wartosciami 30 → obie dostana rang 4, a nastepna ranga to 6 (nie 5).

| Zawodnik | Suma | Ranking |
|----------|------|---------|
| Beata | 33 | 1 |
| Diana | 32 | 2 |
| Emil | 31 | 3 |
| Adam | 30 | 4 |
| Fiona | 30 | 4 |
| Czarek | 28 | 6 |

- Najwyzsza suma: =MAX(F2:F7) = **33** (Beata)
- Srednia suma: =AVERAGE(F2:F7) = (30+33+28+32+31+30)/6 = 184/6 = **30.67**

**Wyjasnienie**: RANK(wartosc; zakres; kolejnosc) zwraca pozycje wartosci w rankingu. Trzeci argument: 0 = malejaco (najwieksza = 1), 1 = rosnaco. Zakres `$F$2:$F$7` musi byc bezwzgledny (z $), aby po skopiowaniu formuly zawsze odwolywal sie do calej kolumny sum.
</details>

---

### Cwiczenie 18.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: ogolna

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
