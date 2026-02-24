# Test-Tutor Fixes Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Fix 3 top issues from test-tutor report + 1 probna.md bug to raise overall score from 89.5 to 90+.

**Architecture:** Targeted edits in 3 skill files. No new files, no code changes, no refactoring.

**Tech Stack:** Markdown skill files, bash for validation.

---

### Task 1: Add hint separation warning (SKILL.md F.4)

**Files:**
- Modify: `.claude/skills/matura/SKILL.md:311`

**Step 1: Add explicit prohibition after line 311**

In `.claude/skills/matura/SKILL.md`, find the line:

```
   - Jesli hinty dostepne → podaj nastepna wskazowke z `wskazowki[]`:
```

Insert the following **immediately after** that line (before the `* **Poziom 1**` line):

```
     **WAZNE: NIGDY nie lacZ pytania sokratejskiego z wskazowka w jednej wiadomosci.**
     Sekwencja to 2 OSOBNE wiadomosci: (1) pytanie → czekaj na odpowiedz ucznia → (2) wskazowka.
```

**Step 2: Verify edit**

Read `.claude/skills/matura/SKILL.md` lines 311-325 and confirm the warning appears between line 311 and the `Poziom 1` bullet.

---

### Task 2: Add krok-po-kroku gate check (SKILL.md E)

**Files:**
- Modify: `.claude/skills/matura/SKILL.md:243`

**Step 1: Replace "Podaj swoje rozwiazanie" with gate check**

In `.claude/skills/matura/SKILL.md`, find the line:

```
3. Popros: "Podaj swoje rozwiazanie."
```

Replace it with:

```
3. **[GATE]** Przed prosba o rozwiazanie sprawdz:
   - Jesli `trudnosc` >= `srednie-trudne` ORAZ `typ` in (`sledzenie_algorytmu`, `projektowanie_algorytmu`):
     → Przejdz do "Tryb krok-po-kroku" ponizej (sekcja E). NIE mow "Podaj rozwiazanie".
   - W przeciwnym razie: "Podaj swoje rozwiazanie."
```

**Step 2: Verify edit**

Read `.claude/skills/matura/SKILL.md` lines 240-270 and confirm the gate check appears at step 3, before the krok-po-kroku subsection.

---

### Task 3: Fix probna.md --hint bug

**Files:**
- Modify: `.claude/skills/matura/probna.md:40`

**Step 1: Add --hint 0 to progress blad template**

In `.claude/skills/matura/probna.md`, find the line:

```
   `./matura progress blad --exercise-id {rok}.{zad}.{podzad} --typ {typ} --kod {kod}`
```

Replace with:

```
   `./matura progress blad --exercise-id {rok}.{zad}.{podzad} --typ {typ} --kod {kod} --hint 0`
```

**Step 2: Verify edit**

Read `.claude/skills/matura/probna.md` line 40 and confirm `--hint 0` is present.

---

### Task 4: Extend probna scenario script (test-tutor SKILL.md)

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md:319-345`

**Step 1: Replace probna fixed script and checkpoints**

In `.claude/skills/test-tutor/SKILL.md`, find the section `### 5.6 probna` (line 319) and replace everything from line 319 through line 345 (the closing ` ``` ` of Binary checkpoints) with:

```
### 5.6 probna

**Fixed student script:**
```
wymiana_1_uczen: "probna 2023"
wymiana_2_uczen: "[poprawna odpowiedz — zadanie 1.1, pelne punkty]"
wymiana_3_uczen: "[bledna odpowiedz — zadanie 1.2, off_by_one]"
wymiana_4_uczen: "[czesciowo poprawna — zadanie 2.1, 50% punktow]"
wymiana_5_uczen: "pomin"
wymiana_6_uczen: "[poprawna odpowiedz — zadanie 3.1]"
wymiana_7_uczen: "przerwij"
wymiana_8_uczen: "tak"
```

**Binary checkpoints:**
```
CLI compliance:
[ ] exam meta --rok 2023
[ ] exam task --rok 2023 --zadanie N (per zadanie: 1, 2, 3)
[ ] START_TS i ELAPSED
[ ] progress blad przy bledzie (wymiana_3) z --hint 0
[ ] exam save --rok 2023 --results '[...]' --czas M

Coaching reaction:
[ ] coaching_actions (N/A — egzamin)

Scenario-specific:
[ ] Brak hintow (tryb egzaminacyjny)
[ ] Podsumowanie per-zadanie + per-kategoria
[ ] Wyswietlenie zasad egzaminu
[ ] Obsluga "pomin" (0 pkt, przejscie do nastepnego podzadania)
[ ] Obsluga "przerwij" (koniec egzaminu → podsumowanie)
```
```

**Step 2: Verify edit**

Read `.claude/skills/test-tutor/SKILL.md` lines 319-360 and confirm:
- Fixed script has 8 exchanges (wymiana_1 through wymiana_8)
- Checkpoints have 11 items including `pomin` and `przerwij`

---

### Task 5: Add krok-po-kroku checkpoint to cke_unlock (test-tutor SKILL.md)

**Files:**
- Modify: `.claude/skills/test-tutor/SKILL.md:311-316`

**Step 1: Add krok-po-kroku checkpoint**

In `.claude/skills/test-tutor/SKILL.md`, find the `Scenario-specific:` section of cke_unlock (line 311). Add a new checkpoint after `Ogloszenie formatu "=== SPRAWDZIAN TYPU ==="` (line 316):

```
[ ] Tryb krok-po-kroku dla srednie-trudne (sekcja E)
```

**Step 2: Verify edit**

Read `.claude/skills/test-tutor/SKILL.md` lines 311-320 and confirm 7 scenario-specific checkpoints (was 6, added krok-po-kroku).

---

### Task 6: Run validation

**Step 1: Run SKILL lint (test_qa.sh Layer 3)**

Run:
```bash
cd /Users/blt1wz/priv/informa/matura_informatyka_rozszerzona/analiza && ./test_qa.sh --layer 3
```

Expected: All Layer 3 tests PASS.

**Step 2: Verify all edits are consistent**

Read the 3 modified files and confirm:
- SKILL.md F.4: warning present between line 311 and Poziom 1
- SKILL.md E.3: gate check present before "Podaj rozwiazanie"
- probna.md line 40: `--hint 0` present
- test-tutor SKILL.md 5.6 probna: 8 exchanges, 11 checkpoints
- test-tutor SKILL.md 5.5 cke_unlock: krok-po-kroku checkpoint present
