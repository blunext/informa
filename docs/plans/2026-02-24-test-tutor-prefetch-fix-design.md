# Design: Naprawa parsowania JSON w test-tutor pre-fetch

**Data**: 2026-02-24
**Status**: Zatwierdzony

## Problem

Test-tutor pre-fetch (sekcja 2 SKILL.md) uzywa wzorca `echo "$VAR" | python3 -c "import sys,json; ..."` do ekstrakcji ID z JSON. W zsh, `echo` interpretuje escape sequences (`\n`, `\t`) wewnatrz zmiennych, co zamienia poprawne JSON escapes na prawdziwe control characters. Python `json.loads()` odrzuca control characters w stringach JSON.

Skutek: pre-fetch zmienne (`EX_TEORIA_ID`, `EX_IMPL_ID`, etc.) sa puste. Sub-agenty nie dostaja hintow/odpowiedzi, ida "na zywiol" i probuja sqlite3 z bledna nazwa tabeli (`exercises` zamiast `cwiczenia`).

## Root cause

```
Go CLI → valid JSON: "tresc": "...\n..."
  ↓
zsh $() capture → variable contains literal \n
  ↓
echo "$VAR" → zsh interprets \n as newline (0x0A)
  ↓
python3 json.loads() → "Invalid control character" → crash
```

## Naprawa

Zamienic `echo "$VAR" | python3` na `printf '%s' "$VAR" | python3` we WSZYSTKICH wystapieniach w sekcji 2 test-tutor SKILL.md.

Dotyczy linii: 55, 56, 84, 94 (kazda z `echo "$EX_..." | python3 -c`).

## Weryfikacja

Po naprawie, pelny pre-fetch chain powinien zwrocic poprawne ID dla wszystkich 4 typow cwiczen:

```bash
rm -rf /tmp/test-prefetch && mkdir -p /tmp/test-prefetch && cp matura.db /tmp/test-prefetch/
EX=$(./matura --db-dir /tmp/test-prefetch exercise question --typ sledzenie_algorytmu --trudnosc latwe)
printf '%s' "$EX" | python3 -c "import sys,json; print('OK:', json.loads(sys.stdin.read())['id'])"
# Oczekiwany wynik: OK: 1.X
```
