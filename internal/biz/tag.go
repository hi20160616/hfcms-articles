package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Tag struct {
	TagId      int
	Name, Code string
	UpdateTime *timestamppb.Timestamp
}

type Tags struct {
	Collection    []*Tag
	NextPageToken string
}

type TagRepo interface {
	ListTags(ctx context.Context) (*Tags, error)
	GetTag(ctx context.Context, name string) (*Tag, error)
	CreateTag(ctx context.Context, tag *Tag) (*Tag, error)
	UpdateTag(ctx context.Context, tag *Tag) (*Tag, error)
	DeleteTag(ctx context.Context, name string) error
}

type TagUsecase struct {
	repo TagRepo
}

func NewTagUsecase(repo TagRepo, logger *log.Logger) *TagUsecase {
	return &TagUsecase{repo: repo}
}
