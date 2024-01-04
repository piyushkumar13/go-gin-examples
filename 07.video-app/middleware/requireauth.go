package middleware

import (
	"fmt"
	"github.com/gin-gonic/gin"
	"github.com/golang-jwt/jwt/v5"
	"log"
	"net/http"
	"os"
	"time"
	"video-app/config"
	"video-app/domain/entity"
)

func RequireAuth(ctx *gin.Context) {

	tokenString, err := ctx.Cookie("Authorization")

	if err != nil {
		log.Println("Authorization header is not present in cookie")
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	token, err := jwt.Parse(tokenString, func(token *jwt.Token) (interface{}, error) {

		if _, ok := token.Method.(*jwt.SigningMethodHMAC); !ok {
			return nil, fmt.Errorf("Unexpected signing method: %v", token.Header["alg"])
		}

		return []byte(os.Getenv("JWT_SECRET")), nil
	})

	if err != nil {
		log.Println("Token is invalid")
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}

	if claims, ok := token.Claims.(jwt.MapClaims); ok {

		subject := claims["sub"].(string)
		exp := claims["exp"].(float64)

		if float64(time.Now().Unix()) > exp {
			log.Println("Token is expired")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		var user entity.User
		config.DB.First(&user, "email = ?", subject)

		if user.Email == "" {
			log.Println("User with subject not found")
			ctx.AbortWithStatus(http.StatusUnauthorized)
		}

		ctx.Next()

	} else {
		log.Println("Token claims are invalid")
		ctx.AbortWithStatus(http.StatusUnauthorized)
	}
}
