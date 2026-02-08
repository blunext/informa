# Wzorce i Pułapki - Matura 2014

## 🎯 Kluczowe Wzorce Kodu (2014)

### 1. Rekurencja → Iteracja (Zadanie 1)

**Wzorzec rekurencyjny:**
```
Korale(n):
  if n = 1: dodaj czarny, return
  if n parzyste: Korale(n/2), dodaj biały, return
  if n nieparzyste: Korale((n-1)/2), dodaj czarny, return
```

**Konwersja na iterację:**
```cpp
// Kluczowa obserwacja: algorytm to reprezentacja binarna!
void KoraleBis(int n) {
    while (n > 0) {
        if (n % 2 == 0) {
            nawlecz_bialy_na_lewy_koniec();
        } else {
            nawlecz_czarny_na_lewy_koniec();
        }
        n = n / 2;
    }
}
```

**Insight:** Liczba koralików = ⌊log₂(n)⌋ + 1 = liczba bitów w reprezentacji binarnej

---

### 2. Bisekcja - Znajdowanie Miejsca Zerowego (Zadanie 2)

```cpp
double bisekcja(double a, double b, double d, function<double(double)> f) {
    double x = (a + b) / 2.0;

    while (b - a >= d) {
        if (f(a) * f(x) < 0) {
            b = x;  // zero w lewej połowie
        } else {
            a = x;  // zero w prawej połowie
        }
        x = (a + b) / 2.0;
    }

    return x;
}
```

**Wzór na liczbę kroków:** aby przedział <0, L> zmniejszyć do długości < d:
- Po k krokach: długość = L / 2^k
- Potrzebne k: L / 2^k < d  →  k > log₂(L/d)
- Dla L=2, d=0.1: k > log₂(20) ≈ 4.32 → **k = 5** (ale odpowiedź to 6, bo liczymy od kroku 1!)

---

### 3. Śledzenie Algorytmu z Akumulacją (Zadanie 3)

```cpp
// Algorytm oblicza: 1! + 2! + 3! + ... + n!
int suma_silni(int n) {
    if (n == 1) return 1;

    int suma = 1 + n;  // start: 1 + n
    int i = n - 1;

    while (i > 1) {
        suma = 1 + i * suma;  // wzór rekurencyjny
        i = i - 1;
    }

    return suma;
}

// Przykład dla n=4:
// Start: suma = 1 + 4 = 5, i = 3
// Krok 1: suma = 1 + 3*5 = 16, i = 2
// Krok 2: suma = 1 + 2*16 = 33, i = 1
// STOP (i = 1)
// Wynik: 33 = 1! + 2! + 3! + 4!
```

**Wzór matematyczny:** 1! + 2! + 3! + ... + n!

---

### 4. Konwersje Systemów Liczbowych (Zadanie 3c)

```cpp
// Binarny → Hex (grupowanie po 4 bity)
// 101011111100 → 1010 1111 1100 → A F C

string bin_to_hex(string bin) {
    // Uzupełnij zerami z lewej do wielokrotności 4
    while (bin.length() % 4 != 0) {
        bin = "0" + bin;
    }

    string hex = "";
    for (int i = 0; i < bin.length(); i += 4) {
        string nibble = bin.substr(i, 4);
        int val = stoi(nibble, nullptr, 2);

        if (val < 10) hex += to_string(val);
        else hex += char('A' + val - 10);
    }

    return hex;
}
```

---

### 5. Obsługa Plików i Przetwarzanie Stringów (Zadanie 5)

```cpp
#include <fstream>
#include <string>
#include <map>
#include <vector>

// 5a: Zlicz napisy o długości 6
int zlicz_dlugosc_6(string filename) {
    ifstream file(filename);
    string napis;
    int count = 0;

    while (file >> napis) {
        if (napis.length() == 6) {
            count++;
        }
    }

    return count;
}

// 5b: Grupowanie po sumie kodów ASCII
map<int, vector<string>> grupuj_po_sumie_ascii(string filename) {
    ifstream file(filename);
    string napis;
    map<int, vector<string>> grupy;

    while (file >> napis) {
        int suma = 0;
        for (char c : napis) {
            suma += (int)c;
        }
        grupy[suma].push_back(napis);
    }

    // Sortuj każdą grupę
    for (auto& [key, vec] : grupy) {
        sort(vec.begin(), vec.end());
    }

    return grupy;
}

// 5c: Maksymalna różnica kodów znaków
vector<string> max_roznica_kodow(string filename) {
    ifstream file(filename);
    string napis;
    int max_roznica = 0;
    vector<string> wynik;

    while (file >> napis) {
        int min_kod = 256, max_kod = 0;
        for (char c : napis) {
            min_kod = min(min_kod, (int)c);
            max_kod = max(max_kod, (int)c);
        }

        int roznica = max_kod - min_kod;

        if (roznica > max_roznica) {
            max_roznica = roznica;
            wynik.clear();
            wynik.push_back(napis);
        } else if (roznica == max_roznica) {
            wynik.push_back(napis);
        }
    }

    return wynik;
}
```

---

### 6. SQL - Zapytania z JOIN i Agregacją (Zadanie 6)

```sql
-- 6a: Nazwiska dziewczynek z dzielnicy Praga
SELECT nazwisko
FROM dzieci
WHERE plec = 'K' AND dzielnica = 'Praga'
ORDER BY nazwisko;

-- 6b: Dzieci z imieniem/nazwiskiem na 'A', rok 2010
SELECT imie, nazwisko
FROM dzieci
WHERE (imie LIKE 'A%' OR nazwisko LIKE 'A%')
  AND rok_urodzenia = 2010
ORDER BY nazwisko, imie;

-- 6c: Przedszkole z największą liczbą chętnych (1. preferencja)
SELECT p.nazwa, COUNT(*) as liczba_chetnych
FROM preferencje pr
JOIN przedszkola p ON pr.id_przedszkola = p.id
WHERE pr.preferencja = 1
GROUP BY p.id, p.nazwa
ORDER BY liczba_chetnych DESC
LIMIT 1;

-- 6d: 3 przedszkola o najmniejszej średniej chętnych na miejsce
SELECT
    p.nazwa,
    p.miejsca,
    COUNT(pr.id_dziecka) as liczba_chetnych,
    ROUND(COUNT(pr.id_dziecka) * 1.0 / p.miejsca, 2) as srednia
FROM przedszkola p
LEFT JOIN preferencje pr ON p.id = pr.id_przedszkola
GROUP BY p.id, p.nazwa, p.miejsca
ORDER BY srednia ASC
LIMIT 3;
```

---

## ⚠️ Typowe Pułapki (2014)

### Pułapka 1: Rekurencja - Kolejność Dodawania
❌ **Błąd:** Myślenie, że rekurencja dodaje elementy w kolejności wywołania
✅ **Poprawnie:** Rekurencja dodaje elementy **od końca** (nawlekanie PRZED zakończeniem)

### Pułapka 2: Bisekcja - Warunek Pętli
❌ **Błąd:** `while (b - a > d)` - pomija przypadek b-a = d
✅ **Poprawnie:** `while (b - a >= d)` - zatrzymuje się gdy przedział < d

### Pułapka 3: Zaokrąglenia Float
❌ **Błąd:** Porównywanie float/double z `==`
✅ **Poprawnie:** Używać epsilon: `abs(a - b) < 1e-9`

### Pułapka 4: Kody ASCII - Suma vs Max-Min
❌ **Błąd:** Mylenie sumy kodów ASCII z różnicą max-min
✅ **Poprawnie:** Zad 5b to SUMA kodów, Zad 5c to MAX-MIN

### Pułapka 5: SQL - LEFT JOIN vs INNER JOIN
❌ **Błąd:** INNER JOIN pomija przedszkola bez chętnych
✅ **Poprawnie:** LEFT JOIN zachowuje wszystkie przedszkola

### Pułapka 6: Grupowanie w SQL
❌ **Błąd:** Zapominanie o `GROUP BY` przy COUNT/SUM/AVG
✅ **Poprawnie:** `GROUP BY p.id, p.nazwa` dla każdego nieagregowanego pola w SELECT

### Pułapka 7: Konwersja Binarna → Hex
❌ **Błąd:** Grupowanie od lewej zamiast uzupełnienia zerami
✅ **Poprawnie:** 101011111100 → **0**1010 1111 1100 → A F C

---

## 📊 Strategia Punktowa (2014)

### Quick Wins (6 pkt, ~10 min):
- **3c, 3d, 3e** (3 pkt) - pytania P/F
- **2a** (1 pkt) - uzupełnienie tabeli bisekcji
- **6a** (2 pkt) - proste SQL SELECT WHERE

### Średnie (28 pkt, ~90 min):
- **1a, 1b** (5 pkt) - śledzenie rekurencji + wzór
- **2b** (2 pkt) - obliczenie liczby kroków
- **3a, 3b** (3 pkt) - śledzenie algorytmu silni
- **4a, 4b, 4d** (6 pkt) - Excel: MAX, SUM, wykres
- **5a** (4 pkt) - zliczanie napisów
- **6b, 6c** (5 pkt) - SQL średniej trudności

### Trudne (16 pkt, ~110 min):
- **1c** (3 pkt) - projekt algorytmu iteracyjnego
- **2c** (3 pkt) - implementacja bisekcji
- **4c** (3 pkt) - suma skumulowana w Excel
- **5b, 5c** (6 pkt) - grupowanie stringów
- **6d** (4 pkt) - SQL z JOIN + agregacją

**Optymalna kolejność:** 3c,3d,3e,2a,6a → 1a,1b,2b,3a,3b,4a,4b,4d,5a,6b,6c → 1c,2c,4c,5b,5c,6d

---

## 🔑 Kluczowe Wzory Matematyczne

1. **Liczba koralików:** ⌊log₂(n)⌋ + 1
2. **Silnia sumy:** 1! + 2! + 3! + ... + n!
3. **NWD:** NWD(1310, 524) = 262 (algorytm Euklidesa)
4. **Kombinatoryka:** Uściski n osób = n(n-1)/2 = C(n, 2)
5. **Bisekcja kroki:** k > log₂(L/d)

---

## 💡 Lekcje na Przyszłość

1. **Rekurencja ≈ Reprezentacja binarna** - wiele zadań rekurencyjnych to przetwarzanie bitów
2. **Bisekcja = Binary search** - znajomy wzorzec, tylko dla funkcji ciągłych
3. **SQL agregacja** - zawsze GROUP BY wszystkie nieagregowane pola
4. **Excel** - używać nazw zakresów dla czytelności
5. **Stringi** - kody ASCII to int, można sumować i porównywać

---

**Czas na pełne rozwiązanie:** ~180-210 minut
**Rekomendowane tempo:** Quick wins → Średnie → Trudne
**Cel:** 45-50/50 punktów 🎯
