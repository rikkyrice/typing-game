#!/bin/bash

BASE_DIR=/docker-entrypoint-initdb.d
SQL_DIR=${BASE_DIR}/sql
DATA_DIR=${BASE_DIR}/data

## create tables
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} < ${SQL_DIR}/01_users_create.sql
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} < ${SQL_DIR}/02_token_create.sql
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} < ${SQL_DIR}/03_wordlists_create.sql
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} < ${SQL_DIR}/04_words_create.sql
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} < ${SQL_DIR}/05_scores_create.sql
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} < ${SQL_DIR}/06_view_create.sql

## insert data
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -c "\copy users from ${DATA_DIR}/users.csv with csv"
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -c "\copy wordlists from ${DATA_DIR}/wordlists.csv with csv"
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -c "\copy words from ${DATA_DIR}/words.csv with csv"
psql -U ${POSTGRES_USER} -d ${POSTGRES_DB} -c "\copy scores from ${DATA_DIR}/scores.csv with csv"
