package httpapp

import (
	"gopertest/internal/errors"
	"net/http"
)

func (h *HTTPServer) healthCheck(w http.ResponseWriter, r *http.Request) {
	err := h.lc.RunHealthCheck(r.Context())
	if err != nil {
		h.WriteErrorResponse(r.Context(), w, err)
		return
	}
	w.WriteHeader(http.StatusNoContent)
}

func (h *HTTPServer) statusCheck(w http.ResponseWriter, r *http.Request) {
	h.WriteJSON(r.Context(), w, h.lc.RunStatusCheck(r.Context()))
}

func (s *HTTPServer) notFound(w http.ResponseWriter, r *http.Request) {
	s.WriteErrorResponse(r.Context(), w, errors.ErrNotFound)
}

func (s *HTTPServer) methodNotAllowed(w http.ResponseWriter, r *http.Request) {
	s.WriteErrorResponse(r.Context(), w, errors.ErrMethodNotAllowed)
}
