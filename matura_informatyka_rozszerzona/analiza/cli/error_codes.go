package main

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
