package database

import (
	"time"
	"urlShortner/pkg/snowflake"

	"gorm.io/gorm"
)

type URL struct {
	ID        int64
	Long      string
	Short     string
	CreatedAt time.Time
}

func (u *URL) BeforeCreate(tx *gorm.DB) (err error) {
	u.ID = snowflake.GenerateID()
	u.CreatedAt = time.Time{}
	return nil
}
