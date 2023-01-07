package port

import (
	"B/pb"
	"context"
)

type MoviceService interface {
	CreateMovice(context.Context,*pb.CreateMovieRequest)error
	GetMovice(context.Context,[]*pb.ReadMoviesRequest)([]*pb.Movie,error)
	GetMoviceId(context.Context,[]*pb.ReadMovieRequest,string)([]*pb.Movie,error)
	UpdateMovice(context.Context,*pb.UpdateMovieRequest,string)error
	DeleteMovice(context.Context,*pb.DeleteMovieRequest,string)error
}
