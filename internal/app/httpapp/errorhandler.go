package httpapp

import (
	"context"
	"encoding/json"
	e "errors"
	"gopertest/internal/errors"
	"net/http"
)

func (h *HTTPServer) ProcessError(ctx context.Context, err error) (int, []byte, error) {
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
