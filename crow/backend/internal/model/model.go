package model

import (
	"time"

	"github.com/google/uuid"
)

// Model contains common columns shared by all database tables.
type Model struct {
	// ID is a global version 4 unique identifier serving as the primary key.
	ID        uuid.UUID `json:"id"         gorm:"type:uuid; primaryKey; unique; default: gen_random_uuid()"`
	CreatedAt time.Time `json:"created_at"`
	UpdatedAt time.Time `json:"updated_at"`
}
