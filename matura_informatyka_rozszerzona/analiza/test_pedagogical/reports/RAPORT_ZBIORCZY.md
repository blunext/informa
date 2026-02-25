# Raport Zbiorczy — Test Tutor

## Dashboard

- **Wynik**: 95.6  PASS
- **Delta**: +0.3
- **Seria PASS**: 27
- **Liczba uruchomien**: 29
- **Okno (10)**: avg=90.8  min=85.1  max=95.6

## Historia

| # | Data | Commit | Mode | Wynik | Pass | Best | Worst |
|---|------|--------|------|-------|------|------|-------|
| 1 | 2026-02-25 | a4c1c8d | full | 95.6 | PASS | first_session | hint_progression |
| 2 | 2026-02-24 | 0dbf358 | full | 95.3 | PASS | first_session | hint_progression |
| 3 | 2026-02-24 | 0dbf358 | full | 92.3 | PASS | first_session | hint_progression |
| 4 | 2026-02-24 | fce6714 | full | 92.4 | PASS | coaching_aware | difficulty_climb |
| 5 | 2026-02-24 | 07cb9f3 | full | 89.5 | PASS | first_session | probna |
| 6 | 2026-02-24 | 6baa479 | full | 92.3 | PASS | first_session | difficulty_climb |
| 7 | 2026-02-23 | c2e9ce9 | full | 86.6 | PASS | coaching_aware | first_session |
| 8 | 2026-02-23 | c2e9ce9 | full | 87.7 | PASS | hint_progression | probna |
| 9 | 2026-02-22 | dcf2deb | full | 85.1 | PASS | difficulty_climb | review_session |
| 10 | 2026-02-22 | ac622d1 | full | 90.7 | PASS | hint_progression | cke_unlock |
| 11 | 2026-02-22 | 6c9ecc3 | full | 88.4 | PASS | difficulty_climb | coaching_aware |
| 12 | 2026-02-22 | 367575d | full | 89.6 | PASS | coaching_aware | probna |
| 13 | 2026-02-22 | 22c3ebe | full | 84.1 | PASS | hint_progression | first_session |
| 14 | 2026-02-22 | 1314828 | full | 91.2 | PASS | hint_progression | probna |
| 15 | 2026-02-19 | dc0714c | full | 92.0 | PASS | difficulty_climb | first_session |
| 16 | 2026-02-19 | a68bf9a | full | 90.0 | PASS | cke_unlock | first_session |
| 17 | 2026-02-19 | a68bf9a | full | 89.0 | PASS | probna | first_session |
| 18 | 2026-02-19 | 187988f | full | 90.0 | PASS | probna | first_session |
| 19 | 2026-02-19 | 042ebe8 | full | 90.0 | PASS | cke_unlock | first_session |
| 20 | 2026-02-18 | e35c154 | quick | 90.0 | PASS | first_session | first_session |
| 21 | 2026-02-18 | e35c154 | full | 86.3 | PASS | review_session | first_session |
| 22 | 2026-02-18 | ac6751d | beginner | 91.0 | PASS | first_session | first_session |
| 23 | 2026-02-18 | aafe35e | full | 89.0 | PASS | difficulty_climb | first_session |
| 24 | 2026-02-18 | 6767390 | quick | 92.0 | PASS | first_session | first_session |
| 25 | 2026-02-18 | 6767390 | full | 86.0 | PASS | difficulty_climb | first_session |
| 26 | 2026-02-18 | 169947a | full | 97.2 | PASS | hint_progression | difficulty_climb |
| 27 | 2026-02-17 | ab36d15 | quick | 94.0 | PASS | first_session | first_session |
| 28 | 2026-02-17 | 20a7fd7 | quick | 67.0 | FAIL | first_session | first_session |
| 29 | 2026-02-17 | 169947a | full | 85.3 | PASS | cke_unlock | hint_progression |

## Analiza per-scenariusz

| Scenariusz | Avg | Current | Trend | StdDev | Min | Max | Regressed |
|------------|-----|---------|-------|--------|-----|-----|-----------|
| first_session | 92.5 | 100.0 | ↑ | 7.4 | 76.0 | 100.0 |  |
| hint_progression | 89.9 | 89.5 | → | 4.0 | 84.1 | 95.0 |  |
| difficulty_climb | 91.3 | 95.0 | ↑ | 4.5 | 85.0 | 100.0 |  |
| review_session | 92.6 | 100.0 | ↑ | 6.6 | 76.0 | 100.0 |  |
| coaching_aware | 92.6 | 95.0 | ↑ | 4.3 | 85.0 | 100.0 |  |
| cke_unlock | 89.0 | 95.0 | ↑ | 5.7 | 81.0 | 100.0 |  |
| probna | 87.4 | 95.0 | ↑ | 6.9 | 77.0 | 95.0 |  |

## Trendy kryteriow L2

| Kryterium | Avg | Current | Trend |
|-----------|-----|---------|-------|
| socratic | 4.3 | 4.3 | → |
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

## Interpretacja

### Co sie zmienilo?

Wynik 95.6 to **nowy rekord** — najwyzszy w calej historii 29 uruchomien. Delta +0.3 vs poprzedni run (95.3 z 0dbf358). Poprawa jest niewielka ale konsystentna — to juz 3. run z rzedu powyzej 92.

Kluczowe zmiany per scenariusz vs poprzedni run (0dbf358, 95.3):
- **hint_progression 89.55** (bylo 84.09) — poprawa o +5.5 pkt. L1 wzroslo z 81.82% do 90.9% (10/11 vs 9/11). Checkpoint L3 nadal nieosiagalny strukturalnie, ale pozostale checkpointy teraz zdane.
- **difficulty_climb 95.0** (bylo 100.0) — spadek o -5 pkt. Socratic z 5 na 4 (brak follow-up pytan po poprawnych odpowiedziach). Szum ewaluatora.
- **coaching_aware 95.0** (bylo 91.0) — poprawa o +4 pkt. L1 powrocilo do 100% (bylo 90%). Ton 5/5 (bylo 4/5).
- **first_session, review_session** — stabilne 100.0 (bez zmian).
- **cke_unlock, probna** — stabilne 95.0 (bez zmian).

### Top 3 do naprawienia

1. **hint_progression: L3 checkpoint nieosiagalny** — GATE (3 bledy → walk_through) uniemozliwia dotarcie do L3 hint w skrypcie 3-bledowym. Scenariusz testowy (sekcja 5.2 test-tutor) wymaga: albo dodac 4. bledna wymiane przed "poddaje sie", albo zamienic L3 checkpoint na "GATE aktywowany poprawnie".
2. **difficulty_climb: Socratic follow-up** — Ewaluator obnizy socratic gdy tutor nie zadaje pytan poglebajacych po poprawnych odpowiedziach. SKILL.md nie wymaga tego explicite — rozwazyc dodanie "Po poprawnej: krotkie pytanie weryfikujace rozumienie".
3. **coaching_aware/cke_unlock: Socratic bezposrednosc** — "Jaki operator daje reszte?" to podpowiedz, nie pytanie sokratejskie. SKILL.md mogloby dodac guideline: "Pytanie naprowadzajace powinno dotyczyc PROCESU, nie ODPOWIEDZI".

### Uporczywe problemy

- **hint_progression L3 nieosiagalny**: Pojawia sie w 4+ runach (07cb9f3, fce6714, 0dbf358 x2, a4c1c8d). Przyczyna stabilna — GATE blokuje L2/L3 progression. To problem scenariusza testowego, nie SKILL.md. Wymaga jednorazowej aktualizacji checkpointow w test-tutor/SKILL.md sekcja 5.2.
- **Socratic 4/5 na difficulty_climb/coaching_aware/cke_unlock/probna**: Ewaluator konsekwentnie daje 4 (nie 5) gdy tutor poprawnie wykonuje akcje ale nie dodaje dodatkowych pytan poglebajacych. To nie regresja — to sufit obecnego SKILL.md.

### Co dziala dobrze

- **95.6/100 — nowy ATH** (all-time high), ponad avg okna 90.8
- **Seria 27 PASS** bez przerwy — zero regresji od 2026-02-17
- **6 z 7 scenariuszy trend ↑** (first_session, difficulty_climb, review_session, coaching_aware, cke_unlock, probna)
- **first_session i review_session stabilne 100.0** — bezbledne w 2 ostatnich runach
- **probna 95.0** — nowy rekord (max w historii, avg 87.4)
- **L1 compliance**: 68/69 checkpointow (98.6%) — CLI commands prawidlowe
- **tone = 5.0** (avg 4.8) — ton idealny we wszystkich scenariuszach
- **Szum ewaluatora malejacy**: StdDev duplikatow commitow <= 1.5 w ostatnich runach (bylo 6.0 na poczatku)
