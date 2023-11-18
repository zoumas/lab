package main

import (
	"encoding/xml"
	"net/http"
	"time"
)

type RssFeed struct {
	Channel struct {
		Title       string    `xml:"title"`
		Link        string    `xml:"link"`
		Description *string   `xml:"description"`
		Items       []RssItem `xml:"item"`
	} `xml:"channel"`
}

type RssItem struct {
	Title       string  `xml:"title"`
	Link        string  `xml:"link"`
	Description *string `xml:"description"`
	PubDate     string  `xml:"pubDate"`
}

func ParseRssFeed(url string) (RssFeed, error) {
	client := http.Client{
		Timeout: 10 * time.Second,
	}
	resp, err := client.Get(url)
	if err != nil {
		return RssFeed{}, err
	}

	rssFeed := RssFeed{}

	err = xml.NewDecoder(resp.Body).Decode(&rssFeed)
	if err != nil {
		return RssFeed{}, err
	}
	defer resp.Body.Close()

	return rssFeed, nil
}
