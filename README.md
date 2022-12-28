# Kratos Base Project

本项目为一个使用 [kratos框架](https://github.com/go-kratos/kratos) 创建的，基础的微服务项目模板，
以便于后续快速开发。

目前已经实现管理后台管理员和权限系统，已经实现按钮、接口级权限. 

[演示地址](http://kratos.niu12.com/#/system/adminstrator)

[项目前端地址](https://austin-v2-frontend)

PS:演示数据由定时任务服务，每10分钟初始化恢复demo数据库，可能会导致你的数据丢失
## 项目目录

```
├── api // api proto文件和生成文件
├── app // 服务集合
│   ├── administrator // 管理员服务
│   ├── authorization // 权限服务
│   ├── jobs // 分布式定时任务
│   ├── └── initDB // 初始化数据库
│   └── project // api接口服务
│       └── // api接口服务
├── deploy // 部署文件
├── pkg // 自定义包
├── third_party // 第三方包
```

#### 管理后台服务
- [x]  管理后台

#### 管理员服务
- [x]  登录退出
- [x]  管理员管理

#### 权限服务
- [x]  角色管理
- [x]  菜单管理
- [x]  权限管理
- [x]  api管理

#### 文件服务
- [x]  文件上传(后端直传)
- [x]  前端直传(获取token)

#### 通知服务
- [ ] 短信服务
- [ ] 邮件服务


#### 运行方式

##### 组件
管理后台web框架: vue-element-admin

服务注册与发现： ETCD

链路追踪：jaeger

数据库：mysql

缓存：redis

orm: GORM



### 安装
##### 数据库
1.导入sql
```
文件位于docs/initSql
```

##### 后端
1.下载
```
$ go clone austin-v2
```

2.安装依赖
```
$ cd austin-v2 && go mod tidy
```

3.设置配置 以管理员服务为例, 配置文件位于service/configs/  初始化sql文件位于 docs/initSql
```
$ vim ./app/administrator/service/configs/config.yaml
```

4.运行项目
```
$ kratos run
```

##### 前端
1.安装依赖
```
$ cd web && npm install
```

2.启动项目
```
$ npm run dev
```

### 部署(docker)
##### 后端
可以参考kratos部署 (https://go-kratos.dev/docs/devops/docker)

1.服务部署 以管理员服务为例 app/
```
$ cd app/administrator/service
```
2.make打包docker镜像
```
# PS:如果是打包admin镜像 app/project/admin/service 请执行 make dockerAdmin
$ make docker
```
3.运行容器 
```
# 注意端口映射设置， docker部署容器9000端口， 本地开发端口不能全是9000
docker run -p 9000:9000 --name austin-v2-administrator --restart=always -v /data/project/austin-v2/app/administrator/service/configs:/data/conf -d austin-v2/administrator-service:0.1.0
```
##### 前端
1.进入前端目录
```
$ cd web
```
2.编译
```
$ npm run build:prod
```
3.将dist文件夹上传至服务器

##### nginx示例
```
server {
  listen 80;
  listen [::]:80;
  server_name kratos.niu12.com;
  index index.html;
  root /data/project/austin-v2/web/dist;

  # 管理后台接口转发代理
  location  /api/ {
  # nginx代理设置Header
      proxy_set_header            X-real-ip $remote_addr;

      proxy_pass                  http://127.0.0.1:8000/;
  }

}
```

* 有任何建议，请扫码添加我微信进行交流。

![扫码提建议](https://austin-v2.oss-cn-hangzhou.aliyuncs.com/f8f5dacdf87cf358c98c9eb60ce2a13.jpg)
