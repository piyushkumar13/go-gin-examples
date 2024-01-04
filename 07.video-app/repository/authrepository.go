package repository

import (
	"video-app/config"
	"video-app/domain/entity"
)

type AuthRepository interface {
	Save(user *entity.User)
	Find(email string) *entity.User
}

type authRepository struct {
}

func NewAuthRepository() AuthRepository {

	return &authRepository{}
}

func (authRepo *authRepository) Save(user *entity.User) {

	config.DB.Save(user)
}

func (authRepo *authRepository) Find(email string) *entity.User {

	var user entity.User
	config.DB.First(&user, "email = ?", email)

	return &user
}
