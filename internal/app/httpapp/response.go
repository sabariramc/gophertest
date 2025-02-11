package httpapp

import (
	"context"
	"encoding/json"
	"gopertest/internal/app/base"
	"gopertest/internal/errors"
	"net/http"
)

const HttpHeaderContentType = "Content-Type"

func (h *HTTPServer) WriteJSONWithStatusCode(ctx context.Context, w http.ResponseWriter, statusCode int, responseBody any) {
	var err error
	blob, ok := responseBody.([]byte)
	if !ok {
		blob, err = json.Marshal(responseBody)
		if err != nil {
			h.log.Criticalf(ctx, "WriteJSONWithStatusCode: error in response json marshall: %v", err)
			h.WriteErrorResponse(ctx, w, errors.ErrInternalServerError)
			return
		}
	}
	w.Header().Set(HttpHeaderContentType, "application/json")
	w.WriteHeader(statusCode)
	w.Write(blob)
}

func (h *HTTPServer) WriteJSON(ctx context.Context, w http.ResponseWriter, responseBody any) {
	h.WriteJSONWithStatusCode(ctx, w, http.StatusOK, responseBody)
}

func (h *HTTPServer) WriteResponseWithStatusCode(ctx context.Context, w http.ResponseWriter, statusCode int, contentType string, responseBody []byte) {
	w.Header().Set(HttpHeaderContentType, contentType)
	w.WriteHeader(statusCode)
	w.Write(responseBody)
}

func (h *HTTPServer) WriteResponse(ctx context.Context, w http.ResponseWriter, contentType string, responseBody []byte) {
	h.WriteResponseWithStatusCode(ctx, w, http.StatusOK, contentType, responseBody)
}

func (h *HTTPServer) WriteErrorResponse(ctx context.Context, w http.ResponseWriter, err error) {
	h.log.Errorf(ctx, "WriteErrorResponse: error handling request : %v", err)
	statusCode, body, parseErr := base.ProcessError(ctx, err)
	if parseErr != nil {
		h.log.Criticalf(ctx, "WriteErrorResponse: error marshalling response: %v", parseErr)
	}
	h.WriteJSONWithStatusCode(ctx, w, statusCode, body)
}
