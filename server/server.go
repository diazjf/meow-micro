package main

import (
	"log"
	"net"

	"github.com/diazjf/meow-micro/chat"
	"google.golang.org/grpc"
)

const (
	grpcPort = ":5001"
)

func main() {
	lis, err := net.Listen("tcp", grpcPort)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	server := chat.Server{}
	grpcServer := grpc.NewServer()
	log.Printf("Server Starting")

	chat.RegisterChatServiceServer(grpcServer, &server)
	log.Printf("Registering Server")

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
