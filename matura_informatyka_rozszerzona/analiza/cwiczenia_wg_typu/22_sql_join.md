# 22. SQL — JOIN (laczenie tabel)

Typ zadania: **sql_join**
Czestotliwosc: 8/11 lat | Laczna punktacja: 21 pkt
Kategoria: SQL

---

### Cwiczenie 22.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2016 (Uniwersytet)

Tabela **Autorzy**:

| id_autora | imie | nazwisko | kraj |
|-----------|------|----------|------|
| 1 | Adam | Mickiewicz | Polska |
| 2 | Juliusz | Slowacki | Polska |
| 3 | Fiodor | Dostojewski | Rosja |
| 4 | Victor | Hugo | Francja |

Tabela **Ksiazki**:

| id_ksiazki | tytul | id_autora | rok | cena |
|------------|-------|-----------|-----|------|
| 1 | Pan Tadeusz | 1 | 1834 | 29.90 |
| 2 | Dziady | 1 | 1823 | 24.50 |
| 3 | Balladyna | 2 | 1839 | 19.90 |
| 4 | Zbrodnia i kara | 3 | 1866 | 34.90 |
| 5 | Kordian | 2 | 1834 | 22.00 |
| 6 | Nedznicy | 4 | 1862 | 39.90 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli tytul ksiazki oraz imie i nazwisko jej autora. Posortuj alfabetycznie wg nazwiska autora, a nastepnie wg tytulu.

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT K.tytul, A.imie, A.nazwisko
FROM Ksiazki K
JOIN Autorzy A ON K.id_autora = A.id_autora
ORDER BY A.nazwisko, K.tytul;
```

**Weryfikacja**:

| tytul | imie | nazwisko |
|-------|------|----------|
| Zbrodnia i kara | Fiodor | Dostojewski |
| Nedznicy | Victor | Hugo |
| Dziady | Adam | Mickiewicz |
| Pan Tadeusz | Adam | Mickiewicz |
| Balladyna | Juliusz | Slowacki |
| Kordian | Juliusz | Slowacki |

**Wyjasnienie**: INNER JOIN (lub samo JOIN) laczy kazdy wiersz z Ksiazki z wierszem z Autorzy, gdzie id_autora sie zgadza. Jesli autor nie ma ksiazek, nie pojawi sie w wyniku (i odwrotnie). ORDER BY z dwoma kolumnami sortuje najpierw wg nazwiska, a w ramach tego samego nazwiska — wg tytulu. Typowy blad maturalny: brak aliasow (K, A) przy niejednoznacznych nazwach kolumn — jesli obie tabele maja kolumne o tej samej nazwie (np. id_autora), trzeba uzyc prefiksu tabeli.
</details>

---

### Cwiczenie 22.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 (Pilka reczna)

Tabela **Oddzialy**:

| id_oddzialu | nazwa_oddzialu |
|-------------|----------------|
| 1 | Chirurgia |
| 2 | Pediatria |
| 3 | Kardiologia |

Tabela **Lekarze**:

| id_lekarza | imie | nazwisko | id_oddzialu | specjalizacja |
|------------|------|----------|-------------|---------------|
| 1 | Jan | Kowalski | 1 | chirurg |
| 2 | Anna | Nowak | 2 | pediatra |
| 3 | Piotr | Maj | 1 | ortopeda |
| 4 | Ewa | Lis | 3 | kardiolog |

Tabela **Wizyty**:

| id_wizyty | id_lekarza | pacjent | data_wizyty |
|-----------|------------|---------|-------------|
| 1 | 1 | Tomek Bak | 2024-03-01 |
| 2 | 2 | Kasia Zak | 2024-03-01 |
| 3 | 1 | Ola Wrobel | 2024-03-02 |
| 4 | 4 | Marek Krol | 2024-03-02 |
| 5 | 3 | Adam Wilk | 2024-03-03 |
| 6 | 2 | Zofia Ptak | 2024-03-03 |
| 7 | 4 | Jan Duda | 2024-03-04 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwe oddzialu, imie i nazwisko lekarza oraz nazwe pacjenta dla kazdej wizyty. Posortuj wg daty wizyty.

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT O.nazwa_oddzialu, L.imie, L.nazwisko, W.pacjent, W.data_wizyty
FROM Wizyty W
JOIN Lekarze L ON W.id_lekarza = L.id_lekarza
JOIN Oddzialy O ON L.id_oddzialu = O.id_oddzialu
ORDER BY W.data_wizyty;
```

**Weryfikacja**:

| nazwa_oddzialu | imie | nazwisko | pacjent | data_wizyty |
|----------------|------|----------|---------|-------------|
| Chirurgia | Jan | Kowalski | Tomek Bak | 2024-03-01 |
| Pediatria | Anna | Nowak | Kasia Zak | 2024-03-01 |
| Chirurgia | Jan | Kowalski | Ola Wrobel | 2024-03-02 |
| Kardiologia | Ewa | Lis | Marek Krol | 2024-03-02 |
| Chirurgia | Piotr | Maj | Adam Wilk | 2024-03-03 |
| Pediatria | Anna | Nowak | Zofia Ptak | 2024-03-03 |
| Kardiologia | Ewa | Lis | Jan Duda | 2024-03-04 |

**Wyjasnienie**: Laczenie 3 tabel wymaga dwoch JOIN-ow: Wizyty → Lekarze (po id_lekarza) i Lekarze → Oddzialy (po id_oddzialu). Kolejnosc JOIN-ow nie ma znaczenia dla wyniku, ale logicznie laczymy "od srodka" — kazda wizyta ma lekarza, kazdy lekarz nalezy do oddzialu. Typowy blad maturalny: pominiecie jednego JOIN-a i proba bezposredniego laczenia Wizyty z Oddzialy — te tabele nie maja wspolnej kolumny.
</details>

---

### Cwiczenie 22.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2015 (Formula 1)

Tabela **Rowery**:

| id_roweru | model | typ | cena_za_godzine |
|-----------|-------|-----|-----------------|
| 1 | City Cruiser | miejski | 15 |
| 2 | Mountain Pro | gorski | 25 |
| 3 | Speed Racer | szosowy | 30 |
| 4 | Kids Fun | dzieciecy | 10 |
| 5 | Electric One | elektryczny | 35 |

Tabela **Wypozyczenia**:

| id_wypozyczenia | id_roweru | klient | godziny | data |
|-----------------|-----------|--------|---------|------|
| 1 | 1 | Tomek | 3 | 2024-06-01 |
| 2 | 2 | Anna | 2 | 2024-06-01 |
| 3 | 1 | Kasia | 1 | 2024-06-02 |
| 4 | 3 | Piotr | 4 | 2024-06-02 |
| 5 | 2 | Tomek | 3 | 2024-06-03 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli model kazdego roweru oraz laczna liczbe jego wypozyczen. Pokaz WSZYSTKIE rowery, nawet te nigdy nie wypozyczone (dla nich liczba = 0). Posortuj malejaco wg liczby wypozyczen.

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT R.model, COUNT(W.id_wypozyczenia) AS liczba_wypozyczen
FROM Rowery R
LEFT JOIN Wypozyczenia W ON R.id_roweru = W.id_roweru
GROUP BY R.model
ORDER BY liczba_wypozyczen DESC;
```

**Weryfikacja**:
- City Cruiser (id=1): wyp. 1, 3 → **2**
- Mountain Pro (id=2): wyp. 2, 5 → **2**
- Speed Racer (id=3): wyp. 4 → **1**
- Kids Fun (id=4): brak wypozyczen → **0**
- Electric One (id=5): brak wypozyczen → **0**

| model | liczba_wypozyczen |
|-------|-------------------|
| City Cruiser | 2 |
| Mountain Pro | 2 |
| Speed Racer | 1 |
| Electric One | 0 |
| Kids Fun | 0 |

**Wyjasnienie**: LEFT JOIN zachowuje wszystkie wiersze z lewej tabeli (Rowery), nawet te bez dopasowania w Wypozyczenia. Kluczowy szczegol: uzywamy `COUNT(W.id_wypozyczenia)` zamiast `COUNT(*)`. Roznica: COUNT(*) liczy WSZYSTKIE wiersze, wlacznie z tymi gdzie W jest NULL (daloby 1 zamiast 0 dla niewypozyczonych). COUNT(kolumna) pomija wartosci NULL. Typowy blad maturalny: uzycie COUNT(*) z LEFT JOIN — dla rowerow bez wypozyczen daje 1 zamiast 0.
</details>

---

### Cwiczenie 22.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2022 (Ewidencja uczniow)

Tabela **Nauczyciele**:

| id_nauczyciela | imie | nazwisko |
|----------------|------|----------|
| 1 | Jan | Kowalski |
| 2 | Anna | Nowak |
| 3 | Ewa | Maj |

Tabela **Przedmioty**:

| id_przedmiotu | nazwa_przedmiotu |
|---------------|------------------|
| 1 | Matematyka |
| 2 | Fizyka |
| 3 | Angielski |
| 4 | Informatyka |

Tabela **Lekcje**:

| id_lekcji | id_nauczyciela | id_przedmiotu | klasa | dzien_tygodnia |
|-----------|----------------|---------------|-------|----------------|
| 1 | 1 | 1 | 1A | poniedzialek |
| 2 | 1 | 1 | 2B | wtorek |
| 3 | 2 | 3 | 1A | poniedzialek |
| 4 | 2 | 3 | 1B | sroda |
| 5 | 3 | 4 | 2A | czwartek |
| 6 | 1 | 2 | 1A | sroda |
| 7 | 3 | 4 | 1A | piatek |
| 8 | 2 | 3 | 2A | wtorek |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli imie i nazwisko nauczyciela, nazwe przedmiotu oraz klase — ale tylko dla lekcji w klasie '1A'. Posortuj wg dnia tygodnia.

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT N.imie, N.nazwisko, P.nazwa_przedmiotu, L.klasa, L.dzien_tygodnia
FROM Lekcje L
JOIN Nauczyciele N ON L.id_nauczyciela = N.id_nauczyciela
JOIN Przedmioty P ON L.id_przedmiotu = P.id_przedmiotu
WHERE L.klasa = '1A'
ORDER BY L.dzien_tygodnia;
```

**Weryfikacja**:
Lekcje w klasie 1A: id_lekcji = 1, 3, 6, 7

| imie | nazwisko | nazwa_przedmiotu | klasa | dzien_tygodnia |
|------|----------|------------------|-------|----------------|
| Jan | Kowalski | Fizyka | 1A | sroda |
| Anna | Nowak | Angielski | 1A | poniedzialek |
| Jan | Kowalski | Matematyka | 1A | poniedzialek |
| Ewa | Maj | Informatyka | 1A | piatek |

**Uwaga**: Sortowanie alfabetyczne wg dnia tygodnia (czwartek < piatek < poniedzialek < sroda < wtorek) — to sortowanie tekstowe, nie chronologiczne. Gdyby matura wymagala sortowania chronologicznego, potrzebna bylaby kolumna z numerem dnia lub CASE WHEN.

**Wyjasnienie**: Trzy tabele laczone dwoma JOIN-ami, z dodatkowym filtrem WHERE. Kolejnosc: najpierw JOIN laczy wszystkie powiazane wiersze, potem WHERE filtruje wynik. Typowy blad maturalny: umieszczenie warunku filtrowania w ON zamiast w WHERE — przy INNER JOIN wynik jest taki sam, ale przy LEFT JOIN daje rozne wyniki. Na maturze bezpieczniej jest filtrowac w WHERE.
</details>

---

### Cwiczenie 22.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 (Woda na Marsie)

Tabela **Bazy**:

| id_bazy | nazwa_bazy | sektor |
|---------|------------|--------|
| 1 | Alfa | polnocny |
| 2 | Beta | poludniowy |
| 3 | Gamma | polnocny |

Tabela **Astronauci**:

| id_astronauty | imie | nazwisko | id_bazy |
|---------------|------|----------|---------|
| 1 | Jan | Kowalski | 1 |
| 2 | Anna | Nowak | 2 |
| 3 | Piotr | Maj | 1 |
| 4 | Ewa | Lis | 3 |
| 5 | Kasia | Zak | 2 |

Tabela **Projekty**:

| id_projektu | nazwa_projektu | data_rozpoczecia |
|-------------|----------------|-----------------|
| 1 | Terraformacja | 2024-01-15 |
| 2 | Gornictwo | 2024-03-01 |
| 3 | Energia solarna | 2024-06-10 |
| 4 | Szklarnie | 2024-02-20 |

Tabela **Przydzialy**:

| id_przydzialu | id_astronauty | id_projektu | rola |
|---------------|---------------|-------------|------|
| 1 | 1 | 1 | lider |
| 2 | 2 | 1 | czlonek |
| 3 | 3 | 2 | lider |
| 4 | 1 | 3 | czlonek |
| 5 | 4 | 2 | czlonek |
| 6 | 5 | 4 | lider |
| 7 | 2 | 3 | czlonek |
| 8 | 3 | 4 | czlonek |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwe bazy, imie i nazwisko astronauty, nazwe projektu oraz role — ale tylko dla projektow rozpoczetych przed '2024-04-01'. Kazda kombinacja baza-projekt powinna wystapic co najwyzej raz (uzyj DISTINCT). Posortuj wg nazwy bazy, nastepnie wg nazwy projektu.

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT DISTINCT B.nazwa_bazy, A.imie, A.nazwisko, P.nazwa_projektu, PR.rola
FROM Bazy B
JOIN Astronauci A ON B.id_bazy = A.id_bazy
JOIN Przydzialy PR ON A.id_astronauty = PR.id_astronauty
JOIN Projekty P ON PR.id_projektu = P.id_projektu
WHERE P.data_rozpoczecia < '2024-04-01'
ORDER BY B.nazwa_bazy, P.nazwa_projektu;
```

**Weryfikacja**:
Projekty rozpoczete przed 2024-04-01:
- Terraformacja (2024-01-15 ✓)
- Gornictwo (2024-03-01 ✓)
- Szklarnie (2024-02-20 ✓)
- Energia solarna (2024-06-10 ✗)

Przydzialy do tych projektow:
- Terraformacja: Jan Kowalski (Alfa, lider), Anna Nowak (Beta, czlonek)
- Gornictwo: Piotr Maj (Alfa, lider), Ewa Lis (Gamma, czlonek)
- Szklarnie: Kasia Zak (Beta, lider), Piotr Maj (Alfa, czlonek)

| nazwa_bazy | imie | nazwisko | nazwa_projektu | rola |
|------------|------|----------|----------------|------|
| Alfa | Jan | Kowalski | Terraformacja | lider |
| Alfa | Piotr | Maj | Gornictwo | lider |
| Alfa | Piotr | Maj | Szklarnie | czlonek |
| Beta | Anna | Nowak | Terraformacja | czlonek |
| Beta | Kasia | Zak | Szklarnie | lider |
| Gamma | Ewa | Lis | Gornictwo | czlonek |

**Wyjasnienie**: Laczenie 4 tabel wymaga 3 JOIN-ow tworzacych lancuch: Bazy → Astronauci → Przydzialy → Projekty. DISTINCT zapewnia, ze nie ma duplikatow w wynikach (tutaj nie ma ich naturalnie, ale na maturze czesto wymagane jest wyrazne uzycie DISTINCT). Porownanie dat: `data_rozpoczecia < '2024-04-01'` dziala poprawnie, bo daty w formacie ISO (RRRR-MM-DD) mozna porownywac jako tekst. Typowy blad maturalny: bledna kolejnosc JOIN-ow lub pominiecie jednej tabeli posredniczacej — np. proba polaczenia Bazy bezposrednio z Projekty bez przejscia przez Astronauci i Przydzialy.
</details>
