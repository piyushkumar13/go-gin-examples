package exception

import (
	"github.com/gin-gonic/gin"
	"net/http"
)

func HasHandleAppError(err error, ctx *gin.Context) bool {

	if err == nil {
		return false
	}

	videoError, isVideoErr := err.(*VideoAppError)

	if isVideoErr {
		ctx.AbortWithStatusJSON(videoError.GetHttpStatus(), gin.H{
			"errorDescription": videoError.Error(),
		})
		return true
	}

	ctx.AbortWithStatus(http.StatusInternalServerError)
	return true
}

func HasHandleErrorWithCode(err error, httpStatus int, ctx *gin.Context) bool {

	if err == nil {
		return false
	}

	ctx.AbortWithStatus(httpStatus)
	return true
}

func HasHandleErrorWithJson(err error, httpStatus int, ctx *gin.Context) bool {

	if err == nil {
		return false
	}

	ctx.AbortWithStatusJSON(httpStatus, gin.H{
		"errorDescription": err.Error(),
	})
	return true
}
