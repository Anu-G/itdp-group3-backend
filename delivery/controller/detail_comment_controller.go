package controller

import (
	"errors"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
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
	routeDetailComment.GET("/", controller.readDetailComment)
	routeDetailComment.POST("/create", controller.createDetailComment)
	routeDetailComment.DELETE("/delete", controller.deleteComment)

	return controller
}

func (cm *DetailCommentController) readDetailComment(ctx *gin.Context) {
	var readComment entity.DetailComment
	err := cm.ParseBodyRequest(ctx, &readComment)
	if err != nil {
		cm.FailedResponse(ctx, err)
		return
	}
	if readComment.ID == 0 {
		cm.FailedResponse(ctx, errors.New("no comment found"))
		return
	}
	err = cm.cmUC.Read(&readComment)
	if err != nil {
		cm.FailedResponse(ctx, err)
		return
	}
	cm.SuccessResponse(ctx, readComment)
}

func (cm *DetailCommentController) createDetailComment(ctx *gin.Context) {
	var requestCreateComment dto.RequestCreateComment
	var createDetailComment entity.DetailComment
	err := cm.ParseBodyRequest(ctx, &requestCreateComment)
	if err != nil {
		cm.FailedResponse(ctx, err)
		return
	}
	createDetailComment.FeedID = requestCreateComment.FeedID
	createDetailComment.CommentFill = requestCreateComment.CommentFill
	if createDetailComment.FeedID == 0 {
		cm.FailedResponse(ctx, errors.New("no feed found"))
		return
	}
	err = cm.cmUC.Create(&createDetailComment)
	if err != nil {
		cm.FailedResponse(ctx, err)
		return
	}
	cm.SuccessResponse(ctx, createDetailComment)
}

func (cm *DetailCommentController) deleteComment(ctx *gin.Context) {
	var requestDeleteComment dto.RequestDeleteComment
	var deleteDetailComment entity.DetailComment
	err := cm.ParseBodyRequest(ctx, &deleteDetailComment)
	if err != nil {
		cm.FailedResponse(ctx, err)
		return
	}
	deleteDetailComment.ID = requestDeleteComment.CommentID
	if deleteDetailComment.ID == 0 {
		cm.FailedResponse(ctx, errors.New("no feed found"))
		return
	}
	err = cm.cmUC.Delete(&deleteDetailComment)
	if err != nil {
		cm.FailedResponse(ctx, err)
		return
	}
	cm.SuccessResponse(ctx, deleteDetailComment)
}
