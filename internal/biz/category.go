package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Category struct {
	CategoryId int
	Name, Code string
	UpdateTime *timestamppb.Timestamp
}

type Categories struct {
	Collection    []*Category
	NextPageToken string
}

type CategoryRepo interface {
	ListCategories(ctx context.Context) (*Categories, error)
	GetCategory(ctx context.Context, name string) (*Category, error)
	CreateCategory(ctx context.Context, category *Category) (*Category, error)
	UpdateCategory(ctx context.Context, category *Category) (*Category, error)
	DeleteCategory(ctx context.Context, name string) (*emptypb.Empty, error)
}

type CategoryUsecase struct {
	repo CategoryRepo
}

func NewCategoryUsecase(repo CategoryRepo, logger *log.Logger) *CategoryUsecase {
	return &CategoryUsecase{repo: repo}
}
