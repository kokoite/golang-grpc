package main

import (
	"context"
	"io"
	"time"

	pb "github.com/pranjal/grpc-golang/proto"
)

func callServerClientStreaming(client pb.GreetServiceClient, names *pb.NameList) {
	stream, err := client.SayHelloBiDirectionalStreamin(context.Background())
	if err != nil {
		println("something is wrong", err)
		return
	}
	waitChanel := make(chan struct{})
	go func() {
		for {
			message, err := stream.Recv()
			if err == io.EOF {
				println("end of line reached")
				break
			}
			if err != nil {
				println("error occured client side", err)
				break
			}
			println("response from server stream", message.Message)
			time.Sleep(time.Second * 2)
		}
		close(waitChanel)
	}()

	for _, name := range names.Names {
		req := &pb.HelloRequest{
			Name: name,
		}
		if err := stream.Send(req); err != nil {
			println("error while sending data", err)
			break
		}
	}
	stream.CloseSend()
	<-waitChanel
	println("bidrectional streaming finished")

}
