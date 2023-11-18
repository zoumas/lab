package main

import (
	"time"

	"github.com/google/uuid"
)

type Base struct {
	ID        *uuid.UUID `json:"id"         gorm:"type:uuid;primaryKey;default: gen_random_uuid()"`
	CreatedAt time.Time  `json:"created_at" gorm:"autoCreateTime;"`
	UpdatedAt time.Time  `json:"updated_at" gorm:"autoUpdateTime;"`
}

type User struct {
	Base
	Name   string `json:"name"    gorm:"not null;check: name <> ''"`
	ApiKey string `json:"api_key" gorm:"type:VARCHAR(64);unique;not null;default: encode(sha256(random()::text::bytea), 'hex')"`

	FeedFollows []Feed `json:"-" gorm:"many2many:feed_follows;"`
}

type Feed struct {
	Base
	Name string `json:"name" gorm:"not null;check: name <> ''"`
	Url  string `json:"url"  gorm:"unique;not null;check: url <> ''"`

	UserID *uuid.UUID `json:"user_id"`
	User   User       `json:"-"       gorm:"foreignKey:UserID;not null;constraint:OnDelete:CASCADE"`

	LastFetchedAt *time.Time `json:"last_fetched_at"`
}

type FeedFollow struct {
	UserID *uuid.UUID `json:"user_id" gorm:"primaryKey"`
	FeedID *uuid.UUID `json:"feed_id" gorm:"primaryKey"`
}

type Post struct {
	Base
	Title       string    `json:"title"        gorm:"not null"`
	Url         string    `json:"url"          gorm:"unique;not null"`
	Description *string   `json:"description"`
	PublishedAt time.Time `json:"published_at"`

	FeedID *uuid.UUID `json:"feed_id"`
	Feed   Feed       `json:"-"       gorm:"foreignKey:FeedID;not null;constraint:OnDelete:CASCADE"`
}
