package util

import (
	"github.com/golang-jwt/jwt/v5"
	"golang.org/x/crypto/bcrypt"
	"log"
	"net/http"
	"os"
	"video-app/domain/entity"
	exception "video-app/exception"
	"video-app/repository"
)

var authRepository = repository.NewAuthRepository()

func GenerateToken(claims jwt.Claims) (string, error) {
	token := jwt.NewWithClaims(jwt.SigningMethodHS256, claims)

	tokenString, err := token.SignedString([]byte(os.Getenv("JWT_SECRET")))

	if err != nil {
		log.Println("Failed to generate token")
		return "", exception.NewVideoAppError(http.StatusUnauthorized, "Failed to sign token")

	}
	return tokenString, nil
}

func ValidateEmailAndReturnRetrievedUser(user *entity.User) (*entity.User, error) {

	retrievedUser := authRepository.Find(user.Email)

	if retrievedUser == nil && retrievedUser.Email != user.Email {
		log.Println("incorrect credentials or user does not exists, please sign up")

		return nil, exception.NewVideoAppError(http.StatusUnauthorized, "incorrect credentials or user does not exists, please sign up")
	}

	return retrievedUser, nil
}

func ValidatePassword(user, retrievedUser *entity.User) error {

	err := bcrypt.CompareHashAndPassword([]byte(retrievedUser.Password), []byte(user.Password))

	if err != nil {
		log.Println("incorrect password")
		return exception.NewVideoAppError(http.StatusUnauthorized, "incorrect password")
	}

	return nil
}
