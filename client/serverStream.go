package main

import (
	"context"
	"io"
	"log"
	"time"

	pb "github.com/pranjal/grpc-golang/proto"
)

func callSayServerStreaming(client pb.GreetServiceClient, nameList *pb.NameList) {
	log.Printf("Call started for sending server streaming")
	context, cancel := context.WithTimeout(context.Background(), time.Second*5)
	defer cancel()
	stream, err := client.SayHelloServerStreaming(context, nameList)
	if err != nil {
		println("something went wrong while fetching stream from server", err)
		return
	}
	for {
		message, err := stream.Recv()
		if err == io.EOF {
			break
		}

		if err != nil {
			println("something went wrong", err)
			return
		}
		println("streamed response is ", message.Message)
	}
}
