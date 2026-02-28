# Raport Zbiorczy — Test Tutor

## Dashboard

- **Wynik**: 92.1  PASS
- **Delta**: +4.4
- **Seria PASS**: 29
- **Liczba uruchomien**: 31
- **Okno (10)**: avg=91.2  min=86.6  max=95.6

## Historia

| # | Data | Commit | Mode | Wynik | Pass | Best | Worst |
|---|------|--------|------|-------|------|------|-------|
| 1 | 2026-02-28 | 4362947 | full | 92.1 | PASS | probna | hint_progression |
| 2 | 2026-02-28 | c339ce7 | full | 87.7 | PASS | cke_unlock | coaching_aware |
| 3 | 2026-02-25 | a4c1c8d | full | 95.6 | PASS | first_session | hint_progression |
| 4 | 2026-02-24 | 0dbf358 | full | 95.3 | PASS | first_session | hint_progression |
| 5 | 2026-02-24 | 0dbf358 | full | 92.3 | PASS | first_session | hint_progression |
| 6 | 2026-02-24 | fce6714 | full | 92.4 | PASS | coaching_aware | difficulty_climb |
| 7 | 2026-02-24 | 07cb9f3 | full | 89.5 | PASS | first_session | probna |
| 8 | 2026-02-24 | 6baa479 | full | 92.3 | PASS | first_session | difficulty_climb |
| 9 | 2026-02-23 | c2e9ce9 | full | 86.6 | PASS | coaching_aware | first_session |
| 10 | 2026-02-23 | c2e9ce9 | full | 87.7 | PASS | hint_progression | probna |
| 11 | 2026-02-22 | dcf2deb | full | 85.1 | PASS | difficulty_climb | review_session |
| 12 | 2026-02-22 | ac622d1 | full | 90.7 | PASS | hint_progression | cke_unlock |
| 13 | 2026-02-22 | 6c9ecc3 | full | 88.4 | PASS | difficulty_climb | coaching_aware |
| 14 | 2026-02-22 | 367575d | full | 89.6 | PASS | coaching_aware | probna |
| 15 | 2026-02-22 | 22c3ebe | full | 84.1 | PASS | hint_progression | first_session |
| 16 | 2026-02-22 | 1314828 | full | 91.2 | PASS | hint_progression | probna |
| 17 | 2026-02-19 | dc0714c | full | 92.0 | PASS | difficulty_climb | first_session |
| 18 | 2026-02-19 | a68bf9a | full | 90.0 | PASS | cke_unlock | first_session |
| 19 | 2026-02-19 | a68bf9a | full | 89.0 | PASS | probna | first_session |
| 20 | 2026-02-19 | 187988f | full | 90.0 | PASS | probna | first_session |
| 21 | 2026-02-19 | 042ebe8 | full | 90.0 | PASS | cke_unlock | first_session |
| 22 | 2026-02-18 | e35c154 | quick | 90.0 | PASS | first_session | first_session |
| 23 | 2026-02-18 | e35c154 | full | 86.3 | PASS | review_session | first_session |
| 24 | 2026-02-18 | ac6751d | beginner | 91.0 | PASS | first_session | first_session |
| 25 | 2026-02-18 | aafe35e | full | 89.0 | PASS | difficulty_climb | first_session |
| 26 | 2026-02-18 | 6767390 | quick | 92.0 | PASS | first_session | first_session |
| 27 | 2026-02-18 | 6767390 | full | 86.0 | PASS | difficulty_climb | first_session |
| 28 | 2026-02-18 | 169947a | full | 97.2 | PASS | hint_progression | difficulty_climb |
| 29 | 2026-02-17 | ab36d15 | quick | 94.0 | PASS | first_session | first_session |
| 30 | 2026-02-17 | 20a7fd7 | quick | 67.0 | FAIL | first_session | first_session |
| 31 | 2026-02-17 | 169947a | full | 85.3 | PASS | cke_unlock | hint_progression |

## Analiza per-scenariusz

| Scenariusz | Avg | Current | Trend | StdDev | Min | Max | Regressed |
|------------|-----|---------|-------|--------|-----|-----|-----------|
| first_session | 94.1 | 95.0 | → | 6.7 | 76.0 | 100.0 |  |
| hint_progression | 88.8 | 78.6 | ↓ | 5.2 | 78.6 | 95.0 | ⚠ |
| difficulty_climb | 90.3 | 92.0 | ↑ | 4.4 | 85.0 | 100.0 |  |
| review_session | 94.1 | 95.0 | → | 3.7 | 89.0 | 100.0 |  |
| coaching_aware | 92.0 | 95.0 | ↑ | 7.1 | 73.0 | 100.0 |  |
| cke_unlock | 90.5 | 89.0 | ↓ | 5.6 | 81.0 | 100.0 |  |
| probna | 88.3 | 100.0 | ↑ | 8.4 | 77.0 | 100.0 |  |

## Trendy kryteriow L2

| Kryterium | Avg | Current | Trend |
|-----------|-----|---------|-------|
| socratic | 4.1 | 4.1 | → |
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

## Interpretacja

### Co sie zmienilo?

Wynik wzrosl z 87.7 (c339ce7) do **92.1** (4362947), delta **+4.4 pkt**. Glowne zrodla poprawy:

- **coaching_aware**: 73.0 → 95.0 (+22.0) — commit 4362947 naprawil retry logic dla `progress blad` (CLI odrzucenie kodow) i powiazanie bledu z leech tagiem. Wszystkie 10 checkpointow PASS.
- **probna**: 77.0 → 100.0 (+23.0) — naprawiono wzorce wizualizacji i linkowanie leech tagow. Pelny wynik: L1 11/11, L2 socratic 5/5, tone 5/5.
- **cke_unlock**: 96.0 → 89.0 (-7.0) — regresja wynika z surowszego ewaluatora: `typ intro` nie wywolany (nowy checkpoint FAIL) i `krok-po-kroku` nie wymuszony (PARTIAL). To nowe ustalenia ewaluatora, nie regresja kodu.

### Top 3 do naprawienia

1. **hint_progression: redizajn scenariusza testowego** (78.6, ↓ trend, ⚠ regressed) — 3 checkpointy (HINT_LOCKED, L2→L3, cheatsheet_excerpt) sa strukturalnie nieosiagalne: hint_delay=1 uniemozliwia HINT_LOCKED, a 3 bledy triggeruja walk_through zanim L2/L3 moga nastapic. Potrzebna zmiana: hint_delay=2 w scenariuszu + 4-5 wymian zamiast 3.

2. **cke_unlock: `typ intro` przed cwiczeniem** (FAIL) — SKILL.md wymaga `typ intro --typ X` przy first_in_type aby sprawdzic czy uczen widzial juz dany typ. Tutor tego nie wywoluje. Mozliwe rozwiazanie: dodac przypomnienie w sekcji D SKILL.md.

3. **Pytania sokratejskie lekko naprowadzajace** (socratic avg 4.1, stabilne) — Powtarza sie w wielu scenariuszach: tutor ujawnia czesciowe odpowiedzi w pytaniach (np. "reszta z dzielenia" zdradza mod, "1234 mod 10 = 4" zamiast kazac uczniowi obliczyc). Mozliwe rozwiazanie: dodac anty-wzorce do SKILL.md sekcja F.

### Uporczywe problemy (3+ kolejne uruchomienia)

- **hint_progression structural FAILs** — obecne w 5/5 ostatnich uruchomien. Scenariusz wymaga redizajnu (patrz Top 1).
- **Worked example jako demonstracja** — pojawia sie sporadycznie od 02-22. Tutor prezentuje worked example zamiast rozwiazywac go wspolnie z uczniem.
- **Pytanie sokratejskie lekko podajace** — stabilne 4/5 od poczatku. Sufit bez zmian w SKILL.md.
- **feedback_czasowy parafrazowany** — difficulty_climb od 02-24. Tutor nie wyswietla komunikatu doslownie.

### Co dziala dobrze

- **first_session** (95.0, → stabilne, avg 94.1) — najstabilniejszy scenariusz, pelne L1 w wielu uruchomieniach.
- **review_session** (95.0, → stabilne, avg 94.1, StdDev 3.7) — najnizszy szum ewaluatora, konsekwentnie wysoki wynik.
- **probna** (100.0, ↑ trend) — po naprawie w 4362947 osiagnela idealne 100/100. Tryb egzaminowy dziala bezblednie.
- **coaching_aware** (95.0, ↑ trend) — duza poprawa po naprawie retry logic i leech linkowania.
- **Ton i jezyk** (avg 4.8, current 4.9) — stabilnie wysoki, polski, "ty", bez emoji, zachecajacy.
- **29 kolejnych PASS** — system jest stabilny i niezawodny.
