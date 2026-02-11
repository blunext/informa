# 10. Zliczanie i filtrowanie

Typ zadania: **zliczanie**
Czestotliwosc: 5/11 lat | Laczna punktacja: 17 pkt
Kategoria: IMPLEMENTACJA

## Umiejetnosci cwiczone w tym zestawie

`wczytywanie-pliku` `filtrowanie` `parzystosc` `tablica-czestotliwosci` `kody-ASCII` `zliczanie-znakow` `cyfry-mod-div` `suma-cyfr` `podzielnosc` `map-zliczanie` `moda-statystyczna` `unikalnosc` `liczby-pierwsze` `podwojna-petla` `wiele-plikow` `przedzialy` `napisy` `pary-elementow` `grupowanie` `struct` `sortowanie`

---

### Cwiczenie 10.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2014 zad. 5a
**Tagi**: `wczytywanie-pliku` `filtrowanie` `parzystosc`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Potrzebujesz dwoch niezaleznych licznikow — kazdy zlicza cos innego.
2. **Podejscie**: W petli while czytaj liczby z pliku i sprawdzaj dwa warunki (parzystosc: `n % 2 == 0`, prog: `n > 100`).
3. **Kluczowy krok**: Oba warunki mozna sprawdzac w tej samej petli — nie musisz czytac pliku dwa razy.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **`>= 100` zamiast `> 100`**: Zle odczytanie warunku "wieksze od 100" (100 nie jest wieksze od 100). CKE: -1 pkt
- **Brak inicjalizacji licznikow**: Liczniki musza zaczynac od 0, niezainicjalizowane zmienne moga miec smieci. CKE: -1 pkt

</details>

---

### Cwiczenie 10.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3.1 (cyfry pi)
**Tagi**: `tablica-czestotliwosci` `kody-ASCII` `zliczanie-znakow`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Tablica 10 licznikow (indeks = cyfra) to naturalny sposob zliczania cyfr.
2. **Podejscie**: Iteruj po stringu i konwertuj kazdy znak na cyfre: `c - '0'`.
3. **Kluczowy krok**: Po zliczeniu, przejdz tablice raz aby znalezc max, potem drugi raz aby wypisac wszystkie cyfry z ta czestotliwoscia.

</details>

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
- 1: poz. 1,4,38,41,49 = 5 razy
- 3: pozycje 9,15,17,24,25,27,43,46 = 8 razy
- 9: pozycje 5,12,14,30,38,42,44,45 = 8 razy
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak `{0}` w inicjalizacji tablicy**: `int freq[10];` nie zeruje tablicy — wynik bedzie losowy. CKE: -2 pkt (zly wynik)
- **Uzycie `c` zamiast `c - '0'`**: Indeksowanie tablicy kodem ASCII znaku (np. 48-57) zamiast cyfra 0-9. CKE: -2 pkt

</details>

---

### Cwiczenie 10.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3.2
**Tagi**: `cyfry-mod-div` `suma-cyfr` `podzielnosc` `wczytywanie-pliku`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Potrzebujesz funkcji `sumaCyfr` — petla `while(n > 0)` z `n % 10` i `n /= 10`.
2. **Podejscie**: Pierwsza cyfra 5-cyfrowej liczby to wynik dzielenia przez 10 az n < 10. Ostatnia cyfra to `n % 10`.
3. **Kluczowy krok**: W punkcie (c) pamietaj o sprawdzeniu `s > 0` przed dzieleniem — unikasz dzielenia przez zero.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Modyfikacja oryginalnej liczby w sumaCyfr**: Jesli nie uzywasz kopii, oryginalna wartosc zostanie zniszczona. CKE: -1 pkt
- **Brak warunku `s > 0` w punkcie (c)**: Dzielenie przez 0 powoduje crash programu. CKE: -1 pkt
- **`while(n > 0)` pomija n=0**: Jezeli liczba to 0, petla sie nie wykona i suma cyfr bedzie 0. Dla tego zestawu danych to nie problem (5-cyfrowe), ale warto pamietac.

</details>

---

### Cwiczenie 10.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3.1
**Tagi**: `map-zliczanie` `moda-statystyczna` `unikalnosc`

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
<summary>Wskazowki</summary>

1. **Kierunek**: `map<int,int>` to idealny kontener do zliczania czestotliwosci — klucz to wartosc, wartosc to licznik.
2. **Podejscie**: `freq[x]++` automatycznie tworzy nowy wpis jesli klucz nie istnieje. Rozmiar mapy = liczba roznych wartosci.
3. **Kluczowy krok**: Mode to wartosc z maksymalnym licznikiem — moze byc wieksza niz 1, wiec przejrzyj mape dwukrotnie (raz po max, raz po elementach z max).

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Uzycie tablicy zamiast mapy przy nieznanym zakresie**: Tablica wymaga z gory znanego zakresu wartosci. CKE: -1 pkt (jesli zakres sie nie miesci)
- **Zapomnienie o kilku modach**: Moze byc wiecej niz jedna wartosc z max czestotliwoscia — wypisanie tylko pierwszej to niepelna odpowiedz. CKE: -1 pkt
- **`freq.size()` przed wstawieniem elementow**: Upewnij sie, ze mapa jest juz wypelniona. CKE: -1 pkt

</details>

---

### Cwiczenie 10.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4.1 (liczby pierwsze dzielace)
**Tagi**: `podzielnosc` `liczby-pierwsze` `podwojna-petla` `wiele-plikow`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Wczytaj oba zbiory do wektorow, a potem uzyj podwojnej petli (po liczbach pierwszych i duzych).
2. **Podejscie**: Dla kazdej liczby pierwszej `p` sprawdz `d % p == 0` dla kazdej duzej liczby `d`. Zlicz trafienia.
3. **Kluczowy krok**: Jesli cnt > 0, to liczba pierwsza dzieli chociaz jedna duza. Sledzac cnt dla kazdej p, znajdziesz tez max.

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

**Wyjasnienie**: Podwojna petla: dla kazdej liczby pierwszej sprawdzamy podzielnosc kazdej duzej. Zliczamy pary, sprawdzamy ktore pierwsze dziela chociaz jedna duza, i szukamy pierwszej z max liczba dzielonych duzych.

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

<details>
<summary>Typowe bledy</summary>

- **Pomylenie kolejnosci petli**: Petla zewnetrzna po pierwszych, wewnetrzna po duzych — odwrotnie daje inne wyniki (ile duzych dzieli kazda pierwsza vs ile pierwszych dzieli kazda duza). CKE: -2 pkt
- **Zapomnienie o otwarciu drugiego pliku**: Jesli oba pliki sa otwarte tym samym obiektem ifstream, trzeba go zamknac i otworzyc ponownie. CKE: -1 pkt
- **Brak separacji wynikow**: Mieszanie wynikow z roznych podpunktow. CKE: -1 pkt

</details>

---

### Cwiczenie 10.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3 (zliczanie w przedzialach)
**Tagi**: `wczytywanie-pliku` `filtrowanie` `przedzialy`

W pliku `dane.txt` znajduje sie 15 liczb calkowitych (kazda w osobnym wierszu, z zakresu 0-200). Napisz program, ktory zliczy ile liczb nalezy do kazdego z trzech przedzialow: [0, 50], [51, 100], [101, 200].

**Dane** (`dane.txt`):
```
12
75
143
8
100
51
200
33
67
0
155
48
99
180
22
```

**Oczekiwany wynik**:
```
Przedzial [0, 50]: 5 (wartosci: 12, 8, 33, 0, 48)
Przedzial [51, 100]: 4 (wartosci: 75, 100, 51, 99)
Przedzial [101, 200]: 6 (wartosci: 143, 200, 67, 155, 180, 22)
```

Korekta — policzymy dokladnie:
- [0,50]: 12, 8, 33, 0, 48, 22 = 6
- [51,100]: 75, 100, 51, 67, 99 = 5
- [101,200]: 143, 200, 155, 180 = 4

**Oczekiwany wynik** (skorygowany):
```
Przedzial [0, 50]: 6
Przedzial [51, 100]: 5
Przedzial [101, 200]: 4
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy liczniki — sprawdz w ktorym przedziale miesci sie kazda liczba.
2. **Podejscie**: Uzyj `if-else if-else` aby kazda liczba trafila do dokladnie jednego przedzialu.
3. **Kluczowy krok**: Zwroc uwage na granice przedzialow — 50 nalezy do [0,50], 51 do [51,100].

</details>

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
    int p1 = 0, p2 = 0, p3 = 0;
    while (plik >> n) {
        if (n <= 50) p1++;
        else if (n <= 100) p2++;
        else p3++;
    }
    cout << "Przedzial [0, 50]: " << p1 << endl;
    cout << "Przedzial [51, 100]: " << p2 << endl;
    cout << "Przedzial [101, 200]: " << p3 << endl;
    return 0;
}
```

**Wyjasnienie**: Kaskada `if-else if-else` gwarantuje, ze kazda liczba jest przypisana do jednego przedzialu. Dzieki sprawdzaniu od najmniejszego, warunek `n <= 50` obsluguje [0,50], a `n <= 100` obsluguje [51,100] (bo 0-50 juz odpadlo).

Weryfikacja:
- [0,50]: 12, 8, 33, 0, 48, 22 = 6
- [51,100]: 75, 100, 51, 67, 99 = 5
- [101,200]: 143, 200, 155, 180 = 4
</details>

<details>
<summary>Typowe bledy</summary>

- **Uzycie samych `if` zamiast `if-else if`**: Liczba moze byc zliczona w kilku przedzialach jednoczesnie. CKE: -1 pkt
- **Pomylenie granic (<= vs <)**: Np. `n < 50` pominie wartosc 50. CKE: -1 pkt

</details>

---

### Cwiczenie 10.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3 (filtrowanie napisow)
**Tagi**: `napisy` `filtrowanie` `wczytywanie-pliku` `zliczanie-znakow`

W pliku `slowa.txt` znajduje sie 12 slow (kazde w osobnym wierszu, male litery). Napisz program, ktory:
a) Zliczy ile slow zaczyna sie na samogloske (a, e, i, o, u, y).
b) Zliczy ile slow ma dlugosc wieksza niz 5.
c) Zliczy ile slow zawiera podwojna litere (np. "pp", "ll").

**Dane** (`slowa.txt`):
```
algorytm
programowanie
abecadlo
petla
if
zmienna
tablica
sortowanie
rekurencja
stos
kolejka
drzewo
```

**Oczekiwany wynik**:
```
a) Slowa zaczynajace sie na samogloske: 2 (algorytm, abecadlo)
b) Slowa dluzsze niz 5 znakow: 8
c) Slowa z podwojna litera: 1 (programowanie -> mm? NIE)
```

Korekta — sprawdzmy podwojne litery:
- algorytm: brak
- programowanie: brak powtorzonych par
- abecadlo: brak
- petla: brak
- if: brak
- zmienna: nn -> TAK
- tablica: brak
- sortowanie: brak
- rekurencja: brak
- stos: brak
- kolejka: brak
- drzewo: brak

**Oczekiwany wynik** (skorygowany):
```
a) Slowa zaczynajace sie na samogloske: 2 (algorytm, abecadlo)
b) Slowa dluzsze niz 5 znakow: 8 (algorytm, programowanie, abecadlo, zmienna, tablica, sortowanie, rekurencja, kolejka)
c) Slowa z podwojna litera: 1 (zmienna)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: (a) Sprawdz pierwszy znak: `s[0]` i porownaj z samogloskami. (b) Sprawdz `s.length() > 5`.
2. **Podejscie**: Dla (c) przeiteruj po slowie i sprawdz `s[i] == s[i+1]` dla kazdego i.
3. **Kluczowy krok**: Pamietaj o petli do `s.length()-1` (nie `s.length()`) aby nie wyjsc poza zakres.

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

bool samogloska(char c) {
    return c == 'a' || c == 'e' || c == 'i' || c == 'o' || c == 'u' || c == 'y';
}

bool maPodwojna(const string &s) {
    for (int i = 0; i + 1 < (int)s.length(); i++)
        if (s[i] == s[i + 1]) return true;
    return false;
}

int main() {
    ifstream plik("slowa.txt");
    string s;
    int ileA = 0, ileB = 0, ileC = 0;
    vector<string> samogl, dluzsze, podwojne;

    while (plik >> s) {
        if (samogloska(s[0])) { ileA++; samogl.push_back(s); }
        if (s.length() > 5) { ileB++; dluzsze.push_back(s); }
        if (maPodwojna(s)) { ileC++; podwojne.push_back(s); }
    }

    cout << "a) Slowa zaczynajace sie na samogloske: " << ileA << " (";
    for (int i = 0; i < (int)samogl.size(); i++) {
        if (i > 0) cout << ", ";
        cout << samogl[i];
    }
    cout << ")" << endl;

    cout << "b) Slowa dluzsze niz 5 znakow: " << ileB << endl;
    cout << "c) Slowa z podwojna litera: " << ileC << " (";
    for (int i = 0; i < (int)podwojne.size(); i++) {
        if (i > 0) cout << ", ";
        cout << podwojne[i];
    }
    cout << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Trzy niezalezne filtry stosowane w jednym przebiegu pliku. `samogloska()` sprawdza pierwszy znak. `maPodwojna()` szuka par identycznych sasiadow.

Weryfikacja:
- Samogloska na poczatku: algorytm (a), abecadlo (a) = 2
- Dluzsze niz 5: algorytm(8), programowanie(14), abecadlo(8), zmienna(7), tablica(7), sortowanie(10), rekurencja(10), kolejka(7) = 8
- Podwojna litera: zmienna (nn) = 1
</details>

<details>
<summary>Typowe bledy</summary>

- **Petla `i < s.length()` w maPodwojna**: Dostep do `s[i+1]` poza zakresem. CKE: -1 pkt (crash lub UB)
- **Brak `y` w samogloskach polskich**: W jezyku polskim "y" to samogloska — pominiecie jej to blad. CKE: -1 pkt
- **Uzycie `==` na calym stringu zamiast na znaku**: `s == "a"` to nie to samo co `s[0] == 'a'`. CKE: -1 pkt

</details>

---

### Cwiczenie 10.8 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2019 zad. 4 (pary elementow)
**Tagi**: `pary-elementow` `wczytywanie-pliku` `filtrowanie` `podzielnosc`

W pliku `dane.txt` znajduje sie 12 liczb calkowitych (kazda w osobnym wierszu). Napisz program, ktory:
a) Zliczy ile par sasiadujacych elementow (T[i], T[i+1]) ma te same parzystosc (oba parzyste lub oba nieparzyste).
b) Zliczy ile par sasiadujacych elementow ma sume podzielna przez 5.

**Dane** (`dane.txt`):
```
14
27
35
22
18
41
60
55
32
49
10
73
```

**Oczekiwany wynik**:
```
a) Pary o tej samej parzystosci: 4
   (27,35) - oba nieparzyste
   (22,18) - oba parzyste
   (55,32) - nie... 55 nieparzyste
   Korekta:
   (14,27): P+N -> rozne
   (27,35): N+N -> TAK
   (35,22): N+P -> rozne
   (22,18): P+P -> TAK
   (18,41): P+N -> rozne
   (41,60): N+P -> rozne
   (60,55): P+N -> rozne
   (55,32): N+P -> rozne
   (32,49): P+N -> rozne
   (49,10): N+P -> rozne
   (10,73): P+N -> rozne
   Ilosc: 2

b) Pary z suma podzielna przez 5:
   (14,27): 41 -> NIE
   (27,35): 62 -> NIE
   (35,22): 57 -> NIE
   (22,18): 40 -> TAK
   (18,41): 59 -> NIE
   (41,60): 101 -> NIE
   (60,55): 115 -> TAK
   (55,32): 87 -> NIE
   (32,49): 81 -> NIE
   (49,10): 59 -> NIE
   (10,73): 83 -> NIE
   Ilosc: 2
```

**Oczekiwany wynik** (czytelny):
```
a) Pary o tej samej parzystosci: 2
   (27,35) - oba nieparzyste
   (22,18) - oba parzyste

b) Pary z suma podzielna przez 5: 2
   (22,18) - suma=40
   (60,55) - suma=115
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wczytaj wszystko do wektora, potem iteruj po parach sasiadow (i, i+1).
2. **Podejscie**: "Ta sama parzystosc" to `T[i] % 2 == T[i+1] % 2`. Suma podzielna przez 5 to `(T[i] + T[i+1]) % 5 == 0`.
3. **Kluczowy krok**: Petla od 0 do n-2 (wlacznie), nie do n-1 — bo sprawdzamy element i+1.

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

    // a)
    cout << "a) Pary o tej samej parzystosci:" << endl;
    int ileA = 0;
    for (int i = 0; i < n - 1; i++) {
        if (T[i] % 2 == T[i + 1] % 2) {
            string typ = (T[i] % 2 == 0) ? "oba parzyste" : "oba nieparzyste";
            cout << "   (" << T[i] << "," << T[i + 1] << ") - " << typ << endl;
            ileA++;
        }
    }
    cout << "   Ilosc: " << ileA << endl;

    // b)
    cout << endl << "b) Pary z suma podzielna przez 5:" << endl;
    int ileB = 0;
    for (int i = 0; i < n - 1; i++) {
        int s = T[i] + T[i + 1];
        if (s % 5 == 0) {
            cout << "   (" << T[i] << "," << T[i + 1] << ") - suma=" << s << endl;
            ileB++;
        }
    }
    cout << "   Ilosc: " << ileB << endl;
    return 0;
}
```

**Wyjasnienie**: Iterujemy po parach sasiadow. Parzystosc sprawdzamy operatorem modulo. Suma podzielna przez 5 to warunek `(a+b) % 5 == 0`.

Weryfikacja:
- Ta sama parzystosc: (27,35)=N+N, (22,18)=P+P -> 2
- Suma % 5 == 0: (22,18)=40, (60,55)=115 -> 2
</details>

<details>
<summary>Typowe bledy</summary>

- **Petla do `n` zamiast `n-1`**: Dostep do `T[n]` poza zakresem. CKE: -1 pkt (crash)
- **Sprawdzanie `T[i] % 2 == 0 && T[i+1] % 2 == 0`**: To znajdzie tylko pary parzyste, pomijajac pary nieparzyste. CKE: -1 pkt
- **Zapomnienie o liczbach ujemnych**: `(-3) % 2` moze zwrocic -1 w C++. Dla tego zestawu nie ma ujemnych, ale generalnie lepiej porownywac `% 2` obu stron.

</details>

---

### Cwiczenie 10.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3 (rekordy z grupowaniem)
**Tagi**: `map-zliczanie` `grupowanie` `struct` `wczytywanie-pliku`

W pliku `uczniowie.txt` znajduje sie 10 rekordow w formacie: klasa imie ocena (oddzielone spacjami, klasa to string np. "3A", ocena to int 1-6). Napisz program, ktory:
a) Zliczy ile rekordow ma ocene >= 4.
b) Dla kazdej klasy policzy srednia ocen.
c) Poda klase z najwyzsza srednia.

**Dane** (`uczniowie.txt`):
```
3A Jan 5
3A Maria 4
3A Piotr 3
3B Anna 6
3B Tomek 5
3B Ewa 4
3C Kacper 2
3C Ola 5
3C Adam 3
3C Zofia 4
```

**Oczekiwany wynik**:
```
a) Uczniowie z ocena >= 4: 6

b) Srednie ocen:
   3A: 4.00
   3B: 5.00
   3C: 3.50

c) Klasa z najwyzsza srednia: 3B (5.00)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dwa kontenery map: jeden na sume ocen, drugi na liczbe uczniow w klasie.
2. **Podejscie**: `suma[klasa] += ocena; ile[klasa]++;` — potem srednia to `suma/ile` (uzyj rzutowania na double).
3. **Kluczowy krok**: `map<string,int>` automatycznie sortuje klucze leksykograficznie (3A < 3B < 3C).

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <map>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("uczniowie.txt");
    string klasa, imie;
    int ocena;
    map<string, int> suma, ile;
    int dobre = 0;

    while (plik >> klasa >> imie >> ocena) {
        if (ocena >= 4) dobre++;
        suma[klasa] += ocena;
        ile[klasa]++;
    }

    cout << "a) Uczniowie z ocena >= 4: " << dobre << endl;

    cout << endl << "b) Srednie ocen:" << endl;
    cout << fixed << setprecision(2);
    string bestKlasa;
    double bestSr = 0;
    for (auto &p : suma) {
        double sr = (double)p.second / ile[p.first];
        cout << "   " << p.first << ": " << sr << endl;
        if (sr > bestSr) { bestSr = sr; bestKlasa = p.first; }
    }

    cout << endl << "c) Klasa z najwyzsza srednia: " << bestKlasa
         << " (" << bestSr << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Dwie mapy: `suma[klasa]` gromadzi sume ocen, `ile[klasa]` liczbe uczniow. Srednia = suma / ile. Najwyzsza srednia szukana liniowo.

Weryfikacja:
- 3A: 5+4+3=12, 3 uczniow, sr=4.00
- 3B: 6+5+4=15, 3 uczniow, sr=5.00
- 3C: 2+5+3+4=14, 4 uczniow, sr=3.50
- Ocena >= 4: Jan(5), Maria(4), Anna(6), Tomek(5), Ewa(4), Ola(5), Zofia(4) -> 7? Sprawdzmy: 5>=4 TAK, 4>=4 TAK, 3>=4 NIE, 6>=4 TAK, 5>=4 TAK, 4>=4 TAK, 2>=4 NIE, 5>=4 TAK, 3>=4 NIE, 4>=4 TAK -> 7

Korekta oczekiwanego wyniku:
```
a) Uczniowie z ocena >= 4: 7
```
</details>

<details>
<summary>Typowe bledy</summary>

- **Dzielenie calkowite zamiast zmiennoprzecinkowego**: `12 / 3 = 4` jest OK, ale `14 / 4 = 3` (nie 3.5). Uzyj `(double)suma / ile`. CKE: -1 pkt
- **Zapomnienie o klasach z jednym uczniem**: Srednia wtedy rowna ocenie — nie pomijaj ich. CKE: -1 pkt
- **Bledne wczytywanie formatu**: Jesli imie zawiera spacje, `plik >> imie` nie wystarczy — ale w tym zadaniu imiona sa jednoczlonowe.

</details>

---

### Cwiczenie 10.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 4 (wielokryterialne zliczanie)
**Tagi**: `map-zliczanie` `podwojna-petla` `sortowanie` `filtrowanie` `struct`

W pliku `zamowienia.txt` znajduje sie 12 rekordow w formacie: klient produkt ilosc cena_jednostkowa (oddzielone spacjami). Napisz program, ktory:
a) Zliczy laczna wartosc zamowien (ilosc * cena) dla kazdego klienta.
b) Poda klienta z najwieksza laczna wartoscia.
c) Zliczy ile roznych produktow zamowil kazdy klient.
d) Znajdzie produkt zamawiany przez najwiecej roznych klientow.

**Dane** (`zamowienia.txt`):
```
Kowalski Laptop 1 3500
Kowalski Mysz 2 50
Nowak Laptop 1 3500
Nowak Klawiatura 1 200
Nowak Monitor 2 1200
Wisniewski Mysz 5 50
Wisniewski Klawiatura 2 200
Kowalski Monitor 1 1200
Nowak Mysz 3 50
Wisniewski Laptop 1 3500
Kowalski Klawiatura 1 200
Wisniewski Monitor 1 1200
```

**Oczekiwany wynik**:
```
a) Wartosc zamowien:
   Kowalski: 5100
   Nowak: 5250
   Wisniewski: 5450

b) Klient z max wartoscia: Wisniewski (5450)

c) Rozne produkty:
   Kowalski: 4
   Nowak: 4
   Wisniewski: 4

d) Produkt zamawiany przez najwiecej klientow: Laptop (3 klientow)
   Rowniez: Mysz (3 klientow)
   Rowniez: Klawiatura (3 klientow)
   Rowniez: Monitor (3 klientow)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wczytaj rekordy do vectora structow, potem przetwarzaj wielokrotnie.
2. **Podejscie**: (a) `map<string,int>` wartosc_klienta. (c) `map<string, set<string>>` produkty_klienta. (d) `map<string, set<string>>` klienci_produktu.
3. **Kluczowy krok**: `set` automatycznie usuwa duplikaty — rozmiar seta to liczba unikalnych elementow.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <map>
#include <set>
#include <vector>
using namespace std;

int main() {
    ifstream plik("zamowienia.txt");
    string klient, produkt;
    int ilosc, cena;

    map<string, int> wartosc;
    map<string, set<string>> prodKlienta;
    map<string, set<string>> kliProduktu;

    while (plik >> klient >> produkt >> ilosc >> cena) {
        wartosc[klient] += ilosc * cena;
        prodKlienta[klient].insert(produkt);
        kliProduktu[produkt].insert(klient);
    }

    // a)
    cout << "a) Wartosc zamowien:" << endl;
    string bestKli; int bestVal = 0;
    for (auto &p : wartosc) {
        cout << "   " << p.first << ": " << p.second << endl;
        if (p.second > bestVal) { bestVal = p.second; bestKli = p.first; }
    }

    // b)
    cout << endl << "b) Klient z max wartoscia: " << bestKli
         << " (" << bestVal << ")" << endl;

    // c)
    cout << endl << "c) Rozne produkty:" << endl;
    for (auto &p : prodKlienta)
        cout << "   " << p.first << ": " << p.second.size() << endl;

    // d)
    int maxKli = 0;
    for (auto &p : kliProduktu)
        if ((int)p.second.size() > maxKli) maxKli = p.second.size();

    cout << endl << "d) Produkt zamawiany przez najwiecej klientow: ";
    bool first = true;
    for (auto &p : kliProduktu) {
        if ((int)p.second.size() == maxKli) {
            if (!first) cout << "   Rowniez: ";
            cout << p.first << " (" << maxKli << " klientow)" << endl;
            first = false;
        }
    }
    return 0;
}
```

**Wyjasnienie**: Trzy mapy robia ciezka prace: `wartosc` sumuje kwoty, `prodKlienta` zbiera unikalne produkty per klient, `kliProduktu` zbiera unikalnych klientow per produkt. `set` gwarantuje unikalnosc.

Weryfikacja:
- Kowalski: 1*3500 + 2*50 + 1*1200 + 1*200 = 3500+100+1200+200 = 5000. Sprawdzmy: Laptop 3500, Mysz 100, Monitor 1200, Klawiatura 200 = 5000
- Nowak: 3500 + 200 + 2400 + 150 = 6250? Laptop 3500, Klawiatura 200, Monitor 2*1200=2400, Mysz 3*50=150 = 6250. Hmm, to nie zgadza sie z oczekiwanym wynikiem.

Korekta weryfikacji:
- Kowalski: 3500 + 100 + 1200 + 200 = 5000
- Nowak: 3500 + 200 + 2400 + 150 = 6250
- Wisniewski: 250 + 400 + 3500 + 1200 = 5350

Korekta oczekiwanego wyniku:
```
a) Wartosc zamowien:
   Kowalski: 5000
   Nowak: 6250
   Wisniewski: 5350

b) Klient z max wartoscia: Nowak (6250)
```
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie o mnozeniu ilosc*cena**: Dodawanie samej ceny zamiast wartosci. CKE: -2 pkt (zly wynik we wszystkich podpunktach)
- **Brak `set` — uzycie wektora bez usuwania duplikatow**: Jesli klient zamowil ten sam produkt dwa razy, vector go zliczy dwukrotnie. CKE: -1 pkt
- **Bledne wczytywanie**: Format "klient produkt ilosc cena" wymaga dokladnie 4 operacji >> na wiersz. CKE: -1 pkt

</details>

---

## Samoocena

| Poziom | Opis | Kryteria |
|--------|------|----------|
| Podstawowy | Rozumiem zliczanie z jednym warunkiem | Cwiczenia 1, 6 bez pomocy |
| Dobry | Radze sobie z tablica czestotliwosci i map | Cwiczenia 2-4 bez pomocy |
| Bardzo dobry | Umiem filtrowac rekordy i zliczac pary | Cwiczenia 5, 7-8 bez pomocy |
| Doskonaly | Radze sobie z grupowaniem i wielokryterialnym zliczaniem | Cwiczenia 9-10 bez pomocy |

**Co dalej?**
- Jesli masz problem z wczytywaniem plikow -> patrz `cheatsheet_cpp.md` sekcja "Wczytywanie danych"
- Jesli chcesz cwiczycmapy -> patrz cwiczenia `09_zlozone.md` i `20_sql_group_by.md`
- Jesli chcesz cwiczys filtrowanie -> patrz `07_cyfry_liczby.md` i `08_napisy.md`
