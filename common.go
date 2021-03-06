package gorm_db_env_connector

import (
	"gorm.io/gorm"
	"math"
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
	Code    int    `json:"code,omitempty"`
	Message string `json:"message,omitempty"`
}

func (e Error) Error() string {
	return e.Message
}

type Page struct {
	TotalElements int64         `json:"total_elements"`
	PageElements  int           `json:"page_elements"`
	Results       []interface{} `json:"results"`
	PageNumber    int           `json:"page_number"`
	PageSize      int           `json:"page_size"`
	TotalPages    int           `json:"total_pages"`
}

func (p *Page) New(pageNumber, pageSize int, totalElements int64, results []interface{}) {
	p.PageNumber = pageNumber
	p.PageSize = pageSize
	p.TotalElements = totalElements
	p.Results = results
	p.PageElements = len(results)
	p.TotalPages = int(math.Ceil(float64(totalElements) / float64(pageSize)))
}
