package data

import (
	"context"
	"errors"
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

var _ biz.AttributeRepo = new(attributeRepo)

type attributeRepo struct {
	data *Data
	log  *log.Logger
}

func NewAttributeRepo(data *Data, logger *log.Logger) biz.AttributeRepo {
	return &attributeRepo{
		data: data,
		log:  log.Default(),
	}
}

func (ar *attributeRepo) ListAttributes(ctx context.Context) (*biz.Attributes, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	as, err := ar.data.DBClient.DatabaseClient.QueryAttribute().All(ctx)
	if err != nil {
		return nil, err
	}
	bizas := &biz.Attributes{Collection: []*biz.Attribute{}}
	for _, a := range as.Collection {
		bizas.Collection = append(bizas.Collection, &biz.Attribute{
			Id:          a.Id,
			Path:        a.Path,
			Description: a.Description,
			UserId:      a.UserId,
			ArticleId:   a.ArticleId,
			UpdateTime:  timestamppb.New(a.UpdateTime),
		})
	}
	return bizas, nil
}

func (ar *attributeRepo) GetAttribute(ctx context.Context, name string) (*biz.Attribute, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^attributes/([\d.]+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	id := x[1]
	ps := [][4]string{{"id", "=", id}}
	a, err := ar.data.DBClient.DatabaseClient.QueryAttribute().
		Where(ps...).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Attribute{
		Id:          a.Id,
		Path:        a.Path,
		Description: a.Description,
		UserId:      a.UserId,
		ArticleId:   a.ArticleId,
		UpdateTime:  timestamppb.New(a.UpdateTime),
	}, nil
}

// TODO: search via diff opts, by userid, articleid and fuzz search description
func (ar *attributeRepo) SearchAttributes(ctx context.Context, name string) (*biz.Attributes, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^attributes/(.+)/search$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	kws := strings.Split(
		strings.TrimSpace(strings.ReplaceAll(x[1], "　", " ")), " ")
	cs := [][4]string{}
	for _, kw := range kws {
		cs = append(cs,
			// cs will be filtered by Where(clauses...)
			// the last `or` `and` in clause will cut off.
			// so, every clause need `or` or `and` for last element.
			[4]string{"description", "like", kw, "or"},
			[4]string{"content", "like", kw, "or"},
		)
	}
	as, err := ar.data.DBClient.DatabaseClient.QueryAttribute().
		Where(cs...).All(ctx)
	if err != nil {
		return nil, err
	}
	bizas := &biz.Attributes{Collection: []*biz.Attribute{}}
	for _, a := range as.Collection {
		bizas.Collection = append(bizas.Collection, &biz.Attribute{
			Id:          a.Id,
			Path:        a.Path,
			Description: a.Description,
			UserId:      a.UserId,
			ArticleId:   a.ArticleId,
			UpdateTime:  timestamppb.New(a.UpdateTime),
		})
	}
	return bizas, nil
}

func (ar *attributeRepo) CreateAttribute(ctx context.Context, attribute *biz.Attribute) (*biz.Attribute, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := ar.data.DBClient.DatabaseClient.
		InsertAttribute(ctx, &mariadb.Attribute{
			Path:        attribute.Path,
			Description: attribute.Description,
			UserId:      attribute.UserId,
			ArticleId:   attribute.ArticleId,
		}); err != nil {
		return nil, err
	}
	return attribute, nil
}

func (ar *attributeRepo) UpdateAttribute(ctx context.Context, attribute *biz.Attribute) (*biz.Attribute, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := ar.data.DBClient.DatabaseClient.
		UpdateAttribute(ctx, &mariadb.Attribute{
			Id:          attribute.Id,
			Path:        attribute.Path,
			Description: attribute.Description,
			UserId:      attribute.UserId,
			ArticleId:   attribute.ArticleId,
		}); err != nil {
		return nil, err
	}
	return attribute, nil
}

func (ar *attributeRepo) DeleteAttribute(ctx context.Context, name string) error {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^attributes/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return errors.New("name cannot match regex express")
	}
	return ar.data.DBClient.DatabaseClient.DeleteAttribute(ctx, x[1])
}
