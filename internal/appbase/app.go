package appbase

import (
	"context"
	"sync"

	"gitlab.com/engineering/products/api_security/go-common/log"
)

type AppBase struct {
	log           log.Log
	shutdownHooks []ShutdownHook
	healthHooks   []HealthCheckHook
	statusHooks   []StatusCheckHook
	shutdownWg    sync.WaitGroup
}

func New(ctx context.Context, option ...Option) *AppBase {
	config := GetDefaultConfig(ctx)
	for _, opt := range option {
		opt(config)
	}
	b := &AppBase{
		shutdownHooks: make([]ShutdownHook, 0, 10),
		log:           config.Log,
	}
	b.shutdownWg.Add(1)
	return b
}

func (b *AppBase) WaitForCompleteShutDown() {
	b.shutdownWg.Wait()
}
