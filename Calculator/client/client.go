package main

import (
	"context"
	"fmt"
	pb "grpc/protoc"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

func main() {

	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	client := pb.NewCalculateClient(conn)

	res, _ := client.Add(context.Background(), &pb.AddRequest{Number1: 1, Number2: 2})

	fmt.Println(res.Result)

}
