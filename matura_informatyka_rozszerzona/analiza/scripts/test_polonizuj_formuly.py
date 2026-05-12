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
