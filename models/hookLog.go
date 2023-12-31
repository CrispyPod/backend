package models

import (
	"time"

	"github.com/google/uuid"
)

type HookLog struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	HooksID    uuid.UUID
	Hooks      Hook
	Status     int
	Response   string // this should include whole http response,including header, maybe we do parse in front end
	CreateTime time.Time
	Duration   time.Duration
}
