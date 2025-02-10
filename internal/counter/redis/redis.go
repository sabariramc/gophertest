package counter

import (
	"context"
	"fmt"

	"github.com/redis/go-redis/v9"
)

type counter struct {
	client  *redis.Client
	key     string
	metrics *lookupMetric
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
		client:  cfg.RedisClient,
		key:     cfg.Key,
		metrics: newMetric(cfg.MetricsEnabled),
	}
	return c, nil
}

func (c *counter) Get(ctx context.Context) (res int64, err error) {
	metric := c.metrics.start(ctx, &err)
	defer metric.End()
	res, err = c.client.Incr(ctx, c.key).Result()
	if err != nil {
		return 0, fmt.Errorf("error fetching counter: %w", err)
	}
	return
}
