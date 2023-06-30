#!/bin/sh

/go/src/app/wait-for redis.ms.local:6379 -- echo "redis is up"
/go/src/app/app
