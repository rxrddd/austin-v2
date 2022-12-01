# Kratos Base Project

本项目为一个使用 [kratos框架](https://github.com/go-kratos/kratos) 创建的，基础的微服务项目模板，
以便于后续快速开发。

## 项目目录

```
├── api // api proto文件和生成文件
├── app // 服务集合
│   ├── administrator // 管理员服务
│   ├── authorization // 权限服务
│   └── project // api接口服务
│       └── // api接口服务
├── deploy // 部署文件
├── pkg // 自定义包
├── third_party // 第三方包
├── web // 管理后台前端页面 基于vue-element-admin
```


#### 管理员服务
- [x]  登录退出
- [x]  管理员管理

#### 权限服务
- [x]  角色管理
- [ ]  菜单管理
- [x]  权限管理
- [x]  api管理
