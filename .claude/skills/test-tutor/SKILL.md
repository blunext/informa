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
- **intermediate**: difficulty_climb, review_session
- **advanced**: cke_unlock, probna

## 2. Pre-fetch danych

Na poczatku pobierz dane potrzebne agentom (unikaj powtarzania CLI calls w agentach):

```bash
CLI_DIR="matura_informatyka_rozszerzona/analiza/cli"
MATURA="$CLI_DIR/matura"

# Testowany skill
SKILL_CONTENT=$(cat .claude/skills/matura/SKILL.md)

# Przykladowe cwiczenia (po 1 na typ)
EX_TEORIA=$($MATURA exercise get --typ sledzenie_algorytmu --trudnosc latwe)
EX_IMPL=$($MATURA exercise get --typ cyfry_liczby --trudnosc latwe)
EX_SQL=$($MATURA exercise get --typ sql_group_by --trudnosc latwe)
EX_ARKUSZ=$($MATURA exercise get --typ agregacja_warunkowa --trudnosc latwe)

# Typ intro
INTRO_TEORIA=$($MATURA typ intro --typ sledzenie_algorytmu)

# Progress status (fresh)
STATUS=$($MATURA --db-dir /tmp/test-tutor-$$ progress status)

# Cheatsheet excerpt
CHEAT_TEORIA=$($MATURA cheatsheet get --kategoria TEORIA --sekcja "archetyp")

# Raport metadata
REPORT_DATE=$(date +%Y-%m-%d)
COMMIT_HASH=$(git -C /Users/blt1wz/priv/informa rev-parse --short HEAD)
REPORT_DIR="matura_informatyka_rozszerzona/analiza/test_pedagogical/reports"
REPORT_FILE="${REPORT_DIR}/${REPORT_DATE}_${COMMIT_HASH}.md"
```

Utworz temp DB dla kazdego agenta (izolacja):
```bash
mkdir -p /tmp/test-tutor-$$
cp "$CLI_DIR/matura.db" /tmp/test-tutor-$$/matura.db
```

## 3. Persony

### Beginner (Poczatkujacy)
- **Accuracy**: 50% — co drugie cwiczenie blednie
- **Hint probability**: 70% — czesto prosi o wskazowke
- **Typowe bledy**: `mylenie_div_mod`, `brak_group_by`, `zle_adresowanie`, `brak_tabeli_sledzenia`
- **Zachowanie**: pisze krotkie odpowiedzi, czesto "nie wiem", myli podstawowe pojecia, nie zna SQL, myli div/mod
- **Tempo**: wolne, 2x-3x benchmark CKE
- **Sesja**: 3-5 cwiczen

### Intermediate (Sredniozaawansowany)
- **Accuracy**: 70% — wiekszosc poprawnie, ale z lukami
- **Hint probability**: 35% — czasem potrzebuje naprowadzenia
- **Typowe bledy**: `pominiecie_bazy_rekurencji`, `brak_having`, `off_by_one`, `mylenie_avg_worst`
- **Zachowanie**: podaje rozumowanie, ale pomija edge cases, zna skladnie SQL ale myli GROUP BY/HAVING, rozumie rekurencje ale gubi sie w bazach
- **Tempo**: w normie, 1x-1.5x benchmark
- **Sesja**: 5-8 cwiczen

### Advanced (Zaawansowany)
- **Accuracy**: 90% — rzadko sie myli
- **Hint probability**: 10% — prawie nigdy nie potrzebuje
- **Typowe bledy**: `mylenie_avg_worst`, subtelne edge cases, `brak_uzasadnienia_zlozonosc`
- **Zachowanie**: podaje kompletne rozwiazania, uzasadnia, ale czasem pomija uzasadnienie przy P/F, myli avg/worst case
- **Tempo**: szybkie, < benchmark
- **Sesja**: 8-12 cwiczen

## 4. Rubryka oceny (7 kryteriow)

Kazde kryterium oceniane 0-5 pkt, z waga procentowa:

| # | Kryterium | Waga | 5 (wzorcowe) | 3 (dostateczne) | 1 (nieakceptowalne) |
|---|-----------|------|--------------|-----------------|---------------------|
| 1 | Metoda sokratejska | 25% | Tutor ZAWSZE pyta przed podaniem odpowiedzi, uczen probuje pierwszy, pytania naprowadzajace | Tutor czasem podaje odpowiedz bez pytania | Tutor podaje gotowe rozwiazania |
| 2 | Progresja hintow | 20% | L1→L2→L3 w scislej kolejnosci, cheatsheet przy L2, konsolidacja po walk_through | Kolejnosc zachowana ale brak cheatsheet | Brak progresji, hinty pominięte lub w zlej kolejnosci |
| 3 | Sledzenie bledow | 15% | `progress blad --kod X` po kazdym bledzie, `diagnose` co 5 cw., analiza wzorcow | Bledy rejestrowane ale bez diagnozy | Brak rejestrowania bledow |
| 4 | Adaptacja trudnosci | 15% | Streak 3→srednie, 5→sr-trudne, 8→trudne. Walk_through→latwe. Progi przestrzegane | Adaptacja obecna ale progi nieścisłe | Brak adaptacji trudnosci |
| 5 | Powtorki SR | 10% | Review priorytet gdy zaleglosci, exercise next uzywany prawidlowo | SR obecne ale bez priorytetu | Brak sprawdzania zaleglosci |
| 6 | Ton i jezyk | 10% | Polski, "ty", bez emoji, zachecanie, cierpliwosc, feedback czasowy | Poprawny jezyk ale bez zachecania | Angielski, formalny, emoji, lub brak feedbacku |
| 7 | Integralnosc CLI | 5% | Wszystkie komendy poprawne, brak halucynacji cwiczen, prawidlowe ID | Drobne bledy w komendach | Halucynowane cwiczenia, bledne komendy |

**Scoring**: score = sum(kryterium_score * waga). Max = 5.0. Przelicz na 0-100: score/5*100.
**Prog zdania**: >= 70/100.

## 5. Scenariusze

### 5.1 first_session
- Uczen nowy (cwiczenia_lacznie == 0)
- **Oczekiwania tutora**: sprawdza `progress status`, widzi 0, przedstawia 4 bloki, pyta o wybor
- **Uczen**: wybiera TEORIA
- **Oczekiwania**: tutor wywoluje `typ intro`, przedstawia kategorie, daje 1 cwiczenie intro
- **Przebieg**: 3 cwiczenia latwe, uczen odpowiada wg accuracy persony

### 5.2 hint_progression
- Uczen pracuje nad cwiczeniami, trafia na trudne
- **Przebieg**: uczen 3x odpowiada blednie na to samo cwiczenie
- **Oczekiwania**: L1 (pytanie sokratejskie) → L2 (cheatsheet section + wskazowka) → L3 (kluczowy krok) → walk_through
- **Po walk_through**: konsolidacja ("wyjasniej swoimi slowami")
- **Kluczowe**: tutor rejestruje blad (`progress blad`) na kazdym etapie

### 5.3 difficulty_climb
- Uczen odpowiada 3x poprawnie bez pomocy
- **Oczekiwania**: `progress update --wynik poprawne_bez_pomocy` → streak 3 → tutor przechodzi na srednie
- **Kluczowe**: po update sprawdz ze nowe cwiczenie ma trudnosc `srednie`

### 5.4 review_session
- Uczen wraca po przerwie, ma zaleglosci SR
- **Oczekiwania**: tutor sprawdza status, widzi zaleglosci, proponuje powtorke
- **Przebieg**: `exercise review` → uczen rozwiazuje 2-3 zaleglosci
- **Kluczowe**: review ma priorytet nad nowym materialem

### 5.5 cke_unlock
- Uczen ma streak 8, osiagnal poziom trudne
- **Oczekiwania**: tutor ogłasza odblokowanie sprawdzianu CKE
- **Przebieg**: uczen prosi o `sprawdzian`, tutor pobiera `cke get`
- **Kluczowe**: brak hintow na sprawdzianie, ocena czesciowa wg zasady_oceniania

### 5.6 probna
- Uczen prosi o probna mature (skrocona: 3 zadania)
- **Oczekiwania**: tutor pobiera `exam meta`, wyswietla zasady, prowadzi sekwencyjnie
- **Przebieg**: 3 zadania, uczen odpowiada z rozna trafnoscia
- **Kluczowe**: brak hintow, podsumowanie per-zadanie + per-kategoria, zapis `exam save`

## 6. Orchestracja agentow

Dla kazdej pary (persona, scenariusz) uruchom agenta Task z promptem:

```
Jestes agentem testujacym jakosc korepetytora maturalnego.

## Twoje zadanie
Przeprowadz symulacje sesji korepetycji, grajac OBIE role:
- **Tutor**: postepuje DOKLADNIE wg ponizszego SKILL.md
- **Uczen**: postepuje wg specyfikacji persony

## Testowany SKILL.md
<skill>
{SKILL_CONTENT}
</skill>

## Persona: {PERSONA_NAME}
{PERSONA_DESCRIPTION}

## Scenariusz: {SCENARIO_NAME}
{SCENARIO_DESCRIPTION}

## Pre-fetched data
- Exercise (TEORIA): {EX_TEORIA}
- Exercise (IMPL): {EX_IMPL}
- Typ intro: {INTRO_TEORIA}
- Progress status: {STATUS}

## Instrukcje
1. Symuluj dialog tutor↔uczen (8-15 wymian). Tutor postepuje wg SKILL.md,
   uczen wg persony (accuracy, typowe bledy, zachowanie).
2. Przy kazdej akcji tutora ZAPISZ komende CLI ktora tutor POWINIEN wywolac
   (np. `./matura exercise get --typ X`, `./matura progress update --id Y --wynik Z`).
3. Po symulacji OCEN transkrypt wg ponizszej rubryki.
4. Zwroc wynik DOKLADNIE w formacie JSON ponizej.

## Rubryka
{RUBRIC_TABLE}

## Format odpowiedzi (DOKLADNIE ten JSON, nic wiecej)
```json
{
  "persona": "{PERSONA_NAME}",
  "scenario": "{SCENARIO_NAME}",
  "scores": {
    "metoda_sokratejska": {"score": N, "max": 5, "uwagi": "..."},
    "progresja_hintow": {"score": N, "max": 5, "uwagi": "..."},
    "sledzenie_bledow": {"score": N, "max": 5, "uwagi": "..."},
    "adaptacja_trudnosci": {"score": N, "max": 5, "uwagi": "..."},
    "powtorki_sr": {"score": N, "max": 5, "uwagi": "..."},
    "ton_i_jezyk": {"score": N, "max": 5, "uwagi": "..."},
    "integralnosc_cli": {"score": N, "max": 5, "uwagi": "..."}
  },
  "weighted_score": M,
  "pass": true/false,
  "transcript_excerpt": "... (3-5 kluczowych wymian) ...",
  "issues": ["issue1", "issue2"],
  "cli_commands_used": ["cmd1", "cmd2"]
}
```

Score oblicz: sum(score * waga) / 5 * 100, gdzie wagi:
metoda_sokratejska=0.25, progresja_hintow=0.20, sledzenie_bledow=0.15,
adaptacja_trudnosci=0.15, powtorki_sr=0.10, ton_i_jezyk=0.10, integralnosc_cli=0.05.

Pass = weighted_score >= 70.

WAZNE: Nie oceniaj lagodnie. Jesli SKILL.md nie mowi tutorowi zeby cos zrobil,
tutor tego NIE robi — i stracisz punkty. Badz surowy ale sprawiedliwy.
```

Spawns agentow **rownolegle** (Task tool z subagent_type=general-purpose).

## 7. Zbieranie raportow

Po zakonczeniu wszystkich agentow:

1. Parsuj JSON z kazdego agenta (wyciagnij z odpowiedzi blok JSON)
2. Oblicz overall score = srednia wazona ze wszystkich uruchomien
3. Wygeneruj raport markdown

## 8. Raport — format

```markdown
# Test Tutor Report — {REPORT_DATE} (commit {COMMIT_HASH})

## Per-scenario results

### Persona: {persona} | Scenario: {scenario}
| Kryterium | Score | Uwagi |
|-----------|-------|-------|
| Metoda sokratejska | {s}/5 | {uwagi} |
| Progresja hintow | {s}/5 | {uwagi} |
| Sledzenie bledow | {s}/5 | {uwagi} |
| Adaptacja trudnosci | {s}/5 | {uwagi} |
| Powtorki SR | {s}/5 | {uwagi} |
| Ton i jezyk | {s}/5 | {uwagi} |
| Integralnosc CLI | {s}/5 | {uwagi} |
| **SCORE** | **{weighted}/100** | **{PASS/FAIL}** |

[...powtorzone dla kazdej pary...]

## Summary
- **Overall**: {avg_score}/100 ({PASS/FAIL})
- **Pass rate**: {passed}/{total}
- **Weakest**: {najslabsze_kryterium} (avg {x}/5)
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

## 10. Cleanup

```bash
rm -rf /tmp/test-tutor-$$
```
