package service

import (
	"context"

	_ "github.com/mattn/go-sqlite3"
	"github.com/zHenriqueGN/AppPC/internal/database"
	"github.com/zHenriqueGN/AppPC/internal/pb"
)

type CategoryService struct {
	pb.UnimplementedCategoryServiceServer
	CategoryDB database.Category
}

func NewCategoryService(categoryDB database.Category) *CategoryService {
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
