package dbModels

import (
	"crispypod.com/crispypod-backend/graph/model"
	"github.com/google/uuid"
)

type SiteConfig struct {
	ID              uuid.UUID `gorm:"type:uuid;primary_key"`
	SiteName        string
	SiteDescription string
	SiteUrl         string
	SetupComplete   bool `gorm:"default:false"`
}

func (siteConfig *SiteConfig) ToGQLSiteConfig(includeSetupComplete bool) *model.SiteConfig {
	if includeSetupComplete {
		return &model.SiteConfig{
			ID:              siteConfig.ID.String(),
			SiteName:        siteConfig.SiteName,
			SiteDescription: siteConfig.SiteDescription,
			SiteURL:         siteConfig.SiteUrl,
			SetupComplete:   siteConfig.SetupComplete,
		}
	} else {
		return &model.SiteConfig{
			ID:              siteConfig.ID.String(),
			SiteName:        siteConfig.SiteName,
			SiteDescription: siteConfig.SiteDescription,
			SiteURL:         siteConfig.SiteUrl,
		}
	}

}
