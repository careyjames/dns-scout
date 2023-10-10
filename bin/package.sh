#!/bin/bash

# Package the DNS-Scout directory into a tarball
tar --exclude='.git' --exclude='.gitignore' --exclude='.vscode' --exclude='.tmp-history' -czvf "DNS-Scout-v5.8.tar.gz" ./DNS-Scout/
