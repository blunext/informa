"""Testy dla polonizuj_formuly.py."""
import json
import shutil
import sys
from pathlib import Path

import pytest

SCRIPTS_DIR = Path(__file__).resolve().parent
sys.path.insert(0, str(SCRIPTS_DIR))
import polonizuj_formuly as pf

FIXTURES = SCRIPTS_DIR / "tests" / "fixtures"


def test_warstwa1_polonizuje_caly_plik_md(tmp_path):
    """Warstwa 1: globalna podmiana wszystkich formul w pliku."""
    src = FIXTURES / "fixture_warstwa1.md"
    dst = tmp_path / "fixture.md"
    shutil.copy(src, dst)

    mapa = pf.load_mapa()
    changes = pf.polonizuj_md_warstwa1(dst, mapa)

    content = dst.read_text()
    assert "=SUMA(A1:A10)" in content
    assert "JEŻELI(B1>0" in content
    assert "ILE.LICZB(B:B)" in content
    assert "WYSZUKAJ.PIONOWO(C1" in content
    assert "SUMA.WARUNKÓW(D:D" in content
    assert "ZAOKR(x" in content
    assert "ROK(data)" in content
    assert "SUM(" not in content
    assert "IF(" not in content
    assert "COUNT(" not in content
    assert "VLOOKUP(" not in content
    assert "SUMIFS(" not in content
    assert "ROUND(" not in content
    assert "YEAR(" not in content
    assert changes > 0


def test_warstwa1_nie_rusza_nazw_bez_nawiasow(tmp_path):
    """Funkcja SUMA jest największa — bez nawiasu zostaje."""
    f = tmp_path / "test.md"
    f.write_text("Tekst o funkcji SUM jest dluzszy. Ale =SUM(A:A) tak.")
    mapa = pf.load_mapa()
    pf.polonizuj_md_warstwa1(f, mapa)
    content = f.read_text()
    assert "funkcji SUM jest" in content
    assert "=SUMA(A:A)" in content


def test_warstwa2_stosuje_tylko_liste_deterministyczna(tmp_path):
    """Warstwa 2: zamienia tylko fragmenty z polonizacja_warstwa2.json."""
    src = FIXTURES / "fixture_warstwa2.md"
    dst = tmp_path / "fixture.md"
    shutil.copy(src, dst)

    zamiany = [
        {
            "plik": str(dst),
            "linia": 16,
            "stary": "- [ ] `=SUMIF(A:A; \"x\"; B:B)` — to ma byc zmienione",
            "nowy": "- [ ] `=SUMA.JEŻELI(A:A; \"x\"; B:B)` — to ma byc zmienione"
        },
        {
            "plik": str(dst),
            "linia": 17,
            "stary": "- [ ] `=COUNTIF(C:C; \">0\")` — to ma byc zmienione",
            "nowy": "- [ ] `=LICZ.JEŻELI(C:C; \">0\")` — to ma byc zmienione"
        }
    ]
    changes = pf.polonizuj_md_warstwa2(zamiany)

    content = dst.read_text()
    assert "SELECT COUNT(*), SUM(price), MAX(date)" in content
    assert "if (x > 0) sum += abs(x);" in content
    assert "=SUMA.JEŻELI(A:A;" in content
    assert "=LICZ.JEŻELI(C:C;" in content
    assert "=SUMIF(" not in content
    assert "=COUNTIF(" not in content
    assert changes == 2


def test_warstwa2_przerwa_z_bledem_gdy_stary_brak(tmp_path):
    """Jesli STARY fragment nie istnieje w pliku — exit code 2."""
    f = tmp_path / "test.md"
    f.write_text("Inna tresc bez tego fragmentu")
    zamiany = [{"plik": str(f), "linia": 1, "stary": "NIEISTNIEJACE", "nowy": "X"}]
    with pytest.raises(SystemExit) as exc:
        pf.polonizuj_md_warstwa2(zamiany)
    assert exc.value.code == 2


def test_warstwa3_polonizuje_json_cwiczenia(tmp_path):
    """Warstwa 3: zamienia formuly w tresc/odpowiedz/wskazowki/typowe_bledy + tagi."""
    src = FIXTURES / "fixture_cwiczenie.json"
    dst = tmp_path / "test.json"
    shutil.copy(src, dst)
    mapa = pf.load_mapa()
    changes = pf.polonizuj_json_cwiczenie(dst, mapa)

    data = json.loads(dst.read_text())
    assert "SUMA.JEŻELI(B2:B11" in data["tresc"]
    assert "SUMA.JEŻELI(B:B" in data["odpowiedz"]
    assert "SUMA.JEŻELI ma 3 argumenty" in data["wskazowki"][0]["tekst"]
    assert "Pomylenie SUMA.JEŻELI z LICZ.JEŻELI" in data["typowe_bledy"][0]["opis"]
    assert "SUMA.JEŻELI" in data["tagi"]
    assert "SUMIF" not in data["tagi"]
    assert changes > 0


def test_warstwa3_polonizuje_meta_json(tmp_path):
    """_meta.json: tagi_globalne tez rename'owane."""
    f = tmp_path / "_meta.json"
    f.write_text(json.dumps({
        "tagi_globalne": ["SUMIF", "SUMIFS", "AVERAGEIFS", "warunek-liczbowy"]
    }))
    mapa = pf.load_mapa()
    pf.polonizuj_json_meta(f, mapa)
    data = json.loads(f.read_text())
    assert data["tagi_globalne"] == ["SUMA.JEŻELI", "SUMA.WARUNKÓW", "ŚREDNIA.WARUNKÓW", "warunek-liczbowy"]


def test_warstwa3_polonizuje_cwiczenia_tagi_w_meta(tmp_path):
    """_meta.json zawiera tez cwiczenia[i].tagi — validator sprawdza spojnosc z X.json."""
    f = tmp_path / "_meta.json"
    f.write_text(json.dumps({
        "tagi_globalne": ["SUMIF", "warunek-tekstowy"],
        "cwiczenia": [
            {"id": "15.1", "trudnosc": "latwe", "tagi": ["COUNTIF", "warunek-tekstowy"]},
            {"id": "15.2", "trudnosc": "srednie", "tagi": ["SUMIFS", "AVERAGEIFS"]}
        ]
    }))
    mapa = pf.load_mapa()
    changes = pf.polonizuj_json_meta(f, mapa)
    data = json.loads(f.read_text())
    assert data["cwiczenia"][0]["tagi"] == ["LICZ.JEŻELI", "warunek-tekstowy"]
    assert data["cwiczenia"][1]["tagi"] == ["SUMA.WARUNKÓW", "ŚREDNIA.WARUNKÓW"]
    assert "SUMA.JEŻELI" in data["tagi_globalne"]
    assert changes >= 4  # 1 tagi_globalne + 3 tag-renames in cwiczenia


def test_warstwa4_rename_w_tagi_rejestr(tmp_path):
    """tagi_rejestr.json: rename 7 specyficznych wpisow."""
    f = tmp_path / "tagi_rejestr.json"
    f.write_text(json.dumps({
        "_meta": "central tag registry",
        "tagi": ["AVERAGEIF", "AVERAGEIFS", "COUNTIF", "COUNTIFS", "SUMIF", "SUMIFS", "VLOOKUP",
                 "warunek-tekstowy", "warunek-liczbowy", "JOIN", "GROUP_BY"]
    }))
    mapa = pf.load_mapa()
    changes = pf.polonizuj_tagi_rejestr(f, mapa)
    data = json.loads(f.read_text())
    for stary in ["AVERAGEIF", "AVERAGEIFS", "COUNTIF", "COUNTIFS", "SUMIF", "SUMIFS", "VLOOKUP"]:
        assert stary not in data["tagi"]
    for nowy in ["ŚREDNIA.JEŻELI", "ŚREDNIA.WARUNKÓW", "LICZ.JEŻELI", "LICZ.WARUNKI",
                 "SUMA.JEŻELI", "SUMA.WARUNKÓW", "WYSZUKAJ.PIONOWO"]:
        assert nowy in data["tagi"]
    assert "warunek-tekstowy" in data["tagi"]
    assert "JOIN" in data["tagi"]
    assert "GROUP_BY" in data["tagi"]
    assert changes == 7


def test_separator_normalizuje_przecinek_w_formule(tmp_path):
    """Separator argumentow ',' -> ';' w formulach arkusza."""
    f = tmp_path / "test.md"
    f.write_text("Test: =SUMA(A:A, \"X\", B:B) oraz =ZAOKR(3,14; 2)")
    pf.normalizuj_separator(f)
    content = f.read_text()
    assert "=SUMA(A:A; \"X\"; B:B)" in content
    assert "=ZAOKR(3,14; 2)" in content


def test_separator_obsluguje_warunki_porownan(tmp_path):
    """Comma after comparison digit (e.g. x>0, "tak") should be separator, not decimal."""
    f = tmp_path / "test.md"
    f.write_text('=IF(x>0, "tak", "nie")\n=ZAOKR(3,14, 2)\n=SUMA(A1, B1, 1.5)')
    pf.normalizuj_separator(f)
    content = f.read_text()
    # IF: comparison >0 followed by separator
    assert '=IF(x>0; "tak"; "nie")' in content
    # ZAOKR: 3,14 is decimal (stays), but 14, 2 should be separator
    assert '=ZAOKR(3,14; 2)' in content
    # SUMA: 1.5 has period (English decimal), so comma is separator
    assert '=SUMA(A1; B1; 1.5)' in content


def test_whitelist_guard_odmawia_edycji_pliku_sql(tmp_path):
    """Skrypt nie moze edytowac plikow spoza whitelisty."""
    src = FIXTURES / "fixture_sql_NOT_TOUCH.md"
    dst = tmp_path / "fixture_sql.md"
    shutil.copy(src, dst)
    assert not pf.is_whitelisted(str(dst), "all")


def test_polonizuj_cwiczenie_obsluguje_wskazowki_jako_stringi(tmp_path):
    """Realny schemat: wskazowki to list[str] (nie list[dict]). Regression test po T11."""
    f = tmp_path / "test.json"
    f.write_text(json.dumps({
        "id": "TEST.X",
        "tagi": ["SUMIF"],
        "tresc": "=SUMIF(A:A;\"X\";B:B)",
        "odpowiedz": "=SUMIF(A:A;\"X\";B:B)",
        "wskazowki": [
            "SUMIF ma 3 argumenty: zakres_kryt, kryt, zakres_sum",
            "Pamietaj o cudzyslowach: SUMIF(A:A; \"tekst\"; B:B)"
        ],
        "typowe_bledy": [
            {"opis": "Pomylenie SUMIF z COUNTIF", "kara": -1}
        ]
    }, ensure_ascii=False))
    mapa = pf.load_mapa()
    pf.polonizuj_json_cwiczenie(f, mapa)
    data = json.loads(f.read_text())
    assert data["wskazowki"][0].startswith("SUMA.JEŻELI ma 3 argumenty")
    assert "SUMA.JEŻELI(A:A" in data["wskazowki"][1]
    assert "Pomylenie SUMA.JEŻELI z LICZ.JEŻELI" in data["typowe_bledy"][0]["opis"]
    assert data["typowe_bledy"][0]["kara"] == -1
    assert "SUMA.JEŻELI" in data["tagi"]
