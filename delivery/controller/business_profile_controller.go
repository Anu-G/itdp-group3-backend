package controller

import (
	"errors"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"
	"itdp-group3-backend/utils"
	"strings"

	"github.com/gin-gonic/gin"
)

type BusinessProfileController struct {
	router  *gin.Engine
	usecase usecase.BusinessProfileUseCaseInterface
	api.BaseApi
}

func NewBusinessProfileController(router *gin.Engine, uc usecase.BusinessProfileUseCaseInterface) *BusinessProfileController {
	controller := BusinessProfileController{
		router:  router,
		usecase: uc,
	}

	routeBusinessProfile := controller.router.Group("/business-profile")
	routeBusinessProfile.POST("/add/profile", controller.addBusinessProfile)
	routeBusinessProfile.POST("/add/profile-image", controller.addProfileImage)
	routeBusinessProfile.POST("/get/business-profile", controller.getBusinessProfile)

	return &controller
}

func (b *BusinessProfileController) addBusinessProfile(ctx *gin.Context) {
	var (
		businessProfileReq dto.BusinessProfileRequest
		createdBp          entity.BusinessProfile
	)

	err := b.ParseBodyRequest(ctx, &businessProfileReq)
	if businessProfileReq.AccountID == ""{
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	}else if businessProfileReq.CategoryID == ""{
		b.FailedResponse(ctx, utils.RequiredError("category"))
		return
	}else if businessProfileReq.Address == "" {
		b.FailedResponse(ctx, utils.RequiredError("address"))
		return
	}else if businessProfileReq.DisplayName == ""{
		b.FailedResponse(ctx, utils.RequiredError("display_name"))
		return
	}

	createdBp, err = b.usecase.CreateBusinessProfile(&businessProfileReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}
	b.SuccessResponse(ctx, createdBp)
}

func (b *BusinessProfileController) addProfileImage(ctx *gin.Context) {
	accountId := ctx.PostForm("account_id")

	file, fileHeader, err := ctx.Request.FormFile("profile_image")
	if err != nil {
		b.FailedResponse(ctx, errors.New("failed get file"))
		return
	}

	fileName := strings.Split(fileHeader.Filename, ".")
	if len(fileName) != 2 {
		b.FailedResponse(ctx, errors.New("Unrecognized file extension"))
	}

	fileLocation, err := b.usecase.CreateProfileImage(accountId, file, fileName[1])
	if err != nil {
		b.FailedResponse(ctx, errors.New("failed while saving file"))
		return
	}

	b.SuccessResponse(ctx, fileLocation)
}

func (b *BusinessProfileController) getBusinessProfile(ctx *gin.Context) {
	var (
		businessProfileReq dto.BusinessProfileRequest
		businessProfileRes dto.BusinessProfileResponse
	)

	err := b.ParseBodyRequest(ctx, &businessProfileReq)
	if businessProfileReq.AccountID == ""{
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	}

	businessProfileRes, err = b.usecase.GetBusinessProfile(&businessProfileReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	b.SuccessResponse(ctx, businessProfileRes)	
}