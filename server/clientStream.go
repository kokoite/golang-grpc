package main

import (
	"io"

	pb "github.com/pranjal/grpc-golang/proto"
)

func (s *helloServer) SayHelloClientStreaming(stream pb.GreetService_SayHelloClientStreamingServer) error {

	var messages []string
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			stream.SendAndClose(&pb.MessageList{Messages: messages})
			break
		}
		if err != nil {
			println("some error occured while streaming request")
			return err
		}

		data := message.Name
		if data != "" {
			messages = append(messages, data)
		}
	}
	return nil
}
