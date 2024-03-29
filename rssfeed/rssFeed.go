package rssfeed

import (
	"fmt"
	"os"
	"path/filepath"
	"time"

	"crispypod.com/crispypod-backend/db"
	"crispypod.com/crispypod-backend/dbModels"
	"github.com/gorilla/feeds"
)

func GenerateRSSFeed() {
	var dbSiteConfig dbModels.SiteConfig
	db.DB.First(&dbSiteConfig)
	feed := feeds.Feed{
		Title:       dbSiteConfig.SiteName,
		Link:        &feeds.Link{Href: dbSiteConfig.SiteUrl},
		Description: dbSiteConfig.SiteDescription,
		Created:     time.Now(),
	}

	var episodes []dbModels.Episode
	if err := db.DB.Find(&episodes, dbModels.Episode{EpisodeStatus: dbModels.EpisodeStatus_Published}).Error; err != nil {
		fmt.Println("Failed to get episodes.")
	}

	for _, e := range episodes {
		feed.Items = append(feed.Items, &feeds.Item{
			Title:       e.Title,
			Link:        &feeds.Link{Href: dbSiteConfig.SiteUrl + "/api/audioFile/" + e.AudioFileName.String},
			Description: e.Description,
			Created:     e.CreateTime,
		})
	}

	rssFeed, _ := (&feeds.Rss{Feed: &feed}).ToRss()

	rssFolder := "Feed"
	if _, err := os.Stat(rssFolder); os.IsNotExist(err) {
		if err := os.Mkdir(rssFolder, os.ModePerm); err != nil {
			fmt.Println("Failed to create Rss feed folder")
		}
	}

	os.WriteFile(filepath.Join(rssFolder, "rss.xml"), []byte(rssFeed), os.ModePerm)

}
