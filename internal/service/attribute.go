package service

import (
	"context"
	"log"

	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"github.com/hi20160616/hfcms-articles/internal/data"
	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
)

type AttributeService struct {
	pb.UnimplementedAttributesAPIServer
	ac *biz.AttributeUsecase
}

func NewAttributeService() (*AttributeService, error) {
	dbc, err := mariadb.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewAttributeRepo(db, log.Default())
	ac := biz.NewAttributeUsecase(repo, *log.Default())
	return &AttributeService{ac: ac}, nil
}

func (as *AttributeService) ListAttributes(ctx context.Context) (*pb.ListAttributesResponse, error) {
	return nil, nil
}

func (as *AttributeService) GetAttribute(ctx context.Context, in *pb.GetAttributeRequest) (*pb.Attribute, error) {
	return nil, nil
}

func (as *AttributeService) SearchAttributes(ctx context.Context, in *pb.SearchAttributesRequest) (*pb.SearchAttributesResponse, error) {
	return nil, nil
}

func (as *AttributeService) CreateAttribute(ctx context.Context, in *pb.CreateAttributeRequest) (*pb.Attribute, error) {
	return nil, nil
}

func (as *AttributeService) UpdateAttribute(ctx context.Context, in *pb.UpdateAttributeRequest) (*pb.Attribute, error) {
	return nil, nil
}

func (as *AttributeService) DeleteAttribute(ctx context.Context, in *pb.DeleteAttributeRequest) error {
	return nil
}
