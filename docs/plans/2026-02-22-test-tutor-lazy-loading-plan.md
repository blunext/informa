# test-tutor Lazy Loading Update — Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Update test-tutor SKILL.md to reflect lazy loading exercises (question/hints/answer) and coaching field.

**Architecture:** Single file edit — `.claude/skills/test-tutor/SKILL.md`. Six localized changes: pre-fetch, persona assignment, rubric, hint_progression scenario, new coaching_aware scenario, agent prompt + JSON format.

**Tech Stack:** Markdown (SKILL.md), Bash (pre-fetch snippets)

---

### Task 1: Update pre-fetch section (sekcja 2)

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md:48-52` (add hints+answer fetches after exercise question)

**Step 1: Add hints + answer pre-fetch after existing exercise question lines**

After line 52 (`EX_ARKUSZ=...`), insert:

```bash
# Exercise IDs (for hints/answer fetch)
EX_TEORIA_ID=$(echo "$EX_TEORIA" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])")
EX_IMPL_ID=$(echo "$EX_IMPL" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])")

# Hinty i odpowiedzi (lazy loading — osobne wywolania)
HINTS_TEORIA=$($MATURA --db-dir /tmp/test-tutor-$$ exercise hints --id $EX_TEORIA_ID)
HINTS_IMPL=$($MATURA --db-dir /tmp/test-tutor-$$ exercise hints --id $EX_IMPL_ID)
ANSWER_TEORIA=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $EX_TEORIA_ID)
ANSWER_IMPL=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $EX_IMPL_ID)
```

**Step 2: Add coaching_aware pre-fetch before the `# Raport metadata` block (before line 63)**

Insert:

```bash
# Coaching_aware: zasymuluj progressed studenta
sqlite3 /tmp/test-tutor-$$/matura_progress.db "
INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'srednie', 4);
INSERT INTO progress_tagi (tag, lapses, stability, last_review) VALUES ('cyfry-mod-div', 4, 1.0, '$(date -v-30d +%Y-%m-%d)');
INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, data) VALUES ('7.1', 'cyfry_liczby', 'mylenie_div_mod', 'Pomylenie div z mod', '$(date +%Y-%m-%d)');
INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.1','cyfry_liczby','$(date +%Y-%m-%d)','poprawne_z_pomoca_1');
"
EX_COACHING=$($MATURA --db-dir /tmp/test-tutor-$$ exercise question --typ cyfry_liczby)
EX_COACHING_ID=$(echo "$EX_COACHING" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])")
HINTS_COACHING=$($MATURA --db-dir /tmp/test-tutor-$$ exercise hints --id $EX_COACHING_ID)
ANSWER_COACHING=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $EX_COACHING_ID)
```

**Step 3: Verify — read file, confirm pre-fetch has all 3 blocks (question, hints+answer, coaching_aware)**

---

### Task 2: Update persona scenario assignments (sekcja 1)

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md:29-31`

**Step 1: Change intermediate persona default scenarios**

Replace line 30:
```
- **intermediate**: difficulty_climb, review_session
```
With:
```
- **intermediate**: difficulty_climb, review_session, coaching_aware
```

---

### Task 3: Update rubric (sekcja 4)

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md:96-117`

**Step 1: Change section header from "7 kryteriow" to "8 kryteriow"**

Replace line 96:
```
## 4. Rubryka oceny (7 kryteriow)
```
With:
```
## 4. Rubryka oceny (8 kryteriow)
```

**Step 2: Update criterion #2 description (line 103)**

Replace:
```
| 2 | Progresja hintow | 20% | L1→L2→L3 w scislej kolejnosci, cheatsheet przy L2, konsolidacja po walk_through | Kolejnosc zachowana ale brak cheatsheet | Brak progresji, hinty pominięte lub w zlej kolejnosci |
```
With:
```
| 2 | Progresja hintow | 20% | Hinty lazy (`exercise hints --id`), odpowiedz lazy (`exercise answer --id`), hint_delay respektowany, L1→L2→L3, cheatsheet przy L2, konsolidacja po walk_through | Lazy loading obecne ale hint_delay ignorowany, lub kolejnosc niedokladna | Hinty/odpowiedz podane z gory lub brak progresji |
```

**Step 3: Update criterion #6 weight (line 107)**

Replace:
```
| 6 | Ton i jezyk | 10% | Polski, "ty", bez emoji, zachecanie, cierpliwosc, feedback czasowy | Poprawny jezyk ale bez zachecania | Angielski, formalny, emoji, lub brak feedbacku |
```
With:
```
| 6 | Ton i jezyk | 5% | Polski, "ty", bez emoji, zachecanie, cierpliwosc, feedback czasowy | Poprawny jezyk ale bez zachecania | Angielski, formalny, emoji, lub brak feedbacku |
```

**Step 4: Add criterion #8 after line 108 (after criterion #7)**

Insert:
```
| 8 | Coaching | 5% | Tutor reaguje na coaching.leech_tags (ostrzega o slabych tagach), coaching.past_mistakes (proaktywnie wspomina wczesniejsze bledy), coaching.hint_delay (respektuje opoznienie) | Coaching czesciowo wykorzystany — np. hint_delay ok ale leech_tags ignorowane | Coaching calkowicie ignorowany |
```

**Step 5: Update N/A rule (lines 113-117)**

Replace:
```
**Kryteria N/A**: Jesli scenariusz nie testuje danego kryterium (np. Powtorki SR
w first_session/hint_progression/cke_unlock/probna), ocen na **4/5** z uwaga
"Minimalne wymagania spelnione (progress status sprawdzony), scenariusz nie testuje
tego kryterium bezposrednio." Nie dawaj 5/5 (brak dowodu) ani 3/5 (brak naruszenia).
Wyjatek: jesli tutor NIE sprawdzil progress status na starcie → 2/5.
```
With:
```
**Kryteria N/A**: Jesli scenariusz nie testuje danego kryterium (np. Powtorki SR
w first_session/hint_progression/cke_unlock/probna, lub Coaching w first_session/probna
gdzie uczen jest fresh), ocen na **4/5** z uwaga
"Minimalne wymagania spelnione (progress status sprawdzony), scenariusz nie testuje
tego kryterium bezposrednio." Nie dawaj 5/5 (brak dowodu) ani 3/5 (brak naruszenia).
Wyjatek: jesli tutor NIE sprawdzil progress status na starcie → 2/5.
```

---

### Task 4: Update hint_progression scenario (sekcja 5.2)

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md:128-133`

**Step 1: Replace hint_progression scenario**

Replace lines 128-133:
```
### 5.2 hint_progression
- Uczen pracuje nad cwiczeniami, trafia na trudne
- **Przebieg**: uczen 3x odpowiada blednie na to samo cwiczenie
- **Oczekiwania**: L1 (pytanie sokratejskie) → L2 (cheatsheet section + wskazowka) → L3 (kluczowy krok) → walk_through
- **Po walk_through**: konsolidacja ("wyjasniej swoimi slowami")
- **Kluczowe**: tutor rejestruje blad (`progress blad`) na kazdym etapie
```
With:
```
### 5.2 hint_progression
- Uczen pracuje nad cwiczeniami, trafia na trudne
- **Przebieg**: uczen 3x odpowiada blednie na to samo cwiczenie
- **Oczekiwania (lazy loading)**:
  - Tutor prezentuje cwiczenie z `exercise question` (bez hintow/odpowiedzi)
  - Po 1. bledzie: tutor sprawdza `coaching.hint_delay`:
    * hint_delay=1 → pobiera `exercise hints --id`, podaje L1
    * hint_delay=2 → tylko pytanie sokratejskie (bez hints)
    * hint_delay=3 → tylko pytanie sokratejskie (bez hints)
  - Po kolejnych bledach: tutor pobiera hinty jesli jeszcze nie pobral, progresja L1→L2→L3
  - Po walk_through: pobiera `exercise answer --id`, wyswietla odpowiedz, konsolidacja
- **Kluczowe**: `progress blad` na kazdym etapie, lazy loading respektowany
```

---

### Task 5: Add coaching_aware scenario (sekcja 5.7)

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md` (after line 156, after scenario 5.6)

**Step 1: Insert new scenario 5.7 after probna (5.6)**

After line 156, insert:
```

### 5.7 coaching_aware
- Uczen z historia — ma leech_tags i past_mistakes w coaching
- **Setup**: pre-fetch z progressed DB (uczen ma cwiczenia, 3+ lapses na tagu, bledy w sesjach)
- **Przebieg**:
  1. Tutor pobiera `exercise question` — coaching zawiera leech_tags i past_mistakes
  2. Uczen rozwiazuje cwiczenie z tagiem obecnym w leech_tags
  3. Tutor powinien proaktywnie ostrzec o slabym tagu
  4. Uczen popelnia blad z kodem obecnym w past_mistakes
  5. Tutor powinien powiazac blad z historia ("Ostatnio miales problem z X")
- **Oczekiwania**:
  - Tutor czyta coaching.leech_tags i reaguje (ostrzezenie, dodatkowa uwaga)
  - Tutor czyta coaching.past_mistakes i proaktywnie wspomina
  - hint_delay respektowany (progressed student = familiar/mastered = hint_delay 2-3)
- **Kluczowe**: coaching nie moze byc ignorowany — to glowny cel tego scenariusza
```

---

### Task 6: Update agent prompt and JSON format (sekcja 6)

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md:181-222`

**Step 1: Update pre-fetched data block (lines 181-185)**

Replace:
```
## Pre-fetched data
- Exercise (TEORIA): {EX_TEORIA}
- Exercise (IMPL): {EX_IMPL}
- Typ intro: {INTRO_TEORIA}
- Progress status: {STATUS}
```
With:
```
## Pre-fetched data
- Question (TEORIA): {EX_TEORIA}
- Hints (TEORIA): {HINTS_TEORIA}
- Answer (TEORIA): {ANSWER_TEORIA}
- Question (IMPL): {EX_IMPL}
- Hints (IMPL): {HINTS_IMPL}
- Answer (IMPL): {ANSWER_IMPL}
- Question (COACHING): {EX_COACHING} (only for coaching_aware scenario)
- Hints (COACHING): {HINTS_COACHING}
- Answer (COACHING): {ANSWER_COACHING}
- Typ intro: {INTRO_TEORIA}
- Progress status: {STATUS}
```

**Step 2: Update instructions (lines 187-193)**

Replace:
```
## Instrukcje
1. Symuluj dialog tutor↔uczen (8-15 wymian). Tutor postepuje wg SKILL.md,
   uczen wg persony (accuracy, typowe bledy, zachowanie).
2. Przy kazdej akcji tutora ZAPISZ komende CLI ktora tutor POWINIEN wywolac
   (np. `./matura exercise question --typ X`, `./matura progress update --id Y --wynik Z`).
3. Po symulacji OCEN transkrypt wg ponizszej rubryki.
4. Zwroc wynik DOKLADNIE w formacie JSON ponizej.
```
With:
```
## Instrukcje
1. Symuluj dialog tutor↔uczen (8-15 wymian). Tutor postepuje wg SKILL.md,
   uczen wg persony (accuracy, typowe bledy, zachowanie).
2. Przy kazdej akcji tutora ZAPISZ komende CLI ktora tutor POWINIEN wywolac
   (np. `./matura exercise question --typ X`, `./matura exercise hints --id Y`,
   `./matura exercise answer --id Y`, `./matura progress update --id Y --wynik Z`).
3. WAZNE — lazy loading: tutor NIE widzi hintow ani odpowiedzi na starcie.
   Musi pobrac je osobnymi komendami. Jesli tutor podaje hint bez wczesniejszego
   `exercise hints --id` — to blad integralnosci.
4. Po symulacji OCEN transkrypt wg ponizszej rubryki.
5. Zwroc wynik DOKLADNIE w formacie JSON ponizej.
```

**Step 3: Add coaching to JSON scores (line 210)**

After line 210 (`"integralnosc_cli": ...`), insert:
```
    "coaching": {"score": N, "max": 5, "uwagi": "..."}
```

**Step 4: Update weight formula (lines 220-222)**

Replace:
```
Score oblicz: sum(score * waga) / 5 * 100, gdzie wagi:
metoda_sokratejska=0.25, progresja_hintow=0.20, sledzenie_bledow=0.15,
adaptacja_trudnosci=0.15, powtorki_sr=0.10, ton_i_jezyk=0.10, integralnosc_cli=0.05.
```
With:
```
Score oblicz: sum(score * waga) / 5 * 100, gdzie wagi:
metoda_sokratejska=0.25, progresja_hintow=0.20, sledzenie_bledow=0.15,
adaptacja_trudnosci=0.15, powtorki_sr=0.10, ton_i_jezyk=0.05, integralnosc_cli=0.05, coaching=0.05.
```

---

### Task 7: Update report format (sekcja 8)

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md:250-256`

**Step 1: Add coaching row to report table template**

After line 256 (`| Integralnosc CLI | ...`), insert:
```
| Coaching | {s}/5 | {uwagi} |
```

---

### Task 8: Verify + commit

**Step 1: Run test_qa.sh layer 3 (SKILL lint)**

Run: `matura_informatyka_rozszerzona/analiza/test_qa.sh --layer 3`
Expected: all PASS (CLI commands in updated SKILL.md are valid)

**Step 2: Grep for stale `exercise get` references**

Run: `grep -r "exercise get" .claude/skills/`
Expected: no matches

**Step 3: Commit**

```bash
git add .claude/skills/test-tutor/SKILL.md
git commit -m "Update test-tutor for lazy loading: coaching criterion, hint_progression, coaching_aware scenario"
```
