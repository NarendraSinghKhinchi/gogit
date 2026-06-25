package main

import (
	"context"
	"log"
	"net"

	pb "grpcserver/proto/gen"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(ctx context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	return &pb.AddResponse{
		Sum: req.A + req.B,
	}, nil
}

func main() {

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		log.Fatal("Failed to listen: ", err)
	}

	grpcServer := grpc.NewServer()
	pb.RegisterCalculateServer(grpcServer, &server{})

	log.Println("Server is running on port 50051")
	err = grpcServer.Serve(lis)
	if err != nil {
		log.Fatal("Failed to ser: ", err)
	}
}
