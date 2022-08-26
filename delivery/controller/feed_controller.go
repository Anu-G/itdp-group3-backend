package controller

import (
	"errors"
	"fmt"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"
	"itdp-group3-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type FeedController struct {
	router     *gin.Engine
	fUC        usecase.FeedUsecase
	fmUC       usecase.DetailMediaFeedUsecase
	middleware middleware.AuthTokenMiddleware
	api.BaseApi
}

func NewFeedController(router *gin.Engine, fUC usecase.FeedUsecase, fmUC usecase.DetailMediaFeedUsecase, md middleware.AuthTokenMiddleware) *FeedController {
	controller := FeedController{
		router:     router,
		fUC:        fUC,
		fmUC:       fmUC,
		middleware: md,
	}

	routeFeed := controller.router.Group("/feed")
	routeFeed.Use(md.RequireToken())
	routeFeed.GET("/", controller.readFeed)
	routeFeed.GET("/account", controller.readAccountFeed)
	routeFeed.GET("/category", controller.readCategoryFeed)
	routeFeed.GET("/paged", controller.readByPageFeed)
	routeFeed.GET("/followed", controller.readFollowedFeed)
	routeFeed.POST("/create", controller.createFeed)
	routeFeed.PUT("/update", controller.updateFeed)
	routeFeed.DELETE("/", controller.deleteFeed)

	return &controller
}

func (f *FeedController) readFeed(ctx *gin.Context) {
	var readFeed []entity.Feed
	err := f.fUC.Read(&readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	f.SuccessResponse(ctx, readFeed)
}

func (f *FeedController) readAccountFeed(ctx *gin.Context) {
	var readFeed dto.ReadPage
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	resFeed, err := f.fUC.ReadByAccountID(readFeed.ID, readFeed.Page, readFeed.PageLim)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	f.SuccessResponse(ctx, resFeed)
}

func (f *FeedController) readFollowedFeed(ctx *gin.Context) {
	var readFeed dto.ReadPage
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	resFeed, err := f.fUC.ReadByFollowerAccountID(readFeed.ID, readFeed.Page, readFeed.PageLim)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	f.SuccessResponse(ctx, resFeed)
}

func (f *FeedController) readCategoryFeed(ctx *gin.Context) {
	var readFeed dto.ReadPage
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	resFeed, err := f.fUC.ReadByProfileCategory(readFeed.Cat, readFeed.Page, readFeed.PageLim)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	f.SuccessResponse(ctx, resFeed)
}

func (f *FeedController) readByPageFeed(ctx *gin.Context) {
	var readFeed dto.ReadPage
	var responseFeed []dto.ResponseFeed
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	resFeed, err := f.fUC.ReadByPage(readFeed.Page, readFeed.PageLim)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	var holdFeed dto.ResponseFeed
	for _, feed := range resFeed {
		holdFeed.AccountID = feed.AccountID
		holdFeed.CaptionPost = feed.CaptionPost
		links := strings.Split(feed.DetailMediaFeeds, ",")
		for _, link := range links[0 : len(links)-2] {
			holdFeed.MediaLinks = append(holdFeed.MediaLinks, link)
		}
		responseFeed = append(responseFeed, holdFeed)
	}
	f.SuccessResponse(ctx, responseFeed)
}

func (f *FeedController) deleteFeed(ctx *gin.Context) {
	var readFeed dto.DeleteFeed
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	err = f.fUC.Delete(readFeed.ID)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	delMessage := fmt.Sprintf("Delete feed id %d success", readFeed.ID)
	f.SuccessResponse(ctx, delMessage)
}

func (f *FeedController) createFeed(ctx *gin.Context) {
	var createFeed dto.RequestFeed
	var feedInput entity.Feed
	err := f.ParseBodyRequest(ctx, &createFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	if createFeed.AccountID == 0 {
		f.FailedResponse(ctx, utils.RequiredError("account ID"))
		return
	} else if createFeed.CaptionPost == "" {
		f.FailedResponse(ctx, utils.RequiredError("feed caption"))
		return
	} else if createFeed.MediaLinks == nil {
		f.FailedResponse(ctx, utils.RequiredError("photos/videos"))
		return
	}
	feedInput.AccountID = createFeed.AccountID
	feedInput.CaptionPost = createFeed.CaptionPost
	var holdLink string
	for _, link := range createFeed.MediaLinks {
		holdLink = holdLink + link + ","
	}
	feedInput.DetailMediaFeeds = holdLink
	err = f.fUC.Create(&feedInput)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	f.SuccessResponse(ctx, feedInput)
}

func (f *FeedController) updateFeed(ctx *gin.Context) {
	var requestUpdateFeed dto.RequestUpdateFeed
	var updateFeed entity.Feed
	err := f.ParseBodyRequest(ctx, &requestUpdateFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	if requestUpdateFeed.FeedID == 0 {
		f.FailedResponse(ctx, errors.New("no feed found"))
		return
	}
	updateFeed.ID = requestUpdateFeed.FeedID
	err = f.fUC.ReadByID(&updateFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	var holdLink string
	for _, link := range requestUpdateFeed.MediaLinks {
		holdLink = holdLink + link + ","
	}
	updateFeed.CaptionPost = requestUpdateFeed.CaptionPost
	updateFeed.DetailMediaFeeds = holdLink
	err = f.fUC.Update(&updateFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	f.SuccessResponse(ctx, updateFeed)
}
