// server.go
package main

import (
	"context"
	"log"
	"net"

	// pb "your/module/ping" // adjust path
	pb "go-cassandra-demo-service/proutput/servicepb" // adjust path

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedUserRpcServer
}

func (s *server) AddUser(ctx context.Context, req *pb.UserInfo) (*pb.Success, error) {
	return &pb.Success{Message: "pong from Go"}, nil
}

func main() {
	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatalf("failed to listen: %v", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterUserRpcServer(grpcServer, &server{})

	log.Println("gRPC server listening on port 50051")
	grpcServer.Serve(lis)
}
