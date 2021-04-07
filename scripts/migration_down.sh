#!/usr/bin/env bash

# Absolute path to project root
base="$( cd "$(dirname "$0")" ; pwd -P )/.."
cd ${base}

migration_folder="core/db/migrations"
database="database.db"

# downgrades an already performed migration
goose -dir ${migration_folder} sqlite3 ./${database} down