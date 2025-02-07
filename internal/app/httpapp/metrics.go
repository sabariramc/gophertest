package httpapp

import (
	"net/http"
	"strconv"
	"sync"
	"time"

	mt "gopertest/internal/metrics"

	"gitlab.com/engineering/products/api_security/go-common/metrics"
)

type responseWriter struct {
	statusCode int
	http.ResponseWriter
}

func (w *responseWriter) reset() {
	w.statusCode = 0
}

func (w *responseWriter) WriteHeader(code int) {
	w.statusCode = code
	w.ResponseWriter.WriteHeader(code)
}

func (w *responseWriter) Write(body []byte) (int, error) {
	return w.ResponseWriter.Write(body)
}

type responseWriterObjectPool struct {
	pool *sync.Pool
}

func (p *responseWriterObjectPool) Get() *responseWriter {
	return p.pool.Get().(*responseWriter)
}

func (p *responseWriterObjectPool) Put(rw *responseWriter) {
	rw.reset()
	p.pool.Put(rw)
}

var rwPool *responseWriterObjectPool

func init() {
	rwPool = &responseWriterObjectPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &responseWriter{}
			},
		},
	}
}

type metricRecorderObjectPool struct {
	pool *sync.Pool
}

func (p *metricRecorderObjectPool) Get() *apiMetricRecorder {
	return p.pool.Get().(*apiMetricRecorder)
}

func (p *metricRecorderObjectPool) Put(m *apiMetricRecorder) {
	m.reset()
	p.pool.Put(m)
}

var metricsPool *metricRecorderObjectPool

type serverMetric struct {
	latency *metrics.PrometheusHistogram
	enabled bool
}

func newServerMetric(enabled bool) *serverMetric {
	sm := &serverMetric{enabled: enabled}
	if sm.enabled {
		sm.latency = metrics.NewHistogram(mt.MetricAPIServerLatency, "latency of api server access", []string{"path", "method", "statusCode"}, []float64{.5, 1, 5, 10, 15, 20})
		metricsPool = &metricRecorderObjectPool{
			pool: &sync.Pool{
				New: func() interface{} {
					return &apiMetricRecorder{
						latency: sm.latency,
					}
				},
			},
		}
	}
	return sm
}

func (m *serverMetric) start(w *responseWriter, r *http.Request, path string) mt.MetricRecorder {
	if !m.enabled {
		return mt.NewDummyMetricRecorder()
	}
	mr := metricsPool.Get()
	mr.st = time.Now()
	mr.w = w
	mr.r = r
	mr.path = path
	return mr
}

const millisecondAsFloat = float64(time.Millisecond)

type apiMetricRecorder struct {
	st      time.Time
	w       *responseWriter
	r       *http.Request
	path    string
	latency *metrics.PrometheusHistogram
}

func (m *apiMetricRecorder) End() {
	t := float64(time.Since(m.st)) / millisecondAsFloat
	if m.w.statusCode == 0 {
		m.w.statusCode = http.StatusOK
	}
	labels := map[string]string{"path": m.path, "method": m.r.Method, "statusCode": strconv.Itoa(m.w.statusCode)}
	m.latency.Observe(t, labels)
	m.reset()
	metricsPool.Put(m)
}

func (m *apiMetricRecorder) reset() {
	m.st = time.Time{}
	m.w = nil
	m.r = nil
	m.path = ""
}
