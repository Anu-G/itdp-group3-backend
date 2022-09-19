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
	flUC       usecase.DetailLikeUsecase
	middleware middleware.AuthTokenMiddleware
	api.BaseApi
}

func NewFeedController(router *gin.Engine, fUC usecase.FeedUsecase, fmUC usecase.DetailMediaFeedUsecase, flUC usecase.DetailLikeUsecase, md middleware.AuthTokenMiddleware) *FeedController {
	controller := FeedController{
		router:     router,
		fUC:        fUC,
		fmUC:       fmUC,
		flUC:       flUC,
		middleware: md,
	}

	routeFeed := controller.router.Group("/feed")
	routeFeed.Use(md.RequireToken())
	routeFeed.GET("/", controller.readFeed)
	routeFeed.POST("/account", controller.readAccountFeed)
	routeFeed.POST("/category", controller.readCategoryFeed)
	routeFeed.POST("/paged", controller.readByPageFeed)
	routeFeed.POST("/timeline", controller.readForTimeline)
	routeFeed.POST("/followed", controller.readFollowedFeed)
	routeFeed.POST("/create", controller.createFeed)
	routeFeed.POST("/update", controller.updateFeed)
	routeFeed.POST("/delete", controller.deleteFeed)
	routeFeed.POST("/like", controller.likeFeed)
	routeFeed.POST("/unlike", controller.unlikeFeed)

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
	var responseFeedTimeline []dto.FeedDetailResponse
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	resFeed, err := f.fUC.ReadByAccountID(readFeed.ID)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	for _, feed := range resFeed {
		links := strings.Split(feed.DetailMediaFeeds, ",")
		responseFeedTimeline = append(responseFeedTimeline, dto.FeedDetailResponse{
			AccountID:        feed.AccountID,
			PostID:           feed.PostID,
			DisplayName:      feed.DisplayName,
			CaptionPost:      feed.CaptionPost,
			ProfileImage:     feed.ProfileImage,
			CreatedAt:        feed.CreatedAt,
			DetailComment:    feed.DetailComment,
			DetailMediaFeeds: links,
			DetailLike:       feed.DetailLike,
			TotalLike:        len(feed.DetailLike),
		})
	}
	f.SuccessResponse(ctx, responseFeedTimeline)
}

func (f *FeedController) readForTimeline(ctx *gin.Context) {
	var readFeed dto.ReadPage
	var responseFeedTimeline []dto.FeedDetailResponse
	err := f.ParseBodyRequest(ctx, &readFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	resFeed, err := f.fUC.ReadForTimeline(readFeed.Page, readFeed.PageLim)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	for _, feed := range resFeed {
		links := strings.Split(feed.DetailMediaFeeds, ",")
		responseFeedTimeline = append(responseFeedTimeline, dto.FeedDetailResponse{
			AccountID:        feed.AccountID,
			PostID:           feed.PostID,
			DisplayName:      feed.DisplayName,
			CaptionPost:      feed.CaptionPost,
			ProfileImage:     feed.ProfileImage,
			CreatedAt:        feed.CreatedAt,
			DetailComment:    feed.DetailComment,
			DetailMediaFeeds: links,
			DetailLike:       feed.DetailLike,
			TotalLike:        len(feed.DetailLike),
		})
	}
	f.SuccessResponse(ctx, responseFeedTimeline)
}

func (f *FeedController) readFollowedFeed(ctx *gin.Context) {
	var readFeed dto.ReadPage
	var responseFollowedFeed []dto.FeedDetailResponse
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
	for _, feed := range resFeed {
		links := strings.Split(feed.DetailMediaFeeds, ",")
		responseFollowedFeed = append(responseFollowedFeed, dto.FeedDetailResponse{
			PostID:           feed.PostID,
			DisplayName:      feed.DisplayName,
			CaptionPost:      feed.CaptionPost,
			ProfileImage:     feed.ProfileImage,
			CreatedAt:        feed.CreatedAt,
			DetailComment:    feed.DetailComment,
			DetailLike:       feed.DetailLike,
			DetailMediaFeeds: links,
			TotalLike:        len(feed.DetailLike),
		})
	}
	f.SuccessResponse(ctx, responseFollowedFeed)
}

func (f *FeedController) readCategoryFeed(ctx *gin.Context) {
	var readFeed dto.ReadPage
	var responseCategoryFeed []dto.FeedDetailResponse
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
	for _, feed := range resFeed {
		links := strings.Split(feed.DetailMediaFeeds, ",")
		responseCategoryFeed = append(responseCategoryFeed, dto.FeedDetailResponse{
			AccountID:        feed.AccountID,
			PostID:           feed.PostID,
			DisplayName:      feed.DisplayName,
			CaptionPost:      feed.CaptionPost,
			ProfileImage:     feed.ProfileImage,
			CreatedAt:        feed.CreatedAt,
			DetailComment:    feed.DetailComment,
			DetailLike:       feed.DetailLike,
			DetailMediaFeeds: links,
			TotalLike:        len(feed.DetailLike),
		})
	}
	f.SuccessResponse(ctx, responseCategoryFeed)
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
	holdLink = strings.Join(createFeed.MediaLinks, ",")
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
	holdLink = strings.Join(requestUpdateFeed.MediaLinks, ",")
	updateFeed.CaptionPost = requestUpdateFeed.CaptionPost
	updateFeed.DetailMediaFeeds = holdLink
	err = f.fUC.Update(&updateFeed)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	f.SuccessResponse(ctx, updateFeed)
}

func (f *FeedController) likeFeed(ctx *gin.Context) {
	var requestLike dto.LikeRequest
	err := f.ParseBodyRequest(ctx, &requestLike)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	err = f.flUC.Like(&requestLike)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	f.SuccessResponse(ctx, requestLike)
}

func (f *FeedController) unlikeFeed(ctx *gin.Context) {
	var requestLike dto.LikeRequest
	err := f.ParseBodyRequest(ctx, &requestLike)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	err = f.flUC.Unlike(&requestLike)
	if err != nil {
		f.FailedResponse(ctx, err)
		return
	}
	f.SuccessResponse(ctx, requestLike)
}
