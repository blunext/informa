# 04. Test prawda/falsz

Typ zadania: **test_prawda_falsz**
Czestotliwosc: 11/12 lat | Laczna punktacja: 25 pkt
Kategoria: TEORIA

## Umiejetnosci cwiczone w tym zestawie

`systemy-liczbowe` `konwersja-bin-hex-oct` `SQL-skladnia` `WHERE-vs-HAVING` `JOIN` `ORDER-BY` `sortowanie-zlozonosc` `stabilnosc-sortowania` `rekurencja` `struktury-danych` `LIFO-FIFO` `BST` `sieci-komputerowe` `HTTPS-szyfrowanie` `IPv4` `DNS` `formaty-danych` `kompresja` `drzewa` `grafy` `tablice-haszujace`

---

### Cwiczenie 4.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 3.2, Matura 2019 zad. 3.1
**Tagi**: `systemy-liczbowe` `konwersja-bin-hex-oct`

Ocen prawdziwosc ponizszych zdan. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Kazda liczba w systemie osemkowym zawiera wylacznie cyfry 0-7. | |
| b) | Liczba 1111(2) jest rowna 16(10). | |
| c) | Jedna cyfra szesnastkowa odpowiada dokladnie 4 bitom. | |
| d) | Liczba A3(16) jest mniejsza niz 200(10). | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Przypomnij sobie zasady systemow pozycyjnych — jakie cyfry sa dopuszczalne w systemie o podstawie p?
2. **Podejscie**: Dla b) oblicz 1111(2) recznie: 1*8+1*4+1*2+1*1. Dla d) przelicz A3(16) na dziesietny.
3. **Kluczowy krok**: A(16) = 10(10). Wiec A3(16) = 10*16 + 3 = 163.

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **P** | W systemie o podstawie 8 dopuszczalne cyfry to 0, 1, 2, 3, 4, 5, 6, 7. |
| b) | **F** | 1111(2) = 1*8 + 1*4 + 1*2 + 1*1 = 15, nie 16. Liczba 10000(2) = 16. |
| c) | **P** | 4 bity moga reprezentowac wartosci 0-15, co odpowiada dokladnie jednej cyfrze szesnastkowej (0-9, A-F). |
| d) | **P** | A3(16) = 10*16 + 3 = 163(10). Poniewaz 163 < 200, zdanie jest prawdziwe — A3(16) jest rzeczywiscie mniejsza niz 200(10). |

Odpowiedzi: a) P, b) F, c) P, d) P
</details>

<details>
<summary>Typowe bledy</summary>

- **1111(2) = 16**: Czesty blad — 1111(2) = 15 (2^4 - 1). Liczba 16 to 10000(2) = 2^4. CKE: -0.5 pkt
- **Zapomnienie ze A=10 w hex**: W systemie szesnastkowym: A=10, B=11, C=12, D=13, E=14, F=15. CKE: -0.5 pkt

</details>

---

### Cwiczenie 4.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 3, Matura 2019 zad. 3.3, Matura 2022 zad. 3.3
**Tagi**: `SQL-skladnia` `WHERE-vs-HAVING` `JOIN` `ORDER-BY`

Ocen prawdziwosc ponizszych zdan dotyczacych jezyka SQL. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Klauzula WHERE filtruje rekordy po wykonaniu agregacji GROUP BY. | |
| b) | LEFT JOIN zwraca wartosc NULL w kolumnach prawej tabeli dla rekordow, ktore nie maja dopasowania. | |
| c) | ORDER BY domyslnie sortuje wyniki malejaco (DESC). | |
| d) | Klauzula HAVING moze byc uzyta tylko w polaczeniu z GROUP BY. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Przypomnij sobie kolejnosc wykonywania klauzul SQL: FROM -> WHERE -> GROUP BY -> HAVING -> SELECT -> ORDER BY.
2. **Podejscie**: WHERE dziala PRZED GROUP BY, HAVING dziala PO. LEFT JOIN zachowuje wszystkie wiersze z lewej tabeli.
3. **Kluczowy krok**: ORDER BY domyslnie = ASC (rosnaco), nie DESC.

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **F** | WHERE filtruje rekordy PRZED agregacja (przed GROUP BY). Do filtrowania po agregacji sluzy HAVING. Kolejnosc: WHERE -> GROUP BY -> HAVING. |
| b) | **P** | LEFT JOIN zachowuje wszystkie rekordy z lewej tabeli. Jezeli rekord nie ma dopasowania w prawej tabeli, kolumny prawej tabeli wypelniane sa wartoscia NULL. |
| c) | **F** | ORDER BY domyslnie sortuje rosnaco (ASC). Aby sortowac malejaco, nalezy jawnie uzyc DESC. |
| d) | **P** | HAVING filtruje grupy po agregacji, wiec wymaga GROUP BY (lub calej tabeli jako jednej grupy). W standardowym SQL HAVING stosuje sie po GROUP BY. |

Odpowiedzi: a) F, b) P, c) F, d) P
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomylenie WHERE i HAVING**: WHERE filtruje WIERSZE przed grupowaniem, HAVING filtruje GRUPY po agregacji. CKE: -1 pkt
- **ORDER BY DESC jako domyslne**: Domyslne to ASC (rosnaco). Czesty blad na maturze. CKE: -0.5 pkt
- **LEFT JOIN vs INNER JOIN**: LEFT zachowuje wszystkie wiersze z lewej tabeli (NULL dla brakujacych), INNER — tylko dopasowane. CKE: -0.5 pkt

</details>

---

### Cwiczenie 4.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 3.1
**Tagi**: `sortowanie-zlozonosc` `stabilnosc-sortowania`

Ocen prawdziwosc ponizszych zdan dotyczacych algorytmow sortowania i zlozonosci obliczeniowej. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Sortowanie babelkowe ma pesymistyczna zlozonosc czasowa O(n log n). | |
| b) | Sortowanie przez wstawianie (insertion sort) jest algorytmem stabilnym. | |
| c) | Algorytm quicksort w najgorszym przypadku ma zlozonosc O(n^2). | |
| d) | Kazdy algorytm sortowania oparty na porownywaniu elementow ma zlozonosc co najmniej O(n log n) w przypadku pesymistycznym. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Przypomnij sobie zlozonosci: babelkowe O(n^2), quicksort O(n^2) pesym. / O(n log n) sred., merge sort O(n log n).
2. **Podejscie**: Stabilnosc = elementy o rownych kluczach zachowuja kolejnosc. Insertion sort i merge sort sa stabilne, quicksort i heap sort nie.
3. **Kluczowy krok**: Dolna granica Omega(n log n) dotyczy sortowan opartych na porownaniach (twierdzenie o drzewie decyzyjnym).

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **F** | Sortowanie babelkowe ma pesymistyczna zlozonosc O(n^2), nie O(n log n). Wymaga dwoch zagniezdznonych petli po n elementach. |
| b) | **P** | Insertion sort jest stabilny — elementy o rownych kluczach zachowuja swoja wzgledna kolejnosc. Wstawiajac element, przesuwamy go w lewo tylko dopoki napotykamy elementy scisle wieksze. |
| c) | **P** | Quicksort w najgorszym przypadku (np. juz posortowana tablica przy zlym wyborze pivota) ma zlozonosc O(n^2). Srednia zlozonosc to O(n log n). |
| d) | **P** | Dolne ograniczenie zlozonosci sortowan opartych na porownywaniu wynosi Omega(n log n). Wynika to z faktu, ze drzewo decyzyjne o n! lisciach ma wysokosc co najmniej log2(n!) = Theta(n log n). |

Odpowiedzi: a) F, b) P, c) P, d) P
</details>

<details>
<summary>Typowe bledy</summary>

- **Babelkowe O(n log n)**: Babelkowe to ZAWSZE O(n^2) pesymistycznie. Jedynie optymistycznie (z flaga) moze byc O(n). CKE: -0.5 pkt
- **Quicksort zawsze O(n log n)**: Quicksort MA najgorszy przypadek O(n^2)! Tylko srednia to O(n log n). CKE: -0.5 pkt
- **Mylenie stabilnosci z zlozonoscia**: Stabilnosc to cecha zachowania kolejnosci rownych elementow, nie ma zwiazku z szybkoscia. CKE: -0.5 pkt

</details>

---

### Cwiczenie 4.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 3, Matura 2021 zad. 3
**Tagi**: `rekurencja` `struktury-danych` `LIFO-FIFO` `BST`

Ocen prawdziwosc ponizszych zdan dotyczacych rekurencji i struktur danych. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Kazda funkcje rekurencyjna mozna zapisac w postaci iteracyjnej (z uzyciem petli). | |
| b) | Stos (stack) dziala na zasadzie FIFO (First In First Out). | |
| c) | Rekurencja ogonowa (tail recursion) moze byc zawsze zamieniona na petle. | |
| d) | Drzewo BST (Binary Search Tree) gwarantuje wyszukiwanie elementu w czasie O(log n) dla dowolnych danych wejsciowych. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Kazda rekurencja mozna zamienic na iteracje (ewentualnie z jawnym stosem). Stos to LIFO, kolejka to FIFO.
2. **Podejscie**: BST moze sie zdegenerowac do listy — kiedy to sie zdarza?
3. **Kluczowy krok**: BST ma O(log n) tylko gdy jest zrownowazony. Wstawianie posortowanych danych daje drzewo-liste z O(n).

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **P** | Kazda funkcja rekurencyjna moze byc przeksztalcona na iteracyjna — w najgorszym przypadku symulujemy stos wywolan explicite za pomoca struktury stosu. |
| b) | **F** | Stos dziala na zasadzie LIFO (Last In First Out), nie FIFO. Struktura FIFO to kolejka (queue). |
| c) | **P** | Rekurencja ogonowa (gdy wywolanie rekurencyjne jest ostatnia operacja) moze byc mechanicznie zamieniona na petle. Kompilatory czesto dokonuja tej optymalizacji automatycznie (tail call optimization). |
| d) | **F** | BST gwarantuje O(log n) tylko gdy drzewo jest zrownowazne. W najgorszym przypadku (np. wstawianie elementow w kolejnosci rosnacej) drzewo degeneruje sie do listy i wyszukiwanie ma zlozonosc O(n). Gwarancje O(log n) daja drzewa zrownowazone, np. AVL lub czerwono-czarne. |

Odpowiedzi: a) P, b) F, c) P, d) F
</details>

<details>
<summary>Typowe bledy</summary>

- **Stos = FIFO**: CZESTY blad! Stos = LIFO (ostatni wchodzi, pierwszy wychodzi). Kolejka = FIFO. CKE: -0.5 pkt
- **BST zawsze O(log n)**: BST moze degenerowac do listy! Tylko zrownowazony BST (AVL, czerwono-czarny) gwarantuje O(log n). CKE: -0.5 pkt
- **Nie kazda rekurencja da sie zamienic na iteracje**: To falszywe twierdzenie — kazda rekurencja DA sie zamienic (z jawnym stosem). CKE: -0.5 pkt

</details>

---

### Cwiczenie 4.5 (trudnosc: trudne, ~3 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 3 (DNS), Matura 2018 zad. 3
**Tagi**: `sieci-komputerowe` `HTTPS-szyfrowanie` `IPv4` `DNS` `formaty-danych`

Ocen prawdziwosc ponizszych zdan z zakresu sieci komputerowych, grafiki i formatow danych. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Protokol HTTPS szyfruje dane za pomoca klucza symetrycznego, ktory jest uzgadniany z wykorzystaniem kryptografii asymetrycznej (klucza publicznego). | |
| b) | Format BMP przechowuje obrazy z kompresja stratna, podobnie jak JPEG. | |
| c) | Adres IPv4 sklada sie z 4 oktetow (grup po 8 bitow), co daje lacznie 32 bity. | |
| d) | Serwer DNS zamienia nazwy domen (np. www.example.com) na adresy MAC urzadzen sieciowych. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: HTTPS uzywa kryptografii hybrydowej. BMP to format bezstratny. DNS zamienia domeny na adresy IP.
2. **Podejscie**: Rozroznij: BMP (bezstratny), JPEG (stratny), PNG (bezstratny). Rozroznij: IP (warstwa sieciowa) vs MAC (warstwa lacza).
3. **Kluczowy krok**: DNS -> IP, ARP -> MAC. IPv4 = 32 bity, IPv6 = 128 bitow.

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **P** | HTTPS (TLS/SSL) korzysta z kryptografii hybrydowej: klucz asymetryczny (RSA/ECDH) sluzy do bezpiecznego uzgodnienia wspolnego klucza sesji (symetrycznego, np. AES), ktory nastepnie szyfruje wlasciwa komunikacje. |
| b) | **F** | Format BMP przechowuje obrazy BEZ kompresji (lub z kompresja bezstratna, np. RLE). BMP zachowuje pelna informacje o kazdym pikselu, w przeciwienstwie do JPEG, ktory uzywa kompresji stratnej (DCT). |
| c) | **P** | Adres IPv4 to 32-bitowa liczba zapisywana jako 4 liczby dziesietne oddzielone kropkami (np. 192.168.1.1), kazda odpowiadajaca jednemu oktetowi (8 bitow). 4 * 8 = 32 bity. |
| d) | **F** | DNS (Domain Name System) zamienia nazwy domen na adresy IP (np. 93.184.216.34), NIE na adresy MAC. Adresy MAC sa adresami warstwy lacza danych (fizycznymi) i sa rozwiazywane przez protokol ARP. |

Odpowiedzi: a) P, b) F, c) P, d) F
</details>

<details>
<summary>Typowe bledy</summary>

- **DNS -> MAC**: DNS zamienia domeny na IP, nie MAC. MAC to warstwa lacza danych (ARP). CKE: -0.5 pkt
- **BMP = kompresja stratna**: BMP jest bezstratny. Stratny = JPEG, bezstratny = BMP, PNG, TIFF. CKE: -0.5 pkt
- **HTTPS = tylko szyfrowanie asymetryczne**: HTTPS uzywa asymetrycznego TYLKO do uzgodnienia klucza. Dane sa szyfrowane symetrycznie (AES). CKE: -0.5 pkt

</details>

---

### Cwiczenie 4.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 zad. 3, styl nowej formuly
**Tagi**: `kompresja` `formaty-danych` `systemy-liczbowe`

Ocen prawdziwosc ponizszych zdan. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Format PNG wykorzystuje kompresje stratna, dlatego pliki PNG sa mniejsze niz BMP. | |
| b) | Liczba FF(16) jest rowna 256(10). | |
| c) | Kodowanie UTF-8 moze uzywac od 1 do 4 bajtow na jeden znak. | |
| d) | Kompresja bezstratna pozwala odtworzyc oryginalne dane w 100% bez utraty informacji. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: PNG to kompresja BEZstratna (nie stratna). FF(16) = 15*16 + 15 = ?
2. **Podejscie**: UTF-8: znaki ASCII to 1 bajt, polskie znaki 2 bajty, emoji do 4 bajtow.
3. **Kluczowy krok**: FF(16) = 255, nie 256. 256(10) = 100(16).

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **F** | PNG wykorzystuje kompresje BEZstratna (algorytm Deflate = LZ77 + Huffman). Pliki PNG sa mniejsze niz BMP dzieki kompresji, ale jest to kompresja bezstratna, nie stratna. |
| b) | **F** | FF(16) = 15*16 + 15 = 240 + 15 = 255, nie 256. Liczba 256(10) = 100(16). |
| c) | **P** | UTF-8 jest kodowaniem o zmiennej dlugosci: znaki ASCII (0-127) zajmuja 1 bajt, wieksznosc znakow europejskich 2 bajty, znaki azjatyckie 3 bajty, emoji i rzadkie znaki 4 bajty. |
| d) | **P** | Kompresja bezstratna (np. ZIP, PNG, FLAC) pozwala odtworzyc identyczna kopie oryginalnych danych. W przeciwienstwie do stratnej (JPEG, MP3), ktora traci czesc informacji. |

Odpowiedzi: a) F, b) F, c) P, d) P
</details>

<details>
<summary>Typowe bledy</summary>

- **PNG = kompresja stratna**: PNG to BEZstratna kompresja! Stratna = JPEG, MP3, MP4. CKE: -0.5 pkt
- **FF(16) = 256**: FF(16) = 255 = 2^8 - 1. Liczba 256 = 2^8 = 100(16). CKE: -0.5 pkt
- **UTF-8 = zawsze 1 bajt na znak**: To dotyczy tylko ASCII. UTF-8 ma zmienna dlugosc (1-4 bajty). CKE: -0.5 pkt

</details>

---

### Cwiczenie 4.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3 (algorytmy i struktury)
**Tagi**: `sortowanie-zlozonosc` `drzewa` `grafy` `struktury-danych`

Ocen prawdziwosc ponizszych zdan. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Algorytm sortowania przez scalanie (merge sort) ma zlozonosc O(n log n) zarowno w przypadku optymistycznym, jak i pesymistycznym. | |
| b) | Graf nieskierowany o n wierzcholkach i n-1 krawedziach jest zawsze drzewem. | |
| c) | W drzewie o n wierzcholkach jest dokladnie n-1 krawedzi. | |
| d) | Algorytm BFS (przeszukiwanie wszerz) uzywa stosu jako struktury pomocniczej. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Merge sort zawsze dzieli na polowy i scala — niezaleznie od danych. BFS uzywa kolejki, DFS uzywa stosu.
2. **Podejscie**: Graf o n wierzcholkach i n-1 krawedziach to drzewo TYLKO jezeli jest spojny.
3. **Kluczowy krok**: BFS = kolejka (FIFO), DFS = stos (LIFO). To CZESTY blad na maturze.

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **P** | Merge sort zawsze dzieli tablice na polowy, sortuje kazda rekurencyjnie i scala. Koszt scalania na kazdym poziomie to O(n), a poziomow jest log2(n). Stad O(n log n) w KAZDYM przypadku (optymistycznym, srednim, pesymistycznym). |
| b) | **F** | Graf o n wierzcholkach i n-1 krawedziach jest drzewem TYLKO jezeli jest spojny. Niespojny graf moze miec n wierzcholkow i n-1 krawedzi (np. cykl na 3 wierzcholkach + izolowany wierzcholek = 4 wierzcholki, 3 krawedzie). |
| c) | **P** | Kazde drzewo (spojny graf acykliczny) o n wierzcholkach ma dokladnie n-1 krawedzi. To podstawowa wlasnosc drzew. |
| d) | **F** | BFS uzywa KOLEJKI (FIFO), nie stosu. To DFS (przeszukiwanie w glab) uzywa stosu (LIFO, lub niejawnie — przez rekurencje). |

Odpowiedzi: a) P, b) F, c) P, d) F
</details>

<details>
<summary>Typowe bledy</summary>

- **BFS = stos, DFS = kolejka**: Odwrotnie! BFS = kolejka (FIFO), DFS = stos (LIFO). CKE: -0.5 pkt
- **n wierzcholkow + n-1 krawedzi = drzewo**: Brakuje warunku spojnosci! CKE: -0.5 pkt
- **Merge sort O(n^2) w pesymistycznym**: Merge sort ma O(n log n) ZAWSZE. To quicksort ma O(n^2) pesymistycznie. CKE: -0.5 pkt

</details>

---

### Cwiczenie 4.8 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2025 zad. 3 (logika, twierdzenia)
**Tagi**: `rekurencja` `zlozonosc-czasowa` `tablice-haszujace` `struktury-danych`

Ocen prawdziwosc ponizszych zdan. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Rekurencyjna implementacja obliczania n-tej liczby Fibonacciego ma zlozonosc O(n). | |
| b) | Tablica haszujaca (hash table) gwarantuje wyszukiwanie elementu w czasie O(1) w najgorszym przypadku. | |
| c) | Algorytm Dijkstry znajduje najkrotsze sciezki w grafie z wagami ujemnymi. | |
| d) | Przeszukiwanie binarne wymaga, aby tablica byla posortowana. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Fibonacci rekurencyjnie (bez memoizacji) to O(2^n), nie O(n). Hash table moze miec kolizje.
2. **Podejscie**: Dijkstra NIE dziala z wagami ujemnymi (do tego sluzy Bellman-Ford). Binsearch wymaga posortowania.
3. **Kluczowy krok**: Hash table: sredni O(1), pesymistyczny O(n) (gdy wszystkie klucze trafiaja do jednego kubla).

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **F** | Naiwna rekurencyjna implementacja Fibonacciego ma zlozonosc O(2^n) (wykladnicza), bo wielokrotnie oblicza te same wartosci. Zlozonosc O(n) ma wersja iteracyjna lub rekurencyjna z memoizacja. |
| b) | **F** | Hash table ma O(1) SREDNIO (amortyzowane), ale w najgorszym przypadku (wiele kolizji) O(n). Gwarancja O(1) pesymistycznego wymaga doskonalej funkcji haszujacej (perfect hashing), co jest rzadkie w praktyce. |
| c) | **F** | Algorytm Dijkstry NIE dziala poprawnie z ujemnymi wagami krawedzi. Do grafow z ujemnymi wagami nalezy uzyc algorytmu Bellmana-Forda. |
| d) | **P** | Przeszukiwanie binarne (binary search) wymaga, aby elementy byly posortowane. Algorytm porownuje element srodkowy z szukanym i odrzuca polowe zakresu — to dziala TYLKO gdy dane sa uporzadkowane. |

Odpowiedzi: a) F, b) F, c) F, d) P
</details>

<details>
<summary>Typowe bledy</summary>

- **Fibonacci rekurencyjnie O(n)**: To O(2^n)! Iteracyjnie lub z memoizacja = O(n). CKE: -0.5 pkt
- **Hash table zawsze O(1)**: Tylko srednio. Pesymistycznie O(n) przez kolizje. CKE: -0.5 pkt
- **Dijkstra z ujemnymi wagami**: NIE dziala! Uzyj Bellmana-Forda. CKE: -0.5 pkt

</details>

---

### Cwiczenie 4.9 (trudnosc: srednie-trudne, ~3 pkt)
**Zrodlo inspiracji**: Matura 2021 zad. 3 (szyfrowanie), Matura 2018 zad. 3
**Tagi**: `HTTPS-szyfrowanie` `sieci-komputerowe` `formaty-danych`

Ocen prawdziwosc ponizszych zdan. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | W szyfrowaniu asymetrycznym ten sam klucz sluzy do szyfrowania i deszyfrowania. | |
| b) | Adres IPv6 sklada sie z 128 bitow, co daje ponad 3.4 * 10^38 mozliwych adresow. | |
| c) | Protokol TCP gwarantuje dostarczenie pakietow w kolejnosci i bez strat, w przeciwienstwie do UDP. | |
| d) | Maska podsieci 255.255.255.0 oznacza, ze ostatni oktet adresu IP identyfikuje urzadzenie w sieci. | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Szyfrowanie asymetryczne = dwa rozne klucze (publiczny + prywatny). Symetryczne = ten sam klucz.
2. **Podejscie**: IPv6 = 128 bitow. TCP = niezawodny (z potwierdzeniami), UDP = szybki (bez gwarancji).
3. **Kluczowy krok**: Maska /24 (255.255.255.0) = 24 bity sieci + 8 bitow hosta = do 254 urzadzen.

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **F** | W szyfrowaniu ASYMETRYCZNYM sa DWA rozne klucze: publiczny (do szyfrowania) i prywatny (do deszyfrowania). To w szyfrowaniu SYMETRYCZNYM ten sam klucz sluzy do obu operacji. |
| b) | **P** | IPv6 uzywa adresow 128-bitowych. 2^128 ≈ 3.4 * 10^38, co jest ogromna przestrzenia adresowa (w porownaniu z IPv4: 2^32 ≈ 4.3 * 10^9). |
| c) | **P** | TCP (Transmission Control Protocol) zapewnia niezawodna, uporzadkowana i kontrolowana dostarke danych. UDP (User Datagram Protocol) jest szybszy, ale nie gwarantuje dostarczenia ani kolejnosci. |
| d) | **P** | Maska 255.255.255.0 (/24) oznacza, ze pierwsze 3 oktety identyfikuja siec, a ostatni oktet (8 bitow) identyfikuje urzadzenie (hosta) w tej sieci. Daje to do 254 adresow dla urzadzen (2^8 - 2, odejmujemy adres sieci i broadcast). |

Odpowiedzi: a) F, b) P, c) P, d) P
</details>

<details>
<summary>Typowe bledy</summary>

- **Asymetryczne = jeden klucz**: Asymetryczne = DWA klucze. Symetryczne = jeden klucz. CKE: -0.5 pkt
- **TCP vs UDP odwrotnie**: TCP = niezawodny, UDP = szybki ale bez gwarancji. CKE: -0.5 pkt
- **Maska /24 = 256 urzadzen**: Odejmujemy 2 (adres sieci + broadcast), wiec 254. CKE: -0.5 pkt

</details>

---

### Cwiczenie 4.10 (trudnosc: trudne, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 zad. 3 (zaawansowane twierdzenia)
**Tagi**: `zlozonosc-czasowa` `rekurencja` `struktury-danych` `grafy`

Ocen prawdziwosc ponizszych zdan. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Problem plecakowy 0-1 mozna rozwiazac w czasie wielomianowym. | |
| b) | Dla dowolnej liczby naturalnej n, suma pierwszych n liczb naturalnych wynosi n*(n+1)/2. | |
| c) | Algorytm zachlanny (greedy) zawsze daje rozwiazanie optymalne. | |
| d) | Jezeli funkcja f(n) = O(n^2) i g(n) = O(n), to f(n) + g(n) = O(n^2). | |

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Problem plecakowy 0-1 jest NP-trudny — brak algorytmu wielomianowego (chyba ze P=NP). Greedy nie zawsze daje optimum.
2. **Podejscie**: Suma 1+2+...+n = n(n+1)/2 (wzor Gaussa). Dla notacji O: O(n^2) + O(n) = O(n^2) (dominuje wyzszy wyraz).
3. **Kluczowy krok**: Greedy daje optimum np. dla problemu wydawania reszty (standardowe nominaly), ale NIE dla plecakowego 0-1.

</details>

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **F** | Problem plecakowy 0-1 jest NP-trudny. Najlepsze znane algorytmy maja zlozonosc pseudowielomianowa O(nW) gdzie W to pojemnosc plecaka, lub wykladnicza O(2^n). Nie jest znany algorytm scisle wielomianowy (i nie bedzie, chyba ze P=NP). |
| b) | **P** | Wzor Gaussa: 1+2+3+...+n = n*(n+1)/2. Mozna to udowodnic indukcyjnie: baza n=1: 1 = 1*2/2 = 1. Krok: zakladamy dla k, dowod dla k+1: k*(k+1)/2 + (k+1) = (k+1)*(k+2)/2. |
| c) | **F** | Algorytmy zachlanne NIE zawsze daja rozwiazanie optymalne. Dzialaja optymalnie tylko dla problemow o wlasnosci optymalnej podstruktury i zachlannego wyboru (np. minimalne drzewo rozpinajace, kodowanie Huffmana). Kontrprzyklad: plecak 0-1 z przedmiotami o wagach [3,4,5], wartosciach [3,4,5] i pojemnosci 7 — zachlanny wezme 5, ale optimum to 3+4=7. |
| d) | **P** | Przy dodawaniu funkcji w notacji O dominuje wyzszy rzad: O(n^2) + O(n) = O(n^2). Formalnie: istnieja stale c1, c2, n0 takie ze f(n) <= c1*n^2 i g(n) <= c2*n, wiec f(n)+g(n) <= c1*n^2 + c2*n <= (c1+c2)*n^2 = O(n^2). |

Odpowiedzi: a) F, b) P, c) F, d) P
</details>

<details>
<summary>Typowe bledy</summary>

- **Plecak 0-1 wielomianowy**: Rozwiazanie O(nW) wyglada na wielomianowe, ale W jest wykladnicze wzgledem rozmiaru wejscia (log W bitow). CKE: -0.5 pkt
- **Greedy = zawsze optimum**: Kontrprzyklad: plecak 0-1, problem komiwojazera. Greedy dziala dla: MST (Kruskal/Prim), Huffman, wydawanie reszty. CKE: -0.5 pkt
- **O(n^2) + O(n) = O(n^3)**: NIE! Przy dodawaniu bierzemy maksimum rzadow, nie mnozenie. CKE: -0.5 pkt

</details>

---

## Samoocena

Po rozwiazaniu cwiczen bez podgladania odpowiedzi, okresl swoj poziom:

| Poziom | Opis | Wynik |
|--------|------|-------|
| Podstawowy | Znasz podstawy systemow liczbowych i proste fakty o SQL | 1-3 cwiczen bez pomocy |
| Dobry | Radzisz sobie z zlozonoscia sortowania, rekurencja i strukturami danych | 4-6 cwiczen bez pomocy |
| Bardzo dobry | Poprawnie oceniasz twierdzenia o sieciach, szyfrowaniu i grafach | 7-8 cwiczen bez pomocy |
| Doskonaly | Znasz NP-trudnosc, algorytmy zachlanne i zaawansowane wlasnosci O-notacji | 9-10 cwiczen bez pomocy |

**Co dalej?**
- Poziom Podstawowy: Przerob cwiczenia 4.1, 4.2, 4.6 jeszcze raz. Wrocz do `cheatsheet_teoria.md` (sekcja: systemy liczbowe i SQL).
- Poziom Dobry: Skup sie na cwiczeniach 4.3, 4.4, 4.7, 4.8. Przejdz do `03_analiza_algorytmu.md`.
- Poziom Bardzo dobry/Doskonaly: Przejdz do `05_konwersja_systemow_liczbowych.md` i `06_teoria_bezpieczenstwa.md`.
