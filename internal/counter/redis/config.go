package counter

import (
	"fmt"
	e "gopertest/internal/env"

	"github.com/redis/go-redis/v9"
	"gitlab.com/engineering/products/api_security/go-common/environment"
)

var env = environment.EnvironmentReader{}

type Config struct {
	RedisClient *redis.Client
	Key         string
}

func GetDefaultConfig() *Config {
	serviceName := env.ReadString(e.EnvServiceName, "default")
	return &Config{
		Key: env.ReadString(e.EnvMathCounterKey, fmt.Sprintf("%v:math_counter", serviceName)),
	}
}

func (c *Config) Validate() error {
	if c.RedisClient == nil {
		return fmt.Errorf("redis client is required")
	}
	return nil
}

type Option func(*Config)

func WithRedisClient(client *redis.Client) Option {
	return func(c *Config) {
		c.RedisClient = client
	}
}
func WithKey(key string) Option {
	return func(c *Config) {
		c.Key = key
	}
}
