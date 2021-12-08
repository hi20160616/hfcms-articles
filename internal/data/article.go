package data

import (
	"context"
	"log"
	"time"

	_ "github.com/hi20160616/hfcms-articles/api/articles/v1"
	_ "github.com/hi20160616/hfcms-articles/configs"
	"github.com/hi20160616/hfcms-articles/internal/biz"
)

var _ biz.ArticleRepo = new(articleRepo)

type articleRepo struct {
	data *Data
	log  *log.Logger
}

func NewArticleRepo(data *Data, logger log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.Default(),
	}
}

func (ar *articleRepo) ListArticles(ctx context.Context) ([]*biz.Article, error) {
	as := []*biz.Article{}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	return as, nil
}

func (ar *articleRepo) GetArticle(ctx context.Context, name string) (*biz.Article, error) {
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	return nil, nil
}

func (ar *articleRepo) SearchArticles(ctx context.Context, name string) ([]*biz.Article, error) {
	as := []*biz.Article{}
	ctx, cancel := context.WithTimeout(context.Background(), 50*time.Second)
	defer cancel()
	return as, nil
}

func (ar *articleRepo) CreateArticle(ctx context.Context, parent string) (*biz.Article, error) {
	return nil, nil
}

func (ar *articleRepo) UpdateArticle(ctx context.Context, article *biz.Article) (*biz.Article, error) {
	return nil, nil
}

func (ar *articleRepo) DeleteArticle(ctx context.Context, name string) error {
	return nil
}
