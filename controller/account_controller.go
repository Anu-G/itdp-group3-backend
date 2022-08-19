package controller

import (
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"

	"github.com/gin-gonic/gin"
)

type AccountController struct {
	router     *gin.Engine
	accUC      usecase.AccountUsecase
	middleware middleware.AuthTokenMiddleware
	api.BaseApi
}

func NewAccountController(router *gin.Engine, accUc usecase.AccountUsecase, middleware middleware.AuthTokenMiddleware) {
	controller := AccountController{
		router:     router,
		accUC:      accUc,
		middleware: middleware,
	}
	routeAccount := controller.router.Group("/account")
	routeAccount.Use(middleware.RequireToken())
	routeAccount.GET("/", controller.readAccount)
	routeAccount.PUT("/update", controller.createAccount)
}

func (ac *AccountController) readAccount(ctx *gin.Context) {
	var readAccount entity.Account
	userName := ac.middleware.OpenToken(ctx)
	readAccount.Username = userName
	err := ac.accUC.FindByUsername(&readAccount)
	if err != nil {
		ac.FailedResponse(ctx, err)
	}
	ac.SuccessResponse(ctx, readAccount)
}

func (ac *AccountController) createAccount(ctx *gin.Context) {
	var newAccount entity.Account
	var regAccount entity.Account
	userName := ac.middleware.OpenToken(ctx)
	err := ac.ParseBodyRequest(ctx, newAccount)
	if err != nil {
		ac.FailedResponse(ctx, err)
	}
	regAccount = newAccount
	regAccount.Username = userName
	err = ac.accUC.Update(&regAccount)
	if err != nil {
		ac.FailedResponse(ctx, err)
	}
	ac.SuccessResponse(ctx, newAccount)
}
