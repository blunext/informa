# 02. Projektowanie algorytmu

Typ zadania: **projektowanie_algorytmu**
Czestotliwosc: 12/12 lat | Laczna punktacja: 43 pkt
Kategoria: TEORIA

## Umiejetnosci cwiczone w tym zestawie

`mod-div` `cyfry` `petla-while` `pseudokod` `C++` `warunek` `rekurencja` `stos` `drzewo` `konwersja-systemow` `palindrom` `NWD` `liczby-pierwsze` `napisy` `tablica` `podciag` `iteracja` `zlozonosc`

---

### Cwiczenie 2.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 2 (cyfry), styl CKE
**Tagi**: `mod-div` `cyfry` `petla-while` `warunek`

Napisz w pseudokodzie lub C++ funkcje `parzyste_cyfry(n)`, ktora dla danej liczby naturalnej n > 0 zwraca liczbe jej cyfr parzystych (0, 2, 4, 6, 8).

**Ograniczenia**: Uzyj tylko zmiennych calkowitych i operatorow arytmetycznych (mod, div, +, porownania). Nie wolno uzywac stringow, tablic ani funkcji wbudowanych.

**Przyklady**:
- `parzyste_cyfry(24681)` = 4 (cyfry parzyste: 2, 4, 6, 8)
- `parzyste_cyfry(13579)` = 0 (brak cyfr parzystych)
- `parzyste_cyfry(1024)` = 3 (cyfry parzyste: 0, 2, 4)

<details>
<summary>Wskazowki</summary>

1. Zastanow sie, jak wyodrebnic ostatnia cyfre liczby -- jaki operator arytmetyczny daje reszte z dzielenia przez 10?
2. Po sprawdzeniu ostatniej cyfry, jak "usunac" ja z liczby? Uzyj dzielenia calkowitego przez 10. Powtarzaj az liczba stanie sie rowna 0.
3. W kazdym kroku petli: `cyfra = n mod 10`, sprawdz `cyfra mod 2 == 0`, jesli tak -- zwieksz licznik. Na koncu `n = n div 10`.
</details>

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

<details>
<summary>Typowe bledy</summary>

1. **Brak obslugi cyfry 0 jako parzystej** -- cyfra 0 jest parzysta (0 mod 2 = 0). Uczniowie czesto pomijaja ja w warunku lub traktuja jako nieparzysty przypadek brzegowy. CKE: -1 pkt za bledny wynik dla danych z cyfra 0.
2. **Uzycie `n > 0` zamiast petli po cyfrach** -- sprawdzanie `n != 0` zadziala identycznie, ale jesli uczeń uzyje warunku `n >= 10` (myslac o "cyfrach wielocyfrowych"), pominie jednocyfrowe liczby. CKE: -1 pkt.
3. **Porownywanie przez wyliczenie `cyfra==0 || cyfra==2 || ...`** -- poprawne, ale CKE preferuje `cyfra mod 2 = 0` jako elegantsze rozwiazanie. Punktacja zwykle pelna, ale ryzyko bledu przy wyliczaniu.
</details>

---

### Cwiczenie 2.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 1.3 (przestaw iteracyjnie)
**Tagi**: `rekurencja` `petla-while` `cyfry` `pseudokod`

Dana jest funkcja rekurencyjna obliczajaca sume cyfr liczby:

```
funkcja sumaCyfr(n)
    jezeli n = 0:
        zwroc 0
    zwroc sumaCyfr(n div 10) + n mod 10
```

**Polecenie**: Napisz iteracyjna wersje tej funkcji. Uzyj petli `dopoki` (while). Nie uzywaj rekurencji.

<details>
<summary>Wskazowki</summary>

1. Przeanalizuj, co robi rekurencja: w kazdym wywolaniu dodaje ostatnia cyfre (`n mod 10`) i przechodzi do liczby bez ostatniej cyfry (`n div 10`). Jak to przelozyc na petle?
2. Potrzebujesz zmiennej akumulujacej sume. W kazdej iteracji petli dodajesz `n mod 10` do sumy i zmniejszasz `n` przez `n div 10`.
3. Warunek konca petli to `n > 0` (odpowiada warunkowi bazowemu rekurencji `n = 0`). Przed petla zainicjuj `suma := 0`.
</details>

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

<details>
<summary>Typowe bledy</summary>

1. **Zapomnienie inicjalizacji akumulatora** -- zmienna `suma` musi byc zainicjowana na 0 przed petla. Brak inicjalizacji w C++ daje niezdefiniowane zachowanie. CKE: -1 pkt za brak inicjalizacji.
2. **Zamiana kolejnosci operacji** -- pisanie `n := n div 10` PRZED `suma := suma + n mod 10` powoduje utrate ostatniej cyfry. CKE: blad logiczny, -1 pkt.
</details>

---

### Cwiczenie 2.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: styl CKE -- "tylko zmienne calkowite"
**Tagi**: `palindrom` `mod-div` `cyfry` `C++`

Napisz pseudokod funkcji `czyPalindrom(n)` sprawdzajacej, czy liczba naturalna n > 0 jest palindromem (czytana od lewej i od prawej daje ta sama liczbe, np. 12321, 4554, 7).

**Ograniczenia**: Nie wolno uzywac stringow, tablic ani funkcji wbudowanych. Dozwolone sa wylacznie zmienne calkowite i operatory arytmetyczne (mod, div, +, *, porownania).

**Przyklady**:
- `czyPalindrom(12321)` = PRAWDA
- `czyPalindrom(12345)` = FALSZ
- `czyPalindrom(4554)` = PRAWDA
- `czyPalindrom(7)` = PRAWDA

<details>
<summary>Wskazowki</summary>

1. Pomysl o strategii: jesli odwrocisz cyfry liczby i otrzymasz te sama wartosc, to liczba jest palindromem. Jak odwrocic liczbe uzywajac mod i div?
2. Buduj odwrocona liczbe od zera: w kazdym kroku `odwrocona = odwrocona * 10 + n mod 10`, potem `n = n div 10`. Powtarzaj az `n = 0`.
3. Pamietaj, zeby PRZED petla zapisac oryginalna wartosc `n` do osobnej zmiennej (`oryginal := n`), bo petla zmodyfikuje `n`. Na koncu porownaj `odwrocona` z `oryginal`.
</details>

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

<details>
<summary>Typowe bledy</summary>

1. **Brak zapisu oryginalnej wartosci** -- porownywanie z `n` po petli (ktore jest juz rowne 0) zamiast z zapisana wartoscia poczatkowa. CKE: -2 pkt, bo wynik zawsze PRAWDA (0 == 0 po petli jest falszywe, ale logika bledu jest powazna).
2. **Bledna kolejnosc w budowaniu odwroconej** -- pisanie `odwrocona = cyfra * 10 + odwrocona` zamiast `odwrocona * 10 + cyfra` daje bledny wynik. CKE: -2 pkt za bledny algorytm.
3. **Pominiecie jednocyfrowych liczb** -- jednocyfrowa liczba jest zawsze palindromem. Algorytm poprawnie to obsluguje (po jednej iteracji odwrocona = cyfra = oryginal), ale uczniowie czasem dodaja niepotrzebny warunek brzegowy, ktory moze wprowadzic blad.
</details>

---

### Cwiczenie 2.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3 (pseudokod), Matura 2024 zad. 6 (systemy)
**Tagi**: `konwersja-systemow` `mod-div` `petla-while` `pseudokod`

Napisz pseudokod funkcji `konwertuj(n, k)`, ktora zwraca zapis liczby naturalnej n > 0 w systemie o podstawie k (2 <= k <= 9) jako liczbe calkowita.

Na przyklad: `konwertuj(13, 2)` zwraca 1101, bo 13 w systemie dwojkowym to 1101.

**Ograniczenia**: Uzyj tylko zmiennych calkowitych i operatorow arytmetycznych. Wynik ma byc liczba calkowita (nie string). Zakladamy, ze wynik miesci sie w zakresie typu calkowitego.

**Przyklady**:
- `konwertuj(13, 2)` = 1101
- `konwertuj(100, 8)` = 144
- `konwertuj(25, 3)` = 221

<details>
<summary>Wskazowki</summary>

1. Reszta z dzielenia `n mod k` daje kolejne cyfry wyniku (od najmniej znaczacej). Jak "dostawic" te cyfre na odpowiednia pozycje w wyniku?
2. Uzyj zmiennej `mnoznik` (startowo 1), ktora mnozy reszte i przesuwa ja na wlasciwa pozycje. Po kazdej iteracji `mnoznik = mnoznik * 10`.
3. W kazdym kroku: `wynik = wynik + (n mod k) * mnoznik`, potem `n = n div k` i `mnoznik = mnoznik * 10`. Kontynuuj az `n = 0`.
</details>

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

<details>
<summary>Typowe bledy</summary>

1. **Inicjalizacja `mnoznik := 0` zamiast `mnoznik := 1`** -- pierwsza cyfra zostaje pomnozona przez 0 i znika. CKE: -2 pkt za blad w algorytmie.
2. **Uzycie `n mod 10` zamiast `n mod k`** -- skopiowanie szablonu z cwiczen na cyfrach dziesietnych bez zmiany podstawy. CKE: -2 pkt.
3. **Proba konwersji dla k >= 10** -- algorytm z wynikiem calkowitym nie dziala dla systemow o podstawie >= 10 (np. szesnastkowy), bo cyfry >= 10 nie maja jednoznacznej reprezentacji w liczbie calkowitej. To nie jest blad, jesli polecenie wymaga k <= 9, ale warto byc swiadomym ograniczenia.
</details>

---

### Cwiczenie 2.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 1.3 (rozszerzony)
**Tagi**: `rekurencja` `stos` `drzewo` `C++`

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
<summary>Wskazowki</summary>

1. Stos sluzy do zapamietywania wezlow, ktore jeszcze musimy odwiedzic. Zacznij od wlozenia korzenia na stos. W kazdej iteracji zdejmij wezel z wierzcholka stosu.
2. Po zdjeciu wezla ze stosu dodaj jego wartosc do sumy, a nastepnie wloz jego potomkow na stos. Pamietaj o kolejnosci: najpierw PRAWY, potem LEWY (bo stos jest LIFO -- ostatni wlozony bedzie pierwszy zdjety).
3. Petla konczy sie, gdy stos jest pusty. Nie wkladaj na stos pustych wezlow (NULL). Caly algorytm: `push(korzen)`, potem `while(!empty): wezel=pop(), suma+=w, push(R), push(L)`.
</details>

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

<details>
<summary>Typowe bledy</summary>

1. **Bledna kolejnosc push na stos** -- wkladanie najpierw lewego, potem prawego powoduje przetwarzanie w kolejnosci wezel-prawy-lewy zamiast preorder (wezel-lewy-prawy). CKE: -1 pkt za bledna kolejnosc przejscia.
2. **Brak sprawdzenia NULL przed push** -- wkladanie pustych wezlow na stos powoduje blad wykonania (dereferencja nullptr). CKE: program nie kompiluje sie lub wywala -- 0 pkt za implementacje.
3. **Uzycie `stos.front()` zamiast `stos.top()`** -- pomylenie stosu z kolejka. W C++ `stack` nie ma metody `front()`. CKE: -1 pkt za blad skladniowy.
</details>

---

### Cwiczenie 2.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 2 (NWD), styl CKE
**Tagi**: `NWD` `petla-while` `mod-div` `pseudokod`

Napisz w pseudokodzie lub C++ funkcje `NWD(a, b)`, ktora oblicza najwiekszy wspolny dzielnik dwoch liczb naturalnych a > 0 i b > 0, korzystajac z algorytmu Euklidesa z reszta z dzielenia (wersja iteracyjna).

**Ograniczenia**: Uzyj tylko zmiennych calkowitych i operatora mod. Nie wolno uzywac rekurencji ani funkcji wbudowanych.

**Przyklady**:
- `NWD(48, 18)` = 6
- `NWD(100, 25)` = 25
- `NWD(17, 13)` = 1 (liczby wzglednie pierwsze)
- `NWD(12, 12)` = 12

<details>
<summary>Wskazowki</summary>

1. Algorytm Euklidesa opiera sie na fakcie, ze NWD(a, b) = NWD(b, a mod b). Jak to przelozyc na petle?
2. Petla dziala dopoki `b > 0`. W kazdym kroku oblicz `reszta = a mod b`, potem zamien: `a = b`, `b = reszta`. Gdy `b` osiagnie 0, wynikiem jest `a`.
3. Przyklad krok po kroku: NWD(48, 18): 48 mod 18 = 12, potem NWD(18, 12): 18 mod 12 = 6, potem NWD(12, 6): 12 mod 6 = 0, wynik = 6.
</details>

<details>
<summary>Odpowiedz</summary>

**Pseudokod:**

```
funkcja NWD(a, b)
    dopoki b > 0:
        reszta := a mod b
        a := b
        b := reszta
    zwroc a
```

**C++:**

```cpp
int NWD(int a, int b) {
    while (b > 0) {
        int reszta = a % b;
        a = b;
        b = reszta;
    }
    return a;
}
```

**Weryfikacja NWD(48, 18):**

| Krok | a | b | a mod b | a (nowe) | b (nowe) |
|------|---|---|---------|----------|----------|
| 1 | 48 | 18 | 12 | 18 | 12 |
| 2 | 18 | 12 | 6 | 12 | 6 |
| 3 | 12 | 6 | 0 | 6 | 0 |

Wynik: 6. Sprawdzenie: 48 = 6*8, 18 = 6*3.

**Weryfikacja NWD(100, 25):**

| Krok | a | b | a mod b | a (nowe) | b (nowe) |
|------|---|---|---------|----------|----------|
| 1 | 100 | 25 | 0 | 25 | 0 |

Wynik: 25. Sprawdzenie: 100 = 25*4, 25 = 25*1.

**Weryfikacja NWD(17, 13):**

| Krok | a | b | a mod b | a (nowe) | b (nowe) |
|------|---|---|---------|----------|----------|
| 1 | 17 | 13 | 4 | 13 | 4 |
| 2 | 13 | 4 | 1 | 4 | 1 |
| 3 | 4 | 1 | 0 | 1 | 0 |

Wynik: 1. Liczby 17 i 13 sa wzglednie pierwsze.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Bledna kolejnosc przypiasan** -- pisanie `a := reszta; b := a` zamiast uzycia zmiennej tymczasowej `reszta`. Bez zmiennej posredniej wartosc `a` jest nadpisana zanim zostanie przypisana do `b`. CKE: -2 pkt za bledny algorytm.
2. **Warunek petli `b >= 0` zamiast `b > 0`** -- powoduje dzielenie przez zero w nastepnej iteracji, gdy `b` osiagnie 0. CKE: -1 pkt za blad wykonania.
3. **Zwracanie `b` zamiast `a`** -- po zakonczeniu petli `b = 0`, wiec wynik zawsze bylby 0. CKE: -2 pkt.
</details>

---

### Cwiczenie 2.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 2 (liczby pierwsze), styl CKE
**Tagi**: `liczby-pierwsze` `petla-while` `warunek` `C++`

Napisz w pseudokodzie lub C++ funkcje `czyPierwsza(n)`, ktora dla danej liczby naturalnej n >= 2 sprawdza, czy n jest liczba pierwsza. Funkcja zwraca PRAWDA, jesli n jest pierwsza, FALSZ w przeciwnym przypadku.

**Ograniczenia**: Uzyj metody sprawdzania podzielnosci przez kolejne liczby od 2. Zoptymalizuj algorytm tak, aby sprawdzac dzielniki tylko do pierwiastka z n (tj. dopoki `dzielnik * dzielnik <= n`). Nie wolno uzywac funkcji `sqrt` ani tablic.

**Przyklady**:
- `czyPierwsza(2)` = PRAWDA
- `czyPierwsza(17)` = PRAWDA
- `czyPierwsza(15)` = FALSZ (15 = 3 * 5)
- `czyPierwsza(1)` = FALSZ (1 nie jest liczba pierwsza)

<details>
<summary>Wskazowki</summary>

1. Liczba jest pierwsza, jesli nie ma dzielnikow wiekszych od 1 i mniejszych od niej samej. Zacznij sprawdzanie od dzielnika 2 i zwieksaj go o 1.
2. Nie musisz sprawdzac dzielnikow wiekszych niz pierwiastek z n. Zamiast obliczac pierwiastek, uzyj warunku `dzielnik * dzielnik <= n` -- to unika potrzeby uzycia `sqrt`.
3. Jesli znajdziesz jakikolwiek dzielnik (tj. `n mod dzielnik = 0`), zwroc od razu FALSZ. Jesli petla sie skonczy bez znalezienia dzielnika, zwroc PRAWDA. Pamietaj o przypadku brzegowym: n < 2 nie jest pierwsza.
</details>

<details>
<summary>Odpowiedz</summary>

**Pseudokod:**

```
funkcja czyPierwsza(n)
    jezeli n < 2:
        zwroc FALSZ
    dzielnik := 2
    dopoki dzielnik * dzielnik <= n:
        jezeli n mod dzielnik = 0:
            zwroc FALSZ
        dzielnik := dzielnik + 1
    zwroc PRAWDA
```

**C++:**

```cpp
bool czyPierwsza(int n) {
    if (n < 2) return false;
    int dzielnik = 2;
    while (dzielnik * dzielnik <= n) {
        if (n % dzielnik == 0)
            return false;
        dzielnik++;
    }
    return true;
}
```

**Weryfikacja czyPierwsza(17):**

| Krok | dzielnik | dzielnik*dzielnik | <= 17? | 17 % dzielnik |
|------|----------|-------------------|--------|---------------|
| 1 | 2 | 4 | tak | 1 (nie dzieli) |
| 2 | 3 | 9 | tak | 2 (nie dzieli) |
| 3 | 4 | 16 | tak | 1 (nie dzieli) |
| 4 | 5 | 25 | nie | - |

Petla konczy sie, zwroc PRAWDA. 17 jest liczba pierwsza.

**Weryfikacja czyPierwsza(15):**

| Krok | dzielnik | dzielnik*dzielnik | <= 15? | 15 % dzielnik |
|------|----------|-------------------|--------|---------------|
| 1 | 2 | 4 | tak | 1 (nie dzieli) |
| 2 | 3 | 9 | tak | 0 (dzieli!) |

Znaleziono dzielnik 3, zwroc FALSZ. 15 = 3 * 5.

**Weryfikacja czyPierwsza(2):**

dzielnik = 2, dzielnik*dzielnik = 4 > 2, petla sie nie wykonuje, zwroc PRAWDA.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Sprawdzanie `dzielnik <= n` zamiast `dzielnik * dzielnik <= n`** -- algorytm poprawny, ale zbyt wolny (zlozonosc O(n) zamiast O(sqrt(n))). CKE moze odebrac -1 pkt za brak optymalizacji, jesli polecenie wyraznie wymaga efektywnosci.
2. **Brak obslugi n < 2** -- dla n = 1 algorytm bez warunku brzegowego zwroci PRAWDA (petla sie nie wykona), co jest bledne. CKE: -1 pkt.
3. **Uzycie `dzielnik * dzielnik < n` (ostry) zamiast `<=`** -- pomija przypadek, gdy n jest kwadratem liczby pierwszej (np. 25 = 5*5). Dla n = 25 i dzielnik = 5: 5*5 = 25, `25 < 25` jest falsz, wiec petla sie konczy i zwraca PRAWDA (blednie). CKE: -1 pkt za bledny wynik.
</details>

---

### Cwiczenie 2.8 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 5 (napisy), styl CKE
**Tagi**: `napisy` `petla-while` `warunek` `C++`

Napisz w pseudokodzie lub C++ funkcje `najdluzsza_seria(s)`, ktora dla danego napisu `s` (ciagu znakow o dlugosci >= 1) zwraca dlugosc najdluzszej serii identycznych kolejnych znakow.

Na przyklad, w napisie "aaabbcdddd" najdluzsza seria to "dddd" o dlugosci 4.

**Ograniczenia**: Uzyj petli i porownywania znakow. Nie wolno uzywac funkcji wbudowanych do wyszukiwania wzorcow. Zakladamy indeksowanie od 0.

**Przyklady**:
- `najdluzsza_seria("aaabbcdddd")` = 4 (seria 'd')
- `najdluzsza_seria("abcdef")` = 1 (brak powtorzen)
- `najdluzsza_seria("aaa")` = 3
- `najdluzsza_seria("x")` = 1

<details>
<summary>Wskazowki</summary>

1. Przejdz po napisie od drugiego znaku (indeks 1). Porownuj kazdy znak z poprzednim. Jesli sa takie same, zwieksz biezacy licznik serii; jesli rozne, zacznij nowa serie.
2. Potrzebujesz dwoch zmiennych: `biezaca` (dlugosc aktualnej serii) i `najdluzsza` (dotychczasowe maksimum). Gdy `biezaca` staje sie wieksza niz `najdluzsza`, zaktualizuj `najdluzsza`.
3. Na poczatku ustaw `biezaca = 1` i `najdluzsza = 1` (pierwszy znak tworzy serie dlugosci 1). Iteruj od indeksu 1 do konca napisu.
</details>

<details>
<summary>Odpowiedz</summary>

**Pseudokod:**

```
funkcja najdluzsza_seria(s)
    n := dlugosc(s)
    najdluzsza := 1
    biezaca := 1
    dla i := 1, 2, ..., n-1:
        jezeli s[i] = s[i-1]:
            biezaca := biezaca + 1
        w przeciwnym razie:
            biezaca := 1
        jezeli biezaca > najdluzsza:
            najdluzsza := biezaca
    zwroc najdluzsza
```

**C++:**

```cpp
int najdluzsza_seria(const string& s) {
    int n = s.size();
    int najdluzsza = 1;
    int biezaca = 1;
    for (int i = 1; i < n; i++) {
        if (s[i] == s[i - 1])
            biezaca++;
        else
            biezaca = 1;
        if (biezaca > najdluzsza)
            najdluzsza = biezaca;
    }
    return najdluzsza;
}
```

**Weryfikacja najdluzsza_seria("aaabbcdddd"):**

| i | s[i] | s[i-1] | rowne? | biezaca | najdluzsza |
|---|------|--------|--------|---------|------------|
| 1 | a | a | tak | 2 | 2 |
| 2 | a | a | tak | 3 | 3 |
| 3 | b | a | nie | 1 | 3 |
| 4 | b | b | tak | 2 | 3 |
| 5 | c | b | nie | 1 | 3 |
| 6 | d | c | nie | 1 | 3 |
| 7 | d | d | tak | 2 | 3 |
| 8 | d | d | tak | 3 | 3 |
| 9 | d | d | tak | 4 | 4 |

Wynik: 4 (seria 'd').

**Weryfikacja najdluzsza_seria("abcdef"):**

Kazdy znak rozny od poprzedniego, `biezaca` resetuje sie do 1 w kazdym kroku. Wynik: 1.

**Weryfikacja najdluzsza_seria("x"):**

Petla sie nie wykonuje (n=1, i=1 < 1 falsz). Wynik: 1 (wartosc poczatkowa).
</details>

<details>
<summary>Typowe bledy</summary>

1. **Aktualizacja `najdluzsza` tylko w galezi `else`** -- jesli najdluzsza seria konczy sie na koncu napisu (np. "abcdddd"), wartosc `biezaca` nigdy nie jest porownana z `najdluzsza` po ostatnim znaku. CKE: -1 pkt za bledny wynik.
2. **Inicjalizacja `biezaca = 0` i `najdluzsza = 0`** -- dla jednoelementowego napisu petla sie nie wykona i wynik bedzie 0 zamiast 1. CKE: -1 pkt.
3. **Iteracja od i = 0 z porownaniem `s[i] == s[i-1]`** -- dostep do `s[-1]` powoduje odczyt poza zakresem tablicy. CKE: blad wykonania, -2 pkt.
</details>

---

### Cwiczenie 2.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 4 (podciagi), styl CKE
**Tagi**: `tablica` `podciag` `iteracja` `C++`

Dana jest tablica `T` zawierajaca `n` liczb calkowitych (n >= 1). Napisz w pseudokodzie lub C++ funkcje `max_suma_podciagu(T, n)`, ktora znajduje maksymalna sume spojnego podciagu tablicy (tj. ciagu kolejnych elementow). Jesli wszystkie elementy sa ujemne, wynikiem jest najwiekszy (najmniej ujemny) element.

**Ograniczenia**: Algorytm powinien dzialac w czasie O(n), tj. przejsc przez tablice dokladnie raz. Nie wolno uzywac zagniezdzonej petli. (Wskazowka: algorytm Kadane'a.)

**Przyklady**:
- `max_suma_podciagu([2, -1, 3, -2, 5], 5)` = 7 (podciag [2, -1, 3, -2, 5])
- `max_suma_podciagu([-3, -2, -1, -4], 4)` = -1 (element -1)
- `max_suma_podciagu([1, 2, 3], 3)` = 6 (cala tablica)
- `max_suma_podciagu([5, -9, 6, -2, 3], 5)` = 7 (podciag [6, -2, 3])

<details>
<summary>Wskazowki</summary>

1. Pomysl o tym tak: idac po tablicy od lewej, w kazdym momencie zastanow sie, czy "oplacalniej" jest kontynuowac biezacy podciag (dodajac aktualny element) czy zaczac nowy podciag od aktualnego elementu.
2. Utrzymuj dwie zmienne: `biezaca_suma` (najlepsza suma konczaca sie na biezacym elemencie) i `max_suma` (globalne maximum dotychczas). W kazdym kroku: `biezaca_suma = max(T[i], biezaca_suma + T[i])`.
3. Jesli `biezaca_suma + T[i] < T[i]`, to znaczy ze `biezaca_suma < 0` i oplacalniej jest zaczac od nowa. Zaktualizuj `max_suma = max(max_suma, biezaca_suma)`. Zainicjuj obie zmienne wartoscia `T[0]`.
</details>

<details>
<summary>Odpowiedz</summary>

**Pseudokod (algorytm Kadane'a):**

```
funkcja max_suma_podciagu(T, n)
    biezaca_suma := T[0]
    max_suma := T[0]
    dla i := 1, 2, ..., n-1:
        jezeli biezaca_suma + T[i] > T[i]:
            biezaca_suma := biezaca_suma + T[i]
        w przeciwnym razie:
            biezaca_suma := T[i]
        jezeli biezaca_suma > max_suma:
            max_suma := biezaca_suma
    zwroc max_suma
```

**C++:**

```cpp
int max_suma_podciagu(int T[], int n) {
    int biezaca_suma = T[0];
    int max_suma = T[0];
    for (int i = 1; i < n; i++) {
        if (biezaca_suma + T[i] > T[i])
            biezaca_suma = biezaca_suma + T[i];
        else
            biezaca_suma = T[i];
        if (biezaca_suma > max_suma)
            max_suma = biezaca_suma;
    }
    return max_suma;
}
```

**Weryfikacja max_suma_podciagu([2, -1, 3, -2, 5], 5):**

| i | T[i] | biezaca_suma + T[i] | T[i] | biezaca_suma | max_suma |
|---|------|---------------------|------|--------------|----------|
| 0 | 2 | - | - | 2 | 2 |
| 1 | -1 | 2+(-1)=1 | -1 | 1 (kontynuuj) | 2 |
| 2 | 3 | 1+3=4 | 3 | 4 (kontynuuj) | 4 |
| 3 | -2 | 4+(-2)=2 | -2 | 2 (kontynuuj) | 4 |
| 4 | 5 | 2+5=7 | 5 | 7 (kontynuuj) | 7 |

Wynik: 7. Podciag: [2, -1, 3, -2, 5] (cala tablica).

**Weryfikacja max_suma_podciagu([-3, -2, -1, -4], 4):**

| i | T[i] | biezaca_suma + T[i] | T[i] | biezaca_suma | max_suma |
|---|------|---------------------|------|--------------|----------|
| 0 | -3 | - | - | -3 | -3 |
| 1 | -2 | -3+(-2)=-5 | -2 | -2 (nowy) | -2 |
| 2 | -1 | -2+(-1)=-3 | -1 | -1 (nowy) | -1 |
| 3 | -4 | -1+(-4)=-5 | -4 | -4 (nowy) | -1 |

Wynik: -1 (element -1, bo wszystkie sa ujemne).

**Weryfikacja max_suma_podciagu([5, -9, 6, -2, 3], 5):**

| i | T[i] | biezaca_suma + T[i] | T[i] | biezaca_suma | max_suma |
|---|------|---------------------|------|--------------|----------|
| 0 | 5 | - | - | 5 | 5 |
| 1 | -9 | 5+(-9)=-4 | -9 | -4 (kontynuuj) | 5 |
| 2 | 6 | -4+6=2 | 6 | 6 (nowy) | 6 |
| 3 | -2 | 6+(-2)=4 | -2 | 4 (kontynuuj) | 6 |
| 4 | 3 | 4+3=7 | 3 | 7 (kontynuuj) | 7 |

Wynik: 7. Podciag: [6, -2, 3].
</details>

<details>
<summary>Typowe bledy</summary>

1. **Inicjalizacja `max_suma = 0`** -- jesli wszystkie elementy sa ujemne, wynik blednie wyniesie 0 (co nie odpowiada zadnemu podciagowi). Prawidlowa inicjalizacja to `max_suma = T[0]`. CKE: -1 pkt za bledny wynik dla danych ujemnych.
2. **Uzycie zagniezdzonej petli (brute force O(n^2))** -- sprawdzanie kazdego mozliwego podciagu jest poprawne, ale nie spelnia wymagania zlozonosci O(n). CKE: -1 do -2 pkt za brak optymalizacji, jesli polecenie tego wymaga.
3. **Zapomnienie o aktualizacji `max_suma` wewnatrz petli** -- obliczanie `biezaca_suma` bez porownywania z `max_suma` w kazdym kroku moze powodowac pominiecie optymalnego podciagu, ktory nie konczy sie na ostatnim elemencie. CKE: -1 pkt.
</details>

---

### Cwiczenie 2.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3 (rekurencja i iteracja), styl CKE
**Tagi**: `rekurencja` `iteracja` `cyfry` `zlozonosc` `C++`

Dany jest nastepujacy algorytm rekurencyjny obliczajacy "cyfrowy korzen" liczby (tj. powtarzane sumowanie cyfr az do uzyskania jednocyfrowego wyniku):

```
funkcja korzen_cyfrowy(n)
    jezeli n < 10:
        zwroc n
    zwroc korzen_cyfrowy(sumaCyfr(n))
```

gdzie `sumaCyfr(n)` oblicza sume cyfr liczby n (patrz cwiczenie 2.2).

**Polecenie**: Napisz w pelni iteracyjna wersje funkcji `korzen_cyfrowy(n)` -- bez uzywania rekurencji i bez oddzielnej funkcji `sumaCyfr`. Uzyj zagniezdzonej petli: zewnetrzna petla powtarza sumowanie, wewnetrzna oblicza sume cyfr.

**Ograniczenia**: Nie wolno uzywac rekurencji, stringow ani tablic. Tylko zmienne calkowite i operatory arytmetyczne.

**Przyklady**:
- `korzen_cyfrowy(9875)` = 2 (9+8+7+5=29, 2+9=11, 1+1=2)
- `korzen_cyfrowy(493)` = 7 (4+9+3=16, 1+6=7)
- `korzen_cyfrowy(7)` = 7 (juz jednocyfrowa)
- `korzen_cyfrowy(100)` = 1 (1+0+0=1)

<details>
<summary>Wskazowki</summary>

1. Potrzebujesz dwoch petli. Zewnetrzna petla dziala dopoki `n >= 10` (tzn. n ma wiecej niz jedna cyfre). Wewnetrzna petla oblicza sume cyfr biezacego `n`.
2. W wewnetrznej petli oblicz sume cyfr: `suma = 0`, potem `while (n > 0): suma += n mod 10; n = n div 10`. Po wewnetrznej petli przypisz `n = suma`.
3. Zewnetrzna petla sprawdza warunek `n >= 10` PRZED wewnetrzna. Jesli `n < 10`, petla sie konczy i zwracamy `n`. Calosc to `while (n >= 10) { suma=0; while (n>0) { suma+=n%10; n/=10; } n=suma; }`.
</details>

<details>
<summary>Odpowiedz</summary>

**Pseudokod:**

```
funkcja korzen_cyfrowy(n)
    dopoki n >= 10:
        suma := 0
        dopoki n > 0:
            suma := suma + n mod 10
            n := n div 10
        n := suma
    zwroc n
```

**C++:**

```cpp
int korzen_cyfrowy(int n) {
    while (n >= 10) {
        int suma = 0;
        while (n > 0) {
            suma += n % 10;
            n /= 10;
        }
        n = suma;
    }
    return n;
}
```

**Weryfikacja korzen_cyfrowy(9875):**

Iteracja zewnetrzna 1 (n = 9875 >= 10):
| Krok wewn. | n | n%10 | suma | n (po /10) |
|------------|---|------|------|------------|
| 1 | 9875 | 5 | 5 | 987 |
| 2 | 987 | 7 | 12 | 98 |
| 3 | 98 | 8 | 20 | 9 |
| 4 | 9 | 9 | 29 | 0 |
n := 29

Iteracja zewnetrzna 2 (n = 29 >= 10):
| Krok wewn. | n | n%10 | suma | n (po /10) |
|------------|---|------|------|------------|
| 1 | 29 | 9 | 9 | 2 |
| 2 | 2 | 2 | 11 | 0 |
n := 11

Iteracja zewnetrzna 3 (n = 11 >= 10):
| Krok wewn. | n | n%10 | suma | n (po /10) |
|------------|---|------|------|------------|
| 1 | 11 | 1 | 1 | 1 |
| 2 | 1 | 1 | 2 | 0 |
n := 2

n = 2 < 10, petla zewnetrzna konczy sie. Wynik: 2.

**Weryfikacja korzen_cyfrowy(493):**

Iteracja 1: 4+9+3 = 16, n := 16
Iteracja 2: 1+6 = 7, n := 7
n = 7 < 10. Wynik: 7.

**Weryfikacja korzen_cyfrowy(7):**

n = 7 < 10, petla zewnetrzna sie nie wykonuje. Wynik: 7.

**Uwaga o zlozonosci**: Mozna udowodnic, ze korzen cyfrowy mozna obliczyc wzorem `1 + (n - 1) % 9` dla n > 0 (zwiazek z arytmetyka modularną). Jednak na maturze CKE oczekuje algorytmu iteracyjnego z petlami, nie wzoru matematycznego.
</details>

<details>
<summary>Typowe bledy</summary>

1. **Brak reinicjalizacji `suma = 0` na poczatku kazdej iteracji zewnetrznej** -- `suma` z poprzedniej iteracji dodaje sie do nowej, dajac bledne wyniki. CKE: -2 pkt za blad logiczny.
2. **Warunek zewnetrznej petli `n > 0` zamiast `n >= 10`** -- petla nigdy sie nie konczy dla n > 0 lub konczy z blednym wynikiem. Dla n = 7: wewnetrzna petla da suma = 7, n := 7, i petla sie powtarza w nieskonczonosc. CKE: -2 pkt.
3. **Pomylenie kolejnosci petli** -- uzycie jednej petli zamiast zagniezdzonej powoduje, ze algorytm oblicza sume cyfr tylko raz (nie powtarza procesu az do jednej cyfry). CKE: -2 pkt za niekompletny algorytm.
</details>

---

## Samoocena

| Poziom | Opis | Wymaganie |
|--------|------|-----------|
| Podstawowy | Rozumiesz petle while i operatory mod/div, potrafisz napisac proste funkcje na cyfrach | 1-3 cwiczen bez pomocy |
| Dobry | Potrafisz projektowac algorytmy z warunkami, konwersja systemow, obsluga przypadkow brzegowych | 4-6 cwiczen bez pomocy |
| Bardzo dobry | Potrafisz zamienic rekurencje na iteracje, pracujesz z napisami i tablicami, znasz algorytmy klasyczne (NWD, Kadane) | 7-8 cwiczen bez pomocy |
| Doskonaly | Swobodnie projektujesz algorytmy z zagniezdzonymi petlami, stosami i strukturami danych, potrafisz uzasadnic zlozonosc | 9-10 cwiczen bez pomocy |

### Co dalej?
- Jesli poziom **Podstawowy**: Wroc do cwiczen 01 (sledzenie algorytmu) i przesledzreczne kilkanascie algorytmow krok po kroku. Nastepnie powtorz cwiczenia 2.1-2.3 az do bieglego pisania petli while z mod/div.
- Jesli poziom **Dobry**: Przejdz do cwiczen z implementacji (typ 04-06), aby procwiczyz pisanie pelnych programow w C++. Szczegolna uwage zwroc na obsluge plikow i formatowanie wyjscia.
- Jesli poziom **Bardzo dobry**: Skup sie na cwiczeniach trudnych (2.5, 2.9, 2.10) i szukaj alternatywnych rozwiazan. Przejdz do cwiczen z analizy algorytmow (typ 03), aby poglebic rozumienie zlozonosci.
- Jesli poziom **Doskonaly**: Rozwiazuj zadania maturalne z lat 2023-2025 (nowa formula) na czas. Staraj sie zmiesic w 10 minutach na zadanie teoretyczne. Pracuj nad czystoscia pseudokodu i precyzja zapisu.
