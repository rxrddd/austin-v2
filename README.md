# 基于kratos 的聚合消息推送平台

#### 介绍
austin-go项目的v2版本

> v1版本
#### github地址：[https://github.com/rxrddd/austin-go](https://github.com/rxrddd/austin-go)
#### gitee地址：[https://gitee.com/AbelZou/austin-go](https://gitee.com/AbelZou/austin-go)

> v2版本
#### github地址：[https://github.com/rxrddd/austin-v2](https://github.com/rxrddd/austin-v2)

#### 项目描述

1. 基于kratos/grpc/ants/rabbitmq/mysql/redis 写的一个聚合消息推送平台
1. 开发时:
```
cd austin-v2/app/msgpusher && kratos run //启动grpc和http接口

cd austin-v2/app/msgpusher-worker && kratos run //启动消费端
```

1. 如需要测试去重服务则修改`message_template`表中的`deduplication_config`字段
```
{"deduplication_10":{"num":1,"time":300},"deduplication_20":{"num":5}}
```
5. 使用示例
> 邮件消息
```
curl --location --request POST 'http://localhost:8888/send' \
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
curl --location --request POST 'http://localhost:8888/send' \
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
curl --location --request POST 'http://localhost:8888/send' \
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
curl --location --request POST 'http://localhost:8888/send' \
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
curl --location --request POST 'http://localhost:8888/send' \
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
.
├── Makefile
├── README.md
├── api   //grpc 接口定义
├── app  //项目代码
│   ├── administrator //用户信息,登录
│   ├── authorization //授权
│   ├── files //文件上传oss
│   ├── msgpusher //msgpusher的rpc和http接口
│   ├── msgpusher-common //公共文件
│   ├── msgpusher-worker //消费端
│   └── project
│       └── admin  //后端接口项目
├── deploy  //部署文件
├── dev.md
├── docs  //文档
├── go.mod
├── go.sum
├── openapi.yaml
├── pkg  //公共包
└── third_party //三方包 谷歌啥的

```




#### Thanks


kratos：[https://github.com/go-krato/krato](https://github.com/go-kratos/kratos)

austin：[https://gitee.com/zhongfucheng/austin](https://gitee.com/zhongfucheng/austin)

ants：[https://github.com/panjf2000/ants](https://github.com/panjf2000/ants)

gomail：[https://gopkg.in/gomail.v2](https://gopkg.in/gomail.v2)

wechat：[https://github.com/silenceper/wechat](https://github.com/silenceper/wechat)

ding：[https://github.com/wanghuiyt/ding](https://github.com/wanghuiyt/ding)
