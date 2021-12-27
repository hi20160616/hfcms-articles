package mariadb

import (
	"context"
	"fmt"
	"strings"
	"testing"

	"github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

func TestInsertCategory(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	category1 := &Category{
		Name: "category 1",
		Code: "1",
	}
	category2 := &Category{
		Name: "category 2",
		Code: "2",
	}
	category3 := &Category{
		Name: "category 3",
		Code: "3",
	}
	if err := c.DatabaseClient.InsertCategory(context.Background(), category1); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertCategory(context.Background(), category2); err != nil {
		t.Error(err)
	}
	if err := c.DatabaseClient.InsertCategory(context.Background(), category3); err != nil {
		t.Error(err)
	}
}

func TestListCategories(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}
	got, err := c.DatabaseClient.QueryCategory().All(context.Background())
	if err != nil {
		t.Errorf("%v", err)
		return
	}
	for _, e := range got.Collection {
		fmt.Println(e)
	}
}

func TestWhereCategories(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Errorf("%v", c.Err)
		return
	}

	as := [][4]string{
		{"name", "like", "1"},
		{"code", "=", "2"},
		{"NotExist", "=", "3"}, // field not exist
	}

	for _, a := range as {
		fmt.Println("-------------------------------------------")
		fmt.Println("test where: ", a)
		got, err := c.DatabaseClient.QueryCategory().Where(a).All(context.Background())
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

func TestCategoriesPrepareQuery(t *testing.T) {
	aq := &CategoryQuery{query: "SELECT * FROM articles"}
	aq.Where([4]string{"name", "like", "test"})
	if err := aq.prepareQuery(context.Background()); err != nil {
		t.Error(err)
	}
	fmt.Println(aq.query, aq.args)
}

func TestUpdateCategories(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Error(c.Err)
		return
	}
	category := &Category{
		Id:   1,
		Name: "Category 1 Updated",
		Code: "11",
	}
	if err := c.DatabaseClient.UpdateCategory(context.Background(), category); err != nil {
		t.Error(err)
		return
	}
	ps := [4]string{"id", "=", "1"}
	got, err := c.DatabaseClient.QueryCategory().Where(ps).First(context.Background())
	if err != nil {
		t.Error(err)
		return
	}
	fmt.Println(got)
}

func TestDeleteCategory(t *testing.T) {
	c := NewClient()
	if c.Err != nil {
		t.Error(c.Err)
		return
	}
	id := "1"
	if err := c.DatabaseClient.DeleteCategory(context.Background(), id); err != nil {
		t.Error(err)
		return
	}

	ps := [4]string{"id", "=", id}
	got, err := c.DatabaseClient.QueryCategory().Where(ps).First(context.Background())
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
