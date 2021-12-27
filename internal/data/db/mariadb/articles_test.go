package mariadb

import (
	"context"
	"fmt"
	"strings"
	"testing"
	"time"
)

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
	if err := c.Articles.Insert(context.Background(), article1); err != nil {
		t.Error(err)
	}
	if err := c.Articles.Insert(context.Background(), article2); err != nil {
		t.Error(err)
	}
	if err := c.Articles.Insert(context.Background(), article3); err != nil {
		t.Error(err)
	}
}

func TestListArticles(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	got, err := c.Articles.Query().All(context.Background())
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
		got, err := c.Articles.Query().Where(a).All(context.Background())
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

func TestPrepareQuery(t *testing.T) {
	aq := &ArticleQuery{query: "SELECT * FROM articles"}
	aq.Where([4]string{"name", "like", "test"})
	if err := aq.prepareQuery(context.Background()); err != nil {
		t.Error(err)
	}
	fmt.Println(aq.query, aq.args)
}

func TestUpdateArticle(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Error(c.Err)
		return
	}
	article := &Article{
		Id:      "211224161902.97258600003",
		Title:   "Test title update",
		Content: "Test content update",
	}
	if err := c.Articles.Update(context.Background(), article); err != nil {
		t.Error(err)
		return
	}
	ps := [4]string{"id", "=", "211224161902.97258600003"}
	got, err := c.Articles.Query().Where(ps).First(context.Background())
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
	id := "211224161902.97258600003"
	if err := c.Articles.Delete(context.Background(), id); err != nil {
		t.Error(err)
		return
	}

	ps := [4]string{"id", "=", "211224161902.97258600003"}
	got, err := c.Articles.Query().Where(ps).First(context.Background())
	if err != nil {
		if strings.Contains(err.Error(), "Item not found in table") {
			return
		}
		t.Error(err)
		return
	} else {
		fmt.Println("got :", got, "but need nothing!")
	}
}
