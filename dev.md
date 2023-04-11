```
kratos proto client api/msgpusher/v1/msgpusher.proto
kratos proto client api/msgpusher/v1/msgpusher.error.proto
kratos proto client api/msgpusher-manager/v1/msgpusher-manager.proto
kratos proto client api/project/admin/v1/admin.proto
```

- 使用docker-compose安装ectd，redis，mysql等软件
```
docker-compose -f deply/env/docker-compose.yml up -d
```
- 导入austin-v2数据库
- build 镜像
```
docker build -f app/mgr/Dockerfile -t mgr .
docker build -f app/msgpusher/Dockerfile -t msgpusher .
docker build -f app/msgpusher-manager/Dockerfile -t msgpusher-manager .
docker build -f app/msgpusher-worker/Dockerfile -t msgpusher-worker .
```
- run起来
```
docker run --rm -d -p 8000:8000 --network deply_software_network  -v /yourdir/app/mgr/configs:/data/conf mgr
docker run --rm -d -p 7000:7000 -p 8001:8000 --network deply_software_network  -v /yourdir/app/msgpusher/configs:/data/conf msgpusher
docker run --rm -d -p 10000:8000 --network deply_software_network  -v /yourdir/app/msgpusher-manager/configs:/data/conf msgpusher-manager
docker run --rm -d --network deply_software_network  -v /yourdir/app/msgpusher-worker/configs:/data/conf msgpusher-worker
```




