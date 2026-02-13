#!/usr/bin/env python3
"""Thorough validation: compare every MD exercise with its JSON counterpart."""
import re
import json
import glob
import os

SRC_DIR = os.path.join(os.path.dirname(__file__), "wg_typu")
DST_DIR = os.path.join(os.path.dirname(__file__), "json")

issues = []
warnings = []
stats = {"files": 0, "exercises": 0, "fields_checked": 0}


def issue(file, ex_id, msg):
    issues.append(f"[ERROR] {file} / {ex_id}: {msg}")


def warn(file, ex_id, msg):
    warnings.append(f"[WARN]  {file} / {ex_id}: {msg}")


def extract_md_exercises(text):
    """Re-parse MD to get raw exercise data for comparison."""
    parts = re.split(r'(?=^### Cwiczenie \d+\.\d+)', text, flags=re.MULTILINE)
    exercises = []
    for p in parts:
        if not p.strip().startswith('### Cwiczenie'):
            continue
        ex = {}

        # id, trudnosc, punkty
        m = re.match(r'### Cwiczenie (\d+\.\d+)\s*\(trudnosc:\s*([^,]+),\s*~?(\d+)\s*pkt\)', p)
        if m:
            ex['id'] = m.group(1)
            ex['trudnosc'] = m.group(2).strip()
            ex['punkty'] = int(m.group(3))
        else:
            ex['id'] = '???'
            ex['trudnosc'] = '???'
            ex['punkty'] = -1

        # zrodlo
        m = re.search(r'\*\*Zrodlo inspiracji\*\*:\s*(.+)', p)
        ex['zrodlo'] = m.group(1).strip() if m else ""

        # tagi
        m = re.search(r'\*\*Tagi\*\*:\s*(.+)', p)
        ex['tagi'] = sorted(re.findall(r'`([^`]+)`', m.group(1))) if m else []

        # tresc: between Tagi line and first <details>
        tagi_match = re.search(r'\*\*Tagi\*\*:\s*.+\n', p)
        details_match = re.search(r'<details>', p)
        if tagi_match and details_match:
            ex['tresc'] = p[tagi_match.end():details_match.start()].strip()
        else:
            ex['tresc'] = ""

        # wskazowki block
        m = re.search(r'<details>\s*\n\s*<summary>\s*Wskazowki\s*</summary>\s*\n(.*?)</details>', p, re.DOTALL)
        ex['wskazowki_raw'] = m.group(1).strip() if m else ""

        # Count numbered hints
        if ex['wskazowki_raw']:
            ex['wskazowki_count'] = len(re.findall(r'^\d+\.\s+', ex['wskazowki_raw'], re.MULTILINE))
        else:
            ex['wskazowki_count'] = 0

        # odpowiedz block
        m = re.search(r'<details>\s*\n\s*<summary>\s*Odpowiedz\s*</summary>\s*\n(.*?)</details>', p, re.DOTALL)
        ex['odpowiedz'] = m.group(1).strip() if m else ""

        # typowe_bledy - count bullet points
        m = re.search(r'<details>\s*\n\s*<summary>\s*Typowe bledy\s*</summary>\s*\n(.*?)</details>', p, re.DOTALL)
        bledy_raw = m.group(1).strip() if m else ""
        # Count both bullet (- **) and numbered (1. **) error items
        ex['bledy_count'] = len(re.findall(r'^(?:- |\d+\.\s+)\*\*', bledy_raw, re.MULTILINE))
        ex['bledy_raw'] = bledy_raw

        exercises.append(ex)
    return exercises


def validate_file(basename):
    md_path = os.path.join(SRC_DIR, basename + '.md')
    json_path = os.path.join(DST_DIR, basename + '.json')

    if not os.path.exists(json_path):
        issue(basename, '-', f"JSON file missing: {json_path}")
        return

    with open(md_path, 'r', encoding='utf-8') as f:
        md_text = f.read()
    with open(json_path, 'r', encoding='utf-8') as f:
        jdata = json.load(f)

    stats['files'] += 1

    # --- Header validation ---
    # typ
    if jdata['typ'] != basename:
        issue(basename, 'header', f"typ mismatch: JSON={jdata['typ']} expected={basename}")

    # nazwa not empty
    if not jdata['nazwa']:
        issue(basename, 'header', "empty nazwa")

    # kategoria
    valid_kat = {'TEORIA', 'IMPLEMENTACJA', 'ARKUSZ', 'SQL'}
    if jdata['kategoria'] not in valid_kat:
        issue(basename, 'header', f"invalid kategoria: {jdata['kategoria']}")

    # czestotliwosc format
    if not re.match(r'\d+/\d+ lat', jdata['czestotliwosc']):
        issue(basename, 'header', f"bad czestotliwosc format: {jdata['czestotliwosc']}")

    # punkty_lacznie
    if jdata['punkty_lacznie'] <= 0:
        issue(basename, 'header', f"punkty_lacznie={jdata['punkty_lacznie']}")

    # tagi_globalne
    if len(jdata['tagi_globalne']) == 0:
        issue(basename, 'header', "empty tagi_globalne")

    # --- Exercise count ---
    md_exercises = extract_md_exercises(md_text)
    json_exercises = jdata['cwiczenia']

    if len(md_exercises) != len(json_exercises):
        issue(basename, '-', f"exercise count: MD={len(md_exercises)} JSON={len(json_exercises)}")
        return

    # --- Per-exercise validation ---
    for md_ex, json_ex in zip(md_exercises, json_exercises):
        eid = md_ex['id']
        stats['exercises'] += 1

        # id
        if json_ex['id'] != md_ex['id']:
            issue(basename, eid, f"id mismatch: JSON={json_ex['id']}")
        stats['fields_checked'] += 1

        # trudnosc
        if json_ex['trudnosc'] != md_ex['trudnosc']:
            issue(basename, eid, f"trudnosc: JSON='{json_ex['trudnosc']}' MD='{md_ex['trudnosc']}'")
        stats['fields_checked'] += 1

        # punkty
        if json_ex['punkty'] != md_ex['punkty']:
            issue(basename, eid, f"punkty: JSON={json_ex['punkty']} MD={md_ex['punkty']}")
        stats['fields_checked'] += 1

        # zrodlo
        if json_ex['zrodlo'] != md_ex['zrodlo']:
            issue(basename, eid, f"zrodlo mismatch:\n  JSON: {json_ex['zrodlo']}\n  MD:   {md_ex['zrodlo']}")
        stats['fields_checked'] += 1

        # tagi (sorted comparison)
        json_tagi = sorted(json_ex.get('tagi', []))
        md_tagi = md_ex['tagi']
        if json_tagi != md_tagi:
            issue(basename, eid, f"tagi mismatch:\n  JSON: {json_tagi}\n  MD:   {md_tagi}")
        stats['fields_checked'] += 1

        # tresc - compare stripped
        json_tresc = json_ex.get('tresc', '').strip()
        md_tresc = md_ex['tresc'].strip()
        if json_tresc != md_tresc:
            # Check if it's just whitespace differences
            if json_tresc.replace(' ', '').replace('\n', '') == md_tresc.replace(' ', '').replace('\n', ''):
                warn(basename, eid, "tresc: minor whitespace difference")
            else:
                # Find first difference
                min_len = min(len(json_tresc), len(md_tresc))
                diff_pos = min_len
                for i in range(min_len):
                    if json_tresc[i] != md_tresc[i]:
                        diff_pos = i
                        break
                issue(basename, eid,
                      f"tresc CONTENT mismatch at char {diff_pos}:\n"
                      f"  JSON[{diff_pos}:{diff_pos+40}]: {repr(json_tresc[diff_pos:diff_pos+40])}\n"
                      f"  MD  [{diff_pos}:{diff_pos+40}]: {repr(md_tresc[diff_pos:diff_pos+40])}\n"
                      f"  JSON len={len(json_tresc)}, MD len={len(md_tresc)}")
        if not json_tresc:
            issue(basename, eid, "tresc is EMPTY")
        stats['fields_checked'] += 1

        # wskazowki count
        json_hints = json_ex.get('wskazowki', [])
        if len(json_hints) != md_ex['wskazowki_count']:
            issue(basename, eid, f"wskazowki count: JSON={len(json_hints)} MD={md_ex['wskazowki_count']}")
        if len(json_hints) == 0:
            issue(basename, eid, "wskazowki EMPTY")
        # Check each hint is non-trivial (>10 chars)
        for i, h in enumerate(json_hints):
            if len(h) < 10:
                warn(basename, eid, f"wskazowki[{i}] very short ({len(h)} chars): {repr(h[:50])}")
        stats['fields_checked'] += 1

        # odpowiedz - compare content
        json_odp = json_ex.get('odpowiedz', '').strip()
        md_odp = md_ex['odpowiedz'].strip()
        if json_odp != md_odp:
            if json_odp.replace(' ', '').replace('\n', '') == md_odp.replace(' ', '').replace('\n', ''):
                warn(basename, eid, "odpowiedz: minor whitespace difference")
            else:
                min_len = min(len(json_odp), len(md_odp))
                diff_pos = min_len
                for i in range(min_len):
                    if json_odp[i] != md_odp[i]:
                        diff_pos = i
                        break
                issue(basename, eid,
                      f"odpowiedz CONTENT mismatch at char {diff_pos}:\n"
                      f"  JSON len={len(json_odp)}, MD len={len(md_odp)}")
        if not json_odp:
            issue(basename, eid, "odpowiedz is EMPTY")
        # Check reasonable length
        if len(json_odp) < 50:
            warn(basename, eid, f"odpowiedz very short ({len(json_odp)} chars)")
        stats['fields_checked'] += 1

        # typowe_bledy count
        json_bledy = json_ex.get('typowe_bledy', [])
        if len(json_bledy) != md_ex['bledy_count']:
            issue(basename, eid, f"typowe_bledy count: JSON={len(json_bledy)} MD={md_ex['bledy_count']}")
        if len(json_bledy) == 0:
            issue(basename, eid, "typowe_bledy EMPTY")
        # Check each error has opis and kara
        for i, b in enumerate(json_bledy):
            if not b.get('opis'):
                issue(basename, eid, f"typowe_bledy[{i}]: empty opis")
            if not b.get('kara'):
                warn(basename, eid, f"typowe_bledy[{i}]: empty kara — '{b.get('opis', '')[:50]}'")
        stats['fields_checked'] += 1

        # punkty sum sanity
        if json_ex['punkty'] < 1 or json_ex['punkty'] > 10:
            warn(basename, eid, f"unusual punkty value: {json_ex['punkty']}")

        # trudnosc valid values
        valid_trudnosc = {'latwe', 'srednie', 'trudne', 'srednie-trudne'}
        if json_ex['trudnosc'] not in valid_trudnosc:
            warn(basename, eid, f"unusual trudnosc: '{json_ex['trudnosc']}'")


def main():
    md_files = sorted(glob.glob(os.path.join(SRC_DIR, "[0-9]*.md")))
    json_files = sorted(glob.glob(os.path.join(DST_DIR, "[0-9]*.json")))

    print(f"Source MD files: {len(md_files)}")
    print(f"Target JSON files: {len(json_files)}")

    # Check for missing JSONs
    md_basenames = {os.path.basename(f).replace('.md', '') for f in md_files}
    json_basenames = {os.path.basename(f).replace('.json', '') for f in json_files}
    missing = md_basenames - json_basenames
    extra = json_basenames - md_basenames
    if missing:
        for m in sorted(missing):
            issue(m, '-', "MD exists but JSON missing")
    if extra:
        for e in sorted(extra):
            warn(e, '-', "JSON exists but no MD source")

    # Validate each file
    for md_path in md_files:
        basename = os.path.basename(md_path).replace('.md', '')
        validate_file(basename)

    # --- Report ---
    print(f"\nValidated: {stats['files']} files, {stats['exercises']} exercises, {stats['fields_checked']} field checks")

    if issues:
        print(f"\n{'='*60}")
        print(f"ERRORS: {len(issues)}")
        print(f"{'='*60}")
        for i in issues:
            print(i)

    if warnings:
        print(f"\n{'='*60}")
        print(f"WARNINGS: {len(warnings)}")
        print(f"{'='*60}")
        for w in warnings:
            print(w)

    if not issues and not warnings:
        print("\n✅ PERFECT — all fields match between MD and JSON, zero issues.")
    elif not issues:
        print(f"\n✅ No errors. {len(warnings)} minor warnings (see above).")
    else:
        print(f"\n❌ {len(issues)} errors found — need fixing!")

    return len(issues)


if __name__ == '__main__':
    exit(main())
