# Cheatsheet: SQL

## Kolejnosc klauzul (KRYTYCZNE)

```
Pisanie:    SELECT -> FROM -> JOIN ON -> WHERE -> GROUP BY -> HAVING -> ORDER BY -> LIMIT
Wykonanie:  FROM -> JOIN -> WHERE -> GROUP BY -> HAVING -> SELECT -> ORDER BY -> LIMIT
```

> UWAGA: Alias z SELECT nie dziala w WHERE (bo WHERE wykona sie wczesniej). Dziala w ORDER BY.

---

## SELECT + WHERE — wzorce filtrow

```sql
WHERE kolumna = 'wartosc'                  -- dokladne
WHERE kolumna IN ('a', 'b', 'c')           -- lista
WHERE kolumna BETWEEN 10 AND 20            -- zakres (wlaczajacy!)
WHERE kolumna LIKE 'A%'                    -- zaczyna sie od A
WHERE kolumna LIKE '%ski'                  -- konczy sie na ski
WHERE kolumna LIKE '_a%'                   -- 2. litera = a
WHERE kolumna IS NULL / IS NOT NULL        -- NULL (NIE uzyj = NULL!)
WHERE (A OR B) AND C                       -- nawiasy! AND ma wyzszy priorytet
```

---

## JOIN (3 wzorce)

```sql
-- INNER JOIN: tylko pasujace rekordy
SELECT t1.kol, t2.kol
FROM tabela1 t1
JOIN tabela2 t2 ON t1.id = t2.id_foreign;

-- LEFT JOIN: wszystko z lewej (NULL gdy brak dopasowania)
SELECT t1.kol, t2.kol
FROM tabela1 t1
LEFT JOIN tabela2 t2 ON t1.id = t2.id_foreign;

-- LEFT JOIN + IS NULL: rekordy BEZ dopasowania
SELECT t1.kol FROM tabela1 t1
LEFT JOIN tabela2 t2 ON t1.id = t2.id_foreign
WHERE t2.id IS NULL;
```

JOIN 3-4 tabel: dodawaj kolejne `JOIN ... ON ...` po kolei.

---

## GROUP BY + agregacje

```sql
SELECT kolumna, COUNT(*), SUM(kol), AVG(kol), MIN(kol), MAX(kol)
FROM tabela
GROUP BY kolumna
HAVING COUNT(*) > 5           -- filtr PO grupowaniu
ORDER BY COUNT(*) DESC;
```

| Funkcja | Dzialanie |
|---------|-----------|
| `COUNT(*)` | Liczy WSZYSTKIE wiersze (z NULL) |
| `COUNT(kol)` | Liczy tylko NIE-NULL |
| `COUNT(DISTINCT kol)` | Liczy unikalne wartosci |
| `SUM`/`AVG`/`MIN`/`MAX` | Pomijaja NULL |

> UWAGA: WHERE = filtr PRZED grupowaniem, HAVING = filtr PO grupowaniu

> UWAGA: W GROUP BY musza byc WSZYSTKIE kolumny z SELECT ktore NIE sa w agregacji

---

## Podzapytania (3 wzorce)

```sql
-- 1. Filtrowanie po zbiorze
WHERE kol IN (SELECT kol FROM tabela2 WHERE ...)

-- 2. Porownanie ze srednia/max/min
WHERE kol > (SELECT AVG(kol) FROM tabela)

-- 3. Podzapytanie w HAVING
HAVING AVG(ocena) > (SELECT AVG(ocena) FROM Oceny)
```

> UWAGA: `WHERE kol > AVG(kol)` NIE ZADZIALA — trzeba podzapytania!

---

## Daty

```sql
-- Format ISO: 'RRRR-MM-DD' — porownuj jako tekst
SUBSTR(data, 1, 4)    -- rok
SUBSTR(data, 6, 2)    -- miesiac
WHERE data BETWEEN '2024-01-01' AND '2024-12-31'
```

---

## Inne przydatne

```sql
DISTINCT                -- unikalne wyniki
ORDER BY kol ASC/DESC   -- sortowanie
LIMIT n                 -- pierwsze n wynikow
ROUND(wartosc, 2)       -- zaokraglenie
LENGTH(tekst)           -- dlugosc
SUBSTR(tekst, 1, 3)     -- wycinanie (od 1, nie od 0!)
COUNT(*) * 1.0 / n      -- wymuszenie dzielenia zmiennoprzecinkowego

CASE WHEN w > 8000 THEN 'wysoka'
     WHEN w > 5000 THEN 'srednia'
     ELSE 'niska' END AS kategoria
```

---

## 10 pulapek

| # | Pulapka | Poprawnie |
|---|---------|-----------|
| 1 | `COUNT(*)` z LEFT JOIN daje 1 zamiast 0 | `COUNT(t2.kol)` — pomija NULL |
| 2 | `WHERE kol != NULL` | `WHERE kol IS NOT NULL` |
| 3 | `NOT IN` z NULL w podzapytaniu = pusty wynik | `LEFT JOIN + IS NULL` |
| 4 | `SELECT miasto, COUNT(*)` bez GROUP BY | Dodaj `GROUP BY miasto` |
| 5 | `WHERE COUNT(*) > 5` | `HAVING COUNT(*) > 5` |
| 6 | `3/2` = 1 (dzielenie calkowite) | `3 * 1.0 / 2` = 1.5 |
| 7 | `WHERE cena > AVG(cena)` | `WHERE cena > (SELECT AVG(cena) FROM ...)` |
| 8 | `SUBSTR(tekst, 0, 3)` | `SUBSTR(tekst, 1, 3)` — od 1! |
| 9 | `A AND B OR C` = `(A AND B) OR C` | Uzywaj nawiasow: `A AND (B OR C)` |
| 10 | `LIKE 'A'` = dokladnie 'A' | `LIKE 'A%'` = zaczyna sie od A |
