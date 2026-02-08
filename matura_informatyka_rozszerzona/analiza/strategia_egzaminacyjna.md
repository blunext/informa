# 🎯 Strategia Zdobycia 50/50 Punktów na Maturze Rozszerzonej z Informatyki

## 📖 Definicje i Podstawy Teoretyczne

### 1. Sortowanie
**Definicja**: Proces uporządkowania elementów według określonego kryterium (rosnąco lub malejąco).

**Główne zastosowania**:
- Uporządkowanie danych dla lepszej czytelności
- Przygotowanie danych do przeszukiwania binarnego
- Znajdowanie wartości ekstremalnych (min/max)
- Grupowanie podobnych elementów
- Wykrywanie duplikatów

**Kluczowe właściwości**:
- Złożoność: O(n²) dla prostych algorytmów (bubble, selection), O(n log n) dla zaawansowanych (merge, quick)
- **Sortowanie stabilne**: zachowuje względną kolejność równych elementów
- **In-place**: sortuje bez używania dodatkowej pamięci

**Kiedy używać**: Zawsze gdy potrzebujesz uporządkowanych danych lub chcesz użyć przeszukiwania binarnego.

---

### 2. Przeszukiwanie Binarne (Binary Search)
**Definicja**: Algorytm wyszukiwania elementu w **posortowanej** tablicy przez wielokrotne dzielenie zakresu na pół.

**Główne zastosowania**:
- Wyszukiwanie wartości w posortowanej tablicy
- Znajdowanie pierwiastka liczby (2018 - pierwiastek sześcienny)
- Metoda bisekcji dla funkcji ciągłych (2014 - znajdowanie miejsca zerowego)
- Znajdowanie wartości granicznej (pierwsza/ostatnia wartość spełniająca warunek)
- Optymalizacja odpowiedzi (znajdowanie minimum/maksimum przy ograniczeniach)

**Kluczowe właściwości**:
- **Wymaga posortowanych danych**
- Złożoność: O(log n) - bardzo szybki!
- Działa przez "dziel i zwyciężaj"
- Liczba kroków: ⌈log₂(n)⌉

**Kiedy używać**:
- Gdy masz posortowaną tablicę i szukasz konkretnej wartości
- Gdy musisz znaleźć pierwiastek/miejsce zerowe funkcji
- Gdy chcesz zoptymalizować algorytm O(n) do O(log n)

**Warunek konieczny**: Dane muszą być posortowane lub funkcja musi być monotoniczna!

---

### 3. Rekurencja
**Definicja**: Technika, w której funkcja wywołuje samą siebie z mniejszym problemem, aż osiągnie przypadek bazowy.

**Główne zastosowania**:
- Przetwarzanie struktur drzewiastych (drzewa, katalogi)
- Dzielenie problemu na mniejsze podproblemy (dziel i zwyciężaj)
- Obliczenia matematyczne (silnia, Fibonacci, NWD)
- Generowanie permutacji i kombinacji
- Algorytmy backtracking

**Kluczowe właściwości**:
- Wymaga **warunku bazowego** (przypadek, który kończy rekurencję)
- Każde wywołanie rekurencyjne to **mniejszy problem**
- Używa stosu wywołań (może prowadzić do stack overflow dla dużych n)
- **Odwraca kolejność**: akcje wykonują się "od tyłu" przy powrocie z rekurencji

**Kiedy używać**:
- Gdy problem naturalnie dzieli się na mniejsze podobne podproblemy
- Gdy struktura danych jest rekurencyjna (drzewo)
- Gdy iteracyjne rozwiązanie jest bardzo skomplikowane

**Uwaga**: Często na maturze trzeba konwertować rekurencję → iterację lub odwrotnie!

---

### 4. Operacje na Liczbach
**Definicja**: Algorytmy z teorii liczb operujące na liczbach całkowitych: cyfry, dzielniki, liczby pierwsze, NWD/NWW.

**Główne zastosowania**:

**a) Cyfry liczby** (mod 10, div 10):
- Sumowanie/mnożenie cyfr
- Sprawdzanie właściwości cyfr
- Odwracanie liczby
- Konwersje systemów liczbowych

**b) NWD (Największy Wspólny Dzielnik)**:
- Algorytm Euklidesa
- Skracanie ułamków
- Rozwiązywanie równań diofantycznych

**c) Liczby pierwsze**:
- Test pierwszości
- Sito Eratostenesa (generowanie wielu liczb pierwszych)
- Faktoryzacja (rozkład na czynniki pierwsze)

**d) Dzielniki**:
- Znajdowanie wszystkich dzielników
- Suma dzielników
- Liczba dzielników

**Kluczowe właściwości**:
- Operacje mod/div działają w O(log n) - liczba cyfr
- Sito Eratostenesa: O(n log log n)
- NWD: O(log min(a,b))
- **Dzielniki tylko do √n** - kluczowa optymalizacja!

**Kiedy używać**: Teoria liczb jest stałym elementem matury!

---

### 5. Przetwarzanie Plików
**Definicja**: Operacje wejścia/wyjścia (I/O) - czytanie danych z plików i zapisywanie wyników.

**Główne zastosowania**:
- Czytanie danych wejściowych do programu
- Zapisywanie wyników obliczeń
- Przetwarzanie dużych zbiorów danych
- Parsowanie formatów (CSV, JSON-like, custom)

**Kluczowe operacje**:
- `ifstream` - czytanie z pliku
- `ofstream` - zapisywanie do pliku
- `>>` operator - czytanie liczb/słów (pomija whitespace)
- `getline()` - czytanie całych linii
- `stringstream` - parsowanie linii na tokeny

**Kluczowe właściwości**:
- Zawsze zamykaj pliki (lub użyj RAII)
- Sprawdzaj czy plik się otworzył: `if (!file.is_open())`
- `>>` pomija spacje/entery, `getline()` czyta z enterami

**Kiedy używać**: W KAŻDYM zadaniu praktycznym część II!

---

### 6. SQL (Structured Query Language)
**Definicja**: Język zapytań do relacyjnych baz danych służący do manipulacji i pobierania danych.

**Główne zastosowania**:
- Wyszukiwanie danych (SELECT)
- Łączenie tabel (JOIN)
- Agregacje i statystyki (COUNT, SUM, AVG)
- Filtrowanie (WHERE)
- Grupowanie (GROUP BY)
- Sortowanie (ORDER BY)

**Kluczowe operacje**:

**JOIN** - łączenie tabel:
- `INNER JOIN` - tylko rekordy pasujące w obu tabelach
- `LEFT JOIN` - wszystkie z lewej + pasujące z prawej (NULL jeśli brak)
- `RIGHT JOIN` - wszystkie z prawej + pasujące z lewej

**Agregacje** - obliczenia na grupach:
- `COUNT(*)` - liczba rekordów
- `SUM(kolumna)` - suma wartości
- `AVG(kolumna)` - średnia
- `MIN/MAX(kolumna)` - minimum/maksimum

**Kluczowe właściwości**:
- SELECT określa **co** wybrać
- FROM określa **skąd** (które tabele)
- WHERE filtruje **rekordy** (przed agregacją)
- GROUP BY grupuje dane
- HAVING filtruje **grupy** (po agregacji)
- ORDER BY sortuje wynik

**Kiedy używać**: W każdym zadaniu z bazami danych!

**Najczęstszy błąd**: Zapomnienie GROUP BY przy agregacjach!

---

### 7. Operacje na Stringach
**Definicja**: Manipulacja tekstem - wyszukiwanie, modyfikacja, porównywanie ciągów znaków.

**Główne zastosowania**:
- Przetwarzanie tekstów
- Parsowanie danych
- Szyfrowanie/deszyfrowanie
- Operacje na kodach ASCII
- Pattern matching

**Kluczowe operacje**:
- Długość: `s.length()`
- Dostęp do znaku: `s[i]`
- Podciąg: `s.substr(pos, len)`
- Konkatenacja: `s1 + s2`
- Porównywanie: `s1 == s2` (leksykograficzne)
- Kody ASCII: `(int)znak` i `(char)kod`

**Kluczowe właściwości**:
- Stringi są indeksowane od 0
- Kod ASCII: 'A'=65, 'a'=97, '0'=48
- Różnica 'a'-'A' = 32
- Porównywanie stringów jest leksykograficzne (alfabetyczne)

**Kiedy używać**: Często w zadaniach z plikami tekstowymi i szyfrowaniem.

---

### 8. Struktury Danych
**Definicja**: Sposoby organizacji i przechowywania danych umożliwiające efektywny dostęp i modyfikację.

**Główne struktury**:

**Vector** - tablica dynamiczna:
- Szybki dostęp O(1)
- Dodawanie na końcu O(1) amortyzowane
- Automatyczne zarządzanie pamięcią

**Map** - słownik (klucz → wartość):
- Przechowuje pary klucz-wartość
- Szybkie wyszukiwanie O(log n)
- Klucze są unikalne i posortowane

**Set** - zbiór unikalnych elementów:
- Automatycznie usuwa duplikaty
- Elementy posortowane
- Szybkie sprawdzanie przynależności O(log n)

**Queue** - kolejka FIFO:
- First In First Out
- Dodawanie z tyłu, usuwanie z przodu
- BFS (przeszukiwanie wszerz)

**Stack** - stos LIFO:
- Last In First Out
- Dodawanie i usuwanie ze szczytu
- DFS, symulacja rekurencji

**Kiedy używać**:
- Vector: domyślny wybór dla tablic
- Map: grupowanie, liczenie wystąpień, indeksowanie
- Set: usuwanie duplikatów, sprawdzanie unikalności
- Queue: BFS, przetwarzanie kolejkowe
- Stack: DFS, parsowanie, symulacja rekurencji

---

### 9. Geometria i Matematyka
**Definicja**: Algorytmy geometryczne operujące na punktach, odcinkach, wielokątach oraz obliczenia matematyczne.

**Główne zastosowania**:

**Geometria obliczeniowa**:
- Odległość między punktami: √((x₂-x₁)² + (y₂-y₁)²)
- Sortowanie punktów według kąta (2018 - Krajobraz)
- Sprawdzanie kolizji
- Pole wielokąta

**Matematyka dyskretna**:
- Ciągi arytmetyczne/geometryczne
- Kombinatoryka: C(n,k) = n!/(k!(n-k)!)
- Permutacje: n!
- Wzory skrócone

**Kluczowe wzory**:
- Suma ciągu arytmetycznego: n(a₁ + aₙ)/2
- Suma 1+2+...+n = n(n+1)/2
- Uściski n osób: n(n-1)/2
- Kombinacje C(n,2) = n(n-1)/2

**Kiedy używać**: W zadaniach geometrycznych i kombinatorycznych.

---

### 10. Złożoność Algorytmów
**Definicja**: Miara efektywności algorytmu - ile czasu/pamięci potrzebuje w zależności od rozmiaru danych (notacja O).

**Główne klasy złożoności**:

**O(1)** - stała:
- Dostęp do elementu tablicy: `arr[i]`
- Operacje arytmetyczne
- Przykład: sprawdzenie parzystości

**O(log n)** - logarytmiczna:
- Przeszukiwanie binarne
- Operacje na zbilansowanych drzewach
- Dzielenie zakresu na pół

**O(n)** - liniowa:
- Przejście przez tablicę
- Szukanie minimum/maksimum
- Sumowanie elementów

**O(n log n)** - liniowo-logarytmiczna:
- Efektywne sortowanie (merge sort, quick sort)
- Najszybsze możliwe sortowanie przez porównania

**O(n²)** - kwadratowa:
- Dwie zagnieżdżone pętle
- Proste algorytmy sortowania (bubble, selection)
- Sprawdzanie wszystkich par

**O(2ⁿ)** - eksponencjalna:
- Sprawdzanie wszystkich podzbiorów
- UNIKAĆ! Działa tylko dla małych n (<20)

**Optymalizacje**:
- n² → n: usunięcie zagnieżdżonej pętli
- n → log n: posortuj i użyj binary search
- n → √n: dla dzielników - sprawdzaj tylko do √n

**Kiedy używać**: Zawsze! Na maturze często pytają o złożoność.

---

## 📊 TOP 10 Najważniejszych Algorytmów (Wg częstości 2014-2025)

### 1. **Sortowanie** (100% egzaminów)
**Dlaczego**: Podstawa, pojawia się ZAWSZE
**Co umieć**:
```cpp
// Bubble sort (najprostszy)
for (int i = 0; i < n-1; i++)
    for (int j = 0; j < n-i-1; j++)
        if (arr[j] > arr[j+1])
            swap(arr[j], arr[j+1]);

// Selection sort
for (int i = 0; i < n-1; i++) {
    int min_idx = i;
    for (int j = i+1; j < n; j++)
        if (arr[j] < arr[min_idx])
            min_idx = j;
    swap(arr[i], arr[min_idx]);
}

// Sortowanie z kluczem niestandardowym (WAŻNE!)
sort(arr, arr+n, [](int a, int b) {
    return custom_key(a) < custom_key(b);
});
```
**Typowe zastosowania**: Uporządkowanie danych, znajdowanie min/max, grupowanie

---

### 2. **Przeszukiwanie Binarne** (80%+ egzaminów)
**Dlaczego**: Efektywne O(log n), często w przebraniu
**Co umieć**:
```cpp
// Klasyczne binary search
int binary_search(int arr[], int n, int x) {
    int left = 0, right = n-1;
    while (left <= right) {
        int mid = (left + right) / 2;
        if (arr[mid] == x) return mid;
        if (arr[mid] < x) left = mid + 1;
        else right = mid - 1;
    }
    return -1;
}

// Binary search dla funkcji (bisekcja, pierwiastek)
int find_root(int n) {
    int p = 1, q = n;
    while (p < q) {
        int s = (p + q) / 2;
        if (condition(s, n)) p = s + 1;
        else q = s;
    }
    return p;
}
```
**Typowe zastosowania**: Znajdowanie pierwiastka, bisekcja, wyszukiwanie w posortowanej tablicy

---

### 3. **Rekurencja** (70%+ egzaminów)
**Dlaczego**: Eleganckie rozwiązania, często konwersja rekurencja↔iteracja
**Co umieć**:
```cpp
// Wzorzec rekurencyjny
void rekurencja(int n) {
    if (warunek_bazowy) { /* akcja */ return; }
    rekurencja(mniejszy_problem);  // PRZED akcją
    /* akcja */                      // PO rekurencji
}

// Konwersja na iterację (STACK/QUEUE)
void iteracja(int n) {
    stack<int> s;
    while (!s.empty() || warunek) {
        // symuluj rekurencję używając stosu
    }
}
```
**Typowe zastosowania**: Korale (2014), drzewa, podział problemu

---

### 4. **Operacje na Liczbach** (70%+ egzaminów)
**Dlaczego**: Teoria liczb to klasyka
**Co umieć**:
```cpp
// Cyfry liczby
while (n > 0) {
    int cyfra = n % 10;
    n = n / 10;
}

// NWD (algorytm Euklidesa)
int nwd(int a, int b) {
    while (b != 0) {
        int temp = b;
        b = a % b;
        a = temp;
    }
    return a;
}

// Sito Eratostenesa (liczby pierwsze)
bool prime[MAX];
void sieve(int n) {
    fill(prime, prime+n+1, true);
    prime[0] = prime[1] = false;
    for (int i = 2; i*i <= n; i++)
        if (prime[i])
            for (int j = i*i; j <= n; j += i)
                prime[j] = false;
}

// Dzielniki (OPTYMALIZACJA!)
vector<int> divisors(int n) {
    vector<int> div;
    for (int i = 1; i*i <= n; i++) {  // TYLKO do √n !!!
        if (n % i == 0) {
            div.push_back(i);
            if (i != n/i) div.push_back(n/i);
        }
    }
    return div;
}
```

---

### 5. **Przetwarzanie Plików** (70%+ egzaminów)
**Dlaczego**: Część II zawsze wymaga czytania plików
**Co umieć**:
```cpp
#include <fstream>
#include <string>

// Czytanie liczb
ifstream file("dane.txt");
int n;
while (file >> n) {
    // przetwarzaj n
}

// Czytanie linii
string line;
while (getline(file, line)) {
    // przetwarzaj line
}

// Parsowanie CSV
while (getline(file, line)) {
    stringstream ss(line);
    string token;
    vector<string> tokens;
    while (getline(ss, token, ',')) {
        tokens.push_back(token);
    }
}

// Zapisywanie
ofstream out("wyniki.txt");
out << wynik << endl;
```

---

### 6. **SQL (JOIN + Agregacje)** (100% egzaminów z bazami)
**Dlaczego**: Bazy danych to obowiązkowy element
**Co umieć**:
```sql
-- Podstawy
SELECT kolumna1, kolumna2 FROM tabela WHERE warunek;

-- JOIN (NAJWAŻNIEJSZE!)
SELECT t1.col, t2.col
FROM tabela1 t1
INNER JOIN tabela2 t2 ON t1.id = t2.id_foreign
WHERE warunek;

-- Agregacje
SELECT kategoria, COUNT(*), SUM(wartosc), AVG(wartosc)
FROM tabela
GROUP BY kategoria
HAVING COUNT(*) > 5
ORDER BY SUM(wartosc) DESC;

-- LIKE (wzorce)
WHERE nazwa LIKE 'A%'  -- zaczyna się na A

-- Subqueries
SELECT * FROM tabela1
WHERE id IN (SELECT id_foreign FROM tabela2 WHERE warunek);
```

---

### 7. **Operacje na Stringach** (60%+ egzaminów)
**Co umieć**:
```cpp
string s = "Hello";
s.length()          // długość
s[i]                // i-ty znak
s.substr(pos, len)  // podstring
s + "World"         // konkatenacja

// Kody ASCII
int kod = (int)'A';  // 65
char znak = (char)65; // 'A'

// Suma kodów
int suma = 0;
for (char c : s) suma += (int)c;

// Min/Max kod
char min_char = *min_element(s.begin(), s.end());
char max_char = *max_element(s.begin(), s.end());
```

---

### 8. **Struktury Danych** (60%+ egzaminów)
**Co umieć**:
```cpp
// Vector (tablica dynamiczna)
vector<int> v;
v.push_back(x);     // dodaj na koniec
v.size()            // rozmiar
v[i]                // dostęp
sort(v.begin(), v.end());

// Map (słownik)
map<string, int> m;
m["klucz"] = wartość;
m.count("klucz")    // czy istnieje?

// Set (zbiór unikalnych)
set<int> s;
s.insert(x);
s.count(x)          // czy istnieje?

// Queue
queue<int> q;
q.push(x);          // dodaj
q.front();          // pierwszy
q.pop();            // usuń pierwszy

// Stack
stack<int> st;
st.push(x);         // dodaj
st.top();           // szczyt
st.pop();           // usuń ze szczytu
```

---

### 9. **Geometria/Matematyka** (40%+ egzaminów)
**Co umieć**:
```cpp
// Odległość punktów
double dist(int x1, int y1, int x2, int y2) {
    return sqrt((x2-x1)*(x2-x1) + (y2-y1)*(y2-y1));
}

// Kąt od początku układu (2018 - Krajobraz)
double angle(int x, int y) {
    return (double)x / y;  // klucz sortowania
}

// Suma ciągu
int suma_arytm(int a1, int n, int r) {
    return n * (2*a1 + (n-1)*r) / 2;
}

// Kombinatoryka
int newton(int n, int k) {
    // C(n,k) = n! / (k! * (n-k)!)
    // Dla uścisków: n*(n-1)/2
}
```

---

### 10. **Złożoność Algorytmów** (60%+ egzaminów)
**Co umieć**:
- O(1) - stała
- O(log n) - logarytmiczna (binary search)
- O(n) - liniowa (iteracja przez tablicę)
- O(n log n) - sortowanie
- O(n²) - podwójna pętla
- O(2ⁿ) - eksponencjalna (unikać!)

**Optymalizacje**:
- Dzielniki: tylko do √n
- Liczby pierwsze: sito zamiast sprawdzania każdej
- Wyszukiwanie: posortuj i użyj binary search

---

## 🎯 Strategia Punktowa

### ETAP 1: Quick Wins (5-10 pkt, 10-15 min) ⚡
**Cel**: Zdobyć łatwe punkty na start

**Zadania**:
- Pytania P/F (prawda/fałsz)
- Proste obliczenia
- Śledzenie algorytmu dla małych danych

**Strategia**:
1. Przejrzyj cały arkusz (2 min)
2. Zaznacz zadania P/F i proste obliczenia
3. Zrób je WSZYSTKIE najpierw (10 min)
4. **Cel**: 5-10 punktów w 15 minut

---

### ETAP 2: Zadania Standardowe (20-30 pkt, 60-90 min) 📝
**Cel**: Zdobyć większość punktów

**Zadania**:
- Analiza i śledzenie algorytmów
- Projektowanie prostych algorytmów
- SQL queries (JOIN, GROUP BY)
- Excel (formuły, wykresy)
- Programowanie (czytanie plików, proste algorytmy)

**Strategia**:
1. Zacznij od zadań, które ZNASZ
2. SQL i Excel to często szybkie punkty
3. Programowanie: najpierw podzadania łatwe
4. **Cel**: 25-35 punktów łącznie (z Etap 1)

---

### ETAP 3: Zadania Trudne (15-20 pkt, 60-90 min) 🔥
**Cel**: Dobić do maksimum

**Zadania**:
- Złożone algorytmy (DP, greedy)
- Optymalizacje (złożoność O(n²) lub mniejsza)
- Trudne zadania programistyczne
- Złożone zapytania SQL

**Strategia**:
1. Jeśli utkniesz > 15 min, przejdź dalej
2. Wróć na koniec jeśli będzie czas
3. Częściowe rozwiązanie > brak rozwiązania
4. **Cel**: 45-50 punktów łącznie

---

## ⏱️ Time Management

**Całkowity czas**: 180-210 minut (w zależności od formuły)

### Podział czasu (Formuła 2015, 2-częściowa):
- **Część I** (90 min, ~20 pkt):
  - Quick wins: 10-15 min → 5-8 pkt
  - Średnie: 40-50 min → 10-15 pkt
  - Trudne: 25-35 min → 5-10 pkt
  - Bufor: 5 min

- **Część II** (120 min, ~30 pkt):
  - Quick wins: 10 min → 5 pkt
  - Średnie: 60-70 min → 15-20 pkt
  - Trudne: 40-50 min → 10-15 pkt
  - Bufor: 10 min

### Podział czasu (Formuła 2018+, jednoczęściowa?):
- Do ustalenia po pełnej analizie

---

## ⚠️ Checklist Przed Egzaminem

### Dzień przed:
- [ ] Powtórz TOP 10 algorytmów
- [ ] Przejrzyj wzorce kodu (templates)
- [ ] Powtórz SQL (JOIN, GROUP BY, agregacje)
- [ ] Sprawdź czy znasz środowisko (IDE, kompilator)
- [ ] Przygotuj ściągę mentalną (wzory, złożoności)

### Na egzaminie:
- [ ] Przeczytaj CAŁY arkusz (2 min)
- [ ] Zaznacz zadania quick-win
- [ ] Zrób najpierw quick-wins
- [ ] Wracaj do trudnych tylko jeśli zostanie czas
- [ ] Zostaw 10 min na koniec na sprawdzenie

---

## 🐛 TOP 10 Pułapek

### 1. **Int Overflow**
❌ `int result = a * b;` gdy a, b duże
✅ `long long result = (long long)a * b;`

### 2. **Dzielenie Całkowite**
❌ `double avg = suma / n;` gdy suma, n to int
✅ `double avg = (double)suma / n;`

### 3. **Indeksowanie od 0 vs 1**
❌ Tablica w C++ od 0, zadanie od 1
✅ Uważaj na przesunięcie indeksów

### 4. **Porównywanie Float**
❌ `if (a == b)` dla double
✅ `if (abs(a - b) < 1e-9)`

### 5. **SQL: LEFT JOIN vs INNER JOIN**
❌ INNER JOIN pomija NULL
✅ LEFT JOIN zachowuje wszystkie z lewej

### 6. **SQL: GROUP BY**
❌ SELECT bez GROUP BY przy agregacji
✅ GROUP BY wszystkie nieagregowane kolumny

### 7. **Dzielniki tylko do √n**
❌ `for (int i = 1; i <= n; i++)`
✅ `for (int i = 1; i*i <= n; i++)`

### 8. **Sortowanie stabilne**
❌ Mylenie kolejności przy równych kluczach
✅ Użyj stable_sort lub sprawdź dokładnie

### 9. **String: zawinięcie cykliczne**
❌ `(i + k) % n` dla ujemnych
✅ `((i + k) % n + n) % n`

### 10. **Rekurencja: kolejność akcji**
❌ Myślenie, że akcja przed rekurencją = pierwsza
✅ Rekurencja odwraca kolejność!

---

## 📚 Materiały Referencyjne

### C++ Quick Reference
```cpp
// Includes
#include <iostream>   // cin, cout
#include <fstream>    // ifstream, ofstream
#include <string>     // string
#include <vector>     // vector
#include <algorithm>  // sort, min, max
#include <map>        // map
#include <set>        // set
#include <queue>      // queue, priority_queue
#include <stack>      // stack
#include <cmath>      // sqrt, pow, abs

// Useful functions
sort(arr, arr+n);                    // sortowanie
reverse(arr, arr+n);                 // odwracanie
*min_element(arr, arr+n);            // minimum
*max_element(arr, arr+n);            // maksimum
count(arr, arr+n, value);            // zliczanie
find(arr, arr+n, value);             // wyszukiwanie
```

### SQL Quick Reference
```sql
-- Struktura podstawowa
SELECT [DISTINCT] kolumny
FROM tabela1 [AS t1]
[INNER|LEFT|RIGHT] JOIN tabela2 [AS t2] ON warunek_join
WHERE warunek
GROUP BY kolumny
HAVING warunek_na_agregacje
ORDER BY kolumny [ASC|DESC]
LIMIT n;

-- Agregacje
COUNT(*), COUNT(kolumna), SUM(kolumna),
AVG(kolumna), MIN(kolumna), MAX(kolumna)

-- Wzorce LIKE
% - dowolny ciąg znaków
_ - jeden dowolny znak
```

---

## 🎯 Cel: 50/50 Punktów

**Strategia**:
1. Quick wins: 100% accuracy → 8-10 pkt
2. Średnie: 90% accuracy → 25-30 pkt
3. Trudne: 70% accuracy → 12-15 pkt
4. **TOTAL**: 45-50 punktów

**Time management**:
- Nie utknij na jednym zadaniu > 20 min
- Wracaj do trudnych na końcu
- Częściowe rozwiązanie lepsze niż brak

**Mental preparation**:
- Znasz TOP 10 algorytmów
- Masz wzorce kodu w głowie
- Umiesz SQL i Excel
- Jesteś gotowy/gotowa! 🚀

---

**Powodzenia!** 🎓
