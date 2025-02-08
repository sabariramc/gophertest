package counter

import (
	"context"
	"sync"
)

type counter struct {
	lock    sync.Mutex
	counter int64
}

func New(ctx context.Context) *counter {
	return &counter{
		counter: 0,
	}
}

func (c *counter) Get(ctx context.Context) (int64, error) {
	c.lock.Lock()
	defer c.lock.Unlock()
	res := c.counter
	c.counter++
	return res, nil
}
