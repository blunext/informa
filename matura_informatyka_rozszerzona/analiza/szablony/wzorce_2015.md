# Wzorce i Pułapki - Matura 2015

## Kluczowe Wzorce Kodu (2015)

### 1. Activity Selection Problem - Problem Telewidza (Zadanie 1)

**Wzorzec zachłanny - Strategia D (OPTYMALNA):**
```cpp
// Strategia: wybierz film kończący się najwcześniej
struct Film {
    string nazwa;
    int start, koniec;
};

vector<Film> wybierz_filmy_optymalnie(vector<Film> filmy) {
    // Sortuj według czasu zakończenia (rosnąco)
    sort(filmy.begin(), filmy.end(),
         [](Film a, Film b) { return a.koniec < b.koniec; });

    vector<Film> wybrane;
    int ostatni_koniec = -1;

    for (Film f : filmy) {
        // Jeśli film nie koliduje z ostatnio wybranym
        if (f.start >= ostatni_koniec) {
            wybrane.push_back(f);
            ostatni_koniec = f.koniec;
        }
    }

    return wybrane;
}
```

**Insight:** Earliest Finish Time jest jedyną optymalną strategią zachłanną dla tego problemu.

**Kontrprzykłady dla nieptymalnych strategii:**

Strategia A (najdłuższy film):
```
TV1: film1(10:00-12:00, 2h), film2(12:00-14:00, 2h)
TV2: film3(10:00-11:00, 1h), film4(11:00-12:00, 1h)
Wynik A: {film1, film2} = 2 filmy
Optymalnie: {film3, film4, film2} = 3 filmy
```

---

### 2. Rozszerzony Algorytm Euklidesa (Zadanie 3)

**Wzór matematyczny:** NWD(a, b) = a · x + b · y

```cpp
// Rekurencyjna wersja rozszerzonego algorytmu Euklidesa
pair<int, int> rozszerzony_euklides(int a, int b) {
    // Warunek bazowy
    if (b == 0) {
        return {1, 0};  // NWD(a, 0) = a = a·1 + 0·0
    }

    // Rekurencyjne wywołanie
    int r = a % b;
    auto [x_prim, y_prim] = rozszerzony_euklides(b, r);

    // Obliczenie x i y na podstawie x' i y'
    int x = y_prim;
    int y = x_prim - (a / b) * y_prim;

    return {x, y};
}

// Przykład użycia:
// auto [x, y] = rozszerzony_euklides(188, 12);
// NWD(188, 12) = 4 = 188·(-1) + 12·16
```

**Kluczowa zależność:**
```
Dla r = a mod b i NWD(b, r) = b·x' + r·y'
Para (x, y) to:
    x = y'
    y = x' - (a div b) · y'
```

---

### 3. Operacje na Cyfrach Liczby (Zadanie 4)

```cpp
#include <fstream>
#include <string>

// 4.1: Sprawdź czy pierwsza cyfra = ostatnia cyfra
bool pierwsza_rowna_ostatniej(int n) {
    int ostatnia = n % 10;

    // Znajdź pierwszą cyfrę
    int pierwsza = n;
    while (pierwsza >= 10) {
        pierwsza /= 10;
    }

    return pierwsza == ostatnia;
}

// Zliczanie liczb spełniających warunek
int zlicz_z_warunku(string filename) {
    ifstream file(filename);
    int liczba, count = 0;

    while (file >> liczba) {
        if (pierwsza_rowna_ostatniej(liczba)) {
            count++;
        }
    }

    return count;
}

// 4.2: Podzielność przez 8 (uwaga: zadanie 2015 dotyczy liczb BINARNYCH)
bool podzielna_przez_8(int n) {
    // W systemie dwojkowym: liczba dzieli sie przez 8 <=> trzy ostatnie BITY = "000".
    // Jesli liczbe trzymamy juz jako int, wystarczy: n % 8 == 0.
    // Jesli wczytujesz bity jako string s: return s.size() >= 3 && s.substr(s.size()-3) == "000";
    return n % 8 == 0;
}

// 4.3: Min i Max z numerem wiersza
pair<int, int> znajdz_min_max_wiersze(string filename) {
    ifstream file(filename);
    int liczba;
    int min_val = INT_MAX, max_val = INT_MIN;
    int min_wiersz = 0, max_wiersz = 0;
    int wiersz = 1;

    while (file >> liczba) {
        if (liczba < min_val) {
            min_val = liczba;
            min_wiersz = wiersz;
        }
        if (liczba > max_val) {
            max_val = liczba;
            max_wiersz = wiersz;
        }
        wiersz++;
    }

    return {min_wiersz, max_wiersz};
}
```

---

### 4. Arkusz Kalkulacyjny - Demografia (Zadanie 5)

**5.1: Sumowanie po regionach**
```excel
// Komórka dla regionu A (zakładając dane w kolumnach A=region, B=ludność)
=SUMIF(A:A,"A",B:B)

// Dla regionu B
=SUMIF(A:A,"B",B:B)
```

**5.2: Zliczanie województw**
```excel
// Liczba województw w całym kraju
=COUNTA(A2:A20)

// Liczba województw w regionie A
=COUNTIF(A:A,"A")
```

**5.3: Prognoza na 2025**
```excel
// Zakładając wzrost 1% rocznie przez 12 lat (2013-2025)
=B2 * (1.01)^12

// Najliczniejsze województwo
=INDEX(A:A, MATCH(MAX(D:D), D:D, 0))

// Przeludnienie (ludność > powierzchnia)
=COUNTIF(E:E, ">0")
// gdzie E to kolumna z formułą =ludność - powierzchnia
```

---

### 5. SQL - Zapytania z JOIN i Agregacją (Zadanie 6)

```sql
-- 6.1: Najwcześniejszy wyścig w historii
SELECT g.Nazwa, g.Sezon
FROM GrandPrix g
ORDER BY g.Data ASC
LIMIT 1;

-- 6.2: Miejsca bez wyścigów
SELECT m.Nazwa
FROM Miejsca m
LEFT JOIN GrandPrix g ON m.ID = g.ID_Miejsca
WHERE g.ID IS NULL;

-- 6.3: Mistrzowie w latach 2000, 2006, 2012
SELECT
    z.Imie,
    z.Nazwisko,
    w.Sezon,
    SUM(wg.Punkty) as Suma_Punktow
FROM Zawodnicy z
JOIN Wyniki_GrandPrix wg ON z.ID = wg.ID_Zawodnika
JOIN GrandPrix g ON wg.ID_GrandPrix = g.ID
JOIN Sezony w ON g.Sezon = w.Sezon
WHERE w.Sezon IN (2000, 2006, 2012)
GROUP BY z.ID, w.Sezon
HAVING SUM(wg.Punkty) = (
    SELECT MAX(suma) FROM (
        SELECT SUM(punkty) as suma
        FROM Wyniki_GrandPrix
        WHERE sezon = w.Sezon
        GROUP BY ID_Zawodnika
    )
)
ORDER BY w.Sezon;

-- 6.4: Liczba zawodników z poszczególnych krajów, którzy zdobyli punkty
-- Uwaga: zadanie wymaga tylko tych zawodnikow, ktorzy ZDOBYLI PUNKTY
-- w jakimkolwiek wyscigu — dlatego potrzebny JOIN z Wyniki_GrandPrix.
SELECT
    z.Kraj,
    COUNT(DISTINCT z.ID) as Liczba_Zawodnikow
FROM Zawodnicy z
JOIN Wyniki_GrandPrix wg ON z.ID = wg.ID_Zawodnika
WHERE wg.Punkty > 0
GROUP BY z.Kraj
ORDER BY z.Kraj;
```

---

## Typowe Pułapki (2015)

### Pułapka 1: Activity Selection - Wybór Strategii
**Błąd:** Myślenie, że najkrótsze filmy (strategia B) dają optymalne rozwiązanie
**Poprawnie:** Tylko earliest finish time (strategia D) jest optymalna

### Pułapka 2: Kontrprzykłady
**Błąd:** Trudność z konstrukcją kontrprzykładu pokazującego nieptymalność
**Poprawnie:** Dobierz czasy tak, żeby zachłanna strategia wybrała długi film blokujący dwa krótkie

### Pułapka 3: Rozszerzony Euklides - Kolejność Powrotu
**Błąd:** Mylenie kolejności w obliczeniach x i y podczas powrotu z rekurencji
**Poprawnie:** x = y', y = x' - (a div b) · y'

### Pułapka 4: Cyfry Liczby
**Błąd:** Użycie toString() zamiast operacji arytmetycznych
**Poprawnie:** Pierwsza cyfra: while(n >= 10) n /= 10; Ostatnia: n % 10

### Pułapka 5: Podzielność przez 8
**Błąd:** Sprawdzanie tylko ostatniej cyfry (to działa dla 2, nie dla 8)
**Poprawnie:** n % 8 == 0 (trzeba sprawdzić całą liczbę)

### Pułapka 6: Numeracja Wierszy w Pliku
**Błąd:** Liczenie od 0 zamiast od 1
**Poprawnie:** Pierwszy wiersz to wiersz 1 (nie 0)

### Pułapka 7: Plik Testowy vs Pełny
**Błąd:** Przetestowanie tylko na 250 przykładowych wierszach
**Poprawnie:** Pełny plik ma 1000 wierszy - użyj właściwego pliku!

### Pułapka 8: SUMIF w Excel
**Błąd:** Użycie SUM zamiast SUMIF dla warunkowego sumowania
**Poprawnie:** =SUMIF(zakres_kryterium, kryterium, zakres_do_zsumowania)

### Pułapka 9: SQL LEFT JOIN
**Błąd:** Użycie INNER JOIN gdy szukamy rekordów bez dopasowania
**Poprawnie:** LEFT JOIN + WHERE prawaTABELA.ID IS NULL

### Pułapka 10: GROUP BY w SQL
**Błąd:** Zapomnienie GROUP BY przy użyciu SUM/COUNT
**Poprawnie:** Wszystkie nieagregowane kolumny muszą być w GROUP BY

---

## Strategia Punktowa (2015)

### Quick Wins (5 pkt, ~10 min):
- **2.1-2.5** (5 pkt) - pytania P/F (testy wiedzy)

### Średnie (25 pkt, ~90 min):
- **1.1** (2 pkt) - śledzenie algorytmu dla strategii B, C, D
- **3.1** (2 pkt) - uzupełnienie tabeli Euklidesa
- **4.1, 4.2** (6 pkt) - podstawowe operacje na liczbach
- **5.1, 5.2** (7 pkt) - Excel: sumowanie i zliczanie
- **6.1, 6.2** (4 pkt) - proste SQL
- **6.3** (3 pkt) - SQL średniej trudności

### Trudne (20 pkt, ~80 min):
- **1.2** (3 pkt) - konstrukcja kontrprzykładów
- **3.2** (3 pkt) - uzupełnienie rekurencji
- **4.3** (6 pkt) - min/max z numerem wiersza
- **5.3** (6 pkt) - prognozy i złożone obliczenia
- **6.4** (3 pkt) - SQL z grupowaniem

**Optymalna kolejność:** Testy 2.1-2.5 → 1.1, 3.1, 4.1, 4.2, 5.1, 5.2, 6.1, 6.2 → 1.2, 3.2, 4.3, 5.3, 6.3, 6.4

---

## Kluczowe Wzory Matematyczne

1. **NWD rozszerzony:** NWD(a, b) = a · x + b · y
2. **Rekurencyjna zależność:** x = y', y = x' - (a div b) · y'
3. **Pierwsza cyfra:** while (n >= 10) n /= 10
4. **Ostatnia cyfra:** n % 10
5. **Podzielność:** n % k == 0
6. **Wzrost procentowy:** wartość · (1 + stopa)^okres

---

## Lekcje na Przyszłość

1. **Greedy ≠ zawsze optymalny** - tylko niektóre strategie zachłanne są optymalne
2. **Earliest Finish Time** - kluczowa strategia w interval scheduling
3. **Rozszerzony Euklides** - rekurencja z powrotem wartości pary
4. **Operacje na cyfrach** - mod 10 i div 10 są szybsze niż konwersja na string
5. **Excel SUMIF/COUNTIF** - warunkowe operacje są podstawą analizy danych
6. **SQL LEFT JOIN** - do znajdowania rekordów bez dopasowania
7. **Numeracja od 1** - w zadaniach maturalnych wiersze numerujemy od 1

---

**Czas na pełne rozwiązanie:** ~180 minut
**Rekomendowane tempo:** Quick wins (10 min) → Średnie (90 min) → Trudne (80 min)
**Cel:** 45-50/50 punktów
