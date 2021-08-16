#!/bin/bash
set -e

sed -ri "s/#log_statement = 'none'/log_statement = 'all'/g" /var/lib/postgresql/data/postgresql.conf
sed -ri "s/#logging_collector = off/logging_collector = on/g" /var/lib/postgresql/data/postgresql.conf
sed -ri "s/#log_directory = 'log'/log_directory = '\/var\/lib\/postgresql\/data\/log'/g" /var/lib/postgresql/data/postgresql.conf
sed -ri "s/#log_filename = 'postgresql-%Y-%m-%d_%H%M%S.log'/log_filename = 'postgresql-%Y-%m-%d_%H%M%S.log'/g" /var/lib/postgresql/data/postgresql.conf
sed -ri "s/#log_rotation_size = 10MB/log_rotation_size = 10MB/g" /var/lib/postgresql/data/postgresql.conf
sed -ri "s/#log_truncate_on_rotation = off/log_truncate_on_rotation = off/g" /var/lib/postgresql/data/postgresql.conf
sed -ri "s/#log_destination = 'stderr'/log_destination = 'csvlog'/g" /var/lib/postgresql/data/postgresql.conf

mkdir -m 777 /var/lib/postgresql/log

cat /var/lib/postgresql/data/postgresql.conf

psql -v ON_ERROR_STOP=1 --username "$POSTGRES_USER" --dbname "$POSTGRES_DB" <<-EOSQL
	create table if not exists code
(
    id bigserial not null
        constraint code_pk
            primary key,
    phone varchar,
    code integer
);

alter table code owner to $POSTGRES_USER;
GRANT ALL PRIVILEGES ON DATABASE $POSTGRES_DB TO $POSTGRES_USER;
EOSQL
