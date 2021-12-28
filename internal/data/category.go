package data

import (
	"context"
	"errors"
	"log"
	"regexp"
	"time"

	_ "github.com/hi20160616/hfcms-articles/api/articles/v1"
	_ "github.com/hi20160616/hfcms-articles/configs"
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ biz.CategoryRepo = new(categoryRepo)

type categoryRepo struct {
	data *Data
	log  *log.Logger
}

func NewCategoryRepo(data *Data, logger *log.Logger) biz.CategoryRepo {
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
	// name=categories/1
	re := regexp.MustCompile(`^categories/(\d+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	id := x[1]
	ps := [][4]string{{"id", "=", id}}
	c, err := cr.data.DBClient.DatabaseClient.QueryCategory().
		Where(ps...).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Category{
		CategoryId: c.Id,
		Name:       c.Name,
		Code:       c.Code,
		UpdateTime: timestamppb.New(c.UpdateTime),
	}, nil
}

func (cr *categoryRepo) CreateCategory(ctx context.Context, category *biz.Category) (*biz.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := cr.data.DBClient.DatabaseClient.
		InsertCategory(ctx, &mariadb.Category{
			Name: category.Name,
			Code: category.Code,
		}); err != nil {
		return nil, err
	}
	return category, nil
}

func (cr *categoryRepo) UpdateCategory(ctx context.Context, category *biz.Category) (*biz.Category, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := cr.data.DBClient.DatabaseClient.
		UpdateCategory(ctx, &mariadb.Category{
			Id:   category.CategoryId,
			Name: category.Name,
			Code: category.Code,
		}); err != nil {
		return nil, err
	}
	return category, nil
}

func (cr *categoryRepo) DeleteCategory(ctx context.Context, name string) error {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^categories/(\d+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return errors.New("name cannot match regex express")
	}
	return cr.data.DBClient.DatabaseClient.DeleteCategory(ctx, x[1])
}
