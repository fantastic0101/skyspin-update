#!/bin/bash

host=""
hosts="
doudou-test
doudou-prod
dou-ph-prod
dou-idr-prod
"
select host in $hosts ; do
    break
done

echo host $host

rsync -avz -P --delete ./dist root@${host}:/data/game-admin/
