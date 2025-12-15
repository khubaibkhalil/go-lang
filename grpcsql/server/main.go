package main

import (
	"context"
	pb "crudgrpc/proto"
	"fmt"
	"log"
	"net"

	"go.mongodb.org/mongo-driver/v2/bson"
	"go.mongodb.org/mongo-driver/v2/mongo"
	"go.mongodb.org/mongo-driver/v2/mongo/options"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

const connectionString = ""

const dbName = "netflix"
const colName = "watchlist"

var collection *mongo.Collection

func init() {
	clientOption := options.Client().ApplyURI(connectionString)
	client, err := mongo.Connect(clientOption)

	if err != nil {
		log.Fatal(err)
	}
	fmt.Println("MongoDB Connected")

	collection = client.Database(dbName).Collection(colName)
	fmt.Println("Collection instance is Ready")

}

type server struct {
	pb.UnimplementedNetflixServer
}
type Movie struct {
	ID      bson.ObjectID `bson:"_id,omitempty"`
	Name    string        `bson:"movie"`
	Watched bool          `bson:"watched"`
}

// type NetflixClient interface {
// 	CreateMovie(ctx context.Context, in *Movie, opts ...grpc.CallOption) (*Movie, error)
// 	GetAllMovies(ctx context.Context, in *Empty, opts ...grpc.CallOption) (*MovieList, error)
// 	GetMoveById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Movie, error)
// 	UpdateMovieById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Movie, error)
// 	DeleteMovieById(ctx context.Context, in *Id, opts ...grpc.CallOption) (*Empty, error)
// }

func (s *server) CreateMovie(ctx context.Context, in *pb.Movie) (*pb.Movie, error) {

	movie := Movie{
		Name:    in.Movie,
		Watched: in.Watched,
	}
	res, err := collection.InsertOne(ctx, movie)
	if err != nil {
		return nil, err
	}

	id := res.InsertedID.(bson.ObjectID)

	fmt.Printf("id %T", res.InsertedID)

	return &pb.Movie{
		Id:      id.Hex(),
		Movie:   movie.Name,
		Watched: movie.Watched,
	}, nil

}

func (s *server) GetAllMovies(ctx context.Context, in *pb.Empty) (*pb.MovieList, error) {
	cursor, err := collection.Find(ctx, bson.M{})
	if err != nil {
		log.Fatal(err)
	}
	defer cursor.Close(ctx)
	fmt.Println("Connected to Collection")
	var movies []*pb.Movie
	var movie pb.Movie
	for cursor.Next(ctx) {

		err := cursor.Decode(&movie)
		fmt.Println("Print m")
		fmt.Print(movie)
		if err != nil {
			log.Fatal(err)
		}
		movies = append(movies, &pb.Movie{
			Id:      movie.Id,
			Movie:   movie.Movie,
			Watched: movie.Watched,
		})
	}
	if err := cursor.Err(); err != nil {
		return nil, err
	}

	return &pb.MovieList{List: movies}, nil
}

func (s *server) GetMoveById(ctx context.Context, in *pb.Id) (*pb.Movie, error) {
	id, _ := bson.ObjectIDFromHex(in.Id)
	movie := &Movie{}
	filter := bson.M{"_id": id}

	res := collection.FindOne(ctx, filter)

	fmt.Println("Result")
	fmt.Print(res)

	res.Decode(movie)
	return &pb.Movie{
		Id:      movie.ID.Hex(),
		Movie:   movie.Name,
		Watched: movie.Watched,
	}, nil

}
func (s *server) UpdateMovieById(ctx context.Context, in *pb.Movie) (*pb.Movie, error) {
	id, _ := bson.ObjectIDFromHex(in.Id)
	movie := &Movie{

		Name:    in.Movie,
		Watched: in.Watched,
	}

	collection.UpdateByID(ctx, id, movie)

	return &pb.Movie{
		Id:      id.Hex(),
		Movie:   in.Movie,
		Watched: in.Watched,
	}, nil

}

func (s *server) DeleteMovieById(ctx context.Context, in *pb.Id) (*pb.Empty, error) {
	id, _ := bson.ObjectIDFromHex(in.Id)
	filter := bson.M{"_id": id}
	res, err := collection.DeleteOne(ctx, filter)
	if err != nil {
		panic(err)
	}
	fmt.Println("Deleted : ", res.DeletedCount)

	return &pb.Empty{}, nil

}

func main() {
	listener, tcpErr := net.Listen("tcp", ":9000")
	if tcpErr != nil {
		panic(tcpErr)
	}
	srv := grpc.NewServer()
	pb.RegisterNetflixServer(srv, &server{})

	reflection.Register(srv)

	if e := srv.Serve(listener); e != nil {
		panic(e)
	}
}
