package data

import (
	"github.com/hi20160616/hfcms-articles/internal/biz"
	"github.com/hi20160616/hfcms-articles/internal/data/db/mariadb"
)

var _ biz.ArticleRepo = new(articleRepo)

type Data struct {
	DBClient *mariadb.Client
}
