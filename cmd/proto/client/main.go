package main

import (
	"bufio"
	"context"
	"log"
	"os"
	"time"

	"github.com/alextavella/grpc/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:50051", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := proto.NewPingPongClient(conn)

	for {
		log.Println("Enter text: ")

		scanner := bufio.NewScanner(os.Stdin)
		scanner.Scan()
		msg := scanner.Text()

		log.Printf("Sending message: %s", msg)

		ctx, cancel := context.WithTimeout(context.Background(), time.Second)

		r, err := c.Ping(ctx, &proto.PingRequest{Message: msg})
		if err != nil {
			cancel()
			log.Fatal(err)
		}

		log.Printf("Received message: %s", r.GetMessage())
		log.Println("-------------------------")
	}
}
