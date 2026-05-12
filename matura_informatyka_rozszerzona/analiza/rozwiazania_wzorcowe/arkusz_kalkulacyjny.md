# Rozwiazania wzorcowe: Arkusz kalkulacyjny

Dwa pelne rozwiazania prawdziwych zadan maturalnych — z procesem myslowym, formulami i weryfikacja.

---

## [2023] Zadanie 6: Konfitury owocowe (10 pkt)

**Typ**: arkusz_wykres + arkusz_agregacja_warunkowa + arkusz_symulacja | **Czas**: ~50 min | **Trudnosc**: trudne

### Tresc (skrot)

Plik `owoce.txt` (TSV): dostawy owocow do przetworni w okresie 01.05.2020–30.09.2020.
Kolumny: **data**, **maliny** (kg), **truskawki** (kg), **porzeczki** (kg). Lacznie 153 wiersze (dni).

- **6.1** (3 pkt): Zestawienie kg owocow na miesiac (maj-wrzesien) + wykres kolumnowy z opisami.
- **6.2** (1 pkt): Ile dni porzeczki byly na 1. miejscu (najwiecej kg sposrod 3 owocow)?
- **6.3** (3 pkt): Symulacja produkcji konfitur — codziennie produkuje sie konfitur z **2 owocow, ktorych jest najwiecej**. Ile razy produkowano kazdy rodzaj?
- **6.4** (3 pkt): Ile kg konfitur kazdego rodzaju wyprodukowano? (1 kg konfitur = 1 kg owocu A + 1 kg owocu B)

### Podejscie — jak myslec

1. **6.1**: SUMIF po miesiacu (MIESIAC(data)) dla kazdego owocu. Wykres kolumnowy grupowany.
2. **6.2**: COUNTIF z warunkiem `porzeczki = MAX(maliny, truskawki, porzeczki)` — ale uwaga na remisy.
3. **6.3/6.4**: To najtrudniejsze — **symulacja dzen po dniu**:
   - Codziennie: zapas_dzis = zapas_wczoraj + dostawa_dzis
   - Wybierz 2 owoce z 3, ktorych jest najwiecej
   - Produkcja = MIN(owoc_A, owoc_B) — tyle kg konfitur
   - Zuzyte po 1 kg kazdego z 2 owocow na 1 kg konfitur
   - Reszta przechodzi na kolejny dzien

### Rozwiazanie

#### 6.1 — Zestawienie miesiecy + wykres (3 pkt)

**Uklad arkusza** (dane w kolumnach A-D, od wiersza 2):
```
A: data | B: maliny | C: truskawki | D: porzeczki
```

**Zestawienie** (np. w kolumnach F-I):
```
       F          G              H          I
1  Miesiac    Maliny      Truskawki   Porzeczki
2  maj        =SUMA.ILOCZYNÓW((MIESIĄC($A$2:$A$154)=5)*$B$2:$B$154)
3  czerwiec   =SUMA.ILOCZYNÓW((MIESIĄC($A$2:$A$154)=6)*$B$2:$B$154)
...
```

Alternatywnie z SUMIFS:
```
G2: =SUMA.WARUNKÓW($B$2:$B$154,$A$2:$A$154,">="&DATA(2020,5,1),$A$2:$A$154,"<"&DATA(2020,6,1))
```

**Wyniki**:

| Miesiac | Maliny | Truskawki | Porzeczki |
|---|---|---|---|
| Maj | 9238 | 9287 | 3309 |
| Czerwiec | 9485 | 8916 | 5081 |
| Lipiec | 11592 | 11339 | 10567 |
| Sierpien | 11045 | 11386 | 11078 |
| Wrzesien | 6532 | 7476 | 6355 |

**Wykres**: Kolumnowy grupowany, 3 serie (maliny, truskawki, porzeczki).
- Os X: nazwy miesiecy
- Os Y: "Liczba kilogramow"
- Tytul: np. "Dostawy owocow w okresie maj-wrzesien 2020"
- Legenda: maliny, truskawki, porzeczki

**Punktacja wykres**: 1 pkt za typ i dane, 1 pkt za opisy (tytul, legenda, osie). Brak opisow = -1 pkt.

#### 6.2 — Dni z porzeczkami na 1. miejscu (1 pkt)

```
=LICZ.JEŻELI(E2:E154;"porzeczki")
```

gdzie E2 to kolumna pomocnicza:
```
E2: =JEŻELI(ORAZ(D2>=B2, D2>=C2), "porzeczki", "inne")
```

Lub jednolinijkowo:
```
=SUMA.ILOCZYNÓW((D2:D154>=B2:B154)*(D2:D154>=C2:C154)*1)
```

**Wynik**: **19**

#### 6.3 — Symulacja: ktory rodzaj konfitur? (3 pkt)

**Uklad symulacji** — nowe kolumny (np. F-N):

```
F: zapas_maliny    (= zapas_wczoraj + dostawa_dzis)
G: zapas_truskawki
H: zapas_porzeczki
I: owoc_1 (nazwa owocu z najwieksza iloscia)
J: owoc_2 (nazwa owocu z druga najwieksza iloscia)
K: rodzaj_konfitur (np. "malinowo-truskawkowe")
L: produkcja_kg    (= MIN z dwoch wybranych owocow)
M: nowe_F (zapas maliny po produkcji)
N: nowe_G, O: nowe_H
```

Ale prostsze podejscie — **sortujemy 3 wartosci** i bierzemy 2 najwieksze:

```
Dzien 1 (01.05):
  F2: =B2               (zapas maliny = dostawa, bo dzien 1)
  G2: =C2
  H2: =D2

Dzien 2+ (02.05):
  F3: =F2 - [zuzyto_maliny_wczoraj] + B3
```

**Kluczowa logika** — dla kazdego dnia:

Niech m, t, p = zapasy maliny, truskawki, porzeczki.

```
min_z_3 = MIN(m, t, p)   // owoc z najmniejsza iloscia (odpada)
```

Jesli `min_z_3 = p` (porzeczki najmniej) → produkujemy **malinowo-truskawkowe**, ilosc = MIN(m, t).
Jesli `min_z_3 = t` → **malinowo-porzeczkowe**, ilosc = MIN(m, p).
Jesli `min_z_3 = m` → **truskawkowo-porzeczkowe**, ilosc = MIN(t, p).

**Formuly** (uproszczone, wiersz i = dzien i):

```
// Produkcja (kg konfitur)
L_i = MEDIAN(F_i, G_i, H_i)
  -- MEDIAN z 3 wartosci = druga co do wielkosci = MIN z 2 najwiekszych

// Rodzaj
K_i = JEŻELI(MIN(F_i,G_i,H_i)=H_i, "mal-trus",
      JEŻELI(MIN(F_i,G_i,H_i)=G_i, "mal-porz", "trus-porz"))

// Aktualizacja zapasow na kolejny dzien
F_(i+1) = F_i - [zuzyto_maliny] + B_(i+1)
  -- zuzyto_maliny = L_i jesli maliny uzyto, 0 jesli nie
```

Alternatywna sprytna formula: **MEDIAN(m, t, p)** = MIN(2 najwiekszych) = produkcja kg!

**Wyniki**:

| Rodzaj konfitur | Liczba dni |
|---|---|
| malinowo-porzeczkowe | **41** |
| malinowo-truskawkowe | **72** |
| truskawkowo-porzeczkowe | **40** |

(Suma: 153 dni — zgadza sie z okresem maj-wrzesien)

#### 6.4 — Laczna masa konfitur (3 pkt)

1 kg konfitur = 1 kg owocu A + 1 kg owocu B. Wiec ilosc konfitur = MIN(owoc_A, owoc_B) = MEDIAN(m,t,p).

```
Suma konfitur danego rodzaju = SUMA(produkcja_kg) dla dni gdy produkowano ten rodzaj
```

Formuly:
```
=SUMA.ILOCZYNÓW((K2:K154="mal-porz")*L2:L154)    // malinowo-porzeczkowe
=SUMA.ILOCZYNÓW((K2:K154="mal-trus")*L2:L154)    // malinowo-truskawkowe
=SUMA.ILOCZYNÓW((K2:K154="trus-porz")*L2:L154)   // truskawkowo-porzeczkowe
```

**Wyniki**:

| Rodzaj konfitur | Masa (kg) |
|---|---|
| malinowo-porzeczkowe | **18008** |
| malinowo-truskawkowe | **29732** |
| truskawkowo-porzeczkowe | **18382** |
| **Lacznie** | **66122** |

### Weryfikacja

| Podzadanie | Oficjalna odpowiedz CKE |
|---|---|
| 6.1 | Tabela miesiecy + wykres kolumnowy z opisami |
| 6.2 | **19** |
| 6.3 | mal-porz: **41**, mal-trus: **72**, trus-porz: **40** |
| 6.4 | mal-porz: **18008**, mal-trus: **29732**, trus-porz: **18382** |

### Pulapki

- **6.1**: Brak opisu wykresu (tytul, legenda, osie) = -1 pkt. Zawsze dodawaj opisy!
- **6.3/6.4**: Symulacja wymaga **zapasow z poprzedniego dnia** — kazdy blad kumuluje sie.
- **6.4**: 1 kg konfitur = 1 kg owocu A + 1 kg owocu B. Nie 0.5 + 0.5! Wiec produkcja = MIN(A, B), nie MIN(A,B)/2.
- **6.3**: Trzeba poprawnie wybierac pare owocow — owoce z **najwieksza** iloscia, nie np. dwa pierwsze.
- **6.4**: CKE daje 2 pkt za wyniki 2x wieksze (36016, 59464, 36764) — blad interpretacji "2 kg owocow na 1 kg konfitur".

**Wskazowka**: MEDIAN(a, b, c) = MIN z dwoch najwiekszych wartosci — to klucz do prostej formuly produkcji.

---

## [2025] Zadanie 6: Martianeum (11 pkt)

**Typ**: arkusz_agregacja_warunkowa + arkusz_transformacja + arkusz_wykres + arkusz_symulacja | **Czas**: ~50 min | **Trudnosc**: trudne

### Tresc (skrot)

Dane o pracy stacji wydobywczej na Marsie (plik TSV):
Kolumny: **data**, **nazwa_obszaru**, **masa_ladunku** (kg), **zawartosc_martianeum** (%).

Stacja wydobywa martianeum z ladunkow o zawartosci **>= 1%**. Gdy stan magazynu >= 100 kg, transporter zabiera 100 kg na orbite.

- **6.1** (2 pkt): Laczna masa ladunkow + laczna masa wydobytego martianeum.
- **6.2** (1 pkt): Obszar z najmniejsza srednia masa ladunkow.
- **6.3** (2 pkt): Podzial na 7-dniowe okresy (od 03.03.2033). Okres z najwieksza laczna masa ladunkow.
- **6.4** (3 pkt): Zestawienie 30 obszarow x 6 lat + wykres skumulowany kolumnowy.
- **6.5** (3 pkt): Symulacja magazynu — ile razy wyslano 100 kg na orbite? Daty pierwszego i ostatniego transportu.

### Podejscie — jak myslec

1. **6.1**: SUM na masie + SUMPRODUCT z warunkiem >= 1%.
2. **6.2**: AVERAGEIF po obszarze, szukaj MIN.
3. **6.3**: Podzial na okresy = `ZAOKR.DO.CAŁK((data - data_poczatkowa) / 7)`. SUMIFS po numerze okresu.
4. **6.4**: Tabela przestawna (COUNTIFS po obszarze i roku). Wykres skumulowany (stacked bar).
5. **6.5**: Symulacja krokowa: stan += wydobyte, jesli stan >= 100: stan -= 100, transport++.

### Rozwiazanie

#### 6.1 — Laczna masa i martianeum (2 pkt)

Zalozmy dane w A:D od wiersza 2, A=data, B=obszar, C=masa, D=zawartosc%.

```
Laczna masa ladunkow:
=SUMA(C2:C9999)

Laczna masa martianeum (tylko ladunki z zawartoscia >= 1%):
=SUMA.ILOCZYNÓW((D2:D9999>=1)*(C2:C9999)*(D2:D9999/100))
```

Wyjasnienie: dla kazdego ladunku z zawartoscia >= 1%, masa martianeum = masa * zawartosc/100.

**Wyniki**:
- Laczna masa ladunkow: **41498,2 kg**
- Laczna masa martianeum: **3092,2943 kg**

#### 6.2 — Obszar z najmniejsza srednia (1 pkt)

Przygotuj liste unikalnych obszarow (30 nazw). Dla kazdego:

```
=ŚREDNIA.JEŻELI($B$2:$B$9999; F2; $C$2:$C$9999)
```

gdzie F2 = nazwa obszaru. Potem szukaj MIN i odpowiadajacej nazwy:

```
=INDEKS(F2:F31, PODAJ.POZYCJĘ(MIN(G2:G31), G2:G31, 0))
```

**Wynik**: **Thaumasia**

#### 6.3 — Okresy 7-dniowe (2 pkt)

**Kluczowa formula**: numer okresu = `ZAOKR.DO.CAŁK((data - DATA(2033,3,3)) / 7)`

Kolumna pomocnicza (np. E):
```
E2: =ZAOKR.DO.CAŁK((A2-DATA(2033,3,3))/7)
```

Potem SUMIFS po numerze okresu:
```
=SUMA.WARUNKÓW($C$2:$C$9999; $E$2:$E$9999; numer_okresu)
```

Generujemy numery okresow 0, 1, 2, ... i szukamy tego z MAX suma. Data poczatku = DATA(2033,3,3) + numer_okresu * 7.

**Wynik**: **174,5 kg**, poczatek okresu: **13.12.2035**

**Uwaga**: Okresy liczone od 03.03.2033 (nie od poczatku miesiaca!).

#### 6.4 — Zestawienie + wykres skumulowany (3 pkt)

**Tabela**: 30 wierszy (obszary) x 6 kolumn (lata 2033-2038).

```
=SUMA.ILOCZYNÓW(($B$2:$B$9999=$F2)*(ROK($A$2:$A$9999)=G$1))
```

gdzie F2 = nazwa obszaru, G1 = rok (2033).

⚠️ **Uwaga**: w COUNTIFS argument `criteria_range` musi byc zakresem komorek — `ROK(zakres)` jako tablicowa funkcja NIE jest akceptowane (zwroci `#VALUE!`). Dlatego uzywamy SUMPRODUCT, ktore poprawnie obsluguje arytmetyke tablicowa. Alternatywa: kolumna pomocnicza `=ROK(A2)` + COUNTIFS po niej.

**Wykres**: Kolumnowy **skumulowany** (stacked bar chart).
- Os X: nazwy obszarow (30)
- Os Y: liczba przewozow
- Legenda: lata 2033-2038
- Tytul: np. "Liczba przewozow ladunku wg obszaru i roku"

**Wazne**: Wykres musi byc **skumulowany** (stacked), nie grupowany (clustered)! To czesty blad.

#### 6.5 — Symulacja magazynu (3 pkt)

Dane posortowane chronologicznie. Symulacja dzien po dniu:

**Kolumna pomocnicza** — wydobyte martianeum z ladunku:
```
F2: =JEŻELI(D2>=1; C2*D2/100; 0)
```

**Symulacja** (kolumny G = stan magazynu, H = transport):
```
G2: =F2                          // poczatkowy stan = pierwszy ladunek
H2: =JEŻELI(G2>=100; 1; 0)          // czy transport?

G3: =G2 - H2*100 + F3           // stan = poprzedni - wyslane + nowe
H3: =JEŻELI(G3>=100; 1; 0)
...
```

**Uwaga**: Transporter zabiera **dokladnie 100 kg** — nawet jesli magazyn ma 250 kg, zabiera tylko raz 100 kg (zostaje 150 kg). Jesli po zabraniu 100 wciaz jest >= 100, nastepny transport dopiero przy **nastepnym ladunku** (nie w tym samym dniu wielokrotnie).

Hmm, ale to zalezy od interpretacji. Jesli dane sa posortowane i kazdy wiersz to oddzielny ladunek (moze byc wiele w jednym dniu), to sprawdzamy warunek po kazdym ladunku osobno.

**Wyniki**:
- Liczba transportow na orbite: **30**
- Pierwszy transport: **29.05.2033**
- Ostatni transport: **01.09.2038**

### Weryfikacja

| Podzadanie | Oficjalna odpowiedz CKE |
|---|---|
| 6.1 | Masa ladunkow: **41498,2 kg**, martianeum: **3092,2943 kg** |
| 6.2 | **Thaumasia** |
| 6.3 | **174,5 kg**, poczatek: **13.12.2035** |
| 6.4 | Tabela 30x6 + wykres skumulowany kolumnowy |
| 6.5 | **30** transportow, pierwszy: **29.05.2033**, ostatni: **01.09.2038** |

### Pulapki

- **6.1**: Zawartosc >= 1% (nie > 1%) — warunek progowy wlaczajacy.
- **6.1**: Wydobycie = masa * zawartosc/100 — **cala** zawartosc mineralu, nie tylko nadwyzka powyzej 1%.
- **6.3**: Okresy 7-dniowe od **03.03.2033**, nie od poczatku miesiaca. Formula: `ZAOKR.DO.CAŁK((data - start) / 7)`.
- **6.4**: Wykres **skumulowany** (stacked), nie grupowany (clustered) — czesty blad.
- **6.5**: Transporter zabiera **dokladnie 100 kg** — nadmiar zostaje w magazynie.
- **6.5**: Nie kumuluj transportow — sprawdzaj warunek po kazdym ladunku osobno.
- Separator dziesietny — dane uzywaja **przecinka** (format polski), nie kropki.
