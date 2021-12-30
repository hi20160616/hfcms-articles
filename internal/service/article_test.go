package service

import (
	"context"
	"fmt"
	"testing"

	"github.com/hi20160616/hfcms-articles/internal/biz"
)

var us = InitArticleService()

func TestListArticles(t *testing.T) {
	as, err := us.ac.ListArticles(context.Background(), "")
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range as.Collection {
		fmt.Println(a)
	}
}

func TestGetArticle(t *testing.T) {
	id := "211229113754.21503300002"
	a, err := us.ac.GetArticle(context.Background(), "articles/"+id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestSearchArticles(t *testing.T) {
	name := "articles/test3/search"
	as, err := us.ac.SearchArticles(context.Background(), name)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range as.Collection {
		fmt.Println(v)
	}
}

func TestUpdateArticle(t *testing.T) {
	a, err := us.ac.UpdateArticle(context.Background(), &biz.Article{
		ArticleId:  "211229113754.21503300002",
		Title:      "UpdateViaAS",
		Content:    "UpdateViaAS",
		CategoryId: 5,
		UserId:     6,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}
