package service

import (
	"context"
	"fmt"
	"testing"
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
