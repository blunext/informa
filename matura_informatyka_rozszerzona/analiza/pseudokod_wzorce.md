# Pseudokod CKE — Konwencje i Wzorce

Sciagawka konwencji pseudokodu stosowanego przez CKE na maturze rozszerzonej z informatyki.
Oparta na analizie arkuszy 2014-2025.

---

## 1. Slowniczek CKE — pseudokod vs C++

| Pseudokod CKE | C++ | Znaczenie |
|---|---|---|
| `funkcja nazwa(arg)` | `typ nazwa(typ arg) {` | Definicja funkcji |
| `dopoki warunek:` | `while (warunek) {` | Petla while |
| `dla i := 0, 1, ..., n-1:` | `for (int i = 0; i < n; i++) {` | Petla for |
| `jezeli warunek:` | `if (warunek) {` | Warunek |
| `w przeciwnym razie:` | `} else {` | Galaz else |
| `zwroc wartosc` | `return wartosc;` | Zwracanie wyniku |
| `:=` | `=` | Przypisanie |
| `=` | `==` | Porownanie |
| `<>` | `!=` | Rozne |
| `mod` | `%` | Reszta z dzielenia |
| `div` | `/` (dla int) | Dzielenie calkowite |
| `PRAWDA` / `FALSZ` | `true` / `false` | Wartosci logiczne |
| `AND` / `OR` | `&&` / `\|\|` | Operatory logiczne |
| `NULL` | `nullptr` | Pusty wskaznik |
| `T[i]` | `T[i]` | Element tablicy |
| `stos.push(x)` | `stos.push(x)` | Wloz na stos |
| `stos.pop()` | `x = stos.top(); stos.pop();` | Zdejmij ze stosu |
| `wezel.L`, `wezel.R` | `wezel->L`, `wezel->R` | Potomek lewy/prawy |
| `wezel.w` | `wezel->w` | Wartosc wezla |

**Uwaga**: CKE uzywa polskich slow kluczowych. W pseudokodzie nie ma srednikow ani klamer — bloki wyznaczaja wciecia.

---

## 2. Struktura bloku

### Zasady formatowania

- **Wciecia** (zwykle 4 spacje) = blok kodu — odpowiednik `{ }` w C++
- **Dwukropek** na koncu linii otwierajacej blok: `dopoki ...:`, `jezeli ...:`, `dla ...:`, `funkcja ...:`
- Brak srednikow, brak klamer

### Szablony blokow

**Funkcja:**
```
funkcja nazwa(parametry)
    // cialo funkcji
    zwroc wynik
```

**Petla dopoki (while):**
```
dopoki warunek:
    // cialo petli
```

**Petla dla (for):**
```
dla i := 0, 1, ..., n-1:
    // cialo petli
```

**Warunek:**
```
jezeli warunek:
    // blok jezeli
w przeciwnym razie:
    // blok else
```

**Zagniezdzenie (wciecia sie sumuja):**
```
funkcja przyklad(n)
    dla i := 0, 1, ..., n-1:
        jezeli warunek(i):
            // 3 poziomy wciecia = 12 spacji
```

---

## 3. Siedem archetypow pseudokodu CKE

### Archetyp 1: Petla po cyfrach (mod 10 / div 10)

Najczestszy wzorzec — pojawia sie prawie co roku.

**Szablon:**
```
funkcja przetworzCyfry(n)
    wynik := 0
    dopoki n > 0:
        cyfra := n mod 10
        // przetwarzanie cyfry
        n := n div 10
    zwroc wynik
```

**Przyklad — zliczanie cyfr parzystych:**
```
funkcja parzyste_cyfry(n)
    licznik := 0
    dopoki n > 0:
        cyfra := n mod 10
        jezeli cyfra mod 2 = 0:
            licznik := licznik + 1
        n := n div 10
    zwroc licznik
```

**Tabelka sledzenia** dla `parzyste_cyfry(1024)`:

| Krok | n | cyfra | cyfra%2=0? | licznik |
|------|---|-------|------------|---------|
| 1 | 1024 | 4 | tak | 1 |
| 2 | 102 | 2 | tak | 2 |
| 3 | 10 | 0 | tak | 3 |
| 4 | 1 | 1 | nie | 3 |

Wynik: 3

**Zrodla maturalne**: 2024/2 (Cyfry), 2023/3 (cyfry pi), 2019/4, 2015/4

---

### Archetyp 2: Odwracanie liczby (palindrom)

**Szablon:**
```
funkcja odwroc(n)
    odwrocona := 0
    dopoki n > 0:
        cyfra := n mod 10
        odwrocona := odwrocona * 10 + cyfra
        n := n div 10
    zwroc odwrocona
```

**Przyklad — sprawdzanie palindromu:**
```
funkcja czyPalindrom(n)
    oryginal := n
    odwrocona := 0
    dopoki n > 0:
        cyfra := n mod 10
        odwrocona := odwrocona * 10 + cyfra
        n := n div 10
    jezeli odwrocona = oryginal:
        zwroc PRAWDA
    w przeciwnym razie:
        zwroc FALSZ
```

**Tabelka sledzenia** dla `odwroc(1234)`:

| Krok | n | cyfra | odwrocona |
|------|---|-------|-----------|
| 1 | 1234 | 4 | 4 |
| 2 | 123 | 3 | 43 |
| 3 | 12 | 2 | 432 |
| 4 | 1 | 1 | 4321 |

Wynik: 4321

---

### Archetyp 3: Konwersja systemu liczbowego

**Szablon (dec -> base k, wynik jako int):**
```
funkcja konwertuj(n, k)
    wynik := 0
    mnoznik := 1
    dopoki n > 0:
        reszta := n mod k
        wynik := wynik + reszta * mnoznik
        mnoznik := mnoznik * 10
        n := n div k
    zwroc wynik
```

**Tabelka sledzenia** dla `konwertuj(13, 2)`:

| Krok | n | n mod 2 | wynik | mnoznik |
|------|---|---------|-------|---------|
| 1 | 13 | 1 | 1 | 10 |
| 2 | 6 | 0 | 1 | 100 |
| 3 | 3 | 1 | 101 | 1000 |
| 4 | 1 | 1 | 1101 | 10000 |

Wynik: 1101 (= 13 w systemie dwojkowym)

**Schemat Hornera (base k -> dec):**
```
funkcja naDziesietny(cyfry, k, dlug)
    // cyfry — tablica cyfr, dlug — liczba cyfr
    wynik := 0
    dla i := 0, 1, ..., dlug-1:
        wynik := wynik * k + cyfry[i]
    zwroc wynik
```

**Zrodla maturalne**: 2023/6, 2025/5, 2015/2, 2014/3c

---

### Archetyp 4: NWD Euklidesa

**Szablon:**
```
funkcja NWD(a, b)
    dopoki b <> 0:
        temp := b
        b := a mod b
        a := temp
    zwroc a
```

**Tabelka sledzenia** dla `NWD(48, 18)`:

| Krok | a | b | a mod b | temp |
|------|---|---|---------|------|
| 1 | 48 | 18 | 12 | 18 |
| 2 | 18 | 12 | 6 | 12 |
| 3 | 12 | 6 | 0 | 6 |
| 4 | 6 | 0 | — | — |

Wynik: NWD(48, 18) = 6

**Wersja rekurencyjna (czesto na maturze):**
```
funkcja NWD(a, b)
    jezeli b = 0:
        zwroc a
    zwroc NWD(b, a mod b)
```

**Zrodla maturalne**: 2015/3, 2019/4, 2025/3

---

### Archetyp 5: Zliczanie z warunkiem

**Szablon:**
```
funkcja zlicz(T, n)
    licznik := 0
    dla i := 0, 1, ..., n-1:
        jezeli warunek(T[i]):
            licznik := licznik + 1
    zwroc licznik
```

**Przyklad — zliczanie wystapien w tablicy:**
```
funkcja sprawdz(T, n)
    C[0] := 0; C[1] := 0; C[2] := 0; C[3] := 0
    dla i := 0, 1, ..., n-1:
        C[T[i]] := C[T[i]] + 1
    dla j := 0, 1, 2, 3:
        jezeli C[j] <> j:
            zwroc FALSZ
    zwroc PRAWDA
```

**Tabelka sledzenia** dla `sprawdz([3, 1, 2, 1, 3, 2], 6)`:

| Krok (i) | T[i] | C[0] | C[1] | C[2] | C[3] |
|-----------|-------|------|------|------|------|
| start | - | 0 | 0 | 0 | 0 |
| 0 | 3 | 0 | 0 | 0 | 1 |
| 1 | 1 | 0 | 1 | 0 | 1 |
| 2 | 2 | 0 | 1 | 1 | 1 |
| 3 | 1 | 0 | 2 | 1 | 1 |
| 4 | 3 | 0 | 2 | 1 | 2 |
| 5 | 2 | 0 | 2 | 2 | 2 |

C = [0, 2, 2, 2]. C[1]=2 <> 1, wiec wynik: FALSZ

**Zrodla maturalne**: 2022/1, 2023/3, 2024/3

---

### Archetyp 6: Budowanie wyniku z cyfr

Wariant archetypu 1, w ktorym zamiast zliczac — budujemy nowa liczbe z wybranych/przetworzonych cyfr.

**Szablon:**
```
funkcja zbuduj(n)
    wynik := 0
    mnoznik := 1
    dopoki n > 0:
        cyfra := n mod 10
        jezeli warunek(cyfra):
            wynik := wynik + przetworz(cyfra) * mnoznik
            mnoznik := mnoznik * 10
        n := n div 10
    zwroc wynik
```

**Przyklad — "nieparzysty skrot" z matury 2024:**
```
funkcja nieparzyskiSkrot(n)
    wynik := 0
    mnoznik := 1
    dopoki n > 0:
        cyfra := n mod 10
        jezeli cyfra mod 2 = 1:
            wynik := wynik + cyfra * mnoznik
            mnoznik := mnoznik * 10
        n := n div 10
    zwroc wynik
```

**Tabelka sledzenia** dla `nieparzyskiSkrot(24639)`:

| Krok | n | cyfra | nieparzysta? | wynik | mnoznik |
|------|---|-------|-------------|-------|---------|
| 1 | 24639 | 9 | tak | 9 | 10 |
| 2 | 2463 | 3 | tak | 39 | 100 |
| 3 | 246 | 6 | nie | 39 | 100 |
| 4 | 24 | 4 | nie | 39 | 100 |
| 5 | 2 | 2 | nie | 39 | 100 |

Wynik: 39 (cyfry nieparzyste: 9, 3 — w oryginalnej kolejnosci)

**Pulapka**: `mnoznik` rosnie TYLKO gdy cyfra spelnia warunek — inaczej w wyniku pojawiaja sie "dziury".

**Zrodla maturalne**: 2024/2 (Cyfry), 2024/3 (pseudokod)

---

### Archetyp 7: Zamiana rekurencji na iteracje

Standardowe zadanie CKE — "przepisz rekurencyjnie / iteracyjnie".

**Rekurencja prosta (ogonowa) -> petla while:**
```
// Rekurencja:
funkcja suma(n)
    jezeli n = 0:
        zwroc 0
    zwroc suma(n div 10) + n mod 10

// Iteracja:
funkcja sumaIter(n)
    wynik := 0
    dopoki n > 0:
        wynik := wynik + n mod 10
        n := n div 10
    zwroc wynik
```

Regula: warunek bazowy -> warunek konca petli, akumulator zbiera wynik.

**Rekurencja ze stosem (drzewo) -> petla + stos:**
```
// Rekurencja:
funkcja suma_preorder(wezel)
    jezeli wezel = NULL:
        zwroc 0
    lewy := suma_preorder(wezel.L)
    prawy := suma_preorder(wezel.R)
    zwroc wezel.w + lewy + prawy

// Iteracja:
funkcja suma_preorder_iter(korzen)
    jezeli korzen = NULL:
        zwroc 0
    suma := 0
    stos := nowy pusty stos
    stos.push(korzen)
    dopoki stos nie jest pusty:
        wezel := stos.pop()
        suma := suma + wezel.w
        jezeli wezel.R <> NULL:
            stos.push(wezel.R)
        jezeli wezel.L <> NULL:
            stos.push(wezel.L)
    zwroc suma
```

Regula: na stos wkladamy potomkow w ODWROTNEJ kolejnosci (prawy, potem lewy), bo stos jest LIFO.

**Zrodla maturalne**: 2025/1.3, 2014/1c (Korale), 2023/1 (BST)

---

## 4. Tablice i struktury danych

### Tablice

```
// Jednowymiarowa — indeksowanie od 0
T[0], T[1], ..., T[n-1]

// Dwuwymiarowa — wiersz i, kolumna j
A[i, j]    // lub A[i][j] w niektorych zadaniach
```

### Stos (LIFO)

```
stos := nowy pusty stos
stos.push(x)          // wloz na stos
x := stos.pop()       // zdejmij ze stosu (zwraca wartosc)
stos nie jest pusty    // warunek petli
```

C++: `stack<int> s; s.push(x); int x = s.top(); s.pop();`

### Drzewa binarne

```
wezel.w    // wartosc wezla
wezel.L    // lewy potomek (moze byc NULL)
wezel.R    // prawy potomek (moze byc NULL)
```

### Kolejka (FIFO) — rzadziej, ale pojawia sie

```
kolejka := nowa pusta kolejka
kolejka.push(x)       // wloz na koniec
x := kolejka.pop()    // zdejmij z poczatku
```

---

## 5. Ograniczenia CKE — czego NIE wolno

Typowe ograniczenia w zadaniach z pseudokodem:

| Zakaz | Dlaczego | Obejscie |
|---|---|---|
| Stringi / napisy | Pseudokod CKE operuje na liczbach | Uzywaj `mod 10` / `div 10` do ekstrakcji cyfr |
| Tablice | Wymuszenie algorytmu "w locie" | Przetwarzaj cyfra po cyfrze w petli |
| Float / zmiennoprzecinkowe | Ograniczenie do arytmetyki calkowitej | `div` zamiast `/`, `mod` zamiast reszty |
| Funkcje wbudowane | Nie mozna uzyc `sqrt()`, `abs()`, `toString()` | Implementuj recznie (np. `i*i <= n` zamiast `sqrt(n)`) |
| Konwersja na string | `str(n)` nie istnieje w pseudokodzie | Petla `mod 10` / `div 10` |

**Czeste sformulowanie w arkuszu**:
> "Uzyj wylacznie zmiennych calkowitych i operatorow: mod, div, +, -, *, porownania."

---

## 6. Jak rysowac tabelke sledzenia

### Format

```
| Krok | zmienna_1 | zmienna_2 | ... | warunek? | wynik |
|------|-----------|-----------|-----|----------|-------|
| start | wart_poczatkowa | ... | | |
| 1    | ... | ... | ... | ... | ... |
| 2    | ... | ... | ... | ... | ... |
```

### Zasady

1. **Kolumny = zmienne** — jedna kolumna na kazda zmienna (w tym pomocnicze jak `cyfra`, `temp`)
2. **Wiersze = kroki** — jeden wiersz na kazda iteracje petli lub wywolanie rekurencji
3. **Warunek** — kolumna "warunek?" pokazuje decyzje w `jezeli`
4. **Pierwszy wiersz** — wartosci poczatkowe (przed petla)
5. **Ostatni wiersz** — wartosci koncowe (wynik)

### Przyklad: sledzenie `odwroc(4826)`

| Krok | n (pocz.) | cyfra | odwrocona | n (koniec) |
|------|-----------|-------|-----------|------------|
| start | 4826 | — | 0 | — |
| 1 | 4826 | 6 | 6 | 482 |
| 2 | 482 | 2 | 62 | 48 |
| 3 | 48 | 8 | 628 | 4 |
| 4 | 4 | 4 | 6284 | 0 |

Wynik: 6284

### Dla rekurencji — rysuj drzewo wywolan

```
g(12) = g(6) + g(4) + 1
├── g(6) = g(3) + g(2) + 1 = 6
│   ├── g(3) = g(1) + g(1) + 1 = 3
│   └── g(2) = g(1) + g(0) + 1 = 2
└── g(4) = g(2) + g(1) + 1 = 4
    └── g(2) = g(1) + g(0) + 1 = 2
=> g(12) = 6 + 4 + 1 = 11
```

Przy kazdym wezle zapisz: argument i wartosc zwracana.

---

## 7. Szesc pulapek pseudokodowych

### Pulapka 1: Brak warunku stopu

```
// ZLE — petla nieskonczona dla n = 0
funkcja cyfry(n)
    dopoki n > 0:       // nigdy nie wejdzie!
        ...
    zwroc wynik          // zwraca 0, ale nie przetwarza n=0
```

**Rozwiazanie**: Sprawdz przypadek `n = 0` osobno lub uzyj petli `do...while`.

### Pulapka 2: `=` vs `:=`

```
// W pseudokodzie CKE:
x := 5       // przypisanie (ustaw x na 5)
jezeli x = 5:  // porownanie (czy x rowne 5?)
```

Pomylenie tych operatorow to czesty blad przy sledzeniu. W C++ analogicznie: `=` (przypisanie) vs `==` (porownanie).

### Pulapka 3: n = 0 lub jednocyfrowe n

Wiele archetypow (mod/div) nie obsluguje `n = 0`:
- `0 mod 10 = 0`, `0 div 10 = 0` — petla `dopoki n > 0` w ogole nie wchodzi
- Dla `n = 5` (jednocyfrowe) — petla wykonuje sie raz, ale `mnoznik` moze nie byc potrzebny

**Sprawdzaj przypadki brzegowe**: n=0, n jednocyfrowe, n z zerami w srodku (np. 1024).

### Pulapka 4: Wciecia = blok

```
// Poprawne — oba wiersze sa w bloku if:
jezeli cyfra mod 2 = 0:
    wynik := wynik + cyfra
    licznik := licznik + 1

// Bledne odczytanie — licznik poza if:
jezeli cyfra mod 2 = 0:
    wynik := wynik + cyfra
licznik := licznik + 1       // to wykonuje sie ZAWSZE!
```

Na maturze CKE celowo testuje uwazne czytanie wiec — liczy sie kazda spacja.

### Pulapka 5: Nieparzysta liczba cyfr

W architekturze "buduj wynik z cyfr" (archetyp 6), jesli mnoznik rosnie tylko warunkowo, wynikowa liczba moze miec mniej cyfr niz oryginalna. Np.:

```
nieparzyskiSkrot(2468) = 0   // zadna cyfra nieparzysta
nieparzyskiSkrot(24639) = 39  // tylko 2 cyfry z 5
```

To nie blad — to zamierzone dzialanie. Ale na maturze moga zapytac "ile cyfr ma wynik?" i latwo sie pomylic.

### Pulapka 6: Kolejnosc stosu (LIFO)

Przy zamianie rekurencji na iteracje ze stosem — wkladaj na stos w ODWROTNEJ kolejnosci:

```
// Chcemy przetworzyc: lewy, potem prawy
// Wkladamy na stos: PRAWY, potem LEWY
jezeli wezel.R <> NULL:
    stos.push(wezel.R)      // prawy PIERWSZY na stos
jezeli wezel.L <> NULL:
    stos.push(wezel.L)      // lewy DRUGI — wiec zdejmie sie PIERWSZY
```

Pomylenie kolejnosci daje prawidlowa sume, ale zly porzadek odwiedzania wezlow.

---

## Podsumowanie

| Co | Gdzie w tym dokumencie |
|---|---|
| Nie wiem, co znaczy `mod` / `div` / `:=` | Sekcja 1 (slowniczek) |
| Nie wiem, jak czytac pseudokod CKE | Sekcja 2 (struktura bloku) |
| Szukam wzorca na konkretny typ zadania | Sekcja 3 (7 archetypow) |
| Nie pamietam skladni stosu / drzewa | Sekcja 4 (struktury danych) |
| Nie wiem, czego NIE wolno uzywac | Sekcja 5 (ograniczenia CKE) |
| Nie umiem rysowac tabelki sledzenia | Sekcja 6 (poradnik) |
| Robie bledy przy sledzeniu | Sekcja 7 (6 pulapek) |

---

*Powiazane materialy:*
- `strategia_egzaminacyjna.md` — TOP 14 algorytmow z kodem C++
- `cwiczenia_wg_typu/01_sledzenie_algorytmu.md` — cwiczenia ze sledzenia
- `cwiczenia_wg_typu/02_projektowanie_algorytmu.md` — cwiczenia z projektowania pseudokodu
- `cpp_szablony.md` — sciagawka C++ (implementacja tych samych algorytmow)
