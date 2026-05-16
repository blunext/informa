# SQL Access-warning Implementation Plan

> **For agentic workers:** REQUIRED SUB-SKILL: Use superpowers:subagent-driven-development (recommended) or superpowers:executing-plans to implement this plan task-by-task. Steps use checkbox (`- [ ]`) syntax for tracking.

**Goal:** Dodac w SKILL.md skilla `matura` krotki pointer, ktory instruuje Claude'a (tutora) aby ostrzegal ucznia o konstrukcjach SQL niezgodnych z dialektem MS Access (JET/ACE), uzywanym na egzaminie maturalnym.

**Architecture:** Jeden edit w istniejacym pliku `/Users/blt1wz/priv/informa/.claude/skills/matura/SKILL.md` — dodanie ~9-liniowego bloku markdown po linii 310 (za lista pytan sokratejskich dla SQL, przed blokiem `**ARKUSZ:**`). Blok zawiera kontekst (matura=Access, cwiczenia=SQLite), polecenie sprawdzania, kilka przykladow-zahaczek (nie wyczerpujaca lista) i framing "to ostrzezenie, nie blad". Zero nowych plikow, zero zmian w kodzie CLI/Go, zero zmian w weryfikatorach.

**Tech Stack:** Markdown (SKILL.md w formie instrukcji dla LLM). Tekst w ASCII (zgodnie z konwencja istniejacego pliku — brak polskich diakrytykow).

**Spec reference:** `docs/superpowers/specs/2026-04-23-sql-access-warning-design.md`

**User constraints (from MEMORY.md):**
- NEVER run `git commit` — user commits manually. Plan NIE zawiera krokow git commit.

---

## Task 1: Dodanie bloku Access-warning do SKILL.md

**Files:**
- Modify: `/Users/blt1wz/priv/informa/.claude/skills/matura/SKILL.md:310-312`

**Context dla wykonawcy (bez zalozen):** Plik SKILL.md (582 linii) jest prompt-em dla Claude'a pelniacego role korepetytora matury. Sekcja E ("Prezentacja cwiczenia") zawiera podsekcje "Wzorce dekompozycji per kategoria" — tam sa bloki `**TEORIA:**`, `**IMPLEMENTACJA:**`, `**SQL:**`, `**ARKUSZ:**` z sokratejskimi pytaniami dla kazdej kategorii zadan. Wstawiamy nowy blok `**Access-warning (tylko dla typow SQL):**` miedzy koncem listy SQL (linia 310) a blokiem `**ARKUSZ:**` (linia 312), zachowujac istniejacy pusty wiersz jako separator.

- [ ] **Step 1: Zweryfikuj aktualny stan SKILL.md w okolicach insertion point**

Run:
```bash
sed -n '305,315p' /Users/blt1wz/priv/informa/.claude/skills/matura/SKILL.md
```

Expected output (jesli plik nie zostal zmieniony od czasu pisania specu 2026-04-23):
```
**SQL:**
1. "Ktore tabele sa potrzebne?"
2. "Jak je polaczyc? (JOIN, warunek ON)"
3. "Jakie warunki WHERE?"
4. "Czy trzeba grupowac? (GROUP BY + HAVING)"
5. "Co w SELECT? Jakie funkcje agregujace?"

**ARKUSZ:**
1. "Gdzie sa dane zrodlowe? (zakres komorek)"
2. "Jaka formula? (SUMIFS, COUNTIF, VLOOKUP?)"
3. "Jakie referencje? (bezwzgledne $ vs wzgledne)"
```

**Jesli output sie rozni** (np. ktos w miedzyczasie edytowal SKILL.md, zmienily sie numery linii, inne tresci bloku SQL): NIE wykonuj Step 2 na slepo. Zlokalizuj nowa pozycje bloku `**SQL:**` komenda `grep -n '^\*\*SQL:\*\*' SKILL.md` i dostosuj `old_string` w Step 2 do aktualnej tresci konca listy SQL + nastepujacej po nim pustej linii + poczatku bloku `**ARKUSZ:**`.

- [ ] **Step 2: Wstaw blok Access-warning przez tool Edit**

Wykonaj (tool Edit, nie bash sed):
- `file_path`: `/Users/blt1wz/priv/informa/.claude/skills/matura/SKILL.md`
- `old_string`:
```
5. "Co w SELECT? Jakie funkcje agregujace?"

**ARKUSZ:**
```
- `new_string`:
```
5. "Co w SELECT? Jakie funkcje agregujace?"

**Access-warning (tylko dla typow SQL):**
Matura odbywa sie w MS Access (dialekt JET/ACE), a cwiczenia sa weryfikowane
w SQLite — skladnie sie roznia. Gdy uczen pokazuje zapytanie, sprawdz czy
uzyte konstrukcje zadzialaja w Accessie. Typowe pulapki: `SUBSTR`/`LENGTH`
(w Accessie `Mid`/`Len`), `||` konkatenacja (`&`), `LIMIT` (`TOP`),
`COALESCE` (`Nz`), literaly dat `'...'` (`#...#`), JOIN 3+ tabel bez
nawiasow. Gdy znajdziesz niezgodnosc — pokaz wersje Access-kompatybilna
i krotko wyjasnij roznice. Nie traktuj tego jako blad w cwiczeniu — to
ostrzezenie o dialekcie, zapytanie ucznia moze byc semantycznie poprawne.

**ARKUSZ:**
```

**Uwaga o znakach:** uzywamy "pauz" (U+2014, `—`), jak reszta SKILL.md (sprawdz np. linia 319). Polskie znaki pisane bez diakrytykow (konwencja pliku). Myslnik w "Access-kompatybilna" to zwykly minus (U+002D).

- [ ] **Step 3: Zweryfikuj wynik przez ponowne przeczytanie**

Run:
```bash
sed -n '305,325p' /Users/blt1wz/priv/informa/.claude/skills/matura/SKILL.md
```

Expected output: blok SQL (5 linii), pusta linia, nowy blok `**Access-warning (tylko dla typow SQL):**` (9 linii tresci), pusta linia, `**ARKUSZ:**`.

Dodatkowa weryfikacja:
```bash
grep -c "^\*\*Access-warning" /Users/blt1wz/priv/informa/.claude/skills/matura/SKILL.md
```
Expected: `1`

```bash
wc -l /Users/blt1wz/priv/informa/.claude/skills/matura/SKILL.md
```
Expected: `593` (582 + 11 nowych linii: 9 tresci + 1 pusty wiersz + 1 blank trailing — dokladnie 11 z uwagi na zachowanie istniejacej pustej linii).

**Jesli liczba linii nie zgadza sie w granicach +/- 2:** sprawdz czy Edit nie zdublowal/zgubil pustej linii separujacej. Popraw manualnie jesli trzeba.

- [ ] **Step 4: NIE commituj**

Zgodnie z preferencja uzytkownika (MEMORY.md: "NEVER run `git commit` — user commits manually"). Zakoncz task bez `git add` / `git commit`. Poinformuj uzytkownika ze zmiana jest gotowa do przegladu przez `git diff .claude/skills/matura/SKILL.md`.

---

## Task 2: Smoke test w realnej sesji tutoringowej

**Files:**
- Nie zmienia plikow — uzytkownik testuje manualnie.

**Context dla wykonawcy:** Nie ma deterministycznego testu (konsekwencja wyboru opcji B — AI-based check). Testujemy przez realna sesje z tutorem. Ponizsze 3 scenariusze pokrywaja: (a) golden path — flagowanie oczywistej niezgodnosci, (b) regresja — brak falszywych alarmow, (c) edge case — mieszana poprawnosc.

**Wykonawca:** Przedstaw ponizsze 3 scenariusze uzytkownikowi, popros zeby przeszedl przez nie w osobnej sesji Claude Code (nowy tab / nowa rozmowa, zeby skill sie zaladowal od zera). Odbierz wyniki i poinformuj czy iterujemy czy konczymy.

- [ ] **Step 1: Przekaz uzytkownikowi scenariusz 1 (golden path — `SUBSTR`)**

Instrukcja do przekazania uzytkownikowi:

> Uruchom: `/matura SQL`
> Gdy tutor poprosi o rozwiazanie zadania SQL (typ 20-23), wpisz rozwiazanie uzywajace `SUBSTR(...)` lub `LENGTH(...)`. Przykladowe zapytanie (nawet jesli nie pasuje do zadania — chodzi o skladnie):
> ```sql
> SELECT SUBSTR(nazwisko, 1, 3) AS skrot FROM uczniowie WHERE LENGTH(imie) > 5;
> ```
>
> **Oczekiwana reakcja tutora:** Tutor powinien (a) ocenic zapytanie merytorycznie, (b) **dodatkowo** poinformowac ze `SUBSTR` i `LENGTH` nie zadzialaja w Accessie, i podac wersje z `Mid`/`Len`. Nie powinien traktowac tego jako bledu obnizajacego punktacje.
>
> **Czerwona flaga:** Tutor nie wspomina w ogole o Accessie / dialekcie / konwersji — wtedy trigger nie zadzialal.

- [ ] **Step 2: Przekaz uzytkownikowi scenariusz 2 (regresja — zapytanie czyste)**

Instrukcja do przekazania uzytkownikowi:

> W tej samej lub nowej sesji, przy kolejnym zadaniu SQL, wpisz zapytanie uzywajace wylacznie konstrukcji zgodnych z Accessem:
> ```sql
> SELECT imie, nazwisko FROM uczniowie WHERE klasa = '3A' ORDER BY nazwisko;
> ```
>
> **Oczekiwana reakcja tutora:** Ocenia zapytanie merytorycznie. **Nie** wymysla Access-warningu tam gdzie nie ma niezgodnosci. (Dopuszczalne: krotka wzmianka "to zadzialaloby tez w Accessie" — niedopuszczalne: proponowanie "popraw SELECT na TOP" albo inne sztuczne zmiany.)
>
> **Czerwona flaga:** Tutor generuje ostrzezenie tam gdzie nie ma niezgodnosci → instrukcja za agresywna, trzeba zlagodzic framing.

- [ ] **Step 3: Przekaz uzytkownikowi scenariusz 3 (edge case — JOIN 3+ tabel)**

Instrukcja do przekazania uzytkownikowi:

> Przy kolejnym zadaniu SQL (najlepiej typ `sql_join`), wpisz:
> ```sql
> SELECT u.imie, g.nazwa, s.tytul
> FROM uczniowie u INNER JOIN grupy g ON u.grupa_id = g.id INNER JOIN ksiazki s ON u.id = s.wlasciciel_id;
> ```
>
> **Oczekiwana reakcja tutora:** Powinien zflagowac **brak nawiasow** przy 3 tabelach i pokazac wersje Access: `FROM ((uczniowie u INNER JOIN grupy g ON ...) INNER JOIN ksiazki s ON ...)`. Dodatkowo moze wspomniec ze w SQLite/standardzie SQL jest to OK — ale w Accessie wymagane sa nawiasy.
>
> **Czerwona flaga:** Tutor pomija brak nawiasow → zahaczka "JOIN 3+ tabel bez nawiasow" w SKILL.md nie wystarcza modelowi, trzeba ja rozszerzyc (np. dodac krotki przyklad w samej instrukcji).

- [ ] **Step 4: Zbierz feedback i zdecyduj**

Po tym jak uzytkownik przeprowadzi wszystkie 3 scenariusze, zbierz wyniki w formacie:

| Scenariusz | Wynik | Notatka |
|---|---|---|
| 1. SUBSTR/LENGTH | PASS / FAIL / CZESCIOWY | ... |
| 2. Czyste zapytanie | PASS / FAIL / CZESCIOWY | ... |
| 3. JOIN 3+ tabel | PASS / FAIL / CZESCIOWY | ... |

**Decyzja:**
- **3x PASS** → done, feature dziala. Koniec planu.
- **Scenariusz 1 FAIL** → trigger nie dziala. Sprawdz: czy Claude widzi blok Access-warning w SKILL.md (czy sesja sie odswiezyla)? Jesli tak — wzmocnij framing ("ZAWSZE sprawdz Access-zgodnosc dla typow SQL, nawet jesli zapytanie wyglada poprawnie semantycznie").
- **Scenariusz 2 FAIL (false positive)** → zlagodz framing ("Tylko gdy widzisz konkretna niezgodnosc — nie wymyslaj problemow").
- **Scenariusz 3 FAIL** → dopisz konkretny przyklad nawiasowania do bloku (np. `FROM ((A JOIN B ON ...) JOIN C ON ...)`).

**Nie wyprodukuj szumu:** jesli wszystkie PASS, NIE rozszerzaj bloku "na zapas" — YAGNI. Czekaj na realny sygnal z uzytkowania.

---

## Self-review checklist (post-plan)

Wykonano w trakcie pisania planu:

- [x] **Spec coverage:** Wszystkie wymagania specu pokryte — (1) lokalizacja w SKILL.md → Task 1 Step 2; (2) tresc bloku → Task 1 Step 2 `new_string`; (3) brak zmian w CLI/weryfikatorach → Task 1 jest jedyna zmiana plikowa; (4) trigger dla typow SQL → w tresci bloku "(tylko dla typow SQL)" + lokalizacja w podsekcji SQL; (5) framing "ostrzezenie, nie blad" → ostatnie zdanie bloku; (6) smoke test → Task 2 (scenariusze 1-3).
- [x] **Placeholder scan:** Brak TBD/TODO. Wszystkie `old_string`/`new_string` sa pelne. Scenariusze testowe maja konkretne zapytania i oczekiwane zachowania.
- [x] **Type consistency:** N/A (edycja markdown, nie kod).
- [x] **Git constraint:** Task 1 Step 4 wyraznie zabrania `git commit`. Brak krokow commitujacych w planie.
- [x] **ASCII compliance:** Blok Access-warning napisany w ASCII Polish (bez diakrytykow), zgodnie z konwencja SKILL.md (potwierdzone przez `grep -c '[ąćęłńóśźż]' SKILL.md` → 2 na caly plik, czyli konwencja = ASCII).
