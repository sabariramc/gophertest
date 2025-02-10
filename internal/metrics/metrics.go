package metrics

const (
	MetricAPIServerLatency = "api_server_latency" //Histogram
	MetricRedisLatency     = "redis_latency"      //Histogram
)

type MetricRecorder interface {
	End()
}

var dummy = &dummyMetricRecorder{}

type dummyMetricRecorder struct{}

func (d *dummyMetricRecorder) End() {}

func NewDummyMetricRecorder() MetricRecorder {
	return dummy
}
