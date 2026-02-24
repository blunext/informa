package main

import "sort"

// errorCodeWhitelist maps exercise type → allowed error codes.
// CLI rejects any code not in the whitelist for the given type.
var errorCodeWhitelist = map[string][]string{
	"sledzenie_algorytmu": {
		"mylenie_div_mod", "zla_kolejnosc_sledzenia", "pominiecie_bazy_rekurencji",
		"zly_mnoznik", "brak_tabeli_sledzenia", "zla_parzystosc_cyfry", "bledne_wciecia_blok",
	},
	"projektowanie_algorytmu": {
		"zly_algorytm", "brak_warunku_stopu", "bledna_skladnia_pseudokod",
		"niepoprawna_petla", "brak_inicjalizacji",
	},
	"analiza_algorytmu": {
		"zla_zlozonosc_klasa", "brak_uzasadnienia_zlozonosc", "mylenie_avg_worst",
		"zly_kontrprzyklad", "brak_wzoru",
	},
	"test_prawda_falsz": {
		"brak_uzasadnienia_pf", "mylenie_avg_worst_pf", "nieprecyzyjne_uzasadnienie",
		"pomylenie_stabilnosci_sortowania",
	},
	"konwersja_systemow_liczbowych": {
		"zla_baza_konwersji", "zla_kolejnosc_reszt", "brak_zapisu_posredniego",
		"zle_grupowanie_bitow", "blad_uzupelnienia_do_2",
	},
	"teoria_bezpieczenstwa": {
		"mylenie_typow_malware", "mylenie_szyfrowania_sym_asym",
		"mylenie_protokolow", "brak_rozroznienia_klucz_pub_pryw",
	},
	// SQL types share same codes
	"sql_group_by":     {"brak_group_by", "zly_join_warunek", "brak_having", "zla_agregacja", "null_zamiast_is_null", "count_star_vs_kolumna", "zla_kolejnosc_klauzul"},
	"sql_podzapytania": {"brak_group_by", "zly_join_warunek", "brak_having", "zla_agregacja", "null_zamiast_is_null", "count_star_vs_kolumna", "zla_kolejnosc_klauzul"},
	"sql_join":         {"brak_group_by", "zly_join_warunek", "brak_having", "zla_agregacja", "null_zamiast_is_null", "count_star_vs_kolumna", "zla_kolejnosc_klauzul"},
	"sql_select_where": {"brak_group_by", "zly_join_warunek", "brak_having", "zla_agregacja", "null_zamiast_is_null", "count_star_vs_kolumna", "zla_kolejnosc_klauzul"},
	// IMPLEMENTACJA types share same codes
	"cyfry_liczby": {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji", "mylenie_div_mod"},
	"napisy":       {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
	"zlozone":      {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
	"zliczanie":    {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
	"minmax":       {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
	"sekwencje":    {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
	"obrazy_2D":    {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
	"geometryczne": {"brak_inicjalizacji", "zly_warunek_petli", "brak_wczytania", "off_by_one", "dzielenie_calkowite", "zle_indeksowanie", "brak_obslugi_brzegowych", "zla_kolejnosc_operacji"},
	// ARKUSZ types share same codes
	"agregacja_warunkowa":  {"zle_adresowanie", "brak_dolara", "zla_formula_warunkowa", "stala_zamiast_odwolania", "brak_kolumny_pomocniczej"},
	"symulacja":            {"zle_adresowanie", "brak_dolara", "zla_formula_warunkowa", "stala_zamiast_odwolania", "brak_kolumny_pomocniczej"},
	"wykres":               {"zle_adresowanie", "brak_dolara", "zla_formula_warunkowa", "stala_zamiast_odwolania", "brak_kolumny_pomocniczej"},
	"agregacja_podstawowa": {"zle_adresowanie", "brak_dolara", "zla_formula_warunkowa", "stala_zamiast_odwolania", "brak_kolumny_pomocniczej"},
	"transformacja":        {"zle_adresowanie", "brak_dolara", "zla_formula_warunkowa", "stala_zamiast_odwolania", "brak_kolumny_pomocniczej"},
}

// validateErrorCode checks if kod is allowed for the given typ.
// Returns allowed list if invalid.
func validateErrorCode(typ, kod string) (valid bool, allowed []string) {
	codes, ok := errorCodeWhitelist[typ]
	if !ok {
		// Unknown type — allow any code (backwards compat)
		return true, nil
	}
	for _, c := range codes {
		if c == kod {
			return true, nil
		}
	}
	return false, codes
}

// errorCodeDescriptions maps each error code to a human-readable Polish description.
var errorCodeDescriptions = map[string]string{
	// TEORIA - sledzenie
	"mylenie_div_mod":            "Mylenie operatorow div i mod",
	"zla_kolejnosc_sledzenia":    "Zla kolejnosc sledzenia (np. od konca zamiast od poczatku)",
	"pominiecie_bazy_rekurencji": "Pominiecie warunku bazowego rekurencji",
	"zly_mnoznik":                "Zly mnoznik lub potega w obliczeniach",
	"brak_tabeli_sledzenia":      "Brak tabeli sledzenia (tabelka krok-po-kroku)",
	"zla_parzystosc_cyfry":       "Zla ocena parzystosci cyfry",
	"bledne_wciecia_blok":        "Bledna interpretacja bloku kodu (wciecia)",
	// TEORIA - projektowanie
	"zly_algorytm":              "Wybor zlego algorytmu do problemu",
	"brak_warunku_stopu":        "Brak warunku stopu w petli/rekurencji",
	"bledna_skladnia_pseudokod": "Bledna skladnia pseudokodu CKE",
	"niepoprawna_petla":         "Niepoprawna petla (zly zakres, zly krok)",
	"brak_inicjalizacji":        "Brak inicjalizacji zmiennej",
	// TEORIA - analiza
	"zla_zlozonosc_klasa":         "Zla klasa zlozonosci (np. O(n) zamiast O(n^2))",
	"brak_uzasadnienia_zlozonosc": "Brak uzasadnienia klasy zlozonosci",
	"mylenie_avg_worst":           "Mylenie zlozonosci sredniej i pesymistycznej",
	"zly_kontrprzyklad":           "Zly kontrprzyklad",
	"brak_wzoru":                  "Brak wzoru lub wyprowadzenia",
	// TEORIA - P/F
	"brak_uzasadnienia_pf":             "Brak uzasadnienia odpowiedzi P/F",
	"mylenie_avg_worst_pf":             "Mylenie zlozonosci sredniej i pesymistycznej (P/F)",
	"nieprecyzyjne_uzasadnienie":       "Nieprecyzyjne uzasadnienie (za ogolne)",
	"pomylenie_stabilnosci_sortowania": "Pomylenie stabilnosci sortowania",
	// TEORIA - konwersja
	"zla_baza_konwersji":      "Zla baza docelowa lub zrodlowa",
	"zla_kolejnosc_reszt":     "Zla kolejnosc reszt (od dolu zamiast od gory)",
	"brak_zapisu_posredniego": "Brak zapisu posredniego obliczen",
	"zle_grupowanie_bitow":    "Zle grupowanie bitow (bin->hex: grupy po 4)",
	"blad_uzupelnienia_do_2":  "Blad w kodzie uzupelnienia do 2 (U2)",
	// TEORIA - bezpieczenstwo
	"mylenie_typow_malware":            "Mylenie typow malware (trojan, worm, ransomware)",
	"mylenie_szyfrowania_sym_asym":     "Mylenie szyfrowania symetrycznego i asymetrycznego",
	"mylenie_protokolow":               "Mylenie protokolow sieciowych",
	"brak_rozroznienia_klucz_pub_pryw": "Brak rozroznienia klucz publiczny/prywatny",
	// SQL (shared)
	"brak_group_by":         "Brak GROUP BY przy agregacji",
	"zly_join_warunek":      "Zly warunek JOIN (ON)",
	"brak_having":           "Brak HAVING (filtrowanie po agregacji)",
	"zla_agregacja":         "Zla funkcja agregujaca (SUM/COUNT/AVG)",
	"null_zamiast_is_null":  "Uzycie = NULL zamiast IS NULL",
	"count_star_vs_kolumna": "Mylenie COUNT(*) z COUNT(kolumna)",
	"zla_kolejnosc_klauzul": "Zla kolejnosc klauzul SQL",
	// IMPLEMENTACJA (shared)
	"zly_warunek_petli":       "Zly warunek petli (< vs <=, itd.)",
	"brak_wczytania":          "Brak wczytania danych z pliku",
	"off_by_one":              "Blad o jeden (indeks, granica petli)",
	"dzielenie_calkowite":     "Dzielenie calkowite zamiast zmiennoprzecinkowego",
	"zle_indeksowanie":        "Zle indeksowanie (od 0 vs od 1)",
	"brak_obslugi_brzegowych": "Brak obslugi przypadkow brzegowych",
	"zla_kolejnosc_operacji":  "Zla kolejnosc operacji",
	// ARKUSZ (shared)
	"zle_adresowanie":          "Zle adresowanie komorek",
	"brak_dolara":              "Brak $ (adresowanie bezwzgledne)",
	"zla_formula_warunkowa":    "Zla formula warunkowa (SUMIFS, COUNTIF)",
	"stala_zamiast_odwolania":  "Stala zamiast odwolania do komorki",
	"brak_kolumny_pomocniczej": "Brak kolumny pomocniczej",
}

// levenshtein computes the Levenshtein edit distance between two strings.
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

// CodeSuggestion represents a suggested error code with its description.
type CodeSuggestion struct {
	Kod  string `json:"kod"`
	Opis string `json:"opis"`
}

// suggestClosestCodes returns up to 3 closest error codes for a given typ,
// sorted by Levenshtein distance. Returns nil if kod is already valid.
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
	var list []scored
	for _, c := range allowed {
		list = append(list, scored{c, levenshtein(kod, c)})
	}
	sort.Slice(list, func(i, j int) bool { return list[i].dist < list[j].dist })

	limit := 3
	if len(list) < limit {
		limit = len(list)
	}
	result := make([]CodeSuggestion, limit)
	for i := 0; i < limit; i++ {
		result[i] = CodeSuggestion{Kod: list[i].code, Opis: errorCodeDescriptions[list[i].code]}
	}
	return result
}
