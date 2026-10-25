#!/bin/sh
set -eu

PORT_DIR=/mnt/sdcard/roms/ports/2048

cd "$PORT_DIR"

export HOME="$PORT_DIR"
export LC_ALL=C LANG=C
export LD_LIBRARY_PATH=/mnt/sdcard/cubegm/lib:/lib:/usr/lib

exec "$PORT_DIR/2048"
