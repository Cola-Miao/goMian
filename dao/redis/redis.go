package redis

import (
	"context"
	"github.com/redis/go-redis/v9"
	"goMian/config"
)

var DB = new(redisDB)

type redisDB struct {
	db *redis.Client
}

func (rds *redisDB) Init() error {
	r := redis.NewClient(&redis.Options{
		Addr:     config.Cfg.Redis.Addr,
		Password: config.Cfg.Redis.Password,
	})
	ctx := context.Background()
	if err := r.Ping(ctx).Err(); err != nil {
		return err
	}
	rds.db = r
	return nil
}

func (rds *redisDB) Close() error {
	return rds.db.Close()
}
