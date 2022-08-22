package controller

import (
	"errors"
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"
	"strings"

	"github.com/gin-gonic/gin"
	"github.com/google/uuid"
)

type DetailMediaFeedController struct {
	router     *gin.Engine
	fmUC       usecase.DetailMediaFeedUsecase
	middleware middleware.AuthTokenMiddleware
	api.BaseApi
}

func NewDetailMediaFeedController(router *gin.Engine, fmUC usecase.DetailMediaFeedUsecase, md middleware.AuthTokenMiddleware) *DetailMediaFeedController {
	controller := DetailMediaFeedController{
		router:     router,
		fmUC:       fmUC,
		middleware: md,
	}

	routeDetailMediaFeed := controller.router.Group("/mediafeed")
	routeDetailMediaFeed.Use(md.RequireToken())
	routeDetailMediaFeed.GET("/", controller.readDetailMediaFeed)
	routeDetailMediaFeed.POST("/create", controller.createDetailMediaFeed)

	return &controller
}

func (fm *DetailMediaFeedController) readDetailMediaFeed(ctx *gin.Context) {
	var readDetailMediaFeed entity.DetailMediaFeed
	err := fm.ParseBodyRequest(ctx, &readDetailMediaFeed)
	if err != nil {
		fm.FailedResponse(ctx, err)
		return
	}
	err = fm.fmUC.Read(&readDetailMediaFeed)
	if err != nil {
		fm.FailedResponse(ctx, err)
		return
	}
	fm.SuccessResponse(ctx, readDetailMediaFeed)
}

func (fm *DetailMediaFeedController) createDetailMediaFeed(ctx *gin.Context) {
	var detailMediaFeed []dto.DetailMediaFeed

	form, err := ctx.MultipartForm()
	files := form.File["feed_images"]

	if err != nil {
		fm.FailedResponse(ctx, errors.New("failed get file"))
		return
	}

	for _, file := range files {
		newFileName := strings.Split(file.Filename, ".")
		if len(newFileName) != 2 {
			fm.FailedResponse(ctx, errors.New("Unrecognize file extension"))
			return
		}

		path := `E:\` + "img-feed-" + uuid.New().String() + "." + newFileName[1]

		if err := ctx.SaveUploadedFile(file, path); err != nil {
			fm.FailedResponse(ctx, errors.New("failed while saving file"))
			return
		}

		detailMediaFeed = append(detailMediaFeed, dto.DetailMediaFeed{
			MediaLink: path,
		})
	}

	fm.SuccessResponse(ctx, detailMediaFeed)

}
