package controller

import (
	"errors"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type DetailMediaFeedController struct {
	router     *gin.Engine
	fmUC       usecase.DetailMediaFeedUsecase
	middleware middleware.AuthTokenMiddleware
	api.BaseApi
}

func NewDetailMediaFeedController(router *gin.Engine, fmUC usecase.DetailMediaFeedUsecase, md middleware.AuthTokenMiddleware) *DetailMediaFeedController {
	controller := DetailMediaFeedController{
		router:     router,
		fmUC:       fmUC,
		middleware: md,
	}

	routeDetailMediaFeed := controller.router.Group("/mediafeed")
	routeDetailMediaFeed.Use(md.RequireToken())
	routeDetailMediaFeed.GET("/", controller.readDetailMediaFeed)
	routeDetailMediaFeed.POST("/create", controller.createDetailMediaFeed)

	return &controller
}

func (fm *DetailMediaFeedController) readDetailMediaFeed(ctx *gin.Context) {
	var readDetailMediaFeed entity.DetailMediaFeed
	err := fm.ParseBodyRequest(ctx, &readDetailMediaFeed)
	if err != nil {
		fm.FailedResponse(ctx, err)
	}
	err = fm.fmUC.Read(&readDetailMediaFeed)
	if err != nil {
		fm.FailedResponse(ctx, err)
	}
	fm.SuccessResponse(ctx, readDetailMediaFeed)
}

func (fm *DetailMediaFeedController) createDetailMediaFeed(ctx *gin.Context) {
	feedId := ctx.PostForm("feed_id")

	file, fileHeader, err := ctx.Request.FormFile("feed_image")
	if err != nil {
		fm.FailedResponse(ctx, errors.New("failed get file"))
		return
	}

	fileName := strings.Split(fileHeader.Filename, ".")
	if len(fileName) != 2 {
		fm.FailedResponse(ctx, errors.New("Unrecognized file extension"))
	}

	fileLocation, err := fm.fmUC.Create(feedId, file, fileName[1])
	if err != nil {
		fm.FailedResponse(ctx, errors.New("failed while saving file"))
		return
	}

	fm.SuccessResponse(ctx, fileLocation)
}
