# Debug Checklist

> Cos nie dziala? Przejdz te listy zanim stracisz wiecej czasu.

---

## C++ — program nie kompiluje sie

- [ ] Brak `#include` (np. `<fstream>`, `<algorithm>`, `<string>`, `<cmath>`)
- [ ] Brak `using namespace std;` (albo `std::` przed cout/cin/string/ifstream)
- [ ] Brak srednika `;` na koncu linii
- [ ] Brak nawiasu zamykajacego `}` (funkcja, petla, if)
- [ ] Literowka w nazwie zmiennej
- [ ] `=` zamiast `==` w warunku if
- [ ] Niezadeklarowana zmienna (uzycie przed deklaracja)

## C++ — program kompiluje sie ale daje zly wynik

### Wczytywanie pliku
- [ ] Czy sciezka do pliku jest poprawna? (sprawdz folder roboczy!)
- [ ] Czy plik uzywa spacji, tabulatorow, srednikiow? (dopasuj `>>` vs `getline`)
- [ ] Czy plik ma naglowek? (moze trzeba pominac pierwsza linie)
- [ ] Czy kodowanie jest OK? (polskie znaki w danych)
- [ ] Przetestuj na `_przyklad.txt` — czy wczytuje poprawna liczbe rekordow?

### Operacje na cyfrach
- [ ] Czy `n` moze byc 0? (petla `while(n>0)` nie wejdzie!)
- [ ] Czy `n` moze byc ujemne? (mod daje ujemny wynik w C++)
- [ ] Cyfry wychodza od prawej (jednosci najpierw) — czy to wazne?

### Min / Max
- [ ] Inicjalizacja: NIE `min=0`! Uzyj `min=a[0]` lub `min=INT_MAX`
- [ ] Inicjalizacja: NIE `max=0` jesli dane moga byc ujemne! Uzyj `max=INT_MIN`
- [ ] Czy petla startuje od `i=0` czy `i=1`? (jesli init z `a[0]`, start od `i=1`)

### Najdluzsza sekwencja
- [ ] **Czy sprawdzasz PO petli?** `mx = max(mx, cl);` — KONIECZNE!
- [ ] Czy resetujesz `cl=1` (nie `cl=0`) gdy sekwencja sie przerywa?
- [ ] Czy startowa wartosc `cl=1` i `mx=1`?

### Sortowanie
- [ ] Czy zakres sort jest poprawny? `sort(a, a+n)` (nie `a+n-1`!)
- [ ] Czy komparator jest poprawny? `<` = rosnaco, `>` = malejaco
- [ ] Czy sortujesz PRZED przeszukiwaniem binarnym?

### Tablice
- [ ] Indeksowanie od 0 czy od 1? (C++ tablice od 0!)
- [ ] Czy rozmiar tablicy jest wystarczajacy? (nie za maly?)
- [ ] Czy petla nie wychodzi poza tablice? (`i < n`, nie `i <= n`)

### Petla
- [ ] Petla nieskonczona? Sprawdz warunek zakonczenia
- [ ] Off-by-one: `<` vs `<=`, `i=0` vs `i=1`
- [ ] Czy zmienna iteracyjna sie zmienia? (np. `i++` nie brakuje?)

---

## SQL — zapytanie nie dziala lub daje zly wynik

### Bledy skladniowe
- [ ] Kolejnosc klauzul: `SELECT FROM [JOIN ON] WHERE GROUP BY HAVING ORDER BY`
- [ ] Brak `ON` przy JOIN
- [ ] Przecinek pomiedzy kolumnami w SELECT (ale NIE przed FROM)
- [ ] Cudzyslow: tekst w `'apostrofach'`, nie w "podwojnych"

### Zle wyniki
- [ ] **WHERE vs HAVING**: filtrowanie wierszy = WHERE, filtrowanie grup = HAVING
- [ ] **INNER JOIN vs LEFT JOIN**: INNER odrzuci rekordy bez pary! Uzyj LEFT gdy szukasz "nie maja"
- [ ] **Brak DISTINCT**: "ile roznych" wymaga `COUNT(DISTINCT kolumna)`
- [ ] **Brak GROUP BY**: agregacje (COUNT/SUM/AVG) bez GROUP BY daja 1 wiersz
- [ ] **Aliasy**: czy alias jest konsekwentny? (np. `t1`, `t2`)
- [ ] **NULL**: `WHERE x = NULL` nie dziala! Uzyj `WHERE x IS NULL`
- [ ] **LIKE**: `%` = dowolny ciag, `_` = dokaldnie 1 znak. Np. `LIKE 'Jan%'`

### Podzapytania
- [ ] Czy podzapytanie zwraca 1 kolumne (dla IN/NOT IN)?
- [ ] Czy podzapytanie zwraca 1 wartosc (jesli uzywa `=` zamiast `IN`)?
- [ ] Czy NOT IN poprawnie wyklucza? (uwaga na NULL w podzapytaniu!)

---

## Arkusz — formuly daja zly wynik

### Odniesienia
- [ ] **Brak $**: kopiujesz formule i odniesienia sie przesuwaja?
  - `$A$1` = oba stale
  - `$A1` = kolumna stala, wiersz ruchomy
  - `A$1` = kolumna ruchoma, wiersz staly
- [ ] Sprawdz pierwsza i ostatnia komorke po skopiowaniu formuly

### Formuly warunkowe
- [ ] SUMIF: `=SUMIF(zakres_warunku; warunek; zakres_sumy)` — 3 argumenty!
- [ ] SUMIFS: `=SUMIFS(zakres_sumy; zakres_war1; war1; zakres_war2; war2)` — suma PIERWSZA!
- [ ] COUNTIF: `=COUNTIF(zakres; warunek)` — 2 argumenty
- [ ] Warunek tekstowy w cudzyslowach: `">=100"`, `"Tak"`, `"<>"&A1`
- [ ] Separator: `;` (PL) czy `,` (EN)? Sprawdz ustawienia!

### Symulacja
- [ ] Czy wiersz n odwoluje sie do wiersza n-1 (nie do stalej)?
- [ ] Czy wartosc poczatkowa (wiersz 1) jest poprawna?
- [ ] Czy formula dziala na ostatnim wierszu? (zakres nie ucieka?)

### Wykres
- [ ] Zaznaczono poprawny zakres danych (z naglowkami)?
- [ ] Typ wykresu odpowiada tresci zadania
- [ ] Tytul wykresu — jest?
- [ ] Opis osi X i Y — sa?
- [ ] Legenda — jest (jesli wiele serii)?

---

## Sledzenie algorytmu — wynik sie nie zgadza

- [ ] Czy zaczynasz od warunku bazowego rekurencji?
- [ ] Czy rysujesh drzewo wywolan? (kazde wywolanie = wezel)
- [ ] Czy odrozniasz: `return f(n-1) + f(n-2)` od `return f(n-1) * f(n-2)`?
- [ ] Czy pamietasz o operacjach PO wywolaniu rekurencyjnym? (np. `wynik * 2`)
- [ ] Indeksowanie: od 0 czy od 1? (pseudokod CKE czesto od 1!)
- [ ] `div` w pseudokodzie = dzielenie calkowite (zaokraglenie w dol)
- [ ] `mod` w pseudokodzie = reszta z dzielenia
- [ ] Czy `:=` to przypisanie (nie porownanie)?

---

## Ogolna zasada debugowania

```
1. Czy rozumiem co program/zapytanie powinno robic?
2. Czy dane wejsciowe sa poprawne? (wczytaj i wypisz)
3. Czy wynik posredni jest poprawny? (dodaj cout / wypisz krok)
4. Gdzie dokladnie jest rozbieznosc? (porownaj krok po kroku)
5. Napraw JEDEN blad naraz, przetestuj ponownie
```

**Jesli nie wiesz co jest zle — wypisz wartosci posrednie (`cout`)**

---

*Powiazane: `przed_egzaminem.md`, `podczas_egzaminu.md`*
