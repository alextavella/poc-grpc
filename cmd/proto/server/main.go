package main

import (
	"log"
	"net"

	"github.com/alextavella/grpc/internal/handler"
	"github.com/alextavella/grpc/proto"
	"google.golang.org/grpc"
)

func main() {
	l, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal(err)
	}

	s := grpc.NewServer()
	proto.RegisterPingPongServer(s, &handler.Server{})
	log.Printf("Server listening at %v", l.Addr())

	err = s.Serve(l)
	if err != nil {
		log.Fatal(err)
	}
}
