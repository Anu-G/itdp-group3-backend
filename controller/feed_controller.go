package controller

import (
	"fmt"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"

	"github.com/gin-gonic/gin"
)

type FeedController struct {
	router     *gin.Engine
	fUC        usecase.FeedUsecase
	middleware middleware.AuthTokenMiddleware
	api.BaseApi
}

func NewFeedController(router *gin.Engine, fUC usecase.FeedUsecase, md middleware.AuthTokenMiddleware) *FeedController {
	controller := FeedController{
		router:     router,
		fUC:        fUC,
		middleware: md,
	}

	routeFeed := controller.router.Group("/feed")
	routeFeed.Use(md.RequireToken())
	routeFeed.GET("/", controller.readFeed)
	routeFeed.GET("/account", controller.readAccountFeed)
	routeFeed.GET("/category", controller.readCategoryFeed)
	routeFeed.GET("/paged", controller.readByPageFeed)
	routeFeed.POST("/create", controller.createFeed)
	routeFeed.DELETE("/", controller.deleteFeed)

	return &controller
}

func (f *FeedController) readFeed(ctx *gin.Context) {
	var readFeed entity.Feed
	err := f.fUC.Read(&readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	f.SuccessResponse(ctx, readFeed)
}

func (f *FeedController) readAccountFeed(ctx *gin.Context) {
	var readFeed dto.ReadPage
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	resFeed, err := f.fUC.ReadByAccountID(readFeed.ID, readFeed.Page, readFeed.PageLim)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	f.SuccessResponse(ctx, resFeed)
}

func (f *FeedController) readCategoryFeed(ctx *gin.Context) {
	var readFeed dto.ReadPage
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	resFeed, err := f.fUC.ReadByProfileCategory(readFeed.Cat, readFeed.Page, readFeed.PageLim)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	f.SuccessResponse(ctx, resFeed)
}

func (f *FeedController) readByPageFeed(ctx *gin.Context) {
	var readFeed dto.ReadPage
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	resFeed, err := f.fUC.ReadByPage(readFeed.Page, readFeed.PageLim)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	f.SuccessResponse(ctx, resFeed)
}

func (f *FeedController) deleteFeed(ctx *gin.Context) {
	var readFeed dto.DeleteFeed
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	err = f.fUC.Delete(readFeed.ID)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	delMessage := fmt.Sprintf("Delete feed id %d success", readFeed.ID)
	f.SuccessResponse(ctx, delMessage)
}

func (f *FeedController) createFeed(ctx *gin.Context) {
	var createFeed entity.Feed
	err := f.ParseBodyRequest(ctx, &createFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	err = f.fUC.Create(&createFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
	}
	f.SuccessResponse(ctx, createFeed)
}
