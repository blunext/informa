# Checklist: Przed Egzaminem

> Co przygotowac, co powtorzyc, co zabrac. Przejrzyj dzien wczesniej i rano przed wyjsciem.

---

## Co zabrac na egzamin

- [ ] Dowod osobisty
- [ ] Dlugopis (czarny) + zapasowy
- [ ] Kalkulator prosty (nie programowalny, nie graficzny)
- [ ] Linijka (do wykresow w arkuszu)
- [ ] Woda / przekaska
- [ ] Zegarek (telefon zabieraja!)

## Srodowisko — przygotuj dzien wczesniej

- [ ] Wiesz jaki IDE/edytor bedzie na egzaminie (Code::Blocks / Dev-C++ / Visual Studio)
- [ ] Wiesz jak skompilowac i uruchomic program C++ w tym IDE
- [ ] Wiesz jak otworzyc plik .txt w programie (sciezka do danych!)
- [ ] Wiesz jak otworzyc arkusz kalkulacyjny (Calc / Excel)
- [ ] Wiesz jak otworzyc baze danych (Access / Base) i uruchomic zapytanie SQL
- [ ] Wiesz jak zapisac plik wynikowy (TXT, XLSX/ODS, MDB/ODB)

## Co powtorzyc (ostatni wieczor / rano)

### C++ — wzorce wczytywania plikow
```cpp
ifstream plik("dane.txt");
int x;
while (plik >> x) { /* ... */ }
```
- [ ] Wczytywanie liczb (spacje)
- [ ] Wczytywanie linii (`getline`)
- [ ] Wczytywanie par/trojek wartosci

### Algorytmy TOP 5 (najczestsze)
- [ ] Cyfry liczby: `while(n>0) { c=n%10; n/=10; }`
- [ ] NWD Euklides: `while(b) { t=b; b=a%b; a=t; }`
- [ ] Test pierwszosci: `for(i=2; i*i<=n; i++)`
- [ ] Min/max: `if(x>mx) mx=x;` (init: `mx=a[0]` lub `INT_MIN`)
- [ ] Najdluzsza sekwencja: `curr_len/max_len` + **sprawdzenie PO petli**

### SQL — kolejnosc klauzul
```
SELECT ... FROM ... JOIN ... ON ...
WHERE ... GROUP BY ... HAVING ...
ORDER BY ... LIMIT ...
```
- [ ] WHERE = przed grupowaniem, HAVING = po grupowaniu
- [ ] LEFT JOIN + IS NULL = "nie maja"
- [ ] DISTINCT = "ile roznych"

### Arkusz — kluczowe formuly
- [ ] `=SUMA.JEŻELI(zakres_war; warunek; zakres_sum)` / SUMA.WARUNKÓW
- [ ] `=LICZ.JEŻELI(zakres; warunek)`
- [ ] Odniesienia: `$A$1` (bezwzgledne), `$A1` (kolumna stala), `A$1` (wiersz staly)
- [ ] Wykres: tytul + opis osi + legenda

### Systemy liczbowe
- [ ] 10 -> k: dziel z reszta, reszty od konca
- [ ] k -> 10: Horner: `r = r*k + cyfra`
- [ ] bin -> hex: grupuj po 4 bity **od prawej**

## Mentalnosc

- Przeczytaj WSZYSTKIE zadania przed rozpoczeciem
- Zacznij od najlatwiejszych (quick wins)
- Nie tkwij dluzej niz 15 min w jednym zadaniu bez postepow — przejdz dalej
- Czesciowe rozwiazanie = czesciowe punkty. Napisz COKOLWIEK
- 40/50 pkt jest do zdobycia z zadan standardowych (SQL + arkusz + proste programowanie)

---

*Powiazane: `podczas_egzaminu.md`, `debug_checklist.md`*
