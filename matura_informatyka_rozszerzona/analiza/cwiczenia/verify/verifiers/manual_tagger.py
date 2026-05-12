"""Manual tagger — marks exercises that require human review.

Includes sanity checks that catch common issues (empty/short answers,
missing expected content patterns) without full verification.
"""
import re
from .base import VerificationResult

# Map file type prefix to review hints
REVIEW_HINTS = {
    '01': 'Sprawdz krokowa tablice sledzenia algorytmu i poprawnosc kazdego kroku.',
    '02': 'Zweryfikuj poprawnosc zaprojektowanego algorytmu/pseudokodu — czy obsluguje przypadki brzegowe?',
    '03': 'Sprawdz poprawnosc analizy zlozonosci obliczeniowej i odpowiedzi na pytania o algorytm.',
    '04': 'Sprawdz kazde twierdzenie prawda/falsz — czy uzasadnienie jest poprawne?',
    '06': 'Sprawdz poprawnosc odpowiedzi dot. bezpieczenstwa informatycznego.',
    '15': 'Zweryfikuj formuly arkuszowe (SUMIFS, COUNTIFS, adresowanie) — brak silnika do auto-weryfikacji.',
    '16': 'Zweryfikuj formuly symulacji w arkuszu — sprawdz logike krokowa.',
    '17': 'Zweryfikuj formuly do wykresu i typ wykresu.',
    '18': 'Zweryfikuj podstawowe agregacje arkuszowe (SUMA, SREDNIA, MIN, MAX).',
    '19': 'Zweryfikuj formuly transformacji danych w arkuszu.',
}


def _has_markdown_table(text: str) -> bool:
    """Check if text contains at least one markdown table."""
    lines = text.split('\n')
    for i, line in enumerate(lines):
        line = line.strip()
        if (line.startswith('|') and line.endswith('|')
                and i + 1 < len(lines)):
            sep = lines[i + 1].strip()
            if sep.startswith('|') and re.match(r'^\|[\s\-:|]+\|$', sep):
                return True
    return False


def _has_code_block(text: str) -> bool:
    """Check if text contains a fenced code block."""
    return bool(re.search(r'```\w*\s*\n.*?```', text, re.DOTALL))


def _has_structured_content(text: str) -> bool:
    """Check if text contains structured content (tables, code, P/F, numbered steps, bold)."""
    if _has_markdown_table(text):
        return True
    if _has_code_block(text):
        return True
    # P/F answers (full words or bold shorthand)
    if re.search(r'\b(PRAWDA|FALSZ)\b', text) or re.search(r'\*\*[PF]\*\*', text):
        return True
    # Numbered steps with calculations (e.g. "1. sumaCyfr(47) -> ...")
    if len(re.findall(r'^\d+\.\s+\S', text, re.MULTILINE)) >= 2:
        return True
    # Bold markers indicating structured answer (e.g. **a)**, **Wynik:**)
    if len(re.findall(r'\*\*[^*]+\*\*', text)) >= 2:
        return True
    return False


def _count_pf_answers(text: str) -> int:
    """Count PRAWDA/FALSZ answers in text. Also matches **P**/**F** shorthand."""
    count = len(re.findall(r'\b(PRAWDA|FALSZ)\b', text))
    # Also match bold P/F shorthand in tables: **P** or **F**
    count += len(re.findall(r'\*\*([PF])\*\*', text))
    return count


def _count_pf_questions(text: str) -> int:
    """Count statements/questions in tresc that expect P/F answers."""
    count = len(re.findall(r'^\s*[a-z]\)', text, re.MULTILINE))
    count += len(re.findall(r'^\s*\d+[\.\)]\s', text, re.MULTILINE))
    return max(count, 0)


def _has_arkusz_formula(text: str) -> bool:
    """Check if text contains a spreadsheet formula (=FUNCTION or =cell_ref).

    Obsluguje polskie nazwy funkcji MS Excel PL (np. SUMA.JEŻELI, WYSZUKAJ.PIONOWO,
    USUŃ.ZBĘDNE.ODSTĘPY) — klasa znakow zawiera polskie wielkie litery i kropke.
    """
    # Named functions: =SUMA.WARUNKÓW(...), =JEŻELI(...), =VLOOKUP(...)
    if re.search(r'=\s*[A-ZŁŻŚĆŃÓĄĘŹ.]{2,}\s*\(', text):
        return True
    # Cell reference formulas: =C2+B3, =A1*B1
    if re.search(r'=\s*[A-Z]+\d+\s*[\+\-\*/]', text):
        return True
    # Simple cell copy: =A1
    if re.search(r'=\s*[A-Z]+\d+\s*$', text, re.MULTILINE):
        return True
    return False


def _run_sanity_checks(exercise: dict, prefix: str) -> tuple[list[str], list[str]]:
    """Run sanity checks on a MANUAL_REVIEW exercise.

    Returns (passed_checks, failed_checks) as lists of descriptions.
    """
    odpowiedz = exercise.get('odpowiedz', '')
    tresc = exercise.get('tresc', '')
    passed = []
    failed = []

    # === Universal checks (all MANUAL_REVIEW) ===

    # 1. odpowiedz length > 100
    if len(odpowiedz) > 100:
        passed.append('odpowiedz >100 znakow')
    else:
        failed.append(f'odpowiedz za krotka ({len(odpowiedz)} znakow, potrzeba >100)')

    # 2. odpowiedz contains structured content
    if _has_structured_content(odpowiedz):
        passed.append('odpowiedz zawiera tresc strukturalna')
    else:
        failed.append('odpowiedz brak tresci strukturalnej (tabela/kod/P-F/lista/bold)')

    # === Per-type checks ===

    if prefix == '01':
        # sledzenie: odpowiedz should contain table, code block, or numbered steps
        has_steps = len(re.findall(r'^\d+\.\s+\S', odpowiedz, re.MULTILINE)) >= 2
        has_bullet_steps = len(re.findall(r'^- \w+.*->.*', odpowiedz, re.MULTILINE)) >= 2
        if (_has_markdown_table(odpowiedz) or _has_code_block(odpowiedz)
                or has_steps or has_bullet_steps):
            passed.append('sledzenie: odpowiedz zawiera sledzenie krokowe')
        else:
            failed.append('sledzenie: brak tabeli, drzewa lub listy krokow w odpowiedzi')

    elif prefix == '02':
        # projektowanie: odpowiedz should contain code block (pseudocode/C++)
        if _has_code_block(odpowiedz):
            passed.append('projektowanie: odpowiedz zawiera blok kodu')
        else:
            failed.append('projektowanie: brak bloku kodu w odpowiedzi')

    elif prefix == '03':
        # analiza: odpowiedz should reference complexity or analysis terms
        analysis_patterns = [
            r'O\(', r'zlozonosc', r'pesymist', r'optymist',
            r'liniow', r'kwadrat', r'logarytm', r'stabl', r'wykona',
            r'porownani', r'operacj', r'krok',
        ]
        has_analysis = any(
            re.search(p, odpowiedz, re.IGNORECASE) for p in analysis_patterns
        )
        if has_analysis:
            passed.append('analiza: odpowiedz zawiera terminologie analityczna')
        else:
            failed.append('analiza: brak terminologii analitycznej w odpowiedzi')

    elif prefix == '04':
        # test P/F: count of PRAWDA/FALSZ answers >= questions in tresc
        n_answers = _count_pf_answers(odpowiedz)
        n_questions = _count_pf_questions(tresc)
        if n_answers > 0 and n_answers >= n_questions:
            passed.append(f'test P/F: {n_answers} odpowiedzi >= {n_questions} pytan')
        elif n_answers > 0:
            failed.append(
                f'test P/F: {n_answers} odpowiedzi P/F < {n_questions} pytan w tresci')
        else:
            failed.append('test P/F: brak odpowiedzi PRAWDA/FALSZ w odpowiedzi')

    elif prefix == '06':
        # bezpieczenstwo: odpowiedz >= 200 chars (descriptive answers)
        if len(odpowiedz) >= 200:
            passed.append(f'bezpieczenstwo: odpowiedz ma {len(odpowiedz)} znakow (>=200)')
        else:
            failed.append(
                f'bezpieczenstwo: odpowiedz za krotka ({len(odpowiedz)} znakow, potrzeba >=200)')

    elif prefix in ('15', '16', '18'):
        # arkusz formuly: odpowiedz should contain formula
        if _has_arkusz_formula(odpowiedz):
            passed.append('arkusz: odpowiedz zawiera formule')
        else:
            failed.append('arkusz: brak formuly (=FUNKCJA lub =komorka) w odpowiedzi')

    elif prefix == '17':
        # wykresy: odpowiedz should mention "wykres"
        if 'wykres' in odpowiedz.lower():
            passed.append('wykres: odpowiedz wspomina o wykresie')
        else:
            failed.append('wykres: brak slowa "wykres" w odpowiedzi')

    elif prefix == '19':
        # transformacja: odpowiedz >= 100 chars
        if len(odpowiedz) >= 100:
            passed.append(f'transformacja: odpowiedz ma {len(odpowiedz)} znakow (>=100)')
        else:
            failed.append(
                f'transformacja: odpowiedz za krotka ({len(odpowiedz)} znakow, potrzeba >=100)')

    return passed, failed


def tag_manual(exercise: dict, file_typ: str) -> VerificationResult:
    """Tag exercise as requiring manual review, with sanity checks."""
    eid = exercise['id']
    prefix = file_typ.split('_')[0]

    kategoria = 'TEORIA'
    if 15 <= int(prefix) <= 19:
        kategoria = 'ARKUSZ'

    hint = REVIEW_HINTS.get(prefix, 'Wymaga recznej weryfikacji — brak metody automatycznej.')

    # Run sanity checks
    passed, failed = _run_sanity_checks(exercise, prefix)

    if failed:
        # Sanity check failure — report as FAIL
        details = '; '.join(failed)
        return VerificationResult(
            exercise_id=eid,
            file_typ=file_typ,
            kategoria=kategoria,
            status='FAIL',
            method='manual_sanity',
            details=f'Sanity FAIL: {details}',
        )

    # All sanity checks passed
    n_checks = len(passed)
    details = f'{hint} ({n_checks} sanity checks passed)'
    return VerificationResult(
        exercise_id=eid,
        file_typ=file_typ,
        kategoria=kategoria,
        status='MANUAL_REVIEW',
        method='manual_sanity',
        details=details,
    )
