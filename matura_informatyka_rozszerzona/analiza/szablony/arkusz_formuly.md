# Arkusz Kalkulacyjny — Sciagawka na Mature

**Czestotliwosc**: 100% matur (12/12 lat) | **Punkty**: ~10 pkt/rok (112 pkt lacznie w 5 typach zadan)
**Narzedzie**: Excel lub LibreOffice Calc | **Czas**: ~40-50 min na egzaminie

---

## 1. Odniesienia komorkowe ($) — KLUCZOWA SEKCJA

### Tabela 4 typow odniesien

| Typ | Zapis | Kopiowanie w prawo | Kopiowanie w dol | Kiedy uzywac |
|-----|-------|-------------------|------------------|--------------|
| **Bezwzgledne** | `$A$1` | $A$1 | $A$1 | Stala wartosc (wspolczynnik, prog, norma) |
| **Mieszane (kolumna)** | `$A1` | $A1 | $A2 | Etykiety wierszy w tabeli krzyzowej |
| **Mieszane (wiersz)** | `A$1` | B$1 | A$1 | Naglowki kolumn w tabeli krzyzowej |
| **Wzgledne** | `A1` | B1 | A2 | Kopiowanie zwyklych formul |

### Przyklad kopiowania

Formula w B2: `=A2*$E$1`

| | Kopiowanie do C2 (prawo) | Kopiowanie do B3 (dol) |
|---|---|---|
| Przed | `=A2*$E$1` | `=A2*$E$1` |
| Po | `=B2*$E$1` | `=A3*$E$1` |

`A2` jest wzgledne — przesuwa sie. `$E$1` jest bezwzgledne — zostaje.

### Rozszerzajacy zakres (suma narastajaca)

```
C2: =SUM($B$2:B2)
```

| Komorka | Formula po skopiowaniu | Efekt |
|---------|----------------------|-------|
| C2 | `=SUM($B$2:B2)` | Suma B2 |
| C3 | `=SUM($B$2:B3)` | Suma B2:B3 |
| C4 | `=SUM($B$2:B4)` | Suma B2:B4 |
| C5 | `=SUM($B$2:B5)` | Suma B2:B5 |

Poczatek zakotwiczony (`$B$2`), koniec wzgledny (`B2`) — zakres rosnie z kazdym wierszem. Uzywane do **sum kumulacyjnych** (np. matura 2024 — rabaty narastajace).

### Odniesienia w tabeli krzyzowej (pivot)

Formula w F2 kopiowana w prawo i w dol:
```
=SUMIFS($C$2:$C$21; $A$2:$A$21;$E2; $B$2:$B$21;F$1)
```

- `$E2` — kolumna zakotwiczona (etykieta wiersza), wiersz wzgledny
- `F$1` — wiersz zakotwiczony (naglowek kolumny), kolumna wzgledna
- `$C$2:$C$21` — zakresy danych z $ na obu wymiarach (calkowicie nieruchome — dane sa zawsze w tych samych kolumnach)

⚠️ **Wazne**: dla zakresow danych uzywaj $ rowniez **przed litera kolumny**. Inaczej po skopiowaniu w prawo (F2→G2→H2…) kolumny `C`, `A`, `B` przesuna sie odpowiednio na `D`, `B`, `C` — formula wskaze na puste lub niewlasciwe kolumny.

---

## 2. Agregacja warunkowa (38 pkt, 11/12 lat) — NAJWAZNIEJSZY

### COUNTIF — zliczanie z warunkiem

```
=COUNTIF(B2:B11;"Nabial")
```
Zlicza komorki w B2:B11 zawierajace tekst "Nabial".

### SUMIF — suma z warunkiem

```
=SUMIF(B2:B10;"Styczen";C2:C10)
```
**Kolejnosc**: SUMIF(**zakres_kryterium**; kryterium; **zakres_sumy**)

Sumuje wartosci z C, gdzie B = "Styczen".

### AVERAGEIF — srednia z warunkiem

```
=AVERAGEIF(B2:B11;"M";C2:C11)
```
**Kolejnosc**: AVERAGEIF(**zakres_kryterium**; kryterium; **zakres_sredniej**)

Srednia z C, gdzie B = "M".

### SUMIFS — suma z wieloma warunkami

```
=SUMIFS(D2:D13;B2:B13;"Polnoc";C2:C13;"Q3")
```
**Kolejnosc**: SUMIFS(**zakres_sumy**; zakres_kryt1; kryt1; zakres_kryt2; kryt2)

### COUNTIFS — zliczanie z wieloma warunkami

```
=COUNTIFS(B2:B11;"Marzec";C2:C11;">"&F4)
```

### PULAPKA: SUMIF vs SUMIFS — ROZNA KOLEJNOSC!

| Funkcja | Pierwszy argument | Kolejnosc |
|---------|-------------------|-----------|
| **SUMIF** | zakres_kryterium | kryt → zakres_sumy (suma na KONCU) |
| **SUMIFS** | **zakres_sumy** | zakres_kryt1 → kryt1 → ... (suma na POCZATKU) |

```
SUMIF(zakres_kryt; kryterium; zakres_sumy)      ← suma na koncu
SUMIFS(zakres_sumy; zakres_kryt1; kryt1; ...)    ← suma na poczatku
```

### PULAPKA: Operator w kryterium

Zly: `=COUNTIFS(C2:C11;">5")` (porownuje z liczba 5, nie z komorka)
Dobry: `=COUNTIFS(C2:C11;">"&F4)` (laczy operator `">"` z wartoscia komorki F4)

Kryterium z komorka: `">"&F4` — operator `&` skleja tekst `">"` z wartoscia.

### Przyklad z matury 2024 (Hurtownia — jablka zimowe)

Dane: transakcje w kolumnach — NIP klienta (A), odmiana (B), typ Z/L (C), kg (D).
Cel: suma kg jablek zimowych (C="Z") dla kazdego klienta.

```
=SUMIFS($D$2:$D$2501; $A$2:$A$2501;G2; $C$2:$C$2501;"Z")
```
G2 = NIP klienta. Zakresy danych bezwzgledne ($), kryterium NIP wzgledne.

### Przyklad z matury 2025 (Martianeum — zawartosc >= 1%)

Suma masy ladunkow, ktore spelniaja prog zawartosci:
```
=SUMIFS(C2:C1001;D2:D1001;">="&1)
```
Prog `">="&1` oznacza "zawartosc >= 1%".

---

## 3. Symulacja krok-po-kroku (37 pkt, 12/12 lat)

### Wzorzec 1: Akumulator prosty

Saldo = poprzednie + zmiana.

```
C3: =C2+B3
```

| Krok | Operacja | Saldo |
|------|----------|-------|
| start | — | 1000 |
| 1 | +500 | 1500 |
| 2 | -200 | 1300 |

### Wzorzec 2: Wzrost ze stala (populacja, procent)

Populacja rosnie o 2.5% rocznie, wspolczynnik w D1.

```
B3: =B2*$D$1
```

| Rok | Populacja |
|-----|-----------|
| 2024 | 50000 |
| 2025 | 51250 |
| 2026 | 52531 |

`$D$1` — stala bezwzgledna (nie zmienia sie przy kopiowaniu).
Matura 2015: prognoza ludnosci `=B2*(1.01)^12`.

### Wzorzec 3: Stan z IF (alarm, kontrola)

Zapas = poprzedni + przyjecie - wydanie. Alarm gdy zapas < 50.

```
D3: =D2+B3-C3
E3: =IF(D3<50;"NISKI STAN";"")
```

| Dzien | Przyjecie | Wydanie | Zapas | Alarm |
|-------|-----------|---------|-------|-------|
| start | — | — | 200 | |
| 1 | 0 | 80 | 120 | |
| 3 | 0 | 70 | 40 | NISKI STAN |

### Wzorzec 4: Suma narastajaca z rozszerzajacym zakresem

Zakupy kumulacyjne — rabat zalezy od lacznej kwoty:

```
C2: =SUM($B$2:B2)
D2: =IF(C2>=2000;15%;IF(C2>=1000;10%;IF(C2>=500;5%;0%)))
E2: =B2*(1-D2)
```

| Zakup | Kwota | Kumulacja | Rabat | Po rabacie |
|-------|-------|-----------|-------|------------|
| 1 | 350 | 350 | 0% | 350.00 |
| 2 | 280 | 630 | 5% | 266.00 |
| 3 | 420 | 1050 | 10% | 378.00 |
| 6 | 300 | 2100 | 15% | 255.00 |

Matura 2024: rabaty hurtowe (5 gr/kg przy 15000+ kg, 10 gr/kg przy 20000+ kg).

### Wzorzec 5: Akumulator z progiem (magazyn + transport)

Magazyn gromadzi rude. Gdy >= 100 ton, wysyla transport 100 ton.

```
C2: =B2                           (dzien 1, startowy zapas = 0)
C3: =E2+B3                        (od dnia 2)
D2: =IF(C2>=100;"TAK";"NIE")
E2: =IF(D2="TAK";C2-100;C2)
```

| Dzien | Wydobycie | Zapas przed | Transport? | Zapas po |
|-------|-----------|-------------|-----------|----------|
| 1 | 35 | 35 | NIE | 35 |
| 2 | 42 | 77 | NIE | 77 |
| 3 | 28 | 105 | TAK | 5 |
| 4 | 50 | 55 | NIE | 55 |

Matura 2025: Martianeum — wyslij 100 kg na orbite gdy magazyn >= 100 kg.

Zliczanie transportow: `=COUNTIF(D2:D11;"TAK")`

### Matura 2023 (Konfitury): Symulacja z wyborem pary

Codziennie produkuj konfitury z dwoch owocow o najwiekszej ilosci (proporcja 1:1).
- Kolumny pomocnicze: zapas po produkcji = zapas wczora + dostawa - zuzycie
- Zuzycie = MIN(dwa_najwieksze_zapasy) dla kazdego z dwoch wybranych owocow
- Trzeci owoc (najnizsza ilosc) czeka w chlodni

---

## 4. Wykresy (25 pkt, 10/12 lat)

### 6 typow wykresow — kiedy uzywac

| Typ wykresu | Kiedy uzywac | Przyklad z matury |
|-------------|-------------|-------------------|
| **Kolumnowy** (slupkowy) | Jedna seria, kategorie na osi X | 2014/4d: przychody wg pol roku |
| **Kolowy** (pie) | Czesci calosci (udzialy %) | 2017: struktura sprzedazy cukru |
| **Kolumnowy grupowany** | Kilka serii obok siebie, porownanie | 2023/6.1: maliny/truskawki/porzeczki wg miesiecy |
| **Kolumnowy skumulowany** | Czesci skladowe calosci w czasie | 2025/6.4: przewozy z 30 obszarow wg lat |
| **Liniowy** | Trend w czasie | dane pogodowe, temperatury |
| **Kombinowany** (combo) | Rozne skale/jednostki | temperatura (linia) + opady (slupki) |

### Checklist tworzenia wykresu (4 elementy = 1 pkt)

1. **Tytul wykresu** — opisowy (np. "Miesieczna sprzedaz produktow")
2. **Opis osi X** — co oznaczaja kategorie (np. "Miesiac", "Obszar")
3. **Opis osi Y** — jednostka (np. "Sprzedaz (tys. zl)", "Liczba kg")
4. **Legenda** — wymagana przy >= 2 seriach danych

### Wykres skumulowany (stacked) vs grupowany (clustered)

| Cecha | Skumulowany | Grupowany |
|-------|-------------|-----------|
| Uklad | Warstwy jedna na drugiej | Slupki obok siebie |
| Wysokosc slupka | = suma wszystkich serii | = wartosc jednej serii |
| Czytanie | Calkosc + sklad | Porownanie wartosci |
| Matura 2025 | 30 obszarow x 6 lat | — |

### Tworzenie wykresu — kroki

1. Zaznacz dane RAZEM z naglowkami (etykietami)
2. Wstaw > Wykres > wybierz typ
3. Dodaj tytul, opisy osi, legende
4. Sprawdz, czy serie danych sa poprawnie przypisane
5. Jezeli combo — kliknij serie > zmien typ (linia/slupek) > przesuń na os dodatkowa

---

## 5. Agregacja podstawowa (9 pkt, 3/12 lat)

### Funkcje

| Funkcja | Dzialanie | Przyklad |
|---------|-----------|---------|
| `SUM(B2:B9)` | Suma | 110 |
| `AVERAGE(B2:B9)` | Srednia arytmetyczna | 13.75 |
| `MAX(B2:B9)` | Wartosc maksymalna | 19 |
| `MIN(B2:B9)` | Wartosc minimalna | 8 |
| `COUNT(B2:B9)` | Zlicza komorki z **liczbami** | 8 |
| `COUNTA(B2:B9)` | Zlicza komorki **niepuste** (tekst + liczby) | 8 |
| `COUNTBLANK(B2:B9)` | Zlicza komorki **puste** | 2 |
| `RANK(F2;$F$2:$F$7;0)` | Pozycja w rankingu (0=malejaco) | 3 |
| `ABS(D2)` | Wartosc bezwzgledna | 5.3 |

### RANK — szczegoly

```
G2: =RANK(F2;$F$2:$F$7;0)
```
- Trzeci argument: `0` = malejaco (najwyzszy = 1), `1` = rosnaco
- Zakres `$F$2:$F$7` musi byc **bezwzgledny** (z $)
- Remisy: oba elementy dostaja ta sama pozycje, nastepna pozycja pominieta (30, 30 → ranga 4, 4, nastepna 6)

### Roznica COUNT vs COUNTA

| Komorka | COUNT | COUNTA |
|---------|-------|--------|
| 42 | TAK | TAK |
| "Tak" | NIE | TAK |
| (pusta) | NIE | NIE |

COUNT liczy tylko liczby. COUNTA liczy wszystko niepuste. Na maturze czesto trzeba COUNTA do tekstu.

---

## 6. Transformacja danych (3 pkt, 2/12 lat)

### Tabela krzyzowa z SUMIFS

Dane plaska (wiersz na kazda transakcje) → tabela: kategorie (wiersze) x kwartaly (kolumny).

```
F2: =SUMIFS($C$2:$C$21; $A$2:$A$21;$E2; $B$2:$B$21;F$1)
```

Kopiowana w prawo (Q1→Q2→Q3→Q4) i w dol (Elektronika→Odziez→Zywnosc).

### Tabela krzyzowa z COUNTIFS (zliczanie)

```
E2: =COUNTIFS($A$2:$A$16;$D2; $B$2:$B$16;E$1)
```

Zlicza ile razy pytanie P1 otrzymalo odpowiedz A. Te same zasady $ co wyzej.

### Kolumna pomocnicza — grupowanie po okresach

7-dniowe okresy (matura 2025 — Martianeum):

```
C2: =INT((A2-1)/7)+1
```

| Dzien | INT((dzien-1)/7)+1 | Okres |
|-------|--------------------|-------|
| 1-7 | INT(0..6 / 7)+1 | 1 |
| 8-14 | INT(7..13 / 7)+1 | 2 |
| 15-21 | INT(14..20 / 7)+1 | 3 |

Potem SUMIF po kolumnie pomocniczej:
```
=SUMIF(C2:C22;E2;B2:B22)
```

### Grupowanie po miesiacach (z dat)

Jezeli dane maja daty, wyodrebnij miesiac:
```
=MONTH(A2)
```
lub rok:
```
=YEAR(A2)
```

---

## 7. Funkcje dodatkowe (luki z egzaminow 2023-2025)

### VLOOKUP — wyszukiwanie w tabeli (cennik, taryfikator)

```
=VLOOKUP(B2;$G$2:$H$10;2;0)
```

| Argument | Znaczenie |
|----------|-----------|
| `B2` | Wartosc szukana (np. nazwa odmiany) |
| `$G$2:$H$10` | Tabela (bezwzgledna!) — szukana w 1. kolumnie |
| `2` | Numer kolumny wyniku (np. cena w 2. kolumnie) |
| `0` | Dokladne dopasowanie (ZAWSZE uzywaj 0 na maturze) |

Matura 2024: VLOOKUP do cennika jablek — cena za kg na podstawie odmiany.

### INDEX-MATCH — alternatywa VLOOKUP (bardziej elastyczna)

```
=INDEX(C2:C10;MATCH(B2;A2:A10;0))
```

| Czesc | Dzialanie |
|-------|-----------|
| `MATCH(B2;A2:A10;0)` | Znajdz pozycje B2 w zakresie A2:A10 |
| `INDEX(C2:C10; ...)` | Zwroc wartosc z tej pozycji w zakresie C |

Przewaga nad VLOOKUP: kolumna wynikowa moze byc na lewo od szukanej.

Matura 2015: najludniejsze wojewodztwo `=INDEX(A:A;MATCH(MAX(D:D);D:D;0))`

### Daty

| Funkcja | Dzialanie | Przyklad |
|---------|-----------|---------|
| `MONTH(A2)` | Numer miesiaca (1-12) | 5 |
| `YEAR(A2)` | Rok | 2024 |
| `DAY(A2)` | Dzien miesiaca | 15 |
| `MIN(A2:A100)` | Najwczesniejsza data | 2033-03-03 |
| `MAX(A2:A100)` | Najpozniejsza data | 2038-09-01 |

Daty w Excelu to liczby — mozna je porownywac, odejmowac, uzywac MIN/MAX.

Matura 2025: `MIN(data_pomiaru)` i `MAX(data_pomiaru)` dla kazdego lazika.

### Tekstowe

| Funkcja | Dzialanie | Przyklad |
|---------|-----------|---------|
| `LEFT(A2;3)` | Pierwsze 3 znaki | "ABC" z "ABCDEF" |
| `RIGHT(A2;2)` | Ostatnie 2 znaki | "EF" z "ABCDEF" |
| `MID(A2;2;3)` | 3 znaki od pozycji 2 | "BCD" z "ABCDEF" |
| `LEN(A2)` | Dlugosc tekstu | 6 |

### Procenty i zaokraglanie

```
=B2/SUM($B$2:$B$6)*100          udzial procentowy
=ROUND(B2;2)                     zaokraglenie do 2 miejsc
=ROUND(B2*C2;2)                  cena × ilosc, zaokraglona
```

`ROUND(wartosc; liczba_miejsc)` — zaokragla (nie obcina!). Na maturze czesto wymagane "z dokladnoscia do 2 miejsc po przecinku".

### IF i zagniezdzony IF

```
=IF(warunek; wartosc_jesli_tak; wartosc_jesli_nie)

Zagniezdzony (progi rabatowe):
=IF(C2>=2000;15%;IF(C2>=1000;10%;IF(C2>=500;5%;0%)))
```

Sprawdzaj progi **od najwiekszego** — pierwszy speliony warunek konczy ewaluacje.

### AND / OR w warunkach

```
=IF(AND(G2;H2);"OK";"ODRZUT")        oba warunki musza byc TRUE
=IF(OR(G2;H2);"PRZESZEDL";"NIE")     wystarczy jeden TRUE
```

---

## 8. Typowe pulapki (8 pozycji)

### 1. Brak $ w odniesieniach

| Blednie | Poprawnie | Problem |
|---------|-----------|---------|
| `=B2*D1` | `=B2*$D$1` | Po skopiowaniu D1 staje sie D2, D3... |
| `=SUMIF(B:B;E2;C:C)` | `=SUMIF($B:$B;E2;$C:$C)` | Zakresy danych powinny byc nieruchome |

### 2. SUMIF vs SUMIFS — kolejnosc argumentow

```
ZLE:  =SUMIFS(A:A;"kryt";C:C)       ← kolejnosc SUMIF w SUMIFS
OK:   =SUMIFS(C:C;A:A;"kryt")       ← SUMIFS: suma PIERWSZA
```

### 3. Srednia z zerami / pustymi komorkami

```
AVERAGE(1;0;3;0;5) = 9/5 = 1.8      ← zera LICZONE
AVERAGE(1;;3;;5) = 9/3 = 3.0         ← puste POMINIETE
```

Jezeli zera maja byc pominiete: `=AVERAGEIF(B2:B11;"<>0")`

### 4. Operator w kryterium — brak & z komorka

```
ZLE:  =COUNTIF(C:C;">F4")           ← szuka tekstu ">F4"
OK:   =COUNTIF(C:C;">"&F4)          ← laczy ">" z wartoscia F4
```

### 5. Zly typ wykresu

| Dane | Zly wykres | Poprawny wykres |
|------|-----------|-----------------|
| Czesci calosci | Kolumnowy | **Kolowy** |
| Sklad + calkosc w czasie | Grupowany | **Skumulowany** |
| 2 rozne jednostki | 1 os Y | **Combo z 2 osiami Y** |

### 6. Brak elementow wykresu (traci 1 pkt)

Zawsze dodaj: **tytul**, **opis osi X**, **opis osi Y**, **legenda** (przy >= 2 seriach).

### 7. Dzielenie przez 0

```
=IF(B2=0;0;A2/B2)                   zabezpieczenie przed #DIV/0!
=IFERROR(A2/B2;0)                   alternatywa — zwraca 0 przy bledzie
```

### 8. Rozszerzajacy zakres — brak $ na poczatku

```
ZLE:  =SUM(B2:B2)    → po skopiowaniu do wiersza 5: =SUM(B5:B5)
OK:   =SUM($B$2:B2)  → po skopiowaniu do wiersza 5: =SUM($B$2:B5)
```

---

## 9. Schemat rozwiazywania zadania arkuszowego

### Krok 1: Wczytaj dane

- Otworz plik .txt w arkuszu (Dane > Z pliku tekstowego)
- Sprawdz separator (tab, srednik, spacja)
- Sprawdz separator dziesietny (przecinek w danych polskich!)
- Zweryfikuj na danych przykladowych

### Krok 2: Rozpoznaj typ zadania

| Sygnal w tresci | Typ | Sekcja |
|-----------------|-----|--------|
| "ile/suma/srednia... spelniajacych warunek" | Agregacja warunkowa | §2 |
| "symuluj/modeluj/codziennie/krok po kroku" | Symulacja | §3 |
| "utworz wykres/przedstaw graficznie" | Wykres | §4 |
| "policz/zsumuj wszystkie" (bez warunku) | Agregacja podstawowa | §5 |
| "zestawienie/tabela krzyzowa/grupuj po" | Transformacja | §6 |

### Krok 3: Wybierz formule

| Cel | Formula |
|-----|---------|
| Zlicz z warunkiem | `COUNTIF` / `COUNTIFS` |
| Zsumuj z warunkiem | `SUMIF` / `SUMIFS` |
| Srednia z warunkiem | `AVERAGEIF` |
| Wyszukaj wartosc z tabeli | `VLOOKUP` / `INDEX-MATCH` |
| Warunek logiczny | `IF` / `IF(AND(...))` |
| Suma narastajaca | `SUM($B$2:B2)` |
| Grupuj po okresach | `INT((A2-1)/N)+1` + `SUMIF` |
| Wyodrebnij miesiac | `MONTH(data)` |

### Krok 4: Sprawdz $

Przed skopiowaniem formuly w dol/prawo, zadaj sobie pytania:

1. **Czy zakres danych powinien sie zmieniac?** NIE → dodaj `$` (np. `$B$2:$B$100`)
2. **Czy stala/prog/wspolczynnik powinien sie zmieniac?** NIE → dodaj `$` (np. `$D$1`)
3. **Czy etykieta wiersza powinna sie zmieniac w prawo?** NIE → `$E2`
4. **Czy naglowek kolumny powinien sie zmieniac w dol?** NIE → `F$1`

### Krok 5: Zweryfikuj

- Sprawdz na danych przykladowych (pliki `_przyklad.txt`)
- Porownaj z oczekiwanymi wynikami z tresci zadania
- Jezeli symulacja — przelicz recznie pierwsze 3-4 wiersze
- Jezeli SUMIF — zweryfikuj, ze suma czesci = suma calosci

---

## Dodatek: Przyklady z matur 2023-2025

### Matura 2023 — Konfitury owocowe (10 pkt)

| Podzadanie | Typ | Formula |
|-----------|-----|---------|
| 6.1 | Wykres | SUMIF po miesiacu + wykres kolumnowy |
| 6.2 | Agregacja war. | Zlicz dni gdzie porzeczki = MAX(3 owoce) |
| 6.3 | Symulacja | Wybierz 2 owoce z max, produkuj MIN(a,b) kg, reszta na jutro |
| 6.4 | Symulacja | Jak 6.3 ale zlicz kg konfitur (MIN z dwoch najwiekszych) |

### Matura 2024 — Hurtownia jablek (10 pkt)

| Podzadanie | Typ | Formula |
|-----------|-----|---------|
| 7.1 | Agregacja war. | SUMIFS: kg jablek zimowych per klient |
| 7.2 | Agregacja war. | VLOOKUP(cena) * kg, SUMIF po odmianie |
| 7.3 | Wykres | SUMIFS per miesiac + MAX odmiana + wykres kolumnowy |
| 7.4 | Symulacja | SUM($B$2:B2) kumulacyjny + IF progi rabatowe |

### Matura 2025 — Martianeum (11 pkt)

| Podzadanie | Typ | Formula |
|-----------|-----|---------|
| 6.1 | Agregacja war. | SUMIF (masa), SUMIFS z progiem >= 1% |
| 6.2 | Agregacja war. | AVERAGEIF per obszar + MIN |
| 6.3 | Transformacja | INT((data-start)/7)+1 + SUMIF po okresie |
| 6.4 | Wykres | COUNTIFS per obszar per rok + wykres **skumulowany** |
| 6.5 | Symulacja | Akumulator + IF(zapas>=100; wyslij 100) |
