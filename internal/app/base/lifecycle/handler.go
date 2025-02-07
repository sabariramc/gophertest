package lifecycle

import (
	"context"
	"sync"

	"gitlab.com/engineering/products/api_security/go-common/log"
)

type Lifecycle struct {
	log           log.Log
	shutdownHooks []ShutdownHook
	healthHooks   []HealthCheckHook
	statusHooks   []StatusCheckHook
	shutdownWg    sync.WaitGroup
}

func New(ctx context.Context, option ...Option) *Lifecycle {
	config := GetDefaultConfig(ctx)
	for _, opt := range option {
		opt(config)
	}
	b := &Lifecycle{
		shutdownHooks: make([]ShutdownHook, 0, 10),
		log:           config.Log,
	}
	b.shutdownWg.Add(1)
	return b
}

func (b *Lifecycle) WaitForCompleteShutDown() {
	b.shutdownWg.Wait()
}
