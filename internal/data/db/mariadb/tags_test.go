package mariadb

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func TestInsertTag(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	tag1 := &Tag{
		Name: "tag 1",
	}
	tag2 := &Tag{
		Name: "tag 2",
	}
	tag3 := &Tag{
		Name: "tag 3",
	}
	if err := c.DatabaseClient.InsertTag(context.Background(), tag1); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertTag(context.Background(), tag2); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertTag(context.Background(), tag3); err != nil {
		t.Error(err)
	}
}

func TestListTags(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	got, err := c.DatabaseClient.QueryTag().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestWhereTags(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}

	as := [][4]string{
		{"name", "like", "1"},
		{"id", "=", "2"},
		{"NotExist", "=", "3"}, // field not exist
	}

	for _, a := range as {
		fmt.Println("-------------------------------------------")
		fmt.Println("test where: ", a)
		got, err := c.DatabaseClient.QueryTag().Where(a).All(context.Background())
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

func TestUpdateTag(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Error(c.Err)
		return
	}
	tag := &Tag{
		Id:   1,
		Name: "Tag 1 Updated",
	}
	if err := c.DatabaseClient.UpdateTag(context.Background(), tag); err != nil {
		t.Error(err)
		return
	}
	ps := [4]string{"id", "=", "1"}
	got, err := c.DatabaseClient.QueryTag().Where(ps).First(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(got)
}

func TestDeleteTag(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Error(c.Err)
		return
	}
	id := "1"
	if err := c.DatabaseClient.DeleteTag(context.Background(), id); err != nil {
		t.Error(err)
		return
	}

	ps := [4]string{"id", "=", id}
	got, err := c.DatabaseClient.QueryTag().Where(ps).First(context.Background())
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
