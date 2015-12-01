package database

import (
	"database/sql"

	"github.com/jmoiron/sqlx"
)

func ExecStruct(db *sqlx.DB, statement string, data interface{}) (sql.Result, error) {
	return db.NamedExec(statement, data)
}
