#!/bin/bash
set -e
set -o pipefail

cgogen yubikey.yml
cp /usr/src/yubico-c/*.h .
mv yubikey/* .
rm -rf yubikey
