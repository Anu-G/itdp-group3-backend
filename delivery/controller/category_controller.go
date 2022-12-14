package controller

import (
	"errors"
	"fmt"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"

	"github.com/gin-gonic/gin"
)

type CategoryController struct {
	router     *gin.Engine
	catUC      usecase.CategoryUsecase
	middleware middleware.AuthTokenMiddleware
	api.BaseApi
}

func NewCategoryController(router *gin.Engine, catUC usecase.CategoryUsecase, md middleware.AuthTokenMiddleware) *CategoryController {
	controller := CategoryController{
		router:     router,
		catUC:      catUC,
		middleware: md,
	}

	routeCategory := controller.router.Group("/category")
	routeCategory.Use(controller.middleware.RequireToken())
	routeCategory.GET("/", controller.readAll)
	routeCategory.POST("/create", controller.create)

	return &controller
}

func (catc *CategoryController) readAll(ctx *gin.Context) {
	var readCategory []entity.Category
	var resRead dto.ReadCategoryResponse
	var responseRead []dto.ReadCategoryResponse
	err := catc.catUC.ReadAll(&readCategory)
	if err != nil {
		catc.FailedResponse(ctx, err)
		return
	}
	for _, cats := range readCategory {
		resRead.CategoryID = cats.ID
		resRead.CategoryNames = cats.CategoryName
		responseRead = append(responseRead, resRead)
	}
	catc.SuccessResponse(ctx, responseRead)
}

func (catc *CategoryController) create(ctx *gin.Context) {
	var requestCreate dto.CreateCategoryRequest
	var createCategory entity.Category
	err := catc.ParseBodyRequest(ctx, &requestCreate)
	if err != nil {
		catc.FailedResponse(ctx, err)
		return
	}
	if requestCreate.CategoryName == "" {
		catc.FailedResponse(ctx, errors.New("category name must be empty"))
		return
	}
	createCategory.CategoryName = requestCreate.CategoryName
	err = catc.catUC.Create(&createCategory)
	if err != nil {
		catc.FailedResponse(ctx, err)
		return
	}
	responseCreate := fmt.Sprintf("category %s created successfully", requestCreate.CategoryName)
	catc.SuccessResponse(ctx, responseCreate)
}
