# 12. Przetwarzanie sekwencji

Typ zadania: **sekwencje**
Czestotliwosc: 3/12 lat | Laczna punktacja: 13 pkt
Kategoria: IMPLEMENTACJA

## Umiejetnosci cwiczone w tym zestawie

`ciag-rownych` `ciag-rosnacy` `current-max` `gorka` `NWD-Euklidesa` `spojny-podciag` `srednia-segmentu` `sumy-prefiksowe` `wczytywanie-pliku` `ciag-malejacy` `plateau` `oscylacja` `roznica-sasiadow` `skan-liniowy` `ciag-stalej-roznicy`

---

### Cwiczenie 12.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4.3
**Tagi**: `ciag-rownych` `current-max` `skan-liniowy` `wczytywanie-pliku`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Wzorzec current/max — sledz biezacy ciag rownych i najdluzszy znaleziony.
2. **Podejscie**: Jesli `T[i] == T[i-1]` to curDl++, w przeciwnym razie porownaj z max i zresetuj.
3. **Kluczowy krok**: Nie zapomnij sprawdzic po petli — najdluzszy ciag moze konczyc sie na ostatnim elemencie.

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

<details>
<summary>Typowe bledy</summary>

- **Brak sprawdzenia po petli**: Jesli najdluzszy ciag konczy tablice, `curDl > maxDl` nigdy nie zostanie sprawdzone w petli. CKE: -2 pkt
- **Inicjalizacja curDl=0**: Pierwszy element sam w sobie tworzy ciag dlugosci 1, nie 0. CKE: -1 pkt

</details>

---

### Cwiczenie 12.2 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3.3
**Tagi**: `ciag-rosnacy` `current-max` `skan-liniowy` `wczytywanie-pliku`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Ten sam wzorzec current/max co w 12.1, ale warunek to `T[i] > T[i-1]` (scisly wzrost).
2. **Podejscie**: Zmien warunek rownosci na warunek wzrostu.
3. **Kluczowy krok**: Zapamietaj pozycje poczatkowa i dlugosc najdluzszego podciagu — elementy wypisz z wektora.

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

<details>
<summary>Typowe bledy</summary>

- **`>=` zamiast `>`**: "Scisle rosnacy" oznacza bez rownosci. Z `>=` ciag 3,3,5 tez bylby "rosnacy". CKE: -1 pkt
- **Brak aktualizacji po petli**: Jak w 12.1. CKE: -2 pkt

</details>

---

### Cwiczenie 12.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3.3 (rosnaco-malejacy)
**Tagi**: `gorka` `ciag-rosnacy` `ciag-malejacy` `skan-liniowy`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Szukaj faz: najpierw wzrost, potem spadek. Gorka to przejscie wzrost->spadek.
2. **Podejscie**: Dla kazdej pozycji startowej szukaj konca fazy rosnacej (szczyt), potem konca fazy malejacej.
3. **Kluczowy krok**: Minimalna gorka: 2 elementy w fazie rosnacej + 1 dodatkowy w malejacej = 3 elementy lacznie.

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
- Poz.1: 3,5,8,6,4 -> rosnaco: 3,5,8 (szczyt poz.3), malejaco: 8,6,4. Gorka dl.5 (poz.1-5).
- Poz.6: 10,15,20,25,22,18,12,7,1 -> rosnaco: 10,15,20,25 (szczyt poz.9), malejaco: 25,22,18,12,7,1. Gorka dl.9 (poz.6-14).
- Poz.15: 9,13,11,6,2 -> rosnaco: 9,13 (szczyt poz.16), malejaco: 13,11,6,2. Gorka dl.5 (poz.15-19).
Najdluzsza: dl.9 (poz.6-14), szczyt=25 (poz.9).
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak sprawdzenia fazy malejacej**: Ciag tylko rosnacy (bez spadku) nie jest gorka. CKE: -1 pkt
- **Zliczanie elementow plateau jako gorki**: `T[j+1] == T[j]` nie jest ani wzrostem, ani spadkiem — przerywamy faze. CKE: -1 pkt
- **Szukanie gorek od kazdej pozycji**: Mozna zoptymalizowac, ale na danych maturalnych (20-1000 elementow) O(n^2) jest OK.

</details>

---

### Cwiczenie 12.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4.3 (najdluzszy ciag z NWD > 1)
**Tagi**: `NWD-Euklidesa` `spojny-podciag` `current-max` `wczytywanie-pliku`

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
<summary>Wskazowki</summary>

1. **Kierunek**: NWD narastajaco — NWD(a,b,c) = NWD(NWD(a,b), c). Jesli NWD spadnie do 1, przerywamy.
2. **Podejscie**: Podwojna petla: dla kazdego startu oblicz NWD narastajaco. `break` gdy NWD == 1.
3. **Kluczowy krok**: NWD moze tylko malec lub pozostac stalym — nigdy nie wzrosnie po dodaniu elementu.

</details>

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
- Poz.6: 7, NWD(7,14)=7, NWD(7,21)=7, NWD(7,35)=7, NWD(7,49)=7 -> dl.5, NWD=7. NWD(7,11)=1, stop.

Mamy dwa ciagi dlugosci 5. Pierwszy znaleziony: poz.1 (dl.5, NWD=3). Poz.6 tez dl.5 (NWD=7).
bestLen nie zmieni sie przy rownej dlugosci (warunek `len > bestLen`), wiec odpowiedz to poz.1.
</details>

<details>
<summary>Typowe bledy</summary>

- **Bledna implementacja NWD**: Klasyczny blad to `while (a != b)` zamiast algorytmu Euklidesa z modulo — dziala, ale jest wolny. CKE: 0 pkt (poprawne, ale moze TLE)
- **Brak `break` przy NWD==1**: Kontynuowanie po NWD=1 jest bezcelowe (NWD nigdy nie wzrosnie), ale nie jest bledem logicznym. Moze spowodowac timeout.
- **Zapomnienie o NWD == bestNwd przy rownej dlugosci**: Nie jest wymagane — wystarczy pierwszy znaleziony.

</details>

---

### Cwiczenie 12.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4.4 (maksymalna srednia segmentu)
**Tagi**: `srednia-segmentu` `sumy-prefiksowe` `spojny-podciag` `wczytywanie-pliku`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Sumy prefiksowe pozwalaja obliczyc sume kazdego fragmentu w O(1).
2. **Podejscie**: `prefix[i+len] - prefix[i]` to suma fragmentu od i dlugosci len. Srednia = suma/len.
3. **Kluczowy krok**: Sprawdz wszystkie fragmenty dlugosci >= 5. Najwyzsza srednia czesto wystepuje przy najkrotszym fragmencie (dl. 5).

</details>

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
- Poz.9-13 (dl.5): 20+18+25+22+19 = 104, sr=20.80
- Poz.8-13 (dl.6): 15+20+18+25+22+19 = 119, sr=19.83
Najlepsza: Poz.9-13 (dl.5), srednia 20.80.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak sum prefiksowych (O(n^3) zamiast O(n^2))**: Obliczanie sumy w petli za kazdym razem od nowa. Dziala, ale moze timeout na duzych danych. CKE: 0 pkt (poprawne, ale wolne)
- **Dzielenie calkowite**: `suma / len` zamiast `(double)suma / len`. CKE: -1 pkt
- **Minimalna dlugosc 4 zamiast 5**: Zle odczytanie warunku "co najmniej 5". CKE: -1 pkt

</details>

---

### Cwiczenie 12.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4 (ciag malejacy)
**Tagi**: `ciag-malejacy` `current-max` `skan-liniowy` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 15 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory znajdzie najdluzszy spojny podciag scisle malejacy. Podaj pozycje poczatkowa, dlugosc i elementy.

**Dane** (`dane.txt`):
```
10
8
6
4
12
15
13
9
7
3
20
18
16
14
11
```

**Oczekiwany wynik**:
```
Najdluzszy ciag malejacy:
Pozycja poczatkowa: 6, dlugosc: 5
Elementy: 15 13 9 7 3
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Ten sam wzorzec current/max co w 12.2, ale warunek to `T[i] < T[i-1]` (scisly spadek).
2. **Podejscie**: Jedyna zmiana: warunek `T[i] > T[i-1]` zamien na `T[i] < T[i-1]`.
3. **Kluczowy krok**: Nie zapomnij sprawdzic po petli.

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
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    int maxDl = 1, maxStart = 0;
    int curDl = 1, curStart = 0;

    for (int i = 1; i < n; i++) {
        if (T[i] < T[i - 1]) {
            curDl++;
        } else {
            if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }
            curDl = 1; curStart = i;
        }
    }
    if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }

    cout << "Najdluzszy ciag malejacy:" << endl;
    cout << "Pozycja poczatkowa: " << maxStart + 1 << ", dlugosc: " << maxDl << endl;
    cout << "Elementy: ";
    for (int i = maxStart; i < maxStart + maxDl; i++)
        cout << T[i] << " ";
    cout << endl;
    return 0;
}
```

**Wyjasnienie**: Wzorzec current/max z warunkiem `T[i] < T[i-1]` (spadek).

Weryfikacja:
- 10,8,6,4 (dl.4, poz.1)
- 15,13,9,7,3 (dl.5, poz.6) -> najdluzszy
- 20,18,16,14,11 (dl.5, poz.11) -> rowny, ale nie dluzszy
Najdluzszy (pierwszy): dl.5 od poz.6.
</details>

<details>
<summary>Typowe bledy</summary>

- **`<=` zamiast `<`**: "Scisle malejacy" wyklucza rownosc. CKE: -1 pkt
- **Brak aktualizacji po petli**: Ciag 20,18,16,14,11 konczy tablice — bez sprawdzenia po petli moze zostac pominiety.

</details>

---

### Cwiczenie 12.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3 (plateau)
**Tagi**: `plateau` `ciag-rownych` `ciag-rosnacy` `current-max`

W pliku `dane.txt` znajduje sie 20 liczb calkowitych (kazda w osobnym wierszu). Fragment ciagu nazywamy "niemalejacy" jesli `T[i] <= T[i+1]` dla kazdej pary sasiadow w fragmencie (dopuszczamy rownosc). Znajdz najdluzszy spojny podciag niemalejacy.

**Dane** (`dane.txt`):
```
3
3
5
5
5
8
10
10
12
7
4
4
6
8
8
8
11
2
9
9
```

**Oczekiwany wynik**:
```
Najdluzszy ciag niemalejacy:
Pozycja poczatkowa: 1, dlugosc: 9
Elementy: 3 3 5 5 5 8 10 10 12
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Warunek niemalejacosci to `T[i] >= T[i-1]` (dozwolona rownosc).
2. **Podejscie**: Wzorzec current/max z `T[i] >= T[i-1]`.
3. **Kluczowy krok**: Niemalejacy = rosnacy + plateau. To laczy cwiczenia 12.1 i 12.2.

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
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    int maxDl = 1, maxStart = 0;
    int curDl = 1, curStart = 0;

    for (int i = 1; i < n; i++) {
        if (T[i] >= T[i - 1]) {
            curDl++;
        } else {
            if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }
            curDl = 1; curStart = i;
        }
    }
    if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }

    cout << "Najdluzszy ciag niemalejacy:" << endl;
    cout << "Pozycja poczatkowa: " << maxStart + 1 << ", dlugosc: " << maxDl << endl;
    cout << "Elementy: ";
    for (int i = maxStart; i < maxStart + maxDl; i++)
        cout << T[i] << " ";
    cout << endl;
    return 0;
}
```

**Wyjasnienie**: Jedyna zmiana wzgledem 12.2: warunek `>` zamieniony na `>=`. Dopuszczamy rowne elementy obok siebie.

Weryfikacja:
- Poz.1: 3,3,5,5,5,8,10,10,12 (dl.9) -> niemalejacy (kazdy >= poprzedniego)
- Poz.10: 7 (dl.1, bo 12>7 przerwa)
- Poz.11: 4,4,6,8,8,8,11 (dl.7, poz.11-17)
- Poz.18: 2,9,9 (dl.3)
Najdluzszy: dl.9 od poz.1.
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomylenie "niemalejacy" z "rosnacy"**: Niemalejacy dopuszcza rownosc (`>=`), rosnacy nie (`>`). CKE: -1 pkt
- **Brak aktualizacji po petli**: Jak we wszystkich cwiczeniach z wzorcem current/max. CKE: -2 pkt

</details>

---

### Cwiczenie 12.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4 (oscylacje)
**Tagi**: `oscylacja` `roznica-sasiadow` `skan-liniowy` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 15 liczb calkowitych (kazda w osobnym wierszu). Fragment ciagu nazywamy "oscylujacym" jesli roznice miedzy sasiadami naprzemiennie zmieniaja znak (tj. jesli T[i+1]-T[i] > 0, to T[i+2]-T[i+1] < 0, i na odwrot). Znajdz najdluzszy spojny podciag oscylujacy.

**Dane** (`dane.txt`):
```
5
10
3
12
1
15
8
20
6
25
7
30
4
35
2
```

**Oczekiwany wynik**:
```
Najdluzszy ciag oscylujacy:
Pozycja poczatkowa: 1, dlugosc: 15
Elementy: 5 10 3 12 1 15 8 20 6 25 7 30 4 35 2
```

Korekta — sprawdzmy: roznice: +5,-7,+9,-11,+14,-7,+12,-14,+19,-18,+23,-26,+31,-33
Znaki: +,-,+,-,+,-,+,-,+,-,+,-,+,-
Zawsze naprzemienne -> caly ciag jest oscylujacy.

Zmienmy dane aby nie byl caly ciag:

**Dane** (`dane.txt`) — skorygowane:
```
5
10
3
12
1
15
20
25
8
30
7
35
4
2
40
```

**Oczekiwany wynik** (skorygowany):
```
Roznice: +5,-7,+9,-11,+14,+5,+5,-17,+22,-23,+28,-31,-2,+38
Znaki:   +  -  +  -   +   +  +   -   +   -   +   -   -  +
Oscylacje: poz.1-6 (5,10,3,12,1,15) dl.6, potem przerwa (+,+)
           poz.8-13 (25,8,30,7,35,4) dl.6
           poz.13-15 (4,2,40) -> -,+ dl.3

Najdluzszy ciag oscylujacy:
Pozycja poczatkowa: 1, dlugosc: 6
Elementy: 5 10 3 12 1 15
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Oblicz roznice `d[i] = T[i+1] - T[i]`. Oscylacja to naprzemienne znaki roznic.
2. **Podejscie**: Sprawdz `d[i] * d[i-1] < 0` (rozne znaki). Wzorzec current/max.
3. **Kluczowy krok**: Ciag oscylujacy zaczyna sie od pary — minimalny ma dlugosc 2 (jedna roznica).

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
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    // Oblicz roznice
    vector<int> d(n - 1);
    for (int i = 0; i < n - 1; i++)
        d[i] = T[i + 1] - T[i];

    int maxDl = 2, maxStart = 0;
    int curDl = 2, curStart = 0;

    for (int i = 1; i < (int)d.size(); i++) {
        if ((long long)d[i] * d[i - 1] < 0) {
            curDl++;
        } else {
            if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }
            curDl = 2; curStart = i;
        }
    }
    if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }

    cout << "Najdluzszy ciag oscylujacy:" << endl;
    cout << "Pozycja poczatkowa: " << maxStart + 1 << ", dlugosc: " << maxDl << endl;
    cout << "Elementy: ";
    for (int i = maxStart; i < maxStart + maxDl; i++)
        cout << T[i] << " ";
    cout << endl;
    return 0;
}
```

**Wyjasnienie**: Obliczamy roznice sasiadow. Oscylacja to naprzemiennosc znakow: `d[i] * d[i-1] < 0`. Wzorzec current/max na tablicy roznic, ale dlugosc liczymy w elementach (= roznice + 1).

Weryfikacja (dane skorygowane: 5,10,3,12,1,15,20,25,8,30,7,35,4,2,40):
- Roznice: +5,-7,+9,-11,+14,+5,+5,-17,+22,-23,+28,-31,-2,+38
- Naprzemienne: d[0]*d[1]=-35<0 TAK, d[1]*d[2]=-63<0 TAK, d[2]*d[3]=-99<0 TAK, d[3]*d[4]=-154<0 TAK, d[4]*d[5]=70>0 NIE
- Poz.1-6: dl.6. Poz.8-12: 25,8,30,7,35 dl.5? Sprawdzmy: d[7]=-17, d[8]=+22, d[9]=-23, d[10]=+28, d[11]=-31 -> naprzemienne, dl.6 (poz.8-13).
Najdluzsze: dl.6.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak `(long long)` w mnozeniu**: `d[i] * d[i-1]` moze przekroczyc zakres int. CKE: -1 pkt (UB)
- **Roznica == 0 nie jest ani + ani -**: Jesli `T[i+1] == T[i]`, to `d[i] = 0` i `d[i]*d[i-1] = 0`, co nie jest < 0. Oscylacja sie przerywa — to poprawne zachowanie.
- **Dlugosc w roznicach zamiast elementach**: n roznic = n+1 elementow. CKE: -1 pkt

</details>

---

### Cwiczenie 12.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4 (ciag arytmetyczny)
**Tagi**: `ciag-stalej-roznicy` `roznica-sasiadow` `current-max` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 15 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory znajdzie najdluzszy spojny podciag bedacy ciagiem arytmetycznym (stala roznica miedzy sasiadami). Podaj pozycje, dlugosc, elementy i roznice ciagu.

**Dane** (`dane.txt`):
```
5
8
11
14
17
20
3
7
11
15
12
10
8
6
4
```

**Oczekiwany wynik**:
```
Najdluzszy ciag arytmetyczny:
Pozycja poczatkowa: 1, dlugosc: 6
Elementy: 5 8 11 14 17 20
Roznica: 3
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Ciag arytmetyczny ma stala roznice: `d[i] == d[i-1]` dla kazdej pary sasiadnich roznic.
2. **Podejscie**: Oblicz roznice, potem szukaj najdluzszego ciagu rownych roznic (wzorzec z 12.1!).
3. **Kluczowy krok**: Dlugosc ciagu arytmetycznego = dlugosc ciagu rownych roznic + 1.

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
    vector<int> T;
    int x;
    while (plik >> x) T.push_back(x);
    int n = T.size();

    if (n < 2) { cout << "Za malo danych" << endl; return 0; }

    vector<int> d(n - 1);
    for (int i = 0; i < n - 1; i++)
        d[i] = T[i + 1] - T[i];

    // Najdluzszy ciag rownych roznic
    int maxDl = 1, maxStart = 0;
    int curDl = 1, curStart = 0;

    for (int i = 1; i < (int)d.size(); i++) {
        if (d[i] == d[i - 1]) {
            curDl++;
        } else {
            if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }
            curDl = 1; curStart = i;
        }
    }
    if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }

    // Ciag arytmetyczny ma dlugosc = ciag roznic + 1
    int arDl = maxDl + 1;

    cout << "Najdluzszy ciag arytmetyczny:" << endl;
    cout << "Pozycja poczatkowa: " << maxStart + 1 << ", dlugosc: " << arDl << endl;
    cout << "Elementy: ";
    for (int i = maxStart; i < maxStart + arDl; i++)
        cout << T[i] << " ";
    cout << endl;
    cout << "Roznica: " << d[maxStart] << endl;
    return 0;
}
```

**Wyjasnienie**: Ciag arytmetyczny to ciag o stalej roznicy. Sprowadzamy problem do szukania najdluzszego ciagu rownych w tablicy roznic (cwiczenie 12.1). Dlugosc ciagu arytmetycznego = dlugosc ciagu rownych roznic + 1.

Weryfikacja:
- Roznice: 3,3,3,3,3,(-17),4,4,4,(-3),(-2),(-2),(-2),(-2)
- Ciagi rownych roznic: 3,3,3,3,3 (dl.5, poz.1), 4,4,4 (dl.3, poz.7), -2,-2,-2,-2 (dl.4, poz.11)
- Najdluzszy: dl.5 (roznice) -> dl.6 (elementy) od poz.1: 5,8,11,14,17,20, roznica=3
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o +1 przy przeliczaniu dlugosci**: n roznic = n+1 elementow. CKE: -1 pkt
- **Brak obslugi ciagu 2-elementowego**: Kazda para jest ciagiem arytmetycznym. Minimalny ciag arytmetyczny ma dlugosc 2.

</details>

---

### Cwiczenie 12.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4 + 2019 zad. 4 (zlozony)
**Tagi**: `spojny-podciag` `NWD-Euklidesa` `skan-liniowy` `wczytywanie-pliku` `gorka`

W pliku `dane.txt` znajduje sie 20 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory:
a) Znajdzie najdluzszy spojny podciag, w ktorym kazdy element jest podzielny przez 3.
b) Znajdzie najdluzsza "dolinke" (najpierw scisle maleje, potem scisle rosnie — odwrotnosc gorki). Minimalna dlg to 3.
c) Znajdzie najdluzszy spojny podciag, w ktorym bezwzgledna roznica miedzy kazdymi dwoma sasiadami jest <= 5.

**Dane** (`dane.txt`):
```
12
9
15
6
21
18
7
3
6
12
24
5
8
4
2
6
10
14
18
22
```

**Oczekiwany wynik**:
```
a) Najdluzszy ciag podzielnych przez 3:
   Pozycja poczatkowa: 1, dlugosc: 6
   Elementy: 12 9 15 6 21 18

b) Najdluzsza dolinka:
   Pozycja poczatkowa: 11, dlugosc: 7
   Elementy: 24 5 8 4 2 6 10
   Minimum: 2 (pozycja 15)
   Korekta: 24,5 malejaco. 5,8 rosnaco -> zmiana trendu po 2 krokach.
   Szukajmy dokladniej: faza malejaca, potem rosnaca.
   Od poz.11: 24 -> 5 -> malejaco? NIE: 24>5 OK, ale 5<8 -> zmiana na rosnacej.
   Faza malejaca: 24,5 (dl.2). Faza rosnaca: 5,8 (dl.2). Dolinka dl.3 (24,5,8).
   Od poz.12: 5,8 rosnaco -> brak spadku na poczatku.
   Od poz.11: 24,5 spadek. 5,8 wzrost -> dolinka (24,5,8) dl.3.
   Od poz.13: 8,4,2 spadek. 2,6,10,14,18,22 wzrost -> dolinka dl.8 (8,4,2,6,10,14,18,22).

   Najdluzsza dolinka:
   Pozycja poczatkowa: 13, dlugosc: 8
   Elementy: 8 4 2 6 10 14 18 22
   Minimum: 2 (pozycja 15)

c) Najdluzszy ciag z |T[i]-T[i+1]| <= 5:
   Poz.1: |12-9|=3, |9-15|=6 > 5 -> dl.2
   Poz.2: |9-15|=6 > 5 -> dl.1
   Poz.3: |15-6|=9 > 5 -> dl.1
   Poz.4: |6-21|=15 > 5 -> dl.1
   Poz.5: |21-18|=3, |18-7|=11 > 5 -> dl.2
   Poz.7: |7-3|=4, |3-6|=3, |6-12|=6 > 5 -> dl.3
   Poz.14: |4-2|=2, |2-6|=4, |6-10|=4, |10-14|=4, |14-18|=4, |18-22|=4 -> dl.7

   Najdluzszy ciag z |roznica| <= 5:
   Pozycja poczatkowa: 14, dlugosc: 7
   Elementy: 4 2 6 10 14 18 22
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy niezalezne problemy, kazdy ze wzorcem current/max.
2. **Podejscie**: (a) Warunek: `T[i] % 3 == 0`. (b) Szukaj fazy malejacej, potem rosnacej (odwrotnosc gorki). (c) Warunek: `|T[i]-T[i+1]| <= 5`.
3. **Kluczowy krok**: (b) Dolinka to dokładna odwrotnosc gorki z cw. 12.3 — zamien > na < i < na > w warunkach faz.

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

    // a) Najdluzszy ciag podzielnych przez 3
    int maxA = 0, startA = 0, curA = 0, curStartA = 0;
    for (int i = 0; i < n; i++) {
        if (T[i] % 3 == 0) {
            if (curA == 0) curStartA = i;
            curA++;
        } else {
            if (curA > maxA) { maxA = curA; startA = curStartA; }
            curA = 0;
        }
    }
    if (curA > maxA) { maxA = curA; startA = curStartA; }

    cout << "a) Najdluzszy ciag podzielnych przez 3:" << endl;
    cout << "   Pozycja poczatkowa: " << startA + 1 << ", dlugosc: " << maxA << endl;
    cout << "   Elementy: ";
    for (int i = startA; i < startA + maxA; i++) cout << T[i] << " ";
    cout << endl;

    // b) Najdluzsza dolinka (spadek + wzrost)
    int bestDol = 0, bestDolStart = 0, bestDolMin = 0;
    for (int i = 0; i < n - 2; i++) {
        int j = i;
        while (j + 1 < n && T[j + 1] < T[j]) j++;
        if (j == i) continue;
        int valley = j;
        while (j + 1 < n && T[j + 1] > T[j]) j++;
        if (j == valley) continue;
        int len = j - i + 1;
        if (len > bestDol) {
            bestDol = len; bestDolStart = i; bestDolMin = valley;
        }
    }

    cout << endl << "b) Najdluzsza dolinka:" << endl;
    cout << "   Pozycja poczatkowa: " << bestDolStart + 1
         << ", dlugosc: " << bestDol << endl;
    cout << "   Elementy: ";
    for (int i = bestDolStart; i < bestDolStart + bestDol; i++)
        cout << T[i] << " ";
    cout << endl;
    cout << "   Minimum: " << T[bestDolMin]
         << " (pozycja " << bestDolMin + 1 << ")" << endl;

    // c) Najdluzszy ciag z |roznica sasiada| <= 5
    int maxC = 1, startC = 0, curC = 1, curStartC = 0;
    for (int i = 0; i < n - 1; i++) {
        if (abs(T[i + 1] - T[i]) <= 5) {
            curC++;
        } else {
            if (curC > maxC) { maxC = curC; startC = curStartC; }
            curC = 1; curStartC = i + 1;
        }
    }
    if (curC > maxC) { maxC = curC; startC = curStartC; }

    cout << endl << "c) Najdluzszy ciag z |roznica| <= 5:" << endl;
    cout << "   Pozycja poczatkowa: " << startC + 1 << ", dlugosc: " << maxC << endl;
    cout << "   Elementy: ";
    for (int i = startC; i < startC + maxC; i++) cout << T[i] << " ";
    cout << endl;
    return 0;
}
```

**Wyjasnienie**: Trzy niezalezne wyszukiwania, kazde ze wzorcem current/max. (a) Ciag elementow spełniajacych warunek. (b) Odwrotnosc gorki — faza malejaca, potem rosnaca. (c) Ciag par spelniajacych warunek na roznicy.

Weryfikacja:
a) Podzielne przez 3: 12,9,15,6,21,18 (poz.1-6, dl.6). Nastepne: 7 -> nie. 3,6,12,24 (poz.8-11, dl.4). Najdluzszy: dl.6 od poz.1.
b) Dolinka od poz.13: 8 > 4 > 2 < 6 < 10 < 14 < 18 < 22. Spadek: 8,4,2 (dl.3). Wzrost: 2,6,10,14,18,22 (dl.6). Dolinka dl.8 (poz.13-20).
c) Od poz.14: |4-2|=2, |2-6|=4, |6-10|=4, |10-14|=4, |14-18|=4, |18-22|=4 -> dl.7.
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomylenie gorki z dolinka**: Gorka = wzrost+spadek, dolinka = spadek+wzrost. CKE: -2 pkt
- **Warunek `<= 5` zamiast `< 5`**: Zalezy od tresci — "nie przekracza 5" to `<= 5`, "mniejsza niz 5" to `< 5`. CKE: -1 pkt
- **Brak abs() w punkcie (c)**: Bez wartosci bezwzglednej, ujemne roznice (spadek) nie sa sprawdzane. CKE: -1 pkt

</details>

---

## Samoocena

| Poziom | Opis | Kryteria |
|--------|------|----------|
| Podstawowy | Rozumiem wzorzec current/max | Cwiczenia 1-2, 6 bez pomocy |
| Dobry | Radze sobie z gorkami i ciagami arytmetycznymi | Cwiczenia 3, 7, 9 bez pomocy |
| Bardzo dobry | Umiem NWD narastajaco i sumy prefiksowe | Cwiczenia 4-5, 8 bez pomocy |
| Doskonaly | Radze sobie ze zlozonym przetwarzaniem sekwencji | Cwiczenie 10 bez pomocy |

**Co dalej?**
- Jesli masz problem z wzorcem current/max -> cwicz 12.1, 12.2, 12.6 az beda automatyczne
- Jesli chcesz cwiczycNWD -> patrz `07_cyfry_liczby.md` cwiczenia z NWD
- Jesli chcesz cwiczycsumy prefiksowe -> patrz `cheatsheet_cpp.md` sekcja "Sumy prefiksowe"
