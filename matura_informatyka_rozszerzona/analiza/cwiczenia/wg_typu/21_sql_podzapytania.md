# 21. SQL — Podzapytania (subqueries)

Typ zadania: **sql_podzapytania**
Czestotliwosc: 8/12 lat | Laczna punktacja: 25 pkt
Kategoria: SQL

## Umiejetnosci cwiczone w tym zestawie

`podzapytanie-skalarne` `NOT-IN` `IN` `EXISTS` `podzapytanie-w-WHERE` `podzapytanie-w-HAVING` `LEFT-JOIN-IS-NULL` `AVG` `MAX` `MIN` `korelacja-podzapytan` `COUNT-DISTINCT`

---

### Cwiczenie 21.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2019 (Perfumy)
**Tagi**: `podzapytanie-skalarne` `AVG` `podzapytanie-w-WHERE`

Tabela **Pracownicy** zawiera dane o pracownikach firmy:

| id | imie | nazwisko | dzial | wynagrodzenie |
|----|------|----------|-------|--------------|
| 1 | Anna | Kowalska | IT | 8500 |
| 2 | Jan | Nowak | HR | 5200 |
| 3 | Ewa | Maj | IT | 9200 |
| 4 | Piotr | Lis | Sprzedaz | 6100 |
| 5 | Kasia | Zak | HR | 5800 |
| 6 | Tomek | Bak | IT | 7800 |
| 7 | Ola | Wrobel | Sprzedaz | 7200 |
| 8 | Marek | Krol | Sprzedaz | 5500 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli imie i nazwisko pracownikow, ktorych wynagrodzenie jest wyzsze niz srednie wynagrodzenie wszystkich pracownikow.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Nie mozesz uzyc AVG bezposrednio w WHERE — potrzebujesz podzapytania.
2. **Podejscie**: Podzapytanie `(SELECT AVG(wynagrodzenie) FROM Pracownicy)` zwraca jedna wartosc (skalarne).
3. **Kluczowy krok**: Porownaj wynagrodzenie kazdego pracownika z ta wartoscia: `WHERE wynagrodzenie > (SELECT ...)`.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT imie, nazwisko, wynagrodzenie
FROM Pracownicy
WHERE wynagrodzenie > (SELECT AVG(wynagrodzenie) FROM Pracownicy);
```

**Weryfikacja**:
- Srednia wynagrodzen: (8500+5200+9200+6100+5800+7800+7200+5500)/8 = 55300/8 = **6912.50**
- Pracownicy z wynagrodzeniem > 6912.50:
  - Anna Kowalska (8500 ✓)
  - Ewa Maj (9200 ✓)
  - Tomek Bak (7800 ✓)
  - Ola Wrobel (7200 ✓)

| imie | nazwisko | wynagrodzenie |
|------|----------|--------------|
| Anna | Kowalska | 8500 |
| Ewa | Maj | 9200 |
| Tomek | Bak | 7800 |
| Ola | Wrobel | 7200 |

**Wyjasnienie**: Podzapytanie `(SELECT AVG(wynagrodzenie) FROM Pracownicy)` zwraca pojedyncza wartosc (6912.50), ktora jest uzywana w warunku WHERE. To tzw. podzapytanie skalarne — zwraca dokladnie jedna wartosc.
</details>

<details>
<summary>Typowe bledy</summary>

- **AVG bezposrednio w WHERE**: `WHERE wynagrodzenie > AVG(wynagrodzenie)` — blad skladni, funkcje agregujace nie moga byc uzywane w WHERE. CKE: -2 pkt.
- **Brak nawiasow wokol podzapytania**: Podzapytanie musi byc w nawiasach. CKE: -1 pkt (blad skladni).

</details>

---

### Cwiczenie 21.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2021 (Gra strategiczna)
**Tagi**: `NOT-IN` `podzapytanie-w-WHERE`

Tabela **Kursy**:

| id_kursu | nazwa_kursu | kategoria |
|----------|-------------|-----------|
| 1 | Python basics | programowanie |
| 2 | SQL dla poczatkujacych | bazy_danych |
| 3 | Excel zaawansowany | arkusze |
| 4 | HTML i CSS | webdev |
| 5 | Algorytmy | programowanie |

Tabela **Zapisy**:

| id_zapisu | id_kursu | id_ucznia | data_zapisu |
|-----------|----------|-----------|-------------|
| 1 | 1 | 101 | 2024-01-10 |
| 2 | 1 | 102 | 2024-01-12 |
| 3 | 2 | 103 | 2024-01-15 |
| 4 | 3 | 101 | 2024-02-01 |
| 5 | 5 | 104 | 2024-02-10 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwy kursow, na ktore nikt sie nie zapisal.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Szukasz kursow, ktorych id NIE wystepuje w tabeli Zapisy.
2. **Podejscie**: Podzapytanie zwraca zbior id_kursu z Zapisy, a NOT IN wyklucza je.
3. **Kluczowy krok**: `WHERE id_kursu NOT IN (SELECT id_kursu FROM Zapisy)`.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT nazwa_kursu
FROM Kursy
WHERE id_kursu NOT IN (SELECT id_kursu FROM Zapisy);
```

**Weryfikacja**:
- Kursy z zapisami (id_kursu w tabeli Zapisy): 1, 2, 3, 5
- Kursy bez zapisow: id=4 (HTML i CSS)

| nazwa_kursu |
|-------------|
| HTML i CSS |

**Rozwiazanie alternatywne z LEFT JOIN:**
```sql
SELECT K.nazwa_kursu
FROM Kursy K
LEFT JOIN Zapisy Z ON K.id_kursu = Z.id_kursu
WHERE Z.id_zapisu IS NULL;
```

**Wyjasnienie**: Podzapytanie `(SELECT id_kursu FROM Zapisy)` zwraca zbior {1, 2, 3, 5}. NOT IN sprawdza, ktore id_kursu z tabeli Kursy NIE naleza do tego zbioru.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak NOT**: Samo `IN` daloby kursy z zapisami, czyli odwrotnosc oczekiwanego wyniku. CKE: -2 pkt.
- **NOT IN z NULL**: Jesli podzapytanie zwraca NULL wsrod wynikow, NOT IN zwraca pusty zbior. Np. gdyby Zapisy mialy wiersz z id_kursu = NULL, wynik bylby pusty. CKE: nie dotyczy tego zadania, ale warto pamietac.

</details>

---

### Cwiczenie 21.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 (Rejestr wykroczen)
**Tagi**: `LEFT-JOIN-IS-NULL` `NOT-IN` `podzapytanie-w-WHERE`

Tabela **Klienci**:

| id_klienta | imie | nazwisko | miasto |
|------------|------|----------|--------|
| 1 | Anna | Kowalska | Warszawa |
| 2 | Jan | Nowak | Krakow |
| 3 | Ewa | Maj | Gdansk |
| 4 | Piotr | Lis | Warszawa |
| 5 | Kasia | Zak | Poznan |
| 6 | Tomek | Bak | Krakow |

Tabela **Zamowienia**:

| id_zamowienia | id_klienta | data | kwota |
|---------------|------------|------|-------|
| 1 | 1 | 2024-01-10 | 250 |
| 2 | 2 | 2024-01-15 | 180 |
| 3 | 1 | 2024-02-20 | 340 |
| 4 | 3 | 2024-03-05 | 120 |
| 5 | 6 | 2024-03-10 | 95 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli imie i nazwisko klientow, ktorzy nigdy nie zlozyli zamowienia. Uzyj LEFT JOIN.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: LEFT JOIN zachowuje wszystkie wiersze z lewej tabeli, nawet bez dopasowania.
2. **Podejscie**: LEFT JOIN Klienci z Zamowienia, potem szukaj wierszy, gdzie Zamowienia daje NULL.
3. **Kluczowy krok**: `WHERE Z.id_zamowienia IS NULL` — to wylapuje klientow bez zamowien.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT K.imie, K.nazwisko
FROM Klienci K
LEFT JOIN Zamowienia Z ON K.id_klienta = Z.id_klienta
WHERE Z.id_zamowienia IS NULL;
```

**Weryfikacja**:
- Klienci z zamowieniami: Anna (id=1), Jan (id=2), Ewa (id=3), Tomek (id=6)
- Klienci BEZ zamowien: Piotr (id=4), Kasia (id=5)

| imie | nazwisko |
|------|----------|
| Piotr | Lis |
| Kasia | Zak |

**Rozwiazanie alternatywne z NOT IN:**
```sql
SELECT imie, nazwisko
FROM Klienci
WHERE id_klienta NOT IN (SELECT id_klienta FROM Zamowienia);
```

**Wyjasnienie**: LEFT JOIN zachowuje WSZYSTKIE wiersze z lewej tabeli (Klienci), nawet jezeli nie maja dopasowania w prawej tabeli (Zamowienia). Dla niedopasowanych wierszy kolumny z Zamowienia maja wartosc NULL.
</details>

<details>
<summary>Typowe bledy</summary>

- **INNER JOIN zamiast LEFT JOIN**: Odrzuca wiersze bez dopasowania — nigdy nie pokaze klientow bez zamowien. CKE: -3 pkt.
- **IS NULL na zlej kolumnie**: `WHERE Z.kwota IS NULL` tez zadziala, ale `WHERE Z.id_klienta IS NULL` jest bezpieczniejsze. CKE: akceptowane.
- **= NULL zamiast IS NULL**: `WHERE Z.id_zamowienia = NULL` zawsze daje FALSE. CKE: -2 pkt.

</details>

---

### Cwiczenie 21.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 (Gra strategiczna)
**Tagi**: `podzapytanie-skalarne` `AVG` `podzapytanie-w-WHERE` `IN`

Tabela **Gracze**:

| id_gracza | nick | poziom | frakcja |
|-----------|------|--------|---------|
| 1 | DragonSlayer | 45 | Ork |
| 2 | SilverMage | 38 | Elf |
| 3 | IronFist | 52 | Ork |
| 4 | StarLight | 41 | Elf |
| 5 | DarkKnight | 47 | Czlowiek |
| 6 | FireBlade | 33 | Ork |
| 7 | MoonArcher | 50 | Elf |
| 8 | StoneGuard | 29 | Czlowiek |

Tabela **Misje**:

| id_misji | id_gracza | nazwa_misji | punkty |
|----------|-----------|-------------|--------|
| 1 | 1 | Smocza jaskinia | 120 |
| 2 | 1 | Mroczny las | 85 |
| 3 | 3 | Smocza jaskinia | 150 |
| 4 | 5 | Mroczny las | 90 |
| 5 | 7 | Smocza jaskinia | 140 |
| 6 | 2 | Mroczny las | 75 |
| 7 | 3 | Lodowa forteca | 200 |
| 8 | 5 | Lodowa forteca | 180 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nicki graczy, ktorzy ukonczyli misje o nazwie 'Smocza jaskinia' i zdobyli w niej wiecej punktow niz srednia punktow ze WSZYSTKICH misji.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dwa warunki w WHERE — jeden na nazwe misji, drugi porownuje z podzapytaniem.
2. **Podejscie**: JOIN + WHERE z nazwa_misji + porownanie punktow ze srednia globalna.
3. **Kluczowy krok**: Podzapytanie `(SELECT AVG(punkty) FROM Misje)` oblicza srednia z CALEJ tabeli, nie tylko z jednej misji.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT G.nick
FROM Gracze G
JOIN Misje M ON G.id_gracza = M.id_gracza
WHERE M.nazwa_misji = 'Smocza jaskinia'
  AND M.punkty > (SELECT AVG(punkty) FROM Misje);
```

**Weryfikacja**:
- Srednia punktow ze wszystkich misji: (120+85+150+90+140+75+200+180)/8 = 1040/8 = **130.0**
- Gracze z misja 'Smocza jaskinia':
  - DragonSlayer: 120 punktow (>130? ✗)
  - IronFist: 150 punktow (>130? ✓)
  - MoonArcher: 140 punktow (>130? ✓)

| nick |
|------|
| IronFist |
| MoonArcher |

**Wyjasnienie**: Zapytanie laczy dwa warunki w WHERE: filtr na nazwe misji i porownanie z wynikiem podzapytania skalarnego. Podzapytanie oblicza srednia z CALEJ tabeli Misje, nie tylko z misji 'Smocza jaskinia'.
</details>

<details>
<summary>Typowe bledy</summary>

- **Ograniczenie podzapytania**: `(SELECT AVG(punkty) FROM Misje WHERE nazwa_misji = 'Smocza jaskinia')` daje srednia = (120+150+140)/3 = 136.67 — inny wynik niz globalna srednia 130. CKE: -2 pkt.
- **Brak JOIN**: Bez polaczenia tabel nie uzyskamy nickow graczy. CKE: -2 pkt.
- **OR zamiast AND**: `WHERE nazwa_misji = '...' OR punkty > (...)` daje tez graczy z innych misji. CKE: -2 pkt.

</details>

---

### Cwiczenie 21.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2019 (Perfumy)
**Tagi**: `podzapytanie-w-HAVING` `COUNT-DISTINCT` `AVG` `ROUND` `IN`

Tabela **Filmy**:

| id_filmu | tytul | rok | gatunek |
|----------|-------|-----|---------|
| 1 | Zielona mila | 1999 | dramat |
| 2 | Gladiator | 2000 | akcja |
| 3 | Incepcja | 2010 | sci-fi |
| 4 | Interstellar | 2014 | sci-fi |
| 5 | Joker | 2019 | dramat |
| 6 | Avengers | 2012 | akcja |
| 7 | Matrix | 1999 | sci-fi |
| 8 | Titanic | 1997 | dramat |

Tabela **Oceny**:

| id_oceny | id_filmu | uzytkownik | ocena |
|----------|----------|------------|-------|
| 1 | 1 | user_A | 9 |
| 2 | 1 | user_B | 8 |
| 3 | 2 | user_A | 7 |
| 4 | 3 | user_C | 10 |
| 5 | 3 | user_A | 9 |
| 6 | 4 | user_B | 10 |
| 7 | 4 | user_C | 9 |
| 8 | 5 | user_A | 8 |
| 9 | 6 | user_B | 6 |
| 10 | 7 | user_C | 9 |
| 11 | 7 | user_A | 8 |
| 12 | 8 | user_A | 7 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli gatunki filmow, w ktorych srednia ocena jest wyzsza niz ogolna srednia ocena ze wszystkich filmow. Dla kazdego takiego gatunku pokaz srednia ocene zaokraglona do 2 miejsc i liczbe ocenionych filmow.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: GROUP BY po gatunku, filtrowanie grup z HAVING i podzapytaniem.
2. **Podejscie**: JOIN Filmy z Ocenami, GROUP BY gatunek, HAVING AVG(ocena) > (SELECT AVG(ocena) FROM Oceny).
3. **Kluczowy krok**: COUNT(DISTINCT id_filmu) liczy unikalne filmy, nie unikalne oceny.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT F.gatunek,
       ROUND(AVG(O.ocena), 2) AS srednia_ocena,
       COUNT(DISTINCT F.id_filmu) AS liczba_filmow
FROM Filmy F
JOIN Oceny O ON F.id_filmu = O.id_filmu
GROUP BY F.gatunek
HAVING AVG(O.ocena) > (SELECT AVG(ocena) FROM Oceny);
```

**Weryfikacja**:

Ogolna srednia ocen: (9+8+7+10+9+10+9+8+6+9+8+7)/12 = 100/12 = **8.33**

Srednia wg gatunku:
- dramat: filmy 1,5,8 → oceny: 9,8,8,7 → srednia = 32/4 = **8.00** (>8.33? ✗)
- akcja: filmy 2,6 → oceny: 7,6 → srednia = 13/2 = **6.50** (>8.33? ✗)
- sci-fi: filmy 3,4,7 → oceny: 10,9,10,9,9,8 → srednia = 55/6 = **9.17** (>8.33? ✓)

| gatunek | srednia_ocena | liczba_filmow |
|---------|---------------|---------------|
| sci-fi | 9.17 | 3 |

**Wyjasnienie**: Podzapytanie w klauzuli HAVING pozwala porownac wynik agregacji grupy ze srednia globalna. COUNT(DISTINCT F.id_filmu) liczy unikalne filmy, a nie unikalne oceny.
</details>

<details>
<summary>Typowe bledy</summary>

- **COUNT(*) zamiast COUNT(DISTINCT id_filmu)**: Policzyloby 6 ocen zamiast 3 filmow. CKE: -1 pkt.
- **Podzapytanie w WHERE zamiast HAVING**: `WHERE AVG(O.ocena) > (...)` — blad skladni. CKE: -2 pkt.
- **Podzapytanie ze srednia per gatunek**: Gdyby podzapytanie filtrowalo po gatunku, porownywaloby grupe z nia sama. CKE: -2 pkt.

</details>

---

### Cwiczenie 21.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2016 (Uniwersytet)
**Tagi**: `podzapytanie-skalarne` `MAX` `podzapytanie-w-WHERE`

Tabela **Produkty**:

| id | nazwa | kategoria | cena | na_stanie |
|----|-------|-----------|------|-----------|
| 1 | Laptop | elektronika | 3500 | 12 |
| 2 | Mysz | akcesoria | 80 | 150 |
| 3 | Monitor | elektronika | 1200 | 30 |
| 4 | Klawiatura | akcesoria | 250 | 85 |
| 5 | Tablet | elektronika | 2800 | 20 |
| 6 | Sluchawki | akcesoria | 350 | 60 |
| 7 | Drukarka | elektronika | 900 | 15 |
| 8 | Pendrive | akcesoria | 40 | 200 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwe i cene najdrozszego produktu. Uzyj podzapytania z MAX.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Nie mozesz uzyc MAX bezposrednio z innymi kolumnami bez GROUP BY.
2. **Podejscie**: Podzapytanie `(SELECT MAX(cena) FROM Produkty)` zwraca najwyzsza cene.
3. **Kluczowy krok**: `WHERE cena = (SELECT MAX(cena) FROM Produkty)`.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT nazwa, cena
FROM Produkty
WHERE cena = (SELECT MAX(cena) FROM Produkty);
```

**Weryfikacja**:
- MAX(cena) = 3500 (Laptop)
- Produkty z cena = 3500: Laptop

| nazwa | cena |
|-------|------|
| Laptop | 3500 |

**Wyjasnienie**: Podzapytanie skalarne zwraca jedna wartosc (3500). WHERE porownuje cene kazdego produktu z ta wartoscia. Gdyby kilka produktow mialo te sama maksymalna cene, wszystkie bylyby wyswietlone.
</details>

<details>
<summary>Typowe bledy</summary>

- **SELECT nazwa, MAX(cena)**: Bez GROUP BY — SQLite wybierze losowa nazwe, inne bazy zglasza blad. CKE: -1 pkt.
- **ORDER BY cena DESC LIMIT 1**: Poprawna alternatywa, ale LIMIT nie jest czescia standardu SQL i na maturze podzapytanie jest bezpieczniejsze. CKE: akceptowane w SQLite.

</details>

---

### Cwiczenie 21.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2022 (Ewidencja uczniow)
**Tagi**: `IN` `podzapytanie-w-WHERE` `NOT-IN`

Tabela **Pracownicy**:

| id | imie | nazwisko | dzial | stanowisko |
|----|------|----------|-------|------------|
| 1 | Anna | Maj | IT | programista |
| 2 | Jan | Krol | IT | tester |
| 3 | Ewa | Lis | HR | rekruter |
| 4 | Piotr | Zak | IT | programista |
| 5 | Kasia | Wrobel | HR | kierownik |
| 6 | Tomek | Bak | Finanse | analityk |

Tabela **Szkolenia**:

| id | id_pracownika | nazwa_szkolenia | data |
|----|---------------|-----------------|------|
| 1 | 1 | SQL zaawansowany | 2024-03-10 |
| 2 | 2 | Testowanie | 2024-03-15 |
| 3 | 1 | Python | 2024-04-01 |
| 4 | 4 | SQL zaawansowany | 2024-03-10 |
| 5 | 6 | Excel | 2024-04-05 |

**Polecenie**:
1. Wyswietl imiona i nazwiska pracownikow, ktorzy ukonczyli szkolenie 'SQL zaawansowany'.
2. Wyswietl imiona i nazwiska pracownikow z dzialu 'IT', ktorzy NIE ukonczyli zadnego szkolenia.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dwa osobne zapytania — jedno z IN, drugie z NOT IN.
2. **Podejscie**: Podzapytanie zwraca zbior id_pracownika spelniajacych warunek.
3. **Kluczowy krok**: W zapytaniu 2 laczysz warunek na dzial (WHERE) z NOT IN.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytania SQL:**
```sql
-- 1. Pracownicy ze szkoleniem SQL zaawansowany
SELECT imie, nazwisko
FROM Pracownicy
WHERE id IN (SELECT id_pracownika FROM Szkolenia WHERE nazwa_szkolenia = 'SQL zaawansowany');

-- 2. Pracownicy IT bez zadnego szkolenia
SELECT imie, nazwisko
FROM Pracownicy
WHERE dzial = 'IT'
  AND id NOT IN (SELECT id_pracownika FROM Szkolenia);
```

**Weryfikacja**:

1. Szkolenie 'SQL zaawansowany' ukonczyli: id=1 (Anna), id=4 (Piotr)

| imie | nazwisko |
|------|----------|
| Anna | Maj |
| Piotr | Zak |

2. Pracownicy IT: id=1 (Anna), id=2 (Jan), id=4 (Piotr). Z nich w Szkoleniach: id=1, id=2, id=4 — wszyscy maja szkolenia. Brak wynikow.

| imie | nazwisko |
|------|----------|
| (pusty wynik) | |

**Wyjasnienie**: IN z podzapytaniem sprawdza przynaleznosc do zbioru. NOT IN wyklucza elementy zbioru. Mozna laczyc NOT IN z innymi warunkami w WHERE za pomoca AND.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak WHERE w podzapytaniu (zad. 1)**: `SELECT id_pracownika FROM Szkolenia` zwraca wszystkich — bez filtrowania po nazwie szkolenia. CKE: -2 pkt.
- **OR zamiast AND (zad. 2)**: `WHERE dzial = 'IT' OR id NOT IN (...)` daje pracownikow IT PLUS pracownikow bez szkolen z innych dzialow. CKE: -2 pkt.

</details>

---

### Cwiczenie 21.8 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2017 (Pilka reczna)
**Tagi**: `podzapytanie-skalarne` `MIN` `MAX` `podzapytanie-w-WHERE`

Tabela **Zawodnicy**:

| id | imie | nazwisko | druzyna | wiek | wzrost |
|----|------|----------|---------|------|--------|
| 1 | Adam | Kowalski | Orly | 25 | 185 |
| 2 | Jan | Nowak | Wilki | 30 | 192 |
| 3 | Ewa | Maj | Orly | 22 | 175 |
| 4 | Piotr | Lis | Rysie | 28 | 188 |
| 5 | Kasia | Zak | Wilki | 24 | 170 |
| 6 | Tomek | Bak | Rysie | 32 | 195 |
| 7 | Ola | Wrobel | Orly | 26 | 178 |
| 8 | Marek | Krol | Wilki | 21 | 180 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli imie, nazwisko i druzyne zawodnikow, ktorych wzrost jest wiekszy niz najwyzszy wzrost wsrod zawodnikow druzyny 'Wilki'.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Podzapytanie z MAX ograniczone do jednej druzyny.
2. **Podejscie**: `(SELECT MAX(wzrost) FROM Zawodnicy WHERE druzyna = 'Wilki')` daje najwyzszy wzrost Wilkow.
3. **Kluczowy krok**: Porownaj wzrost kazdego zawodnika z tym wynikiem.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT imie, nazwisko, druzyna
FROM Zawodnicy
WHERE wzrost > (SELECT MAX(wzrost) FROM Zawodnicy WHERE druzyna = 'Wilki');
```

**Weryfikacja**:
- MAX wzrost Wilkow: Jan (192), Kasia (170), Marek (180) → MAX = **192**
- Zawodnicy z wzrostem > 192:
  - Tomek Bak: 195 > 192 ✓

| imie | nazwisko | druzyna |
|------|----------|---------|
| Tomek | Bak | Rysie |

**Wyjasnienie**: Podzapytanie z WHERE ogranicza MAX do jednej druzyny. Wynik (192) jest porownywany ze wzrostem WSZYSTKICH zawodnikow (nie tylko Wilkow). Gdybys chcial wykluczyc samych Wilkow z wynikow, dodalbys `AND druzyna != 'Wilki'` w zapytaniu zewnetrznym.
</details>

<details>
<summary>Typowe bledy</summary>

- **MAX bez WHERE**: `(SELECT MAX(wzrost) FROM Zawodnicy)` daje MAX ze wszystkich — wynik bylby pusty (nikt nie jest wyzszy niz najwyzszy). CKE: -2 pkt.
- **`>=` zamiast `>`**: "Wiekszy niz" to `>`, nie `>=`. Z `>=` Jan Nowak (192) tez bylby w wynikach. CKE: -1 pkt.

</details>

---

### Cwiczenie 21.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2025 (Woda na Marsie)
**Tagi**: `korelacja-podzapytan` `podzapytanie-skalarne` `AVG` `podzapytanie-w-WHERE`

Tabela **Uczniowie**:

| id | imie | nazwisko | klasa |
|----|------|----------|-------|
| 1 | Anna | Maj | 3A |
| 2 | Jan | Krol | 3A |
| 3 | Ewa | Lis | 3B |
| 4 | Piotr | Zak | 3B |
| 5 | Kasia | Wrobel | 3A |
| 6 | Tomek | Bak | 3B |

Tabela **Wyniki**:

| id | id_ucznia | przedmiot | punkty |
|----|-----------|-----------|--------|
| 1 | 1 | Matematyka | 85 |
| 2 | 2 | Matematyka | 70 |
| 3 | 3 | Matematyka | 92 |
| 4 | 4 | Matematyka | 60 |
| 5 | 5 | Matematyka | 78 |
| 6 | 6 | Matematyka | 88 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli imie, nazwisko, klase i punkty uczniow, ktorych wynik z Matematyki jest wyzszy niz srednia wynikow z Matematyki W ICH KLASIE (nie w calej szkole). Uzyj podzapytania skorelowanego.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Podzapytanie musi odwolywac sie do klasy ucznia z zapytania zewnetrznego — to podzapytanie skorelowane.
2. **Podejscie**: `WHERE W.punkty > (SELECT AVG(W2.punkty) FROM Wyniki W2 JOIN Uczniowie U2 ON ... WHERE U2.klasa = U.klasa)`.
3. **Kluczowy krok**: Alias w podzapytaniu (U2, W2) musi byc INNY niz w zapytaniu glownym (U, W), ale odwolanie `U2.klasa = U.klasa` laczy podzapytanie z zapytaniem zewnetrznym.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT U.imie, U.nazwisko, U.klasa, W.punkty
FROM Uczniowie U
JOIN Wyniki W ON U.id = W.id_ucznia
WHERE W.przedmiot = 'Matematyka'
  AND W.punkty > (
    SELECT AVG(W2.punkty)
    FROM Wyniki W2
    JOIN Uczniowie U2 ON U2.id = W2.id_ucznia
    WHERE W2.przedmiot = 'Matematyka'
      AND U2.klasa = U.klasa
  );
```

**Weryfikacja**:

Srednia 3A: Anna (85), Jan (70), Kasia (78) → (85+70+78)/3 = **77.67**
Srednia 3B: Ewa (92), Piotr (60), Tomek (88) → (92+60+88)/3 = **80.00**

- Anna: 85 > 77.67 ✓
- Jan: 70 > 77.67 ✗
- Ewa: 92 > 80.00 ✓
- Piotr: 60 > 80.00 ✗
- Kasia: 78 > 77.67 ✓
- Tomek: 88 > 80.00 ✓

| imie | nazwisko | klasa | punkty |
|------|----------|-------|--------|
| Anna | Maj | 3A | 85 |
| Ewa | Lis | 3B | 92 |
| Kasia | Wrobel | 3A | 78 |
| Tomek | Bak | 3B | 88 |

**Wyjasnienie**: Podzapytanie skorelowane jest wykonywane dla KAZDEGO wiersza z zapytania zewnetrznego. Odwolanie `U2.klasa = U.klasa` powoduje, ze srednia jest obliczana osobno dla kazdej klasy. To rozni sie od podzapytania nieskorelowanego, ktore jest wykonywane raz.
</details>

<details>
<summary>Typowe bledy</summary>

- **Nieskorelowane podzapytanie**: `(SELECT AVG(punkty) FROM Wyniki)` oblicza srednia globalna zamiast per klasa. CKE: -3 pkt.
- **Te same aliasy**: `U` w podzapytaniu i w zapytaniu glownym — niejednoznacznosc. CKE: -1 pkt (blad skladni).
- **Brak filtra na przedmiot w podzapytaniu**: Gdyby tabela Wyniki miala tez inne przedmioty, srednia bylaby bledna. CKE: -1 pkt.

</details>

---

### Cwiczenie 21.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2023 (Gry planszowe)
**Tagi**: `podzapytanie-w-HAVING` `podzapytanie-skalarne` `COUNT-DISTINCT` `AVG` `ROUND` `IN`

Tabela **Biblioteki**:

| id_bib | nazwa | miasto |
|--------|-------|--------|
| 1 | Miejska | Warszawa |
| 2 | Akademicka | Krakow |
| 3 | Dzielnicowa | Warszawa |
| 4 | Uczelniana | Gdansk |

Tabela **Ksiazki**:

| id_ks | tytul | gatunek | id_bib |
|-------|-------|---------|--------|
| 1 | Pan Tadeusz | poezja | 1 |
| 2 | Lalka | proza | 1 |
| 3 | Quo Vadis | proza | 2 |
| 4 | Potop | proza | 1 |
| 5 | Ferdydurke | proza | 3 |
| 6 | Dziady | poezja | 2 |
| 7 | Zbrodnia i kara | proza | 4 |
| 8 | Trans-Atlantyk | proza | 3 |

Tabela **Wypozyczenia**:

| id_wyp | id_ks | czytelnik | data_wyp | data_zwrotu |
|--------|-------|-----------|----------|-------------|
| 1 | 1 | Jan | 2024-01-10 | 2024-01-25 |
| 2 | 2 | Anna | 2024-01-15 | 2024-02-01 |
| 3 | 3 | Ewa | 2024-02-01 | 2024-02-15 |
| 4 | 1 | Piotr | 2024-02-10 | 2024-02-28 |
| 5 | 5 | Anna | 2024-03-01 | 2024-03-20 |
| 6 | 4 | Jan | 2024-03-05 | NULL |
| 7 | 6 | Ewa | 2024-03-10 | 2024-03-25 |
| 8 | 2 | Kasia | 2024-03-15 | NULL |
| 9 | 7 | Tomek | 2024-04-01 | 2024-04-10 |
| 10 | 8 | Anna | 2024-04-05 | NULL |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwe i miasto bibliotek, w ktorych liczba wypozyczen (zakonczonych — data_zwrotu IS NOT NULL) jest wieksza niz srednia liczba zakonczonych wypozyczen na biblioteke. Dla kazdej takiej biblioteki podaj tez te liczbe.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Potrzebujesz: (a) policzyc zakonczone wypozyczenia per biblioteka, (b) obliczyc srednia z tych liczb, (c) porownac.
2. **Podejscie**: JOIN trzech tabel, WHERE data_zwrotu IS NOT NULL, GROUP BY po bibliotece, HAVING z podzapytaniem.
3. **Kluczowy krok**: Podzapytanie w HAVING musi obliczac srednia liczbe wypozyczen na biblioteke — to wymaga zagniezdzonego podzapytania ze swoim GROUP BY.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT B.nazwa, B.miasto, COUNT(*) AS liczba_wypozyczen
FROM Biblioteki B
JOIN Ksiazki K ON B.id_bib = K.id_bib
JOIN Wypozyczenia W ON K.id_ks = W.id_ks
WHERE W.data_zwrotu IS NOT NULL
GROUP BY B.id_bib, B.nazwa, B.miasto
HAVING COUNT(*) > (
  SELECT AVG(cnt) FROM (
    SELECT COUNT(*) AS cnt
    FROM Ksiazki K2
    JOIN Wypozyczenia W2 ON K2.id_ks = W2.id_ks
    WHERE W2.data_zwrotu IS NOT NULL
    GROUP BY K2.id_bib
  )
);
```

**Weryfikacja**:

Zakonczone wypozyczenia per biblioteka:
- Miejska (id=1): ks.1 (wyp.1 ✓, wyp.4 ✓), ks.2 (wyp.2 ✓, wyp.8 ✗ NULL), ks.4 (wyp.6 ✗ NULL) → **3**
- Akademicka (id=2): ks.3 (wyp.3 ✓), ks.6 (wyp.7 ✓) → **2**
- Dzielnicowa (id=3): ks.5 (wyp.5 ✓), ks.8 (wyp.10 ✗ NULL) → **1**
- Uczelniana (id=4): ks.7 (wyp.9 ✓) → **1**

Srednia: (3+2+1+1)/4 = 7/4 = **1.75**

Biblioteki z liczba > 1.75:
- Miejska: 3 > 1.75 ✓
- Akademicka: 2 > 1.75 ✓

| nazwa | miasto | liczba_wypozyczen |
|-------|--------|-------------------|
| Miejska | Warszawa | 3 |
| Akademicka | Krakow | 2 |

**Wyjasnienie**: Zagniezdzone podzapytanie (podzapytanie w podzapytaniu) najpierw oblicza COUNT per biblioteka, potem AVG z tych COUNT-ow. Wewnetrzne podzapytanie zwraca tabele (cnt per biblioteka), a zewnetrzne oblicza srednia z tych wartosci. To jeden z najtrudniejszych wzorcow SQL na maturze.
</details>

<details>
<summary>Typowe bledy</summary>

- **AVG(COUNT(*))**: Nie mozna zagniezdzyc funkcji agregujacych bezposrednio — `AVG(COUNT(*))` to blad. Potrzeba podzapytania. CKE: -2 pkt.
- **Brak WHERE data_zwrotu IS NOT NULL w podzapytaniu**: Podzapytanie liczyloby tez niezakonczone wypozyczenia. CKE: -1 pkt.
- **Brak filtra na data_zwrotu w zapytaniu glownym**: Policzyloby wszystkie wypozyczenia (zakonczenia i niezakonczone). CKE: -2 pkt.

</details>

---

## Samoocena

| Poziom | Opis | Wymaganie |
|--------|------|-----------|
| Podstawowy | Podzapytanie skalarne z AVG/MAX, NOT IN | 1-3 cwiczen bez pomocy |
| Dobry | LEFT JOIN IS NULL, podzapytanie z wieloma warunkami | 4-6 cwiczen bez pomocy |
| Bardzo dobry | Podzapytanie skorelowane, HAVING z podzapytaniem | 7-8 cwiczen bez pomocy |
| Doskonaly | Zagniezdzone podzapytania, AVG z COUNT | 9-10 cwiczen bez pomocy |

### Co dalej?
- Jesli poziom **Podstawowy**: Powtorz cwiczenia 21.1-21.2, przejdz do `cheatsheet_sql.md` (sekcja podzapytania).
- Jesli poziom **Dobry**: Przejdz do cwiczen `20_sql_group_by.md` i `22_sql_join.md`.
- Jesli poziom **Bardzo dobry**: Sprobuj cwiczen 21.9-21.10 bez wskazowek, potem przejdz do `23_sql_select_where.md` (cwiczenia trudne).
- Jesli poziom **Doskonaly**: Przejdz do arkuszy maturalnych — zacznij od 2023+.
