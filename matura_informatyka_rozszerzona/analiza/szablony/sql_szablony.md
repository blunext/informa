# Szablony SQL — Sciagawka na Mature

Samowystarczalny dokument. Kazdy szablon jest kompletnym zapytaniem gotowym do uruchomienia.

---

## 1. Kolejnosc klauzul SQL (KLUCZOWA SEKCJA)

### Kolejnosc pisania zapytania:

```sql
SELECT kolumny          -- 5. co wyswietlic
FROM tabela             -- 1. skad brac dane
JOIN tabela2 ON ...     -- 2. laczenie tabel
WHERE warunek           -- 3. filtr na wiersze (PRZED grupowaniem)
GROUP BY kolumny        -- 4. grupowanie
HAVING warunek_grupy    -- 6. filtr na grupy (PO grupowaniu)
ORDER BY kolumna        -- 7. sortowanie
LIMIT n                 -- 8. ograniczenie wynikow
```

### Kolejnosc wykonania (baza danych przetwarza w tej kolejnosci):

```
FROM → JOIN → WHERE → GROUP BY → HAVING → SELECT → ORDER BY → LIMIT
```

**Dlaczego to wazne**: Nie mozna uzyc aliasu z SELECT w WHERE (bo WHERE wykonuje sie wczesniej). Mozna uzyc aliasu w ORDER BY (bo ORDER BY wykonuje sie pozniej).

---

## 2. SELECT + WHERE (10 pkt, 4/12 lat)

### A. Filtrowanie wielowarunkowe (AND / OR)

```sql
-- Uczniowie z klasy 3A ze srednia >= 4.5
SELECT imie, nazwisko
FROM Uczniowie
WHERE klasa = '3A' AND srednia_ocen >= 4.5;
```

**Pulapka**: AND ma wyzszy priorytet niz OR. Nawiasy sa obowiazkowe:
```sql
-- ZLE: opady NOT NULL AND opady > 0 AND wiatr > 10 OR amplituda > 15
--       interpretowane jako: (... AND ... AND wiatr > 10) OR (amplituda > 15)
-- DOBRZE:
WHERE opady IS NOT NULL AND opady > 0
  AND (wiatr > 10 OR (temp_max - temp_min) > 15);
```

### B. LIKE — wzorce tekstowe

```sql
-- Produkty zaczynajace sie od 'K' lub konczace na 'a'
SELECT nazwa, cena
FROM Produkty
WHERE nazwa LIKE 'K%' OR nazwa LIKE '%a'
ORDER BY nazwa;
```

| Wzorzec | Znaczenie | Przyklad |
|---------|-----------|----------|
| `'A%'` | zaczyna sie na A | Adam, Anna |
| `'%ski'` | konczy sie na ski | Kowalski |
| `'_a%'` | druga litera to a | Jan, Kasia |
| `'%owa%'` | zawiera owa | Kowalska |

### C. BETWEEN — zakres wartosci

```sql
-- Pojazdy z lat 2018-2021
SELECT marka, model
FROM Pojazdy
WHERE rok_produkcji BETWEEN 2018 AND 2021;
```

**Pulapka**: BETWEEN jest wlaczajace (inclusive) — obejmuje obie granice. `BETWEEN 2018 AND 2021` to to samo co `>= 2018 AND <= 2021`.

### D. IN — lista wartosci

```sql
-- Pojazdy o pojemnosci 1400, 1600, 1800 lub 2000
SELECT marka, model
FROM Pojazdy
WHERE pojemnosc IN (1400, 1600, 1800, 2000);
```

Rownowazne: `WHERE pojemnosc = 1400 OR pojemnosc = 1600 OR ...` — ale IN jest krotsze.

### E. IS NULL / IS NOT NULL

```sql
-- Pomiary gdzie opady sa znane (nie NULL)
SELECT stacja, data, opady
FROM Pomiary
WHERE opady IS NOT NULL;
```

**Pulapka**: `WHERE opady != NULL` NIE DZIALA. Jedyny poprawny sposob to `IS NULL` / `IS NOT NULL`.

### F. LENGTH — dlugosc tekstu

```sql
-- Filmy z tytulem dluzszym niz 10 znakow
SELECT tytul, LENGTH(tytul) AS dlugosc
FROM Filmy
WHERE LENGTH(tytul) > 10;
```

### G. SUBSTR — wycinanie fragmentu tekstu

```sql
-- Pierwsze 3 litery tytulu
SELECT tytul, SUBSTR(tytul, 1, 3) AS poczatek
FROM Filmy
WHERE rezyser = 'Pawlikowski';
```

**Pulapka**: W SQL pozycje liczymy od 1 (nie od 0 jak w C++!). `SUBSTR('Hello', 1, 3)` daje `'Hel'`.

---

## 3. JOIN — laczenie tabel (21 pkt, 9/12 lat)

### A. INNER JOIN — 2 tabele

```sql
-- Tytul ksiazki i jej autor
SELECT K.tytul, A.imie, A.nazwisko
FROM Ksiazki K
JOIN Autorzy A ON K.id_autora = A.id_autora
ORDER BY A.nazwisko, K.tytul;
```

### B. INNER JOIN — 3 tabele

```sql
-- Wizyta: oddział → lekarz → wizyta
SELECT O.nazwa_oddzialu, L.imie, L.nazwisko, W.pacjent
FROM Wizyty W
JOIN Lekarze L ON W.id_lekarza = L.id_lekarza
JOIN Oddzialy O ON L.id_oddzialu = O.id_oddzialu
ORDER BY W.data_wizyty;
```

### C. INNER JOIN — 4 tabele (matura 2025: Mars)

```sql
-- Bazy → Astronauci → Przydzialy → Projekty
SELECT B.nazwa_bazy, A.imie, A.nazwisko, P.nazwa_projektu, PR.rola
FROM Bazy B
JOIN Astronauci A ON B.id_bazy = A.id_bazy
JOIN Przydzialy PR ON A.id_astronauty = PR.id_astronauty
JOIN Projekty P ON PR.id_projektu = P.id_projektu
WHERE P.data_rozpoczecia < '2024-04-01'
ORDER BY B.nazwa_bazy, P.nazwa_projektu;
```

### D. LEFT JOIN + IS NULL — rekordy BEZ dopasowania

```sql
-- Klienci, ktorzy nigdy nie zlozyli zamowienia
SELECT K.imie, K.nazwisko
FROM Klienci K
LEFT JOIN Zamowienia Z ON K.id_klienta = Z.id_klienta
WHERE Z.id_zamowienia IS NULL;
```

**Jak to dziala**: LEFT JOIN zachowuje WSZYSTKIE wiersze z lewej tabeli. Dla klientow bez zamowien kolumny z Zamowienia maja wartosc NULL. `WHERE Z.id IS NULL` wybiera wlasnie tych klientow.

Alternatywa z NOT IN:
```sql
SELECT imie, nazwisko
FROM Klienci
WHERE id_klienta NOT IN (SELECT id_klienta FROM Zamowienia);
```

### E. LEFT JOIN + COUNT — zliczanie wlaczajac zera

```sql
-- Kazdy rower i liczba wypozyczen (rowniez 0 dla niewypozyczonych)
SELECT R.model, COUNT(W.id_wypozyczenia) AS liczba_wypozyczen
FROM Rowery R
LEFT JOIN Wypozyczenia W ON R.id_roweru = W.id_roweru
GROUP BY R.model
ORDER BY liczba_wypozyczen DESC;
```

**PULAPKA COUNT(*) vs COUNT(kolumna)**:

| Zapytanie | Rower bez wypozyczen | Dlaczego |
|-----------|---------------------|----------|
| `COUNT(*)` | daje **1** | liczy wiersz z NULL-ami |
| `COUNT(W.id_wypozyczenia)` | daje **0** | pomija NULL-e |

**Zasada**: Z LEFT JOIN ZAWSZE uzywaj `COUNT(kolumna_z_prawej_tabeli)`.

---

## 4. GROUP BY + agregacje (36 pkt, 9/12 lat — NAJWAZNIEJSZY)

### A. COUNT + GROUP BY — zliczanie w grupach

```sql
-- Ile ksiazek w kazdym gatunku
SELECT gatunek, COUNT(*) AS liczba
FROM Ksiazki
GROUP BY gatunek
ORDER BY liczba DESC;
```

### B. Wiele funkcji agregujacych naraz

```sql
-- Dla kazdej kategorii: liczba produktow i srednia cena
SELECT kategoria, COUNT(*) AS liczba, ROUND(AVG(cena), 2) AS srednia_cena
FROM Produkty
GROUP BY kategoria
ORDER BY kategoria;
```

### C. HAVING — filtr PO grupowaniu

```sql
-- Klasy ze srednia ocen powyzej 4.0
SELECT K.nazwa_klasy, ROUND(AVG(U.srednia_ocen), 2) AS srednia_klasy
FROM Klasy K
JOIN Uczniowie U ON K.id_klasy = U.id_klasy
GROUP BY K.nazwa_klasy
HAVING AVG(U.srednia_ocen) > 4.0
ORDER BY srednia_klasy DESC;
```

**PULAPKA WHERE vs HAVING**:

| Klauzula | Kiedy filtruje | Przyklad |
|----------|---------------|----------|
| WHERE | PRZED grupowaniem (wiersze) | `WHERE rok > 2020` |
| HAVING | PO grupowaniu (grupy) | `HAVING COUNT(*) > 5` |

```sql
-- ZLE: WHERE AVG(ocena) > 4.0   (blad! AVG wymaga zgrupowanych danych)
-- DOBRZE: HAVING AVG(ocena) > 4.0
```

### D. SUM z JOIN + HAVING

```sql
-- Druzyny ze > 3 bramkami strzalonymi u siebie
SELECT D.nazwa, SUM(M.bramki_gosp) AS bramki_u_siebie
FROM Druzyny D
JOIN Mecze M ON D.id_druzyny = M.id_gospodarzy
GROUP BY D.nazwa
HAVING SUM(M.bramki_gosp) > 3
ORDER BY bramki_u_siebie DESC;
```

### E. Grupowanie po 2 kolumnach

```sql
-- Laczna kwota zamowien dla kazdego miasta i miesiaca
SELECT K.miasto, SUBSTR(Z.data_zamowienia, 6, 2) AS miesiac, SUM(Z.kwota) AS suma
FROM Klienci K
JOIN Zamowienia Z ON K.id_klienta = Z.id_klienta
GROUP BY K.miasto, SUBSTR(Z.data_zamowienia, 6, 2)
HAVING SUM(Z.kwota) > 2000
ORDER BY K.miasto, miesiac;
```

**Pulapka**: Grupowanie po jednej kolumnie zamiast po obu — wtedy np. Gdansk-styczeń i Gdansk-marzec zostana zlaczone w jedna grupe "Gdansk".

### F. COUNT(DISTINCT) — ile roznych wartosci

```sql
-- Ile roznych filmow oceniono w kazdym gatunku
SELECT F.gatunek, COUNT(DISTINCT F.id_filmu) AS liczba_filmow
FROM Filmy F
JOIN Oceny O ON F.id_filmu = O.id_filmu
GROUP BY F.gatunek;
```

**Pulapka**: `COUNT(*)` policzy liczbe OCEN (film z 3 ocenami = 3). `COUNT(DISTINCT F.id_filmu)` policzy liczbe FILMOW (film z 3 ocenami = 1).

### G. ROUND — zaokraglanie

```sql
ROUND(AVG(cena), 2)       -- srednia zaokraglona do 2 miejsc
ROUND(wartosc, 0)          -- zaokraglenie do liczby calkowitej
ROUND(COUNT(*) * 1.0 / n, 2) -- stosunek zaokraglony do 2 miejsc
```

---

## 5. Podzapytania (25 pkt, 8/12 lat)

### A. Podzapytanie skalarne — porownanie ze srednia

```sql
-- Pracownicy zarabiajacy wiecej niz srednia
SELECT imie, nazwisko, wynagrodzenie
FROM Pracownicy
WHERE wynagrodzenie > (SELECT AVG(wynagrodzenie) FROM Pracownicy);
```

**Pulapka**: `WHERE wynagrodzenie > AVG(wynagrodzenie)` NIE ZADZIALA. Funkcje agregujace nie moga byc uzywane bezposrednio w WHERE — trzeba podzapytania.

### B. WHERE ... IN (SELECT ...) — filtrowanie po zbiorze

```sql
-- Nazwy kursow, na ktore ktos sie zapisal
SELECT nazwa_kursu
FROM Kursy
WHERE id_kursu IN (SELECT id_kursu FROM Zapisy);
```

### C. WHERE ... NOT IN (SELECT ...) — wykluczanie zbioru

```sql
-- Kursy bez zapisow
SELECT nazwa_kursu
FROM Kursy
WHERE id_kursu NOT IN (SELECT id_kursu FROM Zapisy);
```

**PULAPKA NOT IN z NULL**: Jesli podzapytanie zwraca NULL wsrod wynikow, cale NOT IN zwraca pusty wynik! Bezpieczniejsza alternatywa: LEFT JOIN + IS NULL.

```sql
-- ZLE (gdy Zapisy.id_kursu moze byc NULL):
WHERE id_kursu NOT IN (SELECT id_kursu FROM Zapisy)
-- wynik: PUSTY (nic nie zwroci!)

-- DOBRZE (odporne na NULL):
SELECT K.nazwa_kursu
FROM Kursy K
LEFT JOIN Zapisy Z ON K.id_kursu = Z.id_kursu
WHERE Z.id_zapisu IS NULL;
```

### D. Podzapytanie w HAVING — srednia grupy vs srednia globalna

```sql
-- Gatunki filmow ze srednia ocena wyzsza niz globalna srednia
SELECT F.gatunek, ROUND(AVG(O.ocena), 2) AS srednia_ocena
FROM Filmy F
JOIN Oceny O ON F.id_filmu = O.id_filmu
GROUP BY F.gatunek
HAVING AVG(O.ocena) > (SELECT AVG(ocena) FROM Oceny);
```

### E. Laczenie podzapytania z warunkiem WHERE

```sql
-- Gracze z misji 'Smocza jaskinia' z wynikiem powyzej globalnej sredniej
SELECT G.nick
FROM Gracze G
JOIN Misje M ON G.id_gracza = M.id_gracza
WHERE M.nazwa_misji = 'Smocza jaskinia'
  AND M.punkty > (SELECT AVG(punkty) FROM Misje);
```

---

## 6. Funkcje i operacje dodatkowe

### A. CASE WHEN (matura 2017, 2023)

```sql
-- Klasyfikacja wynagrodzenia
SELECT imie, nazwisko, wynagrodzenie,
    CASE
        WHEN wynagrodzenie > 8000 THEN 'wysokie'
        WHEN wynagrodzenie > 5000 THEN 'srednie'
        ELSE 'niskie'
    END AS kategoria
FROM Pracownicy;
```

```sql
-- Zliczanie warunkowe (ile meczy wygranych/przegranych)
SELECT D.nazwa,
    SUM(CASE WHEN M.bramki_gosp > M.bramki_gosci THEN 1 ELSE 0 END) AS wygrane,
    SUM(CASE WHEN M.bramki_gosp < M.bramki_gosci THEN 1 ELSE 0 END) AS przegrane
FROM Druzyny D
JOIN Mecze M ON D.id_druzyny = M.id_gospodarzy
GROUP BY D.nazwa;
```

### B. Operacje na datach

```sql
-- Wyciaganie roku i miesiaca z daty w formacie 'RRRR-MM-DD'
SUBSTR(data, 1, 4)    -- rok:    '2024'
SUBSTR(data, 6, 2)    -- miesiac: '03'
SUBSTR(data, 9, 2)    -- dzien:  '15'

-- Porownywanie dat (format ISO pozwala porownywac jako tekst)
WHERE data > '2023-01-01'
WHERE data BETWEEN '2024-01-01' AND '2024-12-31'

-- Rok z daty (alternatywna skladnia)
WHERE SUBSTR(data_pomiaru, 1, 4) = '2060'
```

### C. DISTINCT — unikalne wartosci

```sql
-- Lista unikalnych producentow (matura 2025)
SELECT DISTINCT P.nazwa
FROM Producent P
JOIN Laziki L ON P.kod_producenta = L.kod_producenta
JOIN Pomiary M ON L.nr_lazika = M.nr_lazika
JOIN Obszary O ON M.kod_obszaru = O.kod_obszaru
WHERE O.nazwa_obszaru = 'Arcadia'
  AND SUBSTR(M.data_pomiaru, 1, 4) = '2060';
```

### D. ORDER BY + LIMIT

```sql
-- Najstarszy wyscig (pierwszy wynik posortowany wg daty)
SELECT g.Nazwa, g.Sezon
FROM GrandPrix g
ORDER BY g.Data ASC
LIMIT 1;

-- 3 najtansze produkty
SELECT nazwa, cena
FROM Produkty
ORDER BY cena ASC
LIMIT 3;
```

### E. Aliasy tabel i kolumn

```sql
-- Aliasy tabel (skroty): K, Z zamiast Klienci, Zamowienia
SELECT K.imie, Z.kwota
FROM Klienci K
JOIN Zamowienia Z ON K.id_klienta = Z.id_klienta;

-- Alias kolumny: AS nazwa_wynikowa
SELECT COUNT(*) AS liczba, AVG(cena) AS srednia_cena
FROM Produkty;
```

### F. Dzielenie calkowite → zmiennoprzecinkowe

```sql
-- ZLE: 3/2 = 1 (dzielenie calkowite!)
SELECT COUNT(*) / miejsca AS stosunek FROM tabela;

-- DOBRZE: mnozenie przez 1.0 wymusza dzielenie zmiennoprzecinkowe
SELECT COUNT(*) * 1.0 / miejsca AS stosunek FROM tabela;

-- Alternatywa: CAST
SELECT CAST(COUNT(*) AS REAL) / miejsca AS stosunek FROM tabela;

-- Przyklad z matury 2014 (chetni na miejsce w przedszkolu)
SELECT p.nazwa, ROUND(COUNT(pr.id_dziecka) * 1.0 / p.miejsca, 2) AS srednia
FROM przedszkola p
LEFT JOIN preferencje pr ON p.id = pr.id_przedszkola
GROUP BY p.id, p.nazwa, p.miejsca
ORDER BY srednia ASC
LIMIT 3;
```

---

## 7. Schematy baz danych z matur CKE

### 2014: Przedszkola (3 tabele, relacja wiele-do-wielu)

```
dzieci (id, imie, nazwisko, plec, dzielnica, rok_urodzenia)
przedszkola (id, nazwa, miejsca)
preferencje (id_dziecka, id_przedszkola, preferencja)
```

Relacja: dziecko ↔ przedszkole przez tabele preferencje. Typowe zadania: GROUP BY + COUNT, LEFT JOIN + ROUND do stosunku chetnych/miejsca.

### 2015: Formula 1 (4-5 tabel)

```
Zawodnicy (ID, Imie, Nazwisko, Kraj)
GrandPrix (ID, Nazwa, Data, Sezon, ID_Miejsca)
Wyniki_GrandPrix (ID_Zawodnika, ID_GrandPrix, Punkty)
Miejsca (ID, Nazwa)
Sezony (Sezon, ...)
```

Typowe zadania: LEFT JOIN + IS NULL (miejsca bez wyscigow), 4-table JOIN, COUNT DISTINCT, podzapytanie w HAVING.

### 2017: Pilka reczna (3-4 tabele)

```
Druzyny (id_druzyny, nazwa, miasto)
Mecze (id_meczu, id_gospodarzy, id_gosci, bramki_gosp, bramki_gosci, id_sedziego, rok)
Sedziowie (id_sedziego, nazwisko)
```

Typowe zadania: SUM z JOIN po id_gospodarzy vs id_gosci, CASE WHEN (wygrane/remisy/przegrane), LEFT JOIN (sedziowie bez meczy).

### 2019: Perfumy (3+ tabel)

```
Perfumy (id, nazwa, marka, rodzina_zapachow, cena)
Sklady (id_perfum, skladnik)
Sklepy (id, nazwa, ...)
```

Typowe zadania: JOIN 3 tabel z LIKE '%paczula%', GROUP BY z MIN(cena), obliczenie nowej ceny (cena * 0.85).

### 2023: Gry planszowe (3 tabele + dodatkowa)

```
gry (id_gry, nazwa, kategoria)
gracze (id_gracza, imie, nazwisko, wiek)
oceny (id_gry, id_gracza, stan, ocena)
sklep (id_gry, cena, promocja)  -- dodatkowa tabela w podzadaniu SQL
```

Typowe zadania: JOIN + WHERE + SUM (suma cen promocyjnych gier logicznych).

### 2024: Wykroczenia drogowe (3 tabele + dodatkowa)

```
kierowcy (id_kierowcy, imie, nazwisko, data_urodzenia, plec, miejscowosc)
taryfikator (id_wykroczenia, opis, punkty_karne, mandat)
rejestr (id_zdarzenia, id_kierowcy, id_wykroczenia, data_zdarzenia, predkosc)
fotoradar (id_fotoradaru, miejscowosc, dozwolona_predkosc)  -- dodatkowa
```

Typowe zadania: GROUP BY z 3 tabelami, podzapytania NOT IN (fotoradary bez wykroczen).

### 2025: Woda na Marsie (4 tabele)

```
Producent (kod_producenta, nazwa, kraj)
Laziki (nr_lazika, nazwa_lazika, rok_wyslania, wsp_ladowania, kod_producenta)
Pomiary (nr_lazika, data_pomiaru, kod_obszaru, wspolrzedne, glebokosc, ilosc)
Obszary (kod_obszaru, nazwa_obszaru)
```

Typowe zadania: JOIN 4 tabel z DISTINCT, SUBSTR na datach, filtrowanie po roku i nazwie obszaru.

---

## 8. Typowe pulapki SQL na maturze

| # | Pulapka | Bledny kod | Poprawny kod |
|---|---------|-----------|-------------|
| 1 | COUNT(*) z LEFT JOIN | `COUNT(*)` → daje 1 zamiast 0 | `COUNT(t2.id)` → daje 0 dla NULL |
| 2 | Porownanie z NULL | `WHERE opady != NULL` | `WHERE opady IS NOT NULL` |
| 3 | NOT IN z NULL w podzapytaniu | `NOT IN (SELECT kol...)` → pusty wynik | `LEFT JOIN + IS NULL` |
| 4 | Brak GROUP BY przy agregacji | `SELECT miasto, COUNT(*)` bez GROUP BY | Dodaj `GROUP BY miasto` |
| 5 | WHERE zamiast HAVING | `WHERE COUNT(*) > 5` | `HAVING COUNT(*) > 5` |
| 6 | Dzielenie calkowite | `3/2` → daje 1 | `3 * 1.0 / 2` → daje 1.5 |
| 7 | AVG w WHERE | `WHERE cena > AVG(cena)` | `WHERE cena > (SELECT AVG(cena) FROM ...)` |
| 8 | SUBSTR od 0 | `SUBSTR(tekst, 0, 3)` | `SUBSTR(tekst, 1, 3)` (od 1!) |
| 9 | AND/OR priorytet | `A AND B OR C` = `(A AND B) OR C` | Uzyj nawiasow: `A AND (B OR C)` |
| 10 | LIKE bez % | `WHERE nazwa LIKE 'A'` = dokladnie 'A' | `WHERE nazwa LIKE 'A%'` = zaczyna sie od A |

### Rozwiniety przyklad pulapki #1 (najczestszy blad):

```sql
-- ZADANIE: Pokaz kazdy rower i liczbe wypozyczen (0 dla niewypozyczonych)

-- ZLE:
SELECT R.model, COUNT(*) AS ile
FROM Rowery R
LEFT JOIN Wypozyczenia W ON R.id_roweru = W.id_roweru
GROUP BY R.model;
-- Wynik: Electric One → 1 (BLAD! powinno byc 0)

-- DOBRZE:
SELECT R.model, COUNT(W.id_wypozyczenia) AS ile
FROM Rowery R
LEFT JOIN Wypozyczenia W ON R.id_roweru = W.id_roweru
GROUP BY R.model;
-- Wynik: Electric One → 0 (POPRAWNIE)
```

### Rozwiniety przyklad pulapki #3 (NOT IN z NULL):

```sql
-- ZADANIE: Znajdz kierowcow bez wykroczen
-- Tabela rejestr: id_kierowcy moze byc NULL (np. niezidentyfikowany)

-- ZLE:
SELECT imie, nazwisko FROM kierowcy
WHERE id_kierowcy NOT IN (SELECT id_kierowcy FROM rejestr);
-- Jesli rejestr zawiera wiersz z id_kierowcy = NULL,
-- NOT IN zwraca pustą tablice (ZERO wynikow!)

-- DOBRZE:
SELECT K.imie, K.nazwisko
FROM kierowcy K
LEFT JOIN rejestr R ON K.id_kierowcy = R.id_kierowcy
WHERE R.id_zdarzenia IS NULL;
-- Zawsze dziala poprawnie, niezaleznie od NULL-i
```

---

## 9. Schemat rozwiazywania zadania SQL na maturze

1. **Przeczytaj schemat bazy** — zidentyfikuj tabele i klucze obce (strzalki/relacje)
2. **Okresl ile tabel potrzebujesz** — jesli dane sa w wielu tabelach, potrzebujesz JOIN
3. **Wybierz typ JOIN**:
   - Potrzebujesz WSZYSTKICH rekordow z jednej strony (nawet bez dopasowania) → LEFT JOIN
   - Potrzebujesz tylko dopasowanych → INNER JOIN (zwykly JOIN)
4. **Czy jest grupowanie?** — jesli w poleceniu jest "dla kazdego...", "ile...", "srednia..." → GROUP BY
5. **Czy jest warunek na grupe?** — "tylko grupy, gdzie..." → HAVING
6. **Napisz zapytanie w kolejnosci**: SELECT → FROM → JOIN → WHERE → GROUP BY → HAVING → ORDER BY → LIMIT
7. **Sprawdz pulapki**: COUNT(*) vs COUNT(kol), WHERE vs HAVING, dzielenie calkowite

### Import danych z pliku TXT do bazy

Na maturze dane sa w plikach `.txt` (TSV/CSV). Schemat importu:
1. Otworz program bazodanowy (np. DB Browser for SQLite, Access, MySQL Workbench)
2. Utworz tabele wg schematu z arkusza
3. Importuj dane: File → Import → Table from CSV/TXT file
4. Ustaw separator (tabulator lub przecinek) i kodowanie (UTF-8)
5. Sprawdz czy dane wczytaly sie poprawnie (kilka pierwszych wierszy)
