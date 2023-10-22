#!/bin/bash

# Make sure to install Go 1.21 before running this script
# or update the path to the Go binary accordingly.

# Step 1: Compile the Go code
echo "Compiling Go code..."
go build -v -o ./bin/dns-scout

GOOS=darwin GOARCH=amd64 go build -v -o ./dns-scout-macos-amd64-intel-v6.0/dns-scout
# produces a binary for macOS running on Intel x86_64 architecture (Intel Macs). It does not produce a binary for the newer Apple Silicon Macs, which use the ARM64 architecture.

GOOS=darwin GOARCH=arm64 go build -v -o ./dns-scout-macos-arm64-silicon-v6.0/dns-scout
# This will produce a binary (dns-scout-macos-arm64) that runs on macOS systems with ARM64 architecture (Apple Silicon).

GOOS=linux GOARCH=amd64 go build -v -o ./dns-scout-linux-amd64-ubuntu-kali-v6.0/dns-scout
# This will generate a binary (dns-scout-linux-amd64) that is suitable for most Kali and Ubuntu installations on AMD64/x86_64 hardware.

GOOS=linux GOARCH=arm64 go build -v -o ./dns-scout-linux-arm64-raspberry-pi-v6.0/dns-scout
# Raspberry Pi 64-bit ARM

GOOS=linux GOARCH=386 go build -v -o ./dns-scout-linux-386-v6.0/dns-scout
# If you want to support older 32-bit machines or other architectures, you'll need to specify different GOARCH values. For example, for 32-bit x86:

# Create Debian packages for Linux builds
for arch in amd64 arm64 386; do
  deb_folder="./dns-scout-linux-${arch}-v6.0-deb"
  deb_name="dns-scout-linux-${arch}-v6.0-deb.deb"  # Initialize with a default name

  mkdir -p "${deb_folder}/usr/local/bin"
  mkdir -p "${deb_folder}/usr/share/doc/dns-scout"
  mkdir -p "${deb_folder}/DEBIAN"

  cp "./bin/dns-scout" "${deb_folder}/usr/local/bin/"
  cp README.md "${deb_folder}/usr/share/doc/dns-scout/"
  cp setup-api-token.sh "${deb_folder}/usr/share/doc/dns-scout/"

  echo "Package: dns-scout
Version: 6.0
Section: net
Priority: optional
Architecture: ${arch}
Essential: no
Installed-Size: $(du -s "${deb_folder}" | cut -f1)
Maintainer: Carey Balboa
Description: DNS Scout for Linux/MacOS
 DNS Scout pulls and displays DNS records in a color-coded console output.
 It stands out by filtering out non-essential information, presenting users
 with a cleaner, more focused view of the DNS data. The tool is optimized
 for clarity and relevance, making it ideal for easy DNS reconnaissance
 and troubleshooting." > "${deb_folder}/DEBIAN/control"

  # Rename the Debian packages for specific distributions
  case "${arch}" in
    amd64)
      deb_name="dns-scout-linux-amd64-ubuntu-kali-v6.0.deb"
      ;;
    arm64)
      deb_name="dns-scout-linux-arm64-raspberry-pi-v6.0.deb"
      ;;
    386)
      deb_name="dns-scout-linux-386-v6.0.deb"
      ;;
  esac

  dpkg-deb --build "${deb_folder}" "${deb_name}"
done

tar czvf dns-scout-macos-amd64-intel-v6.0.tar.gz --transform 's,^./dns-scout-macos-amd64-intel-v6.0/dns-scout,dns-scout,' ./dns-scout-macos-amd64-intel-v6.0/dns-scout ./README.md ./setup-api-token.sh

tar czvf dns-scout-macos-arm64-silicon-v6.0.tar.gz --transform 's,^./dns-scout-macos-arm64-silicon-v6.0/dns-scout,dns-scout,' ./dns-scout-macos-arm64-silicon-v6.0/dns-scout ./README.md ./setup-api-token.sh

tar czvf dns-scout-linux-amd64-ubuntu-kali-v6.0.tar.gz --transform 's,^./dns-scout-linux-amd64-ubuntu-kali-v6.0/dns-scout,dns-scout,' ./dns-scout-linux-amd64-ubuntu-kali-v6.0/dns-scout ./README.md ./setup-api-token.sh

tar czvf dns-scout-linux-arm64-raspberry-pi-v6.0.tar.gz --transform 's,^./dns-scout-linux-arm64-raspberry-pi-v6.0/dns-scout,dns-scout,' ./dns-scout-linux-arm64-raspberry-pi-v6.0/dns-scout ./README.md ./setup-api-token.sh

tar czvf dns-scout-linux-386-v6.0.tar.gz --transform 's,^./dns-scout-linux-386-v6.0/dns-scout,dns-scout,' ./dns-scout-linux-386-v6.0/dns-scout ./README.md ./setup-api-token.sh

shasum -a 256 ./dns-scout-macos-amd64-intel-v6.0/dns-scout ./dns-scout-macos-arm64-silicon-v6.0/dns-scout ./dns-scout-linux-amd64-ubuntu-kali-v6.0/dns-scout ./dns-scout-linux-386-v6.0/dns-scout ./dns-scout-linux-arm64-raspberry-pi-v6.0/dns-scout ./dns-scout-linux-amd64-ubuntu-kali-v6.0.deb ./dns-scout-linux-arm64-raspberry-pi-v6.0.deb ./dns-scout-linux-386-v6.0.deb

echo "Build complete."