package service

import (
	"context"
	"io"

	"github.com/zHenriqueGN/AppPC/internal/database"
	"github.com/zHenriqueGN/AppPC/internal/entity"
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

func (c *CourseService) CreateCourse(ctx context.Context, in *pb.CreateCourseRequest) (*pb.CourseResponse, error) {
	course := entity.NewCourse(in.Name, in.Description, in.CategoryID)
	err := c.CourseDB.Create(course)
	if err != nil {
		return nil, err
	}
	return &pb.CourseResponse{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryID:  course.CategoryID,
	}, nil
}

func (c *CourseService) CreateCourseStream(stream pb.CourseService_CreateCourseStreamServer) error {
	courses := &pb.CourseListResponse{}
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(courses)
		}
		if err != nil {
			return err
		}
		course := entity.NewCourse(in.Name, in.Description, in.CategoryID)
		err = c.CourseDB.Create(course)
		if err != nil {
			return err
		}
		courses.Courses = append(courses.Courses, &pb.CourseResponse{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryID:  course.CategoryID,
		})
	}
}

func (c *CourseService) CreateCourseStreamBidirectional(stream pb.CourseService_CreateCourseStreamBidirectionalServer) error {
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return nil
		}
		if err != nil {
			return err
		}
		course := entity.NewCourse(in.Name, in.Description, in.CategoryID)
		err = c.CourseDB.Create(course)
		if err != nil {
			return err
		}
		err = stream.Send(&pb.CourseResponse{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryID:  course.CategoryID,
		})
		if err != nil {
			return err
		}
	}
}

func (c *CourseService) GetCourse(ctx context.Context, in *pb.GetCourseRequest) (*pb.CourseResponse, error) {
	course, err := c.CourseDB.FindByID(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.CourseResponse{
		Id:          course.ID,
		Name:        course.Name,
		Description: course.Description,
		CategoryID:  course.CategoryID,
	}, nil
}

func (c *CourseService) ListCourses(ctx context.Context, in *pb.BlankRequest) (*pb.CourseListResponse, error) {
	courses, err := c.CourseDB.FindAll()
	if err != nil {
		return nil, err
	}
	var coursesList []*pb.CourseResponse
	for _, course := range courses {
		coursesList = append(coursesList, &pb.CourseResponse{
			Id:          course.ID,
			Name:        course.Name,
			Description: course.Description,
			CategoryID:  course.CategoryID,
		})
	}
	return &pb.CourseListResponse{
		Courses: coursesList,
	}, nil
}
