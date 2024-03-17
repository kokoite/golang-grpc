package main

import (
	"context"
	"log"
	"time"

	pb "github.com/pranjal/grpc-golang/proto"
)

func callSayHello(client pb.GreetServiceClient) {
	context, cancel := context.WithTimeout(context.Background(), 2*time.Second)
	defer cancel()

	resp, err := client.SayHello(context, &pb.NoParam{})
	if err != nil {
		println("something went wrong", err.Error())
		return
	}
	log.Printf("response is %s", resp.Message)
}
