package main

import (
	"context"
	"google.golang.org/grpc"
	pb "grpc-go/helloworld"
	"log"
	"os"
	"time"
)

const (
	address     = "localhost:50051"
	defaultname = "world"
)

func main() {
	con, err := grpc.Dial(address, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		log.Fatalf("did not connect: %v", err)
	}

	defer con.Close()
	c := pb.NewGreeterClient(con)

	name := defaultname
	if len(os.Args) > 1{
		name = os.Args[1]
	}
	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.SayHello(ctx, &pb.HelloRequest{Name: name})
	if err != nil {
		log.Fatalf("could not greet: %v", err)
	}
	log.Printf("Greeting: %s", r.GetMessage())
}
