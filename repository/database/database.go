package database

import (
	"context"

	"urlShortner/storage"
)

type Database interface {
	GetLongByShort(ctx context.Context, short string) (string, error)
	CreateNewURL(ctx context.Context, url *URL) error
}

type PostgresDB struct {
	db storage.Database
}

func NewPostgresDB(db storage.Database) *PostgresDB {
	return &PostgresDB{db: db}
}
