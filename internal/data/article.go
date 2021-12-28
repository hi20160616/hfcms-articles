package data

import (
	"context"
	"errors"
	"log"
	"regexp"
	"strconv"
	"time"

	_ "github.com/hi20160616/hfcms-articles/api/articles/v1"
	_ "github.com/hi20160616/hfcms-articles/configs"
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"google.golang.org/protobuf/types/known/timestamppb"
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

func (ar *articleRepo) ListArticles(ctx context.Context) (*biz.Articles, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	articles, err := ar.data.DBClient.DatabaseClient.QueryArticle().All(ctx)
	if err != nil {
		return nil, err
	}
	as := &biz.Articles{Collection: []*biz.Article{}}
	for _, a := range articles.Collection {
		as.Collection = append(as.Collection, &biz.Article{
			ArticleId:  a.Id,
			Title:      a.Title,
			Content:    a.Content,
			CategoryId: strconv.Itoa(a.CategoryId),
			UserId:     strconv.Itoa(a.UserId),
			UpdateTime: timestamppb.New(a.UpdateTime),
		})
	}
	return as, nil
}

func (ar *articleRepo) GetArticle(ctx context.Context, name string) (*biz.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	// name=articles/211228101711.111111000001
	re := regexp.MustCompile(`^articles/([\d.]+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	id := x[1]
	ps := [][4]string{{"id", "=", id}}
	a, err := ar.data.DBClient.DatabaseClient.QueryArticle().
		Where(ps...).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Article{
		ArticleId:  a.Id,
		Title:      a.Title,
		Content:    a.Content,
		CategoryId: strconv.Itoa(a.CategoryId),
		UserId:     strconv.Itoa(a.UserId),
		UpdateTime: timestamppb.New(a.UpdateTime),
	}, nil
}

func (ar *articleRepo) SearchArticles(ctx context.Context, name string) (*biz.Articles, error) {
	as := &biz.Articles{Collection: []*biz.Article{}}
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
