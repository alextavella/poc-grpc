package handler

import (
	"context"
	"log"

	"github.com/alextavella/grpc/proto"
)

type Server struct {
	proto.UnimplementedPingPongServer
}

func (s *Server) Ping(_ context.Context, in *proto.PingRequest) (*proto.PingResponse, error) {
	r := &proto.PingResponse{}
	m := in.GetMessage()
	log.Println("Received message:", m)

	switch m {
	case "ping", "Ping", "PING":
		r.Message = "pong"
	default:
		r.Message = "I don't understand"
	}

	log.Println("Sending message:", r.Message)

	return r, nil
}
