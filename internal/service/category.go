package service

import (
	"context"
	"log"

	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"github.com/hi20160616/hfcms-articles/internal/data"
	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryService struct {
	pb.UnimplementedCategoriesAPIServer
	cc *biz.CategoryUsecase
}

func NewCategoryService() (*CategoryService, error) {
	dbc, err := mariadb.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewCategoryRepo(db, log.Default())
	cc := biz.NewCategoryUsecase(repo, log.Default())
	return &CategoryService{cc: cc}, nil
}

func (cs *CategoryService) ListCategories(ctx context.Context, in *pb.ListCategoriesRequest) (*pb.ListCategoriesResponse, error) {
	return nil, nil
}

func (cs *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	return nil, nil
}

func (cs *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	return nil, nil
}

func (cs *CategoryService) UpdateCategory(ctx context.Context, in *pb.UpdateCategoryRequest) (*pb.Category, error) {
	return nil, nil
}

func (cs *CategoryService) DeleteCategory(ctx context.Context, in *pb.DeleteCategoryRequest) (*emptypb.Empty, error) {
	return new(emptypb.Empty), nil
}
