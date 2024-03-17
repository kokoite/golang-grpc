package main

import (
	"context"

	pb "github.com/pranjal/grpc-golang/proto"
)

func (s *helloServer) SayHello(ctx context.Context, request *pb.NoParam) (*pb.HelloResponse, error) {
	return &pb.HelloResponse{Message: "hello is not world"}, nil
}
