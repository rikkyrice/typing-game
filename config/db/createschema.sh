#!/bin/bash

/database/config/db2inst1/sqllib/bin/db2 connect to USERDB user db2inst1 using password

# create tables
/database/config/db2inst1/sqllib/bin/db2 -tvf /var/custom/sql/users_create.sql
/database/config/db2inst1/sqllib/bin/db2 -tvf /var/custom/sql/wordlists_create.sql
/database/config/db2inst1/sqllib/bin/db2 -tvf /var/custom/sql/words_create.sql
/database/config/db2inst1/sqllib/bin/db2 -tvf /var/custom/sql/scores_create.sql
/database/config/db2inst1/sqllib/bin/db2 -tvf /var/custom/sql/token_create.sql

# insert data
/database/config/db2inst1/sqllib/bin/db2 import from /var/custom/data/users.csv of del insert into Users
/database/config/db2inst1/sqllib/bin/db2 import from /var/custom/data/wordlists.csv of del insert into WordLists
/database/config/db2inst1/sqllib/bin/db2 import from /var/custom/data/words.csv of del insert into Words
/database/config/db2inst1/sqllib/bin/db2 import from /var/custom/data/scores.csv of del insert into Scores

/database/config/db2inst1/sqllib/bin/db2 terminate
