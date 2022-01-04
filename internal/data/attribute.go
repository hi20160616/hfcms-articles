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
	"google.golang.org/protobuf/types/known/emptypb"
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

// by article_id: `attributes/aid/211229113754.21500200001/search`
// by user_id: `attributes/uid/1/search`
// fuzz description: `attributes/desc/kw1,kw2,kw3/search`
// by article_id and user_id: `attributes/aid/211229113754.21500200001/uid/1/search`
// by user_id and article_id and fuzz:
// `attributes/aid/211229113754.21500200001/uid/1/desc/kw1,kw2,kw3/search`
func (ar *attributeRepo) SearchAttributes(
	ctx context.Context, name string) (*biz.Attributes, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	e := `^attributes/(aid/([\d.]+))?(/)?(uid/(\d+))?(/)?(desc/([^^^#]+))?/search$`
	re := regexp.MustCompile(e)
	x := re.FindStringSubmatch(name)
	// aid: x[2], uid: x[5], kws: x[8]
	if len(x) != 9 {
		return nil, errors.New("name cannot match regex express")
	}
	aid, uid, kws := x[2], x[5], strings.Split(strings.TrimSpace(strings.ReplaceAll(
		x[8], "　", " ")), ",")
	cs := [][4]string{}
	if aid != "" {
		cs = append(cs, [4]string{"article_id", "=", aid, "and"})
	}
	if uid != "" {
		cs = append(cs, [4]string{"user_id", "=", uid, "and"})
	}
	for _, kw := range kws {
		cs = append(cs,
			// cs will be filtered by Where(clauses...)
			// the last `or` `and` in clause will cut off.
			// so, every clause need `or` or `and` for last element.
			[4]string{"description", "like", kw, "or"},
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

func (ar *attributeRepo) CreateAttribute(
	ctx context.Context, attribute *biz.Attribute) (*biz.Attribute, error) {
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

func (ar *attributeRepo) UpdateAttribute(
	ctx context.Context, attribute *biz.Attribute) (*biz.Attribute, error) {
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

func (ar *attributeRepo) DeleteAttribute(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^attributes/([\d.]+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{}, errors.New("name cannot match regex express")
	}
	return &emptypb.Empty{},
		ar.data.DBClient.DatabaseClient.DeleteAttribute(ctx, x[1])
}
