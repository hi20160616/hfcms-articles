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

var tr = func() biz.TagRepo {
	dc, err := mariadb.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return NewTagRepo(&Data{DBClient: dc}, log.Default())

}()

func TestCreateTag(t *testing.T) {
	cs := []*biz.Tag{
		{Name: "Tag1"},
		{Name: "Tag2"},
		{Name: "Tag3"},
	}
	create := func(c *biz.Tag) {
		a, err := tr.CreateTag(context.Background(), c)
		if err != nil {
			t.Error(err)
		}
		fmt.Println(a)
	}

	for _, v := range cs {
		create(v)
	}
}

func TestGetTag(t *testing.T) {
	a, err := tr.GetTag(context.Background(), "tags/4")
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestUpdateTag(t *testing.T) {
	a, err := tr.UpdateTag(context.Background(), &biz.Tag{
		TagId: 3,
		Name:  "Tag3+",
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a)
}

func TestListTags(t *testing.T) {
	cs, err := tr.ListTags(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	for _, c := range cs.Collection {
		fmt.Println(c)
	}
}

func TestDeleteTag(t *testing.T) {
	id := "2"
	name := "tags/" + id + "/delete"
	if _, err := tr.DeleteTag(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := tr.GetTag(context.Background(), "tags/"+id)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if a != nil {
		t.Error(fmt.Errorf("DeleteTag failed."))
	}
}
