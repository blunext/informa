# Design: test-tutor SKILL.md — aktualizacja dla lazy loading

Data: 2026-02-22

## Problem

Po wdrożeniu lazy loading exercises (exercise get → exercise question/hints/answer + coaching),
test-tutor SKILL.md nie odzwierciedla zmian. Agenci symulujący sesje nie mają kompletnych danych,
scenariusze nie testują lazy loading ani coaching, rubryka nie weryfikuje nowych zachowań.

## Decyzje

1. **Pre-fetch**: pobierać hints + answer osobno (agenci potrzebują pełnych danych do symulacji)
2. **Rubryka**: nowe kryterium #8 "Coaching" (5%), waga z #6 "Ton i język" (10%→5%)
3. **Scenariusze**: zmodyfikować hint_progression + dodać nowy coaching_aware
4. **coaching_aware**: przypisany do intermediate (progressed student)

## Zmiany

### 1. Pre-fetch danych (sekcja 2)

Dodaj osobne wywołania hints + answer po exercise question:

```bash
EX_TEORIA_ID=$(echo "$EX_TEORIA" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])")
EX_IMPL_ID=$(echo "$EX_IMPL" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])")
HINTS_TEORIA=$($MATURA --db-dir /tmp/test-tutor-$$ exercise hints --id $EX_TEORIA_ID)
HINTS_IMPL=$($MATURA --db-dir /tmp/test-tutor-$$ exercise hints --id $EX_IMPL_ID)
ANSWER_TEORIA=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $EX_TEORIA_ID)
ANSWER_IMPL=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $EX_IMPL_ID)
```

Dla coaching_aware — zasymuluj progressed studenta:

```bash
sqlite3 /tmp/test-tutor-$$/matura_progress.db "
INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'srednie', 4);
INSERT INTO progress_tagi (tag, lapses, stability, last_review) VALUES ('cyfry-mod-div', 4, 1.0, '$(date -v-30d +%Y-%m-%d)');
INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, data) VALUES ('7.1', 'cyfry_liczby', 'mylenie_div_mod', 'Pomylenie div z mod', '$(date +%Y-%m-%d)');
INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.1','cyfry_liczby','$(date +%Y-%m-%d)','poprawne_z_pomoca_1');
"
EX_COACHING=$($MATURA --db-dir /tmp/test-tutor-$$ exercise question --typ cyfry_liczby)
HINTS_COACHING=$($MATURA --db-dir /tmp/test-tutor-$$ exercise hints --id $(echo "$EX_COACHING" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])"))
ANSWER_COACHING=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $(echo "$EX_COACHING" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])"))
```

### 2. Prompt agenta (sekcja 6)

Pre-fetched data:
```
- Question (TEORIA): {EX_TEORIA}
- Hints (TEORIA): {HINTS_TEORIA}
- Answer (TEORIA): {ANSWER_TEORIA}
- Question (IMPL): {EX_IMPL}
- Hints (IMPL): {HINTS_IMPL}
- Answer (IMPL): {ANSWER_IMPL}
- Typ intro: {INTRO_TEORIA}
- Progress status: {STATUS}
```

Instrukcje (dodać punkt o lazy loading):
```
3. WAZNE — lazy loading: tutor NIE widzi hintow ani odpowiedzi na starcie.
   Musi pobrac je osobnymi komendami. Jesli tutor podaje hint bez wczesniejszego
   `exercise hints --id` — to blad integralnosci.
```

### 3. Rubryka — zmiana wag + kryterium #8

| # | Kryterium | Waga |
|---|-----------|------|
| 1 | Metoda sokratejska | 25% |
| 2 | Progresja hintów | 20% |
| 3 | Śledzenie błędów | 15% |
| 4 | Adaptacja trudności | 15% |
| 5 | Powtórki SR | 10% |
| 6 | Ton i język | **5%** |
| 7 | Integralność CLI | 5% |
| **8** | **Coaching** | **5%** |

Kryterium #2 — zaktualizowany opis:
- 5/5: hinty lazy (`exercise hints --id`), odpowiedź lazy (`exercise answer --id`), hint_delay respektowany, L1→L2→L3, cheatsheet przy L2, konsolidacja po walk_through
- 3/5: lazy loading obecne ale hint_delay ignorowany, lub kolejność niedokładna
- 1/5: hinty/odpowiedź podane z góry lub brak progresji

Kryterium #8 — Coaching:
- 5/5: tutor reaguje na leech_tags (ostrzega), past_mistakes (proaktywnie wspomina), hint_delay (respektuje)
- 3/5: coaching częściowo wykorzystany
- 1/5: coaching ignorowany
- N/A: fresh student → 4/5 "fresh student, coaching minimalne"

### 4. Scenariusz hint_progression — aktualizacja

```
### 5.2 hint_progression
- Uczen pracuje nad cwiczeniami, trafia na trudne
- Przebieg: uczen 3x odpowiada blednie na to samo cwiczenie
- Oczekiwania (lazy loading):
  - Tutor prezentuje cwiczenie z `exercise question` (bez hintow/odpowiedzi)
  - Po 1. bledzie: tutor sprawdza coaching.hint_delay:
    * hint_delay=1 → pobiera `exercise hints --id`, podaje L1
    * hint_delay=2 → tylko pytanie sokratejskie (bez hints)
    * hint_delay=3 → tylko pytanie sokratejskie (bez hints)
  - Po kolejnych bledach: pobiera hinty jesli nie pobral, progresja L1→L2→L3
  - Po walk_through: pobiera `exercise answer --id`, wyswietla odpowiedz, konsolidacja
- Kluczowe: progress blad na kazdym etapie, lazy loading respektowany
```

### 5. Nowy scenariusz coaching_aware

```
### 5.7 coaching_aware
- Uczen z historia — ma leech_tags i past_mistakes w coaching
- Setup: pre-fetch z progressed DB (uczen ma 10+ cwiczen, 3+ lapses na tagu)
- Przebieg:
  1. Tutor pobiera exercise question — coaching zawiera leech_tags i past_mistakes
  2. Uczen rozwiazuje cwiczenie z tagiem obecnym w leech_tags
  3. Tutor powinien proaktywnie ostrzec o slabym tagu
  4. Uczen popelnia blad z kodem obecnym w past_mistakes
  5. Tutor powinien powiazac blad z historia
- Oczekiwania:
  - Tutor czyta coaching.leech_tags i reaguje
  - Tutor czyta coaching.past_mistakes i proaktywnie wspomina
  - hint_delay respektowany (familiar/mastered = 2-3)
- Kluczowe: coaching nie moze byc ignorowany
```

Przypisanie: intermediate (difficulty_climb, review_session, coaching_aware).

### 6. Format JSON agenta

Dodaj `coaching` do scores + zaktualizuj wagi:
```
coaching=0.05, ton_i_jezyk=0.05 (bylo 0.10)
```
