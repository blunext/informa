# Design: CLI Guardrails + SKILL.md + Test-Tutor Redesign

**Data**: 2026-02-23
**Commit bazowy**: c2e9ce9
**Podejscie**: D — "CLI as Guardrail Layer"

## Kontekst

Analiza regresyjna 8 pelnych uruchomien test-tutor (F7) wykazala:
- **Plateau**: srednia 87.9/100, odch.std 2.5, zakres 84-91
- **Szum ewaluatorski**: 1-2 pkt overall, 5-14 pkt per-scenariusz
- **Trwale problemy**: exercise question zamiast next (5/8), nietrafne kody bledow (6/8), first_session wariancja (odch.std 7.4)

### Cel
- Podniesc srednia >92/100
- Zmniejszyc wariancje (odch.std <1.5)
- Zachowac naturalnosc tutoringu

### Kluczowy insight
Przeniesc logike walidacyjna z SKILL.md (LLM musi pamietac) do CLI (Go binary wymusza). LLM skupia sie na jakosci dialogu, CLI na compliance.

---

## 1. CLI Guardrails (nowe zachowania Go binary)

### 1.1 Lazy loading enforcement

CLI sledzi per-exercise stan w `matura_progress.db`:

```sql
CREATE TABLE IF NOT EXISTS active_exercises (
    exercise_id TEXT PRIMARY KEY,
    fetched_at TEXT NOT NULL,
    attempt_count INTEGER DEFAULT 0,
    hints_fetched BOOLEAN DEFAULT FALSE,
    answer_fetched BOOLEAN DEFAULT FALSE
);
```

| Komenda | Warunek | Przy naruszeniu |
|---------|---------|-----------------|
| `exercise answer --id X` | `attempt_count >= 1` | `ERROR: lazy_loading — student hasn't attempted yet. Record attempt first via 'progress blad' or 'progress update'.` |
| `exercise hints --id X` | `attempt_count >= hint_delay` | `{"status":"HINT_LOCKED","attempt":N,"hint_delay":D,"action":"Zadaj pytanie sokratejskie BEZ hintow"}` |

`exercise next` / `exercise review` automatycznie rejestruja nowe cwiczenie w `active_exercises`.

### 1.2 Hint delay enforcement

CLI odmawia `exercise hints` dopoki `attempt_count < hint_delay`:
- `progress blad` inkrementuje `attempt_count`
- `exercise hints` sprawdza `attempt_count >= hint_delay` (hint_delay z coaching)
- Przy locku zwraca structured JSON z instrukcja

### 1.3 Error code validation

`progress blad --exercise-id X --typ Y --kod Z`:
- CLI sprawdza Z vs whitelist per typ Y (hardcoded w Go)
- Przy blednym kodzie: `ERROR: kod 'Z' niedostepny dla typ 'Y'. Dozwolone: [lista]`
- `--hint N` wymagany parametr (fail bez tego)

Whitelist per typ (przeniesiony z SKILL.md sekcja F):
- sledzenie_algorytmu: mylenie_div_mod, zla_kolejnosc_sledzenia, pominiecie_bazy_rekurencji, ...
- sql_group_by: brak_having, brak_group_by, zla_kolejnosc_klauzul, ...
- cyfry_liczby: mylenie_div_mod, off_by_one, brak_inicjalizacji, ...
- (pelna lista w kodzie Go, nie w SKILL.md)

### 1.4 Coaching actions (structured output)

`exercise next` i `exercise review` zwracaja nowe pole `coaching_actions`:

```json
{
  "id": "7.23",
  "coaching": { ... },
  "coaching_actions": [
    "WARN_LEECH: Tag 'cyfry-mod-div' sprawia Ci trudnosc — zwroc uwage",
    "MENTION_PAST: Ostatnio mialeS problem z 'mylenie_div_mod'",
    "HINT_DELAY: 2 (Od teraz mniej podpowiedzi — rozwijasz samodzielnosc)"
  ]
}
```

Generowanie w Go:
- `WARN_LEECH` — jesli coaching.leech_tags niepuste
- `MENTION_PAST` — jesli coaching.past_mistakes niepuste
- `HINT_DELAY` — jesli hint_delay >= 2 i to pierwszy raz z tym delay

### 1.5 Usuniecie `exercise question` ze SKILL.md

- SKILL.md odwoluje sie WYLACZNIE do `exercise next`
- `exercise question` zostaje w CLI (do debug/test), ale nie jest w SKILL.md
- `exercise next` automatycznie zarzadza: review > interleave > new, auto-difficulty, pool warnings

### 1.6 Auto-diagnose w progress update

`progress update` response zawiera pole `diagnose` gdy count % 5 == 0:

```json
{
  "status": "ok",
  "diagnose": {
    "top_errors": [{"kod": "mylenie_div_mod", "count": 3}],
    "recommendation": "Powtorz sledzenie z div/mod"
  }
}
```

Tutor nie musi pamietac o wywoaniu `progress diagnose` co 5 cwiczen — CLI robi to automatycznie.

### 1.7 Kompatybilnosc z /clear pattern

Wszystkie guardrails sa DB-backed (nie in-memory). CLI jest bezstanowy (kazdy call = nowy proces). `/clear` + `/matura` dziala identycznie jak teraz — `progress status` odczytuje caly stan z DB.

---

## 2. SKILL.md Redesign

### 2.1 Nowa struktura

```
A. Rola i jezyk (ton, Socratic, polski, first-type intro)
B. CLI Reference (zaktualizowana tabela — bez exercise question)
C. Start sesji (3 scenariusze, dedykowany first_session checklist)
D. Pobieranie cwiczen (TYLKO exercise next, auto-difficulty)
E. Prezentacja (coaching_actions jako gotowe instrukcje)
F. Ocena odpowiedzi (6-krokowy checklist, CLI wymusza lazy/delay/kody)
   F.1 Checklist
   F.2 Wizualizacje proaktywne
G. Komendy ucznia (wskazowka, poddaje sie, wyjasniej)
H. Sprawdzian CKE / Probna matura
I. Reset kontekstu
```

### 2.2 Uproszczony checklist sekcja F (9 -> 6 krokow)

Obecny (9 krokow):
1. Pobierz odpowiedz (lazy) — CLI WYMUSZA
2. Porownaj
3. Jesli poprawna -> update
4. Jesli bledna -> progress blad PRZED hintem
5. Sprawdz hint_delay, zdecyduj — CLI WYMUSZA
6. Jesli hint -> pobierz exercise hints (lazy)
7. Wizualizacja
8. Progress update
9. Diagnoza co 5 cw. — CLI AUTO

Nowy (6 krokow):
1. Porownaj odpowiedz ucznia z wzorcowa (`exercise answer --id X` — CLI zablokuje jesli za wczesnie)
2. POPRAWNA: `progress update --id X --wynik Y --czas Z`
3. BLEDNA: `progress blad --exercise-id X --typ Y --kod Z --hint N` (CLI waliduje kod i wymaga --hint)
4. Sprobuj `exercise hints --id X` — CLI zwroci hinty LUB HINT_LOCKED z instrukcja
5. Jesli hint odblokowany: pokaz wskazowke + pytanie sokratejskie
6. Wizualizacja proaktywna (jak wczesniej)

### 2.3 Sekcja C: First session checklist

```
Scenariusz 1 — Pierwszy kontakt (cwiczenia_lacznie == 0)
[CHECKLIST]
1. progress status -> widzi 0
2. Przedstaw 4 bloki (TEORIA, IMPLEMENTACJA, ARKUSZ, SQL)
3. Uczen wybiera -> typ intro --typ {wybrany}
4. Worked example z cheatsheet
5. exercise next --typ {wybrany}
6. Po odpowiedzi ucznia -> exercise answer (CLI zablokuje jesli za wczesnie)
```

### 2.4 Sekcja E: Coaching jako gotowe instrukcje

```
Przeczytaj coaching_actions z exercise next/review.
Kazda akcje wlacz naturalnie w dialog PRZED podaniem tresci cwiczenia:
- WARN_LEECH: "Uwaga — temat X sprawia Ci trudnosc"
- MENTION_PAST: "Ostatnio miales problem z Y"
- HINT_DELAY: "Od teraz mniej podpowiedzi"
```

---

## 3. Test-Tutor Redesign

### 3.1 Dwuwarstwowy scoring

**Layer 1 — Binary checkpoints (60% wagi)**

Per scenariusz: lista TAK/NIE pytan. Ewaluator sprawdza konkretne fakty, nie ocenia holistycznie.

**Layer 2 — Holistic quality (40% wagi)**

Tradycyjna ocena 1-5 dla kryteriow jakosciowych (metoda sokratejska, ton).

**Laczny score**: `0.6 * L1_percent + 0.4 * (L2_avg/5 * 100)`

### 3.2 Redukcja kryteriow: 8 -> 5

| Nowe kryterium | Typ | Waga | Zastepuje |
|----------------|-----|------|-----------|
| CLI compliance | L1 binary | 35% | Integralnosc CLI + Progresja hintow + Sledzenie bledow + Adaptacja trudnosci + Powtorki SR |
| Metoda sokratejska | L2 holistic | 25% | Metoda sokratejska |
| Ton i jezyk | L2 holistic | 15% | Ton i jezyk |
| Coaching reaction | L1 binary | 15% | Coaching |
| Scenario-specific | L1 binary | 10% | Nowe — checkpoints unikalne per scenariusz |

### 3.3 Binary checkpoints per scenariusz

**first_session**:
```
[ ] progress status NA STARCIE
[ ] Przedstawienie 4 blokow
[ ] typ intro po wyborze ucznia
[ ] exercise next (nie exercise question)
[ ] exercise answer NIE pobrane przed proba ucznia
[ ] progress blad z --kod i --hint N
[ ] progress update z --wynik i --czas
[ ] coaching_actions zrealizowane (jesli obecne)
```

**hint_progression**:
```
[ ] exercise next (nie question)
[ ] exercise answer NIE pobrane przed proba
[ ] exercise hints ZABLOKOWANE przy probie < hint_delay
[ ] progress blad PRZED exercise hints
[ ] Progresja L1 -> L2 -> L3 (jesli 3 bledy)
[ ] Konsolidacja po walk_through (pytanie "wyjasniej swoimi slowami")
[ ] cheatsheet przy L2
```

**difficulty_climb**:
```
[ ] 3x progress update --wynik poprawne_bez_pomocy
[ ] Nastepne cwiczenie ma wyzszy poziom trudnosci (exercise next auto)
[ ] progress blad z poprawnym kodem przy bledzie
```

**review_session**:
```
[ ] progress status -> zaleglosci > 0
[ ] exercise review uzyty (nie exercise next)
[ ] Review PRIORYTET nad nowym materialem
[ ] progress blad PRZED hintem przy bledzie
[ ] Leech tag ostrzezenie jesli obecny
```

**coaching_aware**:
```
[ ] coaching_actions WARN_LEECH zrealizowane PRZED cwiczeniem
[ ] coaching_actions MENTION_PAST zrealizowane PO bledzie
[ ] hint_delay respektowany (exercise hints zablokowane)
[ ] Komunikat "Od teraz mniej podpowiedzi" obecny
[ ] progress blad PRZED exercise hints
```

**cke_unlock**:
```
[ ] Ogloszenie odblokowania w formacie "*** ODBLOKOWANO ***"
[ ] cke worked-example PRZED sprawdzianem
[ ] Pytanie o pulapki po worked-example
[ ] cke get z --exclude
[ ] Brak hintow na sprawdzianie ("To sprawdzian")
[ ] cke save z punktami
[ ] START_TS i ELAPSED
```

**probna**:
```
[ ] exam meta wywolane
[ ] exam task per zadanie
[ ] START_TS i ELAPSED
[ ] Brak hintow (tryb egzaminacyjny)
[ ] progress blad przy bledzie
[ ] exam save z pelnym JSON
[ ] Podsumowanie per-zadanie i per-kategoria
```

### 3.4 Deterministyczne scenariusze (fixed student scripts)

Zamiast ewaluatora improwizujacego odpowiedzi ucznia, kazdy scenariusz ma fixed script:

```
first_session:
  wymiana_1_uczen: "Chce zaczac od TEORII"
  wymiana_2_uczen: "[poprawna odpowiedz — podaj dokladna wartosc z pre-fetched answer]"
  wymiana_3_uczen: "[bledna odpowiedz — mylenie_div_mod: np. 256 mod 10 = 25]"
  wymiana_4_uczen: "[poprawna po pytaniu sokratejskim]"
  wymiana_5_uczen: "[poprawna — 2. cwiczenie]"
```

Ewaluator NIE improwizuje odpowiedzi ucznia — dostaje gotowy script. Ocenia TYLKO zachowanie tutora. To redukuje wariancje per-scenariusz o ~50%.

### 3.5 Anchor examples w promptach ewaluatorow

Dla L2 holistic kryteriow:

```
Metoda sokratejska:
  5/5: Tutor pyta "Co sie stanie gdy n=0?" zamiast podac odpowiedz.
       Uczen sam dochodzi do rozwiazania.
  3/5: Tutor mowi "Podpowiedz: sprawdz warunek bazowy" — daje kierunek ale nie pyta.
  1/5: Tutor mowi "Odpowiedz to 13, bo..." — gotowe rozwiazanie.

Ton i jezyk:
  5/5: Polski, "ty", zachecajacy, cierpliwy przez 3 proby, brak emoji.
  3/5: Poprawny jezyk ale bez zachecania, lub sporadyczny emoji.
  1/5: Angielski, formalny, lub agresywny ton.
```

---

## 4. Oczekiwane zyski

| Metryka | Obecna (F7) | Oczekiwana | Zrodlo zysku |
|---------|-------------|------------|--------------|
| Srednia overall | 87.9 | **92-95** | CLI wymusza compliance, redukcja N/A |
| Odch.std overall | 2.5 | **<1.5** | Deterministyczne scripty, binary scoring |
| first_session zakres | 72-95 (23) | **85-95 (10)** | Fixed script + CLI guardrails |
| probna srednia | 82.1 | **88-92** | Dedicated checkpoints zamiast 5x N/A |
| Integralnosc CLI zakres | 3.86-5.00 | **4.5-5.0** | CLI waliduje — bledy niemozliwe |

## 5. Plan implementacji (orientacyjna kolejnosc)

1. **CLI guardrails** (Go code): active_exercises table, lazy loading enforcement, hint delay enforcement, error code validation, coaching_actions generation, auto-diagnose
2. **CLI tests**: Rozszerzenie main_test.go o testy guardrails
3. **SKILL.md rewrite**: Nowa struktura A-I z uproszczonym checklistem
4. **test-tutor rewrite**: 5 kryteriow, binary checkpoints, fixed scripts, anchor examples
5. **test_qa.sh update**: Layer 3 (SKILL lint) dostosowany do nowej struktury
6. **Walidacja**: 2-3 uruchomienia test-tutor na nowym designie, porownanie z baseline
