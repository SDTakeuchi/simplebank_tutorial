#!/bin/sh

set -e

echo "start the app"
# means takes all params passed to the script, and run it
# /app/main will be run
exec "$@"
