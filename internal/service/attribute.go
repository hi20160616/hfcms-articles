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

type AttributeService struct {
	pb.UnimplementedAttributesAPIServer
	au *biz.AttributeUsecase
}

func NewAttributeService() (*AttributeService, error) {
	dbc, err := mariadb.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewAttributeRepo(db, log.Default())
	au := biz.NewAttributeUsecase(repo, *log.Default())
	return &AttributeService{au: au}, nil
}

func (as *AttributeService) ListAttributes(ctx context.Context, in *pb.ListAttributesRequest) (*pb.ListAttributesResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from ListAttributes")
		}
	}()
	bizas, err := as.au.ListAttributes(ctx)
	if err != nil {
		return nil, err
	}
	resp := []*pb.Attribute{}
	for _, a := range bizas.Collection {
		resp = append(resp, &pb.Attribute{
			AttributeId: int32(a.Id),
			Path:        a.Path,
			Description: a.Description,
			UserId:      int32(a.UserId),
			ArticleId:   a.ArticleId,
			UpdateTime:  a.UpdateTime,
		})
	}

	return &pb.ListAttributesResponse{Attributes: resp}, nil
}

func (as *AttributeService) GetAttribute(ctx context.Context, in *pb.GetAttributeRequest) (*pb.Attribute, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from GetAttribute: %s\n%v\n", in.GetName(), err)
		}
	}()
	biza, err := as.au.GetAttribute(ctx, in.GetName())
	if err != nil {
		return nil, err
	}
	return &pb.Attribute{
		AttributeId: int32(biza.Id),
		Path:        biza.Path,
		Description: biza.Description,
		UserId:      int32(biza.UserId),
		ArticleId:   biza.ArticleId,
		UpdateTime:  biza.UpdateTime,
	}, nil
}

func (as *AttributeService) SearchAttributes(ctx context.Context, in *pb.SearchAttributesRequest) (*pb.SearchAttributesResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from SearchAttributes: %s\n%v\n", in.GetName(), err)
		}
	}()
	bizas, err := as.au.SearchAttributes(ctx, in.GetName())
	if err != nil {
		return nil, err
	}
	resp := &pb.SearchAttributesResponse{}
	for _, a := range bizas.Collection {
		resp.Attributes = append(resp.Attributes, &pb.Attribute{
			AttributeId: int32(a.Id),
			Path:        a.Path,
			Description: a.Description,
			UserId:      int32(a.UserId),
			ArticleId:   a.ArticleId,
			UpdateTime:  a.UpdateTime,
		})
	}
	return resp, nil
}

func (as *AttributeService) CreateAttribute(ctx context.Context, in *pb.CreateAttributeRequest) (*pb.Attribute, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from CreateAttribute: %s\n%v\n", in.GetName(), err)
		}
	}()
	biza, err := as.au.CreateAttribute(ctx, &biz.Attribute{
		Path:        in.Attribute.Path,
		Description: in.Attribute.Description,
		UserId:      int(in.Attribute.UserId),
		ArticleId:   in.Attribute.ArticleId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Attribute{
		AttributeId: int32(biza.Id),
		Path:        in.Attribute.Path,
		Description: in.Attribute.Description,
		UserId:      in.Attribute.UserId,
		ArticleId:   in.Attribute.ArticleId,
		UpdateTime:  in.Attribute.UpdateTime,
	}, nil
}

func (as *AttributeService) UpdateAttribute(ctx context.Context, in *pb.UpdateAttributeRequest) (*pb.Attribute, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from UpdateAttribute: \n%v\n", err)
		}
	}()
	biza, err := as.au.UpdateAttribute(ctx, &biz.Attribute{
		Id:          int(in.Attribute.AttributeId),
		Path:        in.Attribute.Path,
		Description: in.Attribute.Description,
		UserId:      int(in.Attribute.UserId),
		ArticleId:   in.Attribute.ArticleId,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Attribute{
		AttributeId: int32(biza.Id),
		Path:        biza.Path,
		Description: biza.Description,
		UserId:      int32(biza.UserId),
		ArticleId:   biza.ArticleId,
		UpdateTime:  biza.UpdateTime,
	}, nil
}

func (as *AttributeService) DeleteAttribute(ctx context.Context, in *pb.DeleteAttributeRequest) (*emptypb.Empty, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from DeleteAttribute: %s\n%v\n", in.GetName(), err)
		}
	}()
	return as.au.DeleteAttribute(ctx, in.Name)
}
