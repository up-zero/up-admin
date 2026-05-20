# up-admin-server

> Gin、Gorm 后台服务

## 安装

1. 首次启动时，程序会自动初始化默认配置文件

2. 通过启动参数 `--config` 指定配置文件路径，例如：

```bash
go run . --config ./config.yaml
```

3. 打开配置文件，至少完善 `database.dsn` 配置

```text
user:pass@tcp(127.0.0.1:3306)/dbname?charset=utf8mb4&parseTime=True&loc=Local
```

4. 根据自己的环境按需修改 `redis`、`jwt`、`server`、`log` 等配置

## 配置说明

- 默认配置文件名：`config.yaml`
- 配置文件修改后会自动热重载

## 特殊错误码

+ 60403 (登录信息无效)
