package main

import (
	"context"
	"errors"
	"fmt"
	pb "grpc/protoc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedCalculateServer
}

func (s *server) Add(c context.Context, req *pb.AddRequest) (*pb.AddResponse, error) {
	fmt.Println("Recived Request from the client ", req)
	fmt.Print("Server....")

	return &pb.AddResponse{Result: req.Number1 + req.Number2}, errors.New("")

}

func main() {
	listner, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}

	srv := grpc.NewServer()

	pb.RegisterCalculateServer(srv, &server{})

	reflection.Register(srv)

	if err := srv.Serve(listner); err != nil {
		panic(err)
	}

}
