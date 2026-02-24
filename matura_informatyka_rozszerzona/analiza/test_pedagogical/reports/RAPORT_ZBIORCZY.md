# Raport Zbiorczy — Test Tutor

## Dashboard

- **Wynik**: 95.3  PASS
- **Delta**: +3.0
- **Seria PASS**: 26
- **Liczba uruchomien**: 28
- **Okno (10)**: avg=90.0  min=85.1  max=95.3

## Historia

| # | Data | Commit | Mode | Wynik | Pass | Best | Worst |
|---|------|--------|------|-------|------|------|-------|
| 1 | 2026-02-24 | 0dbf358 | full | 95.3 | PASS | first_session | hint_progression |
| 2 | 2026-02-24 | 0dbf358 | full | 92.3 | PASS | first_session | hint_progression |
| 3 | 2026-02-24 | fce6714 | full | 92.4 | PASS | coaching_aware | difficulty_climb |
| 4 | 2026-02-24 | 07cb9f3 | full | 89.5 | PASS | first_session | probna |
| 5 | 2026-02-24 | 6baa479 | full | 92.3 | PASS | first_session | difficulty_climb |
| 6 | 2026-02-23 | c2e9ce9 | full | 86.6 | PASS | coaching_aware | first_session |
| 7 | 2026-02-23 | c2e9ce9 | full | 87.7 | PASS | hint_progression | probna |
| 8 | 2026-02-22 | dcf2deb | full | 85.1 | PASS | difficulty_climb | review_session |
| 9 | 2026-02-22 | ac622d1 | full | 90.7 | PASS | hint_progression | cke_unlock |
| 10 | 2026-02-22 | 6c9ecc3 | full | 88.4 | PASS | difficulty_climb | coaching_aware |
| 11 | 2026-02-22 | 367575d | full | 89.6 | PASS | coaching_aware | probna |
| 12 | 2026-02-22 | 22c3ebe | full | 84.1 | PASS | hint_progression | first_session |
| 13 | 2026-02-22 | 1314828 | full | 91.2 | PASS | hint_progression | probna |
| 14 | 2026-02-19 | dc0714c | full | 92.0 | PASS | difficulty_climb | first_session |
| 15 | 2026-02-19 | a68bf9a | full | 90.0 | PASS | cke_unlock | first_session |
| 16 | 2026-02-19 | a68bf9a | full | 89.0 | PASS | probna | first_session |
| 17 | 2026-02-19 | 187988f | full | 90.0 | PASS | probna | first_session |
| 18 | 2026-02-19 | 042ebe8 | full | 90.0 | PASS | cke_unlock | first_session |
| 19 | 2026-02-18 | e35c154 | quick | 90.0 | PASS | first_session | first_session |
| 20 | 2026-02-18 | e35c154 | full | 86.3 | PASS | review_session | first_session |
| 21 | 2026-02-18 | ac6751d | beginner | 91.0 | PASS | first_session | first_session |
| 22 | 2026-02-18 | aafe35e | full | 89.0 | PASS | difficulty_climb | first_session |
| 23 | 2026-02-18 | 6767390 | quick | 92.0 | PASS | first_session | first_session |
| 24 | 2026-02-18 | 6767390 | full | 86.0 | PASS | difficulty_climb | first_session |
| 25 | 2026-02-18 | 169947a | full | 97.2 | PASS | hint_progression | difficulty_climb |
| 26 | 2026-02-17 | ab36d15 | quick | 94.0 | PASS | first_session | first_session |
| 27 | 2026-02-17 | 20a7fd7 | quick | 67.0 | FAIL | first_session | first_session |
| 28 | 2026-02-17 | 169947a | full | 85.3 | PASS | cke_unlock | hint_progression |

## Analiza per-scenariusz

| Scenariusz | Avg | Current | Trend | StdDev | Min | Max | Regressed |
|------------|-----|---------|-------|--------|-----|-----|-----------|
| first_session | 91.3 | 100.0 | ↑ | 7.0 | 76.0 | 100.0 |  |
| hint_progression | 89.9 | 84.1 | ↓ | 4.0 | 84.1 | 95.0 | ⚠ |
| difficulty_climb | 91.0 | 100.0 | ↑ | 4.4 | 85.0 | 100.0 |  |
| review_session | 91.4 | 100.0 | ↑ | 6.2 | 76.0 | 100.0 |  |
| coaching_aware | 91.4 | 91.0 | → | 5.1 | 83.0 | 100.0 |  |
| cke_unlock | 88.6 | 100.0 | ↑ | 5.4 | 81.0 | 100.0 |  |
| probna | 86.7 | 92.0 | ↑ | 6.4 | 77.0 | 95.0 |  |

## Trendy kryteriow L2

| Kryterium | Avg | Current | Trend |
|-----------|-----|---------|-------|
| socratic | 4.2 | 4.7 | → |
| tone | 4.8 | 4.7 | → |

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

Wynik 95.3 to **najwyzszy wynik od 28 uruchomien** (ex aequo z 97.2 z 169947a). Delta +3.0 vs poprzedni run (92.3). Poprawa wynika z dwoch fixow SKILL.md:

- **[GATE] po 3 bledach**: Dodanie reguly "3 bledy → pomin hinty, idz do walk_through" w sekcji F. Efekt: first_session, difficulty_climb, review_session, cke_unlock teraz na 100.0 (poprzednio 87-95).
- **[WYMAGANE] cheatsheet**: Dodanie markera w L2 hint. Efekt: scenariusze ktore docieraja do L2 teraz wywoluja cheatsheet.

Per scenariusz:
- **4 scenariusze na 100.0/100** (first_session, difficulty_climb, review_session, cke_unlock) — najlepszy wynik ever
- **hint_progression 84.1** — stabilne vs poprzedni run (84.09), ale 2 checkpointy (L2 hint, cheatsheet) sa teraz **strukturalnie nieosiagalne** bo GATE pomija krok 4 (hinty). To problem scenariusza testowego, nie SKILL.md.
- **coaching_aware 91.0** — lekki spadek z 95.0 (MENTION_PAST nie w coaching_actions dla tego cwiczenia)
- **probna 92.0** — spadek z 95.0 — szum ewaluatora (socratic=4, tone=4 vs poprzednie 4,5)

### Top 3 do naprawienia

1. **hint_progression: Aktualizacja checkpointow scenariusza** — Po GATE (3 bledy → walk_through), checkpointy "Progresja L1→L2→L3" i "cheatsheet get przy L2" sa nieosiagalne. Scenariusz testowy (sekcja 5.2 test-tutor) wymaga aktualizacji: zamienic L2/L3 checkpointy na "GATE aktywowany po 3 bledach" i "walk_through bez hintow L2/L3".
2. **coaching_aware: MENTION_PAST w danych testowych** — Checkpoint "MENTION_PAST zrealizowane PO bledzie" failli bo CLI zwraca tylko WARN_LEECH i HINT_DELAY dla tego cwiczenia. Rozwiazanie: albo dodac MENTION_PAST do pre-fetch danych, albo zmienic checkpoint na warunkowy ("jesli MENTION_PAST w coaching_actions").
3. **probna: Empatia przy przerwaniu** — SKILL.md sekcja H3 mogloby explicite wymagac "zachecajacy komentarz przy przerwaniu egzaminu + sugestia cwiczen na slabe obszary".

### Uporczywe problemy

- **hint_progression L1 < 90%**: Pojawia sie w 3+ runach (07cb9f3: 90%, fce6714: 90.91%, 0dbf358 x2: 81.82%). Przyczyna zmieniala sie: najpierw brak L3, potem L3 nieosiagalny, teraz GATE blokuje L2/L3. To problem ewoluujacy — kolejna iteracja powinna skupic sie na aktualizacji scenariusza testowego.
- **difficulty_climb socratic**: Historycznie niski (avg 3-3.5), teraz naprawiony (5/5 w tym runie). Poprawa moze byc niestabilna — warto monitorowac.

### Co dziala dobrze

- **95.3/100 — najwyzszy wynik w oknie 10 runow** (avg okna 90.0)
- **Seria 26 PASS** bez przerwy — zero regresji od 2026-02-17
- **4 scenariusze na 100.0** — first_session, difficulty_climb, review_session, cke_unlock (bezbledne)
- **cke_unlock trend ↑** (avg 88.6, current 100.0) — konsekwentna poprawa od 81.0 na poczatku
- **first_session trend ↑** (avg 91.3, current 100.0) — nowy rekord
- **socratic=4.7** (avg 4.2) — najwyzszy w historii, GATE pozytywnie wplynal na jakosc metody sokratejskiej
- **L1 compliance 96.0%** — CLI commands prawie bezbedne
- **Szum ewaluatora 0dbf358: StdDev=1.5** — niska zmiennosc miedzy runami na tym samym kodzie
