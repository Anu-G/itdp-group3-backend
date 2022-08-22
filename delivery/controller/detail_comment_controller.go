package controller

import (
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"

	"github.com/gin-gonic/gin"
)

type DetailCommentController struct {
	router     *gin.Engine
	cmUC       usecase.DetailCommentUsecase
	middleware middleware.AuthTokenMiddleware
	api.BaseApi
}

func NewCommentController(router *gin.Engine, cmUC usecase.DetailCommentUsecase, md middleware.AuthTokenMiddleware) *DetailCommentController {
	controller := &DetailCommentController{
		router:     router,
		cmUC:       cmUC,
		middleware: md,
	}

	routeDetailComment := controller.router.Group("/comment")
	routeDetailComment.Use(md.RequireToken())
	routeDetailComment.GET("/")
	routeDetailComment.POST("/create")

	return controller
}

func (cm *DetailCommentController) readDetailComment(ctx *gin.Context) {
	var readComment entity.DetailComment
	err := cm.ParseBodyRequest(ctx, &readComment)
	if err != nil {
		cm.FailedResponse(ctx, err)
	}
	err = cm.cmUC.Read(&readComment)
	if err != nil {
		cm.FailedResponse(ctx, err)
	}
	cm.SuccessResponse(ctx, readComment)
}

func (cm *DetailCommentController) createDetailComment(ctx *gin.Context) {
	var createDetailComment entity.DetailComment
	err := cm.ParseBodyRequest(ctx, &createDetailComment)
	if err != nil {
		cm.FailedResponse(ctx, err)
	}
	err = cm.cmUC.Create(&createDetailComment)
	if err != nil {
		cm.FailedResponse(ctx, err)
	}
	cm.SuccessResponse(ctx, createDetailComment)
}
