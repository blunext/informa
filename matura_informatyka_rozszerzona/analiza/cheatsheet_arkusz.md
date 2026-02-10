# Cheatsheet: Arkusz kalkulacyjny

## Odniesienia $ (4 typy)

| Zapis | Kopiuj w prawo | Kopiuj w dol | Uzycie |
|-------|----------------|--------------|--------|
| `$A$1` | $A$1 | $A$1 | Stala (prog, wspolczynnik) |
| `$A1` | $A1 | $A2 | Etykieta wiersza (pivot) |
| `A$1` | B$1 | A$1 | Naglowek kolumny (pivot) |
| `A1` | B1 | A2 | Zwykla formula |

> UWAGA: "$" od strony etykiet — `$E2` (kolumna stala), `F$1` (wiersz staly)

---

## Tabela krzyzowa (pivot) — szablon

```
=SUMIFS(C$2:C$21; A$2:A$21;$E2; B$2:B$21;F$1)
```

`$E2` = etykieta wiersza | `F$1` = naglowek kolumny | zakresy z `$` na wierszach

---

## Suma narastajaca (kumulacyjna)

```
C2: =SUM($B$2:B2)    -> C3: =SUM($B$2:B3)    -> C4: =SUM($B$2:B4)
```

Poczatek zakotwiczony `$B$2`, koniec wzgledny — zakres rosnie.

---

## Formuly warunkowe

| Funkcja | Skladnia | Kolejnosc |
|---------|----------|-----------|
| `COUNTIF` | `COUNTIF(zakres; kryterium)` | — |
| `COUNTIFS` | `COUNTIFS(zakr1; kryt1; zakr2; kryt2)` | — |
| **`SUMIF`** | `SUMIF(zakr_kryt; kryt; zakr_sumy)` | **suma na KONCU** |
| **`SUMIFS`** | `SUMIFS(zakr_sumy; zakr_kryt1; kryt1; ...)` | **suma na POCZATKU** |
| `AVERAGEIF` | `AVERAGEIF(zakr_kryt; kryt; zakr_sredniej)` | jak SUMIF |

> UWAGA: SUMIF i SUMIFS maja ODWROTNA kolejnosc argumentow!

Kryterium z komorka: `">"&F4` (operator `&` skleja `">"` z wartoscia F4)

---

## Wyszukiwanie

```
=VLOOKUP(szukana; $tabela; nr_kolumny; 0)       0 = dokladne!
=INDEX(zakres_wynikow; MATCH(szukana; zakres_szukania; 0))
```

VLOOKUP: szukana musi byc w 1. kolumnie tabeli.
INDEX-MATCH: kolumna wynikowa moze byc GDZIEKOLWIEK.

---

## IF / zagniezdzony IF

```
=IF(warunek; jesli_tak; jesli_nie)
=IF(C2>=2000;15%;IF(C2>=1000;10%;IF(C2>=500;5%;0%)))
```

Progi od NAJWIEKSZEGO — pierwszy spelniony konczy ewaluacje.

---

## Symulacja — wzorce

```
Akumulator:         C3 = C2 + B3                (saldo = poprz + zmiana)
Wzrost procentowy:  B3 = B2 * $D$1              (stala bezwzgledna)
Magazyn z progiem:  D2 = IF(C2>=100;"TAK";"NIE")
                    E2 = IF(D2="TAK";C2-100;C2)
```

---

## Inne przydatne

| Funkcja | Dzialanie |
|---------|-----------|
| `MAX` / `MIN` | Wartosc ekstr. |
| `AVERAGE` | Srednia (puste pomija, zera liczy!) |
| `COUNT` / `COUNTA` | Liczby / niepuste |
| `RANK(F2;$F$2:$F$7;0)` | Pozycja (0=malej.) |
| `MOD(A2;B2)` | Reszta |
| `INT(A2)` | Czesc calkowita |
| `LEFT`/`RIGHT`/`MID`/`LEN` | Tekstowe |
| `MONTH(data)`/`YEAR(data)` | Z dat |
| `ROUND(wart;2)` | Zaokraglenie |
| `AND()`/`OR()` | Logiczne (w IF) |

---

## Wykres — jaki typ?

| Dane | Typ wykresu |
|------|-------------|
| Jedna seria, kategorie | **Kolumnowy** (slupkowy) |
| Czesci calosci (udzialy %) | **Kolowy** (pie) |
| Porownanie kilku serii | **Kolumnowy grupowany** |
| Sklad + calkosc w czasie | **Kolumnowy skumulowany** |
| Trend w czasie | **Liniowy** |

Checklist: **tytul** + **opis osi X** + **opis osi Y** + **legenda** (przy >=2 seriach)

---

## Pulapki

1. **Brak $** — `=B2*D1` po skopiowaniu zmieni D1 na D2 -> `=B2*$D$1`
2. **SUMIF vs SUMIFS** — rozna kolejnosc argumentow (suma na koncu vs poczatku)
3. **AVERAGE z zerami** — zera sa LICZONE, puste POMINIETE
4. **Kryterium z komorka** — `">"&F4` nie `">F4"` (szuka tekstu!)
5. **Rozszerzajacy zakres** — `=SUM($B$2:B2)` nie `=SUM(B2:B2)`
6. **Zly typ wykresu** — czesci calosci = kolowy, nie kolumnowy
7. **Brak elementow wykresu** — tytul, osie, legenda = 1 pkt
8. **Dzielenie przez 0** — `=IF(B2=0;0;A2/B2)` lub `=IFERROR(A2/B2;0)`
