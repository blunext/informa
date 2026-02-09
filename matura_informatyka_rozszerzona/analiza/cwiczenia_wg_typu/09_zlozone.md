# 09. Przetwarzanie zlozone (wieloetapowe)

Typ zadania: **zlozone**
Czestotliwosc: 4/11 lat | Laczna punktacja: 24 pkt
Kategoria: IMPLEMENTACJA

---

### Cwiczenie 9.1 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 4

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

---

### Cwiczenie 9.2 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 4

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

---

### Cwiczenie 9.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 7 (laczenie tabel)

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

---

### Cwiczenie 9.4 (trudnosc: srednie, ~5 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 4.3 (trojki dzielnikowe)

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

---

### Cwiczenie 9.5 (trudnosc: trudne, ~6 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 7 (przetwarzanie wieloplikowe)

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
    // Wczytanie uczestnikow
    ifstream f1("uczestnicy.txt");
    map<int, string> uczestnicy;
    int id; string imie;
    while (f1 >> id >> imie) uczestnicy[id] = imie;

    // Wczytanie kursow
    ifstream f2("kursy.txt");
    map<int, string> kursy;
    string nazwa;
    while (f2 >> id >> nazwa) kursy[id] = nazwa;

    int liczbaKursow = kursy.size();

    // Wczytanie zapisow
    ifstream f3("zapisy.txt");
    map<int, vector<int>> ocenyUcz;  // ucz_id -> lista ocen
    map<int, vector<int>> ocenyKurs; // kurs_id -> lista ocen
    map<int, set<int>> kursyUcz;     // ucz_id -> zbior kursow
    int uid, kid, ocena;
    while (f3 >> uid >> kid >> ocena) {
        ocenyUcz[uid].push_back(ocena);
        ocenyKurs[kid].push_back(ocena);
        kursyUcz[uid].insert(kid);
    }

    // a) Srednia ocena kazdego uczestnika
    cout << "a) Srednia ocena kazdego uczestnika:" << endl;
    cout << fixed << setprecision(2);
    for (auto &u : uczestnicy) {
        double suma = 0;
        for (int o : ocenyUcz[u.first]) suma += o;
        cout << "   " << u.second << ": "
             << suma / ocenyUcz[u.first].size() << endl;
    }

    // b) Kurs z najwyzsza srednia
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

    // c) Uczestnicy na wszystkich kursach
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
Fizyka i Informatyka maja ta sama srednia 4.25. Fizyka zostanie znaleziona pierwsza (id 102 < 103).

Korekta: kurs z najwyzsza srednia to Fizyka lub Informatyka (obie 4.25). W kodzie Fizyka zostanie znaleziona jako pierwsza.

c) Uczestnicy na wszystkich 3 kursach:
- Anna: {101,102,103} -> 3 kursy -> TAK
- Jan: {101,103} -> 2 kursy -> NIE
- Ewa: {101,102,103} -> 3 kursy -> TAK
- Piotr: {102} -> 1 kurs -> NIE
- Maria: {101,102,103} -> 3 kursy -> TAK
</details>

---
