package httpapp

import (
	"context"
	"encoding/json"
	"fmt"
	"net/http"
)

const HttpHeaderContentType = "Content-Type"

func (h *HTTPServer) WriteJSONWithStatusCode(ctx context.Context, w http.ResponseWriter, statusCode int, responseBody any) {
	var err error
	blob, ok := responseBody.([]byte)
	if !ok {
		blob, err = json.Marshal(responseBody)
		if err != nil {
			h.log.Criticalf(ctx, "Error in response json marshall", fmt.Errorf("HttpServer.WriteJsonWithStatusCode: error marshalling response: %w", err), responseBody)
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
	h.log.Errorf(ctx, "error handling request : %v", err)
	statusCode, body, parseErr := h.ProcessError(ctx, err)
	if parseErr != nil {
		h.log.Criticalf(ctx, "error marshalling response: %v", parseErr)
	}
	h.WriteJSONWithStatusCode(ctx, w, statusCode, body)
}
