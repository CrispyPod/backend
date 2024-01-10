package dbModels

import (
	"database/sql"
	"time"

	"github.com/google/uuid"
)

type DeployLog struct {
	ID       uuid.UUID `gorm:"type:uuid;primary_key"`
	Log      sql.NullString
	Status   int
	BuildAt  time.Time
	Duration time.Duration
}
