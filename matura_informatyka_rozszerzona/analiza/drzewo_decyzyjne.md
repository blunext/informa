# Schemat Decyzyjny — Drzewo Rozwiazywania Zadan Maturalnych

> Dokument prowadzi od opisu problemu do konkretnego algorytmu i wzorca kodu.
> Oparty na analizie 12 lat matur (2014-2025).

---

## Sekcja 1: Flowchart — "Jaki to typ zadania?"

```
START: Czytasz tresc zadania
  |
  |-- Czy jest baza danych / tabele / SQL?
  |     TAK --> [SQL] Idz do: "SQL" w Sekcji 2
  |
  |-- Czy wymaga arkusza kalkulacyjnego / Excel / formul?
  |     TAK --> Czy trzeba zrobic wykres?
  |               TAK --> typ: arkusz_wykres
  |               NIE --> Czy symulacja krok-po-kroku (wiersz n zalezy od n-1)?
  |                         TAK --> typ: arkusz_symulacja
  |                         NIE --> typ: arkusz_agregacja_warunkowa
  |
  |-- Czy trzeba napisac program / kod / wczytac plik z danymi?
  |     TAK --> [IMPLEMENTACJA] Idz do: "Implementacja" w Sekcji 2
  |             Podtyp rozpoznaj wg danych:
  |               - Liczby/cyfry --> cyfry_liczby / minmax / zliczanie
  |               - Napisy/tekst --> napisy
  |               - Ciagi/sekwencje --> sekwencje
  |               - Siatka/piksele --> obrazy_2D
  |               - Wiele krokow, zlozony algorytm --> zlozone
  |
  |-- Czy jest pseudokod / schemat blokowy / algorytm do przesledzenia?
  |     TAK --> Czy trzeba podac wynik dla danych wejsciowych?
  |               TAK --> typ: sledzenie_algorytmu
  |               NIE --> Czy trzeba ocenic zlozonosc / wlasciwosci?
  |                         TAK --> typ: analiza_algorytmu
  |                         NIE --> typ: projektowanie_algorytmu
  |
  |-- Czy to test Prawda/Falsz?
  |     TAK --> typ: test_prawda_falsz
  |
  |-- Czy dotyczy systemow liczbowych (bin/oct/hex)?
  |     TAK --> typ: konwersja_systemow_liczbowych
  |
  |-- Czy to krotkie pytanie za 1-2 pkt (bezpieczenstwo, protokoly)?
        TAK --> typ: quick_win (teoria_bezpieczenstwa)
```

**Wskazowka**: W nowej formule (2023+) zadania sa przemieszane — nie zakladaj kolejnosci! Przeczytaj WSZYSTKIE zadania i zacznij od najlatwiejszych.

---

## Sekcja 2: "Mam problem X — uzyj algorytmu Y"

### Operacje na liczbach

| Widzisz w zadaniu | Algorytm | C++ |
|---|---|---|
| "cyfry liczby", "suma cyfr", "iloczyn cyfr" | Petla mod/div | `while(n>0){c=n%10; /*...*/ n/=10;}` |
| "NWD", "najwiekszy wspolny dzielnik" | Euklides | `while(b){t=b; b=a%b; a=t;}` |
| "NWW", "najmniejsza wspolna wielokrotnosc" | NWW = a*b/NWD(a,b) | `nww = a / nwd(a,b) * b;` |
| "liczba pierwsza", "czy pierwsza" | Test do sqrt(n) | `for(i=2;i*i<=n;i++) if(n%i==0) return false;` |
| "dzielniki liczby", "czynniki" | Petla do sqrt(n) | `for(i=1;i*i<=n;i++) if(n%i==0)...` |
| "rozklad na czynniki pierwsze" | Faktoryzacja | `for(d=2;d*d<=n;d++) while(n%d==0){/*d*/; n/=d;}` |
| "wiele liczb pierwszych", "sito" | Sito Eratostenesa | `sito[i]=true; for(i=2;i*i<=N;i++) if(sito[i]) for(j=i*i;j<=N;j+=i) sito[j]=false;` |
| "potega", "k-ta potega" | Szybkie potegowanie / petla | `while(k>0){if(k%2) w*=a; a*=a; k/=2;}` |

### Wyszukiwanie i porzadkowanie

| Widzisz w zadaniu | Algorytm | C++ |
|---|---|---|
| "najmniejszy/najwiekszy element" | Petla min/max | `if(x>mx){mx=x; idx=i;}` |
| "n-ty po posortowaniu", "mediana" | sort + indeks | `sort(a,a+n); cout<<a[k];` |
| "wyszukaj w posortowanej" | Binary search | `while(l<=r){mid=(l+r)/2; if(a[mid]==v) ...; else if(a[mid]<v) l=mid+1; else r=mid-1;}` |
| "ile spelnia warunek" | Licznik z warunkiem | `if(warunek(x)) cnt++;` |
| "posortuj wg klucza" | sort z komparatorem | `sort(a,a+n,[](auto&a,auto&b){return a.key<b.key;});` |

### Ciagi i sekwencje

| Widzisz w zadaniu | Algorytm | C++ |
|---|---|---|
| "najdluzszy ciag/podciag spl. warunek" | curr_len / max_len | `if(war) cl++; else {mx=max(mx,cl); cl=1;}` **+ `mx=max(mx,cl);` PO petli!** |
| "najdluzszy wspolny podciag" (LCS) | DP 2D | `if(a[i]==b[j]) dp[i][j]=dp[i-1][j-1]+1; else dp[i][j]=max(dp[i-1][j],dp[i][j-1]);` |
| "bloki jedynek/zer", "serie" | Zliczanie zmian | `if(a[i]!=a[i-1]) bloki++;` |
| "podciag rosnacy/malejacy" | Porownanie sasiednich | `if(a[i]>a[i-1]) rosnacy++;` |
| "sliding window", "okno" | Dwa wskazniki / suma okna | `sum+=a[r]; while(sum>k) sum-=a[l++];` |

### Napisy i tekst

| Widzisz w zadaniu | Algorytm | C++ |
|---|---|---|
| "palindrom" | Porownanie od krancow | `for(i=0;i<n/2;i++) if(s[i]!=s[n-1-i]) return false;` |
| "szyfr Cezara", "przesuniecie" | Arytmetyka mod 26 | `c = (c-'A'+shift)%26 + 'A';` |
| "kody ASCII", "wartosc znaku" | Rzutowanie char<->int | `int k=(int)c; char z=(char)(k+1);` |
| "podciag/podslowo" | substr / petla | `s.substr(pos,len)` lub reczna petla |
| "zamiana znakow", "filtrowanie" | Petla po znakach | `for(char c:s) if(warunek(c)) wynik+=c;` |

### Rekurencja vs iteracja

| Widzisz w zadaniu | Podejscie | Jak? |
|---|---|---|
| "przesledzic funkcje rekurencyjna" | Tabelka z drzewem wywolan | Rysuj drzewo: kazde wywolanie = wezel, zapisuj parametry i wynik |
| "zamien na iteracje" | Petla while + akumulator | Baze rekurencji -> warunek petli; wywolanie -> aktualizacja zmiennych |
| "drzewo BST", "preorder/inorder" | Rekurencja naturalna | Lewy -> Prawy; drzewo wywolan odwzorowuje strukture drzewa |
| "ile wywolan", "zlozonosc rekurencji" | Drzewo wywolan + wzor | Fib: O(2^n), dziel na pol: O(log n), liniowa: O(n) |

### Systemy liczbowe

| Widzisz w zadaniu | Algorytm | C++ |
|---|---|---|
| "zamien na system k" (10->k) | Dzielenie z reszta | `while(n>0){res=char('0'+n%k)+res; n/=k;}` |
| "zamien z systemu k na 10" (k->10) | Schemat Hornera | `for(char c:s) r = r*k + (c-'0');` |
| "bin na hex" (2->16) | Grupowanie po 4 bity **od prawej** | 0000=0, ..., 1001=9, 1010=A, ..., 1111=F |
| "hex na bin" (16->2) | Kazda cyfra hex = 4 bity | A=1010, F=1111, 3=0011 |
| "dodawanie/odejmowanie w systemie" | Kolumna po kolumnie od prawej | Jak pisemnie: suma cyfr + przeniesienie, dziel przez podstawe |
| "kod Graya" | XOR z przesuniecia | `gray = n ^ (n >> 1)` |

### Tablice 2D / siatki

| Widzisz w zadaniu | Algorytm | C++ |
|---|---|---|
| "sciezka na planszy", "min koszt" | DP 2D | `dp[i][j] = a[i][j] + min(dp[i-1][j], dp[i][j-1]);` |
| "polaczone obszary", "wyspy" | DFS/BFS (flood fill) | Rekurencyjny DFS lub kolejka BFS, oznaczaj odwiedzone |
| "obraz, piksele, siatka" | Iteracja 2D | `for(i) for(j) przetworz(tab[i][j]);` |
| "sasiedzi komorki" | 4 lub 8 kierunkow | `int dx[]={-1,1,0,0}; int dy[]={0,0,-1,1};` |

### SQL — bazy danych

| Widzisz w zadaniu | Wzorzec SQL | Kod |
|---|---|---|
| "dla kazdego X podaj sume/liczbe Y" | GROUP BY + agregacja | `SELECT x, COUNT(*) FROM t GROUP BY x` |
| "ktorzy NIE maja / NIE figuruja" | NOT IN lub LEFT JOIN+NULL | `WHERE id NOT IN (SELECT id FROM ...)` |
| "dane z 2-3 tabel" | JOIN (INNER) | `FROM t1 JOIN t2 ON t1.id = t2.fk` |
| "filtruj grupy" (np. >5) | HAVING (nie WHERE!) | `GROUP BY x HAVING COUNT(*) > 5` |
| "brak rekordu w drugiej tabeli" | LEFT JOIN + IS NULL | `LEFT JOIN t2 ON ... WHERE t2.id IS NULL` |
| "wartosc z zakresu", "pomiedzy" | BETWEEN lub >= AND <= | `WHERE rok BETWEEN 2020 AND 2025` |
| "tekst zawiera", "zaczyna sie" | LIKE | `WHERE name LIKE 'Jan%'` |
| "rozne wartosci", "unikalne" | DISTINCT | `SELECT DISTINCT kolumna FROM ...` |
| "warunkowo w kolumnie" | CASE WHEN | `CASE WHEN x>10 THEN 'duzo' ELSE 'malo' END` |
| "n pierwszych / ostatnich" | ORDER BY + LIMIT | `ORDER BY col DESC LIMIT 5` |

### Arkusz kalkulacyjny

| Widzisz w zadaniu | Narzedzie | Formula |
|---|---|---|
| "suma/liczba wg warunku" | SUMIF / COUNTIF | `=SUMIFS(D:D; B:B; "X"; C:C; ">100")` |
| "srednia wg warunku" | AVERAGEIF(S) | `=AVERAGEIFS(D:D; B:B; "X")` |
| "zlicz unikalne" | COUNTIF + pomocnicza | `=1/COUNTIF(zakres; wartosc)` -> SUMA |
| "symulacja dzien po dniu" | Formula z odwolaniem wyzej | Wiersz n odwoluje sie do wiersza n-1 |
| "wykres" | Zaznacz + wstaw | Kolumnowy / kolowy / liniowy + tytul + os + legenda |
| "grupowanie po okresach" | Tabela przestawna / SUMIFS | Podzial dat na okresy -> agregacja |
| "odniesienie bezwzgledne" | Dolary ($) | `$A$1` (stale) vs `A1` (wzgledne) vs `$A1` (kolumna stala) |

---

## Sekcja 3: Pulapki — "Za co tracisz punkty"

### TOP 10 najczestszych bledow (na podstawie zasad oceniania CKE)

**1. Sledzenie rekurencji — kumulacja bledow**
Jeden zly krok = caly lancuch wywolan zle. Rysuj drzewo wywolan, sprawdzaj warunek bazowy NAJPIERW.

**2. Projektowanie algorytmu — uzycie zakazanych builtinow**
Jezeli zadanie mowi "bez uzycia funkcji bibliotecznych na stringach" — nie mozesz uzyc `strlen()`, `substr()` itd. Pisz petlami.

**3. Sekwencje — brak sprawdzenia PO petli**
Klasyczny blad: `if(war) cl++; else {mx=max(mx,cl); cl=1;}` — ale po petli ostatnia sekwencja nie zostala porownana! Dodaj: `mx = max(mx, cl);` PO zakonczeniu petli.

**4. SQL — INNER JOIN zamiast LEFT JOIN**
Gdy szukasz rekordow, ktore NIE MAJA odpowiednika w drugiej tabeli, musisz uzyc `LEFT JOIN ... WHERE t2.id IS NULL`. INNER JOIN odfiltruje wlasnie te rekordy, ktore cie interesuja.

**5. SQL — WHERE zamiast HAVING na agregacie**
`WHERE COUNT(*) > 5` to blad skladniowy! Filtrowanie agregatow wymaga `HAVING`. WHERE filtruje PRZED grupowaniem, HAVING filtruje PO.

**6. SQL — brak DISTINCT**
Gdy pytanie mowi "ile roznych..." lub "unikalne wartosci", musisz dodac `DISTINCT`. Bez tego policzysz duplikaty.

**7. Arkusz — brak $ w odniesieniach bezwzglednych**
Gdy kopiujesz formule w dol/bok, odniesienia sie przesuwaja. Uzyj `$` do zablokowania: `$A$1` (oba), `$A1` (kolumna), `A$1` (wiersz).

**8. Arkusz — brak tytulu/legendy/osi na wykresie**
CKE odejmuje punkty za brak: tytulu wykresu, opisu osi, legendy (gdy wiele serii). To latwy punkt do stracenia (-1 pkt).

**9. Systemy liczbowe — grupowanie bitow od lewej zamiast od prawej**
Przy konwersji bin->hex grupujesz po 4 bity **OD PRAWEJ**. Np. `10111` = `0001 0111` = `17₁₆`, NIE `1011 1000`.

**10. Min/max — zla inicjalizacja**
Nie inicjalizuj `min = 0` (bo dane moga byc > 0 i nigdy nie zaktualizujesz!). Uzywaj `min = INT_MAX` / `max = INT_MIN` lub `min = a[0]` (pierwszym elementem).

### Dodatkowe pulapki:

**11. Wczytywanie pliku — zla sciezka lub format**
Sprawdz czy plik uzywa spacji, tabulatorow, czy srednikow jako separatora. Przetestuj na pliku `_przyklad.txt` NAJPIERW.

**12. Konwersja systemow — pomieszanie kierunku reszty**
Reszty z dzielenia odczytujemy **od konca** (od ostatniej do pierwszej). Zapisuj w `string` dodajac na poczatek, albo odwroc na koniec.

---

## Sekcja 4: Kolejnosc na egzaminie

### Strategia czasowa (210 min, 50 pkt)

```
Etap 1: ROZPOZNANIE (10 min)
  Przeczytaj WSZYSTKIE zadania. Oznacz trudnosc: latwe / srednie / trudne.

Etap 2: QUICK WINS (5-10 min, ~3-5 pkt)
  - Pytania P/F
  - Krotkie pytania za 1-2 pkt (bezpieczenstwo, protokoly, systemy liczbowe)
  - Proste sledzenie algorytmu

Etap 3: SQL (30-40 min, ~8-10 pkt)
  - Zacznij od najprostszego (SELECT WHERE)
  - Potem GROUP BY + agregacje
  - Na koniec JOIN i podzapytania
  - Testuj kazde zapytanie na danych przykladowych

Etap 4: ARKUSZ KALKULACYJNY (40-50 min, ~10 pkt)
  - Formuly agregujace (SUMIF, COUNTIF)
  - Symulacje (formuly z odwolaniami)
  - Wykres NA KONIEC (tytul + osie + legenda!)

Etap 5: PROSTE PODZADANIA PROGRAMISTYCZNE (30-40 min)
  - Zliczanie, filtrowanie, min/max
  - Operacje na cyfrach (mod/div)
  - Zawsze testuj na pliku _przyklad.txt

Etap 6: TEORIA I ALGORYTMY (20-30 min)
  - Analiza zlozonosci
  - Projektowanie algorytmu (pseudokod)
  - Sledzenie rekurencji (tabelka!)

Etap 7: TRUDNE PODZADANIA (reszta czasu)
  - Zlozone programowanie (DP, DFS/BFS, LCS)
  - Trudne SQL (wielokrotne JOIN, podzapytania zagniezdzone)
  - Jesli nie umiesz — napisz COKOLWIEK (czesciowe punkty!)
```

### Zasada 80/20
- ~40 pkt (80%) mozna zdobyc z zadan standardowych (SQL + arkusz + proste programowanie + quick wins)
- ~10 pkt (20%) wymaga algorytmiki i trudnego programowania
- **NIE trwon czasu na trudne zadanie, jesli nie zrobiles latweych!**

### Checklist przed oddaniem:
- [ ] Czy program kompiluje sie i uruchamia?
- [ ] Czy wyniki zgadzaja sie z plikiem `_przyklad.txt`?
- [ ] Czy SQL-e zwracaja sensowne wyniki?
- [ ] Czy wykres ma tytul, osie, legende?
- [ ] Czy odniesienia w arkuszu sa bezwzgledne tam, gdzie trzeba ($)?
- [ ] Czy zapisales WSZYSTKIE pliki?

---

## Pokrycie 23 typow zadan

| # | Typ zadania | Sekcja 2 — gdzie szukac |
|---|---|---|
| 1 | sledzenie_algorytmu | Rekurencja vs iteracja: tabelka z drzewem wywolan |
| 2 | projektowanie_algorytmu | Cala Sekcja 2 — dobierz algorytm do problemu |
| 3 | analiza_algorytmu | Rekurencja: zlozonosc; Wyszukiwanie: O(log n) vs O(n) |
| 4 | test_prawda_falsz | Quick win — przeczytaj uwazanie, sprawdz kazde zdanie |
| 5 | konwersja_systemow | Systemy liczbowe: 6 wzorcow konwersji |
| 6 | teoria_bezpieczenstwa | Quick win — wiedza ogolna |
| 7 | cyfry_liczby | Operacje na liczbach: mod/div, NWD, l.pierwsze, sito |
| 8 | napisy | Napisy i tekst: palindrom, szyfr, ASCII, filtrowanie |
| 9 | zlozone | Ciagi + Tablice 2D + Operacje na liczbach (kombinacja) |
| 10 | zliczanie | Wyszukiwanie: licznik z warunkiem |
| 11 | minmax | Wyszukiwanie: petla min/max, sort+indeks |
| 12 | sekwencje | Ciagi i sekwencje: curr_len/max_len, bloki |
| 13 | obrazy_2D | Tablice 2D: DFS/BFS, iteracja 2D, sasiedzi |
| 14 | geometryczne | Operacje na liczbach + Tablice 2D |
| 15 | arkusz_agregacja_warunkowa | Arkusz: SUMIF, COUNTIF, AVERAGEIF |
| 16 | arkusz_symulacja | Arkusz: formula z odwolaniem wyzej |
| 17 | arkusz_wykres | Arkusz: zaznacz + wstaw + tytul/osie/legenda |
| 18 | arkusz_agregacja_podstawowa | Arkusz: SUM, COUNT, AVERAGE |
| 19 | arkusz_transformacja | Arkusz: tabela przestawna / SUMIFS |
| 20 | sql_group_by | SQL: GROUP BY + agregacja |
| 21 | sql_podzapytania | SQL: NOT IN, LEFT JOIN+NULL |
| 22 | sql_join | SQL: JOIN (INNER/LEFT) |
| 23 | sql_select_where | SQL: SELECT z WHERE, BETWEEN, LIKE |

---

*Ostatnia aktualizacja: 2026-02-08*
*Zrodlo danych: analiza 12 lat matur (2014-2025)*
