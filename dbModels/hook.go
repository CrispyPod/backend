package dbModels

import (
	"database/sql"
	"time"

	"crispypod.com/crispypod-backend/graph/model"
	"github.com/google/uuid"
)

type HookTriggerType int

const (
	HookTriggerType_SiteConfig_Updated       HookTriggerType = 0
	HookTriggerType_EpisodeVisibilityChanged HookTriggerType = 1
)

type Hook struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Name    string
	Trigger HookTriggerType

	WebURL     string
	Method     string
	Headers    sql.NullString // in json format
	AppendBody sql.NullString
	CreateTime time.Time
	ModifyTime time.Time
}

func (h *Hook) ToGQLHook() *model.Hook {
	rtHook := model.Hook{
		ID:         h.ID.String(),
		Name:       h.Name,
		Trigger:    int(h.Trigger),
		WebURL:     h.WebURL,
		Method:     h.Method,
		Headers:    h.Headers.String,
		AppendBody: h.AppendBody.String,
	}

	return &rtHook
}
