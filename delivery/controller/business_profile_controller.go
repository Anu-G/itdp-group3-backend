package controller

import (
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/usecase"

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
	routeBusinessProfile.POST("/add/profile-image")
	routeBusinessProfile.POST("/add/business-profile")

	return &controller
}

// func (b *BusinessProfileController) addBusinessProfile(ctx *gin.Context) {
// 	var (
// 		businessProfileReq dto.BusinessProfileRequest
// 		createdBp          entity.BusinessProfile
// 	)

// 	err := b.ParseBodyRequest(ctx, &businessProfileReq)
// 	if businessProfileReq.Address == "" {
// 		b.FailedResponse(ctx, utils.RequiredError("address"))
// 		return
// 	}

// 	createdBp, err = b.usecase.CreateBusinessProfile(&businessProfileReq, file, fileName[1])
// 	if err != nil {
// 		b.FailedResponse(ctx, err)
// 		return
// 	}
// 	b.SuccessResponse(ctx, createdBp)
// }

// func (b *BusinessProfileController) addProfileImage(ctx *gin.Context) {

// 	file, fileHeader, err := ctx.Request.FormFile("profileImage")
// 	if err != nil {
// 		b.FailedResponse(ctx, errors.New("failed get file"))
// 		return
// 	}

// 	fileName := strings.Split(fileHeader.Filename, ".")
// 	if len(fileName) != 2 {
// 		b.FailedResponse(ctx, errors.New("Unrecognized file extension"))
// 	}

// }
