package controller

import (
	"github.com/gin-gonic/gin"
	"golang.org/x/crypto/bcrypt"
	"net/http"
	"video-app/domain/entity"
	"video-app/exception"
	"video-app/service"
)

type AuthController interface {
	SignUp(ctx *gin.Context)
	Login(ctx *gin.Context)

	Validate(ctx *gin.Context)
}

type authController struct {
	as service.AuthService
}

func NewAuthController(authService service.AuthService) AuthController {
	return &authController{
		as: authService,
	}
}

func (authCtrl *authController) SignUp(ctx *gin.Context) {

	var user entity.User

	err := ctx.ShouldBind(&user)
	if exception.HasHandleErrorWithJson(err, http.StatusBadRequest, ctx) {
		return
	}

	hashPassword, err := bcrypt.GenerateFromPassword([]byte(user.Password), 10)
	if exception.HasHandleErrorWithJson(err, http.StatusBadRequest, ctx) {
		return
	}

	user.Password = string(hashPassword)

	authCtrl.as.Save(&user)

	ctx.JSON(http.StatusNoContent, gin.H{})
}

func (authCtrl *authController) Login(ctx *gin.Context) {

	var user entity.User

	err := ctx.ShouldBind(&user)

	if exception.HasHandleErrorWithJson(err, http.StatusBadRequest, ctx) {
		return
	}

	tokenString, err := authCtrl.as.ValidateAndReturnToken(&user)

	if exception.HasHandleAppError(err, ctx) {
		return
	}

	ctx.SetCookie("Authorization", tokenString, 3600, "", "", false, true)
}

func (authCtrl *authController) Validate(ctx *gin.Context) {

	ctx.JSON(http.StatusOK, gin.H{
		"msg": "I am logged in",
	})
}
