"""Manual tagger — marks exercises that require human review."""
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


def tag_manual(exercise: dict, file_typ: str) -> VerificationResult:
    """Tag exercise as requiring manual review."""
    eid = exercise['id']
    prefix = file_typ.split('_')[0]

    kategoria = 'TEORIA'
    if int(prefix) >= 15 and int(prefix) <= 19:
        kategoria = 'ARKUSZ'

    hint = REVIEW_HINTS.get(prefix, 'Wymaga recznej weryfikacji — brak metody automatycznej.')

    return VerificationResult(
        exercise_id=eid,
        file_typ=file_typ,
        kategoria=kategoria,
        status='MANUAL_REVIEW',
        method='manual',
        details=hint,
    )
