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

type TagService struct {
	pb.UnimplementedTagsAPIServer
	tu *biz.TagUsecase
}

func NewTagService() (*TagService, error) {
	dbc, err := mariadb.NewClient()
	if err != nil {
		return nil, err
	}

	db := &data.Data{DBClient: dbc}
	repo := data.NewTagRepo(db, log.Default())
	tu := biz.NewTagUsecase(repo, log.Default())
	return &TagService{tu: tu}, nil
}

func (ts *TagService) ListTags(ctx context.Context, in *pb.ListTagsRequest) (*pb.ListTagsResponse, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from ListTags: \n%v\n", err)
		}
	}()
	bizts, err := ts.tu.ListTags(ctx)
	if err != nil {
		return nil, err
	}
	resp := &pb.ListTagsResponse{}
	for _, tag := range bizts.Collection {
		resp.Tags = append(resp.Tags, &pb.Tag{
			TagId:      int32(tag.TagId),
			TagName:    tag.TagName,
			UpdateTime: tag.UpdateTime,
		})
	}
	return resp, nil
}

func (ts *TagService) GetTag(ctx context.Context, in *pb.GetTagRequest) (*pb.Tag, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from GetTag: %s\n%v\n", in.GetName(), err)
		}
	}()
	bizt, err := ts.tu.GetTag(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Tag{
		TagId:      int32(bizt.TagId),
		Name:       bizt.TagName,
		UpdateTime: bizt.UpdateTime,
	}, nil
}

func (ts *TagService) CreateTag(ctx context.Context, in *pb.CreateTagRequest) (*pb.Tag, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from CreateTag: \n%v\n", err)
		}
	}()
	bizt, err := ts.tu.CreateTag(ctx, &biz.Tag{
		TagName: in.Tag.TagName,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Tag{
		TagId:      int32(bizt.TagId),
		TagName:    bizt.TagName,
		UpdateTime: bizt.UpdateTime,
	}, nil
}

func (ts *TagService) UpdateTag(ctx context.Context, in *pb.UpdateTagRequest) (*pb.Tag, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from UpdateTag: \n%v\n", err)
		}
	}()
	bizt, err := ts.tu.UpdateTag(ctx, &biz.Tag{
		TagId:   int(in.Tag.TagId),
		TagName: in.Tag.TagName,
	})
	if err != nil {
		return nil, err
	}
	return &pb.Tag{
		TagId:      int32(bizt.TagId),
		TagName:    bizt.TagName,
		UpdateTime: bizt.UpdateTime,
	}, nil
}

func (ts *TagService) DeleteTag(ctx context.Context, in *pb.DeleteTagRequest) (*emptypb.Empty, error) {
	defer func() {
		if err := recover(); err != nil {
			glog.Errorf("Recovered from DeleteTag: %s\n%v\n", in.Name, err)
		}
	}()
	return ts.tu.DeleteTag(ctx, in.Name)
}
