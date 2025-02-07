package math

import (
	"context"
	"gopertest/internal/counter"

	"gitlab.com/engineering/products/api_security/go-common/log"
)

type Math struct {
	counter counter.Counter
	log     log.Log
}

func New(ctx context.Context, options ...Option) (*Math, error) {
	cfg := GetDefaultConfig()
	for _, opt := range options {
		opt(cfg)
	}
	if err := cfg.Validate(); err != nil {
		return nil, err
	}
	return &Math{
		counter: cfg.Counter,
	}, nil
}

func (m *Math) Add(ctx context.Context, inp *Input) *Output {
	return &Output{
		Result: inp.A + inp.B + m.next(ctx),
	}
}

func (m *Math) Subtract(ctx context.Context, inp *Input) *Output {
	return &Output{
		Result: inp.A - inp.B + m.next(ctx),
	}
}

func (m *Math) Multiply(ctx context.Context, inp *Input) *Output {
	return &Output{(inp.A * inp.B) - m.next(ctx)}
}

func (m *Math) next(ctx context.Context) int64 {
	val, err := m.counter.Get(ctx)
	if err != nil {
		m.log.Errorf(ctx, "Error getting counter: %v", err)
		return 0
	}
	return val
}
