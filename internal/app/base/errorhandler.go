package base

import (
	"context"
	"fmt"
	"runtime/debug"
)

func PanicRecovery(ctx context.Context, rec any) (string, error) {
	stackTrace := string(debug.Stack())
	err, ok := rec.(error)
	if !ok {
		err = fmt.Errorf("panic: %v", rec)
	}
	return stackTrace, err
}
