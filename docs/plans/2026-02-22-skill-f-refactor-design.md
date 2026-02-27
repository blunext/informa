# Design: Refaktor sekcji F SKILL.md + poprawki probna.md

Data: 2026-02-22
Kontekst: test-tutor raport 22c3ebe (84.1/100) zidentyfikowal 5 problemow systemowych

## Problem

Test-tutor agenci pomijaja kluczowe kroki z SKILL.md:
1. **probna**: brak `progress status` na starcie (SKILL.md C wymaga ZAWSZE)
2. **probna**: uzycie `exercise answer` na ID egzaminowych (YYYY.N) — powinno byc wg `zasady_oceniania` z `exam task`
3. **probna**: brak `progress blad` przy bledach ucznia
4. **SKILL.md**: `--weight-reset` pomijany w 2/7 scenariuszy mimo 4 wzmianek
5. **SKILL.md**: komunikat hint_delay=2 pomijany mimo obecnosci w sekcji F

Glowna przyczyna: sekcja F ma narracyjna strukture — kluczowe kroki sa rozproszone w podsekcjach, agenci je pomijaja.

## Podejscie

**B: Refaktor sekcji F na checklist + nowa sekcja w probna.md**

### Zmiany w probna.md (3 dodane reguly)

1. **Start, przed pkt 1**: `0. ZAWSZE sprawdz status: ./matura progress status`
2. **Przebieg, po pkt 4**: zakaz `exercise answer` na ID egzaminowych — oceniaj wg `zasady_oceniania` z `exam task`
3. **Przebieg, po pkt 4**: regula `progress blad` przy bledach ucznia

### Zmiany w SKILL.md sekcja F (refaktor)

**USUWANE podsekcje** (tresc wchodzi do checklisty):
- "System hintow (coaching-driven)" (linie 303-326)
- "Przebieg hintow (3 poziomy)" (linie 329-360)

**NOWA podsekcja** na gorze sekcji F — "CHECKLIST — po odpowiedzi ucznia":

```
### CHECKLIST — po odpowiedzi ucznia

1. Pobierz odpowiedz: `exercise answer --id {id}` (lazy — DOPIERO teraz)
2. Porownaj z odpowiedzia ucznia
3. Jesli POPRAWNA → przejdz do kroku 8
4. Jesli BLEDNA — dla KAZDEGO bledu osobno:
   `progress blad --exercise-id {id} --typ {typ} --kod {kod}`
   (PRZED czymkolwiek innym!)
5. Sprawdz `coaching.hint_delay` vs liczba prob:
   - proba < hint_delay → pytanie sokratejskie (BEZ hintow, BEZ exercise hints)
   - proba >= hint_delay → `exercise hints --id {id}`, podaj wskazowke:
     - L1: wskazowki[0] + pytanie sokratejskie + dodaj --hint 1 do progress blad
     - L2: `cheatsheet get --sekcja` + wskazowki[1] + --hint 2
     - L3: wskazowki[2] + rozpisz krok po kroku + --hint 3
6. Przy PIERWSZYM cwiczeniu z hint_delay >= 2:
   - hint_delay=2: "Od teraz mniej podpowiedzi — rozwijasz samodzielnosc."
   - hint_delay=3: "Bez podpowiedzi — jak na prawdziwym egzaminie."
7. Po 3 probach / "poddaje sie" → walk_through:
   - `exercise answer --id {id}`
   - Wyswietl pelna odpowiedz + typowe_bledy
   - Konsolidacja: "Wyjasniej swoimi slowami..."
   - Wizualizacja (jesli typ z listy wizualizacji)
8. Zapis wyniku:
   - ELAPSED=$(($(date +%s) - START_TS))
   - `progress update --id {id} --wynik {wynik} --czas $ELAPSED`
```

**NIEZMIENIONE podsekcje** (zostaja po checklistie jako referencja):
- Lazy loading (obecne linie 281-287) — przeniesione jako komentarz do kroku 1
- Punktacja czesciowa TEORIA (289-301)
- Wizualizacje proaktywne (363-401)
- Zapis wyniku (403-427) — skrocone, glowna logika w kroku 8
- Kody bledow (428-462)
- Dobor kodu bledu (451-462)
- Proaktywna detekcja wzorcow (465-478)

### Zmiany w SKILL.md sekcja D (--weight-reset)

Bez zmian — juz jest 4x wspomniane + [WYMAGANE]. Problem nie jest w tekscie, a w parsowaniu przez agentow. Checklist w F poprawi ogolna compliance.

## Metryki sukcesu

Po wdrozeniu: test-tutor powinien osiagnac:
- integralnosc_cli >= 4.0 (bylo 3.86)
- coaching >= 4.0 (bylo 3.86)
- first_session >= 80 (bylo 72)
- probna >= 80 (bylo 76)
- Overall >= 86 (bylo 84.1)

## Ryzyko

- Refaktor F moze spowodowac regresje w scenariuszach ktore dzialaja dobrze (hint_progression 97, difficulty_climb 93)
- Mitygacja: checklist jest superset obecnych regul, nie usuwa zadnej — tylko konsoliduje i uporzadkowuje
