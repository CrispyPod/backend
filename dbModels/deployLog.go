package dbModels

import (
	"database/sql"
	"time"

	"crispypod.com/crispypod-backend/graph/model"
	"github.com/google/uuid"
)

type DeployLog struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	Log        sql.NullString
	Status     int
	CreateTime time.Time
	Duration   time.Duration
}

func (l *DeployLog) ToGQLDeployLog() *model.DeployLog {
	rtDeployLog := model.DeployLog{
		ID:         l.ID.String(),
		Log:        l.Log.String,
		Status:     l.Status,
		CreateTime: int(l.CreateTime.Unix()),
		Duration:   int(l.Duration),
	}
	return &rtDeployLog
}
