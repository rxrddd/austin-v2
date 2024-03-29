# 基于kratos 的聚合消息推送平台

### 注意：当前项目只用于交流学习，请勿直接使用，随时可能会大修改
#### 介绍

基于kratos 的聚合消息推送平台 austin-go项目的v2版本 

> v2版本
#### github地址：[https://github.com/rxrddd/austin-v2](https://github.com/rxrddd/austin-v2)

> v1版本
#### github地址：[https://github.com/rxrddd/austin-go](https://github.com/rxrddd/austin-go)
#### gitee地址：[https://gitee.com/AbelZou/austin-go](https://gitee.com/AbelZou/austin-go)


#### 项目描述

1. 基于kratos/grpc/ants/asynq/mysql/redis 写的一个聚合消息推送平台
使用docker快速部署
 - 使用docker-compose安装ectd，redis，mysql等软件
```
docker-compose -f deply/env/docker-compose.yml up -d
```
- 修改每个服务的config.yaml中的mysql地址，redis地址，etcd地址等信息
- 导入数据库`docs/sql/austin-v2.sql`

- 运行脚本
```
sh deployment-shell.sh
```


1. 如需要测试去重服务则修改`message_template`表中的`deduplication_config`字段
```
{"deduplication_10":{"num":1,"time":300},"deduplication_20":{"num":5}}
```
5. 使用示例
> 邮件消息
```
curl --location --request POST 'http://localhost:8001/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "code": "send",
    "messageParam": {
        "receiver": "test@qq.com",
        "variables": {
            "title": "测试操作",
            "content": "Hello <b>Bob</b> and <i>Cora</i>!"
        }
    },
    "messageTemplateId": 2
}'
```

> 微信公众号消息
```
curl --location --request POST 'http://localhost:8001/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "messageParam": {
        "receiver": "okEEF6WB92HO14qdy0Nosq62OVyY",
        "variables": {
            "data": {
                "order_no": "DD12345678", //模板参数
                "time": "2022-01-11 10:00:00" //模板参数
            }
        }
    },
    "messageTemplateId": 4
}'

//参数带颜色的
curl --location --request POST 'http://localhost:8001/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "messageParam": {
        "receiver": "openId",
        "variables": {
            "data": {
                "name":"张三12333|#0000FF"
            }
        }
    },
    "messageTemplateId": 4
}'
```

> 钉钉自定义机器人
```
//艾特某些手机号
curl --location --request POST 'http://localhost:8001/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "messageParam": {
        "receiver": "13588888888,13588888887",
        "variables": {
            "content": "测试\n换行"
        }
    },
    "messageTemplateId": 5
}'

//艾特全部人
curl --location --request POST 'http://localhost:8001/send' \
--header 'Content-Type: application/json' \
--data-raw '{
    "messageParam": {
        "receiver": "@all",
        "variables": {
            "content": "测试\n换行"
        }
    },
    "messageTemplateId": 5
}'
```




#### 目录说明

```
├── Makefile
├── README.md
├── api   //proto 接口定义
├── app
│   ├── mgr //admin管理端接口  内含rbac 调用msgpusher-manager rpc服务接口
│   ├── msgpusher-manager //聚合消息后台curd rpc服务
│   ├── msgpusher //消息推送服务对外http接口
│   ├── msgpusher-common //公用文件
│   ├── msgpusher-worker //消息推送消费端
├── deploy  //部署文件
│   ├── env 
│   │   ├── docker-compose.yml 所需mysql，etcd，redis组件环境 
├── dev.md
├── docs  //sql文件
├── deployment-shell.sh 一键部署所有模块的脚本
├── go.mod
├── go.sum
├── openapi.yaml
├── pkg  //公共包
└── third_party 
```




#### Thanks


kratos：[https://github.com/go-krato/krato](https://github.com/go-kratos/kratos)

austin：[https://gitee.com/zhongfucheng/austin](https://gitee.com/zhongfucheng/austin)

ants：[https://github.com/panjf2000/ants](https://github.com/panjf2000/ants)

asynq：[https://github.com/hibiken/asynq](https://github.com/hibiken/asynq)

gomail：[https://gopkg.in/gomail.v2](https://gopkg.in/gomail.v2)

wechat：[https://github.com/silenceper/wechat](https://github.com/silenceper/wechat)

ding：[https://github.com/wanghuiyt/ding](https://github.com/wanghuiyt/ding)
