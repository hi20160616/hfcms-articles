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

type Attribute struct {
	Id                           int64
	UserId                       int
	Path, Description, ArticleId string
	UpdateTime                   time.Time
}

type Attributes struct {
	Collection []*Attribute
}

type AttributeQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertAttribute(ctx context.Context,
	attribute *Attribute) error {
	q := `INSERT INTO attributes(
		path, description, user_id, article_id
		) VALUES (?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		path=?, description=?, user_id=?, article_id=?`
	aq := &AttributeQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query,
		attribute.Path, attribute.Description, attribute.UserId,
		attribute.ArticleId,
		attribute.Path, attribute.Description, attribute.UserId,
		attribute.ArticleId)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}

func (dc *DatabaseClient) UpdateAttribute(ctx context.Context,
	attribute *Attribute) error {
	q := `UPDATE attributes SET path=?, description=?, user_id=?, article_id=?
		WHERE id=?`
	aq := &AttributeQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, attribute.Path, attribute.Description,
		attribute.UserId, attribute.ArticleId, attribute.Id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) DeleteAttribute(ctx context.Context, id string) error {
	q := `DELETE FROM attributes WHERE id=?`
	aq := &AttributeQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryAttribute() *AttributeQuery {
	return &AttributeQuery{db: dc.db, query: "SELECT * FROM attributes"}
}

func mkAttribute(rows *sql.Rows) (*Attributes, error) {
	var id int64
	var user_id int
	var path, description, article_id sql.NullString
	var update_time sql.NullTime
	var attributes = &Attributes{}
	for rows.Next() {
		if err := rows.Scan(
			&id, &path, &description, &user_id, &article_id, &update_time,
		); err != nil {
			return nil, errors.WithMessage(err, "mkAttribute rows.Scan error")
		}
		attributes.Collection = append(attributes.Collection, &Attribute{
			Id:          id,
			Path:        path.String,
			Description: description.String,
			UserId:      user_id,
			ArticleId:   article_id.String,
			UpdateTime:  update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkAttribute error")
	}
	return attributes, nil
}

func (aq *AttributeQuery) All(ctx context.Context) (*Attributes, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := aq.db.Query(aq.query, aq.args...)
	// rows, err := aq.db.Query("SELECT * FROM attributes WHERE title like ?", "%%test%%")
	if err != nil {
		return nil, err
	}
	return mkAttribute(rows)
}

func (aq *AttributeQuery) First(ctx context.Context) (*Attribute, error) {
	nodes, err := aq.Limit(1).All(ctx)
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
func (aq *AttributeQuery) Where(cs ...[4]string) *AttributeQuery {
	aq.clauses = append(aq.clauses, cs...)
	return aq
}

func (aq *AttributeQuery) Order(condition string) *AttributeQuery {
	aq.order = condition
	return aq
}

func (aq *AttributeQuery) Limit(limit int) *AttributeQuery {
	aq.limit = &limit
	return aq
}

func (aq *AttributeQuery) Offset(offset int) *AttributeQuery {
	aq.offset = &offset
	return aq
}

func (aq *AttributeQuery) prepareQuery(ctx context.Context) error {
	if aq.clauses != nil {
		aq.query += " WHERE "
		for i, c := range aq.clauses {
			// the last `or` or `and` in clause will cut off there.
			// so, every clause need `or` or `and` for last element.
			if i == len(aq.clauses)-1 {
				aq.query += fmt.Sprintf(" %s %s ?", c[0], c[1])
			} else {
				aq.query += fmt.Sprintf(" %s %s ? %s", c[0], c[1], c[3])
			}
			if strings.ToLower(c[1]) == "like" {
				c[2] = fmt.Sprintf("%%%s%%", c[2])
			} else {
				c[2] = fmt.Sprintf("%s", c[2])
			}
			aq.args = append(aq.args, c[2])
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
