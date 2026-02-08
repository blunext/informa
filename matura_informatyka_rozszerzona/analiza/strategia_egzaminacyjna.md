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

### 11. Programowanie Dynamiczne (DP)
**Definicja**: Technika rozwiązywania problemów przez podział na **nakładające się podproblemy** i zapamiętywanie ich wyników, aby uniknąć wielokrotnych obliczeń.

**Różnica rekurencja vs DP**:
- Rekurencja: rozwiązuje te same podproblemy wielokrotnie → wolne O(2ⁿ)
- DP: zapamiętuje wyniki podproblemów w tablicy → szybkie O(n) lub O(n²)

**Główne zastosowania**:
- Najdłuższy wspólny podciąg (LCS)
- Problem plecakowy (knapsack)
- Ścieżki na planszy/siatce (2024 - Plansza!)
- Najdłuższy rosnący podciąg (LIS)
- Fibonacci (klasyczny przykład)
- Podział zbioru na podzbiory

**Dwa podejścia**:
1. **Top-down (memoizacja)**: rekurencja + tablica cache
2. **Bottom-up (tabulacja)**: wypełnianie tablicy od najmniejszych podproblemów

**Kluczowe właściwości**:
- Wymaga **optymalnej podstruktury** (optymalne rozwiązanie składa się z optymalnych podrozwiązań)
- Wymaga **nakładających się podproblemów**
- Złożoność zależy od rozmiaru tablicy DP: O(n), O(n²), O(n×m)

**Kiedy używać**:
- Gdy widzisz słowa: "najdłuższy", "minimalny koszt", "liczba sposobów", "czy jest możliwe"
- Gdy rekurencja jest za wolna (drzewo wywołań rośnie eksponencjalnie)
- Gdy problem można rozłożyć na podproblemy z powtórzeniami

**Przykład maturalny**: Matura 2024 zad. 1 - sprawdzanie istnienia ścieżki na planszy (DP 2D)

---

### 12. Algorytmy Zachłanne (Greedy)
**Definicja**: Algorytm, który w każdym kroku wybiera **lokalnie najlepszą opcję**, licząc na to, że prowadzi do globalnie optymalnego rozwiązania.

**Główne zastosowania**:
- Problem doboru aktywności (activity selection) - matura 2015!
- Problem wydawania reszty (monety)
- Algorytm Kruskala/Prima (minimalne drzewo rozpinające)
- Kompresja Huffmana
- Scheduling (planowanie zadań)
- Problemy z deadline'ami

**Kluczowe właściwości**:
- Prosty w implementacji (zazwyczaj: posortuj + iteruj)
- **NIE ZAWSZE daje optymalne rozwiązanie** - trzeba udowodnić poprawność
- Złożoność: często O(n log n) ze względu na sortowanie
- Wymaga **właściwości zachłannej**: lokalny wybór prowadzi do globalnego optimum

**Schemat algorytmu zachłannego**:
1. Posortuj dane według odpowiedniego kryterium
2. Iteruj przez posortowane dane
3. Na każdym kroku podejmij zachłanną decyzję (weź lub pomiń)

**Kiedy używać**:
- Gdy problem ma strukturę wyboru: "weź lub pomiń"
- Gdy po posortowaniu widać naturalną strategię
- Gdy trzeba znaleźć minimum/maksimum i problem ma podstruktury zachłanne

**Przykład maturalny**: Matura 2015 zad. 1 - problem doboru aktywności (sortowanie po czasie zakończenia)

---

### 13. Systemy Liczbowe
**Definicja**: Sposoby reprezentacji liczb w różnych bazach (dziesiętny, binarny, ósemkowy, szesnastkowy, dowolny).

**Główne systemy**:
- **Binarny (base 2)**: 0, 1 → używany w komputerach
- **Ósemkowy (base 8)**: 0-7
- **Dziesiętny (base 10)**: 0-9 → codziennego użytku
- **Szesnastkowy (base 16)**: 0-9, A-F → adresy pamięci, kolory

**Główne operacje**:
- **Konwersja base 10 → base k**: dzielenie z resztą
- **Konwersja base k → base 10**: sumowanie cyfra × pozycja (schemat Hornera)
- **Dodawanie/odejmowanie w systemie k**: jak dziesiętne, ale przeniesienie przy k
- **Relacje**: 1 cyfra hex = 4 bity, 1 cyfra oct = 3 bity

**Schemat Hornera** (szybka konwersja base k → base 10):
- Zamiast: d₃×k³ + d₂×k² + d₁×k + d₀
- Oblicz: ((d₃×k + d₂)×k + d₁)×k + d₀

**Kluczowe właściwości**:
- Liczba cyfr w systemie k: ⌊log_k(n)⌋ + 1
- Konwersja bin→hex: grupuj po 4 bity od prawej
- Konwersja bin→oct: grupuj po 3 bity od prawej

**Kiedy używać**: Pojawia się na maturze regularnie, zarówno w zadaniach P/F jak i implementacyjnych.

**Przykład maturalny**: Matura 2023 zad. 6 - dodawanie/odejmowanie w systemach 3 i 9

---

## Przewodnik po Typach Zadan

Analiza 11 lat matur (2014-2025) wykazala 23 rozne typy zadan w 4 kategoriach.
Pelna macierz punktow: `ranking_typow_zadan.csv`

---

### KATEGORIA: TEORIA I ANALIZA

#### 1. sledzenie_algorytmu
**Co to jest**: Przesledzic algorytm krok po kroku dla podanych danych, podac wynik.
**Czestotliwosc**: 11/11 lat, 45 pkt lacznie — PEWNE na egzaminie!
**Jak podejsc**:
1. Zrob tabele zmiennych (kolumny = zmienne, wiersze = kroki)
2. Wykonuj instrukcje linia po linii, zapisujac wartosci
3. Uwazaj na kolejnosc operacji w rekurencji (odwrotna!)
4. Sprawdz wynik dla danych przykladowych
**Typowe pulapki**:
- Pomylenie kolejnosci w rekurencji (akcje przed vs po wywolaniu)
- Pominiecien kroku w petli (off-by-one)
- Bledne sledzenie warunkow if/else
**Przyklad**: 2014/1a (Korale), 2015/1.1 (strategie), 2023/1.1 (BST)

#### 2. projektowanie_algorytmu
**Co to jest**: Napisac algorytm/pseudokod rozwiazujacy problem. Czesto ograniczenia: tylko int, brak builtinow.
**Czestotliwosc**: 11/11 lat, 43 pkt lacznie — PEWNE na egzaminie!
**Jak podejsc**:
1. Zrozum dokladnie co algorytm ma robic (wejscie/wyjscie)
2. Pomysl o przypadkach brzegowych
3. Napisz pseudokod lub C++ (to co umiesz lepiej)
4. Przetestuj mentalnie na danych przykladowych
**Typowe pulapki**:
- Uzycie niedozwolonych operacji (np. float gdy dozwolone tylko int)
- Brak warunku stopu w petli/rekurencji
- Niepoprawna konwersja rekurencja <-> iteracja
**Ograniczenia**: Czesto "uzyj tylko zmiennych calkowitych", "nie uzywaj funkcji wbudowanych"
**Przyklad**: 2014/1c (KoraleBis), 2015/3.2 (Rozszerzony Euklides), 2024/3 (pseudokod)

#### 3. analiza_algorytmu
**Co to jest**: Okresl zlozonosc, udowodnij wlasciwosci, podaj min/max, znajdz kontrprzyklad.
**Czestotliwosc**: 10/11 lat, 37 pkt lacznie
**Jak podejsc**:
1. Zidentyfikuj petle zagniezdzenie — to daje zlozonosc
2. Dla dowodow: uzyj kontrprzykadow (udowodnij ze NIE dziala)
3. Dla min/max: pomysl o najgorszym/najlepszym przypadku
4. Sprawdz warunki brzegowe (n=0, n=1, dane posortowane)
**Typowe pulapki**:
- Mylenie O(n) z O(n^2) gdy petla wewnetrzna zalezy od i
- Zapomnienie o koszcie sortowania w zlozonosci
- Bledne kontrprzyklady (nie spelniajace warunkow zadania)
**Przyklad**: 2014/1b (ile koralikow), 2015/1.2 (kontrprzyklady), 2022/2 (ab-slowo)

#### 4. test_prawda_falsz
**Co to jest**: Ocen prawdziwosc 4 zdan (P/F). Tematy: algorytmy, SQL, sieci, grafika.
**Czestotliwosc**: 10/11 lat, 25 pkt lacznie
**Jak podejsc**:
1. Czytaj DOKLADNIE — jedno slowo moze zmienic odpowiedz
2. Jezeli nie jestes pewien, szukaj kontrprzykladu
3. Czesto 2P + 2F (ale nie zawsze!)
4. Tematy: SQL, systemy liczbowe, sieci, formaty plikow
**Typowe pulapki**:
- Pospiesz = bledy (kazdy punkt sie liczy)
- Mylenie "zawsze" z "czasami"
- Nieznajmosc terminologii (ADWARE, BMP vs JPG, LEFT JOIN)
**Przyklad**: 2014/3b-e (test mieszany), 2016/3 (DNS, rekurencja), 2019/3

#### 5. konwersja_systemow_liczbowych
**Co to jest**: Konwersje miedzy bazami (bin/oct/hex/dec), arytmetyka w roznych systemach.
**Czestotliwosc**: 9/11 lat, 12 pkt lacznie
**Jak podejsc**:
1. bin->hex: grupuj po 4 bity od prawej
2. bin->oct: grupuj po 3 bity od prawej
3. base k -> dec: schemat Hornera
4. dec -> base k: dzielenie z reszta
**Typowe pulapki**:
- Grupowanie bitow od LEWEJ strony zamiast prawej
- Bledy w arytmetyce (przeniesienie/pozyczka w systemie k)
- Zapomnienie o dopelnieniu zerami przy grupowaniu
**Przyklad**: 2014/3c (bin->hex), 2015/2.1 (mnozenie w syst. 4), 2025/5 (dodawanie bin)

#### 6. teoria_bezpieczenstwa
**Co to jest**: Szyfrowanie, protokoly sieciowe, bezpieczenstwo (quick wins za 1 pkt).
**Czestotliwosc**: 2/11 lat, 2 pkt lacznie — pojawia sie od 2023
**Jak podejsc**:
1. Znaj roznice: szyfrowanie symetryczne vs asymetryczne
2. Znaj podstawowe protokoly (HTTP, HTTPS, FTP, DHCP)
3. Znaj typy zagrozne (keylogger, phishing, malware)
**Przyklad**: 2023/4 (szyfrowanie asymetryczne), 2025/4 (keylogger)

---

### KATEGORIA: IMPLEMENTACJA

#### 7. przetwarzanie_cyfry_liczby
**Co to jest**: Analiza cyfr (mod/div), NWD, potegi, faktoryzacja, podzielnosc.
**Czestotliwosc**: 6/11 lat, 36 pkt lacznie
**Jak podejsc**:
1. Wzorzec: `while(n>0) { cyfra = n%10; n /= 10; }`
2. NWD: algorytm Euklidesa (while b!=0)
3. Pierwszosc: sprawdzaj dzielniki do sqrt(n)
4. Faktoryzacja: dziel przez kolejne liczby od 2
**Typowe pulapki**:
- Pierwsza cyfra wymaga petli (nie n%10 ale n po wielokrotnym /10)
- Overflow przy mnozeniu duzych liczb (uzyj long long)
- NWD(0, x) = x (nie 0!)
**Przyklad**: 2015/4.1-4.2 (cyfry, podzielnosc), 2019/4 (potegi 3, NWD), 2024/4 (l. pierwsze)

#### 8. przetwarzanie_napisy
**Co to jest**: Palindromy, szyfry, ASCII, manipulacja tekstem.
**Czestotliwosc**: 4/11 lat, 25 pkt lacznie
**Jak podejsc**:
1. Kody ASCII: A=65, a=97, 0=48, roznica a-A=32
2. Palindrom: porownaj s[i] z s[n-1-i]
3. Szyfr Cezara: (kod - 'A' + przesuniecie) % 26 + 'A'
4. Suma kodow: for(char c : s) sum += (int)c;
**Typowe pulapki**:
- Zawijanie cykliczne: ((x % n) + n) % n dla ujemnych
- Sortowanie leksykograficzne vs numeryczne
- Wielkosc liter (A vs a)
**Przyklad**: 2014/5 (napisy ASCII), 2016/6 (szyfr Cezara), 2021/4 (DOPISZ/USUN)

#### 9. przetwarzanie_zlozone
**Co to jest**: Wieloetapowy algorytm na danych z pliku — wymaga kilku krokow logiki.
**Czestotliwosc**: 4/11 lat, 24 pkt lacznie
**Jak podejsc**:
1. Rozbij problem na etapy (wczytaj -> przetworz -> wypisz)
2. Kazdy etap testuj osobno
3. Czesto wymaga map/set do grupowania
4. Uzyj struktur (struct) dla zlozonych danych
**Typowe pulapki**:
- Proba zrobienia wszystkiego w jednej petli
- Brak testowania na danych przykladowych
- Bledne parsowanie pliku
**Przyklad**: 2021/4 (operacje na napisach), 2022/4.3 (trojki czynnikow), 2025/2 (wzorce 2D)

#### 10. przetwarzanie_zliczanie
**Co to jest**: Zlicz elementy spelniajace warunek, filtruj dane z pliku.
**Czestotliwosc**: 5/11 lat, 17 pkt lacznie
**Jak podejsc**:
1. Wzorzec: `int count=0; for(auto x : data) if(warunek(x)) count++;`
2. Wczytaj dane z pliku do tablicy/vectora
3. Zastosuj warunek filtrowania
4. Wypisz wynik
**Typowe pulapki**:
- Off-by-one w warunkach (< vs <=)
- Bledy w czytaniu pliku (pominiety ostatni element)
- Liczenie od 0 vs od 1
**Przyklad**: 2014/5a (napisy dl. 6), 2023/3.1-3.2 (cyfry pi), 2024/3.2 (nieparzysty skrot)

#### 11. przetwarzanie_minmax
**Co to jest**: Znajdz min/max, posortuj, analizuj rozklad danych.
**Czestotliwosc**: 5/11 lat, 17 pkt lacznie
**Jak podejsc**:
1. Inicjalizuj min = INT_MAX, max = INT_MIN (lub pierwszym elementem)
2. Iteruj i aktualizuj
3. Zapamietaj INDEKS min/max jezeli potrzebny
4. Dla kilku min/max: uzyj vectora
**Typowe pulapki**:
- Zla inicjalizacja (min=0 gdy dane moga byc ujemne)
- Zapomnienie o zapisaniu indeksu
- Wielu elementow o tej samej wartosci min/max
**Przyklad**: 2015/4.3 (wiersz min/max), 2023/3.3 (sekwencje rosnaco-malejace), 2024/4.2 (max czynnik)

#### 12. przetwarzanie_sekwencje
**Co to jest**: Najdluzszy podciag, bloki, wzorce w ciagach.
**Czestotliwosc**: 3/11 lat, 13 pkt lacznie
**Jak podejsc**:
1. Uzyj zmiennych: current_len, max_len
2. Iteruj i sprawdzaj warunek kontynuacji sekwencji
3. Jezeli warunek zlamany: porownaj current z max, resetuj
4. Na koncu: jeszcze raz porownaj (sekwencja moze konczyc sie na ostatnim elemencie)
**Typowe pulapki**:
- Zapomnienie o ostatnim porownaniu po petli
- Bledna definicja "kontynuacji" sekwencji
- Puste dane / dane jednoelementowe
**Przyklad**: 2019/4.3 (najdluzszy ciag o NWD>1), 2023/3.3-3.4 (rosnaco-malejace)

#### 13. przetwarzanie_obrazy_2D
**Co to jest**: Piksele, siatki 2D, connected components (DFS/BFS).
**Czestotliwosc**: 2/11 lat, 11 pkt lacznie
**Jak podejsc**:
1. Wczytaj dane do tablicy 2D
2. DFS/BFS do przeszukiwania sasiadow
3. 4-sasiedztwo: gora, dol, lewo, prawo
4. Liczenie spojnych skladowych: iteruj po pikselach, DFS nieodwiedzonych
**Typowe pulapki**:
- Wyjscie poza granice tablicy (sprawdzaj x>=0, x<N, y>=0, y<M)
- Zapomnienie o oznaczeniu jako odwiedzone (visited)
- Stack overflow przy duzych obszarach (uzyj BFS zamiast rekurencyjnego DFS)
**Przyklad**: 2017/6 (obraz rastrowy, connected components), 2025/2.4 (wzorce 2D)

#### 14. obliczenia_geometryczne
**Co to jest**: Odleglosci, srodki odcinkow, pola, Monte Carlo.
**Czestotliwosc**: 1/11 lat, 4 pkt lacznie
**Jak podejsc**:
1. Odleglosc: sqrt((x2-x1)^2 + (y2-y1)^2)
2. Srodek: ((x1+x2)/2, (y1+y2)/2)
3. Pole trojkata: |x1(y2-y3) + x2(y3-y1) + x3(y1-y2)| / 2
**Typowe pulapki**:
- Dzielenie calkowite zamiast zmiennoprzecinkowego
- Precision issues z double
**Przyklad**: 2025/3 (Dron - NWD + srodek odcinka)

---

### KATEGORIA: ARKUSZ KALKULACYJNY

#### 15. arkusz_agregacja_warunkowa
**Co to jest**: SUMIF, COUNTIF, AVERAGEIF, SUMIFS — agregacja z warunkami.
**Czestotliwosc**: 9/11 lat, 38 pkt lacznie
**Jak podejsc**:
1. Zidentyfikuj ZAKRES danych, KRYTERIUM, i ZAKRES_SUMOWANIA
2. SUMIF(zakres_kryt, kryterium, zakres_sum)
3. Dla wielu warunkow: SUMIFS (warunki w parach zakres+kryterium)
4. Kopiowanie: uzyj $ do zablokowania odniesien
**Typowe pulapki**:
- Pomylenie SUMIF z SUMIFS (kolejnosc argumentow!)
- Brak $ w odniesieniach bezwzglednych
- Kryterium tekstowe bez cudzyslowow
**Przyklad**: 2014/4a (max przychod wieczorem), 2019/5 (pogoda), 2021/5 (wodociagi)

#### 16. arkusz_symulacja
**Co to jest**: Symulacje krokowe, formuly dynamiczne, prognozy.
**Czestotliwosc**: 9/11 lat, 37 pkt lacznie
**Jak podejsc**:
1. Zrozum model (wzor przyrostu, rabat progresywny, etc.)
2. Utworz kolumny pomocnicze dla kolejnych krokow
3. Pierwsza komorka = recznie, potem kopiuj formule w dol
4. Testuj na danych przykladowych
**Typowe pulapki**:
- Blad w pierwszym wierszu propaguje sie do wszystkich
- Zle odniesienia przy kopiowaniu (brak $)
- Zaokraglenia w symulacjach finansowych
**Przyklad**: 2015/5.3 (prognoza ludnosci 2025), 2017/4 (rabaty cukier), 2023/6 (konfitury)

#### 17. arkusz_wykres
**Co to jest**: Tworzenie wykresow: kolumnowy, kolowy, liniowy.
**Czestotliwosc**: 8/11 lat, 25 pkt lacznie
**Jak podejsc**:
1. Zaznacz dane RAZEM z etykietami
2. Wstaw wykres odpowiedniego typu
3. Dodaj tytul, etykiety osi, legende
4. Sprawdz czy dane sa poprawnie przypisane do serii
**Typowe pulapki**:
- Zly typ wykresu (kolowy dla wartosci ujemnych)
- Brak etykiet lub tytulu
- Zaznaczenie zlego zakresu danych
**Przyklad**: 2014/4d (slupkowy), 2015/5.1 (kolumnowy), 2025/6 (skumulowany)

#### 18. arkusz_agregacja_podstawowa
**Co to jest**: SUM, COUNT, AVERAGE, MAX/MIN — proste agregacje bez warunkow.
**Czestotliwosc**: 3/11 lat, 9 pkt lacznie
**Jak podejsc**:
1. SUM(zakres), AVERAGE(zakres), COUNT(zakres)
2. MAX(zakres), MIN(zakres)
3. Zakres moze byc kolumna lub wiersz
**Typowe pulapki**:
- COUNT liczy niepuste komorki (uzyj COUNTA dla tekstu)
- Puste komorki pomijane w AVERAGE
**Przyklad**: 2014/4b (laczna wartosc przychodow), 2015/5.1 (mieszkancy regionow)

#### 19. arkusz_transformacja
**Co to jest**: Grupowanie, pivoty, restrukturyzacja danych.
**Czestotliwosc**: 2/11 lat, 3 pkt lacznie
**Jak podejsc**:
1. Zrozum strukture wejsciowa i docelowa
2. Uzyj SUMIF/COUNTIF do przeorganizowania danych
3. Tabele przestawne (pivot tables) jezeli dostepne
**Przyklad**: 2017/4.1 (grupowanie cukru), 2025/6 (transformacja danych)

---

### KATEGORIA: SQL

#### 20. sql_group_by
**Co to jest**: GROUP BY z COUNT/SUM/AVG/MAX/MIN, czesto z HAVING.
**Czestotliwosc**: 8/11 lat, 36 pkt lacznie
**Jak podejsc**:
1. SELECT kolumna_grupujaca, funkcja_agregujaca(kolumna)
2. FROM tabela
3. GROUP BY kolumna_grupujaca
4. HAVING warunek_na_grupe (filtrowanie PO agregacji)
5. ORDER BY do posortowania wyniku
**Typowe pulapki**:
- Brak GROUP BY przy agregacji
- WHERE zamiast HAVING (WHERE = przed, HAVING = po agregacji)
- Brak wszystkich nieagregowanych kolumn w GROUP BY
**Przyklad**: 2014/6d (srednia chetnch), 2016/5.3 (liczba studentow), 2022/6 (czas pobytu)

#### 21. sql_podzapytania
**Co to jest**: Podzapytania zagniezdzone, NOT IN, EXISTS, IN.
**Czestotliwosc**: 7/11 lat, 25 pkt lacznie
**Jak podejsc**:
1. Napisz najpierw podzapytanie (wewnetrzne)
2. Przetestuj podzapytanie osobno
3. Uzyj IN/NOT IN/EXISTS w zapytaniu zewnetrznym
4. Alternatywa: LEFT JOIN + IS NULL (zamiast NOT IN)
**Typowe pulapki**:
- NULL w NOT IN (daje puste wyniki!)
- Podzapytanie zwracajace wiecej niz 1 kolumne
- Korelacja bledna (alias z zewnetrznego zapytania)
**Przyklad**: 2015/6.2 (miejsce bez GrandPrix), 2016/5.4-5.5 (NOT IN), 2024/8 (NOT IN)

#### 22. sql_join
**Co to jest**: Laczenie 2-3 tabel przez INNER JOIN lub LEFT JOIN.
**Czestotliwosc**: 8/11 lat, 21 pkt lacznie
**Jak podejsc**:
1. Zidentyfikuj tabele i klucze laczenia
2. INNER JOIN: tylko pasujace rekordy
3. LEFT JOIN: wszystkie z lewej + pasujace z prawej (NULL jezeli brak)
4. Lacz po kluczach: t1.id = t2.id_foreign
**Typowe pulapki**:
- INNER zamiast LEFT (tracisz rekordy bez dopasowania)
- Bledny warunek ON (zle klucze)
- Duplikaty przy JOIN many-to-many
**Przyklad**: 2014/6c (przedszkola), 2015/6.3 (mistrzowie F1), 2019/6 (perfumy 3 tabele)

#### 23. sql_select_where
**Co to jest**: Prosty SELECT z WHERE na 1 tabeli — najlatwiejszy typ SQL.
**Czestotliwosc**: 4/11 lat, 10 pkt lacznie
**Jak podejsc**:
1. SELECT kolumny FROM tabela WHERE warunek
2. Operatory: =, <>, <, >, LIKE, BETWEEN, IN
3. LIKE: % = dowolny ciag, _ = jeden znak
4. ORDER BY do posortowania, LIMIT do ograniczenia
**Typowe pulapki**:
- Brak cudzyslowow dla tekstu w WHERE
- LIKE case-sensitive (zalezy od bazy)
- Zapomnienie o ORDER BY gdy wymagane
**Przyklad**: 2014/6a (dziewczynki z Pragi), 2015/6.1 (najwczesniejsze GP)

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

### 11. **Programowanie Dynamiczne** (40%+ egzaminów - rośnie!)
**Dlaczego**: Od 2023 pojawia się coraz częściej (Plansza 2024!)
**Co umieć**:
```cpp
// DP 1D - Fibonacci (wzorzec podstawowy)
int dp[MAX];
dp[0] = 0; dp[1] = 1;
for (int i = 2; i <= n; i++)
    dp[i] = dp[i-1] + dp[i-2];

// DP 2D - ścieżki na planszy (matura 2024!)
bool dp[N][M];
dp[0][0] = true;
for (int i = 0; i < N; i++)
    for (int j = 0; j < M; j++) {
        if (plansza[i][j] == blokada) { dp[i][j] = false; continue; }
        if (i > 0 && dp[i-1][j]) dp[i][j] = true;  // z góry
        if (j > 0 && dp[i][j-1]) dp[i][j] = true;  // z lewej
    }
// Odpowiedź: dp[N-1][M-1]

// DP - najdłuższy rosnący podciąg (LIS)
int lis[MAX];
fill(lis, lis+n, 1);
for (int i = 1; i < n; i++)
    for (int j = 0; j < i; j++)
        if (arr[j] < arr[i])
            lis[i] = max(lis[i], lis[j] + 1);
int ans = *max_element(lis, lis+n);

// DP - problem plecakowy 0/1
int dp[MAX_W+1] = {0};
for (int i = 0; i < n; i++)
    for (int w = W; w >= weight[i]; w--)
        dp[w] = max(dp[w], dp[w - weight[i]] + value[i]);
```
**Typowe zastosowania**: Ścieżki na planszy, problem plecakowy, ciągi

---

### 12. **Algorytmy Zachłanne** (30%+ egzaminów)
**Dlaczego**: Eleganckie i szybkie rozwiązania
**Co umieć**:
```cpp
// Problem doboru aktywności (matura 2015!)
// Posortuj aktywności po czasie zakończenia
sort(activities.begin(), activities.end(),
    [](Activity a, Activity b) { return a.end < b.end; });

int count = 1;
int last_end = activities[0].end;
for (int i = 1; i < n; i++) {
    if (activities[i].start >= last_end) {
        count++;
        last_end = activities[i].end;
    }
}

// Problem wydawania reszty (zachłanny - monety od największej)
vector<int> coins = {200, 100, 50, 20, 10, 5, 2, 1};
int reszta = kwota;
for (int coin : coins) {
    int ile = reszta / coin;
    reszta -= ile * coin;
    // ile monet o nominale coin
}
```
**Typowe zastosowania**: Scheduling, reszta, pokrycie zbioru

---

### 13. **Systemy Liczbowe** (50%+ egzaminów)
**Dlaczego**: Stały element matury - pytania P/F i implementacja
**Co umieć**:
```cpp
// Konwersja base 10 → base k
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

// Konwersja base k → base 10 (Schemat Hornera)
int fromBase(string s, int k) {
    int result = 0;
    for (char c : s) {
        int digit;
        if (c >= '0' && c <= '9') digit = c - '0';
        else digit = c - 'A' + 10;
        result = result * k + digit;  // Horner!
    }
    return result;
}

// Dodawanie w systemie k
string addBase(string a, string b, int k) {
    string result = "";
    int carry = 0;
    int i = a.size()-1, j = b.size()-1;
    while (i >= 0 || j >= 0 || carry) {
        int sum = carry;
        if (i >= 0) sum += a[i--] - '0';
        if (j >= 0) sum += b[j--] - '0';
        result = char('0' + sum % k) + result;
        carry = sum / k;
    }
    return result;
}
```
**Typowe zastosowania**: Konwersje, dodawanie/odejmowanie w dowolnym systemie, bin↔hex

---

### 14. **BFS/DFS (Przeszukiwanie Grafów)** (30%+ egzaminów)
**Dlaczego**: Pojawia się w zadaniach z obrazami, planszami, grafami
**Co umieć**:
```cpp
// BFS - przeszukiwanie wszerz (najkrótsza ścieżka, connected components)
void bfs(int start, vector<vector<int>>& adj, vector<bool>& visited) {
    queue<int> q;
    q.push(start);
    visited[start] = true;
    while (!q.empty()) {
        int v = q.front(); q.pop();
        for (int u : adj[v]) {
            if (!visited[u]) {
                visited[u] = true;
                q.push(u);
            }
        }
    }
}

// DFS - przeszukiwanie w głąb (connected components, flood fill)
void dfs(int x, int y, vector<vector<int>>& grid, vector<vector<bool>>& vis) {
    if (x < 0 || x >= N || y < 0 || y >= M) return;
    if (vis[x][y] || grid[x][y] == 0) return;
    vis[x][y] = true;
    dfs(x+1, y, grid, vis);  // dół
    dfs(x-1, y, grid, vis);  // góra
    dfs(x, y+1, grid, vis);  // prawo
    dfs(x, y-1, grid, vis);  // lewo
}

// Liczenie connected components (matura 2017!)
int count_components(vector<vector<int>>& grid) {
    vector<vector<bool>> vis(N, vector<bool>(M, false));
    int count = 0;
    for (int i = 0; i < N; i++)
        for (int j = 0; j < M; j++)
            if (!vis[i][j] && grid[i][j] == 1) {
                dfs(i, j, grid, vis);
                count++;
            }
    return count;
}
```
**Typowe zastosowania**: Connected components (2017), flood fill, najkrótsza ścieżka, drzewa BST (2023)

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

**Dlaczego zaczynac od Quick Wins?**

1. **Psychologia egzaminu** — Zdobycie 5-10 pkt w pierwszych 15 min buduje pewnosc siebie i redukuje stres. Unikasz sytuacji, gdzie utkniesz na trudnym zadaniu na starcie i tracisz czas + motywacje.

2. **Gwarancja punktow** — Zadania P/F i proste sledzenie algorytmu maja najwyzszy stosunek punktow do czasu. Nawet jesli zabraknie czasu na koncu, te punkty juz masz zabezpieczone.

3. **Dane z analizy 11 lat matur potwierdzaja to**:
   - `sledzenie_algorytmu` — 45 pkt lacznie, 11/11 lat, a pierwsze podzadania sa zwykle proste
   - `test_prawda_falsz` — 25 pkt lacznie, 10/11 lat, wymaga tylko wiedzy (zero programowania)
   - `konwersja_systemow` — 12 pkt, 9/11 lat, czysto mechaniczne obliczenia

4. **Rozpoznanie arkusza** — Przegladajac caly arkusz w 2 min na starcie, wiesz co Cie czeka i mozesz zaplanowac czas na reszte.

5. **Matematyka ryzyka** — Jesli zaczniesz od trudnego zadania za 5 pkt i poswiecisz 40 min bez rozwiazania, straciles czas, ktory mogl dac Ci 10 pkt latwych. Quick wins to ~1-2 min/punkt, trudne zadania to ~5-8 min/punkt.

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

### Podział czasu (Formuła 2023 - AKTUALNA!):
**1 arkusz, 210 minut, 50 punktów, 7 zadań**

Typowa struktura (na podstawie analizy 2023):

| Zad | Typ | Pkt | Szacowany czas | Priorytet |
|---|---|---|---|---|
| 4-5 | Quick-win P/F + systemy liczbowe | 2 | 5-10 min | ETAP 1 |
| 1.1 | Śledzenie algorytmu (teoria) | 2 | 10 min | ETAP 1 |
| 7.1 | Prosty SQL (SELECT + JOIN) | 1 | 5 min | ETAP 1 |
| 1.2-1.3 | Analiza algorytmu (teoria) | 5 | 15-20 min | ETAP 2 |
| 2.1 | Algorytm (pseudokod) | 3 | 15 min | ETAP 2 |
| 2.2-2.3 | Programowanie: pliki (proste) | 4 | 15-20 min | ETAP 2 |
| 3.1-3.2 | Programowanie: pliki (srednie) | 5 | 20-25 min | ETAP 2 |
| 6.1-6.2 | Arkusz: zestawienie + wykres | 4 | 15-20 min | ETAP 2 |
| 7.2-7.3 | SQL sredni (JOIN + GROUP BY) | 4 | 15-20 min | ETAP 2 |
| 2.4-2.5 | XOR + programowanie binarne | 4 | 20 min | ETAP 3 |
| 3.3-3.4 | Programowanie: ciagi rosnaco-malejace | 5 | 25-30 min | ETAP 3 |
| 6.3-6.4 | Arkusz/Prog: symulacja produkcji | 6 | 25-30 min | ETAP 3 |
| 7.4-7.5 | SQL trudny (kategorie wiekowe, SUM) | 5 | 15-20 min | ETAP 3 |

**Sugerowana kolejnosc**:
1. **0-15 min**: Zad 4, 5, 1.1, 7.1 → ~5 pkt
2. **15-90 min**: Zad 1.2-1.3, 2.1-2.3, 3.1-3.2, 6.1-6.2, 7.2-7.3 → ~25 pkt
3. **90-200 min**: Zad 2.4-2.5, 3.3-3.4, 6.3-6.4, 7.4-7.5 → ~20 pkt
4. **200-210 min**: Sprawdzenie, poprawki → bufor

**TOTAL**: ~50 pkt w 210 min

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
