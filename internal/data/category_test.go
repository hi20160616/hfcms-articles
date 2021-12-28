package data

import (
	"context"
	"fmt"
	"log"
	"testing"

	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
)

var cr = NewCategoryRepo(&Data{DBClient: mariadb.NewClient()}, *log.Default())

func TestListCategories(t *testing.T) {
	cs, err := cr.ListCategories(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range cs.Collection {
		fmt.Println(c)
	}
}
