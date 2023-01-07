package main

import (
	"A/Interceptor"
	"A/pb"
	"context"
	"fmt"
	grpc_middleware "github.com/grpc-ecosystem/go-grpc-middleware"
	"google.golang.org/grpc"
)

func main() {
	conn, err := grpc.Dial(
		":8080",
		grpc.WithInsecure(),
		grpc.WithUnaryInterceptor(
			grpc_middleware.ChainUnaryClient(
				Interceptor.DateLogClientInterceptor,
				Interceptor.MethodLogClientInterceptor,
			),
		),
	)
	if err != nil {
		fmt.Println(err)
	}
	defer conn.Close()

	BookServ := pb.NewMovieServiceClient(conn)
	fmt.Println("start call gRpc")
	//var movie Movie
	/*
	data := &pb.Movie{
		Title: "Tk33",
		Genre: "Gee33",
	}
	res,err1 := BookServ.CreateMovie(context.Background(),&pb.CreateMovieRequest{Movie: data})
	if err != nil{
		fmt.Println("err ",err1)
	}
	fmt.Println("insert succesfull",res.Movie)
	*/
	/*
	fmt.Println("Read data")
	 arr,_:=BookServ.GetMovies(context.Background(),&pb.ReadMoviesRequest{})
	 fmt.Println("data : ",arr)

	 */
	/*
	res,_:=BookServ.GetMovie(context.Background(),&pb.ReadMovieRequest{Id: "001"})
	fmt.Println("data ",res)
	 */
	/*
	res,_:= BookServ.DeleteMovie(context.Background(),&pb.DeleteMovieRequest{Id: "004"})
	fmt.Println("Delete :  ",res)
	 */
	res,_:= BookServ.UpdateMovie(context.Background(),&pb.UpdateMovieRequest{Movie: &pb.Movie{
		Id:    "004",
		Title: "New",
		Genre: "New",
	}})
	fmt.Println("Update ",res)
}
type Movie struct {
	ID    string `json:"id"`
	Title string `json:"Title"`
	Genre string `json:"genre"`
}
