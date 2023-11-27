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

func (s *UserService) GetAll() ([]user.User, error) {
	return s.repo.GetAll()
}

func (s *UserService) GetByName(name string) (user.User, error) {
	return s.repo.GetByName(name)
}

func (s *UserService) Delete(u user.User) error {
	return s.repo.Delete(u)
}
