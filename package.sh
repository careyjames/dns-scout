#!/bin/bash

# Package the dns-scout directory into a tarball
tar --exclude='.git' --exclude='.gitignore' --exclude='.vscode' --exclude='.tmp-history' -czvf "dns-scout-v6.0.tar.gz" ./bin/
