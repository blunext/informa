# 20. SQL — GROUP BY z agregacjami

Typ zadania: **sql_group_by**
Czestotliwosc: 8/11 lat | Laczna punktacja: 36 pkt
Kategoria: SQL

## Umiejetnosci cwiczone w tym zestawie

`GROUP-BY` `COUNT` `SUM` `AVG` `MIN` `MAX` `ROUND` `HAVING` `WHERE-vs-HAVING` `JOIN-z-GROUP-BY` `SUBSTR` `ORDER-BY` `wielokolumnowe-grupowanie` `aliasy` `funkcje-daty`

---

### Cwiczenie 20.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 (Gry planszowe)
**Tagi**: `GROUP-BY` `COUNT` `ORDER-BY`

Tabela **Ksiazki** zawiera dane o ksiazkach w bibliotece szkolnej:

| id | tytul | autor | gatunek | rok_wydania |
|----|-------|-------|---------|-------------|
| 1 | Potop | Sienkiewicz | historyczna | 1886 |
| 2 | Lalka | Prus | realistyczna | 1890 |
| 3 | Quo Vadis | Sienkiewicz | historyczna | 1896 |
| 4 | Faraon | Prus | historyczna | 1897 |
| 5 | Krzyzacy | Sienkiewicz | historyczna | 1900 |
| 6 | Placowka | Prus | realistyczna | 1886 |
| 7 | Ogniem i mieczem | Sienkiewicz | historyczna | 1884 |
| 8 | Emancypantki | Prus | realistyczna | 1894 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli kazdy gatunek oraz liczbe ksiazek w tym gatunku. Posortuj malejaco wg liczby ksiazek.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Potrzebujesz pogrupowac ksiazki wg jednej kolumny i policzyc wiersze w kazdej grupie.
2. **Podejscie**: Uzyj GROUP BY na kolumnie gatunek, a do zliczania — funkcji COUNT(*).
3. **Kluczowy krok**: Pamietaj o ORDER BY z DESC, aby posortowac malejaco.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT gatunek, COUNT(*) AS liczba
FROM Ksiazki
GROUP BY gatunek
ORDER BY liczba DESC;
```

**Weryfikacja**:
- historyczna: Potop, Quo Vadis, Faraon, Krzyzacy, Ogniem i mieczem = **5**
- realistyczna: Lalka, Placowka, Emancypantki = **3**

| gatunek | liczba |
|---------|--------|
| historyczna | 5 |
| realistyczna | 3 |

**Wyjasnienie**: GROUP BY grupuje wiersze o tej samej wartosci w kolumnie `gatunek`. COUNT(*) zlicza wiersze w kazdej grupie. ORDER BY z DESC sortuje od najwiekszej wartosci.
</details>

<details>
<summary>Typowe bledy</summary>

- **Uzycie WHERE zamiast GROUP BY**: WHERE filtruje wiersze PRZED grupowaniem, a do filtrowania grup sluzy HAVING — tu jednak nie filtrujemy grup, wiec wystarczy GROUP BY. CKE: -2 pkt (brak wyniku)
- **Brak aliasu w COUNT(*)**: Bez `AS liczba` kolumna wynikowa moze miec niezrozumiala nazwe. CKE: dopuszczalne, ale -0.5 pkt za czytelnosc w niektorych schematach.

</details>

---

### Cwiczenie 20.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)
**Tagi**: `GROUP-BY` `COUNT` `AVG` `ROUND` `ORDER-BY`

Tabela **Produkty** zawiera dane o produktach w sklepie z odzieza:

| id | nazwa | kategoria | cena | rozmiar |
|----|-------|-----------|------|---------|
| 1 | Koszulka polo | meska | 89.99 | M |
| 2 | Sukienka letnia | damska | 149.99 | S |
| 3 | Spodnie jeans | meska | 179.99 | L |
| 4 | Bluzka jedwabna | damska | 199.99 | M |
| 5 | T-shirt basic | meska | 39.99 | S |
| 6 | Spodnica midi | damska | 129.99 | L |
| 7 | Koszula flanelowa | meska | 119.99 | M |
| 8 | Tunika | damska | 89.99 | S |
| 9 | Polo sportowe | meska | 69.99 | L |
| 10 | Sukienka koktajlowa | damska | 249.99 | M |

**Polecenie**: Napisz zapytanie SQL, ktore dla kazdej kategorii wyswietli: liczbe produktow, srednia cene zaokraglona do 2 miejsc po przecinku. Posortuj alfabetycznie wg kategorii.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Grupuj po kategorii i zastosuj dwie funkcje agregujace jednoczesnie.
2. **Podejscie**: COUNT(*) i AVG(cena) w jednym SELECT, pamietaj o ROUND.
3. **Kluczowy krok**: Skladnia ROUND to ROUND(wartosc, liczba_miejsc_dziesietnych).

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT kategoria, COUNT(*) AS liczba, ROUND(AVG(cena), 2) AS srednia_cena
FROM Produkty
GROUP BY kategoria
ORDER BY kategoria;
```

**Weryfikacja**:
- damska: Sukienka letnia (149.99), Bluzka jedwabna (199.99), Spodnica midi (129.99), Tunika (89.99), Sukienka koktajlowa (249.99) → 5 produktow, srednia = (149.99+199.99+129.99+89.99+249.99)/5 = 819.95/5 = **163.99**
- meska: Koszulka polo (89.99), Spodnie jeans (179.99), T-shirt basic (39.99), Koszula flanelowa (119.99), Polo sportowe (69.99) → 5 produktow, srednia = (89.99+179.99+39.99+119.99+69.99)/5 = 499.95/5 = **99.99**

| kategoria | liczba | srednia_cena |
|-----------|--------|-------------|
| damska | 5 | 163.99 |
| meska | 5 | 99.99 |

**Wyjasnienie**: ROUND(AVG(cena), 2) oblicza srednia i zaokragla do 2 miejsc po przecinku. Bez ROUND wynik moze miec wiele miejsc dziesietnych. Na maturze czesto wymagane jest zaokraglanie — warto pamietac skladnie ROUND(wartosc, liczba_miejsc).
</details>

<details>
<summary>Typowe bledy</summary>

- **Brak ROUND**: Wynik sredniej moze miec wiele cyfr po przecinku. CKE: -1 pkt jesli polecenie wymaga zaokraglenia.
- **ROUND z jednym argumentem**: `ROUND(AVG(cena))` zaokragla do calkowitej — nie do 2 miejsc. CKE: -1 pkt.

</details>

---

### Cwiczenie 20.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2022 (Ewidencja uczniow)
**Tagi**: `GROUP-BY` `AVG` `ROUND` `HAVING` `JOIN-z-GROUP-BY` `WHERE-vs-HAVING`

Tabela **Klasy**:

| id_klasy | nazwa_klasy | wychowawca |
|----------|-------------|------------|
| 1 | 1A | Kowalski |
| 2 | 1B | Nowak |
| 3 | 2A | Wisniewski |
| 4 | 2B | Lewandowska |

Tabela **Uczniowie**:

| id_ucznia | imie | nazwisko | id_klasy | srednia_ocen |
|-----------|------|----------|----------|-------------|
| 1 | Anna | Maj | 1 | 4.8 |
| 2 | Jan | Krol | 1 | 3.5 |
| 3 | Ewa | Lis | 2 | 5.0 |
| 4 | Piotr | Zak | 2 | 4.2 |
| 5 | Kasia | Wrobel | 3 | 4.6 |
| 6 | Tomek | Bak | 3 | 3.9 |
| 7 | Ola | Rak | 3 | 4.1 |
| 8 | Marek | Duda | 4 | 2.8 |
| 9 | Zofia | Ptak | 1 | 4.3 |
| 10 | Adam | Wilk | 4 | 3.6 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwe klasy oraz srednia ocen uczniow w tej klasie (zaokraglona do 2 miejsc), ale tylko dla klas, w ktorych srednia ocen jest wieksza niz 4.0.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Musisz polaczyc dwie tabele, pogrupowac i odfiltrowac grupy.
2. **Podejscie**: JOIN laczy Klasy z Uczniami, GROUP BY grupuje wg klasy, a filtrowanie grup wymaga HAVING (nie WHERE).
3. **Kluczowy krok**: HAVING AVG(...) > 4.0 — uzyj funkcji agregajacej bezposrednio w HAVING, nie aliasu.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT K.nazwa_klasy, ROUND(AVG(U.srednia_ocen), 2) AS srednia_klasy
FROM Klasy K
JOIN Uczniowie U ON K.id_klasy = U.id_klasy
GROUP BY K.nazwa_klasy
HAVING AVG(U.srednia_ocen) > 4.0
ORDER BY srednia_klasy DESC;
```

**Weryfikacja**:
- 1A: Anna (4.8), Jan (3.5), Zofia (4.3) → (4.8+3.5+4.3)/3 = 12.6/3 = **4.20** (>4.0 ✓)
- 1B: Ewa (5.0), Piotr (4.2) → (5.0+4.2)/2 = 9.2/2 = **4.60** (>4.0 ✓)
- 2A: Kasia (4.6), Tomek (3.9), Ola (4.1) → (4.6+3.9+4.1)/3 = 12.6/3 = **4.20** (>4.0 ✓)
- 2B: Marek (2.8), Adam (3.6) → (2.8+3.6)/2 = 6.4/2 = **3.20** (>4.0 ✗)

| nazwa_klasy | srednia_klasy |
|-------------|--------------|
| 1B | 4.60 |
| 1A | 4.20 |
| 2A | 4.20 |

**Wyjasnienie**: Kluczowa roznica: WHERE filtruje wiersze PRZED grupowaniem, HAVING filtruje grupy PO grupowaniu. Tutaj nie mozna uzyc `WHERE AVG(...) > 4.0`, bo AVG wymaga juz zgrupowanych danych.
</details>

<details>
<summary>Typowe bledy</summary>

- **WHERE zamiast HAVING**: `WHERE AVG(srednia_ocen) > 4.0` powoduje blad skladni — funkcje agregujace nie dzialaja w WHERE. CKE: -2 pkt (blad logiczny).
- **Alias w HAVING**: `HAVING srednia_klasy > 4.0` — w SQLite dziala, ale w wielu bazach nie. Bezpieczniej powtorzyc: `HAVING AVG(U.srednia_ocen) > 4.0`. CKE: akceptowane, ale ryzykowne.
- **Brak JOIN**: Proba grupowania bez polaczenia tabel — nie uzyskamy nazwy klasy. CKE: -2 pkt.

</details>

---

### Cwiczenie 20.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2017 (Pilka reczna)
**Tagi**: `GROUP-BY` `SUM` `HAVING` `JOIN-z-GROUP-BY`

Tabela **Druzyny**:

| id_druzyny | nazwa | miasto |
|------------|-------|--------|
| 1 | Orly | Warszawa |
| 2 | Wilki | Krakow |
| 3 | Rysie | Gdansk |
| 4 | Sokoly | Poznan |

Tabela **Mecze**:

| id_meczu | id_gospodarzy | id_gosci | bramki_gosp | bramki_gosci |
|----------|---------------|----------|-------------|-------------|
| 1 | 1 | 2 | 3 | 1 |
| 2 | 3 | 4 | 2 | 2 |
| 3 | 2 | 3 | 4 | 0 |
| 4 | 4 | 1 | 1 | 3 |
| 5 | 1 | 3 | 2 | 1 |
| 6 | 2 | 4 | 0 | 2 |
| 7 | 3 | 1 | 1 | 1 |
| 8 | 4 | 2 | 3 | 3 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli nazwe kazdej druzyny oraz laczna liczbe bramek strzelonych jako gospodarze (SUM bramki_gosp). Pokaz tylko druzyny, ktore strzelily jako gospodarze wiecej niz 3 bramki lacznie. Posortuj malejaco.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Musisz polaczyc Druzyny z Meczami, ale po odpowiedniej kolumnie (gospodarze).
2. **Podejscie**: JOIN na `id_druzyny = id_gospodarzy`, potem GROUP BY i SUM.
3. **Kluczowy krok**: Filtrowanie po SUM wymaga HAVING, nie WHERE.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT D.nazwa, SUM(M.bramki_gosp) AS bramki_u_siebie
FROM Druzyny D
JOIN Mecze M ON D.id_druzyny = M.id_gospodarzy
GROUP BY D.nazwa
HAVING SUM(M.bramki_gosp) > 3
ORDER BY bramki_u_siebie DESC;
```

**Weryfikacja**:
- Orly (id=1) jako gospodarze: mecz 1 (3), mecz 5 (2) → suma = **5** (>3 ✓)
- Wilki (id=2) jako gospodarze: mecz 3 (4), mecz 6 (0) → suma = **4** (>3 ✓)
- Rysie (id=3) jako gospodarze: mecz 2 (2), mecz 7 (1) → suma = **3** (>3 ✗)
- Sokoly (id=4) jako gospodarze: mecz 4 (1), mecz 8 (3) → suma = **4** (>3 ✓)

| nazwa | bramki_u_siebie |
|-------|----------------|
| Orly | 5 |
| Wilki | 4 |
| Sokoly | 4 |

**Wyjasnienie**: JOIN laczy druzyny z meczami po id_gospodarzy — to kluczowy element, bo kazda druzyna gra raz jako gospodarz, raz jako gosc. Uzywamy `M.id_gospodarzy` (nie `id_gosci`), bo interesuja nas tylko bramki strzelone u siebie.
</details>

<details>
<summary>Typowe bledy</summary>

- **Pomylenie id_gospodarzy z id_gosci**: Daje bramki stracone zamiast strzelonych u siebie. CKE: -3 pkt (bledny wynik).
- **Uzycie bramki_gosci zamiast bramki_gosp**: Analogiczny blad — bramki_gosci to bramki strzelone przez gosci, nie przez gospodarzy. CKE: -3 pkt.
- **Brak HAVING**: Bez filtrowania Rysie (3 bramki) tez pojawia sie w wyniku. CKE: -1 pkt.

</details>

---

### Cwiczenie 20.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)
**Tagi**: `GROUP-BY` `SUM` `HAVING` `JOIN-z-GROUP-BY` `wielokolumnowe-grupowanie` `SUBSTR` `funkcje-daty`

Tabela **Klienci**:

| id_klienta | nazwa | miasto | typ |
|------------|-------|--------|-----|
| 1 | Sklep ABC | Warszawa | detaliczny |
| 2 | Market XYZ | Krakow | hurtowy |
| 3 | Delikatesy Jan | Warszawa | detaliczny |
| 4 | Supermarket Ola | Gdansk | hurtowy |
| 5 | Sklep u Marka | Krakow | detaliczny |

Tabela **Zamowienia**:

| id_zamowienia | id_klienta | data_zamowienia | kwota |
|---------------|------------|-----------------|-------|
| 1 | 1 | 2024-01-15 | 1200 |
| 2 | 2 | 2024-01-20 | 5800 |
| 3 | 1 | 2024-02-10 | 800 |
| 4 | 3 | 2024-02-14 | 1500 |
| 5 | 4 | 2024-03-01 | 4200 |
| 6 | 2 | 2024-03-15 | 3600 |
| 7 | 5 | 2024-01-25 | 950 |
| 8 | 3 | 2024-03-20 | 2100 |
| 9 | 1 | 2024-03-28 | 1100 |
| 10 | 4 | 2024-01-30 | 3800 |
| 11 | 2 | 2024-02-22 | 4100 |
| 12 | 5 | 2024-03-05 | 700 |

**Polecenie**: Napisz zapytanie SQL, ktore dla kazdego miasta i kazdego miesiaca (numer miesiaca) wyswietli laczna kwote zamowien. Pokaz tylko te kombinacje miasto-miesiac, w ktorych laczna kwota przekroczyla 2000. Posortuj wg miasta, a nastepnie wg miesiaca.

**Wskazowka**: Uzyj funkcji SUBSTR(data_zamowienia, 6, 2) lub STRFTIME('%m', data_zamowienia) do wyodrebnienia miesiaca.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Grupujesz po dwoch kolumnach jednoczesnie: miasto i miesiac wyciagniety z daty.
2. **Podejscie**: JOIN + GROUP BY z dwoma wyrazeniami + HAVING na SUM.
3. **Kluczowy krok**: SUBSTR(data_zamowienia, 6, 2) wyciaga znaki na pozycjach 6-7 z formatu 'RRRR-MM-DD', dajac numer miesiaca.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT K.miasto, SUBSTR(Z.data_zamowienia, 6, 2) AS miesiac, SUM(Z.kwota) AS suma_kwot
FROM Klienci K
JOIN Zamowienia Z ON K.id_klienta = Z.id_klienta
GROUP BY K.miasto, SUBSTR(Z.data_zamowienia, 6, 2)
HAVING SUM(Z.kwota) > 2000
ORDER BY K.miasto, miesiac;
```

**Weryfikacja** (grupowanie miasto + miesiac):

Gdansk:
- 01: klient 4 → 3800 → **3800** (>2000 ✓)
- 03: klient 4 → 4200 → **4200** (>2000 ✓)

Krakow:
- 01: klient 2 (5800) + klient 5 (950) → **6750** (>2000 ✓)
- 02: klient 2 (4100) → **4100** (>2000 ✓)
- 03: klient 2 (3600) + klient 5 (700) → **4300** (>2000 ✓)

Warszawa:
- 01: klient 1 (1200) → **1200** (>2000 ✗)
- 02: klient 1 (800) + klient 3 (1500) → **2300** (>2000 ✓)
- 03: klient 3 (2100) + klient 1 (1100) → **3200** (>2000 ✓)

| miasto | miesiac | suma_kwot |
|--------|---------|-----------|
| Gdansk | 01 | 3800 |
| Gdansk | 03 | 4200 |
| Krakow | 01 | 6750 |
| Krakow | 02 | 4100 |
| Krakow | 03 | 4300 |
| Warszawa | 02 | 2300 |
| Warszawa | 03 | 3200 |

**Wyjasnienie**: GROUP BY z wieloma kolumnami (miasto, miesiac) tworzy grupy dla kazdej unikalnej kombinacji wartosci. SUBSTR(data, 6, 2) wyciaga znaki 6-7 z daty w formacie 'RRRR-MM-DD', dajac numer miesiaca jako tekst ('01', '02', '03'). Alternatywnie: STRFTIME('%m', data).
</details>

<details>
<summary>Typowe bledy</summary>

- **Grupowanie tylko po jednej kolumnie**: `GROUP BY K.miasto` bez miesiaca — lacza sie np. Gdansk-01 i Gdansk-03 w jedna grupe. CKE: -2 pkt.
- **SUBSTR z blednym offsetem**: `SUBSTR(data, 5, 2)` wyciaga '-M' zamiast 'MM'. W SQL pozycje liczymy od 1. CKE: -2 pkt (bledne grupowanie).
- **Brak HAVING**: Warszawa-01 (1200) pojawia sie w wyniku mimo ze < 2000. CKE: -1 pkt.

</details>

---

### Cwiczenie 20.6 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2025 (Woda na Marsie)
**Tagi**: `GROUP-BY` `MIN` `MAX` `ORDER-BY`

Tabela **Stacje** zawiera dane o stacjach meteorologicznych:

| id | nazwa | region | wysokosc_npm | rok_zalozenia |
|----|-------|--------|-------------|---------------|
| 1 | Kasprowy Wierch | gory | 1991 | 1936 |
| 2 | Zakopane | gory | 844 | 1950 |
| 3 | Warszawa | niziny | 106 | 1921 |
| 4 | Suwalki | niziny | 184 | 1945 |
| 5 | Sniezka | gory | 1603 | 1900 |
| 6 | Wroclaw | niziny | 120 | 1948 |
| 7 | Hala Gasienicowa | gory | 1520 | 1965 |
| 8 | Poznan | niziny | 86 | 1953 |

**Polecenie**: Napisz zapytanie SQL, ktore dla kazdego regionu wyswietli minimalna i maksymalna wysokosc n.p.m. stacji. Posortuj alfabetycznie wg regionu.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Grupuj po regionie i zastosuj funkcje agregujace MIN i MAX.
2. **Podejscie**: SELECT region, MIN(wysokosc_npm), MAX(wysokosc_npm) z GROUP BY.
3. **Kluczowy krok**: Pamietaj o aliasach — `AS min_wysokosc` i `AS max_wysokosc`.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT region, MIN(wysokosc_npm) AS min_wysokosc, MAX(wysokosc_npm) AS max_wysokosc
FROM Stacje
GROUP BY region
ORDER BY region;
```

**Weryfikacja**:
- gory: Kasprowy (1991), Zakopane (844), Sniezka (1603), Hala Gasienicowa (1520) → MIN=**844**, MAX=**1991**
- niziny: Warszawa (106), Suwalki (184), Wroclaw (120), Poznan (86) → MIN=**86**, MAX=**184**

| region | min_wysokosc | max_wysokosc |
|--------|-------------|-------------|
| gory | 844 | 1991 |
| niziny | 86 | 184 |

**Wyjasnienie**: MIN i MAX to funkcje agregujace — dzialaja na grupach wierszy. Mozna je laczyc w jednym zapytaniu z innymi funkcjami (COUNT, SUM, AVG). Kazda kolumna w SELECT musi albo byc w GROUP BY, albo byc funkcja agregujaca.
</details>

<details>
<summary>Typowe bledy</summary>

- **Dodanie kolumny nie-agregatowej do SELECT bez GROUP BY**: np. `SELECT region, nazwa, MIN(wysokosc_npm)` — ktora nazwa stacji powinna sie wyswietlic? SQLite wybiera losowo, inne bazy zglaszaja blad. CKE: -1 pkt.
- **Pomylenie MIN/MAX**: Odwrotne nazwy kolumn wynikowych. CKE: -1 pkt.

</details>

---

### Cwiczenie 20.7 (trudnosc: srednie, ~3 pkt)
**Zrodlo inspiracji**: Matura 2019 (Perfumy)
**Tagi**: `GROUP-BY` `COUNT` `HAVING` `WHERE-vs-HAVING`

Tabela **Zamowienia** zawiera dane o zamowieniach w restauracji:

| id | stolik | kelner | danie | cena | data |
|----|--------|--------|-------|------|------|
| 1 | 3 | Anna | Pizza | 32 | 2024-05-01 |
| 2 | 1 | Jan | Spaghetti | 28 | 2024-05-01 |
| 3 | 3 | Anna | Tiramisu | 18 | 2024-05-01 |
| 4 | 2 | Anna | Pizza | 32 | 2024-05-02 |
| 5 | 1 | Jan | Risotto | 35 | 2024-05-02 |
| 6 | 4 | Ewa | Spaghetti | 28 | 2024-05-02 |
| 7 | 3 | Anna | Risotto | 35 | 2024-05-02 |
| 8 | 2 | Jan | Tiramisu | 18 | 2024-05-03 |
| 9 | 4 | Ewa | Pizza | 32 | 2024-05-03 |
| 10 | 1 | Jan | Pizza | 32 | 2024-05-03 |

**Polecenie**: Napisz zapytanie SQL, ktore wyswietli imie kelnera oraz liczbe obslugionych zamowien, ale tylko dla kelnerow, ktorzy obsluzyli wiecej niz 2 zamowienia. Posortuj malejaco wg liczby zamowien.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Grupuj po kelnerze, policz zamowienia, odfiltruj grupy.
2. **Podejscie**: GROUP BY kelner + COUNT(*) + HAVING.
3. **Kluczowy krok**: HAVING COUNT(*) > 2 — nie mozesz uzyc WHERE do filtrowania wynikow agregacji.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT kelner, COUNT(*) AS liczba_zamowien
FROM Zamowienia
GROUP BY kelner
HAVING COUNT(*) > 2
ORDER BY liczba_zamowien DESC;
```

**Weryfikacja**:
- Anna: zamowienia 1, 3, 4, 7 → **4** (>2 ✓)
- Jan: zamowienia 2, 5, 8, 10 → **4** (>2 ✓)
- Ewa: zamowienia 6, 9 → **2** (>2 ✗)

| kelner | liczba_zamowien |
|--------|----------------|
| Anna | 4 |
| Jan | 4 |

**Wyjasnienie**: HAVING filtruje grupy po agregacji. Ewa ma dokladnie 2 zamowienia — warunek `> 2` ja wyklucza. Gdyby polecenie mowilo "co najmniej 2", nalezaloby uzyc `>= 2`.
</details>

<details>
<summary>Typowe bledy</summary>

- **`> 2` vs `>= 2`**: "Wiecej niz 2" to `> 2`, a "co najmniej 2" to `>= 2`. Czesta pomylka na maturze. CKE: -1 pkt.
- **WHERE COUNT(*) > 2**: Blad skladni — COUNT nie dziala w WHERE. CKE: -2 pkt.

</details>

---

### Cwiczenie 20.8 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2023 (Gry planszowe)
**Tagi**: `GROUP-BY` `AVG` `ROUND` `HAVING` `JOIN-z-GROUP-BY` `WHERE-vs-HAVING`

Tabela **Uczniowie**:

| id | imie | nazwisko | klasa |
|----|------|----------|-------|
| 1 | Anna | Maj | 3A |
| 2 | Jan | Krol | 3A |
| 3 | Ewa | Lis | 3B |
| 4 | Piotr | Zak | 3B |
| 5 | Kasia | Wrobel | 3A |
| 6 | Tomek | Bak | 3C |

Tabela **Oceny**:

| id | id_ucznia | przedmiot | ocena |
|----|-----------|-----------|-------|
| 1 | 1 | Matematyka | 5 |
| 2 | 1 | Fizyka | 4 |
| 3 | 2 | Matematyka | 3 |
| 4 | 2 | Fizyka | 3 |
| 5 | 3 | Matematyka | 5 |
| 6 | 3 | Fizyka | 5 |
| 7 | 4 | Matematyka | 4 |
| 8 | 4 | Fizyka | 3 |
| 9 | 5 | Matematyka | 4 |
| 10 | 5 | Fizyka | 5 |
| 11 | 6 | Matematyka | 2 |
| 12 | 6 | Fizyka | 3 |

**Polecenie**: Napisz zapytanie SQL, ktore dla kazdego przedmiotu i kazdej klasy wyswietli srednia ocene (zaokraglona do 1 miejsca po przecinku). Pokaz tylko te kombinacje, w ktorych srednia jest wieksza lub rowna 4.0. Posortuj wg przedmiotu, potem malejaco wg sredniej.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Potrzebujesz JOIN dwoch tabel, potem GROUP BY po dwoch kolumnach.
2. **Podejscie**: JOIN Uczniowie z Ocenami, GROUP BY (przedmiot, klasa), HAVING na AVG.
3. **Kluczowy krok**: Grupujesz po dwoch kolumnach — kazda unikalna para (przedmiot, klasa) to osobna grupa.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT O.przedmiot, U.klasa, ROUND(AVG(O.ocena), 1) AS srednia
FROM Uczniowie U
JOIN Oceny O ON U.id = O.id_ucznia
GROUP BY O.przedmiot, U.klasa
HAVING AVG(O.ocena) >= 4.0
ORDER BY O.przedmiot, srednia DESC;
```

**Weryfikacja** (grupowanie przedmiot + klasa):

Fizyka:
- 3A: Anna (4), Jan (3), Kasia (5) → (4+3+5)/3 = **4.0** (>=4.0 ✓)
- 3B: Ewa (5), Piotr (3) → (5+3)/2 = **4.0** (>=4.0 ✓)
- 3C: Tomek (3) → **3.0** (>=4.0 ✗)

Matematyka:
- 3A: Anna (5), Jan (3), Kasia (4) → (5+3+4)/3 = **4.0** (>=4.0 ✓)
- 3B: Ewa (5), Piotr (4) → (5+4)/2 = **4.5** (>=4.0 ✓)
- 3C: Tomek (2) → **2.0** (>=4.0 ✗)

| przedmiot | klasa | srednia |
|-----------|-------|---------|
| Fizyka | 3A | 4.0 |
| Fizyka | 3B | 4.0 |
| Matematyka | 3B | 4.5 |
| Matematyka | 3A | 4.0 |

**Wyjasnienie**: Wielokolumnowe GROUP BY tworzy grupe dla kazdej unikalnej pary (przedmiot, klasa). HAVING z `>=` obejmuje rowniez srednia rowna dokladnie 4.0. ROUND(..., 1) zaokragla do 1 miejsca po przecinku.
</details>

<details>
<summary>Typowe bledy</summary>

- **`>` zamiast `>=`**: Pominiecie grup ze srednia rowna 4.0. Przy tych danych eliminuje 3 z 4 wynikow. CKE: -2 pkt.
- **GROUP BY po jednej kolumnie**: Grupowanie tylko po przedmiocie mieszalby klasy razem. CKE: -2 pkt.
- **Brak JOIN**: Tabela Oceny nie ma kolumny klasa — bez JOIN nie mozna pogrupowac wg klas. CKE: -3 pkt.

</details>

---

### Cwiczenie 20.9 (trudnosc: srednie-trudne, ~4 pkt)
**Zrodlo inspiracji**: Matura 2021 (Gra strategiczna)
**Tagi**: `GROUP-BY` `COUNT` `SUM` `HAVING` `JOIN-z-GROUP-BY` `aliasy`

Tabela **Sklepy**:

| id_sklepu | nazwa | miasto |
|-----------|-------|--------|
| 1 | Elektro-Max | Warszawa |
| 2 | TechWorld | Krakow |
| 3 | Neonet | Warszawa |
| 4 | MediaPark | Gdansk |

Tabela **Sprzedaz**:

| id | id_sklepu | produkt | ilosc | cena_jednostkowa | data |
|----|-----------|---------|-------|-------------------|------|
| 1 | 1 | Laptop | 2 | 3500 | 2024-06-01 |
| 2 | 1 | Mysz | 10 | 80 | 2024-06-01 |
| 3 | 2 | Laptop | 1 | 3200 | 2024-06-02 |
| 4 | 2 | Monitor | 3 | 1200 | 2024-06-02 |
| 5 | 3 | Laptop | 3 | 3800 | 2024-06-03 |
| 6 | 3 | Mysz | 5 | 90 | 2024-06-03 |
| 7 | 4 | Monitor | 2 | 1100 | 2024-06-04 |
| 8 | 1 | Monitor | 1 | 1300 | 2024-06-05 |
| 9 | 4 | Laptop | 1 | 3000 | 2024-06-05 |

**Polecenie**: Napisz zapytanie SQL, ktore dla kazdego sklepu wyswietli nazwe sklepu, liczbe roznych transakcji oraz laczny przychod (SUM ilosc * cena_jednostkowa). Pokaz tylko sklepy z przychodem powyzej 5000. Posortuj malejaco wg przychodu.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: JOIN + GROUP BY, ale uwaga na obliczenie przychodu — to iloczyn dwoch kolumn.
2. **Podejscie**: `SUM(ilosc * cena_jednostkowa)` oblicza laczny przychod per grupa.
3. **Kluczowy krok**: Mnozenie odbywa sie WEWNATRZ SUM — najpierw ilosc*cena dla kazdego wiersza, potem suma w grupie.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT S.nazwa, COUNT(*) AS liczba_transakcji,
       SUM(SP.ilosc * SP.cena_jednostkowa) AS przychod
FROM Sklepy S
JOIN Sprzedaz SP ON S.id_sklepu = SP.id_sklepu
GROUP BY S.nazwa
HAVING SUM(SP.ilosc * SP.cena_jednostkowa) > 5000
ORDER BY przychod DESC;
```

**Weryfikacja**:
- Elektro-Max: (2*3500)+(10*80)+(1*1300) = 7000+800+1300 = **9100**, 3 transakcje (>5000 ✓)
- TechWorld: (1*3200)+(3*1200) = 3200+3600 = **6800**, 2 transakcje (>5000 ✓)
- Neonet: (3*3800)+(5*90) = 11400+450 = **11850**, 2 transakcje (>5000 ✓)
- MediaPark: (2*1100)+(1*3000) = 2200+3000 = **5200**, 2 transakcje (>5000 ✓)

| nazwa | liczba_transakcji | przychod |
|-------|-------------------|----------|
| Neonet | 2 | 11850 |
| Elektro-Max | 3 | 9100 |
| TechWorld | 2 | 6800 |
| MediaPark | 2 | 5200 |

**Wyjasnienie**: `SUM(ilosc * cena_jednostkowa)` oblicza przychod: najpierw mnozy ilosc przez cene dla kazdego wiersza, potem sumuje w grupie. To czesty wzorzec na maturze. Nie mozna uzyc `SUM(ilosc) * SUM(cena_jednostkowa)` — to daloby zupelnie inny (bledny) wynik.
</details>

<details>
<summary>Typowe bledy</summary>

- **`SUM(ilosc) * SUM(cena_jednostkowa)`**: To nie to samo co `SUM(ilosc * cena_jednostkowa)`. Np. dla Elektro-Max: SUM(ilosc)=13, SUM(cena)=4880, iloczyn=63440 — zupelnie bledny. CKE: -3 pkt.
- **Brak HAVING**: Wszystkie sklepy maja przychod > 5000, ale gdyby nie — bledne wyniki. CKE: -1 pkt.
- **Uzycie aliasu w HAVING**: `HAVING przychod > 5000` — w SQLite dziala, ale bezpieczniej powtorzyc wyrazenie. CKE: akceptowane, ale ryzykowne na egzaminie.

</details>

---

### Cwiczenie 20.10 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2015 (Formula 1)
**Tagi**: `GROUP-BY` `COUNT` `AVG` `ROUND` `HAVING` `JOIN-z-GROUP-BY` `wielokolumnowe-grupowanie` `aliasy`

Tabela **Wydzialy**:

| id_wydzialu | nazwa_wydzialu | budynek |
|-------------|----------------|---------|
| 1 | Informatyka | A |
| 2 | Matematyka | A |
| 3 | Fizyka | B |
| 4 | Chemia | B |

Tabela **Studenci**:

| id | imie | nazwisko | id_wydzialu | rok_studiow |
|----|------|----------|-------------|-------------|
| 1 | Anna | Kowalska | 1 | 2 |
| 2 | Jan | Nowak | 1 | 3 |
| 3 | Ewa | Maj | 2 | 1 |
| 4 | Piotr | Lis | 3 | 2 |
| 5 | Kasia | Zak | 1 | 1 |
| 6 | Tomek | Bak | 2 | 2 |
| 7 | Ola | Wrobel | 3 | 1 |
| 8 | Marek | Krol | 4 | 3 |

Tabela **Egzaminy**:

| id | id_studenta | przedmiot | wynik | data_egzaminu |
|----|-------------|-----------|-------|---------------|
| 1 | 1 | Algorytmy | 85 | 2024-01-20 |
| 2 | 1 | Bazy danych | 92 | 2024-01-25 |
| 3 | 2 | Algorytmy | 78 | 2024-01-20 |
| 4 | 3 | Analiza | 95 | 2024-01-22 |
| 5 | 4 | Mechanika | 60 | 2024-01-23 |
| 6 | 5 | Algorytmy | 88 | 2024-01-20 |
| 7 | 6 | Analiza | 72 | 2024-01-22 |
| 8 | 7 | Mechanika | 55 | 2024-01-23 |
| 9 | 8 | Chemia org. | 90 | 2024-01-24 |
| 10 | 2 | Bazy danych | 70 | 2024-01-25 |
| 11 | 4 | Termodynamika | 65 | 2024-01-26 |
| 12 | 6 | Algebra | 80 | 2024-01-27 |

**Polecenie**: Napisz zapytanie SQL, ktore dla kazdego budynku i kazdego roku studiow wyswietli:
- nazwe budynku
- rok studiow
- liczbe studentow (ktory zdawali przynajmniej 1 egzamin)
- sredni wynik egzaminow (zaokraglony do 1 miejsca)

Pokaz tylko grupy, w ktorych sredni wynik jest wyzszy niz 75. Posortuj wg budynku, potem malejaco wg sredniej.

<details>
<summary>Wskazowki</summary>

1. **Kierunek**: Trzy tabele do polaczenia, grupowanie po dwoch kolumnach z roznych tabel.
2. **Podejscie**: JOIN Wydzialy → Studenci → Egzaminy, GROUP BY (budynek, rok_studiow).
3. **Kluczowy krok**: COUNT(DISTINCT id_studenta) liczy unikalnych studentow — jeden student moze miec wiele egzaminow. AVG(wynik) oblicza srednia ze wszystkich egzaminow w grupie.

</details>

<details>
<summary>Odpowiedz</summary>

**Zapytanie SQL:**
```sql
SELECT W.budynek, S.rok_studiow,
       COUNT(DISTINCT S.id) AS liczba_studentow,
       ROUND(AVG(E.wynik), 1) AS sredni_wynik
FROM Wydzialy W
JOIN Studenci S ON W.id_wydzialu = S.id_wydzialu
JOIN Egzaminy E ON S.id = E.id_studenta
GROUP BY W.budynek, S.rok_studiow
HAVING AVG(E.wynik) > 75
ORDER BY W.budynek, sredni_wynik DESC;
```

**Weryfikacja** (grupowanie budynek + rok_studiow):

Budynek A:
- rok 1: Ewa (Analiza 95), Kasia (Algorytmy 88) → 2 studentow, srednia = (95+88)/2 = **91.5** (>75 ✓)
- rok 2: Anna (Algorytmy 85, Bazy 92), Tomek (Analiza 72, Algebra 80) → 2 studentow, srednia = (85+92+72+80)/4 = **82.3** (>75 ✓) (zaokr. 329/4=82.25→82.3)
- rok 3: Jan (Algorytmy 78, Bazy 70) → 1 student, srednia = (78+70)/2 = **74.0** (>75 ✗)

Budynek B:
- rok 1: Ola (Mechanika 55) → 1 student, srednia = **55.0** (>75 ✗)
- rok 2: Piotr (Mechanika 60, Termodynamika 65) → 1 student, srednia = (60+65)/2 = **62.5** (>75 ✗)
- rok 3: Marek (Chemia org. 90) → 1 student, srednia = **90.0** (>75 ✓)

| budynek | rok_studiow | liczba_studentow | sredni_wynik |
|---------|-------------|------------------|-------------|
| A | 1 | 2 | 91.5 |
| A | 2 | 2 | 82.3 |
| B | 3 | 1 | 90.0 |

**Wyjasnienie**: Trzy JOIN-y tworzace lancuch Wydzialy→Studenci→Egzaminy. COUNT(DISTINCT S.id) liczy unikalnych studentow (Anna ma 2 egzaminy, ale to 1 studentka). AVG(E.wynik) oblicza srednia z WSZYSTKICH egzaminow w grupie (nie srednia ze srednich studentow). Wielokolumnowe GROUP BY tworzy grupe per (budynek, rok_studiow).
</details>

<details>
<summary>Typowe bledy</summary>

- **COUNT(*) zamiast COUNT(DISTINCT)**: Anna ma 2 egzaminy — COUNT(*) policzyloby ja 2 razy. CKE: -1 pkt.
- **Brak jednego JOIN**: Pominiecie tabeli Wydzialy — nie mozna uzyskac budynku. CKE: -3 pkt.
- **Srednia ze srednich**: Osobne obliczenie sredniej per student, potem srednia z tych srednich — to inny wynik niz AVG ze wszystkich egzaminow. CKE: -2 pkt (zalezy od interpretacji polecenia).

</details>

---

## Samoocena

| Poziom | Opis | Wymaganie |
|--------|------|-----------|
| Podstawowy | Rozumiesz GROUP BY z jedna funkcja agregujaca | 1-3 cwiczen bez pomocy |
| Dobry | Laczysz GROUP BY z HAVING i JOIN | 4-6 cwiczen bez pomocy |
| Bardzo dobry | Wielokolumnowe GROUP BY, DISTINCT, obliczenia w SUM | 7-8 cwiczen bez pomocy |
| Doskonaly | Wszystkie cwiczenia, w tym lancuchy JOIN i zlozone HAVING | 9-10 cwiczen bez pomocy |

### Co dalej?
- Jesli poziom **Podstawowy**: Powtorz cwiczenia 20.1-20.2, potem przejdz do `cheatsheet_sql.md` (sekcja GROUP BY).
- Jesli poziom **Dobry**: Przejdz do cwiczen `21_sql_podzapytania.md` i `22_sql_join.md`.
- Jesli poziom **Bardzo dobry**: Sprobuj cwiczen `21_sql_podzapytania.md` (cwiczenia trudne) i wrocz do 20.8-20.10 bez wskazowek.
- Jesli poziom **Doskonaly**: Przejdz do arkuszy maturalnych — zacznij od 2023+.
