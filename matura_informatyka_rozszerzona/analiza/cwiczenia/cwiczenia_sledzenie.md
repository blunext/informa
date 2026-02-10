# Cwiczenia ze sledzenia algorytmow (krok po kroku)

Dedykowany zbior **24 cwiczen** pokrywajacych wszystkie 7 archetypow pseudokodu CKE
oraz dodatkowe wzorce maturalne. Kazde cwiczenie zawiera pseudokod, polecenie
z konkretnymi danymi wejsciowymi i pelna tabelke sledzenia w rozwiazaniu.

**Typ zadania**: `sledzenie_algorytmu` — pojawia sie 11/11 lat, lacznie 45 pkt.

**Poziomy trudnosci**: latwe (~2 pkt), srednie (~3 pkt), trudne (~4-5 pkt)

---

## Spis tresci

| Sekcja | Archetyp | Cwiczenia | Trudnosc |
|--------|----------|-----------|----------|
| 1. Petla po cyfrach | mod 10 / div 10 | 1-3 | latwe-trudne |
| 2. Palindrom / odwracanie | odwroc + sprawdz | 4-5 | latwe-srednie |
| 3. Systemy liczbowe | dec->base, Horner, dodawanie | 6-8 | latwe-trudne |
| 4. NWD Euklidesa | iteracyjny, rekurencyjny + NWW | 9-10 | latwe-srednie |
| 5. Zliczanie / tablica | histogram, ciag rosnacy | 11-12 | latwe-srednie |
| 6. Budowanie wyniku z cyfr | filtrowanie, porownywanie | 13-14 | latwe-srednie |
| 7. Rekurencja i stos | silnia, Fibonacci, preorder | 15-17 | latwe-trudne |
| 8. BONUS — wzorce maturalne | bisekcja, BST, kopiec, ... | 18-24 | srednie-trudne |

---

## Sekcja 1: Petla po cyfrach (Archetyp 1)

### Cwiczenie 1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 2

Dana jest funkcja:

```
funkcja sumaNieparzystych(n)
    suma := 0
    dopoki n > 0:
        cyfra := n mod 10
        jezeli cyfra mod 2 = 1:
            suma := suma + cyfra
        n := n div 10
    zwroc suma
```

**Polecenie**: Oblicz wartosc `sumaNieparzystych(n)` dla kazdej z podanych liczb.
Dla liczby 7531 wypelnij pelna tabele sledzenia.

| Lp. | n |
|-----|---|
| a)  | 7531 |
| b)  | 2048 |
| c)  | 19305 |

<details>
<summary>Odpowiedz</summary>

**a) sumaNieparzystych(7531)**

| Krok | n (pocz.) | cyfra | cyfra%2=1? | suma | n (koniec) |
|------|-----------|-------|------------|------|------------|
| start | 7531 | — | — | 0 | — |
| 1 | 7531 | 1 | tak | 0+1=1 | 753 |
| 2 | 753 | 3 | tak | 1+3=4 | 75 |
| 3 | 75 | 5 | tak | 4+5=9 | 7 |
| 4 | 7 | 7 | tak | 9+7=16 | 0 |

Wynik: **16**

**b) sumaNieparzystych(2048)**

- cyfra 8: parzysta, suma = 0
- cyfra 4: parzysta, suma = 0
- cyfra 0: parzysta, suma = 0
- cyfra 2: parzysta, suma = 0

Wynik: **0** (brak cyfr nieparzystych)

**c) sumaNieparzystych(19305)**

- cyfra 5: nieparzysta, suma = 5
- cyfra 0: parzysta, suma = 5
- cyfra 3: nieparzysta, suma = 5+3=8
- cyfra 9: nieparzysta, suma = 8+9=17
- cyfra 1: nieparzysta, suma = 17+1=18

Wynik: **18**
</details>

---

### Cwiczenie 2 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: wariant maturalny

Dana jest funkcja:

```
funkcja iloczynNiezerowych(n)
    iloczyn := 1
    bylaCyfra := FALSZ
    dopoki n > 0:
        cyfra := n mod 10
        jezeli cyfra <> 0:
            iloczyn := iloczyn * cyfra
            bylaCyfra := PRAWDA
        n := n div 10
    jezeli bylaCyfra = PRAWDA:
        zwroc iloczyn
    w przeciwnym razie:
        zwroc 0
```

**Polecenie**: Przesled algorytm krok po kroku dla podanych wartosci n.

| Lp. | n |
|-----|---|
| a)  | 3072 |
| b)  | 10001 |

<details>
<summary>Odpowiedz</summary>

**a) iloczynNiezerowych(3072)**

| Krok | n (pocz.) | cyfra | cyfra<>0? | iloczyn | bylaCyfra | n (koniec) |
|------|-----------|-------|-----------|---------|-----------|------------|
| start | 3072 | — | — | 1 | FALSZ | — |
| 1 | 3072 | 2 | tak | 1*2=2 | PRAWDA | 307 |
| 2 | 307 | 7 | tak | 2*7=14 | PRAWDA | 30 |
| 3 | 30 | 0 | nie | 14 | PRAWDA | 3 |
| 4 | 3 | 3 | tak | 14*3=42 | PRAWDA | 0 |

bylaCyfra = PRAWDA, wiec wynik: **42**

**b) iloczynNiezerowych(10001)**

| Krok | n (pocz.) | cyfra | cyfra<>0? | iloczyn | bylaCyfra | n (koniec) |
|------|-----------|-------|-----------|---------|-----------|------------|
| start | 10001 | — | — | 1 | FALSZ | — |
| 1 | 10001 | 1 | tak | 1*1=1 | PRAWDA | 1000 |
| 2 | 1000 | 0 | nie | 1 | PRAWDA | 100 |
| 3 | 100 | 0 | nie | 1 | PRAWDA | 10 |
| 4 | 10 | 0 | nie | 1 | PRAWDA | 1 |
| 5 | 1 | 1 | tak | 1*1=1 | PRAWDA | 0 |

Wynik: **1** (iloczyn samych jedynek = 1, zera sa pomijane)
</details>

---

### Cwiczenie 3 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 1

Dana jest funkcja:

```
funkcja dopelnij(n)
    wynik := 0
    mnoznik := 1
    dopoki n > 0:
        cyfra := n mod 10
        nowa := 9 - cyfra
        wynik := wynik + nowa * mnoznik
        mnoznik := mnoznik * 10
        n := n div 10
    zwroc wynik
```

Funkcja zamienia kazda cyfre na jej dopelnienie do 9 (np. 3 -> 6, 0 -> 9).

**Polecenie**: Oblicz wartosc `dopelnij(n)` dla kazdej z podanych liczb.
Dla liczby 3816 wypelnij pelna tabele sledzenia.

| Lp. | n |
|-----|---|
| a)  | 3816 |
| b)  | 9000 |
| c)  | 12345 |

<details>
<summary>Odpowiedz</summary>

**a) dopelnij(3816)**

| Krok | n (pocz.) | cyfra | nowa (9-cyfra) | wynik | mnoznik | n (koniec) |
|------|-----------|-------|----------------|-------|---------|------------|
| start | 3816 | — | — | 0 | 1 | — |
| 1 | 3816 | 6 | 3 | 0+3*1=3 | 10 | 381 |
| 2 | 381 | 1 | 8 | 3+8*10=83 | 100 | 38 |
| 3 | 38 | 8 | 1 | 83+1*100=183 | 1000 | 3 |
| 4 | 3 | 3 | 6 | 183+6*1000=6183 | 10000 | 0 |

Wynik: **6183**

Sprawdzenie: 3816 + 6183 = 9999. Suma liczby i jej dopelnienia daje 999...9.

**b) dopelnij(9000)**

- cyfra 0 -> nowa 9, wynik = 9
- cyfra 0 -> nowa 9, wynik = 99
- cyfra 0 -> nowa 9, wynik = 999
- cyfra 9 -> nowa 0, wynik = 0*1000+999 = 999

Wynik: **999** (uwaga: wynik jest 3-cyfrowy, bo wiodaca cyfra to 0!)

Sprawdzenie: 9000 + 0999 = 9999.

**c) dopelnij(12345)**

- cyfra 5 -> nowa 4, wynik = 4
- cyfra 4 -> nowa 5, wynik = 54
- cyfra 3 -> nowa 6, wynik = 654
- cyfra 2 -> nowa 7, wynik = 7654
- cyfra 1 -> nowa 8, wynik = 87654

Wynik: **87654**

Sprawdzenie: 12345 + 87654 = 99999.
</details>

---

## Sekcja 2: Odwracanie liczby / palindrom (Archetyp 2)

### Cwiczenie 4 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: klasyczny

Dana jest funkcja:

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

**Polecenie**: Dla kazdej liczby podaj stan zmiennych `odwrocona` i `n` po kazdej iteracji petli, a nastepnie wynik funkcji.

| Lp. | n |
|-----|---|
| a)  | 12321 |
| b)  | 12345 |
| c)  | 1001 |

<details>
<summary>Odpowiedz</summary>

**a) czyPalindrom(12321)**

oryginal = 12321

| Krok | n (pocz.) | cyfra | odwrocona | n (koniec) |
|------|-----------|-------|-----------|------------|
| start | 12321 | — | 0 | — |
| 1 | 12321 | 1 | 0*10+1=1 | 1232 |
| 2 | 1232 | 2 | 1*10+2=12 | 123 |
| 3 | 123 | 3 | 12*10+3=123 | 12 |
| 4 | 12 | 2 | 123*10+2=1232 | 1 |
| 5 | 1 | 1 | 1232*10+1=12321 | 0 |

odwrocona (12321) = oryginal (12321) -> Wynik: **PRAWDA**

**b) czyPalindrom(12345)**

oryginal = 12345

| Krok | n (pocz.) | cyfra | odwrocona | n (koniec) |
|------|-----------|-------|-----------|------------|
| start | 12345 | — | 0 | — |
| 1 | 12345 | 5 | 5 | 1234 |
| 2 | 1234 | 4 | 54 | 123 |
| 3 | 123 | 3 | 543 | 12 |
| 4 | 12 | 2 | 5432 | 1 |
| 5 | 1 | 1 | 54321 | 0 |

odwrocona (54321) <> oryginal (12345) -> Wynik: **FALSZ**

**c) czyPalindrom(1001)**

oryginal = 1001

| Krok | n (pocz.) | cyfra | odwrocona | n (koniec) |
|------|-----------|-------|-----------|------------|
| start | 1001 | — | 0 | — |
| 1 | 1001 | 1 | 1 | 100 |
| 2 | 100 | 0 | 10 | 10 |
| 3 | 10 | 0 | 100 | 1 |
| 4 | 1 | 1 | 1001 | 0 |

odwrocona (1001) = oryginal (1001) -> Wynik: **PRAWDA**
</details>

---

### Cwiczenie 5 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: wariant "odwroc i dodaj"

Dana jest funkcja:

```
funkcja odwroc(n)
    odwrocona := 0
    dopoki n > 0:
        odwrocona := odwrocona * 10 + n mod 10
        n := n div 10
    zwroc odwrocona

funkcja odwrocIDodaj(n, maxKrokow)
    krok := 0
    dopoki krok < maxKrokow:
        r := odwroc(n)
        jezeli r = n:
            zwroc krok
        n := n + r
        krok := krok + 1
    zwroc -1
```

Funkcja `odwrocIDodaj` wielokrotnie odwraca liczbe i dodaje ja do siebie, az wynik stanie sie palindromem (lub przekroczy limit krokow).

**Polecenie**: Przesled `odwrocIDodaj(n, 5)` dla podanych wartosci n. Podaj wartosc n po kazdym kroku oraz wynik koncowy.

| Lp. | n |
|-----|---|
| a)  | 196 |
| b)  | 407 |

<details>
<summary>Odpowiedz</summary>

**a) odwrocIDodaj(196, 5)**

| krok | n (pocz.) | r = odwroc(n) | r = n? | n := n + r |
|------|-----------|---------------|--------|------------|
| 0 | 196 | 691 | nie | 196+691=887 |
| 1 | 887 | 788 | nie | 887+788=1675 |
| 2 | 1675 | 5761 | nie | 1675+5761=7436 |
| 3 | 7436 | 6347 | nie | 7436+6347=13783 |
| 4 | 13783 | 38731 | nie | 13783+38731=52514 |

Osiagnieto maxKrokow=5, wynik: **-1** (nie znaleziono palindromu w 5 krokach).

Uwaga: 196 to slynna liczba — prawdopodobnie NIGDY nie tworzy palindromu (nieudowodnione).

**b) odwrocIDodaj(407, 5)**

| krok | n (pocz.) | r = odwroc(n) | r = n? | n := n + r |
|------|-----------|---------------|--------|------------|
| 0 | 407 | 704 | nie | 407+704=1111 |
| 1 | 1111 | 1111 | TAK! | — |

r = n (1111 jest palindromem), wiec wynik: **1** (znaleziono po 1 kroku iteracji — ale zwraca wartosc `krok=1`).
</details>

---

## Sekcja 3: Konwersja systemu liczbowego (Archetyp 3)

### Cwiczenie 6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 3c

Dana jest funkcja:

```
funkcja naOsemkowy(n)
    wynik := 0
    mnoznik := 1
    dopoki n > 0:
        reszta := n mod 8
        wynik := wynik + reszta * mnoznik
        mnoznik := mnoznik * 10
        n := n div 8
    zwroc wynik
```

**Polecenie**: Oblicz wartosc `naOsemkowy(n)` dla podanych liczb.
Dla n=100 wypelnij pelna tabele sledzenia.

| Lp. | n |
|-----|---|
| a)  | 100 |
| b)  | 255 |

<details>
<summary>Odpowiedz</summary>

**a) naOsemkowy(100)**

| Krok | n (pocz.) | n mod 8 | wynik | mnoznik | n (koniec) |
|------|-----------|---------|-------|---------|------------|
| start | 100 | — | 0 | 1 | — |
| 1 | 100 | 4 | 0+4*1=4 | 10 | 12 |
| 2 | 12 | 4 | 4+4*10=44 | 100 | 1 |
| 3 | 1 | 1 | 44+1*100=144 | 1000 | 0 |

Wynik: **144** (100 w systemie osemkowym)

Sprawdzenie: 1*64 + 4*8 + 4*1 = 64+32+4 = 100.

**b) naOsemkowy(255)**

| Krok | n (pocz.) | n mod 8 | wynik | mnoznik | n (koniec) |
|------|-----------|---------|-------|---------|------------|
| start | 255 | — | 0 | 1 | — |
| 1 | 255 | 7 | 7 | 10 | 31 |
| 2 | 31 | 7 | 77 | 100 | 3 |
| 3 | 3 | 3 | 377 | 1000 | 0 |

Wynik: **377** (255 w systemie osemkowym)

Sprawdzenie: 3*64 + 7*8 + 7*1 = 192+56+7 = 255.
</details>

---

### Cwiczenie 7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 6

Dana jest funkcja (schemat Hornera):

```
funkcja naDziesietny(cyfry, dlug, baza)
    wynik := 0
    dla i := 0, 1, ..., dlug-1:
        wynik := wynik * baza + cyfry[i]
    zwroc wynik
```

**Polecenie**: Oblicz wartosc `naDziesietny(cyfry, dlug, baza)` dla podanych danych.

| Lp. | cyfry | dlug | baza |
|-----|-------|------|------|
| a)  | [2, 1, 0, 2] | 4 | 3 |
| b)  | [1, 1, 0, 1, 0] | 5 | 2 |

<details>
<summary>Odpowiedz</summary>

**a) naDziesietny([2, 1, 0, 2], 4, 3)** — trojkowy -> dziesietny

| Krok (i) | cyfry[i] | wynik := wynik*3 + cyfry[i] |
|----------|----------|------------------------------|
| start | — | 0 |
| i=0 | 2 | 0*3+2 = 2 |
| i=1 | 1 | 2*3+1 = 7 |
| i=2 | 0 | 7*3+0 = 21 |
| i=3 | 2 | 21*3+2 = 65 |

Wynik: **65**

Sprawdzenie: 2*27 + 1*9 + 0*3 + 2*1 = 54+9+0+2 = 65.

**b) naDziesietny([1, 1, 0, 1, 0], 5, 2)** — dwojkowy -> dziesietny

| Krok (i) | cyfry[i] | wynik := wynik*2 + cyfry[i] |
|----------|----------|------------------------------|
| start | — | 0 |
| i=0 | 1 | 0*2+1 = 1 |
| i=1 | 1 | 1*2+1 = 3 |
| i=2 | 0 | 3*2+0 = 6 |
| i=3 | 1 | 6*2+1 = 13 |
| i=4 | 0 | 13*2+0 = 26 |

Wynik: **26**

Sprawdzenie: 11010(2) = 16+8+0+2+0 = 26.
</details>

---

### Cwiczenie 8 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 5

Dana jest funkcja dodajaca dwie liczby w systemie trojkowym:

```
funkcja dodajTrojkowo(A, B, dlug)
    // A, B — tablice cyfr trojkowych (indeks 0 = najmniej znaczaca cyfra)
    // dlug — dlugosc tablic (obie tej samej dlugosci)
    przeniesienie := 0
    dla i := 0, 1, ..., dlug-1:
        s := A[i] + B[i] + przeniesienie
        C[i] := s mod 3
        przeniesienie := s div 3
    C[dlug] := przeniesienie
    zwroc C
```

**Polecenie**: Oblicz wynik dodawania 212(3) + 121(3).

Dane (najmniej znaczaca cyfra na indeksie 0):
- A = [2, 1, 2] (czyli 212 w systemie trojkowym, od prawej)
- B = [1, 2, 1] (czyli 121 w systemie trojkowym, od prawej)
- dlug = 3

<details>
<summary>Odpowiedz</summary>

**dodajTrojkowo([2,1,2], [1,2,1], 3)**

| Krok (i) | A[i] | B[i] | przeniesienie (wej.) | s = A[i]+B[i]+prz. | C[i] = s mod 3 | przeniesienie (wyj.) = s div 3 |
|----------|------|------|---------------------|---------------------|-----------------|-------------------------------|
| start | — | — | 0 | — | — | 0 |
| i=0 | 2 | 1 | 0 | 2+1+0=3 | 3 mod 3=0 | 3 div 3=1 |
| i=1 | 1 | 2 | 1 | 1+2+1=4 | 4 mod 3=1 | 4 div 3=1 |
| i=2 | 2 | 1 | 1 | 2+1+1=4 | 4 mod 3=1 | 4 div 3=1 |

C[3] := przeniesienie = 1

Wynik C = [0, 1, 1, 1] (od indeksu 0)

Czytajac od najwyzszego indeksu: **1110(3)**

Sprawdzenie dziesietne:
- 212(3) = 2*9 + 1*3 + 2*1 = 18+3+2 = 23
- 121(3) = 1*9 + 2*3 + 1*1 = 9+6+1 = 16
- 23 + 16 = 39
- 1110(3) = 1*27 + 1*9 + 1*3 + 0*1 = 27+9+3+0 = 39. Zgadza sie.
</details>

---

## Sekcja 4: NWD Euklidesa (Archetyp 4)

### Cwiczenie 9 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 zad. 3

Dana jest funkcja:

```
funkcja NWD(a, b)
    dopoki b <> 0:
        temp := b
        b := a mod b
        a := temp
    zwroc a
```

**Polecenie**: Przesled algorytm dla podanych par (a, b). Wypelnij tabelke sledzenia.

| Lp. | a | b |
|-----|---|---|
| a)  | 84 | 36 |
| b)  | 17 | 5 |

<details>
<summary>Odpowiedz</summary>

**a) NWD(84, 36)**

| Krok | a (pocz.) | b (pocz.) | a mod b | temp | a (koniec) | b (koniec) |
|------|-----------|-----------|---------|------|------------|------------|
| 1 | 84 | 36 | 12 | 36 | 36 | 12 |
| 2 | 36 | 12 | 0 | 12 | 12 | 0 |

b = 0, wiec wynik: **NWD(84, 36) = 12**

Sprawdzenie: 84 = 12*7, 36 = 12*3.

**b) NWD(17, 5)**

| Krok | a (pocz.) | b (pocz.) | a mod b | temp | a (koniec) | b (koniec) |
|------|-----------|-----------|---------|------|------------|------------|
| 1 | 17 | 5 | 2 | 5 | 5 | 2 |
| 2 | 5 | 2 | 1 | 2 | 2 | 1 |
| 3 | 2 | 1 | 0 | 1 | 1 | 0 |

b = 0, wiec wynik: **NWD(17, 5) = 1** (liczby sa wzglednie pierwsze)
</details>

---

### Cwiczenie 10 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4

Dane sa funkcje:

```
funkcja NWD(a, b)
    jezeli b = 0:
        zwroc a
    zwroc NWD(b, a mod b)

funkcja NWW(a, b)
    zwroc (a * b) div NWD(a, b)
```

**Polecenie**:
- a) Narysuj stos wywolan rekurencyjnych `NWD(120, 45)`. Podaj wynik.
- b) Oblicz `NWW(120, 45)`.

<details>
<summary>Odpowiedz</summary>

**a) Stos wywolan NWD(120, 45):**

```
NWD(120, 45)    // b<>0, wiec NWD(45, 120 mod 45) = NWD(45, 30)
  NWD(45, 30)   // b<>0, wiec NWD(30, 45 mod 30) = NWD(30, 15)
    NWD(30, 15)  // b<>0, wiec NWD(15, 30 mod 15) = NWD(15, 0)
      NWD(15, 0) // b=0, zwroc 15
```

Powrot: 15 -> 15 -> 15 -> 15

Wynik: **NWD(120, 45) = 15**

Liczba wywolan: **4**

**b) NWW(120, 45):**

NWW(120, 45) = (120 * 45) div NWD(120, 45) = 5400 div 15 = **360**

Sprawdzenie: 360 = 120*3 = 45*8.
Poprawniej: 360/120 = 3, 360/45 = 8. Obie dzielenia sa calkowite.
</details>

---

## Sekcja 5: Zliczanie z warunkiem / tablica (Archetyp 5)

### Cwiczenie 11 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 1

Dana jest funkcja:

```
funkcja maxWystapien(T, n)
    // T — tablica o indeksach 0..n-1
    // wartosci T[i] naleza do {0, 1, 2, 3}
    C[0] := 0; C[1] := 0; C[2] := 0; C[3] := 0
    dla i := 0, 1, ..., n-1:
        C[T[i]] := C[T[i]] + 1
    maks := C[0]
    indMaks := 0
    dla j := 1, 2, 3:
        jezeli C[j] > maks:
            maks := C[j]
            indMaks := j
    zwroc indMaks
```

**Polecenie**: Przesled algorytm dla T = [2, 0, 1, 2, 3, 2, 1, 0]. Podaj:
- a) Stan tablicy C po kazdej iteracji pierwszej petli.
- b) Wartosc zwracana przez funkcje.

<details>
<summary>Odpowiedz</summary>

**a) Pierwsza petla — budowanie histogramu:**

T = [2, 0, 1, 2, 3, 2, 1, 0], n = 8

| Krok (i) | T[i] | Operacja | C[0] | C[1] | C[2] | C[3] |
|----------|------|----------|------|------|------|------|
| start | — | — | 0 | 0 | 0 | 0 |
| i=0 | 2 | C[2]++ | 0 | 0 | 1 | 0 |
| i=1 | 0 | C[0]++ | 1 | 0 | 1 | 0 |
| i=2 | 1 | C[1]++ | 1 | 1 | 1 | 0 |
| i=3 | 2 | C[2]++ | 1 | 1 | 2 | 0 |
| i=4 | 3 | C[3]++ | 1 | 1 | 2 | 1 |
| i=5 | 2 | C[2]++ | 1 | 1 | 3 | 1 |
| i=6 | 1 | C[1]++ | 1 | 2 | 3 | 1 |
| i=7 | 0 | C[0]++ | 2 | 2 | 3 | 1 |

Stan koncowy C = [2, 2, 3, 1]

**Druga petla — szukanie maksimum:**

maks = C[0] = 2, indMaks = 0

| j | C[j] | C[j] > maks? | maks | indMaks |
|---|------|--------------|------|---------|
| 1 | 2 | 2 > 2? NIE | 2 | 0 |
| 2 | 3 | 3 > 2? TAK | 3 | 2 |
| 3 | 1 | 1 > 3? NIE | 3 | 2 |

**b) Wynik: indMaks = 2** (wartosc 2 wystapila najczesciej — 3 razy)
</details>

---

### Cwiczenie 12 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3

Dana jest funkcja:

```
funkcja najdluzszyRosnacy(T, n)
    maxDlug := 1
    aktDlug := 1
    dla i := 1, 2, ..., n-1:
        jezeli T[i] > T[i-1]:
            aktDlug := aktDlug + 1
            jezeli aktDlug > maxDlug:
                maxDlug := aktDlug
        w przeciwnym razie:
            aktDlug := 1
    zwroc maxDlug
```

**Polecenie**: Przesled algorytm dla T = [3, 5, 7, 2, 4, 8, 9, 1]. Podaj wartosc `maxDlug` i `aktDlug` po kazdej iteracji.

<details>
<summary>Odpowiedz</summary>

**najdluzszyRosnacy([3, 5, 7, 2, 4, 8, 9, 1], 8)**

| Krok (i) | T[i-1] | T[i] | T[i]>T[i-1]? | aktDlug | maxDlug |
|----------|--------|------|--------------|---------|---------|
| start | — | — | — | 1 | 1 |
| i=1 | 3 | 5 | tak | 2 | 2 |
| i=2 | 5 | 7 | tak | 3 | 3 |
| i=3 | 7 | 2 | nie | 1 | 3 |
| i=4 | 2 | 4 | tak | 2 | 3 |
| i=5 | 4 | 8 | tak | 3 | 3 |
| i=6 | 8 | 9 | tak | 4 | 4 |
| i=7 | 9 | 1 | nie | 1 | 4 |

Wynik: **maxDlug = 4** (najdluzszy ciag rosnacy: [2, 4, 8, 9])
</details>

---

## Sekcja 6: Budowanie wyniku z cyfr (Archetyp 6)

### Cwiczenie 13 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3

Dana jest funkcja:

```
funkcja parzysteCyfry(n)
    wynik := 0
    mnoznik := 1
    dopoki n > 0:
        cyfra := n mod 10
        jezeli cyfra mod 2 = 0:
            wynik := wynik + cyfra * mnoznik
            mnoznik := mnoznik * 10
        n := n div 10
    zwroc wynik
```

**Polecenie**: Oblicz wartosc `parzysteCyfry(n)` dla podanych liczb.

| Lp. | n |
|-----|---|
| a)  | 31628 |
| b)  | 13579 |

<details>
<summary>Odpowiedz</summary>

**a) parzysteCyfry(31628)**

| Krok | n (pocz.) | cyfra | cyfra%2=0? | wynik | mnoznik | n (koniec) |
|------|-----------|-------|------------|-------|---------|------------|
| start | 31628 | — | — | 0 | 1 | — |
| 1 | 31628 | 8 | tak | 0+8*1=8 | 10 | 3162 |
| 2 | 3162 | 2 | tak | 8+2*10=28 | 100 | 316 |
| 3 | 316 | 6 | tak | 28+6*100=628 | 1000 | 31 |
| 4 | 31 | 1 | nie | 628 | 1000 | 3 |
| 5 | 3 | 3 | nie | 628 | 1000 | 0 |

Wynik: **628** (cyfry parzyste z 31628, w oryginalnej kolejnosci)

**b) parzysteCyfry(13579)**

- cyfra 9: nieparzysta -> pomijamy
- cyfra 7: nieparzysta -> pomijamy
- cyfra 5: nieparzysta -> pomijamy
- cyfra 3: nieparzysta -> pomijamy
- cyfra 1: nieparzysta -> pomijamy

Wynik: **0** (brak cyfr parzystych)
</details>

---

### Cwiczenie 14 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: wariant maturalny

Dana jest funkcja:

```
funkcja wiekszeOdPoprzedniej(n)
    wynik := 0
    mnoznik := 1
    poprzednia := -1
    dopoki n > 0:
        cyfra := n mod 10
        jezeli cyfra > poprzednia:
            wynik := wynik + cyfra * mnoznik
            mnoznik := mnoznik * 10
        poprzednia := cyfra
        n := n div 10
    zwroc wynik
```

Uwaga: przetwarzamy cyfry od prawej do lewej. "Poprzednia" to cyfra na prawo od biezacej (przetworzona wczesniej).

**Polecenie**: Przesled algorytm dla n = 53172. Wypelnij pelna tabele sledzenia.

<details>
<summary>Odpowiedz</summary>

**wiekszeOdPoprzedniej(53172)**

| Krok | n (pocz.) | cyfra | poprzednia | cyfra > poprzednia? | wynik | mnoznik | n (koniec) |
|------|-----------|-------|------------|---------------------|-------|---------|------------|
| start | 53172 | — | -1 | — | 0 | 1 | — |
| 1 | 53172 | 2 | -1 | 2 > -1? tak | 0+2*1=2 | 10 | 5317 |
| 2 | 5317 | 7 | 2 | 7 > 2? tak | 2+7*10=72 | 100 | 531 |
| 3 | 531 | 1 | 7 | 1 > 7? nie | 72 | 100 | 53 |
| 4 | 53 | 3 | 1 | 3 > 1? tak | 72+3*100=372 | 1000 | 5 |
| 5 | 5 | 5 | 3 | 5 > 3? tak | 372+5*1000=5372 | 10000 | 0 |

Wynik: **5372**

Analiza: przetwarzamy od prawej. Cyfry (od prawej): 2, 7, 1, 3, 5.
- 2: wieksza od -1 (zawsze brana) -> bierzemy
- 7: wieksza od 2 -> bierzemy
- 1: NIE wieksza od 7 -> pomijamy
- 3: wieksza od 1 -> bierzemy
- 5: wieksza od 3 -> bierzemy

Wynikowe cyfry (w oryginalnej kolejnosci od lewej): 5, 3, 7, 2 -> **5372**
</details>

---

## Sekcja 7: Rekurencja i stos (Archetyp 7)

### Cwiczenie 15 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: klasyczny

Dana jest funkcja:

```
funkcja silnia(n)
    jezeli n <= 1:
        zwroc 1
    zwroc n * silnia(n - 1)
```

**Polecenie**: Dla n = 5:
- a) Narysuj stos wywolan rekurencyjnych.
- b) Podaj wartosc zwracana na kazdym poziomie.
- c) Podaj koncowy wynik i calkowita liczbe wywolan.

<details>
<summary>Odpowiedz</summary>

**a) i b) Stos wywolan silnia(5):**

```
silnia(5)                    // 5 * silnia(4)
  silnia(4)                  // 4 * silnia(3)
    silnia(3)                // 3 * silnia(2)
      silnia(2)              // 2 * silnia(1)
        silnia(1)            // n<=1, zwroc 1
        => zwraca 1
      => zwraca 2 * 1 = 2
    => zwraca 3 * 2 = 6
  => zwraca 4 * 6 = 24
=> zwraca 5 * 24 = 120
```

**c) Wynik: silnia(5) = 120**

Calkowita liczba wywolan: **5** (silnia(5), silnia(4), silnia(3), silnia(2), silnia(1))
</details>

---

### Cwiczenie 16 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: rozne matury

Dana jest funkcja:

```
funkcja fib(n)
    jezeli n <= 1:
        zwroc n
    zwroc fib(n - 1) + fib(n - 2)
```

**Polecenie**: Dla n = 6:
- a) Narysuj drzewo wywolan (ze wszystkimi galezami).
- b) Policz calkowita liczbe wywolan funkcji.
- c) Podaj wynik koncowy.

<details>
<summary>Odpowiedz</summary>

**a) Drzewo wywolan fib(6):**

```
fib(6) = fib(5) + fib(4)
├── fib(5) = fib(4) + fib(3)
│   ├── fib(4) = fib(3) + fib(2)
│   │   ├── fib(3) = fib(2) + fib(1)
│   │   │   ├── fib(2) = fib(1) + fib(0)
│   │   │   │   ├── fib(1) = 1
│   │   │   │   └── fib(0) = 0
│   │   │   │   => fib(2) = 1
│   │   │   └── fib(1) = 1
│   │   │   => fib(3) = 1 + 1 = 2
│   │   └── fib(2) = fib(1) + fib(0)
│   │       ├── fib(1) = 1
│   │       └── fib(0) = 0
│   │       => fib(2) = 1
│   │   => fib(4) = 2 + 1 = 3
│   └── fib(3) = fib(2) + fib(1)
│       ├── fib(2) = fib(1) + fib(0)
│       │   ├── fib(1) = 1
│       │   └── fib(0) = 0
│       │   => fib(2) = 1
│       └── fib(1) = 1
│       => fib(3) = 1 + 1 = 2
│   => fib(5) = 3 + 2 = 5
└── fib(4) = fib(3) + fib(2)
    ├── fib(3) = fib(2) + fib(1)
    │   ├── fib(2) = fib(1) + fib(0)
    │   │   ├── fib(1) = 1
    │   │   └── fib(0) = 0
    │   │   => fib(2) = 1
    │   └── fib(1) = 1
    │   => fib(3) = 1 + 1 = 2
    └── fib(2) = fib(1) + fib(0)
        ├── fib(1) = 1
        └── fib(0) = 0
        => fib(2) = 1
    => fib(4) = 2 + 1 = 3
=> fib(6) = 5 + 3 = 8
```

**b) Calkowita liczba wywolan:**

| Wywolanie | Ile razy |
|-----------|----------|
| fib(6) | 1 |
| fib(5) | 1 |
| fib(4) | 2 |
| fib(3) | 3 |
| fib(2) | 5 |
| fib(1) | 8 |
| fib(0) | 5 |

Razem: 1+1+2+3+5+8+5 = **25 wywolan**

**c) Wynik: fib(6) = 8**

Ciag Fibonacciego: 0, 1, 1, 2, 3, 5, **8**, 13, ...
</details>

---

### Cwiczenie 17 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 1.3

Dane jest drzewo binarne:

```
         10
        /  \
       5    15
      / \     \
     3   7    20
    /
   1
```

i algorytm preorder iteracyjny ze stosem:

```
funkcja preorder(korzen)
    jezeli korzen = NULL:
        zwroc ""
    wynik := ""
    stos := pusty stos
    stos.push(korzen)
    dopoki stos nie jest pusty:
        wezel := stos.pop()
        wynik := wynik + str(wezel.w) + " "
        jezeli wezel.R <> NULL:
            stos.push(wezel.R)
        jezeli wezel.L <> NULL:
            stos.push(wezel.L)
    zwroc wynik
```

**Polecenie**: Przesled algorytm krok po kroku. Podaj stan stosu i wynik po kazdej iteracji.

<details>
<summary>Odpowiedz</summary>

**preorder(korzen) — korzen = wezel(10)**

| Krok | stos.pop() -> wezel | wynik (po dopisaniu) | push R? | push L? | stos (po kroku, wierzch na prawo) |
|------|---------------------|----------------------|---------|---------|----------------------------------|
| init | — | "" | — | — | [10] |
| 1 | 10 | "10 " | R=15, push | L=5, push | [15, 5] |
| 2 | 5 | "10 5 " | R=7, push | L=3, push | [15, 7, 3] |
| 3 | 3 | "10 5 3 " | R=NULL, nie | L=1, push | [15, 7, 1] |
| 4 | 1 | "10 5 3 1 " | R=NULL, nie | L=NULL, nie | [15, 7] |
| 5 | 7 | "10 5 3 1 7 " | R=NULL, nie | L=NULL, nie | [15] |
| 6 | 15 | "10 5 3 1 7 15 " | R=20, push | L=NULL, nie | [20] |
| 7 | 20 | "10 5 3 1 7 15 20 " | R=NULL, nie | L=NULL, nie | [] |

Stos pusty, koniec petli.

Wynik: **"10 5 3 1 7 15 20"** (porzadek preorder)

Kluczowa obserwacja: prawy potomek wkladamy na stos PRZED lewym, dzieki czemu lewy jest zdejmowany PIERWSZY (LIFO). To gwarantuje porzadek: korzen -> lewe poddrzewo -> prawe poddrzewo.
</details>

---

## Sekcja 8: BONUS — Wzorce maturalne

### Cwiczenie 18 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 2a

Dana jest funkcja (metoda bisekcji):

```
funkcja bisekcja(a, b, eps)
    // szukamy miejsca zerowego f(x) = x*x - 5 w przedziale [a, b]
    // zakladamy f(a) < 0 i f(b) > 0
    dopoki b - a > eps:
        sr := (a + b) div 2     // dzielenie calkowite (wersja uproszczona)
        fsr := sr * sr - 5
        jezeli fsr = 0:
            zwroc sr
        jezeli fsr < 0:
            a := sr
        w przeciwnym razie:
            b := sr
    zwroc a
```

Uwaga: To uproszczona wersja calkowitoliczbowa. W oryginalnej bisekcji uzywa sie zmiennoprzecinkowych.

**Polecenie**: Przesled `bisekcja(0, 100, 1)` (szukamy sqrt(5) w liczbach calkowitych). Podaj wartosci a, b, sr i f(sr) po kazdej iteracji.

<details>
<summary>Odpowiedz</summary>

**bisekcja(0, 100, 1)**

f(x) = x*x - 5

| Krok | a | b | b-a > 1? | sr=(a+b) div 2 | f(sr) = sr*sr-5 | f(sr) < 0? | Nowe a/b |
|------|---|---|----------|----------------|-----------------|------------|----------|
| 1 | 0 | 100 | 100>1 tak | 50 | 2500-5=2495 | nie (>0) | b:=50 |
| 2 | 0 | 50 | 50>1 tak | 25 | 625-5=620 | nie (>0) | b:=25 |
| 3 | 0 | 25 | 25>1 tak | 12 | 144-5=139 | nie (>0) | b:=12 |
| 4 | 0 | 12 | 12>1 tak | 6 | 36-5=31 | nie (>0) | b:=6 |
| 5 | 0 | 6 | 6>1 tak | 3 | 9-5=4 | nie (>0) | b:=3 |
| 6 | 0 | 3 | 3>1 tak | 1 | 1-5=-4 | tak (<0) | a:=1 |
| 7 | 1 | 3 | 2>1 tak | 2 | 4-5=-1 | tak (<0) | a:=2 |
| 8 | 2 | 3 | 1>1? NIE | — | — | — | — |

Wynik: **a = 2** (sqrt(5) ~ 2.236, wiec calkowita czesc to 2)
</details>

---

### Cwiczenie 19 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 2

Dana jest funkcja (partycjonowanie Lomuto):

```
funkcja partycja(T, lewy, prawy)
    pivot := T[prawy]
    i := lewy
    dla j := lewy, lewy+1, ..., prawy-1:
        jezeli T[j] <= pivot:
            zamien T[i] z T[j]
            i := i + 1
    zamien T[i] z T[prawy]
    zwroc i
```

**Polecenie**: Przesled `partycja(T, 0, 6)` dla T = [5, 3, 8, 1, 7, 2, 6].
Podaj stan tablicy i wartosc `i` po kazdej iteracji petli.

<details>
<summary>Odpowiedz</summary>

**partycja([5, 3, 8, 1, 7, 2, 6], 0, 6)**

pivot = T[6] = 6, i = 0

| j | T[j] | T[j]<=6? | Zamiana | i (po) | Stan tablicy T |
|---|------|----------|--------|--------|----------------|
| start | — | — | — | 0 | [5, 3, 8, 1, 7, 2, 6] |
| 0 | 5 | 5<=6 tak | T[0]<->T[0] (bez zmian) | 1 | [5, 3, 8, 1, 7, 2, 6] |
| 1 | 3 | 3<=6 tak | T[1]<->T[1] (bez zmian) | 2 | [5, 3, 8, 1, 7, 2, 6] |
| 2 | 8 | 8<=6 nie | — | 2 | [5, 3, 8, 1, 7, 2, 6] |
| 3 | 1 | 1<=6 tak | T[2]<->T[3]: 8<->1 | 3 | [5, 3, 1, 8, 7, 2, 6] |
| 4 | 7 | 7<=6 nie | — | 3 | [5, 3, 1, 8, 7, 2, 6] |
| 5 | 2 | 2<=6 tak | T[3]<->T[5]: 8<->2 | 4 | [5, 3, 1, 2, 7, 8, 6] |

Po petli: zamien T[i] z T[prawy], czyli T[4]<->T[6]: 7<->6

Stan koncowy: **[5, 3, 1, 2, 6, 8, 7]**

Wynik (pozycja pivota): **i = 4**

Wlasciwosc: wszystkie elementy na lewo od T[4]=6 sa <= 6, wszystkie na prawo sa > 6.
</details>

---

### Cwiczenie 20 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 2

Dana jest funkcja rekurencyjna generujaca napisy z liter {a, b}:

```
funkcja generuj(napis, maxA, maxB)
    jezeli maxA = 0 AND maxB = 0:
        wypisz napis
        zwroc
    jezeli maxA > 0:
        generuj(napis + "a", maxA - 1, maxB)
    jezeli maxB > 0:
        generuj(napis + "b", maxA, maxB - 1)
```

**Polecenie**: Przesled `generuj("", 2, 1)`. Narysuj drzewo wywolan i podaj wszystkie wypisane napisy w kolejnosci.

<details>
<summary>Odpowiedz</summary>

**Drzewo wywolan generuj("", 2, 1):**

```
generuj("", 2, 1)
├── generuj("a", 1, 1)                    [maxA>0, dolacz "a"]
│   ├── generuj("aa", 0, 1)               [maxA>0, dolacz "a"]
│   │   ├── [maxA=0, pomijamy galaz "a"]
│   │   └── generuj("aab", 0, 0)          [maxB>0, dolacz "b"]
│   │       └── maxA=0 AND maxB=0 -> wypisz "aab"
│   └── generuj("ab", 1, 0)               [maxB>0, dolacz "b"]
│       ├── generuj("aba", 0, 0)           [maxA>0, dolacz "a"]
│       │   └── maxA=0 AND maxB=0 -> wypisz "aba"
│       └── [maxB=0, pomijamy galaz "b"]
└── generuj("b", 2, 0)                    [maxB>0, dolacz "b"]
    ├── generuj("ba", 1, 0)                [maxA>0, dolacz "a"]
    │   ├── generuj("baa", 0, 0)           [maxA>0, dolacz "a"]
    │   │   └── maxA=0 AND maxB=0 -> wypisz "baa"
    │   └── [maxB=0, pomijamy galaz "b"]
    └── [maxB=0, pomijamy galaz "b"]
```

**Wypisane napisy (w kolejnosci):**
1. **aab**
2. **aba**
3. **baa**

To wszystkie 3-literowe napisy z dokladnie 2 literami "a" i 1 litera "b".

Liczba takich napisow: C(3,1) = 3 (wybieramy pozycje dla "b").

Calkowita liczba wywolan: **10** (7 rekurencyjnych + 3 bazowe wypisujace)
</details>

---

### Cwiczenie 21 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 2

Dana jest procedura wstawiania do kopca typu max (kopiec jako tablica):

```
funkcja wstawDoKopca(kopiec, rozmiar, wartosc)
    // kopiec — tablica indeksowana od 1
    // rozmiar — liczba elementow przed wstawieniem
    rozmiar := rozmiar + 1
    kopiec[rozmiar] := wartosc
    i := rozmiar
    dopoki i > 1:
        rodzic := i div 2
        jezeli kopiec[i] > kopiec[rodzic]:
            zamien kopiec[i] z kopiec[rodzic]
            i := rodzic
        w przeciwnym razie:
            zwroc rozmiar
    zwroc rozmiar
```

**Polecenie**: Wstaw kolejno wartosci 10, 4, 15, 20, 3, 25 do pustego kopca (zaczynamy od rozmiar = 0). Po kazdym wstawieniu podaj stan tablicy kopca.

<details>
<summary>Odpowiedz</summary>

**Wstawianie kolejnych elementow:**

**Wstaw 10**: kopiec[1] = 10, i=1, i<=1 wiec koniec.
Kopiec: **[10]**

**Wstaw 4**: kopiec[2] = 4, i=2.
- rodzic = 2 div 2 = 1, kopiec[2]=4 > kopiec[1]=10? NIE -> koniec.
Kopiec: **[10, 4]**

**Wstaw 15**: kopiec[3] = 15, i=3.
- rodzic = 3 div 2 = 1, kopiec[3]=15 > kopiec[1]=10? TAK -> zamien.
  Kopiec: [15, 4, 10], i=1.
- i=1, i<=1 wiec koniec petli.
Kopiec: **[15, 4, 10]**

**Wstaw 20**: kopiec[4] = 20, i=4.
- rodzic = 4 div 2 = 2, kopiec[4]=20 > kopiec[2]=4? TAK -> zamien.
  Kopiec: [15, 20, 10, 4], i=2.
- rodzic = 2 div 2 = 1, kopiec[2]=20 > kopiec[1]=15? TAK -> zamien.
  Kopiec: [20, 15, 10, 4], i=1.
- i=1, i<=1 wiec koniec petli.
Kopiec: **[20, 15, 10, 4]**

**Wstaw 3**: kopiec[5] = 3, i=5.
- rodzic = 5 div 2 = 2, kopiec[5]=3 > kopiec[2]=15? NIE -> koniec.
Kopiec: **[20, 15, 10, 4, 3]**

**Wstaw 25**: kopiec[6] = 25, i=6.
- rodzic = 6 div 2 = 3, kopiec[6]=25 > kopiec[3]=10? TAK -> zamien.
  Kopiec: [20, 15, 25, 4, 3, 10], i=3.
- rodzic = 3 div 2 = 1, kopiec[3]=25 > kopiec[1]=20? TAK -> zamien.
  Kopiec: [25, 15, 20, 4, 3, 10], i=1.
- i=1, i<=1 wiec koniec petli.
Kopiec: **[25, 15, 20, 4, 3, 10]**

**Koncowy kopiec (tablica od indeksu 1):**

```
         25
        /  \
      15    20
     / \   /
    4   3 10
```

Wlasciwosc max-kopca: kazdy rodzic >= dzieci. Sprawdzenie:
- 25 >= 15, 25 >= 20
- 15 >= 4, 15 >= 3
- 20 >= 10
</details>

---

### Cwiczenie 22 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 1

Dana jest procedura wstawiania do BST (drzewa poszukiwan binarnych):

```
funkcja wstawBST(korzen, wartosc)
    nowyWezel := utworz wezel z wartoscia wartosc, L=NULL, R=NULL
    jezeli korzen = NULL:
        zwroc nowyWezel
    wezel := korzen
    dopoki PRAWDA:
        jezeli wartosc < wezel.w:
            jezeli wezel.L = NULL:
                wezel.L := nowyWezel
                zwroc korzen
            wezel := wezel.L
        w przeciwnym razie:
            jezeli wezel.R = NULL:
                wezel.R := nowyWezel
                zwroc korzen
            wezel := wezel.R
```

**Polecenie**: Wstaw kolejno wartosci 8, 3, 10, 1, 6, 14, 4, 7, 13 do pustego BST.
- a) Narysuj drzewo po kazdym wstawieniu 4 pierwszych elementow, a nastepnie koncowe drzewo.
- b) Podaj elementy w porzadku inorder (posortowane rosnaco).

<details>
<summary>Odpowiedz</summary>

**a) Budowanie BST krok po kroku:**

Wstaw 8 (korzen = NULL, nowy korzen):
```
8
```

Wstaw 3 (3 < 8, idz L, L=NULL -> wstaw):
```
  8
 /
3
```

Wstaw 10 (10 >= 8, idz R, R=NULL -> wstaw):
```
  8
 / \
3   10
```

Wstaw 1 (1 < 8 -> L=3, 1 < 3 -> L=NULL -> wstaw):
```
    8
   / \
  3   10
 /
1
```

Wstaw 6 (6 < 8 -> L=3, 6 >= 3 -> R=NULL -> wstaw):
```
    8
   / \
  3   10
 / \
1   6
```

Wstaw 14 (14 >= 8 -> R=10, 14 >= 10 -> R=NULL -> wstaw):
```
    8
   / \
  3   10
 / \    \
1   6   14
```

Wstaw 4 (4 < 8 -> 3, 4 >= 3 -> 6, 4 < 6 -> L=NULL -> wstaw):
```
    8
   / \
  3   10
 / \    \
1   6   14
   /
  4
```

Wstaw 7 (4 < 8 -> 3, 7 >= 3 -> 6, 7 >= 6 -> R=NULL -> wstaw):
```
    8
   / \
  3   10
 / \    \
1   6   14
   / \
  4   7
```

Wstaw 13 (13 >= 8 -> 10, 13 >= 10 -> 14, 13 < 14 -> L=NULL -> wstaw):

**Koncowe drzewo:**
```
       8
      / \
    3    10
   / \     \
  1   6    14
     / \   /
    4   7 13
```

**b) Inorder (lewy -> korzen -> prawy):**

Odwiedzamy: 1, 3, 4, 6, 7, 8, 10, 13, 14

Wynik inorder: **1 3 4 6 7 8 10 13 14** (posortowane rosnaco — wlasciwosc BST)
</details>

---

### Cwiczenie 23 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 1

Dana jest funkcja przetwarzajaca pary cyfr:

```
funkcja przetworzPary(n)
    wynik := 0
    mnoznik := 1
    dopoki n > 0:
        para := n mod 100
        suma := (para div 10) + (para mod 10)
        wynik := wynik + suma * mnoznik
        mnoznik := mnoznik * 10
        n := n div 100
    zwroc wynik
```

Funkcja bierze po dwie cyfry (od prawej), oblicza sume cyfr w kazdej parze i buduje wynik.

**Polecenie**: Oblicz wartosc `przetworzPary(n)` dla podanych liczb.

| Lp. | n |
|-----|---|
| a)  | 123456 |
| b)  | 7890 |

<details>
<summary>Odpowiedz</summary>

**a) przetworzPary(123456)**

| Krok | n (pocz.) | para = n mod 100 | cyfra_L = para div 10 | cyfra_R = para mod 10 | suma | wynik | mnoznik | n (koniec) |
|------|-----------|-------------------|----------------------|-----------------------|------|-------|---------|------------|
| start | 123456 | — | — | — | — | 0 | 1 | — |
| 1 | 123456 | 56 | 5 | 6 | 11 | 0+11*1=11 | 10 | 1234 |
| 2 | 1234 | 34 | 3 | 4 | 7 | 11+7*10=81 | 100 | 12 |
| 3 | 12 | 12 | 1 | 2 | 3 | 81+3*100=381 | 1000 | 0 |

Wynik: **381**

Weryfikacja:
- Para 56: 5+6=11
- Para 34: 3+4=7
- Para 12: 1+2=3
- Wynik = 381 (cyfry wyniku: 3, 8, 1 — ale uwaga, 11 to dwie cyfry!)

Pulapka: suma pary moze byc > 9 (np. 5+6=11), co daje dwucyfrowa wartosc w jednej "pozycji". Dlatego mnoznik rosnie o *10 na kazda pare, a nie na kazda cyfre wyniku. Pozycja 1: 11*1=11 (zajmuje 2 cyfry dziesietne), pozycja 2: 7*10=70, pozycja 3: 3*100=300. Razem: 300+70+11 = 381.

**b) przetworzPary(7890)**

| Krok | n (pocz.) | para = n mod 100 | cyfra_L | cyfra_R | suma | wynik | mnoznik | n (koniec) |
|------|-----------|-------------------|---------|---------|------|-------|---------|------------|
| start | 7890 | — | — | — | — | 0 | 1 | — |
| 1 | 7890 | 90 | 9 | 0 | 9 | 0+9*1=9 | 10 | 78 |
| 2 | 78 | 78 | 7 | 8 | 15 | 9+15*10=159 | 100 | 0 |

Wynik: **159**

Weryfikacja: Para 90: 9+0=9. Para 78: 7+8=15. Wynik: 15*10+9 = 159.
</details>

---

### Cwiczenie 24 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 5

Dana jest funkcja dodajaca dwie liczby binarne:

```
funkcja dodajBinarnie(A, B, dlug)
    // A, B — tablice bitow (indeks 0 = najmniej znaczacy bit)
    przeniesienie := 0
    dla i := 0, 1, ..., dlug-1:
        s := A[i] + B[i] + przeniesienie
        C[i] := s mod 2
        przeniesienie := s div 2
    C[dlug] := przeniesienie
    zwroc C
```

**Polecenie**: Oblicz wynik dodawania 1101(2) + 0111(2).

Dane (najmniej znaczacy bit na indeksie 0):
- A = [1, 0, 1, 1] (czytajac normalnie od lewej: 1101)
- B = [1, 1, 1, 0] (czytajac normalnie od lewej: 0111)
- dlug = 4

<details>
<summary>Odpowiedz</summary>

**dodajBinarnie([1,0,1,1], [1,1,1,0], 4)**

| Krok (i) | A[i] | B[i] | przeniesienie (wej.) | s = A[i]+B[i]+prz. | C[i] = s mod 2 | przeniesienie (wyj.) = s div 2 |
|----------|------|------|---------------------|---------------------|-----------------|-------------------------------|
| start | — | — | 0 | — | — | 0 |
| i=0 | 1 | 1 | 0 | 1+1+0=2 | 2 mod 2=0 | 2 div 2=1 |
| i=1 | 0 | 1 | 1 | 0+1+1=2 | 2 mod 2=0 | 2 div 2=1 |
| i=2 | 1 | 1 | 1 | 1+1+1=3 | 3 mod 2=1 | 3 div 2=1 |
| i=3 | 1 | 0 | 1 | 1+0+1=2 | 2 mod 2=0 | 2 div 2=1 |

C[4] := przeniesienie = 1

Wynik C = [0, 0, 1, 0, 1] (od indeksu 0)

Czytajac od najwyzszego indeksu: **10100(2)**

Sprawdzenie dziesietne:
- 1101(2) = 8+4+0+1 = 13
- 0111(2) = 0+4+2+1 = 7
- 13 + 7 = 20
- 10100(2) = 16+0+4+0+0 = 20. Zgadza sie.
</details>

---

## Podsumowanie

### Statystyki zbioru

| Kategoria | Liczba cwiczen |
|-----------|---------------|
| Latwe (~2 pkt) | 7 |
| Srednie (~3 pkt) | 9 |
| Trudne (~4-5 pkt) | 8 |
| **Razem** | **24** |

### Pokrycie archetypow

| Archetyp | Cwiczenia | Pokrycie |
|----------|-----------|----------|
| 1. Petla po cyfrach | 1-3 | suma, iloczyn, dopelnienie |
| 2. Palindrom / odwracanie | 4-5 | sprawdzenie, odwroc-i-dodaj |
| 3. Systemy liczbowe | 6-8 | dec->oct, Horner, dodawanie trojkowe |
| 4. NWD Euklidesa | 9-10 | iteracyjny, rekurencyjny + NWW |
| 5. Zliczanie / tablica | 11-12 | histogram, ciag rosnacy |
| 6. Budowanie wyniku | 13-14 | filtrowanie parzystych, porownywanie sasiednich |
| 7. Rekurencja i stos | 15-17 | silnia, Fibonacci (drzewo), preorder iteracyjny |
| 8. Wzorce maturalne | 18-24 | bisekcja, partycja, generowanie, kopiec, BST, pary cyfr, dodawanie binarne |

### Inspiracje maturalne

| Rok | Cwiczenia |
|-----|-----------|
| 2014 | 6, 18 |
| 2015 | 9 |
| 2016 | 19 |
| 2019 | 10, 20 |
| 2021 | 3, 21 |
| 2022 | 11 |
| 2023 | 7, 12, 17, 22 |
| 2024 | 1, 13 |
| 2025 | 8, 23, 24 |

---

*Powiazane materialy:*
- `pseudokod_wzorce.md` — konwencje pseudokodu CKE, 7 archetypow, poradnik rysowania tabelek
- `cwiczenia_wg_typu/01_sledzenie_algorytmu.md` — 5 dodatkowych cwiczen (inne algorytmy)
- `strategia_egzaminacyjna.md` — TOP 14 algorytmow z kodem C++
