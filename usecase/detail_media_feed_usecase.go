package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
	"mime/multipart"

	"github.com/gin-gonic/gin"
)

type DetailMediaFeedUsecase interface {
	Create(file *multipart.FileHeader, fileName string, ctx *gin.Context) (string, error)
	Read(fm *entity.DetailMediaFeed) error
}

type detailMediaDetailMediaFeedUsecase struct {
	repo     repository.DetailMediaFeedRepository
	fileRepo repository.FileRepository
}

func NewDetailMediaFeedUsecase(repo repository.DetailMediaFeedRepository, fileRepo repository.FileRepository) DetailMediaFeedUsecase {
	return &detailMediaDetailMediaFeedUsecase{
		repo:     repo,
		fileRepo: fileRepo,
	}
}

func (fmc *detailMediaDetailMediaFeedUsecase) Create(file *multipart.FileHeader, fileName string, ctx *gin.Context) (string, error) {
	return fmc.fileRepo.SavefromCtx(file, fileName, ctx)
}

func (fmc *detailMediaDetailMediaFeedUsecase) Read(fm *entity.DetailMediaFeed) error {
	return fmc.repo.Read(fm)
}
