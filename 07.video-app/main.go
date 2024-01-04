package main

import (
	"github.com/gin-gonic/gin"
	"video-app/config"
	"video-app/controller"
	"video-app/middleware"
	"video-app/repository"
	"video-app/service"
)

func init() {

	config.LoadEnvs()
	config.ConnectToDb()
}
func main() {

	videoRepository := repository.NewVideoRepository()
	videoService := service.NewVideoService(videoRepository)
	videoController := controller.NewVideoController(videoService)

	router := gin.Default()

	router.GET("/videos/:id", videoController.Find)
	router.GET("/videos", videoController.FindAll)
	router.POST("/videos", videoController.Save)

	// auth related routes
	authRepository := repository.NewAuthRepository()
	authService := service.NewAuthService(authRepository)
	authController := controller.NewAuthController(authService)

	router.POST("/signup", authController.SignUp)
	router.POST("/login", authController.Login)
	router.GET("/validate", middleware.RequireAuth, authController.Validate)

	//router.Run(":8080")
	router.Run() // It will pick the PORT env var defined in local.env

}
