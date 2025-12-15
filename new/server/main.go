package main

import (
	"context"
	"fmt"
	pb "grpc/proto"
	"net"

	"google.golang.org/grpc"
)

type server struct {
	pb.UnimplementedPrinterServer
}

func (s *server) PrintString(ctx context.Context, in *pb.Print) (*pb.Print, error) {
	fmt.Print("String from Client = ", in.Hello)

	return &pb.Print{Hello: "Im String from Server"}, nil

}

func main() {
	listner, err := net.Listen("tcp", ":9000")

	if err != nil {
		panic(err)
	}
	serv := grpc.NewServer()
	pb.RegisterPrinterServer(serv, &server{})

	fmt.Print("Serving at port 9000")
	serv.Serve(listner)

}

// func (UnimplementedPrinterServer) PrintString(context.Context, *Print) (*Print, error)
