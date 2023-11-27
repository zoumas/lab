package app

import (
	"github.com/zoumas/lab/crow/backend/internal/database"
	"github.com/zoumas/lab/crow/backend/internal/user"
)

type UserService struct {
	repo database.GormUserRepo
}

func NewUserService(repo database.GormUserRepo) *UserService {
	return &UserService{repo}
}

func (s *UserService) Save(user user.User) error {
	return s.repo.Save(user)
}
