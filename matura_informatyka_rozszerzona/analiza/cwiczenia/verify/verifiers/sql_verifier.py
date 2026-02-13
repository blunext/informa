"""SQL exercise verifier — executes queries in SQLite and compares results."""
import re
import sqlite3
from .base import (
    VerificationResult, extract_code_blocks, parse_markdown_tables,
    compare_tables, values_match
)


def _infer_type(values: list[str]) -> str:
    """Infer SQL column type from sample values."""
    for v in values:
        v = v.strip()
        if not v or v.upper() == 'NULL':
            continue
        try:
            int(v)
            continue
        except ValueError:
            pass
        try:
            float(v)
            return "REAL"
        except ValueError:
            return "TEXT"
    # All ints or empty
    all_int = True
    for v in values:
        v = v.strip()
        if not v or v.upper() == 'NULL':
            continue
        try:
            int(v)
        except ValueError:
            all_int = False
            break
    return "INTEGER" if all_int else "TEXT"


def _parse_tables_from_tresc(tresc: str) -> dict[str, list[dict]]:
    """Parse table definitions from exercise tresc.

    Looks for patterns like:
      Tabela **Nazwa**:  (or Tabela **Nazwa** zawiera...)
    followed by a markdown table.

    Returns dict: table_name -> list of row dicts.
    """
    result = {}
    lines = tresc.split('\n')
    i = 0
    pending_name = None

    while i < len(lines):
        line = lines[i].strip()

        # Look for table name patterns
        m = re.search(r'[Tt]abela\s+\*\*(\w+)\*\*', line)
        if m:
            pending_name = m.group(1)
            i += 1
            continue

        # If we have a pending table name and find a markdown table header
        if pending_name and line.startswith('|') and line.endswith('|'):
            if i + 1 < len(lines):
                sep = lines[i + 1].strip()
                if sep.startswith('|') and re.match(r'^\|[\s\-:|]+\|$', sep):
                    headers = [h.strip() for h in line.strip('|').split('|')]
                    headers = [h for h in headers if h]
                    rows = []
                    j = i + 2
                    while j < len(lines):
                        row_line = lines[j].strip()
                        if not row_line.startswith('|') or not row_line.endswith('|'):
                            break
                        cells = [c.strip() for c in row_line.strip('|').split('|')]
                        cells = [c for c in cells if c is not None]
                        if len(cells) == len(headers):
                            row = {headers[k]: cells[k] for k in range(len(headers))}
                            rows.append(row)
                        j += 1
                    if rows:
                        result[pending_name] = rows
                    i = j
                    pending_name = None
                    continue

        # Reset pending name if we hit non-table content
        if pending_name and line and not line.startswith('|'):
            # Keep pending_name — the table might be after some text
            pass

        i += 1

    return result


def _create_table_sql(name: str, rows: list[dict]) -> list[str]:
    """Generate CREATE TABLE and INSERT statements."""
    if not rows:
        return []

    headers = list(rows[0].keys())
    # Collect sample values per column
    col_values = {h: [] for h in headers}
    for row in rows:
        for h in headers:
            col_values[h].append(row.get(h, ''))

    # Infer types
    col_types = {h: _infer_type(col_values[h]) for h in headers}

    stmts = []
    cols_def = ', '.join(f'"{h}" {col_types[h]}' for h in headers)
    stmts.append(f'CREATE TABLE "{name}" ({cols_def});')

    for row in rows:
        vals = []
        for h in headers:
            v = row[h].strip()
            if v.upper() == 'NULL' or v == '':
                vals.append('NULL')
            elif col_types[h] == 'TEXT':
                vals.append("'" + v.replace("'", "''") + "'")
            else:
                # Try numeric, fallback to text
                try:
                    if col_types[h] == 'INTEGER':
                        int(v)
                        vals.append(v)
                    else:
                        float(v)
                        vals.append(v)
                except ValueError:
                    vals.append("'" + v.replace("'", "''") + "'")
        cols_list = ', '.join(f'"{h}"' for h in headers)
        vals_list = ', '.join(vals)
        stmts.append(f'INSERT INTO "{name}" ({cols_list}) VALUES ({vals_list});')

    return stmts


def _is_empty_placeholder(table: list[dict]) -> bool:
    """Check if table is a placeholder for empty results like '(pusty wynik)'."""
    if len(table) != 1:
        return False
    row = table[0]
    for v in row.values():
        v_lower = str(v).strip().lower()
        if 'pusty' in v_lower or 'brak' in v_lower or v_lower.startswith('('):
            return True
    return False


def _find_expected_tables(odpowiedz: str) -> list[list[dict]]:
    """Find all expected result tables in odpowiedz.

    Filters out tables with checkmarks and empty placeholders.
    Returns list of tables.
    """
    tables = parse_markdown_tables(odpowiedz)
    if not tables:
        return []

    valid = []
    for t in tables:
        # Skip tables with check/cross marks
        has_marks = False
        for row in t:
            for v in row.values():
                if any(c in str(v) for c in ['✓', '✗', '✔', '✘']):
                    has_marks = True
                    break
            if has_marks:
                break
        if has_marks:
            continue

        # Skip empty placeholder tables
        if _is_empty_placeholder(t):
            continue

        valid.append(t)

    return valid


def _find_expected_table(odpowiedz: str) -> list[dict] | None:
    """Find the expected result table in odpowiedz (last valid table)."""
    tables = _find_expected_tables(odpowiedz)
    return tables[-1] if tables else None


def _extract_sql(odpowiedz: str) -> str | None:
    """Extract SQL query from odpowiedz."""
    blocks = extract_code_blocks(odpowiedz, 'sql')
    if blocks:
        return blocks[0].strip()
    # Fallback: try without language tag
    blocks = extract_code_blocks(odpowiedz, '')
    for b in blocks:
        b_upper = b.strip().upper()
        if b_upper.startswith(('SELECT', 'WITH', 'INSERT', 'UPDATE', 'DELETE')):
            return b.strip()
    return None


def verify_sql(exercise: dict, file_typ: str) -> VerificationResult:
    """Verify a single SQL exercise."""
    eid = exercise['id']
    tresc = exercise.get('tresc', '')
    odpowiedz = exercise.get('odpowiedz', '')

    # 1. Parse tables from tresc
    tables = _parse_tables_from_tresc(tresc)
    if not tables:
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='SQL',
            status='ERROR', method='sqlite_exec',
            details='No tables found in tresc',
        )

    # 2. Extract SQL from odpowiedz
    sql = _extract_sql(odpowiedz)
    if not sql:
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='SQL',
            status='ERROR', method='sqlite_exec',
            details='No SQL code block found in odpowiedz',
        )

    # 3. Extract expected result tables
    expected_tables = _find_expected_tables(odpowiedz)

    # 4. Execute in SQLite
    try:
        conn = sqlite3.connect(':memory:')
        cur = conn.cursor()

        # Create and populate tables
        for tname, trows in tables.items():
            for stmt in _create_table_sql(tname, trows):
                cur.execute(stmt)

        # Handle multi-query (split by ;)
        queries = [q.strip() for q in sql.split(';') if q.strip()]
        # Collect results from ALL SELECT queries
        all_results = []  # list of (rows, col_names)
        for q in queries:
            q_stripped = re.sub(r'--[^\n]*\n?', '', q).strip()
            if not q_stripped:
                continue
            cur.execute(q)
            if q_stripped.upper().startswith(('SELECT', 'WITH')):
                rows = cur.fetchall()
                col_names = [desc[0] for desc in cur.description]
                all_results.append((rows, col_names))

        conn.close()

        if not all_results:
            return VerificationResult(
                exercise_id=eid, file_typ=file_typ, kategoria='SQL',
                status='ERROR', method='sqlite_exec',
                details='No SELECT query found or no results returned',
            )

        def _rows_to_table(rows, col_names):
            table = []
            for row in rows:
                row_dict = {}
                for i, col in enumerate(col_names):
                    val = row[i]
                    if val is None:
                        row_dict[col] = 'NULL'
                    elif isinstance(val, float) and val == int(val):
                        row_dict[col] = str(int(val))
                    else:
                        row_dict[col] = str(val)
                table.append(row_dict)
            return table

        # Multi-query: match each SELECT result with its corresponding expected table
        if len(all_results) > 1 and len(expected_tables) >= 1:
            all_pass = True
            all_details = []
            for qi, ((rows, col_names), exp_tbl) in enumerate(
                zip(all_results, expected_tables), 1
            ):
                actual_table = _rows_to_table(rows, col_names)
                match, detail_msg = compare_tables(exp_tbl, actual_table)
                all_details.append(f"Query {qi}: {detail_msg}")
                if not match:
                    all_pass = False

            details_str = '; '.join(all_details)
            if all_pass:
                return VerificationResult(
                    exercise_id=eid, file_typ=file_typ, kategoria='SQL',
                    status='PASS', method='sqlite_exec',
                    details=f'All {len(all_results)} queries match ({details_str})',
                )
            else:
                return VerificationResult(
                    exercise_id=eid, file_typ=file_typ, kategoria='SQL',
                    status='FAIL', method='sqlite_exec',
                    details=f'Multi-query mismatch: {details_str}',
                )

        # Single query or single expected table: use last result vs last expected
        results, col_names = all_results[-1]
        actual_table = _rows_to_table(results, col_names)

        # 5. Compare
        expected_table = expected_tables[-1] if expected_tables else None

        if expected_table is None:
            actual_str = _table_to_str(actual_table)
            return VerificationResult(
                exercise_id=eid, file_typ=file_typ, kategoria='SQL',
                status='PASS', method='sqlite_exec',
                details=f'SQL executed OK, {len(results)} rows returned (no expected table to compare)',
                actual=actual_str,
            )

        match, detail_msg = compare_tables(expected_table, actual_table)
        actual_str = _table_to_str(actual_table)
        expected_str = _table_to_str(expected_table)

        if match:
            return VerificationResult(
                exercise_id=eid, file_typ=file_typ, kategoria='SQL',
                status='PASS', method='sqlite_exec',
                details=f'SQL result matches expected ({len(results)} rows)',
                expected=expected_str, actual=actual_str,
            )
        else:
            return VerificationResult(
                exercise_id=eid, file_typ=file_typ, kategoria='SQL',
                status='FAIL', method='sqlite_exec',
                details=f'SQL result mismatch: {detail_msg}',
                expected=expected_str, actual=actual_str,
            )

    except sqlite3.Error as e:
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='SQL',
            status='ERROR', method='sqlite_exec',
            details=f'SQLite error: {e}',
            error=str(e),
        )
    except Exception as e:
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='SQL',
            status='ERROR', method='sqlite_exec',
            details=f'Unexpected error: {e}',
            error=str(e),
        )


def _table_to_str(table: list[dict]) -> str:
    """Convert table to compact string for reporting."""
    if not table:
        return "(empty)"
    headers = list(table[0].keys())
    lines = [' | '.join(headers)]
    for row in table:
        lines.append(' | '.join(str(row.get(h, '')) for h in headers))
    return '\n'.join(lines)
