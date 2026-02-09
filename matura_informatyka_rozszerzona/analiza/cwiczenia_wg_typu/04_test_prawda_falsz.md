# 04. Test prawda/falsz

Typ zadania: **test_prawda_falsz**
Czestotliwosc: 10/11 lat | Laczna punktacja: 25 pkt
Kategoria: TEORIA

---

### Cwiczenie 4.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 3.2, Matura 2019 zad. 3.1

Ocen prawdziwosc ponizszych zdan. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Kazda liczba w systemie osemkowym zawiera wylacznie cyfry 0-7. | |
| b) | Liczba 1111(2) jest rowna 16(10). | |
| c) | Jedna cyfra szesnastkowa odpowiada dokladnie 4 bitom. | |
| d) | Liczba A3(16) jest mniejsza niz 200(10). | |

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

---

### Cwiczenie 4.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 zad. 3, Matura 2019 zad. 3.3, Matura 2022 zad. 3.3

Ocen prawdziwosc ponizszych zdan dotyczacych jezyka SQL. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Klauzula WHERE filtruje rekordy po wykonaniu agregacji GROUP BY. | |
| b) | LEFT JOIN zwraca wartosc NULL w kolumnach prawej tabeli dla rekordow, ktore nie maja dopasowania. | |
| c) | ORDER BY domyslnie sortuje wyniki malejaco (DESC). | |
| d) | Klauzula HAVING moze byc uzyta tylko w polaczeniu z GROUP BY. | |

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

---

### Cwiczenie 4.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2022 zad. 3.1

Ocen prawdziwosc ponizszych zdan dotyczacych algorytmow sortowania i zlozonosci obliczeniowej. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Sortowanie babelkowe ma pesymistyczna zlozonosc czasowa O(n log n). | |
| b) | Sortowanie przez wstawianie (insertion sort) jest algorytmem stabilnym. | |
| c) | Algorytm quicksort w najgorszym przypadku ma zlozonosc O(n^2). | |
| d) | Kazdy algorytm sortowania oparty na porownywaniu elementow ma zlozonosc co najmniej O(n log n) w przypadku pesymistycznym. | |

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

---

### Cwiczenie 4.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 3, Matura 2021 zad. 3

Ocen prawdziwosc ponizszych zdan dotyczacych rekurencji i struktur danych. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Kazda funkcje rekurencyjna mozna zapisac w postaci iteracyjnej (z uzyciem petli). | |
| b) | Stos (stack) dziala na zasadzie FIFO (First In First Out). | |
| c) | Rekurencja ogonowa (tail recursion) moze byc zawsze zamieniona na petle. | |
| d) | Drzewo BST (Binary Search Tree) gwarantuje wyszukiwanie elementu w czasie O(log n) dla dowolnych danych wejsciowych. | |

<details>
<summary>Odpowiedz</summary>

| Lp. | Odpowiedz | Uzasadnienie |
|-----|-----------|-------------|
| a) | **P** | Kazda funkcja rekurencyjna moze byc przeksztalcona na iteracyjna — w najgorszym przypadku symulujemy stos wywolan explicite za pomoca struktury stosu. |
| b) | **F** | Stos dziala na zasadzie LIFO (Last In First Out), nie FIFO. Struktura FIFO to kolejka (queue). |
| c) | **P** | Rekurencja ogonowa (gdy wywolanie rekurencyjne jest ostatnia operacja) moze byc mechanicznie zamieniona na petle. Kompilatory czesto dokonuja tej optymalizacji automatycznie (tail call optimization). |
| d) | **F** | BST gwarantuje O(log n) tylko gdy drzewo jest zrownowazne. W najgorszym przypadku (np. wstawianie elementow w kolejnosci rosnacej) drzewo degeneruje sie do listy i wyszukiwanie ma zlozonosc O(n). Gwarancje O(log n) daja drzewa zrownowazne, np. AVL lub czerwono-czarne. |

Odpowiedzi: a) P, b) F, c) P, d) F
</details>

---

### Cwiczenie 4.5 (trudnosc: trudne, ~3 pkt)
**Zrodlo inspiracji**: Matura 2016 zad. 3 (DNS), Matura 2018 zad. 3

Ocen prawdziwosc ponizszych zdan z zakresu sieci komputerowych, grafiki i formatow danych. Wpisz P (prawda) lub F (falsz).

| Lp. | Zdanie | P/F |
|-----|--------|-----|
| a) | Protokol HTTPS szyfruje dane za pomoca klucza symetrycznego, ktory jest uzgadniany z wykorzystaniem kryptografii asymetrycznej (klucza publicznego). | |
| b) | Format BMP przechowuje obrazy z kompresja stratna, podobnie jak JPEG. | |
| c) | Adres IPv4 sklada sie z 4 oktetow (grup po 8 bitow), co daje lacznie 32 bity. | |
| d) | Serwer DNS zamienia nazwy domen (np. www.example.com) na adresy MAC urzadzen sieciowych. | |

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
