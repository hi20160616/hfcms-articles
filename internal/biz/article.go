package biz

import (
	"context"
	"log"

	"google.golang.org/protobuf/types/known/timestamppb"
)

type Article struct {
	CategoryId, UserId        int
	ArticleId, Title, Content string
	UpdateTime                *timestamppb.Timestamp
	TagIds                    []string
}

type Articles struct {
	Collection    []*Article
	NextPageToken string
}

type ArticleRepo interface {
	ListArticles(ctx context.Context, parent string) (*Articles, error)
	GetArticle(ctx context.Context, name string) (*Article, error)
	SearchArticles(ctx context.Context, name string) (*Articles, error)
	CreateArticle(ctx context.Context, article *Article) (*Article, error)
	UpdateArticle(ctx context.Context, article *Article) (*Article, error)
	DeleteArticle(ctx context.Context, name string) error
}

type ArticleUsecase struct {
	repo ArticleRepo
}

func NewArticleUsecase(repo ArticleRepo, logger log.Logger) *ArticleUsecase {
	return &ArticleUsecase{repo: repo}
}

func (au *ArticleUsecase) CreateArticle(ctx context.Context, article *Article) (*Article, error) {
	return au.repo.CreateArticle(ctx, article)
}

func (au *ArticleUsecase) ListArticles(ctx context.Context, parent string) (*Articles, error) {
	return au.repo.ListArticles(ctx, parent)
}

func (au *ArticleUsecase) GetArticle(ctx context.Context, name string) (*Article, error) {
	return au.repo.GetArticle(ctx, name)
}

func (au *ArticleUsecase) SearchArticles(ctx context.Context, name string) (*Articles, error) {
	return au.repo.SearchArticles(ctx, name)
}

func (au *ArticleUsecase) UpdateArticle(ctx context.Context, article *Article) (*Article, error) {
	return au.repo.UpdateArticle(ctx, article)
}

func (au *ArticleUsecase) DeleteArticle(ctx context.Context, name string) error {
	return au.repo.DeleteArticle(ctx, name)
}
