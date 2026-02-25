# Design: Naprawa wag kontekstu sesji w CLI matura

**Data**: 2026-02-25
**Cel**: Upewnić się że komunikat `/clear` + `/matura` pojawia się po ~9 ćwiczeniach (nie ~16)

## Kontekst

System wag kontekstu (`session_context_weight`) w CLI śledzi ile danych weszło do kontekstu AI. Gdy waga >= 80, `exercise next` zwraca `reset_suggested=true` i tutor sugeruje `/clear`.

Po ostatnich zmianach (Feb 24-25: cheatsheet auto-attach, exercise rubric, error code suggestions) kilka komend CLI nie ma wag lub ma zaniżone wagi. Efekt: /clear pojawia się po ~16 ćwiczeniach zamiast ~9.

Lazy loading (guardrails: HINT_LOCKED, LAZY_LOADING_BLOCKED) działa poprawnie — blokowane ścieżki nie dodają wagi. Problem dotyczy ścieżek które DOSTARCZAJĄ dane ale nie śledzą tego w wagach.

## Zmiany w `commands.go`

| Komenda | Linia | Teraz | Po zmianie | Warunek |
|---------|-------|-------|------------|---------|
| `exercise review` | ~691 | 0 | +4 per exercise | `addWeight(d, 4*len(results))` |
| `exercise rubric` | ~3070 | 0 | +1 | Wymaga dodania `d := db(cmd)` |
| `progress diagnose` | ~1484 | 0 | +2 | Już ma `d := db(cmd)` |
| `progress update` (nie walk_through) | ~872 | 0 | +1 | Dodać przed `if wynik == "walk_through"` |
| `progress blad` (hint<3) | ~1396 | 0 | +1 | Przenieść addWeight przed `if hint == 3` |
| `exercise hints` (z cheatsheet) | ~550 | +1 | +3 jeśli cheatsheet, +1 bez | Warunkowe na `hints.CheatsheetExcerpt` |

### Szczegóły implementacji

**exercise review** (linia ~691):
```go
// Before jsonOut(results):
addWeight(d, 4*len(results))
```

**exercise rubric** (linia ~3062):
```go
RunE: func(cmd *cobra.Command, args []string) error {
    d := db(cmd)  // NEW
    // ...
    addWeight(d, 1)  // NEW
    jsonOut(out)
```

**progress diagnose** (linia ~1484+):
```go
// Before jsonOut(out):
addWeight(d, 2)
```

**progress update** (linia ~872):
```go
// ADD before the walk_through check:
addWeight(d, 1)
// Keep existing walk_through weight:
if wynik == "walk_through" {
    addWeight(d, 5)  // existing — now total +6 for walk_through
}
```

**progress blad** (linia ~1396):
```go
// REPLACE conditional weight with always-add:
addWeight(d, 1)
if hint == 3 {
    addWeight(d, 2)  // extra weight for full hint chain (total +3)
}
```

**exercise hints** (linia ~550):
```go
hints := getExerciseHints(d, id, level)
if hints.CheatsheetExcerpt != "" {
    addWeight(d, 3)
} else {
    addWeight(d, 1)
}
```

## Wpływ na sesje

| Scenariusz | Teraz | Po zmianie |
|------------|-------|------------|
| Poprawna bez pomocy | 4/ćw → /clear po ~20 | 5/ćw → /clear po ~16 |
| 1 błąd + hint (typowa) | 5/ćw → /clear po ~16 | 7-9/ćw → /clear po ~9-11 |
| Walk_through (ciężka) | 15/ćw → /clear po ~5 | 20/ćw → /clear po ~4 |

## Testy

### main_test.go

1. **Aktualizacja `TestAutoWeight`**: Dodać scenariusze z nowymi wagami
2. **Nowy `TestWeightRealisticSession`**: Symulacja 8 ćwiczeń z mieszanką wyników → weryfikacja session_weight w zakresie 70-90
3. **Test warunkowej wagi hints**: hints z/bez cheatsheet_excerpt → różne wagi

### test_qa.sh

Nie wymaga zmian — Layer 5 (Go unit tests) automatycznie uruchomi nowe testy.

## Pliki do modyfikacji

1. `matura_informatyka_rozszerzona/analiza/cli/commands.go` — 6 punktów addWeight
2. `matura_informatyka_rozszerzona/analiza/cli/main_test.go` — nowe/zaktualizowane testy
3. Po zmianach: `./build.sh` (rebuild binaries + reimport matura.db)

## Weryfikacja

1. `cd analiza/cli && go test -v -run TestWeight`
2. `cd analiza/cli && go test -v` (wszystkie 13+ testów)
3. `./test_qa.sh --layer 5` (Go unit tests via QA suite)
4. `./test_qa.sh` (full suite — opcjonalnie)
