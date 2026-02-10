# 12. Przetwarzanie sekwencji

Typ zadania: **sekwencje**
Czestotliwosc: 3/11 lat | Laczna punktacja: 13 pkt
Kategoria: IMPLEMENTACJA

---

### Cwiczenie 12.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4.3

W pliku `dane.txt` znajduje sie 20 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory znajdzie najdluzszy spojny fragment ciagu, w ktorym wszystkie elementy sa rowne. Podaj dlugosc tego fragmentu, wartosc i pozycje poczatkowa.

**Dane** (`dane.txt`):
```
5
3
3
3
7
7
2
2
2
2
4
4
4
1
6
6
6
6
6
8
```

**Oczekiwany wynik**:
```
Najdluzszy ciag rownych elementow:
Wartosc: 6, dlugosc: 5, pozycja poczatkowa: 15
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
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    int maxDl = 1, maxStart = 0;
    int curDl = 1, curStart = 0;

    for (int i = 1; i < n; i++) {
        if (T[i] == T[i - 1]) {
            curDl++;
        } else {
            if (curDl > maxDl) {
                maxDl = curDl;
                maxStart = curStart;
            }
            curDl = 1;
            curStart = i;
        }
    }
    if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }

    cout << "Najdluzszy ciag rownych elementow:" << endl;
    cout << "Wartosc: " << T[maxStart] << ", dlugosc: " << maxDl
         << ", pozycja poczatkowa: " << maxStart + 1 << endl;
    return 0;
}
```

**Wyjasnienie**: Wzorzec current/max: sledzony jest biezacy ciag rownych elementow. Gdy element rozni sie od poprzedniego, porownujemy biezaca dlugosc z najdluzsza i resetujemy.

Weryfikacja:
- 5 (dl.1), 3,3,3 (dl.3, poz.2), 7,7 (dl.2), 2,2,2,2 (dl.4, poz.7), 4,4,4 (dl.3), 1 (dl.1), 6,6,6,6,6 (dl.5, poz.15), 8 (dl.1)
Najdluzszy: 6 powtorzone 5 razy od pozycji 15.
</details>

---

### Cwiczenie 12.2 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3.3

W pliku `dane.txt` znajduje sie 15 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory znajdzie najdluzszy spojny podciag scisle rosnacy. Podaj pozycje poczatkowa, dlugosc i elementy.

**Dane** (`dane.txt`):
```
5
8
12
3
6
9
14
18
25
2
4
7
1
10
15
```

**Oczekiwany wynik**:
```
Najdluzszy ciag rosnacy:
Pozycja poczatkowa: 4, dlugosc: 6
Elementy: 3 6 9 14 18 25
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
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    int maxDl = 1, maxStart = 0;
    int curDl = 1, curStart = 0;

    for (int i = 1; i < n; i++) {
        if (T[i] > T[i - 1]) {
            curDl++;
        } else {
            if (curDl > maxDl) {
                maxDl = curDl;
                maxStart = curStart;
            }
            curDl = 1;
            curStart = i;
        }
    }
    if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }

    cout << "Najdluzszy ciag rosnacy:" << endl;
    cout << "Pozycja poczatkowa: " << maxStart + 1 << ", dlugosc: " << maxDl << endl;
    cout << "Elementy: ";
    for (int i = maxStart; i < maxStart + maxDl; i++)
        cout << T[i] << " ";
    cout << endl;
    return 0;
}
```

**Wyjasnienie**: Analogicznie do ciagu rownych, ale warunek to `T[i] > T[i-1]` (scisly wzrost).

Weryfikacja:
- 5,8,12 (dl.3, poz.1)
- 3,6,9,14,18,25 (dl.6, poz.4) -> najdluzszy
- 2,4,7 (dl.3, poz.10)
- 1,10,15 (dl.3, poz.13)
Najdluzszy: 6 elementow od pozycji 4.
</details>

---

### Cwiczenie 12.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3.3 (rosnaco-malejacy)

W pliku `dane.txt` znajduje sie 20 liczb calkowitych (kazda w osobnym wierszu). Fragment ciagu nazywamy "gorka" jesli najpierw scisle rosnie, a potem scisle maleje (czesc rosnaca i malejaca musza miec co najmniej po 2 elementy, wiec minimalna gorka ma dlugosc 3). Napisz program, ktory znajdzie najdluzsza gorke w ciagu. Podaj pozycje poczatkowa, dlugosc i elementy.

**Dane** (`dane.txt`):
```
3
5
8
6
4
10
15
20
25
22
18
12
7
1
9
13
11
6
2
14
```

**Oczekiwany wynik**:
```
Najdluzsza gorka:
Pozycja poczatkowa: 6, dlugosc: 9
Elementy: 10 15 20 25 22 18 12 7 1
Szczyt: 25 (pozycja 9)
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
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    int bestLen = 0, bestStart = 0, bestPeak = 0;

    for (int i = 0; i < n - 2; i++) {
        // Szukaj wzrostu od pozycji i
        int j = i;
        while (j + 1 < n && T[j + 1] > T[j]) j++;
        if (j == i) continue; // brak wzrostu
        int peak = j;
        // Szukaj spadku od szczytu
        while (j + 1 < n && T[j + 1] < T[j]) j++;
        if (j == peak) continue; // brak spadku
        int len = j - i + 1;
        if (len > bestLen) {
            bestLen = len;
            bestStart = i;
            bestPeak = peak;
        }
    }

    cout << "Najdluzsza gorka:" << endl;
    cout << "Pozycja poczatkowa: " << bestStart + 1
         << ", dlugosc: " << bestLen << endl;
    cout << "Elementy: ";
    for (int i = bestStart; i < bestStart + bestLen; i++)
        cout << T[i] << " ";
    cout << endl;
    cout << "Szczyt: " << T[bestPeak]
         << " (pozycja " << bestPeak + 1 << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Dla kazdej pozycji startowej szukamy fazy rosnacej (dopoki nastepny > biezacy), nastepnie fazy malejacej (dopoki nastepny < biezacy). Jezeli obie fazy maja co najmniej 1 krok (czesc rosnaca >= 2 elementy, malejaca >= 2), mamy gorke.

Weryfikacja:
- Poz.1: 3,5,8,6,4 -> roslaco: 3,5,8 (szczyt poz.3), malejaco: 8,6,4. Gorka dl.5 (poz.1-5).
- Poz.6: 10,15,20,25,22,18,12,7,1 -> rosnaco: 10,15,20,25 (szczyt poz.9), malejaco: 25,22,18,12,7,1. Gorka dl.9 (poz.6-14).
- Poz.15: 9,13,11,6,2 -> rosnaco: 9,13 (szczyt poz.16), malejaco: 13,11,6,2. Gorka dl.5 (poz.15-19).
Najdluzsza: dl.9 (poz.6-14), szczyt=25 (poz.9).
</details>

---

### Cwiczenie 12.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4.3 (najdluzszy ciag z NWD > 1)

W pliku `dane.txt` znajduje sie 15 liczb calkowitych wiekszych od 1 (kazda w osobnym wierszu). Napisz program, ktory znajdzie najdluzszy spojny podciag, w ktorym NWD wszystkich elementow jest wiekszy od 1. Podaj dlugosc, pozycje startowa i NWD tego podciagu.

**Dane** (`dane.txt`):
```
12
18
6
24
15
7
14
21
35
49
11
22
33
44
5
```

**Oczekiwany wynik**:
```
Najdluzszy ciag z NWD > 1:
Pozycja poczatkowa: 1, dlugosc: 5
Elementy: 12 18 6 24 15
NWD: 3
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int nwd(int a, int b) {
    while (b != 0) { int t = b; b = a % b; a = t; }
    return a;
}

int main() {
    ifstream plik("dane.txt");
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    int bestLen = 0, bestStart = 0, bestNwd = 0;

    for (int i = 0; i < n; i++) {
        int g = T[i];
        for (int j = i; j < n; j++) {
            g = nwd(g, T[j]);
            if (g == 1) break;
            int len = j - i + 1;
            if (len > bestLen) {
                bestLen = len;
                bestStart = i;
                bestNwd = g;
            }
        }
    }

    cout << "Najdluzszy ciag z NWD > 1:" << endl;
    cout << "Pozycja poczatkowa: " << bestStart + 1
         << ", dlugosc: " << bestLen << endl;
    cout << "Elementy: ";
    for (int i = bestStart; i < bestStart + bestLen; i++)
        cout << T[i] << " ";
    cout << endl;
    cout << "NWD: " << bestNwd << endl;
    return 0;
}
```

**Wyjasnienie**: Dla kazdej pozycji startowej obliczamy NWD narastajaco (dodajac kolejne elementy). Gdy NWD spadnie do 1, przerywamy. Sledzony jest najdluzszy ciag z NWD > 1.

Weryfikacja:
- Poz.1: NWD(12)=12, NWD(12,18)=6, NWD(6,6)=6, NWD(6,24)=6, NWD(6,15)=3 -> ciag dl.5, NWD=3. NWD(3,7)=1, stop.
- Poz.6: NWD(7)=7, NWD(7,14)=7, NWD(7,21)=7, NWD(7,35)=7, NWD(7,49)=7, NWD(7,11)=1 -> ciag dl.5? Ale poz.6 to 7, nie 14.

Poprawka: T = [12,18,6,24,15,7,14,21,35,49,11,22,33,44,5] (1-indexed)
- Poz.1: 12,18,6,24,15 -> NWD(12,18)=6, NWD(6,6)=6, NWD(6,24)=6, NWD(6,15)=3 -> dl.5, NWD=3
  NWD(3,7)=1, stop
- Poz.6: 7 -> NWD=7, NWD(7,14)=7, NWD(7,21)=7, NWD(7,35)=7, NWD(7,49)=7 -> dl.5, NWD=7
  NWD(7,11)=1, stop

Mamy dwa ciagi dlugosci 5. Pierwszy znaleziony: poz.1 (dl.5, NWD=3). Poz.6 tez dl.5 (NWD=7).
bestLen nie zmieni sie przy rownej dlugosci (warunek `len > bestLen`).

Korekta oczekiwanego wyniku:
```
Pozycja poczatkowa: 1, dlugosc: 5
Elementy: 12 18 6 24 15
NWD: 3
```

Lub jesli szukamy pierwszego najdluzszego, to poz.1.
Alternatywnie zmienmy dane aby odpowiedz byla jednoznaczna.
</details>

---

### Cwiczenie 12.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4.4 (maksymalna srednia segmentu)

W pliku `dane.txt` znajduje sie 20 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory znajdzie spojny podciag o dlugosci co najmniej 5 z maksymalna srednia arytmetyczna. Podaj pozycje startowa, dlugosc i srednia.

**Dane** (`dane.txt`):
```
10
5
3
8
12
7
2
15
20
18
25
22
19
4
6
11
14
9
1
16
```

**Oczekiwany wynik**:
```
Optymalny segment (dlugosc >= 5):
Pozycja poczatkowa: 9, dlugosc: 5
Elementy: 20 18 25 22 19
Srednia: 20.80
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    // Sumy prefiksowe
    vector<int> prefix(n + 1, 0);
    for (int i = 0; i < n; i++)
        prefix[i + 1] = prefix[i] + T[i];

    double bestSr = -1;
    int bestStart = 0, bestLen = 5;

    for (int i = 0; i < n; i++) {
        for (int len = 5; i + len <= n; len++) {
            int suma = prefix[i + len] - prefix[i];
            double sr = (double)suma / len;
            if (sr > bestSr) {
                bestSr = sr;
                bestStart = i;
                bestLen = len;
            }
        }
    }

    cout << "Optymalny segment (dlugosc >= 5):" << endl;
    cout << "Pozycja poczatkowa: " << bestStart + 1
         << ", dlugosc: " << bestLen << endl;
    cout << "Elementy: ";
    for (int i = bestStart; i < bestStart + bestLen; i++)
        cout << T[i] << " ";
    cout << endl;
    cout << "Srednia: " << fixed << setprecision(2) << bestSr << endl;
    return 0;
}
```

**Wyjasnienie**: Sumy prefiksowe pozwalaja obliczyc sume dowolnego fragmentu w O(1). Sprawdzamy wszystkie mozliwe fragmenty o dlugosci >= 5, szukajac tego z najwyzsza srednia.

Weryfikacja (kilka kandydatow):
- Poz.8-12 (dl.5): 15+20+18+25+22 = 100, sr=20.00
- Poz.8-13 (dl.6): 15+20+18+25+22+19 = 119, sr=19.83
- Poz.9-13 (dl.5): 20+18+25+22+19 = 104, sr=20.80
- Poz.9-14 (dl.6): 20+18+25+22+19+4 = 108, sr=18.00

Korekta: Poz.9-13 (dl.5) ma srednia 20.80, co jest wyzsze niz 20.00.
Sprawdzmy dalej: Poz.10-14 (dl.5): 18+25+22+19+4 = 88, sr=17.60
Poz.8-12 (dl.5): 15+20+18+25+22 = 100, sr=20.00
Poz.9-13 (dl.5): 20+18+25+22+19 = 104, sr=20.80 -> to lepsze!

Poprawiony oczekiwany wynik:
Pozycja poczatkowa: 9, dlugosc: 5
Elementy: 20 18 25 22 19
Srednia: 20.80
</details>

---
