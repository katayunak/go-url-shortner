package storage

import (
	"errors"
	"fmt"

	"github.com/golang-migrate/migrate/v4"
	"gorm.io/gorm"

	_ "github.com/golang-migrate/migrate/v4/database/postgres"
	_ "github.com/golang-migrate/migrate/v4/source/file"
	_ "github.com/lib/pq"
	pg "gorm.io/driver/postgres"

	"urlShortner/config"
)

type postgres struct {
	config config.PostgresConfig
	db     *gorm.DB
}

func NewPostgres(conf config.PostgresConfig) Database {
	return &postgres{
		config: conf,
	}
}

func (p *postgres) Connect() error {
	dsn := fmt.Sprintf("host=%s user=%s password=%s dbname=%s port=%s sslmode=%s TimeZone=%s",
		p.config.Host, p.config.Username, p.config.Password, p.config.Database,
		p.config.Port, p.config.SSLMode, p.config.Timezone)

	db, err := gorm.Open(pg.Open(dsn), &gorm.Config{})
	if err != nil {
		return err
	}
	p.db = db

	return nil
}

func (p *postgres) MigrateUp() error {
	dsn := fmt.Sprintf(
		"postgres://%s:%s@%s:%s/%s?sslmode=%s",
		p.config.Username, p.config.Password, p.config.Host, p.config.Port,
		p.config.Database,
		p.config.SSLMode,
	)
	m, err := migrate.New(
		p.config.MigrationFile,
		dsn)
	if err != nil {
		return err
	}

	err = m.Up()
	if err != nil && !errors.Is(err, migrate.ErrNoChange) {
		return err
	}

	return nil
}

func (p *postgres) DB() *gorm.DB {
	return p.db
}
