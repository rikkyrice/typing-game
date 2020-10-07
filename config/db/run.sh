#!/bin/bash
. ./env.list
docker rm -f typing-game
docker run --name typing-game -h db2server --restart=always --detach --privileged=true -p 50000:50000 -p 55000:55000 --env-file env.list typing-game:v1.0
docker ps -q --filter "name=typing-game"
docker logs -f typing-game
# db2_container_id=$(docker ps -q --filter "name=typing-game")
# while true
# do
#   docker exec -i ${db2_container_id} test -e /tmp/end.txt
#   if [ $? -eq 0 ]; then
#     echo "finished."
#     exit
#   fi
#   echo "waiting."
#   sleep 3
# done