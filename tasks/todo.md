# Refaktor: CLI-first architecture (CLI → AI)

## Kontekst

Obecna architektura: AI (SKILL.md) jest kontrolerem, CLI (`matura`) jest bazą danych.
Docelowa architektura: CLI jest gospodarzem (stan, pedagogika, przepływ), AI jest gościem (język naturalny).

### Problem z obecnym modelem (AI → CLI)
- Logika pedagogiczna żyje w prompcie (SKILL.md) — trudna do testowania, wersjonowania, debugowania
- Zależność od platformy — działa tylko w Claude Code
- Niedeterminizm — ten sam uczeń, to samo ćwiczenie, inna sesja = inny przebieg
- Brak stanu sesji — CLI nie wie, że trwa sesja; każde wywołanie jest atomowe
- SKILL.md = spaghetti controller (orkiestracja w prozie)
- Ogromny koszt tokenów (cały SKILL.md w kontekście przy każdym wywołaniu)
- Testowalność wymaga spawnu agentów AI (test-tutor) zamiast `go test`

### Cel docelowy (CLI → AI)
- Logika pedagogiczna w kodzie Go (FSM, testowalna)
- AI wywoływane celowo, tylko do zadań językowych (ocena odpowiedzi, hinty, wyjaśnienia)
- Stan sesji w pamięci procesu
- Działa w dowolnym terminalu (nie tylko Claude Code)
- Tryb offline (degraded mode bez AI)
- Deterministic tests na pedagogice (`go test`)
- Mały koszt tokenów (celowane zapytania do API)

## Architektura docelowa

```
analiza/cli/
├── main.go
├── commands.go
├── database.go
├── importer.go
├── types.go
├── tutor/
│   ├── session.go     # stan sesji, FSM (maszyna stanów)
│   ├── pedagogy.go    # reguły: kiedy hint, kiedy odpowiedź, Socratic rules
│   ├── evaluator.go   # interface Evaluator { Evaluate(answer, exercise) → Feedback }
│   ├── ai_client.go   # implementacja: Claude API calls
│   ├── offline.go     # fallback: pattern matching bez AI
│   └── tutor_test.go  # deterministic tests
```

### FSM sesji korepetytora

```
PRESENT_EXERCISE → WAIT_ANSWER → EVALUATE →
  ├── CORRECT → UPDATE_PROGRESS → NEXT_EXERCISE
  └── INCORRECT → GENERATE_HINT → WAIT_ANSWER (loop, max 3)
```

### Gdzie wchodzi AI (i tylko tam)

1. **Ocena odpowiedzi otwartej** — czy odpowiedź ucznia jest poprawna/częściowa
2. **Generowanie hintu sokratejskiego** — dostosowanego do konkretnego błędu
3. **Tłumaczenie konceptu** — gdy uczeń prosi o wyjaśnienie

### Co pozostaje w kodzie Go (bez AI)

- Wybór ćwiczenia (algorytm, filtry, FSRS)
- Śledzenie postępu (progress update, SR intervals)
- Decyzja "hint czy odpowiedź" (reguły, próg prób)
- Sekwencjonowanie sesji (ile ćwiczeń, kiedy przerwa)
- Prezentacja treści (formatowanie, rendering)

## Plan migracji (przyrostowy)

### Faza 1: Interaktywna pętla CLI (bez AI)
- [ ] `matura tutor --typ {typ}` — interaktywna sesja w terminalu
- [ ] FSM w Go: present → wait → evaluate (porównanie z wzorcem) → next
- [ ] Prosty matching odpowiedzi (exact, normalized)
- [ ] Integracja z istniejącym progress tracking
- [ ] Testy jednostkowe na FSM i pedagogice

### Faza 2: AI evaluation (+Claude API)
- [ ] Interface `Evaluator` z dwoma implementacjami (pattern, ai)
- [ ] `--ai` flag + `ANTHROPIC_API_KEY` env
- [ ] AI ocenia odpowiedzi otwarte (C++, SQL, pseudokod)
- [ ] Fallback na pattern matching gdy brak API key
- [ ] Testy z mockiem AI

### Faza 3: AI hints (Socratic dialogue)
- [ ] AI generuje hinty sokratejskie w pętli
- [ ] Kontekst: ćwiczenie + odpowiedź ucznia + historia hintów
- [ ] Hint progression (ogólny → szczegółowy → odpowiedź)
- [ ] Rate limiting / cost tracking

### Faza 4: SKILL.md jako thin wrapper
- [ ] SKILL.md staje się cienkim wrapperem wołającym `matura tutor`
- [ ] Lub: SKILL.md deprecated, CLI jest jedynym interfejsem
- [ ] test-tutor migruje na `go test` + ewentualnie AI integration tests

## Pytania otwarte

- [ ] Który model Claude do evaluacji? (Haiku dla kosztów vs Sonnet dla jakości)
- [ ] Jak zarządzać API key? (env var, config file, keychain)
- [ ] Czy zachować SKILL.md jako alternatywny interfejs, czy deprecate?
- [ ] Streaming odpowiedzi AI w terminalu? (lepsze UX vs złożoność)
- [ ] Ile kontekstu dawać AI przy evaluacji? (samo ćwiczenie vs historia sesji)
- [ ] Język interfejsu CLI: polski czy angielski? (docelowi użytkownicy to polscy uczniowie)

## Notatki

- JSON files pozostają source of truth — CLI importuje z nich
- matura.db (SQLite) bez zmian — nowy kod korzysta z istniejącego database.go
- Obecny system (SKILL.md + test-tutor) działa równolegle podczas migracji
- Migracja nie blokuje żadnej istniejącej funkcjonalności
