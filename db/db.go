package database

import (
	"github.com/jmoiron/sqlx"

	_ "github.com/lib/pq"
)

const (
	UniqueViolationErr string = "23505"
)

type Options struct {
	Driver, Params string
}

func Connect(options Options) *sqlx.DB {
	return sqlx.MustConnect(options.Driver, options.Params)
}
