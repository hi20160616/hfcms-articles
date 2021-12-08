package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Article struct {
	ArticleId, Title, Content, CategoryId, UserId string
	UpdateTime                                    *timestamppb.Timestamp
	TagIds                                        []string
}

type ArticleRepo interface {
	ListArticles(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, name string) (*Article, error)
	SearchArticles(ctx context.Context, name string) ([]*Article, error)
	CreateArticle(ctx context.Context, parent string) (*Article, error)
	UpdateArticle(ctx context.Context, article *Article) (*Article, error)
	DeleteArticle(ctx context.Context, name string) error
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}
