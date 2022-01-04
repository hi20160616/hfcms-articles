package service

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	v1 "github.com/hi20160616/hfcms-articles/api/articles/v1"
)

var as = func() *ArticleService {
	us, err := NewArticleService()
	if err != nil {
		log.Fatal(err)
	}
	return us
}()

func TestCreateArticles(t *testing.T) {

	a, err := as.CreateArticle(context.Background(), &v1.CreateArticleRequest{
		Article: &v1.Article{
			Title:      "Test CreateArticle Service",
			Content:    "Test CreateArticle Service Content",
			CategoryId: 123,
			UserId:     123,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestListArticles(t *testing.T) {
	as, err := as.ListArticles(context.Background(), &v1.ListArticlesRequest{})
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range as.Articles {
		fmt.Println(a)
	}
}

func TestGetArticle(t *testing.T) {
	id := "211229113754.21503300002"
	a, err := as.GetArticle(context.Background(), &v1.GetArticleRequest{Name: "articles/" + id})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestSearchArticles(t *testing.T) {
	name := "articles/test3/search"
	articles, err := as.SearchArticles(context.Background(), &v1.SearchArticlesRequest{Name: name})
	if err != nil {
		t.Fatal(err)
	}
	for _, v := range articles.Articles {
		fmt.Println(v)
	}
}

func TestUpdateArticle(t *testing.T) {
	a, err := as.UpdateArticle(context.Background(), &v1.UpdateArticleRequest{
		Article: &v1.Article{
			ArticleId:  "211229113754.21503300002",
			Title:      "UpdateViaAS",
			Content:    "UpdateViaAS",
			CategoryId: 5,
			UserId:     6,
		},
	})
	if err != nil {
		t.Fatal(err)
	}
	fmt.Println(a)
}

func TestDeleteArticle(t *testing.T) {
	id := "211229113754.21503300002"
	name := "articles/" + id + "/delete"
	if _, err := as.DeleteArticle(context.Background(), &v1.DeleteArticleRequest{Name: name}); err != nil {
		t.Fatal(err)
	}
	_, err := as.GetArticle(context.Background(), &v1.GetArticleRequest{Name: "articles/" + id})
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
	}
}
