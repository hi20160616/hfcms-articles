package service

import (
	"context"
	"fmt"
	"log"

	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"github.com/hi20160616/hfcms-articles/internal/data"
	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
)

type ArticleService struct {
	pb.UnimplementedArticlesAPIServer
	Acase *biz.ArticleUsecase
}

func NewArticleService() *ArticleService {
	dbc := mariadb.NewClient()
	db := &data.Data{DBClient: dbc}
	repo := data.NewArticleRepo(db, log.Default())
	ac := biz.NewArticleUsecase(repo, *log.Default())
	return &ArticleService{Acase: ac}
}

func (as *ArticleService) ListArticles(ctx context.Context, in *pb.ListArticlesRequest, msTitle string) (*pb.ListArticlesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListArticles: \n%v\n", r)
		}
	}()
	bizas, err := as.Acase.ListArticles(ctx, in.Parent)
	if err != nil {
		return nil, err
	}
	resp := []*pb.Article{}
	for _, a := range bizas.Collection {
		resp = append(resp, &pb.Article{
			ArticleId:  a.ArticleId,
			Title:      a.Title,
			Content:    a.Content,
			CategoryId: int32(a.CategoryId),
			UserId:     int32(a.UserId),
			UpdateTime: a.UpdateTime,
		})
	}
	return &pb.ListArticlesResponse{Articles: resp}, nil
}

func (as *ArticleService) GetArticle(ctx context.Context, in *pb.GetArticleRequest, msTitle string) (*pb.Article, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetArticle: %s\n%v\n", in.Name, r)
		}
	}()
	biza, err := as.Acase.GetArticle(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	return &pb.Article{
		ArticleId:  biza.ArticleId,
		Title:      biza.Title,
		Content:    biza.Content,
		CategoryId: int32(biza.CategoryId),
		UserId:     int32(biza.UserId),
		UpdateTime: biza.UpdateTime,
	}, nil
}

func (as *ArticleService) SearchArticles(ctx context.Context, in *pb.SearchArticlesRequest) (*pb.SearchArticlesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in SearchArticles: \n%v\n", r)
		}
	}()
	bizas, err := as.Acase.SearchArticles(ctx, in.Name)
	if err != nil {
		return nil, err
	}
	respAs := &pb.SearchArticlesResponse{}
	for _, a := range bizas.Collection {
		respAs.Articles = append(respAs.Articles, &pb.Article{
			ArticleId:  a.ArticleId,
			Title:      a.Title,
			Content:    a.Content,
			CategoryId: int32(a.CategoryId),
			UserId:     int32(a.UserId),
			UpdateTime: a.UpdateTime})
	}
	return respAs, nil
}

func (as *ArticleService) UpdateArticle(ctx context.Context, in *pb.UpdateArticleRequest) (*pb.Article, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateArticles: \n%v\n", r)
		}
	}()
	a, err := as.Acase.UpdateArticle(ctx, &biz.Article{
		ArticleId:  in.Article.ArticleId,
		Title:      in.Article.Title,
		Content:    in.Article.Content,
		CategoryId: int(in.Article.CategoryId),
		UserId:     int(in.Article.UserId),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Article{
		ArticleId:  a.ArticleId,
		Title:      a.Title,
		Content:    a.Content,
		CategoryId: int32(a.CategoryId),
		UserId:     int32(a.UserId),
		UpdateTime: a.UpdateTime,
	}, nil
}

func (as *ArticleService) DeleteArticle(ctx context.Context, in *pb.DeleteArticleRequest) error {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateArticles: \n%v\n", r)
		}
	}()
	if err := as.Acase.DeleteArticle(ctx, in.Name); err != nil {
		return err
	}
	return nil
}

func (as *ArticleService) CreateArticle(ctx context.Context, in *pb.CreateArticleRequest) (*pb.Article, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in UpdateArticles: \n%v\n", r)
		}
	}()
	a, err := as.Acase.CreateArticle(ctx, &biz.Article{
		ArticleId:  in.Article.ArticleId,
		Title:      in.Article.Title,
		Content:    in.Article.Content,
		CategoryId: int(in.Article.CategoryId),
		UserId:     int(in.Article.UserId),
	})
	if err != nil {
		return nil, err
	}
	return &pb.Article{
		ArticleId:  a.ArticleId,
		Title:      a.Title,
		Content:    a.Content,
		CategoryId: int32(a.CategoryId),
		UserId:     int32(a.UserId),
		UpdateTime: a.UpdateTime,
	}, nil
}
