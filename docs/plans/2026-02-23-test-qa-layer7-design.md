# Design: Layer 7 — Exercise Consistency Check

**Data**: 2026-02-23
**Cel**: Automatyczne sprawdzenie spójności danych po `/generate-exercises`

## Problem

Po generacji ćwiczeń nie ma jednego polecenia które sprawdzi:
- Czy JSON trafił poprawnie do matura.db (content sync)
- Czy ID są ciągłe (bez luk/duplikatów)
- Czy rozkład trudności jest sensowny

L2 sprawdza tylko count ≥407, L4 sprawdza schema/verification — ale nikt nie porównuje JSON↔DB per ćwiczenie.

## Rozwiązanie

### Część 1: Nowa komenda CLI `matura data verify`

Komenda Go w `commands.go` / `database.go`:

```
matura data verify --source ../cwiczenia/json
```

**Logika:**
1. Wczytaj wszystkie `_meta.json` + `{id}.json` z dysku (407 plików)
2. Wczytaj wszystkie ćwiczenia z attached `matura.db`
3. Porównaj per ćwiczenie: `id`, `trudnosc`, `punkty`, `tagi`, `tresc` (length+hash), `odpowiedz` (length+hash)
4. Raport JSON:
   ```json
   {
     "total_disk": 407,
     "total_db": 407,
     "matched": 407,
     "mismatched": [],
     "missing_in_db": [],
     "missing_on_disk": []
   }
   ```
5. Exit 0 = all OK, Exit 1 = rozbieżności

### Część 2: Layer 7 w test_qa.sh

3 checki w nowej funkcji `run_layer_7()`:

**7a. Data verify (CLI)**
```bash
"$MATURA" data verify --source "$SCRIPT_DIR/cwiczenia/json"
```
PASS jeśli exit 0, FAIL jeśli exit 1.

**7b. Ciągłość ID**
Dla każdego z 23 typów:
- Parsuj `_meta.json` → wyciągnij sekwencje M z id `N.M`
- Sprawdź: posortowane M = 1, 2, 3, ..., max (bez luk, bez duplikatów)
- PASS per typ, FAIL jeśli luka

**7c. Rozkład trudności**
Dla każdego z 23 typów:
- Policz ćwiczenia per trudność
- WARN jeśli typ ma <2 poziomy trudności
- Wyświetl rozkład

### Integracja

```bash
# Dispatch (dodaj na końcu default run):
run_layer_7

# --layer 7 support:
7) run_layer_7 ;;
```

Kolejność: po L6 (L7 wymaga zaimportowanej DB).

## Pliki do modyfikacji

1. `analiza/cli/commands.go` — nowa komenda `data verify`
2. `analiza/cli/database.go` — logika porównania JSON↔DB
3. `analiza/test_qa.sh` — nowa funkcja `run_layer_7()` + dispatch
4. `analiza/cli/main_test.go` — test dla `data verify`

## Weryfikacja

1. `cd analiza/cli && go build -o matura .`
2. `./matura data verify --source ../cwiczenia/json` → exit 0, matched=407
3. `cd analiza && ./test_qa.sh --layer 7` → all PASS
4. `./test_qa.sh` → 130+ tests, all PASS (no regression)
