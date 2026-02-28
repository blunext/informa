# 11. Wyszukiwanie min/max

Typ zadania: **minmax**
Czestotliwosc: 5/12 lat | Laczna punktacja: 17 pkt
Kategoria: IMPLEMENTACJA

## Umiejetnosci cwiczone w tym zestawie

`min-max` `pozycja-elementu` `skan-liniowy` `wczytywanie-pliku` `rekordy` `sortowanie` `k-ty-element` `mediana` `tablica-2D` `srednia` `roznica-bezwzgledna` `spojny-fragment` `vector-par` `warunkowe-min-max` `struct` `wielokryterialne`

---

### Cwiczenie 11.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 zad. 4.3
**Tagi**: `min-max` `pozycja-elementu` `skan-liniowy` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 12 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory znajdzie wartosc minimalna i maksymalna oraz poda ich pozycje (numerowane od 1).

**Dane** (`dane.txt`):
```
45
12
78
34
91
23
67
8
56
82
15
39
```

**Oczekiwany wynik**:
```
Minimum: 8 (pozycja: 8)
Maksimum: 91 (pozycja: 5)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Inicjalizuj min/max pierwszym elementem, potem skanuj reszte.
2. **Podejscie**: Sledzac min/max, pamietaj tez ich pozycje (indeks + 1).
3. **Kluczowy krok**: Warunki aktualizacji: `tab[i] < minVal` i `tab[i] > maxVal`.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    vector<int> tab;
    int x;
    while (plik >> x) tab.push_back(x);

    int minVal = tab[0], maxVal = tab[0];
    int minPos = 1, maxPos = 1;
    for (int i = 1; i < tab.size(); i++) {
        if (tab[i] < minVal) { minVal = tab[i]; minPos = i + 1; }
        if (tab[i] > maxVal) { maxVal = tab[i]; maxPos = i + 1; }
    }
    cout << "Minimum: " << minVal << " (pozycja: " << minPos << ")" << endl;
    cout << "Maksimum: " << maxVal << " (pozycja: " << maxPos << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Klasyczny skan liniowy — inicjalizujemy min/max pierwszym elementem, potem iterujemy i aktualizujemy przy znalezieniu mniejszego/wiekszego.

Weryfikacja: Ciag: 45,12,78,34,91,23,67,8,56,82,15,39. Min=8 (poz.8), Max=91 (poz.5).
</details>

<details>
<summary>Typowe bledy</summary>

- **Inicjalizacja min=0**: Jesli wszystkie liczby sa dodatnie, min=0 nigdy nie zostanie nadpisane. Inicjalizuj pierwszym elementem lub `INT_MAX`. CKE: -1 pkt
- **Indeks od 0 zamiast od 1**: CKE numeruje pozycje od 1. CKE: -1 pkt (za kazda bledna pozycje)

</details>

---

### Cwiczenie 11.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4.2
**Tagi**: `min-max` `rekordy` `skan-liniowy` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 10 rekordow w formacie: nazwa wartosc (oddzielone spacja). Napisz program, ktory poda nazwe o najwyzszej i najnizszej wartosci.

**Dane** (`dane.txt`):
```
Alfa 150
Beta 87
Gamma 203
Delta 45
Epsilon 178
Zeta 92
Eta 310
Theta 56
Iota 134
Kappa 267
```

**Oczekiwany wynik**:
```
Najwyzsza wartosc: Eta (310)
Najnizsza wartosc: Delta (45)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Czytaj nazwe i wartosc w petli. Sledzic trzeba max/min wartosc i towarzyszaca nazwe.
2. **Podejscie**: `plik >> nazwa >> wartosc` wczytuje jedno pole string i jedno int.
3. **Kluczowy krok**: Przy aktualizacji maxVal/minVal, zaktualizuj tez maxNazwa/minNazwa.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    string nazwa, maxN, minN;
    int wartosc, maxV = -1, minV = 1000000;

    while (plik >> nazwa >> wartosc) {
        if (wartosc > maxV) { maxV = wartosc; maxN = nazwa; }
        if (wartosc < minV) { minV = wartosc; minN = nazwa; }
    }
    cout << "Najwyzsza wartosc: " << maxN << " (" << maxV << ")" << endl;
    cout << "Najnizsza wartosc: " << minN << " (" << minV << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Skan liniowy z zapamietywaniem nazwy towarzyszacej ekstremalnej wartosci.

Weryfikacja: Max=Eta(310), Min=Delta(45).
</details>

<details>
<summary>Typowe bledy</summary>

- **Inicjalizacja maxV=0 przy danych ktore moga byc ujemne**: Jesli dane sa ujemne, max=0 nie zostanie nadpisane. CKE: -1 pkt
- **Zapomnienie o zapamietaniu nazwy**: Wypisanie samej wartosci bez nazwy to niepelna odpowiedz. CKE: -1 pkt

</details>

---

### Cwiczenie 11.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4.2 (k-ta liczba pierwsza)
**Tagi**: `sortowanie` `k-ty-element` `mediana` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 20 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory:
a) Poda 5-ta najmniejsza liczbe.
b) Poda 3-cia najwieksza liczbe.
c) Poda mediane (wartosc srodkowa po posortowaniu; dla parzystej liczby elementow — srednia dwoch srodkowych).

**Dane** (`dane.txt`):
```
42
17
85
3
56
91
28
64
12
73
38
95
7
50
21
69
34
88
46
60
```

**Oczekiwany wynik**:
```
Po sortowaniu: 3 7 12 17 21 28 34 38 42 46 50 56 60 64 69 73 85 88 91 95
a) 5-ta najmniejsza: 21
b) 3-cia najwieksza: 88
c) Mediana: 48.00
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Posortuj tablice rosnaco — `sort(tab.begin(), tab.end())`.
2. **Podejscie**: K-ta najmniejsza to `tab[k-1]`. K-ta najwieksza to `tab[n-k]`.
3. **Kluczowy krok**: Mediana dla parzystego n: `(tab[n/2-1] + tab[n/2]) / 2.0` — pamietaj o dzieleniu zmiennoprzecinkowym.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    vector<int> tab;
    int x;
    while (plik >> x) tab.push_back(x);

    sort(tab.begin(), tab.end());
    int n = tab.size();

    cout << "Po sortowaniu: ";
    for (int v : tab) cout << v << " ";
    cout << endl;

    cout << "a) 5-ta najmniejsza: " << tab[4] << endl;
    cout << "b) 3-cia najwieksza: " << tab[n - 3] << endl;

    double mediana;
    if (n % 2 == 0)
        mediana = (tab[n/2 - 1] + tab[n/2]) / 2.0;
    else
        mediana = tab[n/2];
    cout << "c) Mediana: " << fixed << setprecision(2) << mediana << endl;
    return 0;
}
```

**Wyjasnienie**: Sortujemy tablice rosnaco. K-ty najmniejszy element to `tab[k-1]`, k-ty najwiekszy to `tab[n-k]`. Mediana dla parzystej liczby elementow to srednia dwoch srodkowych: `(tab[9] + tab[10]) / 2.0`.

Weryfikacja:
Posortowane: 3,7,12,17,21,28,34,38,42,46,50,56,60,64,69,73,85,88,91,95
a) tab[4] = 21
b) tab[17] = 88
c) tab[9]=46, tab[10]=50 -> mediana = (46+50)/2 = 48.00
</details>

<details>
<summary>Typowe bledy</summary>

- **Dzielenie calkowite w medianie**: `(46 + 50) / 2 = 48` (OK akurat tutaj), ale np. `(3+4)/2 = 3` zamiast 3.5. Uzyj `/2.0`. CKE: -1 pkt
- **Indeks k-tej zamiast k-1**: `tab[5]` to 6-ta najmniejsza, nie 5-ta. CKE: -1 pkt
- **Brak sortowania**: K-ty element bez sortowania to element na k-tej pozycji, nie k-ty co do wielkosci. CKE: -2 pkt

</details>

---

### Cwiczenie 11.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2015 zad. 4.3 (min/max w wierszu tablicy)
**Tagi**: `tablica-2D` `min-max` `srednia` `wczytywanie-pliku`

Dana jest tablica 4x6 (4 wiersze, 6 kolumn) zapisana w pliku `tablica.txt` (kazdy wiersz tablicy w osobnym wierszu, wartosci oddzielone spacjami). Napisz program, ktory:
a) Znajdzie wiersz zawierajacy wartosc maksymalna calej tablicy.
b) W tym wierszu poda wartosc minimalna.
c) Znajdzie kolumne z najmniejsza srednia.

**Dane** (`tablica.txt`):
```
12 45 7 23 56 31
89 14 67 42 8 53
38 71 19 94 26 60
5 33 82 17 48 11
```

**Oczekiwany wynik**:
```
a) Wiersz z wartoscia max: wiersz 3 (wartosc max: 94)
b) Minimum w wierszu 3: 19
c) Kolumna z najmniejsza srednia: kolumna 5 (srednia: 34.50)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wczytaj tablice 2D, potem przetwarzaj wierszami i kolumnami.
2. **Podejscie**: (a) Podwojna petla szukajaca globalnego max + zapamietanie wiersza. (b) Petla po kolumnach jednego wiersza. (c) Petla po kolumnach z obliczeniem sredniej.
3. **Kluczowy krok**: Srednia kolumny = suma elementow w kolumnie / liczba wierszy. Pamietaj o dzieleniu zmiennoprzecinkowym.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("tablica.txt");
    int tab[4][6];
    for (int i = 0; i < 4; i++)
        for (int j = 0; j < 6; j++)
            plik >> tab[i][j];

    // a) Wiersz z wartoscia max
    int maxVal = tab[0][0], maxWiersz = 0;
    for (int i = 0; i < 4; i++)
        for (int j = 0; j < 6; j++)
            if (tab[i][j] > maxVal) { maxVal = tab[i][j]; maxWiersz = i; }

    cout << "a) Wiersz z wartoscia max: wiersz " << maxWiersz + 1
         << " (wartosc max: " << maxVal << ")" << endl;

    // b) Min w tym wierszu
    int minVal = tab[maxWiersz][0];
    for (int j = 1; j < 6; j++)
        if (tab[maxWiersz][j] < minVal) minVal = tab[maxWiersz][j];
    cout << "b) Minimum w wierszu " << maxWiersz + 1 << ": " << minVal << endl;

    // c) Kolumna z najmniejsza srednia
    double minSr = 1e9;
    int minKol = 0;
    for (int j = 0; j < 6; j++) {
        double suma = 0;
        for (int i = 0; i < 4; i++) suma += tab[i][j];
        double sr = suma / 4;
        if (sr < minSr) { minSr = sr; minKol = j; }
    }
    cout << "c) Kolumna z najmniejsza srednia: kolumna " << minKol + 1
         << " (srednia: " << fixed << setprecision(2) << minSr << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Klasyczne operacje na tablicy 2D. (a) Podwojna petla po calej tablicy szukajac max. (b) Petla po jednym wierszu szukajac min. (c) Petla po kolumnach liczac srednie.

Weryfikacja:
a) Max tablicy: 94 w wierszu 3 (indeks 2), pozycja [2][3]
b) Wiersz 3: 38,71,19,94,26,60 -> min=19
c) Srednie kolumn:
   - kol.1: (12+89+38+5)/4 = 144/4 = 36.00
   - kol.2: (45+14+71+33)/4 = 163/4 = 40.75
   - kol.3: (7+67+19+82)/4 = 175/4 = 43.75
   - kol.4: (23+42+94+17)/4 = 176/4 = 44.00
   - kol.5: (56+8+26+48)/4 = 138/4 = 34.50
   - kol.6: (31+53+60+11)/4 = 155/4 = 38.75
   Najmniejsza srednia: kolumna 5 (34.50)
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomylenie wierszy i kolumn**: Indeks `[i][j]` — i to wiersz, j to kolumna. Zamiana powoduje transponowanie. CKE: -2 pkt
- **Dzielenie calkowite w sredniej**: `suma/4` zamiast `suma/4.0`. CKE: -1 pkt
- **Numerowanie od 0**: Wypisanie "wiersz 2" zamiast "wiersz 3". CKE: -1 pkt

</details>

---

### Cwiczenie 11.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2.3 + 2024 zad. 4
**Tagi**: `roznica-bezwzgledna` `spojny-fragment` `min-max` `podwojna-petla`

W pliku `dane.txt` znajduje sie 15 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory:
a) Znajdzie pare indeksow (i, j), gdzie i < j, dla ktorej |T[i] - T[j]| jest maksymalne.
b) Znajdzie pare sasiadujacych elementow (indeksy i, i+1) o najmniejszej roznicy bezwzglednej.
c) Znajdzie najdluzszy spojny fragment tablicy, w ktorym roznica miedzy max a min nie przekracza 10.

**Dane** (`dane.txt`):
```
25
18
22
20
19
24
21
50
45
48
47
46
3
80
5
```

**Oczekiwany wynik**:
```
a) Max roznica: |T[13]-T[14]| = |3-80| = 77 (pozycje 13, 14)

b) Min roznica sasiadow: |T[4]-T[5]| = |20-19| = 1 (pozycje 4, 5)

c) Najdluzszy fragment z max-min <= 10:
   Pozycje 1-7: [25,18,22,20,19,24,21], max=25, min=18, roznica=7 <= 10
   Dlugosc: 7
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: (a) Max roznica = max_element - min_element calego ciagu. (b) Skan par sasiadow. (c) Podwojna petla z sledzeniem min/max fragmentu.
2. **Podejscie**: Dla (a) wystarczy znalezc globalne min i max. Dla (c) rozszerzaj fragment od kazdej pozycji i, sledzac min/max — przerwij gdy roznica > 10.
3. **Kluczowy krok**: W (c) `break` przy roznica > 10 jest OK, bo dodanie kolejnych elementow nie zmniejszy roznicy.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <cmath>
#include <algorithm>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    // a) Max |T[i]-T[j]|, i<j
    int maxDiff = 0, ai = 0, aj = 0;
    for (int i = 0; i < n; i++)
        for (int j = i + 1; j < n; j++)
            if (abs(T[i] - T[j]) > maxDiff) {
                maxDiff = abs(T[i] - T[j]);
                ai = i; aj = j;
            }
    cout << "a) Max roznica: |T[" << ai+1 << "]-T[" << aj+1 << "]| = |"
         << T[ai] << "-" << T[aj] << "| = " << maxDiff
         << " (pozycje " << ai+1 << ", " << aj+1 << ")" << endl;

    // b) Min roznica sasiadow
    int minDiff = abs(T[0] - T[1]);
    int bi = 0;
    for (int i = 1; i < n - 1; i++) {
        int d = abs(T[i] - T[i + 1]);
        if (d < minDiff) { minDiff = d; bi = i; }
    }
    cout << endl << "b) Min roznica sasiadow: |T[" << bi+1 << "]-T[" << bi+2
         << "]| = |" << T[bi] << "-" << T[bi+1] << "| = " << minDiff
         << " (pozycje " << bi+1 << ", " << bi+2 << ")" << endl;

    // c) Najdluzszy fragment z max-min <= 10
    int bestLen = 0, bestStart = 0;
    for (int i = 0; i < n; i++) {
        int curMin = T[i], curMax = T[i];
        for (int j = i; j < n; j++) {
            curMin = min(curMin, T[j]);
            curMax = max(curMax, T[j]);
            if (curMax - curMin <= 10) {
                if (j - i + 1 > bestLen) {
                    bestLen = j - i + 1;
                    bestStart = i;
                }
            } else break;
        }
    }
    cout << endl << "c) Najdluzszy fragment z max-min <= 10:" << endl;
    cout << "   Pozycje " << bestStart+1 << "-" << bestStart+bestLen
         << ": [";
    for (int i = bestStart; i < bestStart + bestLen; i++) {
        if (i > bestStart) cout << ",";
        cout << T[i];
    }
    int fMin = *min_element(T.begin()+bestStart, T.begin()+bestStart+bestLen);
    int fMax = *max_element(T.begin()+bestStart, T.begin()+bestStart+bestLen);
    cout << "], max=" << fMax << ", min=" << fMin
         << ", roznica=" << fMax-fMin << " <= 10" << endl;
    cout << "   Dlugosc: " << bestLen << endl;
    return 0;
}
```

**Wyjasnienie**: (a) Podwojna petla O(n^2) sprawdzajaca wszystkie pary — mozna tez znalezc globalne min i max. (b) Petla po parach sasiadow. (c) Podwojna petla — dla kazdego poczatku rozszerzamy fragment aktualizujac min/max, dopoki roznica <= 10.

Weryfikacja:
a) Globalne min=3 (poz.13), max=80 (poz.14). Para (13,14): |3-80|=77.
b) Roznice sasiadow: |25-18|=7, |18-22|=4, |22-20|=2, |20-19|=1, |19-24|=5, |24-21|=3, |21-50|=29, |50-45|=5, |45-48|=3, |48-47|=1, |47-46|=1, |46-3|=43, |3-80|=77, |80-5|=75
   Min=1 (pozycje 4,5 lub 10,11 lub 11,12). Pierwsza: pozycje 4,5 (20,19).
c) Fragment 1-7: [25,18,22,20,19,24,21], min=18, max=25, diff=7. Element 8 (50) sprawia ze diff=32. Dlugosc 7.
   Fragment 8-12: [50,45,48,47,46], min=45, max=50, diff=5. Dlugosc 5.
   Najdluzszy: 7 (pozycje 1-7).
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak `abs()` w roznicy**: Bez wartosci bezwzglednej, ujemne roznice nie beda porownywane poprawnie. CKE: -1 pkt
- **Inicjalizacja minDiff=0 w punkcie (b)**: Roznica 0 jest mniejsza niz kazda inna, wiec min nigdy nie zostanie nadpisane. CKE: -1 pkt
- **Brak `break` w punkcie (c)**: Bez break program dziala poprawnie ale wolniej (O(n^3) zamiast O(n^2) w praktyce). Nie jest to blad logiczny, ale moze powodowac timeout na duzych danych.

</details>

---

### Cwiczenie 11.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3 (min/max warunkowe)
**Tagi**: `min-max` `filtrowanie` `wczytywanie-pliku` `warunkowe-min-max`

W pliku `dane.txt` znajduje sie 12 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory:
a) Znajdzie najwieksza liczbe parzysta.
b) Znajdzie najmniejsza liczbe nieparzysta.

**Dane** (`dane.txt`):
```
17
42
8
55
36
91
24
63
100
5
78
33
```

**Oczekiwany wynik**:
```
a) Najwieksza parzysta: 100
b) Najmniejsza nieparzysta: 5
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dwa osobne skany — w kazdym sprawdzaj dodatkowy warunek (parzystosc/nieparzystosc).
2. **Podejscie**: Inicjalizuj max parzysta jako -1e9 (lub `INT_MIN`), min nieparzysta jako 1e9 (lub `INT_MAX`).
3. **Kluczowy krok**: Jesli `n % 2 == 0` i `n > maxP` -> aktualizuj maxP.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <climits>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    int n;
    int maxP = INT_MIN, minN = INT_MAX;

    while (plik >> n) {
        if (n % 2 == 0 && n > maxP) maxP = n;
        if (n % 2 != 0 && n < minN) minN = n;
    }

    cout << "a) Najwieksza parzysta: " << maxP << endl;
    cout << "b) Najmniejsza nieparzysta: " << minN << endl;
    return 0;
}
```

**Wyjasnienie**: Skan liniowy z warunkiem — aktualizujemy max/min tylko gdy element spelnia dodatkowe kryterium (parzystosc/nieparzystosc).

Weryfikacja:
- Parzyste: 42, 8, 36, 24, 100, 78. Max = 100.
- Nieparzyste: 17, 55, 91, 63, 5, 33. Min = 5.
</details>

<details>
<summary>Typowe bledy</summary>

- **Inicjalizacja max pierwszym elementem bez sprawdzenia parzystosci**: Jesli pierwszy element jest nieparzysty, to max parzystych zaczyna sie od zlej wartosci. CKE: -1 pkt
- **Brak obslugi przypadku gdy nie ma elementow spelniajacych warunek**: Jesli nie ma parzystych, maxP = INT_MIN to bezsensowny wynik. Trzeba sprawdzic.

</details>

---

### Cwiczenie 11.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4 (k-najlepszych)
**Tagi**: `sortowanie` `rekordy` `vector-par` `wczytywanie-pliku`

W pliku `wyniki.txt` znajduje sie 10 rekordow w formacie: imie wynik (oddzielone spacja, wynik to int). Napisz program, ktory:
a) Wypisze 3 najlepsze wyniki (najwyzsze) z imionami.
b) Wypisze najgorszy wynik z imieniem.
c) Wypisze sredni wynik calego zestawu.

**Dane** (`wyniki.txt`):
```
Anna 85
Bartek 72
Celina 93
Damian 61
Ewa 88
Filip 77
Grazyna 95
Hubert 69
Irena 91
Jan 84
```

**Oczekiwany wynik**:
```
a) Top 3:
   1. Grazyna (95)
   2. Celina (93)
   3. Irena (91)

b) Najgorszy: Damian (61)

c) Sredni wynik: 81.50
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wczytaj do vectora par `{wynik, imie}`, posortuj malejaco po wyniku.
2. **Podejscie**: `sort` z lambda `[](auto &a, auto &b) { return a.first > b.first; }`.
3. **Kluczowy krok**: Po sortowaniu: top 3 to pierwsze 3 elementy, najgorszy to ostatni. Srednia = suma / n.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <string>
#include <algorithm>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("wyniki.txt");
    vector<pair<int, string>> dane;
    string imie;
    int wynik;
    while (plik >> imie >> wynik)
        dane.push_back({wynik, imie});

    sort(dane.begin(), dane.end(), [](auto &a, auto &b) {
        return a.first > b.first;
    });

    cout << "a) Top 3:" << endl;
    for (int i = 0; i < 3; i++)
        cout << "   " << i + 1 << ". " << dane[i].second
             << " (" << dane[i].first << ")" << endl;

    cout << endl << "b) Najgorszy: " << dane.back().second
         << " (" << dane.back().first << ")" << endl;

    int suma = 0;
    for (auto &p : dane) suma += p.first;
    cout << endl << "c) Sredni wynik: " << fixed << setprecision(2)
         << (double)suma / dane.size() << endl;
    return 0;
}
```

**Wyjasnienie**: Sortowanie malejace pozwala latwo wyciagnac top k i najgorszy. Srednia = suma / n.

Weryfikacja:
- Posortowane malejaco: 95(Grazyna), 93(Celina), 91(Irena), 88(Ewa), 85(Anna), 84(Jan), 77(Filip), 72(Bartek), 69(Hubert), 61(Damian)
- Top 3: Grazyna(95), Celina(93), Irena(91)
- Najgorszy: Damian(61)
- Suma: 95+93+91+88+85+84+77+72+69+61 = 815. Srednia: 815/10 = 81.50
</details>

<details>
<summary>Typowe bledy</summary>

- **Sortowanie rosnace zamiast malejacego**: Top 3 bedzie na koncu, nie na poczatku. CKE: -1 pkt
- **Brak `(double)` w sredniej**: 815/10 = 81 (calkowite), nie 81.50. CKE: -1 pkt
- **Pomylenie kolejnosci w parze**: `{imie, wynik}` zamiast `{wynik, imie}` — sortowanie po imieniu zamiast po wyniku. CKE: -2 pkt

</details>

---

### Cwiczenie 11.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4 (min/max roznic)
**Tagi**: `roznica-bezwzgledna` `pary-elementow` `min-max` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 10 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory:
a) Dla kazdego elementu policzy, ile jest od niego wiekszych i ile mniejszych.
b) Znajdzie element, dla ktorego liczba wiekszych i mniejszych jest najbardziej zrownowazona (minimalna roznica |wieksze - mniejsze|).

**Dane** (`dane.txt`):
```
15
42
8
33
27
50
19
36
11
44
```

**Oczekiwany wynik**:
```
a) Statystyki:
   15: wiekszych=7, mniejszych=2, roznica=5
   42: wiekszych=2, mniejszych=7, roznica=5
   8: wiekszych=9, mniejszych=0, roznica=9
   33: wiekszych=4, mniejszych=5, roznica=1
   27: wiekszych=5, mniejszych=4, roznica=1
   50: wiekszych=0, mniejszych=9, roznica=9
   19: wiekszych=7, mniejszych=2, roznica=5
   36: wiekszych=3, mniejszych=6, roznica=3
   11: wiekszych=8, mniejszych=1, roznica=7
   44: wiekszych=1, mniejszych=8, roznica=7

b) Najbardziej zrownowazone: 33 (roznica=1)
   Rowniez: 27 (roznica=1)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dla kazdego elementu policz ile jest wiekszych i mniejszych — petla w petli.
2. **Podejscie**: Wewnetrzna petla porownuje element z kazdym innym (pomijajac siebie).
3. **Kluczowy krok**: Zrownowazona pozycja to taka, gdzie element jest "blisko mediany" — minimalna `|wieksze - mniejsze|`.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <cmath>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    cout << "a) Statystyki:" << endl;
    int minRoz = n;
    vector<int> bestIdx;

    for (int i = 0; i < n; i++) {
        int wiecej = 0, mniej = 0;
        for (int j = 0; j < n; j++) {
            if (j == i) continue;
            if (T[j] > T[i]) wiecej++;
            if (T[j] < T[i]) mniej++;
        }
        int roz = abs(wiecej - mniej);
        cout << "   " << T[i] << ": wiekszych=" << wiecej
             << ", mniejszych=" << mniej << ", roznica=" << roz << endl;
        if (roz < minRoz) { minRoz = roz; bestIdx.clear(); bestIdx.push_back(i); }
        else if (roz == minRoz) bestIdx.push_back(i);
    }

    cout << endl << "b) Najbardziej zrownowazone: ";
    for (int k = 0; k < (int)bestIdx.size(); k++) {
        if (k > 0) cout << "   Rowniez: ";
        cout << T[bestIdx[k]] << " (roznica=" << minRoz << ")" << endl;
    }
    return 0;
}
```

**Wyjasnienie**: Podwojna petla O(n^2): dla kazdego elementu zliczamy ile jest od niego wiekszych i mniejszych. Najlepszy to ten z minimalna roznica tych zliczen.

Weryfikacja:
Posortowane: 8,11,15,19,27,33,36,42,44,50
- 33: wiekszych 4 (36,42,44,50), mniejszych 5 (8,11,15,19,27) -> roznica=1
- 27: wiekszych 5 (33,36,42,44,50), mniejszych 4 (8,11,15,19) -> roznica=1
</details>

<details>
<summary>Typowe bledy</summary>

- **Porownanie elementu z samym soba**: Bez `if (j == i) continue` element jest jednoczesnie wiekszy i mniejszy od siebie. CKE: -1 pkt
- **Zapomnienie o elementach rownych**: Elementy rowne nie sa ani wieksze, ani mniejsze. W tym zestawie nie ma duplikatow, ale generalnie trzeba o tym pamietac.

</details>

---

### Cwiczenie 11.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2 (min/max z wieloma kryteriami)
**Tagi**: `struct` `wielokryterialne` `sortowanie` `wczytywanie-pliku`

W pliku `produkty.txt` znajduje sie 10 rekordow w formacie: nazwa kategoria cena ilosc (oddzielone spacjami). Napisz program, ktory:
a) Znajdzie najdrozszy produkt w kazdej kategorii.
b) Znajdzie produkt z najwieksza laczna wartoscia (cena * ilosc).
c) Znajdzie kategorie z najwieksza laczna wartoscia.

**Dane** (`produkty.txt`):
```
Laptop Elektronika 3500 5
Mysz Elektronika 50 100
Monitor Elektronika 1200 15
Biurko Meble 800 8
Krzeslo Meble 450 20
Szafa Meble 1500 3
Dlugopis Biuro 5 500
Papier Biuro 20 200
Toner Biuro 150 30
Segregator Biuro 15 100
```

**Oczekiwany wynik**:
```
a) Najdrozszy w kategorii:
   Biuro: Toner (150)
   Elektronika: Laptop (3500)
   Meble: Szafa (1500)

b) Produkt z max wartoscia: Laptop (3500 * 5 = 17500)

c) Kategoria z max wartoscia: Elektronika (35500)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy mapy: `maxCena[kat]`, `maxProd[kat]`, `sumWart[kat]`.
2. **Podejscie**: Dla kazdego rekordu: aktualizuj max ceny w kategorii, doloz wartosc do sumy kategorii.
3. **Kluczowy krok**: Wartosc produktu = cena * ilosc. Wartosc kategorii = suma wartosci wszystkich produktow w niej.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <map>
using namespace std;

int main() {
    ifstream plik("produkty.txt");
    string nazwa, kat;
    int cena, ilosc;

    map<string, int> maxCena;
    map<string, string> maxProd;
    map<string, int> sumWart;
    string bestProd; int bestWart = 0;

    while (plik >> nazwa >> kat >> cena >> ilosc) {
        int wart = cena * ilosc;
        if (maxCena.find(kat) == maxCena.end() || cena > maxCena[kat]) {
            maxCena[kat] = cena;
            maxProd[kat] = nazwa;
        }
        sumWart[kat] += wart;
        if (wart > bestWart) { bestWart = wart; bestProd = nazwa; }
    }

    // a)
    cout << "a) Najdrozszy w kategorii:" << endl;
    for (auto &p : maxProd)
        cout << "   " << p.first << ": " << p.second
             << " (" << maxCena[p.first] << ")" << endl;

    // b)
    cout << endl << "b) Produkt z max wartoscia: " << bestProd
         << " (" << bestWart << ")" << endl;

    // c)
    string bestKat; int bestSum = 0;
    for (auto &p : sumWart)
        if (p.second > bestSum) { bestSum = p.second; bestKat = p.first; }
    cout << endl << "c) Kategoria z max wartoscia: " << bestKat
         << " (" << bestSum << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Trzy mapy sledza rozne aspekty danych. `maxCena` i `maxProd` — najdrozszy per kategoria. `sumWart` — laczna wartosc per kategoria. Globalne max po wartosci sledzone na biezaco.

Weryfikacja:
- Elektronika: Laptop 3500*5=17500, Mysz 50*100=5000, Monitor 1200*15=18000. Suma=40500. Max cena: Laptop 3500.
  Korekta: 17500+5000+18000 = 40500. Nie 35500.

Korekta oczekiwanego wyniku:
```
c) Kategoria z max wartoscia: Elektronika (40500)
```

Pelna weryfikacja:
- Elektronika: 17500+5000+18000 = 40500
- Meble: 6400+9000+4500 = 19900
- Biuro: 2500+4000+4500+1500 = 12500
Max: Elektronika (40500)
Max produkt: Monitor (18000), nie Laptop (17500).

Korekta:
```
b) Produkt z max wartoscia: Monitor (1200 * 15 = 18000)
```
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomylenie "najdrozszy" z "najwieksza wartosc"**: Najdrozszy = max cena jednostkowa. Najwieksza wartosc = max(cena*ilosc). CKE: -2 pkt
- **Brak `find()` przy inicjalizacji mapy**: Dla nowej kategorii `maxCena[kat]` wynosi 0, wiec sprawdzenie `cena > maxCena[kat]` moze nie zadzialac dla ceny 0. CKE: -1 pkt

</details>

---

### Cwiczenie 11.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2 + 2024 zad. 4 (zlozony min/max)
**Tagi**: `min-max` `spojny-fragment` `podwojna-petla` `sortowanie` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 15 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory:
a) Znajdzie najdluzszy spojny podciag scisle rosnacy, a nastepnie w tym podciagu znajdzie mediane.
b) Znajdzie pare (niekoniecznie sasiednich) elementow T[i] i T[j] (i < j) taka, ze T[j] - T[i] jest maksymalne (najwiekszy zysk "kup tanio, sprzedaj drogo").
c) Poda ile elementow ciagu jest jednoczesnie wiekszych od swojego lewego sasiada i mniejszych od prawego (lokalne "doliny").

**Dane** (`dane.txt`):
```
30
10
20
40
60
50
55
80
75
5
15
25
35
45
70
```

**Oczekiwany wynik**:
```
a) Najdluzszy ciag rosnacy:
   Pozycje 2-5: [10, 20, 40, 60], dlugosc 4
   Mediana: 30.00

b) Max zysk: kup T[10]=5, sprzedaj T[15]=70, zysk=65

c) Lokalne doliny (T[i-1] > T[i] < T[i+1]):
   Poz.2: 30 > 10 < 20 -> TAK
   Poz.6: 60 > 50 < 55 -> TAK
   Poz.10: 75 > 5 < 15 -> TAK
   Ilosc: 3
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: (a) Wzorzec current/max dla ciagu rosnacego, potem sort+mediana fragmentu. (b) Sledzenie minimum dotychczasowego i max roznicy. (c) Sprawdz warunek na sasiady.
2. **Podejscie**: Dla (b) iteruj raz: sledzac `minDotychczasowy`, oblicz `T[j] - minDotychczasowy` i aktualizuj max.
3. **Kluczowy krok**: (b) to klasyczny problem "Best Time to Buy and Sell Stock" — liniowe O(n) zamiast O(n^2).

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    // a) Najdluzszy ciag rosnacy
    int maxDl = 1, maxStart = 0, curDl = 1, curStart = 0;
    for (int i = 1; i < n; i++) {
        if (T[i] > T[i - 1]) {
            curDl++;
        } else {
            if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }
            curDl = 1; curStart = i;
        }
    }
    if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }

    cout << "a) Najdluzszy ciag rosnacy:" << endl;
    cout << "   Pozycje " << maxStart + 1 << "-" << maxStart + maxDl
         << ": [";
    vector<int> fragment(T.begin() + maxStart, T.begin() + maxStart + maxDl);
    for (int i = 0; i < maxDl; i++) {
        if (i > 0) cout << ", ";
        cout << fragment[i];
    }
    cout << "], dlugosc " << maxDl << endl;

    sort(fragment.begin(), fragment.end());
    double mediana;
    if (maxDl % 2 == 0)
        mediana = (fragment[maxDl / 2 - 1] + fragment[maxDl / 2]) / 2.0;
    else
        mediana = fragment[maxDl / 2];
    cout << "   Mediana: " << fixed << setprecision(2) << mediana << endl;

    // b) Max zysk (buy low, sell high)
    int minSoFar = T[0], minIdx = 0;
    int maxProfit = 0, buyIdx = 0, sellIdx = 0;
    for (int j = 1; j < n; j++) {
        int profit = T[j] - minSoFar;
        if (profit > maxProfit) {
            maxProfit = profit;
            buyIdx = minIdx;
            sellIdx = j;
        }
        if (T[j] < minSoFar) {
            minSoFar = T[j];
            minIdx = j;
        }
    }
    cout << endl << "b) Max zysk: kup T[" << buyIdx + 1 << "]=" << T[buyIdx]
         << ", sprzedaj T[" << sellIdx + 1 << "]=" << T[sellIdx]
         << ", zysk=" << maxProfit << endl;

    // c) Lokalne doliny
    cout << endl << "c) Lokalne doliny (T[i-1] > T[i] < T[i+1]):" << endl;
    int ileDolin = 0;
    for (int i = 1; i < n - 1; i++) {
        if (T[i - 1] > T[i] && T[i] < T[i + 1]) {
            cout << "   Poz." << i + 1 << ": " << T[i - 1] << " > "
                 << T[i] << " < " << T[i + 1] << " -> TAK" << endl;
            ileDolin++;
        }
    }
    cout << "   Ilosc: " << ileDolin << endl;
    return 0;
}
```

**Wyjasnienie**: (a) Wzorzec current/max + mediana posortowanego fragmentu. (b) Liniowy algorytm "buy and sell": sledzac minimum dotychczasowe, w kazdym punkcie sprawdzamy potencjalny zysk. (c) Prosta iteracja ze sprawdzeniem obu sasiadow.

Weryfikacja:
- Ciag: 30,10,20,40,60,50,55,80,75,5,15,25,35,45,70
- Ciagi rosnace: (10,20,40,60) dl.4 od poz.2; (50,55,80) dl.3 od poz.6; (5,15,25,35,45,70) dl.6 od poz.10

Korekta: najdluzszy to (5,15,25,35,45,70) dl.6 od poz.10, nie dl.4.
Mediana: posortowane juz sa: 5,15,25,35,45,70. Mediana = (25+35)/2 = 30.00

Korekta oczekiwanego wyniku:
```
a) Pozycje 10-15: [5, 15, 25, 35, 45, 70], dlugosc 6
   Mediana: 30.00
b) Max zysk: kup T[10]=5, sprzedaj T[15]=70, zysk=65
```
</details>

<details>
<summary>Typowe bledy</summary>

- **O(n^2) w punkcie (b)**: Podwojna petla dziala, ale na duzych danych timeout. Liniowy algorytm z minSoFar jest lepszy. CKE: 0 pkt (poprawne, ale wolne)
- **Zapomnienie o aktualizacji maxDl na koncu petli (a)**: Jesli najdluzszy ciag konczy sie na ostatnim elemencie, trzeba sprawdzic po petli. CKE: -2 pkt
- **Dolina vs szczyt**: Dolina to `T[i-1] > T[i] < T[i+1]`, szczyt to `T[i-1] < T[i] > T[i+1]`. Pomylenie powoduje odwrotny wynik. CKE: -1 pkt

</details>

---

## Samoocena

| Poziom | Opis | Kryteria |
|--------|------|----------|
| Podstawowy | Umiem znalezc min/max w tablicy | Cwiczenia 1-2, 6 bez pomocy |
| Dobry | Radze sobie z sortowaniem i k-tym elementem | Cwiczenia 3, 7 bez pomocy |
| Bardzo dobry | Umiem przetwarzac tablice 2D i pary elementow | Cwiczenia 4-5, 8 bez pomocy |
| Doskonaly | Radze sobie ze zlozonym min/max i optymalizacja | Cwiczenia 9-10 bez pomocy |

**Co dalej?**
- Jesli masz problem z sortowaniem -> patrz `cheatsheet_cpp.md` sekcja "Sortowanie"
- Jesli chcesz cwiczycwyszukiwanie w ciagach -> patrz `12_sekwencje.md`
- Jesli chcesz cwiczycoperacje na rekordach -> patrz `09_zlozone.md`
