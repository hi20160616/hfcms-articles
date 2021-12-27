package data

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
)

func TestListArticles(t *testing.T) {
	data := mariadb.NewClient()
	ar := NewArticleRepo(&Data{DBClient: data}, *log.Default())
	as, err := ar.ListArticles(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range as.Collection {
		fmt.Println(a)
	}
}
