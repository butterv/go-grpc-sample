#!/bin/sh

export GO111MODULE=on

set -e

host="$1"
shift
user="$1"
shift
password="$1"
shift
dbname="$1"
shift
cmd="$@"

echo "waiting db..."
until mysql -h"$host" -u"$user" -p"$password" "$dbname" -e "SHOW TABLES;"
do
  >$2 echo -n "."
  sleep 1
done

>&2 echo "Mysql is up - executing command"

exec $cmd
