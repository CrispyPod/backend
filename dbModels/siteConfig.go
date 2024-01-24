package dbModels

import (
	"database/sql"

	"crispypod.com/crispypod-backend/graph/model"
	"github.com/google/uuid"
)

type SiteConfig struct {
	ID               uuid.UUID `gorm:"type:uuid;primary_key"`
	SiteName         string
	SiteDescription  string
	SiteUrl          string
	SetupComplete    bool `gorm:"default:false"`
	SiteIconFile     sql.NullString
	DefaultThumbnail sql.NullString
}

func (siteConfig *SiteConfig) ToGQLSiteConfig(includeSetupComplete bool) *model.SiteConfig {
	if includeSetupComplete {
		return &model.SiteConfig{
			ID:               siteConfig.ID.String(),
			SiteName:         siteConfig.SiteName,
			SiteDescription:  siteConfig.SiteDescription,
			SiteURL:          siteConfig.SiteUrl,
			SetupComplete:    siteConfig.SetupComplete,
			SiteIconFile:     siteConfig.SiteIconFile.String,
			DefaultThumbnail: siteConfig.DefaultThumbnail.String,
		}
	} else {
		return &model.SiteConfig{
			ID:               siteConfig.ID.String(),
			SiteName:         siteConfig.SiteName,
			SiteDescription:  siteConfig.SiteDescription,
			SiteURL:          siteConfig.SiteUrl,
			SiteIconFile:     siteConfig.SiteIconFile.String,
			DefaultThumbnail: siteConfig.DefaultThumbnail.String,
		}
	}

}
