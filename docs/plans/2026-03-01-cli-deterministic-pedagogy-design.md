# Design: CLI jako decydent pedagogiczny

**Data:** 2026-03-01
**Status:** Zatwierdzony
**Podejście:** TDD — testy przed implementacją

## Problem

AI wprowadza błędy w trzech obszarach:
1. **Kody błędów** — AI zgaduje z 230+ kodów, często trafia źle → diagnosis na śmieciach
2. **Ocena częściowa** — AI przyznaje dowolne punkty niezgodne z rubryką
3. **Guardrails hintów** — AI daje hinty bez czekania na próbę ucznia

## Zasada

Tam gdzie da się zdeterminizować — CLI decyduje sam.
Tam gdzie nie — CLI wymusza rygorystyczny protokół.

## Zmiany

### 1. `exercise suggest-error --id X [--student-answer Y]` (NOWA)

Dwie ścieżki:

**A — auto-detekcja** (typy z jednoznaczną odpowiedzią):
- CLI porównuje `student-answer` z poprawną odpowiedzią z JSON
- Wzorce: off-by-one (diff=±1), odwrocona_logika (P↔F), brak_algorytmu (answer=0/pusty)
- Zwraca: `{auto_detected: true, rekomendowany: "off_by_one", powod: "wynik o 1 za mały"}`

**B — zamknięta lista** (implementacja, SQL, arkusz):
- CLI zwraca wszystkie kody dla danego typu z opisami
- Zwraca: `{auto_detected: false, kody_dla_typu: [{kod, opis}, ...]}`

**~150 LOC**

### 2. `exercise check-answer --id X --answer Y` (NOWA)

Auto-scoring dla typów TEORIA z jednoznaczną odpowiedzią.

Normalizacja: trim whitespace, lowercase, porównanie numeryczne (13==13.0), boolean (PRAWDA/prawda/P).

Wspierane typy:
- sledzenie_algorytmu — TAK
- test_prawda_falsz — TAK
- konwersja_systemow — TAK
- analiza_algorytmu — CZĘŚCIOWO (gdy odpowiedź = wartość)
- projektowanie_algorytmu — NIE
- teoria_bezpieczenstwa — NIE

Zwraca: `{poprawne: bool, wynik: "pelne"|"zero", poprawna_odpowiedz?: string}`

**~200 LOC**

### 3. `progress update --wynik` enum guard (ZMIANA)

5 dozwolonych wartości:
- `pelne` (100%)
- `prawie_pelne` (75%)
- `czesciowe` (50%)
- `minimalne` (25%)
- `zero` (0%)

CLI liczy punkty: exercise.punkty × procent.

Migracja istniejących danych: >0.75→pelne, >0.5→prawie_pelne, >0.25→czesciowe, >0→minimalne, 0→zero.

**~30 LOC**

### 4. `exercise hints` guard no-attempt (ZMIANA)

Nowa kolumna `attempts_since_last_hint` w `active_exercises`.
- Rośnie przy `progress blad`
- Resetuje się przy `exercise hints`
- Hints zablokowane gdy = 0

Zwraca: `{error: "HINT_BLOCKED_NO_ATTEMPT", message: "..."}`

**~30 LOC**

### 5. Coaching actions z pre-formatowanym tekstem (ZMIANA)

`exercise next` zwraca coaching_actions z:
- `typ` — flaga (WARN_LEECH, SUGGEST_SLOWER, etc.)
- `tekst` — gotowe zdanie do wplecenia w dialog
- `priorytet` — wysoki/niski

Template strings w Go per typ akcji.

**~60 LOC**

### 6. SKILL.md HARD GATES (ZMIANA)

Nowe wymagania:
1. Przed `progress blad` → MUSI wywołać `suggest-error`
2. Dla auto-scorable typów → MUSI użyć `check-answer`
3. `--wynik` = enum (pelne/prawie_pelne/czesciowe/minimalne/zero)
4. Między hintami → CLI wymusza próbę ucznia

## Kolejność implementacji (TDD)

1. `progress update` enum guard — najprostsze, natychmiastowy efekt
2. `exercise hints` guard no-attempt — proste, wzmacnia guardrails
3. coaching actions z tekstem — średnie, wzbogaca istniejącą komendę
4. `exercise suggest-error` — złożone, fuzzy matching + auto-detekcja
5. `exercise check-answer` — najzłożoniejsze, normalizacja per typ
6. SKILL.md HARD GATES — na końcu, gdy CLI działa

## Czego NIE zmieniamy

- Metoda sokratejska (AI)
- Wizualizacje ASCII (AI)
- Negocjacja interleave (AI)
- FSRS algorytm (CLI bez zmian)
- Import / build pipeline (bez zmian)

## Szacunek

~470 LOC Go + testy TDD. 6 kroków implementacji.
