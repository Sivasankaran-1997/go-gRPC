package main

import (
	"net"
	"log"
	"google.golang.org/grpc"
	"grpc_curd/users/server/controller"
	pb "grpc_curd/users/proto"
)

func main() {
	listen, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatalf("Failed to listen. %v", err)
	}

	opts := []grpc.ServerOption{}
	srv := grpc.NewServer(opts...)


	pb.RegisterUserServiceServer(srv,controller.NewUserControllerServer())


	if err := srv.Serve(listen); err != nil {
		log.Fatalf("Failed to serve. %v", err)
		
	}

}