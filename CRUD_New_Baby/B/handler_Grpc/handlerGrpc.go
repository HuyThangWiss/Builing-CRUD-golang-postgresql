package handler_Grpc

import (
	"B/core/service"
	"B/pb"
	"context"
	"fmt"
	"google.golang.org/grpc/codes"
	"google.golang.org/grpc/status"
)

type ServerBook struct {
	pb.UnimplementedMovieServiceServer
	BookServer *service.BookService
}

func NewServerGrpc(BookServer *service.BookService)*ServerBook  {
	return &ServerBook{
		BookServer: BookServer,
	}
}
///
func (s *ServerBook) CreateMovie(ctx context.Context,req *pb.CreateMovieRequest) (*pb.CreateMovieResponse, error) {
	fmt.Print("Create Grpc ...")
	err := s.BookServer.CreateMice(ctx,req)
	if err != nil{
		return nil, status.Errorf(codes.Unimplemented, "method CreateMovie not implemented")
	}
	return nil,nil
}
func (s *ServerBook) GetMovie(ctx context.Context,req *pb.ReadMovieRequest) (*pb.ReadMovieResponse, error) {
	fmt.Print("Find Id movie")
	movice := []*pb.ReadMovieRequest{}
	arr,err:=s.BookServer.ReadAllId(ctx,movice,req.Id)
	if err != nil{
		return nil, status.Errorf(codes.Unimplemented, "method GetMovie not implemented")
	}
	return &pb.ReadMovieResponse{Movie: arr},nil
}
////
func (s *ServerBook) GetMovies(ctx context.Context,req *pb.ReadMoviesRequest) (*pb.ReadMoviesResponse, error) {

	fmt.Print("Find All movie ")
	movice := []*pb.ReadMoviesRequest{}
	arr,err := s.BookServer.ReadAll(ctx,movice)

	if err != nil{
		return nil, status.Errorf(codes.Unimplemented, "method GetMovies not implemented")
	}
	return &pb.ReadMoviesResponse{Movies: arr},nil


}
func (s *ServerBook) UpdateMovie(ctx context.Context,req *pb.UpdateMovieRequest) (*pb.UpdateMovieResponse, error) {
	fmt.Print("Update movie ")
	reqMovie := req.GetMovie()
	err:= s.BookServer.UpdateMovie(ctx,req,reqMovie.GetId())
	if err != nil{
		return nil, status.Errorf(codes.Unimplemented, "method UpdateMovie not implemented")
	}
	return &pb.UpdateMovieResponse{Movie: reqMovie},nil
}
func (s *ServerBook) DeleteMovie(ctx context.Context,req *pb.DeleteMovieRequest) (*pb.DeleteMovieResponse, error) {
	fmt.Print("delete ")
	err:= s.BookServer.DeleteMovice(ctx,req,req.GetId())
	if err != nil{
		return nil, status.Errorf(codes.Unimplemented, "method DeleteMovie not implemented")
	}
	return &pb.DeleteMovieResponse{Success: true},nil
}