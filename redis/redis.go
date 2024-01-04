package redis

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"github.com/okysetiawan/go-locker"
	"github.com/okysetiawan/go-locker/errors"
	"github.com/rotisserie/eris"
	"time"
)

type Instance interface {
	redis.Cmdable
	Close() error
}

type redisLocker struct {
	instance Instance
}

func (r *redisLocker) Lock(ctx context.Context, key string, timeout time.Duration) error {
	isLocked, err := r.instance.SetNX(ctx, key, true, timeout).Result()
	if err != nil {
		return eris.Wrap(errors.ErrLock, err.Error())
	}
	if isLocked {
		return errors.ErrEventLocked
	}

	return nil
}

func (r *redisLocker) Unlock(ctx context.Context, key string) error {
	err := r.instance.Del(ctx, key).Err()
	if err != nil {
		return eris.Wrap(errors.ErrUnlock, err.Error())
	}

	return nil
}

func (r *redisLocker) Close() error {
	err := r.instance.Close()
	if err != nil {
		return eris.Wrap(errors.ErrClose, err.Error())
	}

	return nil
}

func NewLockerFromInstance(instance Instance) locker.Locker {
	return &redisLocker{instance: instance}
}

func NewLockerFromRedisClient(conf *redis.Client) locker.Locker {
	return &redisLocker{instance: conf}
}

func NewLockerFromRedisConfig(conf *redis.Options) locker.Locker {
	return &redisLocker{instance: redis.NewClient(conf)}
}

func NewLockerFromRedisClusterClient(conf *redis.ClusterClient) locker.Locker {
	return &redisLocker{instance: conf}
}

func NewLockerFromRedisClusterConfig(conf *redis.ClusterOptions) locker.Locker {
	return &redisLocker{instance: redis.NewClusterClient(conf)}
}
