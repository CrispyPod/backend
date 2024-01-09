package models

import (
	"time"

	"github.com/google/uuid"
)

type DeployLog struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Log      string
	Status   int
	BuildAt  time.Time
	Duration time.Duration
}
