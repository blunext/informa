# Checklist: Podczas Egzaminu

> Strategia czasowa na 210 minut. Kolejnosc zadan. Kontrola czasu.

---

## Etap 0: ROZPOZNANIE (10 min) [0:00 - 0:10]

- [ ] Przeczytaj WSZYSTKIE zadania (tytuly + tresc)
- [ ] Przy kazdym zapisz: L (latwe), S (srednie), T (trudne)
- [ ] Zidentyfikuj typy: SQL / arkusz / programowanie / teoria / quick win
- [ ] Zaplanuj kolejnosc (zaczynaj od L, koniec na T)

**Nowa formula (2023+):** Zadania sa przemieszane — nie zakladaj kolejnosci!

---

## Etap 1: QUICK WINS (10 min) [0:10 - 0:20] ~3-5 pkt

- [ ] Pytania P/F — przeczytaj uwazenie kazde zdanie
- [ ] Krotkie pytania za 1-2 pkt (bezpieczenstwo, protokoly, systemy liczbowe)
- [ ] Proste sledzenie algorytmu (jezeli krotki, <5 krokow)

**Zasada:** Jesli nie jestes pewny w P/F — nie zgaduj, przejdz dalej i wroc.

---

## Etap 2: SQL (30-40 min) [0:20 - 1:00] ~8-10 pkt

Kolejnosc podzadan:
1. [ ] Proste SELECT WHERE
2. [ ] GROUP BY + agregacje (COUNT, SUM, AVG)
3. [ ] JOIN (INNER, potem LEFT)
4. [ ] Podzapytania (NOT IN, EXISTS)

**Przed kazdym zapytaniem:**
- [ ] Przeczytaj schemat bazy danych (jakie tabele, jakie kolumny, klucze)
- [ ] Sprawdz czy potrzebujesz JOIN (dane z >1 tabeli?)
- [ ] Sprawdz czy potrzebujesz GROUP BY (agregacja?)
- [ ] Sprawdz czy WHERE czy HAVING (filtrowanie przed vs po grupowaniu)
- [ ] Uruchom zapytanie — czy wynik ma sens?

---

## Etap 3: ARKUSZ KALKULACYJNY (40-50 min) [1:00 - 1:50] ~10 pkt

Kolejnosc podzadan:
1. [ ] Formuly agregujace (SUMIF, COUNTIF, AVERAGEIF)
2. [ ] Symulacje (formuly z odwolaniami do wyzszych wierszy)
3. [ ] Wykres NA KONIEC

**Przed formulami:**
- [ ] Czy formule bede kopiowac? -> Uzyj $ (odniesienia bezwzgledne)
- [ ] Czy SUMIF czy SUMIFS? (1 warunek vs wiele warunkow)
- [ ] Przetestuj formule na pierwszym wierszu, potem kopiuj

**Wykres — nie zapomnij:**
- [ ] Tytul wykresu
- [ ] Opis osi X i Y
- [ ] Legenda (jesli wiele serii)
- [ ] Odpowiedni typ (kolumnowy / kolowy / liniowy)

---

## Etap 4: PROGRAMOWANIE — LATWE (30-40 min) [1:50 - 2:30]

Kolejnosc:
1. [ ] Wczytaj dane z pliku (przetestuj na `_przyklad.txt`!)
2. [ ] Zrob najlatwiejsze podzadanie (zliczanie / filtrowanie)
3. [ ] Nastepne podzadania wg trudnosci

**Szablon startowy:**
```cpp
#include <iostream>
#include <fstream>
#include <string>
#include <algorithm>
using namespace std;

int main() {
    ifstream plik("dane.txt");
    // wczytywanie...
    plik.close();
    return 0;
}
```

**Przy kazdym podzadaniu:**
- [ ] Czy program kompiluje sie bez bledow?
- [ ] Czy wynik na `_przyklad.txt` zgadza sie z oczekiwanym?
- [ ] Czy wynik na pelnym pliku `dane.txt` jest sensowny?
- [ ] Czy zapisalem plik wynikowy (jesli wymagany)?

---

## Etap 5: TEORIA / ALGORYTMY (20-30 min) [2:30 - 3:00]

- [ ] Sledzenie rekurencji -> **RYSUJ DRZEWO WYWOLAN** (tabelka!)
- [ ] Analiza zlozonosci -> petla w petli = O(n^2), dzielenie na pol = O(log n)
- [ ] Projektowanie algorytmu -> pseudokod CKE (przypisanie: `:=`, `dla i=1,2,...,n`, `dopoki`)

**Sledzenie — kolejnosc:**
1. Sprawdz warunek bazowy
2. Zapisz parametry wywolania
3. Rozwin wywolania rekurencyjne
4. Zbierz wyniki od lisci do korzenia

---

## Etap 6: TRUDNE PODZADANIA (reszta czasu) [3:00 - 3:30]

- [ ] Wroc do pominieteych podzadan
- [ ] DP, DFS/BFS, zlozone algorytmy
- [ ] **Nie zostawiaj pustych odpowiedzi** — czesciowe rozwiazanie = czesciowe punkty
- [ ] Nawet pseudokod lub opis algorytmu moze dac punkt

---

## Kontrola czasu — punkty kontrolne

| Czas | Co powinnismy miec zrobione | Punkty |
|------|---------------------------|--------|
| 0:20 | Rozpoznanie + quick wins | ~3-5 |
| 1:00 | + SQL | ~13-15 |
| 1:50 | + Arkusz | ~23-25 |
| 2:30 | + Programowanie (latwe) | ~33-37 |
| 3:00 | + Teoria/algorytmy | ~40-45 |
| 3:30 | + Trudne + przeglad | ~45-50 |

---

## Przed oddaniem (ostatnie 10 min)

- [ ] Czy WSZYSTKIE pliki sa zapisane?
- [ ] Czy programy kompiluja sie i uruchamiaja?
- [ ] Czy wyniki SQL-i sa sensowne?
- [ ] Czy wykres ma tytul, osie, legende?
- [ ] Czy odniesienia $ w arkuszu sa poprawne?
- [ ] Czy nie zostawiles pustych odpowiedzi? (napisz cokolwiek!)

---

*Powiazane: `przed_egzaminem.md`, `debug_checklist.md`*
