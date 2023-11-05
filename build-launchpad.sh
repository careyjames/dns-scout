#!/bin/bash

# Make sure to install Go 1.21 before running this script
# or update the path to the Go binary accordingly.
# Define the GPG key ID
GPG_KEY_ID="EF536354988BF362947FC6FDBEB7932396E8FB23"

project_root="."
export VERSION=$(grep Version constant/constants.go  | awk -F '=' '{print $2}'| awk -F'"' '{print $2}')

cp -R ./debian "${HOME}/binaries/debian"

# Run dpkg-buildpackage
echo "======== Running Debian packaging process..."
cd ~/binaries/debian
dpkg-buildpackage -k${GPG_KEY_ID}

# Run source changes
echo "======== Running Debian source changes process..."
dpkg-buildpackage -S -k${GPG_KEY_ID}

echo "Build complete."
