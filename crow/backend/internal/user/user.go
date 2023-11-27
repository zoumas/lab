package user

import "github.com/zoumas/lab/crow/backend/internal/model"

type User struct {
	model.Model
	Name     string `json:"name" gorm:"size:20;unique;not null"`
	Password string `json:"-"    gorm:"size:72;not null"`
}

type CreateUserParams struct {
	Name     string
	Password string
}

type UserRepo interface {
	Save(user User) error
}