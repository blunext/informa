# 03. Analiza algorytmu

Typ zadania: **analiza_algorytmu**
Czestotliwosc: 10/11 lat | Laczna punktacja: 37 pkt
Kategoria: TEORIA

---

### Cwiczenie 3.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 3.1 (zlozonosc), styl P/F z matur

Podane sa cztery fragmenty kodu operujace na tablicy T o n elementach. Przypisz kazdemu fragmentowi zlozonosc czasowa ze zbioru: {O(log n), O(n), O(n^2), O(n^3)}.

**Fragment A:**
```
s := 0
dla i := 1, 2, ..., n:
    s := s + T[i]
```

**Fragment B:**
```
s := 0
dla i := 1, 2, ..., n:
    dla j := 1, 2, ..., n:
        s := s + T[i] * T[j]
```

**Fragment C:**
```
p := 1
q := n
dopoki p < q:
    s := (p + q) div 2
    jezeli T[s] < x:
        p := s + 1
    w przeciwnym razie:
        q := s
```

**Fragment D:**
```
s := 0
dla i := 1, 2, ..., n:
    dla j := 1, 2, ..., n:
        dla k := 1, 2, ..., n:
            jezeli T[i] + T[j] = T[k]:
                s := s + 1
```

<details>
<summary>Odpowiedz</summary>

- **Fragment A**: **O(n)** -- jedna petla przechodzaca n elementow
- **Fragment B**: **O(n^2)** -- dwie zagniezdzne petle, kazda po n iteracji, razem n*n operacji
- **Fragment C**: **O(log n)** -- przeszukiwanie binarne, w kazdym kroku zakres dzielony na pol
- **Fragment D**: **O(n^3)** -- trzy zagniezdzne petle, kazda po n iteracji, razem n*n*n operacji

**Uzasadnienie dla fragmentu C:**
Poczatkowo zakres wynosi n. Po kazdej iteracji zmniejsza sie o polowe:
n -> n/2 -> n/4 -> ... -> 1.
Liczba krokow: log2(n). Stad zlozonosc O(log n).
</details>

---

### Cwiczenie 3.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 zad. 1.2 (kontrprzyklady), Matura 2022 zad. 2

Dany jest algorytm sortowania babelkowego z optymalizacja (wczesniejsze zakonczenie):

```
funkcja sortuj(T, n)
    dla i := 1, 2, ..., n-1:
        zamiana := FALSZ
        dla j := 1, 2, ..., n-i:
            jezeli T[j] > T[j+1]:
                zamien(T[j], T[j+1])
                zamiana := PRAWDA
        jezeli zamiana = FALSZ:
            zakoncz petle zewnetrzna
```

Rozpatrz nastepujace twierdzenie:

*"Algorytm wykonuje dokladnie n*(n-1)/2 porownan niezaleznie od danych wejsciowych."*

**Polecenie**:
- a) Czy twierdzenie jest prawdziwe? Jezeli nie, podaj kontrprzyklad (tablice, dla ktorej liczba porownan jest inna niz n*(n-1)/2).
- b) Podaj minimalna liczbe porownan dla tablicy n-elementowej w tym algorytmie i dane wejsciowe, dla ktorych to minimum jest osiagane.

<details>
<summary>Odpowiedz</summary>

**a) Twierdzenie jest FALSZWE.**

Kontrprzyklad: T = [1, 2, 3, 4, 5] (n = 5)

Przebieg algorytmu:
- i=1: porownujemy pary (1,2), (2,3), (3,4), (4,5) -> 4 porownania, 0 zamian
- zamiana = FALSZ -> algorytm konczy sie!

Liczba porownan: 4 = n-1
Oczekiwana wg twierdzenia: 5*4/2 = 10

4 != 10, wiec twierdzenie jest falszywe.

**b) Minimalna liczba porownan: n-1**

Osiagana dla tablicy juz posortowanej rosnaco, np. T = [1, 2, 3, ..., n].

W pierwszym przebiegu (i=1) algorytm wykonuje n-1 porownan i nie dokonuje zadnej zamiany. Flaga `zamiana` pozostaje FALSZ, wiec petla zewnetrzna konczy sie po jednym przebiegu.

Maksymalna liczba porownan (n*(n-1)/2) wystepuje np. dla tablicy posortowanej malejaco [n, n-1, ..., 2, 1], gdy w kazdym przebiegu nastepuje co najmniej jedna zamiana.
</details>

---

### Cwiczenie 3.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 1b (Korale)

Dany jest algorytm przeszukujacy tablice T o n elementach calkowitych w poszukiwaniu wartosci x:

```
funkcja szukaj(T, n, x)
    dla i := 1, 2, ..., n:
        jezeli T[i] = x:
            zwroc i
    zwroc -1
```

**Polecenie**:
- a) Jaka jest minimalna liczba porownan `T[i] = x` wykonywanych przez algorytm? Podaj dane wejsciowe, dla ktorych to minimum jest osiagane.
- b) Jaka jest maksymalna liczba porownan? Podaj dane wejsciowe, dla ktorych to maksimum jest osiagane.
- c) Jaka jest srednia liczba porownan przy zalozeniu, ze x na pewno znajduje sie w tablicy i z rownym prawdopodobienstwem moze byc na kazdej pozycji?

<details>
<summary>Odpowiedz</summary>

**a) Minimalna liczba porownan: 1**

Osiagana, gdy szukany element x znajduje sie na pierwszej pozycji tablicy.

Przyklad: T = [5, 3, 7, 1, 9], x = 5
Algorytm: i=1, T[1]=5=x -> zwraca 1. Jedno porownanie.

**b) Maksymalna liczba porownan: n**

Osiagana w dwoch przypadkach:
1. Element x jest na ostatniej pozycji: T = [3, 7, 1, 9, 5], x = 5
   Algorytm sprawdza wszystkie n pozycji, znajduje na ostatniej.
2. Element x nie wystepuje w tablicy: T = [3, 7, 1, 9, 2], x = 5
   Algorytm sprawdza wszystkie n pozycji, zwraca -1.

**c) Srednia liczba porownan (x na pewno w tablicy):**

Jezeli x jest na pozycji k (z rownym prawdopodobienstwem 1/n dla k = 1, 2, ..., n), to liczba porownan wynosi k.

Srednia = (1 + 2 + 3 + ... + n) / n = n(n+1)/2 / n = **(n+1)/2**

Przyklad: Dla n = 100 srednia liczba porownan wynosi 50.5.
</details>

---

### Cwiczenie 3.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 1 (liczby skojarzone)

Dany jest algorytm szukajacy liczby dzielnikow liczby n:

```
funkcja liczbaDzielnikow(n)
    licznik := 0
    dla i := 1, 2, ..., n:
        jezeli n mod i = 0:
            licznik := licznik + 1
    zwroc licznik
```

**Polecenie**:
- a) Czy algorytm jest poprawny (daje prawidlowy wynik dla kazdego n >= 1)?
- b) Napisz zoptymalizowana wersje algorytmu o zlozonosci O(sqrt(n)). Wyjasnij, na czym polega optymalizacja.
- c) Ile operacji `n mod i` zaoszczedzisz dla n = 10000, porownujac oryginalna wersje z zoptymalizowana?

<details>
<summary>Odpowiedz</summary>

**a) Tak, algorytm jest poprawny.**

Sprawdza kazda liczbe od 1 do n jako potencjalny dzielnik. Jezeli n mod i = 0, to i jest dzielnikiem n. Wszystkie dzielniki zostana znalezione, poniewaz kazdy dzielnik d liczby n spelnia 1 <= d <= n.

**b) Zoptymalizowana wersja O(sqrt(n)):**

```
funkcja liczbaDzielnikowOpt(n)
    licznik := 0
    i := 1
    dopoki i * i <= n:
        jezeli n mod i = 0:
            licznik := licznik + 1
            jezeli i <> n / i:
                licznik := licznik + 1
        i := i + 1
    zwroc licznik
```

**Optymalizacja polega na:**

Jezeli i jest dzielnikiem n (n mod i = 0), to n/i tez jest dzielnikiem n.
Dzielniki wystepuja w parach (i, n/i), gdzie i <= sqrt(n) i n/i >= sqrt(n).
Wystarczy wiec sprawdzic kandydatow tylko do sqrt(n) i liczyc oba dzielniki z pary.

Wyjatkowy przypadek: jezeli i = n/i (czyli i = sqrt(n)), liczymy dzielnik tylko raz.

**c) Oszczednosc dla n = 10000:**

- Oryginalna wersja: 10000 operacji mod (i od 1 do 10000)
- Zoptymalizowana wersja: sqrt(10000) = 100 operacji mod (i od 1 do 100)
- Oszczednosc: 10000 - 100 = **9900 operacji**

Stosunek: 100x szybciej!

**Weryfikacja dla n = 12:**
- Oryginalna: sprawdza 1..12, dzielniki: 1,2,3,4,6,12 -> 6
- Zoptymalizowana: sqrt(12) ~ 3.46, sprawdza i=1,2,3
  - i=1: 12%1=0, para (1,12), licznik=2
  - i=2: 12%2=0, para (2,6), licznik=4
  - i=3: 12%3=0, para (3,4), licznik=6
  - i=4: 4*4=16>12, koniec
  - Wynik: 6. Poprawnie!
</details>

---

### Cwiczenie 3.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 2 (ab-slowo)

Problem: Dany zbiór n liczb calkowitych. Zlicz ile jest par (i, j), gdzie i < j, takich ze |T[i] - T[j]| <= k (roznica bezwzgledna co najwyzej k).

**Algorytm A** (brute force):

```
funkcja zliczA(T, n, k)
    wynik := 0
    dla i := 1, 2, ..., n-1:
        dla j := i+1, i+2, ..., n:
            jezeli |T[i] - T[j]| <= k:
                wynik := wynik + 1
    zwroc wynik
```

**Algorytm B** (z sortowaniem):

```
funkcja zliczB(T, n, k)
    posortuj(T)   // sortowanie rosnace, koszt O(n log n)
    wynik := 0
    dla i := 1, 2, ..., n:
        // szukaj binarnie ostatniego j takiego ze T[j] <= T[i] + k
        j := szukajBinarnie(T, i+1, n, T[i] + k)
        wynik := wynik + (j - i)
    zwroc wynik
```

W algorytmie B, `szukajBinarnie(T, lo, hi, x)` zwraca indeks ostatniego elementu <= x w zakresie [lo, hi] (lub lo-1 jezeli brak). Koszt jednego wyszukiwania: O(log n).

**Polecenie**:
- a) Podaj zlozonosc czasowa algorytmu A i algorytmu B.
- b) Dla T = [1, 5, 3, 8, 2] i k = 2, oblicz wynik obu algorytmow (powinny byc identyczne).
- c) Oblicz dokladna liczbe operacji porownania dla obu algorytmow przy n = 1000. Przyjmij, ze szukanie binarne wykonuje dokladnie ceil(log2(n)) porownan, a sortowanie n*ceil(log2(n)) porownan.

<details>
<summary>Odpowiedz</summary>

**a) Zlozonosc:**

- Algorytm A: **O(n^2)** -- dwie zagniezdzne petle, kazda do n
- Algorytm B: **O(n log n)** -- sortowanie O(n log n) + n wyszukiwan binarnych O(n log n) = O(n log n)

**b) Dla T = [1, 5, 3, 8, 2], k = 2:**

**Algorytm A** (sprawdzamy wszystkie pary i < j):

| Para (i,j) | T[i] | T[j] | |T[i]-T[j]| | <= 2? |
|------------|------|------|------------|-------|
| (1,2) | 1 | 5 | 4 | nie |
| (1,3) | 1 | 3 | 2 | tak |
| (1,4) | 1 | 8 | 7 | nie |
| (1,5) | 1 | 2 | 1 | tak |
| (2,3) | 5 | 3 | 2 | tak |
| (2,4) | 5 | 8 | 3 | nie |
| (2,5) | 5 | 2 | 3 | nie |
| (3,4) | 3 | 8 | 5 | nie |
| (3,5) | 3 | 2 | 1 | tak |
| (4,5) | 8 | 2 | 6 | nie |

Wynik A: **4 pary**

**Algorytm B** (po posortowaniu):

T posortowane: [1, 2, 3, 5, 8]

| i | T[i] | T[i]+k | szukaj ostatni <= T[i]+k | j | wynik += j-i |
|---|------|--------|--------------------------|---|-------------|
| 1 | 1 | 3 | T[3]=3 <= 3 | 3 | 3-1 = 2 |
| 2 | 2 | 4 | T[3]=3 <= 4 | 3 | 3-2 = 1 |
| 3 | 3 | 5 | T[4]=5 <= 5 | 4 | 4-3 = 1 |
| 4 | 5 | 7 | T[4]=5 <= 7, T[5]=8 > 7 | 4 | 4-4 = 0 |
| 5 | 8 | 10 | brak elementow po i=5 | 5 | 5-5 = 0 |

Wynik B: 2 + 1 + 1 + 0 + 0 = **4 pary**

Wyniki sie zgadzaja.

**c) Dokladna liczba porownan dla n = 1000:**

ceil(log2(1000)) = 10

Algorytm A:
- Porownania w petlach: n*(n-1)/2 = 1000*999/2 = **499 500 porownan**

Algorytm B:
- Sortowanie: n * ceil(log2(n)) = 1000 * 10 = 10 000 porownan
- Wyszukiwania binarne: n * ceil(log2(n)) = 1000 * 10 = 10 000 porownan
- Razem: **20 000 porownan**

Algorytm B jest 499500/20000 = **~25 razy szybszy** dla n = 1000.

Prog oplacalnosci: Algorytm B jest lepszy gdy n*(n-1)/2 > 2*n*log2(n), czyli w przyblizeniu n > 4*log2(n), co zachodzi juz dla n >= 16.
</details>
