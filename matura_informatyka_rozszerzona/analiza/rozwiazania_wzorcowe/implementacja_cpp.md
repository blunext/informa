# Rozwiazania wzorcowe: Implementacja (C++)

Trzy pelne rozwiazania prawdziwych zadan maturalnych — z procesem myslowym, kodem i weryfikacja.

---

## [2024] Zadanie 3: Nieparzysty skrot (10 pkt)

**Typ**: cyfry_liczby + projektowanie_algorytmu | **Czas**: ~35 min | **Trudnosc**: trudne

### Tresc (skrot)

Nieparzysty skrot liczby `n` to liczba powstala przez **usuniecie cyfr parzystych** z zapisu dziesietnego `n`.
Np. skrot(24680) nie istnieje (same parzyste), skrot(13579) = 13579, skrot(123456) = 135.

- **3.1** (3 pkt): Pseudokod funkcji `skrot(n)` — tylko operacje calkowitoliczbowe (mod/div), **bez stringow**.
- **3.2** (3 pkt): Plik `skrot.txt` (200 liczb < 30000). Ile liczb nie ma skrotu? Podaj najwieksza taka liczbe.
- **3.3** (4 pkt): Plik `skrot2.txt` (200 liczb). Wypisz liczby, dla ktorych `NWD(liczba, skrot(liczba)) = 7`.

### Podejscie — jak myslec

1. **Rozpoznanie typu**: przetwarzanie cyfr (mod 10 / div 10) + NWD Euklidesa — dwa klasyczne wzorce.
2. **3.1 — pseudokod**: Iterujemy po cyfrach od prawej (mod 10), budujemy wynik od konca. Kluczowa zmienna `p` (potega 10) sluzy do "przyklejania" cyfr nieparzystych na odpowiednia pozycje.
3. **3.2**: Liczba nie ma skrotu, gdy `skrot(n) == 0` (wszystkie cyfry parzyste). Zliczamy takie + szukamy max.
4. **3.3**: Dla kazdej liczby: oblicz skrot, oblicz NWD, sprawdz == 7.

### Rozwiazanie

#### 3.1 — Pseudokod (3 pkt)

```
funkcja skrot(n):
    m <- 0          // wynik (skrot)
    p <- 1          // pozycja (1, 10, 100, ...)
    dopoki n > 0:
        cyfra <- n mod 10
        jesli cyfra mod 2 <> 0:   // cyfra nieparzysta
            m <- m + cyfra * p
            p <- p * 10
        n <- n div 10
    zwroc m
```

**Uwaga CKE**: Za uzycie stringow (str(), substr) — 0 pkt. Tylko arytmetyka calkowitoliczbowa.

#### 3.2 + 3.3 — Program C++ (7 pkt)

```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

// Nieparzysty skrot (arytmetycznie)
int skrot(int n) {
    int m = 0, p = 1;
    while (n > 0) {
        int cyfra = n % 10;
        if (cyfra % 2 != 0) {   // nieparzysta
            m += cyfra * p;
            p *= 10;
        }
        n /= 10;
    }
    return m;
}

// NWD — algorytm Euklidesa
int nwd(int a, int b) {
    while (b) { int t = b; b = a % b; a = t; }
    return a;
}

int main() {
    // --- Zadanie 3.2 ---
    ifstream f1("skrot.txt");
    int x, ile_bez = 0, max_bez = 0;
    while (f1 >> x) {
        if (skrot(x) == 0) {        // brak skrotu = same parzyste cyfry
            ile_bez++;
            if (x > max_bez) max_bez = x;
        }
    }
    cout << "3.2: " << ile_bez << " liczb, max = " << max_bez << endl;
    // Odp: 18 liczb, max = 28422

    // --- Zadanie 3.3 ---
    ifstream f2("skrot2.txt");
    cout << "3.3: ";
    while (f2 >> x) {
        int s = skrot(x);
        if (s > 0 && nwd(x, s) == 7) {
            cout << x << " ";
        }
    }
    cout << endl;
    // Odp: 784, 14196, 2247, 24087, 3871, 10192

    return 0;
}
```

### Weryfikacja

| Podzadanie | Oficjalna odpowiedz CKE |
|---|---|
| 3.2 | 18 liczb, najwieksza: **28422** |
| 3.3 | **784, 14196, 2247, 24087, 3871, 10192** |

### Pulapki

- **3.1**: Uzycie stringow = 0 pkt. Trzeba budowac wynik arytmetycznie.
- **3.1**: Zmienna `p` musi rosnac TYLKO gdy dodajemy cyfre (nie w kazdej iteracji!).
- **3.2**: Pomylenie max z min (czesty blad za 1 pkt zamiast 2).
- **3.3**: NWD musi byc **dokladnie** 7, nie wielokrotnosc 7. Blad: `nwd(x,s) % 7 == 0`.
- **3.3**: Nie zapomnij sprawdzic `s > 0` (skrot istnieje) — plik gwarantuje to, ale bezpiecznie.

---

## [2024] Zadanie 4: Liczby (10 pkt)

**Typ**: zliczanie + minmax + cyfry_liczby + sekwencje | **Czas**: ~45 min | **Trudnosc**: trudne

### Tresc (skrot)

Plik `liczby.txt`:
- Wiersz 1: **3000 liczb pierwszych** z zakresu [2, 2000].
- Wiersz 2: **20 liczb calkowitych** z zakresu [2, 1 000 000 000].

Podzadania:
- **4.1** (2 pkt): Ile liczb z wiersza 1 jest dzielnikiem jakiejkolwiek liczby z wiersza 2?
- **4.2** (2 pkt): 101-sza liczba z wiersza 1 w kolejnosci **od najwiekszej**.
- **4.3** (3 pkt): Ktore z 20 liczb da sie przedstawic jako iloczyn wylacznie liczb z wiersza 1 (z ograniczeniem krotnosci)?
- **4.4** (3 pkt): Spojny fragment wiersza 1 o **co najmniej 50 elementach** z najwieksza srednia.

### Podejscie — jak myslec

1. **4.1**: Dla kazdej liczby pierwszej p sprawdzamy, czy ktoras z 20 liczb jest podzielna przez p. Uzywamy zbioru (set) pierwszych do szybkiego lookup.
2. **4.2**: Sortowanie wiersza 1 malejaco, wybranie elementu [100] (indeks od 0).
3. **4.3**: Rozklad na czynniki pierwsze — dzielimy kazda z 20 liczb kolejno przez liczby z wiersza 1 (zliczajac krotnosci). Jesli po wyczerpaniu sie zostaje 1, to sie da.
4. **4.4**: **Sliding window / sumy prefiksowe** — najtrudniejsze podzadanie. Szukamy spojnego fragmentu >= 50 elem. z max srednia.

### Rozwiazanie

```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <set>
#include <map>
#include <sstream>
using namespace std;

int main() {
    ifstream plik("liczby.txt");
    string linia1, linia2;
    getline(plik, linia1);
    getline(plik, linia2);

    // Parsowanie wiersza 1 (3000 liczb pierwszych)
    vector<int> pierwsze;
    istringstream ss1(linia1);
    int x;
    while (ss1 >> x) pierwsze.push_back(x);

    // Parsowanie wiersza 2 (20 duzych liczb)
    vector<long long> duze;
    istringstream ss2(linia2);
    long long y;
    while (ss2 >> y) duze.push_back(y);

    // --- 4.1: Ile liczb pierwszych jest dzielnikiem jakiejkolwiek duzej ---
    int ile_dzielnikow = 0;
    for (int p : pierwsze) {
        bool jest = false;
        for (long long d : duze) {
            if (d % p == 0) { jest = true; break; }
        }
        if (jest) ile_dzielnikow++;
    }
    cout << "4.1: " << ile_dzielnikow << endl;  // 212

    // --- 4.2: 101-sza od najwiekszej ---
    vector<int> posort = pierwsze;
    sort(posort.begin(), posort.end(), greater<int>());
    cout << "4.2: " << posort[100] << endl;  // 1933

    // --- 4.3: Rozklad na czynniki z listy ---
    // Zlicz krotnosci kazdej liczby pierwszej w wierszu 1
    map<int, int> dostepne;
    for (int p : pierwsze) dostepne[p]++;

    cout << "4.3: ";
    for (long long d : duze) {
        long long kopia = d;
        map<int, int> uzyto;
        bool ok = true;
        // Dzielimy przez kazda unikalna pierwsza z listy
        for (auto& [p, maks] : dostepne) {
            while (kopia % p == 0) {
                uzyto[p]++;
                if (uzyto[p] > maks) { ok = false; break; }
                kopia /= p;
            }
            if (!ok) break;
        }
        if (ok && kopia == 1) {
            cout << d << " ";
        }
    }
    cout << endl;
    // 547839600, 2954285, 573219169, 573549984, 212444924

    // --- 4.4: Spojny fragment >= 50 elem. z max srednia ---
    // Sumy prefiksowe
    int n = pierwsze.size();
    vector<long long> pref(n + 1, 0);
    for (int i = 0; i < n; i++) pref[i + 1] = pref[i] + pierwsze[i];

    double best_avg = 0;
    int best_start = 0, best_len = 0;

    // Sprawdzamy wszystkie fragmenty o dlugosci >= 50
    // Optymalizacja: dla kazdego konca j, szukamy min pref[i] dla i <= j-50
    // Prostsze: iterujemy po dlugosciach od 50 do n
    // Ale O(n^2) = 9*10^6 — akceptowalne
    for (int i = 0; i <= n - 50; i++) {
        for (int j = i + 50; j <= n; j++) {
            double avg = (double)(pref[j] - pref[i]) / (j - i);
            if (avg > best_avg || (avg == best_avg && i < best_start)) {
                best_avg = avg;
                best_start = i;
                best_len = j - i;
            }
        }
    }
    cout << "4.4: srednia=" << best_avg
         << " elementow=" << best_len
         << " poczatek=" << pierwsze[best_start]
         << " (pozycja " << best_start + 1 << ")" << endl;
    // srednia ~1200.70, 61 elementow, poczatek 1847 (pozycja 2797)

    return 0;
}
```

**Uwaga o zlozonosci 4.4**: Powyzszy O(n^2) dziala dla n=3000 (~4.5 mln operacji). Mozna tez uzyc sprytniejszej metody: dla kazdego konca `j`, minimalny `pref[i]` wsrod `i <= j-50` rosnie monotonicznie — wystarczy sledzic jedno minimum.

### Weryfikacja

| Podzadanie | Oficjalna odpowiedz CKE |
|---|---|
| 4.1 | **212** (pulapka: 2788 = liczby NIE bedace dzielnikami) |
| 4.2 | **1933** (pulapka: 31 = 101-sza od najmniejszej) |
| 4.3 | **547839600, 2954285, 573219169, 573549984, 212444924** |
| 4.4 | Srednia: **1200.704918...**, 61 elementow, poczatek: **1847** (poz. 2797) |

### Pulapki

- **4.1**: Odwrotna odpowiedz (2788) za 1 pkt — pytaja o dzielniki, nie "nie-dzielniki".
- **4.2**: Sortowanie od najmniejszej daje 31 (za 1 pkt). Pytaja od **najwiekszej**.
- **4.3**: Ograniczenie krotnosci — jesli liczba pierwsza `p` wystepuje w wierszu 1 trzy razy, mozna uzyc `p` co najwyzej 3 razy w rozkladzie.
- **4.3**: Duze liczby (do 10^9) — uzyj `long long`.
- **4.4**: Warunek **co najmniej 50** elementow — fragmenty 49-elementowe sa niepoprawne.
- **4.4**: Przy remisie — pierwsza srednia (najwczesniej wystepujaca).

---

## [2023] Zadanie 3: Liczba Pi (10 pkt)

**Typ**: zliczanie + minmax + sekwencje | **Czas**: ~45 min | **Trudnosc**: srednie

### Tresc (skrot)

Plik `pi.txt`: 10000 cyfr rozwieniecia dziesietnego liczby pi (kazda cyfra w osobnym wierszu).
Fragment 2-cyfrowy = para kolejnych cyfr (np. cyfry na pozycjach 1-2, 2-3, ..., 9999-10000 = 9999 fragmentow).

- **3.1** (2 pkt): Ile fragmentow 2-cyfrowych ma wartosc **wieksza od 90** (>90, nie >=90)?
- **3.2** (3 pkt): Ktory fragment 2-cyfrowy ma **najmniej** i **najwiecej** wystapien? Przy remisie — mniejsza wartosc.
- **3.3** (3 pkt): Ile ciagow **rosnaco-malejacych** z dokladnie 6 kolejnych cyfr?
- **3.4** (2 pkt): Najdluzszy ciag rosnaco-malejacy — podaj pozycje i ciag.

Ciag rosnaco-malejacy (dlugosc >= 4): istnieje `k` (2 <= k <= n-2) takie ze:
`a[1] < a[2] < ... < a[k]` oraz `a[k] > a[k+1] > ... > a[n]`.
Rownosc na granicy jest dopuszczalna (np. 5,9,9,4,1 — czesc rosnaca 5,9; malejaca 9,4,1).

### Podejscie — jak myslec

1. **3.1**: Proste — tworzymy pary, liczymy ile > 90. Uwaga: `> 90`, nie `>= 90`.
2. **3.2**: Histogram 100 komorek (00-99). Szukamy min i max z warunkiem remisu.
3. **3.3/3.4**: To najtrudniejsze. Trzeba rozpoznac wzorzec "rośnie-maleje":
   - Czesc rosnaca: **ostro rosnaca** (a[i] < a[i+1])
   - Punkt "szczytu": a[k-1] < a[k] >= a[k+1] (rownosc na granicy!)
   - Czesc malejaca: **ostro malejaca** (a[i] > a[i+1])
   - Minimum 2 elementy w czesci rosnacej i 2 w malejacej.

### Rozwiazanie

```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int main() {
    // Wczytaj cyfry
    ifstream plik("pi.txt");
    vector<int> c;
    int x;
    while (plik >> x) c.push_back(x);
    int n = c.size();  // 10000

    // --- 3.1: Fragmenty 2-cyfrowe > 90 ---
    int cnt_gt90 = 0;
    for (int i = 0; i < n - 1; i++) {
        int frag = c[i] * 10 + c[i + 1];
        if (frag > 90) cnt_gt90++;       // > 90, NIE >= 90!
    }
    cout << "3.1: " << cnt_gt90 << endl;  // 902

    // --- 3.2: Histogram fragmentow 2-cyfrowych ---
    int hist[100] = {};
    for (int i = 0; i < n - 1; i++) {
        int frag = c[i] * 10 + c[i + 1];
        hist[frag]++;
    }
    int min_frag = 0, max_frag = 0;
    for (int f = 1; f < 100; f++) {       // od 1, bo fragment 00 tez mozliwy
        if (hist[f] < hist[min_frag]) min_frag = f;
        if (hist[f] > hist[max_frag]) max_frag = f;
    }
    // Przy remisie — mniejsza wartosc (iterujemy od 0, wiec pierwsza znaleziona)
    cout << "3.2: min=" << min_frag << " (" << hist[min_frag] << " wyst.)"
         << "  max=" << max_frag << " (" << hist[max_frag] << " wyst.)" << endl;
    // min=88 (80 wyst.), max=65 (124 wyst.)

    // --- 3.3: Ciagi rosnaco-malejace z dokladnie 6 cyfr ---
    int cnt_rm = 0;
    for (int i = 0; i <= n - 6; i++) {
        // Sprawdzamy 6-elementowy podciag c[i..i+5]
        // Szukamy punktu szczytu k (indeks 1..4, bo min 2 rosnace i 2 malejace)
        // k = pozycja ostatniego elementu czesci rosnacej (wzgl. poczatku okna)
        for (int k = 1; k <= 4; k++) {
            // Czesc rosnaca: c[i]..c[i+k] ostro rosnaca
            bool rosnaca = true;
            for (int j = 0; j < k; j++) {
                if (c[i + j] >= c[i + j + 1]) { rosnaca = false; break; }
            }
            if (!rosnaca) continue;

            // Czesc malejaca: c[i+k]..c[i+5] ostro malejaca
            bool malejaca = true;
            for (int j = k; j < 5; j++) {
                if (c[i + j] <= c[i + j + 1]) { malejaca = false; break; }
            }
            if (!malejaca) continue;

            // Znaleziono rosnaco-malejacy z k elementow rosnacej, 6-k malejacej
            // Warunek: k >= 1 (min 2 w rosnacej: elem 0..k) i k <= 4 (min 2 w malejacej: elem k..5)
            cnt_rm++;
            break;  // Nie liczymy podwojnie tego samego okna
        }
    }
    cout << "3.3: " << cnt_rm << endl;  // 214

    // --- 3.4: Najdluzszy ciag rosnaco-malejacy ---
    int best_pos = -1, best_len = 0;

    for (int i = 0; i < n; i++) {
        // Znajdz dlugosc czesci rosnacej od i
        int j = i;
        while (j + 1 < n && c[j] < c[j + 1]) j++;
        int rosnaca_do = j;  // szczyt na pozycji j

        // Czesc rosnaca musi miec >= 2 elementy (pozycje i..j, j > i)
        if (rosnaca_do == i) continue;

        // Znajdz dlugosc czesci malejacej od j
        while (j + 1 < n && c[j] > c[j + 1]) j++;
        int malejaca_do = j;

        // Czesc malejaca musi miec >= 2 elementy (pozycje rosnaca_do..j, j > rosnaca_do)
        if (malejaca_do == rosnaca_do) continue;

        int len = malejaca_do - i + 1;  // dlugosc calego ciagu
        if (len >= 4 && len > best_len) {
            best_len = len;
            best_pos = i;
        }
    }

    cout << "3.4: pozycja=" << best_pos + 1 << " dlugosc=" << best_len << " ciag=";
    for (int j = best_pos; j < best_pos + best_len; j++) cout << c[j];
    cout << endl;
    // pozycja=2781, ciag=014576540, dlugosc 9

    return 0;
}
```

### Weryfikacja

| Podzadanie | Oficjalna odpowiedz CKE |
|---|---|
| 3.1 | **902** (pulapka: 1014 = fragmenty >= 90) |
| 3.2 | min: **88** (80 wyst.), max: **65** (124 wyst.) |
| 3.3 | **214** |
| 3.4 | Pozycja: **2781**, ciag: **014576540** (dlugosc 9) |

### Pulapki

- **3.1**: `> 90` vs `>= 90` — roznica 902 vs 1014. CKE daje 1 pkt za 1014.
- **3.2**: Przy remisach trzeba brac fragment o **mniejszej wartosci** liczbowej.
- **3.3**: Definicja rosnaco-malejacego jest nietypowa — czesc rosnaca jest **ostro** rosnaca, czesc malejaca **ostro** malejaca, ale na granicy (szczycie) dozwolona jest rownosc. Np. (5,9,9,4,1) jest poprawny.
- **3.3**: Warunek `k >= 2` i `n-k >= 2` — minimum 2 elementy w kazdej czesci.
- **3.4**: Ciag musi byc "dokladnie" rosnaco-malejacy — nie mozna go rozszerzyc w zadna strone (jest maksymalny).
