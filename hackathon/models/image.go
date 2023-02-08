package models

import (
	"encoding/json"
	"time"

	"gorm.io/gorm"
)

type Image struct {
	gorm.Model
	UserID    uint64    `json:"user_id"`
	Name      string    `json:"name" gorm:"unique"`
	Size      uint64    `json:"size"`
	Metadata  string    `json:"metadata"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}

// HttpRequestInfo is metadata of http request
type HttpRequestInfo struct {
	Method      string `json:"method"`
	URL         string `json:"url"`
	ContentType string `json:"content_type"`
	Proto       string `json:"proto"`
	Header      string `json:"header"`
}

func (h HttpRequestInfo) ToString() string {
	bytes, _ := json.Marshal(h)

	return string(bytes)
}
