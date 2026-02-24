# CLI Reliability Migration — Implementation Plan

> **For Claude:** REQUIRED SUB-SKILL: Use superpowers:executing-plans to implement this plan task-by-task.

**Goal:** Migrate 3 fragile/non-deterministic AI-side behaviors to Go CLI for reliability: cheatsheet auto-attach in hints, exercise rubric command, error code suggestions.

**Architecture:** Three independent changes to the Go CLI (`matura_informatyka_rozszerzona/analiza/cli/`). Change 1 enriches `exercise hints` response. Change 2 adds new `exercise rubric` command. Change 3 enriches `progress blad` error response. Each change includes Go code + tests + SKILL.md update.

**Tech Stack:** Go 1.21+, `modernc.org/sqlite`, `github.com/spf13/cobra`, existing test harness (`main_test.go`)

**Key paths:**
- CLI source: `matura_informatyka_rozszerzona/analiza/cli/`
- SKILL.md: `.claude/skills/matura/SKILL.md`
- QA: `matura_informatyka_rozszerzona/analiza/test_qa.sh --layer 5` (Go tests), `--layer 3` (SKILL lint)
- Build: `matura_informatyka_rozszerzona/analiza/cli/build.sh`

---

### Task 1: Cheatsheet auto-attach in `exercise hints` — test

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/main_test.go`

**Step 1: Write the failing test**

Add after the existing `TestExerciseHintsById` test (~line 2485):

```go
func TestExerciseHintsCheatsheetExcerpt(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Pick a cyfry_liczby exercise (IMPLEMENTACJA category)
	results, _ := queryExercises(db, "cyfry_liczby", "latwe", "")
	if len(results) == 0 {
		t.Fatal("no exercises")
	}
	ex := results[0]

	// Register fetch so hints aren't blocked
	registerFetch(db, ex.ID, "cyfry_liczby", 0)

	hints := getExerciseHints(db, ex.ID, "latwe")
	if hints.ID != ex.ID {
		t.Fatalf("got id %q, want %q", hints.ID, ex.ID)
	}
	// New: cheatsheet_excerpt should be populated when exercise has tags
	// that match cheatsheet sections
	if hints.CheatsheetExcerpt == "" {
		t.Error("expected non-empty cheatsheet_excerpt for cyfry_liczby exercise")
	}
}

func TestExerciseHintsCheatsheetExcerptNoMatch(t *testing.T) {
	dir := testDir(t)
	db := openTestDB(t, dir)

	// Create a minimal exercise with tags that won't match any cheatsheet section
	hints := getExerciseHints(db, "nonexistent_id_99", "latwe")
	// For nonexistent exercise, excerpt should be empty (graceful fallback)
	if hints.CheatsheetExcerpt != "" {
		t.Errorf("expected empty excerpt for unknown exercise, got %q", hints.CheatsheetExcerpt)
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestExerciseHintsCheatsheetExcerpt -v -count=1`
Expected: FAIL — `HintsOut` has no field `CheatsheetExcerpt`

**Step 3: Commit**

```
test: add failing tests for cheatsheet auto-attach in hints
```

---

### Task 2: Cheatsheet auto-attach in `exercise hints` — implement

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/types.go` (add field to HintsOut)
- Modify: `matura_informatyka_rozszerzona/analiza/cli/commands.go` (enrich getExerciseHints)

**Step 1: Add `CheatsheetExcerpt` field to `HintsOut`**

In `types.go`, modify `HintsOut` (line ~132):

```go
// HintsOut is what exercise hints returns
type HintsOut struct {
	ID                 string   `json:"id"`
	Wskazowki          []string `json:"wskazowki"`
	MaxHints           int      `json:"max_hints"`
	CheatsheetExcerpt  string   `json:"cheatsheet_excerpt,omitempty"`
}
```

**Step 2: Add `findCheatsheetExcerptByTags` helper in `commands.go`**

Add near the `extractSection` function (~line 2748):

```go
// findCheatsheetExcerptByTags searches the cheatsheet for the exercise's
// category, matching against the exercise's tags. Returns the first matching
// section, or "" if no match.
func findCheatsheetExcerptByTags(d *sql.DB, exerciseID string) string {
	var kategoria, tagiJSON string
	err := d.QueryRow(`SELECT kategoria, tagi FROM data.cwiczenia WHERE id = ?`, exerciseID).Scan(&kategoria, &tagiJSON)
	if err != nil {
		return ""
	}
	var tagi []string
	json.Unmarshal([]byte(tagiJSON), &tagi)
	if len(tagi) == 0 {
		return ""
	}

	var content string
	err = d.QueryRow(`SELECT content FROM data.cheatsheets WHERE kategoria = ?`, kategoria).Scan(&content)
	if err != nil {
		return ""
	}

	// Try each tag against cheatsheet sections
	for _, tag := range tagi {
		if section := extractSection(content, tag); section != "" {
			return section
		}
	}
	return ""
}
```

**Step 3: Modify `getExerciseHints` to call the helper**

In `commands.go` (~line 250), change `getExerciseHints`:

```go
func getExerciseHints(d *sql.DB, id, level string) HintsOut {
	var wskazowkiJSON string
	d.QueryRow("SELECT wskazowki FROM data.cwiczenia WHERE id = ?", id).Scan(&wskazowkiJSON)
	var wskazowki []string
	json.Unmarshal([]byte(wskazowkiJSON), &wskazowki)
	if wskazowki == nil {
		wskazowki = []string{}
	}
	maxHints := calculateMaxHints(level)
	if maxHints >= 0 && len(wskazowki) > maxHints {
		wskazowki = wskazowki[:maxHints]
	}
	if maxHints < 0 {
		maxHints = len(wskazowki)
	}
	out := HintsOut{ID: id, Wskazowki: wskazowki, MaxHints: maxHints}

	// Auto-attach cheatsheet section when hints include Level 2+
	if len(wskazowki) >= 2 {
		out.CheatsheetExcerpt = findCheatsheetExcerptByTags(d, id)
	}

	return out
}
```

**Step 4: Run tests**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestExerciseHints -v -count=1`
Expected: PASS for both new and existing tests

**Step 5: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v -count=1 ./...`
Expected: All tests pass

**Step 6: Commit**

```
feat: auto-attach cheatsheet excerpt in exercise hints
```

---

### Task 3: Exercise rubric command — test

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/main_test.go`

**Step 1: Write the failing test**

```go
func TestExerciseRubricSledzenie(t *testing.T) {
	rubric, err := getRubric("sledzenie_algorytmu")
	if err != nil {
		t.Fatalf("getRubric: %v", err)
	}
	if rubric.Typ != "sledzenie_algorytmu" {
		t.Errorf("got typ %q", rubric.Typ)
	}
	if rubric.Kategoria != "TEORIA" {
		t.Errorf("got kategoria %q", rubric.Kategoria)
	}
	if rubric.Rubric.Full.Opis == "" {
		t.Error("empty full.opis")
	}
	if rubric.Rubric.Half.Opis == "" {
		t.Error("empty half.opis")
	}
	if rubric.Rubric.Zero.Opis == "" {
		t.Error("empty zero.opis")
	}
	if rubric.Rubric.Full.Procent != 100 {
		t.Errorf("full.procent = %d, want 100", rubric.Rubric.Full.Procent)
	}
}

func TestExerciseRubricAllTypes(t *testing.T) {
	allTypes := []string{
		"sledzenie_algorytmu", "projektowanie_algorytmu", "analiza_algorytmu",
		"test_prawda_falsz", "konwersja_systemow_liczbowych", "teoria_bezpieczenstwa",
		"cyfry_liczby", "napisy", "zlozone", "zliczanie", "minmax", "sekwencje", "obrazy_2D", "geometryczne",
		"agregacja_warunkowa", "symulacja", "wykres", "agregacja_podstawowa", "transformacja",
		"sql_group_by", "sql_podzapytania", "sql_join", "sql_select_where",
	}
	for _, typ := range allTypes {
		rubric, err := getRubric(typ)
		if err != nil {
			t.Errorf("getRubric(%q): %v", typ, err)
			continue
		}
		if rubric.Rubric.Full.Opis == "" {
			t.Errorf("%s: empty full.opis", typ)
		}
	}
}

func TestExerciseRubricUnknownType(t *testing.T) {
	_, err := getRubric("nonexistent_type")
	if err == nil {
		t.Error("expected error for unknown type")
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestExerciseRubric -v -count=1`
Expected: FAIL — `getRubric` undefined

**Step 3: Commit**

```
test: add failing tests for exercise rubric command
```

---

### Task 4: Exercise rubric command — implement

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/types.go` (add RubricOut types)
- Modify: `matura_informatyka_rozszerzona/analiza/cli/commands.go` (add rubric data + command)
- Modify: `matura_informatyka_rozszerzona/analiza/cli/main.go` (register command)

**Step 1: Add types in `types.go`**

Add after `AnswerOut` (~line 143):

```go
// RubricLevel describes one scoring level
type RubricLevel struct {
	Opis    string `json:"opis"`
	Procent int    `json:"procent"`
}

// RubricDetail holds all scoring levels for a type
type RubricDetail struct {
	Full  RubricLevel `json:"full"`
	Half  RubricLevel `json:"half"`
	Zero  RubricLevel `json:"zero"`
	Notes string      `json:"notes,omitempty"`
}

// RubricOut is what exercise rubric returns
type RubricOut struct {
	Typ       string       `json:"typ"`
	Kategoria string       `json:"kategoria"`
	Rubric    RubricDetail `json:"rubric"`
}
```

**Step 2: Add rubric data + `getRubric` + command in `commands.go`**

Add the rubric data map and command function. Place the data map near `errorCodeWhitelist` reference area, and the command function near `exerciseAnswerCmd`:

```go
// typKategoriaMap maps exercise types to their category
var typKategoriaMap = map[string]string{
	"sledzenie_algorytmu": "TEORIA", "projektowanie_algorytmu": "TEORIA",
	"analiza_algorytmu": "TEORIA", "test_prawda_falsz": "TEORIA",
	"konwersja_systemow_liczbowych": "TEORIA", "teoria_bezpieczenstwa": "TEORIA",
	"cyfry_liczby": "IMPLEMENTACJA", "napisy": "IMPLEMENTACJA",
	"zlozone": "IMPLEMENTACJA", "zliczanie": "IMPLEMENTACJA",
	"minmax": "IMPLEMENTACJA", "sekwencje": "IMPLEMENTACJA",
	"obrazy_2D": "IMPLEMENTACJA", "geometryczne": "IMPLEMENTACJA",
	"agregacja_warunkowa": "ARKUSZ", "symulacja": "ARKUSZ",
	"wykres": "ARKUSZ", "agregacja_podstawowa": "ARKUSZ",
	"transformacja": "ARKUSZ",
	"sql_group_by": "SQL", "sql_podzapytania": "SQL",
	"sql_join": "SQL", "sql_select_where": "SQL",
}

// rubricData holds CKE-based grading rules per exercise type
var rubricData = map[string]RubricDetail{
	// TEORIA
	"sledzenie_algorytmu": {
		Full:  RubricLevel{"Tabela poprawna + wynik koncowy poprawny", 100},
		Half:  RubricLevel{"Poprawny tok rozumowania, 1-2 bledy rachunkowe w wierszach", 50},
		Zero:  RubricLevel{"Zly algorytm lub brak tabeli", 0},
		Notes: "Kazdy wiersz tabeli jest wart punkty. Bledny wynik przy poprawnym toku = 50%.",
	},
	"projektowanie_algorytmu": {
		Full:  RubricLevel{"Poprawny pseudokod/C++ rozwiazujacy problem", 100},
		Half:  RubricLevel{"Poprawna idea, bledy skladniowe lub drobne luki", 50},
		Zero:  RubricLevel{"Zly algorytm lub brak rozwiazania", 0},
		Notes: "Liczy sie poprawnosc algorytmu, nie skladnia. Brak obslugi przypadkow brzegowych = -25%.",
	},
	"analiza_algorytmu": {
		Full:  RubricLevel{"Poprawna klasa zlozonosci O() + uzasadnienie", 100},
		Half:  RubricLevel{"Poprawna klasa zlozonosci bez uzasadnienia", 50},
		Zero:  RubricLevel{"Zla klasa zlozonosci", 0},
		Notes: "Uzasadnienie musi odwolywac sie do struktury algorytmu (petla, rekurencja).",
	},
	"test_prawda_falsz": {
		Full:  RubricLevel{"Poprawne P/F + poprawne uzasadnienie", 100},
		Half:  RubricLevel{"Poprawne P/F bez uzasadnienia lub z blednym uzasadnieniem", 50},
		Zero:  RubricLevel{"Bledne P/F", 0},
		Notes: "Brak uzasadnienia = ZAWSZE max 50%. CKE wymaga uzasadnienia nawet przy poprawnej odpowiedzi.",
	},
	"konwersja_systemow_liczbowych": {
		Full:  RubricLevel{"Poprawny wynik + zapis posredni obliczen", 100},
		Half:  RubricLevel{"Poprawny wynik bez zapisu posredniego", 50},
		Zero:  RubricLevel{"Bledny wynik", 0},
		Notes: "Zapis posredni: kolumna dzielenia z resztami, grupowanie bitow, itd.",
	},
	"teoria_bezpieczenstwa": {
		Full:  RubricLevel{"Poprawne dopasowanie + definicja/wyjasnienie", 100},
		Half:  RubricLevel{"Poprawne dopasowanie bez definicji", 50},
		Zero:  RubricLevel{"Bledne dopasowanie", 0},
		Notes: "Przy pytaniach otwartych: wymagana precyzyjna definicja, nie ogolniki.",
	},
	// IMPLEMENTACJA (shared rubric)
	"cyfry_liczby":  implRubric(),
	"napisy":        implRubric(),
	"zlozone":       implRubric(),
	"zliczanie":     implRubric(),
	"minmax":        implRubric(),
	"sekwencje":     implRubric(),
	"obrazy_2D":     implRubric(),
	"geometryczne":  implRubric(),
	// ARKUSZ (shared rubric)
	"agregacja_warunkowa":  arkuszRubric(),
	"symulacja":            arkuszRubric(),
	"wykres":               arkuszRubric(),
	"agregacja_podstawowa": arkuszRubric(),
	"transformacja":        arkuszRubric(),
	// SQL (shared rubric)
	"sql_group_by":     sqlRubric(),
	"sql_podzapytania": sqlRubric(),
	"sql_join":         sqlRubric(),
	"sql_select_where": sqlRubric(),
}

func implRubric() RubricDetail {
	return RubricDetail{
		Full:  RubricLevel{"Program kompiluje sie, daje poprawne wyniki dla danych przykladowych i pelnych", 100},
		Half:  RubricLevel{"Poprawny algorytm, ale bledy: off-by-one, brak obslugi brzegowych, bledne I/O", 50},
		Zero:  RubricLevel{"Zly algorytm lub program sie nie kompiluje", 0},
		Notes: "Czesciowe punkty: poprawne wczytanie (25%), poprawny algorytm z drobnymi bledami (50%), poprawne wypisanie (75%).",
	}
}

func arkuszRubric() RubricDetail {
	return RubricDetail{
		Full:  RubricLevel{"Poprawna formula dajaca prawidlowe wyniki + poprawne adresowanie", 100},
		Half:  RubricLevel{"Poprawna idea formuly, ale blad adresowania ($) lub zly zakres", 50},
		Zero:  RubricLevel{"Zla formula lub brak rozwiazania", 0},
		Notes: "Blad $ (brak dolara przy stale) = najczestszy blad. Poprawna formula ze zlym zakresem = 50%.",
	}
}

func sqlRubric() RubricDetail {
	return RubricDetail{
		Full:  RubricLevel{"Poprawne zapytanie SQL dajace prawidlowy wynik", 100},
		Half:  RubricLevel{"Poprawna struktura (tabele, JOIN, GROUP BY), ale blad w warunku lub agregacji", 50},
		Zero:  RubricLevel{"Zla struktura zapytania lub brak rozwiazania", 0},
		Notes: "Kolejnosc kolumn w wyniku nie ma znaczenia. Aliasy opcjonalne. Brak GROUP BY przy agregacji = 0%.",
	}
}

// getRubric returns grading rubric for the given exercise type.
func getRubric(typ string) (RubricOut, error) {
	rubric, ok := rubricData[typ]
	if !ok {
		return RubricOut{}, fmt.Errorf("unknown exercise type: %s", typ)
	}
	kat, ok := typKategoriaMap[typ]
	if !ok {
		kat = ""
	}
	return RubricOut{Typ: typ, Kategoria: kat, Rubric: rubric}, nil
}

func exerciseRubricCmd() *cobra.Command {
	var typ string
	cmd := &cobra.Command{
		Use:   "rubric",
		Short: "Get grading rubric for an exercise type",
		RunE: func(cmd *cobra.Command, args []string) error {
			if typ == "" {
				return fatal("--typ is required")
			}
			out, err := getRubric(typ)
			if err != nil {
				return fatal(err.Error())
			}
			jsonOut(out)
			return nil
		},
	}
	cmd.Flags().StringVar(&typ, "typ", "", "Exercise type (e.g. sledzenie_algorytmu)")
	return cmd
}
```

**Step 3: Register the command in `main.go`**

In `main.go` line 66, add `exerciseRubricCmd()`:

```go
exerciseCmd.AddCommand(exerciseQuestionCmd(), exerciseHintsCmd(), exerciseAnswerCmd(), exerciseReviewCmd(), exerciseNextCmd(), exerciseRubricCmd())
```

**Step 4: Run tests**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestExerciseRubric -v -count=1`
Expected: All 3 tests PASS

**Step 5: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v -count=1 ./...`
Expected: All tests pass

**Step 6: Commit**

```
feat: add exercise rubric command with CKE grading rules
```

---

### Task 5: Error code suggestions — test

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/main_test.go`

**Step 1: Write the failing test**

```go
func TestSuggestClosestCodes(t *testing.T) {
	// Exact match — no suggestions needed
	suggestions := suggestClosestCodes("sql_group_by", "brak_having")
	if len(suggestions) != 0 {
		t.Errorf("expected no suggestions for valid code, got %d", len(suggestions))
	}

	// Typo — should suggest closest
	suggestions = suggestClosestCodes("sql_group_by", "brak_havng")
	if len(suggestions) == 0 {
		t.Fatal("expected suggestions for typo")
	}
	if suggestions[0].Kod != "brak_having" {
		t.Errorf("got %q, want brak_having", suggestions[0].Kod)
	}
	if suggestions[0].Opis == "" {
		t.Error("expected non-empty opis in suggestion")
	}

	// Completely wrong — should still suggest something
	suggestions = suggestClosestCodes("sql_group_by", "completely_wrong")
	if len(suggestions) == 0 {
		t.Fatal("expected suggestions even for completely wrong code")
	}
	// Should return max 3 suggestions
	if len(suggestions) > 3 {
		t.Errorf("expected max 3 suggestions, got %d", len(suggestions))
	}
}

func TestLevenshteinDistance(t *testing.T) {
	tests := []struct {
		a, b string
		want int
	}{
		{"", "", 0},
		{"abc", "", 3},
		{"", "abc", 3},
		{"abc", "abc", 0},
		{"brak_having", "brak_havng", 1},
		{"kitten", "sitting", 3},
	}
	for _, tt := range tests {
		got := levenshtein(tt.a, tt.b)
		if got != tt.want {
			t.Errorf("levenshtein(%q, %q) = %d, want %d", tt.a, tt.b, got, tt.want)
		}
	}
}
```

**Step 2: Run test to verify it fails**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run TestSuggestClosest -v -count=1`
Expected: FAIL — `suggestClosestCodes` undefined

**Step 3: Commit**

```
test: add failing tests for error code suggestions
```

---

### Task 6: Error code suggestions — implement

**Files:**
- Modify: `matura_informatyka_rozszerzona/analiza/cli/error_codes.go` (add descriptions, levenshtein, suggest)
- Modify: `matura_informatyka_rozszerzona/analiza/cli/commands.go` (use suggestions in `progressBladCmd`)

**Step 1: Add error code descriptions + levenshtein + suggest to `error_codes.go`**

Add after the existing `validateErrorCode` function:

```go
// errorCodeDescriptions maps error codes to human-readable descriptions.
var errorCodeDescriptions = map[string]string{
	// TEORIA - sledzenie
	"mylenie_div_mod":           "Mylenie operatorow div i mod",
	"zla_kolejnosc_sledzenia":   "Zla kolejnosc sledzenia (np. od konca zamiast od poczatku)",
	"pominiecie_bazy_rekurencji": "Pominiecie warunku bazowego rekurencji",
	"zly_mnoznik":               "Zly mnoznik lub potega w obliczeniach",
	"brak_tabeli_sledzenia":     "Brak tabeli sledzenia (tabelka krok-po-kroku)",
	"zla_parzystosc_cyfry":      "Zla ocena parzystosci cyfry",
	"bledne_wciecia_blok":       "Bledna interpretacja bloku kodu (wciecia)",
	// TEORIA - projektowanie
	"zly_algorytm":              "Wybor zlego algorytmu do problemu",
	"brak_warunku_stopu":        "Brak warunku stopu w petli/rekurencji",
	"bledna_skladnia_pseudokod": "Bledna skladnia pseudokodu CKE",
	"niepoprawna_petla":         "Niepoprawna petla (zly zakres, zly krok)",
	"brak_inicjalizacji":        "Brak inicjalizacji zmiennej",
	// TEORIA - analiza
	"zla_zlozonosc_klasa":           "Zla klasa zlozonosci (np. O(n) zamiast O(n^2))",
	"brak_uzasadnienia_zlozonosc":   "Brak uzasadnienia klasy zlozonosci",
	"mylenie_avg_worst":             "Mylenie zlozonosci sredniej i pesymistycznej",
	"zly_kontrprzyklad":             "Zly kontrprzyklad",
	"brak_wzoru":                    "Brak wzoru lub wyprowadzenia",
	// TEORIA - P/F
	"brak_uzasadnienia_pf":          "Brak uzasadnienia odpowiedzi P/F",
	"mylenie_avg_worst_pf":          "Mylenie zlozonosci sredniej i pesymistycznej (P/F)",
	"nieprecyzyjne_uzasadnienie":    "Nieprecyzyjne uzasadnienie (za ogolne)",
	"pomylenie_stabilnosci_sortowania": "Pomylenie stabilnosci sortowania",
	// TEORIA - konwersja
	"zla_baza_konwersji":       "Zla baza docelowa lub zrodlowa",
	"zla_kolejnosc_reszt":      "Zla kolejnosc reszt (od dolu zamiast od gory)",
	"brak_zapisu_posredniego":  "Brak zapisu posredniego obliczen",
	"zle_grupowanie_bitow":     "Zle grupowanie bitow (bin->hex: grupy po 4)",
	"blad_uzupelnienia_do_2":   "Blad w kodzie uzupelnienia do 2 (U2)",
	// TEORIA - bezpieczenstwo
	"mylenie_typow_malware":            "Mylenie typow malware (trojan, worm, ransomware)",
	"mylenie_szyfrowania_sym_asym":     "Mylenie szyfrowania symetrycznego i asymetrycznego",
	"mylenie_protokolow":               "Mylenie protokolow sieciowych",
	"brak_rozroznienia_klucz_pub_pryw": "Brak rozroznienia klucz publiczny/prywatny",
	// SQL
	"brak_group_by":           "Brak GROUP BY przy agregacji",
	"zly_join_warunek":        "Zly warunek JOIN (ON)",
	"brak_having":             "Brak HAVING (filtrowanie po agregacji)",
	"zla_agregacja":           "Zla funkcja agregujaca (SUM/COUNT/AVG)",
	"null_zamiast_is_null":    "Uzycie = NULL zamiast IS NULL",
	"count_star_vs_kolumna":   "Mylenie COUNT(*) z COUNT(kolumna)",
	"zla_kolejnosc_klauzul":   "Zla kolejnosc klauzul SQL",
	// IMPLEMENTACJA
	"zly_warunek_petli":        "Zly warunek petli (< vs <=, itd.)",
	"brak_wczytania":           "Brak wczytania danych z pliku",
	"off_by_one":               "Blad o jeden (indeks, granica petli)",
	"dzielenie_calkowite":      "Dzielenie calkowite zamiast zmiennoprzecinkowego",
	"zle_indeksowanie":         "Zle indeksowanie (od 0 vs od 1)",
	"brak_obslugi_brzegowych":  "Brak obslugi przypadkow brzegowych",
	"zla_kolejnosc_operacji":   "Zla kolejnosc operacji",
	// ARKUSZ
	"zle_adresowanie":          "Zle adresowanie komorek",
	"brak_dolara":              "Brak $ (adresowanie bezwzgledne)",
	"zla_formula_warunkowa":    "Zla formula warunkowa (SUMIFS, COUNTIF)",
	"stala_zamiast_odwolania":  "Stala zamiast odwolania do komorki",
	"brak_kolumny_pomocniczej": "Brak kolumny pomocniczej",
}

// levenshtein calculates the Levenshtein edit distance between two strings.
func levenshtein(a, b string) int {
	la, lb := len(a), len(b)
	if la == 0 {
		return lb
	}
	if lb == 0 {
		return la
	}
	prev := make([]int, lb+1)
	curr := make([]int, lb+1)
	for j := 0; j <= lb; j++ {
		prev[j] = j
	}
	for i := 1; i <= la; i++ {
		curr[0] = i
		for j := 1; j <= lb; j++ {
			cost := 1
			if a[i-1] == b[j-1] {
				cost = 0
			}
			curr[j] = min(curr[j-1]+1, min(prev[j]+1, prev[j-1]+cost))
		}
		prev, curr = curr, prev
	}
	return prev[lb]
}

// CodeSuggestion is a suggested error code with description
type CodeSuggestion struct {
	Kod  string `json:"kod"`
	Opis string `json:"opis"`
}

// suggestClosestCodes returns top 3 closest error codes for the given type.
// Returns empty slice if kod is valid (exact match).
func suggestClosestCodes(typ, kod string) []CodeSuggestion {
	valid, allowed := validateErrorCode(typ, kod)
	if valid {
		return nil
	}
	if len(allowed) == 0 {
		return nil
	}

	type scored struct {
		code string
		dist int
	}
	var scored_list []scored
	for _, c := range allowed {
		scored_list = append(scored_list, scored{c, levenshtein(kod, c)})
	}
	sort.Slice(scored_list, func(i, j int) bool {
		return scored_list[i].dist < scored_list[j].dist
	})

	limit := 3
	if len(scored_list) < limit {
		limit = len(scored_list)
	}
	result := make([]CodeSuggestion, limit)
	for i := 0; i < limit; i++ {
		opis := errorCodeDescriptions[scored_list[i].code]
		result[i] = CodeSuggestion{Kod: scored_list[i].code, Opis: opis}
	}
	return result
}
```

Note: This file needs `"sort"` in imports. Since `error_codes.go` currently has no imports, add:

```go
package main

import "sort"
```

**Step 2: Modify `progressBladCmd` in `commands.go` to return suggestions on error**

In `commands.go` (~line 1344), replace the error return with enriched JSON:

Change:
```go
		valid, allowed := validateErrorCode(typ, kod)
		if !valid {
			return fatal(fmt.Sprintf("kod '%s' niedostepny dla typ '%s'. Dozwolone: %v", kod, typ, allowed))
		}
```

To:
```go
		valid, _ := validateErrorCode(typ, kod)
		if !valid {
			suggestions := suggestClosestCodes(typ, kod)
			// Build valid_codes with descriptions
			_, allowedCodes := validateErrorCode(typ, "___placeholder___")
			type codeWithDesc struct {
				Kod  string `json:"kod"`
				Opis string `json:"opis"`
			}
			validCodes := make([]codeWithDesc, len(allowedCodes))
			for i, c := range allowedCodes {
				validCodes[i] = codeWithDesc{c, errorCodeDescriptions[c]}
			}
			out := map[string]interface{}{
				"error":       fmt.Sprintf("kod '%s' niedostepny dla typ '%s'", kod, typ),
				"valid_codes": validCodes,
				"suggestions": suggestions,
			}
			jsonOut(out)
			// Return fatal to get exit code 2
			return fatal(fmt.Sprintf("kod '%s' niedostepny dla typ '%s'", kod, typ))
		}
```

Wait — this would print JSON AND return error. Better approach: print JSON error and exit cleanly with code 2. Looking at the codebase pattern, `fatal()` returns an error that main handles by printing to stderr. But we want structured JSON on stdout.

Revised approach — output JSON and then return fatal (the fatal message goes to stderr, JSON to stdout):

Actually, the simplest approach following existing patterns: change the fatal message to include the suggestion, but output the full structured response via jsonOut. The AI reads stdout (JSON), not stderr.

```go
		valid, _ := validateErrorCode(typ, kod)
		if !valid {
			suggestions := suggestClosestCodes(typ, kod)
			_, allowedCodes := validateErrorCode(typ, "___force_invalid___")
			type codeEntry struct {
				Kod  string `json:"kod"`
				Opis string `json:"opis"`
			}
			validList := make([]codeEntry, len(allowedCodes))
			for i, c := range allowedCodes {
				validList[i] = codeEntry{c, errorCodeDescriptions[c]}
			}
			jsonOut(map[string]interface{}{
				"error":       fmt.Sprintf("kod '%s' niedostepny dla typ '%s'", kod, typ),
				"valid_codes": validList,
				"suggestions": suggestions,
			})
			return fatal(fmt.Sprintf("kod '%s' niedostepny", kod))
		}
```

**Step 3: Run tests**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -run "TestSuggestClosest|TestLevenshtein|TestValidateErrorCode" -v -count=1`
Expected: All PASS

**Step 4: Run full test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v -count=1 ./...`
Expected: All tests pass

**Step 5: Commit**

```
feat: add error code suggestions with Levenshtein matching
```

---

### Task 7: Update SKILL.md

**Files:**
- Modify: `.claude/skills/matura/SKILL.md`

**Step 1: Remove hardcoded cheatsheet mappings (lines 361-364)**

Replace the Level 2 hint section. Change lines 359-364 from:
```
       POTEM: `wskazowki[1]` + **[WYMAGANE]** cytat z cheatsheet:
       `./matura cheatsheet get --kategoria {kat} --sekcja "{temat}"` ← MUSISZ wywolac
       Mapowanie: mod/div→"archetyp", rekurencja→"rekurencj", zlozonosc→"zlozonosc",
       JOIN→"join", GROUP BY→"group", sortowanie→"sort", adresowanie→"adresow",
       szyfrowanie→"bezpieczen", P/F→"prawda", konwersja→"konwersj"
```

To:
```
       POTEM: `wskazowki[1]` + jesli `cheatsheet_excerpt` w odpowiedzi hints jest niepuste → cytuj go uczniowi
```

**Step 2: Remove hardcoded grading rules (lines 382-394)**

Replace the "Punktacja czesciowa" section. Change lines 382-394 from the full table to:
```
### Punktacja czesciowa

Pobierz kryteria punktacji: `./matura exercise rubric --typ {typ}`
Zwraca: `{rubric: {full: {opis, procent}, half: {opis, procent}, zero: {opis, procent}, notes}}`.
Stosuj te kryteria przy ocenie odpowiedzi ucznia.
Regula ogolna: jesli uczen ma poprawny tok rozumowania ale drobny blad rachunkowy -> 50-75% pkt.
```

**Step 3: Simplify error code validation instructions (lines 127, 341, 437-438)**

Change line 127 from:
```
- `progress blad --kod Z` z niepoprawnym kodem → CLI odrzuci i zwroci liste dozwolonych kodow. Wybierz najblizszy z listy.
```
To:
```
- `progress blad --kod Z` z niepoprawnym kodem → CLI zwroci JSON z `suggestions[]` (kody z opisami). Uzyj `suggestions[0].kod` jesli opis pasuje do bledu ucznia.
```

Change line 341 from:
```
   - CLI odrzuci niepoprawny kod i zwroci liste dozwolonych — wybierz najblizszy
```
To:
```
   - CLI odrzuci niepoprawny kod i zwroci `suggestions[]` z opisami — uzyj pierwszej pasajacej sugestii
```

Change lines 437-438 from:
```
CLI waliduje kody bledow per typ — uzyj dowolnego kodu opisujacego pomylke ucznia.
Jesli kod jest niepoprawny, CLI odrzuci i zwroci liste dozwolonych kodow — wybierz najblizszy.
```
To:
```
CLI waliduje kody bledow per typ — uzyj dowolnego kodu opisujacego pomylke ucznia.
Jesli kod jest niepoprawny, CLI zwroci `suggestions[]` z najblizszymi kodami i opisami — uzyj pasujacego.
```

**Step 4: Add `exercise rubric` to CLI reference table (line ~105)**

Add row after the `Zapisz blad` row:
```
| Punktacja typu | `./matura exercise rubric --typ {typ}` |
```

**Step 5: Verify SKILL.md lint**

Run: `matura_informatyka_rozszerzona/analiza/test_qa.sh --layer 3`
Expected: PASS

**Step 6: Commit**

```
docs: update SKILL.md for CLI reliability migration (cheatsheet auto-attach, rubric, suggestions)
```

---

### Task 8: Build + integration verification

**Files:**
- No changes — verification only

**Step 1: Build binaries**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && bash build.sh`
Expected: Builds macOS + Windows binaries, re-imports matura.db

**Step 2: Run full Go test suite**

Run: `cd matura_informatyka_rozszerzona/analiza/cli && go test -v -count=1 ./...`
Expected: All tests pass

**Step 3: Run QA layer 5 (Go tests)**

Run: `matura_informatyka_rozszerzona/analiza/test_qa.sh --layer 5`
Expected: PASS

**Step 4: Run QA layer 3 (SKILL lint)**

Run: `matura_informatyka_rozszerzona/analiza/test_qa.sh --layer 3`
Expected: PASS

**Step 5: Run QA layer 1 (CLI smoke tests)**

Run: `matura_informatyka_rozszerzona/analiza/test_qa.sh --layer 1`
Expected: PASS (new `exercise rubric` should not break existing smoke tests)

**Step 6: Manual spot-check new features**

```bash
CLI_DIR="matura_informatyka_rozszerzona/analiza/cli"
# Test rubric
$CLI_DIR/matura exercise rubric --typ sledzenie_algorytmu
# Expected: JSON with rubric.full, rubric.half, rubric.zero

# Test hints cheatsheet excerpt (need to set up active exercise first)
$CLI_DIR/matura exercise next --typ cyfry_liczby
# Copy the ID, then:
$CLI_DIR/matura exercise hints --id {id}
# Expected: JSON with cheatsheet_excerpt field (may be null if no tag match)

# Test error code suggestions
$CLI_DIR/matura progress blad --exercise-id 7.1 --typ cyfry_liczby --kod "brak_havng" --hint 0
# Expected: JSON on stdout with suggestions[] before error on stderr
```

**Step 7: Commit (if any fixes were needed)**

```
fix: integration fixes for CLI reliability migration
```
