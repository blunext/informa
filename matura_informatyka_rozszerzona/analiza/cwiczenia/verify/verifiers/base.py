"""Base utilities for exercise verification."""
import re
from dataclasses import dataclass, field


@dataclass
class VerificationResult:
    exercise_id: str          # "20.1"
    file_typ: str             # "20_sql_group_by"
    kategoria: str            # SQL / IMPLEMENTACJA / TEORIA / ARKUSZ
    status: str               # PASS / FAIL / MANUAL_REVIEW / ERROR
    method: str               # "sqlite_exec" / "cpp_compile_run" / "python_calc" / "manual"
    details: str              # Human-readable description
    expected: str | None = None
    actual: str | None = None
    error: str | None = None


def extract_code_blocks(text: str, lang: str = "") -> list[str]:
    """Extract fenced code blocks for a given language from markdown text.

    If lang is empty, extracts all code blocks regardless of language tag.
    Returns list of code block contents (without fences).
    """
    if lang:
        pattern = rf'```{re.escape(lang)}\s*\n(.*?)```'
    else:
        pattern = r'```(?:\w*)\s*\n(.*?)```'
    return re.findall(pattern, text, re.DOTALL)


def parse_markdown_tables(text: str) -> list[list[dict]]:
    """Parse all markdown tables from text.

    Returns a list of tables, where each table is a list of row dicts
    with column headers as keys.
    """
    tables = []
    lines = text.split('\n')
    i = 0
    while i < len(lines):
        line = lines[i].strip()
        # Detect table header row
        if line.startswith('|') and line.endswith('|'):
            header_line = line
            # Check for separator row
            if i + 1 < len(lines):
                sep = lines[i + 1].strip()
                if sep.startswith('|') and re.match(r'^\|[\s\-:|]+\|$', sep):
                    # Parse headers
                    headers = [h.strip() for h in header_line.strip('|').split('|')]
                    headers = [h for h in headers if h]  # remove empty
                    rows = []
                    j = i + 2
                    while j < len(lines):
                        row_line = lines[j].strip()
                        if not row_line.startswith('|') or not row_line.endswith('|'):
                            break
                        cells = [c.strip() for c in row_line.strip('|').split('|')]
                        cells = [c for c in cells if c is not None]
                        if len(cells) == len(headers):
                            row = {}
                            for k, h in enumerate(headers):
                                row[h] = cells[k] if k < len(cells) else ""
                            rows.append(row)
                        j += 1
                    if rows:
                        tables.append(rows)
                    i = j
                    continue
        i += 1
    return tables


def _try_numeric(val: str) -> int | float | str:
    """Try to parse a string as int or float, return original string if neither."""
    val = val.strip()
    # Remove bold markers
    val = re.sub(r'\*\*([^*]+)\*\*', r'\1', val)
    val = val.strip()
    try:
        return int(val)
    except ValueError:
        pass
    try:
        return float(val)
    except ValueError:
        pass
    return val


def values_match(expected, actual, tol: float = 0.01) -> bool:
    """Compare two values with numeric tolerance.

    Both values can be strings that look like numbers.
    """
    ev = _try_numeric(str(expected))
    av = _try_numeric(str(actual))

    if isinstance(ev, (int, float)) and isinstance(av, (int, float)):
        if isinstance(ev, float) or isinstance(av, float):
            return abs(float(ev) - float(av)) <= tol
        return ev == av

    # String comparison (case-insensitive, whitespace-normalized)
    return str(ev).strip().lower() == str(av).strip().lower()


def _rows_to_sort_key(row: dict) -> list:
    """Convert row dict to a sortable key."""
    result = []
    for v in row.values():
        n = _try_numeric(str(v))
        if isinstance(n, (int, float)):
            result.append((0, float(n), str(v)))
        else:
            result.append((1, 0.0, str(v).strip().lower()))
    return result


def _compare_rows(expected: list[dict], actual: list[dict],
                  float_tol: float) -> tuple[bool, str]:
    """Row-by-row comparison helper."""
    exp_cols = list(expected[0].keys())
    for row_idx, (exp_row, act_row) in enumerate(zip(expected, actual)):
        exp_vals = list(exp_row.values())
        act_vals = list(act_row.values())
        for col_idx in range(len(exp_vals)):
            if not values_match(exp_vals[col_idx], act_vals[col_idx], float_tol):
                return False, (
                    f"Row {row_idx}, col {col_idx} ({exp_cols[col_idx]}): "
                    f"expected '{exp_vals[col_idx]}', got '{act_vals[col_idx]}'"
                )
    return True, f"All {len(expected)} rows match"


def compare_tables(expected: list[dict], actual: list[dict],
                   float_tol: float = 0.01) -> tuple[bool, str]:
    """Compare two tables (lists of row dicts).

    First tries exact row-by-row match. If that fails, tries sorted comparison
    (handles ORDER BY ties where SQLite may return rows in different order).

    Returns (match, details_message).
    """
    if len(expected) != len(actual):
        return False, f"Row count mismatch: expected {len(expected)}, got {len(actual)}"

    if not expected:
        return True, "Both tables empty"

    exp_cols = list(expected[0].keys())
    act_cols = list(actual[0].keys())

    if len(exp_cols) != len(act_cols):
        return False, f"Column count mismatch: expected {len(exp_cols)}, got {len(act_cols)}"

    # Try exact row-by-row match first
    match, msg = _compare_rows(expected, actual, float_tol)
    if match:
        return True, msg

    # Fallback: sort both tables and compare (handles ORDER BY ties)
    try:
        exp_sorted = sorted(expected, key=_rows_to_sort_key)
        act_sorted = sorted(actual, key=_rows_to_sort_key)
        match2, msg2 = _compare_rows(exp_sorted, act_sorted, float_tol)
        if match2:
            return True, f"All {len(expected)} rows match (after sorting — ORDER BY tie)"
    except Exception:
        pass

    # Return the original (unsorted) mismatch message
    return False, msg


def normalize_whitespace(text: str) -> str:
    """Normalize whitespace for output comparison."""
    return ' '.join(text.split()).strip()
