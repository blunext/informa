# Design: Fix 3 Test-Tutor Findings (2026-02-28)

## Context

Test-tutor run on commit c339ce7 scored 87.7/100 (down from 95.6). Two scenarios regressed:
- **coaching_aware** (73/100): missing proactive visualization + missing leech-tag linking after error
- **probna** (77/100): `progress blad` uses invalid error codes, CLI rejects them, errors not recorded

## Changes

### Fix 1: `progress blad` retry logic (SKILL.md + probna.md)

**SKILL.md F.3** (~line 342): Strengthen existing "uzyj suggestions[0]" to explicit auto-retry:
```
- CLI odrzuci niepoprawny kod i zwroci `suggestions[]` z opisami — **natychmiast** wywolaj
  `progress blad` ponownie z `suggestions[0].kod` (jesli opis pasuje) lub `suggestions[1].kod`.
  NIE kontynuuj bez zapisania bledu.
```

**SKILL.md guardrails** (~line 128): Same strengthening.

**probna.md** (~line 40): Add after `progress blad` call:
```
Jesli CLI odrzuci kod (zwroci `suggestions[]`) — uzyj sugestii i wywolaj ponownie (patrz SKILL.md F.3).
```

### Fix 2: Visualization per-category (SKILL.md F.6)

**SKILL.md** (~line 404): Add after existing 5 TEORIA patterns:
```
- **IMPLEMENTACJA** (cyfry, napisy, ...): tabelka operacji mod/div na przykladowych danych LUB schemat petli/algorytmu
- **ARKUSZ**: schemat formuly z referencjami ($A$1 vs A1) LUB tabela krokow symulacji
- **SQL**: tabela wyniku zapytania (oczekiwana vs uzyskana) LUB schemat JOIN-ow
```

### Fix 3: Leech-tag linking after error (SKILL.md F.3)

**SKILL.md F.3** (~line 343): Add after `progress blad` instructions:
```
- Jesli blad (`--kod`) jest zwiazany z tagiem z `coaching_actions` WARN_LEECH — powiaz explicite:
  "To ten sam problem z {tag}, o ktorym mowilismy na poczatku. Zwroc szczegolna uwage."
```

## Files Modified

| File | Lines | Change |
|------|-------|--------|
| `.claude/skills/matura/SKILL.md` | ~128, ~342-343, ~404 | Strengthen retry, add leech link, add viz patterns |
| `.claude/skills/matura/probna.md` | ~40 | Add cross-reference to retry logic |

## Validation

After changes, run `/test-tutor` (at minimum `quick` or `--scenario coaching_aware`) to verify scores improve.
