package controller

import (
	"github.com/gin-gonic/gin"
	"net/http"
	"strconv"
	"video-app/domain/entity"
	"video-app/service"
)

type VideoController interface {
	Save(ctx *gin.Context)
	Find(ctx *gin.Context)
	FindAll(ctx *gin.Context)
}

type videoController struct {
	vs service.VideoService
}

func NewVideoController(service service.VideoService) VideoController {

	return &videoController{
		vs: service,
	}
}

func (v *videoController) Save(ctx *gin.Context) {

	var video entity.Video

	err := ctx.BindJSON(&video)
	if err != nil {
		ctx.JSON(http.StatusBadRequest, gin.H{
			"errorDescription": err.Error(),
		})

		return
	}

	addedVideo := v.vs.Save(video)

	ctx.JSON(201, addedVideo)
}

func (v *videoController) Find(ctx *gin.Context) {
	strid := ctx.Param("id")

	id, _ := strconv.ParseUint(strid, 10, 32)
	video := v.vs.Find(uint(id))

	ctx.JSON(200, video)

	// we can also do like this
	//id, _ := strconv.Atoi(strid)
	//uintId := uint(id)
	//v.vs.Find(uintId)
}

func (v *videoController) FindAll(ctx *gin.Context) {

	videos := v.vs.FindAll()

	ctx.JSON(200, videos)
}
