# 14. Obliczenia geometryczne

Typ zadania: **geometryczne**
Czestotliwosc: 1/11 lat | Laczna punktacja: 4 pkt
Kategoria: IMPLEMENTACJA

---

### Cwiczenie 14.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3 (Dron)

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

---

### Cwiczenie 14.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3.2a

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

---

### Cwiczenie 14.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2018 zad. 2 (Krajobraz)

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
- (1,2)-(4,6): sqrt(9+16)=5.00
- (1,2)-(7,1): sqrt(36+1)=6.08
- (1,2)-(3,5): sqrt(4+9)=3.61
- (1,2)-(10,3): sqrt(81+1)=sqrt(82)=9.06
- (3,5)-(4,6): sqrt(1+1)=sqrt(2)=1.41
- (1,2)-(8,8): sqrt(49+36)=sqrt(85)=9.22
- (1,2)-(2,9): sqrt(1+49)=sqrt(50)=7.07

Najblizsze: (3,5)-(4,6) = 1.41
Najdalsze: sprawdzmy (1,2)-(8,8) = 9.22 vs (1,2)-(10,3) = 9.06 vs (2,9)-(10,3) = sqrt(64+36) = 10.00

Korekta: (2,9)-(10,3): sqrt(64+36)=sqrt(100)=10.00. To jest wieksze niz 9.22.
Sprawdzmy jeszcze: (2,9)-(7,1): sqrt(25+64)=sqrt(89)=9.43
(7,1)-(2,9): 9.43
(1,2)-(8,8): 9.22

Najdalsze: (2,9)-(10,3) = 10.00
</details>

---

### Cwiczenie 14.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3.2b (srodek odcinka)

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
Trojki (A, B, srodek C):
A=(0,0), B=(4,8), srodek C=(2,4) -> TAK
A=(0,0), B=(2,4), srodek C=(1,2) -> TAK
A=(2,4), B=(4,8), srodek C=(3,6)... nie ma takiego punktu
A=(4,8), B=(6,2), srodek=(5,5)... nie ma
A=(6,2), B=(8,4), srodek=(7,3)... nie ma
A=(0,0), B=(6,2), srodek=(3,1) -> TAK
A=(0,0), B=(8,4), srodek=(4,2)... nie ma
A=(2,4), B=(8,4), srodek=(5,4)... nie ma
A=(6,2), B=(4,8), srodek=(5,5)... nie ma
A=(0,0), B=(10,6)... nie ma B=(10,6)
Znalezione trojki:
(0,0)-(4,8): srodek (2,4)
(0,0)-(2,4): srodek (1,2)
(0,0)-(6,2): srodek (3,1)
Ilosc: 3
```

**Oczekiwany wynik** (czytelna forma):
```
Trojki gdzie jeden punkt jest srodkiem odcinka dwoch pozostalych:
(0,0) i (4,8) -> srodek (2,4) - jest w zbiorze
(0,0) i (2,4) -> srodek (1,2) - jest w zbiorze
(0,0) i (6,2) -> srodek (3,1) - jest w zbiorze
Ilosc: 3
```

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
- (0,0)+(8,4) = (8,4)/2 = (4,2) -> w zbiorze? NIE
- (2,4)+(4,8) = (6,12)/2 = (3,6) -> w zbiorze? NIE
- (2,4)+(6,2) = (8,6)/2 = (4,3) -> w zbiorze? NIE
- (4,8)+(6,2) = (10,10)/2 = (5,5) -> w zbiorze? NIE
- (6,2)+(8,4) = (14,6)/2 = (7,3) -> w zbiorze? NIE
- (1,2)+(3,1): (4,3)/2 = (2, 1.5) -> nie calkowite
- (1,2)+(5,3): (6,5)/2 = (3, 2.5) -> nie calkowite
- (3,1)+(5,3): (8,4)/2 = (4,2) -> w zbiorze? NIE
- (6,2)+(4,8): juz sprawdzone wyzej
</details>

---

### Cwiczenie 14.5 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3

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
- Trojkat (0,0)-(8,0)-(2,8): |0*(0-8) + 8*(8-0) + 2*(0-0)| / 2 = |0 + 64 + 0| / 2 = 32.00
- Trojkat (0,0)-(8,0)-(0,6): |0*(0-6) + 8*(6-0) + 0*(0-0)| / 2 = |0 + 48 + 0| / 2 = 24.00
- Trojkat (8,0)-(0,6)-(2,8): |8*(6-8) + 0*(8-0) + 2*(0-6)| / 2 = |(-16) + 0 + (-12)| / 2 = 28/2 = 14.00

Korekta: sprawdzmy ponownie trojke (8,0)-(0,6)-(2,8):
P = |8*(6-8) + 0*(8-0) + 2*(0-6)| / 2 = |8*(-2) + 0 + 2*(-6)| / 2 = |-16-12|/2 = 28/2 = 14.00

Trzeci co do wielkosci trojkat to nie (8,0)-(0,6)-(2,8) z polem 14. Sprawdzmy inne:
- (0,0)-(7,5)-(2,8): |0*(5-8) + 7*(8-0) + 2*(0-5)|/2 = |0+56-10|/2 = 46/2 = 23.00
- (0,0)-(8,0)-(7,5): |0*(0-5) + 8*(5-0) + 7*(0-0)|/2 = |0+40+0|/2 = 20.00
- (0,0)-(8,0)-(2,8): 32.00 (max)

c) D=(4,3) w trojkacie (0,0)-(8,0)-(2,8):
- pole ABD: A=(0,0), B=(8,0), D=(4,3): |0*(0-3)+8*(3-0)+4*(0-0)|/2 = |0+24+0|/2 = 12.00
- pole ACD: A=(0,0), C=(2,8), D=(4,3): |0*(8-3)+2*(3-0)+4*(0-8)|/2 = |0+6-32|/2 = 26/2 = 13.00
- pole BCD: B=(8,0), C=(2,8), D=(4,3): |8*(8-3)+2*(3-0)+4*(0-8)|/2 = |40+6-32|/2 = 14/2 = 7.00
- Suma: 12+13+7 = 32.00 = pole ABC -> D lezy wewnatrz: TAK
</details>

---
