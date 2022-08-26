package controller

import (
	"itdp-group3-backend/delivery/api"
	"itdp-group3-backend/middleware"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/usecase"
	"itdp-group3-backend/utils"

	"github.com/gin-gonic/gin"
)

type FAQController struct {
	router  *gin.Engine
	usecase usecase.FAQUseCaseInterface
	api.BaseApi
	middleware middleware.AuthTokenMiddleware
}

func NewFAQController(router *gin.Engine, uc usecase.FAQUseCaseInterface, middleware middleware.AuthTokenMiddleware) *FAQController {
	controller := FAQController{
		router:     router,
		usecase:    uc,
		middleware: middleware,
	}

	routeFAQ := controller.router.Group("/faq")
	routeFAQ.Use(middleware.RequireToken())
	routeFAQ.POST("/add/faq", controller.addFAQ)
	routeFAQ.POST("/get/faq", controller.getFAQ)
	routeFAQ.POST("/delete/faq", controller.deleteFAQ)

	return &controller
}

func (b *FAQController) addFAQ(ctx *gin.Context) {
	var (
		faqReq     dto.FAQRequest
		createdFAQ entity.BusinessFAQ
	)

	err := b.ParseBodyRequest(ctx, &faqReq)
	if faqReq.AccountID == "" {
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	} else if faqReq.Answer == "" {
		b.FailedResponse(ctx, utils.RequiredError("answer"))
		return
	} else if faqReq.Question == "" {
		b.FailedResponse(ctx, utils.RequiredError("question"))
		return
	} else if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	createdFAQ, err = b.usecase.CreateFAQ(&faqReq)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}
	b.SuccessResponse(ctx, createdFAQ)
}

func (b *FAQController) getFAQ(ctx *gin.Context) {
	var (
		faqReq dto.FAQRequest
		faqRes []entity.BusinessFAQ
	)

	err := b.ParseBodyRequest(ctx, &faqReq)
	if faqReq.AccountID == "" {
		b.FailedResponse(ctx, utils.RequiredError("account_id"))
		return
	} else if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	faqRes, err = b.usecase.GetFAQByAccount(faqReq.AccountID)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	b.SuccessResponse(ctx, faqRes)
}

func (b *FAQController) deleteFAQ(ctx *gin.Context) {
	var (
		faqReq dto.FAQRequest
	)

	err := b.ParseBodyRequest(ctx, &faqReq)
	if faqReq.FAQID == "" {
		b.FailedResponse(ctx, utils.RequiredError("faq_id"))
		return
	} else if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	err = b.usecase.Delete(faqReq.FAQID)
	if err != nil {
		b.FailedResponse(ctx, err)
		return
	}

	b.SuccessResponse(ctx, "success delete faq accountID "+faqReq.AccountID)
}
