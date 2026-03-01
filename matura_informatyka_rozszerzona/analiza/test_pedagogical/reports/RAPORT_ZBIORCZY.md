# Raport Zbiorczy — Test Tutor

## Dashboard

- **Wynik**: 95.3  PASS
- **Delta**: -1.1
- **Seria PASS**: 33
- **Liczba uruchomien**: 35
- **Okno (10)**: avg=92.5  min=83.1  max=96.4

## Historia

| # | Data | Commit | Mode | Wynik | Pass | Best | Worst |
|---|------|--------|------|-------|------|------|-------|
| 1 | 2026-03-01 | 92ccc1a | full | 95.3 | PASS | first_session | hint_progression |
| 2 | 2026-03-01 | 69554fe | full | 96.4 | PASS | first_session | hint_progression |
| 3 | 2026-03-01 | 98ef87d | full | 83.1 | PASS | first_session | difficulty_climb |
| 4 | 2026-03-01 | 98ef87d | full | 94.4 | PASS | first_session | hint_progression |
| 5 | 2026-02-28 | 4362947 | full | 92.1 | PASS | probna | hint_progression |
| 6 | 2026-02-28 | c339ce7 | full | 87.7 | PASS | cke_unlock | coaching_aware |
| 7 | 2026-02-25 | a4c1c8d | full | 95.6 | PASS | first_session | hint_progression |
| 8 | 2026-02-24 | 0dbf358 | full | 95.3 | PASS | first_session | hint_progression |
| 9 | 2026-02-24 | 0dbf358 | full | 92.3 | PASS | first_session | hint_progression |
| 10 | 2026-02-24 | fce6714 | full | 92.4 | PASS | coaching_aware | difficulty_climb |
| 11 | 2026-02-24 | 07cb9f3 | full | 89.5 | PASS | first_session | probna |
| 12 | 2026-02-24 | 6baa479 | full | 92.3 | PASS | first_session | difficulty_climb |
| 13 | 2026-02-23 | c2e9ce9 | full | 86.6 | PASS | coaching_aware | first_session |
| 14 | 2026-02-23 | c2e9ce9 | full | 87.7 | PASS | hint_progression | probna |
| 15 | 2026-02-22 | dcf2deb | full | 85.1 | PASS | difficulty_climb | review_session |
| 16 | 2026-02-22 | ac622d1 | full | 90.7 | PASS | hint_progression | cke_unlock |
| 17 | 2026-02-22 | 6c9ecc3 | full | 88.4 | PASS | difficulty_climb | coaching_aware |
| 18 | 2026-02-22 | 367575d | full | 89.6 | PASS | coaching_aware | probna |
| 19 | 2026-02-22 | 22c3ebe | full | 84.1 | PASS | hint_progression | first_session |
| 20 | 2026-02-22 | 1314828 | full | 91.2 | PASS | hint_progression | probna |
| 21 | 2026-02-19 | dc0714c | full | 92.0 | PASS | difficulty_climb | first_session |
| 22 | 2026-02-19 | a68bf9a | full | 90.0 | PASS | cke_unlock | first_session |
| 23 | 2026-02-19 | a68bf9a | full | 89.0 | PASS | probna | first_session |
| 24 | 2026-02-19 | 187988f | full | 90.0 | PASS | probna | first_session |
| 25 | 2026-02-19 | 042ebe8 | full | 90.0 | PASS | cke_unlock | first_session |
| 26 | 2026-02-18 | e35c154 | quick | 90.0 | PASS | first_session | first_session |
| 27 | 2026-02-18 | e35c154 | full | 86.3 | PASS | review_session | first_session |
| 28 | 2026-02-18 | ac6751d | beginner | 91.0 | PASS | first_session | first_session |
| 29 | 2026-02-18 | aafe35e | full | 89.0 | PASS | difficulty_climb | first_session |
| 30 | 2026-02-18 | 6767390 | quick | 92.0 | PASS | first_session | first_session |
| 31 | 2026-02-18 | 6767390 | full | 86.0 | PASS | difficulty_climb | first_session |
| 32 | 2026-02-18 | 169947a | full | 97.2 | PASS | hint_progression | difficulty_climb |
| 33 | 2026-02-17 | ab36d15 | quick | 94.0 | PASS | first_session | first_session |
| 34 | 2026-02-17 | 20a7fd7 | quick | 67.0 | FAIL | first_session | first_session |
| 35 | 2026-02-17 | 169947a | full | 85.3 | PASS | cke_unlock | hint_progression |

## Analiza per-scenariusz

| Scenariusz | Avg | Current | Trend | StdDev | Min | Max | Regressed |
|------------|-----|---------|-------|--------|-----|-----|-----------|
| first_session | 98.0 | 100.0 | ↑ | 2.4 | 95.0 | 100.0 |  |
| hint_progression | 85.5 | 81.8 | ↓ | 7.8 | 67.7 | 95.0 |  |
| difficulty_climb | 89.4 | 96.4 | ↑ | 10.8 | 59.5 | 100.0 |  |
| review_session | 95.0 | 96.4 | ↑ | 3.8 | 90.0 | 100.0 |  |
| coaching_aware | 92.5 | 100.0 | ↑ | 7.5 | 73.0 | 100.0 |  |
| cke_unlock | 94.2 | 96.4 | ↑ | 3.2 | 89.0 | 100.0 |  |
| probna | 92.7 | 96.4 | ↑ | 6.0 | 77.0 | 100.0 |  |

## Trendy kryteriow L2

| Kryterium | Avg | Current | Trend |
|-----------|-----|---------|-------|
| socratic | 4.1 | 4.3 | → |
| tone | 4.8 | 4.9 | → |

## Szum ewaluatora (duplikaty commitow)

| Commit | Scores | StdDev |
|--------|--------|--------|
| 169947a | 85.3, 97.2 | 6.0 |
| 6767390 | 86.0, 92.0 | 3.0 |
| e35c154 | 86.3, 90.0 | 1.9 |
| a68bf9a | 89.0, 90.0 | 0.5 |
| c2e9ce9 | 87.7, 86.6 | 0.6 |
| 0dbf358 | 92.3, 95.3 | 1.5 |
| 98ef87d | 94.4, 83.1 | 5.7 |

## Interpretacja

### Co sie zmienilo?

Wynik **95.3** (-1.1 vs poprzedni 96.4) to minimalny spadek w granicach szumu ewaluatora. Commit 92ccc1a ("deterministic pedagogy: 5 CLI features") nie wprowadzil regresji w SKILL.md — roznica wynika z losowego doboru cwiczenia w scenariuszu hint_progression.

Kluczowe zmiany per scenariusz:
- **hint_progression**: 81.8 (spadek z 95.0) — cwiczenie 1.17 ma tylko 1 hint (H1), wiec progresja L1→L2→L3 i cheatsheet_excerpt przy L2 sa niemozliwe. To limitacja danych testowych, nie defekt SKILL.md. Poprzedni run mial cwiczenie z 3 hintami.
- **coaching_aware**: 100.0 (wzrost z 95.0) — all-time high! Poprawki MENTION_PAST i coaching z commitu 69554fe skuteczne.
- **difficulty_climb**, **cke_unlock**, **probna**: 96.4 (wzrost z 95.0) — stabilna poprawa.
- **first_session**: 100.0 — stabilne od wielu runow.
- **L1 = 100%** w 6/7 scenariuszy (wszystkie oprocz hint_progression) — pierwszy raz z tak wysoka sprzedaza L1.

### Top 3 do naprawienia

1. **hint_progression volatility (StdDev=7.8, jedyny ↓ trend)**: Scenariusz jest wrazliwy na losowy dobor cwiczenia — jesli wypadnie cwiczenie z 1 hintem, 2 checkpointy (progresja_hintow, cheatsheet_excerpt_przy_L2) automatycznie FAILuja. Rekomendacja: dodac filtr `--min-hints 3` w pre-fetch lub uzyc deterministycznego ID cwiczenia dla tego scenariusza.

2. **Metoda sokratejska stabilnie 4.3/5 (trend →)**: Evaluatory konsekwentnie odejmuja punkt za zbyt ogolne pytania diagnostyczne. SKILL.md sekcja E mogloby zawierac przyklady pytan sokratejskich per archetyp (np. "Oblicz 427 mod 10 krok po kroku" zamiast "Gdzie jest blad?").

3. **probna: brak empatycznego komentarza przy przerwaniu**: L2 tone = 4/5 w probna (jedyny scenariusz < 5/5). Dodanie w probna.md zdania "Rozumiem, przerwanie to dobra decyzja — zobaczmy co udalo Ci sie zrobic" mogloby to naprawic.

### Uporczywe problemy

- **hint_progression jako "worst" scenariusz** w 9/12 ostatnich full runow — najczestszy weak spot. Glowna przyczyna: zmiennosc danych testowych (liczba hintow per cwiczenie).
- **Socratic 4.x/5** — pojawia sie w kazdym runie od poczatku (35 runow). Strukturalne ograniczenie: evaluator oczekuje granularnych pytan diagnostycznych, ktore sa trudne do wyspecyfikowania w SKILL.md bez overengineering.
- **Evaluator noise StdDev do 6.0** (commit 169947a: 85.3 vs 97.2) — ten sam SKILL.md, rozne wyniki. Czynnik: losowy dobor cwiczen i subiektywnosc L2.

### Co dziala dobrze

- **Seria PASS = 33** — najdluzsza seria w historii, rosnie co run.
- **6/7 scenariuszy z trendem ↑** — jedyny ↓ to hint_progression (z powodu danych, nie SKILL).
- **coaching_aware = 100.0** — all-time high, potwierdza skutecznosc poprawek MENTION_PAST/WARN_LEECH.
- **first_session = 100.0** (StdDev=2.4, avg=98.0) — najstabilniejszy scenariusz, benchmark jakosci.
- **Ton i jezyk = 4.9/5** (avg=4.8) — konsekwentnie wysoki, polski "ty" bez emoji.
- **cke_unlock avg=94.2** (najwyzszy avg ze wszystkich scenariuszy oprocz first_session) — sprawdzianowy flow dziala solidnie.
