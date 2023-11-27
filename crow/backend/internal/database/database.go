package database

import (
	"github.com/zoumas/lab/crow/backend/internal/user"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
)

func New(dsn string) (*gorm.DB, error) {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})
	if err != nil {
		return nil, err
	}

	if err := automigrate(db); err != nil {
		return nil, err
	}

	return db, err
}

func automigrate(db *gorm.DB) error {
	if err := db.AutoMigrate(&user.User{}); err != nil {
		return err
	}

	return nil
}
