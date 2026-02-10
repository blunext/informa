# Rozwiazania wzorcowe: SQL

Dwa pelne rozwiazania prawdziwych zadan maturalnych — z procesem myslowym, zapytaniami i weryfikacja.

---

## [2023] Zadanie 7: Gry planszowe (10 pkt)

**Typ**: sql_join + sql_group_by + sql_podzapytania | **Czas**: ~40 min | **Trudnosc**: trudne

### Tresc (skrot)

Trzy tabele (pliki TSV):
- **gry** (id_gry, nazwa, kategoria)
- **gracze** (id_gracza, imie, nazwisko, wiek)
- **oceny** (id_gry, id_gracza, stan, ocena) — `stan` moze byc np. "posiada", "chce", itp.

Podzadania:
- **7.1** (1 pkt): Tytul gry z najwieksza liczba ocen.
- **7.2** (2 pkt): Srednia ocen kazdej gry z kategorii "imprezowa" (do 2 miejsc po przecinku).
- **7.3** (2 pkt): Ilu graczy nie posiada zadnej gry (stan != "posiada" we wszystkich rekordach), ale wystawilo co najmniej 1 ocene?
- **7.4** (3 pkt): Kategorie wiekowe (juniorzy <=19, seniorzy 20-49, weterani >=50). Dla kazdej: gra z najwieksza liczba ocen od graczy z tej kategorii.
- **7.5** (2 pkt): SQL — ile kosztuje zakup gier logicznych w promocji? (nowa tabela `sklep`)

### Podejscie — jak myslec

1. **7.1**: Proste — GROUP BY id_gry + COUNT, ORDER BY DESC, LIMIT 1. Ale pytaja o **tytul**, wiec JOIN z gry.
2. **7.2**: JOIN gry-oceny, WHERE kategoria='imprezowa', GROUP BY gra, AVG + ROUND.
3. **7.3**: Klucz — NOT IN z podzapytaniem. Gracze ktorzy NIE maja rekordu ze stan='posiada', ALE maja jakis rekord w tabeli oceny.
4. **7.4**: CASE WHEN dla kategorii wiekowych. Najpierw policzyc oceny na (kategoria, gra), potem znalezc MAX w kazdej kategorii.
5. **7.5**: Prosty JOIN + SUM z dwoma warunkami WHERE.

### Rozwiazanie

#### 7.1 — Gra z najwieksza liczba ocen (1 pkt)

```sql
SELECT g.nazwa, COUNT(*) AS ile
FROM oceny o
JOIN gry g ON o.id_gry = g.id_gry
GROUP BY o.id_gry, g.nazwa
ORDER BY ile DESC
LIMIT 1;
```

**Wynik**: K2

**Czemu JOIN?** Pytaja o **tytul** gry, nie id_gry. Bez JOIN dostaniesz id — to moze kosztowac punkt.

#### 7.2 — Srednie ocen gier imprezowych (2 pkt)

```sql
SELECT g.nazwa, ROUND(AVG(o.ocena), 2) AS srednia
FROM oceny o
JOIN gry g ON o.id_gry = g.id_gry
WHERE g.kategoria = 'imprezowa'
GROUP BY g.nazwa
ORDER BY g.nazwa;
```

**Wynik**:

| Gra | Srednia |
|---|---|
| 5 sekund | 8.16 |
| Avalone | 8.25 |
| Colt Express | 7.54 |
| Jenga | 8.16 |
| Koncept | 8.37 |
| Mamy szpiega | 8.22 |
| Przebiegle wielblady | 7.73 |
| Sushi Go | 8.07 |
| Swiatowy Konflikt | 7.80 |
| Szeryf z Nottingham | 7.88 |

**Wazne**: ROUND (zaokraglanie), nie TRUNCATE (obcinanie). Roznica kosztuje 1 pkt.

#### 7.3 — Gracze bez posiadanych gier (2 pkt)

```sql
SELECT COUNT(DISTINCT o.id_gracza)
FROM oceny o
WHERE o.id_gracza NOT IN (
    SELECT id_gracza
    FROM oceny
    WHERE stan = 'posiada'
);
```

**Wynik**: 334

**Dlaczego nie 351?** Bo 351 to gracze bez stanu "posiada" **w calej tabeli gracze** (lacznie z tymi, co nie wystawili zadnej oceny). Warunek "wystawili co najmniej jedna ocene" oznacza: szukamy w tabeli oceny. Czytanie FROM oceny automatycznie odfiltruje graczy bez ocen.

#### 7.4 — Najpopularniejsze gry wg kategorii wiekowej (3 pkt)

To najtrudniejsze podzadanie. Podejscie dwuetapowe:

**Krok 1**: Policzyc oceny na (kategoria_wiekowa, gra):

```sql
-- Widok pomocniczy (lub podzapytanie)
SELECT
    CASE
        WHEN gr.wiek <= 19 THEN 'juniorzy'
        WHEN gr.wiek <= 49 THEN 'seniorzy'
        ELSE 'weterani'
    END AS kategoria_wiekowa,
    g.nazwa,
    COUNT(*) AS ile_ocen
FROM oceny o
JOIN gracze gr ON o.id_gracza = gr.id_gracza
JOIN gry g ON o.id_gry = g.id_gry
GROUP BY kategoria_wiekowa, g.nazwa;
```

**Krok 2**: Dla kazdej kategorii znalezc maksimum. W praktyce maturalnej — najlatwiej zrobic to w arkuszu/programie po imporcie wynikow SQL, lub uzyc podzapytania:

```sql
-- Pelne rozwiazanie SQL (zaawansowane)
SELECT t.kategoria_wiekowa, t.nazwa, t.ile_ocen
FROM (
    SELECT
        CASE
            WHEN gr.wiek <= 19 THEN 'juniorzy'
            WHEN gr.wiek <= 49 THEN 'seniorzy'
            ELSE 'weterani'
        END AS kategoria_wiekowa,
        g.nazwa,
        COUNT(*) AS ile_ocen
    FROM oceny o
    JOIN gracze gr ON o.id_gracza = gr.id_gracza
    JOIN gry g ON o.id_gry = g.id_gry
    GROUP BY kategoria_wiekowa, g.nazwa
) t
WHERE t.ile_ocen = (
    SELECT MAX(t2.ile_ocen)
    FROM (
        SELECT
            CASE
                WHEN gr2.wiek <= 19 THEN 'juniorzy'
                WHEN gr2.wiek <= 49 THEN 'seniorzy'
                ELSE 'weterani'
            END AS kat2,
            COUNT(*) AS ile_ocen
        FROM oceny o2
        JOIN gracze gr2 ON o2.id_gracza = gr2.id_gracza
        GROUP BY kat2, o2.id_gry
    ) t2
    WHERE t2.kat2 = t.kategoria_wiekowa
);
```

**Wyniki**:

| Kategoria | Gra (gry) | Ocen |
|---|---|---|
| juniorzy | Terraformacja Marsa **i** K2 | 6 |
| seniorzy | K2 | 24 |
| weterani | Robinson Crusoe | 28 |

**Uwaga**: Juniorzy maja **dwie** gry z tym samym maksimum — trzeba podac obie!

#### 7.5 — Koszt gier logicznych w promocji (2 pkt)

Nowa tabela: **sklep** (id_gry, cena, promocja). Relacja gry-sklep: jeden do wielu (gra moze miec kilka ofert).

```sql
SELECT SUM(s.cena) AS koszt
FROM gry g
JOIN sklep s ON g.id_gry = s.id_gry
WHERE g.kategoria = 'logiczna'
  AND s.promocja = true;
```

**Kluczowe**: INNER JOIN (nie LEFT JOIN) — interesuja nas tylko gry **dostepne** w sklepie. Filtr `promocja = true` wyklucza oferty w normalnych cenach.

### Weryfikacja

| Podzadanie | Oficjalna odpowiedz CKE |
|---|---|
| 7.1 | **K2** |
| 7.2 | 10 gier imprezowych ze srednimi (tabela powyzej) |
| 7.3 | **334** (pulapka: 351 = bez filtra "wystawili ocene") |
| 7.4 | juniorzy: **Terraformacja Marsa i K2** (po 6), seniorzy: **K2** (24), weterani: **Robinson Crusoe** (28) |
| 7.5 | Zapytanie SQL (oceniane za poprawnosc skladni) |

### Pulapki

- **7.2**: ROUND, nie TRUNCATE — roznica moze byc 1 grosz, ale kosztuje punkt.
- **7.3**: 351 vs 334 — trzeba filtrowac po graczach ktorzy **maja jakis rekord** w tabeli oceny.
- **7.4**: Juniorzy maja **remis** (dwie gry po 6 ocen) — trzeba podac obie.
- **7.4**: Podanie id_gry zamiast nazwy kosztuje 1 pkt.
- **7.5**: Gra moze miec wiele ofert w sklepie (cena normalna i promocyjna) — filtr `promocja = true`.

---

## [2025] Zadanie 7: Poszukiwanie wody na Marsie (10 pkt)

**Typ**: sql_join + sql_podzapytania + przetwarzanie_zlozone | **Czas**: ~40 min | **Trudnosc**: trudne

### Tresc (skrot)

Cztery tabele:
- **Producent** (kod_producenta, nazwa, kraj)
- **Laziki** (nr_lazika, nazwa_lazika, rok_wyslania, wsp_ladowania, kod_producenta)
- **Pomiary** (nr_lazika, data_pomiaru, kod_obszaru, wspolrzedne, glebokosc, ilosc)
- **Obszary** (kod_obszaru, nazwa_obszaru)

`wsp_ladowania` i `wspolrzedne` zawieraja litere N (polnocna) lub S (poludniowa polkula).

Podzadania:
- **7.1** (2 pkt): Obszar z najwieksza laczna iloscia wody na glebokosci **<= 100 m**.
- **7.2** (2 pkt): Lazik z najdluzszym okresem pomiarow (od pierwszego do ostatniego).
- **7.3** (2 pkt): Obszary, na ktorych **zaden** lazik nie wykonal pomiaru w roku wyslania.
- **7.4** (2 pkt): Laziki wyladowane na polkuli S, ale wykonujace pomiary na **obu** polkulach.
- **7.5** (2 pkt): SQL — producenci, ktorych laziki badaly "Arcadia" w 2060 r.

### Podejscie — jak myslec

1. **7.1**: JOIN pomiary-obszary, WHERE glebokosc <= 100, GROUP BY obszar, SUM(ilosc), szukaj MAX.
2. **7.2**: GROUP BY nr_lazika, MAX(data) - MIN(data), szukaj MAX roznice. JOIN z laziki po nazwe.
3. **7.3**: Znajdz pary (lazik, obszar) gdzie lazik mierzyl w roku wyslania. Obszary NIE nalezace do tego zbioru = odpowiedz. To NOT IN / NOT EXISTS.
4. **7.4**: Filtruj laziki z S w wsp_ladowania. Sprawdz czy maja pomiary z N i z S we wspolrzednych.
5. **7.5**: JOIN 4 tabel, WHERE nazwa_obszaru = 'Arcadia' AND YEAR(data) = 2060, SELECT DISTINCT nazwa producenta.

### Rozwiazanie

#### 7.1 — Obszar z najwieksza iloscia wody do 100 m (2 pkt)

```sql
SELECT o.nazwa_obszaru, SUM(p.ilosc) AS laczna_woda
FROM Pomiary p
JOIN Obszary o ON p.kod_obszaru = o.kod_obszaru
WHERE p.glebokosc <= 100    -- do 100 metrow WLACZNIE
GROUP BY o.nazwa_obszaru
ORDER BY laczna_woda DESC
LIMIT 1;
```

**Wynik**: Mare Boreum

**Wazne**: `<= 100`, nie `< 100` — "do 100 metrow wlacznie".

#### 7.2 — Lazik z najdluzszym okresem pomiarow (2 pkt)

```sql
SELECT l.nazwa_lazika,
       MIN(p.data_pomiaru) AS pierwszy,
       MAX(p.data_pomiaru) AS ostatni
FROM Pomiary p
JOIN Laziki l ON p.nr_lazika = l.nr_lazika
GROUP BY l.nr_lazika, l.nazwa_lazika
ORDER BY (MAX(p.data_pomiaru) - MIN(p.data_pomiaru)) DESC
LIMIT 1;
```

Alternatywnie z DATEDIFF:
```sql
ORDER BY DATEDIFF(MAX(p.data_pomiaru), MIN(p.data_pomiaru)) DESC
```

**Wynik**: Spirit 14, pierwszy: 29.08.2066, ostatni: 25.07.2076

**Uwaga**: "Najdluzszy okres" = roznica miedzy datami, NIE liczba pomiarow.

#### 7.3 — Obszary bez pomiarow w roku wyslania lazika (2 pkt)

To zadanie z **negacja** — szukamy obszarow, na ktorych ZADEN lazik nie mierzyl w swoim roku wyslania.

```sql
-- Krok 1: Znajdz obszary, na ktorych JAKIS lazik mierzyl w roku wyslania
-- Krok 2: Wybierz obszary NIE nalezace do tego zbioru

SELECT o.nazwa_obszaru
FROM Obszary o
WHERE o.kod_obszaru NOT IN (
    SELECT DISTINCT p.kod_obszaru
    FROM Pomiary p
    JOIN Laziki l ON p.nr_lazika = l.nr_lazika
    WHERE YEAR(p.data_pomiaru) = l.rok_wyslania
)
ORDER BY o.nazwa_obszaru;
```

**Wynik**: Aeolis, Amazonis, Arabia, Elysium, Eridania, Mare Tyrrhenum, Sinus Sabaeus, Syrtis Major

**Kluczowe**: Pytanie brzmi "na ktorych zaden lazik NIE wykonal pomiaru w tym samym roku, w ktorym zostal wyslany" — to negacja (NOT IN).

#### 7.4 — Laziki z S, pomiary na N i S (2 pkt)

```sql
-- Laziki wyladowane na polkuli poludniowej (S w wsp_ladowania)
-- ktore wykonywaly pomiary na obu polkulach

SELECT DISTINCT l.nazwa_lazika
FROM Laziki l
WHERE l.wsp_ladowania LIKE '%S%'
  AND l.nr_lazika IN (
      SELECT nr_lazika FROM Pomiary WHERE wspolrzedne LIKE '%N%'
  )
  AND l.nr_lazika IN (
      SELECT nr_lazika FROM Pomiary WHERE wspolrzedne LIKE '%S%'
  )
ORDER BY l.nazwa_lazika;
```

Alternatywnie z INTERSECT lub dwoma EXISTS:

```sql
SELECT l.nazwa_lazika
FROM Laziki l
WHERE l.wsp_ladowania LIKE '%S%'
  AND EXISTS (SELECT 1 FROM Pomiary p WHERE p.nr_lazika = l.nr_lazika AND p.wspolrzedne LIKE '%N%')
  AND EXISTS (SELECT 1 FROM Pomiary p WHERE p.nr_lazika = l.nr_lazika AND p.wspolrzedne LIKE '%S%')
ORDER BY l.nazwa_lazika;
```

**Wynik**: Mariner 14, Mariner 15, Mariner 20, Viking 17, Spirit 7, Spirit 12, Rosetta 1, Rosetta 8, Phoenix 3, Phoenix 13

**Kluczowe**: Rozroznienie miedzy `wsp_ladowania` (skad wyladowal) a `wspolrzedne` w Pomiary (gdzie mierzyl).

#### 7.5 — Producenci badajacy Arcadia w 2060 (2 pkt)

```sql
SELECT DISTINCT pr.nazwa
FROM Producent pr
JOIN Laziki l ON pr.kod_producenta = l.kod_producenta
JOIN Pomiary p ON l.nr_lazika = p.nr_lazika
JOIN Obszary o ON p.kod_obszaru = o.kod_obszaru
WHERE o.nazwa_obszaru = 'Arcadia'
  AND YEAR(p.data_pomiaru) = 2060;
```

**Elementy oceny**: JOIN 4 tabel + WHERE z dwoma warunkami + DISTINCT.

### Weryfikacja

| Podzadanie | Oficjalna odpowiedz CKE |
|---|---|
| 7.1 | **Mare Boreum** |
| 7.2 | **Spirit 14** (29.08.2066 – 25.07.2076) |
| 7.3 | 8 obszarow (Aeolis, Amazonis, Arabia, ...) |
| 7.4 | 10 lazikow (Mariner 14, 15, 20, ...) |
| 7.5 | Zapytanie SQL (oceniane za poprawnosc) |

### Pulapki

- **7.1**: Glebkosc `<= 100` (wlacznie), nie `< 100`.
- **7.2**: Najdluzszy **okres** (roznica dat), nie najwiecej pomiarow.
- **7.3**: Negacja — pytaja o obszary BEZ pomiarow w roku wyslania. Latwo dac odwrotna odpowiedz.
- **7.4**: Dwa rozne pola — `wsp_ladowania` (Laziki) vs `wspolrzedne` (Pomiary). Nie pomyl.
- **7.5**: DISTINCT jest konieczne — ten sam producent moze miec wiele lazikow badajacych Arcadia.
- **7.5**: JOIN 4 tabel — pamietaj o Producent-Laziki-Pomiary-Obszary (kazdy ON po kluczu obcym).
