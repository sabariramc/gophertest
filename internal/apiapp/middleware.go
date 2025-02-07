package apiapp

import (
	"fmt"
	"gopertest/internal/appbase"
	"gopertest/internal/errors"
	"net/http"

	"github.com/google/uuid"
	"gitlab.com/engineering/products/api_security/go-common/log/correlation"
)

var pool = appbase.GetEventIDPool()

func (s *HTTPServer) withMiddleware(path string, next http.HandlerFunc) http.Handler {
	return s.SetCorrelationMiddleware(s.MetricsMiddleware(path, s.PanicHandleMiddleware(next)))
}

func (h *HTTPServer) SetCorrelationMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		correlationID := r.Header.Get("X-Correlation-ID")
		if correlationID == "" {
			correlationID = fmt.Sprintf("%v-%v", h.serviceName, uuid.NewString())
		}
		eventID := pool.Get()
		defer pool.Put(eventID)
		eventID.AddCustomPrefixKeyValue("Correlation-Id", correlationID)
		r = r.WithContext(correlation.ContextWithEventIdentifier(r.Context(), eventID))
		next.ServeHTTP(w, r)
	})
}

func (h *HTTPServer) MetricsMiddleware(path string, next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		rw := rwPool.Get()
		defer rwPool.Put(rw)
		rw.ResponseWriter = w
		m := h.metrics.start(rw, r, path)
		defer m.End()
		next.ServeHTTP(rw, r)
	})
}

func (h *HTTPServer) PanicHandleMiddleware(next http.Handler) http.Handler {
	return http.HandlerFunc(func(w http.ResponseWriter, r *http.Request) {
		ctx := r.Context()
		defer func() {
			if rec := recover(); rec != nil {
				stackTrace, err := h.base.PanicRecovery(ctx, rec)
				h.log.Errorf(ctx, "Panic Recovery: err: %v, stackTrace: %s", err, stackTrace)
				h.WriteErrorResponse(ctx, w, errors.ErrInternalServerError)
			}
		}()
		next.ServeHTTP(w, r)
	})
}
