#!/bin/bash

# Make sure to install Go 1.21 before running this script
# or update the path to the Go binary accordingly.
# Define the GPG key ID
GPG_KEY_ID="EF536354988BF362947FC6FDBEB7932396E8FB23"
# Celliwig temp key
GPG_KEY_ID="27B07EC1CEB62E5BBA0947F8D813438DD52A4F0C"
# Define the project root directory
project_root=`realpath $(pwd)`
export VERSION=$(grep Version constant/constants.go  | awk -F '=' '{print $2}'| awk -F'"' '{print $2}')

# Run env subtitute
envsubst < dns-scout.metainfo.xml.tpl > dns-scout.metainfo.xml
#envsubst < debian/files.tpl > debian/files
#envsubst < debian/changelog.tpl > debian/changelog

# Remove existing upstream tarball if it exists
[ -f "../dns-scout_${VERSION}.orig.tar.gz" ] && rm "../dns-scout_${VERSION}.orig.tar.gz"
# Remove debian/patches/
[ -d debian/patches/ ] && rm -rf debian/patches/
# Remove GO vendor modules
[ -d vendor/ ] && rm -rf vendor/

# Update version from changelog (because we might not have set it earlier for debugging)
VERSION=`head -n 1 debian/changelog| sed -E 's|[a-z-]+\s+\(([^\)]+)\)\s+.*|\1|'`

# Create the upstream tarball and place it in the parent directory
echo "Creating upstream tarball..."
tar czvf "../dns-scout_${VERSION}.orig.tar.gz" --exclude='.git' --exclude='.gitattributes' --exclude='.github/*' --exclude='.gitignore' \
						--exclude='*.tpl' \
						--exclude='./bin/*' \
						--exclude='./dns-scout-linux-*' \
						--exclude='./generate_sitemap/*' \
						-C "${project_root}" .

# Download required GO modules
go mod vendor

# Generate a patch file to include vendor module source for LaunchPad build server
# Set editor to bypass having to write description
EDITOR=/bin/true dpkg-source --commit ./ added-vendor-modules

# Build source package
dpkg-buildpackage -S -sa -k${GPG_KEY_ID}

echo "Source build complete."
