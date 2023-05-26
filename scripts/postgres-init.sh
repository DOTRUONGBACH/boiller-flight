#!/bin/bash

set -e

ADDRESS=${1:-""}
ADDRESS_OPTS=
if [ ! -z ${ADDRESS} ]; then
    ADDRESS_OPTS="-h ${ADDRESS}"
fi

BACKEND_API_DATABASE="backend_db"
BACKEND_API_USER="backend_user"
BACKEND_API_PASSWORD="backend_password"

export PGPASSWORD=${POSTGRES_PASSWORD}

psql ${ADDRESS_OPTS} \
    -v ON_ERROR_STOP=1 \
    --username "${POSTGRES_USER:-postgres}" \
    --dbname "${POSTGRES_DB:-postgres}" \
    <<-EOSQL
CREATE USER ${BACKEND_API_USER} WITH ENCRYPTED PASSWORD '${BACKEND_API_PASSWORD}';
CREATE DATABASE ${BACKEND_API_DATABASE};

GRANT ALL PRIVILEGES ON DATABASE ${BACKEND_API_DATABASE} TO ${BACKEND_API_USER};
EOSQL

psql ${ADDRESS_OPTS} \
    -v ON_ERROR_STOP=1 \
    --username "${POSTGRES_USER:-postgres}" \
    --dbname "${BACKEND_API_DATABASE}" \
    <<-EOSQL
CREATE EXTENSION pgcrypto;
CREATE EXTENSION btree_gist;
CREATE EXTENSION pg_trgm;
CREATE EXTENSION "uuid-ossp";
EOSQL

unset PGPASSWORD