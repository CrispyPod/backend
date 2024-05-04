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
		usersCollection, _ := dao.FindCollectionByNameOrId("users")

		userRecord := models.NewRecord(usersCollection)
		userRecord.SetEmail("user@crispy.pod")
		userRecord.SetUsername("defaultuser")
		userRecord.SetVerified(true)
		userRecord.SetPassword("password")

		return dao.SaveRecord(userRecord)
	}, func(db dbx.Builder) error {
		dao := daos.New(db)

		user, _ := dao.FindAuthRecordByEmail("users", "user@crispy.pod")
		if user != nil {
			return dao.DeleteRecord(user)
		}

		return nil
	})
}
