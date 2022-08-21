package controller

import (
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"

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
	var createDetailMediaFeed entity.DetailMediaFeed
	err := fm.ParseBodyRequest(ctx, &createDetailMediaFeed)
	if err != nil {
		fm.FailedResponse(ctx, err)
	}
	err = fm.fmUC.Create(&createDetailMediaFeed)
	if err != nil {
		fm.FailedResponse(ctx, err)
	}
	fm.SuccessResponse(ctx, createDetailMediaFeed)
}
