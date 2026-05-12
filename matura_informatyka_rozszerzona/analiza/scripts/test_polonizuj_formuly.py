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
