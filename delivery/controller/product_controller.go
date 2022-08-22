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
	"github.com/google/uuid"
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
	} else if len(productReq.DetailMediaProducts) == 0 {
		b.FailedResponse(ctx, utils.RequiredError("product must have minimum 1 image"))
	}

	createdProduct, err = b.usecase.CreateProduct(&productReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}
	b.SuccessResponse(ctx, createdProduct)
}

func (b *ProductController) addProductImage(ctx *gin.Context) {
	var detailMediaProducts []dto.DetailMediaProduct

	form, err := ctx.MultipartForm()
	files := form.File["product_images"]

	if err != nil {
		b.FailedResponse(ctx, errors.New("failed get file"))
		return
	}

	for _, file := range files {
		newFileName := strings.Split(file.Filename, ".")
		if len(newFileName) != 2 {
			b.FailedResponse(ctx, errors.New("Unrecognize file extension"))
			return
		}

		path := `E:\ITDP Sinarmas Mining\toktok_dev\img\` + "img-product-" + uuid.New().String() + "." + newFileName[1]

		if err := ctx.SaveUploadedFile(file, path); err != nil {
			b.FailedResponse(ctx, errors.New("failed while saving file"))
			return
		}

		detailMediaProducts = append(detailMediaProducts, dto.DetailMediaProduct{
			MediaLink: path,
		})
	}

	b.SuccessResponse(ctx, detailMediaProducts)
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
