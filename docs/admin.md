# kratos-base-project管理后台接口文档

## 调用方式
使用http请求进行API接口调用，返回数据给是为json格式，对于POST请求传输application/json格式的数据。

### url格式
url的格式为：
```
{{schema}}://{{host}}/api/admin/{{version}}/{{api-name}}
```
schema、host、port和version是可以根据部署环境进行变化的，api-name就是最终索要调用的api的名称。他们代表的含义为：
- **schema** ： 目前是 http
- **host** ： 目前测试环境的host是kratos.niu12.com
- **verson** ：目前的版本是v1
- **api-name** : 要调用的api的名称

### 成功返回格式
```
{
    "code": 0,
    "message": "success",
    "data": {
        
    }
}
```

### 失败返回格式
```
{
    "code": 20002,
    "reason": "ADMINISTRATOR_PASSWORD_ERROR",
    "message": "管理员密码错误",
    "metadata": {}
}
```

## 返回参数解释
| key      | 含义            |
|----------|---------------|
| code     | 返回码 0成功  其他失败 |
| data     | 数据            |
| message  | 提示信息          |
| reason   | 错误原因          |
| metadata | 元数据           |

## 基础描述
测试环境:

url path
```
https://kratos.niu12.com/prod-api/
```

示例： 
```
https://kratos.niu12.com/prod-api/admin/v1/login
```

## 目录
- 管理员
	- [登陆](#登陆)
	- [登陆成功](#登陆成功)
	- [登陆退出](#登陆退出)
	- [当前登录管理员详情](#当前登录管理员详情)
	- [管理员列表](#管理员列表)
    - [管理员详情](#管理员详情)
    - [管理员创建](#管理员创建)
    - [管理员更新](#管理员更新)
    - [管理员删除](#管理员删除)
- 角色
	- [角色列表](#角色列表)
	- [角色创建](#角色创建)
	- [角色更新](#角色更新)
	- [角色删除](#角色删除)
- API
	- [API全部列表](#API全部列表)
	- [API列表](#API列表)
	- [API创建](#API创建)
	- [API更新](#API更新)
	- [API删除](#API删除)
- 菜单
	- [菜单列表](#菜单列表)
	- [菜单列表(树)](#菜单列表(树))
	- [菜单创建](#菜单创建)
	- [菜单更新](#菜单更新)
	- [菜单删除](#菜单删除)
- 角色菜单设置
	- [角色菜单列表](#菜单列表)
	- [角色菜单列表(树)](#菜单列表(树))
	- [角色菜单设置](#角色菜单设置)
- 角色权限设置
  - [角色策略列表](#角色策略列表)
  - [角色策略设置](#角色策略设置)
- 角色用户设置
	- [用户角色列表](#用户角色列表)
	- [角色用户列表](#角色用户列表)
	- [用户角色设置](#用户角色设置)
	- [用户角色删除](#用户角色删除)
	- [用户角色清空](#用户角色清空)
## 登陆

### 接口描述

管理员账号登录管理后台

##### 需要token ： false


### 参数说明

| 参数       | 类型     | 必须   | 说明  |
|----------|--------|------|-----|
| username | string | true | 账号  |
| password | string | true | 密码  |


### 请求示例 

- url:

```
POST {{host}}/admin/v1/login
```

-  请求json
```
{
	"username":"admin",
	"password":"123456"
}
```

### 响应说明

-  返回json

```json
{
    "code": "200",
    "data": {
        "token": "9a4d3fa90d0a637114168d8ea41ff3d7"
    },
    "feedback": {
        "detail": "",
        "message": "Success"
    },
    "time": 1623203898
}
```

-  返回数据说明

| 返回字段  | 类型  |说明  |
| -------- | ------------ | ----------- |
| token | string| 用户授权码|



## 登陆成功

### 接口描述

管理员成功登录管理后台后，记录登录信息

##### 需要token ： true

### 请求示例

- url:

```
POST {{host}}/admin/v1/loginSuccess
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {}
}
```


## 登陆退出

### 接口描述

管理员退出登录

##### 需要token ： true

### 请求示例

- url:

```
POST {{host}}/admin/v1/logout
```


### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {}
}
```

## 当前登录管理员详情

### 接口描述

当前登录的管理员详情信息

##### 需要token ： true

### 请求示例

- url:

```
GET {{host}}/admin/v1/getAdministratorInfo
```


### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"avatar": "https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg",
		"createdAt": "2022-08-19 00:54:16",
		"deletedAt": "",
		"id": "19",
		"lastLoginIp": "",
		"lastLoginTime": "",
		"mobile": "18158445330",
		"nickname": "18158445330",
		"role": "普通管理员",
		"status": "1",
		"updatedAt": "2022-08-19 00:54:16",
		"username": "zhouqi1"
	}
}
```

-  返回数据说明

| 返回字段  | 类型  | 说明  |
| ------------ | ------------ | ------------ |
| avatar | string | 头像url      |
| createdAt | string | 创建时间       |
| deletedAt | string | 删除时间 未删除为空 |
| id | int64  | 主键id       |
| lastLoginIp | string  | 上次登录ip     |
| lastLoginTime | string  | 上次登录时间     |
| mobile | string  | 电话         |
| nickname | string  | 昵称         |
| role | string  | 当前角色       |
| status | string  | 状态 1正常 2冻结 |
| updatedAt | string | 更新时间       |
| username | string | 管理员账号      |




## 管理员列表


##### 需要token ： true

### 请求示例

- url:

```
GET {{host}}/admin/v1/getAdministrators
```

### 参数说明

| 参数               | 类型     | 必须    | 说明           |
|------------------|--------|-------|--------------|
| page          | int | false | 页码           |
| pageSize         | int | false | 页记录数         |
| mobile           | string | false | 手机号          |
| username         | string | false | 用户名          |
| nickname         | string | false | 昵称           |
| status           | int    | false| 用户状态 1正常 2冻结 |
| cerated_at_start | string    | false| 创建开始时间 Y-m-d |
| cerated_at_end   | string    | false| 创建结束时间 Y-m-d |

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"list": [
			{
				"avatar": "https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg",
				"createdAt": "2022-08-17 16:15:17",
				"deletedAt": "",
				"id": "18",
				"lastLoginIp": "127.0.0.1:61113",
				"lastLoginTime": "2022-11-10 17:46:50",
				"mobile": "18158445331",
				"nickname": "卡牌",
				"role": "超级管理员",
				"status": "1",
				"updatedAt": "2022-08-17 16:15:17",
				"username": "zhouqi"
			}
		],
		"total": "3"
	}
}
```

-  返回数据说明

| 返回字段          | 类型     | 说明         |
|---------------|--------|------------|
| list.avatar   | string | 头像url      |
| list.createdAt     | string | 创建时间       |
| list.deletedAt     | string | 删除时间 未删除为空 |
| list.id            | int64  | 主键id       |
| list.lastLoginIp   | string  | 上次登录ip     |
| list.lastLoginTime | string  | 上次登录时间     |
| list.mobile        | string  | 电话         |
| list.nickname      | string  | 昵称         |
| list.role          | string  | 当前角色       |
| list.status        | int64  | 状态 1正常 2冻结 |
| list.updatedAt     | string | 更新时间       |
| list.username      | string | 管理员账号      |
| total      | int64 | 总数         |


## 管理员详情

### 接口描述

管理员详情信息

##### 需要token ： true

### 请求示例

- url:

```
GET {{host}}/admin/v1/getAdministrator?id=18
```

### 参数说明

| 参数               | 类型     | 必须   | 说明   |
|------------------|--------|------|------|
| id          | int | true | 记录id |

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"avatar": "https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg",
		"createdAt": "2022-08-19 00:54:16",
		"deletedAt": "",
		"id": "19",
		"lastLoginIp": "",
		"lastLoginTime": "",
		"mobile": "18158445330",
		"nickname": "18158445330",
		"role": "普通管理员",
		"status": "1",
		"updatedAt": "2022-08-19 00:54:16",
		"username": "zhouqi1"
	}
}
```

-  返回数据说明

| 返回字段  | 类型  | 说明  |
| ------------ | ------------ | ------------ |
| avatar | string | 头像url      |
| createdAt | string | 创建时间       |
| deletedAt | string | 删除时间 未删除为空 |
| id | int64  | 主键id       |
| lastLoginIp | string  | 上次登录ip     |
| lastLoginTime | string  | 上次登录时间     |
| mobile | string  | 电话         |
| nickname | string  | 昵称         |
| role | string  | 当前角色       |
| status | string  | 状态 1正常 2冻结 |
| updatedAt | string | 更新时间       |
| username | string | 管理员账号      |




## 管理员创建

### 接口描述

管理员创建

##### 需要token ： true


### 参数说明

| 参数  | 类型    | 必须         | 说明    |
| ------------  | ------------ |------------|-------|
| username | string | true| 管理员账号 |
| password | string | true| 登录密码       |
| avatar | string | true| 头像url      |
| mobile | string  | true| 电话         |
| nickname | string  | true| 昵称         |
| status | string  | true| 状态 1正常 2冻结 |


### 请求示例

- url:

```
POST {{host}}/admin/v1/administrator
```

-  请求json
```
{
    "username": "zhouqi3111",
    "password": "123456",
    "mobile": "18158445363",
    "nickname": "卡牌",
    "avatar": "https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg",
    "status": "1"
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"avatar": "https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg",
		"createdAt": "2022-11-10 23:41:20",
		"deletedAt": "",
		"id": "22",
		"lastLoginIp": "",
		"lastLoginTime": "",
		"mobile": "18158445363",
		"nickname": "卡牌",
		"role": "",
		"status": "1",
		"updatedAt": "2022-11-10 23:41:20",
		"username": "zhouqi3111"
	}
}
```

-  返回数据说明

| 返回字段  | 类型  | 说明  |
| ------------ | ------------ | ------------ |
| avatar | string | 头像url      |
| createdAt | string | 创建时间       |
| deletedAt | string | 删除时间 未删除为空 |
| id | int64  | 主键id       |
| lastLoginIp | string  | 上次登录ip     |
| lastLoginTime | string  | 上次登录时间     |
| mobile | string  | 电话         |
| nickname | string  | 昵称         |
| role | string  | 当前角色       |
| status | string  | 状态 1正常 2冻结 |
| updatedAt | string | 更新时间       |
| username | string | 管理员账号      |


## 管理员更新

### 接口描述

管理员更新

##### 需要token ： true


### 参数说明

| 参数       | 类型     | 必须         | 说明         |
|----------|--------|------------|------------|
| id       | int64  | true| 管理员id      |
| username | string | true| 管理员账号      |
| password | string | true| 登录密码       |
| avatar   | string | true| 头像url      |
| mobile   | string | true| 电话         |
| nickname | string | true| 昵称         |
| status   | string | true| 状态 1正常 2冻结 |


### 请求示例

- url:

```
PUT {{host}}/admin/v1/administrator
```

-  请求json
```
{
	"id": "22",
    "username": "zhouqi3111",
    "password": "123456",
    "mobile": "18158445363",
    "nickname": "卡牌",
    "avatar": "https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg",
    "status": "1"
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"avatar": "https://kratos-base-project.oss-cn-hangzhou.aliyuncs.com/3441660123117_.pic.jpg",
		"createdAt": "2022-11-10 23:41:20",
		"deletedAt": "",
		"id": "22",
		"lastLoginIp": "",
		"lastLoginTime": "",
		"mobile": "18158445363",
		"nickname": "卡牌",
		"role": "",
		"status": "1",
		"updatedAt": "2022-11-10 23:41:20",
		"username": "zhouqi3111"
	}
}
```

-  返回数据说明

| 返回字段  | 类型  | 说明  |
| ------------ | ------------ | ------------ |
| avatar | string | 头像url      |
| createdAt | string | 创建时间       |
| deletedAt | string | 删除时间 未删除为空 |
| id | int64  | 主键id       |
| lastLoginIp | string  | 上次登录ip     |
| lastLoginTime | string  | 上次登录时间     |
| mobile | string  | 电话         |
| nickname | string  | 昵称         |
| role | string  | 当前角色       |
| status | string  | 状态 1正常 2冻结 |
| updatedAt | string | 更新时间       |
| username | string | 管理员账号      |



## 管理员删除

### 接口描述

管理员删除

##### 需要token ： true

### 参数说明

| 参数               | 类型     | 必须   | 说明   |
|------------------|--------|------|------|
| id          | int | true | 记录id |

### 请求示例

- url:

```
DELETE {{host}}/admin/v1/administrator?id=18
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {}
}
```




## 角色列表

### 接口描述

角色列表(不分页)

##### 需要token ： true

### 请求示例

- url:

```
GET {{host}}/admin/v1/role
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"list": [
			{
				"children": [],
				"createdAt": "2022-09-07 01:15:45",
				"id": "5",
				"name": "普通管理员",
				"parentId": "0",
				"updatedAt": "2022-09-07 01:27:17"
			}
		]
	}
}
```

-  返回数据说明

| 返回字段           | 类型     | 说明    |
|----------------|--------|-------|
| list.children  | []role | 子角色信息 |
| list.createdAt | string | 创建时间  |
| list.id        | int64  | 主键id  |
| list.name  | int64  | 角色名称  |
| list.parentId  | int64  | 父角色id |
| list.updatedAt | string | 更新时间  |

## 角色创建

### 接口描述

角色创建

##### 需要token ： true


### 参数说明

| 参数  | 类型    | 必须        | 说明    |
| ------------  | ------------ |-----------|-------|
| name  | int64  | true| 角色名称      |
| parentId  | int64  | true| 父角色id     |


### 请求示例

- url:

```
POST {{host}}/admin/v1/role
```

-  请求json
```
{
    "parent_id": 1,
    "name": "管理员"
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"children": [],
		"createdAt": "2022-11-11 00:25:40",
		"id": "20",
		"name": "管理员",
		"parentId": "1",
		"updatedAt": "2022-11-11 00:25:40"
	}
}
```

-  返回数据说明

| 返回字段           | 类型  | 说明  |
|----------------|--------|-------|
| children       | []role | 子角色信息 |
| createdAt      | string | 创建时间  |
| id             | int64  | 主键id  |
| name           | int64  | 角色名称  |
| parentId       | int64  | 父角色id |
| updatedAt | string | 更新时间  |


## 角色更新

### 接口描述

角色更新

##### 需要token ： true


### 参数说明

| 参数       | 类型     | 必须         | 说明         |
|----------|--------|------------|------------|
| id       | int64  | true| 角色id      |
| name  | int64  | true| 角色名称      |
| parentId  | int64  | true| 父角色id     |


### 请求示例

- url:

```
PUT {{host}}/admin/v1/role
```

-  请求json
```
{
	"id":1,
    "parent_id": 1,
    "name": "管理员"
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"children": [],
		"createdAt": "2022-11-11 00:25:40",
		"id": "20",
		"name": "管理员",
		"parentId": "1",
		"updatedAt": "2022-11-11 00:25:40"
	}
}
```

-  返回数据说明

| 返回字段  | 类型  | 说明  |
|----------------|--------|-------|
| children       | []role | 子角色信息 |
| createdAt      | string | 创建时间  |
| id             | int64  | 主键id  |
| name           | int64  | 角色名称  |
| parentId       | int64  | 父角色id |
| updatedAt | string | 更新时间  |

## 角色删除

### 接口描述

角色删除

##### 需要token ： true

### 参数说明

| 参数               | 类型     | 必须   | 说明   |
|------------------|--------|------|------|
| id          | int | true | 记录id |

### 请求示例

- url:

```
DELETE {{host}}/admin/v1/role?id=18
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {}
}
```





## API全部列表

### 接口描述

API全部列表 不分页

##### 需要token ： true

### 请求示例

- url:

```
GET {{host}}/admin/v1/apiAll
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"list": [
			{
				"createdAt": "2022-09-19 17:27:45",
				"group": "测2试",
				"id": "16",
				"method": "POST",
				"name": "测1试",
				"path": "1",
				"updatedAt": "2022-09-19 17:37:01"
			}
		]
	}
}
```

-  返回数据说明

| 返回字段           | 类型     | 说明                       |
|----------------|--------|--------------------------|
| list.createdAt | string | 创建时间                     |
| list.group  | string  | 分组名称                     |
| list.id        | int64  | 主键id                     |
| list.method  | string  | 请求方式 POST、GET、DELETE、PUT |
| list.name  | string  | API名称                    |
| list.path  | string  | 路径                       |
| list.updatedAt | string | 更新时间                     |



## API列表

### 接口描述

API列表 分页

##### 需要token ： true

### 参数说明

| 参数               | 类型     | 必须    | 说明    |
|------------------|--------|-------|-------|
| page          | int | false | 页码    |
| pageSize         | int | false | 页记录数  |
| group           | string | false | 分组名称  |
| name           | string | false | api名称 |
| path           | string | false | 请求路径  |
| method           | string | false | 请求方式  |

### 请求示例

- url:

```
GET {{host}}/admin/v1/api
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"list": [
			{
				"createdAt": "2022-09-19 17:27:45",
				"group": "测2试",
				"id": "16",
				"method": "POST",
				"name": "测1试",
				"path": "1",
				"updatedAt": "2022-09-19 17:37:01"
			}
		],
		"total": "1"
	}
}
```

-  返回数据说明

| 返回字段           | 类型     | 说明                       |
|----------------|--------|--------------------------|
| list.createdAt | string | 创建时间                     |
| list.group     | string  | 分组名称                     |
| list.id        | int64  | 主键id                     |
| list.method    | string  | 请求方式 POST、GET、DELETE、PUT |
| list.name      | string  | API名称                    |
| list.path      | string  | 路径                       |
| list.updatedAt | string | 更新时间                     |
| total          | int64 | 总数                       |


## API创建

### 接口描述

API创建

##### 需要token ： true


### 参数说明

| 参数  | 类型    | 必须                       | 说明   |
| ------------  | ------------ |--------------------------|------|
| group  | string  | true| 分组名称 |
| method  | string  | true| 请求方式 POST、GET、DELETE、PUT |
| name  | string  | true| API名称                    |
| path  | string  | true| 路径                       |


### 请求示例

- url:

```
POST {{host}}/admin/v1/api
```

-  请求json
```
{
    "group":"测2试1",
    "name":"测1试",
    "path":"1",
    "method":"POST",
    "created_at":"2022-09-19 17:27:45",
    "updated_at":"2022-09-19 17:37:01"
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"createdAt": "2022-11-11 00:45:21",
		"group": "测2试1",
		"id": "20",
		"method": "POST",
		"name": "测1试",
		"path": "1",
		"updatedAt": "2022-11-11 00:45:21"
	}
}
```

-  返回数据说明

| 返回字段           | 类型  | 说明  |
|----------------|--------|-------|
| createdAt | string | 创建时间                     |
| group  | string  | 分组名称                     |
| id        | int64  | 主键id                     |
| method  | string  | 请求方式 POST、GET、DELETE、PUT |
| name  | string  | API名称                    |
| path  | string  | 路径                       |
| updatedAt | string | 更新时间                     |


## API更新

### 接口描述

API更新

##### 需要token ： true


### 参数说明

| 参数       | 类型     | 必须         | 说明         |
|----------|--------|------------|------------|
| id       | int64  | true| APIid      |
| group  | string  | true| 分组名称 |
| method  | string  | true| 请求方式 POST、GET、DELETE、PUT |
| name  | string  | true| API名称                    |
| path  | string  | true| 路径                       |


### 请求示例

- url:

```
PUT {{host}}/admin/v1/api
```

-  请求json
```
{
    "id": 19,
    "group":"测2试111",
    "name":"测1试",
    "path":"1",
    "method":"POST",
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"createdAt": "2022-11-11 00:45:22",
		"group": "测2试111",
		"id": "20",
		"method": "POST",
		"name": "测1试",
		"path": "1",
		"updatedAt": "2022-11-11 00:47:25"
	}
}
```

-  返回数据说明

| 返回字段  | 类型  | 说明  |
|----------------|--------|-------|
| createdAt | string | 创建时间                     |
| group  | string  | 分组名称                     |
| id        | int64  | 主键id                     |
| method  | string  | 请求方式 POST、GET、DELETE、PUT |
| name  | string  | API名称                    |
| path  | string  | 路径                       |
| updatedAt | string | 更新时间                     |

## API删除

### 接口描述

API删除

##### 需要token ： true

### 参数说明

| 参数               | 类型     | 必须   | 说明   |
|------------------|--------|------|------|
| id          | int | true | 记录id |

### 请求示例

- url:

```
DELETE {{host}}/admin/v1/role?id=18
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {}
}
```





## 菜单列表

### 接口描述

所有菜单列表 非树状，不含子对象

##### 需要token ： true

### 请求示例

- url:

```
GET {{host}}/admin/v1/menuAll
```

### 响应说明

-  返回json

```json
{
	"code":0,
	"message":"success",
	"data":{
		"list":[
			{
				"children":[

				],
				"component":"view/system/index.vue",
				"createdAt":"2022-09-28 11:34:51",
				"hidden":"0",
				"icon":"system",
				"id":"13",
				"menuBtns":[
					{
						"createdAt":"2022-09-28 11:34:51",
						"description":"描述1",
						"id":"14",
						"menuId":"13",
						"name":"创建12",
						"updatedAt":"2022-09-28 11:34:51"
					}
				],
				"name":"dashborard",
				"parentId":"0",
				"path":"",
				"sort":"1",
				"title":"首页",
				"updatedAt":"2022-09-28 11:34:51"
			}
		]
	}
}
```

-  返回数据说明

| 返回字段                     | 类型     | 说明       |
|--------------------------|--------|----------|
| list.children            | []menu | 菜单子对象    |
| list.component           | string | 组件地址     |
| list.createdAt           | string | 创建时间     |
| list.hidden              | int64  | 是否隐藏1是0否 |
| list.icon                | string | icon     |
| list.id                  | int64 | 菜单id     |
| list.name                  | string | 菜单名称     |
| list.parentId                  | string | 父菜单id    |
| list.path                  | string | 菜单path   |
| list.sort                  | string | 排序       |
| list.title                  | string | 菜单title  |
| list.updatedAt           | string | 更新时间     |
| list.menuBtns            | []menuBtn  | 菜单按钮集合   |
| list.menuBtn.createdAt   | string | 创建时间     |
| list.menuBtn.description | string | 按钮描述     |
| list.menuBtn.id          | int64 | 按钮id     |
| list.menuBtn.menuId      | int64 | 菜单id     |
| list.menuBtn.name        | string | 按钮名称     |
| list.menuBtn.updatedAt   | int64 | 更新时间     |


## 菜单列表(树)

### 接口描述

菜单列表 树状结构

##### 需要token ： true

### 请求示例

- url:

```
GET {{host}}/admin/v1/menuTree
```

### 响应说明

-  返回json

```json
{
	"code":0,
	"message":"success",
	"data":{
		"list":[
			{
				"children":[
					{
						"children":[

						],
						"component":"view/system/administrator/index.vue",
						"createdAt":"2022-09-16 16:16:10",
						"hidden":"0",
						"icon":"user",
						"id":"10",
						"menuBtns":[
							{
								"createdAt":"2022-09-18 23:28:42",
								"description":"描述1",
								"id":"1",
								"menuId":"10",
								"name":"创建12",
								"updatedAt":"2022-09-18 23:28:42"
							}
						],
						"name":"administrator",
						"parentId":"9",
						"path":"1",
						"sort":"2",
						"title":"管理员管理",
						"updatedAt":"2022-09-18 23:28:42"
					}
				],
				"component":"view/system/index.vue",
				"createdAt":"2022-09-16 15:40:58",
				"hidden":"0",
				"icon":"system",
				"id":"9",
				"menuBtns":[

				],
				"name":"system",
				"parentId":"0",
				"path":"system",
				"sort":"1",
				"title":"系统管理",
				"updatedAt":"2022-09-16 15:40:58"
			}
		]
	}
}
```

-  返回数据说明

| 返回字段                     | 类型     | 说明       |
|--------------------------|--------|----------|
| list.children            | []menu | 菜单子对象    |
| list.component           | string | 组件地址     |
| list.createdAt           | string | 创建时间     |
| list.hidden              | int64  | 是否隐藏1是0否 |
| list.icon                | string | icon     |
| list.id                  | int64 | 菜单id     |
| list.name                  | string | 菜单名称     |
| list.parentId                  | string | 父菜单id    |
| list.path                  | string | 菜单path   |
| list.sort                  | string | 排序       |
| list.title                  | string | 菜单title  |
| list.updatedAt           | string | 更新时间     |
| list.menuBtns            | []menuBtn  | 菜单按钮集合   |
| list.menuBtn.createdAt   | string | 创建时间     |
| list.menuBtn.description | string | 按钮描述     |
| list.menuBtn.id          | int64 | 按钮id     |
| list.menuBtn.menuId      | int64 | 菜单id     |
| list.menuBtn.name        | string | 按钮名称     |
| list.menuBtn.updatedAt   | int64 | 更新时间     |


## 菜单创建

### 接口描述

菜单创建

##### 需要token ： true


### 参数说明

| 参数  | 类型    | 必须                       | 说明   |
| ------------  | ------------ |--------------------------|------|
| component           | string | true|组件地址     |
| hidden              | int64  | true|是否隐藏1是0否 |
| icon                | string | true|icon     |
| name                  | string |true| 菜单名称     |
| parentId                  | string | true| 父菜单id    |
| path                  | string |true|  菜单path   |
| sort                  | string | true| 排序       |
| title                  | string | true| 菜单title  |
| menuBtns            | []menuBtn  | 菜单按钮集合   |
| menuBtn.description | string | 按钮描述     |
| menuBtn.id          | int64 | 按钮id     |
| menuBtn.menuId      | int64 | 菜单id     |
| menuBtn.name        | string | 按钮名称     |


### 请求示例

- url:

```
POST {{host}}/admin/v1/menu
```

-  请求json
```
{
    "component":"view/system/index.vue",
    "hidden":"0",
    "icon":"system",
    "menuBtns":[
        {
            "description":"描述1",
            "menuId":"10",
            "name":"创建12"
        }
    ],
    "name":"dashborard",
    "parentId":"0",
    "path":"1232",
    "sort":"1",
    "title":"首页"
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"children": [],
		"component": "view/system/index.vue",
		"createdAt": "",
		"hidden": "0",
		"icon": "system",
		"id": "14",
		"menuBtns": [
			{
				"createdAt": "2022-11-11 09:56:25",
				"description": "描述1",
				"id": "15",
				"menuId": "14",
				"name": "创建12",
				"updatedAt": "2022-11-11 09:56:25"
			}
		],
		"name": "dashborard",
		"parentId": "0",
		"path": "1232",
		"sort": "1",
		"title": "首页",
		"updatedAt": ""
	}
}
```

-  返回数据说明

| 返回字段           | 类型  | 说明  |
|----------------|--------|-------|
| children            | []menu | 菜单子对象    |
| component           | string | 组件地址     |
| createdAt           | string | 创建时间     |
| hidden              | int64  | 是否隐藏1是0否 |
| icon                | string | icon     |
| id                  | int64 | 菜单id     |
| name                  | string | 菜单名称     |
| parentId                  | string | 父菜单id    |
| path                  | string | 菜单path   |
| sort                  | string | 排序       |
| title                  | string | 菜单title  |
| updatedAt           | string | 更新时间     |
| menuBtns            | []menuBtn  | 菜单按钮集合   |
| menuBtn.createdAt   | string | 创建时间     |
| menuBtn.description | string | 按钮描述     |
| menuBtn.id          | int64 | 按钮id     |
| menuBtn.menuId      | int64 | 菜单id     |
| menuBtn.name        | string | 按钮名称     |
| menuBtn.updatedAt   | int64 | 更新时间     |


## 菜单更新

### 接口描述

菜单更新

##### 需要token ： true


### 参数说明

| 参数       | 类型     | 必须         | 说明         |
|----------|--------|------------|------------|
| id       | int64  | true| 菜单id      |


### 参数说明

| 参数  | 类型    | 必须                       | 说明   |
| ------------  | ------------ |--------------------------|------|
| component           | string | true|组件地址     |
| hidden              | int64  | true|是否隐藏1是0否 |
| icon                | string | true|icon     |
| name                  | string |true| 菜单名称     |
| parentId                  | string | true| 父菜单id    |
| path                  | string |true|  菜单path   |
| sort                  | string | true| 排序       |
| title                  | string | true| 菜单title  |
| menuBtns            | []menuBtn  | 菜单按钮集合   |
| menuBtn.description | string | 按钮描述     |
| menuBtn.id          | int64 | 按钮id     |
| menuBtn.menuId      | int64 | 菜单id     |
| menuBtn.name        | string | 按钮名称     |


### 请求示例

- url:

```
PUT {{host}}/admin/v1/menu
```

-  请求json
```
{
        "children": [],
        "component": "view/system/index.vue",
        "createdAt": "",
        "hidden": "0",
        "icon": "system",
        "id": "12",
        "menuBtns": [
            {
                "createdAt": "2022-09-28 11:15:14",
                "description": "描述1",
                "id": "13",
                "menuId": "12",
                "name": "创建12",
                "updatedAt": "2022-09-28 11:15:14"
            }
        ],
        "name": "dashborard",
        "parentId": "0",
        "path": "1232",
        "sort": "1",
        "title": "首页22",
        "updatedAt": ""
    }
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"children": [],
		"component": "view/system/index.vue",
		"createdAt": "",
		"hidden": "0",
		"icon": "system",
		"id": "14",
		"menuBtns": [
			{
				"createdAt": "2022-11-11 09:56:25",
				"description": "描述1",
				"id": "15",
				"menuId": "14",
				"name": "创建12",
				"updatedAt": "2022-11-11 09:56:25"
			}
		],
		"name": "dashborard",
		"parentId": "0",
		"path": "1232",
		"sort": "1",
		"title": "首页",
		"updatedAt": ""
	}
}
```

-  返回数据说明

| 返回字段           | 类型  | 说明  |
|----------------|--------|-------|
| children            | []menu | 菜单子对象    |
| component           | string | 组件地址     |
| createdAt           | string | 创建时间     |
| hidden              | int64  | 是否隐藏1是0否 |
| icon                | string | icon     |
| id                  | int64 | 菜单id     |
| name                  | string | 菜单名称     |
| parentId                  | string | 父菜单id    |
| path                  | string | 菜单path   |
| sort                  | string | 排序       |
| title                  | string | 菜单title  |
| updatedAt           | string | 更新时间     |
| menuBtns            | []menuBtn  | 菜单按钮集合   |
| menuBtn.createdAt   | string | 创建时间     |
| menuBtn.description | string | 按钮描述     |
| menuBtn.id          | int64 | 按钮id     |
| menuBtn.menuId      | int64 | 菜单id     |
| menuBtn.name        | string | 按钮名称     |
| menuBtn.updatedAt   | int64 | 更新时间     |

## 菜单删除

### 接口描述

菜单删除

##### 需要token ： true

### 参数说明

| 参数               | 类型     | 必须   | 说明   |
|------------------|--------|------|------|
| id          | int | true | 记录id |

### 请求示例

- url:

```
DELETE {{host}}/admin/v1/menu?id=18
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {}
}
```



## 角色菜单列表

### 接口描述

某个角色拥有的所有菜单列表 非树状，不含子对象

##### 需要token ： true

### 参数说明

| 参数               | 类型     | 必须   | 说明   |
|------------------|--------|------|------|
| id          | int | true | 记录id |

### 请求示例

- url:

```
GET {{host}}/admin/v1/roleMenu?role_id=1
```

### 响应说明

-  返回json

```json
{
	"code":0,
	"message":"success",
	"data":{
		"list":[
			{
				"children":[

				],
				"component":"view/system/index.vue",
				"createdAt":"2022-09-28 11:34:51",
				"hidden":"0",
				"icon":"system",
				"id":"13",
				"menuBtns":[
					{
						"createdAt":"2022-09-28 11:34:51",
						"description":"描述1",
						"id":"14",
						"menuId":"13",
						"name":"创建12",
						"updatedAt":"2022-09-28 11:34:51"
					}
				],
				"name":"dashborard",
				"parentId":"0",
				"path":"",
				"sort":"1",
				"title":"首页",
				"updatedAt":"2022-09-28 11:34:51"
			}
		]
	}
}
```

-  返回数据说明

| 返回字段                     | 类型     | 说明       |
|--------------------------|--------|----------|
| list.children            | []menu | 菜单子对象    |
| list.component           | string | 组件地址     |
| list.createdAt           | string | 创建时间     |
| list.hidden              | int64  | 是否隐藏1是0否 |
| list.icon                | string | icon     |
| list.id                  | int64 | 菜单id     |
| list.name                  | string | 菜单名称     |
| list.parentId                  | string | 父菜单id    |
| list.path                  | string | 菜单path   |
| list.sort                  | string | 排序       |
| list.title                  | string | 菜单title  |
| list.updatedAt           | string | 更新时间     |
| list.menuBtns            | []menuBtn  | 菜单按钮集合   |
| list.menuBtn.createdAt   | string | 创建时间     |
| list.menuBtn.description | string | 按钮描述     |
| list.menuBtn.id          | int64 | 按钮id     |
| list.menuBtn.menuId      | int64 | 菜单id     |
| list.menuBtn.name        | string | 按钮名称     |
| list.menuBtn.updatedAt   | int64 | 更新时间     |


## 角色菜单列表(树)

### 接口描述

某个角色拥有的菜单列表 树状结构

##### 需要token ： true

### 参数说明

| 参数               | 类型     | 必须   | 说明   |
|------------------|--------|------|------|
| id          | int | true | 记录id |

### 请求示例

- url:

```
GET {{host}}/admin/v1/roleMenuTree?id=1
```

### 响应说明

-  返回json

```json
{
	"code":0,
	"message":"success",
	"data":{
		"list":[
			{
				"children":[
					{
						"children":[

						],
						"component":"view/system/administrator/index.vue",
						"createdAt":"2022-09-16 16:16:10",
						"hidden":"0",
						"icon":"user",
						"id":"10",
						"menuBtns":[
							{
								"createdAt":"2022-09-18 23:28:42",
								"description":"描述1",
								"id":"1",
								"menuId":"10",
								"name":"创建12",
								"updatedAt":"2022-09-18 23:28:42"
							}
						],
						"name":"administrator",
						"parentId":"9",
						"path":"1",
						"sort":"2",
						"title":"管理员管理",
						"updatedAt":"2022-09-18 23:28:42"
					}
				],
				"component":"view/system/index.vue",
				"createdAt":"2022-09-16 15:40:58",
				"hidden":"0",
				"icon":"system",
				"id":"9",
				"menuBtns":[

				],
				"name":"system",
				"parentId":"0",
				"path":"system",
				"sort":"1",
				"title":"系统管理",
				"updatedAt":"2022-09-16 15:40:58"
			}
		]
	}
}
```

-  返回数据说明

| 返回字段                     | 类型     | 说明       |
|--------------------------|--------|----------|
| list.children            | []menu | 菜单子对象    |
| list.component           | string | 组件地址     |
| list.createdAt           | string | 创建时间     |
| list.hidden              | int64  | 是否隐藏1是0否 |
| list.icon                | string | icon     |
| list.id                  | int64 | 菜单id     |
| list.name                  | string | 菜单名称     |
| list.parentId                  | string | 父菜单id    |
| list.path                  | string | 菜单path   |
| list.sort                  | string | 排序       |
| list.title                  | string | 菜单title  |
| list.updatedAt           | string | 更新时间     |
| list.menuBtns            | []menuBtn  | 菜单按钮集合   |
| list.menuBtn.createdAt   | string | 创建时间     |
| list.menuBtn.description | string | 按钮描述     |
| list.menuBtn.id          | int64 | 按钮id     |
| list.menuBtn.menuId      | int64 | 菜单id     |
| list.menuBtn.name        | string | 按钮名称     |
| list.menuBtn.updatedAt   | int64 | 更新时间     |



## 角色菜单设置

### 接口描述

设置角色拥有的菜单

##### 需要token ： true


### 参数说明

| 参数      | 类型    | 必须  | 说明 |
|---------|-------|--------------------|--------|
| role_id | int64 | true| 角色id |
| menu_ids | []int64 | true| 菜单id数组 |


### 请求示例

- url:

```
POST {{host}}/admin/v1/roleMenu
```

-  请求json
```
{
    "role_id":1,
    "menu_ids": [9, 10, 13]
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"isSuccess": true
	}
}
```


## 角色策略列表

### 接口描述

查看角色策略，查看拥有访问哪些接口的权限

##### 需要token ： true


### 参数说明

| 参数                 | 类型      | 必须  | 说明   |
|--------------------|---------|--------------------|------|
| role  | string  | true| 角色名称 |


### 请求示例

- url:

```
POST {{host}}/admin/v1/getPolicies?role=超级管理员
```

-  请求json
```
{
    "policyRules":[
        {
            "path":"/api.authorization.v1.Authorization/UpdatePolicies",
            "method":"POST"
        },
        {
            "path":"/api.authorization.v1.Authorization/GetRoleList",
            "method":"GET"
        }
    ],
    "role":"超级管理员"
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"isSuccess": true
	}
}
```


## 角色策略设置

### 接口描述

设置角色策略，设置拥有访问哪些接口的权限

##### 需要token ： true


### 参数说明

| 参数                 | 类型      | 必须  | 说明   |
|--------------------|---------|--------------------|------|
| role               | string  | true| 角色名称 |
| policyRules.path   | string  | true| 请求路径 |
| policyRules.method | string  | true| 请求方法 |


### 请求示例

- url:

```
POST {{host}}/admin/v1/updatePolicies
```

-  请求json
```
{
    "policyRules":[
        {
            "path":"/api.authorization.v1.Authorization/UpdatePolicies",
            "method":"POST"
        },
        {
            "path":"/api.authorization.v1.Authorization/GetRoleList",
            "method":"GET"
        }
    ],
    "role":"超级管理员"
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"isSuccess": true
	}
}
```




## 用户角色列表

### 接口描述

查看用户拥有的角色列表

##### 需要token ： true


### 参数说明

| 参数       | 类型       | 必须  | 说明     |
|----------|----------|--------------------|--------|
| username | string   | true| 账号名称   |


### 请求示例

- url:

```
GET {{host}}/admin/v1/getRolesForUser?username=zhouqi
```

-  请求json
```
{
    "code": 0,
    "message": "success",
    "data": {
        "roles": [
            "超级管理员"
        ]
    }
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"roles": [
			"超级管理员"
		]
	}
}
```


-  返回数据说明

| 返回字段          | 类型       | 说明     |
|---------------|----------|--------|
| roles | []string | 角色名称数组 |


## 角色用户列表

### 接口描述

查看角色拥有的用户列表

##### 需要token ： true


### 参数说明

| 参数       | 类型       | 必须  | 说明     |
|----------|----------|--------------------|--------|
| role     | string | true| 角色名称 |


### 请求示例

- url:

```
GET {{host}}/admin/v1/getUsersForRole?role=超级管理员
```

-  请求json
```
{
    "code": 0,
    "message": "success",
    "data": {
        "users": [
            "zhouqi"
        ]
    }
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"roles": [
			"超级管理员"
		]
	}
}
```


-  返回数据说明

| 返回字段          | 类型       | 说明     |
|---------------|----------|--------|
| users | []string | 用户名称数组 |


## 用户角色设置

### 接口描述

设置用户拥有的角色列表

##### 需要token ： true


### 参数说明

| 参数       | 类型       | 必须  | 说明     |
|----------|----------|--------------------|--------|
| username | string   | true| 账号名称   |
| role     | []string | true| 角色名称数组 |


### 请求示例

- url:

```
POST {{host}}/admin/v1/addRolesForUser
```

-  请求json
```
{
    "username": "zhouqi",
    "roles": ["超级管理员"]
}
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"isSuccess": true
	}
}
```


## 用户角色删除

### 接口描述

删除用户某个角色

##### 需要token ： true


### 参数说明

| 参数       | 类型       | 必须  | 说明     |
|----------|----------|--------------------|--------|
| username | string   | true| 账号名称   |
| role     | string | true| 角色名称 |


### 请求示例

- url:

```
DELETE {{host}}/admin/v1/deleteRoleForUser?username=zhouqi&role=超级管理员
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"isSuccess": true
	}
}
```



## 用户角色清空

### 接口描述

删除用户所有角色

##### 需要token ： true


### 参数说明

| 参数       | 类型       | 必须  | 说明     |
|----------|----------|--------------------|--------|
| username | string   | true| 账号名称   |
| role     | string | true| 角色名称 |


### 请求示例

- url:

```
DELETE {{host}}/admin/v1/deleteRolesForUser?username=zhouqi
```

### 响应说明

-  返回json

```json
{
	"code": 0,
	"message": "success",
	"data": {
		"isSuccess": true
	}
}
```
