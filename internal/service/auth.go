package service

import (
	"github.com/ZhanserikKalmukhambet/Trello/internal/entity"
	"github.com/ZhanserikKalmukhambet/Trello/pkg/util"
)

func (a *AuthService) Register(user entity.User) (int, error) {
	hashedPassword, err := util.HashPassword(user.Password)
	if err != nil {
		return 0, err
	}

	user.Password = hashedPassword
	return a.repos.CreateUser(user)
}

func (a *AuthService) Login(username, password string) (entity.User, error) {
	return a.repos.GetUser(username, password)
}
