package rssfeed

import (
	"os"
	"path/filepath"
	"time"

	"github.com/gorilla/feeds"
	"github.com/pocketbase/pocketbase"
)

func GenerateRssFeed(app *pocketbase.PocketBase) error {
	siteConfigRecord, _ := app.Dao().FindFirstRecordByFilter("site_config", "id!=''")
	feed := feeds.Feed{
		Title:       siteConfigRecord.GetString("site_name"),
		Link:        &feeds.Link{Href: siteConfigRecord.GetString("site_url")},
		Description: siteConfigRecord.GetString("site_description"),
		Created:     time.Now(),
	}

	episodes, err := app.Dao().FindRecordsByFilter("episodes", "status=\"published\"", "-updated", 0, 0)
	if err != nil {
		return err
	}
	for _, e := range episodes {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       e.GetString("title"),
			Link:        &feeds.Link{Href: siteConfigRecord.GetString("site_url") + "/episode/" + e.GetString("audio_file")},
			Description: e.GetString("description"),
			Created:     e.Created.Time(),
		})
	}

	rssFeed, _ := (&feeds.Rss{Feed: &feed}).ToRss()

	rssFolder := filepath.Join(".", "pb_public")
	err = os.MkdirAll(rssFolder, os.ModePerm)
	if err != nil {
		return err
	}

	os.WriteFile(filepath.Join("pb_public", "rss.xml"), []byte(rssFeed), os.ModePerm)

	return nil
}
