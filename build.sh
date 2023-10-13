#!/bin/bash

# Make sure to install Go 1.21 before running this script
# or update the path to the Go binary accordingly.

# Step 1: Compile the Go code
echo "Compiling Go code..."
go build -v -o ./bin/dns-scout main.go

echo "Build complete."
