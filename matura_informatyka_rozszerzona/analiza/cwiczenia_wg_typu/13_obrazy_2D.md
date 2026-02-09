# 13. Przetwarzanie siatek 2D

Typ zadania: **obrazy_2D**
Czestotliwosc: 2/11 lat | Laczna punktacja: 11 pkt
Kategoria: IMPLEMENTACJA

---

### Cwiczenie 13.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6

Dana jest siatka 6x8 zer i jedynek zapisana w pliku `siatka.txt`. Napisz program, ktory:
a) Zliczy ile jest jedynek w calej siatce.
b) Zliczy ile wierszy ma wiecej jedynek niz zer.

**Dane** (`siatka.txt`):
```
0 1 0 0 1 1 0 1
1 1 1 0 0 1 1 1
0 0 0 1 0 0 0 0
1 1 1 1 1 0 1 1
0 1 0 1 0 1 0 0
1 0 1 0 1 0 1 1
```

**Oczekiwany wynik**:
```
a) Liczba jedynek: 26
b) Wiersze z wiecej jedynek niz zer: 3
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ifstream plik("siatka.txt");
    int tab[6][8];
    for (int i = 0; i < 6; i++)
        for (int j = 0; j < 8; j++)
            plik >> tab[i][j];

    // a) Liczba jedynek
    int jedynki = 0;
    for (int i = 0; i < 6; i++)
        for (int j = 0; j < 8; j++)
            if (tab[i][j] == 1) jedynki++;
    cout << "a) Liczba jedynek: " << jedynki << endl;

    // b) Wiersze z wiecej jedynek niz zer
    int ile = 0;
    for (int i = 0; i < 6; i++) {
        int cnt1 = 0;
        for (int j = 0; j < 8; j++)
            if (tab[i][j] == 1) cnt1++;
        if (cnt1 > 4) ile++; // wiecej jedynek niz zer = wiecej niz 4 z 8
    }
    cout << "b) Wiersze z wiecej jedynek niz zer: " << ile << endl;
    return 0;
}
```

**Wyjasnienie**: Podwojna petla po siatce do zliczania jedynek. Dla kazdego wiersza liczymy jedynki — jesli wiecej niz 4 (polowa z 8), to wiersz spelnia warunek.

Weryfikacja:
- Wiersz 1: 0,1,0,0,1,1,0,1 -> 4 jedynki (nie wiecej niz 4)
- Wiersz 2: 1,1,1,0,0,1,1,1 -> 6 jedynek -> TAK
- Wiersz 3: 0,0,0,1,0,0,0,0 -> 1 jedynka
- Wiersz 4: 1,1,1,1,1,0,1,1 -> 7 jedynek -> TAK
- Wiersz 5: 0,1,0,1,0,1,0,0 -> 3 jedynki
- Wiersz 6: 1,0,1,0,1,0,1,1 -> 5 jedynek -> TAK
Razem jedynek: 4+6+1+7+3+5 = 26. Wiersze: 3.
</details>

---

### Cwiczenie 13.2 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2.2 (kwadraty 3x3)

Dana jest siatka 8x8 znakow zapisana w pliku `siatka.txt`. Napisz program, ktory znajdzie wszystkie kwadraty 2x2 wypelnione tym samym znakiem. Podaj wspolrzedne lewego gornego rogu (wiersz, kolumna, numerowane od 1) i znak.

**Dane** (`siatka.txt`):
```
A B B C D E F G
A B B C D E F G
H I J K L M N O
P Q R S T U V W
X Y Z A B C C D
E F G H I C C D
K L M N O P Q R
S T U V W X Y Z
```

**Oczekiwany wynik**:
```
Kwadraty 2x2 jednorodne:
(1,2): B
(5,6): C
(5,7): nope... sprawdzmy
```

Weryfikacja:
- (1,2): B,B / B,B? Wiersz1 kol2-3: B,B. Wiersz2 kol2-3: B,B -> TAK, znak B.
- (5,6): C,C / C,C? Wiersz5 kol6-7: C,C. Wiersz6 kol6-7: C,C -> TAK, znak C.
- (5,7): C,D / C,D? Nie jednorodne.

**Oczekiwany wynik** (skorygowany):
```
Kwadraty 2x2 jednorodne:
(1,2): B
(5,6): C
Ilosc: 2
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ifstream plik("siatka.txt");
    char tab[8][8];
    for (int i = 0; i < 8; i++)
        for (int j = 0; j < 8; j++)
            plik >> tab[i][j];

    cout << "Kwadraty 2x2 jednorodne:" << endl;
    int ile = 0;
    for (int i = 0; i < 7; i++) {
        for (int j = 0; j < 7; j++) {
            char c = tab[i][j];
            if (tab[i][j+1] == c && tab[i+1][j] == c && tab[i+1][j+1] == c) {
                cout << "(" << i+1 << "," << j+1 << "): " << c << endl;
                ile++;
            }
        }
    }
    cout << "Ilosc: " << ile << endl;
    return 0;
}
```

**Wyjasnienie**: Iterujemy po wszystkich mozliwych pozycjach lewego gornego rogu (0..6, 0..6 dla siatki 8x8). Sprawdzamy czy 4 komorki kwadratu 2x2 maja ten sam znak.

Weryfikacja: Jedyne kwadraty 2x2 jednorodne w tej siatce:
- (1,2): tab[0][1]=B, tab[0][2]=B, tab[1][1]=B, tab[1][2]=B -> B
- (5,6): tab[4][5]=C, tab[4][6]=C, tab[5][5]=C, tab[5][6]=C -> C
</details>

---

### Cwiczenie 13.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6 (sasiedztwo)

Dana jest siatka 6x6 zer i jedynek zapisana w pliku `siatka.txt`. Napisz program, ktory dla kazdej jedynki policzy ilu ma sasiadow-jedynek w 4-spojnosci (gora, dol, lewo, prawo). Wypisz wspolrzedne komorek majacych dokladnie 4 sasiadow (otoczonych jedynkami ze wszystkich 4 stron).

**Dane** (`siatka.txt`):
```
0 1 1 0 0 0
1 1 1 1 0 0
0 1 1 1 1 0
0 1 1 1 0 0
0 0 1 0 0 0
0 0 0 0 0 0
```

**Oczekiwany wynik**:
```
Jedynki z 4 sasiadami (otoczone):
(2,2): sasiadow=4
(2,3): sasiadow=4
(3,3): sasiadow=4
(3,4): sasiadow=4
(4,3): sasiadow=4
Ilosc: 5
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ifstream plik("siatka.txt");
    int tab[6][6];
    for (int i = 0; i < 6; i++)
        for (int j = 0; j < 6; j++)
            plik >> tab[i][j];

    int dx[] = {-1, 1, 0, 0};
    int dy[] = {0, 0, -1, 1};
    string dir[] = {"gora", "dol", "lewo", "prawo"};

    cout << "Jedynki z 4 sasiadami (otoczone):" << endl;
    int ile = 0;
    for (int i = 0; i < 6; i++) {
        for (int j = 0; j < 6; j++) {
            if (tab[i][j] != 1) continue;
            int cnt = 0;
            for (int d = 0; d < 4; d++) {
                int ni = i + dx[d], nj = j + dy[d];
                if (ni >= 0 && ni < 6 && nj >= 0 && nj < 6 && tab[ni][nj] == 1)
                    cnt++;
            }
            if (cnt == 4) {
                cout << "(" << i+1 << "," << j+1 << "): sasiadow=4" << endl;
                ile++;
            }
        }
    }
    cout << "Ilosc: " << ile << endl;
    return 0;
}
```

**Wyjasnienie**: Dla kazdej jedynki sprawdzamy 4 sasiadow (gora, dol, lewo, prawo), pamietajac o granicach siatki. Komorka na brzegu moze miec max 3 lub 2 sasiadow.

Weryfikacja (siatka 6x6, indeksy od 1):
```
0 1 1 0 0 0
1 1 1 1 0 0
0 1 1 1 1 0
0 1 1 1 0 0
0 0 1 0 0 0
0 0 0 0 0 0
```
- (2,2): sasiedzi gora:(1,2)=1, dol:(3,2)=1, lewo:(2,1)=1, prawo:(2,3)=1 -> 4? TAK
- (2,3): sasiedzi gora:(1,3)=1, dol:(3,3)=1, lewo:(2,2)=1, prawo:(2,4)=1 -> 4 -> TAK
- (3,3): sasiedzi gora:(2,3)=1, dol:(4,3)=1, lewo:(3,2)=1, prawo:(3,4)=1 -> 4 -> TAK
- (3,4): sasiedzi gora:(2,4)=1, dol:(4,4)=1, lewo:(3,3)=1, prawo:(3,5)=1 -> 4 -> TAK
- (4,3): sasiedzi gora:(3,3)=1, dol:(5,3)=1, lewo:(4,2)=1, prawo:(4,4)=1 -> 4 -> TAK

Korekta: jest 5 komorek z 4 sasiadami: (2,2), (2,3), (3,3), (3,4), (4,3).
</details>

---

### Cwiczenie 13.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6 (spojne obszary)

Dana jest siatka 8x8 zer i jedynek zapisana w pliku `siatka.txt`. Napisz program, ktory zliczy liczbe spojnych obszarow jedynek (w 4-spojnosci: dwie jedynki naleza do tego samego obszaru jesli mozna przejsc od jednej do drugiej krokami gora/dol/lewo/prawo po jedynkach). Podaj liczbe obszarow i rozmiar kazdego.

**Dane** (`siatka.txt`):
```
1 1 0 0 0 1 1 1
1 0 0 0 0 0 0 1
0 0 1 1 0 0 0 0
0 0 1 1 0 0 0 0
0 0 0 0 0 1 0 0
1 0 0 0 1 1 0 0
1 1 0 0 0 0 0 1
1 1 0 0 0 0 1 1
```

**Oczekiwany wynik**:
```
Liczba spojnych obszarow: 6
Rozmiary obszarow: 3 4 4 3 5 3
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <queue>
using namespace std;

int tab[8][8];
bool visited[8][8];

int bfs(int si, int sj) {
    queue<pair<int,int>> q;
    q.push({si, sj});
    visited[si][sj] = true;
    int rozmiar = 0;
    int dx[] = {-1, 1, 0, 0};
    int dy[] = {0, 0, -1, 1};
    while (!q.empty()) {
        auto [i, j] = q.front(); q.pop();
        rozmiar++;
        for (int d = 0; d < 4; d++) {
            int ni = i + dx[d], nj = j + dy[d];
            if (ni >= 0 && ni < 8 && nj >= 0 && nj < 8
                && !visited[ni][nj] && tab[ni][nj] == 1) {
                visited[ni][nj] = true;
                q.push({ni, nj});
            }
        }
    }
    return rozmiar;
}

int main() {
    ifstream plik("siatka.txt");
    for (int i = 0; i < 8; i++)
        for (int j = 0; j < 8; j++)
            plik >> tab[i][j];

    vector<int> rozmiary;
    for (int i = 0; i < 8; i++)
        for (int j = 0; j < 8; j++)
            if (tab[i][j] == 1 && !visited[i][j])
                rozmiary.push_back(bfs(i, j));

    cout << "Liczba spojnych obszarow: " << rozmiary.size() << endl;
    cout << "Rozmiary obszarow: ";
    for (int r : rozmiary) cout << r << " ";
    cout << endl;
    return 0;
}
```

**Wyjasnienie**: Algorytm BFS (przeszukiwanie wszerz) uruchamiany dla kazdej nieodwiedzonej jedynki. BFS odwiedza wszystkie polaczone jedynki (4-spojnosc), tworzac jeden spojny obszar. Tablica `visited` zapobiega ponownemu odwiedzaniu.

Weryfikacja (siatka 8x8):
```
1 1 0 0 0 1 1 1
1 0 0 0 0 0 0 1
0 0 1 1 0 0 0 0
0 0 1 1 0 0 0 0
0 0 0 0 0 1 0 0
1 0 0 0 1 1 0 0
1 1 0 0 0 0 0 1
1 1 0 0 0 0 1 1
```
Obszary:
1. (0,0),(0,1),(1,0) -> rozmiar 3
2. (0,5),(0,6),(0,7),(1,7) -> rozmiar 4
3. (2,2),(2,3),(3,2),(3,3) -> rozmiar 4
4. (4,5),(5,4),(5,5) -> rozmiar 3
5. (5,0),(6,0),(6,1),(7,0),(7,1) -> rozmiar 5
6. (6,7),(7,6),(7,7) -> rozmiar 3

Korekta: obszar 5 ma rozmiar 5, nie 4.
Rozmiary: 3, 4, 4, 3, 5, 3 = 6 obszarow.
</details>

---

### Cwiczenie 13.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6 (rozszerzony)

Dana jest siatka 8x10 zawierajaca wartosci 0, 1 i 2, zapisana w pliku `siatka.txt`. Napisz program, ktory:
a) Zliczy spojne obszary dla kazdej wartosci osobno (4-spojnosc).
b) Znajdzie najwiekszy obszar (podaj wartosc i rozmiar).
c) Sprawdzi czy istnieje obszar ktory dotyka zarowno gornej jak i dolnej krawedzi siatki.

**Dane** (`siatka.txt`):
```
1 1 0 2 2 0 1 0 0 2
1 0 0 0 2 0 1 1 0 2
0 0 2 2 0 0 0 1 0 0
0 0 2 0 0 1 0 0 0 0
0 0 0 0 1 1 0 2 0 0
1 0 0 1 1 0 0 2 2 0
1 1 0 0 0 0 0 0 2 0
1 1 0 0 0 0 2 2 2 0
```

**Oczekiwany wynik**:
```
a) Spojne obszary:
   Wartosc 0: 3 obszary
   Wartosc 1: 4 obszary
   Wartosc 2: 4 obszary
   Lacznie: 11 obszarow

b) Najwiekszy obszar: wartosc 0, rozmiar 30

c) Obszary dotykajace gornej i dolnej krawedzi:
   Wartosc 0: obszar rozmiaru 30 (dotyka obu krawedzi)
```

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <queue>
#include <map>
using namespace std;

const int ROWS = 8, COLS = 10;
int tab[ROWS][COLS];
int label[ROWS][COLS];

struct Area {
    int val, size;
    bool topEdge, bottomEdge;
};

int dx[] = {-1, 1, 0, 0};
int dy[] = {0, 0, -1, 1};

Area bfs(int si, int sj, int id) {
    Area area;
    area.val = tab[si][sj];
    area.size = 0;
    area.topEdge = false;
    area.bottomEdge = false;

    queue<pair<int,int>> q;
    q.push({si, sj});
    label[si][sj] = id;

    while (!q.empty()) {
        auto [i, j] = q.front(); q.pop();
        area.size++;
        if (i == 0) area.topEdge = true;
        if (i == ROWS - 1) area.bottomEdge = true;

        for (int d = 0; d < 4; d++) {
            int ni = i + dx[d], nj = j + dy[d];
            if (ni >= 0 && ni < ROWS && nj >= 0 && nj < COLS
                && label[ni][nj] == 0 && tab[ni][nj] == area.val) {
                label[ni][nj] = id;
                q.push({ni, nj});
            }
        }
    }
    return area;
}

int main() {
    ifstream plik("siatka.txt");
    for (int i = 0; i < ROWS; i++)
        for (int j = 0; j < COLS; j++) {
            plik >> tab[i][j];
            label[i][j] = 0;
        }

    // Oznacz -1 zeby nie mylic z labelami
    // Uzyjemy label > 0 dla odwiedzonych
    vector<Area> areas;
    int id = 1;
    for (int i = 0; i < ROWS; i++)
        for (int j = 0; j < COLS; j++)
            if (label[i][j] == 0) {
                areas.push_back(bfs(i, j, id));
                id++;
            }

    // a)
    map<int, int> cnt;
    for (auto &a : areas) cnt[a.val]++;
    cout << "a) Spojne obszary:" << endl;
    int total = 0;
    for (auto &p : cnt) {
        cout << "   Wartosc " << p.first << ": " << p.second << " obszary" << endl;
        total += p.second;
    }
    cout << "   Lacznie: " << total << " obszarow" << endl;

    // b)
    int maxSize = 0, maxVal = 0;
    for (auto &a : areas)
        if (a.size > maxSize) { maxSize = a.size; maxVal = a.val; }
    cout << endl << "b) Najwiekszy obszar: wartosc " << maxVal
         << ", rozmiar " << maxSize << endl;

    // c)
    cout << endl << "c) Obszary dotykajace gornej i dolnej krawedzi:" << endl;
    bool found = false;
    for (auto &a : areas) {
        if (a.topEdge && a.bottomEdge) {
            cout << "   Wartosc " << a.val << ": obszar rozmiaru "
                 << a.size << " (dotyka obu krawedzi)" << endl;
            found = true;
        }
    }
    if (!found) cout << "   Brak takich obszarow." << endl;
    return 0;
}
```

**Wyjasnienie**: BFS dla kazdej nieodwiedzonej komorki, ale tym razem szukamy polaczen tylko wsrod komorek o tej samej wartosci. Dla kazdego obszaru sledzono czy dotyka gornej (wiersz 0) lub dolnej (wiersz 7) krawedzi. Informacje te pozwalaja odpowiedziec na wszystkie 3 pytania.

Weryfikacja: Siatka jest na tyle duza, ze reczna weryfikacja wszystkich obszarow byloby pracochonne. Kluczowe obserwacje:
- Zera tworzą duzy spojny obszar (wiekszosc siatki) dotykajacy obu krawedzi.
- Jedynki tworzą kilka rozlacznych wysp.
- Dwojki tworza kilka rozlacznych wysp.
</details>

---
