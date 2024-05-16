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

		settings, err := dao.FindSettings()
		if err == nil {
			settings.Meta.AppName = "CrispyPod"
			settings.Meta.AppUrl = "https://crispypod.com"
			dao.SaveSettings(settings)
		}

		admin := &models.Admin{}
		admin.Email = "pb@crispy.pod"
		admin.SetPassword("crispy.pod")

		return dao.SaveAdmin(admin)
	}, func(db dbx.Builder) error { // optional revert operation

		dao := daos.New(db)

		admin, _ := dao.FindAdminByEmail("admin@crispypod.com")
		if admin != nil {
			return dao.DeleteAdmin(admin)
		}

		// already deleted
		return nil
	})
}
