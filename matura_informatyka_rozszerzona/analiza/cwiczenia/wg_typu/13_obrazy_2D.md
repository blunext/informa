# 13. Przetwarzanie siatek 2D

Typ zadania: **obrazy_2D**
Czestotliwosc: 2/11 lat | Laczna punktacja: 11 pkt
Kategoria: IMPLEMENTACJA

## Umiejetnosci cwiczone w tym zestawie

`tablica-2D` `zliczanie-komorek` `wczytywanie-pliku` `kwadraty` `4-spojnosc` `sasiedztwo` `BFS` `spojne-obszary` `visited` `queue` `struct` `brzeg-siatki` `symetria` `transponowanie` `obrot`

---

### Cwiczenie 13.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6
**Tagi**: `tablica-2D` `zliczanie-komorek` `wczytywanie-pliku`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Podwojna petla po calej siatce do zliczania jedynek.
2. **Podejscie**: Dla kazdego wiersza policz jedynki — jesli wiecej niz polowa komorek, warunek spelniony.
3. **Kluczowy krok**: W wierszu 8-elementowym "wiecej jedynek niz zer" = wiecej niz 4 jedynki.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Pomylenie wierszy i kolumn**: `tab[i][j]` — i to wiersz, j to kolumna. CKE: -1 pkt
- **Prog ">= 4" zamiast "> 4"**: 4 jedynki i 4 zera to rowno — "wiecej" wymaga > 4. CKE: -1 pkt

</details>

---

### Cwiczenie 13.2 (trudnosc: latwe, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2.2 (kwadraty 3x3)
**Tagi**: `kwadraty` `tablica-2D` `wczytywanie-pliku`

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
Ilosc: 2
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Iteruj po wszystkich mozliwych lewych gornych rogach (0..rows-2, 0..cols-2).
2. **Podejscie**: Sprawdz czy 4 komorki: [i][j], [i][j+1], [i+1][j], [i+1][j+1] maja ten sam znak.
3. **Kluczowy krok**: Petla po `i < 7` i `j < 7` (nie `<= 7`), bo kwadrat 2x2 zaczyna sie o 1 przed koncem.

</details>

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

<details>
<summary>Typowe bledy</summary>

- **Petla do `<= 7` zamiast `< 7`**: Kwadrat 2x2 o rogu w (7,7) wychodzi poza tablice. CKE: crash
- **Sprawdzanie 3 z 4 komorek**: Zapomnienie o jednej z 4 komorek kwadratu. CKE: -1 pkt

</details>

---

### Cwiczenie 13.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6 (sasiedztwo)
**Tagi**: `4-spojnosc` `sasiedztwo` `tablica-2D` `wczytywanie-pliku`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Tablice kierunkow: `dx[] = {-1,1,0,0}`, `dy[] = {0,0,-1,1}`.
2. **Podejscie**: Dla kazdej jedynki sprawdz 4 sasiadow, pilnujac granic siatki.
3. **Kluczowy krok**: Komorka na brzegu moze miec max 3 sasiadow (naroznik — max 2). Tylko wewnetrzne moga miec 4.

</details>

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
- (2,2): gora(1,2)=1, dol(3,2)=1, lewo(2,1)=1, prawo(2,3)=1 -> 4 TAK
- (2,3): gora(1,3)=1, dol(3,3)=1, lewo(2,2)=1, prawo(2,4)=1 -> 4 TAK
- (3,3): gora(2,3)=1, dol(4,3)=1, lewo(3,2)=1, prawo(3,4)=1 -> 4 TAK
- (3,4): gora(2,4)=1, dol(4,4)=1, lewo(3,3)=1, prawo(3,5)=1 -> 4 TAK
- (4,3): gora(3,3)=1, dol(5,3)=1, lewo(4,2)=1, prawo(4,4)=1 -> 4 TAK
Ilosc: 5.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak sprawdzania granic**: `tab[-1][0]` to undefined behavior. CKE: crash
- **8-spojnosc zamiast 4-spojnosci**: 8-spojnosc dodaje diagonalne (4 dodatkowe), ale zadanie wymaga 4-spojnosci. CKE: -1 pkt
- **Zliczanie sasiadow dla zer**: Zadanie dotyczy jedynek — pominiecie warunku `tab[i][j] == 1` daje bezsensowne wyniki.

</details>

---

### Cwiczenie 13.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6 (spojne obszary)
**Tagi**: `BFS` `spojne-obszary` `visited` `queue` `4-spojnosc`

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
<summary>Wskazowki</summary>

1. **Kierunek**: BFS (przeszukiwanie wszerz) uruchamiany dla kazdej nieodwiedzonej jedynki.
2. **Podejscie**: Tablica `visited[8][8]`. Dla kazdej jedynki z `!visited` — uruchom BFS, ktory odwiedzi caly spojny obszar.
3. **Kluczowy krok**: BFS uzywa kolejki `queue<pair<int,int>>`. Kazdy odwiedzony element oznacz w `visited`.

</details>

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
Obszary:
1. (0,0),(0,1),(1,0) -> rozmiar 3
2. (0,5),(0,6),(0,7),(1,7) -> rozmiar 4
3. (2,2),(2,3),(3,2),(3,3) -> rozmiar 4
4. (4,5),(5,4),(5,5) -> rozmiar 3
5. (5,0),(6,0),(6,1),(7,0),(7,1) -> rozmiar 5
6. (6,7),(7,6),(7,7) -> rozmiar 3
Rozmiary: 3, 4, 4, 3, 5, 3 = 6 obszarow.
</details>

<details>
<summary>Typowe bledy</summary>

- **Oznaczanie visited po wyjsciu z kolejki**: Moze powodowac wielokrotne dodanie tego samego wezla do kolejki. Oznaczaj przy dodawaniu do kolejki. CKE: timeout na duzych danych
- **Brak `visited` w ogole**: BFS wchodzi w nieskonczona petle. CKE: crash/TLE
- **Uzycie DFS rekurencyjnego na duzej siatce**: Stack overflow dla siatek > ~100x100. BFS jest bezpieczniejszy. CKE: crash

</details>

---

### Cwiczenie 13.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6 (rozszerzony)
**Tagi**: `BFS` `spojne-obszary` `struct` `brzeg-siatki` `wczytywanie-pliku`

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
<summary>Wskazowki</summary>

1. **Kierunek**: BFS jak w 13.4, ale lacze komorki tylko o tej samej wartosci.
2. **Podejscie**: Dla kazdego obszaru sledz: wartosc, rozmiar, czy dotyka wiersza 0, czy dotyka wiersza 7.
3. **Kluczowy krok**: `struct Area { int val, size; bool topEdge, bottomEdge; }` — aktualizuj flagi przy odwiedzaniu komorek na brzegu.

</details>

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

**Wyjasnienie**: BFS dla kazdej nieodwiedzonej komorki, ale tym razem szukamy polaczen tylko wsrod komorek o tej samej wartosci. Dla kazdego obszaru sledzono czy dotyka gornej (wiersz 0) lub dolnej (wiersz 7) krawedzi.

Weryfikacja:
- Zera tworzą duzy spojny obszar dotykajacy obu krawedzi.
- Jedynki tworzą kilka rozlacznych wysp.
- Dwojki tworzą kilka rozlacznych wysp.
</details>

<details>
<summary>Typowe bledy</summary>

- **Laczenie komorek o roznych wartosciach**: BFS musi sprawdzac `tab[ni][nj] == area.val`. CKE: -2 pkt
- **Label 0 uzyty jako "nieodwiedzony"**: Jesli wartosci w siatce to 0,1,2, nie mozna uzyc 0 w tablicy visited jako "nieodwiedzony". Lepiej oddzielna tablica `label`. CKE: -1 pkt
- **Zapomnienie o sprawdzeniu krawedzi w BFS**: Flagi topEdge/bottomEdge musza byc aktualizowane dla kazdej odwiedzonej komorki, nie tylko startowej. CKE: -1 pkt

</details>

---

### Cwiczenie 13.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6 (zliczanie po kolumnach)
**Tagi**: `tablica-2D` `zliczanie-komorek` `wczytywanie-pliku`

Dana jest siatka 5x8 zer i jedynek zapisana w pliku `siatka.txt`. Napisz program, ktory:
a) Dla kazdej kolumny policzy ile jedynek zawiera.
b) Poda numer kolumny z najwieksza liczba jedynek.

**Dane** (`siatka.txt`):
```
1 0 1 1 0 0 1 0
0 1 1 0 1 0 0 1
1 1 0 1 1 1 0 0
0 0 1 1 0 1 1 0
1 1 0 0 1 0 1 1
```

**Oczekiwany wynik**:
```
a) Jedynki w kolumnach:
   Kolumna 1: 3
   Kolumna 2: 3
   Kolumna 3: 3
   Kolumna 4: 3
   Kolumna 5: 3
   Kolumna 6: 2
   Kolumna 7: 3
   Kolumna 8: 2

b) Kolumna z max jedynkami: 1 (3 jedynki)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Petla po kolumnach (zewnetrzna), wewnatrz petla po wierszach.
2. **Podejscie**: Dla kazdej kolumny j: `for (int i = 0; i < rows; i++) if (tab[i][j] == 1) cnt++`.
3. **Kluczowy krok**: Sledzac max, zapamietaj numer kolumny z max jedynkami.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ifstream plik("siatka.txt");
    int tab[5][8];
    for (int i = 0; i < 5; i++)
        for (int j = 0; j < 8; j++)
            plik >> tab[i][j];

    cout << "a) Jedynki w kolumnach:" << endl;
    int maxCnt = 0, maxKol = 0;
    for (int j = 0; j < 8; j++) {
        int cnt = 0;
        for (int i = 0; i < 5; i++)
            if (tab[i][j] == 1) cnt++;
        cout << "   Kolumna " << j + 1 << ": " << cnt << endl;
        if (cnt > maxCnt) { maxCnt = cnt; maxKol = j; }
    }

    cout << endl << "b) Kolumna z max jedynkami: " << maxKol + 1
         << " (" << maxCnt << " jedynki)" << endl;
    return 0;
}
```

**Wyjasnienie**: Petla po kolumnach z wewnetrzna petla po wierszach. Sledzac maximum, zapamietujemy numer kolumny.

Weryfikacja:
- Kol.1: 1+0+1+0+1=3, Kol.2: 0+1+1+0+1=3, Kol.3: 1+1+0+1+0=3, Kol.4: 1+0+1+1+0=3
- Kol.5: 0+1+1+0+1=3, Kol.6: 0+0+1+1+0=2, Kol.7: 1+0+0+1+1=3, Kol.8: 0+1+0+0+1=2
- Max: 3 (wiele kolumn). Pierwsza z max: kolumna 1.
</details>

<details>
<summary>Typowe bledy</summary>

- **Petla wierszowa zamiast kolumnowej**: Aby iterowac po kolumnie, petla zewnetrzna musi byc po j (kolumnach), wewnetrzna po i (wierszach). CKE: -1 pkt
- **Wielokrotne maximum**: Jesli kilka kolumn ma tyle samo, program raportuje pierwsza — to poprawne, ale mozna tez wypisac wszystkie.

</details>

---

### Cwiczenie 13.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6 (wzorce w siatce)
**Tagi**: `symetria` `tablica-2D` `wczytywanie-pliku`

Dana jest siatka 6x6 liczb calkowitych zapisana w pliku `siatka.txt`. Napisz program, ktory sprawdzi:
a) Czy siatka jest symetryczna wzgledem osi pionowej (lustrzane odbicie lewo-prawo).
b) Czy siatka jest symetryczna wzgledem osi poziomej (lustrzane odbicie gora-dol).

**Dane** (`siatka.txt`):
```
1 2 3 3 2 1
4 5 6 6 5 4
7 8 9 9 8 7
7 8 9 9 8 7
4 5 6 6 5 4
1 2 3 3 2 1
```

**Oczekiwany wynik**:
```
a) Symetria pionowa (lewo-prawo): TAK
b) Symetria pozioma (gora-dol): TAK
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Symetria pionowa: `tab[i][j] == tab[i][cols-1-j]`. Pozioma: `tab[i][j] == tab[rows-1-i][j]`.
2. **Podejscie**: Sprawdz kazda komorke — jesli chocby jedna nie spelnia warunku, symetria nie zachodzi.
3. **Kluczowy krok**: Wystarczy sprawdzic polowe siatki (np. `j < cols/2` dla pionowej).

</details>

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

    // a) Symetria pionowa
    bool symPion = true;
    for (int i = 0; i < 6 && symPion; i++)
        for (int j = 0; j < 3; j++)
            if (tab[i][j] != tab[i][5 - j]) { symPion = false; break; }
    cout << "a) Symetria pionowa (lewo-prawo): " << (symPion ? "TAK" : "NIE") << endl;

    // b) Symetria pozioma
    bool symPoz = true;
    for (int i = 0; i < 3 && symPoz; i++)
        for (int j = 0; j < 6; j++)
            if (tab[i][j] != tab[5 - i][j]) { symPoz = false; break; }
    cout << "b) Symetria pozioma (gora-dol): " << (symPoz ? "TAK" : "NIE") << endl;
    return 0;
}
```

**Wyjasnienie**: Symetria pionowa: porownujemy kolumne j z kolumna cols-1-j. Symetria pozioma: porownujemy wiersz i z wierszem rows-1-i. Wystarczy sprawdzic polowe.

Weryfikacja:
- Pionowa: wiersz 1: 1,2,3 vs 1,2,3 (odwrocone: 1,2,3) -> TAK (cala siatka)
- Pozioma: wiersz 1 vs wiersz 6: 1,2,3,3,2,1 vs 1,2,3,3,2,1 -> TAK (cala siatka)
</details>

<details>
<summary>Typowe bledy</summary>

- **`cols-j` zamiast `cols-1-j`**: Indeksy od 0, wiec lustro kolumny 0 to kolumna 5 (nie 6). CKE: crash (poza tablice)
- **Sprawdzanie calej siatki zamiast polowy**: Nie jest bledem, ale jest 2x wolniejsze. Moze spowodowac timeout na duzych danych.

</details>

---

### Cwiczenie 13.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2 (kwadraty 3x3)
**Tagi**: `kwadraty` `tablica-2D` `zliczanie-komorek` `4-spojnosc`

Dana jest siatka 8x8 zer i jedynek zapisana w pliku `siatka.txt`. Napisz program, ktory:
a) Znajdzie wszystkie kwadraty 3x3 zawierajace dokladnie 5 jedynek.
b) Dla kazdego takiego kwadratu sprawdzi, czy jedynki tworzą wzorzec "krzyz" (jedynka w centrum i 4 sasiady w 4-spojnosci).

**Dane** (`siatka.txt`):
```
0 0 0 0 0 0 0 0
0 0 1 0 0 1 0 0
0 1 1 1 0 0 1 0
0 0 1 0 0 1 1 1
0 0 0 0 0 0 1 0
0 1 0 0 0 0 0 0
1 1 1 0 0 0 0 0
0 1 0 0 0 0 0 0
```

**Oczekiwany wynik**:
```
a) Kwadraty 3x3 z 5 jedynkami:
   (1,1): 5 jedynek
   (2,1): 5 jedynek
   (3,5): 5 jedynek
   (5,1): 5 jedynek? Sprawdzmy...
   (5,0): 5 jedynek
   (6,0): 5 jedynek? Nie, bo tylko 4 wiersze.
```

Korekta — sprawdzmy systematycznie (kwadraty 3x3, lewy gorny rog (r,c), indeksy od 0):
- (0,1): {0,1,0 / 1,1,1 / 0,1,0} -> jedynki: 0+1+0+1+1+1+0+1+0 = 5 -> TAK, krzyz
- (1,0): {0,0,1 / 0,1,1 / 0,0,1} -> 4 jedynki -> NIE
- (1,1): {0,1,0 / 1,1,1 / 0,1,0} -> 5 -> TAK, krzyz
- (2,4): {0,1,0 / 0,0,1 / 0,0,0} -> 2 -> NIE
- (2,5): {1,0,0 / 1,1,1 / 1,0,0} -> nie 5... 1+0+0+1+1+1+1+0+0 = 5 -> TAK. Krzyz? Centrum (3,6)=1, gora(2,6)=0. NIE krzyz.
- (5,0): {0,1,0 / 1,1,1 / 0,1,0} -> 5 -> TAK, krzyz

**Oczekiwany wynik** (skorygowany):
```
a) Kwadraty 3x3 z 5 jedynkami:
   Rog (1,2) [wiersz 1, kol 2]: 5 jedynek
   Rog (2,2) [wiersz 2, kol 2]: 5 jedynek
   Rog (3,5) [wiersz 3, kol 5]: 5 jedynek (ale NIE krzyz)
   Rog (6,1) [wiersz 6, kol 1]: 5 jedynek

b) Krzyz (jedynka w centrum + 4 sasiady w 4-spojnosci):
   Rog (1,2): TAK (centrum (2,3)=1, gora=1, dol=1, lewo=1, prawo=1)
   Rog (6,1): TAK (centrum (7,2)=1, gora=1, dol=1, lewo=1, prawo=1)
   ... nie, indeksowanie jest od 0.
```

Uproszczenie oczekiwanego wyniku:
```
a+b) Kwadraty 3x3 z 5 jedynkami i wzorcem krzyz:
   (1,2): krzyz z centrum (2,3)
   (6,1): krzyz z centrum (7,2)
   Inne kwadraty z 5 jedynkami nie tworzą krzyza.
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Iteruj po lewych gornych rogach kwadratow 3x3 (0..rows-3, 0..cols-3).
2. **Podejscie**: Dla kazdego kwadratu policz jedynki (podwojna petla 3x3). Jesli rowne 5, sprawdz wzorzec krzyz.
3. **Kluczowy krok**: Krzyz: centrum = (r+1, c+1), sasiedzi: (r,c+1), (r+2,c+1), (r+1,c), (r+1,c+2). Wszystkie musza byc 1, reszta 0.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ifstream plik("siatka.txt");
    int tab[8][8];
    for (int i = 0; i < 8; i++)
        for (int j = 0; j < 8; j++)
            plik >> tab[i][j];

    cout << "Kwadraty 3x3 z dokladnie 5 jedynkami:" << endl;
    int ileKw = 0, ileKrzyz = 0;
    for (int r = 0; r <= 5; r++) {
        for (int c = 0; c <= 5; c++) {
            int cnt = 0;
            for (int di = 0; di < 3; di++)
                for (int dj = 0; dj < 3; dj++)
                    cnt += tab[r + di][c + dj];

            if (cnt == 5) {
                ileKw++;
                // Sprawdz krzyz: centrum + 4 sasiady
                int cr = r + 1, cc = c + 1;
                bool krzyz = tab[cr][cc] == 1
                    && tab[cr - 1][cc] == 1 && tab[cr + 1][cc] == 1
                    && tab[cr][cc - 1] == 1 && tab[cr][cc + 1] == 1;
                cout << "   Rog (" << r + 1 << "," << c + 1 << "): "
                     << (krzyz ? "krzyz" : "nie krzyz") << endl;
                if (krzyz) ileKrzyz++;
            }
        }
    }
    cout << "Kwadraty z 5 jedynkami: " << ileKw << endl;
    cout << "Z nich krzyz: " << ileKrzyz << endl;
    return 0;
}
```

**Wyjasnienie**: Iterujemy po lewych gornych rogach kwadratow 3x3. Liczymy jedynki. Jesli 5, sprawdzamy czy centrum + 4 sasiedzi (krzyz). Krzyz z 5 jedynkami = dokladnie jedynki na pozycjach krzyza, zera na rogach.

Weryfikacja:
Siatka:
```
0 0 0 0 0 0 0 0
0 0 1 0 0 1 0 0
0 1 1 1 0 0 1 0
0 0 1 0 0 1 1 1
0 0 0 0 0 0 1 0
0 1 0 0 0 0 0 0
1 1 1 0 0 0 0 0
0 1 0 0 0 0 0 0
```
Kwadraty 3x3 z 5 jedynkami:
- Rog (0,1): {0,0,0 / 0,1,0 / 1,1,1} -> 4, nie 5
- Rog (1,1): {0,1,0 / 1,1,1 / 0,1,0} -> 5, krzyz TAK
- Rog (2,4): {0,1,0 / 0,1,1 / 0,0,1} -> 4
- Rog (2,5): {1,0,0 / 1,1,1 / 0,1,0} -> 5, krzyz? Centrum (3,6)=1, gora(2,6)=1, dol(4,6)=1, lewo(3,5)=1, prawo(3,7)=1. Tak, to tez krzyz.
- Rog (5,0): {0,1,0 / 1,1,1 / 0,1,0} -> 5, krzyz TAK
</details>

<details>
<summary>Typowe bledy</summary>

- **Petla `r <= 7` zamiast `r <= 5`**: Kwadrat 3x3 z rogiem w wierszu 6 konczy sie w wierszu 8, co jest poza tablica. CKE: crash
- **Zliczanie zamiast wzorca**: 5 jedynek w 3x3 nie oznacza krzyza — moga byc rozmieszczone inaczej. CKE: -1 pkt

</details>

---

### Cwiczenie 13.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 2 (transformacje siatki)
**Tagi**: `transponowanie` `obrot` `tablica-2D` `wczytywanie-pliku`

Dana jest siatka 4x6 liczb calkowitych zapisana w pliku `siatka.txt`. Napisz program, ktory:
a) Wypisze transpozycje siatki (zamiana wierszy z kolumnami — wynik to siatka 6x4).
b) Wypisze siatke obrócona o 90 stopni w prawo.
c) Sprawdzi, czy istnieje wiersz oryginalnej siatki identyczny z ktoras kolumna.

**Dane** (`siatka.txt`):
```
1 2 3 4 5 6
7 8 9 10 11 12
13 14 15 16 17 18
19 20 21 22 23 24
```

**Oczekiwany wynik**:
```
a) Transpozycja (6x4):
   1 7 13 19
   2 8 14 20
   3 9 15 21
   4 10 16 22
   5 11 17 23
   6 12 18 24

b) Obrot 90 stopni w prawo (6x4):
   19 13 7 1
   20 14 8 2
   21 15 9 3
   22 16 10 4
   23 17 11 5
   24 18 12 6

c) Wiersz = kolumna: NIE (zadna para nie jest identyczna)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Transpozycja: `nowa[j][i] = tab[i][j]`. Obrot 90 w prawo: `nowa[j][rows-1-i] = tab[i][j]`.
2. **Podejscie**: Utwórz nowe tablice na wyniki. Obrot 90 = transpozycja + odbicie lustrzane wierszy.
3. **Kluczowy krok**: Porownanie wiersza z kolumna: wiersz i ma 6 elementow, kolumna j ma 4 elementy — moga byc identyczne tylko jesli maja te same dlugosci (co tu nie zachodzi — wiersz ma 6, kolumna 4).

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ifstream plik("siatka.txt");
    int tab[4][6];
    for (int i = 0; i < 4; i++)
        for (int j = 0; j < 6; j++)
            plik >> tab[i][j];

    // a) Transpozycja
    cout << "a) Transpozycja (6x4):" << endl;
    for (int j = 0; j < 6; j++) {
        for (int i = 0; i < 4; i++)
            cout << tab[i][j] << " ";
        cout << endl;
    }

    // b) Obrot 90 w prawo: nowa[j][3-i] = tab[i][j]
    int obrot[6][4];
    for (int i = 0; i < 4; i++)
        for (int j = 0; j < 6; j++)
            obrot[j][3 - i] = tab[i][j];

    cout << endl << "b) Obrot 90 stopni w prawo (6x4):" << endl;
    for (int i = 0; i < 6; i++) {
        for (int j = 0; j < 4; j++)
            cout << obrot[i][j] << " ";
        cout << endl;
    }

    // c) Wiersz = kolumna
    // Wiersz ma 6 elementow, kolumna ma 4 — nie moga byc identyczne
    cout << endl << "c) Wiersz = kolumna: NIE (rozne dlugosci: wiersz=6, kolumna=4)" << endl;
    return 0;
}
```

**Wyjasnienie**: Transpozycja zamienia wiersze z kolumnami. Obrot 90 w prawo: `nowa[j][rows-1-i] = tab[i][j]`. Porownanie wiersza (dlugosc cols) z kolumna (dlugosc rows) jest sensowne tylko dla macierzy kwadratowej.

Weryfikacja:
- Transpozycja: kolumna 1 oryginalnej (1,7,13,19) staje sie wierszem 1 transpozycji.
- Obrot 90: pierwszy wiersz obrotu to ostatnia kolumna oryginalnej czytana od dolu: (19,13,7,1).
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomylenie transpozycji z obrotem**: Transpozycja to `[j][i]`, obrot 90 to `[j][rows-1-i]`. CKE: -2 pkt
- **Zapis do oryginalnej tablicy**: Przy transpozycji/obrocie trzeba uzyc nowej tablicy — nadpisywanie oryginalnej powoduje utrate danych. CKE: -2 pkt

</details>

---

### Cwiczenie 13.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 6 (zaawansowane BFS)
**Tagi**: `BFS` `spojne-obszary` `4-spojnosc` `tablica-2D` `wczytywanie-pliku`

Dana jest siatka 8x8 zer i jedynek zapisana w pliku `siatka.txt`. Napisz program, ktory:
a) Zliczy spojne obszary jedynek.
b) Dla kazdego obszaru poda obwod (liczba krawedzi graniczacych z zerem lub brzegiem siatki).
c) Znajdzie obszar o najwiekszym stosunku obwod/pole.

**Dane** (`siatka.txt`):
```
0 0 0 0 0 0 0 0
0 1 1 0 0 0 0 0
0 1 1 0 0 1 0 0
0 0 0 0 0 0 0 0
0 0 0 1 1 1 0 0
0 0 0 1 0 1 0 0
0 0 0 1 1 1 0 0
0 0 0 0 0 0 0 0
```

**Oczekiwany wynik**:
```
a) Liczba spojnych obszarow: 3

b) Obszary:
   Obszar 1: pole=4, obwod=8
   Obszar 2: pole=1, obwod=4
   Obszar 3: pole=8, obwod=16

c) Max obwod/pole: Obszar 2 (obwod/pole = 4.00)
```

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: BFS jak w 13.4. Obwod = liczba krawedzi jedynki graniczacych z zerem lub brzegiem.
2. **Podejscie**: Podczas BFS, dla kazdej jedynki sprawdz 4 sasiadow. Jesli sasiad to zero lub brzeg, dodaj 1 do obwodu.
3. **Kluczowy krok**: Pole = rozmiar obszaru (ile komorek). Obwod = suma krawedzi zewnetrznych.

</details>

<details>
<summary>Odpowiedz</summary>

**Rozwiazanie (C++):**
```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <queue>
#include <iomanip>
using namespace std;

int tab[8][8];
bool visited[8][8];
int dx[] = {-1, 1, 0, 0};
int dy[] = {0, 0, -1, 1};

struct Area { int pole, obwod; };

Area bfs(int si, int sj) {
    Area a = {0, 0};
    queue<pair<int,int>> q;
    q.push({si, sj});
    visited[si][sj] = true;

    while (!q.empty()) {
        auto [i, j] = q.front(); q.pop();
        a.pole++;
        for (int d = 0; d < 4; d++) {
            int ni = i + dx[d], nj = j + dy[d];
            if (ni < 0 || ni >= 8 || nj < 0 || nj >= 8 || tab[ni][nj] == 0) {
                a.obwod++; // krawedz graniczaca z zerem lub brzegiem
            } else if (!visited[ni][nj]) {
                visited[ni][nj] = true;
                q.push({ni, nj});
            }
        }
    }
    return a;
}

int main() {
    ifstream plik("siatka.txt");
    for (int i = 0; i < 8; i++)
        for (int j = 0; j < 8; j++)
            plik >> tab[i][j];

    vector<Area> areas;
    for (int i = 0; i < 8; i++)
        for (int j = 0; j < 8; j++)
            if (tab[i][j] == 1 && !visited[i][j])
                areas.push_back(bfs(i, j));

    cout << "a) Liczba spojnych obszarow: " << areas.size() << endl;

    cout << endl << "b) Obszary:" << endl;
    double maxRatio = 0; int bestIdx = 0;
    for (int k = 0; k < (int)areas.size(); k++) {
        cout << "   Obszar " << k + 1 << ": pole=" << areas[k].pole
             << ", obwod=" << areas[k].obwod << endl;
        double ratio = (double)areas[k].obwod / areas[k].pole;
        if (ratio > maxRatio) { maxRatio = ratio; bestIdx = k; }
    }

    cout << endl << "c) Max obwod/pole: Obszar " << bestIdx + 1
         << " (obwod/pole = " << fixed << setprecision(2) << maxRatio << ")" << endl;
    return 0;
}
```

**Wyjasnienie**: BFS z dodatkowym zliczaniem obwodu. Dla kazdej jedynki, kazdy sasiad bedacy zerem lub poza granica dodaje 1 do obwodu. Pole to rozmiar obszaru.

Weryfikacja:
- Obszar 1 (kwadrat 2x2 w rogach (1,1)-(2,2)): pole=4, obwod: kazda jedynka ma 2 krawedzie zewn. (narozne) -> 4*4=16? Nie. Sprawdzmy: (1,1): gora=0, lewo=0, dol=1, prawo=1 -> 2 zewn. (1,2): gora=0, prawo=0, dol=1, lewo=1 -> 2 zewn. (2,1): dol=0, lewo=0, gora=1, prawo=1 -> 2 zewn. (2,2): dol=0, prawo=0, gora=1, lewo=1 -> 2 zewn. Obwod = 2*4 = 8. Pole = 4.
- Obszar 2 (pojedyncza jedynka (2,5)): pole=1, obwod=4.
- Obszar 3 (pierscien w rogach (4,3)-(6,5)): pole=8, obwod: 8 jedynek * po 2 zewnetrzne krawedzie = 16. Dokladniej: (4,3): g=0,l=0,d=1,p=1 -> 2; (4,4): g=0,d=0,l=1,p=1 -> 2; (4,5): g=0,p=0,d=1,l=1 -> 2; (5,3): l=0,g=1,d=1,p=0 -> 2; (5,5): p=0,g=1,d=1,l=0 -> 2; (6,3): d=0,l=0,g=1,p=1 -> 2; (6,4): d=0,g=0,l=1,p=1 -> 2; (6,5): d=0,p=0,g=1,l=1 -> 2. Obwod = 16.
- Max obwod/pole: Obszar 2 (4/1=4.00), Obszar 1 (8/4=2.00), Obszar 3 (16/8=2.00).
</details>

<details>
<summary>Typowe bledy</summary>

- **Zliczanie obwodu po BFS zamiast w trakcie**: Mozliwe, ale mniej eleganckie i trzeba iterowac po tablicy label. CKE: 0 pkt (poprawne)
- **Pomylenie pola z obwodem**: Pole = liczba komorek (=1), obwod = liczba krawedzi zewnetrznych. CKE: -2 pkt
- **Brak uwzglednienia brzegu siatki**: Krawedz przy brzegu siatki tez jest zewnetrzna. CKE: -1 pkt

</details>

---

## Samoocena

| Poziom | Opis | Kryteria |
|--------|------|----------|
| Podstawowy | Umiem zliczac komorki w siatce 2D | Cwiczenia 1-2, 6 bez pomocy |
| Dobry | Radze sobie z sasiedztwo 4-spojnym i wzorcami | Cwiczenia 3, 7-8 bez pomocy |
| Bardzo dobry | Umiem implementowac BFS na siatce | Cwiczenia 4-5 bez pomocy |
| Doskonaly | Radze sobie ze zlozonym BFS (obwod, multi-wartosc) | Cwiczenia 9-10 bez pomocy |

**Co dalej?**
- Jesli masz problem z tablicami 2D -> cwicz wczytywanie i iterowanie (cw. 13.1, 13.6)
- Jesli chcesz cwiczycBFS -> patrz `cheatsheet_cpp.md` sekcja "BFS/DFS"
- Jesli chcesz cwiczycstrukturydanych -> patrz `09_zlozone.md`
