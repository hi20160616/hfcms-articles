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

var cr = func() biz.CategoryRepo {
	dc, err := mariadb.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return NewCategoryRepo(&Data{DBClient: dc}, log.Default())
}()

func TestCreateCategory(t *testing.T) {
	cs := []*biz.Category{
		{Name: "Root1", Code: "1"},
		{Name: "Root2", Code: "2"},
	}
	create := func(c *biz.Category) {
		a, err := cr.CreateCategory(context.Background(), c)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(a)
	}

	for _, v := range cs {
		create(v)
	}
}

func TestGetCategory(t *testing.T) {
	a, err := cr.GetCategory(context.Background(), "categories/4")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestUpdateCategory(t *testing.T) {
	a, err := cr.UpdateCategory(context.Background(), &biz.Category{
		CategoryId: 3,
		Name:       "Root+",
		Code:       "1+",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a)
}

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

func TestDeleteCategory(t *testing.T) {
	id := "2"
	name := "categories/" + id + "/delete"
	if err := cr.DeleteCategory(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := cr.GetCategory(context.Background(), "categories/"+id)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if a != nil {
		t.Error(fmt.Errorf("DeleteCategory failed."))
	}
}
