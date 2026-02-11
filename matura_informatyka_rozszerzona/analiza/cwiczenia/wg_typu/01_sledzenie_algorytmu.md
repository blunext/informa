# 01. Sledzenie algorytmu

Typ zadania: **sledzenie_algorytmu**
Czestotliwosc: 11/11 lat | Laczna punktacja: 45 pkt
Kategoria: TEORIA

## Umiejetnosci cwiczone w tym zestawie

`mod-div` `cyfry-liczby` `sledzenie-iteracyjne` `rekurencja` `drzewo-wywolan` `tablica-zliczanie` `plansza-2D` `programowanie-dynamiczne` `NWD-Euklidesa` `stos` `sortowanie` `operacje-bitowe` `kolejka` `tablica-pomocnicza`

---

### Cwiczenie 1.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 2 (Cyfry)
**Tagi**: `mod-div` `cyfry-liczby` `sledzenie-iteracyjne`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Zastanow sie, co robi `n mod 10` i `n div 10` — wyodrebniaja cyfry od konca.
2. **Podejscie**: Dla kazdej cyfry sprawdz parzystosc i zastosuj odpowiednia regule. Pamietaj o mnozniku pozycyjnym.
3. **Kluczowy krok**: Przerob kazda cyfre osobno: parzysta -> /2, nieparzysta -> 1. Zlozyc wynik uzywajac mnoznika (1, 10, 100, ...).

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Przetwarzanie cyfr od lewej zamiast od prawej**: Algorytm uzywa mod 10, wiec zaczyna od ostatniej cyfry. CKE: -1 pkt
- **Zapomnienie o mnozniku**: Wynik sklada sie z cyfr na odpowiednich pozycjach — bez mnoznika cyfry sie nakladaja. CKE: -1 pkt
- **Cyfra 0 jako nieparzysta**: 0 mod 2 = 0, wiec 0 jest parzysta (0/2 = 0). CKE: -0.5 pkt

</details>

---

### Cwiczenie 1.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 1.1 (uproszczony)
**Tagi**: `rekurencja` `mod-div` `drzewo-wywolan`

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
<summary>Wskazowki</summary>

1. **Kierunek**: To rekurencja liniowa — kazde wywolanie generuje dokladnie jedno kolejne.
2. **Podejscie**: Rozpisz stos wywolan: sumaCyfr(1234) -> sumaCyfr(123) + 4 -> sumaCyfr(12) + 3 -> ...
3. **Kluczowy krok**: Warunek bazowy to n=0 (zwraca 0). Liczba wywolan = liczba cyfr + 1 (jedno dodatkowe dla n=0).

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Pominiecie wywolania bazowego (n=0)**: Czesto uczniowie licza tylko wywolania z n>0. Dla n=47 to 3, nie 2 wywolania. CKE: -1 pkt
- **Mylenie div i mod**: `n div 10` obcina ostatnia cyfre, `n mod 10` ja wyodrebnia. CKE: -1 pkt

</details>

---

### Cwiczenie 1.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 1 (n-permutacja)
**Tagi**: `tablica-zliczanie` `sledzenie-iteracyjne` `tablica-pomocnicza`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Tablica C to tablica zliczajaca — liczy ile razy kazda wartosc wystapila w T.
2. **Podejscie**: Iteruj po T element po elemencie, zwiekszajac odpowiedni licznik w C. Potem sprawdz warunek C[j]=j.
3. **Kluczowy krok**: Funkcja zwraca PRAWDA gdy: wartosc 0 wystepuje 0 razy, wartosc 1 — 1 raz, wartosc 2 — 2 razy, wartosc 3 — 3 razy. Razem 0+1+2+3=6 elementow.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie ze indeksowanie C zaczyna sie od 0**: C[T[i]] to indeksowanie wartoscia elementu, nie pozycja. CKE: -1 pkt
- **Blad w zliczaniu**: Np. pominiecie jednego kroku iteracji. Sprawdz, ze suma C[0]+...+C[3] = n. CKE: -0.5 pkt
- **Niepoprawny przyklad w c)**: Tablica musi miec dokladnie 6 elementow i spelnic C[j]=j. CKE: -1 pkt

</details>

---

### Cwiczenie 1.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 2, Matura 2014 zad. 1a (Korale)
**Tagi**: `rekurencja` `drzewo-wywolan` `sledzenie-iteracyjne`

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
<summary>Wskazowki</summary>

1. **Kierunek**: To rekurencja drzewiasta — kazde wywolanie generuje dwa kolejne (div 2 i div 3).
2. **Podejscie**: Zacznij od korzenia g(12), rozwin galaz lewa g(6) i prawa g(4), potem kazda dalej.
3. **Kluczowy krok**: Warunek bazowy: n<=1 zwraca n. Kazdy wezel niebazowy dodaje +1 do sumy dzieci.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Pominiety +1 w formule**: Kazde niebazowe wywolanie dodaje 1 do sumy. CKE: -1 pkt
- **Bledne div (zaokraglanie w gore zamiast w dol)**: 4 div 3 = 1 (nie 1.33 ani 2). CKE: -1 pkt
- **Niedokonczone drzewo**: Kazda galaz musi dojsc do warunku bazowego (n<=1). CKE: -2 pkt

</details>

---

### Cwiczenie 1.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 1 (Plansza)
**Tagi**: `plansza-2D` `programowanie-dynamiczne` `sledzenie-iteracyjne`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Wypelniaj tablice wiersz po wierszu, od lewej do prawej. Pierwszy wiersz i pierwsza kolumna maja specjalne reguly.
2. **Podejscie**: Dla pierwszego wiersza/kolumny: jesli napotkasz czarne pole, cala reszta wiersza/kolumny to 0. Dla reszty: P[i][j]=1 gdy mozna dojsc z gory LUB z lewej i pole jest biale.
3. **Kluczowy krok**: P[i][j] = (P[i-1][j] OR P[i][j-1]) AND (1-plansza[i][j]). Czarne pole zawsze daje 0.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Pomylenie AND i OR**: Pierwszy wiersz/kolumna uzywa AND (jedna przeszkoda blokuje dalsze pola), reszta uzywa OR (wystarczy dojsc z gory LUB z lewej). CKE: -2 pkt
- **Zapomnienie o (1-plansza)**: Nawet jesli mozna dojsc do pola, czarne pole daje P=0. CKE: -1 pkt
- **Bledna kolejnosc wypelniania**: Trzeba wypelniac wiersz po wierszu (od gory), w kazdym wierszu od lewej. CKE: -1 pkt

</details>

---

### Cwiczenie 1.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 1 (Algorytm Euklidesa)
**Tagi**: `NWD-Euklidesa` `sledzenie-iteracyjne`

Dany jest algorytm Euklidesa:

```
funkcja NWD(a, b)
    dopoki b <> 0:
        r := a mod b
        a := b
        b := r
    zwroc a
```

**Polecenie**: Dla kazdej pary (a, b) przesled algorytm, wypelniajac tabele wartosci zmiennych w kazdym kroku petli. Podaj wynik (NWD).

| Lp. | a | b |
|-----|---|---|
| a)  | 48 | 18 |
| b)  | 105 | 42 |
| c)  | 121 | 33 |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: W kazdym kroku oblicz reszte z dzielenia a przez b, potem przesun: a←b, b←r.
2. **Podejscie**: Kontynuuj az b = 0. Wtedy a to NWD.
3. **Kluczowy krok**: Przyklad: 48 mod 18 = 12, potem 18 mod 12 = 6, potem 12 mod 6 = 0 — NWD = 6.

</details>

<details>
<summary>Odpowiedz</summary>

**a) NWD(48, 18)**

| Krok | a | b | r = a mod b |
|------|---|---|-------------|
| 1 | 48 | 18 | 12 |
| 2 | 18 | 12 | 6 |
| 3 | 12 | 6 | 0 |

b = 0, NWD = **6**

**b) NWD(105, 42)**

| Krok | a | b | r = a mod b |
|------|---|---|-------------|
| 1 | 105 | 42 | 21 |
| 2 | 42 | 21 | 0 |

b = 0, NWD = **21**

**c) NWD(121, 33)**

| Krok | a | b | r = a mod b |
|------|---|---|-------------|
| 1 | 121 | 33 | 22 |
| 2 | 33 | 22 | 11 |
| 3 | 22 | 11 | 0 |

b = 0, NWD = **11**
</details>

<details>
<summary>Typowe bledy</summary>

- **Bledna kolejnosc przypisamien a←b, b←r**: Trzeba najpierw obliczyc r, POTEM przypisac a:=b, b:=r. CKE: -1 pkt
- **Zapomnienie, ze wynik to a (nie b) po zakonczeniu petli**: Petla konczy sie gdy b=0, wynik jest w a. CKE: -0.5 pkt

</details>

---

### Cwiczenie 1.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 1 (sledzenie sortowania)
**Tagi**: `sortowanie` `sledzenie-iteracyjne` `tablica-pomocnicza`

Dany jest algorytm sortowania babelkowego:

```
funkcja babelkowe(T, n)
    // T - tablica o indeksach 0..n-1
    dla i := 0, 1, ..., n-2:
        dla j := 0, 1, ..., n-2-i:
            jezeli T[j] > T[j+1]:
                zamien(T[j], T[j+1])
```

**Polecenie**: Przesled algorytm dla tablicy T = [5, 3, 8, 1, 4] (n=5). Podaj:
- a) Stan tablicy po kazdym pelnym przebiegu petli zewnetrznej (i=0, i=1, ...).
- b) Laczna liczbe wykonanych zamian (swap).
- c) Po ktorym przebiegu petli zewnetrznej tablica jest juz posortowana?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: W kazdym przebiegu petli zewnetrznej najwiekszy nieposortowany element "wyplywa" na koniec.
2. **Podejscie**: Dla i=0 porownujesz sasiednie pary j=0..3; dla i=1 j=0..2 (ostatni juz na miejscu) itd.
3. **Kluczowy krok**: Zapisz stan tablicy po KAZDYM przebiegu i licz zamiany. Sortowanie moze skonczyc sie wczesniej niz n-1 przebiegow.

</details>

<details>
<summary>Odpowiedz</summary>

Poczatkowo: T = [5, 3, 8, 1, 4]

**Przebieg i=0** (porownania j=0..3):
- j=0: T[0]=5 > T[1]=3? TAK -> zamien -> [**3**, **5**, 8, 1, 4] (zamiana 1)
- j=1: T[1]=5 > T[2]=8? NIE
- j=2: T[2]=8 > T[3]=1? TAK -> zamien -> [3, 5, **1**, **8**, 4] (zamiana 2)
- j=3: T[3]=8 > T[4]=4? TAK -> zamien -> [3, 5, 1, **4**, **8**] (zamiana 3)

Stan po i=0: [3, 5, 1, 4, 8] — element 8 na swoim miejscu.

**Przebieg i=1** (porownania j=0..2):
- j=0: T[0]=3 > T[1]=5? NIE
- j=1: T[1]=5 > T[2]=1? TAK -> zamien -> [3, **1**, **5**, 4, 8] (zamiana 4)
- j=2: T[2]=5 > T[3]=4? TAK -> zamien -> [3, 1, **4**, **5**, 8] (zamiana 5)

Stan po i=1: [3, 1, 4, 5, 8] — elementy 5, 8 na swoich miejscach.

**Przebieg i=2** (porownania j=0..1):
- j=0: T[0]=3 > T[1]=1? TAK -> zamien -> [**1**, **3**, 4, 5, 8] (zamiana 6)
- j=1: T[1]=3 > T[2]=4? NIE

Stan po i=2: [1, 3, 4, 5, 8] — **tablica posortowana!**

**Przebieg i=3** (porownania j=0..0):
- j=0: T[0]=1 > T[1]=3? NIE

Stan po i=3: [1, 3, 4, 5, 8] — bez zmian.

**a)** Stany po przebiegach: [3,5,1,4,8] → [3,1,4,5,8] → [1,3,4,5,8] → [1,3,4,5,8]

**b)** Laczna liczba zamian: **6**

**c)** Tablica jest posortowana po przebiegu **i=2** (trzecim przebiegu).
</details>

<details>
<summary>Typowe bledy</summary>

- **Bledny zakres petli wewnetrznej**: j idzie do n-2-i (nie n-2 ani n-1). CKE: -1 pkt
- **Zapomnienie o dalszych porownaniach po zamianie**: Po zamianie T[j] i T[j+1] kontynuujemy z j+1. CKE: -1 pkt
- **Mylenie sortowania rosnacego z malejacym**: Warunek T[j]>T[j+1] daje sortowanie rosnace. CKE: -0.5 pkt

</details>

---

### Cwiczenie 1.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 2 (operacje bitowe)
**Tagi**: `operacje-bitowe` `sledzenie-iteracyjne` `mod-div`

Dana jest funkcja konwertujaca liczbe na zapis binarny i zliczajaca jedynki:

```
funkcja policzJedynki(n)
    licznik := 0
    bity := ""
    dopoki n > 0:
        bit := n mod 2
        bity := bit + bity   // dopisz bit na poczatek napisu
        jezeli bit = 1:
            licznik := licznik + 1
        n := n div 2
    zwroc (bity, licznik)
```

**Polecenie**: Przesled algorytm dla kazdej z podanych wartosci n. Podaj:
- zapis binarny (wartosc zmiennej `bity`),
- liczbe jedynek (`licznik`).

Dla n = 53 wypelnij pelna tabele wartosci zmiennych w kazdym kroku.

| Lp. | n |
|-----|---|
| a)  | 53 |
| b)  | 100 |
| c)  | 255 |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Algorytm dzieli n przez 2 i zbiera reszty — to standardowa konwersja na system binarny.
2. **Podejscie**: W kazdym kroku: reszta z dzielenia przez 2 to kolejny bit (od prawej), potem n := n div 2.
3. **Kluczowy krok**: Bity zbieramy od konca do poczatku (dlatego `bit + bity`). Jedynki liczymy na biezaco.

</details>

<details>
<summary>Odpowiedz</summary>

**a) policzJedynki(53)**

| Krok | n (pocz.) | bit = n mod 2 | bity | licznik | n (koniec) |
|------|-----------|----------------|------|---------|------------|
| 1 | 53 | 1 | "1" | 1 | 26 |
| 2 | 26 | 0 | "01" | 1 | 13 |
| 3 | 13 | 1 | "101" | 2 | 6 |
| 4 | 6 | 0 | "0101" | 2 | 3 |
| 5 | 3 | 1 | "10101" | 3 | 1 |
| 6 | 1 | 1 | "110101" | 4 | 0 |

Wynik: bity = **"110101"**, licznik = **4**
Sprawdzenie: 32+16+4+1 = 53. Poprawne.

**b) policzJedynki(100)**

100 div 2: 0, 0, 1, 0, 0, 1, 1 (kolejne reszty)
- 100→50(r=0), 50→25(r=0), 25→12(r=1), 12→6(r=0), 6→3(r=0), 3→1(r=1), 1→0(r=1)

Wynik: bity = **"1100100"**, licznik = **3**
Sprawdzenie: 64+32+4 = 100. Poprawne.

**c) policzJedynki(255)**

255 = 2^8 - 1, wiec w zapisie binarnym to osiem jedynek.
- 255→127(r=1), 127→63(r=1), 63→31(r=1), 31→15(r=1), 15→7(r=1), 7→3(r=1), 3→1(r=1), 1→0(r=1)

Wynik: bity = **"11111111"**, licznik = **8**
</details>

<details>
<summary>Typowe bledy</summary>

- **Odwrocona kolejnosc bitow**: Algorytm generuje bity od LSB do MSB, ale sklada je w poprawnej kolejnosci (bit + bity). Jesli zapisujesz bity w odwrotnej kolejnosci, wynik bedzie bledny. CKE: -1 pkt
- **Blad przy n=0**: Algorytm nie obsluguje n=0 (petla sie nie wykona). Na maturze CKE nie sprawdza tego przypadku.
- **Pomylka div i mod**: n mod 2 to reszta (0 lub 1), n div 2 to czesc calkowita. CKE: -1 pkt

</details>

---

### Cwiczenie 1.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2018 zad. 1 (stos/odwrotna notacja polska)
**Tagi**: `stos` `sledzenie-iteracyjne` `tablica-pomocnicza`

Dany jest algorytm obliczajacy wartosc wyrazenia w odwrotnej notacji polskiej (ONP) z uzyciem stosu:

```
funkcja obliczONP(wyrazenie)
    // wyrazenie to tablica tokenow (liczb i operatorow)
    stos := pusty stos
    dla kazdego tokenu t w wyrazeniu:
        jezeli t jest liczba:
            poloz(stos, t)
        w przeciwnym razie:  // t jest operatorem (+, -, *, /)
            b := zdejmij(stos)
            a := zdejmij(stos)
            wynik := a (operator t) b
            poloz(stos, wynik)
    zwroc zdejmij(stos)
```

**Polecenie**: Przesled algorytm dla wyrazenia ONP: `3 4 2 * + 7 1 - /`

Podaj:
- a) Stan stosu po przetworzeniu kazdego tokenu.
- b) Wynik koncowy.
- c) Zapisz to wyrazenie w notacji infiksowej (standardowej, ze nawiasami).

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Liczby kladziemy na stos. Operator zdejmuje dwa elementy (uwaga na kolejnosc: a to glabszy, b to wierzcholek), oblicza wynik i kladzie go z powrotem.
2. **Podejscie**: Przerob token po tokenie, rysujac stan stosu po kazdym kroku.
3. **Kluczowy krok**: Przy dzieleniu i odejmowaniu kolejnosc ma znaczenie: a-b i a/b, gdzie a zdejmujemy drugie, b pierwsze.

</details>

<details>
<summary>Odpowiedz</summary>

Wyrazenie: `3 4 2 * + 7 1 - /`

| Token | Operacja | Stos (dol → gora) |
|-------|----------|-------------------|
| 3 | poloz 3 | [3] |
| 4 | poloz 4 | [3, 4] |
| 2 | poloz 2 | [3, 4, 2] |
| * | b=2, a=4; 4*2=8; poloz 8 | [3, 8] |
| + | b=8, a=3; 3+8=11; poloz 11 | [11] |
| 7 | poloz 7 | [11, 7] |
| 1 | poloz 1 | [11, 7, 1] |
| - | b=1, a=7; 7-1=6; poloz 6 | [11, 6] |
| / | b=6, a=11; 11/6=1 (dzielenie calkowite); poloz 1 | [1] |

**a)** Stany stosu — jak w tabeli powyzej.

**b)** Wynik koncowy: **1** (przy dzieleniu calkowitym) lub **11/6 ≈ 1.83** (przy dzieleniu rzeczywistym).

**c)** Notacja infiksowa: **(3 + 4 * 2) / (7 - 1)**

Weryfikacja: (3 + 8) / 6 = 11 / 6 = 1 (calkowite) lub ~1.83 (rzeczywiste).
</details>

<details>
<summary>Typowe bledy</summary>

- **Odwrocona kolejnosc operandow**: Przy `-` i `/` kolejnosc ma znaczenie! Najpierw zdejmujemy b (wierzcholek), potem a (glabszy). Obliczamy a op b, nie b op a. CKE: -2 pkt
- **Brak informacji o typie dzielenia**: Na maturze CKE wymaga zaznaczenia czy dzielenie calkowite czy rzeczywiste. CKE: -0.5 pkt
- **Bledne odczytanie ONP**: Tokeny czytamy od lewej do prawej, nie od prawej. CKE: -1 pkt

</details>

---

### Cwiczenie 1.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2015 zad. 1 (listy, sledzenie zlozonych algorytmow)
**Tagi**: `rekurencja` `drzewo-wywolan` `tablica-pomocnicza` `sledzenie-iteracyjne`

Dana jest funkcja rekurencyjna operujaca na tablicy:

```
funkcja tajemnicza(T, lewy, prawy)
    jezeli lewy >= prawy:
        zwroc T[lewy]
    srodek := (lewy + prawy) div 2
    L := tajemnicza(T, lewy, srodek)
    P := tajemnicza(T, srodek + 1, prawy)
    jezeli L > P:
        zwroc L
    w przeciwnym razie:
        zwroc P
```

**Polecenie**: Dla tablicy T = [3, 7, 2, 9, 1, 5, 8, 4] (indeksy 0..7):
- a) Narysuj pelne drzewo wywolan `tajemnicza(T, 0, 7)`. Przy kazdym wezle zapisz zakres [lewy, prawy] i wartosc zwracana.
- b) Podaj calkowita liczbe wywolan funkcji.
- c) Jaki problem rozwiazuje ta funkcja? Co zwroci dla dowolnej tablicy?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Funkcja dzieli tablice na polowy (lewy..srodek i srodek+1..prawy) i laczy wyniki. To schemat "dziel i zwyciezaj".
2. **Podejscie**: Zacznij od tajemnicza(T, 0, 7), srodek = 3. Lewa galaz: tajemnicza(T, 0, 3), prawa: tajemnicza(T, 4, 7). Kazda dalej sie dzieli.
3. **Kluczowy krok**: Warunek bazowy to lewy >= prawy (jednoelementowy zakres — zwroc T[lewy]). Operacja laczenia to `max(L, P)`.

</details>

<details>
<summary>Odpowiedz</summary>

T = [3, 7, 2, 9, 1, 5, 8, 4]

**a) Drzewo wywolan:**

```
tajemnicza(0,7) srodek=3
├── tajemnicza(0,3) srodek=1
│   ├── tajemnicza(0,1) srodek=0
│   │   ├── tajemnicza(0,0) -> T[0]=3
│   │   └── tajemnicza(1,1) -> T[1]=7
│   │   => max(3,7) = 7
│   └── tajemnicza(2,3) srodek=2
│       ├── tajemnicza(2,2) -> T[2]=2
│       └── tajemnicza(3,3) -> T[3]=9
│       => max(2,9) = 9
│   => max(7,9) = 9
└── tajemnicza(4,7) srodek=5
    ├── tajemnicza(4,5) srodek=4
    │   ├── tajemnicza(4,4) -> T[4]=1
    │   └── tajemnicza(5,5) -> T[5]=5
    │   => max(1,5) = 5
    └── tajemnicza(6,7) srodek=6
        ├── tajemnicza(6,6) -> T[6]=8
        └── tajemnicza(7,7) -> T[7]=4
        => max(8,4) = 8
    => max(5,8) = 8
=> max(9,8) = 9
```

**b) Calkowita liczba wywolan:**

- Poziom 0: 1 wywolanie (0,7)
- Poziom 1: 2 wywolania (0,3), (4,7)
- Poziom 2: 4 wywolania (0,1), (2,3), (4,5), (6,7)
- Poziom 3: 8 wywolan (bazowe) (0,0), (1,1), (2,2), (3,3), (4,4), (5,5), (6,6), (7,7)

Razem: 1 + 2 + 4 + 8 = **15 wywolan**

**c)** Funkcja znajduje **maksymalny element** tablicy metoda dziel i zwyciezaj. Dla dowolnej tablicy zwroci jej element najwiekszy.

Weryfikacja: max(3, 7, 2, 9, 1, 5, 8, 4) = 9. Poprawne.
</details>

<details>
<summary>Typowe bledy</summary>

- **Bledny obliczenie srodka**: srodek = (lewy+prawy) div 2, nie (lewy+prawy)/2. Dla (0+7) div 2 = 3, nie 3.5. CKE: -1 pkt
- **Pomylenie zakresu prawej polowy**: Prawa polowa to [srodek+1, prawy], nie [srodek, prawy] (co daloby nieskonczona rekurencje). CKE: -2 pkt
- **Niedokonczone drzewo**: Kazda galaz musi dojsc do warunku bazowego lewy>=prawy. CKE: -1 pkt
- **Pomylenie max i min**: Warunek `L > P` oznacza, ze zwracamy wieksza wartosc (maximum). CKE: -1 pkt

</details>

---

## Samoocena

Po rozwiazaniu cwiczen bez podgladania odpowiedzi, okresl swoj poziom:

| Poziom | Opis | Wynik |
|--------|------|-------|
| Podstawowy | Potrafisz sledzic proste petle z mod/div i prosta rekurencje liniowa | 1-3 cwiczen bez pomocy |
| Dobry | Radzisz sobie z tablicami zliczajacymi, drzewami wywolan i sortowaniem | 4-6 cwiczen bez pomocy |
| Bardzo dobry | Sledzisz algorytmy z programowaniem dynamicznym, stosem i ONP | 7-8 cwiczen bez pomocy |
| Doskonaly | Bezblednie sledzisz dziel-i-zwyciezaj, zlozona rekurencje i plansze 2D | 9-10 cwiczen bez pomocy |

**Co dalej?**
- Poziom Podstawowy: Przerob cwiczenia 1.1, 1.2, 1.6 jeszcze raz. Wrocz do `cheatsheet_teoria.md` (sekcja: sledzenie algorytmu).
- Poziom Dobry: Skup sie na cwiczeniach 1.4, 1.7, 1.8. Przejdz do `02_projektowanie_algorytmu.md`.
- Poziom Bardzo dobry/Doskonaly: Przejdz do `03_analiza_algorytmu.md` i `09_zlozone.md` (implementacja zlozonych algorytmow).
