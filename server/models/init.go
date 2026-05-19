package models

import (
	"context"
	"fmt"
	"time"

	"gitee.com/up-zero/up-admin/define"
	"github.com/go-redis/redis/v8"
	"gorm.io/driver/mysql"
	"gorm.io/gorm"
)

var DB *gorm.DB

var RDB *redis.Client

func ReloadDatastores(cfg define.Config) error {
	newDB, err := openGormDB(cfg.Database.DSN)
	if err != nil {
		return err
	}
	newRDB, err := openRedisDB(cfg.Redis.Addr, cfg.Redis.Username, cfg.Redis.Password)
	if err != nil {
		closeGormDB(newDB)
		return err
	}

	oldDB := DB
	oldRDB := RDB
	DB = newDB
	RDB = newRDB
	closeGormDB(oldDB)
	closeRedisDB(oldRDB)
	return nil
}

func openGormDB(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(mysql.Open(dsn), &gorm.Config{
		DisableForeignKeyConstraintWhenMigrating: true,
	})
	if err != nil {
		return nil, fmt.Errorf("初始化 MySQL 连接失败: %w", err)
	}
	if err = db.AutoMigrate(&UserBasic{}, &RoleBasic{}, &RoleMenu{}, &RoleFunction{}, &MenuBasic{}, &FunctionBasic{}); err != nil {
		closeGormDB(db)
		return nil, fmt.Errorf("执行 MySQL 自动迁移失败: %w", err)
	}
	return db, nil
}

func openRedisDB(addr, username, password string) (*redis.Client, error) {
	client := redis.NewClient(&redis.Options{
		Addr:     addr,
		Username: username,
		Password: password,
	})
	ctx, cancel := context.WithTimeout(context.Background(), 5*time.Second)
	defer cancel()
	if err := client.Ping(ctx).Err(); err != nil {
		_ = client.Close()
		return nil, fmt.Errorf("初始化 Redis 连接失败: %w", err)
	}
	return client, nil
}

func closeGormDB(db *gorm.DB) {
	if db == nil {
		return
	}
	sqlDB, err := db.DB()
	if err != nil {
		return
	}
	_ = sqlDB.Close()
}

func closeRedisDB(client *redis.Client) {
	if client == nil {
		return
	}
	_ = client.Close()
}
