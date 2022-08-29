package controller

import (
	"errors"
	"fmt"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"
	"strings"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	router     *gin.Engine
	accUC      usecase.AccountUsecase
	flUC       usecase.FollowUsecase
	middleware middleware.AuthTokenMiddleware
	api.BaseApi
}

func NewAccountController(router *gin.Engine, accUc usecase.AccountUsecase, middleware middleware.AuthTokenMiddleware, flUC usecase.FollowUsecase) *AccountController {
	controller := AccountController{
		router:     router,
		accUC:      accUc,
		flUC:       flUC,
		middleware: middleware,
	}
	routeAccount := controller.router.Group("/account")
	routeAccount.Use(middleware.RequireToken())
	routeAccount.POST("/product", controller.readAccountForProductDetail)
	routeAccount.POST("/feed", controller.readAccountForFeedDetail)
	routeAccount.PUT("/update", controller.createAccount)
	routeAccount.POST("/follow", controller.follow)
	routeAccount.POST("/unfollow", controller.unfollow)
	routeAccount.GET("/list", controller.showFollowList)
	routeAccount.PUT("/activate-business", controller.activateBusinessAccount)

	return &controller
}

func (ac *AccountController) readAccountForProductDetail(ctx *gin.Context) {
	var readAccount entity.Account
	var requestAccountID dto.AccountFillRequest
	var responseAccountHold dto.ProductDetailResponse
	var responseAccount []dto.ProductDetailResponse
	err := ac.ParseBodyRequest(ctx, &requestAccountID)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	readAccount.ID = requestAccountID.AccountID
	err = ac.accUC.ReadForProductDetail(&readAccount)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	for i := range readAccount.Products {
		responseAccountHold.ProductID = readAccount.Products[i].ID
		responseAccountHold.ProfileImage = readAccount.BusinessProfile.ProfileImage
		responseAccountHold.Name = readAccount.BusinessProfile.DisplayName
		responseAccountHold.ProductName = readAccount.Products[i].ProductName
		responseAccountHold.ProductPrice = readAccount.Products[i].Price
		links := strings.Split(readAccount.Products[i].DetailMediaProducts, ",")
		for i, link := range links {
			if i == len(links)-1 {
				break
			}
			responseAccountHold.DetailMediaProducts = append(responseAccountHold.DetailMediaProducts, link)
		}
		responseAccount = append(responseAccount, responseAccountHold)
	}

	ac.SuccessResponse(ctx, responseAccount)
}

func (ac *AccountController) readAccountForFeedDetail(ctx *gin.Context) {
	var readAccount entity.Account
	var requestAccountID dto.AccountFillRequest
	var responseAccountHold dto.FeedDetailResponse
	var responseAccount []dto.FeedDetailResponse
	err := ac.ParseBodyRequest(ctx, &requestAccountID)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	readAccount.ID = requestAccountID.AccountID
	err = ac.accUC.ReadForFeedDetail(&readAccount)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	for i := range readAccount.Feeds {
		responseAccountHold.PostID = readAccount.Feeds[i].ID
		responseAccountHold.ProfileImage = readAccount.BusinessProfile.ProfileImage
		responseAccountHold.CaptionPost = readAccount.Feeds[i].CaptionPost
		responseAccountHold.CreatedAt = readAccount.Feeds[i].CreatedAt
		links := strings.Split(readAccount.Feeds[i].DetailMediaFeeds, ",")
		for _, link := range links[0 : len(links)-2] {
			responseAccountHold.DetailMediaFeeds = append(responseAccountHold.DetailMediaFeeds, link)
		}
		responseAccountHold.DisplayName = readAccount.BusinessProfile.DisplayName
		responseAccountHold.DetailComment = readAccount.Feeds[i].DetailComments
		responseAccount = append(responseAccount, responseAccountHold)
	}

	ac.SuccessResponse(ctx, responseAccount)
}

func (ac *AccountController) showFollowList(ctx *gin.Context) {
	var requestFollowList dto.FollowListRequest
	err := ac.ParseBodyRequest(ctx, &requestFollowList)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	responseFollowList, err := ac.accUC.FollowList(requestFollowList)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	ac.SuccessResponse(ctx, responseFollowList)
}

func (ac *AccountController) createAccount(ctx *gin.Context) {
	var newAccount entity.Account
	err := ac.ParseBodyRequest(ctx, &newAccount)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	err = ac.accUC.Update(&newAccount)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	ac.SuccessResponse(ctx, newAccount)
}

func (ac *AccountController) activateBusinessAccount(ctx *gin.Context) {
	var requestAccountID dto.ActivateBusinessAccountRequest
	err := ac.ParseBodyRequest(ctx, &requestAccountID)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	responseAccountID, err := ac.accUC.UpdateByID(requestAccountID.AccountID)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	ac.SuccessResponse(ctx, responseAccountID)
}

func (ac *AccountController) follow(ctx *gin.Context) {
	var requestFollow dto.FollowRequest
	err := ac.ParseBodyRequest(ctx, &requestFollow)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	followCheck, err := ac.flUC.FindForVerif(&requestFollow)
	if followCheck.AccountID != 0 {
		ac.FailedResponse(ctx, errors.New("already followed"))
		return
	}
	err = ac.flUC.Create(&requestFollow)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	responseFollow := fmt.Sprintf("%d follow %d success", requestFollow.FollowerAccounID, requestFollow.FollowedAccountID)
	ac.SuccessResponse(ctx, responseFollow)
}

func (ac *AccountController) unfollow(ctx *gin.Context) {
	var requestUnfollow dto.UnfollowRequest
	err := ac.ParseBodyRequest(ctx, &requestUnfollow)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	followCheck, err := ac.flUC.FindForVerif(&requestUnfollow.FollowRequest)
	if followCheck.AccountID != 0 {
		ac.FailedResponse(ctx, errors.New("already followed"))
		return
	}
	err = ac.flUC.Delete(&requestUnfollow)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	responseFollow := fmt.Sprintf("%d unfollow %d success", requestUnfollow.FollowerAccounID, requestUnfollow.FollowedAccountID)
	ac.SuccessResponse(ctx, responseFollow)
}
