package dbModels

import (
	"database/sql"
	"time"

	"crispypod.com/crispypod-backend/graph/model"
	"github.com/google/uuid"
)

type HookLog struct {
	ID         uuid.UUID `gorm:"type:uuid;primary_key"`
	HooksID    uuid.UUID
	Hooks      Hook
	Status     int
	Response   sql.NullString // this should include whole http response,including header, maybe we do parse in front end
	CreateTime time.Time
	Duration   time.Duration
}

func (l *HookLog) ToGQLHookLog() *model.HookLog {
	rtHookLog := model.HookLog{
		ID:         l.ID.String(),
		HookID:     l.HooksID.String(),
		Status:     l.Status,
		Response:   l.Response.String,
		CreateTime: int(l.CreateTime.Unix()),
		Duration:   int(l.Duration),
	}
	return &rtHookLog
}
