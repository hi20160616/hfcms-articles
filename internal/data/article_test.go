package data

import (
	"context"
	"fmt"
	"log"
	"testing"

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
