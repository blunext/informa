# Szablony C++ — Sciagawka na Mature

Samowystarczalny dokument. Kazdy szablon kompiluje sie samodzielnie.

---

## 1. Wczytywanie plikow CKE

### A. Liczby oddzielone spacjami (jedna lub wiele na linie)

Matura 2024 `liczby.txt`: `2 3 5 7 3 5 7 ...`

```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int main() {
    ifstream plik("liczby.txt");
    vector<int> tab;
    int x;
    while (plik >> x) {
        tab.push_back(x);
    }
    cout << "Wczytano: " << tab.size() << " liczb" << endl;
    return 0;
}
```

### B. Jedna wartosc na linie

Matura 2023 `pi.txt`: kolejne cyfry, kazda w osobnym wierszu.

```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <string>
using namespace std;

int main() {
    ifstream plik("pi.txt");
    string linia;
    vector<int> cyfry;
    while (getline(plik, linia)) {
        if (!linia.empty())
            cyfry.push_back(stoi(linia));
    }
    cout << "Wczytano: " << cyfry.size() << " cyfr" << endl;
    return 0;
}
```

### C. CSV ze srednikiem + naglowek

Matura 2024 `kierowcy.txt`: `IdOsoby;Imie;Nazwisko;NrRejestracyjny`

```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <vector>
using namespace std;

int main() {
    ifstream plik("kierowcy.txt");
    string linia;
    getline(plik, linia); // pomin naglowek

    vector<string> imiona;
    vector<string> nazwiska;
    vector<int> id;

    while (getline(plik, linia)) {
        stringstream ss(linia);
        string token;
        vector<string> kol;
        while (getline(ss, token, ';')) {
            kol.push_back(token);
        }
        // kol[0]=Id, kol[1]=Imie, kol[2]=Nazwisko, kol[3]=NrRej
        id.push_back(stoi(kol[0]));
        imiona.push_back(kol[1]);
        nazwiska.push_back(kol[2]);
    }
    cout << "Wczytano: " << id.size() << " kierowcow" << endl;
    return 0;
}
```

### D. TSV (tab) + naglowek

Matura 2023 `gracze.txt`: `id_gracza\timie\tnazwisko\twiek`

```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <vector>
using namespace std;

int main() {
    ifstream plik("gracze.txt");
    string linia;
    getline(plik, linia); // pomin naglowek

    vector<string> imiona;
    vector<int> wiek;

    while (getline(plik, linia)) {
        stringstream ss(linia);
        string sid, imie, nazwisko;
        int w;
        ss >> sid >> imie >> nazwisko >> w;
        imiona.push_back(imie);
        wiek.push_back(w);
    }
    cout << "Wczytano: " << imiona.size() << " graczy" << endl;
    return 0;
}
```

**Alternatywa z getline + tab:**

```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <vector>
using namespace std;

int main() {
    ifstream plik("gracze.txt");
    string linia;
    getline(plik, linia); // pomin naglowek

    while (getline(plik, linia)) {
        stringstream ss(linia);
        string token;
        vector<string> kol;
        while (getline(ss, token, '\t')) {
            kol.push_back(token);
        }
        // kol[0]=id, kol[1]=imie, kol[2]=nazwisko, kol[3]=wiek
        cout << kol[1] << " " << kol[3] << endl;
    }
    return 0;
}
```

### E. Pary/trojki na linie

Matura 2025 `dron.txt`: `217 98` (x y na linie)

```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int main() {
    ifstream plik("dron.txt");
    int x, y;
    vector<pair<int,int>> punkty;
    while (plik >> x >> y) {
        punkty.push_back({x, y});
    }
    cout << "Wczytano: " << punkty.size() << " punktow" << endl;
    return 0;
}
```

**Wersja z trojkami:**

```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

struct Rekord { int a, b, c; };

int main() {
    ifstream plik("dane.txt");
    vector<Rekord> tab;
    int a, b, c;
    while (plik >> a >> b >> c) {
        tab.push_back({a, b, c});
    }
    cout << "Wczytano: " << tab.size() << " rekordow" << endl;
    return 0;
}
```

### F. Wiele plikow naraz (JOIN)

Matura 2024, 2025: czytanie 2-3 plikow i laczenie po kluczu.

```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <map>
#include <vector>
using namespace std;

int main() {
    // Plik 1: id -> nazwa
    ifstream f1("kierowcy.txt");
    string linia;
    getline(f1, linia); // pomin naglowek
    map<int, string> kierowcy;
    while (getline(f1, linia)) {
        stringstream ss(linia);
        string token;
        vector<string> kol;
        while (getline(ss, token, ';')) kol.push_back(token);
        kierowcy[stoi(kol[0])] = kol[1] + " " + kol[2];
    }

    // Plik 2: id, dane
    ifstream f2("rejestr.txt");
    getline(f2, linia); // pomin naglowek
    map<int, vector<int>> dane;
    while (getline(f2, linia)) {
        stringstream ss(linia);
        string token;
        vector<string> kol;
        while (getline(ss, token, ';')) kol.push_back(token);
        int id = stoi(kol[0]);
        int wartosc = stoi(kol[1]);
        dane[id].push_back(wartosc);
    }

    // JOIN: wypisz dane z nazwami
    for (auto &p : dane) {
        if (kierowcy.count(p.first)) {
            cout << kierowcy[p.first] << ": ";
            for (int v : p.second) cout << v << " ";
            cout << endl;
        }
    }
    return 0;
}
```

### G. Liczby dziesietne z europejskim przecinkiem

Matura 2024 `cennik.txt`: `Alwa\t2,9` (przecinek zamiast kropki)

```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <sstream>
#include <vector>
using namespace std;

int main() {
    ifstream plik("cennik.txt");
    string linia;
    while (getline(plik, linia)) {
        stringstream ss(linia);
        string nazwa, cena_str;
        ss >> nazwa >> cena_str;
        // zamien przecinek na kropke
        for (char &c : cena_str)
            if (c == ',') c = '.';
        double cena = stod(cena_str);
        cout << nazwa << ": " << cena << endl;
    }
    return 0;
}
```

### H. Siatka znakow (obraz 2D)

Matura 2025 `symbole.txt`: siatka `+`, `*`, `o` (znaki oddzielone spacjami)

```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <string>
using namespace std;

int main() {
    ifstream plik("symbole.txt");
    vector<vector<char>> siatka;
    string linia;
    while (getline(plik, linia)) {
        vector<char> wiersz;
        for (char c : linia) {
            if (c != ' ' && c != '\t') wiersz.push_back(c);
        }
        if (!wiersz.empty()) siatka.push_back(wiersz);
    }
    int wiersze = siatka.size();
    int kolumny = siatka[0].size();
    cout << "Siatka: " << wiersze << "x" << kolumny << endl;
    return 0;
}
```

---

## 2. Przetwarzanie cyfr i liczb

### Suma/iloczyn/liczba cyfr

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

int liczbaCyfr(int n) {
    int cnt = 0;
    while (n > 0) { cnt++; n /= 10; }
    return cnt;
}

int main() {
    ifstream plik("dane.txt");
    int n;
    while (plik >> n) {
        cout << n << " -> suma cyfr: " << sumaCyfr(n)
             << ", liczba cyfr: " << liczbaCyfr(n) << endl;
    }
    return 0;
}
```

### NWD (Euklides)

```cpp
#include <iostream>
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
    cout << "NWD(48, 18) = " << nwd(48, 18) << endl; // 6
    return 0;
}
```

### Czy pierwsza

```cpp
#include <iostream>
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

int main() {
    for (int i = 2; i <= 30; i++)
        if (czyPierwsza(i)) cout << i << " ";
    cout << endl;
    return 0;
}
```

### Rozklad na czynniki pierwsze

```cpp
#include <iostream>
#include <vector>
using namespace std;

int main() {
    int n = 60;
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
    for (int i = 0; i < (int)czynniki.size(); i++) {
        if (i > 0) cout << " * ";
        cout << czynniki[i];
    }
    cout << endl; // 60 = 2 * 2 * 3 * 5
    return 0;
}
```

### Sito Eratostenesa

```cpp
#include <iostream>
using namespace std;

const int MAX = 100001;
bool pierwsza[MAX];

void sito(int n) {
    fill(pierwsza, pierwsza + n + 1, true);
    pierwsza[0] = pierwsza[1] = false;
    for (int i = 2; i * i <= n; i++)
        if (pierwsza[i])
            for (int j = i * i; j <= n; j += i)
                pierwsza[j] = false;
}

int main() {
    sito(100);
    for (int i = 2; i <= 100; i++)
        if (pierwsza[i]) cout << i << " ";
    cout << endl;
    return 0;
}
```

### Dzielniki liczby

```cpp
#include <iostream>
#include <vector>
using namespace std;

vector<int> dzielniki(int n) {
    vector<int> div;
    for (int i = 1; i * i <= n; i++) {
        if (n % i == 0) {
            div.push_back(i);
            if (i != n / i) div.push_back(n / i);
        }
    }
    return div;
}

int main() {
    for (int d : dzielniki(36)) cout << d << " ";
    cout << endl;
    return 0;
}
```

### Pulapki

- `int` overflow przy duzych liczbach → uzyj `long long`
- Dzielenie calkowite: `7 / 2 == 3`, nie 3.5 → rzutuj: `(double)7 / 2`
- `n == 0` → `sumaCyfr(0)` zwraca 0, `liczbaCyfr(0)` zwraca 0 (obsluz osobno jesli trzeba)

---

## 3. Przetwarzanie napisow

### Palindrom

```cpp
#include <iostream>
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
    string s = "kajak";
    cout << s << (czyPalindrom(s) ? " jest" : " nie jest")
         << " palindromem" << endl;
    return 0;
}
```

### Szyfr Cezara

```cpp
#include <iostream>
#include <string>
using namespace std;

int main() {
    string napis = "KHOOR";
    int k = 3; // przesuniecie
    string wynik = "";
    for (char c : napis) {
        // deszyfrowanie: przesuniecie w lewo
        wynik += (char)((c - 'A' - k + 26) % 26 + 'A');
    }
    cout << wynik << endl; // HELLO
    return 0;
}
```

### Zliczanie znakow

```cpp
#include <iostream>
#include <string>
#include <map>
using namespace std;

int main() {
    string s = "abrakadabra";

    // Sposob 1: tablica (szybki, dla ASCII)
    int freq[256] = {0};
    for (char c : s) freq[(int)c]++;
    cout << "a: " << freq['a'] << endl;

    // Sposob 2: mapa (uniwersalny)
    map<char, int> m;
    for (char c : s) m[c]++;
    for (auto &p : m)
        cout << p.first << ": " << p.second << endl;
    return 0;
}
```

### Konwersja na male/duze litery

```cpp
#include <iostream>
#include <string>
using namespace std;

string toLower(string s) {
    for (char &c : s)
        if (c >= 'A' && c <= 'Z') c = c - 'A' + 'a';
    return s;
}

int main() {
    cout << toLower("KaJaK") << endl; // kajak
    return 0;
}
```

### Operacje na stringach — skrocona sciagawka

```cpp
#include <iostream>
#include <string>
#include <algorithm>
using namespace std;

int main() {
    string s = "Hello";
    s.length();              // 5
    s[0];                    // 'H'
    s.substr(1, 3);          // "ell"
    s + " World";            // "Hello World"
    s.find("ll");            // 2 (pozycja) lub string::npos
    s.erase(1, 2);           // "Hlo" (usun 2 znaki od poz 1)

    // kody ASCII
    int kod = (int)'A';      // 65
    char znak = (char)97;    // 'a'

    // to_string / stoi
    string num = to_string(42); // "42"
    int val = stoi("42");       // 42
    return 0;
}
```

### Pulapki

- `s.length()` zwraca `size_t` (unsigned) — porownanie z -1 nie dziala jak oczekiwano
- Indeksowanie od 0: `s[0]` to pierwszy znak
- `s.find()` zwraca `string::npos` (nie -1) gdy nie znaleziono

---

## 4. Wzorce zliczania i wyszukiwania

### Zliczanie z warunkiem

```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    int n;
    int ile = 0;
    while (plik >> n) {
        if (n % 2 == 0) ile++; // zmien warunek wg zadania
    }
    cout << "Ilosc: " << ile << endl;
    return 0;
}
```

### Min/max z pozycja

```cpp
#include <iostream>
#include <fstream>
#include <vector>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    vector<int> tab;
    int x;
    while (plik >> x) tab.push_back(x);

    int minVal = tab[0], maxVal = tab[0];
    int minPos = 0, maxPos = 0;
    for (int i = 1; i < (int)tab.size(); i++) {
        if (tab[i] < minVal) { minVal = tab[i]; minPos = i; }
        if (tab[i] > maxVal) { maxVal = tab[i]; maxPos = i; }
    }
    cout << "Min: " << minVal << " (poz " << minPos + 1 << ")" << endl;
    cout << "Max: " << maxVal << " (poz " << maxPos + 1 << ")" << endl;
    return 0;
}
```

### Najdluzszy podciag spelniajacy warunek

Wzorzec current/max — najczesciej uzywany na maturze.

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
        if (T[i] > T[i - 1]) { // warunek: scisle rosnacy
            curDl++;
        } else {
            if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }
            curDl = 1;
            curStart = i;
        }
    }
    if (curDl > maxDl) { maxDl = curDl; maxStart = curStart; }

    cout << "Najdluzszy ciag: dlugosc " << maxDl
         << ", start poz " << maxStart + 1 << endl;
    return 0;
}
```

Warianty warunku w linii `if`:
- `T[i] == T[i-1]` — ciag rownych
- `T[i] > T[i-1]` — scisle rosnacy
- `T[i] >= T[i-1]` — niemalejacy
- `T[i] < T[i-1]` — scisle malejacy

### Zliczanie z mapa czestotliwosci

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

    cout << "Roznych wartosci: " << freq.size() << endl;

    int maxF = 0;
    for (auto &p : freq)
        if (p.second > maxF) maxF = p.second;

    cout << "Najczestsze (mode): ";
    for (auto &p : freq)
        if (p.second == maxF) cout << p.first << " ";
    cout << "(" << maxF << " razy)" << endl;
    return 0;
}
```

### Pulapki

- Inicjalizacja min: uzyj `tab[0]`, nie `0` (bo min moze byc ujemne)
- Inicjalizacja max: uzyj `tab[0]`, nie `0` (bo max moze byc 0)
- Po petli: sprawdz ostatni segment (`if (curDl > maxDl)`)

---

## 5. Sortowanie z kluczem

### sort z lambda

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    vector<int> tab = {5, 2, 8, 1, 9};

    // Rosnaco (domyslnie)
    sort(tab.begin(), tab.end());

    // Malejaco
    sort(tab.begin(), tab.end(), [](int a, int b) {
        return a > b;
    });

    // Wg reszty z dzielenia przez 3
    sort(tab.begin(), tab.end(), [](int a, int b) {
        return (a % 3) < (b % 3);
    });

    for (int x : tab) cout << x << " ";
    cout << endl;
    return 0;
}
```

### sort ze struct

```cpp
#include <iostream>
#include <fstream>
#include <vector>
#include <algorithm>
#include <string>
using namespace std;

struct Osoba {
    string imie;
    int wiek;
};

int main() {
    ifstream plik("dane.txt");
    vector<Osoba> osoby;
    string imie; int wiek;
    while (plik >> imie >> wiek) {
        osoby.push_back({imie, wiek});
    }

    // Sortuj po wieku malejaco, przy rownym po imieniu rosnaco
    sort(osoby.begin(), osoby.end(), [](Osoba &a, Osoba &b) {
        if (a.wiek != b.wiek) return a.wiek > b.wiek;
        return a.imie < b.imie;
    });

    for (auto &o : osoby)
        cout << o.imie << " " << o.wiek << endl;
    return 0;
}
```

### Pulapka: stabilnosc sortowania

`sort` nie gwarantuje stabilnosci (kolejnosc rownych elementow moze sie zmienic).
Jesli stabilnosc jest potrzebna, uzyj `stable_sort`.

---

## 6. Formatowanie wyjscia

### Zaokraglanie

```cpp
#include <iostream>
#include <iomanip>
using namespace std;

int main() {
    double x = 3.14159;
    cout << fixed << setprecision(2) << x << endl; // 3.14
    cout << fixed << setprecision(4) << x << endl; // 3.1416
    return 0;
}
```

### Kolumny

```cpp
#include <iostream>
#include <iomanip>
using namespace std;

int main() {
    cout << setw(10) << "Imie" << setw(10) << "Wiek" << endl;
    cout << setw(10) << "Anna" << setw(10) << 25 << endl;
    cout << setw(10) << "Jan" << setw(10) << 30 << endl;
    return 0;
}
```

### Zapis do pliku wyjsciowego

```cpp
#include <iostream>
#include <fstream>
using namespace std;

int main() {
    ofstream out("wyniki.txt");
    out << "Wynik: " << 42 << endl;
    out.close();
    cout << "Zapisano do wyniki.txt" << endl;
    return 0;
}
```

---

## 7. Algorytmy zaawansowane

### Przeszukiwanie binarne

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

int main() {
    vector<int> tab = {2, 5, 8, 12, 16, 23, 38, 56, 72, 91};
    int x = 23;

    int left = 0, right = (int)tab.size() - 1;
    int result = -1;
    while (left <= right) {
        int mid = (left + right) / 2;
        if (tab[mid] == x) { result = mid; break; }
        if (tab[mid] < x) left = mid + 1;
        else right = mid - 1;
    }
    if (result != -1)
        cout << x << " na pozycji " << result << endl;
    else
        cout << x << " nie znaleziono" << endl;
    return 0;
}
```

### Rekurencja — wzorzec

```cpp
#include <iostream>
using namespace std;

void rekurencja(int n) {
    if (n <= 0) return; // warunek bazowy
    cout << n << " ";   // przed wywolaniem
    rekurencja(n - 1);  // zmniejsz problem
    cout << n << " ";   // po wywolaniu (odwrotna kolejnosc)
}

int main() {
    rekurencja(3); // 3 2 1 1 2 3
    cout << endl;
    return 0;
}
```

### Konwersje miedzy bazami

```cpp
#include <iostream>
#include <string>
using namespace std;

// base 10 -> base k
string toBase(int n, int k) {
    if (n == 0) return "0";
    string result = "";
    while (n > 0) {
        int reszta = n % k;
        if (reszta < 10) result = char('0' + reszta) + result;
        else result = char('A' + reszta - 10) + result;
        n /= k;
    }
    return result;
}

// base k -> base 10 (schemat Hornera)
int fromBase(string s, int k) {
    int result = 0;
    for (char c : s) {
        int digit;
        if (c >= '0' && c <= '9') digit = c - '0';
        else digit = c - 'A' + 10;
        result = result * k + digit;
    }
    return result;
}

int main() {
    cout << "255 w bin: " << toBase(255, 2) << endl;   // 11111111
    cout << "255 w hex: " << toBase(255, 16) << endl;   // FF
    cout << "FF z hex:  " << fromBase("FF", 16) << endl; // 255
    return 0;
}
```

### Dodawanie w dowolnym systemie

```cpp
#include <iostream>
#include <string>
using namespace std;

string addBase(string a, string b, int k) {
    string result = "";
    int carry = 0;
    int i = a.size() - 1, j = b.size() - 1;
    while (i >= 0 || j >= 0 || carry) {
        int sum = carry;
        if (i >= 0) sum += a[i--] - '0';
        if (j >= 0) sum += b[j--] - '0';
        result = char('0' + sum % k) + result;
        carry = sum / k;
    }
    return result;
}

int main() {
    cout << addBase("1101", "1011", 2) << endl; // 11000
    return 0;
}
```

### Programowanie dynamiczne (DP)

```cpp
#include <iostream>
using namespace std;

int main() {
    // DP 1D: Fibonacci
    const int N = 20;
    long long dp[N + 1];
    dp[0] = 0; dp[1] = 1;
    for (int i = 2; i <= N; i++)
        dp[i] = dp[i - 1] + dp[i - 2];
    cout << "Fib(" << N << ") = " << dp[N] << endl;
    return 0;
}
```

### DP 2D: sciezki na planszy

```cpp
#include <iostream>
using namespace std;

int main() {
    const int R = 4, C = 5;
    int plansza[R][C] = {
        {0, 0, 0, 1, 0},
        {0, 1, 0, 0, 0},
        {0, 0, 0, 1, 0},
        {0, 0, 0, 0, 0}
    };
    // 0 = wolne, 1 = blokada

    bool dp[R][C] = {};
    dp[0][0] = (plansza[0][0] == 0);
    for (int i = 0; i < R; i++)
        for (int j = 0; j < C; j++) {
            if (plansza[i][j] == 1) { dp[i][j] = false; continue; }
            if (i > 0 && dp[i - 1][j]) dp[i][j] = true;
            if (j > 0 && dp[i][j - 1]) dp[i][j] = true;
        }
    cout << "Mozna dojsc: " << (dp[R-1][C-1] ? "TAK" : "NIE") << endl;
    return 0;
}
```

### Algorytm zachlanny: dobor aktywnosci

```cpp
#include <iostream>
#include <vector>
#include <algorithm>
using namespace std;

struct Akt { int start, end; };

int main() {
    vector<Akt> a = {{1,4},{3,5},{0,6},{5,7},{3,9},{5,9},{6,10},{8,11},{8,12},{2,14},{12,16}};

    sort(a.begin(), a.end(), [](Akt &x, Akt &y) {
        return x.end < y.end;
    });

    int count = 1;
    int lastEnd = a[0].end;
    for (int i = 1; i < (int)a.size(); i++) {
        if (a[i].start >= lastEnd) {
            count++;
            lastEnd = a[i].end;
        }
    }
    cout << "Max aktywnosci: " << count << endl;
    return 0;
}
```

### BFS (graf / siatka)

```cpp
#include <iostream>
#include <vector>
#include <queue>
using namespace std;

const int N = 6;
int grid[N][N];
bool vis[N][N];
int dx[] = {-1, 1, 0, 0};
int dy[] = {0, 0, -1, 1};

int bfs(int si, int sj) {
    queue<pair<int,int>> q;
    q.push({si, sj});
    vis[si][sj] = true;
    int rozmiar = 0;
    while (!q.empty()) {
        auto [i, j] = q.front(); q.pop();
        rozmiar++;
        for (int d = 0; d < 4; d++) {
            int ni = i + dx[d], nj = j + dy[d];
            if (ni >= 0 && ni < N && nj >= 0 && nj < N
                && !vis[ni][nj] && grid[ni][nj] == 1) {
                vis[ni][nj] = true;
                q.push({ni, nj});
            }
        }
    }
    return rozmiar;
}

int main() {
    // ... wczytaj siatke do grid[N][N]
    int obszary = 0;
    for (int i = 0; i < N; i++)
        for (int j = 0; j < N; j++)
            if (grid[i][j] == 1 && !vis[i][j]) {
                int r = bfs(i, j);
                cout << "Obszar " << ++obszary << ": rozmiar " << r << endl;
            }
    return 0;
}
```

### DFS (flood fill)

```cpp
#include <iostream>
#include <vector>
using namespace std;

const int N = 6;
int grid[N][N];
bool vis[N][N];

void dfs(int x, int y, int &rozmiar) {
    if (x < 0 || x >= N || y < 0 || y >= N) return;
    if (vis[x][y] || grid[x][y] == 0) return;
    vis[x][y] = true;
    rozmiar++;
    dfs(x + 1, y, rozmiar);
    dfs(x - 1, y, rozmiar);
    dfs(x, y + 1, rozmiar);
    dfs(x, y - 1, rozmiar);
}

int main() {
    // ... wczytaj siatke
    int obszary = 0;
    for (int i = 0; i < N; i++)
        for (int j = 0; j < N; j++)
            if (grid[i][j] == 1 && !vis[i][j]) {
                int r = 0;
                dfs(i, j, r);
                cout << "Obszar " << ++obszary << ": rozmiar " << r << endl;
            }
    return 0;
}
```

### Struktury danych — skrocona sciagawka

```cpp
#include <iostream>
#include <vector>
#include <map>
#include <set>
#include <queue>
#include <stack>
using namespace std;

int main() {
    // vector
    vector<int> v = {1, 2, 3};
    v.push_back(4);       // dodaj
    v.size();             // 4
    v[0];                 // 1

    // map (slownik)
    map<string, int> m;
    m["a"] = 1;
    m.count("a");         // 1 (istnieje)

    // set (zbior)
    set<int> s;
    s.insert(5);
    s.count(5);           // 1

    // queue (kolejka FIFO)
    queue<int> q;
    q.push(10); q.push(20);
    q.front();            // 10
    q.pop();              // usun 10

    // stack (stos LIFO)
    stack<int> st;
    st.push(10); st.push(20);
    st.top();             // 20
    st.pop();             // usun 20

    return 0;
}
```

### Geometria: odleglosc + pole trojkata

```cpp
#include <iostream>
#include <cmath>
using namespace std;

double dist(int x1, int y1, int x2, int y2) {
    return sqrt((double)(x2-x1)*(x2-x1) + (y2-y1)*(y2-y1));
}

double poleTrojkata(int x1, int y1, int x2, int y2, int x3, int y3) {
    return abs(x1*(y2-y3) + x2*(y3-y1) + x3*(y1-y2)) / 2.0;
}

int main() {
    cout << "Odl: " << dist(0, 0, 3, 4) << endl;         // 5
    cout << "Pole: " << poleTrojkata(0,0,4,0,0,3) << endl; // 6
    return 0;
}
```

---

## 8. Typowe pulapki maturalne

| Pulapka | Problem | Rozwiazanie |
|---------|---------|-------------|
| `int` overflow | `100000 * 100000` przekracza `int` | `long long` |
| Dzielenie calkowite | `7 / 2 == 3` | `(double)7 / 2` lub `7.0 / 2` |
| Off-by-one | indeksy od 0, pozycje od 1 | `tab[i]` → pozycja `i + 1` |
| Porownanie `double` | `0.1 + 0.2 != 0.3` | `abs(a - b) < 1e-9` |
| `size_t` vs `int` | `s.length() - 1` przy pustym stringu = ogromna liczba | rzutuj: `(int)s.length()` |
| Niezamkniecie pliku | plik nie zostanie zrzucony na dysk | `plik.close()` lub RAII |
| Brak `endl` / `"\n"` | brak nowej linii na koncu | dodaj `endl` po kazdym wyniku |
| Bledna sciezka | plik nie w tym samym katalogu | sprawdz gdzie uruchamiasz program |
| Porownanie `char` | `'a' != 'A'` | toLower() przed porownaniem |
| Pusta linia na koncu | `getline` wczyta pusta linie | `if (!linia.empty())` |
