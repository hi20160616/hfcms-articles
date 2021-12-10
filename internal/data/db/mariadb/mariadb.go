package mariadb

import (
	"database/sql"

	"github.com/hi20160616/hfcms-articles/configs"
)

type Client struct {
	db             *sql.DB
	ArticlesClient *ArticlesClient
	Err            error
}

type ArticlesClient struct {
	db *sql.DB
}

func open(cfg *configs.Config) (*sql.DB, error) {
	return sql.Open(cfg.Database.Driver, cfg.Database.Source)
}

func NewClient() *Client {
	cfg := configs.NewConfig("hfcms-articles")
	db, err := open(cfg)
	return &Client{db, &ArticlesClient{db}, err}
}
