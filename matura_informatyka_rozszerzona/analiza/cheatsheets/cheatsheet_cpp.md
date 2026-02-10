# Cheatsheet: C++ (implementacja)

## Wczytywanie plikow (4 wzorce)

```cpp
// A. Liczby oddzielone spacjami/enterami
ifstream plik("dane.txt");
int x; vector<int> T;
while (plik >> x) T.push_back(x);

// B. Linia-po-linii (string -> int)
ifstream plik("dane.txt");
string linia;
while (getline(plik, linia)) { int v = stoi(linia); }

// C. CSV ze srednikiem (+ pomin naglowek)
ifstream plik("dane.txt");
string linia;
getline(plik, linia); // pomin naglowek
while (getline(plik, linia)) {
    stringstream ss(linia);
    string token; vector<string> kol;
    while (getline(ss, token, ';')) kol.push_back(token);
    // kol[0], kol[1], ... — uzyj stoi() dla liczb
}

// D. Pary/trojki na linie
ifstream plik("dane.txt");
int a, b;
while (plik >> a >> b) { /* a, b */ }
```

---

## Przetwarzanie cyfr

```cpp
// Suma cyfr
int suma = 0;
while (n > 0) { suma += n % 10; n /= 10; }

// Odwracanie liczby
int odwr = 0;
while (n > 0) { odwr = odwr * 10 + n % 10; n /= 10; }

// Liczba cyfr
int cnt = 0;
while (n > 0) { cnt++; n /= 10; }
```

---

## NWD / NWW / Pierwszosc / Sito

```cpp
int nwd(int a, int b) {
    while (b) { int t = b; b = a % b; a = t; } return a;
}
int nww(int a, int b) { return a / nwd(a, b) * b; }

bool pierwsza(int n) {
    if (n < 2) return false;
    for (int i = 2; i * i <= n; i++) if (n % i == 0) return false;
    return true;
}

// Sito Eratostenesa
bool sito[MAX]; // fill(sito, sito+MAX, true); sito[0]=sito[1]=false;
for (int i = 2; i * i < MAX; i++)
    if (sito[i]) for (int j = i*i; j < MAX; j += i) sito[j] = false;
```

---

## Systemy liczbowe

```cpp
// dec -> base k (wynik jako string)
string toBase(int n, int k) {
    string r = ""; while (n > 0) { r = char('0'+n%k) + r; n /= k; } return r;
}
// base k -> dec (schemat Hornera)
int fromBase(string s, int k) {
    int r = 0; for (char c : s) r = r * k + (c - '0'); return r;
}
```

---

## Napisy

```cpp
s.length()           // dlugosc
s[i]                 // i-ty znak (od 0)
s.substr(pos, len)   // podciag
s.find("abc")        // pozycja lub string::npos
s + "xyz"            // laczenie
stoi("42")           // string -> int
to_string(42)        // int -> string
(int)'A'             // 65;  (char)97 -> 'a';  'a'-'A' = 32
```

---

## Sortowanie

```cpp
sort(v.begin(), v.end());                               // rosnaco
sort(v.begin(), v.end(), [](int a, int b){ return a > b; }); // malejaco
// Wg klucza custom:
sort(v.begin(), v.end(), [](auto &a, auto &b){
    if (a.x != b.x) return a.x < b.x;  // priorytet 1
    return a.y > b.y;                    // priorytet 2
});
```

---

## Zliczanie / Min-Max / Najdluzszy ciag

```cpp
// Zliczanie
int ile = 0; for (auto x : T) if (warunek(x)) ile++;

// Min/Max z pozycja
int mn = T[0], idx = 0;
for (int i = 1; i < n; i++) if (T[i] < mn) { mn = T[i]; idx = i; }

// Najdluzszy ciag spelniajacy warunek (current/max)
int cur = 1, mx = 1;
for (int i = 1; i < n; i++) {
    if (T[i] > T[i-1]) cur++;
    else { mx = max(mx, cur); cur = 1; }
}
mx = max(mx, cur); // NIE ZAPOMNIJ po petli!

// Mapa czestotliwosci
map<int, int> freq;
for (auto x : T) freq[x]++;
```

---

## Kontenery STL — 1 linijka uzycia

```cpp
vector<int> v;   v.push_back(x); v.size(); v[i];
map<string,int> m; m["k"]=1; m.count("k"); // 1=jest
set<int> s;      s.insert(x); s.count(x);
stack<int> st;   st.push(x); st.top(); st.pop();
queue<int> q;    q.push(x); q.front(); q.pop();
```

---

## Pulapki

| Pulapka | Rozwiazanie |
|---------|-------------|
| `int` overflow (`100000*100000`) | `long long` |
| Dzielenie calkowite (`7/2=3`) | `(double)7/2` lub `7.0/2` |
| Off-by-one (indeks 0 vs pozycja 1) | `tab[i]` = pozycja `i+1` |
| `s.length()-1` przy pustym stringu | rzutuj: `(int)s.length()` |
| Plik nie otwarty | sprawdz `if (!plik.is_open())` |
| Min init = 0 (dane ujemne!) | init z `T[0]` |
| Brak sprawdzenia po petli | `mx = max(mx, cur)` po for |
