package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Tag struct {
	TagId      int
	TagName    string
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
	DeleteTag(ctx context.Context, name string) (*emptypb.Empty, error)
}

type TagUsecase struct {
	repo TagRepo
}

func NewTagUsecase(repo TagRepo, logger *log.Logger) *TagUsecase {
	return &TagUsecase{repo: repo}
}

func (tu *TagUsecase) ListTags(ctx context.Context) (*Tags, error) {
	return tu.repo.ListTags(ctx)
}
func (tu *TagUsecase) GetTag(ctx context.Context, name string) (*Tag, error) {
	return tu.repo.GetTag(ctx, name)
}
func (tu *TagUsecase) CreateTag(ctx context.Context, tag *Tag) (*Tag, error) {
	return tu.repo.CreateTag(ctx, tag)
}
func (tu *TagUsecase) UpdateTag(ctx context.Context, tag *Tag) (*Tag, error) {
	return tu.repo.UpdateTag(ctx, tag)
}
func (tu *TagUsecase) DeleteTag(ctx context.Context, name string) (*emptypb.Empty, error) {
	return tu.repo.DeleteTag(ctx, name)
}
