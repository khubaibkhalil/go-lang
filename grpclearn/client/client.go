package main

import (
	"context"
	pb "grpclearn/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))
	if err != nil {
		panic(err)
	}
	client := pb.NewExampleClient(conn)

	req := &pb.HelloRequest{Name: "Hello from Khubaib Client"}

	client.SayHello(context.TODO(), req)
}
