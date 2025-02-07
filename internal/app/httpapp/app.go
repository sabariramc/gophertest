package httpapp

import (
	"context"
	"gopertest/internal/app/base/lifecycle"
	"gopertest/internal/service/math"
	"net/http"
	"time"

	"github.com/julienschmidt/httprouter"
	"gitlab.com/engineering/products/api_security/go-common/api/server"
	"gitlab.com/engineering/products/api_security/go-common/log"
)

type HTTPServer struct {
	serviceName string
	lc          *lifecycle.Lifecycle
	server      *server.RESTAPIServer
	router      *httprouter.Router
	math        *math.Math
	log         log.Log
	metrics     *serverMetric
	port        uint16
}

func New(ctx context.Context, options ...Option) (*HTTPServer, error) {
	config, err := GetDefaultConfig(ctx)
	if err != nil {
		return nil, err
	}
	for _, opt := range options {
		err := opt(config)
		if err != nil {
			return nil, err
		}
	}
	err = config.Validate()
	if err != nil {
		return nil, err
	}
	router := httprouter.New()
	h := &HTTPServer{
		serviceName: config.ServiceName,
		lc:          config.Lifecycle,
		server:      config.Server,
		log:         config.Log,
		port:        config.Port,
		metrics:     newServerMetric(config.MetricsEnabled),
		router:      router,
		math:        config.Math,
	}
	h.server.Handler = h
	h.lc.RegisterHooks(h)
	h.configureRoutes()
	return h, nil
}

func (s *HTTPServer) ServeHTTP(w http.ResponseWriter, r *http.Request) {
	st := time.Now()
	s.router.ServeHTTP(w, r)
	s.log.Debugf(r.Context(), "Request processed in %v", time.Since(st).String())
}

func (h *HTTPServer) Start() {
	ctx := context.Background()
	h.log.Noticef(ctx, "Server starting at %v", h.port)
	h.lc.StartSignalMonitor(ctx)
	err := h.server.ListenAndServe()
	if err != nil && err != http.ErrServerClosed {
		h.log.Criticalf(ctx, "Server crashed: err: %s", err)
		h.lc.Shutdown(ctx)
	}
	h.lc.WaitForCompleteShutDown()
}

func (s *HTTPServer) Name(ctx context.Context) string {
	return "HTTPServer"
}

func (s *HTTPServer) Shutdown(ctx context.Context) error {
	return s.server.Server.Shutdown(ctx)
}

func (s *HTTPServer) configureRoutes() {
	s.configureMetaRoutes()
	s.configureEchoRoutes()
	s.configureMath()
}

func (s *HTTPServer) configureMetaRoutes() {
	s.router.NotFound = http.HandlerFunc(s.notFound)
	s.router.MethodNotAllowed = http.HandlerFunc(s.methodNotAllowed)
	s.router.HandlerFunc(http.MethodGet, "/meta/health", s.healthCheck)
	s.router.HandlerFunc(http.MethodGet, "/meta/status", s.statusCheck)
	s.registerWithMiddleware(http.MethodGet, "/meta/bench", s.bench)
}

func (s *HTTPServer) configureEchoRoutes() {
	path := "/echo/:message"
	echo := s.withMiddleware(path, s.echo)
	s.router.Handler(http.MethodGet, path, echo)
	s.router.Handler(http.MethodPatch, path, echo)
	s.router.Handler(http.MethodPost, path, echo)
	s.router.Handler(http.MethodDelete, path, echo)
	s.router.Handler(http.MethodPut, path, echo)
}

func (s *HTTPServer) configureMath() {
	s.registerWithMiddleware(http.MethodPost, "/math/add", s.add)
}

func (s *HTTPServer) registerWithMiddleware(method, path string, handler http.HandlerFunc) {
	s.router.Handler(method, path, s.withMiddleware(path, handler))
}
