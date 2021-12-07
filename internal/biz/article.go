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

// TODO: fix this repo
type ArticleRepo interface {
	ListArticles(ctx context.Context) ([]*Article, error)
	GetArticle(ctx context.Context, id string) (*Article, error)
	SearchArticles(ctx context.Context, keyword ...string) ([]*Article, error)
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}
