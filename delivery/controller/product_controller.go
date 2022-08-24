package controller

import (
	"errors"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"
	"itdp-group3-backend/utils"

	"github.com/gin-gonic/gin"
)

type ProductController struct {
	router  *gin.Engine
	usecase usecase.ProductUseCaseInterface
	api.BaseApi
	middleware middleware.AuthTokenMiddleware
}

func NewProductController(router *gin.Engine, uc usecase.ProductUseCaseInterface, middleware middleware.AuthTokenMiddleware) *ProductController {
	controller := ProductController{
		router:     router,
		usecase:    uc,
		middleware: middleware,
	}

	routeProduct := controller.router.Group("/product")
	routeProduct.Use(middleware.RequireToken())
	routeProduct.POST("/add/product", controller.addProduct)
	routeProduct.POST("/add/product-image", controller.addProductImage)
	routeProduct.POST("/get/by-account", controller.getByAccount)
	routeProduct.POST("/get/by-product", controller.getByProduct)
	routeProduct.POST("/delete/product", controller.deleteProduct)

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
	} else if err != nil {
		b.FailedResponse(ctx, err)
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
	form, err := ctx.MultipartForm()
	files := form.File["product_images"]

	if err != nil {
		b.FailedResponse(ctx, errors.New("failed get file"))
		return
	}

	detailMediaProducts, err := b.usecase.CreateProductImage(files, ctx)

	b.SuccessResponse(ctx, detailMediaProducts)
}

func (b *ProductController) getByAccount(ctx *gin.Context) {
	var (
		productReq dto.ProductRequest
		productRes []dto.ProductResponse
	)

	err := b.ParseBodyRequest(ctx, &productReq)
	if productReq.AccountID == "" {
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	} else if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	productRes, err = b.usecase.GetByAccount(productReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	b.SuccessResponse(ctx, productRes)
}

func (b *ProductController) getByProduct(ctx *gin.Context) {
	var (
		productReq dto.ProductRequest
		productRes dto.ProductResponse
	)

	err := b.ParseBodyRequest(ctx, &productReq)
	if productReq.AccountID == "" {
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	} else if productReq.ProductID == "" {
		b.FailedResponse(ctx, utils.RequiredError("product_id"))
	} else if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	productRes, err = b.usecase.GetByProduct(productReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	b.SuccessResponse(ctx, productRes)
}

func (b *ProductController) deleteProduct(ctx *gin.Context) {
	var (
		productReq dto.ProductRequest
	)

	err := b.ParseBodyRequest(ctx, &productReq)
	if productReq.ProductID == "" {
		b.FailedResponse(ctx, utils.RequiredError("product_id"))
	} else if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	err = b.usecase.Delete(productReq.ProductID)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	b.SuccessResponse(ctx, "success delete productID "+productReq.ProductID)
}
