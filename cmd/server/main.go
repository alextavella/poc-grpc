package main

import (
	"log"
	"net"
	"net/rpc"

	"github.com/alextavella/grpc/internal/handler"
)

func main() {
	addr, err := net.ResolveTCPAddr("tcp", "0.0.0.0:52648")
	if err != nil {
		log.Fatal(err)
	}

	conn, err := net.ListenTCP("tcp", addr)
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()

	h := new(handler.Handler)
	log.Printf("Server listening at %v", conn.Addr())
	s := rpc.NewServer()
	s.Register(h)
	s.Accept(conn)
}
