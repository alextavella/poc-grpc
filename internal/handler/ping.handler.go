package handler

import (
	"fmt"
)

type Args struct {
	Message string
}

type Handler int

func (h *Handler) Ping(args *Args, reply *Args) error {
	fmt.Println("Received message: ", args.Message)

	switch args.Message {
	case "ping", "Ping", "PING":
		reply.Message = "pong"
	default:
		reply.Message = "I don't understand"
	}

	fmt.Println("Sending message: ", reply.Message)
	return nil
}
