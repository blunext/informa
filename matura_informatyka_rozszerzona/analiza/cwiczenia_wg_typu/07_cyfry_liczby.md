# 07. Przetwarzanie cyfr i liczb

Typ zadania: **cyfry_liczby**
Czestotliwosc: 6/11 lat | Laczna punktacja: 36 pkt
Kategoria: IMPLEMENTACJA

---

### Cwiczenie 7.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 zad. 4 (cyfry liczb)

W pliku `dane.txt` znajduje sie 10 liczb calkowitych dodatnich (kazda w osobnym wierszu). Napisz program, ktory wczyta te liczby, obliczy sume cyfr kazdej z nich i wypisze te liczby, ktorych suma cyfr jest parzysta.

**Dane** (`dane.txt`):
```
4821
13507
296
88412
5039
77164
621
45008
9273
30456
```

**Oczekiwany wynik**:
```
4821 (suma cyfr: 15) - NIE
13507 (suma cyfr: 16) - TAK
296 (suma cyfr: 17) - NIE
88412 (suma cyfr: 23) - NIE
5039 (suma cyfr: 17) - NIE
77164 (suma cyfr: 25) - NIE
621 (suma cyfr: 9) - NIE
45008 (suma cyfr: 17) - NIE
9273 (suma cyfr: 21) - NIE
30456 (suma cyfr: 18) - TAK
Liczby z parzysta suma cyfr: 2
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int sumaCyfr(int n) {
    int suma = 0;
    while (n > 0) {
        suma += n % 10;
        n /= 10;
    }
    return suma;
}

int main() {
    ifstream plik("dane.txt");
    int liczba;
    int ile = 0;
    while (plik >> liczba) {
        int s = sumaCyfr(liczba);
        cout << liczba << " (suma cyfr: " << s << ") - ";
        if (s % 2 == 0) {
            cout << "TAK" << endl;
            ile++;
        } else {
            cout << "NIE" << endl;
        }
    }
    cout << "Liczby z parzysta suma cyfr: " << ile << endl;
    return 0;
}
```

**Wyjasnienie**: Funkcja `sumaCyfr` wyodrębnia kolejne cyfry operacjami `% 10` (ostatnia cyfra) i `/ 10` (usunięcie ostatniej cyfry), sumujac je. Nastepnie sprawdzamy parzystosc sumy.

Weryfikacja:
- 4821: 4+8+2+1=15 (nieparzysta)
- 13507: 1+3+5+0+7=16 (parzysta) -> TAK
- 296: 2+9+6=17 (nieparzysta)
- 88412: 8+8+4+1+2=23 (nieparzysta)
- 5039: 5+0+3+9=17 (nieparzysta)
- 77164: 7+7+1+6+4=25 (nieparzysta)
- 621: 6+2+1=9 (nieparzysta)
- 45008: 4+5+0+0+8=17 (nieparzysta)
- 9273: 9+2+7+3=21 (nieparzysta)
- 30456: 3+0+4+5+6=18 (parzysta) -> TAK
</details>

---

### Cwiczenie 7.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4 (NWD)

W pliku `pary.txt` znajduje sie 8 par liczb calkowitych dodatnich (kazda para w osobnym wierszu, liczby oddzielone spacja). Napisz program, ktory obliczy NWD kazdej pary algorytmem Euklidesa i wypisze te pary, ktore nie sa wzglednie pierwsze (NWD > 1).

**Dane** (`pary.txt`):
```
48 18
35 22
120 45
17 13
56 42
99 55
64 24
31 29
```

**Oczekiwany wynik**:
```
48 18 -> NWD = 6
35 22 -> NWD = 1
120 45 -> NWD = 15
17 13 -> NWD = 1
56 42 -> NWD = 14
99 55 -> NWD = 11
64 24 -> NWD = 8
31 29 -> NWD = 1
Pary nie wzglednie pierwsze: 5
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int nwd(int a, int b) {
    while (b != 0) {
        int temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}

int main() {
    ifstream plik("pary.txt");
    int a, b;
    int ile = 0;
    while (plik >> a >> b) {
        int g = nwd(a, b);
        cout << a << " " << b << " -> NWD = " << g << endl;
        if (g > 1) ile++;
    }
    cout << "Pary nie wzglednie pierwsze: " << ile << endl;
    return 0;
}
```

**Wyjasnienie**: Algorytm Euklidesa oblicza NWD zamieniajac pare (a, b) na (b, a mod b) dopoki b != 0. Para jest wzglednie pierwsza gdy NWD = 1.

Weryfikacja:
- 48, 18: 48%18=12, 18%12=6, 12%6=0 -> NWD=6
- 35, 22: 35%22=13, 22%13=9, 13%9=4, 9%4=1, 4%1=0 -> NWD=1
- 120, 45: 120%45=30, 45%30=15, 30%15=0 -> NWD=15
- 17, 13: 17%13=4, 13%4=1, 4%1=0 -> NWD=1
- 56, 42: 56%42=14, 42%14=0 -> NWD=14
- 99, 55: 99%55=44, 55%44=11, 44%11=0 -> NWD=11
- 64, 24: 64%24=16, 24%16=8, 16%8=0 -> NWD=8
- 31, 29: 31%29=2, 29%2=1, 2%1=0 -> NWD=1
</details>

---

### Cwiczenie 7.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4.1 (liczby pierwsze)

W pliku `liczby.txt` znajduje sie 12 liczb calkowitych wiekszych od 1 (kazda w osobnym wierszu). Napisz program, ktory:
a) Wypisze wszystkie liczby pierwsze sposrod danych.
b) Dla kazdej liczby zlozonej podaj jej najmniejszy dzielnik wlasciwy (wiekszy od 1).

**Dane** (`liczby.txt`):
```
17
24
31
45
53
78
91
2
100
67
49
83
```

**Oczekiwany wynik**:
```
a) Liczby pierwsze: 17 31 53 2 67 83
   Ilosc: 6

b) Liczby zlozone i ich najmniejsze dzielniki:
   24 -> 2
   45 -> 3
   78 -> 2
   91 -> 7
   100 -> 2
   49 -> 7
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

bool czyPierwsza(int n) {
    if (n < 2) return false;
    if (n == 2) return true;
    if (n % 2 == 0) return false;
    for (int i = 3; i * i <= n; i += 2) {
        if (n % i == 0) return false;
    }
    return true;
}

int najmniejszyDzielnik(int n) {
    for (int i = 2; i * i <= n; i++) {
        if (n % i == 0) return i;
    }
    return n;
}

int main() {
    ifstream plik("liczby.txt");
    int n;
    vector<int> tab;
    while (plik >> n) tab.push_back(n);

    cout << "a) Liczby pierwsze: ";
    int ile = 0;
    for (int x : tab) {
        if (czyPierwsza(x)) {
            cout << x << " ";
            ile++;
        }
    }
    cout << endl << "   Ilosc: " << ile << endl;

    cout << endl << "b) Liczby zlozone i ich najmniejsze dzielniki:" << endl;
    for (int x : tab) {
        if (!czyPierwsza(x)) {
            cout << "   " << x << " -> " << najmniejszyDzielnik(x) << endl;
        }
    }
    return 0;
}
```

**Wyjasnienie**: Test pierwszosci sprawdza podzielnosc od 2 do sqrt(n). Jesli zaden dzielnik nie zostal znaleziony, liczba jest pierwsza. Najmniejszy dzielnik wlasciwy szukamy analogicznie — pierwszy znaleziony dzielnik jest najmniejszy.

Weryfikacja:
- 17: pierwsza (brak dzielnikow do sqrt(17)~4)
- 24: zlozona, 24/2=12, najmniejszy dzielnik = 2
- 31: pierwsza
- 45: zlozona, 45/3=15, najmniejszy dzielnik = 3
- 53: pierwsza
- 78: zlozona, 78/2=39, najmniejszy dzielnik = 2
- 91: zlozona, 91/7=13, najmniejszy dzielnik = 7
- 2: pierwsza
- 100: zlozona, 100/2=50, najmniejszy dzielnik = 2
- 67: pierwsza
- 49: zlozona, 49/7=7, najmniejszy dzielnik = 7
- 83: pierwsza
</details>

---

### Cwiczenie 7.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 4.2 (faktoryzacja)

W pliku `dane.txt` znajduje sie 8 liczb calkowitych wiekszych od 1 (kazda w osobnym wierszu). Napisz program, ktory dla kazdej liczby:
a) Wypisze jej rozklad na czynniki pierwsze.
b) Poda laczna liczbe czynnikow pierwszych (z powtorzeniami).

Na koniec program powinien podac, ktora liczba ma najwieksza liczbe czynnikow.

**Dane** (`dane.txt`):
```
60
17
84
128
45
97
150
72
```

**Oczekiwany wynik**:
```
60 = 2 * 2 * 3 * 5 (4 czynniki)
17 = 17 (1 czynnik)
84 = 2 * 2 * 3 * 7 (4 czynniki)
128 = 2 * 2 * 2 * 2 * 2 * 2 * 2 (7 czynnikow)
45 = 3 * 3 * 5 (3 czynniki)
97 = 97 (1 czynnik)
150 = 2 * 3 * 5 * 5 (4 czynniki)
72 = 2 * 2 * 2 * 3 * 3 (5 czynnikow)
Najwiecej czynnikow: 128 (7 czynnikow)
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
    int n;
    int maxCzynnikow = 0, maxLiczba = 0;

    while (plik >> n) {
        int oryg = n;
        vector<int> czynniki;
        int d = 2;
        while (d * d <= n) {
            while (n % d == 0) {
                czynniki.push_back(d);
                n /= d;
            }
            d++;
        }
        if (n > 1) czynniki.push_back(n);

        cout << oryg << " = ";
        for (int i = 0; i < czynniki.size(); i++) {
            if (i > 0) cout << " * ";
            cout << czynniki[i];
        }
        int ile = czynniki.size();
        cout << " (" << ile << " czynnik";
        if (ile == 1) cout << ")";
        else if (ile < 5) cout << "i)";
        else cout << "ow)";
        cout << endl;

        if (ile > maxCzynnikow) {
            maxCzynnikow = ile;
            maxLiczba = oryg;
        }
    }
    cout << "Najwiecej czynnikow: " << maxLiczba << " (" << maxCzynnikow << " czynnikow)" << endl;
    return 0;
}
```

**Wyjasnienie**: Faktoryzacja probna: dzielimy przez kolejne liczby od 2. Jesli d dzieli n, dodajemy d do czynnikow i dzielimy n. Kontynuujemy do d*d > n. Jesli na koncu n > 1, to n jest ostatnim czynnikiem pierwszym.

Weryfikacja:
- 60: 60/2=30, 30/2=15, 15/3=5, 5/5=1 -> 2*2*3*5 (4)
- 17: pierwsza -> 17 (1)
- 84: 84/2=42, 42/2=21, 21/3=7, 7/7=1 -> 2*2*3*7 (4)
- 128: 2^7 -> 2*2*2*2*2*2*2 (7)
- 45: 45/3=15, 15/3=5, 5/5=1 -> 3*3*5 (3)
- 97: pierwsza -> 97 (1)
- 150: 150/2=75, 75/3=25, 25/5=5, 5/5=1 -> 2*3*5*5 (4)
- 72: 72/2=36, 36/2=18, 18/2=9, 9/3=3, 3/3=1 -> 2*2*2*3*3 (5)
</details>

---

### Cwiczenie 7.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3.3 (skrot liczby + NWD)

W pliku `dane.txt` znajduje sie 15 liczb calkowitych dodatnich (kazda w osobnym wierszu). "Skrotem" liczby nazywamy liczbe utworzona z jej cyfr nieparzystych (w tej samej kolejnosci). Na przyklad skrotem liczby 24837 jest 37 (cyfry nieparzyste to 3 i 7), a skrotem 2468 jest 0 (brak cyfr nieparzystych).

Napisz program, ktory:
a) Dla kazdej liczby wypisze jej skrot.
b) Znajdzie wszystkie liczby, dla ktorych NWD(liczba, skrot) = 7.

**Dane** (`dane.txt`):
```
24837
1470
35291
8624
77742
5019
63154
28007
91356
42175
11368
50743
7826
39501
14287
```

**Oczekiwany wynik**:
```
a) Skroty:
24837 -> 37
1470 -> 17
35291 -> 3591
8624 -> 0
77742 -> 777
5019 -> 519
63154 -> 315
28007 -> 7
91356 -> 9135
42175 -> 175
11368 -> 113
50743 -> 573
7826 -> 7
39501 -> 3951
14287 -> 17

b) Liczby z NWD(liczba, skrot) = 7:
63154 -> skrot 315, NWD(63154, 315) = 7 -> TAK
28007 -> skrot 7, NWD(28007, 7) = 7 -> TAK
7826 -> skrot 7, NWD(7826, 7) = 7 -> TAK
Ilosc: 3
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
using namespace std;

int nwd(int a, int b) {
    while (b != 0) {
        int temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}

int skrot(int n) {
    string s = to_string(n);
    string wynik = "";
    for (char c : s) {
        int cyfra = c - '0';
        if (cyfra % 2 == 1) {
            wynik += c;
        }
    }
    if (wynik.empty()) return 0;
    return stoi(wynik);
}

int main() {
    ifstream plik("dane.txt");
    int n;
    vector<int> liczby;
    while (plik >> n) liczby.push_back(n);

    cout << "a) Skroty:" << endl;
    vector<int> skroty;
    for (int x : liczby) {
        int sk = skrot(x);
        skroty.push_back(sk);
        cout << x << " -> " << sk << endl;
    }

    cout << endl << "b) Liczby z NWD(liczba, skrot) = 7:" << endl;
    int ile = 0;
    for (int i = 0; i < liczby.size(); i++) {
        if (skroty[i] > 0) {
            int g = nwd(liczby[i], skroty[i]);
            if (g == 7) {
                cout << liczby[i] << " -> skrot " << skroty[i]
                     << ", NWD(" << liczby[i] << ", " << skroty[i]
                     << ") = " << g << " -> TAK" << endl;
                ile++;
            }
        }
    }
    cout << "Ilosc: " << ile << endl;
    return 0;
}
```

**Wyjasnienie**: Wieloetapowe przetwarzanie: (1) ekstrakcja cyfr nieparzystych i budowanie skrotu, (2) obliczenie NWD liczby i skrotu, (3) filtrowanie po warunku NWD = 7.

Weryfikacja skrotow:
- 24837: cyfry 2,4,8,3,7 -> nieparzyste: 3,7 -> skrot 37
- 1470: cyfry 1,4,7,0 -> nieparzyste: 1,7 -> skrot 17
- 35291: cyfry 3,5,2,9,1 -> nieparzyste: 3,5,9,1 -> skrot 3591
- 8624: cyfry 8,6,2,4 -> brak nieparzystych -> skrot 0
- 77742: cyfry 7,7,7,4,2 -> nieparzyste: 7,7,7 -> skrot 777
- 5019: cyfry 5,0,1,9 -> nieparzyste: 5,1,9 -> skrot 519
- 63154: cyfry 6,3,1,5,4 -> nieparzyste: 3,1,5 -> skrot 315
- 28007: cyfry 2,8,0,0,7 -> nieparzyste: 7 -> skrot 7
- 91356: cyfry 9,1,3,5,6 -> nieparzyste: 9,1,3,5 -> skrot 9135
- 42175: cyfry 4,2,1,7,5 -> nieparzyste: 1,7,5 -> skrot 175
- 11368: cyfry 1,1,3,6,8 -> nieparzyste: 1,1,3 -> skrot 113
- 50743: cyfry 5,0,7,4,3 -> nieparzyste: 5,7,3 -> skrot 573
- 7826: cyfry 7,8,2,6 -> nieparzyste: 7 -> skrot 7
- 39501: cyfry 3,9,5,0,1 -> nieparzyste: 3,9,5,1 -> skrot 3951
- 14287: cyfry 1,4,2,8,7 -> nieparzyste: 1,7 -> skrot 17

Weryfikacja NWD = 7:
- 63154, skrot 315: NWD(63154, 315) = NWD(315, 154) = NWD(154, 7) = NWD(7, 0) = 7 -> TAK
- 28007, skrot 7: NWD(28007, 7) = 7 (bo 28007 = 4001*7) -> TAK
- 7826, skrot 7: NWD(7826, 7) = 7 (bo 7826 = 1118*7) -> TAK

Pozostale: NWD(24837,37)=1, NWD(1470,17)=1, NWD(35291,3591)=1, NWD(77742,777)=21, NWD(5019,519)=3, NWD(91356,9135)=3, NWD(42175,175)=175, NWD(11368,113)=1, NWD(50743,573)=1, NWD(39501,3951)=9, NWD(14287,17)=1
</details>

---
