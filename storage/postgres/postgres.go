package postgres

import (
	"database/sql"
	"fmt"
	_ "github.com/lib/pq"
)

type PostgresPostStorage struct {
	db *sql.DB
}

type PostgresCommentStorage struct {
	db *sql.DB
}

func NewPostgresPostStorage(connStr string) (*PostgresPostStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}
	return &PostgresPostStorage{db: db}, nil
}

func NewPostgresCommentStorage(connStr string) (*PostgresCommentStorage, error) {
	db, err := sql.Open("postgres", connStr)
	if err != nil {
		return nil, fmt.Errorf("failed to connect to PostgreSQL: %w", err)
	}
	if err := db.Ping(); err != nil {
		return nil, fmt.Errorf("failed to ping PostgreSQL: %w", err)
	}
	return &PostgresCommentStorage{db: db}, nil
}
