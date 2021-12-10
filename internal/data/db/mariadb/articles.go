package mariadb

import (
	"context"
	"database/sql"
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
	q := "INSERT INTO udp_packets(id, name, title, content, update_time) VALUES(?,?,?,?,?)" +
		" ON DUPLICATE KEY UPDATE id=?, name=?, title=?, content=?, update_time=?"
	aq := &ArticleQuery{db: ac.db, query: q}
	_, err := aq.db.Exec(aq.query,
		article.Id, article.Name, article.Title, article.Content,
		article.UpdateTime, article.CategoryId, article.UserId,
		article.Id, article.Name, article.Title, article.Content,
		article.UpdateTime, article.CategoryId, article.UserId)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}
