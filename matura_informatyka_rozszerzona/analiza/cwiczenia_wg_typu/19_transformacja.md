# 19. Transformacja danych w arkuszu kalkulacyjnym

Typ zadania: **arkusz_transformacja**
Czestotliwosc: 2/11 lat | Laczna punktacja: 3 pkt
Kategoria: ARKUSZ

---

### Cwiczenie 19.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna

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

---

### Cwiczenie 19.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: ogolna

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

---

### Cwiczenie 19.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 (Martianeum)

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

---

### Cwiczenie 19.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: ogolna

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

---

### Cwiczenie 19.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)

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
- Suma wszystkich wartosci zrodlowych: =SUM(C2:C21) = 15000+8000+12000+18000+9500+11000+13000+11000+13500+20000+7500+14000+22000+10000+10500+16000+8500+12500+19000+9000 = 260000 ✓

**Wyjasnienie odniesien w SUMIFS:**
- `$E2` — zakotwiczona kolumna E (kategoria), wiersz wzgledny → kopiowanie w dol zmienia E2→E3→E4
- `F$1` — zakotwiczony wiersz 1 (kwartal), kolumna wzgledna → kopiowanie w prawo zmienia F→G→H→I
- `C$2:C$21`, `A$2:A$21`, `B$2:B$21` — zakresy danych z $ na wierszach, nie zmieniaja sie

To jest reczna tabela przestawna (pivot table) zbudowana z formul. Na maturze dane zrodlowe czesto maja wiele wierszy z ta sama kombinacja kategorii — SUMIFS poprawnie je sumuje. Kluczowe jest prawidlowe uzycie $ w odniesieniach, aby jedna formula dzialala po skopiowaniu w cala tabele.
</details>
