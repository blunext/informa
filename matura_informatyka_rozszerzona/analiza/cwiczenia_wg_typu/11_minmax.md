# 11. Wyszukiwanie min/max

Typ zadania: **minmax**
Czestotliwosc: 5/11 lat | Laczna punktacja: 17 pkt
Kategoria: IMPLEMENTACJA

---

### Cwiczenie 11.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 zad. 4.3

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

---

### Cwiczenie 11.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4.2

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

---

### Cwiczenie 11.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4.2 (k-ta liczba pierwsza)

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

---

### Cwiczenie 11.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2015 zad. 4.3 (min/max w wierszu tablicy)

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

---

### Cwiczenie 11.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2.3 + 2024 zad. 4

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

**Wyjasnienie**: (a) Podwojna petla O(n^2) sprawdzajaca wszystkie pary — mozna tez znalezc globalne min i max. (b) Petla po parach sasiadow. (c) Podwojna petla — dla kazdego poczatku rozszerzamy fragment aktualizujac min/max, dopoki roznica <= 10 (uwaga: `break` dziala poprawnie gdy fragment sie laczy, ale nie zawsze — pelna wersja bez break jest bezpieczniejsza).

Weryfikacja:
a) Globalne min=3 (poz.13), max=80 (poz.14). Para (13,14): |3-80|=77.
b) Rożnice sasiadow: |25-18|=7, |18-22|=4, |22-20|=2, |20-19|=1, |19-24|=5, |24-21|=3, |21-50|=29, |50-45|=5, |45-48|=3, |48-47|=1, |47-46|=1, |46-3|=43, |3-80|=77, |80-5|=75
   Min=1 (pozycje 4,5 lub 10,11 lub 11,12). Pierwsza: pozycje 4,5 (20,19).
c) Fragment 1-7: [25,18,22,20,19,24,21], min=18, max=25, diff=7. Element 8 (50) sprawia ze diff=32. Dlugosc 7.
   Fragment 8-12: [50,45,48,47,46], min=45, max=50, diff=5. Dlugosc 5.
   Najdluzszy: 7 (pozycje 1-7).
</details>

---
