package gorm_db_env_connector

import (
	"gorm.io/gorm"
	"os"
	"time"
)

func resolveEnvOrDefault(envVar string, defaultValue string) string {

	res := os.Getenv(envVar)
	if res == "" {
		res = defaultValue
	}
	return res

}

type StringModel struct {
	ID        string         `gorm:"primarykey;size:36" json:"id,omitempty"`
	CreatedAt time.Time      `json:"created_at"`
	UpdatedAt time.Time      `json:"updated_at"`
	DeletedAt gorm.DeletedAt `gorm:"index" json:"deleted_at"`
}

func Paginate(page int, size int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		offset := page * size
		return db.Offset(offset).Limit(size)
	}
}

type Error struct {
	Code    int
	Message string
}

func (e Error) Error() string {
	return e.Message
}
