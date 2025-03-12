package server

import (
	"fmt"
	"github.com/entonekryzhovnik/user-service/gen/go/userpb"
	"github.com/entonekryzhovnik/user-service/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
	"log"
	"net"
)

func StartGRPCServer(userService *service.UserService) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	userpb.RegisterUserServiceServer(grpcServer, userService)

	fmt.Println("gRPC Server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
