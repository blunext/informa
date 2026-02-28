# Probna matura — symulacja egzaminu

## Wyzwalanie

Komenda `probna [argument]`:
- **rok** (np. `probna 2024`): konkretny egzamin
- **`losowa`**: `./matura exam list --random`
- **`nowa`**: `./matura exam list --formula nowa --random`
- **`stara`**: `./matura exam list --formula stara --random`
- **bez argumentu**: `./matura exam list` (lista dostepnych lat ze statusem i sugestia)

## Start

0. **[WYMAGANE]** Sprawdz status: `./matura progress status` (SKILL.md C wymaga ZAWSZE)
1. Pobierz metadane: `./matura exam meta --rok {rok}`
2. Zapisz timestamp startu: `date +%s`
3. Wyswietl:
```
--- PROBNA MATURA {rok} ---
Czas: {czas_minuty} min | {total_punkty} pkt | Zadan: {n} | Formula: {formula}

Zasady:
- Zadania podawane sekwencyjnie
- Brak hintow — jak na prawdziwym egzaminie
- Komendy: odpowiedz / pomin (0 pkt) / przerwij (koniec)

Zaczynamy? (tak / nie)
```

## Przebieg

Dla kazdego zadania sekwencyjnie:
1. Pobierz zadanie: `./matura exam task --rok {rok} --zadanie {n}`
2. Wyswietl kontekst + kazde podzadanie po kolei
3. Brak hintow — jesli uczen poprosi: "To probna matura — na egzaminie nie ma hintow. Podaj odpowiedz, `pomin` lub `przerwij`."
4. Ocen wg `zasady_oceniania`, przyznaj punkty czesciowe, krotki feedback (1 zdanie)
5. **[NIGDY]** nie uzywaj `exercise answer` na ID egzaminowych (YYYY.N.M).
   Oceniaj WYLACZNIE wg `zasady_oceniania` z odpowiedzi `exam task`.
6. **Jesli uczen popelni blad** — zapisz:
   `./matura progress blad --exercise-id {rok}.{zad}.{podzad} --typ {typ} --kod {kod} --hint 0`
   Jesli CLI odrzuci kod (zwroci `suggestions[]`) — wywolaj ponownie z sugestia (patrz SKILL.md F.3).
7. Prowadz bufor wynikow: `Zad 1.1: 2/3 pkt | Zad 1.2: 1/1 pkt | ...`

Komendy w trakcie: `pomin` (0 pkt za podzadanie), `przerwij` (koniec egzaminu → podsumowanie)

**Reguly behawioralne:**
- **Niezaleznosc podzadan**: po ocenie NIE odwoluj sie do poprzednich podzadan — traktuj kazde osobno
- **Porcjowanie**: co 3 zadania (nie podzadania) wyswietl mini-podsumowanie z dotychczasowymi punktami
- **Bufor wynikow**: prowadz jako tekst inline — nie polegaj na pamieci z poczatku konwersacji

## Podsumowanie

1. Oblicz czas: `date +%s` minus start_timestamp
2. Wyswietl:
```
--- WYNIK PROBNEJ MATURY {rok} ---
{zdobyte} / {total_punkty} pkt ({procent}%)
Czas: {elapsed_min} min / {limit_min} min

Per zadanie:
  Zad. 1: {tytul}
    1.1 ({typ}): {zdobyte}/{max} pkt {v|~|x|-}
    ...

Per kategoria:
  TEORIA: {pkt}/{max} | IMPLEMENTACJA: {pkt}/{max}
  ARKUSZ: {pkt}/{max} | SQL: {pkt}/{max}

Pulapki, na ktore wpadl(a/e)s:
  - Zad. 1.1: {pulapka}
  ...

Mocne strony: {kategorie z pelnym wynikiem}
Do poprawy: {kategorie z <50%}
```
Gdzie status: `v` (pelne pkt), `~` (czesciowe), `x` (0 pkt), `-` (pominiete).

3. **Pelne rozwiazania**: zapytaj "Chcesz zobaczyc pelne rozwiazania? (tak / konkretne zadanie / nie)"
   - **tak**: wyswietlaj po 3 zadania, po kazdej porcji pytaj "Dalej?"
   - **numer** (np. "1.2"): tylko to podzadanie
   - **nie**: zakoncz

4. Oblicz czas:
   ```
   ELAPSED_MIN=$(( ($(date +%s) - START_TS) / 60 ))
   ```
5. Zapisz: `./matura exam save --rok {rok} --results '[{"id":"2024.1.1","pkt":2,"max":3},...]' --czas $ELAPSED_MIN`
