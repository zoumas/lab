package main

import (
	"context"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/aggrss/internal/database"
)

func startScraping(db *database.Queries, workers int, interval time.Duration) {
	t := time.NewTicker(interval)
	for ; ; <-t.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(workers))
		if err != nil {
			log.Println(err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(wg, feed, db)
		}
		wg.Wait()
	}
}

func scrapeFeed(wg *sync.WaitGroup, feed database.Feed, db *database.Queries) {
	defer wg.Done()

	db.MarkFeedAsFetched(context.Background(), feed.ID)

	rssFeed, err := ParseRss(feed.Url)
	if err != nil {
		log.Printf("On feed: %q. Error: %q\n", feed.Name, err)
		return
	}

	for _, item := range rssFeed.Channel.Items {
		now := time.Now().UTC()
		var pubDate *time.Time
		if item.PubDate != nil {
			t, err := time.Parse(time.RFC1123Z, *item.PubDate)
			if err != nil {
				log.Println(err)
				continue
			}
			pubDate = &t
		}

		_, err := db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Title:       item.Title,
			Url:         item.Link,
			Description: item.Description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "unique") {
				continue
			}
			log.Println(err)
		}
	}
}
