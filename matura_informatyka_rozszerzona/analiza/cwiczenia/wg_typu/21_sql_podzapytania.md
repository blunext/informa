# 21. SQL — Podzapytania (subqueries)

Typ zadania: **sql_podzapytania**
Czestotliwosc: 7/11 lat | Laczna punktacja: 25 pkt
Kategoria: SQL

---

### Cwiczenie 21.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2019 (Perfumy)

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

**Wyjasnienie**: Podzapytanie `(SELECT AVG(wynagrodzenie) FROM Pracownicy)` zwraca pojedyncza wartosc (6912.50), ktora jest uzywana w warunku WHERE. To tzw. podzapytanie skalarne — zwraca dokladnie jedna wartosc. Typowy blad maturalny: proba uzycia AVG bezposrednio w WHERE bez podzapytania, np. `WHERE wynagrodzenie > AVG(wynagrodzenie)` — to nie zadziala, bo funkcje agregujace nie moga byc uzywane w WHERE.
</details>

---

### Cwiczenie 21.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2021 (Gra strategiczna)

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

**Wyjasnienie**: Podzapytanie `(SELECT id_kursu FROM Zapisy)` zwraca zbior {1, 2, 3, 5}. NOT IN sprawdza, ktore id_kursu z tabeli Kursy NIE naleza do tego zbioru. Alternatywne rozwiazanie z LEFT JOIN: `SELECT K.nazwa_kursu FROM Kursy K LEFT JOIN Zapisy Z ON K.id_kursu = Z.id_kursu WHERE Z.id_zapisu IS NULL`. Typowy blad maturalny: zapomnienie o NOT — samo IN daloby kursy z zapisami, czyli odwrotnosc oczekiwanego wyniku.
</details>

---

### Cwiczenie 21.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 (Rejestr wykroczen)

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

**Wyjasnienie**: LEFT JOIN zachowuje WSZYSTKIE wiersze z lewej tabeli (Klienci), nawet jezeli nie maja dopasowania w prawej tabeli (Zamowienia). Dla niedopasowanych wierszy kolumny z Zamowienia maja wartosc NULL. Warunek `WHERE Z.id_zamowienia IS NULL` wybiera wlasnie te niedopasowane wiersze. Typowy blad maturalny: uzycie INNER JOIN zamiast LEFT JOIN — INNER JOIN odrzuca wiersze bez dopasowania, wiec nigdy nie pokaze klientow bez zamowien.
</details>

---

### Cwiczenie 21.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 (Gra strategiczna)

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

**Wyjasnienie**: Zapytanie laczy dwa warunki w WHERE: filtr na nazwe misji (porownanie tekstowe) i porownanie z wynikiem podzapytania skalarnego (srednia). Podzapytanie oblicza srednia z CALEJ tabeli Misje, nie tylko z misji 'Smocza jaskinia'. Typowy blad maturalny: ograniczenie podzapytania do jednej misji, np. `AVG(punkty) FROM Misje WHERE nazwa_misji = 'Smocza jaskinia'` — to daloby srednia = (120+150+140)/3 = 136.67, co zmienilloby wynik.
</details>

---

### Cwiczenie 21.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2019 (Perfumy)

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

Liczba filmow:
- sci-fi: Incepcja (3), Interstellar (4), Matrix (7) = 3 filmy

| gatunek | srednia_ocena | liczba_filmow |
|---------|---------------|---------------|
| sci-fi | 9.17 | 3 |

**Wyjasnienie**: Podzapytanie w klauzuli HAVING pozwala porownac wynik agregacji grupy ze srednia globalna. COUNT(DISTINCT F.id_filmu) liczy unikalne filmy, a nie unikalne oceny — jeden film moze miec wiele ocen. Typowy blad maturalny: uzycie COUNT(*) zamiast COUNT(DISTINCT id_filmu) — COUNT(*) policzyloby 6 ocen zamiast 3 filmow. Drugi blad: umieszczenie podzapytania w WHERE zamiast HAVING.
</details>
