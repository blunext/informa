# 02. Projektowanie algorytmu

Typ zadania: **projektowanie_algorytmu**
Czestotliwosc: 11/11 lat | Laczna punktacja: 43 pkt
Kategoria: TEORIA

---

### Cwiczenie 2.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 2 (cyfry), styl CKE

Napisz w pseudokodzie lub C++ funkcje `parzyste_cyfry(n)`, ktora dla danej liczby naturalnej n > 0 zwraca liczbe jej cyfr parzystych (0, 2, 4, 6, 8).

**Ograniczenia**: Uzyj tylko zmiennych calkowitych i operatorow arytmetycznych (mod, div, +, porownania). Nie wolno uzywac stringow, tablic ani funkcji wbudowanych.

**Przyklady**:
- `parzyste_cyfry(24681)` = 4 (cyfry parzyste: 2, 4, 6, 8)
- `parzyste_cyfry(13579)` = 0 (brak cyfr parzystych)
- `parzyste_cyfry(1024)` = 3 (cyfry parzyste: 0, 2, 4)

<details>
<summary>Odpowiedz</summary>

**Pseudokod:**

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

**C++:**

```cpp
int parzyste_cyfry(int n) {
    int licznik = 0;
    while (n > 0) {
        int cyfra = n % 10;
        if (cyfra % 2 == 0)
            licznik++;
        n /= 10;
    }
    return licznik;
}
```

**Weryfikacja**:
- parzyste_cyfry(24681): cyfry 1,8,6,4,2 -> parzyste: 8,6,4,2 -> wynik 4
- parzyste_cyfry(13579): cyfry 9,7,5,3,1 -> parzyste: brak -> wynik 0
- parzyste_cyfry(1024): cyfry 4,2,0,1 -> parzyste: 4,2,0 -> wynik 3
</details>

---

### Cwiczenie 2.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 1.3 (przestaw iteracyjnie)

Dana jest funkcja rekurencyjna obliczajaca sume cyfr liczby:

```
funkcja sumaCyfr(n)
    jezeli n = 0:
        zwroc 0
    zwroc sumaCyfr(n div 10) + n mod 10
```

**Polecenie**: Napisz iteracyjna wersje tej funkcji. Uzyj petli `dopoki` (while). Nie uzywaj rekurencji.

<details>
<summary>Odpowiedz</summary>

**Pseudokod:**

```
funkcja sumaCyfrIter(n)
    suma := 0
    dopoki n > 0:
        suma := suma + n mod 10
        n := n div 10
    zwroc suma
```

**C++:**

```cpp
int sumaCyfrIter(int n) {
    int suma = 0;
    while (n > 0) {
        suma += n % 10;
        n /= 10;
    }
    return suma;
}
```

**Wyjasinienie konwersji**:

Rekurencja ogonowa (gdzie wynik rekurencyjnego wywolania jest bezposrednio zwracany z dodanym skladnikiem) zamieniamy mechanicznie:
1. Warunek bazowy `n = 0` staje sie warunkiem konca petli `n > 0`
2. Akumulator (suma) przechowuje wynik czesciowy
3. Zmiana argumentu `n div 10` staje sie aktualizacja zmiennej `n := n div 10`
4. Dodawanie `n mod 10` odbywa sie w kazdej iteracji

**Weryfikacja**:
- sumaCyfrIter(47): 7 + 4 = 11
- sumaCyfrIter(305): 5 + 0 + 3 = 8
- sumaCyfrIter(1234): 4 + 3 + 2 + 1 = 10
</details>

---

### Cwiczenie 2.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: styl CKE -- "tylko zmienne calkowite"

Napisz pseudokod funkcji `czyPalindrom(n)` sprawdzajacej, czy liczba naturalna n > 0 jest palindromem (czytana od lewej i od prawej daje ta sama liczbe, np. 12321, 4554, 7).

**Ograniczenia**: Nie wolno uzywac stringow, tablic ani funkcji wbudowanych. Dozwolone sa wylacznie zmienne calkowite i operatory arytmetyczne (mod, div, +, *, porownania).

**Przyklady**:
- `czyPalindrom(12321)` = PRAWDA
- `czyPalindrom(12345)` = FALSZ
- `czyPalindrom(4554)` = PRAWDA
- `czyPalindrom(7)` = PRAWDA

<details>
<summary>Odpowiedz</summary>

**Pseudokod (metoda: odwroc liczbe i porownaj):**

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

**C++:**

```cpp
bool czyPalindrom(int n) {
    int oryginal = n;
    int odwrocona = 0;
    while (n > 0) {
        odwrocona = odwrocona * 10 + n % 10;
        n /= 10;
    }
    return odwrocona == oryginal;
}
```

**Sledzenie dla n = 12321:**

| Krok | n | cyfra | odwrocona |
|------|---|-------|-----------|
| poczatek | 12321 | - | 0 |
| 1 | 1232 | 1 | 1 |
| 2 | 123 | 2 | 12 |
| 3 | 12 | 3 | 123 |
| 4 | 1 | 2 | 1232 |
| 5 | 0 | 1 | 12321 |

odwrocona (12321) = oryginal (12321) -> PRAWDA

**Sledzenie dla n = 12345:**

odwrocona po petli: 54321
54321 != 12345 -> FALSZ
</details>

---

### Cwiczenie 2.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3 (pseudokod), Matura 2024 zad. 6 (systemy)

Napisz pseudokod funkcji `konwertuj(n, k)`, ktora zwraca zapis liczby naturalnej n > 0 w systemie o podstawie k (2 <= k <= 9) jako liczbe calkowita.

Na przyklad: `konwertuj(13, 2)` zwraca 1101, bo 13 w systemie dwojkowym to 1101.

**Ograniczenia**: Uzyj tylko zmiennych calkowitych i operatorow arytmetycznych. Wynik ma byc liczba calkowita (nie string). Zakladamy, ze wynik miesci sie w zakresie typu calkowitego.

**Przyklady**:
- `konwertuj(13, 2)` = 1101
- `konwertuj(100, 8)` = 144
- `konwertuj(25, 3)` = 221

<details>
<summary>Odpowiedz</summary>

**Pseudokod:**

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

**C++:**

```cpp
int konwertuj(int n, int k) {
    int wynik = 0;
    int mnoznik = 1;
    while (n > 0) {
        int reszta = n % k;
        wynik += reszta * mnoznik;
        mnoznik *= 10;
        n /= k;
    }
    return wynik;
}
```

**Weryfikacja konwertuj(13, 2):**

| Krok | n | n%2 | wynik | mnoznik | n (po /2) |
|------|---|-----|-------|---------|-----------|
| 1 | 13 | 1 | 0+1*1=1 | 10 | 6 |
| 2 | 6 | 0 | 1+0*10=1 | 100 | 3 |
| 3 | 3 | 1 | 1+1*100=101 | 1000 | 1 |
| 4 | 1 | 1 | 101+1*1000=1101 | 10000 | 0 |

Wynik: 1101. Sprawdzenie: 1*8 + 1*4 + 0*2 + 1*1 = 13.

**Weryfikacja konwertuj(100, 8):**

| Krok | n | n%8 | wynik | mnoznik | n (po /8) |
|------|---|-----|-------|---------|-----------|
| 1 | 100 | 4 | 4 | 10 | 12 |
| 2 | 12 | 4 | 44 | 100 | 1 |
| 3 | 1 | 1 | 144 | 1000 | 0 |

Wynik: 144. Sprawdzenie: 1*64 + 4*8 + 4*1 = 100.

**Weryfikacja konwertuj(25, 3):**

| Krok | n | n%3 | wynik | mnoznik | n (po /3) |
|------|---|-----|-------|---------|-----------|
| 1 | 25 | 1 | 1 | 10 | 8 |
| 2 | 8 | 2 | 21 | 100 | 2 |
| 3 | 2 | 2 | 221 | 1000 | 0 |

Wynik: 221. Sprawdzenie: 2*9 + 2*3 + 1*1 = 25.
</details>

---

### Cwiczenie 2.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 1.3 (rozszerzony)

Dana jest funkcja rekurencyjna przetwarzajaca drzewo binarne w porzadku preorder. Kazdy wezel drzewa ma: wartosc calkowita `w`, lewy potomek `L` i prawy potomek `R` (mogace byc puste, oznaczone jako NULL).

```
funkcja suma_preorder(wezel)
    jezeli wezel = NULL:
        zwroc 0
    lewy := suma_preorder(wezel.L)
    prawy := suma_preorder(wezel.R)
    zwroc wezel.w + lewy + prawy
```

**Polecenie**: Napisz iteracyjna wersje funkcji `suma_preorder` z uzyciem stosu. Stos przechowuje wezly do odwiedzenia. Nie uzywaj rekurencji.

Wskazowka: W porzadku preorder najpierw przetwarzamy wezel, potem lewe poddrzewo, potem prawe. Na stos wkladamy potomkow w odwrotnej kolejnosci (najpierw prawy, potem lewy).

<details>
<summary>Odpowiedz</summary>

**Pseudokod:**

```
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

**C++:**

```cpp
struct Wezel {
    int w;
    Wezel* L;
    Wezel* R;
};

int suma_preorder_iter(Wezel* korzen) {
    if (korzen == nullptr) return 0;
    int suma = 0;
    stack<Wezel*> stos;
    stos.push(korzen);
    while (!stos.empty()) {
        Wezel* wezel = stos.top();
        stos.pop();
        suma += wezel->w;
        if (wezel->R != nullptr)
            stos.push(wezel->R);
        if (wezel->L != nullptr)
            stos.push(wezel->L);
    }
    return suma;
}
```

**Wyjasinienie:**

1. Stos symuluje stos wywolan rekurencji. Zamiast wywolywac funkcje rekurencyjnie, wkladamy wezly na stos.

2. Kolejnosc wkladania na stos jest ODWROTNA do porzadku przetwarzania:
   - Chcemy przetworzyc: wezel, lewy, prawy (preorder)
   - Wkladamy na stos: prawy, lewy (bo stos jest LIFO)
   - Sciagamy ze stosu: lewy (przetwarzany jako pierwszy potomek)

3. Dodajemy wartosc wezla natychmiast po sciagnieciu ze stosu (preorder = wezel przed potomkami).

**Weryfikacja na przykladowym drzewie:**

```
       5
      / \
     3   8
    / \   \
   1   4   9
```

Stos (od dolu):
1. push(5). Stos: [5]
2. pop(5), suma=5. push(8), push(3). Stos: [8, 3]
3. pop(3), suma=8. push(4), push(1). Stos: [8, 4, 1]
4. pop(1), suma=9. (brak potomkow). Stos: [8, 4]
5. pop(4), suma=13. (brak potomkow). Stos: [8]
6. pop(8), suma=21. push(9). Stos: [9]
7. pop(9), suma=30. (brak potomkow). Stos: []

Wynik: 30. Sprawdzenie: 5+3+1+4+8+9 = 30.
</details>
