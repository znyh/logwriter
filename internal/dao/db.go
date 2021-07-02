package dao

import (
	"github.com/go-kratos/kratos/pkg/conf/paladin"
	"github.com/go-kratos/kratos/pkg/database/sql"
)

func NewDB() (db *sql.DB, cf func(), err error) {
	var cfg struct {
		Client *sql.Config
	}
	if err = paladin.Get("db.txt").UnmarshalTOML(&cfg); err != nil {
		return
	}
	db = sql.NewMySQL(cfg.Client)
	cf = func() { _ = db.Close() }
	return
}
