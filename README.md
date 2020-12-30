## 项目功能概述
这个项目是用golang的gin框架写的一个server服务，功能是基本的账号注册、账号登陆、查看账号信息
、修改账号密码。(CURD boy) 

各个接口的请求的body或head示例如下：
```
 - using env:   export GIN_MODE=release
 - using code:  gin.SetMode(gin.ReleaseMode)

[GIN-debug] POST   /api/v1/auth/signup       --> server/handler/v1.Signup (3 handlers)
[GIN-debug] POST   /api/v1/auth/signin       --> server/handler/v1.Signin (3 handlers)
[GIN-debug] POST   /api/v1/user/getInfo      --> server/handler/v1.GetUserInfo (4 handlers)
[GIN-debug] PATCH  /api/v1/user/alterPwd     --> server/handler/v1.AlterUserPwd (4 handlers)
[GIN-debug] Listening and serving HTTP on :8080


signup接口:json：
{
    "name": "name",
    "email": "email",
    "password": "password",
    "re_password": "password"
}

signin接口：json：
{
  "name": "name",
  "password": "password"
}

getinfo接口：header需要：Authorization+token


alterpwd接口：header需要：Authorization+token
json：
{
    "old_password": "oldpwd",
    "new_password": "newpwd"
}
```

项目使用了以下几个库，通过本项目你可以学习以下几个库的基础使用：

* gorm(orm)
* viper(用于对配置文件的读、写、热操作)
* jwt-go(jwt鉴权)
* validata(接口数据有效性验证)

使用方法：config目录下有一个配置文件，你可以在里面更改你的mysql数据库配置信息，（自己需要新建一个数据库，无需建表，代码中已利用gorm库自动建表）。

如果对一些库的使用操作有疑问可以左转到 我的另一个仓库 `goStudyNote` 中寻找，库笔记一般放在playground这个文件夹下面。

项目的目录结构如下：

```
.
├── config
│   └── config.toml
├── dao
│   ├── auth.go
│   └── user.go
├── extend
│   ├── conf
│   │   └── viper.go
│   ├── standardCode
│   │   ├── codeFormatter.go
│   │   └── code.go
│   ├── token
│   │   └── jwt.go
│   ├── utils
│   │   └── utils.go
│   └── validata
│       └── validata.go
├── go.mod
├── go.sum
├── handler
│   └── v1
│       ├── auth.go
│       └── user.go
├── main.go
├── middleware
│   └── jwtAuth.go
├── model
│   ├── initDB.go
│   └── user.go
├── router
    └── router.go
```

## 结语
这个项目是用go写的第一个项目，代码手法啥都都很初级，很高兴你能看到这里，祝大家变得更强。🤞