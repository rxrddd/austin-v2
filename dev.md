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
```
sh deployment-shell.sh
```