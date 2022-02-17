package main

import (
	"context"
	pb "grpc-web-sample/helloworld/proto"
	"log"
	"net"

	"google.golang.org/grpc"
)

type HelloworldHandler struct {
	pb.UnimplementedGreeterServer
}

func (h HelloworldHandler) SayHello(context.Context, *pb.HelloReply) {

}

func main() {
	port := ":9090"
	lis, err := net.Listen("tcp", port)
	if err != nil {
		log.Fatal("faild to listen %v", err)
	}
	server := grpc.NewServer()
	pb.RegisterGreeterServer(server, &HelloworldHandler{})
}
