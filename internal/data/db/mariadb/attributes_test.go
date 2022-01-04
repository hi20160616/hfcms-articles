package mariadb

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func TestInsertAttribute(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	attribute1 := &Attribute{
		Path:        "./Upload/1/test1.jpg",
		Description: "attribute 1 for test",
		UserId:      1,
		ArticleId:   "211227174018.87977400001",
	}
	attribute2 := &Attribute{
		Path:        "./Upload/2/test2.jpg",
		Description: "attribute 2 for test",
		UserId:      2,
		ArticleId:   "211227174018.87980300002",
	}
	attribute3 := &Attribute{
		Path:        "./Upload/3/test3.jpg",
		Description: "attribute 3 for test",
		UserId:      3,
		ArticleId:   "211227174018.87980400003",
	}
	if err := c.DatabaseClient.InsertAttribute(context.Background(), attribute1); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertAttribute(context.Background(), attribute2); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertAttribute(context.Background(), attribute3); err != nil {
		t.Error(err)
	}
}

func TestListAttributes(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	got, err := c.DatabaseClient.QueryAttribute().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestWhereAttributes(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}

	as := [][4]string{
		{"path", "like", "1"},
		{"id", "=", "2"},
		{"NotExist", "=", "3"}, // field not exist
	}

	for _, a := range as {
		fmt.Println("-------------------------------------------")
		fmt.Println("test where: ", a)
		got, err := c.DatabaseClient.QueryAttribute().Where(a).All(context.Background())
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

func TestUpdateAttribute(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	attribute := &Attribute{
		Id:          1,
		Path:        "./Upload/1/test1Updated.jpg",
		Description: "test update attribute 1",
		UserId:      2,
		ArticleId:   "211227174018.87977411111",
	}
	if err := c.DatabaseClient.UpdateAttribute(context.Background(), attribute); err != nil {
		t.Error(err)
		return
	}
	ps := [4]string{"id", "=", "1"}
	got, err := c.DatabaseClient.QueryAttribute().Where(ps).First(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(got)
}

func TestDeleteAttribute(t *testing.T) {
	c, err := NewClient()
	if err != nil {
		t.Fatal(err)
	}
	id := "1"
	if err := c.DatabaseClient.DeleteAttribute(context.Background(), id); err != nil {
		t.Error(err)
		return
	}

	ps := [4]string{"id", "=", id}
	got, err := c.DatabaseClient.QueryAttribute().Where(ps).First(context.Background())
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
