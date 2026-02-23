# Jak uzywac korepetytora `/matura`

> Masz do dyspozycji interaktywnego korepetytora, ktory poprowadzi Cie przez
> cwiczenia, sprawdziany i probne matury. Wszystko oparte na 11 latach
> prawdziwych egzaminow CKE (2014-2025). Ponizej znajdziesz opis wszystkich
> trybow i komend.

---

## Szybki start

1. Wpisz `/matura` — korepetytor przywita Cie i zapyta, od czego chcesz zaczac.
2. Wpisz `/matura SQL` — od razu przechodzisz do cwiczen z SQL.
3. Wpisz `/matura 07_cyfry_liczby` — od razu przechodzisz do konkretnego typu.

Jesli to Twoja pierwsza sesja, korepetytor krotko przedstawi 4 bloki tematyczne
i poprosi Cie o wybor. Jesli juz wczesniej cwiczyles — przypomni Twoj postep
i zaleglosci powtorkowe.

---

## Tryby nauki

### 1. Cwiczenia (tryb domyslny)

To glowny tryb pracy. Korepetytor wybiera cwiczenia dopasowane do Twojego
poziomu i prowadzi Cie metoda sokratejska — nie podaje gotowych odpowiedzi,
tylko naprowadza pytaniami.

- **583 cwiczen** z 23 typow zadan, pogrupowanych w 4 kategorie
- **System hintow** — 3 poziomy wskazowek:
  - Poziom 1: pytanie sokratejskie naprowadzajace na poprawny tok myslenia
  - Poziom 2: cytat z odpowiedniej sciagawki + konkretne pojecie
  - Poziom 3: kluczowy krok rozwiazania — zostajesz z ostatnim krokiem
- **Jesli nie umiesz** — po 3 probach dostajesz pelne rozwiazanie + analize
  typowych bledow CKE
- **Trudnosc rosnie automatycznie**: latwe → srednie → srednie-trudne → trudne

Korepetytor sam wybiera kolejnosc cwiczen — priorytet maja zaleglosci
powtorkowe, potem nowy material na Twoim aktualnym poziomie.

### 2. Sprawdzian typu

Test gotowosci egzaminacyjnej na prawdziwym zadaniu z matury CKE.

- **Odblokowany** po osiagnieciu poziomu "trudne" w danym typie (3 poprawne
  odpowiedzi bez pomocy na poziomie srednie-trudne)
- **Prawdziwe zadanie CKE** z archiwum (nie cwiczenie!)
- **BEZ hintow** — jak na egzaminie
- **Ocena czesciowa** wg oficjalnych zasad CKE (np. "1 pkt za poprawne
  zapytanie, 1 pkt za wynik")
- Po ocenie widzisz pulapki CKE ukryte w zadaniu

Komenda: `sprawdzian sql_group_by` lub `sprawdzian 07_cyfry_liczby`

Bez argumentu: lista odblokowanych typow z informacja ile zadan masz do
dyspozycji.

### 3. Probna matura

Pelna symulacja egzaminu z wybranego roku — wszystkie zadania po kolei.

- **11 lat** do wyboru: 2014-2019, 2021-2025
- Zadania podawane sekwencyjnie, jedno po drugim
- **BEZ hintow** — jak na prawdziwym egzaminie
- Mozesz pominac zadanie (`pomin`) lub przerwac egzamin (`przerwij`)
- **Na koncu**: wynik X/Y pkt, analiza per kategoria, lista pulapek CKE,
  na ktore wpadlas, mocne strony i obszary do poprawy
- Mozesz przejrzec pelne rozwiazania po zakonczeniu

Komendy:
- `probna 2024` — konkretny rok
- `probna nowa` — losowy rok z nowej formuly (2023-2025)
- `probna stara` — losowy rok ze starej formuly (2015-2022)
- `probna losowa` — dowolny losowy rok
- `probna` (bez argumentu) — lista dostepnych lat

### 4. Pulapki CKE

Trening rozpoznawania typowych pulapek z prawdziwych egzaminow. Pulapki to
najczestszy powod utraty punktow na maturze.

- **Tryb quizowy**: korepetytor pokazuje zadanie, Ty szukasz pulapek. Potem
  porownanie z oficjalnymi pulapkami CKE
- **Tryb przegladowy**: zestawienie wszystkich pulapek dla danego typu lub
  kategorii

Komendy:
- `pulapki sql_group_by` — quiz z pulapek danego typu
- `pulapki TEORIA` — quiz z pulapek calej kategorii
- `pulapki lista SQL` — zestawienie (bez quizu)
- `pulapki` (bez argumentu) — pulapki z aktualnie cwiczonego typu

---

## Komendy w trakcie sesji

W dowolnym momencie mozesz wpisac jedna z ponizszych komend. Mozesz tez
rozmawiac naturalnie — korepetytor zrozumie.

| Komenda | Co robi |
|---------|---------|
| `wskazowka` | Nastepny poziom hintu (1 → 2 → 3 → pelna odpowiedz) |
| `poddaje sie` | Pelna odpowiedz + analiza bledow. Trudnosc spada o 1 |
| `wyjasniej [temat]` | Wyjasnienie tematu z odpowiedniej sciagawki |
| `nastepny` / `dalej` | Przejdz do nastepnego cwiczenia |
| `zmien temat` | Wybor innego bloku lub typu zadan |
| `podsumowanie` | Ile cwiczen w tej sesji, jakie wyniki |
| `strategia` | Porady egzaminacyjne (kolejnosc zadan, zarzadzanie czasem) |
| `powtorka` | Pokaz zaleglosci powtorkowe |
| `status` | Pelny raport: typy, poziomy, tagi opanowane i problematyczne |
| `sprawdzian [typ]` | Prawdziwe zadanie CKE (wymaga poziomu "trudne") |
| `probna [rok]` | Symulacja pelnego egzaminu maturalnego |
| `pulapki [typ]` | Quiz z pulapek CKE |

---

## Jak dziala postep

- **Wszystko sie zapisuje** — Twoj postep jest w pliku, wiec mozesz przerwac
  i wrocic pozniej. Korepetytor pamięta, gdzie skonczylismy.
- **Powtorki rozlozone w czasie** (FSRS-5) — algorytm automatycznie dostosowuje
  interwaly powtorkowe do Twojego tempa nauki. Jesli szybko zapominasz dany
  temat — powtorki beda czestsze. Jesli opanowales go dobrze — coraz rzadsze.
- **Tagi umiejetnosci** — kazde cwiczenie ma tagi (np. `mod-div`, `rekurencja`,
  `JOIN`). Kazdy tag ma stabilnosc (sile zapamiętania) i poziom:

| Poziom | Nazwa | Przyblizony interwal |
|--------|-------|----------------------|
| 0 | NOWE | natychmiast |
| 1 | UCZE SIE | ~1-2 dni |
| 2 | CWICZE | ~3-6 dni |
| 3 | PEWNIE | ~7-20 dni |
| 4 | OPANOWANE | 21+ dni |

  Dokladne interwaly sa indywidualne — FSRS uczy sie z Twojej historii.

- **Trudnosc rosnie automatycznie** — 3 poprawne odpowiedzi bez pomocy z rzedu
  = awans na wyzszy poziom. Jesli nie dasz rady (walk-through) — cofniecie o 1.
- **Wykrywanie problemow** — jesli ten sam temat sprawia Ci trudnosc 3+ razy,
  korepetytor zwroci na to uwage i zaproponuje dodatkowe cwiczenia.

---

## Kategorie i typy zadan

23 typy zadan pogrupowane w 4 kategorie — dokladnie jak na maturze CKE.

### TEORIA (6 typow)

| Nr | Typ | Opis |
|----|-----|------|
| 01 | Sledzenie algorytmu | Wykonujesz algorytm krok po kroku, podajesz wynik |
| 02 | Projektowanie algorytmu | Piszesz algorytm/pseudokod rozwiazujacy problem |
| 03 | Analiza algorytmu | Okreslasz zlozonosc, wlasciwosci, poprawnosc |
| 04 | Test prawda/falsz | Oceniasz zdania o algorytmach jako P lub F |
| 05 | Konwersja systemow liczbowych | Zamieniasz miedzy bin/dec/hex i innymi |
| 06 | Teoria bezpieczenstwa | Pytania o szyfrowanie, bezpieczenstwo danych |

### IMPLEMENTACJA (8 typow)

| Nr | Typ | Opis |
|----|-----|------|
| 07 | Cyfry i liczby | Rozkład na cyfry, dzielniki, mod/div |
| 08 | Napisy | Przetwarzanie tekstow, palindromy, szukanie wzorcow |
| 09 | Zlozone | Algorytmy laczace kilka technik |
| 10 | Zliczanie | Ile elementow spelnia warunek |
| 11 | Min/max | Szukanie ekstremow w danych |
| 12 | Sekwencje | Najdluzsze ciagi, serie, podciagi |
| 13 | Obrazy 2D | Operacje na tablicach 2D (mapy, bitmapy) |
| 14 | Obliczenia geometryczne | Odleglosci, pola, wspolrzedne |

### ARKUSZ (5 typow)

| Nr | Typ | Opis |
|----|-----|------|
| 15 | Agregacja warunkowa | SUMIFS, COUNTIFS z warunkami |
| 16 | Symulacja | Modelowanie procesu krok po kroku w arkuszu |
| 17 | Wykres | Tworzenie wykresu z opisami osi, legenda, tytulem |
| 18 | Agregacja podstawowa | SUM, AVG, COUNT, MIN, MAX |
| 19 | Transformacja | Przeksztalcanie danych (VLOOKUP, LEFT, MID) |

### SQL (4 typy)

| Nr | Typ | Opis |
|----|-----|------|
| 20 | GROUP BY | Grupowanie z funkcjami agregujacymi |
| 21 | Podzapytania | Zapytania zagniezdzane (IN, EXISTS, subquery) |
| 22 | JOIN | Laczenie tabel (INNER, LEFT, samozlaczenie) |
| 23 | SELECT/WHERE | Filtrowanie, sortowanie, proste zapytania |

---

## FAQ

**Jak zresetowac postep?**
Usun plik `analiza/cli/matura_progress.db`. Przy nastepnym uruchomieniu `/matura`
korepetytor zacznie od zera.

**Co jesli cwiczenia sie skoncza?**
Korepetytor ostrzeze Cie, gdy w danym typie zostana ostatnie 2 cwiczenia.
Mozesz dogenerowac nowe komenda `/generate-exercises`.

**Jak wrocic do sesji po przerwie?**
Wpisz `/matura` — korepetytor wczyta Twoj postep i zaproponuje kontynuacje
lub zaleglosci powtorkowe.

**Czy moge cwiczic wiecej niz jeden typ na raz?**
Tak — wpisz `zmien temat` w trakcie sesji. Korepetytor sam tez co 3 cwiczenia
wplata powtorke z innego typu (interleaving).

**Skad bierze sie wynik na probnej maturze?**
Ocena odbywa sie wg oficjalnych zasad CKE (te same kryteria, co w kluczu
odpowiedzi). Mozesz dostac punkty czesciowe — tak jak na prawdziwym egzaminie.
