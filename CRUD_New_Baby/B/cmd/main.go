package main

import (
	"B/adapters/database"
	"B/core/port"
	"B/core/service"
	"B/handler_Grpc"
	"B/pb"
	"fmt"
	"google.golang.org/grpc"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"log"
	"net"
)

func main() {
	lis,err := net.Listen("tcp",":8080")
	if err != nil {
		fmt.Println(err)
	}
	defer lis.Close()
	postGresqlCollect := NewConnect()
	postGresql := database.NewPostGresql(postGresqlCollect)
	MoviceRepository := port.MoviceService(postGresql)
	MoviceServices :=service.NewBookService(MoviceRepository)
	MoviceServer := handler_Grpc.NewServerGrpc(MoviceServices)

	s:= grpc.NewServer()

	pb.RegisterMovieServiceServer(s,MoviceServer)
	err = s.Serve(lis)
	if err != nil{
		fmt.Println("err ",err)
	}
}
func NewConnect()*gorm.DB {
	dsn := "host=localhost user=postgres password=1234 dbname=Movie port=5432 sslmode=disable TimeZone=Asia/Shanghai"
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		log.Fatal("Connect fail")
	} else {
		fmt.Print("Connect successfully")
	}
	return db
}