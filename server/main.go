package main

import (
	"log"
	"net"

	pb "github.com/pranjal/grpc-golang/proto"
	"google.golang.org/grpc"
)

type helloServer struct {
	pb.GreetServiceServer
}

func main() {

	listner, err := net.Listen("tcp", ":8080")
	if err != nil {
		log.Fatal("something went wrong", err.Error())
	}
	grpcServer := grpc.NewServer()
	pb.RegisterGreetServiceServer(grpcServer, &helloServer{})
	println("Server started on address ", listner.Addr())
	if err := grpcServer.Serve(listner); err != nil {
		println("something wrong in grpc server")
		return
	}

}
