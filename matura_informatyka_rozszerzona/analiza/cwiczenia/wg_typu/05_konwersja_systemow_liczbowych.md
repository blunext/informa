# 05. Konwersja systemow liczbowych

Typ zadania: **konwersja_systemow_liczbowych**
Czestotliwosc: 10/12 lat | Laczna punktacja: 12 pkt
Kategoria: TEORIA

## Umiejetnosci cwiczone w tym zestawie

`dec-bin` `bin-dec` `bin-hex` `hex-bin` `grupowanie-bitow` `dodawanie-binarne` `odejmowanie-binarne` `przeniesienia` `system-osemkowy` `system-trojkowy` `porownywanie-systemow` `potegi-dwojki` `dzielenie-z-reszta` `uzupelnienie-do-2`

---

### Cwiczenie 5.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: styl maturalny podstawowy
**Tagi**: `dec-bin` `bin-dec` `dzielenie-z-reszta` `potegi-dwojki`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Dziesietny -> binarny: dziel przez 2 i zapisuj reszty. Binarny -> dziesietny: sumuj potegi dwojki.
2. **Podejscie**: Reszty czytaj od dolu do gory. Pamietaj potegi: 1, 2, 4, 8, 16, 32, 64, 128, 256.
3. **Kluczowy krok**: 255 = 2^8 - 1, wiec w binarnym to 8 jedynek. 128 = 2^7, wiec to 1 i 7 zer.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Odwrocona kolejnosc reszt**: Reszty czytamy od DOLU do gory, nie od gory do dolu. CKE: -1 pkt
- **Pominiecie kroku dzielenia**: Trzeba kontynuowac az iloraz = 0. CKE: -0.5 pkt
- **Bledne potegi dwojki**: Zapamietaj: 2^0=1, 2^1=2, 2^2=4, ..., 2^7=128, 2^8=256. CKE: -0.5 pkt

</details>

---

### Cwiczenie 5.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 3c (bin->hex)
**Tagi**: `bin-hex` `grupowanie-bitow`

Zamien podane liczby z systemu binarnego na szesnastkowy, stosujac metode grupowania po 4 bity (od prawej strony).

| Lp. | Liczba binarna |
|-----|----------------|
| a)  | 10110110(2) |
| b)  | 111101001011(2) |
| c)  | 1111111100000000(2) |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Grupuj bity po 4 od prawej. Dopelnij zerami z lewej jezeli trzeba.
2. **Podejscie**: Kazda grupe 4 bitow zamien na cyfre hex: 0000=0, ..., 1001=9, 1010=A, ..., 1111=F.
3. **Kluczowy krok**: Zapamietaj: 1010=A, 1011=B, 1100=C, 1101=D, 1110=E, 1111=F.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Grupowanie od lewej zamiast od prawej**: ZAWSZE grupuj od prawej strony, dopelniaj zerami z lewej. CKE: -1 pkt
- **Pomylka A-F**: A=10, B=11, C=12, D=13, E=14, F=15. CKE: -0.5 pkt
- **Brak sprawdzenia wyniku**: Zawsze konwertuj hex na decimal i porownaj z bin na decimal. CKE: nic, ale warto.

</details>

---

### Cwiczenie 5.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 5 (Dodawanie binarne)
**Tagi**: `dodawanie-binarne` `odejmowanie-binarne` `przeniesienia`

Wykonaj nastepujace dzialania w systemie binarnym. Pokaz przeniesienia/pozyczki w kazdej kolumnie. Sprawdz wynik konwertujac na system dziesietny.

| Lp. | Dzialanie |
|-----|-----------|
| a)  | 10110(2) + 11011(2) |
| b)  | 11111(2) + 1(2) |
| c)  | 100000(2) - 10011(2) |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dodawanie binarne: 0+0=0, 0+1=1, 1+1=10 (zapisz 0, przenosisz 1), 1+1+1=11 (zapisz 1, przenosisz 1).
2. **Podejscie**: Pracuj kolumna po kolumnie od prawej. Nie zapomnij o przeniesieniach.
3. **Kluczowy krok**: Odejmowanie: jezeli gorna cyfra < dolna, musisz pozyczyc z sasiedniej kolumny (jak w dec, ale pozyczka = 2, nie 10).

</details>

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
  10000
- 00122
```

Uzupelniamy: 100000 - 010011

Obliczenia kolumna po kolumnie (od prawej) z pozyczkami:
- kol. 0: 0-1: nie mozna, pozyczamy z kol.1
  - kol.1 = 0, kaskada pozyczek:
  - kol.5: 1 -> 0, daje 1 do kol.4, kol.4 staje sie 10(2)=2
  - kol.4: 2 -> 1, daje 1 do kol.3, kol.3 staje sie 10(2)=2
  - kol.3: 2 -> 1, daje 1 do kol.2, kol.2 staje sie 10(2)=2
  - kol.2: 2 -> 1, daje 1 do kol.1, kol.1 staje sie 10(2)=2
  - kol.1: 2 -> 1, daje 1 do kol.0, kol.0 staje sie 10(2)=2
  - kol.0: 2-1 = 1
- kol.1: 1-1 = 0
- kol.2: 1-0 = 1
- kol.3: 1-0 = 1
- kol.4: 1-1 = 0
- kol.5: 0-0 = 0

Wynik: **001101(2) = 1101(2)**

Sprawdzenie: 100000(2) = 32, 10011(2) = 19, roznica = 13.
1101(2) = 8+4+1 = 13.
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o przeniesieniu**: 1+1 w binarnym to 10, nie 2! CKE: -1 pkt
- **Kaskada pozyczek**: Przy odejmowaniu z zerami (np. 100000-10011) trzeba pozyczyc przez wiele kolumn. CKE: -1 pkt
- **Brak sprawdzenia w dziesietnym**: ZAWSZE konwertuj na dziesietny i sprawdz. CKE: nic, ale ratuje od bledow.

</details>

---

### Cwiczenie 5.4 (trudnosc: srednie, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 5 (porownywanie miedzy systemami)
**Tagi**: `system-osemkowy` `porownywanie-systemow` `grupowanie-bitow`

Dane sa trzy liczby zapisane w roznych systemach liczbowych:

- A = 157(8)
- B = 6F(16)
- C = 1101111(2)

**Polecenie**:
- a) Zamien wszystkie trzy liczby na system dziesietny. Pokaz obliczenia.
- b) Ktora z liczb A, B, C jest najwieksza?
- c) Zamien najwieksza liczbe na system osemkowy.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Kazda liczbe najpierw przelicz na dziesietny. Porownaj wyniki.
2. **Podejscie**: Osemkowy: cyfra*potega_8. Hex: cyfra*potega_16. Binarny: suma poteg 2.
3. **Kluczowy krok**: Bin -> osemkowy: grupuj po 3 bity od prawej.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Bledne potegi 8**: 8^0=1, 8^1=8, 8^2=64, 8^3=512. Czesty blad: 8^2=80. CKE: -0.5 pkt
- **F=16 zamiast F=15**: W hex: F=15, nie 16. CKE: -0.5 pkt
- **Brak dopelnienia przy grupowaniu**: 1101111 -> dopelnij: 001|101|111, nie 110|111|1. CKE: -0.5 pkt

</details>

---

### Cwiczenie 5.5 (trudnosc: trudne, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 6 (system trojkowy)
**Tagi**: `system-trojkowy` `dodawanie-binarne` `przeniesienia`

Wykonaj nastepujace dzialania w systemie trojkowym (o podstawie 3). Przeniesienie nastepuje przy sumie >= 3. Pokaz szczegolowe obliczenia kolumna po kolumnie. Sprawdz wynik konwertujac na system dziesietny.

| Lp. | Dzialanie |
|-----|-----------|
| a)  | 2012(3) + 1221(3) |
| b)  | 10000(3) - 122(3) |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: W systemie trojkowym: cyfry 0, 1, 2. Przeniesienie gdy suma >= 3 (dzielisz przez 3, reszta zostaje).
2. **Podejscie**: Tak jak dodawanie binarne, ale podstawa to 3. Np. 2+1=3=1*3+0, zapisz 0 i przenies 1.
3. **Kluczowy krok**: Potegi trojki: 3^0=1, 3^1=3, 3^2=9, 3^3=27, 3^4=81.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Przeniesienie przy >= 2 zamiast >= 3**: W systemie trojkowym przeniesienie nastepuje przy sumie >= 3, nie >= 2 (to binarny). CKE: -1 pkt
- **Cyfra 3 w zapisie trojkowym**: W systemie trojkowym dopuszczalne cyfry to 0, 1, 2. Cyfra 3 nie istnieje! CKE: -1 pkt
- **Bledne potegi trojki**: 3^0=1, 3^1=3, 3^2=9, 3^3=27, 3^4=81, 3^5=243. CKE: -0.5 pkt

</details>

---

### Cwiczenie 5.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: styl CKE, konwersja hex->dec i dec->hex
**Tagi**: `hex-bin` `dec-bin` `dzielenie-z-reszta`

**Polecenie A**: Zamien podane liczby z systemu szesnastkowego na dziesietny.

| Lp. | Liczba hex |
|-----|------------|
| a)  | 2A(16) |
| b)  | FF(16) |
| c)  | 1C8(16) |

**Polecenie B**: Zamien podane liczby z systemu dziesietnego na szesnastkowy (metoda dzielenia przez 16).

| Lp. | Liczba dziesietna |
|-----|-------------------|
| d)  | 200 |
| e)  | 500 |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Hex -> dec: cyfra * potega 16. Dec -> hex: dziel przez 16, zapisuj reszty.
2. **Podejscie**: Potegi 16: 16^0=1, 16^1=16, 16^2=256. Reszty 10-15 zapisuj jako A-F.
3. **Kluczowy krok**: 1C8(16) = 1*256 + C*16 + 8 = 256 + 12*16 + 8.

</details>

<details>
<summary>Odpowiedz</summary>

**a) 2A(16) -> dziesietny:**
2*16 + 10*1 = 32 + 10 = **42**

**b) FF(16) -> dziesietny:**
15*16 + 15*1 = 240 + 15 = **255**

**c) 1C8(16) -> dziesietny:**
1*256 + 12*16 + 8*1 = 256 + 192 + 8 = **456**

**d) 200(10) -> szesnastkowy:**

| Dzielenie | Iloraz | Reszta |
|-----------|--------|--------|
| 200 / 16 | 12 | 8 |
| 12 / 16 | 0 | 12 = C |

Wynik: **C8(16)**

Sprawdzenie: 12*16 + 8 = 200.

**e) 500(10) -> szesnastkowy:**

| Dzielenie | Iloraz | Reszta |
|-----------|--------|--------|
| 500 / 16 | 31 | 4 |
| 31 / 16 | 1 | 15 = F |
| 1 / 16 | 0 | 1 |

Wynik: **1F4(16)**

Sprawdzenie: 1*256 + 15*16 + 4 = 256 + 240 + 4 = 500.
</details>

<details>
<summary>Typowe bledy</summary>

- **Reszta 12 zapisana jako "12" zamiast "C"**: W hex reszty >= 10 zapisuj jako A-F. CKE: -1 pkt
- **Bledna potega 16**: 16^2 = 256, nie 160. CKE: -0.5 pkt

</details>

---

### Cwiczenie 5.7 (trudnosc: srednie, ~2 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 5, konwersja wieloetapowa
**Tagi**: `bin-hex` `hex-bin` `system-osemkowy` `grupowanie-bitow`

Wykonaj nastepujace konwersje, uzywajac metod grupowania bitow (bez przechodzenia przez system dziesietny).

| Lp. | Konwersja |
|-----|-----------|
| a)  | 3A7(16) -> binarny |
| b)  | 101110011(2) -> osemkowy |
| c)  | 372(8) -> szesnastkowy |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Hex -> bin: kazda cyfre hex zamien na 4 bity. Bin -> oct: grupuj po 3 bity od prawej.
2. **Podejscie**: Oct -> hex: najpierw oct -> bin (grupuj po 3), potem bin -> hex (grupuj po 4).
3. **Kluczowy krok**: 3=0011, A=1010, 7=0111. Dla oct: 3=011, 7=111, 2=010.

</details>

<details>
<summary>Odpowiedz</summary>

**a) 3A7(16) -> binarny:**

Kazda cyfre hex zamieniamy na 4 bity:
- 3 = 0011
- A = 1010
- 7 = 0111

Wynik: **001110100111(2)** = 1110100111(2) (bez wiodacych zer)

Sprawdzenie: 3A7(16) = 3*256 + 10*16 + 7 = 768+160+7 = 935
1110100111(2) = 512+256+128+32+4+2+1 = 935

**b) 101110011(2) -> osemkowy:**

Grupujemy po 3 bity od prawej:
101 | 110 | 011
- 101 = 5
- 110 = 6
- 011 = 3

Wynik: **563(8)**

Sprawdzenie: 101110011(2) = 256+64+32+16+2+1 = 371
563(8) = 5*64 + 6*8 + 3 = 320+48+3 = 371

**c) 372(8) -> szesnastkowy:**

Krok 1: Oct -> bin (grupowanie po 3 bity):
- 3 = 011
- 7 = 111
- 2 = 010

Binarnie: 011111010(2)

Krok 2: Bin -> hex (grupowanie po 4 bity od prawej):
0 | 1111 | 1010
- 0000 (dopelniamy) + 0 = 0
- 1111 = F
- 1010 = A

Wynik: **FA(16)** (pomijamy wiodace zero)

Sprawdzenie: 372(8) = 3*64+7*8+2 = 192+56+2 = 250
FA(16) = 15*16+10 = 250
</details>

<details>
<summary>Typowe bledy</summary>

- **Grupowanie od lewej zamiast od prawej**: Przy bin->oct i bin->hex grupuj ZAWSZE od prawej. CKE: -1 pkt
- **Pominiety krok posredni (oct->hex)**: Nie da sie bezposrednio. Trzeba isc przez binarny. CKE: -0.5 pkt
- **Bledne grupy 3-bitowe**: 5=101, 6=110, 7=111 (nie myliz z 4-bitowymi). CKE: -0.5 pkt

</details>

---

### Cwiczenie 5.8 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 5 (dodawanie w hex)
**Tagi**: `dodawanie-binarne` `hex-bin` `przeniesienia`

Wykonaj nastepujace dzialania w systemie szesnastkowym. Pokaz obliczenia kolumna po kolumnie. Sprawdz wynik w systemie dziesietnym.

| Lp. | Dzialanie |
|-----|-----------|
| a)  | 3F(16) + A5(16) |
| b)  | FF(16) + 1(16) |
| c)  | 1A3(16) - B7(16) |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dodawanie hex: sumuj cyfry jak liczby (0-15). Przeniesienie gdy suma >= 16.
2. **Podejscie**: F+5 = 15+5 = 20 = 1*16 + 4, zapisz 4 i przenies 1.
3. **Kluczowy krok**: FF+1 = 100(16) — analogia do 99+1=100 w dziesietnym.

</details>

<details>
<summary>Odpowiedz</summary>

**a) 3F(16) + A5(16):**

- kol. 0: F+5 = 15+5 = 20 = 1*16 + 4, zapisz 4, przeniesienie 1
- kol. 1: 3+A+1 = 3+10+1 = 14 = E, przeniesienie 0

Wynik: **E4(16)**

Sprawdzenie: 3F = 63, A5 = 165, suma = 228. E4 = 14*16+4 = 228.

**b) FF(16) + 1(16):**

- kol. 0: F+1 = 15+1 = 16 = 1*16 + 0, zapisz 0, przeniesienie 1
- kol. 1: F+0+1 = 15+1 = 16 = 1*16 + 0, zapisz 0, przeniesienie 1
- kol. 2: 1 (przeniesienie)

Wynik: **100(16)**

Sprawdzenie: FF = 255, 1 = 1, suma = 256. 100(16) = 256.

**c) 1A3(16) - B7(16):**

Uzupelniamy: 1A3 - 0B7

- kol. 0: 3-7: nie mozna. Pozyczamy z kol.1: A -> 9, kol.0 dostaje +16. 3+16-7 = 12 = C
- kol. 1: 9-B = 9-11: nie mozna. Pozyczamy z kol.2: 1 -> 0, kol.1 dostaje +16. 9+16-11 = 14 = E
- kol. 2: 0-0 = 0

Wynik: **EC(16)**

Sprawdzenie: 1A3 = 256+160+3 = 419. B7 = 176+7 = 183. 419-183 = 236. EC = 14*16+12 = 236.
</details>

<details>
<summary>Typowe bledy</summary>

- **F+5 = 20, zapisz 20**: W hex zapisujemy reszte z dzielenia przez 16, nie cala wartosc. 20 = 1*16+4, zapisz 4. CKE: -1 pkt
- **Pozyczka = +10 zamiast +16**: W hex pozyczka daje +16 (podstawa systemu), nie +10. CKE: -1 pkt
- **Brak sprawdzenia**: Zawsze konwertuj na dziesietny i weryfikuj. CKE: nic, ale ratuje.

</details>

---

### Cwiczenie 5.9 (trudnosc: srednie-trudne, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 6 (wzorce bitowe, maski)
**Tagi**: `potegi-dwojki` `bin-dec` `porownywanie-systemow`

**Polecenie**: Odpowiedz na ponizsze pytania, uzasadniajac odpowiedz:

- a) Jaka jest najwieksza liczba, ktora mozna zapisac na 16 bitach (bez znaku)?
- b) Jaka jest najwieksza liczba, ktora mozna zapisac na 16 bitach w kodzie uzupelnienia do 2 (ze znakiem)?
- c) Ile roznych wartosci mozna przedstawic za pomoca 3 cyfr szesnastkowych?
- d) Jaki jest zakres wartosci 8-bitowej liczby ze znakiem w kodzie uzupelnienia do 2?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: n bitow bez znaku: zakres 0 do 2^n - 1. Ze znakiem (U2): -2^(n-1) do 2^(n-1) - 1.
2. **Podejscie**: 3 cyfry hex = 12 bitow. 16 bitow bez znaku: max = 2^16 - 1.
3. **Kluczowy krok**: U2 na 8 bitach: -128 do 127. Na 16 bitach: -32768 do 32767.

</details>

<details>
<summary>Odpowiedz</summary>

**a) Najwieksza na 16 bitach (bez znaku):**

16 bitow moze przechowac wartosci od 0 do 2^16 - 1.
2^16 = 65536, wiec max = **65535** = FFFF(16) = 1111111111111111(2)

**b) Najwieksza na 16 bitach (ze znakiem, U2):**

W kodzie U2 na n bitach: zakres od -2^(n-1) do 2^(n-1) - 1.
Dla n=16: od -2^15 do 2^15 - 1 = od -32768 do **32767**.

Najstarszy bit to bit znaku (0 = dodatnia, 1 = ujemna).
Najwieksza dodatnia: 0111111111111111(2) = 7FFF(16) = 32767.

**c) 3 cyfry hex:**

Kazda cyfra hex moze przyjac 16 wartosci (0-F).
3 cyfry: 16^3 = **4096** roznych wartosci (od 000 do FFF, czyli od 0 do 4095).

**d) 8-bitowa U2:**

Zakres: od -2^7 do 2^7 - 1 = od **-128 do 127**.

- Najmniejsza: 10000000(2) = -128
- Najwieksza: 01111111(2) = 127
- Lacznie: 256 wartosci (128 ujemnych + zero + 127 dodatnich)
</details>

<details>
<summary>Typowe bledy</summary>

- **Max bez znaku = 2^n zamiast 2^n - 1**: Na 8 bitach max to 255, nie 256. CKE: -1 pkt
- **U2: zakres symetryczny**: W U2 jest o 1 wiexcej liczb ujemnych niz dodatnich (-128..127, nie -127..127). CKE: -1 pkt
- **3 cyfry hex = 3*16 = 48**: Nie! 16^3 = 4096 (iloczyn, nie suma). CKE: -1 pkt

</details>

---

### Cwiczenie 5.10 (trudnosc: trudne, ~3 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 5 (wzorce bitowe, system niestandadowy)
**Tagi**: `uzupelnienie-do-2` `dodawanie-binarne` `potegi-dwojki`

**Polecenie**: Rozwiaz nastepujace problemy:

- a) Zapisz liczbe -45 w kodzie uzupelnienia do 2 (U2) na 8 bitach.
- b) Wykonaj dodawanie w U2 na 8 bitach: (-45) + 30. Sprawdz wynik.
- c) Dana jest liczba binarna 10110100(2) interpretowana w kodzie U2 na 8 bitach. Jaka wartosc dziesietna ona reprezentuje?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Aby zapisac -x w U2: zapisz x binarnie, zaneguj wszystkie bity, dodaj 1.
2. **Podejscie**: Dodawanie w U2 dziala tak samo jak zwykle dodawanie binarne — po prostu ignorujemy ostatnie przeniesienie.
3. **Kluczowy krok**: Interpretacja U2: jesli najstarszy bit = 1, to liczba ujemna. Wartosc = -(zaneguj bity + 1) lub -2^7 + reszta.

</details>

<details>
<summary>Odpowiedz</summary>

**a) -45 w U2 na 8 bitach:**

Krok 1: 45(10) w binarnym = 00101101(2)
Krok 2: Negacja bitow: 11010010
Krok 3: Dodaj 1: 11010010 + 1 = **11010011(2)**

Sprawdzenie: -128+64+16+2+1 = -128+83 = -45. Poprawne.

**b) (-45) + 30 w U2 na 8 bitach:**

-45 = 11010011(2)
30 = 00011110(2)

```
  przeniesienie: 1 1 1 1 1 0 0 0
                 1 1 0 1 0 0 1 1
               + 0 0 0 1 1 1 1 0
               -------------------
               (1)1 1 1 1 0 0 0 1
```

Odrzucamy przeniesienie poza 8 bitow.

Wynik: **11110001(2)**

Interpretacja U2: -128+64+32+16+1 = -128+113 = **-15**

Sprawdzenie: -45 + 30 = -15. Poprawne.

**c) 10110100(2) w U2:**

Najstarszy bit = 1, wiec liczba jest ujemna.

Metoda 1 (wagi pozycyjne):
-1*128 + 0*64 + 1*32 + 1*16 + 0*8 + 1*4 + 0*2 + 0*1
= -128 + 32 + 16 + 4 = **-76**

Metoda 2 (negacja + 1):
Negacja: 01001011
Dodaj 1: 01001100
01001100(2) = 64+8+4 = 76
Wiec oryginalna wartosc = **-76**
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o +1 po negacji**: U2 = negacja + 1, nie sama negacja (ta daje uzupelnienie do 1). CKE: -1 pkt
- **Nieodrzucenie przeniesienia**: W U2 na n bitach przeniesienie poza n-ty bit ignorujemy. CKE: -1 pkt
- **Bledna interpretacja U2**: Jezeli MSB=1, to wartosc = -2^(n-1) + reszta, NIE po prostu zamien na dec. CKE: -2 pkt

</details>

---

## Samoocena

Po rozwiazaniu cwiczen bez podgladania odpowiedzi, okresl swoj poziom:

| Poziom | Opis | Wynik |
|--------|------|-------|
| Podstawowy | Umiesz zamienic dec<->bin i bin<->hex metoda grupowania | 1-3 cwiczen bez pomocy |
| Dobry | Dodajesz i odejmujesz w binarnym, konwertujesz miedzy dowolnymi systemami | 4-6 cwiczen bez pomocy |
| Bardzo dobry | Radzisz sobie z hex arytmetyka, potegami dwojki i zakresami bitowymi | 7-8 cwiczen bez pomocy |
| Doskonaly | Biegle operujesz kodem U2, dodajesz w U2 i interpretujesz wzorce bitowe | 9-10 cwiczen bez pomocy |

**Co dalej?**
- Poziom Podstawowy: Przerob cwiczenia 5.1, 5.2, 5.6 jeszcze raz. Wrocz do `cheatsheet_teoria.md` (sekcja: systemy liczbowe).
- Poziom Dobry: Skup sie na cwiczeniach 5.3, 5.7, 5.8. Przejdz do `04_test_prawda_falsz.md`.
- Poziom Bardzo dobry/Doskonaly: Przejdz do `01_sledzenie_algorytmu.md` i `07_cyfry_liczby.md`.
