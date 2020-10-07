#!/bin/bash
. ./env.list
docker rm -f typing-game
docker run --name typing-game -h db2server --restart=always --detach --privileged=true -p 56000:56000 --env-file env.list ${REPOSITORY}:${TAG}
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