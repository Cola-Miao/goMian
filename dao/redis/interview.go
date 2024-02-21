package redis

import (
	"context"
	"fmt"
	"github.com/redis/go-redis/v9"
	"goMian/config/inner"
	"goMian/model"
	"time"
)

func (rds *redisDB) RelationInterview(it *model.Interview) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(inner.RedisTimeout))
	defer cancel()
	key := fmt.Sprintf("%d:interview", it.Owner)
	err := rds.db.ZAdd(ctx, key, redis.Z{
		Score:  float64(it.Time),
		Member: it.ID,
	}).Err()
	return err
}

func (rds *redisDB) FillInterview(itsZ []redis.Z, owner any) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(inner.RedisTimeout))
	defer cancel()
	key := fmt.Sprintf("%d:interview", owner)
	err := rds.db.ZAdd(ctx, key, itsZ...).Err()
	return err
}

func (rds *redisDB) FindInterviewsByOwner(owner any) ([]string, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(inner.RedisTimeout))
	defer cancel()
	key := fmt.Sprintf("%d:interview", owner)
	interviews, err := rds.db.ZRange(ctx, key, 0, -1).Result()
	return interviews, err
}

func (rds *redisDB) DeleteInterviewByOwner(owner any) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(inner.RedisTimeout))
	defer cancel()
	key := fmt.Sprintf("%d:interview", owner)
	err := rds.db.Del(ctx, key).Err()
	return err
}
