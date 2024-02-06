package dbModels

import (
	"database/sql"
	"strings"

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
	HeadAnalytics    sql.NullString
	FooterAnalytics  sql.NullString
}

func (siteConfig *SiteConfig) ToGQLSiteConfig(includeSetupComplete bool) *model.SiteConfig {
	iconStr := strings.Clone(siteConfig.SiteIconFile.String)
	thumbnailStr := strings.Clone(siteConfig.DefaultThumbnail.String)

	rtSiteConfig := model.SiteConfig{
		ID:               siteConfig.ID.String(),
		SiteName:         siteConfig.SiteName,
		SiteDescription:  siteConfig.SiteDescription,
		SiteURL:          siteConfig.SiteUrl,
		SiteIconFile:     iconStr,
		DefaultThumbnail: thumbnailStr,
	}

	if includeSetupComplete {
		rtSiteConfig.SetupComplete = siteConfig.SetupComplete
	} else {
		rtSiteConfig.SetupComplete = true
	}

	if siteConfig.HeadAnalytics.Valid {
		strVal := strings.Clone(siteConfig.HeadAnalytics.String)
		rtSiteConfig.HeadAnalytics = &strVal
	}

	if siteConfig.FooterAnalytics.Valid {
		strVal := strings.Clone(siteConfig.FooterAnalytics.String)
		rtSiteConfig.FooterAnalytics = &strVal
	}

	return &rtSiteConfig

}
