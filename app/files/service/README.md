# Kratos Project CRUD Template
基于[kratos-layout](https://github.com/go-kratos/kratos-layout) 生成的crud项目模板,根据文件夹名称生成第一个字母大写服务。

### 中间件
服务发现 consul

链路追踪 jaeger

缓存 redis

数据库 mysql

数据库中间件 gorm


### 使用方式

PS：前提已经安装好kratos

# 开始使用
### 创建项目
通过 kratos 命令创建项目指定模板：

```bash
kratos new administrator -r https://files.git
```

### 使用命令
```bash
# 初始化项目代码
make newServiceInit
```

### 更改数据库配置
```bash
vim ./internal/configs/config.yaml
```

### 更改数据表名称
```bash
vim ./internal/data/entity/serviceName.go TableName()
```

### 数据表结构
```sql
CREATE TABLE `serviceName` (
  `id` int unsigned NOT NULL AUTO_INCREMENT,
  `name` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '名称',
  `created_at` timestamp NOT NULL COMMENT '创建时间',
  `updated_at` timestamp NOT NULL COMMENT '更新时间',
  `deleted_at` char(20) CHARACTER SET utf8mb4 COLLATE utf8mb4_unicode_ci NOT NULL COMMENT '删除时间',
  PRIMARY KEY (`id`) USING BTREE
) ENGINE=InnoDB AUTO_INCREMENT=0 DEFAULT CHARSET=utf8mb4 COLLATE=utf8mb4_unicode_ci COMMENT='数据表备注';
```

### 执行程序
```bash
# 启动端口 :12345
kratos run
```

### 最后

基于以上步骤，修改相关文件(proto、service、biz、data)，基于自己的业务逻辑填充数据库字段，进行crud