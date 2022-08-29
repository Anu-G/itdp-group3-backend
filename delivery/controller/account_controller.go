package controller

import (
	"errors"
	"fmt"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"

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
	routeAccount.GET("/", controller.readAccount)
	routeAccount.PUT("/update", controller.createAccount)
	routeAccount.POST("/follow", controller.follow)
	routeAccount.POST("/unfollow", controller.unfollow)
	routeAccount.GET("/list", controller.showFollowList)
	routeAccount.PUT("/activate-business", controller.activateBusinessAccount)

	return &controller
}

func (ac *AccountController) readAccount(ctx *gin.Context) {
	var readAccount entity.Account
	err := ac.ParseBodyRequest(ctx, &readAccount)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	err = ac.accUC.FindByUsername(&readAccount)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	ac.SuccessResponse(ctx, readAccount)
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
