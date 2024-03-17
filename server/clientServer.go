package main

import (
	"io"

	pb "github.com/pranjal/grpc-golang/proto"
)

func (s *helloServer) SayHelloBiDirectionalStreamin(stream pb.GreetService_SayHelloBiDirectionalStreaminServer) error {
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}
		res := &pb.HelloResponse{Message: "Hello " + message.Name}
		err = stream.Send(res)
		if err != nil {
			println("error occured while sending data through stream")
			return err
		}
		println("Sending to client", res.Message)
	}
	return nil
}
