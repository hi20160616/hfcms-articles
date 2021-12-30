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
	pb.UnimplementedArticleAPIServer
	ac *biz.ArticleUsecase
}

func InitArticleService() *ArticleService {
	dbc := mariadb.NewClient()
	db := &data.Data{DBClient: dbc}
	repo := data.NewArticleRepo(db, log.Default())
	ac := biz.NewArticleUsecase(repo, *log.Default())
	return &ArticleService{ac: ac}
}

func (as *ArticleService) ListArticles(ctx context.Context, in *pb.ListArticlesRequest, msTitle string) (*pb.ListArticlesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListArticles: \n%v\n", r)
		}
	}()
	bizas, err := as.ac.ListArticles(ctx, "")
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

func GetArticle(ctx context.Context, in *pb.GetArticleRequest, msTitle string) (*pb.Article, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in GetArticle: %s\n%v\n", in.Name, r)
		}
	}()

	// ar := data.NewArticleRepo(&data.Data{MsTitle: msTitle}, log.Default)
	// a, err := ar.GetArticle(ctx, in.Id)
	// if err != nil {
	//         return nil, err
	// }
	// return &pb.Article{
	//         Id:            a.Id,
	//         Title:         a.Title,
	//         Content:       a.Content,
	//         WebsiteId:     a.WebsiteId,
	//         WebsiteDomain: a.WebsiteDomain,
	//         WebsiteTitle:  a.WebsiteTitle,
	//         UpdateTime:    a.UpdateTime,
	// }, nil
	return nil, nil
}

func SearchArticles(ctx context.Context, in *pb.SearchArticlesRequest) (*pb.SearchArticlesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in SearchArticles: \n%v\n", r)
		}
	}()
	// ar := data.NewArticleRepo(&data.Data{}, log.Default())
	// as, err := ar.SearchArticles(ctx, in.Keyword)
	// if err != nil {
	//         return nil, err
	// }
	// resp := []*pb.Article{}
	// for _, a := range as {
	//         resp = append(resp, &pb.Article{
	//                 Id:            a.Id,
	//                 Title:         a.Title,
	//                 Content:       a.Content,
	//                 WebsiteId:     a.WebsiteId,
	//                 WebsiteDomain: a.WebsiteDomain,
	//                 WebsiteTitle:  a.WebsiteTitle,
	//                 UpdateTime:    a.UpdateTime,
	//         })
	// }
	// return &pb.SearchArticlesResponse{Articles: resp}, nil
	return nil, nil
}
