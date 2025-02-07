package errors

import (
	"fmt"
	"net/http"
)

var _ error = &CustomError{}
var _ error = &HTTPError{}

type CustomError struct {
	Code        string
	Message     string
	Description any `json:",omitempty"`
}

func (e *CustomError) Error() string {
	return fmt.Sprintf("Code: %s, Message: %s", e.Code, e.Message)
}

type HTTPError struct {
	StatusCode   int `json:"-"`
	*CustomError `json:",inline"`
}

var ErrBadRequest = &HTTPError{
	StatusCode: http.StatusBadRequest,
	CustomError: &CustomError{
		Code:    "BAD_REQUEST",
		Message: "Invalid input",
	},
}

var ErrNotFound = &HTTPError{
	StatusCode: http.StatusNotFound,
	CustomError: &CustomError{
		Code:    "NOT_FOUND",
		Message: "URL Not Found",
	},
}

var ErrMethodNotAllowed = &HTTPError{
	StatusCode: http.StatusMethodNotAllowed,
	CustomError: &CustomError{
		Code:    "METHOD_NOT_ALLOWED",
		Message: "Method Not Allowed",
	},
}

var ErrInternalServerError = &HTTPError{
	StatusCode: http.StatusInternalServerError,
	CustomError: &CustomError{
		Code:        "INTERNAL_SERVER_ERROR",
		Message:     "Internal Server Error",
		Description: "Retry after some time, if persist contact technical team",
	},
}
