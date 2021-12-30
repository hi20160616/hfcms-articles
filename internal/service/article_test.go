package service

import (
	"context"
	"fmt"
	"testing"
)

var us = InitArticleService()

func TestListArticles(t *testing.T) {
	as, err := us.ac.ListArticles(context.Background(), "")
	if err != nil {
		t.Error(err)
		return
	}
	for _, a := range as.Collection {
		fmt.Println(a)
	}
}
