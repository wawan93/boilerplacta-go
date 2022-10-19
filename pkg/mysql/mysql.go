package mysql

import (
	"database/sql"
	"log"
	"time"

	"github.com/Masterminds/squirrel"
	_ "github.com/go-sql-driver/mysql"
)

const (
	_defaultConnMaxLifetime = time.Minute * 3
	_defaultMaxOpenConns    = 10
	_defaultMaxIdleConns    = 10
)

// MySQL -.
type MySQL struct {
	connMaxLifetime time.Duration
	maxOpenConns    int
	maxIdleConns    int

	builder    squirrel.StatementBuilderType
	connection *sql.DB
}

// New -.
func New(url string, opts ...Option) (*MySQL, error) {
	db := &MySQL{
		connMaxLifetime: _defaultConnMaxLifetime,
		maxOpenConns:    _defaultMaxOpenConns,
		maxIdleConns:    _defaultMaxIdleConns,
	}

	for _, opt := range opts {
		opt(db)
	}

	db.builder = squirrel.StatementBuilder.PlaceholderFormat(squirrel.Question)

	conn, err := sql.Open("mysql", url)
	if err != nil {
		return nil, err
	}

	conn.SetConnMaxLifetime(time.Minute * 3)
	conn.SetMaxOpenConns(10)
	conn.SetMaxIdleConns(10)

	db.connection = conn

	return db, nil
}

// Builder -.
func (db *MySQL) Builder() squirrel.StatementBuilderType {
	return db.builder
}

// Connection -.
func (db *MySQL) Connection() *sql.DB {
	return db.connection
}

// Close -.
func (db *MySQL) Close() {
	if db.connection != nil {
		if err := db.connection.Close(); err != nil {
			log.Println(err)
		}
	}
}
