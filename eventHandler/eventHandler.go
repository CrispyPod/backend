package eventhandler

import (
	"fmt"
	"log"

	"crispypod.com/crispypod-backend/db"
	"crispypod.com/crispypod-backend/dbModels"
	"crispypod.com/crispypod-backend/rssfeed"
	"github.com/gookit/event"
)

type EventType string

const (
	EventType_EpisodeVisibilityChanged EventType = "epVisibilityChanged"
	EventType_SiteConfigChanged        EventType = "epSiteConfigChanged"
	EventType_PublishedEpisodeChanged  EventType = "epPublishedEPChanged"
)

func RegisterEvent() {
	event.On(string(EventType_EpisodeVisibilityChanged), event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("Generating RSS Feed...\n")
		rssfeed.GenerateRSSFeed()
		return nil
	}), event.High)

	event.On(string(EventType_SiteConfigChanged), event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("Generating RSS Feed...\n")
		rssfeed.GenerateRSSFeed()
		return nil
	}), event.High)

	event.On(string(EventType_PublishedEpisodeChanged), event.ListenerFunc(func(e event.Event) error {
		fmt.Printf("Generating RSS Feed...\n")
		rssfeed.GenerateRSSFeed()
		return nil
	}), event.High)

	event.On(string(EventType_EpisodeVisibilityChanged), event.ListenerFunc(func(e event.Event) error {
		episode := e.Get("episode").(dbModels.Episode)

		var hooks []dbModels.Hook
		if err := db.DB.Where(dbModels.Hook{Trigger: dbModels.HookTriggerType_EpisodeVisibilityChanged, Enabled: true}).Find(&hooks).Error; err != nil {
			log.Fatal(err)
		}
		for _, h := range hooks {
			go TriggerHook(h, episode)
		}

		return nil
	}), event.Normal)

	event.On(string(EventType_SiteConfigChanged), event.ListenerFunc(func(e event.Event) error {
		siteConfig := e.Get("siteConfig").(dbModels.SiteConfig)
		var hooks []dbModels.Hook
		if err := db.DB.Where(dbModels.Hook{Trigger: dbModels.HookTriggerType_SiteConfigChanged, Enabled: true}).Find(&hooks).Error; err != nil {
			log.Fatal(err)
		}
		for _, h := range hooks {
			go TriggerHook(h, siteConfig)
		}
		return nil
	}), event.Normal)

	event.On(string(EventType_PublishedEpisodeChanged), event.ListenerFunc(func(e event.Event) error {
		episode := e.Get("episode").(dbModels.Episode)

		var hooks []dbModels.Hook
		if err := db.DB.Where(dbModels.Hook{Trigger: dbModels.HookTriggerType_PublishedEpisodeChanged, Enabled: true}).Find(&hooks).Error; err != nil {
			log.Fatal(err)
		}
		for _, h := range hooks {
			go TriggerHook(h, episode)
		}

		return nil
	}), event.Normal)

}
