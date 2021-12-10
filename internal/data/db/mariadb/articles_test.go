package mariadb

import (
	"context"
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
		Id:         time.Now().Format("20060102150405000") + "0000001",
		Name:       "test1 name",
		Title:      "test1 title",
		Content:    "test1 content",
		UpdateTime: time.Now(),
	}
	article2 := &Article{
		Id:         time.Now().Format("20060102150405000") + "0000001",
		Name:       "test2 name",
		Title:      "test2 title",
		Content:    "test2 content",
		UpdateTime: time.Now(),
	}
	article3 := &Article{
		Id:         time.Now().Format("20060102150405000") + "0000001",
		Name:       "test3 name",
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
