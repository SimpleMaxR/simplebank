#!/bin/bash

set -e

echo "run db migration"
goose -dir ./migration postgres "postgres://root:secret@postgres:5432/simple_bank?sslmode=disable" up

echo "start the app"
exec "$@"