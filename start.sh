#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up

echo "start the app"
# means takes all params passed to the script, and run it
# /app/main will be run
exec "$@"