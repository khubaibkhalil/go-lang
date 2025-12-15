package main

import (
	"log"
	"net"

	"github.com/khubaibkhalil/gopostgres/db"
	pb "github.com/khubaibkhalil/gopostgres/proto"
	"github.com/khubaibkhalil/gopostgres/server"
	"google.golang.org/grpc"
)

func main() {
	db.Connect()

	lis, err := net.Listen("tcp", ":9000")
	if err != nil {
		log.Fatal(err)
	}

	serv := grpc.NewServer()
	pb.RegisterMovieServiceServer(serv, &server.MovieServer{})
	log.Println("gRPC server running on :9000")

	serv.Serve(lis)
}
