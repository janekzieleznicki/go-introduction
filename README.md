### Compile proto
```
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative generator/proto/api.proto
```

### Run test and benchmarks
```
go test ./... -bench=. -benchtime=10s -benchmem
```

