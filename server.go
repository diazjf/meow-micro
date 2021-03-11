// Package main implements a server for Greeter service.
package main

import (
	"log"
	"net"

	"github.com/diazjf/meow-micro/chat"
	"google.golang.org/grpc"
)

const (
	port = ":5001"
)

func main() {
	// TODO: Add info logs

	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	s := chat.Server{}
	grpcServer := grpc.NewServer()

	chat.RegisterChatServiceServer(grpcServer, &s)

	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("failed to serve: %v", err)
	}
}
