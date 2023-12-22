package models

import (
	"crispypod.com/crispypod-backend/graph/model"
	"github.com/google/uuid"
)

type SiteConfig struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key"`
	SiteName        string
	SiteDescription string
	SiteUrl         string
}

func (siteConfig *SiteConfig) ToGQLSiteConfig() *model.SiteConfig {
	return &model.SiteConfig{
		ID:              siteConfig.ID.String(),
		SiteName:        siteConfig.SiteName,
		SiteDescription: siteConfig.SiteDescription,
		SiteURL:         siteConfig.SiteUrl,
	}
}
