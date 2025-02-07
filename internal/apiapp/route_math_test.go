package apiapp_test

import (
	"bytes"
	"fmt"
	"net/http"
	"net/http/httptest"
	"runtime"
	"testing"
)

func BenchmarkMathAdd(b *testing.B) {
	for i := 1; i < 8; i += 1 {
		i := i
		b.Run(fmt.Sprintf("goroutines-%d", i*runtime.GOMAXPROCS(0)), func(b *testing.B) {
			b.RunParallel(func(pb *testing.PB) {
				b.SetParallelism(i)
				a, b := 0, 1
				for pb.Next() {
					data := []byte(fmt.Sprintf(`{"a":%d,"b":%d}`, a, b))
					req := httptest.NewRequest(http.MethodPost, "/math/add", bytes.NewReader(data))
					w := httptest.NewRecorder()
					srv.ServeHTTP(w, req)
					cnt++
				}
			})
		})
	}
}
