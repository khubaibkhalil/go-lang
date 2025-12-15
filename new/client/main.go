package main

import (
	"context"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"

	pb "grpc/proto"
)

func main() {

	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}

	client := pb.NewPrinterClient(conn)

	res, _ := client.PrintString(context.Background(), &pb.Print{Hello: "Im string from client"})

	fmt.Println(res.Hello)
}
