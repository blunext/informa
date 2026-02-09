# 01. Sledzenie algorytmu

Typ zadania: **sledzenie_algorytmu**
Czestotliwosc: 11/11 lat | Laczna punktacja: 45 pkt
Kategoria: TEORIA

---

### Cwiczenie 1.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 2 (Cyfry)

Dana jest funkcja:

```
funkcja przetworzCyfry(n)
    wynik := 0
    mnoznik := 1
    dopoki n > 0:
        cyfra := n mod 10
        jezeli cyfra mod 2 = 0:
            wynik := wynik + (cyfra / 2) * mnoznik
        w przeciwnym razie:
            wynik := wynik + 1 * mnoznik
        mnoznik := mnoznik * 10
        n := n div 10
    zwroc wynik
```

Funkcja przetwarza cyfry liczby od prawej strony: cyfry parzyste dzieli przez 2, a nieparzyste zamienia na 1.

**Polecenie**: Oblicz wartosc `przetworzCyfry(n)` dla kazdej z podanych liczb. Dla jednej wybranej liczby wypelnij tabele wartosci zmiennych w kazdym kroku petli.

| Lp. | n |
|-----|---|
| a)  | 4826 |
| b)  | 1357 |
| c)  | 9042 |

<details>
<summary>Odpowiedz</summary>

**a) przetworzCyfry(4826)**

| Krok | n (poczatek) | cyfra | cyfra%2=0? | wynik | mnoznik | n (koniec) |
|------|-------------|-------|------------|-------|---------|-----------|
| 1 | 4826 | 6 | tak (6/2=3) | 0+3*1=3 | 10 | 482 |
| 2 | 482 | 2 | tak (2/2=1) | 3+1*10=13 | 100 | 48 |
| 3 | 48 | 8 | tak (8/2=4) | 13+4*100=413 | 1000 | 4 |
| 4 | 4 | 4 | tak (4/2=2) | 413+2*1000=2413 | 10000 | 0 |

Wynik: **2413**

**b) przetworzCyfry(1357)**

- cyfra 7 (nieparzysta) -> 1, wynik = 1
- cyfra 5 (nieparzysta) -> 1, wynik = 11
- cyfra 3 (nieparzysta) -> 1, wynik = 111
- cyfra 1 (nieparzysta) -> 1, wynik = 1111

Wynik: **1111**

**c) przetworzCyfry(9042)**

- cyfra 2 (parzysta, 2/2=1) -> wynik = 1
- cyfra 4 (parzysta, 4/2=2) -> wynik = 21
- cyfra 0 (parzysta, 0/2=0) -> wynik = 021 = 21
- cyfra 9 (nieparzysta) -> 1, wynik = 1021

Wynik: **1021**
</details>

---

### Cwiczenie 1.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 1.1 (uproszczony)

Dana jest funkcja rekurencyjna:

```
funkcja sumaCyfr(n)
    jezeli n = 0:
        zwroc 0
    zwroc sumaCyfr(n div 10) + n mod 10
```

**Polecenie**: Dla kazdej z podanych wartosci n:
- podaj wynik funkcji `sumaCyfr(n)`,
- podaj calkowita liczbe wywolan funkcji (wliczajac wywolanie poczatkowe).

| Lp. | n |
|-----|---|
| a)  | 47 |
| b)  | 305 |
| c)  | 1234 |

<details>
<summary>Odpowiedz</summary>

**a) sumaCyfr(47)**

Stos wywolan:
1. sumaCyfr(47) -> sumaCyfr(4) + 7
2. sumaCyfr(4) -> sumaCyfr(0) + 4
3. sumaCyfr(0) -> 0 (warunek bazowy)

Powrot: 0 + 4 + 7 = **11**
Liczba wywolan: **3**

**b) sumaCyfr(305)**

Stos wywolan:
1. sumaCyfr(305) -> sumaCyfr(30) + 5
2. sumaCyfr(30) -> sumaCyfr(3) + 0
3. sumaCyfr(3) -> sumaCyfr(0) + 3
4. sumaCyfr(0) -> 0 (warunek bazowy)

Powrot: 0 + 3 + 0 + 5 = **8**
Liczba wywolan: **4**

**c) sumaCyfr(1234)**

Stos wywolan:
1. sumaCyfr(1234) -> sumaCyfr(123) + 4
2. sumaCyfr(123) -> sumaCyfr(12) + 3
3. sumaCyfr(12) -> sumaCyfr(1) + 2
4. sumaCyfr(1) -> sumaCyfr(0) + 1
5. sumaCyfr(0) -> 0 (warunek bazowy)

Powrot: 0 + 1 + 2 + 3 + 4 = **10**
Liczba wywolan: **5**

**Regula**: Dla liczby o k cyfrach funkcja jest wywolywana k+1 razy (k wywolan z niezerowym n + 1 wywolanie bazowe z n=0).
</details>

---

### Cwiczenie 1.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 1 (n-permutacja)

Dany jest algorytm operujacy na tablicy T o elementach ze zbioru {0, 1, 2, 3}:

```
funkcja sprawdz(T, n)
    // T - tablica o indeksach 0..n-1
    // tworzymy tablice C o indeksach 0..3, wypelniona zerami
    C[0] := 0; C[1] := 0; C[2] := 0; C[3] := 0

    dla i := 0, 1, ..., n-1:
        C[T[i]] := C[T[i]] + 1

    dla j := 0, 1, 2, 3:
        jezeli C[j] <> j:
            zwroc FALSZ

    zwroc PRAWDA
```

**Polecenie**: Przesled algorytm krok po kroku dla tablicy T = [3, 1, 2, 1, 3, 2]. Podaj:
- a) Stan tablicy C po kazdej iteracji pierwszej petli.
- b) Wynik koncowy funkcji (PRAWDA lub FALSZ).
- c) Podaj przyklad tablicy 6-elementowej (z wartosciami ze zbioru {0, 1, 2, 3}), dla ktorej funkcja zwroci PRAWDA.

<details>
<summary>Odpowiedz</summary>

**a) Stan tablicy C po kazdej iteracji:**

Tablica T = [3, 1, 2, 1, 3, 2], n = 6

| Krok (i) | T[i] | Operacja | C[0] | C[1] | C[2] | C[3] |
|-----------|-------|----------|------|------|------|------|
| poczatek  | -     | -        | 0    | 0    | 0    | 0    |
| i=0       | 3     | C[3]++   | 0    | 0    | 0    | 1    |
| i=1       | 1     | C[1]++   | 0    | 1    | 0    | 1    |
| i=2       | 2     | C[2]++   | 0    | 1    | 1    | 1    |
| i=3       | 1     | C[1]++   | 0    | 2    | 1    | 1    |
| i=4       | 3     | C[3]++   | 0    | 2    | 1    | 2    |
| i=5       | 2     | C[2]++   | 0    | 2    | 2    | 2    |

Stan koncowy C: [0, 2, 2, 2]

**b) Sprawdzenie warunku C[j] = j:**

| j | C[j] | j | C[j] = j? |
|---|------|---|-----------|
| 0 | 0    | 0 | TAK       |
| 1 | 2    | 1 | NIE       |

Funkcja zwraca **FALSZ** (bo C[1] = 2, a powinno byc 1).

**c) Przyklad tablicy, dla ktorej funkcja zwraca PRAWDA:**

Potrzebujemy: C[0]=0, C[1]=1, C[2]=2, C[3]=3.
Czyli: 0 wystapien 0, 1 wystapienie 1, 2 wystapienia 2, 3 wystapienia 3.
Razem: 0+1+2+3 = 6 elementow.

Przyklad: T = **[3, 2, 3, 1, 2, 3]**

Sprawdzenie: C = [0, 1, 2, 3] -- C[j] = j dla kazdego j, PRAWDA.
</details>

---

### Cwiczenie 1.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 2, Matura 2014 zad. 1a (Korale)

Dana jest funkcja rekurencyjna z dwoma wywolaniami:

```
funkcja g(n)
    jezeli n <= 1:
        zwroc n
    zwroc g(n div 2) + g(n div 3) + 1
```

**Polecenie**:
- a) Narysuj pelne drzewo wywolan dla `g(12)`. Przy kazdym wezle zapisz wartosc argumentu n oraz wartosc zwracana.
- b) Podaj calkowita liczbe wywolan funkcji g.
- c) Podaj wynik koncowy `g(12)`.

<details>
<summary>Odpowiedz</summary>

**a) Drzewo wywolan g(12):**

```
g(12) = g(6) + g(4) + 1
├── g(6) = g(3) + g(2) + 1
│   ├── g(3) = g(1) + g(1) + 1
│   │   ├── g(1) = 1  [bazowy]
│   │   └── g(1) = 1  [bazowy]
│   │   => g(3) = 1 + 1 + 1 = 3
│   └── g(2) = g(1) + g(0) + 1
│       ├── g(1) = 1  [bazowy]
│       └── g(0) = 0  [bazowy]
│       => g(2) = 1 + 0 + 1 = 2
│   => g(6) = 3 + 2 + 1 = 6
└── g(4) = g(2) + g(1) + 1
    ├── g(2) = g(1) + g(0) + 1
    │   ├── g(1) = 1  [bazowy]
    │   └── g(0) = 0  [bazowy]
    │   => g(2) = 1 + 0 + 1 = 2
    └── g(1) = 1  [bazowy]
    => g(4) = 2 + 1 + 1 = 4
=> g(12) = 6 + 4 + 1 = 11
```

**b) Calkowita liczba wywolan:**

Policzymy wszystkie wezly drzewa:
- g(12): 1
- g(6), g(4): 2
- g(3), g(2), g(2), g(1): 4
- g(1), g(1), g(1), g(0), g(1), g(0): 6

Razem: 1 + 2 + 4 + 6 = **13 wywolan**

**c) Wynik koncowy: g(12) = 6 + 4 + 1 = **11****
</details>

---

### Cwiczenie 1.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 1 (Plansza)

Dana jest plansza 5x5 o polach bialych (0) i czarnych (1):

```
plansza:
  0 1 2 3 4
0[0 0 1 0 0]
1[0 0 0 1 0]
2[1 0 0 0 1]
3[0 1 0 0 0]
4[0 0 1 0 0]
```

Algorytm wypelnia tablice P[i][j] rozmiaru 5x5 wedlug nastepujacych regul:
- P[0][0] := 1, jezeli plansza[0][0] = 0; w przeciwnym razie P[0][0] := 0
- Dla i > 0: P[i][0] := P[i-1][0] AND (1 - plansza[i][0])
- Dla j > 0: P[0][j] := P[0][j-1] AND (1 - plansza[0][j])
- Dla i > 0 i j > 0: P[i][j] := (P[i-1][j] OR P[i][j-1]) AND (1 - plansza[i][j])

Wartosc P[i][j] = 1 oznacza, ze istnieje sciezka z pola [0,0] do pola [i,j] przechodzaca wylacznie przez pola biale, poruszajac sie tylko w dol lub w prawo.

**Polecenie**:
- a) Wypelnij cala tablice P[5][5].
- b) Czy istnieje sciezka z [0,0] do [4,4]? Odpowiedz na podstawie wartosci P[4][4].
- c) Podaj jedna sciezke z [0,0] do [4,4] (jezeli istnieje) jako ciag wspolrzednych.

<details>
<summary>Odpowiedz</summary>

**a) Wypelnianie tablicy P:**

Plansza (0 = biale, 1 = czarne):
```
  0 1 2 3 4
0[0 0 1 0 0]
1[0 0 0 1 0]
2[1 0 0 0 1]
3[0 1 0 0 0]
4[0 0 1 0 0]
```

Krok 1: P[0][0] = 1 (bo plansza[0][0] = 0)

Krok 2: Pierwszy wiersz (i=0):
- P[0][1] = P[0][0] AND (1-0) = 1 AND 1 = 1
- P[0][2] = P[0][1] AND (1-1) = 1 AND 0 = 0  (czarne pole!)
- P[0][3] = P[0][2] AND (1-0) = 0 AND 1 = 0
- P[0][4] = P[0][3] AND (1-0) = 0 AND 1 = 0

Krok 3: Pierwsza kolumna (j=0):
- P[1][0] = P[0][0] AND (1-0) = 1 AND 1 = 1
- P[2][0] = P[1][0] AND (1-1) = 1 AND 0 = 0  (czarne pole!)
- P[3][0] = P[2][0] AND (1-0) = 0 AND 1 = 0
- P[4][0] = P[3][0] AND (1-0) = 0 AND 1 = 0

Krok 4: Reszta tablicy:
Wiersz 1 (i=1):
- P[1][1] = (P[0][1] OR P[1][0]) AND (1-0) = (1 OR 1) AND 1 = 1
- P[1][2] = (P[0][2] OR P[1][1]) AND (1-0) = (0 OR 1) AND 1 = 1
- P[1][3] = (P[0][3] OR P[1][2]) AND (1-1) = (0 OR 1) AND 0 = 0  (czarne!)
- P[1][4] = (P[0][4] OR P[1][3]) AND (1-0) = (0 OR 0) AND 1 = 0

Wiersz 2 (i=2):
- P[2][1] = (P[1][1] OR P[2][0]) AND (1-0) = (1 OR 0) AND 1 = 1
- P[2][2] = (P[1][2] OR P[2][1]) AND (1-0) = (1 OR 1) AND 1 = 1
- P[2][3] = (P[1][3] OR P[2][2]) AND (1-0) = (0 OR 1) AND 1 = 1
- P[2][4] = (P[1][4] OR P[2][3]) AND (1-1) = (0 OR 1) AND 0 = 0  (czarne!)

Wiersz 3 (i=3):
- P[3][1] = (P[2][1] OR P[3][0]) AND (1-1) = (1 OR 0) AND 0 = 0  (czarne!)
- P[3][2] = (P[2][2] OR P[3][1]) AND (1-0) = (1 OR 0) AND 1 = 1
- P[3][3] = (P[2][3] OR P[3][2]) AND (1-0) = (1 OR 1) AND 1 = 1
- P[3][4] = (P[2][4] OR P[3][3]) AND (1-0) = (0 OR 1) AND 1 = 1

Wiersz 4 (i=4):
- P[4][1] = (P[3][1] OR P[4][0]) AND (1-0) = (0 OR 0) AND 1 = 0
- P[4][2] = (P[3][2] OR P[4][1]) AND (1-1) = (1 OR 0) AND 0 = 0  (czarne!)
- P[4][3] = (P[3][3] OR P[4][2]) AND (1-0) = (1 OR 0) AND 1 = 1
- P[4][4] = (P[3][4] OR P[4][3]) AND (1-0) = (1 OR 1) AND 1 = 1

**Tablica P:**
```
  0 1 2 3 4
0[1 1 0 0 0]
1[1 1 1 0 0]
2[0 1 1 1 0]
3[0 0 1 1 1]
4[0 0 0 1 1]
```

**b) P[4][4] = 1, wiec sciezka z [0,0] do [4,4] ISTNIEJE.**

**c) Przykladowa sciezka (cofamy sie od [4,4]):**

[0,0] -> [0,1] -> [1,1] -> [1,2] -> [2,2] -> [2,3] -> [3,3] -> [3,4] -> [4,4]

Mozna tez: [0,0] -> [1,0] -> [1,1] -> [2,1] -> [2,2] -> [3,2] -> [3,3] -> [4,3] -> [4,4]
</details>
