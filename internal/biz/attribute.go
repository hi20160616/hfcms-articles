package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

type Attribute struct {
	Id, UserId                   int
	Path, Description, ArticleId string
	UpdateTime                   *timestamppb.Timestamp
}

type Attributes struct {
	Collection    []*Attribute
	NextPageToken string
}

type AttributeRepo interface {
	ListAttributes(ctx context.Context) (*Attributes, error)
	GetAttribute(ctx context.Context, name string) (*Attribute, error)
	SearchAttributes(ctx context.Context, name string) (*Attributes, error)
	CreateAttribute(ctx context.Context, article *Attribute) (*Attribute, error)
	UpdateAttribute(ctx context.Context, article *Attribute) (*Attribute, error)
	DeleteAttribute(ctx context.Context, name string) (*emptypb.Empty, error)
}

type AttributeUsecase struct {
	repo AttributeRepo
}

func NewAttributeUsecase(repo AttributeRepo, logger log.Logger) *AttributeUsecase {
	return &AttributeUsecase{repo: repo}
}

func (au *AttributeUsecase) ListAttributes(ctx context.Context) (*Attributes, error) {
	return au.repo.ListAttributes(ctx)
}
func (au *AttributeUsecase) GetAttribute(ctx context.Context, name string) (*Attribute, error) {
	return au.repo.GetAttribute(ctx, name)
}
func (au *AttributeUsecase) SearchAttributes(ctx context.Context, name string) (*Attributes, error) {
	return au.repo.SearchAttributes(ctx, name)
}
func (au *AttributeUsecase) CreateAttribute(ctx context.Context, article *Attribute) (*Attribute, error) {
	return au.repo.CreateAttribute(ctx, article)
}
func (au *AttributeUsecase) UpdateAttribute(ctx context.Context, article *Attribute) (*Attribute, error) {
	return au.repo.UpdateAttribute(ctx, article)
}
func (au *AttributeUsecase) DeleteAttribute(ctx context.Context, name string) (*emptypb.Empty, error) {
	return au.repo.DeleteAttribute(ctx, name)
}
