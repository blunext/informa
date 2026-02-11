# 14. Obliczenia geometryczne

Typ zadania: **geometryczne**
Czestotliwosc: 1/11 lat | Laczna punktacja: 4 pkt
Kategoria: IMPLEMENTACJA

## Umiejetnosci cwiczone w tym zestawie

`odleglosc-euklidesowa` `sqrt` `wczytywanie-pliku` `vector-par` `punkt-w-prostokacie` `najblizsze-punkty` `srodek-odcinka` `set` `pole-trojkata` `punkt-w-trojkacie` `struct` `sortowanie` `kat-miedzy-wektorami` `wielokat` `otoczka-wypukla` `obwod` `collinear`

---

### Cwiczenie 14.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3 (Dron)
**Tagi**: `odleglosc-euklidesowa` `sqrt` `wczytywanie-pliku` `vector-par`

W pliku `punkty.txt` znajduje sie 6 punktow na plaszczyznie (kazdy w osobnym wierszu w formacie: x y, wspolrzedne calkowite). Punkty tworzą sciezke (lamanę) od pierwszego do ostatniego. Napisz program, ktory obliczy odleglosc miedzy kazdymi dwoma kolejnymi punktami i poda laczna dlugosc sciezki.

**Dane** (`punkty.txt`):
```
0 0
3 4
6 0
6 8
10 8
10 0
```

**Oczekiwany wynik**:
```
Odleglosci miedzy kolejnymi punktami:
(0,0) -> (3,4): 5.00
(3,4) -> (6,0): 5.00
(6,0) -> (6,8): 8.00
(6,8) -> (10,8): 4.00
(10,8) -> (10,0): 8.00
Laczna dlugosc sciezki: 30.00
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wzor na odleglosc euklidesowa: `sqrt((x2-x1)^2 + (y2-y1)^2)`.
2. **Podejscie**: Wczytaj punkty do `vector<pair<int,int>>`, iteruj po parach sasiadow.
3. **Kluczowy krok**: Uzyj `#include <cmath>` dla `sqrt`. Sumuj odleglosci.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <cmath>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("punkty.txt");
    vector<pair<int,int>> P;
    int x, y;
    while (plik >> x >> y) P.push_back({x, y});

    cout << "Odleglosci miedzy kolejnymi punktami:" << endl;
    cout << fixed << setprecision(2);
    double suma = 0;
    for (int i = 0; i < (int)P.size() - 1; i++) {
        double dx = P[i+1].first - P[i].first;
        double dy = P[i+1].second - P[i].second;
        double d = sqrt(dx*dx + dy*dy);
        cout << "(" << P[i].first << "," << P[i].second << ") -> ("
             << P[i+1].first << "," << P[i+1].second << "): " << d << endl;
        suma += d;
    }
    cout << "Laczna dlugosc sciezki: " << suma << endl;
    return 0;
}
```

**Wyjasnienie**: Odleglosc euklidesowa miedzy punktami (x1,y1) i (x2,y2) to sqrt((x2-x1)^2 + (y2-y1)^2). Sumujemy odleglosci kolejnych par.

Weryfikacja:
- (0,0)->(3,4): sqrt(9+16) = sqrt(25) = 5.00
- (3,4)->(6,0): sqrt(9+16) = sqrt(25) = 5.00
- (6,0)->(6,8): sqrt(0+64) = 8.00
- (6,8)->(10,8): sqrt(16+0) = 4.00
- (10,8)->(10,0): sqrt(0+64) = 8.00
Suma: 5+5+8+4+8 = 30.00
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak `#include <cmath>`**: Funkcja `sqrt` wymaga tego naglowka. CKE: blad kompilacji
- **Uzycie int zamiast double dla odleglosci**: `sqrt` zwraca double — przypisanie do int obcina czesc ulamkowa. CKE: -1 pkt
- **Petla do `P.size()` zamiast `P.size()-1`**: Dostep do P[i+1] poza zakresem. CKE: crash

</details>

---

### Cwiczenie 14.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3.2a
**Tagi**: `punkt-w-prostokacie` `wczytywanie-pliku` `filtrowanie`

W pliku `punkty.txt` znajduje sie 10 punktow (x, y) o wspolrzednych calkowitych. Dany jest prostokat o wierzcholkach (1, 1) i (99, 99). Napisz program, ktory zliczy ile punktow lezy scisle wewnatrz prostokata (brzegi nie licza sie).

**Dane** (`punkty.txt`):
```
50 50
1 1
0 45
99 99
100 50
30 70
1 50
50 99
15 85
60 0
```

**Oczekiwany wynik**:
```
Punkty wewnatrz prostokata (1,1)-(99,99) (bez brzegu):
(50,50) - wewnatrz
(30,70) - wewnatrz
(15,85) - wewnatrz
Ilosc: 3
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: "Scisle wewnatrz" = ostro: `x > 1 && x < 99 && y > 1 && y < 99`.
2. **Podejscie**: Wczytuj punkty w petli i sprawdzaj warunek.
3. **Kluczowy krok**: Brzegi nie licza sie — uzyj `>` i `<`, nie `>=` i `<=`.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ifstream plik("punkty.txt");
    int x, y;
    cout << "Punkty wewnatrz prostokata (1,1)-(99,99) (bez brzegu):" << endl;
    int ile = 0;
    while (plik >> x >> y) {
        if (x > 1 && x < 99 && y > 1 && y < 99) {
            cout << "(" << x << "," << y << ") - wewnatrz" << endl;
            ile++;
        }
    }
    cout << "Ilosc: " << ile << endl;
    return 0;
}
```

**Wyjasnienie**: Punkt lezy scisle wewnatrz prostokata gdy 1 < x < 99 i 1 < y < 99 (ostro, bez rownosci).

Weryfikacja:
- (50,50): 1<50<99, 1<50<99 -> TAK
- (1,1): nie, na brzegu
- (0,45): nie, x=0 < 1
- (99,99): nie, na brzegu
- (100,50): nie, x=100 > 99
- (30,70): 1<30<99, 1<70<99 -> TAK
- (1,50): nie, x=1 na brzegu
- (50,99): nie, y=99 na brzegu
- (15,85): 1<15<99, 1<85<99 -> TAK
- (60,0): nie, y=0 < 1
</details>

<details>
<summary>Typowe bledy</summary>

- **`>=` zamiast `>`**: Punkt na brzegu nie jest "scisle wewnatrz". CKE: -1 pkt
- **Pomylenie wspolrzednych prostokata**: (1,1)-(99,99) to lewy dolny i prawy gorny rog. CKE: -1 pkt

</details>

---

### Cwiczenie 14.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2018 zad. 2 (Krajobraz)
**Tagi**: `najblizsze-punkty` `odleglosc-euklidesowa` `podwojna-petla` `wczytywanie-pliku`

W pliku `punkty.txt` znajduje sie 8 punktow (x, y) o wspolrzednych calkowitych. Napisz program, ktory:
a) Znajdzie pare punktow o najmniejszej odleglosci euklidesowej.
b) Znajdzie pare punktow o najwiekszej odleglosci.

**Dane** (`punkty.txt`):
```
1 2
4 6
7 1
3 5
10 3
8 8
2 9
6 4
```

**Oczekiwany wynik**:
```
a) Najblizsze punkty: (3,5) i (4,6), odleglosc: 1.41

b) Najdalsze punkty: (2,9) i (10,3), odleglosc: 10.00
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Podwojna petla O(n^2) po wszystkich parach.
2. **Podejscie**: Sledzac min i max odleglosc, zapamietaj indeksy odpowiednich par.
3. **Kluczowy krok**: Mozna porownywac kwadraty odleglosci (bez sqrt) dla szybkosci, a sqrt obliczyc tylko na koncu.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <cmath>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("punkty.txt");
    vector<pair<int,int>> P;
    int x, y;
    while (plik >> x >> y) P.push_back({x, y});
    int n = P.size();

    double minDist = 1e18, maxDist = 0;
    int mi1, mi2, ma1, ma2;

    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            double dx = P[i].first - P[j].first;
            double dy = P[i].second - P[j].second;
            double d = sqrt(dx*dx + dy*dy);
            if (d < minDist) { minDist = d; mi1 = i; mi2 = j; }
            if (d > maxDist) { maxDist = d; ma1 = i; ma2 = j; }
        }
    }

    cout << fixed << setprecision(2);
    cout << "a) Najblizsze punkty: (" << P[mi1].first << "," << P[mi1].second
         << ") i (" << P[mi2].first << "," << P[mi2].second
         << "), odleglosc: " << minDist << endl;
    cout << endl << "b) Najdalsze punkty: (" << P[ma1].first << "," << P[ma1].second
         << ") i (" << P[ma2].first << "," << P[ma2].second
         << "), odleglosc: " << maxDist << endl;
    return 0;
}
```

**Wyjasnienie**: Podwojna petla O(n^2) sprawdzajaca odleglosci miedzy wszystkimi parami punktow.

Weryfikacja:
- (3,5)-(4,6): sqrt(1+1) = sqrt(2) = 1.41
- (2,9)-(10,3): sqrt(64+36) = sqrt(100) = 10.00
</details>

<details>
<summary>Typowe bledy</summary>

- **Niezainicjalizowane zmienne mi1, mi2**: Jesli zaden if nie jest spelniony (niemozliwe dla n >= 2), wynik bezsensowny. Dobra praktyka: zainicjalizuj. CKE: potencjalny UB
- **Porownywanie odleglosci z `==`**: Odleglosci sa double — unikaj `==` przez epsilony. CKE: 0 pkt (tu nie dotyczy)

</details>

---

### Cwiczenie 14.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3.2b (srodek odcinka)
**Tagi**: `srodek-odcinka` `set` `podwojna-petla` `wczytywanie-pliku`

W pliku `punkty.txt` znajduje sie 8 punktow o wspolrzednych calkowitych. Napisz program, ktory znajdzie wszystkie trojki punktow (A, B, C) takie, ze jeden z nich jest srodkiem odcinka laczacego dwa pozostale. Srodek odcinka AB to punkt ((xA+xB)/2, (yA+yB)/2) — obie wspolrzedne musza byc calkowite.

**Dane** (`punkty.txt`):
```
0 0
2 4
4 8
1 2
6 2
3 1
8 4
5 3
```

**Oczekiwany wynik**:
```
Trojki gdzie jeden punkt jest srodkiem odcinka dwoch pozostalych:
(0,0) i (4,8) -> srodek (2,4) - jest w zbiorze
(0,0) i (2,4) -> srodek (1,2) - jest w zbiorze
(0,0) i (6,2) -> srodek (3,1) - jest w zbiorze
Ilosc: 3
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wczytaj punkty i wstaw do `set<pair<int,int>>` dla szybkiego wyszukiwania.
2. **Podejscie**: Podwojna petla po parach. Srodek ma calkowite wspolrzedne gdy `(xA+xB) % 2 == 0` i `(yA+yB) % 2 == 0`.
3. **Kluczowy krok**: Sprawdz `zbior.count(mid) && mid != P[i] && mid != P[j]` — srodek musi istniec w zbiorze i byc inny od koncow.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <set>
using namespace std;

int main() {
    ifstream plik("punkty.txt");
    vector<pair<int,int>> P;
    set<pair<int,int>> zbior;
    int x, y;
    while (plik >> x >> y) {
        P.push_back({x, y});
        zbior.insert({x, y});
    }
    int n = P.size();

    cout << "Trojki gdzie jeden punkt jest srodkiem odcinka dwoch pozostalych:" << endl;
    int ile = 0;
    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            int sx = P[i].first + P[j].first;
            int sy = P[i].second + P[j].second;
            // Srodek ma calkowite wspolrzedne gdy sumy sa parzyste
            if (sx % 2 == 0 && sy % 2 == 0) {
                pair<int,int> mid = {sx / 2, sy / 2};
                if (zbior.count(mid) && mid != P[i] && mid != P[j]) {
                    cout << "(" << P[i].first << "," << P[i].second << ") i ("
                         << P[j].first << "," << P[j].second << ") -> srodek ("
                         << mid.first << "," << mid.second << ") - jest w zbiorze" << endl;
                    ile++;
                }
            }
        }
    }
    cout << "Ilosc: " << ile << endl;
    return 0;
}
```

**Wyjasnienie**: Dla kazdej pary punktow obliczamy srodek odcinka. Jesli obie wspolrzedne srodka sa calkowite i srodek wystepuje wsrod danych punktow, mamy trojke. Uzywamy `set` do szybkiego sprawdzania przynaleznosci.

Weryfikacja:
- (0,0)+(4,8) = (4,8)/2 = (2,4) -> w zbiorze? TAK
- (0,0)+(2,4) = (2,4)/2 = (1,2) -> w zbiorze? TAK
- (0,0)+(6,2) = (6,2)/2 = (3,1) -> w zbiorze? TAK
- Inne pary: srodki nie sa w zbiorze.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak sprawdzenia parzystosci sum**: Dzielenie `5/2 = 2` w C++ (calkowite!) — wynik bedzie zly. CKE: -2 pkt
- **Brak `mid != P[i] && mid != P[j]`**: Srodek (2,4) miedzy (2,4) i (2,4) to ten sam punkt. CKE: -1 pkt
- **Uzycie wektora zamiast seta**: `find` w wektorze to O(n), w secie O(log n). Na 8 punktach nie ma roznicy, ale na duzych danych tak.

</details>

---

### Cwiczenie 14.5 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3
**Tagi**: `pole-trojkata` `punkt-w-trojkacie` `struct` `sortowanie` `wczytywanie-pliku`

W pliku `punkty.txt` znajduje sie 6 punktow o wspolrzednych calkowitych. Napisz program, ktory:
a) Obliczy pole trojkata dla kazdej mozliwej trojki punktow (uzyj wzoru z wspolrzednych: P = |x1(y2-y3) + x2(y3-y1) + x3(y1-y2)| / 2).
b) Znajdzie trojkat o najwiekszym polu.
c) Sprawdzi, czy punkt D=(4,3) lezy wewnatrz trojkata o najwiekszym polu (metoda: punkt lezy wewnatrz trojkata ABC jesli suma pol trojkatow ABD + ACD + BCD = pole ABC).

**Dane** (`punkty.txt`):
```
0 0
8 0
0 6
4 3
7 5
2 8
```

**Oczekiwany wynik**:
```
a) Najwieksze trojkaty (top 3):
   (0,0)-(8,0)-(2,8): pole = 32.00
   (0,0)-(8,0)-(0,6): pole = 24.00
   (0,0)-(7,5)-(2,8): pole = 23.00

b) Trojkat o max polu: (0,0)-(8,0)-(2,8), pole = 32.00

c) Czy D=(4,3) lezy wewnatrz trojkata (0,0)-(8,0)-(2,8)?
   Pole ABD = 12.00, pole ACD = 13.00, pole BCD = 7.00
   Suma = 32.00, pole ABC = 32.00
   Suma == pole -> D lezy wewnatrz: TAK
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wzor na pole: `P = |x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)| / 2.0`.
2. **Podejscie**: Potrojna petla po trojkach. Sortuj po polu malejaco.
3. **Kluczowy krok**: Punkt wewnatrz trojkata: suma pol 3 pod-trojkatow = pole calego trojkata (z dokladnoscia do epsilon).

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <cmath>
#include <iomanip>
#include <algorithm>
using namespace std;

struct Pt { int x, y; };

double pole(Pt a, Pt b, Pt c) {
    return abs(a.x * (b.y - c.y) + b.x * (c.y - a.y) + c.x * (a.y - b.y)) / 2.0;
}

int main() {
    ifstream plik("punkty.txt");
    vector<Pt> P;
    int x, y;
    while (plik >> x >> y) P.push_back({x, y});
    int n = P.size();

    // a) i b) Oblicz pola wszystkich trojkatow
    struct Tri { int i, j, k; double p; };
    vector<Tri> trojkaty;
    for (int i = 0; i < n; i++)
        for (int j = i+1; j < n; j++)
            for (int k = j+1; k < n; k++) {
                double p = pole(P[i], P[j], P[k]);
                trojkaty.push_back({i, j, k, p});
            }
    sort(trojkaty.begin(), trojkaty.end(), [](Tri &a, Tri &b) { return a.p > b.p; });

    cout << fixed << setprecision(2);
    cout << "a) Najwieksze trojkaty (top 3):" << endl;
    for (int t = 0; t < 3 && t < (int)trojkaty.size(); t++) {
        auto &tr = trojkaty[t];
        cout << "   (" << P[tr.i].x << "," << P[tr.i].y << ")-("
             << P[tr.j].x << "," << P[tr.j].y << ")-("
             << P[tr.k].x << "," << P[tr.k].y << "): pole = "
             << tr.p << endl;
    }

    auto &best = trojkaty[0];
    cout << endl << "b) Trojkat o max polu: ("
         << P[best.i].x << "," << P[best.i].y << ")-("
         << P[best.j].x << "," << P[best.j].y << ")-("
         << P[best.k].x << "," << P[best.k].y << "), pole = "
         << best.p << endl;

    // c) Punkt D wewnatrz?
    Pt D = {4, 3};
    Pt A = P[best.i], B = P[best.j], C = P[best.k];
    double pABC = pole(A, B, C);
    double pABD = pole(A, B, D);
    double pACD = pole(A, C, D);
    double pBCD = pole(B, C, D);
    double sumaP = pABD + pACD + pBCD;

    cout << endl << "c) Czy D=(" << D.x << "," << D.y << ") lezy wewnatrz trojkata ("
         << A.x << "," << A.y << ")-(" << B.x << "," << B.y << ")-("
         << C.x << "," << C.y << ")?" << endl;
    cout << "   Pole ABD = " << pABD << ", pole ACD = " << pACD
         << ", pole BCD = " << pBCD << endl;
    cout << "   Suma = " << sumaP << ", pole ABC = " << pABC << endl;
    if (abs(sumaP - pABC) < 0.001)
        cout << "   Suma == pole -> D lezy wewnatrz: TAK" << endl;
    else
        cout << "   Suma != pole -> D lezy na zewnatrz: NIE" << endl;
    return 0;
}
```

**Wyjasnienie**: Pole trojkata z wspolrzednych: P = |x1(y2-y3) + x2(y3-y1) + x3(y1-y2)| / 2. Punkt D lezy wewnatrz ABC wtedy i tylko wtedy, gdy pole(ABD) + pole(ACD) + pole(BCD) = pole(ABC).

Weryfikacja:
- Trojkat (0,0)-(8,0)-(2,8): |0*(0-8) + 8*(8-0) + 2*(0-0)| / 2 = |0+64+0| / 2 = 32.00
- D=(4,3):
  - pole ABD: |0*(0-3)+8*(3-0)+4*(0-0)|/2 = 24/2 = 12.00
  - pole ACD: |0*(8-3)+2*(3-0)+4*(0-8)|/2 = |0+6-32|/2 = 26/2 = 13.00
  - pole BCD: |8*(8-3)+2*(3-0)+4*(0-8)|/2 = |40+6-32|/2 = 14/2 = 7.00
  - Suma: 12+13+7 = 32.00 = pole ABC -> D wewnatrz: TAK
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak `abs()` we wzorze na pole**: Iloczyn wektorowy moze byc ujemny — bez wartosci bezwzglednej pole bedzie ujemne. CKE: -2 pkt
- **Porownanie double z `==`**: `sumaP == pABC` moze nie zadzialac przez bledy zmiennoprzecinkowe. Uzyj `abs(a-b) < epsilon`. CKE: -1 pkt
- **Zapomnienie `/2.0`**: Pole jest polowa iloczynu wektorowego. CKE: -1 pkt

</details>

---

### Cwiczenie 14.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3 (odleglosc od poczatku)
**Tagi**: `odleglosc-euklidesowa` `sortowanie` `wczytywanie-pliku`

W pliku `punkty.txt` znajduje sie 8 punktow (x, y) o wspolrzednych calkowitych. Napisz program, ktory:
a) Obliczy odleglosc kazdego punktu od poczatku ukladu wspolrzednych (0, 0).
b) Wypisze punkty posortowane rosnaco wedlug odleglosci od (0, 0).

**Dane** (`punkty.txt`):
```
3 4
1 1
5 0
0 7
6 8
2 3
4 4
8 1
```

**Oczekiwany wynik**:
```
a) Odleglosci od (0,0):
   (3,4): 5.00
   (1,1): 1.41
   (5,0): 5.00
   (0,7): 7.00
   (6,8): 10.00
   (2,3): 3.61
   (4,4): 5.66
   (8,1): 8.06

b) Posortowane:
   (1,1): 1.41
   (2,3): 3.61
   (3,4): 5.00
   (5,0): 5.00
   (4,4): 5.66
   (0,7): 7.00
   (8,1): 8.06
   (6,8): 10.00
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Odleglosc od (0,0) to `sqrt(x*x + y*y)`.
2. **Podejscie**: Wczytaj do vectora, oblicz odleglosci, posortuj z lambda po odleglosci.
3. **Kluczowy krok**: `sort` z komparatorem `sqrt(a.x^2+a.y^2) < sqrt(b.x^2+b.y^2)`. Mozna porownywac kwadraty odleglosci (bez sqrt).

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
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("punkty.txt");
    vector<pair<int,int>> P;
    int x, y;
    while (plik >> x >> y) P.push_back({x, y});

    cout << fixed << setprecision(2);
    cout << "a) Odleglosci od (0,0):" << endl;
    for (auto &p : P) {
        double d = sqrt(p.first * p.first + p.second * p.second);
        cout << "   (" << p.first << "," << p.second << "): " << d << endl;
    }

    sort(P.begin(), P.end(), [](auto &a, auto &b) {
        return a.first * a.first + a.second * a.second
             < b.first * b.first + b.second * b.second;
    });

    cout << endl << "b) Posortowane:" << endl;
    for (auto &p : P) {
        double d = sqrt(p.first * p.first + p.second * p.second);
        cout << "   (" << p.first << "," << p.second << "): " << d << endl;
    }
    return 0;
}
```

**Wyjasnienie**: Odleglosc od poczatku to sqrt(x^2 + y^2). Sortowanie po kwadratach odleglosci (bez sqrt) jest szybsze i unika bledow zaokraglen.

Weryfikacja:
- (1,1): sqrt(2) = 1.41
- (2,3): sqrt(13) = 3.61
- (3,4): sqrt(25) = 5.00
- (5,0): sqrt(25) = 5.00
- (4,4): sqrt(32) = 5.66
- (0,7): sqrt(49) = 7.00
- (8,1): sqrt(65) = 8.06
- (6,8): sqrt(100) = 10.00
</details>

<details>
<summary>Typowe bledy</summary>

- **Sortowanie z sqrt zamiast kwadratow**: Dziala, ale jest wolniejsze. Nie jest bledem.
- **Brak stabilnosci sortowania**: Punkty o tej samej odleglosci (3,4) i (5,0) moga byc w dowolnej kolejnosci — to OK.

</details>

---

### Cwiczenie 14.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3 (wspolliniowosc)
**Tagi**: `collinear` `pole-trojkata` `wczytywanie-pliku`

W pliku `punkty.txt` znajduje sie 8 punktow (x, y) o wspolrzednych calkowitych. Napisz program, ktory:
a) Sprawdzi, ktore trojki punktow sa wspollinowe (leza na jednej prostej).
b) Znajdzie najdluzsza grupe punktow wspollinowych.

**Dane** (`punkty.txt`):
```
0 0
1 2
2 4
3 6
5 1
7 3
9 5
4 8
```

**Oczekiwany wynik**:
```
a) Trojki wspollinowe (pole trojkata = 0):
   (0,0)-(1,2)-(2,4): pole=0 -> wspollinowe
   (0,0)-(1,2)-(3,6): pole=0 -> wspollinowe
   (0,0)-(2,4)-(3,6): pole=0 -> wspollinowe
   (1,2)-(2,4)-(3,6): pole=0 -> wspollinowe
   (5,1)-(7,3)-(9,5): pole=0 -> wspollinowe

b) Najdluzsza grupa wspollinowa: 4 punkty
   Punkty: (0,0), (1,2), (2,4), (3,6)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy punkty sa wspollinowe gdy pole trojkata = 0. Wzor: `x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2) == 0`.
2. **Podejscie**: Dla kazdej pary (A, B) policz ile innych punktow lezy na prostej AB (pole = 0).
3. **Kluczowy krok**: Najdluzsza grupa wspollinowa = max(ile punktow na jednej prostej).

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

struct Pt { int x, y; };

int iloczyn(Pt a, Pt b, Pt c) {
    return a.x * (b.y - c.y) + b.x * (c.y - a.y) + c.x * (a.y - b.y);
}

int main() {
    ifstream plik("punkty.txt");
    vector<Pt> P;
    int x, y;
    while (plik >> x >> y) P.push_back({x, y});
    int n = P.size();

    // a) Trojki wspollinowe
    cout << "a) Trojki wspollinowe (pole trojkata = 0):" << endl;
    for (int i = 0; i < n; i++)
        for (int j = i + 1; j < n; j++)
            for (int k = j + 1; k < n; k++)
                if (iloczyn(P[i], P[j], P[k]) == 0)
                    cout << "   (" << P[i].x << "," << P[i].y << ")-("
                         << P[j].x << "," << P[j].y << ")-("
                         << P[k].x << "," << P[k].y << "): wspollinowe" << endl;

    // b) Najdluzsza grupa wspollinowa
    int maxGrupa = 2;
    int bestI = 0, bestJ = 1;
    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            int cnt = 2;
            for (int k = 0; k < n; k++) {
                if (k == i || k == j) continue;
                if (iloczyn(P[i], P[j], P[k]) == 0) cnt++;
            }
            if (cnt > maxGrupa) { maxGrupa = cnt; bestI = i; bestJ = j; }
        }
    }

    cout << endl << "b) Najdluzsza grupa wspollinowa: " << maxGrupa << " punkty" << endl;
    cout << "   Punkty: ";
    bool first = true;
    for (int k = 0; k < n; k++) {
        if (k == bestI || k == bestJ || iloczyn(P[bestI], P[bestJ], P[k]) == 0) {
            if (!first) cout << ", ";
            cout << "(" << P[k].x << "," << P[k].y << ")";
            first = false;
        }
    }
    cout << endl;
    return 0;
}
```

**Wyjasnienie**: Wspolliniowosc trojki = iloczyn wektorowy == 0 (pole trojkata = 0). Dla najdluzszej grupy: dla kazdej pary (A,B) sprawdzamy ile punktow lezy na prostej AB.

Weryfikacja:
- (0,0),(1,2),(2,4),(3,6): kazda trojka ma pole 0. Rownanie prostej: y = 2x. Sprawdzmy: (5,1) y!=2*5, (7,3) y!=2*7. Ale (5,1),(7,3),(9,5): 5*(3-5)+7*(5-1)+9*(1-3) = -10+28-18 = 0 -> wspollinowe na innej prostej (y = x-4).
- Najdluzsza grupa: 4 punkty na prostej y=2x.
</details>

<details>
<summary>Typowe bledy</summary>

- **Uzycie double zamiast int w iloczynie**: Dla wspolrzednych calkowitych iloczyn wektorowy jest dokladny. Double moze wprowadzic bledy zaokraglen. CKE: -1 pkt
- **Zapomnienie o parze (bestI, bestJ)**: Para definiujaca prosta musi byc zawarta w wyniku. CKE: -1 pkt

</details>

---

### Cwiczenie 14.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3 (obwod wielokata)
**Tagi**: `wielokat` `obwod` `odleglosc-euklidesowa` `wczytywanie-pliku`

W pliku `punkty.txt` znajduje sie 6 punktow tworzacych wielokat wypukly (podane w kolejnosci obejscia). Napisz program, ktory:
a) Obliczy obwod wielokata (suma dlugosci bokow, wlaczajac odcinek zamykajacy od ostatniego do pierwszego punktu).
b) Obliczy pole wielokata (wzor sznurowadla: P = |sum(x_i * y_{i+1} - x_{i+1} * y_i)| / 2).
c) Obliczy srodek ciezkosci (centroid): (srednia x, srednia y).

**Dane** (`punkty.txt`):
```
0 0
6 0
8 3
6 6
2 6
0 3
```

**Oczekiwany wynik**:
```
a) Obwod: 22.65

b) Pole: 36.00

c) Centroid: (3.67, 3.00)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Obwod = suma odleglosci kolejnych par + odleglosc ostatni->pierwszy. Pole = wzor sznurowadla.
2. **Podejscie**: Wzor sznurowadla (Shoelace): `P = |sum(x[i]*y[i+1] - x[i+1]*y[i])| / 2` (indeksy modulo n).
3. **Kluczowy krok**: Centroid = (sum(x)/n, sum(y)/n) — prosta srednia wspolrzednych wierzcholkow.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <cmath>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("punkty.txt");
    vector<pair<int,int>> P;
    int x, y;
    while (plik >> x >> y) P.push_back({x, y});
    int n = P.size();

    // a) Obwod
    double obwod = 0;
    for (int i = 0; i < n; i++) {
        int j = (i + 1) % n;
        double dx = P[j].first - P[i].first;
        double dy = P[j].second - P[i].second;
        obwod += sqrt(dx * dx + dy * dy);
    }
    cout << fixed << setprecision(2);
    cout << "a) Obwod: " << obwod << endl;

    // b) Pole (wzor sznurowadla)
    int sum = 0;
    for (int i = 0; i < n; i++) {
        int j = (i + 1) % n;
        sum += P[i].first * P[j].second - P[j].first * P[i].second;
    }
    double pole = abs(sum) / 2.0;
    cout << endl << "b) Pole: " << pole << endl;

    // c) Centroid
    double cx = 0, cy = 0;
    for (auto &p : P) { cx += p.first; cy += p.second; }
    cx /= n; cy /= n;
    cout << endl << "c) Centroid: (" << cx << ", " << cy << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Obwod to suma odleglosci kolejnych wierzcholkow (z zamknieciem). Pole przez wzor sznurowadla. Centroid to srednia arytmetyczna wspolrzednych.

Weryfikacja:
- Obwod: (0,0)->(6,0): 6. (6,0)->(8,3): sqrt(4+9)=sqrt(13)=3.61. (8,3)->(6,6): sqrt(4+9)=3.61. (6,6)->(2,6): 4. (2,6)->(0,3): sqrt(4+9)=3.61. (0,3)->(0,0): 3. Suma: 6+3.61+3.61+4+3.61+3 = 23.83

Korekta: (0,3)->(0,0) to sqrt(0+9)=3. Suma: 6+3.61+3.61+4+3.61+3 = 23.83, nie 22.65.

Korekta oczekiwanego wyniku:
```
a) Obwod: 23.83
```

Pole (sznurowadlo):
- (0*0 - 6*0) + (6*3 - 8*0) + (8*6 - 6*3) + (6*6 - 2*6) + (2*3 - 0*6) + (0*0 - 0*3)
= 0 + 18 + 30 + 24 + 6 + 0 = 78. |78|/2 = 39.

Korekta: Pole = 39.00

Centroid: ((0+6+8+6+2+0)/6, (0+0+3+6+6+3)/6) = (22/6, 18/6) = (3.67, 3.00)
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o zamknieciu wielokata**: Ostatni bok laczy punkt n z punktem 1. Uzyj `(i+1) % n`. CKE: -2 pkt
- **Brak `abs()` we wzorze sznurowadla**: Wynik moze byc ujemny w zaleznosci od kierunku obejscia. CKE: -1 pkt
- **Centroid != srodek wielokata wazonego polem**: Prosta srednia wspolrzednych to centroid wierzcholkow, nie pole-wazony centroid. Dla wielu zadan maturalnych prosta srednia wystarczy.

</details>

---

### Cwiczenie 14.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3 (odleglosc punkt-odcinek)
**Tagi**: `odleglosc-euklidesowa` `pole-trojkata` `struct` `wczytywanie-pliku`

W pliku `punkty.txt` znajduje sie 6 punktow (x, y) o wspolrzednych calkowitych. Pierwsze 2 punkty definiuja odcinek AB. Pozostale 4 to punkty do analizy. Napisz program, ktory:
a) Obliczy odleglosc kazdego z 4 punktow od prostej przechodzacej przez A i B.
b) Poda ktory punkt jest najblizej prostej AB.

Wzor na odleglosc punktu P od prostej przez A i B:
d = |pole_trojkata(A, B, P) * 2| / |AB|

**Dane** (`punkty.txt`):
```
0 0
10 0
3 4
7 2
5 6
1 1
```

**Oczekiwany wynik**:
```
Prosta AB: (0,0)-(10,0), dlugosc |AB| = 10.00

a) Odleglosci od prostej AB:
   (3,4): d = 4.00
   (7,2): d = 2.00
   (5,6): d = 6.00
   (1,1): d = 1.00

b) Najblizej prostej: (1,1) (d=1.00)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Odleglosc punktu od prostej = 2 * pole trojkata / dlugosc podstawy.
2. **Podejscie**: Pole trojkata ze wzoru: `|Ax*(By-Py) + Bx*(Py-Ay) + Px*(Ay-By)| / 2`. Odleglosc: `pole * 2 / |AB|`.
3. **Kluczowy krok**: Jesli AB jest poziomy (y=0), odleglosc punktu (x,y) to po prostu |y|.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <cmath>
#include <iomanip>
using namespace std;

struct Pt { int x, y; };

int main() {
    ifstream plik("punkty.txt");
    vector<Pt> P;
    int x, y;
    while (plik >> x >> y) P.push_back({x, y});

    Pt A = P[0], B = P[1];
    double dAB = sqrt(pow(B.x - A.x, 2) + pow(B.y - A.y, 2));

    cout << fixed << setprecision(2);
    cout << "Prosta AB: (" << A.x << "," << A.y << ")-("
         << B.x << "," << B.y << "), dlugosc |AB| = " << dAB << endl;

    cout << endl << "a) Odleglosci od prostej AB:" << endl;
    double minD = 1e18; int bestIdx = 2;

    for (int i = 2; i < (int)P.size(); i++) {
        Pt Q = P[i];
        double pole2 = abs(A.x * (B.y - Q.y) + B.x * (Q.y - A.y) + Q.x * (A.y - B.y));
        double d = pole2 / dAB;
        cout << "   (" << Q.x << "," << Q.y << "): d = " << d << endl;
        if (d < minD) { minD = d; bestIdx = i; }
    }

    cout << endl << "b) Najblizej prostej: (" << P[bestIdx].x << ","
         << P[bestIdx].y << ") (d=" << minD << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Odleglosc punktu od prostej = 2 * pole trojkata / dlugosc podstawy. Uzywamy wzoru na pole trojkata z wspolrzednych.

Weryfikacja (prosta y=0, A=(0,0), B=(10,0)):
- (3,4): d = |0*(0-4)+10*(4-0)+3*(0-0)| / 10 = 40/10 = 4.00
- (7,2): d = |0*(0-2)+10*(2-0)+7*(0-0)| / 10 = 20/10 = 2.00
- (5,6): d = |0*(0-6)+10*(6-0)+5*(0-0)| / 10 = 60/10 = 6.00
- (1,1): d = |0*(0-1)+10*(1-0)+1*(0-0)| / 10 = 10/10 = 1.00
Najblizej: (1,1) z d=1.00
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomylenie odleglosci od prostej z odlegloscia od odcinka**: Odleglosc od prostej jest zawsze prostopadla. Odleglosc od odcinka moze byc do konca odcinka. CKE: zalezy od tresci
- **Dzielenie przez 0**: Jesli A == B, |AB| = 0 i dzielenie sie nie powiedzie. Sprawdz dane.

</details>

---

### Cwiczenie 14.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3 (pelna geometria)
**Tagi**: `wielokat` `pole-trojkata` `punkt-w-trojkacie` `obwod` `struct` `wczytywanie-pliku`

W pliku `wielokat.txt` znajduje sie n (pierwsza linia) oraz n punktow (wspolrzedne calkowite) tworzacych wielokat wypukly podanych w kolejnosci obejscia. W pliku `zapytania.txt` znajduje sie m (pierwsza linia) oraz m punktow-zapytan. Napisz program, ktory:
a) Obliczy pole i obwod wielokata.
b) Dla kazdego punktu-zapytania sprawdzi, czy lezy wewnatrz wielokata (metoda: punkt lezy wewnatrz wielokata wypuklego jesli lezy po tej samej stronie kazdego boku).
c) Poda ile punktow-zapytan lezy wewnatrz.

**Dane** (`wielokat.txt`):
```
4
0 0
8 0
8 6
0 6
```

**Dane** (`zapytania.txt`):
```
5
4 3
9 3
0 0
8 6
4 7
```

**Oczekiwany wynik**:
```
a) Wielokat: 4 wierzcholkow
   Pole: 48.00
   Obwod: 28.00

b) Zapytania:
   (4,3): WEWNATRZ
   (9,3): NA ZEWNATRZ
   (0,0): NA BRZEGU
   (8,6): NA BRZEGU
   (4,7): NA ZEWNATRZ

c) Wewnatrz (scisle): 1
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Pole = wzor sznurowadla. Obwod = suma dlugosci bokow. Punkt wewnatrz wielokata = po tej samej stronie kazdego boku.
2. **Podejscie**: Iloczyn wektorowy boku (A->B) i wektora do punktu (A->P) mowi po ktorej stronie lezy P. Jesli wszystkie iloczyny maja ten sam znak, punkt jest wewnatrz.
3. **Kluczowy krok**: Jesli iloczyn = 0, punkt lezy na boku (brzeg). Jesli znaki sa rozne, punkt jest na zewnatrz.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <cmath>
#include <iomanip>
using namespace std;

struct Pt { int x, y; };

int cross(Pt a, Pt b, Pt p) {
    return (b.x - a.x) * (p.y - a.y) - (b.y - a.y) * (p.x - a.x);
}

int main() {
    ifstream f1("wielokat.txt"), f2("zapytania.txt");
    int n;
    f1 >> n;
    vector<Pt> W(n);
    for (int i = 0; i < n; i++) f1 >> W[i].x >> W[i].y;

    // a) Pole i obwod
    int sumPole = 0;
    double obwod = 0;
    for (int i = 0; i < n; i++) {
        int j = (i + 1) % n;
        sumPole += W[i].x * W[j].y - W[j].x * W[i].y;
        double dx = W[j].x - W[i].x;
        double dy = W[j].y - W[i].y;
        obwod += sqrt(dx * dx + dy * dy);
    }
    double pole = abs(sumPole) / 2.0;

    cout << fixed << setprecision(2);
    cout << "a) Wielokat: " << n << " wierzcholkow" << endl;
    cout << "   Pole: " << pole << endl;
    cout << "   Obwod: " << obwod << endl;

    // b) Zapytania
    int m;
    f2 >> m;
    cout << endl << "b) Zapytania:" << endl;
    int ileWewnatrz = 0;

    for (int q = 0; q < m; q++) {
        Pt P;
        f2 >> P.x >> P.y;

        bool naBrzegu = false;
        bool wewnatrz = true;
        int firstSign = 0;

        for (int i = 0; i < n; i++) {
            int j = (i + 1) % n;
            int c = cross(W[i], W[j], P);
            if (c == 0) { naBrzegu = true; break; }
            if (firstSign == 0) firstSign = (c > 0) ? 1 : -1;
            else if ((c > 0 ? 1 : -1) != firstSign) { wewnatrz = false; break; }
        }

        cout << "   (" << P.x << "," << P.y << "): ";
        if (naBrzegu) cout << "NA BRZEGU" << endl;
        else if (wewnatrz) { cout << "WEWNATRZ" << endl; ileWewnatrz++; }
        else cout << "NA ZEWNATRZ" << endl;
    }

    cout << endl << "c) Wewnatrz (scisle): " << ileWewnatrz << endl;
    return 0;
}
```

**Wyjasnienie**: Iloczyn wektorowy boku i wektora do punktu mowi po ktorej stronie lezy punkt. Jesli punkt jest po tej samej stronie kazdego boku (wszystkie iloczyny tego samego znaku), jest wewnatrz. Jesli iloczyn = 0, punkt lezy na boku.

Weryfikacja (prostokat (0,0)-(8,0)-(8,6)-(0,6)):
- Pole: |0*0-8*0 + 8*6-8*0 + 8*6-0*6 + 0*0-0*6| / 2 = |0+48+48-0+0-0+0-0|... Sznurowadlo:
  (0*0 - 8*0) + (8*6 - 8*0) + (8*6 - 0*6) + (0*0 - 0*6) = 0+48+48+0 = 96? /2 = 48. OK.
- Obwod: 8+6+8+6 = 28.
- (4,3): wewnatrz prostokata -> TAK
- (9,3): x=9 > 8 -> NA ZEWNATRZ
- (0,0): naroznik -> NA BRZEGU
- (8,6): naroznik -> NA BRZEGU
- (4,7): y=7 > 6 -> NA ZEWNATRZ
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak obslugi brzegu (cross == 0)**: Punkt na boku nie jest ani scisle wewnatrz, ani na zewnatrz. CKE: -1 pkt
- **Zapomnienie o zamknieciu wielokata w petli**: Ostatni bok laczy punkt n z punktem 1. CKE: -2 pkt
- **Zly kierunek obejscia**: Jesli punkty sa podane w odwrotnej kolejnosci, wszystkie iloczyny zmienia znak — algorytm dalej dziala, bo sprawdzamy stalosc znaku.

</details>

---

## Samoocena

| Poziom | Opis | Kryteria |
|--------|------|----------|
| Podstawowy | Umiem liczyc odleglosc euklidesowa i sprawdzac punkt w prostokacie | Cwiczenia 1-2, 6 bez pomocy |
| Dobry | Radze sobie z najblizszymi/najdalszymi punktami i srodkiem | Cwiczenia 3-4, 7 bez pomocy |
| Bardzo dobry | Umiem obliczac pole trojkata i sprawdzac punkt w trojkacie | Cwiczenie 5, 8-9 bez pomocy |
| Doskonaly | Radze sobie z wielokatami i zlozonym point-in-polygon | Cwiczenie 10 bez pomocy |

**Co dalej?**
- Jesli masz problem z `sqrt` i `cmath` -> patrz `cheatsheet_cpp.md` sekcja "Funkcje matematyczne"
- Jesli chcesz cwiczcpodwojna petle -> patrz `11_minmax.md` i `10_zliczanie.md`
- Jesli chcesz cwiczycBFS/DFS na siatkach -> patrz `13_obrazy_2D.md`
