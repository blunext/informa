# Design: Lazy Loading Exercises + Coaching

**Data**: 2026-02-22
**Status**: Zaakceptowany

## Problem

CLI (`exercise get`) zwraca AI wszystko naraz: treść, hinty, pełną odpowiedź (~2KB).
Skutki:
- AI widzi odpowiedź od razu — hinty sokratejskie są sztuczne (podprowadzanie do znanej odpowiedzi)
- Marnowanie tokenów — hinty i odpowiedź ładują się nawet gdy uczeń sam rozwiąże zadanie
- ~60% payloadu jest zbędne na etapie nauczania

## Rozwiązanie

Rozbić `exercise get` na 3 komendy. AI pobiera dane przyrostowo — treść na starcie, hinty gdy uczeń się zaciął, odpowiedź przy ocenie. CLI dostarcza `coaching` z kontekstem ucznia (obliczany z progress.db).

## Nowe komendy CLI

### `exercise question --typ X [--trudnosc T] [--exclude ids]`

Zastępuje `exercise get`. Zwraca treść + coaching, bez hintów i odpowiedzi.

```json
{
  "id": "8.1",
  "typ_nazwa": "napisy",
  "kategoria": "IMPLEMENTACJA",
  "trudnosc": "latwe",
  "punkty": 2,
  "zrodlo": "Matura 2025 zad. 2.1",
  "tagi": ["palindrom", "wczytywanie-pliku"],
  "tresc": "W pliku napisy.txt...",
  "coaching": {
    "student_level": "familiar",
    "hint_delay": 2,
    "leech_tags": ["cyfry-mod-div"],
    "past_mistakes": ["s[n-i] zamiast s[n-1-i]"]
  }
}
```

### `exercise hints --id X`

Pobiera hinty na żądanie. Respektuje `max_hints` (obliczane z poziomu ucznia).

```json
{
  "id": "8.1",
  "wskazowki": ["Kierunek: ...", "Podejście: ...", "Kluczowy krok: ..."],
  "max_hints": 3
}
```

### `exercise answer --id X`

Pobiera pełną odpowiedź przy ocenie.

```json
{
  "id": "8.1",
  "odpowiedz": "**Rozwiązanie (C++):** ...",
  "typowe_bledy": [{"opis": "...", "kara": "-2 pkt"}]
}
```

## Pole `coaching`

Obliczane w locie z `progress.db`. Żadnych nowych tabel — dane już istnieją.

### `student_level`

Interpretowany przez CLI na podstawie streak + zrobione:

| Level | Warunek |
|-------|---------|
| `new` | 0 zrobionych danego typu |
| `learning` | 1-3 zrobione, streak < 3 |
| `familiar` | 4+ zrobione LUB streak >= 3 |
| `mastered` | poziom_trudnosci = trudne |

### `hint_delay`

Sugerowana liczba prób ucznia przed podaniem hinta:

| student_level | hint_delay |
|---------------|------------|
| new | 1 |
| learning | 1 |
| familiar | 2 |
| mastered | 3 |

SKILL.md traktuje `hint_delay` jako sugestię — jeśli uczeń jawnie prosi o pomoc, AI pobiera hint natychmiast.

### `leech_tags`

Tagi z FSRS-5 gdzie uczeń wielokrotnie oblał i wciąż nie opanował:

```sql
SELECT tag FROM progress_tagi
WHERE lapses >= 3
  AND retrievability(elapsed_days, stability) < 0.85
```

Self-healing: gdy uczeń opanuje tag (retrievability rośnie), znika z leech.

### `past_mistakes`

Konkretne błędy z ostatnich sesji, filtrowane po tagach bieżącego ćwiczenia:

```sql
SELECT DISTINCT blad_opis FROM progress_bledy
WHERE typ = ?
  AND blad_kod IN (tagi bieżącego ćwiczenia)
  AND data IN (
    SELECT DISTINCT data FROM progress_zrobione
    ORDER BY data DESC LIMIT 5
  )
ORDER BY data DESC LIMIT 3
```

Skaluje się z rytmem ucznia (ostatnie 5 sesji, nie dni kalendarzowych).

### `previous_result` (tylko review)

Wynik ucznia z ostatniego podejścia do tego ćwiczenia:

```sql
SELECT wynik FROM progress_zrobione WHERE id = ?
```

Tylko w `exercise review` / `exercise next` (mode=review). Pomijane dla nowych ćwiczeń.

## Nowe typy Go

```go
type QuestionOut struct {
    ID        string   `json:"id"`
    TypNazwa  string   `json:"typ_nazwa"`
    Kategoria string   `json:"kategoria"`
    Trudnosc  string   `json:"trudnosc"`
    Punkty    int      `json:"punkty"`
    Zrodlo    string   `json:"zrodlo"`
    Tagi      []string `json:"tagi"`
    Tresc     string   `json:"tresc"`
    Coaching  Coaching `json:"coaching"`
}

type Coaching struct {
    StudentLevel   string   `json:"student_level"`
    HintDelay      int      `json:"hint_delay"`
    LeechTags      []string `json:"leech_tags"`
    PastMistakes   []string `json:"past_mistakes"`
    PreviousResult string   `json:"previous_result,omitempty"`
}

type HintsOut struct {
    ID        string   `json:"id"`
    Wskazowki []string `json:"wskazowki"`
    MaxHints  int      `json:"max_hints"`
}

type AnswerOut struct {
    ID          string        `json:"id"`
    Odpowiedz   string       `json:"odpowiedz"`
    TypoweBledy []CommonError `json:"typowe_bledy"`
}
```

`ExerciseOut` zostaje wewnętrznie (import, queryExercises) — nie wystawiany na stdout.

## Zmiany w istniejących komendach

### `exercise review`

Zwraca `QuestionOut` + coaching (z `previous_result`) zamiast pełnego `ExerciseOut`.

### `exercise next`

Zwraca `QuestionOut` + coaching zamiast pełnego `ExerciseOut`. Wrapuje question/review decyzję jak dotąd.

## Zmiany w SKILL.md

Nowy flow:
1. `exercise question` → AI dostaje treść + coaching
2. AI uczy sokratejsko z własnej wiedzy
3. `coaching.hint_delay` prób minęło LUB uczeń prosi → `exercise hints --id X`
4. Uczeń daje odpowiedź → `exercise answer --id X` → AI ocenia

Reguła: `hint_delay` to sugestia, nie blokada. Uczeń proszący o pomoc = natychmiastowy override.

## Weryfikacja odpowiedzi bez pełnej odpowiedzi

| Kategoria | Jak AI weryfikuje | Potrzebuje `exercise answer`? |
|-----------|-------------------|-------------------------------|
| IMPLEMENTACJA | `tresc` zawiera `**Oczekiwany wynik**` (100% ćwiczeń) | Tylko do oceny jakości kodu |
| TEORIA | AI sama prześledzi algorytm / obliczy konwersję | Tak, przy ocenie |
| SQL | AI sama oceni query | Tak, przy ocenie |
| ARKUSZ | AI sama oceni formułę | Tak, przy ocenie |

## Co się NIE zmienia

- JSON-y ćwiczeń — zero edycji
- `cke get` — bez zmian (małe payloady, inna struktura)
- `progress update` — bez zmian
- `progress status` — bez zmian
- Import (`data import`) — bez zmian
- `matura.db` schema — bez zmian
- `progress.db` schema — bez zmian
- Walidator i verifier — bez zmian

## Zakres zmian

| Plik | Zmiana |
|------|--------|
| `commands.go` | Nowe komendy (question/hints/answer), coaching logic, zmiana review/next output |
| `types.go` | Nowe typy: QuestionOut, Coaching, HintsOut, AnswerOut |
| `main.go` | Rejestracja nowych komend, usunięcie exercise get |
| `main_test.go` | Nowe testy coaching, hints, answer; aktualizacja istniejących |
| `SKILL.md` | Nowy lazy flow, usunięcie reguł interpretacji surowych danych |
| `test_qa.sh` | Aktualizacja smoke testów (L1), journey testów (L6) |
