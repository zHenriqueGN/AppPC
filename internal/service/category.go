package service

import (
	"context"
	"io"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zHenriqueGN/AppPC/internal/database"
	"github.com/zHenriqueGN/AppPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.CategoryDB
}

func NewCategoryService(categoryDB database.CategoryDB) *CategoryService {
	return &CategoryService{
		CategoryDB: categoryDB,
	}
}

func (c *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.Create(in.Name, in.Description)
	if err != nil {
		return nil, err
	}
	return &pb.CategoryResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.CategoryResponse, error) {
	category, err := c.CategoryDB.GetCategory(in.Id)
	if err != nil {
		return nil, err
	}
	return &pb.CategoryResponse{
		Id:          category.ID,
		Name:        category.Name,
		Description: category.Description,
	}, nil
}

func (c *CategoryService) ListCategories(ctx context.Context, in *pb.BlankRequest) (*pb.CategoryListResponse, error) {
	categories, err := c.CategoryDB.FindAll()
	if err != nil {
		return nil, err
	}
	var categoriesList []*pb.CategoryResponse
	for _, category := range categories {
		categoriesList = append(categoriesList, &pb.CategoryResponse{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		},
		)
	}
	return &pb.CategoryListResponse{
		Categories: categoriesList,
	}, nil
}

func (c *CategoryService) CreateCategoryStream(stream pb.CategoryService_CreateCategoryStreamServer) error {
	categories := &pb.CategoryListResponse{}
	for {
		in, err := stream.Recv()
		if err == io.EOF {
			return stream.SendAndClose(categories)
		}
		if err != nil {
			return err
		}
		category, err := c.CategoryDB.Create(in.Name, in.Description)
		if err != nil {
			return err
		}
		categories.Categories = append(categories.Categories, &pb.CategoryResponse{
			Id:          category.ID,
			Name:        category.Name,
			Description: category.Description,
		})
	}
}
