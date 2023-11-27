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

func (r *GormUserRepo) GetAll() ([]user.User, error) {
	var users []user.User
	if err := r.db.Find(&users).Error; err != nil {
		return nil, err
	}
	return users, nil
}

func (r *GormUserRepo) GetByName(name string) (user.User, error) {
	var u user.User
	if err := r.db.Where("name = ?", name).First(&u).Error; err != nil {
		return user.User{}, nil
	}
	return u, nil
}

func (r *GormUserRepo) Delete(u user.User) error {
	return r.db.Delete(&u).Error
}
