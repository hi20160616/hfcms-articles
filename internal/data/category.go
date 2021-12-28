package data

import (
	"context"
	"log"
	"time"

	_ "github.com/hi20160616/hfcms-articles/api/articles/v1"
	_ "github.com/hi20160616/hfcms-articles/configs"
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ biz.CategoryRepo = new(categoryRepo)

type categoryRepo struct {
	data *Data
	log  *log.Logger
}

func NewCategoryRepo(data *Data, logger log.Logger) biz.CategoryRepo {
	return &categoryRepo{
		data: data,
		log:  log.Default(),
	}
}

func (ar *categoryRepo) ListCategories(ctx context.Context) (*biz.Categories, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	as, err := ar.data.DBClient.DatabaseClient.QueryCategory().All(ctx)
	if err != nil {
		return nil, err
	}
	bizas := &biz.Categories{Collection: []*biz.Category{}}
	for _, a := range as.Collection {
		bizas.Collection = append(bizas.Collection, &biz.Category{
			CategoryId: a.Id,
			Name:       a.Name,
			Code:       a.Code,
			UpdateTime: timestamppb.New(a.UpdateTime),
		})
	}
	return bizas, nil
}

func (cr *categoryRepo) GetCategory(ctx context.Context, name string) (*biz.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	return nil, nil
}

func (cr *categoryRepo) CreateCategory(ctx context.Context, category *biz.Category) (*biz.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	return nil, nil
}

func (cr *categoryRepo) UpdateCategory(ctx context.Context, category *biz.Category) (*biz.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	return nil, nil
}

func (cr *categoryRepo) DeleteCategory(ctx context.Context, name string) error {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	return nil
}
