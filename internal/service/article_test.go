package service

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/hi20160616/hfcms-articles/internal/biz"
)

var us = NewArticleService()

func TestCreateArticles(t *testing.T) {
	a, err := us.Acase.CreateArticle(context.Background(), &biz.Article{
		Title:      "Test CreateArticle Service",
		Content:    "Test CreateArticle Service Content",
		CategoryId: 123,
		UserId:     123,
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestListArticles(t *testing.T) {
	as, err := us.Acase.ListArticles(context.Background(), "")
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
	a, err := us.Acase.GetArticle(context.Background(), "articles/"+id)
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestSearchArticles(t *testing.T) {
	name := "articles/test3/search"
	as, err := us.Acase.SearchArticles(context.Background(), name)
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range as.Collection {
		fmt.Println(v)
	}
}

func TestUpdateArticle(t *testing.T) {
	a, err := us.Acase.UpdateArticle(context.Background(), &biz.Article{
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

func TestDeleteArticle(t *testing.T) {
	id := "211229113754.21503300002"
	name := "articles/" + id + "/delete"
	if err := us.Acase.DeleteArticle(context.Background(), name); err != nil {
		t.Fatal(err)
	}
	_, err := us.Acase.GetArticle(context.Background(), "articles/"+id)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
}
