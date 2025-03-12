package server

import (
	"fmt"
	"log"
	"net"
	"user-service/gen/go/userpb"
	"user-service/internal/service"

	"google.golang.org/grpc"
)

func StartGRPCServer(userService *service.UserService) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	userpb.RegisterUserServiceServer(grpcServer, userService)

	fmt.Println("gRPC Server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
