# `/matura` — interaktywny korepetytor do matury rozszerzonej z informatyki

Skill do [Claude Code](https://docs.claude.com/en/docs/claude-code/quickstart), który prowadzi Cię przez maturę rozszerzoną z informatyki metodą sokratejską. Oparty o analizę 12 lat prawdziwych egzaminów CKE (2014–2025) — 30 sesji, 641 podzadań, 937 ćwiczeń treningowych.

Dla zdających **maturę rozszerzoną z informatyki** w formule 2023+ (210 min / 50 pkt).

---

## TL;DR

```bash
git clone git@github.com:blunext/informa.git
cd informa
claude        # uruchom Claude Code w tym katalogu
```

Potem w Claude Code wpisz `/matura`. Korepetytor sam przeprowadzi Cię przez resztę.

---

## Co dostajesz

- **Korepetytor 1-na-1** — prowadzi sokratejsko: pytaniami naprowadza, nie podaje gotowych odpowiedzi.
- **937 ćwiczeń treningowych** w 23 typach × 4 kategorie. Z 3-poziomowymi hintami i analizą typowych błędów CKE.
- **641 prawdziwych zadań CKE** z lat 2014–2025 jako sprawdziany i próbne matury.
- **Powtórki rozłożone w czasie** (algorytm FSRS-5) — uczy się Twojego tempa zapamiętywania. Nie zapominasz tego, co już opanowałeś.
- **Automatyczna progresja trudności** — łatwe → średnie → średnie-trudne → trudne. Awansujesz dopiero gdy faktycznie umiesz.
- **TIER 1 algorytmów = 93.2% punktów** na maturze (16 algorytmów z analizy 12 lat). System priorytetyzuje to, co naprawdę liczy się na egzaminie.
- **Pełne symulacje próbnej matury** z dowolnego roku 2014–2025, z oceną wg oficjalnych kryteriów CKE.

---

## Wymagania

- [Claude Code](https://docs.claude.com/en/docs/claude-code/quickstart) (CLI w terminalu lub IDE)
- Konto z subskrypcją Claude (Pro/Team/Enterprise) — szczegóły w [dokumentacji Claude Code](https://docs.claude.com/en/docs/claude-code/overview)
- Git

---

## Wspierane systemy

Korepetytor `/matura` pod spodem wywołuje binarkę `matura` (Go) — to ona trzyma bazę ćwiczeń, postępu i obsługuje powtórki FSRS-5. Repo zawiera gotowe binarki dla dwóch platform; reszta wymaga jednorazowej kompilacji (~30 sekund).

| System | Status | Co robić |
|--------|--------|----------|
| **macOS Apple Silicon** (M1/M2/M3/M4) | ✓ działa od ręki | nic, binarka `matura` (Mach-O arm64) jest w repo |
| **Windows x86-64** | ✓ działa od ręki | nic, binarka `matura.exe` (PE32+) jest w repo |
| **macOS Intel (x86-64)** | wymaga kompilacji | patrz "Kompilacja z źródła" poniżej |
| **Linux** (dowolna architektura) | wymaga kompilacji | patrz "Kompilacja z źródła" poniżej |

### Kompilacja z źródła

Wymagany [Go 1.26+](https://go.dev/doc/install). 

**Linux / macOS Intel:**

```bash
# 1. Zainstaluj Go (jeśli nie masz):
#    Linux:    https://go.dev/doc/install  lub  apt install golang  /  pacman -S go  /  brew install go (Linuxbrew)
#    macOS:    brew install go
go version   # sprawdź, że masz >= 1.26

# 2. Sklonuj repo i zbuduj:
git clone git@github.com:blunext/informa.git
cd informa/matura_informatyka_rozszerzona/analiza/cli
./build.sh   # buduje matura (dla Twojego systemu) + matura.exe + reimportuje baze
```

`build.sh` jednocześnie cross-compiluje wersję Windows i robi reimport bazy z plików JSON do `matura.db`. Jeśli nie chcesz cross-compilacji, ręcznie:

```bash
cd matura_informatyka_rozszerzona/analiza/cli
CGO_ENABLED=0 go build -o matura .
```

Binarka wyląduje obok już istniejącej `matura` z repo — **nadpisz ją** swoją wersją dla Twojego systemu. Po tym `/matura` w Claude Code zacznie działać.

**Weryfikacja:**

```bash
./matura data stats
# Powinno wypisać: 937 cwiczen, 641 zadan CKE, 4 cheatsheety
```

Jeśli widzisz statystyki — gotowe, wracaj do [Szybkiego startu](#szybki-start).

---

## Szybki start

1. **Zainstaluj Claude Code** zgodnie z [oficjalną instrukcją](https://docs.claude.com/en/docs/claude-code/quickstart).
2. **Sklonuj repo**:
   ```bash
   git clone git@github.com:blunext/informa.git
   cd informa
   ```
3. **Uruchom Claude Code** w katalogu repo:
   ```bash
   claude
   ```
4. **Wpisz `/matura`** — korepetytor przywita Cię i zapyta, od czego chcesz zacząć.

Skill znajduje się w `.claude/skills/matura/` — Claude Code wykryje go automatycznie, nic dodatkowo nie konfigurujesz.

**Skróty komend startowych:**

- `/matura` — pełny powitalny flow (wybór kategorii)
- `/matura SQL` — od razu do bloku SQL
- `/matura cyfry_liczby` — od razu do konkretnego typu

---

## Jak wygląda egzamin

**Formuła 2023+ (aktualna):**

- 210 minut (3.5 h), 50 punktów, 7–8 zadań, jeden arkusz.

**Punkty per kategoria (typowo):**

| Kategoria | Punkty | Narzędzie |
|-----------|--------|-----------|
| Teoria (algorytmy, śledzenie, P/F, systemy liczbowe) | ~10 | kartka + długopis |
| Implementacja (C++) | ~20 | Code::Blocks / Dev-C++ |
| Arkusz kalkulacyjny | ~10 | LibreOffice Calc / Excel |
| SQL (bazy danych) | ~10 | MS Access |

Stara formuła 2015–2022 (Część I + Część II) też w pełni pokryta w bazie zadań CKE.

---

## Jak działa korepetytor

### 4 tryby

| Tryb | Kiedy używać | Komenda |
|------|--------------|---------|
| **Ćwiczenia** | Główna nauka — codzienna praktyka | `/matura` lub `/matura {typ}` |
| **Sprawdzian typu** | Test gotowości na prawdziwym zadaniu CKE, bez hintów | `sprawdzian sql_group_by` |
| **Próbna matura** | Pełna symulacja egzaminu z wybranego roku | `probna 2024` |
| **Pułapki CKE** | Quiz z najczęstszych błędów na maturze | `pulapki SQL` |

Sprawdzian odblokowuje się dopiero po osiągnięciu poziomu "trudne" w danym typie. Próbna matura podaje zadania sekwencyjnie i na końcu ocenia wg oficjalnych kryteriów CKE.

### System hintów (3 poziomy)

Hinty pojawiają się **dopiero po Twojej próbie**. Korepetytor fizycznie blokuje hint bez wcześniejszej odpowiedzi — żeby nie skracać myślenia.

- **Poziom 1**: pytanie sokratejskie ("A co gdyby w kroku 3 było `n = 0`?")
- **Poziom 2**: cytat z odpowiedniego cheatsheet + konkretne pojęcie
- **Poziom 3**: kluczowy krok rozwiązania — ostatni krok zostawiasz dla siebie

Po 3 błędnych próbach lub na komendę `poddaje sie` dostajesz pełną odpowiedź + analizę typowych pułapek CKE.

### Powtórki w czasie (FSRS-5)

**FSRS-5** (Free Spaced Repetition Scheduler, wersja 5) to otwarty algorytm powtórek rozłożonych w czasie — następca SuperMemo/Anki SM-2. Zamiast jednego ogólnego harmonogramu, modeluje **każdą umiejętność osobno** trzema parametrami:

- **Stability (stabilność)** — jak długo pamiętasz dany temat zanim zaczniesz zapominać.
- **Difficulty (trudność)** — jak bardzo dany temat jest dla Ciebie wymagający.
- **Retrievability (szansa przypomnienia)** — prawdopodobieństwo, że poprawnie przypomnisz sobie temat w danym momencie.

Po każdej Twojej odpowiedzi algorytm aktualizuje te parametry i wylicza następną datę powtórki tak, żeby trafiła w moment kiedy temat już *prawie* zaczynasz zapominać — to optymalny punkt dla utrwalenia pamięci długotrwałej.

W praktyce: szybko zapominasz dany temat → częstsze powtórki. Opanowałeś → coraz rzadsze. Algorytm uczy się Twojego indywidualnego tempa zapamiętywania.

Więcej: [open-spaced-repetition.github.io/fsrs4anki](https://open-spaced-repetition.github.io/fsrs4anki/).

| Poziom tagu | Nazwa | Przybliżony interwał |
|-------------|-------|----------------------|
| 0 | NOWE | natychmiast |
| 1 | UCZĘ SIĘ | ~1–2 dni |
| 2 | ĆWICZĘ | ~3–6 dni |
| 3 | PEWNIE | ~7–20 dni |
| 4 | OPANOWANE | 21+ dni |

Interwały są indywidualne — FSRS dostosowuje je do Twojej historii.

### Progresja trudności

Łatwe → średnie → średnie-trudne → trudne. 3 czyste odpowiedzi z rzędu = awans o poziom. Walk-through (poddanie się) = cofnięcie o 1.

### Komendy w trakcie sesji

| Komenda | Działanie |
|---------|-----------|
| `wskazowka` | Następny poziom hintu (1 → 2 → 3 → pełna odpowiedź) |
| `poddaje sie` | Pełna odpowiedź + analiza błędów |
| `wyjasniej [temat]` | Wyjaśnienie z cheatsheet metodą sokratejską |
| `nastepny` / `dalej` | Następne ćwiczenie |
| `zmien temat` | Wybór innego typu/kategorii |
| `status` | Dashboard: retencja, zaległości, słabe punkty |
| `powtorka` | Zaległe powtórki |
| `sprawdzian [typ]` | Prawdziwe zadanie CKE |
| `probna [rok]` | Symulacja egzaminu |
| `pulapki [typ]` | Quiz z pułapek |
| `strategia` | Porady egzaminacyjne (kolejność, czas) |

Możesz też po prostu rozmawiać z korepetytorem naturalnie — zrozumie kontekst.

---

## 23 typy zadań — czego się uczysz

### TEORIA (6 typów)

| Typ | Opis |
|-----|------|
| `sledzenie_algorytmu` | Wykonujesz algorytm krok po kroku, podajesz wynik |
| `projektowanie_algorytmu` | Piszesz algorytm/pseudokod rozwiązujący problem |
| `analiza_algorytmu` | Określasz złożoność, własności, poprawność |
| `test_prawda_falsz` | Oceniasz zdania o algorytmach jako P lub F |
| `konwersja_systemow_liczbowych` | Zamieniasz między bin/dec/hex |
| `teoria_bezpieczenstwa` | Pytania o szyfrowanie, bezpieczeństwo danych |

### IMPLEMENTACJA (8 typów, C++)

| Typ | Opis |
|-----|------|
| `cyfry_liczby` | Rozkład na cyfry, dzielniki, mod/div |
| `napisy` | Przetwarzanie tekstów, palindromy, wzorce |
| `zlozone` | Algorytmy łączące kilka technik |
| `zliczanie` | Ile elementów spełnia warunek |
| `minmax` | Szukanie ekstremów |
| `sekwencje` | Najdłuższe ciągi, serie, podciągi |
| `obrazy_2D` | Operacje na tablicach 2D (mapy, bitmapy) |
| `geometryczne` | Odległości, pola, współrzędne |

### ARKUSZ (5 typów)

| Typ | Opis |
|-----|------|
| `agregacja_warunkowa` | SUMA.WARUNKÓW, LICZ.WARUNKI |
| `symulacja` | Modelowanie procesu krok po kroku |
| `wykres` | Wykres z osiami, legendą, tytułem |
| `agregacja_podstawowa` | SUMA, ŚREDNIA, ILE.LICZB, MIN, MAX |
| `transformacja` | WYSZUKAJ.PIONOWO, LEWY, FRAGMENT.TEKSTU |

### SQL (4 typy)

| Typ | Opis |
|-----|------|
| `sql_group_by` | Grupowanie z funkcjami agregującymi |
| `sql_podzapytania` | Zapytania zagnieżdżone (IN, EXISTS) |
| `sql_join` | Łączenie tabel (INNER, LEFT, samozłączenie) |
| `sql_select_where` | Filtrowanie, sortowanie, proste zapytania |

---

## Ranking algorytmów CKE 2014–2025

Z analizy **641 podzadań w 30 sesjach CKE** (1500 pkt łącznie, 1456 wystąpień tagów). Korepetytor priorytetyzuje cwiczenia wg tego rankingu — dzięki temu uczysz się tego, co naprawdę liczy się na egzaminie.

### TOP 5 (wg liczby wystąpień)

| # | Algorytm | Kategoria | Wystąpień | Punktów | Lat |
|---|---|---|---:|---:|---:|
| 1 | `iteracja-po-pliku` | wzorce | 150 | 435 | 12/12 |
| 2 | `SQL-JOIN` | wzorce | 85 | 191 | 12/12 |
| 3 | `SQL-aggregacja` | wzorce | 85 | 177 | 11/12 |
| 4 | `akumulator-licznik` | wzorce | 74 | 212 | 12/12 |
| 5 | `sledzenie-pseudokod` | wzorce | 74 | 136 | 12/12 |

### TIER 1 — Must Have (16 algorytmów)

**Znajomość TIER 1 pozwala podejść do 575/641 podzadań (89.7%) i zdobyć 1398/1500 punktów (93.2%).** Te algorytmy musisz znać na 100%.

| Algorytm | Kategoria | Wystąpień | Lat |
|---|---|---:|---:|
| `iteracja-po-pliku` | wzorce | 150 | 12/12 |
| `SQL-JOIN` | wzorce | 85 | 12/12 |
| `SQL-aggregacja` | wzorce | 85 | 11/12 |
| `akumulator-licznik` | wzorce | 74 | 12/12 |
| `sledzenie-pseudokod` | wzorce | 74 | 12/12 |
| `arkusz-agregacja-warunkowa` | wzorce | 72 | 12/12 |
| `SQL-GROUP-BY` | wzorce | 72 | 11/12 |
| `current-max` | wzorce | 68 | 12/12 |
| `SQL-WHERE` | wzorce | 64 | 12/12 |
| `iteracja-po-cyfrach` | wzorce | 59 | 11/12 |
| `konwersja-systemow` | klasyczne | 56 | 11/12 |
| `arkusz-symulacja-iteracyjna` | wzorce | 53 | 11/12 |
| `rekurencja` | techniki | 48 | 12/12 |
| `SQL-ORDER-BY` | wzorce | 43 | 11/12 |
| `analiza-zlozonosci` | techniki | 41 | 12/12 |
| `znajdz-i-policz` | wzorce | 32 | 11/12 |

### TIER 2 — Powinno się znać (16 algorytmów)

Razem TIER 1 + 2 = 590/641 podzadań (92.0%) i 1435/1500 punktów (95.7%).

`porownywanie-tekstow`, `akumulator-suma`, `wyszukiwanie-liniowe`, `current-min`, `SQL-podzapytanie-niezalezne`, `SQL-LIKE`, `tablica-2D`, `Horner`, `SQL-HAVING`, `SQL-LEFT-JOIN-NULL`, `najdluzszy-podciag-niemalejacy`, `prefix-sum`, `przeszukiwanie-binarne`, `drzewo`, `SQL-funkcje-tekstowe`, `bisekcja`.

### TIER 3 — Nice to have (29 algorytmów)

Pojawiają się rzadko (1–9 wystąpień). Pełna lista w [`analiza/RANKING_ALGORYTMOW.md`](matura_informatyka_rozszerzona/analiza/RANKING_ALGORYTMOW.md).

### TOP 10 kombinacji 2-algorytmowych

Pary algorytmów, które najczęściej występują razem w jednym podzadaniu — jeśli umiesz jedno, naucz się też drugiego.

| # | Algorytm A | Algorytm B | Wspólnie |
|---|---|---|---:|
| 1 | `SQL-GROUP-BY` | `SQL-aggregacja` | 64 |
| 2 | `SQL-JOIN` | `SQL-aggregacja` | 61 |
| 3 | `akumulator-licznik` | `iteracja-po-pliku` | 56 |
| 4 | `SQL-GROUP-BY` | `SQL-JOIN` | 55 |
| 5 | `current-max` | `iteracja-po-pliku` | 43 |
| 6 | `SQL-JOIN` | `SQL-WHERE` | 40 |
| 7 | `SQL-WHERE` | `SQL-aggregacja` | 36 |
| 8 | `SQL-JOIN` | `SQL-ORDER-BY` | 31 |
| 9 | `iteracja-po-cyfrach` | `iteracja-po-pliku` | 30 |
| 10 | `iteracja-po-pliku` | `znajdz-i-policz` | 27 |

### Główne wnioski edukacyjne

1. **Praca z plikiem dominuje** — `iteracja-po-pliku` jest w prawie każdej sesji CKE. Umiejętność niezbędna.
2. **SQL ma własny TIER 1** — `SQL-JOIN`, `SQL-aggregacja`, `SQL-GROUP-BY`, `SQL-WHERE` to filary zadania bazodanowego (i jak widać w kombinacjach — występują razem).
3. **Śledzenie pseudokodu** równie częste co programowanie — wymagana zarówno teoria, jak i praktyka.
4. **Konwersje systemów liczbowych** + **Horner** występują regularnie — fundamenty teorii.
5. **Programowanie dynamiczne** pojawiło się pierwszy raz w 2024 — nowy obszar wymagający uwagi.

### Algorytmy z podstawy programowej NIE testowane przez CKE (2014–2025)

Wymienione w podstawie, ale nigdy nie pojawiły się na maturze — **niski priorytet nauki**:

- `fraktale-rekurencyjne` (drzewo binarne, płatek Kocha, dywan Sierpińskiego)
- `metoda-wstepujaca-zstepujaca` (top-down vs bottom-up)
- `najdluzszy-wspolny-podciag` (LCS, programowanie dynamiczne)
- `podciag-najwieksza-suma` (algorytm Kadane'a)

Pełny ranking z heatmapą rok × algorytm, rozbiciem per kategoria i podstawą programową: [`analiza/RANKING_ALGORYTMOW.md`](matura_informatyka_rozszerzona/analiza/RANKING_ALGORYTMOW.md). W trakcie sesji wpisz `ranking algorytmow` — korepetytor pokaże dokument.

---

## Pułapki specyficzne dla narzędzi egzaminacyjnych

Korepetytor ostrzega o nich automatycznie, ale warto je znać z góry.

**SQL w MS Access** — dialekt JET/ACE różni się od SQLite/MySQL:
- `Mid` / `Len` zamiast `SUBSTR` / `LENGTH`
- `&` zamiast `||` (konkatenacja)
- `TOP` zamiast `LIMIT`
- `Nz` zamiast `COALESCE`
- daty: `#2024-01-15#` zamiast `'2024-01-15'`
- JOIN 3+ tabel wymaga nawiasów

**Arkusz Excel PL** — polskie nazwy funkcji:
- `SUMA.WARUNKÓW` zamiast `SUMIFS`
- `LICZ.WARUNKI` zamiast `COUNTIFS`
- `WYSZUKAJ.PIONOWO` zamiast `VLOOKUP`
- `FRAGMENT.TEKSTU` zamiast `MID`

**C++** — egzamin w Code::Blocks / Dev-C++. Standard C++11 powinien działać; przed egzaminem sprawdź swój kod na tym środowisku.

W trakcie sesji wpisz `pulapki SQL` (lub `pulapki ARKUSZ`, `pulapki TEORIA`) żeby przećwiczyć rozpoznawanie.

---

## FAQ

**Jak zresetować postęp?**
Usuń `matura_informatyka_rozszerzona/analiza/cli/matura_progress.db`. Przy następnym `/matura` zaczniesz od zera.

**Jak wrócić do sesji po przerwie?**
Wpisz `/matura`. Korepetytor wczyta postęp z bazy i zaproponuje zaległe powtórki lub kontynuację.

**Czy mogę ćwiczyć kilka typów na raz?**
Tak. Komenda `zmien temat` w trakcie sesji + automatyczny *interleaving* (co 3. ćwiczenie korepetytor wplata powtórkę z innego typu — to utrwala wiedzę).

**Co jeśli skończą mi się ćwiczenia w danym typie?**
Korepetytor ostrzeże gdy zostaną 2 ostatnie. Możesz dogenerować nowe komendą `/generate-exercises`.

**Czy korepetytor pamięta moje wyniki między sesjami?**
Tak — postęp jest w `matura_progress.db`. Możesz przerwać i wrócić nawet po tygodniach.

**Ile czasu zajmie przygotowanie się do matury z `/matura`?**
Zależy od poziomu startu. Komenda `strategia` w trakcie sesji pokazuje 3 warianty planu nauki (3+ tygodnie / <7 dni / <24h).

---

## Inne skille w repo

- **`/generate-exercises`** — generowanie nowych ćwiczeń gdy w jakimś typie zaczyna brakować.
- **`/test-tutor`** — testowanie jakości pedagogicznej skilla (dla osób ulepszających `/matura`).

---

## Źródła i licencja

**Materiały CKE** (arkusze, klucze, dane do zadań) pochodzą z [arkusze.pl](https://arkusze.pl/informatyka-matura-poziom-rozszerzony/) — oficjalne źródło. Są własnością Centralnej Komisji Egzaminacyjnej, używane edukacyjnie.

**Kod własny** (skille, baza ćwiczeń, CLI, analiza): repo [`github.com/blunext/informa`](https://github.com/blunext/informa).

**Brak gwarancji.** To nieoficjalne narzędzie do nauki — nie zastępuje oficjalnej dokumentacji CKE ani konsultacji z nauczycielem.

---

## Powodzenia na egzaminie

Wpisz `/matura` i zaczynaj. Reszta sama się ułoży.
