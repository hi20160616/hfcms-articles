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

type ArticleAttribute struct {
	Id, AttributeId int
	ArticleId       string
}

type ArticleAttributes struct {
	Collection []*ArticleAttribute
}

type ArticleAttributeQuery struct {
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

func (dc *DatabaseClient) InsertArticleAttribute(ctx context.Context, articleAttribute *ArticleAttribute) error {
	q := `INSERT INTO article_attributes(article_id, attribute_id) VALUES (?, ?)`
	aaq := &ArticleAttributeQuery{db: dc.db, query: q}
	_, err := aaq.db.Exec(aaq.query,
		articleAttribute.ArticleId, articleAttribute.AttributeId)
	if err != nil {
		return errors.WithMessage(err, "mariadb: Insert error")
	}
	return nil
}

func (dc *DatabaseClient) UpdateArticleAttribute(ctx context.Context, ArticleAttribute *ArticleAttribute) error {
	q := `UPDATE article_attributes SET article_id=?, attribute_id=? WHERE id=?`
	aaq := &ArticleAttributeQuery{db: dc.db, query: q}
	_, err := aaq.db.Exec(aaq.query, ArticleAttribute.ArticleId, ArticleAttribute.AttributeId,
		ArticleAttribute.Id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) DeleteArticleAttribute(ctx context.Context, id string) error {
	q := `DELETE FROM article_attributes WHERE id=?`
	aaq := &ArticleAttributeQuery{db: dc.db, query: q}
	_, err := aaq.db.Exec(aaq.query, id)
	if err != nil {
		return err
	}
	return nil
}

func (dc *DatabaseClient) QueryArticleAttribute() *ArticleAttributeQuery {
	return &ArticleAttributeQuery{db: dc.db, query: "SELECT * FROM article_attributes"}
}

func mkArticleAttribute(rows *sql.Rows) (*ArticleAttributes, error) {
	var id, attributeId int
	var articleId sql.NullString
	var articleAttributes = &ArticleAttributes{}
	for rows.Next() {
		if err := rows.Scan(&id, &articleId, &attributeId); err != nil {
			return nil, errors.WithMessage(err,
				"mkArticleAttribute rows.Scan error")
		}
		articleAttributes.Collection = append(articleAttributes.Collection, &ArticleAttribute{
			Id:          id,
			ArticleId:   articleId.String,
			AttributeId: attributeId,
		})
	}
	// TODO: to confirm code below can make sence.
	if err := rows.Err(); err != nil {
		return nil, errors.WithMessage(err, "mkArticleAttribute error")
	}
	return articleAttributes, nil
}

func (aaq *ArticleAttributeQuery) All(ctx context.Context) (*ArticleAttributes, error) {
	if err := aaq.prepareQuery(ctx); err != nil {
		return nil, err
	}
	rows, err := aaq.db.Query(aaq.query, aaq.args...)
	// rows, err := aaq.db.Query("SELECT * FROM article_attributes WHERE title like ?", "%%test%%")
	if err != nil {
		return nil, err
	}
	return mkArticleAttribute(rows)
}

func (aaq *ArticleAttributeQuery) First(ctx context.Context) (*ArticleAttribute, error) {
	nodes, err := aaq.Limit(1).All(ctx)
	if err != nil {
		return nil, err
	}
	if len(nodes.Collection) == 0 {
		return nil, ErrNotFound
	}
	return nodes.Collection[0], nil
}

// ps: {["name", "=", "jack", "and"], ["title", "like", "anything", ""]}
func (aaq *ArticleAttributeQuery) Where(ps ...[4]string) *ArticleAttributeQuery {
	aaq.clauses = append(aaq.clauses, ps...)
	return aaq
}

func (aaq *ArticleAttributeQuery) Order(condition string) *ArticleAttributeQuery {
	aaq.order = condition
	return aaq
}

func (aaq *ArticleAttributeQuery) Limit(limit int) *ArticleAttributeQuery {
	aaq.limit = &limit
	return aaq
}

func (aaq *ArticleAttributeQuery) Offset(offset int) *ArticleAttributeQuery {
	aaq.offset = &offset
	return aaq
}

func (aaq *ArticleAttributeQuery) prepareQuery(ctx context.Context) error {
	if aaq.clauses != nil {
		aaq.query += " WHERE "
		for _, p := range aaq.clauses {
			aaq.query += fmt.Sprintf(" %s %s ? %s", p[0], p[1], p[3])
			if strings.ToLower(p[1]) == "like" {
				p[2] = fmt.Sprintf("%%%s%%", p[2])
			} else {
				p[2] = fmt.Sprintf("%s", p[2])
			}
			aaq.args = append(aaq.args, p[2])
		}
	}
	if aaq.order != "" {
		aaq.query += " ORDER BY ?"
		aaq.args = append(aaq.args, aaq.order)
	}
	if aaq.limit != nil {
		aaq.query += " LIMIT ?"
		a := strconv.Itoa(*aaq.limit)
		aaq.args = append(aaq.args, a)
	}
	if aaq.offset != nil {
		aaq.query += ", ?"
		a := strconv.Itoa(*aaq.offset)
		aaq.args = append(aaq.args, a)
	}
	return nil
}
