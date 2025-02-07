package appbase

import (
	"sync"

	"gitlab.com/engineering/products/api_security/go-common/logger"
)

type eventIDObjectPool struct {
	pool *sync.Pool
}

func (p *eventIDObjectPool) Get() *logger.DefaultLogPrefix {
	return p.pool.Get().(*logger.DefaultLogPrefix)
}

func (p *eventIDObjectPool) Put(eventID *logger.DefaultLogPrefix) {
	eventID.SetAccountID(0)
	eventID.SetHost("")
	eventID.SetMethod("")
	eventID.SetRawURL("")
	eventID.ResetSetCustomPrefixKeyValues()
	p.pool.Put(eventID)
}

var eventIDPool *eventIDObjectPool

func GetEventIDPool() *eventIDObjectPool {
	return eventIDPool
}

func init() {
	eventIDPool = &eventIDObjectPool{
		pool: &sync.Pool{
			New: func() interface{} {
				return &logger.DefaultLogPrefix{}
			},
		},
	}
}
