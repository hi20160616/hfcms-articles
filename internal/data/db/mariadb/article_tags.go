package mariadb

import (
	"context"
	"database/sql"
	"fmt"
	"strconv"
	"strings"

	_ "github.com/go-sql-driver/mysql"
	"github.com/pkg/errors"
)

type ArticleTag struct {
	Id        int64
	TagId     int
	ArticleId string
}

type ArticleTags struct {
	Collection []*ArticleTag
}

type ArticleTagQuery struct {
	db     *sql.DB
	limit  *int
	offset *int
	query  string
	// clauses: [ ["name", "=", "jack", "and"], ["title", "like", "anything", ""] ]
	clauses  [][4]string
	order    string
	args     []interface{}
	keywords []string
}

func (dc *DatabaseClient) InsertArticleTag(ctx context.Context, articleTag *ArticleTag) error {
	q := `INSERT INTO article_tags(article_id, tag_id) VALUES (?, ?)`
	atq := &ArticleTagQuery{db: dc.db, query: q}
	_, err := atq.db.Exec(atq.query, articleTag.ArticleId, articleTag.TagId)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}

func (dc *DatabaseClient) UpdateArticleTag(ctx context.Context, ArticleTag *ArticleTag) error {
	q := `UPDATE article_tags SET article_id=?, tag_id=? WHERE id=?`
	atq := &ArticleTagQuery{db: dc.db, query: q}
	_, err := atq.db.Exec(atq.query, ArticleTag.ArticleId, ArticleTag.TagId,
		ArticleTag.Id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) DeleteArticleTag(ctx context.Context, id string) error {
	q := `DELETE FROM article_tags WHERE id=?`
	atq := &ArticleTagQuery{db: dc.db, query: q}
	_, err := atq.db.Exec(atq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryArticleTag() *ArticleTagQuery {
	return &ArticleTagQuery{db: dc.db, query: "SELECT * FROM article_tags"}
}

func mkArticleTag(rows *sql.Rows) (*ArticleTags, error) {
	var id int64
	var tagId int
	var articleId sql.NullString
	var articleTags = &ArticleTags{}
	for rows.Next() {
		if err := rows.Scan(&id, &articleId, &tagId); err != nil {
			return nil, errors.WithMessage(err,
				"mkArticleTag rows.Scan error")
		}
		articleTags.Collection = append(articleTags.Collection, &ArticleTag{
			Id:        id,
			ArticleId: articleId.String,
			TagId:     tagId,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkArticleTag error")
	}
	return articleTags, nil
}

func (atq *ArticleTagQuery) All(ctx context.Context) (*ArticleTags, error) {
	if err := atq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := atq.db.Query(atq.query, atq.args...)
	// rows, err := atq.db.Query("SELECT * FROM article_tags WHERE title like ?", "%%test%%")
	if err != nil {
		return nil, err
	}
	return mkArticleTag(rows)
}

func (atq *ArticleTagQuery) First(ctx context.Context) (*ArticleTag, error) {
	nodes, err := atq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes.Collection) == 0 {
		return nil, ErrNotFound
	}
	return nodes.Collection[0], nil
}

// ps: {["name", "=", "jack", "and"], ["title", "like", "anything", ""]}
func (atq *ArticleTagQuery) Where(ps ...[4]string) *ArticleTagQuery {
	atq.clauses = append(atq.clauses, ps...)
	return atq
}

func (atq *ArticleTagQuery) Order(condition string) *ArticleTagQuery {
	atq.order = condition
	return atq
}

func (atq *ArticleTagQuery) Limit(limit int) *ArticleTagQuery {
	atq.limit = &limit
	return atq
}

func (atq *ArticleTagQuery) Offset(offset int) *ArticleTagQuery {
	atq.offset = &offset
	return atq
}

func (atq *ArticleTagQuery) prepareQuery(ctx context.Context) error {
	if atq.clauses != nil {
		atq.query += " WHERE "
		for _, p := range atq.clauses {
			atq.query += fmt.Sprintf(" %s %s ? %s", p[0], p[1], p[3])
			if strings.ToLower(p[1]) == "like" {
				p[2] = fmt.Sprintf("%%%s%%", p[2])
			} else {
				p[2] = fmt.Sprintf("%s", p[2])
			}
			atq.args = append(atq.args, p[2])
		}
	}
	if atq.order != "" {
		atq.query += " ORDER BY ?"
		atq.args = append(atq.args, atq.order)
	}
	if atq.limit != nil {
		atq.query += " LIMIT ?"
		a := strconv.Itoa(*atq.limit)
		atq.args = append(atq.args, a)
	}
	if atq.offset != nil {
		atq.query += ", ?"
		a := strconv.Itoa(*atq.offset)
		atq.args = append(atq.args, a)
	}
	return nil
}
