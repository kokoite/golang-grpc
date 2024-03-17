package main

import (
	"log"

	pb "github.com/pranjal/grpc-golang/proto"
	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

const (
	port = ":8080"
)

func main() {

	conn, err := grpc.Dial("localhost"+port, grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		log.Fatal("client not connected")
	}
	defer conn.Close()
	client := pb.NewGreetServiceClient(conn)

	names := &pb.NameList{Names: []string{"alice", "bob", "john"}}
	// callSayHello(client)

	// callSayServerStreaming(client, names)

	// callClientStreaming(client, names)

	callServerClientStreaming(client, names)
}
