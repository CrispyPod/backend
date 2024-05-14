package helpers

import (
	"regexp"
	"strings"

	"github.com/pocketbase/dbx"
	"github.com/pocketbase/pocketbase"
	"github.com/pocketbase/pocketbase/models"
)

func appendDashTilNoRepeat(slug string, table string, app *pocketbase.PocketBase, id string) string {
	l1Reg := regexp.MustCompile(`[\\?\\!_.,: ]`)
	l2Reg := regexp.MustCompile("-+")

	l1Result := l1Reg.ReplaceAllString(slug, "-")
	slug = l2Reg.ReplaceAllString(l1Result, "-")
	slug = strings.ToLower(slug)

	for _, err := app.Dao().FindFirstRecordByFilter(table, "link='{:link}' and id!={:id}", dbx.Params{"link": slug, "id": id}); err == nil; {
		slug += "-"
	}

	return slug
}

func EpisodeGenerateSlug(record *models.Record, app *pocketbase.PocketBase) {
	slug := appendDashTilNoRepeat(record.GetString("title"), "episodes", app, record.Id)
	record.Set("slug", slug)

	app.Dao().SaveRecord(record)
}
