package mariadb

import (
	"database/sql"

	"github.com/hi20160616/hfcms-articles/configs"
)

type Client struct {
	db             *sql.DB
	DatabaseClient *DatabaseClient
}

type DatabaseClient struct {
	db *sql.DB
}

func open(cfg *configs.Config) (*sql.DB, error) {
	return sql.Open(cfg.Database.Driver, cfg.Database.Source)
	// return sql.Open("mysql", "hfcms_article_user:hfcms_article_user_pwd@tcp(127.0.0.1:3306)/hfcms_articles?loc=Asia%2FShanghai&parseTime=true")
}

func NewClient() (*Client, error) {
	cfg := configs.NewConfig("hfcms-articles")
	if cfg.Err != nil {
		return &Client{nil, nil}, cfg.Err
	}
	db, err := open(cfg)
	return &Client{db, &DatabaseClient{db}}, err
}
