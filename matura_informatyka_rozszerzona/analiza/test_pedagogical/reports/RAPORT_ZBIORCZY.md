# Raport Zbiorczy — Test Tutor

## Dashboard

- **Wynik**: 96.4  PASS
- **Delta**: +13.3
- **Seria PASS**: 32
- **Liczba uruchomien**: 34
- **Okno (10)**: avg=91.9  min=83.1  max=96.4

## Historia

| # | Data | Commit | Mode | Wynik | Pass | Best | Worst |
|---|------|--------|------|-------|------|------|-------|
| 1 | 2026-03-01 | 69554fe | full | 96.4 | PASS | first_session | hint_progression |
| 2 | 2026-03-01 | 98ef87d | full | 83.1 | PASS | first_session | difficulty_climb |
| 3 | 2026-03-01 | 98ef87d | full | 94.4 | PASS | first_session | hint_progression |
| 4 | 2026-02-28 | 4362947 | full | 92.1 | PASS | probna | hint_progression |
| 5 | 2026-02-28 | c339ce7 | full | 87.7 | PASS | cke_unlock | coaching_aware |
| 6 | 2026-02-25 | a4c1c8d | full | 95.6 | PASS | first_session | hint_progression |
| 7 | 2026-02-24 | 0dbf358 | full | 95.3 | PASS | first_session | hint_progression |
| 8 | 2026-02-24 | 0dbf358 | full | 92.3 | PASS | first_session | hint_progression |
| 9 | 2026-02-24 | fce6714 | full | 92.4 | PASS | coaching_aware | difficulty_climb |
| 10 | 2026-02-24 | 07cb9f3 | full | 89.5 | PASS | first_session | probna |

## Analiza per-scenariusz

| Scenariusz | Avg | Current | Trend | StdDev | Min | Max | Regressed |
|------------|-----|---------|-------|--------|-----|-----|-----------|
| first_session | 98.0 | 100.0 | ↑ | 2.4 | 95.0 | 100.0 |  |
| hint_progression | 85.9 | 95.0 | ↑ | 7.7 | 67.7 | 95.0 |  |
| difficulty_climb | 88.5 | 95.0 | ↑ | 10.6 | 59.5 | 100.0 |  |
| review_session | 94.8 | 100.0 | ↑ | 3.8 | 90.0 | 100.0 |  |
| coaching_aware | 92.0 | 95.0 | ↑ | 7.2 | 73.0 | 100.0 |  |
| cke_unlock | 93.1 | 95.0 | ↑ | 4.1 | 85.0 | 100.0 |  |
| probna | 91.0 | 95.0 | ↑ | 7.2 | 77.0 | 100.0 |  |

## Trendy kryteriow L2

| Kryterium | Avg | Current | Trend |
|-----------|-----|---------|-------|
| socratic | 4.1 | 4.3 | → |
| tone | 4.8 | 5.0 | → |

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

Wynik **96.4** to najwyzszy wynik od czasu 97.2 (commit 169947a, 2026-02-18) i drugi najlepszy w calej historii. Delta +13.3 vs poprzedni run (83.1) jest duza, ale porownanie z ostatnim full runem na tym samym SKILL (98ef87d: 94.4) pokazuje stabilna poprawe o +2.0 pkt. Commit 69554fe wprowadzil 5 poprawek z test-tutor findings (arrow normalization, multi-part answers, MENTION_PAST, SKILL.md coaching) — efekty widoczne we wszystkich scenariuszach.

Kluczowe zmiany:
- **hint_progression**: 95.0 (vs 89.5 i 67.7 w poprzednich runach) — L1 wzroslo z 90.9%/63.6% do 100%. Poprawki do SKILL.md sekcji F (hint flow) wyeliminowaly problemy z progresja L1→L2→L3.
- **probna**: 95.0 (vs 94.5 i 87.5) — stabilna poprawa, exam save i progress blad teraz konsekwentne.
- **first_session** i **review_session**: 100.0 — idealne wyniki, utrzymane od kilku runow.

### Top 3 do naprawienia

1. **check-answer false negatives for multi-part exercises**: CLI `normalize.go` `splitMultiPart` porownuje pelny tekst (z tabelami sledzenia) zamiast samych wynikow liczbowych. Exercise 1.17: student daje "a) 5, b) 3, c) 4" — poprawne, ale check-answer zwraca trafione_parts=0/3. Wymaga fix w normalize.go.

2. **suggest-error HARD GATE vs probna.md**: SKILL.md F.3 mowi "MUSISZ wywolac suggest-error PRZED progress blad" globalnie, ale probna.md nie wspomina o suggest-error. Dla CKE ID (YYYYM.Z.S) suggest-error moze nie dzialac. Rekomendacja: dodac w probna.md explicite wyjtek lub alias suggest-error dla CKE IDs.

3. **Metoda sokratejska granulacja**: L2 socratic srednio 4.3/5 — konsekwentnie tracony 1 pkt za zbyt malo granularne pytania diagnostyczne (np. "gdzie blad?" zamiast "oblicz 12 mod 3 krok po kroku"). SKILL.md sekcja E mogloby zawierac przyklady pytan per archetyp.

### Uporczywe problemy

- **Socratic 4/5** w 5/7 scenariuszy — pojawia sie w kazdym runie od poczatku. To nie jest blad SKILL.md, lecz ograniczenie symulacji (evaluator ocenia hipotetyczne zachowanie tutora).
- **hint_progression StdDev=7.7** — najwyzsza zmiennosc ze wszystkich scenariuszy. Flow proba→sokratejskie→hint jest zlozony i rozni evaluatorzy interpretuja go inaczej.
- **check-answer false negatives** — zglaszane w 3+ runach (difficulty_climb issues). Wymaga fix w CLI.

### Co dziala dobrze

- **L1 = 100%** we wszystkich 7 scenariuszach — pierwszy raz w historii! Wszystkie checkpointy CLI, coaching i scenario-specific spelnione.
- **Ton i jezyk = 5.0/5** we wszystkich scenariuszach — stabilne od 10+ runow.
- **first_session** i **review_session** = 100.0 — konsekwentnie najlepsze scenariusze.
- **Seria PASS = 32** — od drugiego runu w historii.
- **Wszystkie trendy ↑** — kazdy scenariusz ponad srednia historyczna.
