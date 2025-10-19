package storage

import "gorm.io/gorm"

type Database interface {
	Connect() error
	MigrateUp() error
	DB() *gorm.DB
}
