package main

import (
	"context"
	"database/sql"
	"log"
	"strings"
	"sync"
	"time"

	"github.com/google/uuid"
	"github.com/zoumas/lab/boot.dev/rssaggr/internal/database"
)

func startScraping(db *database.Queries, workers int, interval time.Duration) {
	log.Printf("Scraping on %v goroutines with interval %s", workers, interval)

	ticker := time.NewTicker(interval)
	for ; ; <-ticker.C {
		feeds, err := db.GetNextFeedsToFetch(context.Background(), int32(workers))
		if err != nil {
			log.Println("error fetching feeds:", err)
			continue
		}

		wg := &sync.WaitGroup{}
		for _, feed := range feeds {
			wg.Add(1)
			go scrapeFeed(wg, db, feed)
		}
		wg.Wait()
	}
}

func scrapeFeed(wg *sync.WaitGroup, db *database.Queries, feed database.Feed) {
	defer wg.Done()

	_, err := db.MarkFeedAsFetched(context.Background(), feed.ID)
	if err != nil {
		log.Printf("error marking feed with id %s as fetched\n", feed.ID)
		return
	}

	rssFeed, err := ParseRssFeed(feed.Url)
	if err != nil {
		log.Printf("error fetching feed with id %s\n", feed.ID)
		return
	}

	for _, item := range rssFeed.Channel.Item {
		now := time.Now().UTC()
		description := sql.NullString{}
		if item.Description != "" {
			description.String = item.Description
			description.Valid = true
		}
		pubDate, err := time.Parse(time.RFC1123Z, item.PubDate)
		if err != nil {
			log.Printf("couldn't parse date: %v. Feed: %q. error: %q", item.PubDate, feed.Name, err)
			continue
		}

		_, err = db.CreatePost(context.Background(), database.CreatePostParams{
			ID:          uuid.New(),
			CreatedAt:   now,
			UpdatedAt:   now,
			Url:         item.Link,
			Title:       item.Title,
			Description: description,
			PublishedAt: pubDate,
			FeedID:      feed.ID,
		})
		if err != nil {
			if strings.Contains(err.Error(), "duplicate key") {
				continue
			}
			log.Printf("failed to create post with title: %s. error: %q", item.Title, err)
		}
	}
	log.Printf("Feed %s collected, %v posts found", feed.Name, len(rssFeed.Channel.Item))
}
