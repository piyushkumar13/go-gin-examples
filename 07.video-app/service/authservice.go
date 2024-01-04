package service

import (
	"github.com/golang-jwt/jwt/v5"
	"github.com/google/uuid"
	"time"
	"video-app/domain/entity"
	"video-app/repository"
	"video-app/util"
)

type AuthService interface {
	Save(user *entity.User)
	Find(email string) *entity.User
	ValidateAndReturnToken(user *entity.User) (string, error)
}

type authService struct {
	ar repository.AuthRepository
}

func NewAuthService(authRepository repository.AuthRepository) AuthService {
	return &authService{
		ar: authRepository,
	}
}

func (as *authService) Save(user *entity.User) {

	as.ar.Save(user)
}

func (as *authService) Find(email string) *entity.User {

	return as.ar.Find(email)
}

func (as *authService) ValidateAndReturnToken(user *entity.User) (string, error) {

	retrievedUser, err := util.ValidateEmailAndReturnRetrievedUser(user)
	if err != nil {
		return "", err
	}

	if err = util.ValidatePassword(user, retrievedUser); err != nil {
		return "", err
	}

	tokenString, err := util.GenerateToken(generateClaims(retrievedUser))
	if err != nil {
		return "", err
	}

	return tokenString, nil
}

func generateClaims(user *entity.User) *entity.CustomJwtClaims {

	claims := &entity.CustomJwtClaims{}

	claims.Subject = user.Email
	claims.ExpiresAt = jwt.NewNumericDate(time.Now().Add(time.Minute * 5))
	claims.ID = uuid.New().String()
	claims.IssuedAt = jwt.NewNumericDate(time.Now())
	claims.Issuer = "self"
	claims.Authorities = "ROLE_READ,ROLE_WRITE,SOME_AUTHORITY"
	claims.Scope = "READ,WRITE"

	return claims
}
