#!/bin/bash

# Make sure to install Go 1.21 before running this script
# or update the path to the Go binary accordingly.
# Define the GPG key ID
GPG_KEY_ID="EF536354988BF362947FC6FDBEB7932396E8FB23"
# Define the project root directory
project_root="."
export VERSION=$(grep Version constant/constants.go  | awk -F '=' '{print $2}'| awk -F'"' '{print $2}')

# Run env subtitute
envsubst < dns-scout.metainfo.xml.tpl > dns-scout.metainfo.xml
envsubst < debian/files.tpl > debian/files
envsubst < debian/changelog.tpl > debian/changelog

# Clean up the old generated binaries and artifacts
echo "======== Cleaning up generated binaries and artifacts..."
rm -f "${HOME}/binaries/dns-scout-linux-*.tar.gz"
# Move out binary files
# mkdir -p "$HOME/binaries"
# cp "${project_root}/bin/dns-scout" "${project_root}/../binaries/"

# Remove existing upstream tarball if it exists
[ -f "${HOME}/binaries/dns-scout_${VERSION}.orig.tar.gz" ] && rm "${HOME}/binaries/dns-scout_${VERSION}.orig.tar.gz"


# Create the upstream tarball and place it in the correct directory
echo "Creating upstream tarball..."
tar czvf "${HOME}/binaries/dns-scout_${VERSION}.orig.tar.gz" --exclude='.git' --exclude='./dns-scout-linux-*' -C "${project_root}" .

# Step 1: Compile the Go code
echo "=======Compiling Go code..."
CGO_ENABLED=0 go build -a -v -o "$HOME/binaries/dns-scout"

CGO_ENABLED=0 GOOS=darwin GOARCH=amd64 go build -a -v -o "$HOME/binaries/dns-scout-macos-amd64-intel-v${VERSION}/dns-scout"
# produces a binary for macOS running on Intel x86_64 architecture (Intel Macs). It does not produce a binary for the newer Apple Silicon Macs, which use the ARM64 architecture.

CGO_ENABLED=0 GOOS=darwin GOARCH=arm64 go build -a -v -o "$HOME/binaries/dns-scout-macos-arm64-silicon-v${VERSION}/dns-scout"
# This will produce a binary (dns-scout-macos-arm64) that runs on macOS systems with ARM64 architecture (Apple Silicon).

CGO_ENABLED=0 GOOS=linux GOARCH=amd64 go build -a -v -o "$HOME/binaries/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}/dns-scout"
# This will generate a binary (dns-scout-linux-amd64) that is suitable for most Kali and Ubuntu installations on AMD64/x86_64 hardware.

CGO_ENABLED=0 GOOS=linux GOARCH=arm64 go build -a -v -o "$HOME/binaries/dns-scout-linux-arm64-raspberry-pi-v${VERSION}/dns-scout"
# Raspberry Pi 64-bit ARM

CGO_ENABLED=0 GOOS=linux GOARCH=386 go build -a -v -o "$HOME/binaries/dns-scout-linux-386-v${VERSION}/dns-scout"
# If you want to support older 32-bit machines or other architectures, you'll need to specify different GOARCH values. For example, for 32-bit x86:

# Create Debian packages for Linux builds
for arch in amd64 arm64 386; do
  deb_folder="$HOME/binaries/dns-scout-linux-${arch}-v${VERSION}-deb"
  deb_name="dns-scout-linux-${arch}-v${VERSION}-1debian1.deb"

  mkdir -p "${deb_folder}/usr/local/bin"
  mkdir -p "${deb_folder}/usr/share/doc/dns-scout"
  mkdir -p "${deb_folder}/DEBIAN"

  # Copy the binary to the package directory and set permissions
  case "${arch}" in
    amd64)
      binary_path="$HOME/binaries/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}/dns-scout"
      ;;
    arm64)
      binary_path="$HOME/binaries/dns-scout-linux-arm64-raspberry-pi-v${VERSION}/dns-scout"
      ;;
    386)
      binary_path="$HOME/binaries/dns-scout-linux-386-v${VERSION}/dns-scout"
      ;;
  esac

  if [ -f "${binary_path}" ]; then
    cp "${binary_path}" "${deb_folder}/usr/local/bin/"
    chmod 755 "${deb_folder}/usr/local/bin/dns-scout"
  else
    echo "The binary for architecture ${arch} does not exist at the expected path: ${binary_path}"
    exit 1
  fi

  # Copy documentation to the package directory
  cp README.md "${deb_folder}/usr/share/doc/dns-scout/"
  cp setup-api-token.sh "${deb_folder}/usr/share/doc/dns-scout/"

  # Create the control file for the package
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

  # Build the Debian package and specify the output directory
  dpkg-deb --build "${deb_folder}" "$HOME/binaries/${deb_name}"
  if [ $? -ne 0 ]; then
    echo "Failed to build package for architecture: ${arch}"
    exit 1
  fi
done

tar czvf "${HOME}/binaries/dns-scout-macos-amd64-intel-v${VERSION}.tar.gz" --transform 's,^./dns-scout-macos-amd64-intel-v${VERSION}/dns-scout,dns-scout,' -C "${HOME}/binaries" "dns-scout-macos-amd64-intel-v${VERSION}/dns-scout" -C "${HOME}/dns-scout" README.md setup-api-token.sh

tar czvf "${HOME}/binaries/dns-scout-macos-arm64-silicon-v${VERSION}.tar.gz" --transform 's,^./dns-scout-macos-arm64-silicon-v${VERSION}/dns-scout,dns-scout,' -C "${HOME}/binaries" "dns-scout-macos-arm64-silicon-v${VERSION}/dns-scout" -C "${HOME}/dns-scout" README.md setup-api-token.sh

tar czvf "${HOME}/binaries/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}.tar.gz" --transform 's,^./dns-scout-linux-amd64-ubuntu-kali-v${VERSION}/dns-scout,dns-scout,' -C "${HOME}/binaries" "dns-scout-linux-amd64-ubuntu-kali-v${VERSION}/dns-scout" -C "${HOME}/dns-scout" README.md setup-api-token.sh

tar czvf "${HOME}/binaries/dns-scout-linux-arm64-raspberry-pi-v${VERSION}.tar.gz" --transform 's,^./dns-scout-linux-arm64-raspberry-pi-v${VERSION}/dns-scout,dns-scout,' -C "${HOME}/binaries" "dns-scout-linux-arm64-raspberry-pi-v${VERSION}/dns-scout" -C "${HOME}/dns-scout" README.md setup-api-token.sh

tar czvf "${HOME}/binaries/dns-scout-linux-386-v${VERSION}.tar.gz" --transform 's,^./dns-scout-linux-386-v${VERSION}/dns-scout,dns-scout,' -C "${HOME}/binaries" "dns-scout-linux-386-v${VERSION}/dns-scout" -C "${HOME}/dns-scout" README.md setup-api-token.sh

for file in \
"${HOME}/binaries/dns-scout-macos-amd64-intel-v${VERSION}/dns-scout" \
"${HOME}/binaries/dns-scout-macos-arm64-silicon-v${VERSION}/dns-scout" \
"${HOME}/binaries/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}/dns-scout" \
"${HOME}/binaries/dns-scout-linux-386-v${VERSION}/dns-scout" \
"${HOME}/binaries/dns-scout-linux-arm64-raspberry-pi-v${VERSION}/dns-scout" \
"${HOME}/binaries/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}-1debian1.deb" \
"${HOME}/binaries/dns-scout-linux-arm64-raspberry-pi-v${VERSION}-1debian1.deb" \
"${HOME}/binaries/dns-scout-linux-386-v${VERSION}-1debian1.deb"
do
    if [ -f "$file" ]; then
        shasum -a 256 "$file"
    else
        echo "File not found: $file"
    fi
done

echo "looking at folders..."
ls -lart "${HOME}/binaries" | grep dns-scout

# The following move operations are not needed as binaries and packages are already directed to the ~/binaries directory
# echo "======== Moving generated binaries and packages to debian/binaries/"
# mv -f "${project_root}/dns-scout-linux-386-v${VERSION}-1debian1.deb" "$HOME/binaries/"
# mv -f "${project_root}/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}-1debian1.deb" "$HOME/binaries/"
# mv -f "${project_root}/dns-scout-linux-arm64-raspberry-pi-v${VERSION}-1debian1.deb" "$HOME/binaries/"
# mv -f "${project_root}/dns-scout-linux-386-v${VERSION}.tar.gz" "${project_root}/../binaries/"
# mv -f "${project_root}/dns-scout-linux-amd64-ubuntu-kali-v${VERSION}.tar.gz" "${project_root}/../binaries/"
# mv -f "${project_root}/dns-scout-linux-arm64-raspberry-pi-v${VERSION}.tar.gz" "${project_root}/../binaries/"
# mv -f "${project_root}/dns-scout-macos-amd64-intel-v${VERSION}.tar.gz" "${project_root}/../binaries/"
# mv -f "${project_root}/dns-scout-macos-arm64-silicon-v${VERSION}.tar.gz" "${project_root}/../binaries/"

# Run dpkg-buildpackage
echo "======== Running Debian packaging process..."
cd ~/binaries
dpkg-buildpackage -k${GPG_KEY_ID}

# Run source changes
echo "======== Running Debian source changes process..."
dpkg-buildpackage -S -k${GPG_KEY_ID}

echo "Build complete."
