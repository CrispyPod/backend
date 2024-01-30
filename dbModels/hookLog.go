package dbModels

import (
	"database/sql"
	"strings"
	"time"

	"crispypod.com/crispypod-backend/graph/model"
	"github.com/google/uuid"
)

type HookLogStatusType int

const (
	HookLogStatusType_Started  HookLogStatusType = 0
	HookLogStatusType_Finished HookLogStatusType = 1
)

type HookLog struct {
	ID             uuid.UUID `gorm:"type:uuid;primary_key"`
	HooksID        uuid.UUID
	Hooks          Hook
	Status         HookLogStatusType
	ResponseHeader sql.NullString
	ResponseBody   sql.NullString // this should include whole http response,including header, maybe we do parse in front end
	CreateTime     time.Time
	Duration       time.Duration
}

func (l *HookLog) ToGQLHookLog() *model.HookLog {
	headerStr := strings.Clone(l.ResponseHeader.String)
	bodyStr := strings.Clone(l.ResponseBody.String)
	rtHookLog := model.HookLog{
		ID:             l.ID.String(),
		HookID:         l.HooksID.String(),
		Status:         int(l.Status),
		ResponseHeader: headerStr,
		ResponseBody:   bodyStr,
		CreateTime:     int(l.CreateTime.Unix()),
		Duration:       int(l.Duration),
	}
	return &rtHookLog
}
