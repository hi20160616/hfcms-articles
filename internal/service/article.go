package service

import (
	"context"
	"fmt"
	"log"

	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms-articles/internal/data"
)

type ArticleService struct {
}

func ListArticles(ctx context.Context, in *pb.ListArticlesRequest, msTitle string) (*pb.ListArticlesResponse, error) {
	defer func() {
		if r := recover(); r != nil {
			fmt.Printf("Recovered in ListArticles: \n%v\n", r)
		}
	}()
	ar := data.NewArticleRepo(&data.Data{}, log.Default())
	as, err := ar.ListArticles(ctx, "")
	// as, err := ar.ListArticles(ctx, "categories/"+msTitle+"/articles")
	if err != nil {
		return nil, err
	}
	resp := []*pb.Article{}
	for _, a := range as.Collection {
		resp = append(resp, &pb.Article{
			Title:      a.Title,
			Content:    a.Content,
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
