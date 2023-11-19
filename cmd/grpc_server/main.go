package main

import (
	"database/sql"
	"log"
	"net"

	"github.com/zHenriqueGN/AppPC/internal/database"
	"github.com/zHenriqueGN/AppPC/internal/pb"
	"github.com/zHenriqueGN/AppPC/internal/service"
	"google.golang.org/grpc"
	"google.golang.org/grpc/reflection"
)

func main() {
	db, err := sql.Open("sqlite3", "./sqlite.db")
	if err != nil {
		panic(err)
	}
	defer db.Close()

	categoryDB := database.NewCategory(db)
	categoryService := service.NewCategoryService(*categoryDB)

	courseDB := database.NewCourse(db)
	courseService := service.NewCourseService(*courseDB)

	grpcServer := grpc.NewServer()
	reflection.Register(grpcServer)
	pb.RegisterCategoryServiceServer(grpcServer, categoryService)
	pb.RegisterCourseServiceServer(grpcServer, courseService)

	lis, err := net.Listen("tcp", ":50051")
	if err != nil {
		panic(err)
	}

	log.Printf("gRPC server is running on addr %s", lis.Addr().String())

	if err := grpcServer.Serve(lis); err != nil {
		panic(err)
	}

}
