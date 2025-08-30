#!/bin/bash
FILE="PHPUnit_GoScan"

# Build for Windows
GOOS=windows GOARCH=amd64 go build -o ${FILE}_amd64_windows.exe ${FILE}.go

# Build for Linux
GOOS=linux GOARCH=amd64 go build -o ${FILE}_amd64_linux ${FILE}.go

echo "Build complete:"
echo "  - ${FILE}_amd64_windows.exe"
echo "  - ${FILE}_amd64_linux"
