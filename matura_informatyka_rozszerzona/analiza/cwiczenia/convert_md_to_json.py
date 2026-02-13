#!/usr/bin/env python3
"""Convert exercise MD files to structured JSON."""
import re
import json
import glob
import os

SRC_DIR = os.path.join(os.path.dirname(__file__), "wg_typu")
DST_DIR = os.path.join(os.path.dirname(__file__), "json")


def parse_header(text):
    """Parse the file header before exercises."""
    header = {}

    # typ: from "Typ zadania: **xxx**"
    m = re.search(r'Typ zadania:\s*\*\*(\S+)\*\*', text)
    header['typ'] = m.group(1) if m else ""

    # nazwa: from "# XX. Name"
    m = re.search(r'^#\s+\d+\.\s+(.+)', text, re.MULTILINE)
    header['nazwa'] = m.group(1).strip() if m else ""

    # kategoria: from "Kategoria: XXX"
    m = re.search(r'Kategoria:\s*(\S+)', text)
    header['kategoria'] = m.group(1).strip() if m else ""

    # czestotliwosc: from "Czestotliwosc: X/11 lat"
    m = re.search(r'Czestotliwosc:\s*([\d/]+\s*lat)', text)
    header['czestotliwosc'] = m.group(1).strip() if m else ""

    # punkty_lacznie: from "Laczna punktacja: N pkt"
    m = re.search(r'Laczna punktacja:\s*(\d+)\s*pkt', text)
    header['punkty_lacznie'] = int(m.group(1)) if m else 0

    # tagi_globalne: from "## Umiejetnosci" section - backtick-enclosed tags
    # Some files have blank line after header, some don't — use \n+ to match both
    m = re.search(r'## Umiejetnosci.*?\n+(.+?)(\n\n|\n---)', text, re.DOTALL)
    if m:
        header['tagi_globalne'] = re.findall(r'`([^`]+)`', m.group(1))
    else:
        header['tagi_globalne'] = []

    return header


def split_exercises(text):
    """Split text into individual exercise blocks."""
    # Split on "### Cwiczenie X.Y"
    parts = re.split(r'(?=^### Cwiczenie \d+\.\d+)', text, flags=re.MULTILINE)
    return [p for p in parts if p.strip().startswith('### Cwiczenie')]


def extract_details_block(text, summary_name):
    """Extract content from a <details> block with given summary."""
    # Match <details>\n<summary>NAME</summary>\n...\n</details>
    pattern = (
        r'<details>\s*\n\s*<summary>\s*'
        + re.escape(summary_name)
        + r'\s*</summary>\s*\n(.*?)</details>'
    )
    m = re.search(pattern, text, re.DOTALL)
    if m:
        return m.group(1).strip()
    return ""


def parse_wskazowki(raw):
    """Parse hints into a list of 3 strings."""
    if not raw:
        return []
    hints = []
    # Match numbered items: 1. **Label**: text...
    # They may span multiple lines until the next numbered item or end
    items = re.split(r'\n\d+\.\s+', '\n' + raw)
    for item in items:
        item = item.strip()
        if item:
            hints.append(item)
    return hints


def parse_typowe_bledy(raw):
    """Parse common errors into list of {opis, kara}."""
    if not raw:
        return []
    errors = []
    # Errors can start with "- **..." (bullet) or "1. **..." (numbered)
    lines = raw.split('\n')
    current = ""
    for line in lines:
        line = line.strip()
        # Detect new error item: bullet or numbered list
        is_new_item = False
        content = ""
        if line.startswith('- '):
            is_new_item = True
            content = line[2:]
        elif re.match(r'^\d+\.\s+', line):
            is_new_item = True
            content = re.sub(r'^\d+\.\s+', '', line)

        if is_new_item:
            if current:
                errors.append(parse_single_error(current))
            current = content
        elif current and line:
            current += " " + line
    if current:
        errors.append(parse_single_error(current))
    return errors


def parse_single_error(text):
    """Parse a single error line into {opis, kara}."""
    # Try to extract CKE penalty — multiple formats:
    # "CKE: -1 pkt" / "CKE: -1 pkt za bledny wynik" / "CKE: blad logiczny, -2 pkt"
    m = re.search(r'CKE:\s*(.+?)\.?\s*$', text)
    kara = ""
    if m:
        kara_text = m.group(1).strip().rstrip('.')
        # Extract the penalty number if present
        penalty = re.search(r'-[\d.]+\s*pkt', kara_text)
        kara = kara_text if penalty else ""

    # Description is everything, cleaned up
    opis = text.strip()
    # Remove trailing CKE part for opis
    if m:
        opis = text[:m.start()].strip().rstrip('.')

    return {"opis": opis, "kara": kara}


def parse_exercise(block):
    """Parse a single exercise block."""
    ex = {}

    # Header line: ### Cwiczenie X.Y (trudnosc: LEVEL, ~N pkt)
    m = re.match(
        r'### Cwiczenie (\d+\.\d+)\s*\(trudnosc:\s*([^,]+),\s*~(\d+)\s*pkt\)',
        block
    )
    if m:
        ex['id'] = m.group(1)
        ex['trudnosc'] = m.group(2).strip()
        ex['punkty'] = int(m.group(3))
    else:
        # Fallback: try without tilde
        m = re.match(
            r'### Cwiczenie (\d+\.\d+)\s*\(trudnosc:\s*([^,]+),\s*(\d+)\s*pkt\)',
            block
        )
        if m:
            ex['id'] = m.group(1)
            ex['trudnosc'] = m.group(2).strip()
            ex['punkty'] = int(m.group(3))
        else:
            ex['id'] = ""
            ex['trudnosc'] = ""
            ex['punkty'] = 0

    # Zrodlo: **Zrodlo inspiracji**: text
    m = re.search(r'\*\*Zrodlo inspiracji\*\*:\s*(.+)', block)
    ex['zrodlo'] = m.group(1).strip() if m else ""

    # Tagi: **Tagi**: `tag1` `tag2`
    m = re.search(r'\*\*Tagi\*\*:\s*(.+)', block)
    if m:
        ex['tagi'] = re.findall(r'`([^`]+)`', m.group(1))
    else:
        ex['tagi'] = []

    # Tresc: everything between Tagi line and first <details>
    # Find end of Tagi line
    tagi_match = re.search(r'\*\*Tagi\*\*:\s*.+\n', block)
    details_match = re.search(r'<details>', block)
    if tagi_match and details_match:
        tresc = block[tagi_match.end():details_match.start()].strip()
        ex['tresc'] = tresc
    else:
        ex['tresc'] = ""

    # Wskazowki
    raw_wskazowki = extract_details_block(block, "Wskazowki")
    ex['wskazowki'] = parse_wskazowki(raw_wskazowki)

    # Odpowiedz
    ex['odpowiedz'] = extract_details_block(block, "Odpowiedz")

    # Typowe bledy
    raw_bledy = extract_details_block(block, "Typowe bledy")
    ex['typowe_bledy'] = parse_typowe_bledy(raw_bledy)

    return ex


def convert_file(md_path, json_path):
    """Convert a single MD file to JSON."""
    with open(md_path, 'r', encoding='utf-8') as f:
        text = f.read()

    # Get the file number prefix for typ field
    basename = os.path.basename(md_path).replace('.md', '')

    header = parse_header(text)
    header['typ'] = basename  # e.g. "01_sledzenie_algorytmu"

    exercises = split_exercises(text)
    parsed = []
    for ex_block in exercises:
        parsed.append(parse_exercise(ex_block))

    result = {
        "typ": header['typ'],
        "nazwa": header['nazwa'],
        "kategoria": header['kategoria'],
        "czestotliwosc": header['czestotliwosc'],
        "punkty_lacznie": header['punkty_lacznie'],
        "tagi_globalne": header['tagi_globalne'],
        "cwiczenia": parsed
    }

    with open(json_path, 'w', encoding='utf-8') as f:
        json.dump(result, f, ensure_ascii=False, indent=2)

    return len(parsed)


def main():
    os.makedirs(DST_DIR, exist_ok=True)

    md_files = sorted(glob.glob(os.path.join(SRC_DIR, "[0-9]*.md")))
    print(f"Found {len(md_files)} MD files to convert.\n")

    total_exercises = 0
    issues = []

    for md_path in md_files:
        basename = os.path.basename(md_path).replace('.md', '')
        json_path = os.path.join(DST_DIR, basename + '.json')

        count = convert_file(md_path, json_path)
        total_exercises += count

        status = "OK" if count == 10 else f"WARNING: {count} exercises"
        print(f"  {basename}: {count} exercises — {status}")

        if count != 10:
            issues.append((basename, count))

        # Validate non-empty fields
        with open(json_path, 'r', encoding='utf-8') as f:
            data = json.load(f)
        for i, ex in enumerate(data['cwiczenia']):
            problems = []
            if not ex.get('tresc'):
                problems.append('empty tresc')
            if not ex.get('wskazowki'):
                problems.append('empty wskazowki')
            if not ex.get('odpowiedz'):
                problems.append('empty odpowiedz')
            if not ex.get('id'):
                problems.append('empty id')
            if problems:
                issues.append((f"{basename} ex#{i+1}", ', '.join(problems)))

    print(f"\nTotal: {len(md_files)} files, {total_exercises} exercises")

    if issues:
        print(f"\n⚠ Issues ({len(issues)}):")
        for name, detail in issues:
            print(f"  - {name}: {detail}")
    else:
        print("\n✓ All files converted successfully with 10 exercises each.")


if __name__ == '__main__':
    main()
