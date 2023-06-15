#!/bin/sh

/go/src/app/wait-for db.credsystem.local:3306 -- echo "mariadb is up"
/go/src/app/wait-for redis.credsystem.local:6379 -- echo "redis is up"
/go/src/app/app migrate
/go/src/app/app api
