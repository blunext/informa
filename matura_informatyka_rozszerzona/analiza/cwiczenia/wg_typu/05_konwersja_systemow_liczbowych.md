# 05. Konwersja systemow liczbowych

Typ zadania: **konwersja_systemow_liczbowych**
Czestotliwosc: 9/11 lat | Laczna punktacja: 12 pkt
Kategoria: TEORIA

---

### Cwiczenie 5.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: styl maturalny podstawowy

**Polecenie A**: Zamien podane liczby z systemu dziesietnego na binarny. Pokaz pelne obliczenia metoda dzielenia z reszta.

| Lp. | Liczba dziesietna |
|-----|-------------------|
| a)  | 45 |
| b)  | 100 |
| c)  | 255 |

**Polecenie B**: Zamien podane liczby z systemu binarnego na dziesietny. Pokaz obliczenia.

| Lp. | Liczba binarna |
|-----|----------------|
| d)  | 110101(2) |
| e)  | 11111111(2) |
| f)  | 10000000(2) |

<details>
<summary>Odpowiedz</summary>

**a) 45(10) -> binarny:**

| Dzielenie | Iloraz | Reszta |
|-----------|--------|--------|
| 45 / 2 | 22 | 1 |
| 22 / 2 | 11 | 0 |
| 11 / 2 | 5 | 1 |
| 5 / 2 | 2 | 1 |
| 2 / 2 | 1 | 0 |
| 1 / 2 | 0 | 1 |

Czytamy reszty od dolu: 45(10) = **101101(2)**

Sprawdzenie: 1*32 + 0*16 + 1*8 + 1*4 + 0*2 + 1*1 = 32+8+4+1 = 45

**b) 100(10) -> binarny:**

| Dzielenie | Iloraz | Reszta |
|-----------|--------|--------|
| 100 / 2 | 50 | 0 |
| 50 / 2 | 25 | 0 |
| 25 / 2 | 12 | 1 |
| 12 / 2 | 6 | 0 |
| 6 / 2 | 3 | 0 |
| 3 / 2 | 1 | 1 |
| 1 / 2 | 0 | 1 |

100(10) = **1100100(2)**

Sprawdzenie: 64+32+4 = 100

**c) 255(10) -> binarny:**

| Dzielenie | Iloraz | Reszta |
|-----------|--------|--------|
| 255 / 2 | 127 | 1 |
| 127 / 2 | 63 | 1 |
| 63 / 2 | 31 | 1 |
| 31 / 2 | 15 | 1 |
| 15 / 2 | 7 | 1 |
| 7 / 2 | 3 | 1 |
| 3 / 2 | 1 | 1 |
| 1 / 2 | 0 | 1 |

255(10) = **11111111(2)**

Sprawdzenie: 128+64+32+16+8+4+2+1 = 255

**d) 110101(2) -> dziesietny:**

1*32 + 1*16 + 0*8 + 1*4 + 0*2 + 1*1 = 32 + 16 + 4 + 1 = **53**

**e) 11111111(2) -> dziesietny:**

128 + 64 + 32 + 16 + 8 + 4 + 2 + 1 = **255**

**f) 10000000(2) -> dziesietny:**

1*128 = **128**
</details>

---

### Cwiczenie 5.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 3c (bin->hex)

Zamien podane liczby z systemu binarnego na szesnastkowy, stosujac metode grupowania po 4 bity (od prawej strony).

| Lp. | Liczba binarna |
|-----|----------------|
| a)  | 10110110(2) |
| b)  | 111101001011(2) |
| c)  | 1111111100000000(2) |

<details>
<summary>Odpowiedz</summary>

**Metoda**: Grupujemy bity po 4 od prawej strony. Jezeli ostatnia grupa (od lewej) ma mniej niz 4 bity, dopelniamy zerami z lewej. Kazda grupe zamieniamy na cyfre szesnastkowa.

Tabela konwersji:
```
0000=0  0100=4  1000=8  1100=C
0001=1  0101=5  1001=9  1101=D
0010=2  0110=6  1010=A  1110=E
0011=3  0111=7  1011=B  1111=F
```

**a) 10110110(2):**

Grupowanie: `1011 | 0110`
- 1011 = B
- 0110 = 6

Wynik: **B6(16)**

Sprawdzenie: B6(16) = 11*16 + 6 = 182. 10110110(2) = 128+32+16+4+2 = 182.

**b) 111101001011(2):**

Grupowanie: `1111 | 0100 | 1011`
- 1111 = F
- 0100 = 4
- 1011 = B

Wynik: **F4B(16)**

Sprawdzenie: F4B(16) = 15*256 + 4*16 + 11 = 3840+64+11 = 3915.
111101001011(2) = 2048+1024+512+256+64+8+2+1 = 3915.

**c) 1111111100000000(2):**

Grupowanie: `1111 | 1111 | 0000 | 0000`
- 1111 = F
- 1111 = F
- 0000 = 0
- 0000 = 0

Wynik: **FF00(16)**

Sprawdzenie: FF00(16) = 15*4096 + 15*256 = 61440 + 3840 = 65280.
1111111100000000(2) = 32768+16384+8192+4096+2048+1024+512+256 = 65280.
</details>

---

### Cwiczenie 5.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 5 (Dodawanie binarne)

Wykonaj nastepujace dzialania w systemie binarnym. Pokaz przeniesienia/pozyczki w kazdej kolumnie. Sprawdz wynik konwertujac na system dziesietny.

| Lp. | Dzialanie |
|-----|-----------|
| a)  | 10110(2) + 11011(2) |
| b)  | 11111(2) + 1(2) |
| c)  | 100000(2) - 10011(2) |

<details>
<summary>Odpowiedz</summary>

**a) 10110(2) + 11011(2):**

```
  przeniesienie: 1 1 1 1 0
                 1 0 1 1 0
               + 1 1 0 1 1
               ---------
               1 1 0 0 0 1
```

Obliczenia kolumna po kolumnie (od prawej):
- kol. 0: 0+1=1, przeniesienie 0
- kol. 1: 1+1=10, zapisz 0, przeniesienie 1
- kol. 2: 1+0+1=10, zapisz 0, przeniesienie 1
- kol. 3: 0+1+1=10, zapisz 0, przeniesienie 1
- kol. 4: 1+1+1=11, zapisz 1, przeniesienie 1
- kol. 5: 1 (przeniesienie)

Wynik: **110001(2)**

Sprawdzenie: 10110(2) = 22, 11011(2) = 27, suma = 49.
110001(2) = 32+16+1 = 49.

**b) 11111(2) + 1(2):**

```
  przeniesienie: 1 1 1 1 1
                 1 1 1 1 1
               +         1
               ---------
               1 0 0 0 0 0
```

Obliczenia:
- kol. 0: 1+1=10, zapisz 0, przeniesienie 1
- kol. 1: 1+0+1=10, zapisz 0, przeniesienie 1
- kol. 2: 1+0+1=10, zapisz 0, przeniesienie 1
- kol. 3: 1+0+1=10, zapisz 0, przeniesienie 1
- kol. 4: 1+0+1=10, zapisz 0, przeniesienie 1
- kol. 5: 1 (przeniesienie)

Wynik: **100000(2)**

Sprawdzenie: 11111(2) = 31, 1(2) = 1, suma = 32.
100000(2) = 32.

**c) 100000(2) - 10011(2):**

```
  pozyczka:      0 1 1 0 0
                 1 0 0 0 0 0
               -   1 0 0 1 1
               -----------
                 0 0 1 1 0 1
```

Obliczenia kolumna po kolumnie (od prawej):
- kol. 0: 0-1, pozyczamy: 10-1=1
- kol. 1: 0-1-1(pozyczka)=-2, pozyczamy: 10-1-1=0. Ale mamy pozyczke z kol.0: 0-1=-1, pozyczamy: 10-1=0. Przeliczymy systematycznie:

Metoda krok po kroku:
```
  100000
-  10011
```
Uzupelniamy: 100000 - 010011

- kol. 0: 0-1: nie mozna, pozyczamy z kol.1. Ale kol.1 = 0, wiec kaskada pozyczek az do kol.5 (1):
  - kol.5: 1 -> 0, daje 1 do kol.4
  - kol.4: 0+10=10 -> 1, daje 1 do kol.3
  - kol.3: 0+10=10 -> 1, daje 1 do kol.2
  - kol.2: 0+10=10 -> 1, daje 1 do kol.1
  - kol.1: 0+10=10 -> 1, daje 1 do kol.0
  - kol.0: 0+10=10

Teraz odejmowanie:
- kol. 0: 10-1 = 1
- kol. 1: 1-1 = 0
- kol. 2: 1-0 = 1
- kol. 3: 1-0 = 1
- kol. 4: 1-1 = 0
- kol. 5: 0-0 = 0

Wynik: **001101(2) = 1101(2)**

Sprawdzenie: 100000(2) = 32, 10011(2) = 19, roznica = 13.
1101(2) = 8+4+1 = 13.
</details>

---

### Cwiczenie 5.4 (trudnosc: srednie, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 5 (porownywanie miedzy systemami)

Dane sa trzy liczby zapisane w roznych systemach liczbowych:

- A = 157(8)
- B = 6F(16)
- C = 1101111(2)

**Polecenie**:
- a) Zamien wszystkie trzy liczby na system dziesietny. Pokaz obliczenia.
- b) Ktora z liczb A, B, C jest najwieksza?
- c) Zamien najwieksza liczbe na system osemkowy.

<details>
<summary>Odpowiedz</summary>

**a) Konwersja na system dziesietny:**

A = 157(8):
A = 1*64 + 5*8 + 7*1 = 64 + 40 + 7 = **111(10)**

B = 6F(16):
B = 6*16 + 15*1 = 96 + 15 = **111(10)**

C = 1101111(2):
C = 64 + 32 + 8 + 4 + 2 + 1 = **111(10)**

**b) Wszystkie trzy liczby sa rowne: A = B = C = 111(10).**

Zaden z nich nie jest "najwiekszy" — sa identyczne.

**c) 111(10) w systemie osemkowym:**

| Dzielenie | Iloraz | Reszta |
|-----------|--------|--------|
| 111 / 8 | 13 | 7 |
| 13 / 8 | 1 | 5 |
| 1 / 8 | 0 | 1 |

111(10) = **157(8)** (co potwierdza wartosc A)

Alternatywna konwersja C -> osemkowy przez grupowanie po 3 bity:
1101111(2) = 001 | 101 | 111 = 1 | 5 | 7 = 157(8)
</details>

---

### Cwiczenie 5.5 (trudnosc: trudne, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 6 (system trojkowy)

Wykonaj nastepujace dzialania w systemie trojkowym (o podstawie 3). Przeniesienie nastepuje przy sumie >= 3. Pokaz szczegolowe obliczenia kolumna po kolumnie. Sprawdz wynik konwertujac na system dziesietny.

| Lp. | Dzialanie |
|-----|-----------|
| a)  | 2012(3) + 1221(3) |
| b)  | 10000(3) - 122(3) |

<details>
<summary>Odpowiedz</summary>

**a) 2012(3) + 1221(3):**

```
  przeniesienie: 1 1 1 0
                 2 0 1 2
               + 1 2 2 1
               ---------
               1 1 0 1 0
```

Obliczenia kolumna po kolumnie (od prawej):
- kol. 0: 2+1 = 3 = 1*3 + 0, zapisz 0, przeniesienie 1
- kol. 1: 1+2+1 = 4 = 1*3 + 1, zapisz 1, przeniesienie 1
- kol. 2: 0+2+1 = 3 = 1*3 + 0, zapisz 0, przeniesienie 1
- kol. 3: 2+1+1 = 4 = 1*3 + 1, zapisz 1, przeniesienie 1
- kol. 4: 1 (przeniesienie)

Wynik: **11010(3)**

Sprawdzenie:
- 2012(3) = 2*27 + 0*9 + 1*3 + 2*1 = 54 + 3 + 2 = 59
- 1221(3) = 1*27 + 2*9 + 2*3 + 1*1 = 27 + 18 + 6 + 1 = 52
- Suma: 59 + 52 = 111
- 11010(3) = 1*81 + 1*27 + 0*9 + 1*3 + 0*1 = 81 + 27 + 3 = 111

**b) 10000(3) - 122(3):**

Uzupelniamy: 10000 - 00122

```
  10000
- 00122
```

Obliczenia kolumna po kolumnie (od prawej) z pozyczkami:
- kol. 0: 0-2: nie mozna, pozyczamy z kol.1
  - kol.1 = 0, kaskada pozyczek:
  - kol.4: 1 -> 0, daje 1 do kol.3, kol.3 staje sie 10(3)=3
  - kol.3: 3 -> 2, daje 1 do kol.2, kol.2 staje sie 10(3)=3
  - kol.2: 3 -> 2, daje 1 do kol.1, kol.1 staje sie 10(3)=3
  - kol.1: 3 -> 2, daje 1 do kol.0, kol.0 staje sie 10(3)=3
  - kol.0: 3-2 = 1
- kol.1: 2-2 = 0
- kol.2: 2-1 = 1
- kol.3: 2-0 = 2
- kol.4: 0-0 = 0

Wynik: **02101(3) = 2101(3)**

Sprawdzenie:
- 10000(3) = 1*81 = 81
- 122(3) = 1*9 + 2*3 + 2*1 = 9 + 6 + 2 = 17
- Roznica: 81 - 17 = 64
- 2101(3) = 2*27 + 1*9 + 0*3 + 1*1 = 54 + 9 + 1 = 64
</details>
