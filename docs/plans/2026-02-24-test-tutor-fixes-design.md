# Test-Tutor Fixes — Design

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Fix 3 top issues from test-tutor report (89.5/100, commit 07cb9f3) + 1 probna.md bug.

**Architecture:** Minimal targeted edits in 3 files. No new files, no refactoring.

**Files touched:**
- `.claude/skills/matura/SKILL.md` (fixes 1, 2)
- `.claude/skills/matura/probna.md` (fix 4)
- `.claude/skills/test-tutor/SKILL.md` (fix 3)

---

## Fix 1: Socratic question + hint separation (SKILL.md F.4)

**Problem:** Tutor combines Socratic question and hint in one message instead of waiting for student response. Reported in hint_progression (86/100), review_session (95/100). Persistent across 3+ runs.

**Root cause:** SKILL.md F.4 lines 312-315 say "NAJPIERW zapytaj... POTEM: wskazowki" but doesn't explicitly forbid combining them. LLM collapses two-step sequence into one message.

**Fix:** Insert explicit prohibition after line 311:
```
⚠ NIGDY nie lacZ pytania sokratejskiego z wskazowka w jednej wiadomosci.
Sekwencja to 2 OSOBNE wiadomosci: (1) pytanie → czekaj na odpowiedz → (2) hint.
```

**Validation:** Re-run `/test-tutor --scenario hint_progression` — socratic score should be 5/5.

---

## Fix 2: krok-po-kroku gate check (SKILL.md E)

**Problem:** Step-by-step mode not activated for exercise 1.11 (sledzenie_algorytmu, srednie-trudne). Reported in cke_unlock (85/100).

**Root cause:** Section E defines krok-po-kroku as a subsection (lines 260-275) but the main exercise presentation flow doesn't cross-reference it. Tutor skips it because it's not in the main checklist path.

**Fix:** Add gate check in section E main flow, before exercise presentation (around line 245):
```
**[GATE]** Przed prezentacja cwiczenia sprawdz:
- Jesli `trudnosc` >= `srednie-trudne` ORAZ `typ` in (`sledzenie_algorytmu`, `projektowanie_algorytmu`):
  → Przejdz do "Tryb krok-po-kroku" ponizej. NIE mow "Podaj rozwiazanie".
```

**Validation:** Re-run `/test-tutor --scenario cke_unlock` — krok-po-kroku checkpoint should PASS.

---

## Fix 3: Probna extended script (test-tutor SKILL.md)

**Problem:** 4-exchange fixed script structurally prevents reaching exam save and summary. 2 checkpoints always FAIL. Probna avg 82.9/100 — lowest scenario.

**Root cause:** Script too short to cover full probna flow (start → tasks → przerwij → summary → save).

**Fix:** Extend to 8 exchanges:
```
wymiana_1_uczen: "probna 2023"
wymiana_2_uczen: "[poprawna — zad 1.1, pelne punkty]"
wymiana_3_uczen: "[bledna — zad 1.2, off_by_one]"
wymiana_4_uczen: "[czesciowo poprawna — zad 2.1, 50% punktow]"
wymiana_5_uczen: "pomin"  (zad 2.2 — 0 pkt)
wymiana_6_uczen: "[poprawna — zad 3.1]"
wymiana_7_uczen: "przerwij"
wymiana_8_uczen: "tak"  (potwierdzenie zakonczenia)
```

Updated checkpoints (11 total, was 9):
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
[ ] Obsluga "pomin" (0 pkt, przejscie dalej)
[ ] Obsluga "przerwij" (koniec → podsumowanie)
```

**Validation:** Re-run `/test-tutor --scenario probna` — exam_save and podsumowanie should PASS.

---

## Fix 4: probna.md --hint bug

**Problem:** Line 40 of probna.md has `progress blad` without `--hint N`. CLI rejects calls without it.

**Fix:** Change line 40 from:
```
./matura progress blad --exercise-id {rok}.{zad}.{podzad} --typ {typ} --kod {kod}
```
to:
```
./matura progress blad --exercise-id {rok}.{zad}.{podzad} --typ {typ} --kod {kod} --hint 0
```

**Validation:** Lint check in test_qa.sh Layer 3 should pass.

---

## Validation plan

After all fixes:
1. `./test_qa.sh --layer 3` — SKILL lint (checks SKILL.md + probna.md + test-tutor)
2. `/test-tutor` full run — expect overall >= 90/100, probna >= 85/100
