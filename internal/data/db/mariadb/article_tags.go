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
	Id, TagId int
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
	tq := &ArticleTagQuery{db: dc.db, query: q}
	_, err := tq.db.Exec(tq.query, articleTag.ArticleId, articleTag.TagId)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}

func (dc *DatabaseClient) UpdateArticleTag(ctx context.Context, ArticleTag *ArticleTag) error {
	q := `UPDATE article_tags SET article_id=?, tag_id=? WHERE id=?`
	tq := &ArticleTagQuery{db: dc.db, query: q}
	_, err := tq.db.Exec(tq.query, ArticleTag.ArticleId, ArticleTag.TagId,
		ArticleTag.Id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) DeleteArticleTag(ctx context.Context, id string) error {
	q := `DELETE FROM article_tags WHERE id=?`
	tq := &ArticleTagQuery{db: dc.db, query: q}
	_, err := tq.db.Exec(tq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryArticleTag() *ArticleTagQuery {
	return &ArticleTagQuery{db: dc.db, query: "SELECT * FROM article_tags"}
}

func mkArticleTag(rows *sql.Rows) (*ArticleTags, error) {
	var id, tagId int
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

func (tq *ArticleTagQuery) All(ctx context.Context) (*ArticleTags, error) {
	if err := tq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := tq.db.Query(tq.query, tq.args...)
	// rows, err := tq.db.Query("SELECT * FROM article_tags WHERE title like ?", "%%test%%")
	if err != nil {
		return nil, err
	}
	return mkArticleTag(rows)
}

func (tq *ArticleTagQuery) First(ctx context.Context) (*ArticleTag, error) {
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
func (tq *ArticleTagQuery) Where(ps ...[4]string) *ArticleTagQuery {
	tq.clauses = append(tq.clauses, ps...)
	return tq
}

func (tq *ArticleTagQuery) Order(condition string) *ArticleTagQuery {
	tq.order = condition
	return tq
}

func (tq *ArticleTagQuery) Limit(limit int) *ArticleTagQuery {
	tq.limit = &limit
	return tq
}

func (tq *ArticleTagQuery) Offset(offset int) *ArticleTagQuery {
	tq.offset = &offset
	return tq
}

func (tq *ArticleTagQuery) prepareQuery(ctx context.Context) error {
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
