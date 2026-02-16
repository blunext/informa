#\!/bin/bash
set -e
cd "$(dirname "$0")"
CGO_ENABLED=0 go build -o matura .
CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -o matura-linux .
CGO_ENABLED=0 GOOS=windows GOARCH=amd64 go build -o matura.exe .
./matura data import --source ../
echo "Build OK: matura ($(uname -m)), matura-linux (linux-amd64), matura.exe (win-amd64), matura.db"
