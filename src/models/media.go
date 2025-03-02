package models

import (
	"time"

	"github.com/gocql/gocql"
)

type Media struct {
	ID         gocql.UUID `json:"id"`
	UserID     gocql.UUID `json:"user_id"`
	FileName   string     `json:"file_name"`
	FileType   string     `json:"file_type"`
	FileURL    string     `json:"file_url"`
	UploadedAt time.Time  `json:"uploaded_at"`
}
