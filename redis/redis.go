package redis

import (
	"context"
	redis "github.com/go-redis/redis/v8"
	"github.com/okysetiawan/go-locker"
	"github.com/okysetiawan/go-locker/errors"
	"github.com/rotisserie/eris"
	"time"
)

type redisLocker struct {
	redis *redis.Client
}

func (r *redisLocker) Lock(ctx context.Context, key string, timeout time.Duration) error {
	isLocked, err := r.redis.SetNX(ctx, key, true, timeout).Result()
	if err != nil {
		return eris.Wrap(errors.ErrLock, err.Error())
	}
	if isLocked {
		return errors.ErrEventLocked
	}

	return nil
}

func (r *redisLocker) Unlock(ctx context.Context, key string) error {
	err := r.redis.Del(ctx, key).Err()
	if err != nil {
		return eris.Wrap(errors.ErrUnlock, err.Error())
	}

	return nil
}

func (r *redisLocker) Close() error {
	err := r.redis.Close()
	if err != nil {
		return eris.Wrap(errors.ErrClose, err.Error())
	}

	return nil
}

func NewLocker() locker.Locker {
	return &redisLocker{}
}
