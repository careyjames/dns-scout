#!/bin/bash

# Make sure to install Go 1.21 before running this script
# or update the path to the Go binary accordingly.

# Step 1: Compile the Go code
echo "Compiling Go code..."
go build -v -o ./bin/dns-scout main.go

GOOS=darwin GOARCH=amd64 go build -v -o ./dns-scout-macos-amd64-intel-v5.8/dns-scout main.go
# produces a binary for macOS running on Intel x86_64 architecture (Intel Macs). It does not produce a binary for the newer Apple Silicon Macs, which use the ARM64 architecture.

GOOS=darwin GOARCH=arm64 go build -v -o ./dns-scout-macos-arm64-silicon-v5.8/dns-scout main.go
# This will produce a binary (dns-scout-macos-arm64) that runs on macOS systems with ARM64 architecture (Apple Silicon).

GOOS=linux GOARCH=amd64 go build -v -o ./dns-scout-linux-amd64-ubuntu-kali-v5.8/dns-scout main.go
# This will generate a binary (dns-scout-linux-amd64) that is suitable for most Kali and Ubuntu installations on AMD64/x86_64 hardware.

GOOS=linux GOARCH=386 go build -v -o ./dns-scout-linux-386-v5.8/dns-scout main.go
# If you want to support older 32-bit machines or other architectures, you'll need to specify different GOARCH values. For example, for 32-bit x86:

tar czvf dns-scout-macos-amd64-intel-v5.8.tar.gz --transform 's,^./dns-scout-macos-amd64-intel-v5.8/dns-scout,dns-scout,' ./dns-scout-macos-amd64-intel-v5.8/dns-scout ./README.md ./setup-api-token.sh

tar czvf dns-scout-macos-arm64-silicon-v5.8.tar.gz --transform 's,^./dns-scout-macos-arm64-silicon-v5.8/dns-scout,dns-scout,' ./dns-scout-macos-arm64-silicon-v5.8/dns-scout ./README.md ./setup-api-token.sh

tar czvf dns-scout-linux-amd64-ubuntu-kali-v5.8.tar.gz --transform 's,^./dns-scout-linux-amd64-ubuntu-kali-v5.8/dns-scout,dns-scout,' ./dns-scout-linux-amd64-ubuntu-kali-v5.8/dns-scout ./README.md ./setup-api-token.sh

tar czvf dns-scout-linux-386-v5.8.tar.gz --transform 's,^./dns-scout-linux-386-v5.8/dns-scout,dns-scout,' ./dns-scout-linux-386-v5.8/dns-scout ./README.md ./setup-api-token.sh

shasum -a 256 ./dns-scout-macos-amd64-intel-v5.8/dns-scout ./dns-scout-macos-arm64-silicon-v5.8/dns-scout ./dns-scout-linux-amd64-ubuntu-kali-v5.8/dns-scout ./dns-scout-linux-386-v5.8/dns-scout

echo "Build complete."