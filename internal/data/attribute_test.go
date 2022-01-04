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

// var attr = NewAttributeRepo(&Data{DBClient: mariadb.NewClient()}, log.Default())
var attr = func() biz.AttributeRepo {
	dc, err := mariadb.NewClient()
	if err != nil {
		log.Fatal(err)
	}
	return NewAttributeRepo(&Data{DBClient: dc}, log.Default())
}()

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
		Id:          2,
		Path:        "Test Update Create attribute title",
		Description: "测试 Update Create attribute content",
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
	tcs := []struct {
		name string
	}{
		{"attributes/aid/211227174018.87980300002/search"},
		{"attributes/uid/2/search"},
		{"attributes/desc/2,test,3/search"},
		{"attributes/aid/211227174018.87980300002/uid/2/search"},
		{"attributes/aid/211227174018.87980300002/uid/2/desc/test,3,测试/search"},
	}

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

	for _, tc := range tcs {
		out(tc.name)
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
	if _, err := attr.DeleteAttribute(context.Background(), name); err != nil {
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
