package lifecycle

import (
	"context"
	"fmt"
	"time"
)

func (b *Lifecycle) RegisterStatusCheckHook(handler StatusCheckHook) {
	b.statusHooks = append(b.statusHooks, handler)
}

func (b *Lifecycle) RunStatusCheck(ctx context.Context) map[string]any {
	b.log.Debug(ctx, "Starting status check")
	n := len(b.statusHooks)
	res := map[string]any{}
	for i, hook := range b.statusHooks {
		name := hook.Name(ctx)
		b.log.Debugf(ctx, "Running status check %v of %v : %v", i+1, n, name)
		hookCtx, _ := context.WithTimeout(ctx, time.Second)
		res[name] = b.processStatusCheck(hookCtx, hook)
	}
	b.log.Debug(ctx, "Completed status check")
	return res
}

func (b *Lifecycle) processStatusCheck(ctx context.Context, hook StatusCheckHook) any {
	var status any
	var err error
	func() {
		defer func() {
			if rec := recover(); rec != nil {
				err = fmt.Errorf("panic during status check: %v", rec)
			}
		}()
		result := make(chan bool)
		go func() {
			status, err = hook.StatusCheck(ctx)
			result <- true
		}()
		select {
		case <-ctx.Done():
			err = context.DeadlineExceeded
		case <-result:
		}
	}()
	if err != nil {
		status = map[string]string{
			"status": "failed",
			"error":  err.Error(),
		}
		b.log.Errorf(ctx, "status check failed for hook: name: %s: %s"+hook.Name(ctx), err)
	} else {
		status = map[string]any{
			"status": "success",
			"data":   status,
		}
	}
	return status
}
