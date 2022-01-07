package service

import (
	"context"
	"log"

	"github.com/golang/glog"
	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"github.com/hi20160616/hfcms-articles/internal/data"
	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/emptypb"
)

type CategoryService struct {
	pb.UnimplementedCategoriesAPIServer
	cu *biz.CategoryUsecase
}

func NewCategoryService() (*CategoryService, error) {
	dbc, err := mariadb.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewCategoryRepo(db, log.Default())
	cu := biz.NewCategoryUsecase(repo, log.Default())
	return &CategoryService{cu: cu}, nil
}

func (cs *CategoryService) ListCategories(ctx context.Context, in *pb.ListCategoriesRequest) (*pb.ListCategoriesResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from ListCategories: \n%v\n", err)
		}
	}()
	bizcs, err := cs.cu.ListCategories(ctx)
	if err != nil {
		return nil, err
	}
	resp := []*pb.Category{}
	for _, c := range bizcs.Collection {
		resp = append(resp, &pb.Category{
			CategoryId:   int32(c.CategoryId),
			CategoryName: c.CategoryName,
			CategoryCode: c.CategoryCode,
			UpdateTime:   c.UpdateTime,
		})
	}
	return &pb.ListCategoriesResponse{Categories: resp}, nil
}

func (cs *CategoryService) GetCategory(ctx context.Context, in *pb.GetCategoryRequest) (*pb.Category, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from GetCategory: %s\n%v\n", in.GetName(), err)
		}
	}()
	bizc, err := cs.cu.GetCategory(ctx, in.GetName())
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		CategoryId:   int32(bizc.CategoryId),
		CategoryName: bizc.CategoryName,
		CategoryCode: bizc.CategoryCode,
		UpdateTime:   bizc.UpdateTime,
	}, nil
}

func (cs *CategoryService) CreateCategory(ctx context.Context, in *pb.CreateCategoryRequest) (*pb.Category, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from CreateCategory: %s\n%v\n", in.GetName(), err)
		}
	}()
	c, err := cs.cu.CreateCategory(ctx, &biz.Category{
		CategoryId:   int(in.Category.CategoryId),
		CategoryName: in.Category.CategoryName,
		CategoryCode: in.Category.CategoryCode,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		CategoryId:   int32(c.CategoryId),
		CategoryName: c.CategoryName,
		CategoryCode: c.CategoryCode,
		UpdateTime:   c.UpdateTime,
	}, nil
}

func (cs *CategoryService) UpdateCategory(ctx context.Context, in *pb.UpdateCategoryRequest) (*pb.Category, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from UpdateCategory: \n%v\n", err)
		}
	}()
	c, err := cs.cu.UpdateCategory(ctx, &biz.Category{
		CategoryId:   int(in.Category.CategoryId),
		CategoryName: in.Category.CategoryName,
		CategoryCode: in.Category.CategoryCode,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Category{
		CategoryId:   int32(c.CategoryId),
		CategoryName: c.CategoryName,
		CategoryCode: c.CategoryCode,
		UpdateTime:   c.UpdateTime,
	}, nil
}

func (cs *CategoryService) DeleteCategory(ctx context.Context, in *pb.DeleteCategoryRequest) (*emptypb.Empty, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from DeleteCategory: %s\n%v\n", in.GetName(), err)
		}
	}()
	return cs.cu.DeleteCategory(ctx, in.Name)
}
