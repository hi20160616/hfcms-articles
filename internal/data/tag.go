package data

import (
	"context"
	"errors"
	"log"
	"regexp"
	"time"

	_ "github.com/hi20160616/hfcms-articles/api/articles/v1"
	_ "github.com/hi20160616/hfcms-articles/configs"
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
	"google.golang.org/protobuf/types/known/emptypb"
	"google.golang.org/protobuf/types/known/timestamppb"
)

var _ biz.TagRepo = new(tagRepo)

type tagRepo struct {
	data *Data
	log  *log.Logger
}

func NewTagRepo(data *Data, logger *log.Logger) biz.TagRepo {
	return &tagRepo{
		data: data,
		log:  log.Default(),
	}
}

func (tr *tagRepo) ListTags(ctx context.Context) (*biz.Tags, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	ts, err := tr.data.DBClient.DatabaseClient.QueryTag().All(ctx)
	if err != nil {
		return nil, err
	}
	bizas := &biz.Tags{Collection: []*biz.Tag{}}
	for _, tag := range ts.Collection {
		bizas.Collection = append(bizas.Collection, &biz.Tag{
			TagId:      tag.Id,
			TagName:    tag.Name,
			UpdateTime: timestamppb.New(tag.UpdateTime),
		})
	}
	return bizas, nil
}

func (tr *tagRepo) GetTag(ctx context.Context, name string) (*biz.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	// name=tags/1
	re := regexp.MustCompile(`^tags/(\d+)$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return nil, errors.New("name cannot match regex express")
	}
	id := x[1]
	cs := [][4]string{{"id", "=", id}}
	c, err := tr.data.DBClient.DatabaseClient.QueryTag().
		Where(cs...).First(ctx)
	if err != nil {
		return nil, err
	}
	return &biz.Tag{
		TagId:      c.Id,
		TagName:    c.Name,
		UpdateTime: timestamppb.New(c.UpdateTime),
	}, nil
}

func (tr *tagRepo) CreateTag(ctx context.Context, tag *biz.Tag) (*biz.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := tr.data.DBClient.DatabaseClient.
		InsertTag(ctx, &mariadb.Tag{Name: tag.TagName}); err != nil {
		return nil, err
	}
	return tag, nil
}

func (tr *tagRepo) UpdateTag(ctx context.Context, tag *biz.Tag) (*biz.Tag, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	if err := tr.data.DBClient.DatabaseClient.
		UpdateTag(ctx, &mariadb.Tag{
			Id:   tag.TagId,
			Name: tag.TagName,
		}); err != nil {
		return nil, err
	}
	return tag, nil
}

func (tr *tagRepo) DeleteTag(ctx context.Context, name string) (*emptypb.Empty, error) {
	ctx, cancel := context.WithTimeout(ctx, 50*time.Second)
	defer cancel()
	re := regexp.MustCompile(`^tags/(\d+)/delete$`)
	x := re.FindStringSubmatch(name)
	if len(x) != 2 {
		return &emptypb.Empty{}, errors.New("name cannot match regex express")
	}
	return &emptypb.Empty{},
		tr.data.DBClient.DatabaseClient.DeleteTag(ctx, x[1])
}
