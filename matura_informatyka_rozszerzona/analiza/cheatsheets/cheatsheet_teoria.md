# Cheatsheet: Teoria (pseudokod + algorytmy)

## Slowniczek CKE -> C++

| CKE | C++ | | CKE | C++ |
|-----|-----|-|-----|-----|
| `:=` | `=` | | `=` | `==` |
| `<>` | `!=` | | `mod` | `%` |
| `div` | `/` (int) | | `AND`/`OR` | `&&`/`\|\|` |
| `dopoki` | `while` | | `jezeli` | `if` |
| `dla i:=0,1,...,n-1` | `for(i=0;i<n;i++)` | | `zwroc` | `return` |
| `stos.pop()` | `x=s.top(); s.pop();` | | `wezel.L`/`.R` | `->L`/`->R` |

> UWAGA: Wciecia = blok kodu (odpowiednik `{ }` w C++). Kazda spacja sie liczy!

---

## 7 archetypow — rozpoznaj wzorzec

| # | Archetyp | Rozpoznanie | Kluczowa operacja |
|---|----------|-------------|-------------------|
| 1 | Petla po cyfrach | `mod 10` / `div 10` | `cyfra := n mod 10; n := n div 10` |
| 2 | Odwracanie/palindrom | `odwrocona * 10 + cyfra` | `odwr := odwr * 10 + cyfra` |
| 3 | Systemy liczbowe | `mod k` / `div k` lub Horner | dec->k: `mod k / div k`; k->dec: `w := w * k + cyfry[i]` |
| 4 | NWD Euklidesa | `a mod b`, zamiana | `dopoki b<>0: temp:=b; b:=a mod b; a:=temp` |
| 5 | Zliczanie z warunkiem | `jezeli ... licznik++` | `C[T[i]] := C[T[i]] + 1` |
| 6 | Budowanie wyniku | `wynik + cyfra * mnoznik` | mnoznik rosnie TYLKO gdy warunek! |
| 7 | Rekurencja/stos | `wezel.L` / `wezel.R` / `push` | stos LIFO: prawy PRZED lewym na stos |

---

## Drzewo decyzyjne — "widzisz X -> uzyj Y"

| Widzisz w zadaniu | Algorytm |
|---|---|
| "cyfry liczby", "suma cyfr" | Petla mod 10 / div 10 |
| "NWD", "dzielnik" | Euklides |
| "czy pierwsza" | Test do sqrt(n): `i*i <= n` |
| "rozklad na czynniki" | Dziel przez d=2,3,...; `d*d <= n` |
| "wiele l. pierwszych" | Sito Eratostenesa |
| "system dwojkowy/hex" | Konwersja mod k / Horner |
| "posortuj" | `sort` + komparator |
| "najdluzszy ciag" | current_len / max_len |
| "sciezka na planszy" | DP 2D: `dp[i][j] = dp[i-1][j] \|\| dp[i][j-1]` |

---

## TOP 10 zlozonosci (do testow P/F)

| Algorytm | Zlozonosc |
|----------|-----------|
| Przeszukiwanie binarne | O(log n) |
| Petla po tablicy / min / max | O(n) |
| Sortowanie (merge/quick) | O(n log n) |
| Bubble sort / selection sort | O(n^2) |
| Sito Eratostenesa | O(n log log n) |
| NWD Euklidesa | O(log min(a,b)) |
| Dzielniki do sqrt(n) | O(sqrt(n)) |
| Cyfry liczby (mod/div) | O(log n) = liczba cyfr |
| Dwie zagniezdzonce petle | O(n^2) |
| Wszystkie podzbiory | O(2^n) |

---

## 6 pulapek sledzenia

1. **Wciecia = blok** — linia bez wciecia jest POZA if/petla (CKE testuje to celowo)
2. **`:=` vs `=`** — `:=` to przypisanie, `=` to porownanie (w C++: `=` vs `==`)
3. **n = 0** — petla `dopoki n > 0` w ogole nie wejdzie; sprawdz osobno
4. **Mnoznik warunkowy** — w archetyp 6: `mnoznik` rosnie TYLKO gdy `jezeli` jest spelniony
5. **Stos LIFO** — prawy potomek push PRZED lewym (bo zdejmiemy lewego PIERWSZEGO)
6. **Rekurencja odwraca** — akcja PO wywolaniu rekurencyjnym = odwrotna kolejnosc

---

## Tabelka sledzenia — szablon

```
| Krok | zmienna_1 | zmienna_2 | warunek? | wynik |
|------|-----------|-----------|----------|-------|
| start| wart_pocz | wart_pocz |          |       |
| 1    | ...       | ...       | tak/nie  |       |
```

Zasady: kolumny = zmienne, wiersze = iteracje, warunek = osobna kolumna.
Dla rekurencji: rysuj drzewo wywolan (argument -> wartosc zwracana).
