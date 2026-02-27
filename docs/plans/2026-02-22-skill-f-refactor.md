# SKILL.md sekcja F refaktor + probna.md poprawki

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Poprawic 5 problemow zidentyfikowanych przez test-tutor (84.1→86+) przez refaktor sekcji F na checklist i dodanie 3 regul do probna.md.

**Architecture:** Chirurgiczne edycje 2 plikow markdown (SKILL.md, probna.md). Sekcja F zostaje przebudowana: podsekcje "System hintow" i "Przebieg hintow" zastapione jednym checklistem. Reszta sekcji F bez zmian. probna.md dostaje 3 nowe reguly w istniejacych sekcjach.

**Tech Stack:** Markdown only. Weryfikacja: test_qa.sh --layer 3 (SKILL lint) + test-tutor.

---

### Task 1: Dodaj 3 reguly do probna.md

**Files:**
- Modify: `.claude/skills/matura/probna.md:13-14` (sekcja Start)
- Modify: `.claude/skills/matura/probna.md:34-36` (sekcja Przebieg)

**Step 1: Dodaj `progress status` w sekcji Start**

W sekcji "## Start", przed punkt 1 ("Pobierz metadane"), dodaj nowy punkt 0:

```markdown
0. **[WYMAGANE]** Sprawdz status: `./matura progress status` (SKILL.md C wymaga ZAWSZE)
```

Linia 14 (obecnie `1. Pobierz metadane`) staje sie poprzedzona nowym punktem 0.

**Step 2: Dodaj zakaz exercise answer + regule progress blad w sekcji Przebieg**

W sekcji "## Przebieg", po punkcie 4 ("Ocen wg zasady_oceniania...") i przed "Komendy w trakcie:", dodaj:

```markdown
5. **[NIGDY]** nie uzywaj `exercise answer` na ID egzaminowych (YYYY.N.M).
   Oceniaj WYLACZNIE wg `zasady_oceniania` z odpowiedzi `exam task`.
6. **Jesli uczen popelni blad** — zapisz:
   `./matura progress blad --exercise-id {rok}.{zad}.{podzad} --typ {typ} --kod {kod}`
```

**Step 3: Zweryfikuj probna.md**

Run: `cat -n .claude/skills/matura/probna.md | head -50`
Expected: punkt 0 (progress status) widoczny przed "Pobierz metadane", punkty 5-6 widoczne po "Ocen wg zasady_oceniania"

---

### Task 2: Zamien sekcje "System hintow" i "Przebieg hintow" na CHECKLIST

**Files:**
- Modify: `.claude/skills/matura/SKILL.md:303-361` (usuniecie 2 podsekcji, wstawienie checklisty)

**Step 1: Zamien linie 303-361 na nowa podsekcje CHECKLIST**

Usun obecne podsekcje:
- "### System hintow (coaching-driven)" (linie 303-327)
- "#### Przebieg hintow (3 poziomy)" (linie 329-361)

Wstaw w ich miejsce:

```markdown
### CHECKLIST — po odpowiedzi ucznia

**[WYMAGANE]** Wykonaj kroki 1-8 w tej kolejnosci:

**1. Pobierz odpowiedz** (lazy — DOPIERO teraz, nie wczesniej):
   `./matura exercise answer --id {id}` → `odpowiedz` + `typowe_bledy[]`

**2. Porownaj** odpowiedz ucznia z polem `odpowiedz`.
   Uwzglednij rownowazne formy (alias SQL, kolejnosc kolumn). Czesciowo poprawna → potwierdz co dobrze, naprowadz na reszte.

**3. Jesli POPRAWNA** → przejdz do kroku 8.

**4. Jesli BLEDNA** — dla KAZDEGO bledu osobno, PRZED czymkolwiek innym:
   `./matura progress blad --exercise-id {id} --typ {typ} --kod {kod}`
   Wiele bledow w jednej odpowiedzi = wiele osobnych komend `progress blad`.

**5. Sprawdz `coaching.hint_delay` vs numer proby ucznia:**

   **proba < hint_delay** → pytanie sokratejskie (BEZ hintow, BEZ `exercise hints`):
   - `hint_delay=1` (new/learning): hint od 1. proby — przejdz nizej
   - `hint_delay=2` (familiar): 1. proba = pytanie sokratejskie, od 2. → hint
   - `hint_delay=3` (mastered): 1-2. proba = pytanie sokratejskie, od 3. → hint
   - Przy **PIERWSZYM** cwiczeniu z hint_delay >= 2, poinformuj:
     * hint_delay=2: "Od teraz mniej podpowiedzi — rozwijasz samodzielnosc."
     * hint_delay=3: "Bez podpowiedzi — jak na prawdziwym egzaminie."

   **proba >= hint_delay** → pobierz hinty i podaj wskazowke:
   - **[NIGDY]** nie podawaj wskazowki bez `exercise hints --id {id}`
   - `./matura exercise hints --id {id}` → `wskazowki[]`, `max_hints`
   - Dopiero PO pobraniu mozesz podac wskazowke z `wskazowki[]`

   **Poziom 1** (proba == hint_delay):
   - Dodaj `--hint 1` do `progress blad` (z kroku 4)
   - Jesli `max_hints >= 1` → podaj `wskazowki[0]` + pytanie sokratejskie
   - Jesli `max_hints == 0` → tylko pytanie sokratejskie

   **Poziom 2** (nastepna bledna proba):
   - `progress blad ... --hint 2`
   - **ZAWSZE** pobierz i **ZACYTUJ** fragment cheatsheet:
     `./matura cheatsheet get --kategoria {kat} --sekcja "{temat}"`
     Mapowanie: mod/div→"archetyp", rekurencja→"rekurencj", zlozonosc→"zlozonosc",
     JOIN→"join", GROUP BY→"group", sortowanie→"sort", adresowanie→"adresow",
     szyfrowanie→"bezpieczen", P/F→"prawda", konwersja→"konwersj"
   - Jesli `max_hints >= 2` → podaj cytat + `wskazowki[1]`

   **Poziom 3** (nastepna bledna proba):
   - `progress blad ... --hint 3`
   - Jesli `max_hints >= 3` → podaj `wskazowki[2]` (kluczowy krok)
   - Rozpisz rozwiazanie krok po kroku, ostatni krok zostaw uczniowi

**6. Walk_through** resetuje poziom do "new" → hint_delay wraca do 1.

**7. Po 3 probach / "poddaje sie"** → wynik = `walk_through`:
   - `./matura exercise answer --id {id}` (jesli nie pobrana w kroku 1)
   - Wyswietl pelna `odpowiedz` + `typowe_bledy` jako wskazowki CKE
   - **[WYMAGANE] Konsolidacja**: "Wyjasniej swoimi slowami, dlaczego to rozwiazanie dziala."
     * Poprawne → krotki pozytywny feedback
     * Bledne → doprecyzuj (2-3 zdania)
     * `dalej`/`nastepny` → pomin (TYLKO na wyrazna prosbe)
   - **[WYMAGANE] Wizualizacja** (typy: sledzenie, projektowanie, analiza, konwersja, bezpieczenstwo):
     narysuj ASCII diagram (tabelka, drzewo, schemat, kolumna dzielenia, wykres)

**8. Zapis wyniku** (WYMAGANE po kazdym cwiczeniu):
   ```
   ELAPSED=$(($(date +%s) - START_TS))
   ./matura progress update --id {id} --wynik {wynik} --czas $ELAPSED
   ```
   - Jesli `blad_warning` w odpowiedzi → `progress blad` natychmiast
   - Jesli `lapses >= 3` → "Ten temat ({tag}) sprawia Ci trudnosc juz po raz {lapses}."
   - Jesli `feedback_czasowy` niepuste → wyswietl uczniowi doslownie
```

**Step 2: Usun zduplikowana podsekcje "Lazy loading"**

Usun obecne linie 281-287 ("### Lazy loading — NIE pobieraj odpowiedzi z gory" + tresc) — ta informacja jest teraz w krokach 1 i 5 checklisty.

**Step 3: Usun zduplikowana podsekcje "Zapis wyniku"**

Zmniejsz obecna podsekcje "### Zapis wyniku" (linie 403-426) do:

```markdown
### Zapis wyniku — szczegoly CLI

CLI automatycznie po `progress update`:
- Zapisuje cwiczenie jako zrobione (z czasem)
- Aktualizuje streak i poziom trudnosci
- Oblicza nastepne daty powtorkowe (FSRS-5 — adaptacyjne interwaly per tag)
- Zwraca nowy poziom, streak, zaktualizowane tagi
- Zwraca `stability` (sila zapamietania tagu) i `lapses` (ile razy tag wypadl)
- Zwraca `feedback_czasowy` — gotowy tekst do wyswietlenia
```

---

### Task 3: Weryfikacja — test_qa.sh layer 3

**Files:** (read-only)
- `matura_informatyka_rozszerzona/analiza/test_qa.sh`

**Step 1: Uruchom SKILL lint**

Run: `cd matura_informatyka_rozszerzona/analiza && ./test_qa.sh --layer 3`
Expected: All PASS — zaden nowy command nie zostal dodany, zaden istniejacy usuniety.

**Step 2: Sprawdz line count**

Run: `wc -l .claude/skills/matura/SKILL.md .claude/skills/matura/probna.md`
Expected: SKILL.md ~590 linii (bylo 586, +checklist -2 podsekcje), probna.md ~90 linii (bylo 81, +9)

---

### Task 4: Weryfikacja — test-tutor (pelny run)

**Step 1: Uruchom test-tutor**

Run: `/test-tutor` (w nowej sesji Claude Code)

**Step 2: Porownaj wyniki**

Expected improvements vs 22c3ebe (84.1):
- integralnosc_cli avg >= 4.0 (bylo 3.86)
- coaching avg >= 4.0 (bylo 3.86)
- probna score >= 80 (bylo 76)
- first_session score >= 78 (bylo 72)
- Overall >= 86

**Step 3: Jesli regresja w hint_progression lub difficulty_climb**

Jesli hint_progression < 95 lub difficulty_climb < 90 → rewertuj checklist i przeanalizuj co poszlo nie tak.

---

## Notatki implementacyjne

- **Kolejnosc edycji**: probna.md (task 1) → SKILL.md (task 2) → weryfikacja (tasks 3-4)
- **Nie ruszaj**: sekcji A, A2, B, C, D, E, H, I w SKILL.md
- **Nie ruszaj**: sekcji "Punktacja czesciowa", "Wizualizacje", "Kody bledow", "Dobor kodu bledu", "Proaktywna detekcja" — zostaja po checklistie jako referencja
- **SKILL.md ma 586 linii** — checklist dodaje ~70 linii ale usuwamy ~60 (System hintow + Przebieg hintow + Lazy loading + Zapis wyniku skrocony), netto ~+10 linii
