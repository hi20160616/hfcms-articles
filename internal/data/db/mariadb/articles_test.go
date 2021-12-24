package mariadb

import (
	"context"
	"fmt"
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
