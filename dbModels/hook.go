package dbModels

import (
	"database/sql"
	"strings"
	"time"

	"crispypod.com/crispypod-backend/graph/model"
	"github.com/google/uuid"
)

type HookTriggerType int

const (
	HookTriggerType_SiteConfigChanged        HookTriggerType = 0
	HookTriggerType_EpisodeVisibilityChanged HookTriggerType = 1
)

type Hook struct {
	ID      uuid.UUID `gorm:"type:uuid;primary_key"`
	Name    string
	Trigger HookTriggerType
	Enabled bool

	WebURL     string
	Method     string
	Headers    sql.NullString // in json format
	AppendBody sql.NullString
	CreateTime time.Time
	ModifyTime time.Time
}

func (h *Hook) ToGQLHook() *model.Hook {
	headerStr := strings.Clone(h.Headers.String)
	bodyStr := strings.Clone(h.AppendBody.String)
	rtHook := model.Hook{
		ID:         h.ID.String(),
		Name:       h.Name,
		Trigger:    int(h.Trigger),
		WebURL:     h.WebURL,
		Method:     h.Method,
		Enabled:    h.Enabled,
		Headers:    headerStr,
		AppendBody: bodyStr,
	}

	return &rtHook
}
