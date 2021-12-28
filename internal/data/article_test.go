package data

import (
	"context"
	"fmt"
	"log"
	"strings"
	"testing"

	"github.com/hi20160616/hfcms-articles/internal/biz"
	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
)

var ar = NewArticleRepo(&Data{DBClient: mariadb.NewClient()}, *log.Default())

func TestListArticles(t *testing.T) {
	as, err := ar.ListArticles(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range as.Collection {
		fmt.Println(a)
	}
}

func TestGetArticle(t *testing.T) {
	a, err := ar.GetArticle(context.Background(), "articles/211227122641.15716700001")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestSearchArticles(t *testing.T) {
	out := func(name string) {
		fmt.Println("name: ", name)
		as, err := ar.SearchArticles(context.Background(), name)
		if err != nil {
			t.Error(err)
			return
		}
		for _, a := range as.Collection {
			fmt.Println(a)
		}
	}

	names := []string{
		"articles/test1/search",
		"articles/test1 test2/search",
	}
	for _, n := range names {
		out(n)
	}
}

func TestCreateArticle(t *testing.T) {
	a, err := ar.CreateArticle(context.Background(), &biz.Article{
		Title:      "Test Create article title",
		Content:    "Test Create article content",
		CategoryId: 1,
		UserId:     1,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a.ArticleId)
}

func TestUpdateArticle(t *testing.T) {
	a, err := ar.UpdateArticle(context.Background(), &biz.Article{
		ArticleId:  "211228160638.1784271",
		Title:      "Test Update article title",
		Content:    "Test Update article content",
		CategoryId: 1,
		UserId:     1,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a.ArticleId)
}

func TestDeleteArticle(t *testing.T) {
	id := "211227122641.15716700001"
	name := "articles/" + id + "/delete"
	if err := ar.DeleteArticle(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := ar.GetArticle(context.Background(), "articles/"+id)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if a != nil {
		t.Error(fmt.Errorf("DeleteArticle failed."))
	}
}
