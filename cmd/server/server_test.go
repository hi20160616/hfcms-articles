package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"github.com/hi20160616/hfcms-articles/configs"
	"google.golang.org/grpc"
)

func TestGRPCServer(t *testing.T) {
	tt, err := time.ParseDuration("1s")
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), tt)
	defer cancel()

	cfg := configs.NewConfig("hfcms-articles")
	// Set up a connection to the server
	conn, err := grpc.Dial(cfg.API.GRPC.Addr, grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	c := pb.NewArticlesAPIClient(conn)
	as, err := c.ListArticles(ctx, &pb.ListArticlesRequest{Parent: ""})
	if err != nil {
		t.Fatal(err)
	}
	for _, a := range as.Articles {
		fmt.Printf("%-30s %-30s %-30s \n", a.ArticleId, a.Title, a.Content)
	}

	id := "211229113754.21503400003"
	a, err := c.GetArticle(context.Background(), &pb.GetArticleRequest{Name: "articles/" + id})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}
