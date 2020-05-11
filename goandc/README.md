Interfacing Go and C
 - native linking via `cgo`
 - RPC, for example via `grpc`

CGO
 - https://blog.golang.org/cgo
 - https://golang.org/cmd/cgo/
 - https://karthikkaranth.me/blog/calling-c-code-from-go/

```
go generate query.go # compiles query.co → query.o
go run query.go -- 1 2 3 a
```

GRPC
 - https://grpc.io/docs/tutorials/basic/go/
 - https://github.com/grpc/grpc-go

```
go generate goserver/server.go
go run goserver/server.go &
go generate goclient/client.go
go run goclient/client.go -- 1 2 3 a
```
