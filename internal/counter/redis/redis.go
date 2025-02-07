package counter

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type counter struct {
	client *redis.Client
	key    string
}

func New(ctx context.Context, options ...Option) (*counter, error) {
	cfg := GetDefaultConfig()
	for _, opt := range options {
		opt(cfg)
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	c := &counter{
		client: cfg.RedisClient,
		key:    cfg.Key,
	}
	if err := c.init(ctx); err != nil {
		return nil, err
	}
	return c, nil
}

func (c *counter) init(ctx context.Context) error {
	_, err := c.client.Get(ctx, c.key).Int64()
	if err == redis.Nil {
		_, err = c.client.Set(ctx, c.key, 0, 0).Result()
	}
	return err
}

func (c *counter) Get(ctx context.Context) (int64, error) {
	res, err := c.client.Incr(ctx, c.key).Result()
	if err != nil {
		return 0, fmt.Errorf("error fetching counter: %w", err)
	}
	return res, nil
}
