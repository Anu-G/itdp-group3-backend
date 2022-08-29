package controller

import (
	"errors"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"
	"itdp-group3-backend/utils"

	"github.com/cloudinary/cloudinary-go"
	"github.com/cloudinary/cloudinary-go/api/uploader"
	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type BusinessProfileController struct {
	router  *gin.Engine
	usecase usecase.BusinessProfileUseCaseInterface
	api.BaseApi
	middleware middleware.AuthTokenMiddleware
}

func NewBusinessProfileController(router *gin.Engine, uc usecase.BusinessProfileUseCaseInterface, middleware middleware.AuthTokenMiddleware) *BusinessProfileController {
	controller := BusinessProfileController{
		router:     router,
		usecase:    uc,
		middleware: middleware,
	}

	routeBusinessProfile := controller.router.Group("/business-profile")
	routeBusinessProfile.Use(middleware.RequireToken())
	routeBusinessProfile.POST("/add/profile", controller.addBusinessProfile)
	routeBusinessProfile.POST("/add/profile-image", controller.addProfileImage)
	routeBusinessProfile.POST("/get/profile", controller.getProfile)

	return &controller
}

func (b *BusinessProfileController) addBusinessProfile(ctx *gin.Context) {
	var (
		businessProfileReq dto.BusinessProfileRequest
		createdBp          entity.BusinessProfile
	)

	err := b.ParseBodyRequest(ctx, &businessProfileReq)
	if businessProfileReq.AccountID == "" {
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	} else if businessProfileReq.CategoryID == "" {
		b.FailedResponse(ctx, utils.RequiredError("category"))
		return
	} else if businessProfileReq.Address == "" {
		b.FailedResponse(ctx, utils.RequiredError("address"))
		return
	} else if businessProfileReq.DisplayName == "" {
		b.FailedResponse(ctx, utils.RequiredError("display_name"))
		return
	} else if err != nil {
		b.FailedResponse(ctx, err)
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
	cld, _ := cloudinary.NewFromParams("ihdiannaja","954945529412874","7mFstMRVYEOlO784FGNo09mfk_4")

	file, _, err := ctx.Request.FormFile("profile_image")
	if err != nil {
		b.FailedResponse(ctx, errors.New("failed get file"))
		return
	}

	newFileName := "img-business-profile" + uuid.New().String()

	result, err := cld.Upload.Upload(ctx, file, uploader.UploadParams{
		PublicID: newFileName,
		Folder: "Business Profile",
	})

	if err != nil {
		b.FailedResponse(ctx, errors.New("Upload to cloudinary failed"))
		return
	}

	b.SuccessResponse(ctx, result.SecureURL)
}

func (b *BusinessProfileController) getProfile(ctx *gin.Context) {
	var (
		businessProfileReq dto.BusinessProfileRequest
		businessProfileRes dto.BusinessProfileResponse
	)

	err := b.ParseBodyRequest(ctx, &businessProfileReq)
	if businessProfileReq.AccountID == "" {
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	} else if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	businessProfileRes, err = b.usecase.GetBusinessProfile(&businessProfileReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	b.SuccessResponse(ctx, businessProfileRes)
}
