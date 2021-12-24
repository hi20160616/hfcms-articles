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
	Id, Name, Title, Content, CategoryId string
	UserId                               int
	UpdateTime                           time.Time
}

type Articles struct {
	Collection []*Article
}

type ArticleQuery struct {
	db         *sql.DB
	limit      *int
	offset     *int
	query      string
	predicates [][4]string // [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	order      string
	args       []interface{}
	keywords   []string
}

func (ac *ArticlesClient) Insert(ctx context.Context, article *Article) error {
	q := `INSERT INTO articles(
		id, title, content, update_time, category_id, user_id
		) VALUES (?, ?, ?, ?, ?, ?)
		ON DUPLICATE KEY UPDATE
		title=?, content=?, update_time=?, category_id=?, user_id=?`
	aq := &ArticleQuery{db: ac.db, query: q}
	_, err := aq.db.Exec(aq.query,
		article.Id, article.Title, article.Content,
		article.UpdateTime, article.CategoryId, article.UserId,
		article.Title, article.Content,
		article.UpdateTime, article.CategoryId, article.UserId)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}

func (ac *ArticlesClient) Update(ctx context.Context, article *Article) error {
	q := `UPDATE articles SET title=?, content=?, category_id=?, user_id=?
		WHERE id=?`
	aq := &ArticleQuery{db: ac.db, query: q}
	_, err := aq.db.Exec(aq.query, article.Title, article.Content,
		article.CategoryId, article.UserId, article.Id)
	if err != nil {
		return err
	}
	return nil
}

func (ac *ArticlesClient) Query() *ArticleQuery {
	return &ArticleQuery{db: ac.db, query: "SELECT * FROM articles"}
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

// ps: {["name", "=", "jack", "and"], ["title", "like", "anything", ""]}
func (aq *ArticleQuery) Where(ps ...[4]string) *ArticleQuery {
	aq.predicates = append(aq.predicates, ps...)
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
		return nil, errors.WithMessage(err, "mkUDPPacket error")
	}
	return articles, nil
}
