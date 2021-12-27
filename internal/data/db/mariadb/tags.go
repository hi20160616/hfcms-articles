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

type Tag struct {
	Id         int
	Name       string
	UpdateTime time.Time
}

type Tags struct {
	Collection []*Tag
}

type TagQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertTag(ctx context.Context, tag *Tag) error {
	q := `INSERT INTO tags(
		name
		) VALUES (?)
		ON DUPLICATE KEY UPDATE
		name=?`
	tq := &TagQuery{db: dc.db, query: q}
	_, err := tq.db.Exec(tq.query,
		tag.Name,
		tag.Name)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}

func (dc *DatabaseClient) UpdateTag(ctx context.Context, tag *Tag) error {
	q := `UPDATE tags SET name=? WHERE id=?`
	tq := &TagQuery{db: dc.db, query: q}
	_, err := tq.db.Exec(tq.query, tag.Name, tag.Id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) DeleteTag(ctx context.Context, id string) error {
	q := `DELETE FROM tags WHERE id=?`
	tq := &TagQuery{db: dc.db, query: q}
	_, err := tq.db.Exec(tq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryTag() *TagQuery {
	return &TagQuery{db: dc.db, query: "SELECT * FROM tags"}
}

func mkTag(rows *sql.Rows) (*Tags, error) {
	var id int
	var name sql.NullString
	var update_time sql.NullTime
	var tags = &Tags{}
	for rows.Next() {
		if err := rows.Scan(&id, &name, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkTag rows.Scan error")
		}
		tags.Collection = append(tags.Collection, &Tag{
			Id:         id,
			Name:       name.String,
			UpdateTime: update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkTag error")
	}
	return tags, nil
}

func (tq *TagQuery) All(ctx context.Context) (*Tags, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := tq.db.Query(tq.query, tq.args...)
	// rows, err := tq.db.Query("SELECT * FROM tags WHERE title like ?", "%%test%%")
	if err != nil {
		return nil, err
	}
	return mkTag(rows)
}

func (tq *TagQuery) First(ctx context.Context) (*Tag, error) {
	nodes, err := tq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes.Collection) == 0 {
		return nil, ErrNotFound
	}
	return nodes.Collection[0], nil
}

// ps: {["name", "=", "jack", "and"], ["title", "like", "anything", ""]}
func (tq *TagQuery) Where(ps ...[4]string) *TagQuery {
	tq.clauses = append(tq.clauses, ps...)
	return tq
}

func (tq *TagQuery) Order(condition string) *TagQuery {
	tq.order = condition
	return tq
}

func (tq *TagQuery) Limit(limit int) *TagQuery {
	tq.limit = &limit
	return tq
}

func (tq *TagQuery) Offset(offset int) *TagQuery {
	tq.offset = &offset
	return tq
}

func (tq *TagQuery) prepareQuery(ctx context.Context) error {
	if tq.clauses != nil {
		tq.query += " WHERE "
		for _, p := range tq.clauses {
			tq.query += fmt.Sprintf(" %s %s ? %s", p[0], p[1], p[3])
			if strings.ToLower(p[1]) == "like" {
				p[2] = fmt.Sprintf("%%%s%%", p[2])
			} else {
				p[2] = fmt.Sprintf("%s", p[2])
			}
			tq.args = append(tq.args, p[2])
		}
	}
	if tq.order != "" {
		tq.query += " ORDER BY ?"
		tq.args = append(tq.args, tq.order)
	}
	if tq.limit != nil {
		tq.query += " LIMIT ?"
		a := strconv.Itoa(*tq.limit)
		tq.args = append(tq.args, a)
	}
	if tq.offset != nil {
		tq.query += ", ?"
		a := strconv.Itoa(*tq.offset)
		tq.args = append(tq.args, a)
	}
	return nil
}
