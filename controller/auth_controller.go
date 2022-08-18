package controller

import (
	"errors"
	"fmt"
	"itdp-group3-backend/auth"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"
	"itdp-group3-backend/utils"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router    *gin.Engine
	authToken auth.Token
	authUC    usecase.AuthUsecase
	userUC    usecase.UserUsecase
	api.BaseApi
}

func NewAuthController(router *gin.Engine, au usecase.AuthUsecase, uu usecase.UserUsecase, at auth.Token) *AuthController {
	controller := AuthController{
		router:    router,
		authUC:    au,
		userUC:    uu,
		authToken: at,
	}
	routerAuth := controller.router.Group(("/auth"))
	routerAuth.POST("/register", controller.createUserAccount)
	routerAuth.POST("/login", controller.loginUser)

	return &controller
}

func (ac *AuthController) createUserAccount(ctx *gin.Context) {
	var (
		userReq     dto.RegisterUserRequest
		createdUser entity.User
	)
	err := ac.ParseBodyRequest(ctx, &userReq)
	if userReq.Username == "" {
		ac.FailedResponse(ctx, utils.RequiredError("username"))
		return
	} else if userReq.Email == "" {
		ac.FailedResponse(ctx, utils.RequiredError("password"))
		return
	} else if userReq.Password == "" {
		ac.FailedResponse(ctx, utils.RequiredError("customer name"))
		return
	} else if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	createdUser.Username = userReq.Username
	createdUser.Password = userReq.Password
	createdUser.Email = userReq.Email

	fmt.Println("controller", userReq)
	createdUser.Encode()
	if err = ac.authUC.CreateUser(&createdUser); err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	ac.SuccessResponse(ctx, createdUser)
}

func (ac *AuthController) loginUser(ctx *gin.Context) {
	var (
		user     entity.User
		realUser entity.User
	)

	err := ac.ParseBodyRequest(ctx, &user)
	if user.Username == "" {
		ac.FailedResponse(ctx, utils.RequiredError("username"))
		return
	} else if user.Password == "" {
		ac.FailedResponse(ctx, utils.RequiredError("password"))
		return
	} else if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	realUser.Username = user.Username
	err = ac.authUC.FindUser(&realUser)
	if err != nil {
		ac.FailedResponse(ctx, errors.New("wrong username"))
		return
	}
	realUser.Decode()
	if realUser.Password != user.Password {
		ac.FailedResponse(ctx, errors.New("wrong password"))
		return
	}
	generateToken, err := ac.authToken.CreateAccessToken(&realUser)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	if err := ac.authToken.StoreAccessToken(user.Username, generateToken); err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	// if userV.Role != uint(readRole) {
	// 	err = errors.New("not a bussiness account")
	// 	ac.FailedResponse(ctx, err)
	// }
	ac.SuccessResponse(ctx, generateToken.AccessToken)
}
