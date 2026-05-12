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
=SUMA.WARUNKÓW(C$2:C$21; A$2:A$21;$E2; B$2:B$21;F$1)
```

`$E2` = etykieta wiersza | `F$1` = naglowek kolumny | zakresy z `$` na wierszach

---

## Suma narastajaca (kumulacyjna)

```
C2: =SUMA($B$2:B2)    -> C3: =SUMA($B$2:B3)    -> C4: =SUMA($B$2:B4)
```

Poczatek zakotwiczony `$B$2`, koniec wzgledny — zakres rosnie.

---

## Formuly warunkowe

| Funkcja | Skladnia | Kolejnosc |
|---------|----------|-----------|
| `COUNTIF` | `LICZ.JEŻELI(zakres; kryterium)` | — |
| `COUNTIFS` | `LICZ.WARUNKI(zakr1; kryt1; zakr2; kryt2)` | — |
| **`SUMIF`** | `SUMA.JEŻELI(zakr_kryt; kryt; zakr_sumy)` | **suma na KONCU** |
| **`SUMIFS`** | `SUMA.WARUNKÓW(zakr_sumy; zakr_kryt1; kryt1; ...)` | **suma na POCZATKU** |
| `AVERAGEIF` | `ŚREDNIA.JEŻELI(zakr_kryt; kryt; zakr_sredniej)` | jak SUMIF |

> UWAGA: SUMIF i SUMIFS maja ODWROTNA kolejnosc argumentow!

Kryterium z komorka: `">"&F4` (operator `&` skleja `">"` z wartoscia F4)

---

## Wyszukiwanie

```
=WYSZUKAJ.PIONOWO(szukana; $tabela; nr_kolumny; 0)       0 = dokladne!
=INDEKS(zakres_wynikow; PODAJ.POZYCJĘ(szukana; zakres_szukania; 0))
```

VLOOKUP: szukana musi byc w 1. kolumnie tabeli.
INDEX-MATCH: kolumna wynikowa moze byc GDZIEKOLWIEK.

---

## IF / zagniezdzony IF

```
=JEŻELI(warunek; jesli_tak; jesli_nie)
=JEŻELI(C2>=2000;15%;JEŻELI(C2>=1000;10%;JEŻELI(C2>=500;5%;0%)))
```

Progi od NAJWIEKSZEGO — pierwszy spelniony konczy ewaluacje.

---

## Symulacja — wzorce

```
Akumulator:         C3 = C2 + B3                (saldo = poprz + zmiana)
Wzrost procentowy:  B3 = B2 * $D$1              (stala bezwzgledna)
Magazyn z progiem:  D2 = JEŻELI(C2>=100;"TAK";"NIE")
                    E2 = JEŻELI(D2="TAK";C2-100;C2)
```

---

## Inne przydatne

| Funkcja | Dzialanie |
|---------|-----------|
| `MAX` / `MIN` | Wartosc ekstr. |
| `AVERAGE` | Srednia (puste pomija, zera liczy!) |
| `COUNT` / `COUNTA` | Liczby / niepuste |
| `POZYCJA(F2;$F$2:$F$7;0)` | Pozycja (0=malej.) |
| `MOD(A2;B2)` | Reszta |
| `ZAOKR.DO.CAŁK(A2)` | Czesc calkowita |
| `LEFT`/`RIGHT`/`MID`/`LEN` | Tekstowe |
| `MIESIĄC(data)`/`ROK(data)` | Z dat |
| `ZAOKR(wart;2)` | Zaokraglenie |
| `ORAZ()`/`LUB()` | Logiczne (w IF) |

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
5. **Rozszerzajacy zakres** — `=SUMA($B$2:B2)` nie `=SUMA(B2:B2)`
6. **Zly typ wykresu** — czesci calosci = kolowy, nie kolumnowy
7. **Brak elementow wykresu** — tytul, osie, legenda = 1 pkt
8. **Dzielenie przez 0** — `=JEŻELI(B2=0;0;A2/B2)` lub `=JEŻELI.BŁĄD(A2/B2;0)`

---

## Pułapki polskich nazw funkcji — uważaj!

| Funkcja | Polska nazwa | Pułapka |
|---|---|---|
| `ROUNDUP` | `ZAOKR.GÓRA` | NIE `ZAOKR.W.GÓRĘ` (to **CEILING** — zaokrąglanie do wielokrotności) |
| `ROUNDDOWN` | `ZAOKR.DÓŁ` | NIE `ZAOKR.W.DÓŁ` (to **FLOOR**) |
| `CONCATENATE` | `ZŁĄCZ.TEKSTY` (z "Y") | `ZŁĄCZ.TEKST` (bez Y) to nowsza `CONCAT` |
| `SUBSTITUTE` vs `REPLACE` | `PODSTAW` vs `ZASTĄP` | `PODSTAW` = po tekście, `ZASTĄP` = po pozycji znaków |
| `MAX` / `MIN` / `MOD` | bez zmian | Excel PL używa tych samych nazw — NIE tłumaczyć |
| Separator argumentów | `;` (średnik) | Polski locale używa średnika, NIE przecinka |
| Separator dziesiętny | `,` (przecinek) | Polski locale używa przecinka (np. `3,14` zamiast `3.14`) |
