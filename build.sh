#!/bin/bash

# Make sure to install Go 1.21 before running this script
# or update the path to the Go binary accordingly.
# Define the GPG key ID
GPG_KEY_ID="EF536354988BF362947FC6FDBEB7932396E8FB23"
# Define the project root directory
project_root=$(pwd)
export VERSION=$(grep Version constant/constants.go  | awk -F '=' '{print $2}'| awk -F'"' '{print $2}')

# Run env subtitute
envsubst < dns-scout.metainfo.xml.tpl > dns-scout.metainfo.xml
envsubst < debian/files.tpl > debian/files
envsubst < debian/changelog.tpl > debian/changelog

# Move out binary files
mkdir -p "${project_root}/../binaries"
mv "${project_root}/bin/dns-scout" "${project_root}/../binaries/"

# Remove existing upstream tarball if it exists
[ -f "${project_root}/../dns-scout_${VERSION}.orig.tar.gz" ] && rm "${project_root}/../dns-scout_${VERSION}.orig.tar.gz"
rm -rf ./debian/binaries/**

# Create the upstream tarball and place it in the parent directory
echo "Creating upstream tarball..."
tar czvf "${project_root}/../dns-scout_${VERSION}.orig.tar.gz" --exclude='.git' --exclude='./bin/*' --exclude='./dns-scout-linux-*' -C "${project_root}" .

# Step 1: Compile the Go code
echo "=======Compiling Go code..."
go build -v -o "${project_root}/bin/dns-scout"

GOOS=darwin GOARCH=amd64 go build -v -o "${project_root}/dns-scout-macos-amd64-intel-v${VERSION}/dns-scout"
# produces a binary for macOS running on Intel x86_64 architecture (Intel Macs). It does not produce a binary for the newer Apple Silicon Macs, which use the ARM64 architecture.

GOOS=darwin GOARCH=arm64 go build -v -o "${project_root}/dns-scout-macos-arm64-silicon-v${VERSION}/dns-scout"
# This will produce a binary (dns-scout-macos-arm64) that runs on macOS systems with ARM64 architecture (Apple Silicon).

GOOS=linux GOARCH=amd64 go build -v -o "${project_root}/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}/dns-scout"
# This will generate a binary (dns-scout-linux-amd64) that is suitable for most Kali and Ubuntu installations on AMD64/x86_64 hardware.

GOOS=linux GOARCH=arm64 go build -v -o "${project_root}/dns-scout-linux-arm64-raspberry-pi-v${VERSION}/dns-scout"
# Raspberry Pi 64-bit ARM

GOOS=linux GOARCH=386 go build -v -o "${project_root}/dns-scout-linux-386-v${VERSION}/dns-scout"
# If you want to support older 32-bit machines or other architectures, you'll need to specify different GOARCH values. For example, for 32-bit x86:

# Create Debian packages for Linux builds
for arch in amd64 arm64 386; do
  deb_folder="${project_root}/dns-scout-linux-${arch}-v${VERSION}-deb"
  deb_name="dns-scout-linux-${arch}-v${VERSION}-1debian1.deb"

  mkdir -p "${deb_folder}/usr/local/bin"
  mkdir -p "${deb_folder}/usr/share/doc/dns-scout"
  mkdir -p "${deb_folder}/DEBIAN"

  case "${arch}" in
  amd64)
    cp "${project_root}/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}/dns-scout" "${deb_folder}/usr/local/bin/"
    chmod 755 "${deb_folder}/usr/local/bin/dns-scout"
    ;;
  arm64)
    cp "${project_root}/dns-scout-linux-arm64-raspberry-pi-v${VERSION}/dns-scout" "${deb_folder}/usr/local/bin/"
    chmod 755 "${deb_folder}/usr/local/bin/dns-scout"
    ;;
  386)
    cp "${project_root}/dns-scout-linux-386-v${VERSION}/dns-scout" "${deb_folder}/usr/local/bin/"
    chmod 755 "${deb_folder}/usr/local/bin/dns-scout"
    ;;
esac

  cp README.md "${deb_folder}/usr/share/doc/dns-scout/"
  cp setup-api-token.sh "${deb_folder}/usr/share/doc/dns-scout/"

  echo "Package: dns-scout
Version: ${VERSION}
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
      deb_name="dns-scout-linux-amd64-ubuntu-kali-v${VERSION}-1debian1.deb"
      ;;
    arm64)
      deb_name="dns-scout-linux-arm64-raspberry-pi-v${VERSION}-1debian1.deb"
      ;;
    386)
      deb_name="dns-scout-linux-386-v${VERSION}-1debian1.deb"
      ;;
  esac

  dpkg-deb --build "${deb_folder}" "${deb_name}"
done

tar czvf "${project_root}/dns-scout-macos-amd64-intel-v${VERSION}.tar.gz" --transform 's,^./dns-scout-macos-amd64-intel-v${VERSION}/dns-scout,dns-scout,' ./dns-scout-macos-amd64-intel-v${VERSION}/dns-scout ./README.md ./setup-api-token.sh

tar czvf "${project_root}/dns-scout-macos-arm64-silicon-v${VERSION}.tar.gz" --transform 's,^./dns-scout-macos-arm64-silicon-v${VERSION}/dns-scout,dns-scout,' ./dns-scout-macos-arm64-silicon-v${VERSION}/dns-scout ./README.md ./setup-api-token.sh

tar czvf "${project_root}/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}.tar.gz" --transform 's,^./dns-scout-linux-amd64-ubuntu-kali-v${VERSION}/dns-scout,dns-scout,' ./dns-scout-linux-amd64-ubuntu-kali-v${VERSION}/dns-scout ./README.md ./setup-api-token.sh

tar czvf "${project_root}/dns-scout-linux-arm64-raspberry-pi-v${VERSION}.tar.gz" --transform 's,^./dns-scout-linux-arm64-raspberry-pi-v${VERSION}/dns-scout,dns-scout,' ./dns-scout-linux-arm64-raspberry-pi-v${VERSION}/dns-scout ./README.md ./setup-api-token.sh

tar czvf "${project_root}/dns-scout-linux-386-v${VERSION}.tar.gz" --transform 's,^./dns-scout-linux-386-v${VERSION}/dns-scout,dns-scout,' ./dns-scout-linux-386-v${VERSION}/dns-scout ./README.md ./setup-api-token.sh

shasum -a 256 "${project_root}/dns-scout-macos-amd64-intel-v${VERSION}/dns-scout" "${project_root}/dns-scout-macos-arm64-silicon-v${VERSION}/dns-scout" "${project_root}/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}/dns-scout" "${project_root}/dns-scout-linux-386-v${VERSION}/dns-scout" "${project_root}/dns-scout-linux-arm64-raspberry-pi-v${VERSION}/dns-scout" "${project_root}/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}-1debian1.deb" "${project_root}/dns-scout-linux-arm64-raspberry-pi-v${VERSION}-1debian1.deb" "${project_root}/dns-scout-linux-386-v${VERSION}-1debian1.deb"

echo "looking at folders..."
ls -lart | grep dns-scout

# Clean up the generated binaries and artifacts
echo "======== Cleaning up generated binaries and artifacts..."
rm -f "${project_root}/dns-scout-linux-*.tar.gz"
# Move generated binaries and packages to debian/binaries/
echo "======== Moving generated binaries and packages to debian/binaries/"

mv -f "${project_root}/dns-scout-linux-386-v${VERSION}-1debian1.deb" "${project_root}/../binaries/"
mv -f "${project_root}/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}-1debian1.deb" "${project_root}/../binaries/"
mv -f "${project_root}/dns-scout-linux-arm64-raspberry-pi-v${VERSION}-1debian1.deb" "${project_root}/../binaries/"


mv -f "${project_root}/dns-scout-linux-386-v${VERSION}.tar.gz" "${project_root}/../binaries/"
mv -f "${project_root}/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}.tar.gz" "${project_root}/../binaries/"
mv -f "${project_root}/dns-scout-linux-arm64-raspberry-pi-v${VERSION}.tar.gz" "${project_root}/../binaries/"
mv -f "${project_root}/dns-scout-macos-amd64-intel-v${VERSION}.tar.gz" "${project_root}/../binaries/"
mv -f "${project_root}/dns-scout-macos-arm64-silicon-v${VERSION}.tar.gz" "${project_root}/../binaries/"

# Before running dpkg-buildpackage, update debian/source/options to include --include-removal
echo "--include-removal" >> "${project_root}/debian/source/options"

# Run dpkg-buildpackage
echo "======== Running Debian packaging process..."
dpkg-buildpackage -k${GPG_KEY_ID}

# Run source changes
echo "======== Running Debian source changes process..."
dpkg-buildpackage -S -k${GPG_KEY_ID}

# Move back go binary
mv -f "${project_root}/../dns-scout_${VERSION}.orig.tar.gz" "${project_root}/../binaries/"
mv -f "${project_root}/../dns-scout_${VERSION}-1debian1.debian.tar.xz" "${project_root}/../binaries/"

mv -f "${project_root}/../dns-scout_${VERSION}-1debian1_amd64.deb" "${project_root}/../binaries/"
mv -f "${project_root}/../dns-scout_${VERSION}-1debian1_amd64.buildinfo" "${project_root}/../binaries/"
mv -f "${project_root}/../dns-scout_${VERSION}-1debian1_amd64.changes" "${project_root}/../binaries/"

mv -f "${project_root}/../dns-scout_${VERSION}-1debian1.dsc" "${project_root}/../binaries/"
mv -f "${project_root}/../dns-scout_${VERSION}-1debian1_source.buildinfo" "${project_root}/../binaries/"
mv -f "${project_root}/../dns-scout_${VERSION}-1debian1_source.changes" "${project_root}/../binaries/"
mv -f "${project_root}/../dns-scout_${VERSION}-1debian1_source.build" "${project_root}/../binaries/"


mv -f "${project_root}/../binaries/dns-scout" "${project_root}/bin/"
echo "Build complete."
