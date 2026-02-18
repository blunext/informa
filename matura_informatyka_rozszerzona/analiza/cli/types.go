package main

// === JSON import types (match source files exactly) ===

// ExerciseMeta matches _meta.json in exercise directories
type ExerciseMeta struct {
	Typ           string        `json:"typ"`
	Nazwa         string        `json:"nazwa"`
	Kategoria     string        `json:"kategoria"`
	Czestotliwosc string        `json:"czestotliwosc"`
	PunktyLacznie int           `json:"punkty_lacznie"`
	TagiGlobalne  []string      `json:"tagi_globalne"`
	Cwiczenia     []ExerciseRef `json:"cwiczenia"`
}

type ExerciseRef struct {
	ID       string   `json:"id"`
	Trudnosc string   `json:"trudnosc"`
	Punkty   int      `json:"punkty"`
	Tagi     []string `json:"tagi"`
}

// ExerciseJSON matches individual exercise {id}.json files
type ExerciseJSON struct {
	ID          string        `json:"id"`
	Trudnosc    string        `json:"trudnosc"`
	Punkty      int           `json:"punkty"`
	Zrodlo      string        `json:"zrodlo"`
	Tagi        []string      `json:"tagi"`
	Tresc       string        `json:"tresc"`
	Wskazowki   []string      `json:"wskazowki"`
	Odpowiedz   string        `json:"odpowiedz"`
	TypoweBledy []CommonError `json:"typowe_bledy"`
}

type CommonError struct {
	Opis string `json:"opis"`
	Kara string `json:"kara"`
}

// MaturaExam matches matura_YYYY.json top-level
type MaturaExam struct {
	Rok         int        `json:"rok"`
	Formula     string     `json:"formula"`
	CzasMinuty  int        `json:"czas_minuty"`
	TotalPunkty int        `json:"total_punkty"`
	LiczbaZadan int        `json:"liczba_zadan"`
	Czesci      []ExamPart `json:"czesci"`
	Zadania     []ExamTask `json:"zadania"`
}

type ExamPart struct {
	Czesc      int    `json:"czesc"`
	CzasMinuty int    `json:"czas_minuty"`
	Punkty     int    `json:"punkty"`
	Opis       string `json:"opis"`
}

type ExamTask struct {
	Numer         int             `json:"numer"`
	Tytul         string          `json:"tytul"`
	Punkty        int             `json:"punkty"`
	Czesc         int             `json:"czesc"`
	Kontekst      *string         `json:"kontekst"`
	Podzadania    []ExamSubtask   `json:"podzadania"`
	SciezkaDanych *string         `json:"sciezka_danych"`
	PlikiDanych   []string        `json:"pliki_danych"`
	Wymiary       *TaskDimensions `json:"wymiary"`
}

type ExamSubtask struct {
	ID              string   `json:"id"`
	Numer           string   `json:"numer"`
	Punkty          int      `json:"punkty"`
	TypZadania      string   `json:"typ_zadania"`
	Kategoria       string   `json:"kategoria"`
	Tresc           string   `json:"tresc"`
	Odpowiedz       string   `json:"odpowiedz"`
	ZasadyOceniania string   `json:"zasady_oceniania"`
	Pulapki         []string `json:"pulapki"`
}

type TaskDimensions struct {
	Tematyczny    []string `json:"tematyczny"`
	Algorytmiczny []string `json:"algorytmiczny"`
	CzasowyMin    int      `json:"czasowy_min"`
}

// === Output types (returned as JSON on stdout) ===

// ExerciseOut is what exercise get returns
type ExerciseOut struct {
	ID          string        `json:"id"`
	TypNazwa    string        `json:"typ_nazwa"`
	Kategoria   string        `json:"kategoria"`
	Trudnosc    string        `json:"trudnosc"`
	Punkty      int           `json:"punkty"`
	Zrodlo      string        `json:"zrodlo"`
	Tagi        []string      `json:"tagi"`
	Tresc       string        `json:"tresc"`
	Wskazowki   []string      `json:"wskazowki"`
	Odpowiedz   string        `json:"odpowiedz"`
	TypoweBledy []CommonError `json:"typowe_bledy"`
}

// ReviewOut is what exercise review returns
type ReviewOut struct {
	Exercise    ExerciseOut `json:"exercise"`
	Tag         string      `json:"tag"`
	DaysOverdue int         `json:"days_overdue"`
}

// ProgressUpdateOut is what progress update returns
type ProgressUpdateOut struct {
	ID              string   `json:"id"`
	NewLevel        string   `json:"new_level"`
	Streak          int      `json:"streak"`
	TagsUpdated     []string `json:"tags_updated"`
	NextReviewDates []string `json:"next_review_dates"`
	CzasSek         *int     `json:"czas_sek,omitempty"`
	BenchmarkSek    *int     `json:"benchmark_sek,omitempty"`
	Tempo           *string  `json:"tempo,omitempty"`
	BladWarning     *string  `json:"blad_warning,omitempty"`
}

// ProgressStatusOut is what progress status returns
type ProgressStatusOut struct {
	Sesje              int                  `json:"sesje"`
	OstatniaSesja      string               `json:"ostatnia_sesja"`
	CwiczeniaLacznie   int                  `json:"cwiczenia_lacznie"`
	PerTyp             []TypStatusOut       `json:"per_typ"`
	PerKategoria       []KategoriaStatusOut `json:"per_kategoria"`
	Zaleglosci         int                  `json:"zaleglosci"`
	TagiOpanowane      []string             `json:"tagi_opanowane"`
	TagiProblematyczne []string             `json:"tagi_problematyczne"`
}

type TypStatusOut struct {
	Typ             string `json:"typ"`
	PoziomTrudnosci string `json:"poziom_trudnosci"`
	Streak          int    `json:"streak"`
	Zrobione        int    `json:"zrobione"`
	Dostepne        int    `json:"dostepne"`
	AvgCzasSek      *int   `json:"avg_czas_sek,omitempty"`
	BenchmarkSek    *int   `json:"benchmark_sek,omitempty"`
}

// CKEOut is what cke get returns
type CKEOut struct {
	ID              string   `json:"id"`
	Rok             int      `json:"rok"`
	NumerZadania    int      `json:"numer_zadania"`
	Tytul           string   `json:"tytul"`
	Kontekst        string   `json:"kontekst"`
	TypZadania      string   `json:"typ_zadania"`
	Kategoria       string   `json:"kategoria"`
	Punkty          int      `json:"punkty"`
	Tresc           string   `json:"tresc"`
	Odpowiedz       string   `json:"odpowiedz"`
	ZasadyOceniania string   `json:"zasady_oceniania"`
	Pulapki         []string `json:"pulapki"`
	SciezkaDanych   string   `json:"sciezka_danych"`
	PlikiDanych     []string `json:"pliki_danych"`
}

// ExamMetaOut is what exam meta returns
type ExamMetaOut struct {
	Rok         int             `json:"rok"`
	Formula     string          `json:"formula"`
	CzasMinuty  int             `json:"czas_minuty"`
	TotalPunkty int             `json:"total_punkty"`
	LiczbaZadan int             `json:"liczba_zadan"`
	Czesci      []ExamPart      `json:"czesci"`
	Zadania     []ExamTaskBrief `json:"zadania"`
}

type ExamTaskBrief struct {
	Numer  int    `json:"numer"`
	Tytul  string `json:"tytul"`
	Punkty int    `json:"punkty"`
	Czesc  int    `json:"czesc"`
}

// ExamTaskOut is what exam task returns (with full content)
type ExamTaskOut struct {
	Numer         int              `json:"numer"`
	Tytul         string           `json:"tytul"`
	Punkty        int              `json:"punkty"`
	Kontekst      string           `json:"kontekst"`
	Podzadania    []ExamSubtaskOut `json:"podzadania"`
	SciezkaDanych string           `json:"sciezka_danych"`
	PlikiDanych   []string         `json:"pliki_danych"`
}

type ExamSubtaskOut struct {
	ID              string   `json:"id"`
	Numer           string   `json:"numer"`
	Punkty          int      `json:"punkty"`
	TypZadania      string   `json:"typ_zadania"`
	Kategoria       string   `json:"kategoria"`
	Tresc           string   `json:"tresc"`
	Odpowiedz       string   `json:"odpowiedz"`
	ZasadyOceniania string   `json:"zasady_oceniania"`
	Pulapki         []string `json:"pulapki"`
}

// TrapOut is what trap list returns
type TrapOut struct {
	SourceID string   `json:"source_id"`
	Rok      int      `json:"rok"`
	Tytul    string   `json:"tytul"`
	Pulapki  []string `json:"pulapki"`
}

// TrapSaveOut is what trap save returns
type TrapSaveOut struct {
	ID        string `json:"id"`
	Typ       string `json:"typ"`
	Trafienia int    `json:"trafienia"`
	Total     int    `json:"total"`
}

// DataStatsOut is what data stats returns
type DataStatsOut struct {
	Cwiczenia   int `json:"cwiczenia"`
	Podzadania  int `json:"podzadania"`
	Cheatsheets int `json:"cheatsheets"`
}

// ExerciseNextOut is what exercise next returns
type ExerciseNextOut struct {
	Mode           string      `json:"mode"`
	Exercise       ExerciseOut `json:"exercise"`
	ReviewTag      *string     `json:"review_tag"`
	DaysOverdue    *int        `json:"days_overdue"`
	PoolWarning    *string     `json:"pool_warning"`
	SessionCount   int         `json:"session_count"`
	ResetSuggested bool        `json:"reset_suggested"`
	ChosenTyp      string      `json:"chosen_typ,omitempty"`
}

// TypIntroOut is what typ intro returns
type TypIntroOut struct {
	Typ                      string         `json:"typ"`
	Kategoria                string         `json:"kategoria"`
	FirstInType              bool           `json:"first_in_type"`
	FirstInCategory          bool           `json:"first_in_category"`
	OtherTypesDoneInCategory []string       `json:"other_types_done_in_category"`
	Level                    string         `json:"level"`
	Streak                   int            `json:"streak"`
	Available                int            `json:"available"`
	Done                     int            `json:"done"`
	SprawdzianUnlocked       bool           `json:"sprawdzian_unlocked"`
	CKEStats                 *CKETypStats   `json:"cke_stats,omitempty"`
	TopPulapki               []string       `json:"top_pulapki,omitempty"`
	Przyklad                 *ExerciseBrief `json:"przyklad,omitempty"`
	CheatsheetExcerpt        string         `json:"cheatsheet_excerpt,omitempty"`
}

type CKETypStats struct {
	Wystapienia int     `json:"wystapienia"`
	LatTotal    int     `json:"lat_total"`
	AvgPunkty   float64 `json:"avg_punkty"`
	TotalPunkty int     `json:"total_punkty"`
}

type ExerciseBrief struct {
	ID        string `json:"id"`
	Tresc     string `json:"tresc"`
	Odpowiedz string `json:"odpowiedz"`
}

// BladOut is what progress blad returns
type BladOut struct {
	ID         int    `json:"id"`
	ExerciseID string `json:"exercise_id"`
	Typ        string `json:"typ"`
	BladKod    string `json:"blad_kod"`
	BladOpis   string `json:"blad_opis"`
	HintLevel  int    `json:"hint_level"`
	Data       string `json:"data"`
}

type DiagnoseEntry struct {
	BladKod  string   `json:"blad_kod"`
	Count    int      `json:"count"`
	Typy     []string `json:"typy"`
	Ostatnio string   `json:"ostatnio"`
}

type DiagnoseOut struct {
	TopBledy []DiagnoseEntry `json:"top_bledy"`
	Total    int             `json:"total"`
}

// KategoriaStatusOut is per-category status
type KategoriaStatusOut struct {
	Kategoria   string  `json:"kategoria"`
	TypyTotal   int     `json:"typy_total"`
	TypyRuszane int     `json:"typy_ruszane"`
	Zrobione    int     `json:"zrobione"`
	Dostepne    int     `json:"dostepne"`
	AvgStreak   float64 `json:"avg_streak"`
}

// CKEStatusOut is what cke status returns
type CKEStatusOut struct {
	Typ          string `json:"typ"`
	Kategoria    string `json:"kategoria"`
	Unlocked     bool   `json:"unlocked"`
	Level        string `json:"level"`
	CKEDone      int    `json:"cke_done"`
	CKEAvailable int    `json:"cke_available"`
}

// ExamListEntry is one year in exam list
type ExamListEntry struct {
	Rok          int      `json:"rok"`
	Formula      string   `json:"formula"`
	CzasMin      int      `json:"czas_min"`
	TotalPkt     int      `json:"total_pkt"`
	Done         bool     `json:"done"`
	WynikProcent *float64 `json:"wynik_procent,omitempty"`
}

// ExamListOut is what exam list returns
type ExamListOut struct {
	Available []ExamListEntry `json:"available"`
	Suggested *ExamSuggestion `json:"suggested,omitempty"`
}

// ExamSuggestion is a suggested exam year
type ExamSuggestion struct {
	Rok     int    `json:"rok"`
	Formula string `json:"formula"`
	Reason  string `json:"reason"`
}
