# 09. Przetwarzanie zlozone (wieloetapowe)

Typ zadania: **zlozone**
Czestotliwosc: 4/11 lat | Laczna punktacja: 24 pkt
Kategoria: IMPLEMENTACJA

## Umiejetnosci cwiczone w tym zestawie

`map-zliczanie` `grupowanie` `wiele-plikow` `JOIN-tabel` `vector-par` `sortowanie` `wieloetapowe` `struct` `iomanip` `trojne-petle`

---

### Cwiczenie 9.1 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 4
**Tagi**: `map-zliczanie` `grupowanie` `wczytywanie-pliku`

W pliku `osoby.txt` znajduje sie 12 rekordow w formacie: imie miasto (oddzielone spacja, kazdy rekord w osobnym wierszu). Napisz program, ktory:
a) Wypisze ile osob pochodzi z kazdego miasta.
b) Poda miasto, z ktorego pochodzi najwiecej osob.

**Dane** (`osoby.txt`):
```
Anna Krakow
Jan Warszawa
Ewa Krakow
Piotr Gdansk
Maria Warszawa
Tomek Krakow
Kasia Gdansk
Adam Warszawa
Ola Krakow
Marek Poznan
Zofia Warszawa
Pawel Gdansk
```

**Oczekiwany wynik**:
```
a) Liczba osob z kazdego miasta:
   Gdansk: 3
   Krakow: 4
   Poznan: 1
   Warszawa: 4

b) Miasto z najwieksza liczba osob: Krakow (4)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Uzyj `map<string, int>` do zliczania wystapien kazdego miasta.
2. **Podejscie**: Wczytuj pary (imie, miasto), inkrementuj `mapa[miasto]++`.
3. **Kluczowy krok**: `map` automatycznie sortuje klucze alfabetycznie. Iteruj po mapie szukajac max.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <map>
using namespace std;

int main() {
    ifstream plik("osoby.txt");
    string imie, miasto;
    map<string, int> zlicz;

    while (plik >> imie >> miasto) {
        zlicz[miasto]++;
    }

    cout << "a) Liczba osob z kazdego miasta:" << endl;
    string maxMiasto;
    int maxIle = 0;
    for (auto &p : zlicz) {
        cout << "   " << p.first << ": " << p.second << endl;
        if (p.second > maxIle) {
            maxIle = p.second;
            maxMiasto = p.first;
        }
    }

    cout << endl << "b) Miasto z najwieksza liczba osob: "
         << maxMiasto << " (" << maxIle << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: Uzywamy `map<string, int>` do zliczania wystapien kazdego miasta. Mapa automatycznie sortuje klucze alfabetycznie. Nastepnie szukamy miasta o maksymalnej wartosci.

Weryfikacja:
- Krakow: Anna, Ewa, Tomek, Ola = 4
- Warszawa: Jan, Maria, Adam, Zofia = 4
- Gdansk: Piotr, Kasia, Pawel = 3
- Poznan: Marek = 1

Uwaga: Krakow i Warszawa maja po 4 osoby. Map iteruje alfabetycznie, wiec Krakow zostanie znaleziony jako pierwszy z maxIle=4.
</details>

<details>
<summary>Typowe bledy</summary>

- **Uzycie `unordered_map` gdy wymagane sortowanie**: `unordered_map` nie gwarantuje kolejnosci kluczy. CKE: -1 pkt jesli wymaga alfabetycznej
- **Zapomnienie o inicjalizacji mapy**: `map[klucz]++` automatycznie tworzy wpis z wartoscia 0 i inkrementuje. OK.
- **Porownanie stringow z `==` przy szukaniu max**: To poprawne — `==` na stringach dziala w C++. CKE: -0 pkt

</details>

---

### Cwiczenie 9.2 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 4
**Tagi**: `map-zliczanie` `grupowanie` `iomanip` `wczytywanie-pliku`

W pliku `produkty.txt` znajduje sie 10 rekordow w formacie: nazwa kategoria cena (oddzielone spacja). Napisz program, ktory:
a) Obliczy srednia cene w kazdej kategorii.
b) Znajdzie najdrozszy produkt w kazdej kategorii.

**Dane** (`produkty.txt`):
```
Chleb Pieczywo 5
Bulka Pieczywo 2
Mleko Nabialy 4
Ser Nabialy 12
Jogurt Nabialy 3
Jablko Owoce 3
Banan Owoce 6
Maslo Nabialy 8
Bagietka Pieczywo 4
Gruszka Owoce 4
```

**Oczekiwany wynik**:
```
a) Srednia cena w kazdej kategorii:
   Nabialy: 6.75
   Owoce: 4.33
   Pieczywo: 3.67

b) Najdrozszy produkt w kazdej kategorii:
   Nabialy: Ser (12)
   Owoce: Banan (6)
   Pieczywo: Chleb (5)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Grupuj ceny i nazwy wg kategorii uzywajac `map`.
2. **Podejscie**: `map<string, vector<int>>` dla cen, `map<string, pair<string,int>>` dla max.
3. **Kluczowy krok**: Srednia = suma / ilosc. Uzywaj `fixed << setprecision(2)` do formatowania.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <map>
#include <vector>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("produkty.txt");
    string nazwa, kategoria;
    int cena;

    map<string, vector<int>> cenywKat;
    map<string, pair<string, int>> maxwKat;

    while (plik >> nazwa >> kategoria >> cena) {
        cenywKat[kategoria].push_back(cena);
        if (maxwKat.find(kategoria) == maxwKat.end() || cena > maxwKat[kategoria].second) {
            maxwKat[kategoria] = {nazwa, cena};
        }
    }

    cout << "a) Srednia cena w kazdej kategorii:" << endl;
    cout << fixed << setprecision(2);
    for (auto &p : cenywKat) {
        double suma = 0;
        for (int c : p.second) suma += c;
        cout << "   " << p.first << ": " << suma / p.second.size() << endl;
    }

    cout << endl << "b) Najdrozszy produkt w kazdej kategorii:" << endl;
    for (auto &p : maxwKat) {
        cout << "   " << p.first << ": " << p.second.first
             << " (" << p.second.second << ")" << endl;
    }
    return 0;
}
```

**Wyjasnienie**: Grupujemy ceny i produkty wg kategorii uzywajac `map`. Dla sredniej sumujemy ceny w kazdej grupie i dzielimy przez liczbe produktow. Dla max sledzony jest najdrozszy produkt w kazdej kategorii.

Weryfikacja:
- Nabialy: Mleko(4), Ser(12), Jogurt(3), Maslo(8) -> suma=27, srednia=27/4=6.75, max=Ser(12)
- Owoce: Jablko(3), Banan(6), Gruszka(4) -> suma=13, srednia=13/3=4.33, max=Banan(6)
- Pieczywo: Chleb(5), Bulka(2), Bagietka(4) -> suma=11, srednia=11/3=3.67, max=Chleb(5)
</details>

<details>
<summary>Typowe bledy</summary>

- **Dzielenie calkowite zamiast zmiennoprzecinkowego**: `27/4 = 6` zamiast 6.75. Rzutuj na double: `suma / (double)rozmiar`. CKE: -1 pkt
- **Brak `setprecision`**: Domyslne formatowanie double moze dac zbyt wiele miejsc po przecinku. CKE: -0 pkt (ale nieczytelne)
- **Uzycie `map[kat]` bez sprawdzenia czy istnieje**: Przy `maxwKat[kat].second` jesli wpis nie istnieje, tworzy domyslny (pusty string, 0). CKE: -1 pkt

</details>

---

### Cwiczenie 9.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 7 (laczenie tabel)
**Tagi**: `wiele-plikow` `JOIN-tabel` `map-zliczanie` `iomanip`

Dane sa dwie tabele. Tabela A (`uczniowie.txt`) zawiera 6 rekordow w formacie: id imie. Tabela B (`wyniki.txt`) zawiera 10 rekordow w formacie: id punkty. Napisz program, ktory:
a) Dopisze wyniki do uczniow (JOIN po id) i wypisze je.
b) Obliczy sredni wynik dla kazdego ucznia.
c) Znajdzie uczniow, ktorzy nie maja zadnego wyniku.

**Dane** (`uczniowie.txt`):
```
1 Anna
2 Jan
3 Ewa
4 Piotr
5 Maria
6 Tomek
```

**Dane** (`wyniki.txt`):
```
1 85
2 72
1 90
3 68
2 88
4 95
1 78
3 74
4 82
2 65
```

**Oczekiwany wynik**:
```
a) Wyniki uczniow:
   Anna: 85 90 78
   Jan: 72 88 65
   Ewa: 68 74
   Piotr: 95 82

b) Srednie wyniki:
   Anna: 84.33
   Jan: 75.00
   Ewa: 71.00
   Piotr: 88.50

c) Uczniowie bez wynikow:
   Maria (id=5)
   Tomek (id=6)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wczytaj obie tabele do osobnych map. Polacz po kluczu `id`.
2. **Podejscie**: `map<int, string>` dla uczniow, `map<int, vector<int>>` dla wynikow. JOIN = iteracja po uczniach i sprawdzenie wynikow.
3. **Kluczowy krok**: Uczniowie bez wynikow to ci, ktorych id nie ma w mapie wynikow — uzyj `count()`.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <map>
#include <vector>
#include <iomanip>
using namespace std;

int main() {
    ifstream f1("uczniowie.txt"), f2("wyniki.txt");
    map<int, string> uczniowie;
    map<int, vector<int>> wyniki;

    int id; string imie;
    while (f1 >> id >> imie) {
        uczniowie[id] = imie;
    }
    int pkt;
    while (f2 >> id >> pkt) {
        wyniki[id].push_back(pkt);
    }

    cout << "a) Wyniki uczniow:" << endl;
    for (auto &u : uczniowie) {
        if (wyniki.count(u.first)) {
            cout << "   " << u.second << ": ";
            for (int p : wyniki[u.first]) cout << p << " ";
            cout << endl;
        }
    }

    cout << endl << "b) Srednie wyniki:" << endl;
    cout << fixed << setprecision(2);
    for (auto &u : uczniowie) {
        if (wyniki.count(u.first)) {
            double suma = 0;
            for (int p : wyniki[u.first]) suma += p;
            cout << "   " << u.second << ": "
                 << suma / wyniki[u.first].size() << endl;
        }
    }

    cout << endl << "c) Uczniowie bez wynikow:" << endl;
    for (auto &u : uczniowie) {
        if (!wyniki.count(u.first)) {
            cout << "   " << u.second << " (id=" << u.first << ")" << endl;
        }
    }
    return 0;
}
```

**Wyjasnienie**: Wczytujemy obie tabele do osobnych map: uczniowie (id->imie), wyniki (id->lista punktow). JOIN realizujemy przegladajac uczniow i szukajac ich wynikow po id. Uczniowie bez wynikow to ci, ktorych id nie wystepuje w mapie wynikow.

Weryfikacja:
- Anna (id=1): 85, 90, 78 -> srednia = 253/3 = 84.33
- Jan (id=2): 72, 88, 65 -> srednia = 225/3 = 75.00
- Ewa (id=3): 68, 74 -> srednia = 142/2 = 71.00
- Piotr (id=4): 95, 82 -> srednia = 177/2 = 88.50
- Maria (id=5): brak wynikow
- Tomek (id=6): brak wynikow
</details>

<details>
<summary>Typowe bledy</summary>

- **Otwarcie obu plikow tym samym obiektem ifstream**: Trzeba osobne obiekty `ifstream f1, f2` lub zamknac i otworzyc ponownie. CKE: -1 pkt
- **Zakladanie ze wyniki sa w kolejnosci id**: Nie sa — id 1 pojawia sie w wierszach 1, 3, 7. CKE: -1 pkt
- **Brak sprawdzenia `wyniki.count(id)`**: Proba dostepu do wynikow nieistniejacego ucznia tworzy pusty wpis w mapie. CKE: -1 pkt

</details>

---

### Cwiczenie 9.4 (trudnosc: srednie, ~5 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 4.3 (trojki dzielnikowe)
**Tagi**: `trojne-petle` `wieloetapowe` `wczytywanie-pliku`

W pliku `dane.txt` znajduje sie 15 liczb calkowitych dodatnich (kazda w osobnym wierszu). Napisz program, ktory znajdzie wszystkie trojki indeksow (i, j, k) takie, ze i < j < k oraz T[k] dzieli sie przez T[i] i T[k] dzieli sie przez T[j] (indeksy numerowane od 1). Podaj liczbe takich trojek i wypisz pierwsze 5.

**Dane** (`dane.txt`):
```
2
3
4
6
5
12
8
9
24
10
36
15
18
7
48
```

**Oczekiwany wynik**:
```
Trojki (i,j,k) gdzie T[k] % T[i] == 0 i T[k] % T[j] == 0:
(1,2,4): T[1]=2, T[2]=3, T[4]=6 (6%2=0, 6%3=0)
(1,2,6): T[1]=2, T[2]=3, T[6]=12 (12%2=0, 12%3=0)
(1,3,6): T[1]=2, T[3]=4, T[6]=12 (12%2=0, 12%4=0)
(1,2,9): T[1]=2, T[2]=3, T[9]=24 (24%2=0, 24%3=0)
(1,3,9): T[1]=2, T[3]=4, T[9]=24 (24%2=0, 24%4=0)
...
Laczna liczba trojek: 93
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Potrojna petla po indeksach i < j < k.
2. **Podejscie**: Dla kazdej trojki sprawdz dwa warunki podzielnosci: `T[k] % T[i] == 0` i `T[k] % T[j] == 0`.
3. **Kluczowy krok**: Zlozonosc O(n^3) — dla n=15 to 455 trojek do sprawdzenia. Wypisz tylko 5 pierwszych.

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

    cout << "Trojki (i,j,k) gdzie T[k] % T[i] == 0 i T[k] % T[j] == 0:" << endl;
    int ile = 0;
    for (int i = 0; i < n; i++) {
        for (int j = i + 1; j < n; j++) {
            for (int k = j + 1; k < n; k++) {
                if (T[k] % T[i] == 0 && T[k] % T[j] == 0) {
                    ile++;
                    if (ile <= 5) {
                        cout << "(" << i+1 << "," << j+1 << "," << k+1 << "): "
                             << "T[" << i+1 << "]=" << T[i] << ", "
                             << "T[" << j+1 << "]=" << T[j] << ", "
                             << "T[" << k+1 << "]=" << T[k]
                             << " (" << T[k] << "%" << T[i] << "=0, "
                             << T[k] << "%" << T[j] << "=0)" << endl;
                    }
                }
            }
        }
    }
    cout << "..." << endl;
    cout << "Laczna liczba trojek: " << ile << endl;
    return 0;
}
```

**Wyjasnienie**: Potrojna petla O(n^3) przegladajaca wszystkie trojki indeksow i<j<k. Dla kazdej trojki sprawdzamy dwa warunki podzielnosci. Wypisujemy pierwsze 5 trojek i laczna liczbe.

Weryfikacja pierwszych trojek:
- (1,2,4): T={2,3,6}: 6%2=0, 6%3=0 -> TAK
- (1,2,6): T={2,3,12}: 12%2=0, 12%3=0 -> TAK
- (1,3,6): T={2,4,12}: 12%2=0, 12%4=0 -> TAK
- (1,2,9): T={2,3,24}: 24%2=0, 24%3=0 -> TAK
- (1,3,9): T={2,4,24}: 24%2=0, 24%4=0 -> TAK
</details>

<details>
<summary>Typowe bledy</summary>

- **Indeksowanie od 1 w kodzie**: Tablice w C++ sa 0-indeksowane. `T[i]` zamiast `T[i-1]`. CKE: -1 pkt
- **Brak warunku `i < j < k`**: Petla `for(j = 0; ...)` zamiast `for(j = i+1; ...)` liczy te same trojki wielokrotnie. CKE: -2 pkt
- **Dzielenie przez zero**: Jesli T[i] = 0, `T[k] % T[i]` powoduje blad. Tu dane sa dodatnie, ale warto sprawdzic. CKE: -1 pkt

</details>

---

### Cwiczenie 9.5 (trudnosc: trudne, ~6 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 7 (przetwarzanie wieloplikowe)
**Tagi**: `wiele-plikow` `JOIN-tabel` `map-zliczanie` `grupowanie` `iomanip`

Dane sa 3 tabele (podane inline w kodzie lub w osobnych plikach):

**Tabela `uczestnicy.txt`** (id imie):
```
1 Anna
2 Jan
3 Ewa
4 Piotr
5 Maria
```

**Tabela `kursy.txt`** (id nazwa):
```
101 Matematyka
102 Fizyka
103 Informatyka
```

**Tabela `zapisy.txt`** (ucz_id kurs_id ocena):
```
1 101 4
1 102 5
1 103 4
2 101 3
2 103 5
3 101 5
3 102 4
3 103 5
4 102 3
5 101 4
5 102 5
5 103 3
```

Napisz program, ktory:
a) Obliczy srednia ocene kazdego uczestnika.
b) Znajdzie kurs z najwyzsza srednia ocena.
c) Wypisze uczestnikow zapisanych na wszystkie 3 kursy.

**Oczekiwany wynik**:
```
a) Srednia ocena kazdego uczestnika:
   Anna: 4.33
   Jan: 4.00
   Ewa: 4.67
   Piotr: 3.00
   Maria: 4.00

b) Kurs z najwyzsza srednia ocena:
   Fizyka (srednia: 4.25)

c) Uczestnicy zapisani na wszystkie 3 kursy:
   Anna (kursy: 3)
   Ewa (kursy: 3)
   Maria (kursy: 3)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy mapy: uczniowie, kursy, zapisy. Polacz je po kluczach id.
2. **Podejscie**: Agreguj oceny wg uczestnika (srednia) i wg kursu (srednia). Zlicz kursy kazdego uczestnika za pomoca `set`.
3. **Kluczowy krok**: Uczestnik na wszystkich kursach = `kursyUcz[id].size() == liczbaKursow`.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <map>
#include <vector>
#include <set>
#include <iomanip>
using namespace std;

int main() {
    ifstream f1("uczestnicy.txt");
    map<int, string> uczestnicy;
    int id; string imie;
    while (f1 >> id >> imie) uczestnicy[id] = imie;

    ifstream f2("kursy.txt");
    map<int, string> kursy;
    string nazwa;
    while (f2 >> id >> nazwa) kursy[id] = nazwa;

    int liczbaKursow = kursy.size();

    ifstream f3("zapisy.txt");
    map<int, vector<int>> ocenyUcz;
    map<int, vector<int>> ocenyKurs;
    map<int, set<int>> kursyUcz;
    int uid, kid, ocena;
    while (f3 >> uid >> kid >> ocena) {
        ocenyUcz[uid].push_back(ocena);
        ocenyKurs[kid].push_back(ocena);
        kursyUcz[uid].insert(kid);
    }

    cout << "a) Srednia ocena kazdego uczestnika:" << endl;
    cout << fixed << setprecision(2);
    for (auto &u : uczestnicy) {
        double suma = 0;
        for (int o : ocenyUcz[u.first]) suma += o;
        cout << "   " << u.second << ": "
             << suma / ocenyUcz[u.first].size() << endl;
    }

    cout << endl << "b) Kurs z najwyzsza srednia ocena:" << endl;
    double maxSr = 0;
    string maxKurs;
    for (auto &k : kursy) {
        double suma = 0;
        for (int o : ocenyKurs[k.first]) suma += o;
        double sr = suma / ocenyKurs[k.first].size();
        if (sr > maxSr) { maxSr = sr; maxKurs = k.second; }
    }
    cout << "   " << maxKurs << " (srednia: " << maxSr << ")" << endl;

    cout << endl << "c) Uczestnicy zapisani na wszystkie " << liczbaKursow << " kursy:" << endl;
    for (auto &u : uczestnicy) {
        if ((int)kursyUcz[u.first].size() == liczbaKursow) {
            cout << "   " << u.second << " (kursy: " << kursyUcz[u.first].size() << ")" << endl;
        }
    }
    return 0;
}
```

**Wyjasnienie**: Trzy tabele laczony sa przez klucze id. Uzywamy map do przechowywania danych i grupowania ocen wg uczestnika i kursu. Zbior kursow kazdego uczestnika pozwala sprawdzic, kto jest zapisany na wszystkie.

Weryfikacja:
a) Srednie:
- Anna (id=1): 4+5+4=13, 13/3=4.33
- Jan (id=2): 3+5=8, 8/2=4.00
- Ewa (id=3): 5+4+5=14, 14/3=4.67
- Piotr (id=4): 3, 3/1=3.00
- Maria (id=5): 4+5+3=12, 12/3=4.00

b) Srednie kursow:
- Matematyka (101): Anna(4)+Jan(3)+Ewa(5)+Maria(4)=16, 16/4=4.00
- Fizyka (102): Anna(5)+Ewa(4)+Piotr(3)+Maria(5)=17, 17/4=4.25
- Informatyka (103): Anna(4)+Jan(5)+Ewa(5)+Maria(3)=17, 17/4=4.25
Fizyka i Informatyka maja ta sama srednia 4.25. Fizyka zostanie znaleziona jako pierwsza.

c) Uczestnicy na wszystkich 3 kursach:
- Anna: {101,102,103} -> 3 kursy -> TAK
- Jan: {101,103} -> 2 kursy -> NIE
- Ewa: {101,102,103} -> 3 kursy -> TAK
- Piotr: {102} -> 1 kurs -> NIE
- Maria: {101,102,103} -> 3 kursy -> TAK
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomieszanie kluczy ucz_id i kurs_id**: Oba to int, latwo pomylic ktory jest ktory. CKE: -2 pkt
- **Brak `set` do sledzenia unikalnych kursow**: Uzywajac `vector` zamiast `set`, ten sam kurs moze byc policzony wielokrotnie. CKE: -1 pkt
- **Zapomnienie o zamknieciu plikow**: `ifstream` zamyka plik automatycznie w destruktorze. Nie jest bledem, ale dobra praktyka.

</details>

---

### Cwiczenie 9.6 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 5 (sortowanie rekordow)
**Tagi**: `struct` `sortowanie` `wczytywanie-pliku` `grupowanie`

W pliku `wyniki.txt` znajduje sie 10 rekordow w formacie: imie punkty (oddzielone spacja). Napisz program, ktory:
a) Wypisze rekordy posortowane malejaco wg punktow.
b) Wypisze rekordy posortowane rosnaco wg imienia (alfabetycznie).
c) Poda srednia punktow i ile osob ma wynik powyzej sredniej.

**Dane** (`wyniki.txt`):
```
Anna 85
Jan 72
Ewa 90
Piotr 65
Maria 88
Tomek 70
Kasia 95
Adam 78
Ola 82
Marek 60
```

**Oczekiwany wynik**:
```
a) Wg punktow (malejaco):
   Kasia: 95
   Ewa: 90
   Maria: 88
   Anna: 85
   Ola: 82
   Adam: 78
   Jan: 72
   Tomek: 70
   Piotr: 65
   Marek: 60

b) Wg imienia (rosnaco):
   Adam: 78
   Anna: 85
   Ewa: 90
   Jan: 72
   Kasia: 95
   Maria: 88
   Marek: 60
   Ola: 82
   Piotr: 65
   Tomek: 70

c) Srednia: 78.50, powyzej sredniej: 5
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Wczytaj dane do `vector<pair<string, int>>` lub `vector<struct>`.
2. **Podejscie**: `sort` z custom komparatorem: (a) `[](auto &a, auto &b) { return a.second > b.second; }`, (b) `a.first < b.first`.
3. **Kluczowy krok**: Pamietaj ze `sort` modyfikuje wektor — zrob kopie przed drugim sortowaniem lub sortuj dwa razy.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <vector>
#include <algorithm>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("wyniki.txt");
    vector<pair<string, int>> dane;
    string imie; int pkt;
    while (plik >> imie >> pkt) dane.push_back({imie, pkt});

    // a) Malejaco wg punktow
    auto kopia = dane;
    sort(kopia.begin(), kopia.end(), [](auto &a, auto &b) {
        return a.second > b.second;
    });
    cout << "a) Wg punktow (malejaco):" << endl;
    for (auto &p : kopia)
        cout << "   " << p.first << ": " << p.second << endl;

    // b) Rosnaco wg imienia
    sort(kopia.begin(), kopia.end(), [](auto &a, auto &b) {
        return a.first < b.first;
    });
    cout << endl << "b) Wg imienia (rosnaco):" << endl;
    for (auto &p : kopia)
        cout << "   " << p.first << ": " << p.second << endl;

    // c) Srednia
    double suma = 0;
    for (auto &p : dane) suma += p.second;
    double sr = suma / dane.size();
    int powyzej = 0;
    for (auto &p : dane) if (p.second > sr) powyzej++;
    cout << endl << "c) Srednia: " << fixed << setprecision(2) << sr
         << ", powyzej sredniej: " << powyzej << endl;
    return 0;
}
```

Weryfikacja:
- Suma: 85+72+90+65+88+70+95+78+82+60 = 785, srednia = 78.50
- Powyzej 78.50: Kasia(95), Ewa(90), Maria(88), Anna(85), Ola(82) = 5
</details>

<details>
<summary>Typowe bledy</summary>

- **Sortowanie malejace z `<` zamiast `>`**: Daje rosnace zamiast malejacego. CKE: -1 pkt
- **Zapomnienie o kopii wektora**: `sort` modyfikuje in-place. Bez kopii drugie sortowanie traci oryginalna kolejnosc. CKE: -0 pkt (jesli nie potrzeba)
- **Porownanie `>=` zamiast `>` przy sredniej**: "Powyzej sredniej" to scisle wieksze. CKE: -1 pkt

</details>

---

### Cwiczenie 9.7 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 5 (macierz + statystyki)
**Tagi**: `wieloetapowe` `grupowanie` `struct` `sortowanie`

W pliku `oceny.txt` znajduje sie 8 rekordow w formacie: imie przedmiot ocena. Uczen moze miec kilka ocen z roznych przedmiotow. Napisz program, ktory:
a) Dla kazdego ucznia wypisze wszystkie przedmioty i oceny.
b) Wypisze uczniow posortowanych wg sredniej oceny (malejaco).
c) Znajdzie przedmiot z najwyzsza srednia ocen.

**Dane** (`oceny.txt`):
```
Anna Matematyka 5
Anna Fizyka 4
Jan Matematyka 3
Jan Informatyka 5
Ewa Matematyka 4
Ewa Fizyka 5
Ewa Informatyka 5
Jan Fizyka 2
```

**Oczekiwany wynik**:
```
a) Oceny uczniow:
   Anna: Matematyka(5), Fizyka(4) -> srednia: 4.50
   Ewa: Matematyka(4), Fizyka(5), Informatyka(5) -> srednia: 4.67
   Jan: Matematyka(3), Informatyka(5), Fizyka(2) -> srednia: 3.33

b) Ranking wg sredniej:
   1. Ewa (4.67)
   2. Anna (4.50)
   3. Jan (3.33)

c) Przedmiot z najwyzsza srednia:
   Informatyka (srednia: 5.00)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Grupuj po uczniu i po przedmiocie osobno.
2. **Podejscie**: `map<string, vector<pair<string,int>>>` dla uczniow (imie -> lista (przedmiot, ocena)).
3. **Kluczowy krok**: Aby posortowac uczniow wg sredniej, przenies dane do wektora par (srednia, imie) i posortuj malejaco.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <map>
#include <vector>
#include <algorithm>
#include <iomanip>
using namespace std;

int main() {
    ifstream plik("oceny.txt");
    string imie, przedmiot;
    int ocena;
    map<string, vector<pair<string,int>>> uczniowie;
    map<string, vector<int>> przedmioty;

    while (plik >> imie >> przedmiot >> ocena) {
        uczniowie[imie].push_back({przedmiot, ocena});
        przedmioty[przedmiot].push_back(ocena);
    }

    cout << fixed << setprecision(2);
    cout << "a) Oceny uczniow:" << endl;
    vector<pair<double, string>> ranking;
    for (auto &u : uczniowie) {
        cout << "   " << u.first << ": ";
        double suma = 0;
        for (int i = 0; i < u.second.size(); i++) {
            if (i > 0) cout << ", ";
            cout << u.second[i].first << "(" << u.second[i].second << ")";
            suma += u.second[i].second;
        }
        double sr = suma / u.second.size();
        cout << " -> srednia: " << sr << endl;
        ranking.push_back({sr, u.first});
    }

    sort(ranking.begin(), ranking.end(), [](auto &a, auto &b) {
        return a.first > b.first;
    });
    cout << endl << "b) Ranking wg sredniej:" << endl;
    for (int i = 0; i < ranking.size(); i++) {
        cout << "   " << i+1 << ". " << ranking[i].second
             << " (" << ranking[i].first << ")" << endl;
    }

    cout << endl << "c) Przedmiot z najwyzsza srednia:" << endl;
    double maxSr = 0;
    string maxPrz;
    for (auto &p : przedmioty) {
        double suma = 0;
        for (int o : p.second) suma += o;
        double sr = suma / p.second.size();
        if (sr > maxSr) { maxSr = sr; maxPrz = p.first; }
    }
    cout << "   " << maxPrz << " (srednia: " << maxSr << ")" << endl;
    return 0;
}
```

Weryfikacja:
- Anna: Mat(5), Fiz(4) -> 9/2 = 4.50
- Jan: Mat(3), Inf(5), Fiz(2) -> 10/3 = 3.33
- Ewa: Mat(4), Fiz(5), Inf(5) -> 14/3 = 4.67
- Ranking: Ewa(4.67), Anna(4.50), Jan(3.33)
- Przedmioty: Matematyka (5+3+4)/3=4.00, Fizyka (4+5+2)/3=3.67, Informatyka (5+5)/2=5.00
- Max: Informatyka (5.00)
</details>

<details>
<summary>Typowe bledy</summary>

- **Sortowanie stringa zamiast double**: `sort` na `pair<string, double>` sortuje po stringu. Trzeba `pair<double, string>`. CKE: -1 pkt
- **Zapomnienie o odwroceniu kolejnosci sortowania**: Domyslne `sort` jest rosnace. Dla malejacego uzyj `>` w komparatorze. CKE: -1 pkt
- **Pomylenie sredniej ucznia ze srednia przedmiotu**: To dwie rozne agregacje na tych samych danych. CKE: -2 pkt

</details>

---

### Cwiczenie 9.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 4 + 2025 zad. 7 (analiza logów)
**Tagi**: `wieloetapowe` `map-zliczanie` `wczytywanie-pliku` `sortowanie`

W pliku `logi.txt` znajduje sie 15 wpisow w formacie: godzina uzytkownik akcja (oddzielone spacjami). Napisz program, ktory:
a) Zliczy ile akcji wykonal kazdy uzytkownik.
b) Znajdzie godzine z najwieksza liczba akcji.
c) Wypisze uzytkownikow, ktorzy wykonali akcje "LOGIN" i "LOGOUT" (obie).

**Dane** (`logi.txt`):
```
08 Anna LOGIN
08 Jan LOGIN
09 Anna VIEW
09 Ewa LOGIN
10 Anna EDIT
10 Jan VIEW
11 Ewa VIEW
11 Jan LOGOUT
12 Anna LOGOUT
12 Ewa EDIT
13 Ewa LOGOUT
13 Anna LOGIN
14 Anna VIEW
14 Jan LOGIN
15 Jan LOGOUT
```

**Oczekiwany wynik**:
```
a) Akcje wg uzytkownika:
   Anna: 6 akcji
   Ewa: 4 akcje
   Jan: 5 akcji

b) Godzina z najwieksza liczba akcji:
   08, 09, 10, 11, 12, 13, 14 (po 2 akcje kazda)

c) Uzytkownicy z LOGIN i LOGOUT:
   Anna (LOGIN: 2, LOGOUT: 1)
   Ewa (LOGIN: 1, LOGOUT: 1)
   Jan (LOGIN: 2, LOGOUT: 2)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy mapy: `uzytkownik -> ilosc`, `godzina -> ilosc`, `uzytkownik -> set<akcja>`.
2. **Podejscie**: Wczytuj trojki, inkrementuj odpowiednie mapy.
3. **Kluczowy krok**: Uzytkownik ma obie akcje jesli `akcje[user].count("LOGIN") && akcje[user].count("LOGOUT")`.

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
using namespace std;

int main() {
    ifstream plik("logi.txt");
    string godz, user, akcja;
    map<string, int> cntUser, cntGodz;
    map<string, map<string, int>> akcjeUser;

    while (plik >> godz >> user >> akcja) {
        cntUser[user]++;
        cntGodz[godz]++;
        akcjeUser[user][akcja]++;
    }

    cout << "a) Akcje wg uzytkownika:" << endl;
    for (auto &p : cntUser)
        cout << "   " << p.first << ": " << p.second << " akcji" << endl;

    cout << endl << "b) Godzina z najwieksza liczba akcji:" << endl;
    int maxG = 0;
    for (auto &p : cntGodz) if (p.second > maxG) maxG = p.second;
    cout << "   ";
    bool first = true;
    for (auto &p : cntGodz) {
        if (p.second == maxG) {
            if (!first) cout << ", ";
            cout << p.first;
            first = false;
        }
    }
    cout << " (po " << maxG << " akcje kazda)" << endl;

    cout << endl << "c) Uzytkownicy z LOGIN i LOGOUT:" << endl;
    for (auto &u : akcjeUser) {
        if (u.second.count("LOGIN") && u.second.count("LOGOUT")) {
            cout << "   " << u.first << " (LOGIN: " << u.second["LOGIN"]
                 << ", LOGOUT: " << u.second["LOGOUT"] << ")" << endl;
        }
    }
    return 0;
}
```

Weryfikacja:
- Anna: LOGIN, VIEW, EDIT, LOGOUT, LOGIN, VIEW = 6
- Ewa: LOGIN, VIEW, EDIT, LOGOUT = 4
- Jan: LOGIN, VIEW, LOGOUT, LOGIN, LOGOUT = 5
- Godziny: 08(2), 09(2), 10(2), 11(2), 12(2), 13(2), 14(2), 15(1) -> max 2 (wszystkie oprocz 15)
- LOGIN+LOGOUT: Anna(L:2,LO:1), Ewa(L:1,LO:1), Jan(L:2,LO:2) -> wszyscy
</details>

<details>
<summary>Typowe bledy</summary>

- **Wczytywanie godziny jako `int`**: Godziny "08", "09" sa poprawne jako string, ale `int 08` tez dziala. CKE: -0 pkt
- **Uzycie `set` zamiast `map` dla akcji**: Set nie zlicza ile razy — jesli potrzebujesz ile, uzyj `map<string, int>`. CKE: -1 pkt
- **Brak obslugi wielu godzin z tym samym max**: Jesli kilka godzin ma max, wypisz wszystkie. CKE: -1 pkt

</details>

---

### Cwiczenie 9.9 (trudnosc: srednie-trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 4 + 2025 zad. 7 (transakcje)
**Tagi**: `wieloetapowe` `wiele-plikow` `JOIN-tabel` `sortowanie` `iomanip`

Dane sa 2 tabele: klienci (`klienci.txt`: id imie miasto) i zamowienia (`zamowienia.txt`: id_klienta produkt kwota). Napisz program, ktory:
a) Obliczy laczna kwote zamowien kazdego klienta.
b) Obliczy laczna kwote zamowien z kazdego miasta.
c) Znajdzie klienta z najwyzsza srednia kwota zamowienia.

**Dane** (`klienci.txt`):
```
1 Anna Krakow
2 Jan Warszawa
3 Ewa Krakow
4 Piotr Gdansk
```

**Dane** (`zamowienia.txt`):
```
1 Laptop 3500
2 Telefon 2000
1 Monitor 1200
3 Tablet 1500
2 Sluchawki 300
4 Laptop 4000
3 Drukarka 800
1 Klawiatura 200
4 Monitor 1500
2 Laptop 3800
```

**Oczekiwany wynik**:
```
a) Laczna kwota wg klienta:
   Anna: 4900
   Jan: 6100
   Ewa: 2300
   Piotr: 5500

b) Laczna kwota wg miasta:
   Gdansk: 5500
   Krakow: 7200
   Warszawa: 6100

c) Klient z najwyzsza srednia kwota:
   Piotr (srednia: 2750.00, zamowienia: 2)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: JOIN klientow z zamowieniami po id. Grupuj kwoty wg klienta i wg miasta.
2. **Podejscie**: Mapa klientow (id -> {imie, miasto}). Mapa zamowien (id_klienta -> lista kwot).
3. **Kluczowy krok**: Kwota wg miasta = suma kwot klientow z danego miasta. Srednia = suma / ilosc zamowien.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <map>
#include <vector>
#include <iomanip>
using namespace std;

int main() {
    ifstream f1("klienci.txt"), f2("zamowienia.txt");
    map<int, pair<string, string>> klienci; // id -> (imie, miasto)
    int id;
    string imie, miasto;
    while (f1 >> id >> imie >> miasto) klienci[id] = {imie, miasto};

    map<int, vector<int>> zamowienia;
    string produkt;
    int kid, kwota;
    while (f2 >> kid >> produkt >> kwota) zamowienia[kid].push_back(kwota);

    // a) Wg klienta
    cout << "a) Laczna kwota wg klienta:" << endl;
    for (auto &k : klienci) {
        int suma = 0;
        for (int q : zamowienia[k.first]) suma += q;
        cout << "   " << k.second.first << ": " << suma << endl;
    }

    // b) Wg miasta
    cout << endl << "b) Laczna kwota wg miasta:" << endl;
    map<string, int> kwotaMiasto;
    for (auto &k : klienci) {
        int suma = 0;
        for (int q : zamowienia[k.first]) suma += q;
        kwotaMiasto[k.second.second] += suma;
    }
    for (auto &m : kwotaMiasto)
        cout << "   " << m.first << ": " << m.second << endl;

    // c) Najwyzsza srednia
    cout << endl << "c) Klient z najwyzsza srednia kwota:" << endl;
    double maxSr = 0;
    string maxImie;
    int maxZam = 0;
    for (auto &k : klienci) {
        if (zamowienia[k.first].empty()) continue;
        double suma = 0;
        for (int q : zamowienia[k.first]) suma += q;
        double sr = suma / zamowienia[k.first].size();
        if (sr > maxSr) {
            maxSr = sr;
            maxImie = k.second.first;
            maxZam = zamowienia[k.first].size();
        }
    }
    cout << "   " << maxImie << " (srednia: " << fixed << setprecision(2)
         << maxSr << ", zamowienia: " << maxZam << ")" << endl;
    return 0;
}
```

Weryfikacja:
- Anna (id=1): 3500+1200+200 = 4900
- Jan (id=2): 2000+300+3800 = 6100
- Ewa (id=3): 1500+800 = 2300
- Piotr (id=4): 4000+1500 = 5500
- Krakow (Anna+Ewa): 4900+2300 = 7200
- Warszawa (Jan): 6100
- Gdansk (Piotr): 5500
- Srednie: Anna=4900/3=1633, Jan=6100/3=2033, Ewa=2300/2=1150, Piotr=5500/2=2750
- Max: Piotr (2750.00)
</details>

<details>
<summary>Typowe bledy</summary>

- **Wczytanie "Laptop" jako kwota**: Format to `id_klienta produkt kwota` — trzeba wczytac string miedzy nimi. CKE: -2 pkt
- **Pomieszanie id klienta z id zamowienia**: Tu nie ma id zamowienia — kazdy wiersz to oddzielne zamowienie. CKE: -1 pkt
- **Niezamkniecie pliku przed otwarciem drugiego**: Nie blad w C++ (osobne obiekty), ale w Pascalu bylo. CKE: -0 pkt

</details>

---

### Cwiczenie 9.10 (trudnosc: trudne, ~6 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 7 (rozszerzony)
**Tagi**: `wieloetapowe` `wiele-plikow` `JOIN-tabel` `grupowanie` `sortowanie` `iomanip`

Dane sa 3 tabele: zawodnicy (`zawodnicy.txt`: id imie kraj), zawody (`zawody.txt`: id nazwa), wyniki (`wyniki.txt`: zaw_id zaw_id2 czas). Czas jest w sekundach (int). Napisz program, ktory:
a) Obliczy najlepszy (najkrotszy) czas kazdego zawodnika.
b) Wypisze ranking zawodnikow wg najlepszego czasu (rosnaco).
c) Obliczy sredni najlepszy czas dla kazdego kraju.

**Dane** (`zawodnicy.txt`):
```
1 Kowalski POL
2 Nowak POL
3 Mueller GER
4 Schmidt GER
5 Smith USA
```

**Dane** (`zawody.txt`):
```
101 Sprint
102 Maraton
```

**Dane** (`wyniki.txt`):
```
1 101 12
1 102 240
2 101 11
2 102 235
3 101 13
3 102 250
4 101 12
4 102 238
5 101 10
5 102 230
```

**Oczekiwany wynik**:
```
a) Najlepszy czas kazdego zawodnika:
   Kowalski: 12s (Sprint)
   Nowak: 11s (Sprint)
   Mueller: 13s (Sprint)
   Schmidt: 12s (Sprint)
   Smith: 10s (Sprint)

b) Ranking wg najlepszego czasu:
   1. Smith (10s)
   2. Nowak (11s)
   3. Kowalski (12s)
   4. Schmidt (12s)
   5. Mueller (13s)

c) Sredni najlepszy czas wg kraju:
   GER: 12.50s
   POL: 11.50s
   USA: 10.00s
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: JOIN trzech tabel. Dla kazdego zawodnika znajdz min czas.
2. **Podejscie**: Mapa min czasow: `map<int, int>` (id -> min czas). Potem JOIN z zawodnikami.
3. **Kluczowy krok**: Aby znalezc nazwe zawodow dla min czasu, potrzebna jest dodatkowa mapa `id -> (zaw_id2, czas)`.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <map>
#include <vector>
#include <algorithm>
#include <iomanip>
using namespace std;

int main() {
    ifstream f1("zawodnicy.txt"), f2("zawody.txt"), f3("wyniki.txt");
    map<int, pair<string, string>> zaw; // id -> (imie, kraj)
    int id;
    string imie, kraj;
    while (f1 >> id >> imie >> kraj) zaw[id] = {imie, kraj};

    map<int, string> zawody; // id -> nazwa
    string nazwa;
    while (f2 >> id >> nazwa) zawody[id] = nazwa;

    map<int, int> minCzas;
    map<int, int> minZawody;
    int zid, zawid, czas;
    while (f3 >> zid >> zawid >> czas) {
        if (!minCzas.count(zid) || czas < minCzas[zid]) {
            minCzas[zid] = czas;
            minZawody[zid] = zawid;
        }
    }

    // a) Najlepszy czas
    cout << "a) Najlepszy czas kazdego zawodnika:" << endl;
    for (auto &z : zaw) {
        cout << "   " << z.second.first << ": " << minCzas[z.first]
             << "s (" << zawody[minZawody[z.first]] << ")" << endl;
    }

    // b) Ranking
    vector<pair<int, string>> ranking;
    for (auto &z : zaw) ranking.push_back({minCzas[z.first], z.second.first});
    sort(ranking.begin(), ranking.end());

    cout << endl << "b) Ranking wg najlepszego czasu:" << endl;
    for (int i = 0; i < ranking.size(); i++) {
        cout << "   " << i+1 << ". " << ranking[i].second
             << " (" << ranking[i].first << "s)" << endl;
    }

    // c) Sredni najlepszy czas wg kraju
    map<string, vector<int>> krajCzasy;
    for (auto &z : zaw) krajCzasy[z.second.second].push_back(minCzas[z.first]);

    cout << endl << "c) Sredni najlepszy czas wg kraju:" << endl;
    cout << fixed << setprecision(2);
    for (auto &k : krajCzasy) {
        double suma = 0;
        for (int c : k.second) suma += c;
        cout << "   " << k.first << ": " << suma / k.second.size() << "s" << endl;
    }
    return 0;
}
```

Weryfikacja:
- Kowalski: min(12, 240) = 12 (Sprint)
- Nowak: min(11, 235) = 11 (Sprint)
- Mueller: min(13, 250) = 13 (Sprint)
- Schmidt: min(12, 238) = 12 (Sprint)
- Smith: min(10, 230) = 10 (Sprint)
- Ranking: Smith(10), Nowak(11), Kowalski(12), Schmidt(12), Mueller(13)
- POL: (12+11)/2=11.50, GER: (13+12)/2=12.50, USA: 10/1=10.00
</details>

<details>
<summary>Typowe bledy</summary>

- **Szukanie min bez inicjalizacji**: `minCzas[id]` tworzy wpis z wartoscia 0. Trzeba sprawdzic czy klucz istnieje. CKE: -2 pkt
- **Sortowanie po imieniu zamiast po czasie**: `pair<string, int>` sortuje po stringu. Uzyj `pair<int, string>`. CKE: -1 pkt
- **Pomieszanie "sredni czas" ze "sredni najlepszy czas"**: Srednia najlepszych czasow != srednia wszystkich czasow. CKE: -2 pkt
- **Brak JOIN z tabela zawodow**: Aby podac nazwe zawodow, trzeba polaczyc po `zawid`. CKE: -1 pkt

</details>

---

## Samoocena

| Poziom | Opis | Zakres cwiczen |
|--------|------|---------------|
| Podstawowy | Zliczam z map, grupuje dane z pliku | 1-3 bez pomocy |
| Dobry | Lacze dwa pliki (JOIN), uzywam sort z komparatorem | 4-6 bez pomocy |
| Bardzo dobry | Wieloetapowe analizy, ranking, agregacje po kilku kluczach | 7-8 bez pomocy |
| Doskonaly | JOIN trzech tabel, zlozone agregacje, rozne perspektywy danych | 9-10 bez pomocy |

**Co dalej?**
- Jesli masz problemy z cwiczeniami 1-3, wrocz do `cheatsheet_cpp.md` — sekcja "map i vector"
- Jesli opanowales 1-6, przejdz do cwiczen `10_zliczanie.md` lub `11_minmax.md`
- Jesli rozwiazales 9-10, jestes gotowy na prawdziwe zadania maturalne — przejdz do `rozwiazania_wzorcowe/`
