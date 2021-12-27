package mariadb

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
)

func TestPrepareQuery(t *testing.T) {
	qc := &ArticleQuery{query: "SELECT * FROM articles"}
	qc.Where([4]string{"name", "like", "test"})
	if err := qc.prepareQuery(context.Background()); err != nil {
		t.Error(err)
	}
	fmt.Println(qc.query, qc.args)
}

func TestInsert(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	article1 := &Article{
		Id:         time.Now().Format("060102150405.000000") + "00001",
		Title:      "test1 title",
		Content:    "test1 content",
		UpdateTime: time.Now(),
	}
	article2 := &Article{
		Id:         time.Now().Format("060102150405.000000") + "00002",
		Title:      "test2 title",
		Content:    "test2 content",
		UpdateTime: time.Now(),
	}
	article3 := &Article{
		Id:         time.Now().Format("060102150405.000000") + "00003",
		Title:      "test3 title",
		Content:    "test3 content",
		UpdateTime: time.Now(),
	}
	if err := c.DatabaseClient.InsertArticle(context.Background(), article1); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertArticle(context.Background(), article2); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertArticle(context.Background(), article3); err != nil {
		t.Error(err)
	}
}

func TestListArticles(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	got, err := c.DatabaseClient.QueryArticle().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestWhere(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}

	as := [][4]string{
		{"title", "like", "1"},
		{"title", "=", "test2 title"},
	}

	for _, a := range as {
		fmt.Println("-------------------------------------------")
		fmt.Println("test where: ", a)
		got, err := c.DatabaseClient.QueryArticle().Where(a).All(context.Background())
		if err != nil {
			t.Errorf("%v", err)
			return
		}
		for _, e := range got.Collection {
			fmt.Println(e)
		}
		fmt.Println("===========================================")
	}
}

func TestUpdateArticle(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Error(c.Err)
		return
	}
	article := &Article{
		Id:      "211227122641.15719600002",
		Title:   "Test title update",
		Content: "Test content update",
	}
	if err := c.DatabaseClient.UpdateArticle(context.Background(), article); err != nil {
		t.Error(err)
		return
	}
	ps := [4]string{"id", "=", article.Id}
	got, err := c.DatabaseClient.QueryArticle().Where(ps).First(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(got)
}

func TestDeleteArticle(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Error(c.Err)
		return
	}
	id := "211227122641.15719600002"
	if err := c.DatabaseClient.DeleteArticle(context.Background(), id); err != nil {
		t.Error(err)
		return
	}

	ps := [4]string{"id", "=", id}
	got, err := c.DatabaseClient.QueryArticle().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	}
	if got != nil {
		t.Error(errors.New("Delete failed."))
	}
}
