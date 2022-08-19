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
	routeBusinessProfile.POST("/add/profile-image", controller.addProfileImage)
	routeBusinessProfile.POST("/add/business-profile", controller.addBusinessProfile)

	return &controller
}

func (b *BusinessProfileController) addBusinessProfile(ctx *gin.Context) {
	var (
		businessProfileReq dto.BusinessProfileRequest
		createdBp          entity.BusinessProfile
	)

	err := b.ParseBodyRequest(ctx, &businessProfileReq)
	if businessProfileReq.Address == "" {
		b.FailedResponse(ctx, utils.RequiredError("address"))
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
