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
	db         *sql.DB
	limit      *int
	offset     *int
	query      string
	predicates [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order      string
	args       []interface{}
	keywords   []string
}

func (dc *DatabaseClient) InsertCategory(ctx context.Context, category *Category) error {
	q := `INSERT INTO categories(
		id, name, code
		) VALUES (?, ?, ?)
		ON DUPLICATE KEY UPDATE
		id=?, name=?, code=?`
	aq := &CategoryQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query,
		category.Id, category.Name, category.Code,
		category.Id, category.Name, category.Code)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}

func (dc *DatabaseClient) UpdateCategory(ctx context.Context, category *Category) error {
	q := `UPDATE categories SET name=?, code=? WHERE id=?`
	aq := &CategoryQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, category.Name, category.Code, category.Id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) DeleteCategory(ctx context.Context, id string) error {
	q := `DELETE FROM categories WHERE id=?`
	aq := &CategoryQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryCategory() *CategoryQuery {
	return &CategoryQuery{db: dc.db, query: "SELECT * FROM categories"}
}

func (aq *CategoryQuery) All(ctx context.Context) (*Categories, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := aq.db.Query(aq.query, aq.args...)
	// rows, err := aq.db.Query("SELECT * FROM categories WHERE title like ?", "%%test%%")
	if err != nil {
		return nil, err
	}
	return mkCategory(rows)
}

func (aq *CategoryQuery) First(ctx context.Context) (*Category, error) {
	nodes, err := aq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes.Collection) == 0 {
		return nil, ErrNotFound
	}
	return nodes.Collection[0], nil
}

// ps: {["name", "=", "jack", "and"], ["title", "like", "anything", ""]}
func (aq *CategoryQuery) Where(ps ...[4]string) *CategoryQuery {
	aq.predicates = append(aq.predicates, ps...)
	return aq
}

func (aq *CategoryQuery) Order(condition string) *CategoryQuery {
	aq.order = condition
	return aq
}

func (aq *CategoryQuery) Limit(limit int) *CategoryQuery {
	aq.limit = &limit
	return aq
}

func (aq *CategoryQuery) Offset(offset int) *CategoryQuery {
	aq.offset = &offset
	return aq
}

func (aq *CategoryQuery) prepareQuery(ctx context.Context) error {
	if aq.predicates != nil {
		aq.query += " WHERE "
		for _, p := range aq.predicates {
			aq.query += fmt.Sprintf(" %s %s ? %s", p[0], p[1], p[3])
			if strings.ToLower(p[1]) == "like" {
				p[2] = fmt.Sprintf("%%%s%%", p[2])
			} else {
				p[2] = fmt.Sprintf("%s", p[2])
			}
			aq.args = append(aq.args, p[2])
		}
	}
	if aq.order != "" {
		aq.query += " ORDER BY ?"
		aq.args = append(aq.args, aq.order)
	}
	if aq.limit != nil {
		aq.query += " LIMIT ?"
		a := strconv.Itoa(*aq.limit)
		aq.args = append(aq.args, a)
	}
	if aq.offset != nil {
		aq.query += ", ?"
		a := strconv.Itoa(*aq.offset)
		aq.args = append(aq.args, a)
	}
	return nil
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
		return nil, errors.WithMessage(err, "mkUDPPacket error")
	}
	return categories, nil
}
