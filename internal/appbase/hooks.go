package appbase

import "context"

type Name interface {
	Name(ctx context.Context) string
}

type ShutdownHook interface {
	Name
	Shutdown(ctx context.Context) error
}

type HealthCheckHook interface {
	Name
	HealthCheck(ctx context.Context) error
}

type StatusCheckHook interface {
	Name
	StatusCheck(ctx context.Context) (any, error)
}

func (b *AppBase) RegisterHooks(hook any) {
	if hHook, ok := hook.(HealthCheckHook); ok {
		b.RegisterHealthCheckHook(hHook)
	}
	if sHook, ok := hook.(ShutdownHook); ok {
		b.RegisterOnShutdownHook(sHook)
	}
	if sHook, ok := hook.(StatusCheckHook); ok {
		b.RegisterStatusCheckHook(sHook)
	}
}
