package cache

import (
	"context"
	"time"
)

func Set(ctx context.Context, key string, value []byte, expiration time.Duration) error {

	handle, err := getHandle()
	if err != nil {
		return err
	}

	cmd := handle.Conn.Set(ctx, key, value, expiration)
	return cmd.Err()
}

func Get(ctx context.Context, key string) ([]byte, error) {

	handle, err := getHandle()
	if err != nil {
		return nil, err
	}

	cmd := handle.Conn.Get(ctx, key)
	return cmd.Bytes()
}
