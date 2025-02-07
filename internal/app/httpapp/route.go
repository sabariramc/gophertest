package httpapp

import (
	"encoding/json"
	"gopertest/internal/errors"
	"gopertest/internal/service/math"
	"io"
	"net/http"
)

type EchoData struct {
	Method string
	Body   any
	Header map[string][]string
	URL    string
}

func (s *HTTPServer) echo(w http.ResponseWriter, r *http.Request) {
	data, _ := io.ReadAll(r.Body)
	defer r.Body.Close()
	res := EchoData{
		Method: r.Method,
		Body:   string(data),
		Header: r.Header,
		URL:    r.URL.String(),
	}
	s.WriteJSON(r.Context(), w, res)
}

func (s *HTTPServer) bench(w http.ResponseWriter, r *http.Request) {
	w.WriteHeader(http.StatusOK)
	w.Write([]byte("OK"))
}

func (s *HTTPServer) add(w http.ResponseWriter, r *http.Request) {
	inp := &math.Input{}
	err := json.NewDecoder(r.Body).Decode(inp)
	defer r.Body.Close()
	if err != nil {
		s.log.Errorf(r.Context(), "Error decoding input: %v", err)
		s.WriteErrorResponse(r.Context(), w, errors.ErrBadRequest)
		return
	}
	s.WriteJSON(r.Context(), w, s.math.Add(r.Context(), inp))
}
