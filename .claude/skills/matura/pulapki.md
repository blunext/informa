# Pulapki CKE — tryb rozpoznawania pulapek

## Wyzwalanie

Komenda `pulapki [typ|kategoria]`. Bez argumentu → pulapki z typu, nad ktorym uczen aktualnie pracuje.

## Pobieranie

`./matura trap list --typ {typ}` lub `--kategoria {kat}`

## Tryb quizowy

1. Wyswietl skrocona tresc zadania CKE (max 5-6 linii)
2. Zapytaj: "Jakie pulapki widzisz w tym zadaniu? Co moze pojsc nie tak?"
3. Porownaj odpowiedz ucznia z `pulapki[]`
4. Wyswietl feedback:
```
--- Pulapki CKE (Matura {rok}, Zad. {numer}) ---
Twoje trafienia: {N}/{total}
  v {pulapka_1} — trafione!
  x {pulapka_2} — przeoczone
```
5. Zapisz: `./matura trap save --id {id} --typ {typ} --trafienia N --total M`
6. Zapytaj: "Nastepne zadanie czy konczymy?"

## Tryb przegladowy

Komenda `pulapki lista [typ|kategoria]` — wyswietl zestawienie pulapek pogrupowane tematycznie.
