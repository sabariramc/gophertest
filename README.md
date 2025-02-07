```shell
go1.23.5 test -run=^$ -bench=BenchmarkMathAdd -cpuprofile cpu.out -benchmem -parallel=14 -trace trace.out -memprofile mem.out ./cmd/apiserver/app
go1.23.5 tool pprof -http=localhost:8080 app.test cpu.out
go1.23.5 tool trace -http=localhost:8080  trace.out
go1.23.5 tool pprof -http=localhost:8080 app.test mem.out
```