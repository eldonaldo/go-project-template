#!/usr/bin/env bash

# Absolute path to project root
base="$( cd "$(dirname "$0")" ; pwd -P )/.."
cd ${base}

# name of the migration
name="not-set"
if [[ -z ${1+x} ]]; then
    echo "Migration name not set, aborting (usage: ./scripts/migration_create.sh migration_name)";
    exit
else
    # concat all args into a single string
    name="$*"

    # convert whitespaces to underscores
    name=${name// /_}
fi

migration_folder="core/db/migrations"
database="database.db"

goose -dir ${migration_folder} sqlite3 ./${database} create ${name} sql