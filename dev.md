```
kratos proto client api/msgpusher/v1/msgpusher.proto
kratos proto client api/msgpusher/v1/msgpusher.error.proto
kratos proto client api/msgpusher-manager/v1/msgpusher-manager.proto
kratos proto client api/project/admin/v1/admin.proto
```

```
kratos proto server api/msgpusher/v1/msgpusher.proto -t app/msgpusher/internal/service
```
```
protoc --proto_path=.  --proto_path=../../third_party  --go_out=paths=source_relative:. internal/conf/conf.proto
```



