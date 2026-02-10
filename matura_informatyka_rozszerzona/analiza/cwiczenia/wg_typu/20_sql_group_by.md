# 20. SQL — GROUP BY z agregacjami

Typ zadania: **sql_group_by**
Czestotliwosc: 8/11 lat | Laczna punktacja: 36 pkt
Kategoria: SQL

---

### Cwiczenie 20.1 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2023 (Gry planszowe)

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

**Wyjasnienie**: GROUP BY grupuje wiersze o tej samej wartosci w kolumnie `gatunek`. COUNT(*) zlicza wiersze w kazdej grupie. ORDER BY z DESC sortuje od najwiekszej wartosci. Typowy blad maturalny: uzycie WHERE zamiast GROUP BY do filtrowania grup — WHERE filtruje wiersze PRZED grupowaniem, a do filtrowania grup sluzy HAVING.
</details>

---

### Cwiczenie 20.2 (trudnosc: latwe, ~2 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)

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

---

### Cwiczenie 20.3 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2022 (Ewidencja uczniow)

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

**Wyjasnienie**: Kluczowa roznica: WHERE filtruje wiersze PRZED grupowaniem, HAVING filtruje grupy PO grupowaniu. Tutaj nie mozna uzyc `WHERE AVG(...) > 4.0`, bo AVG wymaga juz zgrupowanych danych. Typowy blad maturalny: uzycie WHERE zamiast HAVING przy warunkach na funkcje agregujace.
</details>

---

### Cwiczenie 20.4 (trudnosc: srednie, ~4 pkt)
**Zrodlo inspiracji**: Matura 2017 (Pilka reczna)

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
- Wilki (id=2) jako gospodarze: mecz 1 nie (goscie), mecz 3 (4), mecz 6 (0) → suma = **4** (>3 ✓)
- Rysie (id=3) jako gospodarze: mecz 2 (2), mecz 7 (1) → suma = **3** (>3 ✗)
- Sokoly (id=4) jako gospodarze: mecz 4 (1), mecz 8 (3) → suma = **4** (>3 ✓)

| nazwa | bramki_u_siebie |
|-------|----------------|
| Orly | 5 |
| Wilki | 4 |
| Sokoly | 4 |

**Wyjasnienie**: JOIN laczy druzyny z meczami po id_gospodarzy — to kluczowy element, bo kazda druzyna gra raz jako gospodarz, raz jako gosc. Uzywamy `M.id_gospodarzy` (nie `id_gosci`), bo interesuja nas tylko bramki strzelone u siebie. HAVING SUM(...) > 3 filtruje grupy po agregacji. Typowy blad: pomylenie id_gospodarzy z id_gosci, co daje bramki stracone zamiast strzelonych.
</details>

---

### Cwiczenie 20.5 (trudnosc: trudne, ~5 pkt)
**Zrodlo inspiracji**: Matura 2024 (Hurtownia)

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

**Wyjasnienie**: GROUP BY z wieloma kolumnami (miasto, miesiac) tworzy grupy dla kazdej unikalnej kombinacji wartosci. SUBSTR(data, 6, 2) wyciaga znaki 6-7 z daty w formacie 'RRRR-MM-DD', dajac numer miesiaca jako tekst ('01', '02', '03'). Alternatywnie: STRFTIME('%m', data). Typowy blad maturalny: grupowanie tylko po jednej kolumnie zamiast po obu — wtedy np. Gdansk-01 i Gdansk-03 zostana zlaczone w jedna grupe "Gdansk".
</details>
