package counter

import (
	"context"
)

type Counter interface {
	Get(ctx context.Context) (int64, error) // First value is always 1
}
