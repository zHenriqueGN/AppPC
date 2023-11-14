package service

import (
	"github.com/zHenriqueGN/AppPC/internal/database"
	"github.com/zHenriqueGN/AppPC/internal/pb"
)

type CourseService struct {
	pb.UnimplementedCourseServiceServer
	CourseDB database.CourseDB
}

func NewCourseService(courseDB database.CourseDB) *CourseService {
	return &CourseService{
		CourseDB: courseDB,
	}
}
