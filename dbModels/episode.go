package dbModels

import (
	"database/sql"
	"regexp"
	"strings"
	"time"

	"crispypod.com/crispypod-backend/graph/model"
	"github.com/google/uuid"
)

type EpisodeStatusType int

const (
	EpisodeStatus_Draft     EpisodeStatusType = 0
	EpisodeStatus_Published EpisodeStatusType = 1
)

type Episode struct {
	ID            uuid.UUID `gorm:"type:uuid;primary_key"`
	Title         string
	CreateTime    time.Time
	PublishTime   sql.NullTime
	Description   string
	EpisodeStatus EpisodeStatusType
	NamedLink     string `gorm:"uniqueIndex"`

	ThumbnailFileName   sql.NullString
	ThumbnailUploadName sql.NullString

	AudioFileName       sql.NullString
	AudioFileUploadName sql.NullString
	AudioFileDuration   sql.NullInt64

	UserID uuid.UUID
	User   DbUser
}

func (e *Episode) ToGQLEpisode() *model.Episode {
	es := new(int)
	*es = int(e.EpisodeStatus)
	rtEpisode := model.Episode{
		ID:            e.ID.String(),
		Title:         e.Title,
		CreateTime:    int(e.CreateTime.Unix()),
		Description:   e.Description,
		EpisodeStatus: es,
		NamedLink:     e.NamedLink,
	}

	if e.ThumbnailFileName.Valid {
		strPtr := strings.Clone(e.ThumbnailFileName.String)
		rtEpisode.ThumbnailFileName = &strPtr
	}

	if e.ThumbnailUploadName.Valid {
		strPtr := strings.Clone(e.ThumbnailUploadName.String)
		rtEpisode.ThumbnailUploadName = &strPtr
	}

	if e.AudioFileName.Valid {
		strPtr := strings.Clone(e.AudioFileName.String)
		rtEpisode.AudioFileName = &strPtr
	}

	if e.AudioFileUploadName.Valid {
		strPtr := strings.Clone(e.AudioFileUploadName.String)
		rtEpisode.AudioFileUploadName = &strPtr
	}

	if e.PublishTime.Valid {
		pt := new(int)
		*pt = int(e.PublishTime.Time.Unix())
		rtEpisode.PublishTime = pt
	}

	if e.AudioFileDuration.Valid {
		ad := new(int)
		*ad = int(e.AudioFileDuration.Int64)
		rtEpisode.AudioFileDuration = ad
	}

	return &rtEpisode
}

func (e *Episode) GenerateNamedLink() {
	rexp := regexp.MustCompile(`[\\?\\!_.,: ]`)
	e.NamedLink = rexp.ReplaceAllString(e.Title, "-")
}
