package appbase

import (
	"context"
	"os"
	"os/signal"
	"syscall"
	"time"
)

func (b *AppBase) StartSignalMonitor(ctx context.Context) {
	c := make(chan os.Signal, 1)
	signal.Notify(c, syscall.SIGTERM, os.Interrupt)
	go b.monitorSignals(ctx, c)
}

func (b *AppBase) monitorSignals(ctx context.Context, ch chan os.Signal) {
	<-ch
	b.Shutdown(ctx)
}

func (b *AppBase) RegisterOnShutdownHook(handler ShutdownHook) {
	b.shutdownHooks = append(b.shutdownHooks, handler)
}

func (b *AppBase) Shutdown(ctx context.Context) {
	b.log.Notice(ctx, "Gracefully shutting down server")
	hooksCount := len(b.shutdownHooks)
	for i, hook := range b.shutdownHooks {
		shutdownCtx, _ := context.WithTimeout(ctx, time.Second)
		b.log.Noticef(ctx, "starting step %v of %v - %v", i+1, hooksCount, hook.Name(ctx))
		b.processShutdownHook(shutdownCtx, hook)
	}
	b.log.Notice(ctx, "server shutdown completed")
	b.shutdownWg.Done()
}

func (b *AppBase) processShutdownHook(ctx context.Context, handler ShutdownHook) {
	defer func() {
		if rec := recover(); rec != nil {
			b.log.Errorf(ctx, "panic during shutting down: name: %s: %v"+handler.Name(ctx), rec)
		}
	}()
	errCh := make(chan error)
	go func() {
		errCh <- handler.Shutdown(ctx)
	}()
	var err error
	select {
	case <-ctx.Done():
		err = context.DeadlineExceeded
	case err = <-errCh:
	}
	if err != nil {
		b.log.Errorf(ctx, "error shutting down: name: %s: %s", handler.Name(ctx), err)
	}
}
