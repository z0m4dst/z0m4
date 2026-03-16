#!/usr/bin/env bash

BASE="$HOME/.z0m4sandbox"
TMP="$BASE/tmp"

mkdir -p "$TMP"

SANDBOX=$(mktemp -d "$TMP/sbox.XXXX")

mkdir -p "$SANDBOX/bin"

cp /bin/sh "$SANDBOX/bin/"
cp /bin/echo "$SANDBOX/bin/"
cp /bin/ls "$SANDBOX/bin/"

SCRIPT="$SANDBOX/run.sh"

echo "paste script — CTRL+D to run"
echo ""

cat > "$SCRIPT"

chmod +x "$SCRIPT"

echo "[sandbox running]"

chroot "$SANDBOX" /bin/sh /run.sh

echo "[sandbox destroyed]"

rm -rf "$SANDBOX"
