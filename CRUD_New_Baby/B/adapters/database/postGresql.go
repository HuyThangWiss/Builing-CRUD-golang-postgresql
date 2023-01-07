package database

import (
	"B/adapters/Model"
	"B/pb"
	"context"
	"errors"
	"fmt"
	"github.com/google/uuid"
	"gorm.io/gorm"
)

type MoviceServiceImpl struct {
	BookCollection *gorm.DB
}

func NewPostGresql(impl *gorm.DB)*MoviceServiceImpl  {
	return &MoviceServiceImpl{
		BookCollection: impl,
	}
}

func (u *MoviceServiceImpl)CreateMovice(ctx context.Context,req *pb.CreateMovieRequest)error{
	movice := req.GetMovie()
	movice.Id=uuid.New().String()

	data :=Model.Movie{
		ID:    movice.GetId(),
		Title: movice.GetTitle(),
		Genre: movice.GetGenre(),
	}

	res := u.BookCollection.Create(data)

	if res.RowsAffected ==0 {
		fmt.Print("error create postgetsql")
		return nil
	}
	return nil
}
func (u *MoviceServiceImpl)GetMovice(ctx context.Context,req []*pb.ReadMoviesRequest)([]*pb.Movie,error){
	movices := []*pb.Movie{}
	res:= u.BookCollection.Find(&movices)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}
	return movices,nil
}
func (u *MoviceServiceImpl)GetMoviceId(ctx context.Context,req []*pb.ReadMovieRequest,Id string)([]*pb.Movie,error){
	var movice []*pb.Movie
	res := u.BookCollection.Find(&movice,"id = ? ",Id)
	if res.RowsAffected == 0 {
		return nil, errors.New("movie not found")
	}
	return movice,nil
}
func (u *MoviceServiceImpl)UpdateMovice(ctx context.Context,req *pb.UpdateMovieRequest,id string)error{
	fmt.Println("Update")

	reqMovice := req.GetMovie()

	res := u.BookCollection.Where("id = ?",id).Updates(Model.Movie{
		Title: reqMovice.Title,
		Genre: reqMovice.Genre,
	})
	if res.RowsAffected == 0 {
		return errors.New("movies not found")
	}
	return nil
}
func (u *MoviceServiceImpl)DeleteMovice(ctx context.Context,req *pb.DeleteMovieRequest,id string)error{
	fmt.Println("Delete movice")
	var movice Model.Movie
	res := u.BookCollection.Where("id = ? ",req.GetId()).Delete(&movice)

	if res.RowsAffected == 0 {
		return errors.New("movie delete not found")
	}
	return nil
}














