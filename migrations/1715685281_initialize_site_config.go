package migrations

import (
	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase/daos"
	m "github.com/pocketbase/pocketbase/migrations"
	"github.com/pocketbase/pocketbase/models"
)

func init() {
	m.Register(func(db dbx.Builder) error {
		dao := daos.New(db)
		siteConfigRecord, err := dao.FindFirstRecordByFilter("site_config", "")
		if err == nil && siteConfigRecord != nil {
			dao.DeleteRecord(siteConfigRecord)
		}

		siteConfigCollection, _ := dao.FindCollectionByNameOrId("site_config")

		siteConfigRecord = models.NewRecord(siteConfigCollection)
		siteConfigRecord.Set("site_name", "Crispy Pod")
		siteConfigRecord.Set("site_description", "Awesome podcast!")
		siteConfigRecord.Set("site_url", "https://crispypod.com")

		return dao.SaveRecord(siteConfigRecord)
	}, func(db dbx.Builder) error {
		// add down queries...

		return nil
	})
}
