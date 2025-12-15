package server

import (
	"context"
	"log"

	"github.com/khubaibkhalil/gopostgres/db"
	"github.com/khubaibkhalil/gopostgres/models"
	pb "github.com/khubaibkhalil/gopostgres/proto"
)

type MovieServer struct {
	pb.UnimplementedMovieServiceServer
}

type SeriesServer struct {
	pb.UnimplementedSeriesServiceServer
}

func (s *MovieServer) CreateMovie(ctx context.Context, in *pb.Movie) (*pb.Movie, error) {
	movie := models.Movie{
		Name:    in.Name,
		Watched: in.Watched,
	}

	if err := db.DB.Create(&movie).Error; err != nil {
		return nil, err
	}

	return &pb.Movie{
		Id:      uint32(movie.ID),
		Name:    movie.Name,
		Watched: movie.Watched,
	}, nil

}
func (s *MovieServer) GetAllMovies(ctx context.Context, in *pb.Empty) (*pb.MovieList, error) {
	var movies []models.Movie

	if err := db.DB.Find(&movies).Error; err != nil {
		return nil, err
	}

	var pbMovies []*pb.Movie

	for _, m := range movies {
		pbMovies = append(pbMovies, &pb.Movie{
			Id:      uint32(m.ID),
			Name:    m.Name,
			Watched: m.Watched,
		})

	}

	return &pb.MovieList{Movies: pbMovies}, nil

}
func (s *MovieServer) UpdateWatched(ctx context.Context, in *pb.MovieId) (*pb.Movie, error) {
	var movie models.Movie
	if err := db.DB.First(&movie, in.Id).Error; err != nil {
		return nil, err
	}
	movie.Watched = true
	db.DB.Save(&movie)
	return &pb.Movie{
		Id:      uint32(movie.ID),
		Name:    movie.Name,
		Watched: movie.Watched,
	}, nil

}

func (s *MovieServer) DeleteMovie(ctx context.Context, in *pb.MovieId) (*pb.Empty, error) {

	if err := db.DB.Delete(&models.Movie{}, in.Id).Error; err != nil {
		return nil, err
	}
	return &pb.Empty{}, nil
}

func (s *SeriesServer) CreateSeries(ctx context.Context, in *pb.Series) (*pb.Series, error) {
	series := models.Series{
		Name:    in.Name,
		Watched: in.Watched,
	}

	if err := db.DB.Create(&series).Error; err != nil {
		log.Fatal("Something went wrong while creating datda", err)
	}

	return &pb.Series{Id: uint32(series.ID), Name: series.Name, Watched: series.Watched}, nil
}
func (s *SeriesServer) GetAllSeries(ctx context.Context, in *pb.Empty) (*pb.SeriesList, error) {
	var series []models.Series
	if err := db.DB.Find(&series).Error; err != nil {
		log.Fatal(err)
	}
	var pbSeries []*pb.Series
	for _, s := range series {
		pbSeries = append(pbSeries, &pb.Series{Id: uint32(s.ID), Name: s.Name, Watched: s.Watched})
	}

	return &pb.SeriesList{Series: pbSeries}, nil
}
func (s *SeriesServer) UpdateWatched(ctx context.Context, in *pb.SeriesId) (*pb.Series, error) {
	var series models.Series
	if err := db.DB.First(&series, in.Id).Error; err != nil {
		log.Fatal(err)
	}
	series.Watched = true

	db.DB.Save(&series)
	return &pb.Series{Id: uint32(series.ID), Name: series.Name, Watched: series.Watched}, nil

}
func (s *SeriesServer) DeleteSeries(ctx context.Context, in *pb.SeriesId) (*pb.Empty, error) {
	if err := db.DB.Delete(pb.Series{}, in.Id).Error; err != nil {
		log.Fatal("Error while deleting")
	}

	return &pb.Empty{}, nil

}
