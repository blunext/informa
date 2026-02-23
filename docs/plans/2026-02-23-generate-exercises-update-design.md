# Design: Aktualizacja `/generate-exercises` skill

**Data**: 2026-02-23
**Cel**: Uzupełnić skill o brakujące kroki i reorganizować sekcje po audycie

## Problem

Skill `/generate-exercises` nie był aktualizowany od dodania 97 ćwiczeń. Audyt znalazł 9 problemów w 3 kategoriach (3 krytyczne, 3 ważne, 3 drobne).

## Znalezione problemy

### Krytyczne
1. **Brak `zrodlo` w przykładowym JSON** — walidator wymaga, skill pomija
2. **Brak kroku CLI re-import** — po dodaniu ćwiczeń `matura.db` nieaktualny
3. **Brak info o field matching** — walidator wymusza zgodność `_meta.json` ↔ plik ćwiczenia

### Ważne
4. **Brak baseline update** — `test_qa.sh --update-baseline`
5. **Brak mapowania weryfikator↔katalog** — nie wiadomo co auto-verify
6. **Brak sekcji "Top błędy generacji"** — historia: 190 złych kara, 28 złych output

### Drobne
7. Podwójna numeracja `4.` w Kroku 2
8. Brak info: punkty 1-10, kara ułamkowa (`-0.5 pkt`)
9. Brak wzmianki o konwersjach (katalog 05)

## Podejście: Reorganizacja + łatki

### Zmiany per krok

| Krok | Zmiana |
|------|--------|
| 2 | Fix numeracji 4→5 |
| 3 | +`zrodlo` w przykładzie JSON, +punkty 1-10, +kara ułamkowa, +konwersje(05) |
| 3.5 | bez zmian |
| 4 | +warning: field matching _meta↔plik |
| 5 | **Połącz**: walidacja + re-import + baseline + tabela weryfikatorów |
| NEW | Sekcja "Typowe błędy generacji" (top 8) |
| 6 | +status re-import w podsumowaniu |

### Nowy Krok 5 (połączony)

```
## Krok 5: Walidacja + Re-import (OBOWIĄZKOWY)

### 5a. Schema + weryfikacja merytoryczna
python3 analiza/cwiczenia/validate_json.py --file NN_nazwa
python3 analiza/cwiczenia/verify/verify_all.py --file NN_nazwa --verbose

### Mapowanie weryfikatorów
| Katalogi | Weryfikator | Status |
|----------|-------------|--------|
| 01-04, 06 | manual_sanity | MANUAL_REVIEW |
| 05 | numconv | PASS/FAIL |
| 07-14 | cpp | PASS/FAIL/ERROR |
| 15-19 | manual_sanity | MANUAL_REVIEW |
| 20-23 | sql | PASS/FAIL/ERROR |

### 5b. Re-import do CLI
cd analiza/cli && ./matura data import --source ../

### 5c. Aktualizacja baseline
cd analiza && ./test_qa.sh --update-baseline
```

### Nowa sekcja: Typowe błędy generacji

Top 8 z historii:
1. Zły format `kara` — zawsze `-N pkt` lub `-N.N pkt`
2. Format `**Dane**` — MUSI być dokładnie `**Dane** (\`plik.txt\`):`
3. Notatki robocze w oczekiwanym wyniku — TYLKO czyste linie
4. SQL: ostatnia tabela w odpowiedzi = wynik (bez ✓/✗)
5. Tag spoza rejestru — dodaj do `tagi_rejestr.json` NAJPIERW
6. Niezgodność `_meta.json` ↔ plik — trudnosc/punkty/tagi muszą się zgadzać
7. Brak `zrodlo` — pole wymagane
8. Zły ID — format NN.M, NN = numer katalogu

## Plik do modyfikacji

- `.claude/skills/generate-exercises/SKILL.md`

## Weryfikacja

Po edycji: przeczytać skill od nowa, sprawdzić spójność kroków 1→2→3→3.5→4→5→6.
