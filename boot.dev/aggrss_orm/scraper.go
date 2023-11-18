package main

import (
	"log"
	"strings"
	"sync"
	"time"
)

func (cfg *ApiConfig) StartScraping(workers int, interval time.Duration) {
	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		feeds, err := cfg.GetNextFeedsToFetch(workers)
		if err != nil {
			log.Println(err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go cfg.ScrapeFeed(wg, feed)
		}
		wg.Wait()
	}
}

func (cfg *ApiConfig) ScrapeFeed(wg *sync.WaitGroup, feed Feed) {
	defer wg.Done()

	err := cfg.MarkFeedAsFetched(feed)
	if err != nil {
		log.Println(err)
		return
	}

	rssFeed, err := ParseRssFeed(feed.Url)
	if err != nil {
		log.Println(err)
		return
	}

	log.Printf("Feed %q contains %d posts\n", rssFeed.Channel.Title, len(rssFeed.Channel.Items))
	for _, item := range rssFeed.Channel.Items {
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Println("Failed to parse time for:", item.Title)
			continue
		}

		post := Post{
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		}

		if err := cfg.DB.Create(&post).Error; err != nil {
			if !strings.Contains(err.Error(), "unique") {
				log.Println(err)
			}
			continue
		}
	}
}
