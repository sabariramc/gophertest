package httpapp_test

import (
	"context"
	"gopertest/internal/app/httpapp"
	"gopertest/internal/component"
	inmCtr "gopertest/internal/counter/inmemory"
	redCtr "gopertest/internal/counter/redis"
	"gopertest/internal/service/math"
	"os"
	"testing"

	"github.com/redis/go-redis/v9"
)

var cnt int
var blob []byte
var statusCode int

var srv *httpapp.HTTPServer

func createInMemory() error {
	ctx := context.Background()
	counter := inmCtr.New(ctx)
	mathSvc, err := math.New(ctx, math.WithCounter(counter))
	if err != nil {
		return err
	}
	srv, err = httpapp.New(ctx, httpapp.WithMath(mathSvc))
	if err != nil {
		return err
	}
	return nil
}

func createRedis() (func(), error) {
	comp := component.NewTestDependencyManager(true)
	comp.Setup()
	ctx := context.Background()
	counter, err := redCtr.New(ctx, redCtr.WithRedisClient(redis.NewClient(&redis.Options{
		Addr: "localhost:6379",
	})))
	if err != nil {
		return nil, err
	}
	mathSvc, err := math.New(ctx, math.WithCounter(counter))
	if err != nil {
		return nil, err
	}
	srv, err = httpapp.New(ctx, httpapp.WithMath(mathSvc))
	if err != nil {
		return nil, err
	}
	return func() {
		comp.Teardown()
	}, nil
}

func testInmemory(m *testing.M) {
	err := createInMemory()
	if err != nil {
		os.Exit(1)
	}
	code := m.Run()
	if code != 0 {
		os.Exit(code)
	}
}

func testRedis(m *testing.M) {
	teardown, err := createRedis()
	if err != nil {
		os.Exit(1)
	}
	code := m.Run()
	teardown()
	if code != 0 {
		os.Exit(code)
	}
}

func TestMain(m *testing.M) {
	// testInmemory(m)
	testRedis(m)
	os.Exit(0)
}
