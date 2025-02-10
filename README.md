```shell
go1.23.5 test -v -run=^$ -bench=BenchmarkMathAdd -cpuprofile cpu.out -benchmem -parallel=14 -trace trace.out -memprofile mem.out ./internal/app/httpapp
go1.23.5 tool pprof -http=localhost:8080 httpapp.test cpu.out
go1.23.5 tool trace -http=localhost:8080 trace.out
go1.23.5 tool pprof -http=localhost:8080 httpapp.test mem.out
```


```
docker build --build-arg PRIVATE_KEY="$(cat  ~/.ssh/id_ed25519)" --add-host=gitlab:10.100.40.72 --platform linux/amd64 .
```