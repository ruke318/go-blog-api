## 技术宅男子API  (go iris框架版本)

### 目录结构

```
├── README.md
├── articles //这是写的文章
│   └── 1.md
├── blog //编译后的可执行文件
├── config //配置目录 这个里面的json文件是不会编译进去的, 上传的收记得将配置文件一起传上去
│   ├── config.go
│   └── config.json
├── controllers //控制器目录
│   ├── controller.go //控制器的基础文件
│   ├── link.go
│   ├── posts.go
│   └── user.go
├── database //数据库目录, 包括 mysql redis之类的
│   └── database.go
├── main.go //入口文件
├── models //数据库表模型文件
│   ├── links.go
│   ├── model.go
│   ├── posts.go
│   └── users.go
├── routers //路由目录
│   └── dispath.go
└── tools //工具目录
    └── tools.go
```

### 使用了那些东西

```
iris
gorm
github.com/jinzhu/configor
github.com/olivere/elastic --elastsearch
github.com/thoas/go-funk  --帮助函数

```

### 大致介绍

这个是我自学`go`的第一个项目, 整个项目的大致流程是这样的

1. 加载配置文件
2. 连接数据库, es, redis
3. 路由
    1. 中间件
    2. 分发到`controllers`
    3. 找到对应的 `models`
    4. 数据库操作
    5. 返回json