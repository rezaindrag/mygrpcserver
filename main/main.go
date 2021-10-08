package main

import (
	"fmt"
	"log"
	"net"

	"github.com/rezaindrag/mygrpcserver/model"
	"github.com/rezaindrag/mygrpcserver/server"
	"google.golang.org/grpc"
)

func main() {
	lis, err := net.Listen("tcp", fmt.Sprintf("localhost:%d", 8080))
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}
	var opts []grpc.ServerOption
	grpcServer := grpc.NewServer(opts...)

	// Register services
	model.RegisterUserServiceServer(grpcServer, server.NewUserService())

	log.Fatal(grpcServer.Serve(lis))
}
