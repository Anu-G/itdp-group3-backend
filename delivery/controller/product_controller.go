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

type ProductController struct {
	router  *gin.Engine
	usecase usecase.ProductUseCaseInterface
	api.BaseApi
}

func NewProductController(router *gin.Engine, uc usecase.ProductUseCaseInterface) *ProductController {
	controller := ProductController{
		router:  router,
		usecase: uc,
	}

	routeProduct := controller.router.Group("/product")
	routeProduct.POST("/add/product", controller.addProduct)
	routeProduct.POST("/add/product-image", controller.addProductImage)
	// routeProduct.POST("/get/profile", controller.getProfile)
	// routeProduct.POST("/get/profile-image", controller.getProfileImage)

	return &controller
}

func (b *ProductController) addProduct(ctx *gin.Context) {
	var (
		productReq     dto.ProductRequest
		createdProduct entity.Product
	)

	err := b.ParseBodyRequest(ctx, &productReq)
	if productReq.AccountID == "" {
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	} else if productReq.ProductName == "" {
		b.FailedResponse(ctx, utils.RequiredError("product_name"))
		return
	} else if productReq.Price == "" {
		b.FailedResponse(ctx, utils.RequiredError("price"))
		return
	} else if productReq.Description == "" {
		b.FailedResponse(ctx, utils.RequiredError("description"))
		return
	}

	createdProduct, err = b.usecase.CreateProduct(&productReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}
	b.SuccessResponse(ctx, createdProduct)
}

func (b *ProductController) addProductImage(ctx *gin.Context) {
	file, fileHeader, err := ctx.Request.FormFile("product_image")
	if err != nil {
		b.FailedResponse(ctx, errors.New("failed get file"))
		return
	}

	fileName := strings.Split(fileHeader.Filename, ".")
	if len(fileName) != 2 {
		b.FailedResponse(ctx, errors.New("Unrecognized file extension"))
	}

	fileLocation, err := b.usecase.CreateProductImage(file, fileName[1])
	if err != nil {
		b.FailedResponse(ctx, errors.New("failed while saving file"))
		return
	}

	b.SuccessResponse(ctx, fileLocation)
}

// func (b *ProductController) getProfile(ctx *gin.Context) {
// 	var (
// 		productReq dto.ProductRequest
// 		productRes dto.ProductResponse
// 	)

// 	err := b.ParseBodyRequest(ctx, &productReq)
// 	if productReq.AccountID == "" {
// 		b.FailedResponse(ctx, utils.RequiredError("account_id"))
// 		return
// 	}

// 	productRes, err = b.usecase.GetProduct(&productReq)
// 	if err != nil {
// 		b.FailedResponse(ctx, err)
// 		return
// 	}

// 	b.SuccessResponse(ctx, productRes)
// }

// func (b *ProductController) getProfileImage(ctx *gin.Context) {
// 	var (
// 		productReq dto.ProductRequest
// 		productRes dto.ProductResponse
// 	)

// 	err := b.ParseBodyRequest(ctx, &productReq)
// 	if productReq.AccountID == "" {
// 		b.FailedResponse(ctx, utils.RequiredError("account_id"))
// 		return
// 	}

// 	productRes, err = b.usecase.GetProduct(&productReq)
// 	if err != nil {
// 		b.FailedResponse(ctx, err)
// 		return
// 	}

// 	b.SuccessDownload(ctx, productRes.Product.ProfileImage)
// }
