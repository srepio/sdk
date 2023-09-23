package client

import (
	"context"
	"errors"
	"time"
)

type Effector func(ctx context.Context) (any, error)

func Retry(e Effector, tries int, delay time.Duration) Effector {
	return func(ctx context.Context) (any, error) {
		for i := 0; i < tries; i++ {
			out, err := e(ctx)
			if err == nil || i >= tries-1 {
				return out, err
			}

			select {
			case <-time.After(delay):
			case <-ctx.Done():
				return nil, ctx.Err()
			}
		}
		return nil, errors.New("retry error")
	}
}
