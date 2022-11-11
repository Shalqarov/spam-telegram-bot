package service

import "github.com/Shalqarov/spam-telegram-bot/internal/models"

type userService struct {
	repo models.Repo
}

func NewUserService(userRepo models.Repo) models.Service {
	return &userService{repo: userRepo}
}

func (u *userService) CreateUser(user *models.User) (int64, error) {
	return u.repo.CreateUser(user)
}

func (u *userService) GetUserById(id int64) (*models.User, error) {
	return u.repo.GetUserById(id)
}
