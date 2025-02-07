package lifecycle

import (
	"context"
	"fmt"
	"gopertest/internal/errors"
	"time"
)

func (b *Lifecycle) RegisterHealthCheckHook(handler HealthCheckHook) {
	b.healthHooks = append(b.healthHooks, handler)
}

func (b *Lifecycle) RunHealthCheck(ctx context.Context) error {
	b.log.Debug(ctx, "Starting health check")
	n := len(b.healthHooks)
	description := map[string]string{}
	var err error
	for i, hook := range b.healthHooks {
		name := hook.Name(ctx)
		b.log.Debugf(ctx, "Running health check %v of %v : %v", i+1, n, name)
		hookCtx, _ := context.WithTimeout(ctx, time.Second)
		err := b.processHealthCheck(hookCtx, hook)
		if err != nil {
			description[name] = err.Error()
		}
	}
	if len(description) > 0 {
		err = &errors.CustomError{
			Code:        "HEALTH_CHECK_FAILED",
			Message:     "Health check failed, check description for details",
			Description: description,
		}
	}
	b.log.Debug(ctx, "Completed health check")
	return err
}

func (b *Lifecycle) processHealthCheck(ctx context.Context, hook HealthCheckHook) error {
	var err error
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				err = fmt.Errorf("panic during health check: %v", rec)
			}
		}()
		result := make(chan bool)
		go func() {
			err = hook.HealthCheck(ctx)
			result <- true
		}()
		select {
		case <-ctx.Done():
			err = context.DeadlineExceeded
		case <-result:
		}
	}()
	if err != nil {
		b.log.Errorf(ctx, "health check failed for hook: name: %s: %s", hook.Name(ctx), err)
	}
	return err
}
