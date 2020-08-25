package main

import (
	"context"
	"log"
	"net"

	"google.golang.org/grpc"
	pb "grpc-go/helloworld"
)

const (
	port = ":50051"
)

type server struct {
	pb.UnimplementedGreeterServer
}

func (s *server) SayHello(ctx context.Context, in *pb.HelloRequest) (*pb.HelloReply, error) {
	log.Printf("Received: %v", in.GetName())
	return &pb.HelloReply{Message: "Hello " + in.GetName()}, nil
}

func main() {
	listen, err := net.Listen("tcp", port)
	if err != nil{
		log.Fatalf("failed to listen: %v", err)
	}
	s := grpc.NewServer()
	pb.RegisterGreeterServer(s, &server{})
	if err := s.Serve(listen); err != nil{
		log.Fatalf("failed to serve: %v", err)
	}
}