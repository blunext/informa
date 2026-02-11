# 03. Analiza algorytmu

Typ zadania: **analiza_algorytmu**
Czestotliwosc: 10/11 lat | Laczna punktacja: 37 pkt
Kategoria: TEORIA

## Umiejetnosci cwiczone w tym zestawie

`zlozonosc-czasowa` `O-notacja` `petle-zagniezdzne` `przeszukiwanie-binarne` `sortowanie` `kontrprzyklad` `dzielniki` `optymalizacja-sqrt` `rekurencja` `rownanie-rekurencyjne` `indukcja` `niezmiennik-petli` `drzewo-rekursji` `analiza-srednia`

---

### Cwiczenie 3.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 3.1 (zlozonosc), styl P/F z matur
**Tagi**: `zlozonosc-czasowa` `O-notacja` `petle-zagniezdzne` `przeszukiwanie-binarne`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Policz ile razy wykonuje sie operacja dominujaca (np. dodawanie, porownanie) w kazdym fragmencie.
2. **Podejscie**: Jedna petla = O(n), dwie zagniezdzne = O(n^2), trzy = O(n^3). Szukanie binarne polowi zakres.
3. **Kluczowy krok**: Fragment C: zakres [p,q] zmniejsza sie o polowe w kazdym kroku. Ilosc krokow = log2(n).

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Mylenie O(n) z O(n^2)**: Dwie petle NIEzagniezdzne (jedna po drugiej) to O(n)+O(n)=O(n), nie O(n^2). O(n^2) wymaga zagniezdenia. CKE: -1 pkt
- **Zapomnienie o log n**: Przeszukiwanie binarne to O(log n), nie O(n). Jesli zakres dzieli sie na pol, to logarytm. CKE: -1 pkt

</details>

---

### Cwiczenie 3.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 zad. 1.2 (kontrprzyklady), Matura 2022 zad. 2
**Tagi**: `sortowanie` `kontrprzyklad` `zlozonosc-czasowa`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Optymalizacja polega na wczesnym zakonczeniu, gdy brak zamian — kiedy to sie zdarzy?
2. **Podejscie**: Zastanow sie, jaka tablica spowoduje zakonczenie po JEDNYM przebiegu petli zewnetrznej.
3. **Kluczowy krok**: Tablica juz posortowana -> 0 zamian w pierwszym przebiegu -> flaga FALSZ -> koniec po n-1 porownaniach.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Podanie kontrprzykladu bez obliczen**: Trzeba policzyc faktyczna liczbe porownan, nie tylko napisac "inaczej". CKE: -1 pkt
- **Kontrprzyklad z bledna tablica**: Np. [5,4,3,2,1] nie jest kontrprzykladem — daje n*(n-1)/2 porownan. CKE: -1 pkt
- **Zapomnienie o fladze zamiana**: Algorytm konczy sie wczesniej dzieki fladze. Bez niej zawsze bylby n*(n-1)/2. CKE: -0.5 pkt

</details>

---

### Cwiczenie 3.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 1b (Korale)
**Tagi**: `zlozonosc-czasowa` `analiza-srednia` `przeszukiwanie-binarne`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Minimum = najlepszy przypadek, maksimum = najgorszy przypadek. Kiedy algorytm konczy sie najszybciej/najwolniej?
2. **Podejscie**: Najlepszy: x jest na poczatku. Najgorszy: x jest na koncu lub go nie ma. Srednia: policz oczekiwana wartosc.
3. **Kluczowy krok**: Srednia = (1+2+...+n)/n = n(n+1)/(2n) = (n+1)/2.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Srednia != mediana**: Srednia to (n+1)/2, nie n/2. Czesta pomylka. CKE: -1 pkt
- **Pomieszanie "element nie istnieje" z "element na koncu"**: Oba daja n porownan, ale to rozne sytuacje. CKE: -0.5 pkt
- **Brak zalozenia w c)**: Trzeba jasno okreslic, ze x jest w tablicy. Jezeli x moze nie byc, srednia jest inna. CKE: -0.5 pkt

</details>

---

### Cwiczenie 3.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 1 (liczby skojarzone)
**Tagi**: `dzielniki` `optymalizacja-sqrt` `zlozonosc-czasowa`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Dzielniki wystepuja w parach (i, n/i). Jesli i jest dzielnikiem, to n/i tez.
2. **Podejscie**: Wystarczy sprawdzic i od 1 do sqrt(n). Dla kazdego dzielnika i, dodaj 2 (lub 1 jesli i=sqrt(n)).
3. **Kluczowy krok**: sqrt(10000)=100, wiec zamiast 10000 operacji wystarczy 100.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Podwojne liczenie sqrt(n)**: Dla n=16, i=4: 4*4=16=n, wiec n/i=4=i. Liczymy dzielnik raz, nie dwa razy. CKE: -1 pkt
- **Warunek i*i <= n zamiast i <= sqrt(n)**: Oba sa poprawne, ale i*i unika bledu zaokraglania sqrt. CKE: brak kary, ale warto wiedziec.
- **Zapomnienie o parze**: Liczyc tylko i bez n/i daje polowe dzielnikow. CKE: -2 pkt

</details>

---

### Cwiczenie 3.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 2 (ab-slowo)
**Tagi**: `zlozonosc-czasowa` `sortowanie` `przeszukiwanie-binarne` `analiza-srednia`

Problem: Dany zbior n liczb calkowitych. Zlicz ile jest par (i, j), gdzie i < j, takich ze |T[i] - T[j]| <= k (roznica bezwzgledna co najwyzej k).

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
<summary>Wskazowki</summary>

1. **Kierunek**: Algorytm A sprawdza kazda pare — ile jest par? Algorytm B sortuje i uzywa binsearch — ile razy?
2. **Podejscie**: A: dwie petle => O(n^2). B: sortowanie + n wyszukiwan binarnych => O(n log n) + O(n log n).
3. **Kluczowy krok**: Dla b) posortuj tablice i recznie przesled oba algorytmy. Wyniki musza sie zgadzac.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o koszcie sortowania**: Algorytm B ma O(n log n) za sortowanie — to nie jest "za darmo". CKE: -1 pkt
- **Bledna liczba par w algorytmie A**: Liczba par to n*(n-1)/2, nie n^2. CKE: -1 pkt
- **Pominecie warunku i < j**: Pary (i,j) i (j,i) to ta sama para. Liczymy kazda raz. CKE: -0.5 pkt

</details>

---

### Cwiczenie 3.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 1 (zlozonosc), typowe pytanie CKE
**Tagi**: `zlozonosc-czasowa` `O-notacja` `petle-zagniezdzne`

Przypisz zlozonosc czasowa kazdemu z ponizszych fragmentow kodu. Wybierz sposrod: {O(1), O(log n), O(n), O(n log n), O(n^2)}.

**Fragment A:**
```
s := T[1] + T[n]
```

**Fragment B:**
```
s := 0
dla i := 1, 2, ..., n:
    s := s + T[i]
dla j := 1, 2, ..., n:
    s := s - T[j]
```

**Fragment C:**
```
s := 0
k := n
dopoki k > 0:
    s := s + T[k]
    k := k div 2
```

**Fragment D:**
```
s := 0
dla i := 1, 2, ..., n:
    dla j := i, i+1, ..., n:
        s := s + 1
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Fragment A nie ma petli. Fragment B ma dwie NIEzagniezdzne petle. Fragment C polowi k. Fragment D ma trojkatna petle.
2. **Podejscie**: Policz ile razy wykonuje sie operacja wewnatrz kazdej petli/grupy petli.
3. **Kluczowy krok**: Fragment D: j idzie od i do n, wiec ilosc iteracji to n + (n-1) + ... + 1 = n(n+1)/2 = O(n^2).

</details>

<details>
<summary>Odpowiedz</summary>

- **Fragment A**: **O(1)** — brak petli, stala liczba operacji (2 odczyty i 1 dodawanie)
- **Fragment B**: **O(n)** — dwie petle sekwencyjne (niezagniezdzne), kazda O(n). Razem O(n) + O(n) = O(n)
- **Fragment C**: **O(log n)** — k zaczyna od n i w kazdym kroku jest dzielone przez 2: n, n/2, n/4, ..., 1. Liczba krokow = floor(log2(n)) + 1
- **Fragment D**: **O(n^2)** — petle zagniezdzne, ale petla wewnetrzna zaczyna od i (nie od 1). Laczna liczba iteracji: n + (n-1) + ... + 1 = n(n+1)/2 = O(n^2)

**Pulapka we fragmencie B**: Dwie petle pod rzad to O(n)+O(n) = O(2n) = O(n), NIE O(n^2)! Zagniezdzenie jest potrzebne do O(n^2).

**Pulapka we fragmencie D**: Mimo ze petla wewnetrzna nie zawsze robi n iteracji, suma jest n(n+1)/2 co jest O(n^2).
</details>

<details>
<summary>Typowe bledy</summary>

- **Dwie petle sekwencyjne != O(n^2)**: O(n)+O(n) = O(n), nie O(n^2). Trzeba ZAGNIEZDZENIA. CKE: -1 pkt
- **Trojkatna petla to nadal O(n^2)**: n(n+1)/2 = O(n^2), nie O(n). CKE: -1 pkt
- **O(1) pomieszane z O(n)**: Brak petli = stala ilosc operacji = O(1). CKE: -0.5 pkt

</details>

---

### Cwiczenie 3.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 2 (rekurencja)
**Tagi**: `rekurencja` `rownanie-rekurencyjne` `drzewo-rekursji` `zlozonosc-czasowa`

Dana jest funkcja rekurencyjna:

```
funkcja f(n)
    jezeli n <= 1:
        zwroc 1
    zwroc f(n-1) + f(n-2)
```

**Polecenie**:
- a) Ile razy funkcja `f` jest wywolywana podczas obliczania `f(6)`? (Wliczajac wywolanie poczatkowe.)
- b) Jaka zlozonosc czasowa ma ten algorytm obliczania f(n)?
- c) Podaj wersje iteracyjna tego algorytmu o zlozonosci O(n) i wyjasnij, dlaczego jest szybsza.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: f(n) oblicza n-ta liczbe Fibonacciego. Rekurencyjna wersja ma powtarzajace sie obliczenia.
2. **Podejscie**: Narysuj drzewo wywolan dla f(6). Policz wezly. Zauwazyc powtorzenia: f(3) jest liczone wiele razy.
3. **Kluczowy krok**: Liczba wywolan rosnie wykladniczo. Iteracyjna wersja uzywa dwoch zmiennych i petli.

</details>

<details>
<summary>Odpowiedz</summary>

**a) Drzewo wywolan f(6):**

```
f(6)
├── f(5)
│   ├── f(4)
│   │   ├── f(3)
│   │   │   ├── f(2)
│   │   │   │   ├── f(1) = 1
│   │   │   │   └── f(0) = 1
│   │   │   └── f(1) = 1
│   │   └── f(2)
│   │       ├── f(1) = 1
│   │       └── f(0) = 1
│   └── f(3)
│       ├── f(2)
│       │   ├── f(1) = 1
│       │   └── f(0) = 1
│       └── f(1) = 1
└── f(4)
    ├── f(3)
    │   ├── f(2)
    │   │   ├── f(1) = 1
    │   │   └── f(0) = 1
    │   └── f(1) = 1
    └── f(2)
        ├── f(1) = 1
        └── f(0) = 1
```

Liczba wywolan: **25**

Wynik: f(6) = 13

**b) Zlozonosc: O(2^n)** (wykladnicza)

Dokladniej O(phi^n), gdzie phi = (1+sqrt(5))/2 ≈ 1.618 (zloty podzial). Kazde wywolanie generuje dwa kolejne, wiec drzewo rosnie wykladniczo. Wiele wartosci jest obliczanych wielokrotnie (np. f(3) jest liczone 3 razy).

**c) Wersja iteracyjna O(n):**

```
funkcja f_iter(n)
    jezeli n <= 1:
        zwroc 1
    a := 1   // f(0)
    b := 1   // f(1)
    dla i := 2, 3, ..., n:
        c := a + b
        a := b
        b := c
    zwroc b
```

Jest szybsza, bo kazda wartosc f(i) oblicza dokladnie raz (w jednym kroku petli), bez powtorzen. Uzywa O(1) pamieci dodatkowej (tylko 3 zmienne) zamiast O(n) stosu rekursji.
</details>

<details>
<summary>Typowe bledy</summary>

- **Bledna liczba wywolan**: Trzeba dokladnie policzyc wezly drzewa. Dla f(6) jest ich 25, nie np. 15. CKE: -1 pkt
- **Zlozonosc O(n^2) zamiast O(2^n)**: Fibonacci rekurencyjny rosnie wykladniczo, nie kwadratowo. CKE: -1 pkt
- **Brak wyjasnnienia dlaczego iteracyjna jest szybsza**: Klucz: brak powtarzajacych sie obliczen. CKE: -0.5 pkt

</details>

---

### Cwiczenie 3.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3 (analiza kosztu)
**Tagi**: `zlozonosc-czasowa` `O-notacja` `niezmiennik-petli`

Dany jest algorytm:

```
funkcja tajemnicza(n)
    wynik := 0
    i := 1
    dopoki i <= n:
        dla j := 1, 2, ..., i:
            wynik := wynik + 1
        i := i * 2
    zwroc wynik
```

**Polecenie**:
- a) Przesled algorytm dla n = 16. Podaj wartosci i w kazdej iteracji petli zewnetrznej i ile razy wykonuje sie petla wewnetrzna.
- b) Jaki jest wynik `tajemnicza(16)`?
- c) Podaj zlozonosc czasowa algorytmu. Uzasadnij.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Petla zewnetrzna podwaja i (i*=2), wiec i przechodzi wartosci 1, 2, 4, 8, 16, ... Ile razy to sie zdarzy?
2. **Podejscie**: Petla wewnetrzna wykonuje i iteracji. Suma iteracji = 1 + 2 + 4 + 8 + 16 = ?
3. **Kluczowy krok**: Suma poteg dwojki 1+2+4+...+2^k = 2^(k+1) - 1. Ile jest takich wyrazow dla i <= n?

</details>

<details>
<summary>Odpowiedz</summary>

**a) Sledzenie dla n = 16:**

| Iteracja zew. | i | Petla wew. j=1..i | wynik += |
|---------------|---|-------------------|----------|
| 1 | 1 | j=1..1 -> 1 raz | wynik = 1 |
| 2 | 2 | j=1..2 -> 2 razy | wynik = 3 |
| 3 | 4 | j=1..4 -> 4 razy | wynik = 7 |
| 4 | 8 | j=1..8 -> 8 razy | wynik = 15 |
| 5 | 16 | j=1..16 -> 16 razy | wynik = 31 |
| (koniec) | 32 | 32 > 16, petla konczy sie | |

**b) Wynik: tajemnicza(16) = 1 + 2 + 4 + 8 + 16 = **31****

**c) Zlozonosc: O(n)**

Uzasadnienie:
- Petla zewnetrzna wykonuje sie log2(n) + 1 razy (bo i podwaja sie: 1, 2, 4, ..., n)
- Petla wewnetrzna w k-tej iteracji robi 2^(k-1) operacji
- Suma: 1 + 2 + 4 + ... + n = 2n - 1 = O(n)

Mimo ze mamy zagniezdzne petle, zlozonosc to O(n), nie O(n log n) ani O(n^2)! Petla wewnetrzna robi rozna liczbe iteracji, a ich suma to szereg geometryczny.

**Wzor ogolny**: tajemnicza(n) = 2n - 1 (dla n bedacych potega 2).
</details>

<details>
<summary>Typowe bledy</summary>

- **Zlozonosc O(n^2) lub O(n log n)**: Pozornie zagniezdzne petle, ale suma 1+2+4+...+n = 2n-1 = O(n). CKE: -2 pkt
- **Zapomnienie ze i podwaja sie, a nie rosnie o 1**: i := i*2, nie i := i+1. Petla zewnetrzna ma log2(n) iteracji. CKE: -1 pkt
- **Bledna suma szeregu**: 1+2+4+...+2^k = 2^(k+1)-1, nie 2^k. CKE: -0.5 pkt

</details>

---

### Cwiczenie 3.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 2 (specyfikacja i analiza)
**Tagi**: `indukcja` `niezmiennik-petli` `zlozonosc-czasowa`

Dany jest algorytm:

```
funkcja potega(x, n)
    // n >= 0 (calkowite)
    wynik := 1
    baza := x
    wyk := n
    dopoki wyk > 0:
        jezeli wyk mod 2 = 1:
            wynik := wynik * baza
        baza := baza * baza
        wyk := wyk div 2
    zwroc wynik
```

**Polecenie**:
- a) Przesled algorytm dla x=3, n=13. Podaj wartosci zmiennych w kazdym kroku.
- b) Ile mnozen wykonuje algorytm dla n=13? Ile wykona prosta petla `wynik := x * x * ... * x` (n-1 razy)?
- c) Podaj zlozonosc czasowa algorytmu w funkcji n. Uzasadnij.
- d) Podaj niezmiennik petli: jaki zwiazek zachodzi miedzy `wynik`, `baza`, `wyk` a oryginalnym `x^n`?

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: To szybkie potegowanie (binary exponentiation). Wykladnik jest analizowany bit po bicie.
2. **Podejscie**: Rozpisz n=13 w systemie binarnym: 13 = 1101₂. Bity odpowiadaja krokom algorytmu.
3. **Kluczowy krok**: Niezmiennik: wynik * baza^wyk = x^n jest prawdziwy na poczatku kazdej iteracji.

</details>

<details>
<summary>Odpowiedz</summary>

**a) Sledzenie dla x=3, n=13:**

13 w systemie binarnym: 1101₂

| Krok | wyk | wyk%2 | wynik | baza | operacje |
|------|-----|-------|-------|------|----------|
| pocz. | 13 | - | 1 | 3 | - |
| 1 | 13 | 1 (nieparzyste) | 1*3=3 | 3*3=9 | wyk=13 div 2=6 |
| 2 | 6 | 0 (parzyste) | 3 (bez zmiany) | 9*9=81 | wyk=6 div 2=3 |
| 3 | 3 | 1 (nieparzyste) | 3*81=243 | 81*81=6561 | wyk=3 div 2=1 |
| 4 | 1 | 1 (nieparzyste) | 243*6561=1594323 | 6561*6561=... | wyk=1 div 2=0 |

Wynik: 3^13 = **1594323**

Sprawdzenie: 3^13 = 1594323. Poprawne.

**b) Liczba mnozen:**

- Szybkie potegowanie: W kazdym kroku 1 mnozenie (baza*baza) + ewentualnie 1 (wynik*baza). Dla n=13: 4 kroki, 4 mnozenia baza + 3 mnozenia wynik = **7 mnozen**
- Prosta petla: n-1 = **12 mnozen**

Oszczednosc: prawie 2x mniej mnozen. Dla duzych n roznica jest drastyczna.

**c) Zlozonosc: O(log n)**

Uzasadnienie: W kazdym kroku wyk jest dzielone przez 2 (wyk := wyk div 2). Liczba iteracji to floor(log2(n)) + 1. Kazda iteracja wykonuje stala liczbe operacji (1-2 mnozenia). Stad O(log n).

Porownanie: prosta petla to O(n). Dla n=1000000 szybkie potegowanie robi ~20 krokow zamiast 999999.

**d) Niezmiennik petli:**

**wynik * baza^wyk = x^n**

Na poczatku: 1 * x^n = x^n. Prawda.
W kazdym kroku:
- Jezeli wyk nieparzyste: wynik' = wynik*baza, baza' = baza^2, wyk' = (wyk-1)/2
  wynik' * baza'^wyk' = (wynik*baza) * (baza^2)^((wyk-1)/2) = wynik * baza * baza^(wyk-1) = wynik * baza^wyk
- Jezeli wyk parzyste: wynik' = wynik, baza' = baza^2, wyk' = wyk/2
  wynik' * baza'^wyk' = wynik * (baza^2)^(wyk/2) = wynik * baza^wyk

Na koncu (wyk=0): wynik * baza^0 = wynik = x^n.
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o mnozeniu baza*baza w KAZDYM kroku**: Baza jest podnoszona do kwadratu niezaleznie od parzystosci wyk. CKE: -1 pkt
- **Bledne obliczenie 3^13**: Duze liczby — latwo o blad. Warto weryfikowac krok po kroku. CKE: -1 pkt
- **Niezmiennik bez dowodu**: Trzeba pokazac, ze zachowuje sie w kazdym kroku, nie tylko podac go. CKE: -1 pkt

</details>

---

### Cwiczenie 3.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 2 (analiza zlozonosci rekurencji)
**Tagi**: `rekurencja` `rownanie-rekurencyjne` `zlozonosc-czasowa` `indukcja`

Dany jest algorytm sortowania przez scalanie:

```
funkcja mergeSort(T, lewy, prawy)
    jezeli lewy >= prawy:
        zakoncz
    srodek := (lewy + prawy) div 2
    mergeSort(T, lewy, srodek)
    mergeSort(T, srodek + 1, prawy)
    scal(T, lewy, srodek, prawy)   // koszt scalania: prawy - lewy + 1 porownan
```

**Polecenie**:
- a) Dla n = 8 elementow narysuj drzewo rekursji (jakie zakresy sa przetwarzane na kazdym poziomie).
- b) Ile razy jest wywolywana funkcja mergeSort (wliczajac wywolania bazowe, gdy lewy >= prawy)?
- c) Ile lacznie porownan wykonuje operacja `scal` na kazdym poziomie drzewa? Jaka jest laczna liczba porownan?
- d) Udowodnij ze zlozonosc T(n) = O(n log n), uzywajac rownania rekurencyjnego T(n) = 2T(n/2) + n.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Kazdy poziom drzewa przerabia lacznie n elementow (rozne fragmenty tablicy). Ile jest poziomow?
2. **Podejscie**: Drzewo ma log2(n) + 1 poziomow. Na kazdym poziomie laczny koszt scalania = n.
3. **Kluczowy krok**: Razem: n * (log2(n) + 1) = O(n log n). Rownanie T(n) = 2T(n/2) + n rozwiazuje sie przez rozwijanie.

</details>

<details>
<summary>Odpowiedz</summary>

**a) Drzewo rekursji dla n=8 (indeksy 0..7):**

```
Poziom 0: mergeSort(0,7)               -> scal 8 elem.
          /                  \
Poziom 1: mergeSort(0,3)     mergeSort(4,7)     -> scal 4+4=8 elem.
         /          \        /          \
Poziom 2: mS(0,1)  mS(2,3)  mS(4,5)  mS(6,7)   -> scal 2+2+2+2=8 elem.
         / \       / \       / \       / \
Poziom 3: mS(0) mS(1) mS(2) mS(3) mS(4) mS(5) mS(6) mS(7) -> bazy (brak scalania)
```

**b) Liczba wywolan mergeSort:**

- Poziom 0: 1
- Poziom 1: 2
- Poziom 2: 4
- Poziom 3 (bazowe): 8

Razem: 1 + 2 + 4 + 8 = **15 wywolan**

Ogolnie: 2n - 1 wywolan (dla n bedacego potega 2).

**c) Porownania na kazdym poziomie:**

Koszt scalania = liczba elementow w zakresie.

| Poziom | Zakresy | Laczna liczba elementow do scalenia | Porownania |
|--------|---------|--------------------------------------|------------|
| 0 | (0,7) | 8 | 8 |
| 1 | (0,3), (4,7) | 4 + 4 = 8 | 8 |
| 2 | (0,1), (2,3), (4,5), (6,7) | 2+2+2+2 = 8 | 8 |
| 3 | bazowe | 0 | 0 |

Laczna liczba porownan: 8 + 8 + 8 = 8 * 3 = **24 porownania**

Ogolnie: n * log2(n) (dla n bedacego potega 2). Tu: 8 * 3 = 24.

**d) Dowod O(n log n) przez rozwijanie rownania:**

T(n) = 2T(n/2) + n
     = 2[2T(n/4) + n/2] + n = 4T(n/4) + 2n
     = 4[2T(n/8) + n/4] + 2n = 8T(n/8) + 3n
     = ...
     = 2^k * T(n/2^k) + k*n

Dla k = log2(n): 2^k = n, T(1) = O(1)

T(n) = n * O(1) + log2(n) * n = **O(n log n)**

Jest to najlepsza mozliwa zlozonosc dla sortowania opartego na porownaniach (dolna granica to Omega(n log n)).
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o kosztcie scalania**: scal() kosztuje O(rozmiar zakresu), nie O(1). CKE: -2 pkt
- **Bledna liczba poziomow**: Dla n=8 sa 4 poziomy (0,1,2,3), nie 3. Poziom 3 to bazy (bez scalania). CKE: -1 pkt
- **Bledne rozwijanie rownania**: Kazdy krok podstawia T(n/2^k), a koszt rosnie o n na poziom. CKE: -1 pkt
- **Brak warunku bazowego T(1)**: Rozwijanie musi dojsc do T(1) = stala. CKE: -0.5 pkt

</details>

---

## Samoocena

Po rozwiazaniu cwiczen bez podgladania odpowiedzi, okresl swoj poziom:

| Poziom | Opis | Wynik |
|--------|------|-------|
| Podstawowy | Potrafisz przypisac zlozonosc prostym petlom i rozumiesz O-notacje | 1-3 cwiczen bez pomocy |
| Dobry | Radzisz sobie z kontrprzykladami, analiza srednia i optymalizacja sqrt | 4-6 cwiczen bez pomocy |
| Bardzo dobry | Analizujesz zlozonosc rekurencji, szybkie potegowanie i niezmienniki petli | 7-8 cwiczen bez pomocy |
| Doskonaly | Potrafisz dowodzic zlozonosc O(n log n), rozwijac rownania rekurencyjne | 9-10 cwiczen bez pomocy |

**Co dalej?**
- Poziom Podstawowy: Przerob cwiczenia 3.1, 3.2, 3.6 jeszcze raz. Wrocz do `cheatsheet_teoria.md` (sekcja: notacja O).
- Poziom Dobry: Skup sie na cwiczeniach 3.4, 3.7, 3.8. Przejdz do `01_sledzenie_algorytmu.md`.
- Poziom Bardzo dobry/Doskonaly: Przejdz do `02_projektowanie_algorytmu.md` i `09_zlozone.md`.
