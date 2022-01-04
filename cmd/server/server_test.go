package main

import (
	"context"
	"fmt"
	"testing"
	"time"

	pb "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"google.golang.org/grpc"
)

func TestGRPCServer(t *testing.T) {
	tt, err := time.ParseDuration("1s")
	if err != nil {
		t.Error(err)
	}
	ctx, cancel := context.WithTimeout(context.Background(), tt)
	defer cancel()

	// Set up a connection to the server
	conn, err := grpc.Dial(":9090", grpc.WithInsecure(), grpc.WithBlock())
	if err != nil {
		t.Error(err)
	}
	defer conn.Close()

	c := pb.NewArticlesAPIClient(conn)
	as, err := c.ListArticles(ctx, &pb.ListArticlesRequest{Parent: ""})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(as.Articles)
	// for _, a := range as.Articles {
	//         fmt.Printf("%s %20s %20s \n", a.ArticleId, a.Title, a.Content)
	// }

	// id := "211229113754.21503400003"
	// a, err := c.GetArticle(context.Background(), &pb.GetArticleRequest{Name: "articles/" + id})
	// if err != nil {
	//         t.Fatal(err)
	// }
	// fmt.Println(a)
}
