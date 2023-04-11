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
docker rmi mgr:v1
docker rmi msgpusher:v1
docker rmi msgpusher-manager:v1
docker rmi msgpusher-worker:v1

#删除none镜像
docker rmi $(docker images | grep "none" | awk '{print $3}')

#构建服务
docker build -t mgr:v1 -f app/mgr/Dockerfile  .
docker build -t msgpusher:v1 -f app/msgpusher/Dockerfile .
docker build -t msgpusher-manager:v1 -f app/msgpusher-manager/Dockerfile .
docker build -t msgpusher-worker:v1 -f app/msgpusher-worker/Dockerfile .

#启动服务
docker run -itd --net=host -v /www/wwwroot/austin-v2/app/mgr/configs:/data/conf --name=mgr mgr:v1
docker run -itd --net=host -v /www/wwwroot/austin-v2/app/msgpusher/configs:/data/conf --name=msgpusher msgpusher:v1
docker run -itd --net=host -v /www/wwwroot/austin-v2/app/msgpusher-manager/configs:/data/conf --name=msgpusher-manager msgpusher-manager:v1
docker run -itd --net=host -v /www/wwwroot/austin-v2/app/msgpusher-worker/configs:/data/conf --name=msgpusher-worker msgpusher-worker:v1