package main

import (
	"context"
	"gopertest/internal/apiapp"
	counter "gopertest/internal/counter/inmemory"
	"gopertest/internal/service/math"
)

func main() {
	ctx := context.Background()
	counter := counter.New(ctx)
	mathSvc, err := math.New(ctx, math.WithCounter(counter))
	if err != nil {
		panic(err)
	}
	srv, err := apiapp.New(ctx, apiapp.WithMath(mathSvc))
	if err != nil {
		panic(err)
	}
	srv.Start()
}
