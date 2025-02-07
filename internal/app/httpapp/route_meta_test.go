package httpapp_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"
)

func BenchmarkBenchRoute(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/meta/bench", nil)
	for i := 0; i < b.N; i++ {
		w := httptest.NewRecorder()
		srv.ServeHTTP(w, req)
		cnt++
	}
	fmt.Println(cnt)
}

func BenchmarkBenchRouteParallel(b *testing.B) {
	req := httptest.NewRequest(http.MethodGet, "/meta/bench", nil)
	for i := 1; i < 8; i += 1 {
		i := i
		b.Run(fmt.Sprintf("goroutines-%d", i*runtime.GOMAXPROCS(0)), func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				b.SetParallelism(i)
				for pb.Next() {
					w := httptest.NewRecorder()
					srv.ServeHTTP(w, req)
					cnt++
				}
			})
		})
	}
}

func generateRandomBytesSlices(n int) [][]byte {
	slices := make([][]byte, n)
	for i := range slices {
		slices[i] = make([]byte, 1024) // Adjust the size of each slice as needed
		for j := range slices[i] {
			slices[i][j] = byte(65 + j%26) // Random bytes, here just using ASCII letters for simplicity
		}
	}
	return slices
}

var randomBytesSlices = generateRandomBytesSlices(1000)

func BenchmarkEchoRoute(b *testing.B) {
	idx := 0
	next := func() []byte {
		if idx >= len(randomBytesSlices) {
			idx = 0
		}
		data := randomBytesSlices[idx]
		idx++
		return data
	}
	cnt = 0
	b.ResetTimer()
	for i := 1; i < 8; i += 1 {
		i := i
		b.Run(fmt.Sprintf("goroutines-%d", i*runtime.GOMAXPROCS(0)), func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				b.SetParallelism(i)
				for pb.Next() {
					data := next()
					req := httptest.NewRequest(http.MethodGet, "/echo/bench", bytes.NewBuffer(data))
					w := httptest.NewRecorder()
					srv.ServeHTTP(w, req)
					blob = w.Body.Bytes()
					cnt++
				}
			})
		})
	}
	fmt.Println(cnt)
}
