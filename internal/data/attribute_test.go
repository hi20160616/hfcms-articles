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

var attr = NewAttributeRepo(&Data{DBClient: mariadb.NewClient()}, log.Default())
var aid = "211229114147.23586100001"
var attrid = "4"

func TestCreateAttribute(t *testing.T) {
	a, err := attr.CreateAttribute(context.Background(), &biz.Attribute{
		Path:        "Test Create attribute title",
		Description: "Test Create attribute content",
		ArticleId:   aid,
		UserId:      1,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a)
}

func TestUpdateAttribute(t *testing.T) {
	a, err := attr.UpdateAttribute(context.Background(), &biz.Attribute{
		Id:          4,
		Path:        "Test Update Create attribute title",
		Description: "Test Update Create attribute content",
		ArticleId:   aid,
		UserId:      1,
	})
	if err != nil {
		t.Error(err)
	}
	fmt.Println(a)
}

func TestGetAttribute(t *testing.T) {
	a, err := attr.GetAttribute(context.Background(), "attributes/"+attrid)
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(a)
}

func TestSearchAttributes(t *testing.T) {
	out := func(name string) {
		fmt.Println("name: ", name)
		as, err := attr.SearchAttributes(context.Background(), name)
		if err != nil {
			t.Error(err)
			return
		}
		for _, a := range as.Collection {
			fmt.Println(a)
		}
	}

	names := []string{
		"attributes/test1/search",
		"attributes/test1 test2/search",
	}
	for _, n := range names {
		out(n)
	}
}

func TestListAttributes(t *testing.T) {
	as, err := attr.ListAttributes(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range as.Collection {
		fmt.Println(a)
	}
}

func TestDeleteAttribute(t *testing.T) {
	name := "attributes/" + attrid + "/delete"
	if err := attr.DeleteAttribute(context.Background(), name); err != nil {
		t.Error(err)
		return
	}
	a, err := attr.GetAttribute(context.Background(), "attributes/"+attrid)
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if a != nil {
		t.Error(fmt.Errorf("DeleteAttribute failed."))
	}
}
