package main

import (
	"context"
	"errors"
	"fmt"
	pb "grpclearn/protoc"
	"net"

	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

type server struct {
	pb.UnimplementedExampleServer
}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer()
	pb.RegisterExampleServer(srv, &server{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}

}

func (s *server) SayHello(c context.Context, req *pb.HelloRequest) (*pb.HelloResponse, error) {
	fmt.Println("Recived request from client", req.Name)
	fmt.Println("Hello from Server")
	return &pb.HelloResponse{}, errors.New("")

}
