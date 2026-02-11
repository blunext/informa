# 23. SQL — SELECT z WHERE

Typ zadania: **sql_select_where**
Czestotliwosc: 4/11 lat | Laczna punktacja: 10 pkt
Kategoria: SQL

## Umiejetnosci cwiczone w tym zestawie

`SELECT` `WHERE` `AND` `OR` `LIKE` `BETWEEN` `IN` `NOT` `IS-NULL` `IS-NOT-NULL` `LENGTH` `SUBSTR` `ORDER-BY` `operatory-porownania` `obliczenia-w-SELECT` `DISTINCT` `CASE-WHEN`

---

### Cwiczenie 23.1 (trudnosc: latwe, ~1 pkt)
**Zrodlo inspiracji**: Matura 2023 (Gry planszowe)
**Tagi**: `SELECT` `WHERE` `AND` `operatory-porownania`

Tabela **Uczniowie**:

| id | imie | nazwisko | klasa | wiek | srednia_ocen |
|----|------|----------|-------|------|-------------|
| 1 | Anna | Kowalska | 3A | 18 | 4.8 |
| 2 | Jan | Nowak | 2B | 17 | 3.5 |
| 3 | Ewa | Maj | 3A | 18 | 5.0 |
| 4 | Piotr | Lis | 1C | 16 | 4.2 |
| 5 | Kasia | Zak | 2B | 17 | 3.9 |
| 6 | Tomek | Bak | 3A | 19 | 4.5 |
| 7 | Ola | Wrobel | 1C | 15 | 4.1 |
| 8 | Marek | Krol | 2B | 17 | 2.8 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli imie i nazwisko uczniow z klasy '3A', ktorych srednia ocen jest wieksza lub rowna 4.5.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dwa warunki do polaczenia — klasa i srednia ocen.
2. **Podejscie**: WHERE z dwoma warunkami polaczonymi AND.
3. **Kluczowy krok**: `>=` oznacza "wieksze lub rowne" — uwzglednia tez dokladnie 4.5.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT imie, nazwisko
FROM Uczniowie
WHERE klasa = '3A' AND srednia_ocen >= 4.5;
```

**Weryfikacja**:
- Anna Kowalska: klasa=3A ✓, srednia=4.8 ≥ 4.5 ✓
- Ewa Maj: klasa=3A ✓, srednia=5.0 ≥ 4.5 ✓
- Tomek Bak: klasa=3A ✓, srednia=4.5 ≥ 4.5 ✓

| imie | nazwisko |
|------|----------|
| Anna | Kowalska |
| Ewa | Maj |
| Tomek | Bak |

**Wyjasnienie**: WHERE laczy dwa warunki operatorem AND — oba musza byc spelnione jednoczesnie. Operator `>=` oznacza "wieksze lub rowne".
</details>

<details>
<summary>Typowe bledy</summary>

- **`>` zamiast `>=`**: Tomek Bak (srednia = 4.5) zostalby pominiiety. CKE: -1 pkt.
- **Brak apostrofow wokol '3A'**: `WHERE klasa = 3A` — blad skladni, tekst wymaga apostrofow. CKE: -1 pkt.

</details>

---

### Cwiczenie 23.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2016 (Uniwersytet)
**Tagi**: `SELECT` `WHERE` `LIKE` `OR` `ORDER-BY`

Tabela **Produkty**:

| id | nazwa | kategoria | producent | cena |
|----|-------|-----------|-----------|------|
| 1 | Laptop Pro 15 | elektronika | TechCorp | 4599 |
| 2 | Mysz bezprzewodowa | elektronika | ClickMax | 89 |
| 3 | Klawiatura mechaniczna | elektronika | KeyMaster | 349 |
| 4 | Monitor 27 cali | elektronika | ViewPro | 1299 |
| 5 | Sluchawki Bluetooth | elektronika | SoundWave | 199 |
| 6 | Pendrive 64GB | elektronika | DataStore | 39 |
| 7 | Kabel USB-C | akcesoria | CablePlus | 29 |
| 8 | Ladowarka szybka | akcesoria | PowerUp | 79 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwe i cene produktow, ktorych nazwa zaczyna sie od 'K' lub konczy sie na 'a'. Posortuj alfabetycznie.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dwa wzorce do sprawdzenia — poczatek i koniec nazwy.
2. **Podejscie**: LIKE z `%` — `'K%'` = zaczyna sie od K, `'%a'` = konczy sie na 'a'.
3. **Kluczowy krok**: OR laczy dwa warunki — wystarczy spelnienie jednego.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT nazwa, cena
FROM Produkty
WHERE nazwa LIKE 'K%' OR nazwa LIKE '%a'
ORDER BY nazwa;
```

**Weryfikacja**:
- Zaczyna sie od 'K': Klawiatura mechaniczna ✓, Kabel USB-C ✓
- Konczy sie na 'a': Mysz bezprzewodowa ✓, Klawiatura mechaniczna ✓, Ladowarka szybka ✓

| nazwa | cena |
|-------|------|
| Kabel USB-C | 29 |
| Klawiatura mechaniczna | 349 |
| Ladowarka szybka | 79 |
| Mysz bezprzewodowa | 89 |

**Wyjasnienie**: LIKE pozwala na wyszukiwanie wzorcow: `%` zastepuje dowolny ciag znakow, `_` zastepuje dokladnie jeden znak. OR oznacza, ze wystarczy spelnienie jednego z warunkow.
</details>

<details>
<summary>Typowe bledy</summary>

- **`=` zamiast LIKE**: `nazwa = 'K%'` szukalyby dokladnego tekstu "K%" — nie znajdzie zadnych wynikow. CKE: -2 pkt.
- **AND zamiast OR**: Z AND oba warunki musza byc spelnione — tylko Klawiatura mechaniczna (zaczyna sie od K i konczy na 'a'). CKE: -1 pkt.

</details>

---

### Cwiczenie 23.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 (Rejestr wykroczen)
**Tagi**: `WHERE` `BETWEEN` `IN` `NOT` `LIKE` `AND`

Tabela **Pojazdy**:

| id | marka | model | rok_produkcji | pojemnosc | kolor | nr_rejestracyjny |
|----|-------|-------|---------------|-----------|-------|-------------------|
| 1 | Toyota | Corolla | 2019 | 1800 | bialy | WA 12345 |
| 2 | Ford | Focus | 2017 | 1600 | czarny | KR 98765 |
| 3 | BMW | 320i | 2021 | 2000 | srebrny | GD 55555 |
| 4 | Fiat | 500 | 2020 | 1200 | czerwony | WA 67890 |
| 5 | Volkswagen | Golf | 2018 | 1400 | bialy | PO 11111 |
| 6 | Audi | A4 | 2022 | 2000 | czarny | WA 22222 |
| 7 | Toyota | Yaris | 2016 | 1000 | niebieski | KR 33333 |
| 8 | Skoda | Octavia | 2019 | 1600 | srebrny | GD 44444 |
| 9 | Honda | Civic | 2020 | 1500 | bialy | WA 55566 |
| 10 | Renault | Clio | 2015 | 900 | czerwony | PO 77788 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli marke, model i rok produkcji pojazdow spelniajacych WSZYSTKIE warunki:
- rok produkcji miedzy 2018 a 2021 (wlacznie)
- pojemnosc silnika w zbiorze {1400, 1600, 1800, 2000}
- nr rejestracyjny NIE zaczyna sie od 'WA'

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy warunki polaczone AND — wszystkie musza byc spelnione.
2. **Podejscie**: BETWEEN dla zakresu, IN dla zbioru wartosci, NOT LIKE dla wykluczenia wzorca.
3. **Kluczowy krok**: BETWEEN jest wlaczajace (inclusive) — `BETWEEN 2018 AND 2021` obejmuje 2018 i 2021.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT marka, model, rok_produkcji
FROM Pojazdy
WHERE rok_produkcji BETWEEN 2018 AND 2021
  AND pojemnosc IN (1400, 1600, 1800, 2000)
  AND nr_rejestracyjny NOT LIKE 'WA%';
```

**Weryfikacja**:

| id | marka | model | rok 2018-2021? | pojemnosc w zbiorze? | nie WA? | Wynik |
|----|-------|-------|---------------|---------------------|---------|-------|
| 1 | Toyota | Corolla | 2019 ✓ | 1800 ✓ | WA ✗ | ✗ |
| 3 | BMW | 320i | 2021 ✓ | 2000 ✓ | GD ✓ | ✓ |
| 5 | VW | Golf | 2018 ✓ | 1400 ✓ | PO ✓ | ✓ |
| 8 | Skoda | Octavia | 2019 ✓ | 1600 ✓ | GD ✓ | ✓ |

| marka | model | rok_produkcji |
|-------|-------|---------------|
| BMW | 320i | 2021 |
| Volkswagen | Golf | 2018 |
| Skoda | Octavia | 2019 |

**Wyjasnienie**: BETWEEN a AND b jest rownowazne `>= a AND <= b`. IN sprawdza przynaleznosc do zbioru. NOT LIKE neguje wzorzec.
</details>

<details>
<summary>Typowe bledy</summary>

- **Zapomnienie, ze BETWEEN jest wlaczajace**: Zarowno 2018 jak i 2021 sa uwzgledniane. CKE: informacyjne — nie blad, ale warto wiedziec.
- **OR zamiast AND**: Zmienia logike — wystarczyloby spelnienie jednego warunku. CKE: -2 pkt.
- **Brak NOT przed LIKE**: Bez NOT warunek wybiera pojazdy z WA zamiast je wykluczac. CKE: -2 pkt.

</details>

---

### Cwiczenie 23.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2016 (Uniwersytet — LENGTH)
**Tagi**: `LENGTH` `SUBSTR` `WHERE` `operatory-porownania`

Tabela **Filmy**:

| id | tytul | rezyser | rok | ocena | kraj |
|----|-------|---------|-----|-------|------|
| 1 | Ogniem i mieczem | Hoffman | 1999 | 7.2 | Polska |
| 2 | Ida | Pawlikowski | 2013 | 7.8 | Polska |
| 3 | Kler | Smarzowski | 2018 | 6.5 | Polska |
| 4 | Parasite | Bong | 2019 | 8.6 | Korea |
| 5 | Zimna wojna | Pawlikowski | 2018 | 7.9 | Polska |
| 6 | Boze Cialo | Komasa | 2019 | 7.6 | Polska |
| 7 | Quo Vadis | Kawalerowicz | 2001 | 5.8 | Polska |
| 8 | Komornik | Smarzowski | 2005 | 6.9 | Polska |

**Polecenie**:
1. Napisz zapytanie, ktore wyswietli tytul i dlugosc tytulu (funkcja LENGTH) dla filmow, ktorych tytul ma wiecej niz 10 znakow.
2. Napisz zapytanie, ktore wyswietli tytul, rezyser oraz pierwsze 3 litery tytulu (funkcja SUBSTR) dla filmow rezysera 'Pawlikowski'.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Funkcje tekstowe LENGTH i SUBSTR dzialaja na stringach.
2. **Podejscie**: LENGTH(tekst) zwraca liczbe znakow. SUBSTR(tekst, start, dlugosc) wyciaga fragment.
3. **Kluczowy krok**: W SQL pozycje w SUBSTR liczymy od 1 (nie od 0 jak w C++).

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytania SQL:**

```sql
-- 1. Filmy z dlugim tytulem
SELECT tytul, LENGTH(tytul) AS dlugosc
FROM Filmy
WHERE LENGTH(tytul) > 10;

-- 2. Filmy Pawlikowskiego z poczatkiem tytulu
SELECT tytul, rezyser, SUBSTR(tytul, 1, 3) AS poczatek
FROM Filmy
WHERE rezyser = 'Pawlikowski';
```

**Weryfikacja**:

1. Dlugosci tytulow:
   - 'Ogniem i mieczem' = 17 znakow (>10 ✓)
   - 'Ida' = 3 (✗), 'Kler' = 4 (✗), 'Parasite' = 8 (✗)
   - 'Zimna wojna' = 11 (>10 ✓)
   - 'Boze Cialo' = 10 (>10 ✗, rowne 10, nie wiecej)

| tytul | dlugosc |
|-------|---------|
| Ogniem i mieczem | 17 |
| Zimna wojna | 11 |

2. Filmy Pawlikowskiego:

| tytul | rezyser | poczatek |
|-------|---------|----------|
| Ida | Pawlikowski | Ida |
| Zimna wojna | Pawlikowski | Zim |

**Wyjasnienie**: LENGTH liczy znaki (wliczajac spacje). SUBSTR(tekst, 1, 3) wyciaga znaki od pozycji 1 do 3. Pozycje w SQL liczymy od 1. 'Boze Cialo' ma dokladnie 10 znakow — `> 10` go nie obejmuje.
</details>

<details>
<summary>Typowe bledy</summary>

- **SUBSTR od pozycji 0**: `SUBSTR(tytul, 0, 3)` — w SQLite dziala (traktuje 0 jak 1), ale w standardzie SQL to blad. CKE: -1 pkt.
- **`>= 10` zamiast `> 10`**: Dodaje 'Boze Cialo' (10 znakow) do wyniku, niezgodnie z poleceniem "wiecej niz 10". CKE: -1 pkt.

</details>

---

### Cwiczenie 23.5 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 (Perfumy)
**Tagi**: `IS-NOT-NULL` `IS-NULL` `AND` `OR` `obliczenia-w-SELECT` `ORDER-BY`

Tabela **Pomiary**:

| id | stacja | data | temp_min | temp_max | opady | wiatr |
|----|--------|------|----------|----------|-------|-------|
| 1 | Warszawa | 2024-01-15 | -5 | 2 | 0 | 15 |
| 2 | Krakow | 2024-01-15 | -8 | -1 | 3 | 10 |
| 3 | Gdansk | 2024-01-15 | -2 | 4 | NULL | 20 |
| 4 | Warszawa | 2024-07-20 | 18 | 32 | 0 | 5 |
| 5 | Krakow | 2024-07-20 | 16 | 29 | 12 | 8 |
| 6 | Gdansk | 2024-07-20 | 15 | 25 | NULL | 18 |
| 7 | Warszawa | 2024-03-10 | 2 | 10 | 5 | 12 |
| 8 | Krakow | 2024-03-10 | 0 | 8 | 0 | 7 |
| 9 | Gdansk | 2024-03-10 | 3 | 11 | 8 | 25 |
| 10 | Warszawa | 2024-12-05 | -10 | -3 | 15 | 22 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli stacje, date, temperature minimalna i maksymalna oraz roznice temperatur (temp_max - temp_min) jako kolumne `amplituda` — ale tylko dla pomiarow spelniajacych WSZYSTKIE warunki:
- opady nie sa NULL
- opady sa wieksze od 0
- wiatr jest wiekszy niz 10 LUB amplituda temperatur (temp_max - temp_min) jest wieksza niz 15

Posortuj malejaco wg amplitudy.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy warunki, ale trzeci jest zlozony (OR wewnatrz AND).
2. **Podejscie**: IS NOT NULL + warunek na opady + nawiasy wokol OR.
3. **Kluczowy krok**: Nawiasy sa kluczowe: `AND (wiatr > 10 OR amplituda > 15)`. Bez nawiasow AND ma wyzszy priorytet niz OR.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT stacja, data, temp_min, temp_max,
       (temp_max - temp_min) AS amplituda
FROM Pomiary
WHERE opady IS NOT NULL
  AND opady > 0
  AND (wiatr > 10 OR (temp_max - temp_min) > 15)
ORDER BY amplituda DESC;
```

**Weryfikacja**:

| id | stacja | opady NOT NULL? | opady > 0? | wiatr>10 OR ampl>15? | amplituda | Wynik |
|----|--------|----------------|-----------|---------------------|-----------|-------|
| 2 | Krakow | ✓ (3) | ✓ | wiatr=10 ✗, ampl=7 ✗ | 7 | ✗ |
| 5 | Krakow | ✓ (12) | ✓ | wiatr=8 ✗, ampl=13 ✗ | 13 | ✗ |
| 7 | Warszawa | ✓ (5) | ✓ | wiatr=12 >10 ✓ | 8 | ✓ |
| 9 | Gdansk | ✓ (8) | ✓ | wiatr=25 >10 ✓ | 8 | ✓ |
| 10 | Warszawa | ✓ (15) | ✓ | wiatr=22 >10 ✓ | 7 | ✓ |

| stacja | data | temp_min | temp_max | amplituda |
|--------|------|----------|----------|-----------|
| Warszawa | 2024-03-10 | 2 | 10 | 8 |
| Gdansk | 2024-03-10 | 3 | 11 | 8 |
| Warszawa | 2024-12-05 | -10 | -3 | 7 |

**Wyjasnienie**: Kluczowe koncepty:
1. **IS NOT NULL** — jedyny poprawny sposob sprawdzenia, czy wartosc nie jest NULL.
2. **Nawiasy w zlozonym WHERE** — `AND (... OR ...)` wymaga nawiasow, bo AND ma wyzszy priorytet niz OR.
3. **Kolumna obliczana** — `(temp_max - temp_min) AS amplituda` tworzy nowa kolumne w wyniku.
</details>

<details>
<summary>Typowe bledy</summary>

- **`!= NULL` zamiast `IS NOT NULL`**: `opady != NULL` zawsze zwraca NULL/FALSE — nic nie przejdzie tego filtra. CKE: -2 pkt.
- **Brak nawiasow wokol OR**: Bez nawiasow wyrazenie jest interpretowane inaczej z powodu priorytetu AND > OR. CKE: -2 pkt.
- **Alias w WHERE**: `WHERE amplituda > 15` — w wielu bazach nie mozna uzyc aliasu w WHERE. Trzeba powtorzyc wyrazenie. CKE: -1 pkt.

</details>

---

### Cwiczenie 23.6 (trudnosc: latwe, ~1 pkt)
**Zrodlo inspiracji**: Matura 2025 (Woda na Marsie)
**Tagi**: `SELECT` `WHERE` `ORDER-BY` `operatory-porownania`

Tabela **Planety**:

| id | nazwa | srednica_km | odleglosc_au | liczba_ksiezycow | typ |
|----|-------|-------------|-------------|-----------------|-----|
| 1 | Merkury | 4879 | 0.39 | 0 | skalista |
| 2 | Wenus | 12104 | 0.72 | 0 | skalista |
| 3 | Ziemia | 12756 | 1.00 | 1 | skalista |
| 4 | Mars | 6792 | 1.52 | 2 | skalista |
| 5 | Jowisz | 142984 | 5.20 | 95 | gazowa |
| 6 | Saturn | 120536 | 9.54 | 146 | gazowa |
| 7 | Uran | 51118 | 19.19 | 27 | lodowa |
| 8 | Neptun | 49528 | 30.07 | 16 | lodowa |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwe i liczbe ksiezycow planet, ktore maja wiecej niz 10 ksiezycow. Posortuj malejaco wg liczby ksiezycow.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Prosty filtr WHERE z jednym warunkiem.
2. **Podejscie**: `WHERE liczba_ksiezycow > 10` + ORDER BY DESC.
3. **Kluczowy krok**: Pamietaj, ze `> 10` nie obejmuje 10.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT nazwa, liczba_ksiezycow
FROM Planety
WHERE liczba_ksiezycow > 10
ORDER BY liczba_ksiezycow DESC;
```

**Weryfikacja**:
- Saturn: 146 > 10 ✓
- Jowisz: 95 > 10 ✓
- Uran: 27 > 10 ✓
- Neptun: 16 > 10 ✓

| nazwa | liczba_ksiezycow |
|-------|-----------------|
| Saturn | 146 |
| Jowisz | 95 |
| Uran | 27 |
| Neptun | 16 |

**Wyjasnienie**: Prosty SELECT z WHERE i ORDER BY DESC. To podstawowy wzorzec filtrowania danych w SQL.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak DESC w ORDER BY**: Domyslnie ORDER BY sortuje rosnaco (ASC). CKE: -0.5 pkt.
- **`>= 10` zamiast `> 10`**: Gdyby jakas planeta miala 10 ksiezycow, bylaby wlaczona blednie. CKE: -0.5 pkt (zalezy od danych).

</details>

---

### Cwiczenie 23.7 (trudnosc: srednie, ~2 pkt)
**Zrodlo inspiracji**: Matura 2017 (Pilka reczna)
**Tagi**: `WHERE` `LIKE` `NOT` `AND` `OR`

Tabela **Miasta**:

| id | nazwa | kraj | populacja | powierzchnia_km2 | stolica |
|----|-------|------|-----------|------------------|---------|
| 1 | Warszawa | Polska | 1790658 | 517 | tak |
| 2 | Krakow | Polska | 779115 | 327 | nie |
| 3 | Berlin | Niemcy | 3644826 | 892 | tak |
| 4 | Monachium | Niemcy | 1471508 | 311 | nie |
| 5 | Praga | Czechy | 1309000 | 496 | tak |
| 6 | Brno | Czechy | 382405 | 230 | nie |
| 7 | Gdansk | Polska | 470907 | 262 | nie |
| 8 | Wroclaw | Polska | 641928 | 293 | nie |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwe i populacje miast spelniajacych warunki:
- nazwa NIE konczy sie na 'w'
- populacja > 500000
- miasto jest z Polski LUB jest stolica

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy warunki, trzeci jest zlozony (OR).
2. **Podejscie**: NOT LIKE '%w' + warunek na populacje + nawiasy wokol OR.
3. **Kluczowy krok**: Nawiasy: `AND (kraj = 'Polska' OR stolica = 'tak')`.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT nazwa, populacja
FROM Miasta
WHERE nazwa NOT LIKE '%w'
  AND populacja > 500000
  AND (kraj = 'Polska' OR stolica = 'tak');
```

**Weryfikacja**:

| id | nazwa | NOT LIKE '%w'? | pop > 500k? | Polska OR stolica? | Wynik |
|----|-------|----------------|-------------|-------------------|-------|
| 1 | Warszawa | ✓ (konczy na 'a') | ✓ (1.79M) | Polska ✓ | ✓ |
| 3 | Berlin | ✓ (konczy na 'n') | ✓ (3.64M) | stolica ✓ | ✓ |
| 4 | Monachium | ✓ (konczy na 'm') | ✓ (1.47M) | ani Polska, ani stolica ✗ | ✗ |
| 5 | Praga | ✓ (konczy na 'a') | ✓ (1.31M) | stolica ✓ | ✓ |
| 2 | Krakow | ✗ (konczy na 'w') | — | — | ✗ |
| 8 | Wroclaw | ✗ (konczy na 'w') | — | — | ✗ |

| nazwa | populacja |
|-------|-----------|
| Warszawa | 1790658 |
| Berlin | 3644826 |
| Praga | 1309000 |

**Wyjasnienie**: NOT LIKE '%w' wyklucza nazwy konczace sie na 'w' (Krakow, Wroclaw). Nawiasy wokol OR zapewniaja poprawna logike — bez nich AND bedzie mialo priorytet.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak nawiasow wokol OR**: `AND kraj = 'Polska' OR stolica = 'tak'` zmienia logike — Praga przejdzie nawet bez spelnienia warunkow populacji i NOT LIKE. CKE: -2 pkt.
- **LIKE zamiast NOT LIKE**: Odwrotny wynik — tylko miasta konczace sie na 'w'. CKE: -2 pkt.

</details>

---

### Cwiczenie 23.8 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2022 (Ewidencja uczniow)
**Tagi**: `DISTINCT` `WHERE` `BETWEEN` `LIKE` `ORDER-BY`

Tabela **Loty**:

| id | nr_lotu | skad | dokad | data | godzina | cena | status |
|----|---------|------|-------|------|---------|------|--------|
| 1 | LO101 | Warszawa | Berlin | 2024-06-15 | 08:30 | 450 | planowy |
| 2 | LO102 | Warszawa | Paryz | 2024-06-15 | 10:00 | 680 | planowy |
| 3 | FR201 | Krakow | Londyn | 2024-06-16 | 06:15 | 320 | planowy |
| 4 | LO103 | Warszawa | Berlin | 2024-06-16 | 08:30 | 450 | opozniony |
| 5 | FR202 | Gdansk | Berlin | 2024-06-17 | 14:00 | 280 | planowy |
| 6 | LO104 | Warszawa | Rzym | 2024-06-17 | 12:00 | 550 | odwolany |
| 7 | FR203 | Krakow | Paryz | 2024-06-18 | 09:45 | 390 | planowy |
| 8 | LO105 | Warszawa | Londyn | 2024-06-18 | 16:30 | 720 | planowy |

**Polecenie**:
1. Wyswietl UNIKALNE miasta docelowe (DISTINCT) posortowane alfabetycznie.
2. Wyswietl nr lotu, trase (skad → dokad) i cene lotow, ktore: maja nr lotu zaczynajacy sie od 'LO', kosztuja miedzy 400 a 600, i NIE sa odwolane.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Dwa osobne zapytania — jedno z DISTINCT, drugie z wieloma warunkami.
2. **Podejscie**: DISTINCT eliminuje duplikaty. BETWEEN, LIKE i != laczymy AND-em.
3. **Kluczowy krok**: `status != 'odwolany'` lub `status <> 'odwolany'` — oba sa poprawne.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytania SQL:**
```sql
-- 1. Unikalne miasta docelowe
SELECT DISTINCT dokad
FROM Loty
ORDER BY dokad;

-- 2. Loty LOT w przedziale cenowym, nieodwolane
SELECT nr_lotu, skad || ' → ' || dokad AS trasa, cena
FROM Loty
WHERE nr_lotu LIKE 'LO%'
  AND cena BETWEEN 400 AND 600
  AND status != 'odwolany';
```

**Weryfikacja**:

1. Unikalne miasta docelowe: Berlin, Londyn, Paryz, Rzym

| dokad |
|-------|
| Berlin |
| Londyn |
| Paryz |
| Rzym |

2. Loty LO, cena 400-600, nie odwolane:
- LO101: LO ✓, 450 ∈ [400,600] ✓, planowy ✓ → ✓
- LO102: LO ✓, 680 ∉ [400,600] ✗
- LO103: LO ✓, 450 ∈ [400,600] ✓, opozniony ✓ → ✓
- LO104: LO ✓, 550 ∈ [400,600] ✓, odwolany ✗
- LO105: LO ✓, 720 ∉ [400,600] ✗

| nr_lotu | trasa | cena |
|---------|-------|------|
| LO101 | Warszawa → Berlin | 450 |
| LO103 | Warszawa → Berlin | 450 |

**Wyjasnienie**: DISTINCT eliminuje powtorzenia w wynikach. Operator `||` w SQLite laczy teksty (konkatenacja). `!=` i `<>` sa rownowazne.
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak DISTINCT**: Berlin pojawi sie 3 razy (loty 1, 4, 5). CKE: -1 pkt.
- **`= 'odwolany'` zamiast `!= 'odwolany'`**: Odwrotny filtr — pokaze tylko odwolane loty. CKE: -2 pkt.
- **Konkatenacja z `+`**: W SQLite operator `+` dodaje liczby, nie laczy teksty. Trzeba uzyc `||`. CKE: -1 pkt.

</details>

---

### Cwiczenie 23.9 (trudnosc: srednie-trudne, ~3 pkt)
**Zrodlo inspiracji**: Matura 2021 (Gra strategiczna)
**Tagi**: `CASE-WHEN` `obliczenia-w-SELECT` `WHERE` `ORDER-BY`

Tabela **Wyniki_egzaminu**:

| id | imie | nazwisko | teoria | praktyka | projekt |
|----|------|----------|--------|----------|---------|
| 1 | Anna | Maj | 38 | 42 | 15 |
| 2 | Jan | Krol | 25 | 30 | 10 |
| 3 | Ewa | Lis | 45 | 48 | 18 |
| 4 | Piotr | Zak | 30 | 35 | 12 |
| 5 | Kasia | Wrobel | 20 | 25 | 8 |
| 6 | Tomek | Bak | 42 | 45 | 17 |
| 7 | Ola | Rak | 35 | 38 | 14 |
| 8 | Marek | Duda | 15 | 20 | 5 |

Maksymalne punkty: teoria 50, praktyka 50, projekt 20. Lacznie 120.

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli imie, nazwisko, sume punktow oraz ocene slowna wg schematu:
- suma >= 96 (80%): 'celujacy'
- suma >= 72 (60%): 'bardzo dobry'
- suma >= 48 (40%): 'dostateczny'
- suma < 48: 'niedostateczny'

Posortuj malejaco wg sumy.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Obliczenia w SELECT i CASE WHEN do klasyfikacji.
2. **Podejscie**: `(teoria + praktyka + projekt) AS suma` i CASE WHEN z progami.
3. **Kluczowy krok**: CASE WHEN sprawdza warunki PO KOLEI — pierwszy spelniony wygrywa. Zacznij od najwyzszego progu.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT imie, nazwisko,
       (teoria + praktyka + projekt) AS suma,
       CASE
         WHEN (teoria + praktyka + projekt) >= 96 THEN 'celujacy'
         WHEN (teoria + praktyka + projekt) >= 72 THEN 'bardzo dobry'
         WHEN (teoria + praktyka + projekt) >= 48 THEN 'dostateczny'
         ELSE 'niedostateczny'
       END AS ocena
FROM Wyniki_egzaminu
ORDER BY suma DESC;
```

**Weryfikacja**:
- Ewa: 45+48+18 = **111** → celujacy (>=96)
- Tomek: 42+45+17 = **104** → celujacy (>=96)
- Anna: 38+42+15 = **95** → bardzo dobry (>=72, <96)
- Ola: 35+38+14 = **87** → bardzo dobry (>=72)
- Piotr: 30+35+12 = **77** → bardzo dobry (>=72)
- Jan: 25+30+10 = **65** → dostateczny (>=48, <72)
- Kasia: 20+25+8 = **53** → dostateczny (>=48)
- Marek: 15+20+5 = **40** → niedostateczny (<48)

| imie | nazwisko | suma | ocena |
|------|----------|------|-------|
| Ewa | Lis | 111 | celujacy |
| Tomek | Bak | 104 | celujacy |
| Anna | Maj | 95 | bardzo dobry |
| Ola | Rak | 87 | bardzo dobry |
| Piotr | Zak | 77 | bardzo dobry |
| Jan | Krol | 65 | dostateczny |
| Kasia | Wrobel | 53 | dostateczny |
| Marek | Duda | 40 | niedostateczny |

**Wyjasnienie**: CASE WHEN dziala jak if/else — sprawdza warunki po kolei, pierwszy spelniony determinuje wynik. Dlatego kolejnosc warunkow jest wazna (od najwyzszego progu do najnizszego). ELSE obsluguje wszystkie przypadki nie spelniajace zadnego warunku.
</details>

<details>
<summary>Typowe bledy</summary>

- **Odwrotna kolejnosc warunkow**: Jesli pierwszy warunek to `>= 48`, WSZYSTKIE osoby poza Markiem dostana 'dostateczny'. CKE: -3 pkt.
- **Brak ELSE**: Osoby z suma < 48 dostana NULL zamiast 'niedostateczny'. CKE: -1 pkt.
- **Alias w CASE WHEN**: `WHEN suma >= 96` — w wielu bazach alias nie jest dostepny w tej samej klauzuli SELECT. Trzeba powtorzyc wyrazenie. CKE: -1 pkt.

</details>

---

### Cwiczenie 23.10 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2024 (Rejestr wykroczen)
**Tagi**: `IS-NULL` `IS-NOT-NULL` `CASE-WHEN` `obliczenia-w-SELECT` `SUBSTR` `LENGTH` `LIKE` `ORDER-BY`

Tabela **Zgloszenia**:

| id | nr_sprawy | kategoria | data_zgloszenia | data_zamkniecia | priorytet | opis |
|----|-----------|-----------|-----------------|-----------------|-----------|------|
| 1 | ZG-2024-001 | awaria | 2024-01-10 | 2024-01-12 | wysoki | Wyciek wody |
| 2 | ZG-2024-002 | reklamacja | 2024-01-15 | NULL | niski | Uszkodzony produkt |
| 3 | ZG-2024-003 | awaria | 2024-02-01 | 2024-02-03 | wysoki | Brak pradu |
| 4 | ZG-2024-004 | zapytanie | 2024-02-10 | 2024-02-10 | niski | Pytanie o cene |
| 5 | ZG-2024-005 | awaria | 2024-03-01 | NULL | wysoki | Zepsuta rura |
| 6 | ZG-2024-006 | reklamacja | 2024-03-05 | 2024-03-15 | sredni | Bledne zamowienie |
| 7 | ZG-2024-007 | zapytanie | 2024-03-10 | 2024-03-10 | niski | Godziny otwarcia |
| 8 | ZG-2024-008 | awaria | 2024-04-01 | NULL | wysoki | Brak ogrzewania |
| 9 | ZG-2024-009 | reklamacja | 2024-04-10 | 2024-04-20 | sredni | Wadliwy towar |
| 10 | ZG-2024-010 | zapytanie | 2024-04-15 | 2024-04-15 | niski | Regulamin |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli:
- nr sprawy
- kategorie
- miesiac zgloszenia (numer, uzyj SUBSTR)
- status: 'otwarte' jezeli data_zamkniecia IS NULL, 'zamkniete' w przeciwnym razie
- dlugosc opisu (LENGTH)

Pokaz tylko zgloszenia, ktore sa OTWARTE lub ktorych opis ma wiecej niz 15 znakow. Posortuj wg miesiaca, potem wg priorytetu (wysoki najpierw — uzyj CASE WHEN w ORDER BY).

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: CASE WHEN w SELECT (status) i w ORDER BY (priorytet).
2. **Podejscie**: CASE WHEN data_zamkniecia IS NULL THEN 'otwarte' ELSE 'zamkniete'. SUBSTR(data, 6, 2) wyciaga miesiac.
3. **Kluczowy krok**: ORDER BY CASE WHEN priorytet = 'wysoki' THEN 1 ... END — mapuje tekst na liczby do sortowania.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT nr_sprawy, kategoria,
       SUBSTR(data_zgloszenia, 6, 2) AS miesiac,
       CASE WHEN data_zamkniecia IS NULL THEN 'otwarte' ELSE 'zamkniete' END AS status,
       LENGTH(opis) AS dl_opisu
FROM Zgloszenia
WHERE data_zamkniecia IS NULL
   OR LENGTH(opis) > 15
ORDER BY miesiac,
         CASE priorytet
           WHEN 'wysoki' THEN 1
           WHEN 'sredni' THEN 2
           WHEN 'niski' THEN 3
         END;
```

**Weryfikacja**:

| id | nr_sprawy | otwarte? | dl_opisu > 15? | Wynik |
|----|-----------|----------|----------------|-------|
| 1 | ZG-2024-001 | zamkniete | 'Wyciek wody'=11 ✗ | ✗ |
| 2 | ZG-2024-002 | otwarte ✓ | — | ✓ |
| 3 | ZG-2024-003 | zamkniete | 'Brak pradu'=10 ✗ | ✗ |
| 4 | ZG-2024-004 | zamkniete | 'Pytanie o cene'=14 ✗ | ✗ |
| 5 | ZG-2024-005 | otwarte ✓ | — | ✓ |
| 6 | ZG-2024-006 | zamkniete | 'Bledne zamowienie'=17 ✓ | ✓ |
| 7 | ZG-2024-007 | zamkniete | 'Godziny otwarcia'=17 ✓ | ✓ |
| 8 | ZG-2024-008 | otwarte ✓ | — | ✓ |
| 9 | ZG-2024-009 | zamkniete | 'Wadliwy towar'=14 ✗ | ✗ |
| 10 | ZG-2024-010 | zamkniete | 'Regulamin'=9 ✗ | ✗ |

| nr_sprawy | kategoria | miesiac | status | dl_opisu |
|-----------|-----------|---------|--------|----------|
| ZG-2024-002 | reklamacja | 01 | otwarte | 19 |
| ZG-2024-005 | awaria | 03 | otwarte | 12 |
| ZG-2024-006 | reklamacja | 03 | zamkniete | 17 |
| ZG-2024-007 | zapytanie | 03 | zamkniete | 17 |
| ZG-2024-008 | awaria | 04 | otwarte | 16 |

**Wyjasnienie**: CASE WHEN w SELECT tworzy kolumne obliczana na podstawie warunku. CASE WHEN w ORDER BY pozwala mapowac wartosci tekstowe na liczby — dzieki temu 'wysoki' jest sortowany przed 'sredni' i 'niski'. SUBSTR i LENGTH to funkcje tekstowe uzyte jednoczesnie w roznych celach.
</details>

<details>
<summary>Typowe bledy</summary>

- **`= NULL` zamiast `IS NULL`**: Porownanie z NULL zawsze daje NULL. Trzeba uzyc IS NULL / IS NOT NULL. CKE: -2 pkt.
- **Sortowanie tekstowe po priorytecie**: Alfabetycznie 'niski' < 'sredni' < 'wysoki' — odwrotna kolejnosc niz zamierzona. CKE: -1 pkt.
- **Brak CASE WHEN w ORDER BY**: Nie ma innego sposobu na posortowanie priorytetu w niestandardowej kolejnosci bez dodatkowej kolumny. CKE: -1 pkt.

</details>

---

## Samoocena

| Poziom | Opis | Wymaganie |
|--------|------|-----------|
| Podstawowy | WHERE z AND/OR, operatory porownania, ORDER BY | 1-3 cwiczen bez pomocy |
| Dobry | LIKE, BETWEEN, IN, NOT, DISTINCT | 4-6 cwiczen bez pomocy |
| Bardzo dobry | IS NULL, LENGTH, SUBSTR, obliczenia w SELECT | 7-8 cwiczen bez pomocy |
| Doskonaly | CASE WHEN (w SELECT i ORDER BY), zlozone warunki logiczne | 9-10 cwiczen bez pomocy |

### Co dalej?
- Jesli poziom **Podstawowy**: Powtorz cwiczenia 23.1-23.2, przejdz do `cheatsheet_sql.md` (sekcja SELECT WHERE).
- Jesli poziom **Dobry**: Przejdz do cwiczen `20_sql_group_by.md` (proste GROUP BY).
- Jesli poziom **Bardzo dobry**: Sprobuj cwiczen `21_sql_podzapytania.md` i wrocz do 23.9-23.10 bez wskazowek.
- Jesli poziom **Doskonaly**: Przejdz do arkuszy maturalnych — zacznij od 2023+.
