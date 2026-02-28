# Test-Tutor Findings Fix — Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Fix 3 issues found by test-tutor run 2026-02-28: progress blad retry logic, visualization per-category, leech-tag linking after error.

**Architecture:** Pure text edits in 2 skill files (SKILL.md, probna.md). No Go/CLI changes. ~10 lines added total.

**Tech Stack:** Markdown skill files only.

---

### Task 1: Strengthen `progress blad` retry logic in SKILL.md guardrails

**Files:**
- Modify: `.claude/skills/matura/SKILL.md:128`

**Step 1: Edit the guardrails section**

Replace line 128:
```
- `progress blad --kod Z` z niepoprawnym kodem → CLI zwroci JSON z `suggestions[]` (kody z opisami). Uzyj `suggestions[0].kod` jesli opis pasuje do bledu ucznia.
```

With:
```
- `progress blad --kod Z` z niepoprawnym kodem → CLI zwroci JSON z `suggestions[]` (kody z opisami). **Natychmiast** wywolaj `progress blad` ponownie z `suggestions[0].kod` (jesli opis pasuje) lub kolejna sugestia. NIE kontynuuj bez zapisania bledu.
```

**Step 2: Verify edit**

Read `.claude/skills/matura/SKILL.md` line 128 and confirm the new text is in place.

---

### Task 2: Strengthen `progress blad` retry logic in SKILL.md F.3

**Files:**
- Modify: `.claude/skills/matura/SKILL.md:342`

**Step 1: Edit the F.3 section**

Replace line 342:
```
   - CLI odrzuci niepoprawny kod i zwroci `suggestions[]` z opisami — uzyj pierwszej pasujacej sugestii
```

With:
```
   - CLI odrzuci niepoprawny kod i zwroci `suggestions[]` z opisami — **natychmiast** wywolaj ponownie z `suggestions[0].kod` (jesli opis pasuje) lub kolejna sugestia. NIE kontynuuj bez zapisania bledu.
```

**Step 2: Verify edit**

Read `.claude/skills/matura/SKILL.md` lines 340-345 and confirm.

---

### Task 3: Add leech-tag linking rule in SKILL.md F.3

**Files:**
- Modify: `.claude/skills/matura/SKILL.md:343`

**Step 1: Add the leech-tag linking rule**

After current line 343 (`Wiele bledow = wiele osobnych komend`), insert:
```
   - Jesli `--kod` jest zwiazany z tagiem z `coaching_actions` WARN_LEECH → powiaz explicite:
     "To ten sam problem z {tag}, o ktorym mowilismy na poczatku. Zwroc szczegolna uwage."
```

**Step 2: Verify edit**

Read `.claude/skills/matura/SKILL.md` lines 342-348 and confirm new lines are present.

---

### Task 4: Add visualization patterns for IMPLEMENTACJA, ARKUSZ, SQL

**Files:**
- Modify: `.claude/skills/matura/SKILL.md:394-395` (after teoria_bezpieczenstwa line, before example)

**Step 1: Add 3 category patterns**

After line 394 (`teoria_bezpieczenstwa: narysuj schemat ataku/obrony...`), insert:
```
- **IMPLEMENTACJA** (cyfry, napisy, ...): tabelka operacji mod/div na przykladowych danych LUB schemat petli/algorytmu
- **ARKUSZ**: schemat formuly z referencjami ($A$1 vs A1) LUB tabela krokow symulacji
- **SQL**: tabela wyniku zapytania (oczekiwana vs uzyskana) LUB schemat JOIN-ow
```

**Step 2: Verify edit**

Read `.claude/skills/matura/SKILL.md` lines 386-400 and confirm all 8 patterns are listed (5 TEORIA + 3 new).

---

### Task 5: Add retry cross-reference in probna.md

**Files:**
- Modify: `.claude/skills/matura/probna.md:40`

**Step 1: Add retry instruction after progress blad call**

Replace line 40:
```
   `./matura progress blad --exercise-id {rok}.{zad}.{podzad} --typ {typ} --kod {kod} --hint 0`
```

With:
```
   `./matura progress blad --exercise-id {rok}.{zad}.{podzad} --typ {typ} --kod {kod} --hint 0`
   Jesli CLI odrzuci kod (zwroci `suggestions[]`) — wywolaj ponownie z sugestia (patrz SKILL.md F.3).
```

**Step 2: Verify edit**

Read `.claude/skills/matura/probna.md` lines 39-42 and confirm.

---

### Task 6: Validate with SKILL lint (test_qa.sh L3)

**Step 1: Run SKILL lint**

Run: `cd /Users/blt1wz/priv/informa && ./matura_informatyka_rozszerzona/analiza/test_qa.sh --layer 3`

Expected: All SKILL lint checks PASS.

---
