package locker

import (
	"context"
	"time"
)

type Locker interface {
	Lock(ctx context.Context, key string, timeout time.Duration) error
	Unlock(ctx context.Context, key string) error
	Close() error
}
