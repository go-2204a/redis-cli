package redis_cli

import (
	"context"
	"github.com/go-redis/redis/v8"
)

// InitRedis 初始化并返回一个 Redis 客户端连接
//
// 参数:
// Addr:     Redis 服务器的地址 (例如: "localhost:6379")
// Password: Redis 连接的密码，如果没有密码则为 ""（空字符串）
// Db:       Redis 数据库的编号，通常为 0（默认数据库）
//
// 返回值:
// *redis.Client: 返回一个 Redis 客户端实例，用于与 Redis 进行交互。
// error: 如果连接失败，返回相应的错误信息。
//
// redis包:
// 使用的是v8版本
func InitRedis(Addr string, Password string, Db int) (*redis.Client, error) {
	Rdb := redis.NewClient(&redis.Options{
		Addr:     Addr,
		Password: Password,
		DB:       Db,
	})
	_, err := Rdb.Ping(context.Background()).Result()
	if err != nil {
		return nil, err
	}
	return Rdb, nil
}
