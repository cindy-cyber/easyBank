#!/bin/sh

set -e

echo "run db migration"
/app/migrate -path /app/migration -database "$DB_SOURCE" -verbose up
# call the /app/migrate binary, pass in the path of all migration files, AND
# DB url, up runs all migrations

echo "start the app"
exec "$@"  # take all parameters, pass to the script, and run it, EXPECTED: /app/main in Dockerfile