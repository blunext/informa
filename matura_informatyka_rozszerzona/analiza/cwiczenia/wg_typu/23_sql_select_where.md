# 23. SQL — SELECT z WHERE

Typ zadania: **sql_select_where**
Czestotliwosc: 4/11 lat | Laczna punktacja: 10 pkt
Kategoria: SQL

---

### Cwiczenie 23.1 (trudnosc: latwe, ~1 pkt)
**Zrodlo inspiracji**: Matura 2023 (Gry planszowe)

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

**Wyjasnienie**: WHERE laczy dwa warunki operatorem AND — oba musza byc spelnione jednoczesnie. Operator `>=` oznacza "wieksze lub rowne". Typowy blad maturalny: uzycie `>` zamiast `>=`, przez co Tomek Bak (srednia rowna dokladnie 4.5) zostalby pominiiety.
</details>

---

### Cwiczenie 23.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2016 (Uniwersytet)

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
- Polaczone (bez duplikatow): Kabel USB-C, Klawiatura mechaniczna, Ladowarka szybka, Mysz bezprzewodowa

| nazwa | cena |
|-------|------|
| Kabel USB-C | 29 |
| Klawiatura mechaniczna | 349 |
| Ladowarka szybka | 79 |
| Mysz bezprzewodowa | 89 |

**Wyjasnienie**: Operator LIKE pozwala na wyszukiwanie wzorcow: `%` zastepuje dowolny ciag znakow (rowniez pusty), `_` zastepuje dokladnie jeden znak. `'K%'` = zaczyna sie od K, `'%a'` = konczy sie na a. OR oznacza, ze wystarczy spelnienie jednego z warunkow. LIKE rozroznia wielkosc liter w niektorych bazach danych (np. MySQL domyslnie nie rozroznia, SQLite domyslnie rozroznia dla ASCII). Typowy blad maturalny: uzycie `=` zamiast LIKE z wzorcami — `nazwa = 'K%'` szukalyby dokladnego tekstu "K%".
</details>

---

### Cwiczenie 23.3 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2024 (Rejestr wykroczen)

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
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT marka, model, rok_produkcji
FROM Pojazdy
WHERE rok_produkcji BETWEEN 2018 AND 2021
  AND pojemnosc IN (1400, 1600, 1800, 2000)
  AND nr_rejestracyjny NOT LIKE 'WA%';
```

**Weryfikacja** (sprawdzamy kazdy warunek):

| id | marka | model | rok 2018-2021? | pojemnosc w zbiorze? | nie WA? | Wynik |
|----|-------|-------|---------------|---------------------|---------|-------|
| 1 | Toyota | Corolla | 2019 ✓ | 1800 ✓ | WA ✗ | ✗ |
| 2 | Ford | Focus | 2017 ✗ | — | — | ✗ |
| 3 | BMW | 320i | 2021 ✓ | 2000 ✓ | GD ✓ | ✓ |
| 4 | Fiat | 500 | 2020 ✓ | 1200 ✗ | — | ✗ |
| 5 | VW | Golf | 2018 ✓ | 1400 ✓ | PO ✓ | ✓ |
| 6 | Audi | A4 | 2022 ✗ | — | — | ✗ |
| 7 | Toyota | Yaris | 2016 ✗ | — | — | ✗ |
| 8 | Skoda | Octavia | 2019 ✓ | 1600 ✓ | GD ✓ | ✓ |
| 9 | Honda | Civic | 2020 ✓ | 1500 ✗ | — | ✗ |
| 10 | Renault | Clio | 2015 ✗ | — | — | ✗ |

| marka | model | rok_produkcji |
|-------|-------|---------------|
| BMW | 320i | 2021 |
| Volkswagen | Golf | 2018 |
| Skoda | Octavia | 2019 |

**Wyjasnienie**: BETWEEN a AND b jest rownowazne `>= a AND <= b` (wlacznie obu granic). IN (lista) sprawdza przynaleznosc do zbioru — alternatywa dla wielu OR-ow. NOT LIKE neguje wzorzec. Typowy blad maturalny: zapomnienie, ze BETWEEN jest wlaczajace (inclusive) — `BETWEEN 2018 AND 2021` obejmuje 2018 i 2021.
</details>

---

### Cwiczenie 23.4 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2016 (Uniwersytet — LENGTH)

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
   - 'Ida' = 3 (✗)
   - 'Kler' = 4 (✗)
   - 'Parasite' = 8 (✗)
   - 'Zimna wojna' = 11 (>10 ✓)
   - 'Boze Cialo' = 10 (>10 ✗, rowne 10, nie wiecej)
   - 'Quo Vadis' = 9 (✗)
   - 'Komornik' = 8 (✗)

| tytul | dlugosc |
|-------|---------|
| Ogniem i mieczem | 17 |
| Zimna wojna | 11 |

2. Filmy Pawlikowskiego:

| tytul | rezyser | poczatek |
|-------|---------|----------|
| Ida | Pawlikowski | Ida |
| Zimna wojna | Pawlikowski | Zim |

**Wyjasnienie**: LENGTH(tekst) zwraca liczbe znakow w tekscie (wliczajac spacje). SUBSTR(tekst, pozycja, dlugosc) wyciaga fragment tekstu — pozycje liczymy od 1 (nie od 0 jak w C++!). Mozna uzyc LENGTH w WHERE do filtrowania wg dlugosci. Typowy blad maturalny: pomylenie SUBSTR z indeksowaniem od 0 — w SQL pierwszy znak ma pozycje 1. Rowniez: 'Boze Cialo' ma dokladnie 10 znakow, wiec warunek `> 10` go nie obejmuje (trzeba uwazac na granice warunkow).
</details>

---

### Cwiczenie 23.5 (trudnosc: trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2019 (Perfumy)

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

**Weryfikacja** (krok po kroku):

| id | stacja | opady NOT NULL? | opady > 0? | wiatr>10 OR ampl>15? | amplituda | Wynik |
|----|--------|----------------|-----------|---------------------|-----------|-------|
| 1 | Warszawa | ✓ (0) | ✗ (0) | — | — | ✗ |
| 2 | Krakow | ✓ (3) | ✓ | wiatr=10 ✗, ampl=7 ✗ | 7 | ✗ |
| 3 | Gdansk | ✗ (NULL) | — | — | — | ✗ |
| 4 | Warszawa | ✓ (0) | ✗ (0) | — | — | ✗ |
| 5 | Krakow | ✓ (12) | ✓ | wiatr=8 ✗, ampl=13 ✗ | 13 | ✗ |
| 6 | Gdansk | ✗ (NULL) | — | — | — | ✗ |
| 7 | Warszawa | ✓ (5) | ✓ | wiatr=12 >10 ✓ | 8 | ✓ |
| 8 | Krakow | ✓ (0) | ✗ (0) | — | — | ✗ |
| 9 | Gdansk | ✓ (8) | ✓ | wiatr=25 >10 ✓ | 8 | ✓ |
| 10 | Warszawa | ✓ (15) | ✓ | wiatr=22 >10 ✓ | 7 | ✓ |

| stacja | data | temp_min | temp_max | amplituda |
|--------|------|----------|----------|-----------|
| Warszawa | 2024-03-10 | 2 | 10 | 8 |
| Gdansk | 2024-03-10 | 3 | 11 | 8 |
| Warszawa | 2024-12-05 | -10 | -3 | 7 |

**Wyjasnienie**: Cwiczenie testuje kilka kluczowych konceptow:
1. **IS NOT NULL** — jedyny poprawny sposob sprawdzenia, czy wartosc nie jest NULL. `opady != NULL` NIE zadziala (zwroci zawsze NULL/FALSE).
2. **Nawiasy w zlozonym WHERE** — `AND (... OR ...)` wymaga nawiasow, bo AND ma wyzszy priorytet niz OR. Bez nawiasow wyrazenie bylby interpretowane jako `(opady IS NOT NULL AND opady > 0 AND wiatr > 10) OR (amplituda > 15)`, co daloby inny wynik.
3. **Kolumna obliczana** — `(temp_max - temp_min) AS amplituda` tworzy nowa kolumne w wyniku, ale nie mozna uzyc aliasu `amplituda` w WHERE (w wielu bazach danych). W ORDER BY alias zazwyczaj dziala.

Typowy blad maturalny: `WHERE opady != NULL` zamiast `IS NOT NULL` — to najczestszy blad zwiazany z NULL w SQL.
</details>
