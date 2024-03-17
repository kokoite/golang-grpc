package main

import (
	"context"
	"log"
	"time"

	pb "github.com/pranjal/grpc-golang/proto"
)

func callClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {

	log.Printf("Client streaming started")
	stream, err := client.SayHelloClientStreaming(context.Background())
	if err != nil {
		println("some error occured while creating stream", err)
		return
	}

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			println("some error occured when sending stream")
			break
		}
		println("sent to server", name)
		time.Sleep(time.Second * 1)
	}

	res, err := stream.CloseAndRecv()
	if err != nil {
		println("some error occured")
		return
	}
	println("response from server is", res)

}
