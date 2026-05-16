# SQL Access-warning — design spec

**Data:** 2026-04-23
**Status:** Draft — do akceptacji przez użytkownika
**Kontekst:** Rozmowa brainstormingowa (Claude Code, sesja 2026-04-23)

## Problem

Ćwiczenia SQL (typy 20-23, ~230 pozycji) w skillu `matura` są weryfikowane przez `sql_verifier.py` wykonujące zapytania w Pythonowym `sqlite3` (in-memory). Zapytania-wzorce w polu `odpowiedz` używają składni SQLite-owej (`SUBSTR`, `LENGTH`, `||`, `LIMIT`, `COALESCE`).

Na maturze rozszerzonej z informatyki CKE typowym środowiskiem bazodanowym jest **Microsoft Access (JET/ACE dialect)**. Access nie obsługuje wielu konstrukcji używanych w ćwiczeniach, m.in.:

- Funkcje tekstowe: `SUBSTR` → `Mid`, `LENGTH` → `Len`, `SUBSTRING` → `Mid`
- Operatory: `||` → `&`, dzielenie całkowite `/` → `\`
- Słowa kluczowe: `LIMIT N` → `TOP N`, `COALESCE` → `Nz`, `MOD` (funkcja) → `Mod` (operator)
- Składnia: JOIN ≥3 tabel wymaga nawiasów, literały dat `'YYYY-MM-DD'` → `#YYYY-MM-DD#`
- Brak: `FULL OUTER JOIN`, funkcje okienkowe, CTE w starszych wersjach

Skutek: uczeń, który nauczył się pisać zapytania na istniejących ćwiczeniach, może na maturze dostać „Undefined function 'SUBSTR'" i stracić punkty za nieznajomość dialektu — a nie za błąd algorytmiczny.

## Rozwiązanie

**Opcja B** (wybrana po brainstormingu): live-check w trakcie sesji tutoringowej, realizowany przez AI (Claude czytający SKILL.md). Gdy uczeń pokazuje zapytanie SQL, tutor sprawdza czy użyte konstrukcje są Access-kompatybilne i — jeśli nie — pokazuje wersję dla Accessa z krótkim wyjaśnieniem.

Rozwiązanie **nie** obejmuje:
- Refactoringu istniejących 230 zapytań-wzorców (opcja A — odrzucona jako zbyt kosztowna względem ROI)
- Zmiany weryfikatora `sql_verifier.py` (pozostaje SQLite)
- Kodu w CLI Go (opcja γ/hybrid — odrzucona jako over-engineering dla szybkiego wina)

## Architektura

**Inline w SKILL.md, bez nowych plików.** Jedna sekcja ~5-7 linii markdown.

### Uzasadnienie rezygnacji z osobnego pliku

Początkowa propozycja (osobny `access_dialekt.md` z 5 kategoriami reguł, procedurą tutora, szablonami odpowiedzi — ~80 linii) została odrzucona. Powody:

1. **Dla opcji B zaufanie do modelu jest założeniem bazowym.** Jeśli nie ufamy wiedzy modelu o dialekcie Access, to nie powinniśmy wybierać opcji B, tylko γ (deterministyczny CLI).
2. **Rozdęte reguły zjadają tokeny i maintenance bez realnej wartości.** Model zna dialekt JET — wyliczanie 15 par `X → Y` tylko duplikuje jego wiedzę.
3. **Skill-creator wprost ostrzega:** *"Today's LLMs are smart. They have good theory of mind (...) if you find yourself writing ALWAYS or NEVER in all caps, or using super rigid structures, that's a yellow flag — reframe and explain the reasoning."*
4. **YAGNI.** Jeśli minimalna wersja coś przegapi, iteracja = dopisanie jednego słowa-zahaczki. Maintenance 80-liniowej referencji to przeciwnie — ciężki.

### Co zostaje

Krótki pointer w SKILL.md (~5-7 linii) zawierający:
- **Kontekst:** że matura używa Accessa a ćwiczenia są weryfikowane w SQLite
- **Polecenie:** gdy uczeń pokaże SQL, sprawdź Access-zgodność
- **Kilka przykładów-zahaczek** (`SUBSTR`, `||`, `LIMIT`, brak nawiasów w JOIN-ach, `COALESCE`) — NIE jako wyczerpująca lista, tylko jako zaczepki aktywujące pełną wiedzę modelu
- **Framing:** to nie jest błąd w ćwiczeniu, to ostrzeżenie pedagogiczne (zapytanie ucznia może być semantycznie OK w SQLite)
- **Trigger:** tylko w sesjach z typami SQL (20-23)

### Lokalizacja w SKILL.md

Sekcja **E. Prezentacja ćwiczenia**, podsekcja „Wzorce dekompozycji per kategoria", ma już blok `**SQL:**` (linie 305-310) z pytaniami sokratejskimi dla zadań SQL. Access-warning dodajemy **bezpośrednio za** tym blokiem — przed następnym blokiem `**ARKUSZ:**`. Logika: kiedy tutor pracuje nad zadaniem SQL, ma wszystkie wskazówki w jednym miejscu.

## Propozycja treści (do review w kroku implementacji)

```markdown
**Access-warning (tylko dla typów SQL):**
Matura odbywa się w MS Access (dialekt JET/ACE), a ćwiczenia są weryfikowane
w SQLite — składnie się różnią. Gdy uczeń pokazuje zapytanie, sprawdź czy
użyte konstrukcje zadziałają w Accessie. Typowe pułapki: `SUBSTR`/`LENGTH`
(w Accessie `Mid`/`Len`), `||` konkatenacja (`&`), `LIMIT` (`TOP`),
`COALESCE` (`Nz`), literały dat `'...'` (`#...#`), JOIN 3+ tabel bez
nawiasów. Gdy znajdziesz niezgodność — pokaż wersję Access-kompatybilną
i krótko wyjaśnij różnicę. Nie traktuj tego jako błąd w ćwiczeniu — to
ostrzeżenie o dialekcie, zapytanie ucznia może być semantycznie poprawne.
```

## Zakres zmian (w kolejnym kroku — implementacja)

Jeden edit:
- **Plik:** `/Users/blt1wz/priv/informa/.claude/skills/matura/SKILL.md`
- **Po linii 310** (koniec listy `**SQL:**`)
- **Treść:** blok powyżej, ~7 linii

Nie zmieniamy:
- `pulapki.md`, `probna.md`
- Żadnego pliku JSON z ćwiczeniami
- `sql_verifier.py`
- CLI (Go)
- `test_qa.sh`

## Testowanie i iteracja

Brak deterministycznego testu (konsekwencja wyboru opcji B). Iteracja manualna:

1. **Smoke test:** sesja z ćwiczeniem SQL (np. `/matura SQL`), uczeń wprowadza zapytanie z `SUBSTR` — sprawdzamy czy tutor flaguje to jako Access-warning.
2. **Regression test:** zapytanie bez niezgodności (czyste `SELECT ... WHERE ...`) — tutor NIE powinien wymyślać nieistniejącego problemu.
3. **Edge case:** zapytanie mieszane (poprawne agregacje + JOIN 3 tabel bez nawiasów) — tutor powinien złapać tylko JOIN.

Jeśli w trakcie testów tutor przegapi konstrukcję, której nie ma w „zahaczkach" (np. funkcje okienkowe) — dopisujemy jedno słowo-zahaczkę do bloku. Iteracja tania.

## Tradeoffs — co akceptujemy

- ✅ Szybki win, ~5 linii, zero kodu, zero nowych plików, zero maintenance
- ✅ Zgodne z istniejącą filozofią skilla (CLI dla twardych reguł, markdown dla pedagogiki)
- ✅ Odwracalne: w razie problemów po prostu usuwamy sekcję
- ⚠️ Brak gwarancji że model złapie 100% niezgodności — zdajemy się na jego wiedzę (założenie opcji B)
- ⚠️ Brak formalnego testu — nie da się wpleść do `test_qa.sh` Layer 6 (deterministyczne journeys)
- ⚠️ Nie rozwiązuje problemu „ćwiczenia uczą SQLite a egzamin jest w Accessie" u źródła — tylko łagodzi skutki. Pełny refactor (opcja A) nadal opcjonalny do rozważenia w przyszłości.

## Następny krok

Po akceptacji specu — invokacja skilla `superpowers:writing-plans` do wygenerowania planu implementacyjnego (który dla tej skali = jeden edit + manualny smoke test).
