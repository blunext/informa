#!/usr/bin/env bash
# Download missing CKE exam papers from arkusze.pl
# 12 sessions: 7 czerwiec, 3 próbna, 1 przykładowy, 1 main 2020
set -euo pipefail

BASE="https://arkusze.pl/maturalne"
DIR="$(cd "$(dirname "$0")" && pwd)"

download() {
    local dir="$1" url="$2" filename="$3"
    mkdir -p "$DIR/$dir"
    if [[ -f "$DIR/$dir/$filename" ]]; then
        echo "  SKIP $dir/$filename (exists)"
        return
    fi
    echo "  GET  $dir/$filename"
    if curl -fsSL -o "$DIR/$dir/$filename" "$url"; then
        echo "  OK   $dir/$filename"
    else
        echo "  FAIL $dir/$filename (HTTP error)"
        rm -f "$DIR/$dir/$filename"
    fi
}

download_exam() {
    local dir="$1" slug="$2"
    echo "=== $dir ==="
    download "$dir" "$BASE/$slug.pdf" "arkusz.pdf"
    download "$dir" "$BASE/$slug-2.pdf" "arkusz_cz2.pdf"
    download "$dir" "$BASE/$slug-odpowiedzi.pdf" "odpowiedzi.pdf"
    download "$dir" "$BASE/$slug-odpowiedzi-2.pdf" "odpowiedzi_cz2.pdf"
    download "$dir" "$BASE/$slug-zalaczniki.zip" "zalaczniki.zip"
}

echo "Downloading 12 missing CKE exam sessions..."
echo ""

# 1. Main 2020 exam (termin główny, held in June due to COVID)
download_exam "2020_maj" "informatyka-2020-czerwiec-matura-rozszerzona"

# 2-7. Czerwiec (dodatkowa/resit) sessions
download_exam "2015_czerwiec" "informatyka-2015-czerwiec-matura-rozszerzona"
download_exam "2016_czerwiec" "informatyka-2016-czerwiec-matura-rozszerzona"
download_exam "2017_czerwiec" "informatyka-2017-czerwiec-matura-rozszerzona"
download_exam "2018_czerwiec" "informatyka-2018-czerwiec-matura-rozszerzona"
download_exam "2019_czerwiec" "informatyka-2019-czerwiec-matura-rozszerzona"
download_exam "2020_czerwiec" "informatyka-2020-lipiec-matura-rozszerzona"
download_exam "2021_czerwiec" "informatyka-2021-czerwiec-matura-rozszerzona"

# 8-10. Próbna (practice) exams
download_exam "2014_probna_grudzien" "informatyka-2014-grudzien-probna-rozszerzona"
download_exam "2020_probna_kwiecien" "informatyka-2020-kwiecien-probna-rozszerzona"
download_exam "2021_probna_marzec" "informatyka-2021-marzec-probna-rozszerzona"

# 11. Przykładowy 2015
download_exam "2015_przykladowy" "informatyka-2015-przykladowy-arkusz-cke-rozszerzona"

echo ""
echo "Done. Unzipping attachments..."
for z in "$DIR"/20*_*/zalaczniki.zip; do
    [ -f "$z" ] || continue
    d="$(dirname "$z")"
    echo "  UNZIP $z"
    unzip -qo "$z" -d "$d" 2>/dev/null || echo "  WARN  Failed to unzip $z"
done

echo ""
echo "Summary:"
for d in "$DIR"/20{14_probna,15_czerwiec,15_przykladowy,16_czerwiec,17_czerwiec,18_czerwiec,19_czerwiec,20_maj,20_czerwiec,20_probna,21_czerwiec,21_probna}*/; do
    [ -d "$d" ] || continue
    count=$(ls "$d"/*.pdf 2>/dev/null | wc -l | tr -d ' ')
    echo "  $(basename "$d"): $count PDFs"
done
