# PoC RPC

It defines a RPC server and a client.

The server will receive a request from the client and respond with PONG if the message is PING.

The client will wait for an input from the user and send it to the server

## Running

### Server

```ssh
go run cmd/server/main.go
```

### Client

```ssh
go run cmd/client/main.go
```

## Protocol Buffer

```ssh
protoc --go_out=. --go_opt=paths=source_relative --go-grpc_out=. --go-grpc_opt=paths=source_relative proto/ping_pong.proto
```

### Server

```ssh
go run cmd/proto/server/main.go
```

### Client

```ssh
go run cmd/proto/client/main.go
```
