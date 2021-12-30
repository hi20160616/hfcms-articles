package data

import (
	"context"
	"errors"
	"fmt"
	"log"
	"regexp"
	"strings"
	"time"

	_ "github.com/hi20160616/hfcms-articles/api/articles/v1"
	_ "github.com/hi20160616/hfcms-articles/configs"
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ biz.ArticleRepo = new(articleRepo)

type articleRepo struct {
	data *Data
	log  *log.Logger
}

func NewArticleRepo(data *Data, logger *log.Logger) biz.ArticleRepo {
	return &articleRepo{
		data: data,
		log:  log.Default(),
	}
}

// parent=categories/*/articles
// TODO parent=tags/*/articles
// parent=users/*/articles
func (ar *articleRepo) ListArticles(ctx context.Context, parent string) (*biz.Articles, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	as := &mariadb.Articles{}
	bizas := &biz.Articles{}
	var err error
	re := regexp.MustCompile(`^(categories|tags)/(.+)/articles$`)
	x := re.FindStringSubmatch(parent)
	if len(x) != 3 {
		as, err = ar.data.DBClient.DatabaseClient.QueryArticle().All(ctx)
	} else {
		clause := [4]string{}
		if x[1] == "categories" {
			clause = [4]string{"category_id", "=", x[2]}
		}
		if x[1] == "users" {
			clause = [4]string{"users_id", "=", x[2]}
		}
		as, err = ar.data.DBClient.DatabaseClient.QueryArticle().
			Where(clause).All(ctx)
	}
	if err != nil {
		return nil, err
	}
	for _, a := range as.Collection {
		bizas.Collection = append(bizas.Collection, &biz.Article{
			ArticleId:  a.Id,
			Title:      a.Title,
			Content:    a.Content,
			CategoryId: a.CategoryId,
			UserId:     a.UserId,
			UpdateTime: timestamppb.New(a.UpdateTime),
		})
	}
	return bizas, nil

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
		CategoryId: a.CategoryId,
		UserId:     a.UserId,
		UpdateTime: timestamppb.New(a.UpdateTime),
	}, nil
}

func (ar *articleRepo) SearchArticles(ctx context.Context, name string) (*biz.Articles, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^articles/(.+)/search$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	kws := strings.Split(
		strings.TrimSpace(strings.ReplaceAll(x[1], "　", " ")), ",")
	cs := [][4]string{}
	for _, kw := range kws {
		cs = append(cs,
			// cs will be filtered by Where(clauses...)
			// the last `or` `and` in clause will cut off.
			// so, every clause need `or` or `and` for last element.
			[4]string{"title", "like", kw, "or"},
			[4]string{"content", "like", kw, "or"},
		)
	}
	as, err := ar.data.DBClient.DatabaseClient.QueryArticle().
		Where(cs...).All(ctx)
	if err != nil {
		return nil, err
	}
	bizas := &biz.Articles{Collection: []*biz.Article{}}
	for _, a := range as.Collection {
		bizas.Collection = append(bizas.Collection, &biz.Article{
			ArticleId:  a.Id,
			Title:      a.Title,
			Content:    a.Content,
			CategoryId: a.CategoryId,
			UserId:     a.UserId,
			UpdateTime: timestamppb.New(a.UpdateTime),
		})
	}
	return bizas, nil
}

func (ar *articleRepo) CreateArticle(ctx context.Context, article *biz.Article) (*biz.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	article.ArticleId = time.Now().Format("060102150405.000000") +
		fmt.Sprintf("%05d", article.UserId)
	if err := ar.data.DBClient.DatabaseClient.
		InsertArticle(ctx, &mariadb.Article{
			Id:         article.ArticleId,
			Title:      article.Title,
			Content:    article.Content,
			CategoryId: article.CategoryId,
			UserId:     article.UserId,
		}); err != nil {
		return nil, err
	}
	return article, nil
}

func (ar *articleRepo) UpdateArticle(ctx context.Context, article *biz.Article) (*biz.Article, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := ar.data.DBClient.DatabaseClient.
		UpdateArticle(ctx, &mariadb.Article{
			Id:         article.ArticleId,
			Title:      article.Title,
			Content:    article.Content,
			CategoryId: article.CategoryId,
			UserId:     article.UserId,
		}); err != nil {
		return nil, err
	}
	return article, nil
}

func (ar *articleRepo) DeleteArticle(ctx context.Context, name string) error {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^articles/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return errors.New("name cannot match regex express")
	}
	return ar.data.DBClient.DatabaseClient.DeleteArticle(ctx, x[1])
}
