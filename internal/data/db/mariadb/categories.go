package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"
	"time"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type Category struct {
	Id         int
	Name, Code string
	UpdateTime time.Time
}

type Categories struct {
	Collection []*Category
}

type CategoryQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertCategory(ctx context.Context, category *Category) error {
	q := `INSERT INTO categories(
		name, code
		) VALUES (?, ?)
		ON DUPLICATE KEY UPDATE
		name=?, code=?`
	cq := &CategoryQuery{db: dc.db, query: q}
	_, err := cq.db.Exec(cq.query,
		category.Name, category.Code,
		category.Name, category.Code)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}

func (dc *DatabaseClient) UpdateCategory(ctx context.Context, category *Category) error {
	q := `UPDATE categories SET name=?, code=? WHERE id=?`
	cq := &CategoryQuery{db: dc.db, query: q}
	_, err := cq.db.Exec(cq.query, category.Name, category.Code, category.Id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) DeleteCategory(ctx context.Context, id string) error {
	q := `DELETE FROM categories WHERE id=?`
	cq := &CategoryQuery{db: dc.db, query: q}
	_, err := cq.db.Exec(cq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryCategory() *CategoryQuery {
	return &CategoryQuery{db: dc.db, query: "SELECT * FROM categories"}
}

func mkCategory(rows *sql.Rows) (*Categories, error) {
	var name, code sql.NullString
	var id int
	var update_time sql.NullTime
	var categories = &Categories{}
	for rows.Next() {
		if err := rows.Scan(&id, &name, &code, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkCategory rows.Scan error")
		}
		categories.Collection = append(categories.Collection, &Category{
			Id:         id,
			Name:       name.String,
			Code:       code.String,
			UpdateTime: update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkCategory error")
	}
	return categories, nil
}

func (cq *CategoryQuery) All(ctx context.Context) (*Categories, error) {
	if err := cq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := cq.db.Query(cq.query, cq.args...)
	// rows, err := cq.db.Query("SELECT * FROM categories WHERE title like ?", "%%test%%")
	if err != nil {
		return nil, err
	}
	return mkCategory(rows)
}

func (cq *CategoryQuery) First(ctx context.Context) (*Category, error) {
	nodes, err := cq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes.Collection) == 0 {
		return nil, ErrNotFound
	}
	return nodes.Collection[0], nil
}

// cs: {["name", "=", "jack", "and"], ["title", "like", "anything", ""]}
// the last `or` or `and` in clause will cut off after prepareQuery().
// so, every clause need `or` or `and` for last element.
func (cq *CategoryQuery) Where(cs ...[4]string) *CategoryQuery {
	cq.clauses = append(cq.clauses, cs...)
	return cq
}

func (cq *CategoryQuery) Order(condition string) *CategoryQuery {
	cq.order = condition
	return cq
}

func (cq *CategoryQuery) Limit(limit int) *CategoryQuery {
	cq.limit = &limit
	return cq
}

func (cq *CategoryQuery) Offset(offset int) *CategoryQuery {
	cq.offset = &offset
	return cq
}

func (cq *CategoryQuery) prepareQuery(ctx context.Context) error {
	if cq.clauses != nil {
		cq.query += " WHERE "
		for i, c := range cq.clauses {
			// the last `or` or `and` in clause will cut off there.
			// so, every clause need `or` or `and` for last element.
			if i == len(cq.clauses)-1 {
				cq.query += fmt.Sprintf(" %s %s ?", c[0], c[1])
			} else {
				cq.query += fmt.Sprintf(" %s %s ? %s", c[0], c[1], c[3])
			}
			if strings.ToLower(c[1]) == "like" {
				c[2] = fmt.Sprintf("%%%s%%", c[2])
			} else {
				c[2] = fmt.Sprintf("%s", c[2])
			}
			cq.args = append(cq.args, c[2])
		}
	}
	if cq.order != "" {
		cq.query += " ORDER BY ?"
		cq.args = append(cq.args, cq.order)
	}
	if cq.limit != nil {
		cq.query += " LIMIT ?"
		a := strconv.Itoa(*cq.limit)
		cq.args = append(cq.args, a)
	}
	if cq.offset != nil {
		cq.query += ", ?"
		a := strconv.Itoa(*cq.offset)
		cq.args = append(cq.args, a)
	}
	return nil
}
