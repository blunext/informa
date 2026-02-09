# 08. Przetwarzanie napisow

Typ zadania: **napisy**
Czestotliwosc: 4/11 lat | Laczna punktacja: 25 pkt
Kategoria: IMPLEMENTACJA

---

### Cwiczenie 8.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2.1 (palindromy)

W pliku `napisy.txt` znajduje sie 10 napisow skladajacych sie z malych liter alfabetu lacinskiego (kazdy w osobnym wierszu). Napisz program, ktory wypisze te napisy, ktore sa palindromami (czytane od lewej i od prawej daja ten sam napis).

**Dane** (`napisy.txt`):
```
kajak
programowanie
abcba
level
kotek
racecar
anna
python
ala
rotor
```

**Oczekiwany wynik**:
```
kajak - palindrom
abcba - palindrom
level - palindrom
racecar - palindrom
anna - palindrom
ala - palindrom
rotor - palindrom
Liczba palindromow: 7
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
using namespace std;

bool czyPalindrom(string s) {
    int n = s.length();
    for (int i = 0; i < n / 2; i++) {
        if (s[i] != s[n - 1 - i]) return false;
    }
    return true;
}

int main() {
    ifstream plik("napisy.txt");
    string s;
    int ile = 0;
    while (plik >> s) {
        if (czyPalindrom(s)) {
            cout << s << " - palindrom" << endl;
            ile++;
        }
    }
    cout << "Liczba palindromow: " << ile << endl;
    return 0;
}
```

**Wyjasnienie**: Palindrom sprawdzamy porownujac znaki od poczatku i konca napisu, zbiegajac sie do srodka. Jesli wszystkie pary s[i] i s[n-1-i] sa rowne, napis jest palindromem.

Weryfikacja:
- kajak: k=k, a=a, j (srodek) -> palindrom
- programowanie: p!=e -> nie
- abcba: a=a, b=b, c (srodek) -> palindrom
- level: l=l, e=e, v (srodek) -> palindrom
- kotek: k=k, o!=e -> nie
- racecar: r=r, a=a, c=c, e (srodek) -> palindrom
- anna: a=a, n=n -> palindrom
- python: p!=n -> nie
- ala: a=a, l (srodek) -> palindrom
- rotor: r=r, o=o, t (srodek) -> palindrom
</details>

---

### Cwiczenie 8.2 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 5 (kody ASCII)

W pliku `napisy.txt` znajduje sie 8 napisow skladajacych sie z malych liter alfabetu lacinskiego (kazdy w osobnym wierszu). Napisz program, ktory dla kazdego napisu obliczy sume kodow ASCII wszystkich jego znakow, a nastepnie poda napis o najwiekszej i najmniejszej sumie kodow.

**Dane** (`napisy.txt`):
```
abc
xyz
hello
cat
zoo
bee
ax
wind
```

**Oczekiwany wynik**:
```
abc -> suma ASCII: 294
xyz -> suma ASCII: 363
hello -> suma ASCII: 532
cat -> suma ASCII: 312
zoo -> suma ASCII: 344
bee -> suma ASCII: 300
ax -> suma ASCII: 217
wind -> suma ASCII: 434
Najwieksza suma: hello (532)
Najmniejsza suma: ax (217)
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
    ifstream plik("napisy.txt");
    string s;
    string maxNapis, minNapis;
    int maxSuma = -1, minSuma = 1000000;

    while (plik >> s) {
        int suma = 0;
        for (char c : s) suma += (int)c;
        cout << s << " -> suma ASCII: " << suma << endl;
        if (suma > maxSuma) { maxSuma = suma; maxNapis = s; }
        if (suma < minSuma) { minSuma = suma; minNapis = s; }
    }
    cout << "Najwieksza suma: " << maxNapis << " (" << maxSuma << ")" << endl;
    cout << "Najmniejsza suma: " << minNapis << " (" << minSuma << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Suma kodow ASCII to suma wartosci liczbowych znakow (a=97, b=98, ..., z=122). Szukamy napisu o max i min sumie przegladajac wszystkie napisy.

Weryfikacja:
- abc: 97+98+99 = 294
- xyz: 120+121+122 = 363
- hello: 104+101+108+108+111 = 532
- cat: 99+97+116 = 312
- zoo: 122+111+111 = 344
- bee: 98+101+101 = 300
- ax: 97+120 = 217
- wind: 119+105+110+100 = 434

Max: hello (532), Min: ax (217).
</details>

---

### Cwiczenie 8.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 6 (szyfr Cezara)

W pliku `szyfr.txt` znajduje sie 5 zaszyfrowanych napisow (kazdy w osobnym wierszu). Napisy skladaja sie z duzych liter A-Z i zostaly zaszyfrowane szyfrem Cezara z podanym przesunieciem k (kazda litera tekstu jawnego zostala zastapiona litera o k pozycji dalej w alfabecie, cyklicznie). Kazdy wiersz zawiera liczbe k, a po spacji zaszyfrowany napis. Napisz program, ktory odszyfruje kazdy napis (przesuniecie w lewo o k pozycji).

**Dane** (`szyfr.txt`):
```
3 KHOOR
1 TFDSFU
7 JHZASL
5 YMJWJ
13 PNHFR
```

**Oczekiwany wynik**:
```
KHOOR (k=3) -> HELLO
TFDSFU (k=1) -> SECRET
JHZASL (k=7) -> CASTLE
YMJWJ (k=5) -> THERE
PNHFR (k=13) -> CAUSE
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
    ifstream plik("szyfr.txt");
    int k;
    string napis;
    while (plik >> k >> napis) {
        cout << napis << " (k=" << k << ") -> ";
        for (int i = 0; i < napis.length(); i++) {
            char c = (napis[i] - 'A' - k + 26) % 26 + 'A';
            cout << c;
        }
        cout << endl;
    }
    return 0;
}
```

**Wyjasnienie**: Deszyfrowanie Cezara: przesuwamy kazda litere o k pozycji w lewo w alfabecie. Wzor: `(c - 'A' - k + 26) % 26 + 'A'`. Dodajemy 26 przed operacja modulo aby uniknac ujemnej reszty z dzielenia.

Weryfikacja KHOOR (k=3):
- K(10)-3=7=H, H(7)-3=4=E, O(14)-3=11=L, O(14)-3=11=L, R(17)-3=14=O -> HELLO

Weryfikacja TFDSFU (k=1):
- T(19)-1=18=S, F(5)-1=4=E, D(3)-1=2=C, S(18)-1=17=R, F(5)-1=4=E, U(20)-1=19=T -> SECRET

Weryfikacja JHZASL (k=7):
- J(9)-7=2=C, H(7)-7=0=A, Z(25)-7=18=S, A(0)-7+26=19=T, S(18)-7=11=L, L(11)-7=4=E -> CASTLE

Weryfikacja YMJWJ (k=5):
- Y(24)-5=19=T, M(12)-5=7=H, J(9)-5=4=E, W(22)-5=17=R, J(9)-5=4=E -> THERE

Weryfikacja PNHFR (k=13):
- P(15)-13=2=C, N(13)-13=0=A, H(7)-13+26=20=U, F(5)-13+26=18=S, R(17)-13=4=E -> CAUSE
</details>

---

### Cwiczenie 8.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 4 (DOPISZ/USUN)

Dany jest napis poczatkowy oraz ciag 8 operacji do wykonania na nim. Operacje sa nastepujacych typow:
- `DOPISZ x` — dopisz znak x na koncu napisu
- `USUN k` — usun znak na pozycji k (pozycje numerujemy od 1)
- `ZAMIEN k x` — zamien znak na pozycji k na znak x

Napisz program, ktory zasymuluje wykonanie operacji i wypisze napis po kazdej operacji oraz napis koncowy.

**Dane**:
```
Napis poczatkowy: ALGORYTM
Operacje:
DOPISZ Y
USUN 3
ZAMIEN 2 X
DOPISZ K
USUN 1
ZAMIEN 3 Z
DOPISZ A
USUN 6
```

**Oczekiwany wynik**:
```
Start:    ALGORYTM
Po op. 1: ALGORYTMY
Po op. 2: ALORYTMY
Po op. 3: AXORYTMY
Po op. 4: AXORYTMYK
Po op. 5: XORYTMYK
Po op. 6: XOZYTMYK
Po op. 7: XOZYTMYKA
Po op. 8: XOZYTYKA
Napis koncowy: XOZYTYKA
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <string>
#include <sstream>
using namespace std;

int main() {
    string napis = "ALGORYTM";
    string operacje[] = {
        "DOPISZ Y", "USUN 3", "ZAMIEN 2 X", "DOPISZ K",
        "USUN 1", "ZAMIEN 3 Z", "DOPISZ A", "USUN 6"
    };
    int n = 8;

    cout << "Start:    " << napis << endl;
    for (int i = 0; i < n; i++) {
        istringstream ss(operacje[i]);
        string typ;
        ss >> typ;
        if (typ == "DOPISZ") {
            char znak;
            ss >> znak;
            napis += znak;
        } else if (typ == "USUN") {
            int k;
            ss >> k;
            napis.erase(k - 1, 1);
        } else if (typ == "ZAMIEN") {
            int k; char znak;
            ss >> k >> znak;
            napis[k - 1] = znak;
        }
        cout << "Po op. " << i + 1 << ": " << napis << endl;
    }
    cout << "Napis koncowy: " << napis << endl;
    return 0;
}
```

**Wyjasnienie**: Symulacja operacji na napisie. Uzywamy metod klasy `string`: `+=` do dopisywania, `erase(pos, 1)` do usuwania, indeksowanie `[]` do zamiany. Pozycje sa 1-indeksowane, wiec odejmujemy 1.

Weryfikacja krok po kroku:
1. ALGORYTM + Y -> ALGORYTMY (dl. 9)
2. ALGORYTMY, usun poz.3 ('G') -> ALORYTMY (dl. 8)
3. ALORYTMY, zamien poz.2 na X ('L'->'X') -> AXORYTMY (dl. 8)
4. AXORYTMY + K -> AXORYTMYK (dl. 9)
5. AXORYTMYK, usun poz.1 ('A') -> XORYTMYK (dl. 8)
6. XORYTMYK, zamien poz.3 na Z ('R'->'Z') -> XOZYTMYK (dl. 8)
7. XOZYTMYK + A -> XOZYTMYKA (dl. 9)
8. XOZYTMYKA, usun poz.6 ('M') -> XOZYTYKA (dl. 8)
</details>

---

### Cwiczenie 8.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 5 + 2025 zad. 2

W pliku `napisy.txt` znajduje sie 12 napisow skladajacych sie z malych i duzych liter alfabetu lacinskiego oraz cyfr (kazdy w osobnym wierszu). Napisz program, ktory:
a) Zliczy ile napisow zawiera dokladnie 3 cyfry.
b) Znajdzie najdluzszy napis, ktory nie zawiera powtarzajacych sie znakow (z uwzglednieniem wielkosci liter, tzn. 'A' i 'a' to rozne znaki).
c) Wypisze napisy, ktore sa palindromami po zamianie wszystkich liter na male (case-insensitive palindrom). Cyfry pozostaja bez zmian.

**Dane** (`napisy.txt`):
```
Ab3cD2eF1
KaJaK
helloworld
Aa1Bb2Cc3
NoRepeat
Aba
x5y5z5w5
RacEcAr
ABCBA
test12
Madam
r2d2R
```

**Oczekiwany wynik**:
```
a) Napisy z dokladnie 3 cyframi:
   Ab3cD2eF1 (cyfry: 3)
   Aa1Bb2Cc3 (cyfry: 3)
   Ilosc: 2

b) Najdluzszy napis bez powtorzen:
   Ab3cD2eF1 (dlugosc: 9, wszystkie znaki rozne)

c) Case-insensitive palindromy:
   KaJaK (male: kajak) - palindrom
   Aba (male: aba) - palindrom
   RacEcAr (male: racecar) - palindrom
   ABCBA (male: abcba) - palindrom
   Madam (male: madam) - palindrom
   r2d2R (male: r2d2r) - palindrom
   Ilosc: 6
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <set>
using namespace std;

string toLower(string s) {
    for (char &c : s)
        if (c >= 'A' && c <= 'Z') c = c - 'A' + 'a';
    return s;
}

bool czyPalindrom(string s) {
    int n = s.length();
    for (int i = 0; i < n / 2; i++)
        if (s[i] != s[n - 1 - i]) return false;
    return true;
}

bool bezPowtorzen(string s) {
    set<char> znaki(s.begin(), s.end());
    return (int)znaki.size() == (int)s.size();
}

int countDigits(string s) {
    int cnt = 0;
    for (char c : s)
        if (c >= '0' && c <= '9') cnt++;
    return cnt;
}

int main() {
    ifstream plik("napisy.txt");
    string s;
    vector<string> napisy;
    while (getline(plik, s)) napisy.push_back(s);

    // a) Dokladnie 3 cyfry
    cout << "a) Napisy z dokladnie 3 cyframi:" << endl;
    int ileA = 0;
    for (string &n : napisy) {
        int d = countDigits(n);
        if (d == 3) {
            cout << "   " << n << " (cyfry: 3)" << endl;
            ileA++;
        }
    }
    cout << "   Ilosc: " << ileA << endl;

    // b) Najdluzszy bez powtorzen
    string najdl = "";
    for (string &n : napisy) {
        if (bezPowtorzen(n) && n.length() > najdl.length()) {
            najdl = n;
        }
    }
    cout << endl << "b) Najdluzszy napis bez powtorzen:" << endl;
    cout << "   " << najdl << " (dlugosc: " << najdl.length()
         << ", wszystkie znaki rozne)" << endl;

    // c) Case-insensitive palindromy
    cout << endl << "c) Case-insensitive palindromy:" << endl;
    int ileC = 0;
    for (string &n : napisy) {
        string maly = toLower(n);
        if (czyPalindrom(maly)) {
            cout << "   " << n << " (male: " << maly << ") - palindrom" << endl;
            ileC++;
        }
    }
    cout << "   Ilosc: " << ileC << endl;
    return 0;
}
```

**Wyjasnienie**: Trzy niezalezne analizy: (a) zliczanie cyfr w kazdym napisie, (b) sprawdzenie unikalnosci znakow za pomoca `set` — jesli rozmiar seta rowny dlugosci napisu, to znaki sie nie powtarzaja, (c) konwersja na male litery i test palindromu.

Weryfikacja:
a) Cyfry w kazdym napisie:
- Ab3cD2eF1: 3,2,1 -> 3 cyfry -> TAK
- KaJaK: 0 cyfr
- helloworld: 0 cyfr
- Aa1Bb2Cc3: 1,2,3 -> 3 cyfry -> TAK
- NoRepeat: 0 cyfr
- Aba: 0 cyfr
- x5y5z5w5: 5,5,5,5 -> 4 cyfry
- RacEcAr: 0 cyfr
- ABCBA: 0 cyfr
- test12: 1,2 -> 2 cyfry
- Madam: 0 cyfr
- r2d2R: 2,2 -> 2 cyfry

b) Napisy bez powtorzen znakow:
- Ab3cD2eF1: {A,b,3,c,D,2,e,F,1} = 9 znakow, 9 roznych -> TAK (dl. 9)
- helloworld: 'l' i 'o' sie powtarzaja -> NIE
- NoRepeat: 'e' sie powtarza -> NIE (N,o,R,e,p,e,a,t)
- test12: 't' sie powtarza -> NIE
Najdluzszy: Ab3cD2eF1 (dl. 9)

c) Palindromy (case-insensitive):
- KaJaK -> kajak -> palindrom
- Aba -> aba -> palindrom
- RacEcAr -> racecar -> palindrom
- ABCBA -> abcba -> palindrom
- Madam -> madam -> palindrom
- r2d2R -> r2d2r -> palindrom
</details>

---
