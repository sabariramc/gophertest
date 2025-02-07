package appbase

import (
	"context"

	e "gopertest/internal/env"

	"gitlab.com/engineering/products/api_security/go-common/environment"
	"gitlab.com/engineering/products/api_security/go-common/log"
	"gitlab.com/engineering/products/api_security/go-common/log/factory"
)

var env = environment.EnvironmentReader{}

type Config struct {
	ServiceName string
	Log         log.Log
}

type Option func(*Config)

func GetDefaultConfig(ctx context.Context) *Config {
	return &Config{
		ServiceName: env.ReadString(e.EnvServiceName, "default"),
		Log:         factory.NewLog(ctx, "AppBase"),
	}
}

func WithServiceName(serviceName string) Option {
	return func(c *Config) {
		c.ServiceName = serviceName
	}
}

func WithLog(log log.Log) Option {
	return func(c *Config) {
		c.Log = log
	}
}
