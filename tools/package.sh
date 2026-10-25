
#!/usr/bin/env bash
set -euo pipefail

SCRIPT_DIR="$(cd "$(dirname "${BASH_SOURCE[0]}")" && pwd)"
PROJECT_ROOT="$(cd "$SCRIPT_DIR/.." && pwd)"

OUT_DIR="$PROJECT_ROOT/out"
PACKAGE_DIR="$OUT_DIR/package"
ZIP_FILE="$OUT_DIR/2048.zip"

cd "$PROJECT_ROOT"

echo "==> Preparing package..."

rm -rf "$PACKAGE_DIR"
rm -f "$ZIP_FILE"

PORTS_DIR="$PACKAGE_DIR/roms/ports"
GAME_DIR="$PORTS_DIR/2048"

mkdir -p "$GAME_DIR"
mkdir -p "$PORTS_DIR/.res"

cp "$PROJECT_ROOT/port/2048.sh" "$PORTS_DIR/2048.sh"
cp "$PROJECT_ROOT/port/boxart.png" "$PORTS_DIR/.res/2048.png"
cp "$PROJECT_ROOT/port/run.sh" "$GAME_DIR/run.sh"
cp "$PROJECT_ROOT/port/boxart.png" "$GAME_DIR/boxart.png"
cp "$PROJECT_ROOT/port/config.json" "$GAME_DIR/config.json"
cp "$OUT_DIR/2048" "$GAME_DIR/2048"

chmod +x "$GAME_DIR/run.sh"
chmod +x "$GAME_DIR/2048"
chmod +x "$PORTS_DIR/2048.sh"

echo "==> Creating ZIP archive..."

cd "$PACKAGE_DIR"

if command -v zip >/dev/null 2>&1; then
    echo "==> Using zip"
    zip -qr "$ZIP_FILE" roms

elif command -v 7z >/dev/null 2>&1; then
    echo "==> Using 7z"
    7z a -tzip "$ZIP_FILE" roms >/dev/null

elif command -v 7zz >/dev/null 2>&1; then
    echo "==> Using 7zz"
    7zz a -tzip "$ZIP_FILE" roms >/dev/null

elif command -v bsdtar >/dev/null 2>&1; then
    echo "==> Using bsdtar"
    bsdtar -a -cf "$ZIP_FILE" roms

elif command -v python3 >/dev/null 2>&1; then
    echo "==> Using Python zipfile"
    python3 - "$ZIP_FILE" <<'PY'
import os
import sys
import zipfile

output = sys.argv[1]

with zipfile.ZipFile(output, "w", zipfile.ZIP_DEFLATED) as archive:
    for root, dirs, files in os.walk("roms"):
        for filename in files:
            path = os.path.join(root, filename)
            archive.write(path)
PY

else
    echo "ERROR: No ZIP tool found (zip, 7z, 7zz, bsdtar, python3)." >&2
    exit 1
fi

echo "==> Package created: $ZIP_FILE"
