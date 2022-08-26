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
	"strings"

	"github.com/gin-gonic/gin"
)

type AuthController struct {
	router    *gin.Engine
	authToken auth.Token
	authUC    usecase.AuthUsecase
	userUC    usecase.UserUsecase
	api.BaseApi
}

type authHeader struct {
	AuthorizationHeader string `header:"Authorization"`
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
	routerAuth.POST("/logout", controller.logoutUser)

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
		ac.FailedResponse(ctx, utils.RequiredError("email"))
		return
	} else if userReq.Password == "" {
		ac.FailedResponse(ctx, utils.RequiredError("password"))
		return
	} else if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	if !utils.EmailValidation(userReq.Email) || !utils.PasswordValidation(userReq.Password) {
		ac.FailedResponse(ctx, errors.New("email or password is invalid"))
		return
	}
	createdUser.Username = userReq.Username
	createdUser.Password = userReq.Password
	createdUser.Email = userReq.Email
	createdUser.Account.Username = userReq.Username
	createdUser.Account.RoleID = 1
	createdUser.Encrypt()
	if len(createdUser.Username) > 15 {
		createdUser.Account.PhoneNumber = createdUser.Username[0:15]
	} else {
		createdUser.Account.PhoneNumber = createdUser.Username
	}
	if err = ac.authUC.CreateUser(&createdUser); err != nil {
		ac.FailedResponse(ctx, err)
		return
	}

	responseRegister := fmt.Sprintf("user %s registered successfully", createdUser.Username)
	ac.SuccessResponse(ctx, responseRegister)
}

func (ac *AuthController) loginUser(ctx *gin.Context) {
	var (
		user     entity.User
		realUser entity.User
	)

	err := ac.ParseBodyRequest(ctx, &user)
	if user.Email == "" {
		ac.FailedResponse(ctx, utils.RequiredError("email"))
		return
	} else if user.Password == "" {
		ac.FailedResponse(ctx, utils.RequiredError("password"))
		return
	} else if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	if !utils.EmailValidation(user.Email) || !utils.PasswordValidation(user.Password) {
		ac.FailedResponse(ctx, errors.New("email or password is invalid"))
		return
	}
	realUser.Email = user.Email
	err = ac.authUC.FindUser(&realUser)
	if err != nil {
		ac.FailedResponse(ctx, errors.New("wrong email"))
		return
	}
	realUser.Decrypt()
	if realUser.Password != user.Password {
		ac.FailedResponse(ctx, errors.New("wrong password"))
		return
	}
	generateToken, err := ac.authToken.CreateAccessToken(&realUser)
	if err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	if err := ac.authToken.StoreAccessToken(generateToken.AccessUuid, generateToken); err != nil {
		ac.FailedResponse(ctx, err)
		return
	}
	ac.SuccessResponse(ctx, generateToken.AccessToken)
}

func (ac *AuthController) logoutUser(ctx *gin.Context) {
	h := authHeader{}
	if err := ctx.ShouldBindHeader(&h); err != nil {
		ac.FailedResponse(ctx, errors.New("unauthorized"))
		return
	}

	tokenStr := strings.Replace(h.AuthorizationHeader, "Bearer ", "", -1)
	if tokenStr == "" {
		ac.FailedResponse(ctx, errors.New("unauthorized"))
		return
	}

	token, err := ac.authToken.VerifyAccessToken(tokenStr)
	if err != nil {
		ac.FailedResponse(ctx, errors.New("unauthorized"))
		return
	}
	tokenUuid, err := ac.authToken.DeleteAccessToken(token)
	if err != nil {
		ac.FailedResponse(ctx, errors.New("unauthorized"))
		return
	}
	responseDel := fmt.Sprintf("deleted token : %v", tokenUuid)
	ac.SuccessResponse(ctx, responseDel)
}
