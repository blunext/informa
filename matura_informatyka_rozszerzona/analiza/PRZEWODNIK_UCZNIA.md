# Przewodnik ucznia — matura rozszerzona z informatyki

> Ten dokument to Twoj punkt startu. Masz przed soba kompletny zestaw
> materialow do nauki, oparty na analizie 12 lat prawdziwych egzaminow CKE
> (2014-2025). Znajdziesz tu arkusze egzaminacyjne, sciagawki, szablony kodu,
> 937 cwiczen i rozwiazania wzorcowe. Wszystko poukładane tak, zebys nie tracil
> czasu na szukanie — tylko na nauke.
>
> Przeczytaj ten przewodnik od poczatku do konca (10-15 minut). Potem zacznij
> prace wedlug planu nauki z rozdzialu 5.

---

## 1. Jak wyglada egzamin

Od 2023 roku obowiazuje nowa formula egzaminu:

- **Czas**: 210 minut (3.5 godziny)
- **Punkty**: 50
- **Zadan**: 7-8
- **Jeden arkusz** (bez podzialu na czesci)

Zadania naleza do 4 kategorii:

| Kategoria | Punkty (typowo) | Narzedzie |
|-----------|-----------------|-----------|
| Teoria (algorytmy, sledzenie, P/F) | ~10 pkt | kartka + dlugopis |
| Implementacja (C++) | ~20 pkt | IDE (Code::Blocks / Dev-C++) |
| Arkusz kalkulacyjny | ~10 pkt | LibreOffice Calc / Excel |
| SQL (bazy danych) | ~10 pkt | Access / Base |

Implementacja C++ daje najwiecej punktow. Ale teoria i SQL sa szybsze do
zrobienia — wiec zaczynaj od nich, a C++ zostaw na koniec (wiecej w rozdziale 7).

---

## 2. Co musisz umiec — absolutne minimum

Z analizy 12 lat wynika jasna hierarchia. Te tematy pojawiaja sie prawie zawsze:

**MUSISZ umiec (100% egzaminow):**

1. **SQL** — JOIN, GROUP BY, podzapytania, WHERE z warunkami. Bylo na KAZDYM
   egzaminie bez wyjatku.
2. **Operacje na cyfrach i liczbach** — dzielenie modulo, dzielenie calkowite,
   rozkład na cyfry, suma cyfr. Podstawa wielu zadan programistycznych.
3. **Przetwarzanie plikow w C++** — wczytaj dane z pliku TXT, przetworz,
   wypisz wynik. Kazde zadanie implementacyjne tego wymaga.
4. **Arkusz kalkulacyjny** — SUMIFS, COUNTIFS, VLOOKUP, wykresy z opisami.
   Zawsze jest przynajmniej jedno zadanie arkuszowe.

**Bardzo czeste (75-83% egzaminow):**

5. **Systemy liczbowe** (83%) — konwersje miedzy bin/dec/hex. Czesto w czesci
   teoretycznej, ale tez w zadaniach programistycznych.
6. **Rekurencja** (75%) — sledzenie wywolan, obliczanie wartosci. Musisz umiec
   "rozpisac" rekurencje na kartce.
7. **Sortowanie** (75%) — znajomosc algorytmow, zlozonosc, zastosowanie.
8. **NWD / teoria liczb** (75%) — algorytm Euklidesa, dzielniki, liczby pierwsze.

---

## 3. Mapa materialow — co gdzie znalezc

### Arkusze egzaminacyjne (12 lat)

| Sciezka | Co to jest |
|---------|------------|
| `YYYY_maj/arkusz.pdf` | Tresc egzaminu z danego roku |
| `YYYY_maj/odpowiedzi.pdf` | Klucz odpowiedzi i zasady oceniania |
| `YYYY_maj/Dane*/` | Pliki z danymi do zadan programistycznych |

Lata: 2014-2025 (wszystkie lata, bez przerwy).
2020: egzamin w czerwcu (COVID-19).

Do cwiczen na pelnych arkuszach korzystaj z lat **2023-2025** (nowa formula).
Starsze lata (2014-2022) sa przydatne do cwiczen tematycznych, ale maja inny
format (dwie czesci zamiast jednej).

### Strategia i analiza

| Plik | Co to jest | Kiedy uzywac |
|------|------------|--------------|
| `analiza/strategia_egzaminacyjna.md` | TOP 14 algorytmow + kod C++ + wzorce SQL + strategia | Przeczytaj jako pierwszy dokument |
| `analiza/drzewo_decyzyjne.md` | Schemat: opis problemu --> algorytm | Gdy nie wiesz jak zaczac zadanie |
| `analiza/PODSUMOWANIE_FINAL.md` | Podsumowanie analizy rok po roku | Dla kontekstu, nie do nauki |

### Sciagawki (7 plikow)

Katalog `analiza/cheatsheets/` — wersje skrocone, idealne do wydruku na A4.

| Plik | Zawartosc |
|------|-----------|
| `cheatsheet_cpp.md` | Skladnia C++, wczytywanie plikow, wzorce |
| `cheatsheet_sql.md` | SELECT, JOIN, GROUP BY, podzapytania |
| `cheatsheet_arkusz.md` | Formuly Excel/Calc, wykresy |
| `cheatsheet_teoria.md` | Algorytmy, zlozonosc, systemy liczbowe |
| `przed_egzaminem.md` | Checklist: co zabrac, co powtorzyc |
| `podczas_egzaminu.md` | Strategia rozwiazywania, kolejnosc zadan |
| `debug_checklist.md` | Typowe bledy i jak je naprawic |

### Szablony kodu (6 plikow)

Katalog `analiza/szablony/` — gotowy kod do skopiowania i modyfikacji.

| Plik | Zawartosc |
|------|-----------|
| `cpp_szablony.md` | Szablony C++: wczytywanie, tablice, algorytmy |
| `sql_szablony.md` | Szablony SQL: JOIN, GROUP BY, podzapytania |
| `arkusz_formuly.md` | Szablony formul arkuszowych |
| `pseudokod_wzorce.md` | Wzorce pseudokodu do zadan teoretycznych |
| `wzorce_2014.md` | Wzorce specyficzne dla matury 2014 |
| `wzorce_2015.md` | Wzorce specyficzne dla matury 2015 |

### Cwiczenia (937 cwiczen w 23 typach)

Katalog `analiza/cwiczenia/json/` — 40+ cwiczen na kazdy typ zadania
maturalnego, w formacie JSON z 3-poziomowymi hintami i tagami umiejetnosci.

Najwygodniej korzystac przez korepetytora `/matura` — automatycznie dobiera
trudnosc, prowadzi powtorki rozlozone w czasie i sledzi postep.

Stare pliki MD w `analiza/cwiczenia/wg_typu/` wciaz istnieja, ale system JSON
je zastapil. Pelna lista z priorytetami w `analiza/cwiczenia/wg_typu/README.md`.

Najwazniejsze pliki (posortowane wg punktow na maturach):

| Nr | Plik | Typ | Punkty laczne |
|----|------|-----|---------------|
| 01 | `01_sledzenie_algorytmu.md` | Sledzenie algorytmu | 45 pkt |
| 02 | `02_projektowanie_algorytmu.md` | Projektowanie algorytmu | 43 pkt |
| 15 | `15_agregacja_warunkowa.md` | Arkusz — agregacja warunkowa | 38 pkt |
| 03 | `03_analiza_algorytmu.md` | Analiza algorytmu | 37 pkt |
| 16 | `16_symulacja.md` | Arkusz — symulacja | 37 pkt |
| 07 | `07_cyfry_liczby.md` | Cyfry i liczby (C++) | 36 pkt |
| 20 | `20_sql_group_by.md` | SQL GROUP BY | 36 pkt |
| 04 | `04_test_prawda_falsz.md` | Test prawda/falsz | 25 pkt |
| 17 | `17_wykres.md` | Arkusz — wykres | 25 pkt |
| 22 | `22_sql_join.md` | SQL JOIN | 21 pkt |

### Rozwiazania wzorcowe (4 pliki, 10 zadan)

Katalog `analiza/rozwiazania_wzorcowe/` — pelne rozwiazania prawdziwych zadan
z procesom myslowym: tresc --> podejscie --> kod --> weryfikacja --> pulapki.

| Plik | Zawartosc |
|------|-----------|
| `implementacja_cpp.md` | Zadania programistyczne z kodem C++ |
| `sql_zapytania.md` | Zadania bazodanowe z zapytaniami SQL |
| `teoria_algorytmy.md` | Zadania teoretyczne z rozpisaniem krok po kroku |
| `arkusz_kalkulacyjny.md` | Zadania arkuszowe z formulami i weryfikacja |

---

## 4. Plan nauki — krok po kroku

### Wariant A: Masz 3 tygodnie lub wiecej

**Tydzien 1 — Podstawy teoretyczne i szablony**

- Dzien 1: Przeczytaj `analiza/strategia_egzaminacyjna.md` (rozdzialy 1-6)
  oraz `analiza/drzewo_decyzyjne.md`. Notuj, czego nie rozumiesz.
- Dzien 2-3: Otworz `analiza/cheatsheets/cheatsheet_cpp.md` i
  `analiza/szablony/cpp_szablony.md`. Przepisz recznie kazdy szablon — pisanie
  reka utrwala lepiej niz czytanie. Uruchom kazdy szablon w IDE.
- Dzien 4-5: To samo z SQL: `analiza/cheatsheets/cheatsheet_sql.md` +
  `analiza/szablony/sql_szablony.md`. Napisz kazde zapytanie sam, sprawdz, popraw.
- Dzien 6-7: Arkusz: `analiza/cheatsheets/cheatsheet_arkusz.md` +
  `analiza/szablony/arkusz_formuly.md`. Otworz
  dowolny arkusz w LibreOffice i przetestuj formuly na przykladowych danych.

**Tydzien 2 — Cwiczenia (najczestsze typy)**

Zalecana sciezka: wpisz `/matura` — korepetytor sam dobierze cwiczenia do
Twojego poziomu, da hinty i bedzie sledzil postep. Mozesz tez robic cwiczenia
recznie z katalogu `analiza/cwiczenia/wg_typu/` w tej kolejnosci:

1. `01_sledzenie_algorytmu.md` — sledzenie (45 pkt lacznie na maturach)
2. `02_projektowanie_algorytmu.md` — projektowanie (43 pkt)
3. `03_analiza_algorytmu.md` — analiza algorytmow (37 pkt)
4. `04_test_prawda_falsz.md` — testy P/F (25 pkt)
5. `07_cyfry_liczby.md` — implementacja: cyfry/liczby (36 pkt)
6. `10_zliczanie.md` — implementacja: zliczanie (17 pkt)
7. `11_minmax.md` — implementacja: min/max (17 pkt)
8. `15_agregacja_warunkowa.md` — arkusz: SUMA.WARUNKÓW/LICZ.WARUNKI (38 pkt)
9. `16_symulacja.md` — arkusz: symulacja (37 pkt)
10. `17_wykres.md` — arkusz: wykresy (25 pkt)
11. `20_sql_group_by.md` — SQL GROUP BY (36 pkt)
12. `22_sql_join.md` — SQL JOIN (21 pkt)

Tempo: 5-8 cwiczen dziennie. Zaczynaj od latwych (pierwsze 2 w kazdym pliku).
Po kazdym bledzie: przeczytaj rozwiazanie, zapisz w zeszycie czego nie
wiedziales.

**Tydzien 3 — Rozwiazania wzorcowe i pelne arkusze**

- Dzien 1-3: Przerob 10 zadan z `analiza/rozwiazania_wzorcowe/`. Najpierw
  probuj rozwiazac sam (20-30 min na zadanie), dopiero potem czytaj rozwiazanie.
- Dzien 4-5: Zrob 2-3 pelne arkusze egzaminacyjne z lat 2023-2025 (nowa
  formula). Mierz czas — masz 210 minut. Pliki: `YYYY_maj/arkusz.pdf` + dane
  z katalogu `YYYY_maj/Dane*/`.
- Dzien 6-7: Przejrzyj `analiza/cheatsheets/przed_egzaminem.md` i
  `analiza/cheatsheets/podczas_egzaminu.md`. Powtorz szablony, ktore
  sprawialy Ci problemy.

### Wariant B: Masz mniej niz 7 dni

Nie masz czasu na wszystko. Skup sie na tym, co da najwiecej punktow:

- **Dzien 1**: Przeczytaj `analiza/strategia_egzaminacyjna.md` + 4 sciagawki
  z `analiza/cheatsheets/` (`cheatsheet_cpp.md`, `cheatsheet_sql.md`,
  `cheatsheet_arkusz.md`, `cheatsheet_teoria.md`). Wydrukuj je.
- **Dzien 2-3**: Cwiczenia priorytetowe — zrob przynajmniej latwe i srednie z:
  `01_sledzenie_algorytmu.md`, `02_projektowanie_algorytmu.md`,
  `07_cyfry_liczby.md`, `15_agregacja_warunkowa.md`, `20_sql_group_by.md`.
- **Dzien 4-5**: Rozwiazania wzorcowe — przy tak malej ilosci czasu przeczytaj
  je jak tutorial, nie rob sam. Skup sie na sekcjach "Podejscie" i "Pulapki".
- **Dzien 6**: Pelny arkusz z 2025 roku z pomiarem czasu. Sprawdz z kluczem
  (`2025_maj/odpowiedzi.pdf`).
- **Dzien 7**: Checklisty (`analiza/cheatsheets/przed_egzaminem.md`,
  `analiza/cheatsheets/podczas_egzaminu.md`) + powtorka szablonow
  z `analiza/szablony/`.

---

## 5. Jak korzystac z poszczegolnych materialow

### Cwiczenia (937 cwiczen w 23 typach)

Baza cwiczen zawiera 937 zadan (40+ na typ) z 3-poziomowymi hintami, tagami
umiejetnosci i wzorcowymi odpowiedziami. Cwiczenia sa w formacie JSON
w katalogu `analiza/cwiczenia/json/`.

**Zalecana sciezka**: wpisz `/matura` — korepetytor automatycznie dobiera
trudnosc, daje hinty metoda sokratejska i prowadzi powtorki rozlozone w czasie.
Nie musisz recznie wybierac cwiczen ani szukac plikow.

Zasady pracy (niezaleznie od metody):
- Zawsze najpierw rozwiaz sam — na kartce albo w IDE. Dopiero potem sprawdz.
- Latwe cwiczenia powinny zajac 3-5 minut. Jesli zajmuja dluzej, wroc do
  odpowiedniego cheatsheet'a i powtorz teorie.
- Cel: 80% poprawnych odpowiedzi na poziomie "srednie". Jak to osiagniesz,
  jestes gotowy na ten typ zadania na egzaminie.
- Jesli pomylisz sie w cwiczeniu, zapisz w zeszycie CO bylo bledem. Nie
  wystarczy przeczytac odpowiedz — musisz zrozumiec, dlaczego Twoje
  rozwiazanie bylo bledne.

### Rozwiazania wzorcowe (`analiza/rozwiazania_wzorcowe/`)

To NIE sa cwiczenia — to wzorce myslenia. Czytaj je jak tutorial, nie jak test.

Struktura kazdego rozwiazania:
1. **Tresc** — oryginalne zadanie
2. **Podejscie** — jak myslec o problemie, jaki algorytm wybrac
3. **Kod/rozwiazanie** — pelne rozwiazanie
4. **Weryfikacja** — sprawdzenie na przykladowych danych
5. **Pulapki** — typowe bledy i jak ich uniknac

Przeczytaj sekcje "Podejscie" ZANIM spojrzysz na rozwiazanie. Sprobuj sam
wymyslic algorytm — nawet jesli nie napiszesz kodu, sam proces myslenia jest
cenny.

### Szablony (`analiza/szablony/`)

To gotowy kod do skopiowania i modyfikacji. NIE ucz sie ich na pamiec.
Zamiast tego:
- Naucz sie SZUKAC wlasciwego szablonu dla danego problemu.
- Otworz `analiza/szablony/cpp_szablony.md` obok IDE podczas cwiczen — kopiuj, modyfikuj,
  uruchamiaj.
- Z czasem zaczniesz pamietac najczestsze wzorce naturalnie.

### Sciagawki (`analiza/cheatsheets/`)

Wersje skrocone szablonow — zmieszcza sie na A4. Wydrukuj te 4:
- `analiza/cheatsheets/cheatsheet_cpp.md`
- `analiza/cheatsheets/cheatsheet_sql.md`
- `analiza/cheatsheets/cheatsheet_arkusz.md`
- `analiza/cheatsheets/cheatsheet_teoria.md`

Miej je pod reka podczas cwiczen. Na egzaminie nie mozesz ich miec, ale do
tego czasu bedziesz juz pamietac wiekszość.

### Drzewo decyzyjne (`analiza/drzewo_decyzyjne.md`)

Otwieraj za kazdym razem, gdy nie wiesz jak zaczac zadanie. Prowadzi od opisu
problemu do konkretnego algorytmu: "Czy sa dane liczbowe? Tak --> Czy szukasz
ekstremu? Tak --> minmax". Z czasem te sciezki beda automatyczne.

---

## 6. Strategia na egzaminie

### Kolejnosc rozwiazywania

1. **Przeczytaj WSZYSTKIE zadania** (5 minut). Zaznacz przy kazdym: "umiem" /
   "moze" / "nie wiem". To pozwoli Ci zaplanowac czas.

2. **Quick wins** — zrob NAJPIERW:
   - Konwersje systemow liczbowych (1-2 pkt, 2-3 minuty)
   - Testy prawda/falsz (2-3 pkt, 5 minut)
   - Proste zapytania SQL (SELECT ... WHERE)
   - Sledzenie algorytmu (jesli dobrze rozumiesz pseudokod)

3. **Zadania standardowe** — potem:
   - SQL z GROUP BY i JOIN
   - Arkusz kalkulacyjny (formuly + wykres)
   - Sledzenie/analiza algorytmow

4. **Zadania trudne** — na koniec:
   - Implementacja C++ (zlożone algorytmy)
   - Projektowanie algorytmu (pseudokod)

### Zasady ogolne

- **NIGDY nie zostawiaj pustego miejsca.** Za czesciowe rozwiazanie dostaniesz
  punkty czesciowe. Nawet zapis "wczytaj dane z pliku" w C++ to jest punkt.
- **Zapisuj co 15 minut** (Ctrl+S). Awaria komputera na egzaminie to nie mit.
- **Czytaj tresc dokladnie.** Roznica miedzy `> 90` a `>= 90` to inna
  odpowiedz i stracone punkty.
- **Sprawdzaj na danych przykladowych.** Kazde zadanie ma plik z przykladem
  (np. `liczby_przyklad.txt`). Uzyj go do weryfikacji ZANIM przejdziesz do
  pelnych danych.

Pelna strategia: `analiza/cheatsheets/podczas_egzaminu.md`

---

## 7. Najczestsze bledy uczniow

Te bledy powtarzaja sie co roku. Kazdy z nich to stracone punkty:

**C++:**
- `int` overflow przy duzych liczbach — uzywaj `long long` gdy dane moga
  przekraczac 2 miliardy.
- Zapomnienie `#include <fstream>` lub bledna sciezka do pliku — program
  kompiluje sie, ale nie wczytuje danych.
- Brak `endl` lub `"\n"` na koncu wyjscia — wynik moze sie nie wyswietlic.

**SQL:**
- Uzycie INNER JOIN gdy szukasz "czego NIE ma w tabeli" — potrzebujesz
  LEFT JOIN + WHERE ... IS NULL.
- Zapomnienie GROUP BY przy uzyciu funkcji agregujacych (COUNT, SUM, AVG).
- Aliasy kolumn nierozpoznawane w WHERE — uzyj HAVING lub podzapytania.

**Arkusz kalkulacyjny:**
- Brak opisu wykresu: tytul, legenda, opisy osi — kazdy brak to minus punkt.
- Zly zakres w SUMIFS/COUNTIFS — upewnij sie, ze zakresy maja ten sam rozmiar.
- Brak dolarow ($) w odwolaniach przy kopiowaniu formul w dol.

**Teoria:**
- `> 90` to nie to samo co `>= 90`. Czytaj warunki dokladnie.
- Przy sledzeniu pseudokodu: uzycie stringow tam, gdzie sa zabronione = 0 pkt
  za caly podpunkt.
- Pominiecie przypadku brzegowego (np. pusta tablica, jeden element).

Pelna lista: `analiza/cheatsheets/debug_checklist.md`

---

## 8. Od czego zaczac — teraz

Jesli nie wiesz od czego zaczac, zrob dokladnie to:

1. Otworz `analiza/strategia_egzaminacyjna.md` i przeczytaj rozdzialy 1-6.
2. Otworz `analiza/drzewo_decyzyjne.md` i przejrzyj schemat.
3. Wydrukuj 4 sciagawki z `analiza/cheatsheets/`.
4. Wpisz `/matura` — korepetytor poprowadzi Cie sam: dobierze cwiczenia,
   da hinty i bedzie sledzil Twoj postep.

Alternatywnie, mozesz zaczac recznie od pliku
`analiza/cwiczenia/wg_typu/01_sledzenie_algorytmu.md`.

Nie probuj ogarnac wszystkiego naraz. Jeden typ zadania dziennie to wystarczy,
zeby w 2-3 tygodnie byc gotowym na egzamin.

---

## 9. Interaktywny korepetytor

Masz do dyspozycji interaktywnego korepetytora AI — wpisz `/matura` zeby zaczac
sesje. Co oferuje:

- **937 cwiczen** z 23 typow zadan (40+ na typ) z 3-poziomowymi hintami
- **641 prawdziwych zadan CKE** z 12 lat egzaminow (sprawdziany typu)
- **12 probnych matur** — pelna symulacja egzaminu z ocena wg zasad CKE
- **Spaced repetition (FSRS-5)** — adaptacyjne powtorki rozlozone w czasie (interwaly dobierane indywidualnie)
- **Auto-trudnosc** — poziom rosnie z kazdym sukcesem, spada po porazce
- **Metoda sokratejska** — naprowadza pytaniami, nie podaje gotowych odpowiedzi
- **Tryb pulapek CKE** — trening rozpoznawania typowych pulapek egzaminacyjnych

Postep zapisuje sie automatycznie — mozesz przerwac i wrocic w dowolnym
momencie.

Pelny opis trybow i komend: patrz `JAK_UZYWAC_KOREPETYTORA.md`.
