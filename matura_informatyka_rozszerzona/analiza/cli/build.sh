#\!/bin/bash
set -e
cd "$(dirname "$0")"
CGO_ENABLED=0 go build -o matura .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o matura.exe .
./matura data import --source ../
echo "Build OK: matura ($(uname -m)), matura.exe (win-amd64), matura.db"
