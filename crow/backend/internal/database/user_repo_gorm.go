package database

import (
	"github.com/zoumas/lab/crow/backend/internal/user"
	"gorm.io/gorm"
)

type GormUserRepo struct {
	db *gorm.DB
}

func NewGormUserRepo(db *gorm.DB) *GormUserRepo {
	return &GormUserRepo{db}
}

func (r *GormUserRepo) Save(user user.User) error {
	return r.db.Create(&user).Error
}
