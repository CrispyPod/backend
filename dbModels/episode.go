package dbModels

import (
	"database/sql"
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
		// PublishTime:         pt,
		ThumbnailFileName:   &e.ThumbnailFileName.String,
		ThumbnailUploadName: &e.ThumbnailUploadName.String,
		AudioFileName:       &e.AudioFileName.String,
		AudioFileUploadName: &e.AudioFileUploadName.String,
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
