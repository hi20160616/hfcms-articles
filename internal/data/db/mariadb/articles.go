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

type Article struct {
	Id, Name, Title, Content string
	UserId, CategoryId       int
	UpdateTime               time.Time
}

type Articles struct {
	Collection []*Article
}

type ArticleQuery struct {
	db       *sql.DB
	limit    *int
	offset   *int
	query    string
	clauses  [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertArticle(ctx context.Context, article *Article) error {
	q := `INSERT INTO articles(
		id, title, content, category_id, user_id
		) VALUES (?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		title=?, content=?, category_id=?, user_id=?`
	aq := &ArticleQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query,
		article.Id, article.Title, article.Content, article.CategoryId, article.UserId,
		article.Title, article.Content, article.CategoryId, article.UserId)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}

func (dc *DatabaseClient) UpdateArticle(ctx context.Context, article *Article) error {
	q := `UPDATE articles SET title=?, content=?, category_id=?, user_id=?
		WHERE id=?`
	aq := &ArticleQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, article.Title, article.Content,
		article.CategoryId, article.UserId, article.Id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) DeleteArticle(ctx context.Context, id string) error {
	q := `DELETE FROM articles WHERE id=?`
	aq := &ArticleQuery{db: dc.db, query: q}
	_, err := aq.db.Exec(aq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryArticle() *ArticleQuery {
	return &ArticleQuery{db: dc.db, query: "SELECT * FROM articles"}
}

func (aq *ArticleQuery) All(ctx context.Context) (*Articles, error) {
	if err := aq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := aq.db.Query(aq.query, aq.args...)
	// rows, err := aq.db.Query("SELECT * FROM articles WHERE title like ?", "%%test%%")
	if err != nil {
		return nil, err
	}
	return mkArticle(rows)
}

func (aq *ArticleQuery) First(ctx context.Context) (*Article, error) {
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
func (aq *ArticleQuery) Where(cs ...[4]string) *ArticleQuery {
	aq.clauses = append(aq.clauses, cs...)
	return aq
}

func (aq *ArticleQuery) Order(condition string) *ArticleQuery {
	aq.order = condition
	return aq
}

func (aq *ArticleQuery) Limit(limit int) *ArticleQuery {
	aq.limit = &limit
	return aq
}

func (aq *ArticleQuery) Offset(offset int) *ArticleQuery {
	aq.offset = &offset
	return aq
}

func (aq *ArticleQuery) prepareQuery(ctx context.Context) error {
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

func mkArticle(rows *sql.Rows) (*Articles, error) {
	var id, title, content, category_id sql.NullString
	var update_time sql.NullTime
	var user_id int
	var articles = &Articles{}
	for rows.Next() {
		if err := rows.Scan(&id, &title, &content, &category_id, &user_id, &update_time); err != nil {
			return nil, errors.WithMessage(err, "mkArticle rows.Scan error")
		}
		articles.Collection = append(articles.Collection, &Article{
			Id:         id.String,
			Title:      title.String,
			Content:    content.String,
			UpdateTime: update_time.Time,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkArticle error")
	}
	return articles, nil
}
