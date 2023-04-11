#!/bin/bash

#停止服务
docker stop mgr
docker stop msgpusher
docker stop msgpusher-manager
docker stop msgpusher-worker


#删除容器
docker rm mgr
docker rm msgpusher
docker rm msgpusher-manager
docker rm msgpusher-worker

#删除镜像
docker rm mgr:v1
docker rm msgpusher:v1
docker rm msgpusher-manager:v1
docker rm msgpusher-worker:v1

#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')

#构建服务
docker build -t mgr:v1 -f app/mgr/Dockerfile  .
docker build -t msgpusher:v1 -f app/msgpusher/Dockerfile -t msgpusher .
docker build -t msgpusher-manager:v1 -f app/msgpusher-manager/Dockerfile -t msgpusher-manager .
docker build -t msgpusher-worker:v1 -f app/msgpusher-worker/Dockerfile -t msgpusher-worker .

#启动服务
docker run -itd --net=host --name=mgr mgr:v1
docker run -itd --net=host --name=msgpusher msgpusher:v1
docker run -itd --net=host --name=msgpusher-manager msgpusher-manager:v1
docker run -itd --net=host --name=msgpusher-worker msgpusher-worker:v1