# 说明
这是一个 go web 的脚手架项目，用到的技术栈有：golang、gin、cobra、zap、mysql(gorm)、redis、kafka

# 代码结构
```bash
go-web-scaffold/
├── cmd/
│   ├── root.go
│   ├── server1.go
│   └── server2.go
├── internal/
│   ├── config/
│   │   └── config.go
│   ├── database/
│   │   ├── database.go
│   │   └── redis.go
│   ├── kafka/
│   │   └── kafka.go
│   ├── logging/
│   │   └── logging.go
│   └── handlers/
│       └── handlers.go
├── main.go
└── go.mod
```
