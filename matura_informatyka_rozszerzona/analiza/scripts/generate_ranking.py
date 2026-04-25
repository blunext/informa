#!/usr/bin/env python3
"""
Generator raportu RANKING_ALGORYTMOW.md.

Czyta wszystkie pliki matura_*.json + algorytmy_rejestr.json i produkuje
markdown z:
- Streszczeniem (TOP 5, glowne wnioski)
- Rankingiem glownym (tabela: algorytm | kategoria | wystapien | punkty | %egz | przyklady)
- Rozbiciem per kategoria (klasyczne / techniki / struktury / wzorce)
- Heatmapa rok x algorytm (TOP 20)
- Rekomendacjami TIER 1/2/3 z linkami do podstawy programowej
- Top 5 kombinacjami 2-algorytmowymi
- Algorytmami nieuzywanymi (z podstawy programowej, ale brak w CKE)

Uzycie:
    python3 analiza/scripts/generate_ranking.py
Tworzy plik: analiza/RANKING_ALGORYTMOW.md
"""
from __future__ import annotations
import json
from pathlib import Path
from collections import Counter, defaultdict
from itertools import combinations

ROOT = Path(__file__).resolve().parent.parent
JSON_DIR = ROOT / "json"
REJESTR_PATH = JSON_DIR / "algorytmy_rejestr.json"
OUT_PATH = ROOT / "RANKING_ALGORYTMOW.md"

YEARS = [2014, 2015, 2016, 2017, 2018, 2019, 2020, 2021, 2022, 2023, 2024, 2025]


def load_data():
    with REJESTR_PATH.open(encoding="utf-8") as f:
        rejestr = json.load(f)

    subtasks = []
    for fpath in sorted(JSON_DIR.glob("matura_*.json")):
        if "indeks" in fpath.name:
            continue
        with fpath.open(encoding="utf-8") as f:
            data = json.load(f)
        rok = data.get("rok")
        for z in data.get("zadania", []):
            for p in z.get("podzadania", []):
                subtasks.append({
                    "id": p["id"],
                    "rok": rok,
                    "punkty": p.get("punkty", 0),
                    "typ_zadania": p.get("typ_zadania"),
                    "kategoria_zadania": p.get("kategoria"),
                    "algorytmy": p.get("algorytmy", []),
                })
    return rejestr, subtasks


def compute_stats(rejestr, subtasks):
    algos_info = rejestr["algorytmy"]
    by_alg = defaultdict(lambda: {
        "wystapienia": 0,
        "punkty": 0,
        "lata": set(),
        "examples": [],
        "lat_x_count": Counter(),
    })

    for st in subtasks:
        for tag in st["algorytmy"]:
            entry = by_alg[tag]
            entry["wystapienia"] += 1
            entry["punkty"] += st["punkty"]
            entry["lata"].add(st["rok"])
            entry["lat_x_count"][st["rok"]] += 1
            if len(entry["examples"]) < 3:
                entry["examples"].append(st["id"])

    return by_alg


def compute_combinations(subtasks):
    pair_counter = Counter()
    for st in subtasks:
        tags = sorted(set(st["algorytmy"]))
        for a, b in combinations(tags, 2):
            pair_counter[(a, b)] += 1
    return pair_counter


def fmt_pct(n, total):
    return f"{100 * n / total:.0f}%" if total else "0%"


def build_main_table(rejestr, by_alg) -> list[str]:
    algos_info = rejestr["algorytmy"]
    rows = []
    for nazwa, stats in by_alg.items():
        info = algos_info.get(nazwa, {})
        rows.append({
            "nazwa": nazwa,
            "kategoria": info.get("kategoria", "?"),
            "wystapienia": stats["wystapienia"],
            "punkty": stats["punkty"],
            "lata": len(stats["lata"]),
            "podstawa": info.get("podstawa") or "—",
            "examples": stats["examples"],
        })
    rows.sort(key=lambda r: (-r["wystapienia"], -r["punkty"], r["nazwa"]))
    return rows


def build_heatmap(by_alg, top_n=20) -> tuple[list[str], list[list[str]]]:
    top_algs = sorted(by_alg.items(), key=lambda kv: -kv[1]["wystapienia"])[:top_n]
    header = ["algorytm"] + [str(y) for y in YEARS]
    rows = []
    for nazwa, stats in top_algs:
        row = [nazwa]
        for y in YEARS:
            count = stats["lat_x_count"].get(y, 0)
            row.append(str(count) if count else "·")
        rows.append(row)
    return header, rows


def assign_tier(rows: list[dict], total_egzaminow: int) -> dict:
    tiers = {"TIER 1": [], "TIER 2": [], "TIER 3": []}
    for r in rows:
        wyst = r["wystapienia"]
        if wyst >= 30:
            tiers["TIER 1"].append(r)
        elif wyst >= 10:
            tiers["TIER 2"].append(r)
        elif wyst >= 1:
            tiers["TIER 3"].append(r)
    return tiers


def compute_tier_coverage(tier_algos: list[dict], subtasks: list) -> tuple[int, int]:
    """Zwraca (liczba_podzadan, suma_punktow) ktore zawieraja >=1 tag z tier_algos.
    Liczone unikalnie — jedno podzadanie liczone raz.
    """
    tier_set = {r["nazwa"] for r in tier_algos}
    covered_subs = 0
    covered_pts = 0
    for st in subtasks:
        if any(t in tier_set for t in st["algorytmy"]):
            covered_subs += 1
            covered_pts += st["punkty"]
    return covered_subs, covered_pts


def render_markdown(rejestr, subtasks, by_alg, pair_counter):
    total_subs = len(subtasks)
    tagged = sum(1 for s in subtasks if s["algorytmy"])
    total_pts = sum(s["punkty"] for s in subtasks)
    total_alg_uses = sum(stats["wystapienia"] for stats in by_alg.values())
    used_algs = set(by_alg.keys())
    all_algs = set(rejestr["algorytmy"].keys())
    unused_algs = sorted(all_algs - used_algs)

    rows = build_main_table(rejestr, by_alg)
    top5 = rows[:5]

    lines: list[str] = []
    lines.append("# Ranking algorytmow w zadaniach maturalnych CKE 2014-2025")
    lines.append("")
    lines.append(f"_Wygenerowano automatycznie z `analiza/scripts/generate_ranking.py`_")
    lines.append(f"_Zrodla: 30 plikow `matura_*.json` (641 podzadan) + `algorytmy_rejestr.json` (65 algorytmow)_")
    lines.append("")
    lines.append("## Streszczenie")
    lines.append("")
    lines.append(f"- **641 podzadan** sklasyfikowanych w 30 sesjach CKE (2014-2025), {total_pts} punktow lacznie.")
    lines.append(f"- **{tagged}/{total_subs} ({100*tagged/total_subs:.1f}%)** podzadan ma przynajmniej 1 tag algorytmu.")
    lines.append(f"- **{total_alg_uses}** lacznie wystapien tagow (srednio {total_alg_uses/tagged:.2f} na podzadanie z tagami).")
    lines.append(f"- **{len(used_algs)}/65** algorytmow z rejestru rzeczywiscie pojawia sie w zadaniach.")
    lines.append("")
    lines.append("**TOP 5 algorytmow** (wg liczby wystapien):")
    lines.append("")
    lines.append("| # | Algorytm | Kategoria | Wystapien | Punkty | Lat |")
    lines.append("|---|---|---|---:|---:|---:|")
    for i, r in enumerate(top5, 1):
        lines.append(f"| {i} | `{r['nazwa']}` | {r['kategoria']} | {r['wystapienia']} | {r['punkty']} | {r['lata']}/12 |")
    lines.append("")

    lines.append("**Glowne wnioski edukacyjne**:")
    lines.append("")
    lines.append("1. **Praca z plikiem dominuje**: `iteracja-po-pliku` jest w prawie kazdej sesji CKE — to umiejetnosc niezbedna.")
    lines.append("2. **SQL ma TIER 1** — `SQL-JOIN`, `SQL-aggregacja`, `SQL-GROUP-BY`, `SQL-WHERE` to filary zadania bazodanowego.")
    lines.append("3. **Sledzenie pseudokodu** rownie czeste co programowanie — wymagana zarowno teoria, jak i praktyka.")
    lines.append("4. **Konwersje systemow liczbowych** + **Horner** wystepuje regularnie — fundamenty teorii.")
    lines.append("5. **Programowanie dynamiczne** pojawilo sie po raz pierwszy w 2024 — nowy obszar wymagajacy uwagi.")
    lines.append("")

    lines.append("---")
    lines.append("")
    lines.append("## Ranking glowny")
    lines.append("")
    lines.append("Lista wszystkich algorytmow uzytych w zadaniach CKE 2014-2025, posortowana wg liczby wystapien.")
    lines.append("")
    lines.append("| # | Algorytm | Kategoria | Wystapien | Punkty | Lat | Podstawa | Przyklady |")
    lines.append("|---|---|---|---:|---:|---:|---|---|")
    for i, r in enumerate(rows, 1):
        examples = ", ".join(f"`{e}`" for e in r["examples"])
        podstawa = ", ".join(r["podstawa"]) if isinstance(r["podstawa"], list) else r["podstawa"]
        lines.append(f"| {i} | `{r['nazwa']}` | {r['kategoria']} | {r['wystapienia']} | {r['punkty']} | {r['lata']}/12 | {podstawa} | {examples} |")
    lines.append("")

    lines.append("---")
    lines.append("")
    lines.append("## Rozbicie per kategoria")
    lines.append("")
    by_kat = defaultdict(list)
    for r in rows:
        by_kat[r["kategoria"]].append(r)
    for kat in ["klasyczne", "techniki", "struktury", "wzorce"]:
        kat_rows = by_kat.get(kat, [])
        kat_pts = sum(r["punkty"] for r in kat_rows)
        kat_uses = sum(r["wystapienia"] for r in kat_rows)
        lines.append(f"### {kat.capitalize()} ({len(kat_rows)} algorytmow, {kat_uses} wystapien, {kat_pts} pkt)")
        lines.append("")
        lines.append("| Algorytm | Wystapien | Punkty | Lat |")
        lines.append("|---|---:|---:|---:|")
        for r in kat_rows:
            lines.append(f"| `{r['nazwa']}` | {r['wystapienia']} | {r['punkty']} | {r['lata']}/12 |")
        lines.append("")

    lines.append("---")
    lines.append("")
    lines.append("## Heatmapa: rok x algorytm (TOP 20)")
    lines.append("")
    lines.append("Liczba wystapien algorytmu w danym roku. `·` = brak.")
    lines.append("")
    header, hrows = build_heatmap(by_alg, top_n=20)
    lines.append("| " + " | ".join(header) + " |")
    lines.append("|" + "|".join(["---"] + [":-:" for _ in YEARS]) + "|")
    for r in hrows:
        lines.append("| " + " | ".join(r) + " |")
    lines.append("")

    lines.append("---")
    lines.append("")
    lines.append("## Rekomendacje kolejnosci nauki (TIER 1/2/3)")
    lines.append("")
    tiers = assign_tier(rows, 12)

    # Pokrycie unikalne (jedno podzadanie liczone raz)
    tier1_subs, tier1_pts_uniq = compute_tier_coverage(tiers["TIER 1"], subtasks)
    pct_tier1_subs = 100 * tier1_subs / total_subs
    pct_tier1_pts = 100 * tier1_pts_uniq / total_pts

    lines.append(f"### TIER 1 — Must Have ({len(tiers['TIER 1'])} algorytmow, ≥30 wystapien)")
    lines.append("")
    lines.append(f"**Znajomosc TIER 1 pozwala podejsc do {tier1_subs}/{total_subs} podzadan ({pct_tier1_subs:.1f}%) i {tier1_pts_uniq}/{total_pts} pkt ({pct_tier1_pts:.1f}%).**")
    lines.append(f"Te algorytmy musisz znac na 100%.")
    lines.append("")
    for r in tiers["TIER 1"]:
        lines.append(f"- `{r['nazwa']}` ({r['kategoria']}) — {r['wystapienia']}x, {r['lata']}/12 lat")
    lines.append("")

    tier12_subs, tier12_pts_uniq = compute_tier_coverage(tiers["TIER 1"] + tiers["TIER 2"], subtasks)
    pct_tier12_subs = 100 * tier12_subs / total_subs
    pct_tier12_pts = 100 * tier12_pts_uniq / total_pts

    lines.append(f"### TIER 2 — Powinno sie znac ({len(tiers['TIER 2'])} algorytmow, 10-29 wystapien)")
    lines.append("")
    lines.append(f"Razem TIER 1+2 pozwala podejsc do {tier12_subs}/{total_subs} podzadan ({pct_tier12_subs:.1f}%) i {tier12_pts_uniq}/{total_pts} pkt ({pct_tier12_pts:.1f}%).")
    lines.append("")
    for r in tiers["TIER 2"]:
        lines.append(f"- `{r['nazwa']}` ({r['kategoria']}) — {r['wystapienia']}x, {r['lata']}/12 lat")
    lines.append("")

    lines.append(f"### TIER 3 — Nice to have ({len(tiers['TIER 3'])} algorytmow, 1-9 wystapien)")
    lines.append("")
    lines.append("Rzadziej spotykane, ale mozna na nie trafic.")
    lines.append("")
    for r in tiers["TIER 3"]:
        lines.append(f"- `{r['nazwa']}` ({r['kategoria']}) — {r['wystapienia']}x, {r['punkty']} pkt")
    lines.append("")

    if unused_algs:
        rejestr_algos = rejestr["algorytmy"]
        lines.append(f"### Algorytmy z rejestru NIE testowane przez CKE (2014-2025)")
        lines.append("")
        lines.append(f"Algorytmy z podstawy programowej ktorych CKE nigdy nie pyta — niski priorytet nauki.")
        lines.append("")
        for n in unused_algs:
            info = rejestr_algos.get(n, {})
            podstawa = info.get("podstawa") or "—"
            podstawa = ", ".join(podstawa) if isinstance(podstawa, list) else podstawa
            lines.append(f"- `{n}` (podstawa: {podstawa}) — {info.get('definicja', '')}")
        lines.append("")

    lines.append("---")
    lines.append("")
    lines.append("## TOP 10 kombinacji 2-algorytmowych")
    lines.append("")
    lines.append("Pary tagow ktore najczesciej wystepuja razem w jednym podzadaniu.")
    lines.append("")
    lines.append("| # | Algorytm A | Algorytm B | Wspolnie |")
    lines.append("|---|---|---|---:|")
    for i, ((a, b), n) in enumerate(pair_counter.most_common(10), 1):
        lines.append(f"| {i} | `{a}` | `{b}` | {n} |")
    lines.append("")

    lines.append("---")
    lines.append("")
    lines.append("## Statystyki kategorii")
    lines.append("")
    lines.append("Liczba wystapien tagu = ile razy algorytm pojawia sie w klasyfikacji (1 podzadanie moze miec wiele tagow).")
    lines.append("")
    lines.append("| Kategoria | Algorytmow w rejestrze | Uzywanych | Wystapien | % wszystkich tagow |")
    lines.append("|---|---:|---:|---:|---:|")
    for kat in ["klasyczne", "techniki", "struktury", "wzorce"]:
        kat_in_rej = sum(1 for n, info in rejestr["algorytmy"].items() if info.get("kategoria") == kat)
        kat_used = len(by_kat.get(kat, []))
        kat_uses = sum(r["wystapienia"] for r in by_kat.get(kat, []))
        pct_uses = 100 * kat_uses / total_alg_uses if total_alg_uses else 0
        lines.append(f"| {kat} | {kat_in_rej} | {kat_used} | {kat_uses} | {pct_uses:.1f}% |")
    lines.append("")

    return "\n".join(lines) + "\n"


def main() -> int:
    rejestr, subtasks = load_data()
    by_alg = compute_stats(rejestr, subtasks)
    pair_counter = compute_combinations(subtasks)
    md = render_markdown(rejestr, subtasks, by_alg, pair_counter)
    OUT_PATH.write_text(md, encoding="utf-8")
    print(f"Raport zapisany: {OUT_PATH}")
    print(f"Rozmiar: {len(md)} znakow, {md.count(chr(10))} linii")
    return 0


if __name__ == "__main__":
    import sys
    sys.exit(main())
