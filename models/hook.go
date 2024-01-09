package models

import "github.com/google/uuid"

type Hook struct {
	ID        uuid.UUID `gorm:"type:uuid;primary_key"`
	Name      string
	TableName string
	Trigger   string

	WebURL     string
	Method     string
	Headers    string // in json format
	AppendBody string
}
