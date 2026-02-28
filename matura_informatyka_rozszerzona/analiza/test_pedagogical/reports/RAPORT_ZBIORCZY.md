# Raport Zbiorczy — Test Tutor

## Dashboard

- **Wynik**: 87.7  PASS
- **Delta**: -7.9
- **Seria PASS**: 28
- **Liczba uruchomien**: 30
- **Okno (10)**: avg=90.5  min=85.1  max=95.6

## Historia

| # | Data | Commit | Mode | Wynik | Pass | Best | Worst |
|---|------|--------|------|-------|------|------|-------|
| 1 | 2026-02-28 | c339ce7 | full | 87.7 | PASS | cke_unlock | coaching_aware |
| 2 | 2026-02-25 | a4c1c8d | full | 95.6 | PASS | first_session | hint_progression |
| 3 | 2026-02-24 | 0dbf358 | full | 95.3 | PASS | first_session | hint_progression |
| 4 | 2026-02-24 | 0dbf358 | full | 92.3 | PASS | first_session | hint_progression |
| 5 | 2026-02-24 | fce6714 | full | 92.4 | PASS | coaching_aware | difficulty_climb |
| 6 | 2026-02-24 | 07cb9f3 | full | 89.5 | PASS | first_session | probna |
| 7 | 2026-02-24 | 6baa479 | full | 92.3 | PASS | first_session | difficulty_climb |
| 8 | 2026-02-23 | c2e9ce9 | full | 86.6 | PASS | coaching_aware | first_session |
| 9 | 2026-02-23 | c2e9ce9 | full | 87.7 | PASS | hint_progression | probna |
| 10 | 2026-02-22 | dcf2deb | full | 85.1 | PASS | difficulty_climb | review_session |
| 11 | 2026-02-22 | ac622d1 | full | 90.7 | PASS | hint_progression | cke_unlock |
| 12 | 2026-02-22 | 6c9ecc3 | full | 88.4 | PASS | difficulty_climb | coaching_aware |
| 13 | 2026-02-22 | 367575d | full | 89.6 | PASS | coaching_aware | probna |
| 14 | 2026-02-22 | 22c3ebe | full | 84.1 | PASS | hint_progression | first_session |
| 15 | 2026-02-22 | 1314828 | full | 91.2 | PASS | hint_progression | probna |
| 16 | 2026-02-19 | dc0714c | full | 92.0 | PASS | difficulty_climb | first_session |
| 17 | 2026-02-19 | a68bf9a | full | 90.0 | PASS | cke_unlock | first_session |
| 18 | 2026-02-19 | a68bf9a | full | 89.0 | PASS | probna | first_session |
| 19 | 2026-02-19 | 187988f | full | 90.0 | PASS | probna | first_session |
| 20 | 2026-02-19 | 042ebe8 | full | 90.0 | PASS | cke_unlock | first_session |
| 21 | 2026-02-18 | e35c154 | quick | 90.0 | PASS | first_session | first_session |
| 22 | 2026-02-18 | e35c154 | full | 86.3 | PASS | review_session | first_session |
| 23 | 2026-02-18 | ac6751d | beginner | 91.0 | PASS | first_session | first_session |
| 24 | 2026-02-18 | aafe35e | full | 89.0 | PASS | difficulty_climb | first_session |
| 25 | 2026-02-18 | 6767390 | quick | 92.0 | PASS | first_session | first_session |
| 26 | 2026-02-18 | 6767390 | full | 86.0 | PASS | difficulty_climb | first_session |
| 27 | 2026-02-18 | 169947a | full | 97.2 | PASS | hint_progression | difficulty_climb |
| 28 | 2026-02-17 | ab36d15 | quick | 94.0 | PASS | first_session | first_session |
| 29 | 2026-02-17 | 20a7fd7 | quick | 67.0 | FAIL | first_session | first_session |
| 30 | 2026-02-17 | 169947a | full | 85.3 | PASS | cke_unlock | hint_progression |

## Analiza per-scenariusz

| Scenariusz | Avg | Current | Trend | StdDev | Min | Max | Regressed |
|------------|-----|---------|-------|--------|-----|-----|-----------|
| first_session | 93.0 | 95.0 | ↑ | 7.4 | 76.0 | 100.0 |  |
| hint_progression | 89.9 | 95.0 | ↑ | 4.0 | 84.1 | 95.0 |  |
| difficulty_climb | 90.5 | 87.0 | ↓ | 4.5 | 85.0 | 100.0 |  |
| review_session | 92.2 | 91.0 | ↓ | 6.5 | 76.0 | 100.0 |  |
| coaching_aware | 91.0 | 73.0 | ↓ | 7.3 | 73.0 | 100.0 | ⚠ |
| cke_unlock | 90.1 | 96.0 | ↑ | 5.9 | 81.0 | 100.0 |  |
| probna | 86.5 | 77.0 | ↓ | 7.5 | 77.0 | 95.0 | ⚠ |

## Trendy kryteriow L2

| Kryterium | Avg | Current | Trend |
|-----------|-----|---------|-------|
| socratic | 4.1 | 3.4 | → |
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

Wynik spadl o **-7.9 pkt** (95.6 → 87.7) wzgledem poprzedniego uruchomienia (2026-02-25). Nie zmienialo sie nic w kodzie SKILL.md (ten sam commit c339ce7 co ostatnie zmiany w repo). Spadek wynika z **bardziej surowej oceny** dwoch scenariuszy:

- **coaching_aware**: 95.0 → 73.0 (-22 pkt) — nowy ewaluator dodal dwa dodatkowe checkpointy: wizualizacje proaktywna po bledzie (F.6) i powiazanie aktualnego bledu z leech tagiem po bledzie. Oba FAIL. To nie jest regresja SKILL.md — to **ostrzejsze wymagania testowe**.
- **probna**: 95.0 → 77.0 (-18 pkt) — ewaluator wykryl ze `progress blad` uzywa niepoprawnych kodow bledow (np. `off_by_one` dla `analiza_algorytmu`) odrzucanych przez CLI. Poprzednie testy nie weryfikowaly poprawnosci kodow. To **realne odkrycie** — probna.md powinno explicite wspominac o retry z suggestions[].

Scenariusze first_session (+) i hint_progression (+) lekko wzrosly. cke_unlock stabilny na 96.

### Top 3 do naprawienia

1. **probna.md: dodac instrukcje retry kodow bledow** — Tutor uzywa niepoprawnych kodow w `progress blad` (np. `off_by_one`), CLI je odrzuca, bledy nie sa rejestrowane. probna.md step 6 powinno zawierac: "Jesli CLI odrzuci kod — uzyj sugestii (patrz SKILL.md F.3)."

2. **SKILL.md coaching_aware: wymuszenie powiazania bledu z leech tagiem** — Gdy tutor dostaje WARN_LEECH na starcie a potem uczen popelnia dokladnie ten sam blad, SKILL.md powinno explicite mowic: "Po bledzie zwiazanym z ostrzeganym tagiem — powiaz ('To ten sam problem co wczesniej z X')."

3. **SKILL.md F.6: wizualizacja dla IMPLEMENTACJA** — Regula wymaga wizualizacji po bledzie, ale lista wzorcow (sledzenie, projektowanie, analiza, konwersja, bezpieczenstwo) nie obejmuje IMPLEMENTACJA. Dodac: "Dla IMPLEMENTACJA: tabelka mod vs div, schemat wczytywania, diagram algorytmu."

### Uporczywe problemy

- **Socratic question + hint w jednej wiadomosci** — pojawia sie od 2026-02-18 w roznych scenariuszach (hint_progression, review_session). HARD GATE w sekcji F.4 jest jasny, ale ewaluatorzy oceniaja to niejednolicie. Sugestia: dodac test-case specyficzny na HARD GATE.
- **Worked example ID reuse** — pojawilo sie w difficulty_climb (duplikat ID 1.15). CLI `exercise next` powinno wykluczac `przyklad` ID z `typ intro`.
- **probna scoring** — ewaluator jest niestabilny (77.0-95.0 w roznych uruchomieniach). StdDev=7.5 — najwyzszy z wszystkich scenariuszy.

### Co dziala dobrze

- **first_session**: 95.0, trend ↑ — stabilny ponad srednia (avg=93.0). CLI compliance 100%.
- **hint_progression**: 95.0, trend ↑ — najlepszy wynik w historii. Progresja hintow, konsolidacja i wizualizacja ASCII dzialaja poprawnie.
- **cke_unlock**: 96.0, trend ↑ — najlepszy scenariusz w tym uruchomieniu. Caly flow unlock → worked-example → sprawdzian dziala bezblednie.
- **review_session**: 91.0, stabilny — `exercise review` priorytetyzowane, coaching_actions dostarczane.
- **Tone**: avg=4.7/5, stabilny — jezyk, ton i forma konsekwentnie poprawne.
