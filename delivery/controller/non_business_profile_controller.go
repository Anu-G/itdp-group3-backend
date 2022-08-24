package controller

import (
	"errors"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"
	"itdp-group3-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type NonBusinessProfileController struct {
	router  *gin.Engine
	usecase usecase.NonBusinessProfileUseCaseInterface
	api.BaseApi
	middleware middleware.AuthTokenMiddleware
}

func NewNonBusinessProfileController(router *gin.Engine, uc usecase.NonBusinessProfileUseCaseInterface, middleware middleware.AuthTokenMiddleware) *NonBusinessProfileController {
	controller := NonBusinessProfileController{
		router:     router,
		usecase:    uc,
		middleware: middleware,
	}

	routeNonBusinessProfile := controller.router.Group("/non-business-profile")
	routeNonBusinessProfile.Use(middleware.RequireToken())
	routeNonBusinessProfile.POST("/add/profile", controller.addNonBusinessProfile)
	routeNonBusinessProfile.POST("/add/profile-image", controller.addProfileImage)
	routeNonBusinessProfile.POST("/get/profile", controller.getProfile)
	routeNonBusinessProfile.POST("/get/profile-image", controller.getProfileImage)

	return &controller
}

func (b *NonBusinessProfileController) addNonBusinessProfile(ctx *gin.Context) {
	var (
		nonBusinessProfileReq dto.NonBusinessProfileRequest
		createdBp             entity.NonBusinessProfile
	)

	err := b.ParseBodyRequest(ctx, &nonBusinessProfileReq)
	if nonBusinessProfileReq.AccountID == "" {
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	} else if nonBusinessProfileReq.DisplayName == "" {
		b.FailedResponse(ctx, utils.RequiredError("display_name"))
		return
	} else if err != nil {
		b.FailedResponse(ctx, err)
	}

	createdBp, err = b.usecase.CreateNonBusinessProfile(&nonBusinessProfileReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}
	b.SuccessResponse(ctx, createdBp)
}

func (b *NonBusinessProfileController) addProfileImage(ctx *gin.Context) {
	file, fileHeader, err := ctx.Request.FormFile("profile_image")
	if err != nil {
		b.FailedResponse(ctx, errors.New("failed get file"))
		return
	}

	fileName := strings.Split(fileHeader.Filename, ".")
	if len(fileName) != 2 {
		b.FailedResponse(ctx, errors.New("Unrecognized file extension"))
	}

	fileLocation, err := b.usecase.CreateProfileImage(file, fileName[1])
	if err != nil {
		b.FailedResponse(ctx, errors.New("failed while saving file"))
		return
	}

	b.SuccessResponse(ctx, fileLocation)
}

func (b *NonBusinessProfileController) getProfile(ctx *gin.Context) {
	var (
		nonBusinessProfileReq dto.NonBusinessProfileRequest
		nonBusinessProfileRes dto.NonBusinessProfileResponse
	)

	err := b.ParseBodyRequest(ctx, &nonBusinessProfileReq)
	if nonBusinessProfileReq.AccountID == "" {
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	} else if err != nil {
		b.FailedResponse(ctx, err)
	}

	nonBusinessProfileRes, err = b.usecase.GetNonBusinessProfile(&nonBusinessProfileReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	b.SuccessResponse(ctx, nonBusinessProfileRes)
}

func (b *NonBusinessProfileController) getProfileImage(ctx *gin.Context) {
	var (
		nonBusinessProfileReq dto.NonBusinessProfileRequest
		nonBusinessProfileRes dto.NonBusinessProfileResponse
	)

	err := b.ParseBodyRequest(ctx, &nonBusinessProfileReq)
	if nonBusinessProfileReq.AccountID == "" {
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	} else if err != nil {
		b.FailedResponse(ctx, err)
	}

	nonBusinessProfileRes, err = b.usecase.GetNonBusinessProfile(&nonBusinessProfileReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	b.SuccessDownload(ctx, nonBusinessProfileRes.NonBusinessProfile.ProfileImage)
	// b.SuccessResponse(ctx, nonBusinessProfileRes.NonBusinessProfile.ProfileImage)
}