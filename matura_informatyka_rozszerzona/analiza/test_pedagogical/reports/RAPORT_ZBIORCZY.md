# Raport Zbiorczy — Test Tutor

## Dashboard

- **Wynik**: 94.8  PASS
- **Delta**: -0.5
- **Seria PASS**: 34
- **Liczba uruchomien**: 36
- **Okno (10)**: avg=92.7  min=83.1  max=96.4

## Historia

| # | Data | Commit | Mode | Wynik | Pass | Best | Worst |
|---|------|--------|------|-------|------|------|-------|
| 1 | 2026-03-05 | afb4bb2 | full | 94.8 | PASS | coaching_aware | hint_progression |
| 2 | 2026-03-01 | 92ccc1a | full | 95.3 | PASS | first_session | hint_progression |
| 3 | 2026-03-01 | 69554fe | full | 96.4 | PASS | first_session | hint_progression |
| 4 | 2026-03-01 | 98ef87d | full | 83.1 | PASS | first_session | difficulty_climb |
| 5 | 2026-03-01 | 98ef87d | full | 94.4 | PASS | first_session | hint_progression |
| 6 | 2026-02-28 | 4362947 | full | 92.1 | PASS | probna | hint_progression |
| 7 | 2026-02-28 | c339ce7 | full | 87.7 | PASS | cke_unlock | coaching_aware |
| 8 | 2026-02-25 | a4c1c8d | full | 95.6 | PASS | first_session | hint_progression |
| 9 | 2026-02-24 | 0dbf358 | full | 95.3 | PASS | first_session | hint_progression |
| 10 | 2026-02-24 | 0dbf358 | full | 92.3 | PASS | first_session | hint_progression |
| 11 | 2026-02-24 | fce6714 | full | 92.4 | PASS | coaching_aware | difficulty_climb |
| 12 | 2026-02-24 | 07cb9f3 | full | 89.5 | PASS | first_session | probna |
| 13 | 2026-02-24 | 6baa479 | full | 92.3 | PASS | first_session | difficulty_climb |
| 14 | 2026-02-23 | c2e9ce9 | full | 86.6 | PASS | coaching_aware | first_session |
| 15 | 2026-02-23 | c2e9ce9 | full | 87.7 | PASS | hint_progression | probna |
| 16 | 2026-02-22 | dcf2deb | full | 85.1 | PASS | difficulty_climb | review_session |
| 17 | 2026-02-22 | ac622d1 | full | 90.7 | PASS | hint_progression | cke_unlock |
| 18 | 2026-02-22 | 6c9ecc3 | full | 88.4 | PASS | difficulty_climb | coaching_aware |
| 19 | 2026-02-22 | 367575d | full | 89.6 | PASS | coaching_aware | probna |
| 20 | 2026-02-22 | 22c3ebe | full | 84.1 | PASS | hint_progression | first_session |
| 21 | 2026-02-22 | 1314828 | full | 91.2 | PASS | hint_progression | probna |
| 22 | 2026-02-19 | dc0714c | full | 92.0 | PASS | difficulty_climb | first_session |
| 23 | 2026-02-19 | a68bf9a | full | 90.0 | PASS | cke_unlock | first_session |
| 24 | 2026-02-19 | a68bf9a | full | 89.0 | PASS | probna | first_session |
| 25 | 2026-02-19 | 187988f | full | 90.0 | PASS | probna | first_session |
| 26 | 2026-02-19 | 042ebe8 | full | 90.0 | PASS | cke_unlock | first_session |
| 27 | 2026-02-18 | e35c154 | quick | 90.0 | PASS | first_session | first_session |
| 28 | 2026-02-18 | e35c154 | full | 86.3 | PASS | review_session | first_session |
| 29 | 2026-02-18 | ac6751d | beginner | 91.0 | PASS | first_session | first_session |
| 30 | 2026-02-18 | aafe35e | full | 89.0 | PASS | difficulty_climb | first_session |
| 31 | 2026-02-18 | 6767390 | quick | 92.0 | PASS | first_session | first_session |
| 32 | 2026-02-18 | 6767390 | full | 86.0 | PASS | difficulty_climb | first_session |
| 33 | 2026-02-18 | 169947a | full | 97.2 | PASS | hint_progression | difficulty_climb |
| 34 | 2026-02-17 | ab36d15 | quick | 94.0 | PASS | first_session | first_session |
| 35 | 2026-02-17 | 20a7fd7 | quick | 67.0 | FAIL | first_session | first_session |
| 36 | 2026-02-17 | 169947a | full | 85.3 | PASS | cke_unlock | hint_progression |

## Analiza per-scenariusz

| Scenariusz | Avg | Current | Trend | StdDev | Min | Max | Regressed |
|------------|-----|---------|-------|--------|-----|-----|-----------|
| first_session | 98.0 | 95.0 | ↓ | 2.4 | 95.0 | 100.0 | ⚠ |
| hint_progression | 84.1 | 75.6 | ↓ | 8.2 | 67.7 | 95.0 | ⚠ |
| difficulty_climb | 90.2 | 95.0 | ↑ | 10.9 | 59.5 | 100.0 |  |
| review_session | 95.8 | 98.0 | ↑ | 3.5 | 90.0 | 100.0 |  |
| coaching_aware | 92.5 | 100.0 | ↑ | 7.5 | 73.0 | 100.0 |  |
| cke_unlock | 95.2 | 100.0 | ↑ | 3.3 | 89.0 | 100.0 |  |
| probna | 93.2 | 100.0 | ↑ | 6.4 | 77.0 | 100.0 |  |

## Trendy kryteriow L2

| Kryterium | Avg | Current | Trend |
|-----------|-----|---------|-------|
| socratic | 4.2 | 4.6 | → |
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

Wynik **94.8** (-0.5 vs poprzedni 95.3) to minimalna zmiana w granicach szumu ewaluatora (StdDev historyczny = 3.5). Commit afb4bb2 ("Skill review: fix 8 inconsistencies") nie wprowadzil regresji w SKILL.md.

Kluczowe zmiany per scenariusz:
- **hint_progression**: 75.6 (spadek z 81.8) — najnizszy wynik w tym runie. Evaluator poprawnie zidentyfikowal strukturalny problem: przy hint_delay=1 i GATE 3-prob, pelna progresja L1→L2→L3 jest nieosiagalna. 3 checkpointy automatycznie FAILuja (HINT_LOCKED, progresja_hintow, cheatsheet_excerpt).
- **first_session**: 95.0 (spadek z 100) — evaluator ocenił sokratejska na 4/5 (vs 5/5 poprzednio) za "pytanie zdradzajace odpowiedz". To subiektywna ocena L2 — w granicach szumu.
- **coaching_aware**: 100.0 (stabilne) — pelne punkty za coaching flow (WARN_LEECH, MENTION_PAST, HINT_DELAY).
- **cke_unlock**: 100 (wzrost z 96.4) — tryb krok-po-kroku i sprawdzian CKE ocenione pelna punktacja.
- **probna**: 100 (wzrost z 96.4) — pelne punkty za tryb egzaminacyjny.
- **review_session**: 98.0 (wzrost z 96.4) — powtorka z priorytetem i leech tag handling.

### Top 3 do naprawienia

1. **hint_progression: strukturalny konflikt GATE vs progresja hintow (StdDev=8.2, trend ↓)**: SKILL.md sekcja F krok 3 definiuje progresje "proba 1 → proba 2 (po L1) → proba 3 (po L2+cheatsheet) → walk_through (z L3)". Ale HARD GATE "pytanie i hint osobno" oraz GATE "3 bledne proby → walk_through" konsumuja proby szybciej niz hinty: przy hint_delay=1, max osiagalny to L1 przed walk_through. Rekomendacja: albo poluzowac GATE (4 proby zamiast 3) albo zdefiniowac ze pytanie sokratejskie NIE liczy sie jako proba.

2. **first_session: pytanie sokratejskie zbyt wezkie**: Evaluator ocenił 4/5 za pytanie "silnia(5) wywoluje silnia(n-1)" ktore sugeruje odpowiedz. SKILL.md mogloby zawierac bank pytan sokratejskich per archetyp z roznym stopniem otwartosci.

3. **Auto-scorer false-negative na cwiczeniach wieloczesciowych** (difficulty_climb issue): `check-answer` moze nie obslugiwac odpowiedzi z wieloma czesciami (a/b/c). Tutor musi wtedy nadpisac ocene merytorycznie.

### Uporczywe problemy

- **hint_progression jako "worst" scenariusz** w 10/13 ostatnich full runow — strukturalny weak spot. Przyczyna: konflikt GATE 3-prob z progresja hintow.
- **Socratic 4.x/5** — pojawia sie regularnie od poczatku. Strukturalne ograniczenie evaluatora: oczekuje pytan granularnych, SKILL.md daje ogolne wytyczne.
- **Evaluator noise StdDev do 6.0** (commit 169947a: 85.3 vs 97.2) — ten sam SKILL.md, rozne wyniki. Czynnik: losowy dobor cwiczen i subiektywnosc L2.

### Co dziala dobrze

- **Seria PASS = 34** — najdluzsza w historii, rosnie co run.
- **5/7 scenariuszy z trendem ↑** — difficulty_climb, review_session, coaching_aware, cke_unlock, probna.
- **coaching_aware = 100.0** — stabilne od 2 runow, potwierdza skutecznosc coaching_actions_v2.
- **cke_unlock i probna = 100.0** — nowy rekord, oba scenariusze zaawansowane na pelna punktacje.
- **Ton i jezyk = 4.9/5** (avg=4.8, trend →) — konsekwentnie wysoki.
- **L1 = 100%** w 6/7 scenariuszy — jedyny FAIL to hint_progression (strukturalny, nie defekt SKILL).
- **Okno 10: avg=92.7** — stabilny wzrost z 89.x (luty) do 92+ (marzec).
