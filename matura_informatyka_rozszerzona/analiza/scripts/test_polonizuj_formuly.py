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
