package service

import (
	"context"
	"log"

	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"github.com/hi20160616/hfcms-articles/internal/data"
	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
)

type TagService struct {
	pb.UnimplementedTagsAPIServer
	tc *biz.TagUsecase
}

func NewTagService() (*TagService, error) {
	dbc, err := mariadb.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewTagRepo(db, log.Default())
	tc := biz.NewTagUsecase(repo, log.Default())
	return &TagService{tc: tc}, nil
}

func (ts *TagService) ListTags(ctx context.Context, in *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
	return nil, nil
}

func (ts *TagService) GetTag(ctx context.Context, in *pb.GetTagRequest) (*pb.Tag, error) {
	return nil, nil
}

func (ts *TagService) CreateTag(ctx context.Context, in *pb.CreateTagRequest) (*pb.Tag, error) {
	return nil, nil
}

func (ts *TagService) UpdateTag(ctx context.Context, in *pb.UpdateTagRequest) (*pb.Tag, error) {
	return nil, nil
}

func (ts *TagService) DeleteTag(ctx context.Context, in *pb.DeleteTagRequest) (*pb.Tag, error) {
	return nil, nil
}
