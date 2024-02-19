package redis

import (
	"context"
	"fmt"
	"goMian/config/inner"
	"goMian/model"
	"time"
)

func (rds *redisDB) RelationInterview(it *model.Interview) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(inner.RedisTimeout))
	defer cancel()
	key := fmt.Sprintf("%d:interview", it.Owner)
	err := rds.db.SAdd(ctx, key, it.ID).Err()
	return err
}

func (rds *redisDB) FillInterview(itsID []any, owner int) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(inner.RedisTimeout))
	defer cancel()
	key := fmt.Sprintf("%d:interview", owner)
	err := rds.db.SAdd(ctx, key, itsID...).Err()
	return err
}

func (rds *redisDB) FindInterviewsByOwner(owner int) ([]string, error) {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(inner.RedisTimeout))
	defer cancel()
	key := fmt.Sprintf("%d:interview", owner)
	interviews, err := rds.db.SMembers(ctx, key).Result()
	return interviews, err
}

func (rds *redisDB) DeleteInterviewByOwner(owner int) error {
	ctx, cancel := context.WithDeadline(context.Background(), time.Now().Add(inner.RedisTimeout))
	defer cancel()
	key := fmt.Sprintf("%d:interview", owner)
	err := rds.db.Del(ctx, key).Err()
	return err
}
