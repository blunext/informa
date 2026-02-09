# 10. Zliczanie i filtrowanie

Typ zadania: **zliczanie**
Czestotliwosc: 5/11 lat | Laczna punktacja: 17 pkt
Kategoria: IMPLEMENTACJA

---

### Cwiczenie 10.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 5a

W pliku `dane.txt` znajduje sie 15 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory:
a) Zliczy ile sposrod nich jest parzystych.
b) Zliczy ile jest wiekszych od 100.

**Dane** (`dane.txt`):
```
42
155
7
200
88
13
176
51
99
300
64
111
25
148
33
```

**Oczekiwany wynik**:
```
a) Liczby parzyste: 7
b) Liczby wieksze od 100: 6
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    int n;
    int parzyste = 0, powyzej100 = 0;
    while (plik >> n) {
        if (n % 2 == 0) parzyste++;
        if (n > 100) powyzej100++;
    }
    cout << "a) Liczby parzyste: " << parzyste << endl;
    cout << "b) Liczby wieksze od 100: " << powyzej100 << endl;
    return 0;
}
```

**Wyjasnienie**: Prosta petla z dwoma licznikami. Warunek parzystosci: `n % 2 == 0`. Warunek wiekszosci: `n > 100`.

Weryfikacja:
- Parzyste: 42, 200, 88, 176, 300, 64, 148 = 7
- Wieksze od 100: 155, 200, 176, 300, 111, 148 = 6
</details>

---

### Cwiczenie 10.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3.1 (cyfry pi)

Dany jest ciag 50 cyfr (kolejne cyfry po przecinku liczby pi). Napisz program, ktory zliczy wystapienia kazdej cyfry 0-9 i poda najczesciej wystepujaca cyfre.

**Dane** (ciag cyfr):
```
14159265358979323846264338327950288419716939937510
```

**Oczekiwany wynik**:
```
Czestotliwosc cyfr:
0: 2
1: 5
2: 5
3: 8
4: 4
5: 5
6: 4
7: 4
8: 5
9: 8
Najczesciej: 3 i 9 (8 razy)
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <string>
using namespace std;

int main() {
    string cyfry = "14159265358979323846264338327950288419716939937510";
    int freq[10] = {0};

    for (char c : cyfry) {
        freq[c - '0']++;
    }

    cout << "Czestotliwosc cyfr:" << endl;
    int maxFreq = 0;
    for (int i = 0; i < 10; i++) {
        cout << i << ": " << freq[i] << endl;
        if (freq[i] > maxFreq) maxFreq = freq[i];
    }

    cout << "Najczesciej: ";
    bool first = true;
    for (int i = 0; i < 10; i++) {
        if (freq[i] == maxFreq) {
            if (!first) cout << " i ";
            cout << i;
            first = false;
        }
    }
    cout << " (" << maxFreq << " razy)" << endl;
    return 0;
}
```

**Wyjasnienie**: Tablica czestotliwosci `freq[10]` indeksowana cyfra (0-9). Iterujemy po ciagu znakow, konwertujac kazdy na cyfre `c - '0'` i inkrementujac odpowiedni licznik. Na koniec szukamy maksymalnej czestotliwosci.

Weryfikacja (ciag: 14159265358979323846264338327950288419716939937510):
- 0: pojawiaja sie na poz. 30(0), 50(0) = 2
- 1: poz. 1,4,38,41,49 -> raczej policzmy: 1,4,1,5,9,2,6,5,3,5,8,9,7,9,3,2,3,8,4,6,2,6,4,3,3,8,3,2,7,9,5,0,2,8,8,4,1,9,7,1,6,9,3,9,9,3,7,5,1,0
  Cyfra 1: pozycje 1,4,37,40,49 = 5 razy
- 3: pozycje 9,15,17,24,25,27,43,46 = 8 razy
- 9: pozycje 5,12,14,30,38,42,44,45 = 8 razy
</details>

---

### Cwiczenie 10.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3.2

W pliku `dane.txt` znajduje sie 12 liczb 5-cyfrowych (kazda w osobnym wierszu). Napisz program, ktory:
a) Zliczy ile z nich ma sume cyfr wieksza od 20.
b) Zliczy ile ma pierwsza cyfre wieksza od ostatniej.
c) Zliczy ile jest podzielnych przez sume swoich cyfr.

**Dane** (`dane.txt`):
```
12345
99876
54321
11111
87654
33333
76543
44444
65432
28916
55555
19827
```

**Oczekiwany wynik**:
```
a) Suma cyfr > 20:
   99876 (suma=39)
   87654 (suma=30)
   76543 (suma=25)
   28916 (suma=26)
   55555 (suma=25)
   19827 (suma=27)
   Ilosc: 6

b) Pierwsza cyfra > ostatnia:
   99876 (9>6)
   54321 (5>1)
   87654 (8>4)
   76543 (7>3)
   65432 (6>2)
   Ilosc: 5

c) Podzielne przez sume cyfr:
   12345 (suma cyfr=15, 12345/15=823)
   Ilosc: 1
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int sumaCyfr(int n) {
    int s = 0;
    while (n > 0) { s += n % 10; n /= 10; }
    return s;
}

int main() {
    ifstream plik("dane.txt");
    vector<int> tab;
    int x;
    while (plik >> x) tab.push_back(x);

    // a)
    cout << "a) Suma cyfr > 20:" << endl;
    int ileA = 0;
    for (int n : tab) {
        int s = sumaCyfr(n);
        if (s > 20) {
            cout << "   " << n << " (suma=" << s << ")" << endl;
            ileA++;
        }
    }
    cout << "   Ilosc: " << ileA << endl;

    // b)
    cout << endl << "b) Pierwsza cyfra > ostatnia:" << endl;
    int ileB = 0;
    for (int n : tab) {
        int pierwsza = n;
        while (pierwsza >= 10) pierwsza /= 10;
        int ostatnia = n % 10;
        if (pierwsza > ostatnia) {
            cout << "   " << n << " (" << pierwsza << ">" << ostatnia << ")" << endl;
            ileB++;
        }
    }
    cout << "   Ilosc: " << ileB << endl;

    // c)
    cout << endl << "c) Podzielne przez sume cyfr:" << endl;
    int ileC = 0;
    for (int n : tab) {
        int s = sumaCyfr(n);
        if (s > 0 && n % s == 0) {
            cout << "   " << n << " (suma cyfr=" << s << ", " << n << "/" << s << "=" << n/s << ")" << endl;
            ileC++;
        }
    }
    cout << "   Ilosc: " << ileC << endl;
    return 0;
}
```

**Wyjasnienie**: Trzy niezalezne zliczania. (a) Sumujemy cyfry petla mod/div i porownujemy z 20. (b) Pierwsza cyfra: dzielimy przez 10 dopoki n >= 10. Ostatnia cyfra: n % 10. (c) Sprawdzamy podzielnosc liczby przez sume jej cyfr.

Weryfikacja:
a) Sumy cyfr: 12345(15), 99876(39), 54321(15), 11111(5), 87654(30), 33333(15), 76543(25), 44444(20), 65432(20), 28916(26), 55555(25), 19827(27)
   Wieksze od 20: 99876(39), 87654(30), 76543(25), 28916(26), 55555(25), 19827(27) = 6
</details>

---

### Cwiczenie 10.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3.1

W pliku `dane.txt` znajduje sie 20 liczb calkowitych (kazda w osobnym wierszu, wartosci z zakresu 1-50, moga sie powtarzac). Napisz program, ktory:
a) Zliczy ile jest roznych wartosci.
b) Znajdzie wartosc wystepujaca najczesciej (mode).
c) Zliczy ile wartosci wystepuje dokladnie raz.

**Dane** (`dane.txt`):
```
7
15
3
7
22
15
3
41
7
15
22
3
7
50
15
3
22
7
41
15
```

**Oczekiwany wynik**:
```
a) Roznych wartosci: 6

b) Najczesciej wystepujaca wartosc: 7 (5 razy)
   Rowniez: 15 (5 razy)

c) Wartosci wystepujace dokladnie raz: 1 (wartosc: 50)
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <map>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    int x;
    map<int, int> freq;
    while (plik >> x) freq[x]++;

    // a)
    cout << "a) Roznych wartosci: " << freq.size() << endl;

    // b)
    int maxFreq = 0;
    for (auto &p : freq)
        if (p.second > maxFreq) maxFreq = p.second;

    cout << endl << "b) Najczesciej wystepujaca wartosc: ";
    bool first = true;
    for (auto &p : freq) {
        if (p.second == maxFreq) {
            if (!first) cout << "   Rowniez: ";
            cout << p.first << " (" << maxFreq << " razy)" << endl;
            first = false;
        }
    }

    // c)
    int ileRaz = 0;
    cout << endl << "c) Wartosci wystepujace dokladnie raz: ";
    for (auto &p : freq)
        if (p.second == 1) ileRaz++;
    cout << ileRaz << " (wartosc: ";
    first = true;
    for (auto &p : freq) {
        if (p.second == 1) {
            if (!first) cout << ", ";
            cout << p.first;
            first = false;
        }
    }
    cout << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Mapa czestotliwosci `map<int, int>` zlicza ile razy kazda wartosc wystepuje. Rozmiar mapy to liczba roznych wartosci. Mode to wartosc z maksymalna czestotliwoscia. Wartosci unikalne maja czestotliwosc 1.

Weryfikacja:
- 3: wystepuje 4 razy (poz. 3,7,12,16)
- 7: wystepuje 5 razy (poz. 1,4,9,13,18)
- 15: wystepuje 5 razy (poz. 2,6,10,15,20)
- 22: wystepuje 3 razy (poz. 5,11,17)
- 41: wystepuje 2 razy (poz. 8,19)
- 50: wystepuje 1 raz (poz. 14)
Roznych: 6, max: 7 i 15 (po 5), dokladnie raz: 50
</details>

---

### Cwiczenie 10.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4.1 (liczby pierwsze dzielace)

Dane sa dwa zbiory: 8 liczb pierwszych i 10 duzych liczb calkowitych. Napisz program, ktory:
a) Zliczy ile z podanych liczb pierwszych dzieli chociaz jedna z duzych liczb.
b) Zliczy ile par (pierwsza, duza) spelnia warunek podzielnosci.
c) Znajdzie liczbe pierwsza, ktora dzieli najwiecej duzych liczb.

**Dane**:

Liczby pierwsze (`pierwsze.txt`):
```
2
3
5
7
11
13
17
19
```

Duze liczby (`duze.txt`):
```
210
143
85
66
119
51
34
78
95
231
```

**Oczekiwany wynik**:
```
a) Liczby pierwsze dziealce chociaz jedna duza: 8 (wszystkie)

b) Pary (pierwsza, duza) spelniajace podzielnosc:
   2 dzieli: 210, 66, 34, 78 (4 duze)
   3 dzieli: 210, 66, 51, 78, 231 (5 duzych)
   5 dzieli: 210, 85, 95 (3 duze)
   7 dzieli: 210, 119, 231 (3 duze)
   11 dzieli: 143, 66, 231 (3 duze)
   13 dzieli: 143, 78 (2 duze)
   17 dzieli: 85, 51, 34, 119 (4 duze)
   19 dzieli: 95 (1 duza)
   Laczna liczba par: 4+5+3+3+3+2+4+1 = 25

c) Liczba pierwsza dzielaca najwiecej duzych: 3 (5 duzych)
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
    ifstream f1("pierwsze.txt"), f2("duze.txt");
    vector<int> pierwsze, duze;
    int x;
    while (f1 >> x) pierwsze.push_back(x);
    while (f2 >> x) duze.push_back(x);

    // a) i b) i c)
    int ileA = 0, ileB = 0;
    int maxDzieli = 0, maxPierwsza = 0;

    cout << "b) Pary (pierwsza, duza) spelniajace podzielnosc:" << endl;
    for (int p : pierwsze) {
        int cnt = 0;
        cout << "   " << p << " dzieli: ";
        bool first = true;
        for (int d : duze) {
            if (d % p == 0) {
                if (!first) cout << ", ";
                cout << d;
                first = false;
                cnt++;
            }
        }
        cout << " (" << cnt << " duzych)" << endl;
        ileB += cnt;
        if (cnt > 0) ileA++;
        if (cnt > maxDzieli) { maxDzieli = cnt; maxPierwsza = p; }
    }

    cout << endl << "a) Liczby pierwsze dzielace chociaz jedna duza: " << ileA << endl;
    cout << "   Laczna liczba par: " << ileB << endl;
    cout << endl << "c) Liczba pierwsza dzielaca najwiecej duzych: "
         << maxPierwsza << " (" << maxDzieli << " duzych)" << endl;
    return 0;
}
```

**Wyjasnienie**: Podwojna petla: dla kazdej liczby pierwszej sprawdzamy podzielnosc kazdej duzej. Zliczamy pary, sprawdzamy ktore pierwsze dzielа chociaz jedna duza, i szukamy pierwszej z max liczba dzielonych duzych.

Weryfikacja (rozklady duzych):
- 210 = 2*3*5*7 -> dzielniki: 2,3,5,7
- 143 = 11*13 -> dzielniki: 11,13
- 85 = 5*17 -> dzielniki: 5,17
- 66 = 2*3*11 -> dzielniki: 2,3,11
- 119 = 7*17 -> dzielniki: 7,17
- 51 = 3*17 -> dzielniki: 3,17
- 34 = 2*17 -> dzielniki: 2,17
- 78 = 2*3*13 -> dzielniki: 2,3,13
- 95 = 5*19 -> dzielniki: 5,19
- 231 = 3*7*11 -> dzielniki: 3,7,11

Zliczenia:
- 2: 210,66,34,78 = 4
- 3: 210,66,51,78,231 = 5
- 5: 210,85,95 = 3
- 7: 210,119,231 = 3
- 11: 143,66,231 = 3
- 13: 143,78 = 2
- 17: 85,51,34,119 = 4
- 19: 95 = 1

Laczna liczba par: 4+5+3+3+3+2+4+1 = 25
Max: 3 (5 duzych)
Wszystkie 8 pierwszych dziela chociaz jedna duza: TAK
</details>

---
