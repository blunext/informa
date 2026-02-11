# 22. SQL — JOIN (laczenie tabel)

Typ zadania: **sql_join**
Czestotliwosc: 8/11 lat | Laczna punktacja: 21 pkt
Kategoria: SQL

## Umiejetnosci cwiczone w tym zestawie

`INNER-JOIN` `LEFT-JOIN` `JOIN-wielotabelowy` `aliasy-tabel` `ON-warunek` `COUNT-z-LEFT-JOIN` `self-JOIN` `GROUP-BY-z-JOIN` `WHERE-z-JOIN` `ORDER-BY-wielokolumnowy` `DISTINCT`

---

### Cwiczenie 22.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2016 (Uniwersytet)
**Tagi**: `INNER-JOIN` `aliasy-tabel` `ORDER-BY-wielokolumnowy`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Dane autora sa w innej tabeli niz dane ksiazki — musisz je polaczyc.
2. **Podejscie**: JOIN na wspolnej kolumnie id_autora.
3. **Kluczowy krok**: ORDER BY z dwoma kolumnami: najpierw nazwisko, potem tytul.

</details>

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

**Wyjasnienie**: INNER JOIN (lub samo JOIN) laczy kazdy wiersz z Ksiazki z wierszem z Autorzy, gdzie id_autora sie zgadza. ORDER BY z dwoma kolumnami sortuje najpierw wg nazwiska, a w ramach tego samego nazwiska — wg tytulu.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak aliasow przy niejednoznacznych kolumnach**: Jesli obie tabele maja kolumne `id_autora`, trzeba uzyc prefiksu (K.id_autora lub A.id_autora). CKE: -1 pkt (blad skladni).
- **Brak ON**: `FROM Ksiazki JOIN Autorzy` bez ON tworzy iloczyn kartezjanski (kazda ksiazka z kazdym autorem). CKE: -3 pkt.

</details>

---

### Cwiczenie 22.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 (Pilka reczna)
**Tagi**: `JOIN-wielotabelowy` `aliasy-tabel` `ORDER-BY-wielokolumnowy`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy tabele — potrzebujesz dwoch JOIN-ow.
2. **Podejscie**: Wizyty → Lekarze (po id_lekarza) → Oddzialy (po id_oddzialu).
3. **Kluczowy krok**: Kolejnosc JOIN-ow: od tabeli z danymi (Wizyty) przez tabele posrednia (Lekarze) do tabeli docelowej (Oddzialy).

</details>

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

**Wyjasnienie**: Laczenie 3 tabel wymaga dwoch JOIN-ow: Wizyty → Lekarze (po id_lekarza) i Lekarze → Oddzialy (po id_oddzialu). Kolejnosc JOIN-ow nie ma znaczenia dla wyniku, ale logicznie laczymy "od srodka".
</details>

<details>
<summary>Typowe bledy</summary>

- **Pominiecie jednego JOIN-a**: Proba polaczenia Wizyty bezposrednio z Oddzialy — te tabele nie maja wspolnej kolumny. CKE: -2 pkt.
- **Bledny warunek ON**: np. `ON W.id_lekarza = O.id_oddzialu` — laczenie blednych kolumn. CKE: -3 pkt.

</details>

---

### Cwiczenie 22.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2015 (Formula 1)
**Tagi**: `LEFT-JOIN` `COUNT-z-LEFT-JOIN` `GROUP-BY-z-JOIN`

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
<summary>Wskazowki</summary>

1. **Kierunek**: INNER JOIN odrzuca rowery bez wypozyczen — potrzebujesz LEFT JOIN.
2. **Podejscie**: LEFT JOIN zachowuje WSZYSTKIE wiersze z lewej tabeli (Rowery).
3. **Kluczowy krok**: COUNT(W.id_wypozyczenia) zamiast COUNT(*) — COUNT(*) daje 1 zamiast 0 dla pustych dopasowañ.

</details>

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

**Wyjasnienie**: LEFT JOIN zachowuje wszystkie wiersze z lewej tabeli (Rowery), nawet te bez dopasowania w Wypozyczenia. Kluczowy szczegol: uzywamy `COUNT(W.id_wypozyczenia)` zamiast `COUNT(*)`.
</details>

<details>
<summary>Typowe bledy</summary>

- **COUNT(*) z LEFT JOIN**: Dla rowerow bez wypozyczen daje 1 zamiast 0, bo wiersz z NULL jest liczony. CKE: -1 pkt.
- **INNER JOIN zamiast LEFT JOIN**: Rowery bez wypozyczen nie pojawia sie w wyniku. CKE: -2 pkt.
- **RIGHT JOIN**: Poprawna alternatywa (odwrocona kolejnosc tabel), ale mniej czytelna i na maturze rzadko uzywana. CKE: akceptowane.

</details>

---

### Cwiczenie 22.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2022 (Ewidencja uczniow)
**Tagi**: `JOIN-wielotabelowy` `WHERE-z-JOIN` `ORDER-BY-wielokolumnowy`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy tabele laczone dwoma JOIN-ami, plus filtr WHERE.
2. **Podejscie**: Lekcje → Nauczyciele (po id_nauczyciela) → Przedmioty (po id_przedmiotu), WHERE klasa = '1A'.
3. **Kluczowy krok**: Sortowanie tekstowe po dniu tygodnia — 'czwartek' < 'piatek' < 'poniedzialek' < 'sroda'.

</details>

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
| Ewa | Maj | Informatyka | 1A | piatek |
| Jan | Kowalski | Matematyka | 1A | poniedzialek |
| Anna | Nowak | Angielski | 1A | poniedzialek |
| Jan | Kowalski | Fizyka | 1A | sroda |

**Uwaga**: Sortowanie alfabetyczne wg dnia tygodnia (nie chronologiczne). Gdyby matura wymagala sortowania chronologicznego, potrzebna bylaby kolumna z numerem dnia lub CASE WHEN.

**Wyjasnienie**: Trzy tabele laczone dwoma JOIN-ami, z dodatkowym filtrem WHERE. Kolejnosc: najpierw JOIN laczy wszystkie powiazane wiersze, potem WHERE filtruje wynik.
</details>

<details>
<summary>Typowe bledy</summary>

- **Warunek w ON zamiast WHERE**: `JOIN ... ON ... AND klasa = '1A'` — przy INNER JOIN wynik ten sam, ale przy LEFT JOIN inny. Na maturze bezpieczniej filtrowac w WHERE. CKE: akceptowane, ale ryzykowne.
- **Brak jednego JOIN**: Bez JOIN z Przedmioty nie uzyskamy nazwy przedmiotu. CKE: -2 pkt.

</details>

---

### Cwiczenie 22.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2025 (Woda na Marsie)
**Tagi**: `JOIN-wielotabelowy` `DISTINCT` `WHERE-z-JOIN` `ORDER-BY-wielokolumnowy`

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
<summary>Wskazowki</summary>

1. **Kierunek**: Cztery tabele — potrzebujesz trzech JOIN-ow tworzacych lancuch.
2. **Podejscie**: Bazy → Astronauci → Przydzialy → Projekty, z WHERE na date.
3. **Kluczowy krok**: Daty w formacie ISO (RRRR-MM-DD) mozna porownywac jako tekst: `data_rozpoczecia < '2024-04-01'`.

</details>

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

| nazwa_bazy | imie | nazwisko | nazwa_projektu | rola |
|------------|------|----------|----------------|------|
| Alfa | Jan | Kowalski | Terraformacja | lider |
| Alfa | Piotr | Maj | Gornictwo | lider |
| Alfa | Piotr | Maj | Szklarnie | czlonek |
| Beta | Anna | Nowak | Terraformacja | czlonek |
| Beta | Kasia | Zak | Szklarnie | lider |
| Gamma | Ewa | Lis | Gornictwo | czlonek |

**Wyjasnienie**: Laczenie 4 tabel wymaga 3 JOIN-ow tworzacych lancuch: Bazy → Astronauci → Przydzialy → Projekty. DISTINCT zapewnia brak duplikatow. Porownanie dat w formacie ISO dziala jako porownanie tekstowe.
</details>

<details>
<summary>Typowe bledy</summary>

- **Bledna kolejnosc JOIN-ow**: Proba polaczenia Bazy bezposrednio z Projekty — brak wspolnej kolumny. CKE: -3 pkt.
- **Brak DISTINCT**: Jesli w danych bylyby duplikaty, wynik zawieralby powtorzenia. CKE: -1 pkt.
- **Data bez apostrofow**: `WHERE data < 2024-04-01` — interpretowane jako wyrazenie arytmetyczne (2024 minus 4 minus 1 = 2019). CKE: -2 pkt.

</details>

---

### Cwiczenie 22.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 (Gry planszowe)
**Tagi**: `INNER-JOIN` `aliasy-tabel` `WHERE-z-JOIN`

Tabela **Kategorie**:

| id_kat | nazwa_kategorii |
|--------|-----------------|
| 1 | Strategiczne |
| 2 | Karciane |
| 3 | Familijne |
| 4 | Kooperacyjne |

Tabela **Gry**:

| id_gry | tytul | id_kat | min_graczy | max_graczy | czas_min |
|--------|-------|--------|------------|------------|----------|
| 1 | Catan | 1 | 3 | 4 | 90 |
| 2 | Uno | 2 | 2 | 10 | 30 |
| 3 | Monopoly | 3 | 2 | 6 | 120 |
| 4 | Pandemic | 4 | 2 | 4 | 60 |
| 5 | Risk | 1 | 2 | 6 | 180 |
| 6 | Dixit | 3 | 3 | 6 | 45 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli tytul gry i nazwe kategorii dla gier, w ktore moze grac co najmniej 3 graczy (min_graczy <= 3 AND max_graczy >= 3). Posortuj wg czasu gry.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: JOIN dwoch tabel z dodatkowym filtrem WHERE.
2. **Podejscie**: JOIN Gry z Kategorie, WHERE na min/max graczy.
3. **Kluczowy krok**: "Co najmniej 3 graczy" oznacza, ze 3 musi byc w zakresie [min_graczy, max_graczy].

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT G.tytul, K.nazwa_kategorii, G.czas_min
FROM Gry G
JOIN Kategorie K ON G.id_kat = K.id_kat
WHERE G.min_graczy <= 3 AND G.max_graczy >= 3
ORDER BY G.czas_min;
```

**Weryfikacja**:
- Catan: min=3 <=3 ✓, max=4 >=3 ✓ → ✓
- Uno: min=2 <=3 ✓, max=10 >=3 ✓ → ✓
- Monopoly: min=2 <=3 ✓, max=6 >=3 ✓ → ✓
- Pandemic: min=2 <=3 ✓, max=4 >=3 ✓ → ✓
- Risk: min=2 <=3 ✓, max=6 >=3 ✓ → ✓
- Dixit: min=3 <=3 ✓, max=6 >=3 ✓ → ✓

| tytul | nazwa_kategorii | czas_min |
|-------|-----------------|----------|
| Uno | Karciane | 30 |
| Dixit | Familijne | 45 |
| Pandemic | Kooperacyjne | 60 |
| Catan | Strategiczne | 90 |
| Monopoly | Familijne | 120 |
| Risk | Strategiczne | 180 |

**Wyjasnienie**: Warunek "moze grac 3 graczy" oznacza, ze 3 musi byc w zakresie od min_graczy do max_graczy. To dwa warunki polaczone AND.
</details>

<details>
<summary>Typowe bledy</summary>

- **Tylko jeden warunek**: `WHERE max_graczy >= 3` bez min_graczy — gdyby jakas gra miala min=4, tez by sie pojawila. CKE: -1 pkt.
- **OR zamiast AND**: Zmienilby logike — wystarczylyby spelnienie jednego warunku. CKE: -2 pkt.

</details>

---

### Cwiczenie 22.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)
**Tagi**: `LEFT-JOIN` `COUNT-z-LEFT-JOIN` `GROUP-BY-z-JOIN`

Tabela **Dzialy**:

| id_dzialu | nazwa_dzialu | lokalizacja |
|-----------|-------------|-------------|
| 1 | IT | pietro_3 |
| 2 | HR | pietro_1 |
| 3 | Finanse | pietro_2 |
| 4 | Marketing | pietro_1 |
| 5 | Logistyka | pietro_0 |

Tabela **Pracownicy**:

| id | imie | nazwisko | id_dzialu | data_zatrudnienia |
|----|------|----------|-----------|-------------------|
| 1 | Anna | Kowalska | 1 | 2020-03-15 |
| 2 | Jan | Nowak | 1 | 2021-07-01 |
| 3 | Ewa | Maj | 2 | 2019-11-20 |
| 4 | Piotr | Lis | 3 | 2022-01-10 |
| 5 | Kasia | Zak | 1 | 2023-06-01 |
| 6 | Tomek | Bak | 2 | 2020-09-15 |
| 7 | Ola | Wrobel | 4 | 2021-04-01 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwe dzialu oraz liczbe pracownikow w tym dziale. Pokaz WSZYSTKIE dzialy, nawet te bez pracownikow. Posortuj malejaco wg liczby pracownikow.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: LEFT JOIN zachowuje wszystkie dzialy, nawet puste.
2. **Podejscie**: LEFT JOIN Dzialy z Pracownicy, GROUP BY dzial, COUNT kolumny z prawej tabeli.
3. **Kluczowy krok**: COUNT(P.id) zamiast COUNT(*) — dla pustych dzialow COUNT(*) daje 1, a COUNT(P.id) daje 0.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT D.nazwa_dzialu, COUNT(P.id) AS liczba_pracownikow
FROM Dzialy D
LEFT JOIN Pracownicy P ON D.id_dzialu = P.id_dzialu
GROUP BY D.nazwa_dzialu
ORDER BY liczba_pracownikow DESC;
```

**Weryfikacja**:
- IT: Anna, Jan, Kasia → **3**
- HR: Ewa, Tomek → **2**
- Finanse: Piotr → **1**
- Marketing: Ola → **1**
- Logistyka: (brak) → **0**

| nazwa_dzialu | liczba_pracownikow |
|--------------|--------------------|
| IT | 3 |
| HR | 2 |
| Finanse | 1 |
| Marketing | 1 |
| Logistyka | 0 |

**Wyjasnienie**: LEFT JOIN + COUNT(kolumna z prawej tabeli) + GROUP BY to standardowy wzorzec "policz powiazane rekordy, w tym zerowe". COUNT pomija NULL, wiec puste grupy daja 0.
</details>

<details>
<summary>Typowe bledy</summary>

- **COUNT(*)**: Daje 1 zamiast 0 dla Logistyki — LEFT JOIN tworzy jeden wiersz z NULL, COUNT(*) go liczy. CKE: -1 pkt.
- **INNER JOIN**: Logistyka nie pojawi sie w wyniku. CKE: -2 pkt.

</details>

---

### Cwiczenie 22.8 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 (Perfumy)
**Tagi**: `self-JOIN` `aliasy-tabel` `WHERE-z-JOIN`

Tabela **Pracownicy**:

| id | imie | nazwisko | id_przelozonego | stanowisko |
|----|------|----------|-----------------|------------|
| 1 | Jan | Kowalski | NULL | dyrektor |
| 2 | Anna | Nowak | 1 | kierownik |
| 3 | Piotr | Maj | 1 | kierownik |
| 4 | Ewa | Lis | 2 | specjalista |
| 5 | Kasia | Zak | 2 | specjalista |
| 6 | Tomek | Bak | 3 | specjalista |
| 7 | Ola | Wrobel | 3 | stazystka |
| 8 | Marek | Krol | 3 | specjalista |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli imie i nazwisko kazdego pracownika oraz imie i nazwisko jego przelozonego. Pokaz tylko pracownikow, ktorzy MAJA przelozonego (dyrektor nie ma).

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Tabela odwoluje sie sama do siebie — potrzebujesz self-JOIN.
2. **Podejscie**: JOIN tabeli Pracownicy z nia sama, uzywajac ROZNYCH aliasow.
3. **Kluczowy krok**: `JOIN Pracownicy P2 ON P.id_przelozonego = P2.id` — P to pracownik, P2 to jego przelozony.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT P.imie, P.nazwisko, P2.imie AS imie_przelozonego, P2.nazwisko AS nazwisko_przelozonego
FROM Pracownicy P
JOIN Pracownicy P2 ON P.id_przelozonego = P2.id;
```

**Weryfikacja**:

| imie | nazwisko | imie_przelozonego | nazwisko_przelozonego |
|------|----------|-------------------|-----------------------|
| Anna | Nowak | Jan | Kowalski |
| Piotr | Maj | Jan | Kowalski |
| Ewa | Lis | Anna | Nowak |
| Kasia | Zak | Anna | Nowak |
| Tomek | Bak | Piotr | Maj |
| Ola | Wrobel | Piotr | Maj |
| Marek | Krol | Piotr | Maj |

Jan Kowalski (id_przelozonego = NULL) nie pojawia sie — INNER JOIN pomija NULL.

**Wyjasnienie**: Self-JOIN laczy tabele z sama soba. Kluczowe sa aliasy (P i P2) pozwalajace odroznic "pracownika" od "przelozonego". INNER JOIN automatycznie wyklucza pracownikow z id_przelozonego = NULL.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak aliasow**: Bez roznych aliasow (P, P2) nie mozna odroznic kolumn pracownika od przelozonego. CKE: -2 pkt (blad skladni).
- **Odwrotny warunek ON**: `ON P.id = P2.id_przelozonego` — zamiast przelozonego, dostajemy podwladnych. CKE: -2 pkt.
- **LEFT JOIN bez potrzeby**: Gdyby polecenie mowilo "WSZYSTKICH pracownikow", potrzebny bylby LEFT JOIN — tu wystarczy INNER JOIN. CKE: akceptowane, ale wynik inny (dyrektor z NULL).

</details>

---

### Cwiczenie 22.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 (Gra strategiczna)
**Tagi**: `JOIN-wielotabelowy` `GROUP-BY-z-JOIN` `LEFT-JOIN` `COUNT-z-LEFT-JOIN`

Tabela **Kategorie**:

| id_kat | nazwa |
|--------|-------|
| 1 | Elektronika |
| 2 | Ksiazki |
| 3 | Sport |

Tabela **Produkty**:

| id_prod | nazwa | id_kat | cena |
|---------|-------|--------|------|
| 1 | Laptop | 1 | 3500 |
| 2 | Smartfon | 1 | 2500 |
| 3 | Wiedzmin | 2 | 45 |
| 4 | Dune | 2 | 55 |
| 5 | Pilka | 3 | 80 |
| 6 | Rower | 3 | 1200 |

Tabela **Opinie**:

| id_op | id_prod | uzytkownik | ocena | data |
|-------|---------|------------|-------|------|
| 1 | 1 | Jan | 5 | 2024-01-10 |
| 2 | 1 | Anna | 4 | 2024-01-15 |
| 3 | 2 | Ewa | 3 | 2024-02-01 |
| 4 | 3 | Jan | 5 | 2024-02-10 |
| 5 | 4 | Anna | 4 | 2024-03-01 |
| 6 | 6 | Piotr | 5 | 2024-03-15 |

**Polecenie**: Napisz zapytanie SQL, ktore dla kazdej kategorii wyswietli nazwe kategorii, liczbe produktow z opiniami oraz srednia ocene (zaokraglona do 1 miejsca). Pokaz WSZYSTKIE kategorie — jesli kategoria nie ma opinii, pokaz 0 produktow i NULL jako srednia.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy tabele, LEFT JOIN aby zachowac kategorie bez opinii.
2. **Podejscie**: Kategorie LEFT JOIN Produkty LEFT JOIN Opinie, GROUP BY kategoria.
3. **Kluczowy krok**: COUNT(DISTINCT P.id_prod) liczy unikalne produkty z opiniami (nie liczbe opinii).

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT K.nazwa, COUNT(DISTINCT O.id_prod) AS produkty_z_opiniami,
       ROUND(AVG(O.ocena), 1) AS srednia_ocena
FROM Kategorie K
LEFT JOIN Produkty P ON K.id_kat = P.id_kat
LEFT JOIN Opinie O ON P.id_prod = O.id_prod
GROUP BY K.nazwa
ORDER BY produkty_z_opiniami DESC;
```

**Weryfikacja**:
- Elektronika: Laptop (oceny 5,4), Smartfon (ocena 3) → 2 produkty, srednia = (5+4+3)/3 = **4.0**
- Ksiazki: Wiedzmin (ocena 5), Dune (ocena 4) → 2 produkty, srednia = (5+4)/2 = **4.5**
- Sport: Rower (ocena 5), Pilka (brak) → 1 produkt z opinia, srednia = 5/1 = **5.0**

| nazwa | produkty_z_opiniami | srednia_ocena |
|-------|---------------------|---------------|
| Elektronika | 2 | 4.0 |
| Ksiazki | 2 | 4.5 |
| Sport | 1 | 5.0 |

**Wyjasnienie**: Dwa LEFT JOIN-y zapewniaja, ze kategorie bez produktow i produkty bez opinii tez sie pojawia. COUNT(DISTINCT O.id_prod) liczy unikalne produkty, ktore maja opinie — bez DISTINCT Laptop z 2 opiniami bylby policzony dwa razy.
</details>

<details>
<summary>Typowe bledy</summary>

- **COUNT(*) zamiast COUNT(DISTINCT O.id_prod)**: Liczy opinie, nie produkty. Elektronika daje 3 zamiast 2. CKE: -1 pkt.
- **INNER JOIN zamiast LEFT JOIN**: Kategoria bez opinii nie pojawi sie. CKE: -2 pkt.
- **Tylko jeden LEFT JOIN**: `LEFT JOIN Produkty ... JOIN Opinie` — INNER JOIN na Opinie eliminuje produkty bez opinii. CKE: -1 pkt.

</details>

---

### Cwiczenie 22.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2015 (Formula 1)
**Tagi**: `JOIN-wielotabelowy` `self-JOIN` `GROUP-BY-z-JOIN` `DISTINCT` `WHERE-z-JOIN`

Tabela **Kraje**:

| id_kraju | nazwa_kraju | kontynent |
|----------|-------------|-----------|
| 1 | Polska | Europa |
| 2 | Niemcy | Europa |
| 3 | Brazylia | Ameryka |
| 4 | Japonia | Azja |

Tabela **Kierowcy**:

| id_kier | imie | nazwisko | id_kraju |
|---------|------|----------|----------|
| 1 | Robert | Kubica | 1 |
| 2 | Sebastian | Vettel | 2 |
| 3 | Ayrton | Senna | 3 |
| 4 | Yuki | Tsunoda | 4 |
| 5 | Mick | Schumacher | 2 |
| 6 | Felipe | Massa | 3 |

Tabela **Wyscigi**:

| id_wysc | nazwa | tor | data |
|---------|-------|-----|------|
| 1 | GP Monako | Monte Carlo | 2024-05-26 |
| 2 | GP Japonii | Suzuka | 2024-04-07 |
| 3 | GP Brazylii | Interlagos | 2024-11-03 |

Tabela **Wyniki**:

| id | id_kier | id_wysc | pozycja | punkty |
|----|---------|---------|---------|--------|
| 1 | 1 | 1 | 8 | 4 |
| 2 | 2 | 1 | 1 | 25 |
| 3 | 3 | 1 | 3 | 15 |
| 4 | 4 | 2 | 5 | 10 |
| 5 | 2 | 2 | 2 | 18 |
| 6 | 5 | 2 | 10 | 1 |
| 7 | 3 | 3 | 1 | 25 |
| 8 | 6 | 3 | 4 | 12 |
| 9 | 1 | 3 | 6 | 8 |
| 10 | 2 | 3 | 3 | 15 |

**Polecenie**: Napisz zapytanie SQL, ktore dla kazdego kontynentu wyswietli:
- nazwe kontynentu
- liczbe roznych kierowcow, ktorzy zdobyli punkty
- laczna sume punktow

Pokaz tylko kontynenty z lacza suma punktow > 20. Posortuj malejaco wg sumy.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Cztery tabele: Kraje → Kierowcy → Wyniki (← Wyscigi nie jest potrzebna).
2. **Podejscie**: JOIN trzech tabel, GROUP BY kontynent, HAVING na SUM.
3. **Kluczowy krok**: COUNT(DISTINCT K.id_kier) liczy unikalnych kierowcow, bo jeden moze miec wiele wynikow.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT KR.kontynent,
       COUNT(DISTINCT K.id_kier) AS liczba_kierowcow,
       SUM(W.punkty) AS suma_punktow
FROM Kraje KR
JOIN Kierowcy K ON KR.id_kraju = K.id_kraju
JOIN Wyniki W ON K.id_kier = W.id_kier
GROUP BY KR.kontynent
HAVING SUM(W.punkty) > 20
ORDER BY suma_punktow DESC;
```

**Weryfikacja**:

Europa:
- Kubica: 4 + 8 = 12 pkt
- Vettel: 25 + 18 + 15 = 58 pkt
- Schumacher: 1 pkt
- Kierowcy: 3, Suma: 12 + 58 + 1 = **71** (>20 ✓)

Ameryka:
- Senna: 15 + 25 = 40 pkt
- Massa: 12 pkt
- Kierowcy: 2, Suma: 40 + 12 = **52** (>20 ✓)

Azja:
- Tsunoda: 10 pkt
- Kierowcy: 1, Suma: **10** (>20 ✗)

| kontynent | liczba_kierowcow | suma_punktow |
|-----------|------------------|-------------|
| Europa | 3 | 71 |
| Ameryka | 2 | 52 |

**Wyjasnienie**: Lancuch JOIN-ow: Kraje → Kierowcy → Wyniki. Tabela Wyscigi nie jest potrzebna — nie pytamy o nazwy wyscigow. COUNT(DISTINCT) zapewnia poprawne zliczanie kierowcow. HAVING filtruje grupy po agregacji.
</details>

<details>
<summary>Typowe bledy</summary>

- **Dodatkowy JOIN z Wyscigi**: Niepotrzebny — powieksza wynik (wiele wierszy per wynik), ale COUNT(DISTINCT) i SUM daja ten sam rezultat. Jednak spowalnia zapytanie. CKE: akceptowane, ale zbedne.
- **COUNT(*) zamiast COUNT(DISTINCT)**: Vettel z 3 wynikami bylby policzony 3 razy. Europa daje 6 zamiast 3. CKE: -1 pkt.
- **Brak HAVING**: Azja (10 pkt) pojawia sie w wyniku. CKE: -1 pkt.

</details>

---

## Samoocena

| Poziom | Opis | Wymaganie |
|--------|------|-----------|
| Podstawowy | INNER JOIN dwoch tabel z ON i ORDER BY | 1-3 cwiczen bez pomocy |
| Dobry | LEFT JOIN z COUNT, JOIN trzech tabel | 4-6 cwiczen bez pomocy |
| Bardzo dobry | Self-JOIN, GROUP BY z JOIN, DISTINCT | 7-8 cwiczen bez pomocy |
| Doskonaly | Lancuchy 4 tabel, zlozone GROUP BY z HAVING i DISTINCT | 9-10 cwiczen bez pomocy |

### Co dalej?
- Jesli poziom **Podstawowy**: Powtorz cwiczenia 22.1-22.2, przejdz do `cheatsheet_sql.md` (sekcja JOIN).
- Jesli poziom **Dobry**: Przejdz do cwiczen `20_sql_group_by.md` i `21_sql_podzapytania.md`.
- Jesli poziom **Bardzo dobry**: Sprobuj cwiczen 22.8-22.10 bez wskazowek, potem przejdz do cwiczen GROUP BY (20.8-20.10).
- Jesli poziom **Doskonaly**: Przejdz do arkuszy maturalnych — zacznij od 2023+.
