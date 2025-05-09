package config

import (
	"github.com/redis/go-redis/v9"
)

// RedisConfig 定义Redis配置结构体
type RedisConfig struct {
	Addr     string `yaml:"addr"`
	Password string `yaml:"password"`
	DB       int    `yaml:"db"`
}

// Redis客户端
var rdb *redis.Client

// InitRedis 从配置文件读取配置并初始化Redis客户端
func InitRedis() (*redis.Client, error) {
	// 调用 GetConfig 函数获取配置
	cfg, err := GetConfig()
	if err != nil {
		return nil, err
	}

	// 初始化Redis客户端
	rdb := redis.NewClient(&redis.Options{
		Addr:     cfg.Redis.Addr,
		Password: cfg.Redis.Password,
		DB:       cfg.Redis.DB,
	})

	return rdb, nil
}

func GetRedisClient() *redis.Client {
	return rdb
}
