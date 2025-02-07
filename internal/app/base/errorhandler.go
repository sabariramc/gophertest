package base

import (
	"context"
	"encoding/json"
	e "errors"
	"fmt"
	"gopertest/internal/errors"
	"net/http"
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

func ProcessError(ctx context.Context, err error) (int, []byte, error) {
	var statusCode int
	var body []byte
	statusCode = http.StatusInternalServerError
	var parseErr error
	var customErr *errors.CustomError
	var httpErr *errors.HTTPError
	if e.As(err, &httpErr) {
		statusCode = httpErr.StatusCode
		body, parseErr = json.Marshal(httpErr)
	} else if e.As(err, &customErr) {
		statusCode = http.StatusInternalServerError
		body, parseErr = json.Marshal(customErr)
	} else {
		body, parseErr = json.Marshal(errors.ErrInternalServerError)
	}
	return statusCode, body, parseErr
}
