package counter

import (
	"context"
	mt "gopertest/internal/metrics"
	"sync"
	"time"

	"gitlab.com/engineering/products/api_security/go-common/metrics"
)

const millisecondAsFloat = float64(time.Millisecond)

type lookupMetric struct {
	latency *metrics.PrometheusHistogram
	enabled bool
}

type metricRecorderObjectPool struct {
	pool *sync.Pool
}

func (p *metricRecorderObjectPool) Get() *redisMetricRecorder {
	return p.pool.Get().(*redisMetricRecorder)
}

func (p *metricRecorderObjectPool) Put(m *redisMetricRecorder) {
	m.reset()
	p.pool.Put(m)
}

var metricsPool *metricRecorderObjectPool

func newMetric(enabled bool) *lookupMetric {
	labels := []string{"status"}
	pm := &lookupMetric{enabled: enabled}
	if pm.enabled {
		pm.latency = metrics.NewHistogram(mt.MetricRedisLatency, "latency of redis operations", labels, []float64{.1, .5, 1, 5, 10})
		metricsPool = &metricRecorderObjectPool{
			pool: &sync.Pool{
				New: func() interface{} {
					return &redisMetricRecorder{
						latency: pm.latency,
					}
				},
			},
		}

	}
	return pm
}

func (m *lookupMetric) start(ctx context.Context, err *error) mt.MetricRecorder {
	if !m.enabled {
		return mt.NewDummyMetricRecorder()
	}
	return &redisMetricRecorder{
		st:      time.Now(),
		latency: m.latency,
		err:     err,
	}
}

type redisMetricRecorder struct {
	st      time.Time
	latency *metrics.PrometheusHistogram
	err     *error
}

func (m *redisMetricRecorder) reset() {
	m.st = time.Time{}
	m.err = nil
}

func (m *redisMetricRecorder) End() {
	t := float64(time.Since(m.st)) / millisecondAsFloat
	status := "success"
	if *m.err != nil {
		status = "error"
	}
	labels := map[string]string{"status": status}
	m.latency.Observe(t, labels)
}
