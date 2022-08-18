package controller

import (
	"errors"
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

func NewAuthController(router *gin.Engine, at auth.Token) *AuthController {
	controller := AuthController{
		router:    router,
		authToken: at,
	}
	routerAuth := controller.router.Group(("/auth"))
	routerAuth.POST("/register")
	routerAuth.POST("/login/:business", controller.loginUser)

	return &controller
}

func (ac *AuthController) createUserAccount(ctx *gin.Context) {
	var (
		customerReq     dto.RegisterUserRequest
		createdUserData entity.User
		createdUser     entity.UserCredential
	)
	err := ac.ParseBodyRequest(ctx, &customerReq)
	if customerReq.Username == "" {
		ac.FailedResponse(ctx, utils.RequiredError("username"))
		return
	} else if customerReq.Email == "" {
		ac.FailedResponse(ctx, utils.RequiredError("password"))
		return
	} else if customerReq.Password == "" {
		ac.FailedResponse(ctx, utils.RequiredError("customer name"))
		return
	} else if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	createdUserData.Username = customerReq.Username
	if err = ac.userUC.CreateUser(&createdUserData); err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	createdUser.Username = customerReq.Username
	createdUser.Password = customerReq.Password
	createdUser.Email = customerReq.Email
	createdUser.ID = createdUserData.ID
	createdUser.Encode()
	if err = ac.authUC.CreateUser(&createdUser); err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	ac.SuccessResponse(ctx, createdUser)
}

func (ac *AuthController) loginUser(ctx *gin.Context) {
	readRole, err := utils.StringToInt64(ctx.Param("business"))
	if err != nil {
		ac.FailedResponse(ctx, err)
	}
	var (
		user     entity.UserCredential
		realUser entity.UserCredential
	)

	err = ac.ParseBodyRequest(ctx, &user)
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
	userV, err := ac.authUC.FindUser(&realUser)
	if err != nil {
		ac.FailedResponse(ctx, errors.New("wrong username"))
		return
	}
	realUser.Decode()
	if realUser.Password != user.Password {
		ac.FailedResponse(ctx, errors.New("wrong password"))
		return
	}
	generateToken, err := ac.authToken.CreateAccessToken(&user)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	if err := ac.authToken.StoreAccessToken(user.Username, generateToken); err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	if userV.Role != uint(readRole) {
		err = errors.New("not a bussiness account")
		ac.FailedResponse(ctx, err)
	}
	ac.SuccessResponse(ctx, generateToken.AccessToken)
}
