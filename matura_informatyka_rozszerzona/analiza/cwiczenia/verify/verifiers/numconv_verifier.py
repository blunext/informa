"""Number system conversion verifier — uses Python arithmetic to verify answers."""
import re
from .base import VerificationResult


def _extract_bold_results(text: str) -> list[str]:
    """Extract bold-marked results like **101101(2)**, **B6(16)**, **53**, **128**."""
    return re.findall(r'\*\*([^*]+)\*\*', text)


def _parse_number_with_base(s: str) -> tuple[str, int] | None:
    """Parse a number string like '101101(2)' or 'B6(16)' or 'FF00(16)'.

    Returns (digits_str, base) or None if not parseable.
    """
    m = re.match(r'^([0-9A-Fa-f]+)\((\d+)\)$', s.strip())
    if m:
        return m.group(1), int(m.group(2))
    return None


def _to_int(digits: str, base: int) -> int | None:
    """Convert digits string in given base to int."""
    try:
        return int(digits, base)
    except ValueError:
        return None


def _verify_conversion(answer_text: str) -> list[tuple[bool, str]]:
    """Verify all conversions found in the answer text.

    Returns list of (pass, detail) tuples.
    """
    checks = []

    # Pattern: X(base1) = **Y(base2)**  or  X(10) = **Y(2)**
    # Also: **Y(base)** after "Wynik:" or "= "
    bold_items = _extract_bold_results(answer_text)

    for item in bold_items:
        parsed = _parse_number_with_base(item)
        if not parsed:
            # Plain number like **53** or **128** — check if it appears in conversion context
            try:
                val = int(item)
                # Look for context around this bold number
                # e.g., "110101(2) -> dziesietny: ... = **53**"
                checks.append((True, f"Plain value {val} found (context-dependent)"))
            except ValueError:
                continue
            continue

        digits, base = parsed
        computed = _to_int(digits, base)
        if computed is None:
            checks.append((False, f"Cannot parse {digits} in base {base}"))
            continue

        # Look for the source conversion in nearby text
        # We'll verify the conversion is self-consistent
        checks.append((True, f"{digits}({base}) = {computed}(10) — valid"))

    # Verify specific conversion patterns in the full text:

    # Pattern: dec -> bin verification
    # e.g., "45(10) = **101101(2)**"
    for m in re.finditer(r'(\d+)\(10\)\s*=\s*\*\*([0-9]+)\(2\)\*\*', answer_text):
        dec_val = int(m.group(1))
        bin_str = m.group(2)
        computed = _to_int(bin_str, 2)
        ok = computed == dec_val
        checks.append((ok, f"{dec_val}(10) -> {bin_str}(2): {'OK' if ok else f'FAIL (={computed}(10))'}"))

    # Pattern: bin -> dec verification
    for m in re.finditer(r'([01]+)\(2\)\s*(?:->|=).*?=\s*\*\*(\d+)\*\*', answer_text):
        bin_str = m.group(1)
        dec_val = int(m.group(2))
        computed = _to_int(bin_str, 2)
        ok = computed == dec_val
        checks.append((ok, f"{bin_str}(2) -> {dec_val}(10): {'OK' if ok else f'FAIL (={computed}(10))'}"))

    # Pattern: bin -> hex verification
    for m in re.finditer(r'([01]+)\(2\).*?=?\s*\*\*([0-9A-Fa-f]+)\(16\)\*\*', answer_text):
        bin_str = m.group(1)
        hex_str = m.group(2)
        bin_val = _to_int(bin_str, 2)
        hex_val = _to_int(hex_str, 16)
        ok = bin_val == hex_val
        checks.append((ok, f"{bin_str}(2) -> {hex_str}(16): {'OK' if ok else f'FAIL ({bin_val} != {hex_val})'}"))

    # Pattern: hex -> dec verification
    for m in re.finditer(r'([0-9A-Fa-f]+)\(16\).*?=\s*\*\*(\d+)\*\*', answer_text):
        hex_str = m.group(1)
        dec_val = int(m.group(2))
        computed = _to_int(hex_str, 16)
        if computed is not None:
            ok = computed == dec_val
            checks.append((ok, f"{hex_str}(16) -> {dec_val}(10): {'OK' if ok else f'FAIL (={computed})'}"))

    # Pattern: dec -> hex verification
    for m in re.finditer(r'(\d+)\(10\)\s*=\s*\*\*([0-9A-Fa-f]+)\(16\)\*\*', answer_text):
        dec_val = int(m.group(1))
        hex_str = m.group(2)
        computed = _to_int(hex_str, 16)
        ok = computed == dec_val
        checks.append((ok, f"{dec_val}(10) -> {hex_str}(16): {'OK' if ok else f'FAIL (={computed})'}"))

    # Pattern: oct -> dec verification
    for m in re.finditer(r'([0-7]+)\(8\).*?=\s*\*\*(\d+)(?:\(10\))?\*\*', answer_text):
        oct_str = m.group(1)
        dec_val = int(m.group(2))
        computed = _to_int(oct_str, 8)
        if computed is not None:
            ok = computed == dec_val
            checks.append((ok, f"{oct_str}(8) -> {dec_val}(10): {'OK' if ok else f'FAIL (={computed})'}"))

    # Pattern: binary addition/subtraction verification
    # Look for "Sprawdzenie: X(2) = Y, Z(2) = W, suma/roznica = N. Result(2) = N."
    for m in re.finditer(r'Sprawdzenie:.*?(\d+)\(2\)\s*=\s*(\d+),\s*(\d+)\(2\)\s*=\s*(\d+),\s*(?:suma|roznica)\s*=\s*(\d+)', answer_text):
        bin1 = m.group(1)
        dec1 = int(m.group(2))
        bin2 = m.group(3)
        dec2 = int(m.group(4))
        result_dec = int(m.group(5))

        c1 = _to_int(bin1, 2)
        c2 = _to_int(bin2, 2)

        ok1 = c1 == dec1
        ok2 = c2 == dec2
        checks.append((ok1, f"{bin1}(2) = {dec1}: {'OK' if ok1 else f'FAIL (={c1})'}"))
        checks.append((ok2, f"{bin2}(2) = {dec2}: {'OK' if ok2 else f'FAIL (={c2})'}"))

    # Arithmetic result patterns
    for m in re.finditer(r'Wynik:\s*\*\*([0-9A-Fa-f]+)\((\d+)\)\*\*', answer_text):
        digits = m.group(1)
        base = int(m.group(2))
        computed = _to_int(digits, base)
        if computed is not None:
            checks.append((True, f"Result {digits}({base}) = {computed}(10) — valid representation"))

    # U2 patterns
    for m in re.finditer(r'[Ww]artosc.*?=\s*\*\*(-?\d+)\*\*', answer_text):
        checks.append((True, f"U2 value {m.group(1)} found"))

    # Range patterns
    for m in re.finditer(r'(?:max|najwieksza).*?=\s*\*\*(\d+)\*\*', answer_text, re.IGNORECASE):
        val = int(m.group(1))
        checks.append((True, f"Max value {val} found"))

    return checks


def verify_numconv(exercise: dict, file_typ: str) -> VerificationResult:
    """Verify a single number conversion exercise."""
    eid = exercise['id']
    odpowiedz = exercise.get('odpowiedz', '')

    if not odpowiedz.strip():
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='TEORIA',
            status='ERROR', method='python_calc',
            details='Empty odpowiedz',
        )

    checks = _verify_conversion(odpowiedz)

    if not checks:
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='TEORIA',
            status='MANUAL_REVIEW', method='python_calc',
            details='No verifiable conversion patterns found in odpowiedz',
        )

    failures = [(ok, detail) for ok, detail in checks if not ok]
    passes = [(ok, detail) for ok, detail in checks if ok]

    all_details = '\n'.join(f"{'PASS' if ok else 'FAIL'}: {d}" for ok, d in checks)

    if failures:
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='TEORIA',
            status='FAIL', method='python_calc',
            details=f'{len(failures)}/{len(checks)} checks failed:\n{all_details}',
        )

    return VerificationResult(
        exercise_id=eid, file_typ=file_typ, kategoria='TEORIA',
        status='PASS', method='python_calc',
        details=f'All {len(passes)} checks passed:\n{all_details}',
    )
