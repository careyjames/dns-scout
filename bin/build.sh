#!/bin/bash
- uses: actions/setup-go@v2
  with:
    go-version: 1.21
    modules-directory: src

# Step 1: Compile the Go code
echo "Compiling Go code..."
cd ../src && go build -v -o ../bin/dns-scout main.go && cd -

# Step 2: Create the directory structure
echo "Creating directory structure..."
mkdir -p ../bin/DNS-Scout/share

# Step 4: Copy the shell script to share/
echo "Copying shell script to share/..."
cp ../share/setup-api-token.sh ../bin/DNS-Scout/share/setup-api-token.sh

echo "Build complete."
