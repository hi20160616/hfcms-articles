package mariadb

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"

	"github.com/pkg/errors"
)

var id = "220104181701.65442600123"

func TestPrepareQuery(t *testing.T) {
	qc := &ArticleQuery{query: "SELECT * FROM articles"}
	qc.Where(
		[4]string{"name", "like", "test", "and"},
		[4]string{"name", "like", "test1", "and"},
		[4]string{"name", "like", "test2", "and"},
		[4]string{"name", "like", "test3", "and"},
	)
	if err := qc.prepareQuery(context.Background()); err != nil {
		t.Error(err)
	}
	fmt.Println(qc.query, qc.args)
}

func TestInsertArticle(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	article1 := &Article{
		Id:         time.Now().Format("060102150405.000000") + "00001",
		Title:      "test1 title",
		Content:    "test1 content",
		UserId:     1,
		CategoryId: 1,
		UpdateTime: time.Now(),
	}
	article2 := &Article{
		Id:         time.Now().Format("060102150405.000000") + "00002",
		Title:      "test2 title",
		Content:    "test2 content",
		UserId:     2,
		CategoryId: 2,
		UpdateTime: time.Now(),
	}
	article3 := &Article{
		Id:         time.Now().Format("060102150405.000000") + "00003",
		Title:      "test3 title",
		Content:    "test3 content",
		UserId:     3,
		CategoryId: 3,
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
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
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

func TestWhereArticles(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	out := func(a [4]string) {
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

	outs := func(ps [][4]string) {
		fmt.Println("-------------------------------------------")
		fmt.Println("test where: ", ps)
		got, err := c.DatabaseClient.QueryArticle().Where(ps...).All(context.Background())
		if err != nil {
			t.Error(err)
			return
		}
		// fmt.Println(got.Collection)
		for _, e := range got.Collection {
			fmt.Println(e)
		}
		fmt.Println("===========================================")
	}

	ps1 := [][4]string{
		{"title", "=", "test2 title", "and"},
		{"content", "like", "2", "and"},
		{"user_id", "=", "0"},
	}
	ps2 := [][4]string{
		{"title", "=", "test2 title", "or"},
		{"content", "like", "2", "or"},
		{"user_id", "=", "0"},
	}
	out(ps1[2]) // test one clause
	outs(ps1)
	outs(ps2)
}

func TestUpdateArticle(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	article := &Article{
		Id:         id,
		Title:      "Test title update",
		Content:    "Test content update",
		CategoryId: 5,
		UserId:     2,
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
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
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
