# Raport Zbiorczy — Test Tutor

## Dashboard

- **Wynik**: 83.1  PASS
- **Delta**: -11.3
- **Seria PASS**: 31
- **Liczba uruchomien**: 33
- **Okno (10)**: avg=91.5  min=83.1  max=95.6

## Historia

| # | Data | Commit | Mode | Wynik | Pass | Best | Worst |
|---|------|--------|------|-------|------|------|-------|
| 1 | 2026-03-01 | 98ef87d | full | 83.1 | PASS | first_session | difficulty_climb |
| 2 | 2026-03-01 | 98ef87d | full | 94.4 | PASS | first_session | hint_progression |
| 3 | 2026-02-28 | 4362947 | full | 92.1 | PASS | probna | hint_progression |
| 4 | 2026-02-28 | c339ce7 | full | 87.7 | PASS | cke_unlock | coaching_aware |
| 5 | 2026-02-25 | a4c1c8d | full | 95.6 | PASS | first_session | hint_progression |
| 6 | 2026-02-24 | 0dbf358 | full | 95.3 | PASS | first_session | hint_progression |
| 7 | 2026-02-24 | 0dbf358 | full | 92.3 | PASS | first_session | hint_progression |
| 8 | 2026-02-24 | fce6714 | full | 92.4 | PASS | coaching_aware | difficulty_climb |
| 9 | 2026-02-24 | 07cb9f3 | full | 89.5 | PASS | first_session | probna |
| 10 | 2026-02-24 | 6baa479 | full | 92.3 | PASS | first_session | difficulty_climb |

## Analiza per-scenariusz

| Scenariusz | Avg | Current | Trend | StdDev | Min | Max | Regressed |
|------------|-----|---------|-------|--------|-----|-----|-----------|
| first_session | 97.5 | 100.0 | ↑ | 2.5 | 95.0 | 100.0 |  |
| hint_progression | 85.9 | 67.7 | ↓ | 7.7 | 67.7 | 95.0 | ⚠ |
| difficulty_climb | 87.5 | 59.5 | ↓ | 10.4 | 59.5 | 100.0 | ⚠ |
| review_session | 94.3 | 90.5 | ↓ | 3.4 | 90.0 | 100.0 | ⚠ |
| coaching_aware | 92.0 | 86.0 | ↓ | 7.2 | 73.0 | 100.0 |  |
| cke_unlock | 92.5 | 90.4 | ↓ | 4.2 | 85.0 | 100.0 |  |
| probna | 90.7 | 87.5 | ↓ | 7.1 | 77.0 | 100.0 |  |

## Trendy kryteriow L2

| Kryterium | Avg | Current | Trend |
|-----------|-----|---------|-------|
| socratic | 4.0 | 3.9 | → |
| tone | 4.8 | 4.3 | → |

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

Wynik spadl z 94.4 do **83.1** (-11.3 pkt) na **tym samym commicie** (98ef87d). To **czysto szum ewaluatora** — StdDev=5.7, najwyzszy w historii duplikatow. Dwa scenariusze spadly ponizej progu 70:

- **difficulty_climb** (95.0 → 59.5, -35.5): agent ewaluatora wyczerpal 15 tur na uruchamianie prawdziwych komend CLI zamiast symulowac dialog. Po wznowieniu wyprodukowal surowsza ale mniej poinformowana ocene — nie mial pelnej symulacji do oceny.
- **hint_progression** (89.5 → 67.7, -21.8): surowsza interpretacja — `exercise next` vs pre-fetched cwiczenie, `--hint 0` zamiast `--hint 1` po hincie L1, brak pelnej progresji L1→L2→L3 (walk_through gate po 3 probach).

Pozostale scenariusze: review_session -1.5, coaching_aware -9.0, cke_unlock -4.6, probna -7.0 — wszystkie w granicach typowego szumu.

### Top 3 do naprawienia

1. **hint_progression: L1→L2→L3 tension** — Scenariusz wymaga 3 poziomow hintow, ale SKILL.md GATE po 3 probach wymusza walk_through zanim L2/L3 zostana podane. Fixed script ma 5 wymian ale 3 z nich to bledy → 3-error GATE aktywuje sie przed L2. **Fix**: dodac 4. probe do fixed script LUB zrelaksowac checkpoint ("L1 wystarczy jesli GATE zatrzymal").

2. **difficulty_climb: coaching_actions_v2 wplecenie** — Gdy exercise next zwraca coaching_actions_v2 (np. HINT_DELAY po awansie na higher difficulty), tutor MUSI je wplesc PRZED trescia cwiczenia (SKILL.md E2). Ewaluator surowo ocenia pominiecie. To jest realny problem SKILL.md compliance.

3. **coaching_aware: MENTION_PAST** — Checkpoint wymaga zrealizowania MENTION_PAST, ale coaching_actions_v2 cwiczenia 7.7 nie zawiera tej akcji (tylko WARN_LEECH + HINT_DELAY). **Fix**: albo dodac MENTION_PAST do pre-fetched danych scenariusza, albo uznac F.3 WARN_LEECH linkage za rownowazny.

### Uporczywe problemy (3+ kolejne uruchomienia)

- **hint_progression najslabszy scenariusz**: "worst" w 7 z 10 ostatnich uruchomien. Przyczyna strukturalna: hint_delay=1 + 3-error GATE = L2/L3 nigdy nie podane w pelni. To design issue scenariusza.
- **difficulty_climb niestabilny**: StdDev=10.4, najwyzszy ze wszystkich scenariuszy. Ewaluatorzy oceniaja go bardzo rozmaicie (59.5–100.0).
- **progress_status na starcie probnej**: zgloszone w 2 kolejnych uruchomieniach — latwe do pominiecia gdy uczen wchodzi z komenda "probna YYYY".

### Co dziala dobrze

- **first_session**: 100.0 (↑), najstabilniejszy scenariusz (avg=97.5, StdDev=2.5) — pelna zgodnosc ze SKILL.md
- **31 consecutive PASS** — SKILL.md nie spadla ponizej progu globalnego od poczatku (po jednym FAIL prototypu z 17 lut)
- **cke_unlock stabilny**: 90.4 (avg=92.5) — zlozony scenariusz z 13 checkpointami konsekwentnie dobrze
- **L2 socratic stabilne ~4.0** — metoda sokratejska konsekwentnie realizowana bez regresji
- **L2 tone stabilne ~4.3–4.8** — ton, jezyk, brak emoji bez spadkow
- **probna 87.5**: exam save poprawiony (PASS vs FAIL w poprzednim runie), progress blad w trybie egzaminowym rejestrowany
