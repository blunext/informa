# CLI Reliability Migration — Design

**Data:** 2026-02-24
**Cel:** Przenieść 3 kruche/niedeterministyczne logiki z AI (SKILL.md) do Go CLI, zwiększając niezawodność korepetytora.

## Kontekst

Analiza SKILL.md wykazała 15 możliwości konsolidacji/migracji logiki AI→CLI. Priorytet: niezawodność. Wybrano 3 zmiany o najwyższym stosunku zysk/koszt (łącznie ~80 linii Go):

| # | Zmiana | Podejście | Problem dziś |
|---|--------|-----------|-------------|
| 1 | Cheatsheet auto-attach | Wzbogacenie `exercise hints` | 12 ręcznych mapowań temat→sekcja w SKILL.md |
| 2 | Exercise rubric | Nowa komenda | Reguły punktacji hardcoded w SKILL.md |
| 3 | Error code suggestions | Wzbogacenie `progress blad` error | AI zgaduje "najbliższy" kod |

## Zmiana 1: Cheatsheet auto-attach w `exercise hints`

### Problem
SKILL.md utrzymuje 12 ręcznych mapowań temat→sekcja cheatsheet (`mod/div→"archetyp"`, `rekurencja→"rekurencj"`, `JOIN→"join"`, itd.). Jeśli zmieni się struktura cheatsheet, mapowania się łamią.

### Rozwiązanie
Gdy `exercise hints` zwraca Level 2 hint, CLI automatycznie przeszukuje cheatsheet odpowiedniej kategorii po tagach ćwiczenia i dołącza pasującą sekcję.

### Mechanizm
1. CLI zna `kategoria` ćwiczenia → wie który cheatsheet (TEORIA, IMPLEMENTACJA, ARKUSZ, SQL)
2. CLI zna `tagi` ćwiczenia (np. `["rekurencja", "fibonacci", "zlozonosc"]`)
3. CLI dzieli cheatsheet na sekcje po nagłówkach `##`
4. Dla każdego tagu szuka dopasowania w nagłówkach sekcji (case-insensitive substring)
5. Zwraca najlepszą sekcję jako nowe pole `cheatsheet_excerpt`

### Zmiana odpowiedzi
```json
// Dziś (Level 2):
{"id": "07_001", "wskazowki": ["...", "Sprawdź dzielenie mod 10"], "max_hints": 3}

// Po zmianie (Level 2):
{"id": "07_001", "wskazowki": ["...", "Sprawdź dzielenie mod 10"], "max_hints": 3,
 "cheatsheet_excerpt": "## Archetyp: cyfry liczby\nWczytaj → pętla while n>0 → ..."}
```

### Fallback
Jeśli żaden tag nie pasuje do żadnej sekcji → `cheatsheet_excerpt` = `null` → AI zachowuje się jak dziś.

### Zmiana SKILL.md
- Usunięcie 12 mapowań temat→sekcja
- Usunięcie instrukcji "ręcznie wywołaj cheatsheet get na Level 2"
- Dodanie: "Jeśli hints zwraca `cheatsheet_excerpt`, wyświetl go uczniowi"

---

## Zmiana 2: Nowa komenda `exercise rubric`

### Problem
Reguły punktacji częściowej (full/half/zero) per typ są hardcoded w SKILL.md. AI może je interpretować niespójnie między sesjami.

### Rozwiązanie
Nowa komenda `exercise rubric --typ {typ}` zwraca deterministyczne kryteria punktacji.

### Dane
Stała mapa w Go (nie wymaga DB) — 23 typy × 3 poziomy:

```go
var rubrics = map[string]Rubric{
    "sledzenie_algorytmu": {
        Full:  "Tabela poprawna + wynik końcowy poprawny",
        Half:  "Poprawny tok rozumowania, 1-2 błędy rachunkowe w wierszach",
        Zero:  "Błędny algorytm lub brak tabeli",
        Notes: "Każdy wiersz tabeli jest wart punkty",
    },
    "test_prawda_falsz": {
        Full:  "Poprawna odpowiedź P/F + poprawne uzasadnienie",
        Half:  "Poprawna odpowiedź P/F bez uzasadnienia",
        Zero:  "Błędna odpowiedź P/F",
        Notes: "Brak uzasadnienia = ZAWSZE max 50%",
    },
    // ... 21 remaining types
}
```

### Odpowiedź CLI
```json
{
  "typ": "sledzenie_algorytmu",
  "kategoria": "TEORIA",
  "rubric": {
    "full":  {"opis": "Tabela poprawna + wynik końcowy poprawny", "procent": 100},
    "half":  {"opis": "Poprawny tok, 1-2 błędy rachunkowe", "procent": 50},
    "zero":  {"opis": "Błędny algorytm lub brak tabeli", "procent": 0},
    "notes": "Każdy wiersz tabeli jest wart punkty"
  }
}
```

### Zmiana SKILL.md
- Usunięcie sekcji "Punktacja częściowa" z hardcoded regułami
- Dodanie: "Po ocenie odpowiedzi, wywołaj `exercise rubric --typ {typ}` i zastosuj zwrócone kryteria"

---

## Zmiana 3: Suggestions w `progress blad` error

### Problem
Gdy AI loguje błąd z nieprawidłowym kodem, CLI zwraca listę dozwolonych kodów, a AI musi wybrać "najbliższy" bez algorytmu.

### Rozwiązanie
Wzbogacenie odpowiedzi error o pole `suggestions` — top 2-3 najbliższe kody z opisami.

### Mechanizm
1. CLI już ma listę dozwolonych kodów per typ
2. Gdy `--kod` nie pasuje, CLI oblicza Levenshtein distance do każdego dozwolonego kodu
3. Sortuje po distance, bierze top 2-3
4. Zwraca `suggestions` z kodami i opisami (bez distance — AI nie potrzebuje liczby)

### Zmiana odpowiedzi error
```json
// Dziś:
{"error": "nieznany kod błędu", "valid_codes": ["off_by_one", "bledny_warunek", ...]}

// Po zmianie:
{
  "error": "nieznany kod błędu",
  "valid_codes": [
    {"kod": "off_by_one", "opis": "Błąd o jeden (indeks, granica pętli)"},
    {"kod": "bledny_warunek", "opis": "Zły warunek logiczny (if/while)"},
    {"kod": "brak_warunku_brzegowego", "opis": "Brak obsługi przypadku skrajnego"}
  ],
  "suggestions": [
    {"kod": "off_by_one", "opis": "Błąd o jeden (indeks, granica pętli)"},
    {"kod": "bledny_warunek", "opis": "Zły warunek logiczny (if/while)"}
  ]
}
```

### Zmiana SKILL.md
- Usunięcie instrukcji "wybierz najbliższy z listy"
- Dodanie: "Użyj `suggestions[0].kod` jeśli opis pasuje semantycznie do błędu ucznia"

---

## Przyszłe rozszerzenia (poza zakresem)

Dwa elementy odłożone na później (wymagają autorstwa treści edukacyjnych):

- **Dekompozycja krok-po-kroku** (`exercise decompose --typ --id`) — 23 typy × ~3 poziomy szablonów
- **Teoria na start** (`theory intro --typ`) — kuratorowane wprowadzenia do 6 typów TEORII

---

## Szacowany koszt

| Zmiana | Linie Go | Zmiana SKILL.md | Ryzyko regresji |
|--------|----------|-----------------|-----------------|
| Cheatsheet auto-attach | ~40 | Usunięcie 12 mapowań | Niskie (fallback na null) |
| Exercise rubric | ~40 + ~70 danych | Usunięcie sekcji punktacji | Niskie (nowa komenda) |
| Error code suggestions | ~20 | Uproszczenie 1 instrukcji | Niskie (wzbogacenie error) |
| **Suma** | **~170** | **Netto mniejszy SKILL.md** | **Niskie** |
