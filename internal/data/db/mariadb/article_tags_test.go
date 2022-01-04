package mariadb

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func TestInsertArticleTag(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	tag1 := &ArticleTag{
		ArticleId: "211227174018.87977400001",
		TagId:     1,
	}
	tag2 := &ArticleTag{
		ArticleId: "211227174018.87980300002",
		TagId:     2,
	}
	tag3 := &ArticleTag{
		ArticleId: "211227174018.87980400003",
		TagId:     3,
	}
	if err := c.DatabaseClient.InsertArticleTag(context.Background(), tag1); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertArticleTag(context.Background(), tag2); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertArticleTag(context.Background(), tag3); err != nil {
		t.Error(err)
	}
}

func TestListArticleTags(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryArticleTag().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestWhereArticleTags(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	as := [][4]string{
		{"article_id", "like", "0001"},
		{"id", "=", "2"},
		{"NotExist", "=", "3"}, // field not exist
	}

	for _, a := range as {
		fmt.Println("-------------------------------------------")
		fmt.Println("test where: ", a)
		got, err := c.DatabaseClient.QueryArticleTag().Where(a).All(context.Background())
		if err != nil {
			if driverErr, ok := err.(*mysql.MySQLError); ok {
				if driverErr.Number == 1054 && a[0] == "NotExist" {
					return
				}
			}
			t.Errorf("%v", err)
			return
		}
		for _, e := range got.Collection {
			fmt.Println(e)
		}
		fmt.Println("===========================================")
	}
}

func TestUpdateArticleTag(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	tag := &ArticleTag{
		Id:        1,
		ArticleId: "211227174018.87977400004",
		TagId:     4,
	}
	if err := c.DatabaseClient.UpdateArticleTag(context.Background(), tag); err != nil {
		t.Error(err)
		return
	}
	ps := [4]string{"id", "=", "1"}
	got, err := c.DatabaseClient.QueryArticleTag().Where(ps).First(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(got)
}

func TestDeleteArticleTag(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	id := "1"
	if err := c.DatabaseClient.DeleteArticleTag(context.Background(), id); err != nil {
		t.Error(err)
		return
	}

	ps := [4]string{"id", "=", id}
	got, err := c.DatabaseClient.QueryArticleTag().Where(ps).First(context.Background())
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
