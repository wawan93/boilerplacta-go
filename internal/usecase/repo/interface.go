package repo

import (
	"database/sql"

	"github.com/Masterminds/squirrel"
)

type DB interface {
	Builder() squirrel.StatementBuilderType
	Connection() *sql.DB
}
