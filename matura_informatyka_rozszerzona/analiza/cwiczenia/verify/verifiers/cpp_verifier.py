"""C++ exercise verifier — compiles and runs C++ code, compares output."""
import os
import re
import shutil
import subprocess
import tempfile
from .base import VerificationResult, extract_code_blocks, normalize_whitespace


def _extract_cpp_code(odpowiedz: str) -> str | None:
    """Extract C++ code from odpowiedz."""
    blocks = extract_code_blocks(odpowiedz, 'cpp')
    if blocks:
        # Return the largest block (most likely the full solution)
        return max(blocks, key=len).strip()
    # Fallback: try c++ tag
    blocks = extract_code_blocks(odpowiedz, 'c++')
    if blocks:
        return max(blocks, key=len).strip()
    return None


def _extract_input_files(tresc: str) -> dict[str, str]:
    """Extract input file data from tresc.

    Looks for patterns like:
      **Dane** (`filename.txt`):
      ```
      data here
      ```
    or:
      **Dane** (`filename.txt`):
      ```
      data
      ```

    Returns dict: filename -> content
    """
    files = {}
    # Pattern: **Dane** (`filename.txt`):  or  **Dane** (`filename`):
    # followed by a code block
    pattern = r'\*\*Dane\*\*\s*\(`([^`]+)`\)\s*:\s*\n```[^\n]*\n(.*?)```'
    for m in re.finditer(pattern, tresc, re.DOTALL):
        fname = m.group(1).strip()
        content = m.group(2)
        files[fname] = content

    if not files:
        # Try alternative pattern without parentheses
        pattern = r'\*\*Dane\*\*\s*\(`([^`]+)`\)\s*:\s*\n```\n(.*?)```'
        for m in re.finditer(pattern, tresc, re.DOTALL):
            fname = m.group(1).strip()
            content = m.group(2)
            files[fname] = content

    return files


def _extract_expected_output(tresc: str) -> str | None:
    """Extract expected output from tresc.

    Looks for pattern:
      **Oczekiwany wynik**:
      ```
      output here
      ```
    """
    pattern = r'\*\*Oczekiwany wynik\*\*\s*:\s*\n```[^\n]*\n(.*?)```'
    m = re.search(pattern, tresc, re.DOTALL)
    if m:
        return m.group(1).strip()
    return None


def _patch_file_paths(code: str, workdir: str, files: dict[str, str]) -> str:
    """Replace file path strings in C++ code with absolute paths."""
    for fname in files:
        # Match various patterns: ifstream plik("file"), open("file"), etc.
        escaped = re.escape(fname)
        # Replace quoted filename with absolute path
        abs_path = os.path.join(workdir, fname)
        code = re.sub(
            rf'("){escaped}(")',
            rf'\g<1>{abs_path}\g<2>',
            code
        )
        code = re.sub(
            rf"('){escaped}(')",
            rf"\g<1>{abs_path}\g<2>",
            code
        )
    return code


def _compare_output(expected: str, actual: str) -> tuple[bool, str]:
    """Compare expected and actual outputs with tolerances.

    Normalizes whitespace and allows minor numeric differences.
    """
    exp_lines = [l.strip() for l in expected.strip().split('\n') if l.strip()]
    act_lines = [l.strip() for l in actual.strip().split('\n') if l.strip()]

    if len(exp_lines) != len(act_lines):
        return False, (
            f"Line count mismatch: expected {len(exp_lines)}, got {len(act_lines)}\n"
            f"Expected:\n{expected[:500]}\n\nActual:\n{actual[:500]}"
        )

    for i, (el, al) in enumerate(zip(exp_lines, act_lines)):
        if normalize_whitespace(el) != normalize_whitespace(al):
            # Try numeric-tolerant comparison for lines with numbers
            if _lines_match_numeric(el, al):
                continue
            return False, (
                f"Line {i+1} mismatch:\n"
                f"  expected: {el!r}\n"
                f"  actual:   {al!r}"
            )

    return True, f"All {len(exp_lines)} lines match"


def _lines_match_numeric(expected: str, actual: str) -> bool:
    """Compare two lines allowing for minor numeric formatting differences."""
    # Extract all numbers from both lines
    exp_nums = re.findall(r'-?\d+\.?\d*', expected)
    act_nums = re.findall(r'-?\d+\.?\d*', actual)

    if len(exp_nums) != len(act_nums):
        return False

    # Remove numbers and compare text parts
    exp_text = re.sub(r'-?\d+\.?\d*', '#', expected)
    act_text = re.sub(r'-?\d+\.?\d*', '#', actual)

    if normalize_whitespace(exp_text) != normalize_whitespace(act_text):
        return False

    # Compare numeric values with tolerance
    for en, an in zip(exp_nums, act_nums):
        try:
            ef = float(en)
            af = float(an)
            if abs(ef - af) > 0.01:
                return False
        except ValueError:
            if en != an:
                return False

    return True


def verify_cpp(exercise: dict, file_typ: str, tmp_base: str) -> VerificationResult:
    """Verify a single C++ exercise."""
    eid = exercise['id']
    tresc = exercise.get('tresc', '')
    odpowiedz = exercise.get('odpowiedz', '')

    # 1. Extract C++ code
    code = _extract_cpp_code(odpowiedz)
    if not code:
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='IMPLEMENTACJA',
            status='ERROR', method='cpp_compile_run',
            details='No C++ code block found in odpowiedz',
        )

    # 2. Extract input files and expected output
    input_files = _extract_input_files(tresc)
    expected_output = _extract_expected_output(tresc)

    # 3. Set up working directory
    workdir = os.path.join(tmp_base, f'ex_{eid.replace(".", "_")}')
    os.makedirs(workdir, exist_ok=True)

    try:
        # Write input files
        for fname, content in input_files.items():
            filepath = os.path.join(workdir, fname)
            with open(filepath, 'w', encoding='utf-8') as f:
                f.write(content)

        # Patch file paths in code
        patched_code = _patch_file_paths(code, workdir, input_files)

        # Write source
        source_path = os.path.join(workdir, 'source.cpp')
        with open(source_path, 'w', encoding='utf-8') as f:
            f.write(patched_code)

        # 4. Compile
        prog_path = os.path.join(workdir, 'prog')
        compile_result = subprocess.run(
            ['g++', '-std=c++17', '-o', prog_path, source_path],
            capture_output=True, text=True, timeout=10,
            cwd=workdir,
        )

        if compile_result.returncode != 0:
            return VerificationResult(
                exercise_id=eid, file_typ=file_typ, kategoria='IMPLEMENTACJA',
                status='ERROR', method='cpp_compile_run',
                details=f'Compilation failed',
                error=compile_result.stderr[:500],
            )

        # 5. Run
        run_result = subprocess.run(
            [prog_path],
            capture_output=True, text=True, timeout=5,
            cwd=workdir,
        )

        if run_result.returncode != 0:
            return VerificationResult(
                exercise_id=eid, file_typ=file_typ, kategoria='IMPLEMENTACJA',
                status='ERROR', method='cpp_compile_run',
                details=f'Runtime error (exit code {run_result.returncode})',
                error=run_result.stderr[:500],
            )

        actual_output = run_result.stdout.strip()

        # 6. Compare
        if expected_output is None:
            return VerificationResult(
                exercise_id=eid, file_typ=file_typ, kategoria='IMPLEMENTACJA',
                status='PASS', method='cpp_compile_run',
                details=f'Compiled and ran OK (no expected output to compare)',
                actual=actual_output[:500],
            )

        match, detail_msg = _compare_output(expected_output, actual_output)

        if match:
            return VerificationResult(
                exercise_id=eid, file_typ=file_typ, kategoria='IMPLEMENTACJA',
                status='PASS', method='cpp_compile_run',
                details=f'Output matches expected',
                expected=expected_output[:500], actual=actual_output[:500],
            )
        else:
            return VerificationResult(
                exercise_id=eid, file_typ=file_typ, kategoria='IMPLEMENTACJA',
                status='FAIL', method='cpp_compile_run',
                details=f'Output mismatch: {detail_msg}',
                expected=expected_output[:500], actual=actual_output[:500],
            )

    except subprocess.TimeoutExpired:
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='IMPLEMENTACJA',
            status='ERROR', method='cpp_compile_run',
            details='Timeout expired',
            error='Process timed out',
        )
    except FileNotFoundError:
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='IMPLEMENTACJA',
            status='ERROR', method='cpp_compile_run',
            details='g++ not found — is it installed?',
            error='g++ not found',
        )
    except Exception as e:
        return VerificationResult(
            exercise_id=eid, file_typ=file_typ, kategoria='IMPLEMENTACJA',
            status='ERROR', method='cpp_compile_run',
            details=f'Unexpected error: {e}',
            error=str(e),
        )
