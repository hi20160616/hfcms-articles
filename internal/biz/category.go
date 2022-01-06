package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Category struct {
	CategoryId                 int
	CategoryName, CategoryCode string
	UpdateTime                 *timestamppb.Timestamp
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

func (cu *CategoryUsecase) ListCategories(ctx context.Context) (*Categories, error) {
	return cu.repo.ListCategories(ctx)
}
func (cu *CategoryUsecase) GetCategory(ctx context.Context, name string) (*Category, error) {
	return cu.repo.GetCategory(ctx, name)
}
func (cu *CategoryUsecase) CreateCategory(ctx context.Context, category *Category) (*Category, error) {
	return cu.repo.CreateCategory(ctx, category)
}
func (cu *CategoryUsecase) UpdateCategory(ctx context.Context, category *Category) (*Category, error) {
	return cu.repo.CreateCategory(ctx, category)
}
func (cu *CategoryUsecase) DeleteCategory(ctx context.Context, name string) (*emptypb.Empty, error) {
	return cu.repo.DeleteCategory(ctx, name)
}
