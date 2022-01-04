package main

import (
	"context"
	"flag"
	"fmt"
	"log"
	"testing"
	"time"

	v1 "github.com/hi20160616/hfcms-articles/api/articles/v1"
	"google.golang.org/grpc"
)

func TestMain(t *testing.T) {
	flag.Parse()
	conn, err := grpc.Dial(*addr, grpc.WithInsecure())
	if err != nil {
		log.Fatal(err)
	}
	defer conn.Close()
	c := v1.NewArticlesAPIClient(conn)

	ctx, cancel := context.WithTimeout(context.Background(), time.Second)
	defer cancel()
	r, err := c.ListArticles(ctx, &v1.ListArticlesRequest{Parent: ""})
	if err != nil {
		log.Fatal(err)
	}
	fmt.Println(r.Articles)
}
