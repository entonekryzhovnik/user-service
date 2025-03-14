package server

import (
	"fmt"
	"github.com/entonekryzhovnik/user-service/gen/go/userpb"
	"github.com/entonekryzhovnik/user-service/internal/controller"
	"log"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func StartGRPCServer(controller *controller.UserController) {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("Failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	userpb.RegisterUserServiceServer(grpcServer, controller)

	fmt.Println("gRPC Server is running on port 50051")
	if err := grpcServer.Serve(lis); err != nil {
		log.Fatalf("Failed to serve: %v", err)
	}
}
