# up-admin-server

> Gin、Gorm 后台服务

## 安装

1. 参考如下格式配置`MysqlDSN`环境变量（必要）

```text
user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
```

2. 根据自己的Redis环境，修改`define.go`中的Redis配置（必要）
