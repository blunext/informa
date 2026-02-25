---
name: test-tutor
description: >
  Testuje jakosc pedagogiczna skilla korepetytora (matura/SKILL.md).
  Spawns 3 agentow (beginner/intermediate/advanced), kazdy symuluje sesje,
  ocenia wg rubryki, zwraca raport.
argument-hint: "[quick | beginner | intermediate | advanced | --scenario SCENARIO]"
---

# Test Tutor — Multi-Agent Pedagogical Testing

Skill testujacy jakosc pedagogiczna korepetytora (`matura/SKILL.md`).
Po wywolaniu spawns agentow-symulantow, kazdy symuluje sesje ucznia,
ocenia transkrypt wg rubryki, zwraca raport JSON.

## 1. Parsowanie argumentow

```
/test-tutor                                    # pelny: 3 persony x 2 scenariusze = 6 uruchomien
/test-tutor quick                              # szybki: beginner x first_session = 1 uruchomienie
/test-tutor beginner                           # 1 persona, 2 domyslne scenariusze
/test-tutor intermediate                       # 1 persona, 2 domyslne scenariusze
/test-tutor advanced                           # 1 persona, 2 domyslne scenariusze
/test-tutor --scenario hint_progression        # 1 scenariusz, 3 persony
/test-tutor beginner --scenario first_session  # 1 persona, 1 scenariusz
```

Domyslne scenariusze per persona:
- **beginner**: first_session, hint_progression
- **intermediate**: difficulty_climb, review_session, coaching_aware
- **advanced**: cke_unlock, probna

## 2. Pre-fetch danych

Na poczatku pobierz dane potrzebne agentom (unikaj powtarzania CLI calls w agentach):

```bash
CLI_DIR="matura_informatyka_rozszerzona/analiza/cli"
MATURA="$CLI_DIR/matura"

# Testowany skill
SKILL_CONTENT=$(cat .claude/skills/matura/SKILL.md)

# Utworz temp DB (izolacja od user progress)
mkdir -p /tmp/test-tutor-$$
cp "$CLI_DIR/matura.db" /tmp/test-tutor-$$/matura.db

# Przykladowe cwiczenia (po 1 na typ)
EX_TEORIA=$($MATURA --db-dir /tmp/test-tutor-$$ exercise question --typ sledzenie_algorytmu --trudnosc latwe)
EX_IMPL=$($MATURA --db-dir /tmp/test-tutor-$$ exercise question --typ cyfry_liczby --trudnosc latwe)
EX_SQL=$($MATURA --db-dir /tmp/test-tutor-$$ exercise question --typ sql_group_by --trudnosc latwe)
EX_ARKUSZ=$($MATURA --db-dir /tmp/test-tutor-$$ exercise question --typ agregacja_warunkowa --trudnosc latwe)

# Exercise IDs (for hints/answer fetch)
EX_TEORIA_ID=$(printf '%s' "$EX_TEORIA" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])")
EX_IMPL_ID=$(printf '%s' "$EX_IMPL" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])")

# Unblock guardrails for pre-fetch (exercise question registers with attempt_count=0)
sqlite3 /tmp/test-tutor-$$/matura_progress.db "UPDATE active_exercises SET attempt_count = 99;"

# Hinty i odpowiedzi (lazy loading — osobne wywolania)
HINTS_TEORIA=$($MATURA --db-dir /tmp/test-tutor-$$ exercise hints --id $EX_TEORIA_ID)
HINTS_IMPL=$($MATURA --db-dir /tmp/test-tutor-$$ exercise hints --id $EX_IMPL_ID)
ANSWER_TEORIA=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $EX_TEORIA_ID)
ANSWER_IMPL=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $EX_IMPL_ID)

# Typ intro
INTRO_TEORIA=$($MATURA --db-dir /tmp/test-tutor-$$ typ intro --typ sledzenie_algorytmu)

# Progress status (fresh)
STATUS=$($MATURA --db-dir /tmp/test-tutor-$$ progress status)

# Cheatsheet excerpt
CHEAT_TEORIA=$($MATURA --db-dir /tmp/test-tutor-$$ cheatsheet get --kategoria TEORIA --sekcja "archetyp")

# Coaching_aware: zasymuluj progressed studenta
sqlite3 /tmp/test-tutor-$$/matura_progress.db "
INSERT INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('cyfry_liczby', 'srednie', 4);
INSERT INTO progress_tagi (tag, lapses, stability, last_review) VALUES ('cyfry-mod-div', 4, 1.0, '$(date -v-30d +%Y-%m-%d)');
INSERT INTO progress_bledy (exercise_id, typ, blad_kod, blad_opis, data) VALUES ('7.1', 'cyfry_liczby', 'mylenie_div_mod', 'Pomylenie div z mod', '$(date +%Y-%m-%d)');
INSERT INTO progress_zrobione (id, typ, data, wynik) VALUES ('7.1','cyfry_liczby','$(date +%Y-%m-%d)','poprawne_z_pomoca_1');
"
EX_COACHING=$($MATURA --db-dir /tmp/test-tutor-$$ exercise question --typ cyfry_liczby)
EX_COACHING_ID=$(printf '%s' "$EX_COACHING" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])")
sqlite3 /tmp/test-tutor-$$/matura_progress.db "UPDATE active_exercises SET attempt_count = 99 WHERE exercise_id = '$EX_COACHING_ID';"
HINTS_COACHING=$($MATURA --db-dir /tmp/test-tutor-$$ exercise hints --id $EX_COACHING_ID)
ANSWER_COACHING=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $EX_COACHING_ID)

# cke_unlock: zasymuluj studenta tuz przed progiem trudne (streak=7, srednie-trudne)
sqlite3 /tmp/test-tutor-$$/matura_progress.db "
INSERT OR REPLACE INTO progress_typy (typ, poziom_trudnosci, streak) VALUES ('sledzenie_algorytmu', 'srednie-trudne', 7);
"
EX_CKE_PRE=$($MATURA --db-dir /tmp/test-tutor-$$ exercise question --typ sledzenie_algorytmu --trudnosc srednie-trudne)
EX_CKE_PRE_ID=$(printf '%s' "$EX_CKE_PRE" | python3 -c "import sys,json; print(json.loads(sys.stdin.read())['id'])")
$MATURA --db-dir /tmp/test-tutor-$$ progress blad --exercise-id $EX_CKE_PRE_ID --typ sledzenie_algorytmu --kod zly_wynik --hint 0 > /dev/null 2>&1
ANSWER_CKE_PRE=$($MATURA --db-dir /tmp/test-tutor-$$ exercise answer --id $EX_CKE_PRE_ID)

# Raport metadata
REPORT_DATE=$(date +%Y-%m-%d)
COMMIT_HASH=$(git -C /Users/blt1wz/priv/informa rev-parse --short HEAD)
REPORT_DIR="matura_informatyka_rozszerzona/analiza/test_pedagogical/reports"
REPORT_FILE="${REPORT_DIR}/${REPORT_DATE}_${COMMIT_HASH}.md"
```

## 3. Persony

Persony opisuja profil ucznia. Odpowiedzi ucznia sa **skryptowane** (patrz sekcja 5 — fixed scripts), ale styl/trudnosc tych odpowiedzi wynika z persony.

### Beginner (Poczatkujacy)
- **Accuracy**: 50% — co drugie cwiczenie blednie
- **Typowe bledy**: `mylenie_div_mod`, `brak_group_by`, `zle_adresowanie`, `brak_tabeli_sledzenia`
- **Zachowanie**: pisze krotkie odpowiedzi, czesto "nie wiem", myli podstawowe pojecia
- **Sesja**: 3-5 cwiczen

### Intermediate (Sredniozaawansowany)
- **Accuracy**: 70% — wiekszosc poprawnie, ale z lukami
- **Typowe bledy**: `pominiecie_bazy_rekurencji`, `brak_having`, `off_by_one`, `mylenie_avg_worst`
- **Zachowanie**: podaje rozumowanie, ale pomija edge cases
- **Sesja**: 5-8 cwiczen

### Advanced (Zaawansowany)
- **Accuracy**: 90% — rzadko sie myli
- **Typowe bledy**: `mylenie_avg_worst`, subtelne edge cases, `brak_uzasadnienia_zlozonosc`
- **Zachowanie**: podaje kompletne rozwiazania, uzasadnia
- **Sesja**: 8-12 cwiczen

## 4. Rubryka oceny (2 warstwy, 5 kryteriow)

### Layer 1 — Binary checkpoints (60% wagi)

Per scenariusz: lista TAK/NIE pytan (patrz sekcja 5). Ewaluator sprawdza konkretne fakty w transkrypcie — czy CLI command zostala wywolana, czy odpowiednia akcja nastapila. Nie ocenia holistycznie.

Score L1 = (trafienia / total_checkpoints) * 100

### Layer 2 — Holistic quality (40% wagi)

Ocena 1-5 dla 2 kryteriow jakosciowych z anchor examples:

**Metoda sokratejska (waga 25%)**:
- 5/5: Tutor pyta "Co sie stanie gdy n=0?" zamiast podac odpowiedz. Uczen sam dochodzi do rozwiazania. Pytania naprowadzajace po kazdym bledzie.
- 3/5: Tutor mowi "Podpowiedz: sprawdz warunek bazowy" — daje kierunek ale nie pyta.
- 1/5: Tutor mowi "Odpowiedz to 13, bo..." — gotowe rozwiazanie.

**Ton i jezyk (waga 15%)**:
- 5/5: Polski, "ty", zachecajacy, cierpliwy przez 3 proby, brak emoji. Feedback czasowy. Konsolidacja po walk_through.
- 3/5: Poprawny jezyk ale bez zachecania, lub sporadyczny emoji.
- 1/5: Angielski, formalny, emoji, lub agresywny ton.

Score L2 = (metoda_sokratejska * 0.25 + ton_i_jezyk * 0.15) / (0.25 + 0.15) — sredniawazona L2 kryteriow, znormalizowana do skali 1-5.

### Laczny score

```
overall = 0.6 * L1_percent + 0.4 * (L2_avg / 5 * 100)
```

**Prog zdania**: >= 70/100.

### 5 kryteriow (mapowanie na warstwy)

| # | Kryterium | Warstwa | Waga | Opis |
|---|-----------|---------|------|------|
| 1 | CLI compliance | L1 binary | 35% | Checkpoints CLI: exercise next (nie question), lazy loading, progress blad z --kod i --hint, progress update z --wynik i --czas, exercise hints/answer w odpowiednim momencie |
| 2 | Metoda sokratejska | L2 holistic | 25% | Pytania naprowadzajace, uczen probuje pierwszy, brak gotowych odpowiedzi |
| 3 | Ton i jezyk | L2 holistic | 15% | Polski, "ty", bez emoji, zachecanie, cierpliwosc, feedback czasowy, konsolidacja |
| 4 | Coaching reaction | L1 binary | 15% | coaching_actions zrealizowane: WARN_LEECH, MENTION_PAST, HINT_DELAY |
| 5 | Scenario-specific | L1 binary | 10% | Checkpoints unikalne per scenariusz (patrz sekcja 5) |

## 5. Scenariusze (z fixed scripts i binary checkpoints)

Kazdy scenariusz ma:
- **Fixed student script** — deterministic odpowiedzi ucznia (ewaluator NIE improwizuje)
- **Binary checkpoints** — TAK/NIE pytania do oceny zachowania tutora

### 5.1 first_session

**Fixed student script:**
```
wymiana_1_uczen: "Chce zaczac od TEORII"
wymiana_2_uczen: "[poprawna odpowiedz — dokladna wartosc z pre-fetched ANSWER_TEORIA]"
wymiana_3_uczen: "[bledna odpowiedz — mylenie_div_mod: np. '256 mod 10 = 25']"
wymiana_4_uczen: "[poprawna po pytaniu sokratejskim tutora]"
wymiana_5_uczen: "dalej" (nastepne cwiczenie)
```

**Binary checkpoints:**
```
CLI compliance:
[ ] progress status NA STARCIE
[ ] Przedstawienie 4 blokow (TEORIA, IMPLEMENTACJA, ARKUSZ, SQL)
[ ] typ intro --typ sledzenie_algorytmu po wyborze ucznia
[ ] exercise next --typ sledzenie_algorytmu (NIE exercise question)
[ ] exercise answer --id X NIE pobrane przed proba ucznia
[ ] progress blad --exercise-id X --typ Y --kod Z --hint N (po wymiana_3)
[ ] progress update --id X --wynik Y --czas Z (po kazdym cwiczeniu)

Coaching reaction:
[ ] coaching_actions zrealizowane (jesli obecne w exercise next)

Scenario-specific:
[ ] Worked example z cheatsheet (typ intro first_in_type)
[ ] START_TS=$(date +%s) przed cwiczeniem
```

### 5.2 hint_progression

**Fixed student script:**
```
wymiana_1_uczen: "[bledna odpowiedz — np. brak tabeli sledzenia]"
wymiana_2_uczen: "[bledna odpowiedz — po pytaniu sokratejskim, inna pomylka]"
wymiana_3_uczen: "[bledna odpowiedz — po hincie L1, wciaz zla wartosc]"
wymiana_4_uczen: "poddaje sie"
wymiana_5_uczen: "[konsolidacja — poprawne wyjasnienie swoimi slowami]"
```

**Binary checkpoints:**
```
CLI compliance:
[ ] exercise next (NIE exercise question)
[ ] exercise answer NIE pobrane przed proba ucznia
[ ] progress blad PRZED exercise hints (kazdorazowo)
[ ] progress blad z --hint 0 (przed hintem) i --hint 1/2/3 (po hincie)
[ ] progress update --wynik walk_through --czas Z na koncu

Coaching reaction:
[ ] coaching_actions zrealizowane (jesli obecne)

Scenario-specific:
[ ] exercise hints ZABLOKOWANE przy probie < hint_delay (CLI zwraca HINT_LOCKED)
[ ] Progresja hintow: L1 → L2 (z cheatsheet_excerpt) → L3 (kluczowy krok)
[ ] cheatsheet_excerpt z exercise hints cytowany przy L2
[ ] Konsolidacja po walk_through ("Wyjasniej swoimi slowami...")
[ ] Wizualizacja ASCII po walk_through (sledzenie/drzewo)
```

### 5.3 difficulty_climb

**Fixed student script:**
```
wymiana_1_uczen: "[poprawna odpowiedz — cwiczenie 1]"
wymiana_2_uczen: "[poprawna odpowiedz — cwiczenie 2]"
wymiana_3_uczen: "[poprawna odpowiedz — cwiczenie 3]"
wymiana_4_uczen: "[poprawna odpowiedz — cwiczenie 4 (srednie)]"
```

**Binary checkpoints:**
```
CLI compliance:
[ ] 3x progress update --wynik poprawne_bez_pomocy
[ ] exercise next (auto-difficulty, NIE --trudnosc)
[ ] Nastepne cwiczenie po streak 3 ma trudnosc >= srednie

Coaching reaction:
[ ] coaching_actions zrealizowane (jesli obecne)

Scenario-specific:
[ ] START_TS i ELAPSED per cwiczenie
[ ] Brak hintow (uczen odpowiada poprawnie)
```

### 5.4 review_session

**Fixed student script:**
```
wymiana_1_uczen: "Powtorka" (po wyswietleniu zaleglosci)
wymiana_2_uczen: "[poprawna odpowiedz — powtorka 1]"
wymiana_3_uczen: "[bledna odpowiedz — powtorka 2]"
wymiana_4_uczen: "[poprawna po hincie L1]"
```

**Binary checkpoints:**
```
CLI compliance:
[ ] progress status NA STARCIE → zaleglosci > 0
[ ] exercise review uzyty (NIE exercise next)
[ ] Review PRIORYTET — tutor proponuje powtorke przed nowym materialem
[ ] progress blad PRZED hintem przy bledzie (wymiana_3)
[ ] progress update po kazdym cwiczeniu

Coaching reaction:
[ ] coaching_actions zrealizowane (jesli obecne)
[ ] Leech tag ostrzezenie jesli obecny w coaching_actions

Scenario-specific:
[ ] Propozycja: "Masz N zaleglosci. Powtorka czy nowy material?"
```

### 5.5 cke_unlock

**Fixed student script:**
```
wymiana_0_uczen: "[poprawna odpowiedz na cwiczenie srednie-trudne — streak rosnie do 8]"
wymiana_1_uczen: "sprawdzian sledzenie_algorytmu"
wymiana_2_uczen: "[poprawna odpowiedz na worked-example pytanie o pulapki]"
wymiana_3_uczen: "[czesciowo poprawna odpowiedz na sprawdzianie — 70% punktow]"
```

**Binary checkpoints:**
```
CLI compliance:
[ ] exercise next --typ sledzenie_algorytmu (dla wymiana_0)
[ ] progress update --wynik poprawne_bez_pomocy (po wymiana_0, streak→8)
[ ] cke worked-example --typ X PRZED sprawdzianem
[ ] cke get --typ X --exclude (wyklucz wczesniej robione)
[ ] START_TS i ELAPSED
[ ] cke save --id X --punkty N --max M

Coaching reaction:
[ ] coaching_actions zrealizowane (jesli obecne)

Scenario-specific:
[ ] Ogloszenie odblokowania w formacie "*** ODBLOKOWANO ***" (po progress update gdy streak=8→trudne)
[ ] Pytanie o pulapki po worked-example ("Co zapamiętasz z tych pułapek?")
[ ] Brak hintow na sprawdzianie ("To sprawdzian — na egzaminie tez nie bedzie hintow")
[ ] Ocena czesciowa wg zasady_oceniania
[ ] Ogloszenie formatu "=== SPRAWDZIAN TYPU ==="
[ ] Tryb krok-po-kroku dla srednie-trudne (sekcja E)
```

### 5.6 probna

**Fixed student script:**
```
wymiana_1_uczen: "probna 2023"
wymiana_2_uczen: "[poprawna odpowiedz — 1. podzadanie, pelne punkty]"
wymiana_3_uczen: "[bledna odpowiedz — 2. podzadanie, off_by_one]"
wymiana_4_uczen: "[czesciowo poprawna — 3. podzadanie, 50% punktow]"
wymiana_5_uczen: "pomin"
wymiana_6_uczen: "[poprawna odpowiedz — 5. podzadanie]"
wymiana_7_uczen: "przerwij"
wymiana_8_uczen: "tak"
```

**Binary checkpoints:**
```
CLI compliance:
[ ] exam meta --rok 2023
[ ] exam task --rok 2023 --zadanie N (per zadanie: 1, 2, 3)
[ ] START_TS i ELAPSED
[ ] progress blad przy bledzie (wymiana_3) z --hint 0
[ ] exam save --rok 2023 --results '[...]' --czas M

Coaching reaction:
[ ] coaching_actions (N/A — egzamin)

Scenario-specific:
[ ] Brak hintow (tryb egzaminacyjny)
[ ] Podsumowanie per-zadanie + per-kategoria
[ ] Wyswietlenie zasad egzaminu
[ ] Obsluga "pomin" (0 pkt, przejscie do nastepnego podzadania)
[ ] Obsluga "przerwij" (koniec egzaminu → podsumowanie)
```

### 5.7 coaching_aware

**Fixed student script:**
```
wymiana_1_uczen: "[po tutorowym ostrzezeniu o leech_tag] OK, bede uwazniejszy"
wymiana_2_uczen: "[bledna odpowiedz — mylenie_div_mod, ten sam blad co w historii]"
wymiana_3_uczen: "[poprawna odpowiedz po hincie]"
```

**Binary checkpoints:**
```
CLI compliance:
[ ] exercise next (NIE exercise question)
[ ] progress blad --exercise-id X --typ Y --kod mylenie_div_mod --hint N
[ ] progress update z wynikiem

Coaching reaction:
[ ] coaching_actions WARN_LEECH zrealizowane PRZED cwiczeniem
[ ] coaching_actions MENTION_PAST zrealizowane PO bledzie
[ ] Komunikat "Od teraz mniej podpowiedzi" obecny (jesli HINT_DELAY w actions)
[ ] hint_delay respektowany (exercise hints zablokowane przy probie < hint_delay)
[ ] progress blad PRZED exercise hints

Scenario-specific:
[ ] Tutor powiazal blad z historia ("Ostatnio miales problem z X")
[ ] hint_delay >= 2 (progressed student)
```

## 6. Orchestracja agentow

Dla kazdej pary (persona, scenariusz) uruchom agenta Task z promptem:

```
Jestes agentem testujacym jakosc korepetytora maturalnego.

## Twoje zadanie
Przeprowadz symulacje sesji korepetycji, grajac OBIE role:
- **Tutor**: postepuje DOKLADNIE wg ponizszego SKILL.md
- **Uczen**: postepuje DOKLADNIE wg fixed student script (NIE improwizuj odpowiedzi!)

## Testowany SKILL.md
<skill>
{SKILL_CONTENT}
</skill>

## Persona: {PERSONA_NAME}
{PERSONA_DESCRIPTION}

## Scenariusz: {SCENARIO_NAME}
{SCENARIO_DESCRIPTION}

## Fixed student script
{STUDENT_SCRIPT}

## Binary checkpoints
{CHECKPOINTS}

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
- Question (CKE_PRE): {EX_CKE_PRE} (only for cke_unlock — exercise before unlock)
- Answer (CKE_PRE): {ANSWER_CKE_PRE}
- Typ intro: {INTRO_TEORIA}
- Progress status: {STATUS}

## Instrukcje

### Krok 1: Symuluj dialog
Symuluj dialog tutor↔uczen (8-15 wymian). Tutor postepuje wg SKILL.md.
Uczen postepuje DOKLADNIE wg fixed student script — NIE improwizuj odpowiedzi.
Przy kazdej akcji tutora ZAPISZ komende CLI ktora tutor POWINIEN wywolac.

### Krok 2: Ocen Layer 1 (binary checkpoints)
Dla KAZDEGO checkpointu z listy sprawdz czy zostal spelniony w transkrypcie.
Odpowiedz TAK lub NIE. Policz trafienia.

### Krok 3: Ocen Layer 2 (holistic quality)

**Metoda sokratejska (1-5):**
- 5/5: Tutor pyta "Co sie stanie gdy n=0?" zamiast podac odpowiedz. Uczen sam dochodzi do rozwiazania.
- 3/5: Tutor mowi "Podpowiedz: sprawdz warunek bazowy" — daje kierunek ale nie pyta.
- 1/5: Tutor mowi "Odpowiedz to 13, bo..." — gotowe rozwiazanie.

**Ton i jezyk (1-5):**
- 5/5: Polski, "ty", zachecajacy, cierpliwy przez 3 proby, brak emoji. Feedback czasowy. Konsolidacja po walk_through.
- 3/5: Poprawny jezyk ale bez zachecania, lub sporadyczny emoji.
- 1/5: Angielski, formalny, emoji, lub agresywny ton.

### Krok 4: Oblicz score

```
L1_percent = (trafienia / total_checkpoints) * 100
L2_avg = (metoda_sokratejska * 0.625 + ton_i_jezyk * 0.375)  # znormalizowane wagi: 25/(25+15), 15/(25+15)
overall = 0.6 * L1_percent + 0.4 * (L2_avg / 5 * 100)
pass = overall >= 70
```

### Krok 5: Zwroc JSON

## Format odpowiedzi (DOKLADNIE ten JSON, nic wiecej)
```json
{
  "persona": "{PERSONA_NAME}",
  "scenario": "{SCENARIO_NAME}",
  "layer1": {
    "checkpoints_total": N,
    "checkpoints_passed": M,
    "score_percent": P,
    "details": {
      "checkpoint_name": true/false,
      ...
    }
  },
  "layer2": {
    "metoda_sokratejska": {"score": N, "max": 5, "uwagi": "..."},
    "ton_i_jezyk": {"score": N, "max": 5, "uwagi": "..."}
  },
  "overall_score": M,
  "pass": true/false,
  "transcript_excerpt": "... (3-5 kluczowych wymian) ...",
  "issues": ["issue1", "issue2"],
  "cli_commands_used": ["cmd1", "cmd2"]
}
```

WAZNE: Nie oceniaj lagodnie. Jesli SKILL.md nie mowi tutorowi zeby cos zrobil,
tutor tego NIE robi — i stracisz punkty. Badz surowy ale sprawiedliwy.
Checkpoints sa binarne — spelniony lub nie. Brak argumentu na TAK = NIE.
```

Spawns agentow **rownolegle** (Task tool z subagent_type=general-purpose).

## 7. Zbieranie raportow

Po zakonczeniu wszystkich agentow:

1. Parsuj JSON z kazdego agenta (wyciagnij z odpowiedzi blok JSON)
2. Oblicz overall score = srednia ze wszystkich uruchomien
3. Wygeneruj raport markdown

## 8. Raport — format

```markdown
# Test Tutor Report — {REPORT_DATE} (commit {COMMIT_HASH})

## Per-scenario results

### Persona: {persona} | Scenario: {scenario}

**Layer 1 — Binary checkpoints**: {passed}/{total} ({L1_percent}%)
| Checkpoint | Result |
|------------|--------|
| {name} | PASS/FAIL |
| ... | ... |

**Layer 2 — Holistic quality**:
| Kryterium | Score | Uwagi |
|-----------|-------|-------|
| Metoda sokratejska | {s}/5 | {uwagi} |
| Ton i jezyk | {s}/5 | {uwagi} |

| **OVERALL** | **{overall}/100** | **{PASS/FAIL}** |

[...powtorzone dla kazdej pary...]

## Summary
- **Overall**: {avg_score}/100 ({PASS/FAIL})
- **Pass rate**: {passed}/{total}
- **L1 avg**: {avg_L1}%
- **L2 avg**: {avg_L2}/5
- **Weakest checkpoints**: {najczesciej_failujace_checkpoints}
- **Issues**:
  - {issue1}
  - {issue2}

## Diff vs previous report
- Previous: {prev_file} -> {prev_score}/100
- Delta: {+/-N} points
- Improved: {criteria}
- Regressed: {criteria}
```

## 9. Zapis raportu

1. Zapisz raport do `{REPORT_FILE}` (Write tool)
2. Wyswietl raport w terminalu
3. Jesli istnieje poprzedni raport w `{REPORT_DIR}/`:
   - Przeczytaj ostatni (posortuj po dacie)
   - Porownaj overall score
   - Dodaj sekcje "Diff vs previous report"
4. Podsumuj: "Raport zapisany: `{REPORT_FILE}`"
5. **Append JSONL entry** to `{REPORT_DIR}/historia.json`:
   - For each scenario in this run, construct a JSON object matching this schema:
     ```json
     {"date":"{REPORT_DATE}","commit":"{COMMIT_HASH}","mode":"{MODE}","overall_score":{SCORE},"pass":{PASS},"scenario_count":{N},"scenarios":[{"persona":"...","scenario":"...","score":N.N,"l1_percent":N.N,"l1_total":N,"l1_passed":N,"l2":{"socratic":N,"tone":N},"issues":["..."]}]}
     ```
   - Append as a SINGLE LINE (JSONL format) using Bash: `echo '{JSON}' >> {REPORT_DIR}/historia.json`
   - Do NOT pretty-print. One JSON object per line.
6. **Generate cumulative report**:
   a. Run (Bash): `{CLI_PATH} test-report summary --historia {REPORT_DIR}/historia.json --format md`
   b. Capture the markdown output — this is the deterministic Go-computed section (dashboard, history table, per-scenario trends, L2 trends, evaluator noise)
   c. Write an **Interpretacja** section after the Go output. Read the numbers and trends from the Go output and write:
      - **Co sie zmienilo?** — explain score changes vs previous run using delta and per-scenario trends
      - **Top 3 do naprawienia** — the most impactful issues to fix, drawn from scenario issues in historia.json
      - **Uporczywe problemy** — issues appearing in 3+ consecutive runs (check issues arrays in latest entries of historia.json)
      - **Co dziala dobrze** — stable or improving metrics (scenarios with up-arrow trend, L2 criteria at 4.5+)
   d. Combine: Go markdown output + `## Interpretacja` section
   e. Save to `{REPORT_DIR}/RAPORT_ZBIORCZY.md` using Write tool (overwritten each run)
   f. Display: "Raport zbiorczy zapisany: `{REPORT_DIR}/RAPORT_ZBIORCZY.md`"

## 10. Cleanup

```bash
rm -rf /tmp/test-tutor-$$
```
