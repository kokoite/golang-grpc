package main

import (
	"log"
	"time"

	pb "github.com/pranjal/grpc-golang/proto"
)

func (s *helloServer) SayHelloServerStreaming(req *pb.NameList, stream pb.GreetService_SayHelloServerStreamingServer) error {
	log.Printf("Client streaming recieved %v", req.Names)
	for _, name := range req.Names {
		res := &pb.HelloResponse{
			Message: "Hello " + name,
		}
		_ = stream.Send(res)
		time.Sleep(time.Second * 1)
	}
	return nil
}
