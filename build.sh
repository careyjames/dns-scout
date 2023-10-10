#!/bin/bash

# Step 1: Compile the Go code
echo "Compiling Go code..."
go build -o dns-scout

# Step 2: Create the directory structure
echo "Creating directory structure..."
mkdir -p DNS-Scout/bin
mkdir -p DNS-Scout/share

# Step 3: Move the compiled executable to bin/
echo "Moving compiled executable to bin/..."
mv dns-scout DNS-Scout/bin/

# Step 4: Copy the shell script to share/
echo "Copying shell script to share/..."
cp setup-api-token.sh DNS-Scout/share/

echo "Build complete."
