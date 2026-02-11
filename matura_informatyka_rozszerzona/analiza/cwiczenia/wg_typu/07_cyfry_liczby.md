# 07. Przetwarzanie cyfr i liczb

Typ zadania: **cyfry_liczby**
Czestotliwosc: 6/11 lat | Laczna punktacja: 36 pkt
Kategoria: IMPLEMENTACJA

## Umiejetnosci cwiczone w tym zestawie

`cyfry-mod-div` `suma-cyfr` `NWD-Euklidesa` `test-pierwszosci` `faktoryzacja` `wczytywanie-pliku` `filtrowanie` `systemy-liczbowe` `potegi` `odwracanie-liczby` `NWW` `podzielnosc`

---

### Cwiczenie 7.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2015 zad. 4 (cyfry liczb)
**Tagi**: `cyfry-mod-div` `suma-cyfr` `wczytywanie-pliku` `filtrowanie`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Zastanow sie, jak wyodrebnic poszczegolne cyfry z liczby.
2. **Podejscie**: Uzyj petli `while(n > 0)` z operacjami `n % 10` (ostatnia cyfra) i `n /= 10` (usuniecie ostatniej).
3. **Kluczowy krok**: Po obliczeniu sumy cyfr sprawdz parzystosc: `suma % 2 == 0`.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o n=0**: Petla `while(n > 0)` pomija liczbe 0 (suma cyfr = 0, a nie pusta). Dla tego zadania to OK (liczby dodatnie), ale warto pamietac. CKE: -0 pkt (tu nie dotyczy)
- **Uzycie `int` zamiast wyodrebniania cyfr**: Proba operowania na string bez konwersji. CKE: -1 pkt za bledna metode
- **Brak inicjalizacji licznika `ile`**: Niezainicjalizowana zmienna daje losowy wynik. CKE: -1 pkt

</details>

---

### Cwiczenie 7.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4 (NWD)
**Tagi**: `NWD-Euklidesa` `wczytywanie-pliku` `filtrowanie` `podzielnosc`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Przypomnij sobie algorytm Euklidesa — zamieniaj pare (a, b) na (b, a mod b).
2. **Podejscie**: Petla `while(b != 0)` z zamiana: temp = b, b = a % b, a = temp.
3. **Kluczowy krok**: Para jest wzglednie pierwsza gdy NWD = 1. Zliczaj te z NWD > 1.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Zamiana a i b w zlej kolejnosci**: `a = b; b = a % b;` — po pierwszym przypisaniu stare `a` jest stracone. Trzeba uzyc zmiennej tymczasowej. CKE: -2 pkt
- **Warunek `b > 0` zamiast `b != 0`**: Dla liczb dodatnich to samo, ale dla ogolnosci `!= 0` jest bezpieczniejsze.
- **Porownanie NWD >= 1 zamiast > 1**: Kazda para ma NWD >= 1, wiec warunek bylby zawsze spelniony. CKE: -1 pkt

</details>

---

### Cwiczenie 7.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4.1 (liczby pierwsze)
**Tagi**: `test-pierwszosci` `wczytywanie-pliku` `filtrowanie` `podzielnosc`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Test pierwszosci wymaga sprawdzenia podzielnosci do sqrt(n).
2. **Podejscie**: Petla `for(i = 2; i*i <= n; i++)` — jesli znajdziesz dzielnik, liczba jest zlozona.
3. **Kluczowy krok**: Optymalizacja — sprawdz 2 osobno, potem tylko nieparzyste dzielniki `for(i = 3; i*i <= n; i += 2)`.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Warunek `i < n` zamiast `i*i <= n`**: Program dziala poprawnie ale jest O(n) zamiast O(sqrt(n)). CKE: -0 pkt (ale wolniejszy)
- **Brak obslugi n=2**: Jesli petla zaczyna od i=3, to n=2 nie jest sprawdzane. CKE: -1 pkt
- **Brak `return n` w `najmniejszyDzielnik`**: Dla liczby pierwszej (nie powinno sie wywolac) petla nie znajdzie dzielnika. CKE: -1 pkt

</details>

---

### Cwiczenie 7.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 4.2 (faktoryzacja)
**Tagi**: `faktoryzacja` `cyfry-mod-div` `wczytywanie-pliku` `podzielnosc`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Faktoryzacja probna — dziel przez kolejne liczby od 2.
2. **Podejscie**: Zagniezdzona petla: zewnetrzna po d od 2, wewnetrzna `while(n % d == 0)`. Kontynuuj do `d*d > n`.
3. **Kluczowy krok**: Po petli jesli `n > 1`, to n jest ostatnim czynnikiem pierwszym. Nie zapomnij o tym!

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

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o `if (n > 1)` na koncu**: Jesli po petli n > 1, ostatni czynnik pierwszy jest pomijany. Np. 15 = 3*5, ale bez tego warunku dostaniesz tylko 3. CKE: -2 pkt
- **Warunek `d <= n` zamiast `d*d <= n`**: Poprawne ale O(n) zamiast O(sqrt(n)). Na maturze dla malych danych nie przeszkadza.
- **Brak zapisu oryginalnej wartosci**: Zmienna n jest modyfikowana w petli — trzeba ja skopiowac przed faktoryzacja. CKE: -1 pkt

</details>

---

### Cwiczenie 7.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3.3 (skrot liczby + NWD)
**Tagi**: `cyfry-mod-div` `NWD-Euklidesa` `wczytywanie-pliku` `filtrowanie`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Rozloz zadanie na etapy — najpierw ekstrakcja cyfr nieparzystych, potem NWD.
2. **Podejscie**: Konwertuj liczbe na string, iteruj po znakach, wyciagaj nieparzyste cyfry do nowego stringa, konwertuj z powrotem na int.
3. **Kluczowy krok**: Pamietaj o przypadku brzegowym — gdy nie ma cyfr nieparzystych, skrot = 0. NWD(x, 0) jest niezdefiniowane, wiec pomin taka pare.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Ekstrakcja cyfr w odwrotnej kolejnosci**: Uzywajac mod/div cyfry wychodza od konca (7,3 zamiast 3,7 dla 24837). Uzyj `to_string` lub odwroc wynik. CKE: -2 pkt
- **Brak obslugi skrotu = 0**: NWD(x, 0) moze powodowac dzielenie przez zero lub nieskonczona petle. CKE: -1 pkt
- **Uzycie `atoi` zamiast `stoi`**: `atoi` nie sygnalizuje bledu. Dla pustego stringa lepiej sprawdzic recznie. CKE: -0 pkt (dziala, ale zla praktyka)

</details>

---

### Cwiczenie 7.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3 (cyfry liczby pi)
**Tagi**: `cyfry-mod-div` `wczytywanie-pliku` `filtrowanie`

W pliku `dane.txt` znajduje sie 10 liczb calkowitych dodatnich (kazda w osobnym wierszu). Napisz program, ktory dla kazdej liczby:
a) Poda jej liczbe cyfr.
b) Poda jej pierwsza i ostatnia cyfre.
c) Sprawdzi, czy pierwsza cyfra jest wieksza od ostatniej.

**Dane** (`dane.txt`):
```
7245
918
36470
5
84213
609
12
47856
301
6783
```

**Oczekiwany wynik**:
```
7245: 4 cyfry, pierwsza=7, ostatnia=5, pierwsza > ostatnia: TAK
918: 3 cyfry, pierwsza=9, ostatnia=8, pierwsza > ostatnia: TAK
36470: 5 cyfr, pierwsza=3, ostatnia=0, pierwsza > ostatnia: TAK
5: 1 cyfra, pierwsza=5, ostatnia=5, pierwsza > ostatnia: NIE
84213: 5 cyfr, pierwsza=8, ostatnia=3, pierwsza > ostatnia: TAK
609: 3 cyfry, pierwsza=6, ostatnia=9, pierwsza > ostatnia: NIE
12: 2 cyfry, pierwsza=1, ostatnia=2, pierwsza > ostatnia: NIE
47856: 5 cyfr, pierwsza=4, ostatnia=6, pierwsza > ostatnia: NIE
301: 3 cyfry, pierwsza=3, ostatnia=1, pierwsza > ostatnia: TAK
6783: 4 cyfry, pierwsza=6, ostatnia=3, pierwsza > ostatnia: TAK
Liczby z pierwsza > ostatnia: 6
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Ostatnia cyfra to `n % 10`. Pierwsza cyfra to wynik dzielenia przez 10 dopoki n >= 10.
2. **Podejscie**: Liczba cyfr to dlugosc `to_string(n)` lub zliczanie w petli `while(n > 0)`.
3. **Kluczowy krok**: Pamietaj, ze dla jednocyfrowej liczby pierwsza == ostatnia.

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
    int n;
    int ile = 0;
    while (plik >> n) {
        string s = to_string(n);
        int liczbaCyfr = s.length();
        int pierwsza = s[0] - '0';
        int ostatnia = s[s.length() - 1] - '0';
        cout << n << ": " << liczbaCyfr << " cyfr";
        if (liczbaCyfr == 1) cout << "a";
        else cout << (liczbaCyfr < 5 ? "y" : "");
        cout << ", pierwsza=" << pierwsza << ", ostatnia=" << ostatnia;
        cout << ", pierwsza > ostatnia: ";
        if (pierwsza > ostatnia) {
            cout << "TAK" << endl;
            ile++;
        } else {
            cout << "NIE" << endl;
        }
    }
    cout << "Liczby z pierwsza > ostatnia: " << ile << endl;
    return 0;
}
```

Weryfikacja:
- 7245: 4 cyfry, 7 > 5 -> TAK
- 918: 3 cyfry, 9 > 8 -> TAK
- 36470: 5 cyfr, 3 > 0 -> TAK
- 5: 1 cyfra, 5 > 5 -> NIE (nie scisle)
- 84213: 5 cyfr, 8 > 3 -> TAK
- 609: 3 cyfry, 6 > 9 -> NIE
- 12: 2 cyfry, 1 > 2 -> NIE
- 47856: 5 cyfr, 4 > 6 -> NIE
- 301: 3 cyfry, 3 > 1 -> TAK
- 6783: 4 cyfry, 6 > 3 -> TAK
Razem TAK: 6
</details>

<details>
<summary>Typowe bledy</summary>

- **Wyciaganie pierwszej cyfry petla dzielenia — zapomnienie o kopii**: Petla `while(n >= 10) n /= 10` niszczy oryginalna wartosc. CKE: -1 pkt
- **Uzywanie `n / pow(10, k)` z double**: Bledy zaokraglen float mogą dac zla cyfre. CKE: -1 pkt

</details>

---

### Cwiczenie 7.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3 (konwersja bin/dec)
**Tagi**: `systemy-liczbowe` `cyfry-mod-div` `potegi` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 8 liczb calkowitych dodatnich (kazda w osobnym wierszu, wartosci z zakresu 1-255). Napisz program, ktory:
a) Wypisze kazda liczbe w systemie dwojkowym (bez zer wiodacych).
b) Poda ile jedynek ma jej zapis binarny.
c) Znajdzie liczbe z najwieksza liczba jedynek w zapisie binarnym.

**Dane** (`dane.txt`):
```
42
255
13
100
7
128
63
200
```

**Oczekiwany wynik**:
```
42 -> 101010 (jedynki: 3)
255 -> 11111111 (jedynki: 8)
13 -> 1101 (jedynki: 3)
100 -> 1100100 (jedynki: 3)
7 -> 111 (jedynki: 3)
128 -> 10000000 (jedynki: 1)
63 -> 111111 (jedynki: 6)
200 -> 11001000 (jedynki: 3)
Najwiecej jedynek: 255 (8 jedynek)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Konwersja na system dwojkowy to powtarzane dzielenie przez 2, reszty od konca.
2. **Podejscie**: `while(n > 0) { reszty += (n%2); n /= 2; }`, a potem odwroc string reszt.
3. **Kluczowy krok**: Zliczaj jedynki (reszty rowne 1) w trakcie konwersji.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <algorithm>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    int n;
    int maxJedynki = 0, maxLiczba = 0;

    while (plik >> n) {
        int oryg = n;
        string bin = "";
        int jedynki = 0;
        int tmp = n;
        while (tmp > 0) {
            int bit = tmp % 2;
            bin += to_string(bit);
            if (bit == 1) jedynki++;
            tmp /= 2;
        }
        reverse(bin.begin(), bin.end());
        cout << oryg << " -> " << bin << " (jedynki: " << jedynki << ")" << endl;
        if (jedynki > maxJedynki) {
            maxJedynki = jedynki;
            maxLiczba = oryg;
        }
    }
    cout << "Najwiecej jedynek: " << maxLiczba << " (" << maxJedynki << " jedynek)" << endl;
    return 0;
}
```

Weryfikacja:
- 42: 42/2=21 r0, 21/2=10 r1, 10/2=5 r0, 5/2=2 r1, 2/2=1 r0, 1/2=0 r1 -> 101010, jedynki=3
- 255: 11111111, jedynki=8
- 13: 1101, jedynki=3
- 100: 1100100, jedynki=3
- 7: 111, jedynki=3
- 128: 10000000, jedynki=1
- 63: 111111, jedynki=6
- 200: 11001000, jedynki=3
Max: 255 (8)
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o odwroceniu wyniku**: Reszty z dzielenia sa od konca. Bez `reverse` dostaniesz odwrocony zapis. CKE: -2 pkt
- **Brak obslugi n=0**: `while(n > 0)` daje pusty string dla 0. Dodaj warunek specjalny. CKE: -1 pkt (tu nie dotyczy, n >= 1)
- **Uzycie float/double przy dzieleniu**: `n / 2.0` daje double, traci dokladnosc. Uzywaj `n / 2` (int). CKE: -1 pkt

</details>

---

### Cwiczenie 7.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 4 (palindromy liczbowe)
**Tagi**: `cyfry-mod-div` `odwracanie-liczby` `test-pierwszosci` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 12 liczb calkowitych wiekszych od 10 (kazda w osobnym wierszu). Napisz program, ktory:
a) Sprawdzi, ktore z nich sa palindromami (czytane od lewej i prawej daja ta sama liczbe).
b) Sposrod palindromow znajdzie te, ktore sa jednoczesnie liczbami pierwszymi.

**Dane** (`dane.txt`):
```
121
131
245
1001
353
789
11
4884
929
1234
101
500
```

**Oczekiwany wynik**:
```
a) Palindromy liczbowe:
   121 (odwrocone: 121) - palindrom
   131 (odwrocone: 131) - palindrom
   1001 (odwrocone: 1001) - palindrom
   353 (odwrocone: 353) - palindrom
   11 (odwrocone: 11) - palindrom
   4884 (odwrocone: 4884) - palindrom
   929 (odwrocone: 929) - palindrom
   101 (odwrocone: 101) - palindrom
   Ilosc palindromow: 8

b) Palindromy pierwsze:
   131 - palindrom i pierwsza
   353 - palindrom i pierwsza
   11 - palindrom i pierwsza
   929 - palindrom i pierwsza
   101 - palindrom i pierwsza
   Ilosc: 5
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Odwroc liczbe cyframi (mod/div) i porownaj z oryginalem.
2. **Podejscie**: `while(tmp > 0) { odwr = odwr * 10 + tmp % 10; tmp /= 10; }`. Jesli odwr == oryg, to palindrom.
3. **Kluczowy krok**: Polacz dwa testy — palindrom i pierwszosc — w jednej petli.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int odwroc(int n) {
    int wynik = 0;
    while (n > 0) {
        wynik = wynik * 10 + n % 10;
        n /= 10;
    }
    return wynik;
}

bool czyPierwsza(int n) {
    if (n < 2) return false;
    if (n == 2) return true;
    if (n % 2 == 0) return false;
    for (int i = 3; i * i <= n; i += 2)
        if (n % i == 0) return false;
    return true;
}

int main() {
    ifstream plik("dane.txt");
    int n;
    vector<int> tab;
    while (plik >> n) tab.push_back(n);

    cout << "a) Palindromy liczbowe:" << endl;
    vector<int> palindromy;
    for (int x : tab) {
        int odwr = odwroc(x);
        if (odwr == x) {
            cout << "   " << x << " (odwrocone: " << odwr << ") - palindrom" << endl;
            palindromy.push_back(x);
        }
    }
    cout << "   Ilosc palindromow: " << palindromy.size() << endl;

    cout << endl << "b) Palindromy pierwsze:" << endl;
    int ile = 0;
    for (int x : palindromy) {
        if (czyPierwsza(x)) {
            cout << "   " << x << " - palindrom i pierwsza" << endl;
            ile++;
        }
    }
    cout << "   Ilosc: " << ile << endl;
    return 0;
}
```

Weryfikacja:
- 121: odwrocone 121 -> palindrom. 121 = 11*11 -> nie pierwsza
- 131: palindrom. 131 -> sprawdzamy: 131%2!=0, 131%3!=0, ..., 11*11=121<131, 11*12>131 -> pierwsza
- 245: odwrocone 542 -> nie palindrom
- 1001: palindrom. 1001 = 7*143 = 7*11*13 -> nie pierwsza
- 353: palindrom. 353 -> 353%2!=0, 353%3!=0, ..., 18*18=324<353, 19*19=361>353 -> pierwsza
- 789: odwrocone 987 -> nie palindrom
- 11: palindrom. 11 -> pierwsza
- 4884: palindrom. 4884 = 2*2442 -> nie pierwsza
- 929: palindrom. 929 -> sprawdzamy do sqrt(929)~30: nie dzieli sie -> pierwsza
- 1234: odwrocone 4321 -> nie palindrom
- 101: palindrom. 101 -> pierwsza
- 500: odwrocone 5 (005) -> nie palindrom
</details>

<details>
<summary>Typowe bledy</summary>

- **Porownywanie stringow zamiast liczb**: String "500" odwrocony to "005" co jako string != "500", ale `stoi("005") == 5 != 500`. Obie metody dzialaja, ale trzeba byc konsekwentnym. CKE: -0 pkt
- **Zapomnienie ze 1 nie jest pierwsza**: `czyPierwsza(1)` powinno zwracac false. CKE: -1 pkt
- **Overflow przy odwracaniu duzych liczb**: Dla 10-cyfrowych liczb odwrocona wartosc moze przekroczyc `int`. Tu nie dotyczy (male dane). CKE: -1 pkt na duzych danych

</details>

---

### Cwiczenie 7.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4 + 2022 zad. 4 (NWW i NWD)
**Tagi**: `NWD-Euklidesa` `NWW` `faktoryzacja` `wczytywanie-pliku`

W pliku `trojki.txt` znajduje sie 8 trojek liczb calkowitych dodatnich (kazda trojka w osobnym wierszu, liczby oddzielone spacjami). Napisz program, ktory:
a) Dla kazdej trojki (a, b, c) obliczy NWD(a, b, c) i NWW(a, b, c).
b) Zliczy ile trojek spelnia warunek: NWD(a,b,c) * NWW(a,b,c) == a * b * c (to zachodzi tylko dla par, nie trojek — sprawdz ktore trojki nie spelniaja).

Przypomnienie: NWD(a,b,c) = NWD(NWD(a,b), c), NWW(a,b) = a*b / NWD(a,b), NWW(a,b,c) = NWW(NWW(a,b), c).

**Dane** (`trojki.txt`):
```
12 18 24
5 7 11
6 10 15
8 12 16
3 9 27
14 21 35
4 6 8
20 30 50
```

**Oczekiwany wynik**:
```
12 18 24: NWD=6, NWW=72, NWD*NWW=432, a*b*c=5184, rowne: NIE
5 7 11: NWD=1, NWW=385, NWD*NWW=385, a*b*c=385, rowne: TAK
6 10 15: NWD=1, NWW=30, NWD*NWW=30, a*b*c=900, rowne: NIE
8 12 16: NWD=4, NWW=48, NWD*NWW=192, a*b*c=1536, rowne: NIE
3 9 27: NWD=3, NWW=27, NWD*NWW=81, a*b*c=729, rowne: NIE
14 21 35: NWD=7, NWW=210, NWD*NWW=1470, a*b*c=10290, rowne: NIE
4 6 8: NWD=2, NWW=24, NWD*NWW=48, a*b*c=192, rowne: NIE
20 30 50: NWD=10, NWW=300, NWD*NWW=3000, a*b*c=30000, rowne: NIE
Trojki spelniajace NWD*NWW == a*b*c: 1
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: NWD trzech liczb to NWD(NWD(a,b), c). NWW trzech to NWW(NWW(a,b), c).
2. **Podejscie**: NWW(a,b) = a / NWD(a,b) * b (dziel przed mnozeniem, zeby uniknac overflow).
3. **Kluczowy krok**: Wlasnosc NWD*NWW = a*b zachodzi TYLKO dla par. Dla trojek to nie jest regula — jedyny przypadek to gdy wszystkie trzy sa parami wzglednie pierwsze.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int nwd(int a, int b) {
    while (b != 0) { int t = b; b = a % b; a = t; }
    return a;
}

int nww(int a, int b) {
    return a / nwd(a, b) * b;
}

int main() {
    ifstream plik("trojki.txt");
    int a, b, c;
    int ile = 0;

    while (plik >> a >> b >> c) {
        int g = nwd(nwd(a, b), c);
        int l = nww(nww(a, b), c);
        long long prod = (long long)a * b * c;
        long long gl = (long long)g * l;
        cout << a << " " << b << " " << c << ": NWD=" << g << ", NWW=" << l
             << ", NWD*NWW=" << gl << ", a*b*c=" << prod
             << ", rowne: " << (gl == prod ? "TAK" : "NIE") << endl;
        if (gl == prod) ile++;
    }
    cout << "Trojki spelniajace NWD*NWW == a*b*c: " << ile << endl;
    return 0;
}
```

Weryfikacja:
- 12,18,24: NWD(12,18)=6, NWD(6,24)=6. NWW(12,18)=36, NWW(36,24)=72. 6*72=432, 12*18*24=5184 -> NIE
- 5,7,11: NWD=1, NWW(5,7)=35, NWW(35,11)=385. 1*385=385, 5*7*11=385 -> TAK (parami wzglednie pierwsze)
- 6,10,15: NWD(6,10)=2, NWD(2,15)=1. NWW(6,10)=30, NWW(30,15)=30. 1*30=30, 6*10*15=900 -> NIE
- 8,12,16: NWD(8,12)=4, NWD(4,16)=4. NWW(8,12)=24, NWW(24,16)=48. 4*48=192, 8*12*16=1536 -> NIE
- 3,9,27: NWD=3, NWW(3,9)=9, NWW(9,27)=27. 3*27=81, 3*9*27=729 -> NIE
- 14,21,35: NWD(14,21)=7, NWD(7,35)=7. NWW(14,21)=42, NWW(42,35)=210. 7*210=1470, 14*21*35=10290 -> NIE
- 4,6,8: NWD=2, NWW(4,6)=12, NWW(12,8)=24. 2*24=48, 4*6*8=192 -> NIE
- 20,30,50: NWD=10, NWW(20,30)=60, NWW(60,50)=300. 10*300=3000, 20*30*50=30000 -> NIE
</details>

<details>
<summary>Typowe bledy</summary>

- **Overflow przy a*b*c**: Iloczyn trzech liczb moze przekroczyc `int`. Uzyj `long long`. CKE: -1 pkt
- **NWW obliczane jako a*b/NWD**: Przy duzych a*b moze byc overflow. Lepiej `a / NWD(a,b) * b`. CKE: -1 pkt
- **Bledne NWW trojki**: NWW(a,b,c) != a*b*c / NWD(a,b,c). Trzeba obliczac narastajaco. CKE: -2 pkt

</details>

---

### Cwiczenie 7.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4 + 2015 zad. 4 (zlozony problem cyfrowy)
**Tagi**: `cyfry-mod-div` `suma-cyfr` `test-pierwszosci` `potegi` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 12 liczb calkowitych dodatnich z zakresu 10-99999 (kazda w osobnym wierszu). Definiujemy "wartosc cyfrowa" liczby jako sume kwadratow jej cyfr. Na przyklad wartosc cyfrowa 123 = 1^2 + 2^2 + 3^2 = 14.

Napisz program, ktory:
a) Dla kazdej liczby obliczy jej wartosc cyfrowa.
b) Znajdzie wszystkie liczby, ktorych wartosc cyfrowa jest liczba pierwsza.
c) Znajdzie pare roznych liczb (i, j), i < j, dla ktorych wartosc cyfrowa jest taka sama (pierwsza znaleziona).

**Dane** (`dane.txt`):
```
123
321
456
789
100
999
47
74
256
652
333
811
```

**Oczekiwany wynik**:
```
a) Wartosci cyfrowe:
   123 -> 1+4+9 = 14
   321 -> 9+4+1 = 14
   456 -> 16+25+36 = 77
   789 -> 49+64+81 = 194
   100 -> 1+0+0 = 1
   999 -> 81+81+81 = 243
   47 -> 16+49 = 65
   74 -> 49+16 = 65
   256 -> 4+25+36 = 65
   652 -> 36+25+4 = 65
   333 -> 9+9+9 = 27
   811 -> 64+1+1 = 66

b) Liczby z pierwsza wartoscia cyfrowa:
   456 (wc=77) - 77 nie jest pierwsza (77=7*11)
   Poprawka - sprawdzamy:
   14: 14=2*7 -> NIE
   77: 77=7*11 -> NIE
   194: 194=2*97 -> NIE
   1: nie pierwsza -> NIE
   243: 243=3^5 -> NIE
   65: 65=5*13 -> NIE
   27: 27=3^3 -> NIE
   66: 66=2*3*11 -> NIE
   Ilosc: 0

c) Pierwsza para z ta sama wartoscia cyfrowa:
   123 i 321 (obie: wc=14)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wyodrębniaj cyfry (mod/div), podnoś kazda do kwadratu, sumuj.
2. **Podejscie**: Zapisz wartosci cyfrowe w tablicy. Porownuj parami (podwojna petla) lub uzyj mapy.
3. **Kluczowy krok**: Dla czesci (c) uzyj `map<int, int>` — klucz to wartosc cyfrowa, wartosc to indeks pierwszego wystapienia. Gdy znajdziesz drugi — masz pare.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <map>
using namespace std;

bool czyPierwsza(int n) {
    if (n < 2) return false;
    if (n == 2) return true;
    if (n % 2 == 0) return false;
    for (int i = 3; i * i <= n; i += 2)
        if (n % i == 0) return false;
    return true;
}

int wartoscCyfrowa(int n) {
    int suma = 0;
    while (n > 0) {
        int c = n % 10;
        suma += c * c;
        n /= 10;
    }
    return suma;
}

int main() {
    ifstream plik("dane.txt");
    vector<int> tab;
    int x;
    while (plik >> x) tab.push_back(x);

    vector<int> wc;
    cout << "a) Wartosci cyfrowe:" << endl;
    for (int n : tab) {
        int w = wartoscCyfrowa(n);
        wc.push_back(w);
        // Wypisz rozklad
        string s = to_string(n);
        cout << "   " << n << " -> ";
        for (int i = 0; i < s.length(); i++) {
            int c = s[i] - '0';
            if (i > 0) cout << "+";
            cout << c * c;
        }
        cout << " = " << w << endl;
    }

    cout << endl << "b) Liczby z pierwsza wartoscia cyfrowa:" << endl;
    int ileB = 0;
    for (int i = 0; i < tab.size(); i++) {
        if (czyPierwsza(wc[i])) {
            cout << "   " << tab[i] << " (wc=" << wc[i] << ")" << endl;
            ileB++;
        }
    }
    cout << "   Ilosc: " << ileB << endl;

    cout << endl << "c) Pierwsza para z ta sama wartoscia cyfrowa:" << endl;
    map<int, int> pierwsze;
    bool znaleziono = false;
    for (int i = 0; i < tab.size(); i++) {
        if (pierwsze.count(wc[i])) {
            int j = pierwsze[wc[i]];
            cout << "   " << tab[j] << " i " << tab[i]
                 << " (obie: wc=" << wc[i] << ")" << endl;
            znaleziono = true;
            break;
        }
        pierwsze[wc[i]] = i;
    }
    if (!znaleziono) cout << "   Brak par" << endl;
    return 0;
}
```

Weryfikacja:
- 123: 1+4+9=14, 321: 9+4+1=14 -> ta sama wc -> para (123, 321)
- 456: 16+25+36=77
- 789: 49+64+81=194
- 100: 1+0+0=1
- 999: 81+81+81=243
- 47: 16+49=65, 74: 49+16=65, 256: 4+25+36=65, 652: 36+25+4=65
- 333: 9+9+9=27
- 811: 64+1+1=66
Zadna wc nie jest pierwsza (14=2*7, 77=7*11, 194=2*97, 1 nie, 243=3^5, 65=5*13, 27=3^3, 66=2*3*11).
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o `c * c` — uzycie `c` zamiast kwadratu**: Daje sume cyfr zamiast sumy kwadratow. CKE: -2 pkt (zupelnie inny wynik)
- **Szukanie par podwojna petla O(n^2) zamiast mapy**: Poprawne ale wolne. Na maturze OK dla malych n. CKE: -0 pkt
- **Brak obslugi braku par**: Program moze sie zawiesic lub wypisac smieci. CKE: -1 pkt
- **Sprawdzanie `czyPierwsza(1)` jako true**: 1 nie jest pierwsza. CKE: -1 pkt

</details>

---

## Samoocena

| Poziom | Opis | Zakres cwiczen |
|--------|------|---------------|
| Podstawowy | Rozumiem mod/div, umiem wyodrebnic cyfry i obliczyc sume cyfr | 1-3 bez pomocy |
| Dobry | Sprawnie implementuje NWD, test pierwszosci, faktoryzacje | 4-6 bez pomocy |
| Bardzo dobry | Lacze wiele operacji cyfrowych, konwertuje systemy liczbowe | 7-8 bez pomocy |
| Doskonaly | Rozwiazuje zlozone problemy wieloetapowe z optymalizacja | 9-10 bez pomocy |

**Co dalej?**
- Jesli masz problemy z cwiczeniami 1-3, wrocz do `cheatsheet_cpp.md` — sekcja "Operacje na cyfrach"
- Jesli opanowales 1-6, przejdz do cwiczen `09_zlozone.md` (wieloetapowe przetwarzanie)
- Jesli rozwiazales 9-10, sprobuj cwiczen z `12_sekwencje.md` (zaawansowane wzorce)
