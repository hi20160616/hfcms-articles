package mariadb

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func TestInsertArticleAttribute(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	tag1 := &ArticleAttribute{
		ArticleId:   "211227174018.87977400001",
		AttributeId: 1,
	}
	tag2 := &ArticleAttribute{
		ArticleId:   "211227174018.87980300002",
		AttributeId: 2,
	}
	tag3 := &ArticleAttribute{
		ArticleId:   "211227174018.87980400003",
		AttributeId: 3,
	}
	if err := c.DatabaseClient.InsertArticleAttribute(context.Background(), tag1); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertArticleAttribute(context.Background(), tag2); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertArticleAttribute(context.Background(), tag3); err != nil {
		t.Error(err)
	}
}

func TestListArticleAttributes(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryArticleAttribute().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestWhereArticleAttributes(t *testing.T) {
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
		got, err := c.DatabaseClient.QueryArticleAttribute().Where(a).All(context.Background())
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

func TestUpdateArticleAttribute(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	tag := &ArticleAttribute{
		Id:          1,
		ArticleId:   "211227174018.87977400004",
		AttributeId: 4,
	}
	if err := c.DatabaseClient.UpdateArticleAttribute(context.Background(), tag); err != nil {
		t.Error(err)
		return
	}
	ps := [4]string{"id", "=", "1"}
	got, err := c.DatabaseClient.QueryArticleAttribute().Where(ps).First(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(got)
}

func TestDeleteArticleAttribute(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	id := "1"
	if err := c.DatabaseClient.DeleteArticleAttribute(context.Background(), id); err != nil {
		t.Error(err)
		return
	}

	ps := [4]string{"id", "=", id}
	got, err := c.DatabaseClient.QueryArticleAttribute().Where(ps).First(context.Background())
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
