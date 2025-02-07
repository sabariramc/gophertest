package math

import (
	"fmt"
	"gopertest/internal/counter"
)

type Config struct {
	Counter counter.Counter
}

func GetDefaultConfig() *Config {
	return &Config{}
}

type Option func(*Config)

func WithCounter(counter counter.Counter) Option {
	return func(c *Config) {
		c.Counter = counter
	}
}

func (c *Config) Validate() error {
	if c.Counter == nil {
		return fmt.Errorf("counter is required")
	}
	return nil
}
