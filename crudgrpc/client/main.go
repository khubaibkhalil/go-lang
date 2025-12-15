package main

import (
	"context"
	pb "crudgrpc/proto"
	"fmt"

	"google.golang.org/grpc"
	"google.golang.org/grpc/credentials/insecure"
)

type Moive struct {
	ID      string `json:"movieId"`
	Name    string `json:"movieName"`
	Watched string `json:"isWatched"`
}

func main() {
	conn, err := grpc.NewClient("localhost:9000", grpc.WithTransportCredentials(insecure.NewCredentials()))

	if err != nil {
		panic(err)
	}
	client := pb.NewNetflixClient(conn)

	movies, err := client.GetAllMovies(context.Background(), &pb.Empty{})
	if err != nil {
		fmt.Println("Somethig Went wrong")
	}

	fmt.Println("\nMovies")
	fmt.Println(movies)
	// for _, movie := range movies.List {
	// 	fmt.Printf("ID: %v ,Name: %v ,Watched: %v \n", movie.Id, movie.Movie, movie.Watched)

	// }

	// req := &pb.Movie{Id: "2", Movie: "Harry Potter", Watched: false}

	// res, err := client.CreateMovie(context.Background(), req)
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Println("Movie Created", res)

	// movie, err := client.GetMoveById(context.Background(), &pb.Id{Id: "693bfc420fae579705294dec"})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("\nID: %v \nName: %v\nWatched: %v", movie.Id, movie.Movie, movie.Watched)

	// movie, err := client.UpdateMovieById(context.Background(), &pb.Movie{Id: "693be90153e0741db99a4cb3", Movie: "Harry Potter", Watched: true})
	// if err != nil {
	// 	panic(err)
	// }
	// fmt.Printf("\nMovie Updated:\nID: %v \nName: %v\nWatched: %v", movie.Id, movie.Movie, movie.Watched)

	// _, err = client.DeleteMovieById(context.Background(), &pb.Id{Id: "693bfc420fae579705294dec"})

	// if err != nil {
	// 	panic(err)
	// }

	// fmt.Println("Deleted Movie")

}
