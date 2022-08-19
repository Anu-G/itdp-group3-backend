package controller

import (
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"

	"github.com/gin-gonic/gin"
)

type UserController struct {
	router     *gin.Engine
	uUC        usecase.UserUsecase
	middleware middleware.AuthTokenMiddleware
	api.BaseApi
}

func NewUserController(router *gin.Engine, uUc usecase.UserUsecase, md middleware.AuthTokenMiddleware) {
	controller := UserController{
		router:     router,
		uUC:        uUc,
		middleware: md,
	}
	routeUser := controller.router.Group("/user")
	routeUser.Use(md.RequireToken())
	routeUser.GET("/", controller.readUser)
	routeUser.PUT("/update", controller.updateUser)
}

func (u *UserController) readUser(ctx *gin.Context) {
	var readUser entity.User
	err := u.ParseBodyRequest(ctx, &readUser)
	if err != nil {
		u.FailedResponse(ctx, err)
	}
	err = u.uUC.FindByUsername(&readUser)
	if err != nil {
		u.FailedResponse(ctx, err)
	}
	u.SuccessResponse(ctx, readUser)
}

func (u *UserController) updateUser(ctx *gin.Context) {
	var newUser entity.User
	err := u.ParseBodyRequest(ctx, &newUser)
	if err != nil {
		u.FailedResponse(ctx, err)
	}
	newUser.Encode()
	err = u.uUC.Update(&newUser)
	if err != nil {
		u.FailedResponse(ctx, err)
	}
	u.SuccessResponse(ctx, newUser)
}
