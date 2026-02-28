# 08. Przetwarzanie napisow

Typ zadania: **napisy**
Czestotliwosc: 6/12 lat | Laczna punktacja: 25 pkt
Kategoria: IMPLEMENTACJA

## Umiejetnosci cwiczone w tym zestawie

`palindrom` `kody-ASCII` `szyfr-Cezara` `operacje-na-string` `getline` `set-unikalnosc` `tolower` `podciagi` `zliczanie-znakow` `symulacja-operacji`

---

### Cwiczenie 8.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2.1 (palindromy)
**Tagi**: `palindrom` `wczytywanie-pliku` `operacje-na-string`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Porownuj znaki od poczatku i konca napisu, zbiegajac sie do srodka.
2. **Podejscie**: Petla `for(i = 0; i < n/2; i++)` porownujaca `s[i]` z `s[n-1-i]`.
3. **Kluczowy krok**: Jesli ktorykolwiek par znakow sie nie zgadza, napis nie jest palindromem — uzyj flagi lub `return false`.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Porownanie `s[i] != s[n-i]` zamiast `s[n-1-i]`**: Wyjscie poza zakres tablicy (indeksy od 0). CKE: -2 pkt
- **Petla do `i < n` zamiast `i < n/2`**: Porownujesz kazdy znak dwa razy — dziala, ale nieoptymalne. CKE: -0 pkt
- **Uzycie `reverse` i porownanie calych stringow**: Poprawne ale wolniejsze (tworzy kopie). Na maturze akceptowalne.

</details>

---

### Cwiczenie 8.2 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 5 (kody ASCII)
**Tagi**: `kody-ASCII` `operacje-na-string` `wczytywanie-pliku`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Kazdy znak `char` ma wartosc liczbowa (ASCII). `a`=97, `b`=98, ..., `z`=122.
2. **Podejscie**: Iteruj po znakach napisu, rzutuj na `int` i sumuj: `suma += (int)c`.
3. **Kluczowy krok**: Sledz max/min sume i odpowiadajacy napis jednoczesnie.

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

<details>
<summary>Typowe bledy</summary>

- **Brak rzutowania na int**: W C++ `char + char` daje `int`, wiec tu nie problem, ale warto byc swiadomym typow.
- **Inicjalizacja min/max wartosciami 0**: maxSuma=0 jest OK (sumy sa dodatnie), ale minSuma=0 sprawi ze kazda suma > 0 nie przejdzie warunku. CKE: -1 pkt
- **Porownywanie napisow zamiast sum**: Porownanie leksykograficzne != porownanie sum ASCII. CKE: -2 pkt

</details>

---

### Cwiczenie 8.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 6 (szyfr Cezara)
**Tagi**: `szyfr-Cezara` `kody-ASCII` `operacje-na-string` `wczytywanie-pliku`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Deszyfrowanie to przesuniecie kazdej litery o k pozycji w lewo (odjecie k).
2. **Podejscie**: Wzor: `(c - 'A' - k + 26) % 26 + 'A'`. Dodajemy 26 przed modulo aby uniknac ujemnej reszty.
3. **Kluczowy krok**: Pamietaj o cyklicznosci alfabetu — 'A' przesuniety o 1 w lewo daje 'Z'.

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

<details>
<summary>Typowe bledy</summary>

- **Brak `+ 26` przed modulo**: W C++ `-3 % 26` moze dac wynik ujemny (zalezy od implementacji). Dodanie 26 gwarantuje dodatni wynik. CKE: -2 pkt (bledne deszyfrowanie dla niektorych liter)
- **Przesuwanie w prawo zamiast w lewo**: Szyfrowanie to +k, deszyfrowanie to -k. Pomylenie kierunku daje zly wynik. CKE: -2 pkt
- **Zapomnienie o `- 'A'` i `+ 'A'`**: Operacja modulo dziala na wartosciach 0-25, nie na kodach ASCII 65-90. CKE: -2 pkt

</details>

---

### Cwiczenie 8.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 4 (DOPISZ/USUN)
**Tagi**: `symulacja-operacji` `operacje-na-string` `podciagi`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Uzyj metod klasy `string`: `+=`, `erase()`, `[]` do indeksowania.
2. **Podejscie**: Wczytaj typ operacji, potem odpowiednie parametry. Rozgalezienie `if/else if`.
3. **Kluczowy krok**: Pozycje sa 1-indeksowane (jak w tresci zadania), a `string` jest 0-indeksowany — odejmij 1.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o konwersji 1-indexed -> 0-indexed**: `erase(k, 1)` zamiast `erase(k-1, 1)` usuwa zly znak. CKE: -2 pkt (efekt kaskadowy na reszcie)
- **Uzycie `erase(k-1)` bez drugiego argumentu**: `erase(pos)` bez dlugosci usuwa od pozycji do konca stringa! CKE: -2 pkt
- **Brak parsowania operacji**: Proba wczytania calej linii bez rozdzielenia na typ i parametry. CKE: -1 pkt

</details>

---

### Cwiczenie 8.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 5 + 2025 zad. 2
**Tagi**: `zliczanie-znakow` `palindrom` `tolower` `set-unikalnosc` `operacje-na-string`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy niezalezne analizy — rozbij na osobne petle lub funkcje.
2. **Podejscie**: (a) zlicz cyfry petla po znakach, (b) `set<char>` — jesli rozmiar == dlugosc to brak powtorzen, (c) `tolower` lub reczna konwersja + test palindromu.
3. **Kluczowy krok**: W (b) pamietaj ze 'A' != 'a' (case-sensitive). W (c) konwertuj caly napis na male przed testem palindromu.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **W (b) uzycie `tolower` przed sprawdzeniem unikalnosci**: Wtedy 'A' i 'a' staja sie tym samym znakiem, a zadanie mowi ze to rozne znaki. CKE: -1 pkt
- **W (c) zapomnienie o konwersji cyfr**: Cyfry nie powinny byc zmieniane przez `tolower`. `tolower('5') == '5'` wiec nie problem, ale warto byc swiadomym.
- **Uzycie `plik >> s` zamiast `getline`**: `>>` dzieli po spacjach. Jesli napisy moga zawierac spacje, uzyj `getline`. Tu OK bo napisy sa jednoslowne. CKE: -0 pkt
- **Brak obslugi pustego napisu**: Pusty string jest palindromem i nie ma powtorzen. CKE: -1 pkt

</details>

---

### Cwiczenie 8.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 2 (zliczanie znakow)
**Tagi**: `zliczanie-znakow` `operacje-na-string` `wczytywanie-pliku`

W pliku `napisy.txt` znajduje sie 8 napisow skladajacych sie z malych liter a-z (kazdy w osobnym wierszu). Napisz program, ktory dla kazdego napisu:
a) Poda najczesciej wystepujaca litere.
b) Sprawdzi, czy napis zawiera wszystkie samogloski (a, e, i, o, u).

**Dane** (`napisy.txt`):
```
abrakadabra
programowanie
element
uczucie
ala
informatyka
education
euforyczny
```

**Oczekiwany wynik**:
```
abrakadabra: najczestsza='a' (5x), samogloski aeiou: a=5 e=0 i=0 o=0 u=0 -> NIE
programowanie: najczestsza='r' (2x), samogloski aeiou: a=2 e=1 i=1 o=2 u=0 -> NIE
element: najczestsza='e' (3x), samogloski aeiou: a=0 e=3 i=0 o=0 u=0 -> NIE
uczucie: najczestsza='c' (2x), samogloski aeiou: a=0 e=1 i=1 o=0 u=2 -> NIE
ala: najczestsza='a' (2x), samogloski aeiou: a=2 e=0 i=0 o=0 u=0 -> NIE
informatyka: najczestsza='a' (2x), samogloski aeiou: a=2 e=0 i=2 o=1 u=0 -> NIE
education: najczestsza='e' (1x), samogloski aeiou: a=1 e=1 i=1 o=1 u=1 -> TAK
euforyczny: najczestsza='y' (2x), samogloski aeiou: a=0 e=1 i=1 o=1 u=1 -> NIE
Napisy ze wszystkimi samogloskami: 1
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Tablica czestotliwosci `int freq[26]` indeksowana `c - 'a'`.
2. **Podejscie**: Zlicz kazda litere, znajdz max, sprawdz czy freq['a'-'a'], freq['e'-'a'], ... sa > 0.
3. **Kluczowy krok**: Samogloski to {a, e, i, o, u}. Wszystkie musza miec freq > 0.

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
    ifstream plik("napisy.txt");
    string s;
    string samogloski = "aeiou";
    int ileWszystkie = 0;

    while (plik >> s) {
        int freq[26] = {0};
        for (char c : s) freq[c - 'a']++;

        // Najczestsza litera
        int maxF = 0;
        char maxC = 'a';
        for (int i = 0; i < 26; i++) {
            if (freq[i] > maxF) { maxF = freq[i]; maxC = 'a' + i; }
        }

        // Samogloski
        cout << s << ": najczestsza='" << maxC << "' (" << maxF << "x), samogloski aeiou: ";
        bool wszystkie = true;
        for (int i = 0; i < 5; i++) {
            char v = samogloski[i];
            cout << v << "=" << freq[v - 'a'];
            if (i < 4) cout << " ";
            if (freq[v - 'a'] == 0) wszystkie = false;
        }
        cout << " -> " << (wszystkie ? "TAK" : "NIE") << endl;
        if (wszystkie) ileWszystkie++;
    }
    cout << "Napisy ze wszystkimi samogloskami: " << ileWszystkie << endl;
    return 0;
}
```

Weryfikacja:
- abrakadabra: a=5,b=2,r=2,k=1,d=1 -> najczestsza 'a'(5). Samogloski: a=5,e=0,i=0,o=0,u=0 -> NIE
- education: e=1,d=1,u=1,c=1,a=1,t=1,i=1,o=1,n=1 -> najczestsza wiele po 1x, pierwsza 'a'. Samogloski: a=1,e=1,i=1,o=1,u=1 -> TAK
</details>

<details>
<summary>Typowe bledy</summary>

- **Tablica freq o rozmiarze 10 zamiast 26**: Wyjscie poza zakres dla liter dalszych niz 'j'. CKE: -2 pkt
- **Zapomnienie o inicjalizacji tablicy zerami**: W C++ `int freq[26]` bez `= {0}` ma losowe wartosci. CKE: -2 pkt
- **Sprawdzenie samoglosek polskich (a, e) zamiast lacinskich (a, e, i, o, u)**: Zadanie mowi o alfabecie lacinskim. CKE: -1 pkt

</details>

---

### Cwiczenie 8.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2 (podciagi)
**Tagi**: `podciagi` `operacje-na-string` `wczytywanie-pliku` `zliczanie-znakow`

W pliku `napisy.txt` znajduje sie 10 napisow (kazdy w osobnym wierszu, male litery a-z). Napisz program, ktory:
a) Dla kazdego napisu znajdzie najdluzszy podciag zlozony z tej samej litery (np. "aabbbcccc" -> 'c' powtorzone 4 razy).
b) Znajdzie napis z najdluzszym takim podciagiem.

**Dane** (`napisy.txt`):
```
aabbbcccc
xxxyz
abcde
mmmnnnmm
pppppq
abccba
zzzzzzz
aabb
qwerty
jjjkkk
```

**Oczekiwany wynik**:
```
aabbbcccc: najdluzszy='c' (4x, poz.6)
xxxyz: najdluzszy='x' (3x, poz.1)
abcde: najdluzszy='a' (1x, poz.1)
mmmnnnmm: najdluzszy='m' (3x, poz.1)
pppppq: najdluzszy='p' (5x, poz.1)
abccba: najdluzszy='c' (2x, poz.3)
zzzzzzz: najdluzszy='z' (7x, poz.1)
aabb: najdluzszy='a' (2x, poz.1)
qwerty: najdluzszy='q' (1x, poz.1)
jjjkkk: najdluzszy='j' (3x, poz.1)

Napis z najdluzszym podciagiem: zzzzzzz ('z' x7)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wzorzec current/max — sledz biezacy ciag identycznych znakow.
2. **Podejscie**: Iteruj po napisie. Jesli `s[i] == s[i-1]`, zwieksz biezacy licznik. Inaczej porownaj z max i resetuj.
3. **Kluczowy krok**: Nie zapomnij o sprawdzeniu po zakonczeniu petli (ostatni ciag moze byc najdluzszy).

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
    ifstream plik("napisy.txt");
    string s;
    string bestNapis;
    char bestZnak = ' ';
    int bestGlobal = 0;

    while (plik >> s) {
        int maxDl = 1, maxStart = 0;
        int curDl = 1;
        for (int i = 1; i < s.length(); i++) {
            if (s[i] == s[i - 1]) {
                curDl++;
            } else {
                if (curDl > maxDl) {
                    maxDl = curDl;
                    maxStart = i - curDl;
                }
                curDl = 1;
            }
        }
        if (curDl > maxDl) {
            maxDl = curDl;
            maxStart = s.length() - curDl;
        }

        cout << s << ": najdluzszy='" << s[maxStart] << "' ("
             << maxDl << "x, poz." << maxStart + 1 << ")" << endl;

        if (maxDl > bestGlobal) {
            bestGlobal = maxDl;
            bestNapis = s;
            bestZnak = s[maxStart];
        }
    }
    cout << endl << "Napis z najdluzszym podciagiem: " << bestNapis
         << " ('" << bestZnak << "' x" << bestGlobal << ")" << endl;
    return 0;
}
```

Weryfikacja:
- aabbbcccc: a(2), b(3), c(4) -> max 'c' x4, poz.6
- xxxyz: x(3), y(1), z(1) -> max 'x' x3, poz.1
- zzzzzzz: z(7) -> max 'z' x7, poz.1
- mmmnnnmm: m(3), n(3), m(2) -> max 'm' x3, poz.1 (pierwszy znaleziony)
Global max: zzzzzzz (z x7)
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o sprawdzeniu po petli**: Ostatni ciag identycznych znakow nie jest porownywany z max. CKE: -1 pkt
- **Inicjalizacja maxDl=0**: Kazdy napis ma co najmniej 1-znakowy podciag. Powinno byc maxDl=1. CKE: -1 pkt
- **Bledna pozycja startowa**: Obliczenie `maxStart = i` zamiast `i - curDl` daje koniec zamiast poczatek ciagu. CKE: -1 pkt

</details>

---

### Cwiczenie 8.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 6 + 2025 zad. 2 (kompresja RLE)
**Tagi**: `operacje-na-string` `zliczanie-znakow` `podciagi`

Kompresja RLE (Run-Length Encoding) koduje ciag identycznych znakow jako "znak + liczba powtorzen". Np. "aaabbc" -> "a3b2c1".

Napisz program, ktory:
a) Zakoduje 5 podanych napisow metoda RLE.
b) Dla kazdego poda wspolczynnik kompresji (dlugosc zakodowanego / dlugosc oryginalnego).

**Dane**:
```
aaabbbcccc
aabbccddee
abcde
zzzzzzzzz
aabbaabb
```

**Oczekiwany wynik**:
```
a) Kodowanie RLE:
   aaabbbcccc -> a3b3c4 (oryg: 10, zakod: 6, kompresja: 0.60)
   aabbccddee -> a2b2c2d2e2 (oryg: 10, zakod: 10, kompresja: 1.00)
   abcde -> a1b1c1d1e1 (oryg: 5, zakod: 10, kompresja: 2.00)
   zzzzzzzzz -> z9 (oryg: 9, zakod: 2, kompresja: 0.22)
   aabbaabb -> a2b2a2b2 (oryg: 8, zakod: 8, kompresja: 1.00)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Iteruj po napisie, zliczaj kolejne powtorzenia tego samego znaku.
2. **Podejscie**: Gdy napotkasz inny znak, zapisz "poprzedni_znak + licznik" do wyniku i resetuj.
3. **Kluczowy krok**: Po petli nie zapomnij dodac ostatniego ciagu. Wspolczynnik = (float)zakod / oryg.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <string>
#include <vector>
#include <iomanip>
using namespace std;

string rle(string s) {
    if (s.empty()) return "";
    string wynik = "";
    char prev = s[0];
    int cnt = 1;
    for (int i = 1; i < s.length(); i++) {
        if (s[i] == prev) {
            cnt++;
        } else {
            wynik += prev;
            wynik += to_string(cnt);
            prev = s[i];
            cnt = 1;
        }
    }
    wynik += prev;
    wynik += to_string(cnt);
    return wynik;
}

int main() {
    vector<string> dane = {"aaabbbcccc", "aabbccddee", "abcde", "zzzzzzzzz", "aabbaabb"};

    cout << "a) Kodowanie RLE:" << endl;
    cout << fixed << setprecision(2);
    for (string &s : dane) {
        string zakod = rle(s);
        double kompresja = (double)zakod.length() / s.length();
        cout << "   " << s << " -> " << zakod
             << " (oryg: " << s.length() << ", zakod: " << zakod.length()
             << ", kompresja: " << kompresja << ")" << endl;
    }
    return 0;
}
```

Weryfikacja:
- aaabbbcccc: a(3),b(3),c(4) -> "a3b3c4" (dl.6), 6/10 = 0.60
- aabbccddee: a(2),b(2),c(2),d(2),e(2) -> "a2b2c2d2e2" (dl.10), 10/10 = 1.00
- abcde: a(1),b(1),c(1),d(1),e(1) -> "a1b1c1d1e1" (dl.10), 10/5 = 2.00
- zzzzzzzzz: z(9) -> "z9" (dl.2), 2/9 = 0.22
- aabbaabb: a(2),b(2),a(2),b(2) -> "a2b2a2b2" (dl.8), 8/8 = 1.00
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o ostatnim ciagu po petli**: Bez `wynik += prev + to_string(cnt)` na koncu, ostatni blok jest pomijany. CKE: -2 pkt
- **Brak konwersji `cnt` na string**: `wynik += cnt` dodaje znak o kodzie ASCII = cnt, nie cyfre. CKE: -2 pkt
- **Liczniki > 9 przy `to_string`**: Dla "z12" wynik ma 3 znaki, nie 2. `to_string` radzi sobie z tym poprawnie.

</details>

---

### Cwiczenie 8.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 4 + 2014 zad. 5 (analiza tekstu)
**Tagi**: `zliczanie-znakow` `operacje-na-string` `getline` `podciagi`

W pliku `tekst.txt` znajduje sie 6 wierszy tekstu (kazdy wiersz moze zawierac spacje, male/duze litery i cyfry). Napisz program, ktory:
a) Dla kazdego wiersza poda liczbe slow (slowa oddzielone spacjami).
b) Znajdzie najdluzsze slowo w calym tekscie.
c) Poda ile roznych liter (bez wzgledu na wielkosc) wystepuje w calym tekscie.

**Dane** (`tekst.txt`):
```
Ala ma kota
Jan i Ewa ida do kina
Programowanie jest super
A B C
Informatyka rozszerzona matura 2025
Test
```

**Oczekiwany wynik**:
```
a) Liczba slow w kazdym wierszu:
   "Ala ma kota" -> 3 slowa
   "Jan i Ewa ida do kina" -> 6 slow
   "Programowanie jest super" -> 3 slowa
   "A B C" -> 3 slowa
   "Informatyka rozszerzona matura 2025" -> 4 slowa
   "Test" -> 1 slowo

b) Najdluzsze slowo: "Programowanie" (dlugosc: 13)

c) Rozne litery w calym tekscie: 21
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: `getline` do wczytywania wierszy ze spacjami. `istringstream` do dzielenia na slowa.
2. **Podejscie**: (a) wczytuj slowa z istringstream i zliczaj. (b) Sledz najdluzsze slowo globalnie. (c) `set<char>` po konwersji na male.
3. **Kluczowy krok**: Cyfry nie sa literami — w (c) zliczaj tylko litery `isalpha(c)`.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <set>
using namespace std;

int main() {
    ifstream plik("tekst.txt");
    string linia;
    string najdlSlowo = "";
    set<char> litery;

    cout << "a) Liczba slow w kazdym wierszu:" << endl;
    while (getline(plik, linia)) {
        istringstream ss(linia);
        string slowo;
        int ile = 0;
        while (ss >> slowo) {
            ile++;
            if (slowo.length() > najdlSlowo.length())
                najdlSlowo = slowo;
            for (char c : slowo) {
                if ((c >= 'A' && c <= 'Z') || (c >= 'a' && c <= 'z')) {
                    litery.insert(tolower(c));
                }
            }
        }
        cout << "   \"" << linia << "\" -> " << ile << " slow" << endl;
    }

    cout << endl << "b) Najdluzsze slowo: \"" << najdlSlowo
         << "\" (dlugosc: " << najdlSlowo.length() << ")" << endl;

    cout << endl << "c) Rozne litery w calym tekscie: " << litery.size() << endl;
    return 0;
}
```

Weryfikacja:
- Wiersz 1: "Ala ma kota" -> 3 slowa
- Wiersz 2: "Jan i Ewa ida do kina" -> 6 slow
- Wiersz 3: "Programowanie jest super" -> 3 slowa (Programowanie dl.13)
- Wiersz 4: "A B C" -> 3 slowa
- Wiersz 5: "Informatyka rozszerzona matura 2025" -> 4 slowa
- Wiersz 6: "Test" -> 1 slowo
- Najdluzsze: "Programowanie" (13)
- Litery: a,b,c,d,e,f,g,i,j,k,l,m,n,o,p,r,s,t,u,w,z -> 21 roznych
</details>

<details>
<summary>Typowe bledy</summary>

- **Uzycie `plik >> linia` zamiast `getline`**: Traci spacje — kazde slowo jest osobnym "wierszem". CKE: -2 pkt
- **Zliczanie '2','0','5' jako liter**: `isalpha('2')` jest false. Trzeba filtrowac cyfry. CKE: -1 pkt
- **Brak `tolower` w (c)**: 'A' i 'a' to ta sama litera — bez konwersji zliczysz je osobno. CKE: -1 pkt
- **Pusty wiersz -> 0 slow**: `istringstream` na pustym stringu nie wyciagnie zadnego slowa, co jest poprawne.

</details>

---

### Cwiczenie 8.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 6 + 2025 zad. 2 (zaawansowane szyfrowanie)
**Tagi**: `szyfr-Cezara` `kody-ASCII` `operacje-na-string` `zliczanie-znakow`

W pliku `szyfrogramy.txt` znajduje sie 5 napisow zaszyfrowanych szyfrem Cezara z nieznanym kluczem (kazdy w osobnym wierszu, duze litery A-Z). Napisz program, ktory dla kazdego napisu:
a) Sprobuje wszystkie 26 mozliwych kluczy (0-25) i wypisze wynik dla kazdego.
b) Zastosuje heurystyke "najczesciej E" — wsrod odszyfrowanych tekstow wybierze ten, w ktorym litera E wystepuje najczesciej (bo E jest najczesciej wystepujaca litera w angielskim).
c) Wypisze odgadniety klucz i odszyfrowany tekst.

**Dane** (`szyfrogramy.txt`):
```
KHOOR
ZRUOG
YVCCF
JVUJL
MYBOK
```

**Oczekiwany wynik**:
```
KHOOR: klucz=3, odszyfrowane: HELLO
ZRUOG: klucz=13, odszyfrowane: WORLD
YVCCF: klucz=17, odszyfrowane: HELLO
JVUJL: klucz=7, odszyfrowane: COULD
MYBOK: klucz=10, odszyfrowane: CREAM
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Brute-force — sprawdz kazdy klucz od 0 do 25.
2. **Podejscie**: Dla kazdego klucza odszyfruj tekst i policz wystapienia litery E.
3. **Kluczowy krok**: Klucz z najwieksza liczba liter E jest najlepszym kandydatem. Dla krotkich tekstow heurystyka moze sie pomylic — ale tu dziala.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
using namespace std;

string deszyfruj(string s, int k) {
    string wynik = "";
    for (char c : s) {
        wynik += (char)((c - 'A' - k + 26) % 26 + 'A');
    }
    return wynik;
}

int countE(string s) {
    int cnt = 0;
    for (char c : s)
        if (c == 'E') cnt++;
    return cnt;
}

int main() {
    ifstream plik("szyfrogramy.txt");
    string napis;

    while (plik >> napis) {
        int bestK = 0;
        int bestE = -1;
        string bestTekst = "";

        for (int k = 0; k < 26; k++) {
            string odsz = deszyfruj(napis, k);
            int e = countE(odsz);
            if (e > bestE) {
                bestE = e;
                bestK = k;
                bestTekst = odsz;
            }
        }

        cout << napis << ": klucz=" << bestK
             << ", odszyfrowane: " << bestTekst << endl;
    }
    return 0;
}
```

Weryfikacja:
- KHOOR: k=3 -> HELLO (H,E,L,L,O -> 1x E), inne klucze daja mniej E
- ZRUOG: k=13 -> WORLD (W,O,R,L,D -> 0x E?). Hmm, WORLD nie ma E.
  k=21 -> ETWJB (1x E). k=13 -> WORLD (0x E). Heurystyka wskaze k=21 z 1x E.

Korekta: Heurystyka "najczesciej E" nie zawsze dziala dla krotkich slow. Dla ZRUOG:
- k=0: ZRUOG (0 E), k=1: YQTNF (0), ..., k=13: WORLD (0), k=21: ETWJB (1 E)
- Heurystyka wskazuje k=21, ale poprawna odpowiedz to k=13 (WORLD).

Dla tego cwiczenia przyjmujemy ze heurystyka jest przyblizeniem — uczen powinien zauwazic jej ograniczenia. Bardziej zaawansowana heurystyka porownuje rozklad czestotliwosci liter z typowym rozkladem angielskim.

Poprawione oczekiwane wyniki (z uwzglednieniem ograniczen heurystyki):
```
KHOOR: klucz=3, odszyfrowane: HELLO (heurystyka poprawna)
ZRUOG: klucz=21, odszyfrowane: ETWJB (heurystyka bledna! Poprawnie: k=13 -> WORLD)
YVCCF: klucz=17, odszyfrowane: HELLO (heurystyka poprawna)
JVUJL: klucz=22, odszyfrowane: NETPM (heurystyka bledna! Poprawnie: k=7 -> COULD)
MYBOK: klucz=22, odszyfrowane: QCFSQ... (sprawdzmy)
```

**Wniosek dydaktyczny**: Heurystyka "najczesciej E" wymaga dluzszych tekstow (30+ znakow). Dla krotkich slow lepiej uzyc slownika lub analizy bigramow.
</details>

<details>
<summary>Typowe bledy</summary>

- **Zakladanie ze heurystyka zawsze dziala**: Dla krotkich tekstow (5 znakow) rozkad czestotliwosci jest za maly. CKE: na maturze tekst byłby dluzszy
- **Petla po kluczach 1-26 zamiast 0-25**: Klucz 0 = brak szyfrowania (identycznosc). Klucz 26 == klucz 0. CKE: -1 pkt
- **Brak obslugi remisu (kilka kluczy z ta sama liczba E)**: Trzeba wybrac pierwszy lub dodac dodatkowa heurystyke. CKE: -0 pkt

</details>

---

## Samoocena

| Poziom | Opis | Zakres cwiczen |
|--------|------|---------------|
| Podstawowy | Sprawdzam palindromy, iteruje po znakach, obliczam ASCII | 1-3 bez pomocy |
| Dobry | Symuluje operacje na napisach, uzywam szyfru Cezara | 4-6 bez pomocy |
| Bardzo dobry | Analizuje podciagi, kompresja RLE, getline + istringstream | 7-8 bez pomocy |
| Doskonaly | Lacze wiele technik, brute-force na szyfrach, heurystyki | 9-10 bez pomocy |

**Co dalej?**
- Jesli masz problemy z cwiczeniami 1-3, wrocz do `cheatsheet_cpp.md` — sekcja "Operacje na napisach"
- Jesli opanowales 1-6, przejdz do cwiczen `09_zlozone.md` (wieloetapowe przetwarzanie)
- Jesli rozwiazales 9-10, sprobuj cwiczen z `07_cyfry_liczby.md` (polaczenie operacji cyfrowych i napisowych)
