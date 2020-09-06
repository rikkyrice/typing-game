#!/bin/sh

. ./conf.txt

find ./sql -name '*.sql' | sort > test.txt

#上記4
cat test.txt |\
while read sql
do
mysql -u${USERNAME} -p${PASSWORD} --database=${DATABASE} < $sql
echo $sql;
done