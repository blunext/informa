# 19. Transformacja danych w arkuszu kalkulacyjnym

Typ zadania: **arkusz_transformacja**
Czestotliwosc: 2/11 lat | Laczna punktacja: 3 pkt
Kategoria: ARKUSZ

## Umiejetnosci cwiczone w tym zestawie

`SUMIF` `SUMIFS` `COUNTIFS` `tabela-krzyzowa` `kolumna-pomocnicza` `INT-dzielenie` `odwolania-mieszane` `grupowanie-danych` `VLOOKUP` `INDEX-MATCH` `LEFT-RIGHT-MID` `TEXT` `konkatenacja` `pivot-reczny`

---

### Cwiczenie 19.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `SUMIF` `grupowanie-danych`

W arkuszu mamy dane dziennej sprzedazy z podanym miesiacem:

| | A | B | C |
|---|---|---|---|
| 1 | **Data** | **Miesiac** | **Sprzedaz** |
| 2 | 2024-01-03 | Styczen | 1200 |
| 3 | 2024-01-15 | Styczen | 800 |
| 4 | 2024-01-22 | Styczen | 1500 |
| 5 | 2024-02-05 | Luty | 900 |
| 6 | 2024-02-18 | Luty | 1100 |
| 7 | 2024-03-02 | Marzec | 1400 |
| 8 | 2024-03-10 | Marzec | 600 |
| 9 | 2024-03-21 | Marzec | 1300 |
| 10 | 2024-03-28 | Marzec | 700 |

W osobnej tabeli chcemy uzyskac sumy miesieczne:

| | E | F |
|---|---|---|
| 1 | **Miesiac** | **Suma sprzedazy** |
| 2 | Styczen | ? |
| 3 | Luty | ? |
| 4 | Marzec | ? |

**Polecenie**: Napisz formule w F2, ktora zsumuje sprzedaz ze wszystkich dni nalezacych do danego miesiaca.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Potrzebujesz sumy warunkowej — SUMIF.
2. **Podejscie**: SUMIF(zakres_kryterium; kryterium; zakres_sumy) — filtruj po kolumnie B, sumuj kolumne C.
3. **Kluczowy krok**: Kryterium to E2 (nazwa miesiaca z tabeli docelowej), nie wpisany tekst.

</details>

<details>
<summary>Odpowiedz</summary>

**Formula (F2, kopiowana w dol):**
```
F2: =SUMIF(B2:B10;E2;C2:C10)
```

**Weryfikacja:**
- Styczen: 1200 + 800 + 1500 = **3500**
- Luty: 900 + 1100 = **2000**
- Marzec: 1400 + 600 + 1300 + 700 = **4000**

Kontrola: 3500 + 2000 + 4000 = 9500 = SUM(C2:C10) ✓

**Wyjasnienie**: SUMIF(zakres_kryterium; kryterium; zakres_sumy) sumuje wartosci z kolumny C tylko dla wierszy, gdzie kolumna B odpowiada wartosci z E2. To najprostsza forma transformacji — grupowanie danych szczegolowych w podsumowanie. Kryterium E2 jest wzgledne, wiec po skopiowaniu do F3 stanie sie E3 ("Luty"), a do F4 — E4 ("Marzec").
</details>

<details>
<summary>Typowe bledy</summary>

- **Wpisanie tekstu zamiast odwolania**: =SUMIF(B2:B10;"Styczen";C2:C10) zadziala, ale nie skopiujesz formuly w dol (kazdy wiersz bedzie mial "Styczen"). CKE: -1 pkt
- **Zly porzadek argumentow SUMIF**: SUMIF ma zakres_kryterium PRZED zakresem_sumy (odwrotnie niz SUMIFS!). CKE: -1 pkt

</details>

---

### Cwiczenie 19.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `SUMIFS` `tabela-krzyzowa` `odwolania-mieszane`

W arkuszu mamy liste ocen w formacie: uczen, przedmiot, ocena:

| | A | B | C |
|---|---|---|---|
| 1 | **Uczen** | **Przedmiot** | **Ocena** |
| 2 | Anna | Matematyka | 4 |
| 3 | Anna | Polski | 5 |
| 4 | Anna | Angielski | 4 |
| 5 | Bartek | Matematyka | 3 |
| 6 | Bartek | Polski | 4 |
| 7 | Bartek | Angielski | 5 |
| 8 | Celina | Matematyka | 5 |
| 9 | Celina | Polski | 3 |
| 10 | Celina | Angielski | 4 |

Chcemy utworzyc tabele: uczen (wiersz) × przedmiot (kolumna):

| | E | F | G | H |
|---|---|---|---|---|
| 1 | | **Matematyka** | **Polski** | **Angielski** |
| 2 | **Anna** | ? | ? | ? |
| 3 | **Bartek** | ? | ? | ? |
| 4 | **Celina** | ? | ? | ? |

**Polecenie**: Napisz formule w F2, ktora znajdzie ocene ucznia z wiersza 2 (Anna) z przedmiotu z wiersza 1 (Matematyka). Formula powinna dzialac po skopiowaniu w prawo i w dol.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: SUMIFS z dwoma warunkami (uczen i przedmiot). Odwolania musza byc mieszane ($ na kolumnie LUB wierszu).
2. **Podejscie**: $E2 zakotwicza kolumne E (uczen), F$1 zakotwicza wiersz 1 (przedmiot).
3. **Kluczowy krok**: Zakresy danych (A, B, C) musza miec $ na wierszach, aby nie przesuway sie przy kopiowaniu.

</details>

<details>
<summary>Odpowiedz</summary>

**Formula (F2, kopiowana w prawo i w dol):**
```
F2: =SUMIFS(C$1:C$10;A$1:A$10;$E2;B$1:B$10;F$1)
```

Alternatywnie z INDEX-MATCH (bardziej zaawansowane):
```
F2: =INDEX(C$2:C$10;MATCH($E2&F$1;A$2:A$10&B$2:B$10;0))
```
(Uwaga: formula INDEX-MATCH z konkatenacja wymaga zatwierdzenia Ctrl+Shift+Enter w starszych wersjach Excela)

**Weryfikacja (uzycie SUMIFS):**

SUMIFS sumuje wartosci C, gdzie A=uczen ORAZ B=przedmiot. Poniewaz kazda kombinacja uczen-przedmiot wystepuje dokladnie raz, SUMIFS zwroci te jedna ocene.

| | Matematyka | Polski | Angielski |
|---|---|---|---|
| Anna | 4 | 5 | 4 |
| Bartek | 3 | 4 | 5 |
| Celina | 5 | 3 | 4 |

Sprawdzenie F2: SUMIFS szuka wierszy gdzie A="Anna" i B="Matematyka" → wiersz 2, C=4 ✓
Sprawdzenie G3: SUMIFS szuka wierszy gdzie A="Bartek" i B="Polski" → wiersz 6, C=4 ✓
Sprawdzenie H4: SUMIFS szuka wierszy gdzie A="Celina" i B="Angielski" → wiersz 10, C=4 ✓

**Wyjasnienie odniesien:**
- `$E2` — kolumna E zakotwiczona ($E), wiersz wzgledny → po skopiowaniu w prawo nadal czyta z E (imie ucznia)
- `F$1` — wiersz 1 zakotwiczony ($1), kolumna wzgledna → po skopiowaniu w dol nadal czyta z wiersza 1 (nazwa przedmiotu)
- `C$1:C$10` — zakresy danych zakotwiczone (wiersze z $), nie zmieniaja sie przy kopiowaniu

To jest transformacja z listy "plaskiej" (flat list) do tabeli krzyzowej (cross-tab / pivot).
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak $ lub zle $**: =SUMIFS(C2:C10;A2:A10;E2;B2:B10;F1) — bez $ kopiowanie przesunie zakresy i komorki z kryteriami. CKE: -2 pkt
- **$ na zlym elemencie**: $E$2 (oba zakotwiczone) — nie zmieni sie przy kopiowaniu w dol, wiec kazdy wiersz pokaże dane Anny. CKE: -2 pkt

</details>

---

### Cwiczenie 19.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 (Martianeum)
**Tagi**: `INT-dzielenie` `kolumna-pomocnicza` `SUMIF` `grupowanie-danych`

Fabryka notuje produkcje dziennie przez 21 dni. Chcemy pogrupowac dane w tygodnie (dni 1-7, 8-14, 15-21):

| | A | B |
|---|---|---|
| 1 | **Dzien** | **Produkcja** |
| 2 | 1 | 120 |
| 3 | 2 | 135 |
| 4 | 3 | 110 |
| 5 | 4 | 145 |
| 6 | 5 | 130 |
| 7 | 6 | 95 |
| 8 | 7 | 100 |
| 9 | 8 | 140 |
| 10 | 9 | 125 |
| 11 | 10 | 155 |
| 12 | 11 | 110 |
| 13 | 12 | 130 |
| 14 | 13 | 120 |
| 15 | 14 | 145 |
| 16 | 15 | 160 |
| 17 | 16 | 135 |
| 18 | 17 | 115 |
| 19 | 18 | 150 |
| 20 | 19 | 140 |
| 21 | 20 | 125 |
| 22 | 21 | 130 |

**Polecenie**:
1. W kolumnie C dodaj kolumne pomocnicza "Tydzien" obliczana formula: numer tygodnia = INT((dzien-1)/7)+1.
2. W osobnej tabeli oblicz sume produkcji na kazdy tydzien uzywajac SUMIF.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: INT((dzien-1)/7) dzieli dni na grupy po 7. Dodanie +1 numeruje od 1.
2. **Podejscie**: Kolumna pomocnicza C zamienia numer dnia na numer tygodnia. Potem SUMIF grupuje po C.
3. **Kluczowy krok**: Sprawdz wartosci graniczne: dzien 7 → INT(6/7)+1 = 0+1 = 1, dzien 8 → INT(7/7)+1 = 1+1 = 2.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
C2: =INT((A2-1)/7)+1       (kopiowana w dol)
```

Tabela tygodniowa:
| | E | F |
|---|---|---|
| 1 | **Tydzien** | **Suma** |
| 2 | 1 | =SUMIF(C2:C22;E2;B2:B22) |
| 3 | 2 | =SUMIF(C2:C22;E3;B2:B22) |
| 4 | 3 | =SUMIF(C2:C22;E4;B2:B22) |

**Kolumna pomocnicza C:**

| Dzien | Produkcja | Tydzien |
|-------|-----------|---------|
| 1 | 120 | INT(0/7)+1 = 1 |
| 2 | 135 | INT(1/7)+1 = 1 |
| ... | ... | ... |
| 7 | 100 | INT(6/7)+1 = 1 |
| 8 | 140 | INT(7/7)+1 = 2 |
| ... | ... | ... |
| 14 | 145 | INT(13/7)+1 = 2 |
| 15 | 160 | INT(14/7)+1 = 3 |
| ... | ... | ... |
| 21 | 130 | INT(20/7)+1 = 3 |

**Sumy tygodniowe:**
- Tydzien 1 (dni 1-7): 120+135+110+145+130+95+100 = **835**
- Tydzien 2 (dni 8-14): 140+125+155+110+130+120+145 = **925**
- Tydzien 3 (dni 15-21): 160+135+115+150+140+125+130 = **955**

Kontrola: 835+925+955 = 2715 = SUM(B2:B22) ✓

**Wyjasnienie**: INT((dzien-1)/7) daje numer tygodnia zaczynajac od 0. Dodanie +1 sprawia, ze tygodnie sa numerowane od 1. Dzielenie calkowite INT() zaokragla w dol — wszystkie dni od 1 do 7 daja INT((0..6)/7) = 0, czyli tydzien 1. Kolumna pomocnicza to czesty wzorzec na maturze — tworzymy ja, aby moc uzyc SUMIF do grupowania.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak -1 w formule**: INT(A2/7)+1 da bledne granice — dzien 7 trafi do tygodnia 2 (INT(7/7)+1=2). CKE: -2 pkt
- **ROUND zamiast INT**: ROUND zaokragla do najblizszej, INT zawsze w dol. Dla dzielenia calkowitego potrzebny jest INT. CKE: -1 pkt

</details>

---

### Cwiczenie 19.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `COUNTIFS` `tabela-krzyzowa` `odwolania-mieszane`

Przeprowadzono ankiete z 4 pytaniami. Kazda odpowiedz to A, B, C lub D:

| | A | B |
|---|---|---|
| 1 | **Pytanie** | **Odpowiedz** |
| 2 | P1 | A |
| 3 | P1 | B |
| 4 | P1 | A |
| 5 | P1 | C |
| 6 | P1 | A |
| 7 | P2 | B |
| 8 | P2 | B |
| 9 | P2 | D |
| 10 | P2 | A |
| 11 | P2 | B |
| 12 | P3 | C |
| 13 | P3 | C |
| 14 | P3 | A |
| 15 | P3 | B |
| 16 | P3 | C |

Chcemy utworzyc tabele krzyzowa: pytanie (wiersz) × odpowiedz (kolumna) z liczba wystapien:

| | D | E | F | G | H |
|---|---|---|---|---|---|
| 1 | | **A** | **B** | **C** | **D** |
| 2 | **P1** | ? | ? | ? | ? |
| 3 | **P2** | ? | ? | ? | ? |
| 4 | **P3** | ? | ? | ? | ? |

**Polecenie**: Napisz formule w E2, ktora zliczy ile razy pytanie P1 otrzymalo odpowiedz A. Formula musi dzialac po skopiowaniu w prawo i w dol.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: COUNTIFS z dwoma parami kryteriow — pytanie i odpowiedz.
2. **Podejscie**: $D2 zakotwicza kolumne (pytanie), E$1 zakotwicza wiersz (odpowiedz).
3. **Kluczowy krok**: Weryfikacja: suma kazdego wiersza powinna rownych sie liczbie odpowiedzi na to pytanie.

</details>

<details>
<summary>Odpowiedz</summary>

**Formula (E2, kopiowana w prawo i w dol):**
```
E2: =COUNTIFS(A$2:A$16;$D2;B$2:B$16;E$1)
```

**Wypelniona tabela:**

| | A | B | C | D |
|---|---|---|---|---|
| P1 | 3 | 1 | 1 | 0 |
| P2 | 1 | 3 | 0 | 1 |
| P3 | 1 | 1 | 3 | 0 |

**Weryfikacja:**
- P1: A(wiersze 2,4,6)=3, B(wiersz 3)=1, C(wiersz 5)=1, D=0 → suma=5 ✓
- P2: A(wiersz 10)=1, B(wiersze 7,8,11)=3, C=0, D(wiersz 9)=1 → suma=5 ✓
- P3: A(wiersz 14)=1, B(wiersz 15)=1, C(wiersze 12,13,16)=3, D=0 → suma=5 ✓

**Wyjasnienie odniesien w COUNTIFS:**
- `A$2:A$16` — zakres danych z $ na wierszach (nie zmienia sie przy kopiowaniu)
- `$D2` — kolumna D zakotwiczona (czyta nazwe pytania), wiersz wzgledny
- `B$2:B$16` — zakres odpowiedzi z $ na wierszach
- `E$1` — wiersz 1 zakotwiczony (czyta nazwe odpowiedzi), kolumna wzgledna

COUNTIFS z dwoma parami kryteriow zlicza wiersze spelniajace OBA warunki jednoczesnie. To klasyczna transformacja z listy do tabeli krzyzowej za pomoca formul.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak $ na wierszach zakresow danych**: Po skopiowaniu w dol zakresy A2:A16 i B2:B16 przesuna sie. CKE: -2 pkt
- **$ na obu elementach kryterium**: $D$2 nie zmieni sie w dol — kazdy wiersz pokaże P1. CKE: -2 pkt
- **Brak weryfikacji sum**: Nie sprawdzenie, ze suma wiersza = liczba odpowiedzi na pytanie. CKE: brak kary, ale pomaga znalezc bledy

</details>

---

### Cwiczenie 19.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)
**Tagi**: `SUMIFS` `tabela-krzyzowa` `odwolania-mieszane` `pivot-reczny`

Sklep ma dane sprzedazy z calego roku:

| | A | B | C |
|---|---|---|---|
| 1 | **Kategoria** | **Kwartal** | **Wartosc** |
| 2 | Elektronika | Q1 | 15000 |
| 3 | Odziez | Q1 | 8000 |
| 4 | Zywnosc | Q1 | 12000 |
| 5 | Elektronika | Q2 | 18000 |
| 6 | Odziez | Q2 | 9500 |
| 7 | Zywnosc | Q2 | 11000 |
| 8 | Elektronika | Q1 | 13000 |
| 9 | Odziez | Q3 | 11000 |
| 10 | Zywnosc | Q3 | 13500 |
| 11 | Elektronika | Q3 | 20000 |
| 12 | Odziez | Q1 | 7500 |
| 13 | Zywnosc | Q4 | 14000 |
| 14 | Elektronika | Q4 | 22000 |
| 15 | Odziez | Q4 | 10000 |
| 16 | Zywnosc | Q2 | 10500 |
| 17 | Elektronika | Q2 | 16000 |
| 18 | Odziez | Q3 | 8500 |
| 19 | Zywnosc | Q4 | 12500 |
| 20 | Elektronika | Q3 | 19000 |
| 21 | Odziez | Q4 | 9000 |

Chcemy utworzyc tabele pivot — sumy sprzedazy wg kategorii (wiersze) i kwartalow (kolumny):

| | E | F | G | H | I | J |
|---|---|---|---|---|---|---|
| 1 | | **Q1** | **Q2** | **Q3** | **Q4** | **Razem** |
| 2 | **Elektronika** | ? | ? | ? | ? | ? |
| 3 | **Odziez** | ? | ? | ? | ? | ? |
| 4 | **Zywnosc** | ? | ? | ? | ? | ? |
| 5 | **Razem** | ? | ? | ? | ? | ? |

**Polecenie**:
1. Napisz formule w F2 (suma Elektroniki w Q1) uzywajac SUMIFS — z odniesieniami, ktore pozwola na kopiowanie w prawo i w dol.
2. Napisz formule w J2 (razem dla wiersza) i F5 (razem dla kolumny).
3. Wypelnij cala tabele i zweryfikuj, ze sumy wierszy i kolumn sie zgadzaja.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: SUMIFS z $E2 (kategoria) i F$1 (kwartal) — klasyczne odwolania mieszane do tabeli krzyzowej.
2. **Podejscie**: Wiersz "Razem" to SUM kolumny, kolumna "Razem" to SUM wiersza. Punkt kontrolny: J5 = suma wszystkich zrodlowych.
3. **Kluczowy krok**: Weryfikacja krzyzowa — suma wierszow Razem = suma kolumn Razem = SUM(C2:C21).

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
F2: =SUMIFS(C$2:C$21;A$2:A$21;$E2;B$2:B$21;F$1)
    (kopiowana w prawo do I2 i w dol do I4)

J2: =SUM(F2:I2)          (suma wiersza, kopiowana w dol)
F5: =SUM(F2:F4)          (suma kolumny, kopiowana w prawo)
J5: =SUM(J2:J4)          (suma calosci, lub =SUM(F5:I5) — te same)
```

**Wypelniona tabela:**

Elektronika w Q1: wiersze 2 (15000) i 8 (13000) = 28000
Elektronika w Q2: wiersze 5 (18000) i 17 (16000) = 34000
Elektronika w Q3: wiersze 11 (20000) i 20 (19000) = 39000
Elektronika w Q4: wiersz 14 (22000) = 22000

Odziez w Q1: wiersze 3 (8000) i 12 (7500) = 15500
Odziez w Q2: wiersz 6 (9500) = 9500
Odziez w Q3: wiersze 9 (11000) i 18 (8500) = 19500
Odziez w Q4: wiersze 15 (10000) i 21 (9000) = 19000

Zywnosc w Q1: wiersz 4 (12000) = 12000
Zywnosc w Q2: wiersze 7 (11000) i 16 (10500) = 21500
Zywnosc w Q3: wiersz 10 (13500) = 13500
Zywnosc w Q4: wiersze 13 (14000) i 19 (12500) = 26500

| | Q1 | Q2 | Q3 | Q4 | Razem |
|---|---|---|---|---|---|
| Elektronika | 28000 | 34000 | 39000 | 22000 | 123000 |
| Odziez | 15500 | 9500 | 19500 | 19000 | 63500 |
| Zywnosc | 12000 | 21500 | 13500 | 26500 | 73500 |
| Razem | 55500 | 65000 | 72000 | 67500 | **260000** |

**Weryfikacja krzyzowa:**
- Suma wierszow: 123000 + 63500 + 73500 = 260000 ✓
- Suma kolumn: 55500 + 65000 + 72000 + 67500 = 260000 ✓
- Suma wszystkich wartosci zrodlowych: =SUM(C2:C21) = 260000 ✓

**Wyjasnienie odniesien w SUMIFS:**
- `$E2` — zakotwiczona kolumna E (kategoria), wiersz wzgledny → kopiowanie w dol zmienia E2→E3→E4
- `F$1` — zakotwiczony wiersz 1 (kwartal), kolumna wzgledna → kopiowanie w prawo zmienia F→G→H→I
- `C$2:C$21`, `A$2:A$21`, `B$2:B$21` — zakresy danych z $ na wierszach, nie zmieniaja sie

To jest reczna tabela przestawna (pivot table) zbudowana z formul. Na maturze dane zrodlowe czesto maja wiele wierszy z ta sama kombinacja kategorii — SUMIFS poprawnie je sumuje. Kluczowe jest prawidlowe uzycie $ w odniesieniach, aby jedna formula dzialala po skopiowaniu w cala tabele.
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomieszanie $ w odwolaniach**: Najczestszy blad — $ na zlym elemencie powoduje, ze formula nie kopiuje sie poprawnie. CKE: -2 pkt
- **Brak weryfikacji krzyzowej**: Suma wiersza Razem musi = suma kolumny Razem. Jesli nie, gdzies jest blad. CKE: -1 pkt jesli wynik bledny
- **Uzycie SUMIF zamiast SUMIFS**: SUMIF obsluguje tylko 1 warunek — tu potrzebne sa 2 (kategoria i kwartal). CKE: -2 pkt

</details>

---

### Cwiczenie 19.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `LEFT-RIGHT-MID` `konkatenacja` `TEXT`

W arkuszu mamy dane w formacie wymagajacym transformacji tekstowej:

| | A | B |
|---|---|---|
| 1 | **Kod produktu** | **Cena** |
| 2 | EL-001-2024 | 299.90 |
| 3 | OD-015-2024 | 89.50 |
| 4 | ZY-102-2024 | 12.30 |
| 5 | EL-042-2023 | 450.00 |
| 6 | OD-008-2023 | 65.00 |
| 7 | ZY-055-2024 | 8.90 |

Kody produktow maja format: XX-NNN-RRRR (kategoria-numer-rok).

**Polecenie**:
1. W kolumnie C wyodrebnij kategorie (pierwsze 2 znaki).
2. W kolumnie D wyodrebnij numer produktu (znaki 4-6).
3. W kolumnie E wyodrebnij rok (ostatnie 4 znaki).
4. W kolumnie F utworz nowy format: "Kategoria: XX, Nr: NNN (RRRR)".

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: LEFT wyciaga znaki od lewej, RIGHT od prawej, MID ze srodka.
2. **Podejscie**: LEFT(A2;2) = kategoria, MID(A2;4;3) = numer, RIGHT(A2;4) = rok.
3. **Kluczowy krok**: Konkatenacja operatorem & laczy teksty: "Kategoria: "&C2&", Nr: "&D2&" ("&E2&")".

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
C2: =LEFT(A2;2)
D2: =MID(A2;4;3)
E2: =RIGHT(A2;4)
F2: ="Kategoria: "&C2&", Nr: "&D2&" ("&E2&")"
```

**Wyniki:**

| Kod | Cena | C: Kat. | D: Nr | E: Rok | F: Nowy format |
|-----|------|---------|-------|--------|----------------|
| EL-001-2024 | 299.90 | EL | 001 | 2024 | Kategoria: EL, Nr: 001 (2024) |
| OD-015-2024 | 89.50 | OD | 015 | 2024 | Kategoria: OD, Nr: 015 (2024) |
| ZY-102-2024 | 12.30 | ZY | 102 | 2024 | Kategoria: ZY, Nr: 102 (2024) |
| EL-042-2023 | 450.00 | EL | 042 | 2023 | Kategoria: EL, Nr: 042 (2023) |
| OD-008-2023 | 65.00 | OD | 008 | 2023 | Kategoria: OD, Nr: 008 (2023) |
| ZY-055-2024 | 8.90 | ZY | 055 | 2024 | Kategoria: ZY, Nr: 055 (2024) |

**Wyjasnienie**: LEFT(tekst; n) zwraca n znakow od lewej. MID(tekst; start; n) zwraca n znakow od pozycji start (numerowane od 1). RIGHT(tekst; n) zwraca n znakow od prawej. Operator & laczy (konkatenuje) teksty. Na maturze te funkcje pojawiaja sie przy przetwarzaniu kodow, dat i identyfikatorow.
</details>

<details>
<summary>Typowe bledy</summary>

- **Bledna pozycja poczatkowa w MID**: MID(A2;3;3) zaczyna od myślnika, nie od numeru. Pozycja 4 to pierwszy znak numeru. CKE: -1 pkt
- **Zapomnienie o & miedzy fragmentami**: "Kategoria: " C2 bez & nie zadziala — trzeba "Kategoria: "&C2. CKE: -1 pkt

</details>

---

### Cwiczenie 19.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2018 (Szkolna biblioteka)
**Tagi**: `VLOOKUP` `kolumna-pomocnicza` `grupowanie-danych`

W arkuszu mamy liste wypozyczen z biblioteki:

| | A | B | C |
|---|---|---|---|
| 1 | **ID czytelnika** | **Tytul** | **Data zwrotu** |
| 2 | C001 | Harry Potter | 2024-02-15 |
| 3 | C002 | Lalka | 2024-02-20 |
| 4 | C001 | Solaris | 2024-03-01 |
| 5 | C003 | Ferdydurke | 2024-02-28 |
| 6 | C002 | Pan Tadeusz | 2024-03-10 |
| 7 | C001 | Diune | 2024-03-15 |
| 8 | C003 | Potop | 2024-02-25 |
| 9 | C002 | Hobbit | 2024-03-05 |

W osobnym arkuszu mamy dane czytelnikow:

| | E | F | G |
|---|---|---|---|
| 1 | **ID** | **Imie** | **Klasa** |
| 2 | C001 | Anna | 3A |
| 3 | C002 | Bartek | 3B |
| 4 | C003 | Celina | 3A |

**Polecenie**:
1. W kolumnie D wyszukaj imie czytelnika na podstawie ID (uzyj VLOOKUP).
2. Ile wypozyczen mial kazdy czytelnik (uzyj COUNTIF)?
3. Ile wypozyczen przypadlo na klase 3A?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: VLOOKUP szuka wartosci w pierwszej kolumnie tabeli i zwraca wartosc z innej kolumny.
2. **Podejscie**: VLOOKUP(A2;$E$2:$G$4;2;0) — szuka ID w tabeli E:G, zwraca kolumne 2 (imie).
3. **Kluczowy krok**: Dla klasy 3A — najpierw wyszukaj klase (VLOOKUP z kolumna 3), potem COUNTIF po klasie.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
D2: =VLOOKUP(A2;$E$2:$G$4;2;0)     (kopiowana w dol)

Wypozyczenia C001: =COUNTIF(A2:A9;"C001")
Wypozyczenia C002: =COUNTIF(A2:A9;"C002")
Wypozyczenia C003: =COUNTIF(A2:A9;"C003")

Klasa czytelnika (kolumna pomocnicza H):
H2: =VLOOKUP(A2;$E$2:$G$4;3;0)     (kopiowana w dol)

Wypozyczenia klasy 3A: =COUNTIF(H2:H9;"3A")
```

**Wyniki:**

| ID | Tytul | Data | D: Imie | H: Klasa |
|----|-------|------|---------|----------|
| C001 | Harry Potter | 2024-02-15 | Anna | 3A |
| C002 | Lalka | 2024-02-20 | Bartek | 3B |
| C001 | Solaris | 2024-03-01 | Anna | 3A |
| C003 | Ferdydurke | 2024-02-28 | Celina | 3A |
| C002 | Pan Tadeusz | 2024-03-10 | Bartek | 3B |
| C001 | Diune | 2024-03-15 | Anna | 3A |
| C003 | Potop | 2024-02-25 | Celina | 3A |
| C002 | Hobbit | 2024-03-05 | Bartek | 3B |

- C001 (Anna): **3** wypozyczenia
- C002 (Bartek): **3** wypozyczenia
- C003 (Celina): **2** wypozyczenia

Klasa 3A (Anna + Celina): 3 + 2 = **5** wypozyczen
Klasa 3B (Bartek): **3** wypozyczenia

**Wyjasnienie**: VLOOKUP(szukana; tabela; kolumna; dokladnosc) wyszukuje wartosci w tabeli referencyjnej. Czwarty argument 0 (lub FALSE) oznacza dokladne dopasowanie. Tabela referencjna musi miec klucz (ID) w PIERWSZEJ kolumnie. VLOOKUP + COUNTIF to czesty wzorzec na maturze — laczymy dane z dwoch tabel.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak 0 (FALSE) jako czwarty argument**: Bez niego VLOOKUP uzywa przyblizonego dopasowania, co wymaga posortowanych danych. CKE: -1 pkt
- **Brak $ w zakresie tabeli**: $E$2:$G$4 — bez $ kopiowanie przesunie zakres. CKE: -2 pkt
- **Zly numer kolumny**: Kolumna 1 = ID, kolumna 2 = Imie, kolumna 3 = Klasa. Pomieszanie numerow. CKE: -1 pkt

</details>

---

### Cwiczenie 19.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: ogolna
**Tagi**: `LEFT-RIGHT-MID` `INT-dzielenie` `kolumna-pomocnicza` `SUMIF` `grupowanie-danych`

W arkuszu mamy dane o transakcjach z identyfikatorem w formacie RRRR-MM-NNNN:

| | A | B |
|---|---|---|
| 1 | **ID transakcji** | **Kwota** |
| 2 | 2024-01-0001 | 250 |
| 3 | 2024-01-0015 | 180 |
| 4 | 2024-02-0003 | 420 |
| 5 | 2024-02-0008 | 310 |
| 6 | 2024-03-0002 | 550 |
| 7 | 2024-03-0011 | 190 |
| 8 | 2024-01-0022 | 340 |
| 9 | 2024-02-0017 | 275 |
| 10 | 2024-03-0005 | 610 |
| 11 | 2024-03-0019 | 160 |

**Polecenie**:
1. W kolumnie C wyodrebnij rok z ID transakcji.
2. W kolumnie D wyodrebnij miesiac (jako liczbe).
3. W kolumnie E wyodrebnij numer transakcji (jako liczbe).
4. W kolumnie F przypisz nazwe kwartalu na podstawie miesiaca (1-3 = Q1, 4-6 = Q2, ...).
5. Oblicz sume kwot transakcji na kazdy kwartal.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: LEFT/MID/RIGHT + VALUE do konwersji tekstu na liczbe. Kwartal = INT((miesiac-1)/3)+1.
2. **Podejscie**: C2=LEFT(A2;4), D2=VALUE(MID(A2;6;2)), E2=VALUE(RIGHT(A2;4)).
3. **Kluczowy krok**: "Q"&INT((D2-1)/3)+1 laczy "Q" z obliczonym numerem kwartalu.

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly:**
```
C2: =LEFT(A2;4)                           (rok jako tekst)
D2: =VALUE(MID(A2;6;2))                   (miesiac jako liczba)
E2: =VALUE(RIGHT(A2;4))                   (numer jako liczba)
F2: ="Q"&INT((D2-1)/3)+1                  (kwartal)
```

**Wyniki:**

| ID | Kwota | C: Rok | D: Mies. | E: Nr | F: Kwartal |
|----|-------|--------|----------|-------|------------|
| 2024-01-0001 | 250 | 2024 | 1 | 1 | Q1 |
| 2024-01-0015 | 180 | 2024 | 1 | 15 | Q1 |
| 2024-02-0003 | 420 | 2024 | 2 | 3 | Q1 |
| 2024-02-0008 | 310 | 2024 | 2 | 8 | Q1 |
| 2024-03-0002 | 550 | 2024 | 3 | 2 | Q1 |
| 2024-03-0011 | 190 | 2024 | 3 | 11 | Q1 |
| 2024-01-0022 | 340 | 2024 | 1 | 22 | Q1 |
| 2024-02-0017 | 275 | 2024 | 2 | 17 | Q1 |
| 2024-03-0005 | 610 | 2024 | 3 | 5 | Q1 |
| 2024-03-0019 | 160 | 2024 | 3 | 19 | Q1 |

Uwaga: Wszystkie transakcje sa ze stycznia-marca 2024, wiec wszystkie naleza do Q1.

**Sumy kwartalne:**
- Q1: =SUMIF(F2:F11;"Q1";B2:B11) = 250+180+420+310+550+190+340+275+610+160 = **3285**

Kontrola: SUM(B2:B11) = 3285 ✓

**Wyjasnienie**: VALUE() konwertuje tekst na liczbe — MID zwraca tekst "01", VALUE zamienia go na 1. Bez VALUE nie mozna wykonywac obliczen (INT, porownania). Wzorzec ID→ekstrakcja→kolumna pomocnicza→grupowanie jest bardzo czesty na maturze.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak VALUE**: MID(A2;6;2) zwraca tekst "01", nie liczbe 1. Obliczenia na tekście daja bledy. CKE: -2 pkt
- **Bledna pozycja MID**: Rok to znaki 1-4, myślnik to 5, miesiac to 6-7. CKE: -1 pkt
- **Bledna formula kwartalu**: INT((1-1)/3)+1 = 1 (Q1), INT((4-1)/3)+1 = 2 (Q2) — sprawdz wartosci graniczne. CKE: -1 pkt

</details>

---

### Cwiczenie 19.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2023 (Konfitury)
**Tagi**: `INDEX-MATCH` `VLOOKUP` `tabela-krzyzowa` `odwolania-mieszane`

Sklep ma cennik w formacie tabeli krzyzowej:

| | A | B | C | D |
|---|---|---|---|---|
| 1 | **Rozmiar \ Kolor** | **Czerwony** | **Niebieski** | **Zielony** |
| 2 | S | 49.90 | 54.90 | 44.90 |
| 3 | M | 59.90 | 64.90 | 54.90 |
| 4 | L | 69.90 | 74.90 | 64.90 |
| 5 | XL | 79.90 | 84.90 | 74.90 |

Lista zamowien:

| | F | G | H |
|---|---|---|---|
| 1 | **Rozmiar** | **Kolor** | **Cena** |
| 2 | M | Niebieski | ? |
| 3 | XL | Czerwony | ? |
| 4 | S | Zielony | ? |
| 5 | L | Niebieski | ? |
| 6 | M | Czerwony | ? |

**Polecenie**: Napisz formule w H2, ktora automatycznie wyszuka cene na podstawie rozmiaru (F2) i koloru (G2) z tabeli cennikowej. Formula powinna dzialac po skopiowaniu w dol.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: INDEX+MATCH — INDEX zwraca wartosc z tabeli na przecieciu wiersza i kolumny. MATCH szuka pozycji wartosci.
2. **Podejscie**: INDEX(tabela_cen; MATCH(rozmiar;kolumna_rozmiarow;0); MATCH(kolor;wiersz_kolorow;0)).
3. **Kluczowy krok**: MATCH zwraca NUMER pozycji, nie wartosc. INDEX uzywa tego numeru do pobrania ceny.

</details>

<details>
<summary>Odpowiedz</summary>

**Formula:**
```
H2: =INDEX($B$2:$D$5;MATCH(F2;$A$2:$A$5;0);MATCH(G2;$B$1:$D$1;0))
```

Alternatywnie:
```
H2: =INDEX($A$1:$D$5;MATCH(F2;$A$2:$A$5;0)+1;MATCH(G2;$B$1:$D$1;0)+1)
```

**Wyjanienie formuly:**
- `$B$2:$D$5` — tabela cen (bez naglowkow)
- `MATCH(F2;$A$2:$A$5;0)` — szuka rozmiaru F2 w kolumnie A, zwraca numer wiersza
- `MATCH(G2;$B$1:$D$1;0)` — szuka koloru G2 w wierszu 1, zwraca numer kolumny

**Weryfikacja:**

| Rozmiar | Kolor | MATCH rozmiar | MATCH kolor | Cena |
|---------|-------|---------------|-------------|------|
| M | Niebieski | 2 | 2 | 64.90 |
| XL | Czerwony | 4 | 1 | 79.90 |
| S | Zielony | 1 | 3 | 44.90 |
| L | Niebieski | 3 | 2 | 74.90 |
| M | Czerwony | 2 | 1 | 59.90 |

Kontrola z tabela: M-Niebieski = 64.90 ✓, XL-Czerwony = 79.90 ✓, S-Zielony = 44.90 ✓

**Wyjasnienie**: INDEX+MATCH to potezniejsza alternatywa dla VLOOKUP — pozwala szukac w dwoch wymiarach jednoczesnie. VLOOKUP szuka tylko w jednej kolumnie (lewej), a INDEX+MATCH w dowolnym kierunku. Na maturze INDEX+MATCH pojawia sie przy wyszukiwaniu w tablicach 2D (cenniki, rozmiary, odleglosci).
</details>

<details>
<summary>Typowe bledy</summary>

- **Zakres INDEX wlaczajacy naglowki**: Jesli INDEX obejmuje A1:D5, to MATCH musi zwrocic pozycje +1. Lepiej uzyc B2:D5 bez naglowkow. CKE: -1 pkt
- **Brak $ w zakresach tabeli cennikowej**: Kopiowanie w dol przesunie zakresy. CKE: -2 pkt
- **Uzycie VLOOKUP zamiast INDEX-MATCH**: VLOOKUP nie obsluguje wyszukiwania dwuwymiarowego bez kolumny pomocniczej. CKE: mozliwe, ale wymaga wiecej pracy.

</details>

---

### Cwiczenie 19.10 (trudnosc: trudne, ~6 pkt)
**Zrodlo inspiracji**: Matura 2025 (Martianeum)
**Tagi**: `SUMIFS` `COUNTIFS` `tabela-krzyzowa` `pivot-reczny` `odwolania-mieszane` `kolumna-pomocnicza`

Firma transportowa rejestruje przejazdy:

| | A | B | C | D | E |
|---|---|---|---|---|---|
| 1 | **Nr** | **Skad** | **Dokad** | **Typ** | **Koszt** |
| 2 | 1 | Warszawa | Kraków | Towarowy | 2500 |
| 3 | 2 | Kraków | Wrocław | Osobowy | 800 |
| 4 | 3 | Warszawa | Gdańsk | Towarowy | 3200 |
| 5 | 4 | Gdańsk | Kraków | Osobowy | 1100 |
| 6 | 5 | Wrocław | Warszawa | Towarowy | 2800 |
| 7 | 6 | Kraków | Gdańsk | Osobowy | 950 |
| 8 | 7 | Warszawa | Wrocław | Towarowy | 2100 |
| 9 | 8 | Gdańsk | Warszawa | Osobowy | 1300 |
| 10 | 9 | Kraków | Warszawa | Towarowy | 2600 |
| 11 | 10 | Wrocław | Kraków | Osobowy | 750 |
| 12 | 11 | Warszawa | Kraków | Osobowy | 900 |
| 13 | 12 | Kraków | Wrocław | Towarowy | 1900 |

**Polecenie**:
1. Utworz tabele krzyzowa: miasto poczatkowe (wiersz) × typ przejazdu (kolumna) z SUMA kosztow.
2. Dodaj wiersz i kolumne "Razem".
3. Utworz druga tabele krzyzowa z LICZBA przejazdow (zamiast sumy kosztow).
4. Oblicz sredni koszt przejazdu towarowego i osobowego.
5. Z ktorego miasta wyjezdzalo najwiecej przejazdow towarowych?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dwie tabele krzyzowe — jedna z SUMIFS (koszty), druga z COUNTIFS (liczba). Ten sam wzorzec odwolan.
2. **Podejscie**: Tabela 1: SUMIFS(E;B;miasto;D;typ). Tabela 2: COUNTIFS(B;miasto;D;typ).
3. **Kluczowy krok**: Sredni koszt = suma kosztow (z tabeli 1) / liczba przejazdow (z tabeli 2).

</details>

<details>
<summary>Odpowiedz</summary>

**Formuly dla tabeli kosztow (G2):**
```
G2: =SUMIFS($E$2:$E$13;$B$2:$B$13;$F2;$D$2:$D$13;G$1)
```

**Formuly dla tabeli liczby (M2):**
```
M2: =COUNTIFS($B$2:$B$13;$L2;$D$2:$D$13;M$1)
```

**Tabela 1 — Suma kosztow:**

| Miasto | Towarowy | Osobowy | Razem |
|--------|----------|---------|-------|
| Warszawa | 2500+3200+2100 = 7800 | 900 | 8700 |
| Kraków | 2600+1900 = 4500 | 800+950 = 1750 | 6250 |
| Gdańsk | — (0) | 1100+1300 = 2400 | 2400 |
| Wrocław | 2800 | 750 | 3550 |
| **Razem** | **15100** | **5800** | **20900** |

Weryfikacja: SUM(E2:E13) = 2500+800+3200+1100+2800+950+2100+1300+2600+750+900+1900 = 20900 ✓

**Tabela 2 — Liczba przejazdow:**

| Miasto | Towarowy | Osobowy | Razem |
|--------|----------|---------|-------|
| Warszawa | 3 | 1 | 4 |
| Kraków | 2 | 2 | 4 |
| Gdańsk | 0 | 2 | 2 |
| Wrocław | 1 | 1 | 2 |
| **Razem** | **6** | **6** | **12** |

**Sredni koszt przejazdu:**
- Towarowy: 15100 / 6 = **2516.67 zl**
- Osobowy: 5800 / 6 = **966.67 zl**

**Najwiecej przejazdow towarowych:**
**Warszawa** — 3 przejazdy towarowe

**Wyjasnienie**: Dwie tabele krzyzowe (SUMIFS i COUNTIFS) uzywaja tego samego wzorca odwolan mieszanych — rozni sie tylko funkcja. Sredni koszt obliczamy jako iloraz odpowiednich komorek z obu tabel. Na maturze tabele krzyzowe pojawiaja sie przy analizie danych wielowymiarowych.
</details>

<details>
<summary>Typowe bledy</summary>

- **Pominecie zerowych wartosci**: Gdańsk ma 0 przejazdow towarowych — SUMIFS zwroci 0, nie blad. Trzeba uwzglednic to w analizie. CKE: -1 pkt jesli wplywaja na srednia
- **Dzielenie przez zero**: Jesli miasto ma 0 przejazdow danego typu, srednia nie istnieje. Na szczescie tu kazdy typ ma 6 przejazdow. CKE: -1 pkt w ogolnosci
- **Pomieszanie kolumny B (Skad) z C (Dokad)**: Pytanie mowi o miescie poczatkowym — filtrujemy po B, nie C. CKE: -2 pkt

</details>

---

## Samoocena

| Poziom | Opis | Cwiczenia |
|--------|------|-----------|
| Podstawowy | Uzywam SUMIF do prostego grupowania i LEFT/RIGHT/MID do ekstrakcji tekstu | 19.1, 19.6 bez pomocy |
| Dobry | Tworze tabele krzyzowe z SUMIFS/COUNTIFS i poprawnymi odwolaniami mieszanymi | 19.2-19.4, 19.7-19.8 bez pomocy |
| Bardzo dobry | Lacze tabele z VLOOKUP/INDEX-MATCH i tworze tabele pivot z weryfikacja krzyzowa | 19.5, 19.9 bez pomocy |
| Doskonaly | Tworze wiele tabel krzyzowych (SUMIFS + COUNTIFS) i analizuje dane wielowymiarowe | 19.10 bez pomocy |

**Co dalej?**
- Jesli masz trudnosci z odwolaniami mieszanymi: przejrzyj `cheatsheet_arkusz.md` (sekcja Adresowanie $)
- Jesli opanowales srednie: przejdz do cwiczen 15 (Agregacja warunkowa) — utrwalisz SUMIFS/COUNTIFS
- Jesli zrobiles wszystkie 10: sprobuj cwiczen 16 (Symulacja) — bardziej zlozonego zastosowania arkusza
