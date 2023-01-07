package service

import (
	"B/core/port"
	"B/pb"
	"context"
	"fmt"
	"log"
)

type BookService struct {
	BooksRepository port.MoviceService
}

func NewBookService(service port.MoviceService)*BookService  {
	return &BookService{
		BooksRepository: service,
	}
}

func (u *BookService)CreateMice(ctx context.Context,Movice *pb.CreateMovieRequest)error  {
	err := u.BooksRepository.CreateMovice(ctx,Movice)
	if err != nil{
		log.Fatalf("err : ", err)
		return  err
	}
	return nil
}
func (u *BookService)ReadAll(ctx context.Context,Movice []*pb.ReadMoviesRequest)([]*pb.Movie,error)  {
	arr,_:= u.BooksRepository.GetMovice(ctx,Movice)

	return arr,nil
}
func (u *BookService)ReadAllId(ctx context.Context,Movice []*pb.ReadMovieRequest,Id string)([]*pb.Movie,error)  {
	arr,_:=u.BooksRepository.GetMoviceId(ctx,Movice,Id)

	return arr,nil
}
func (u *BookService)DeleteMovice(ctx context.Context,movice *pb.DeleteMovieRequest,Id string)error  {
	err := u.BooksRepository.DeleteMovice(ctx,movice,Id)
	if err != nil{
		fmt.Print("err service")
		return err
	}
	return nil
}

func (u *BookService)UpdateMovie(ctx context.Context,movice *pb.UpdateMovieRequest,id string)error  {
	err := u.BooksRepository.UpdateMovice(ctx,movice,id)
	if err != nil{
		fmt.Print("err service update")
		return err
	}
	return nil
}












