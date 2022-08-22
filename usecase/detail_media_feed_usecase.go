package usecase

import (
	"fmt"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
	"mime/multipart"
	"strconv"
)

type DetailMediaFeedUsecase interface {
	Create(feedID string, file multipart.File, fileName string) (string, error)
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

func (fmc *detailMediaDetailMediaFeedUsecase) Create(feedID string, file multipart.File, fileName string) (string, error) {
	fileExt := fmt.Sprintf("img-bp-%s.%s", feedID, fileName)
	fileLocation, err := fmc.fileRepo.Save(file, fileExt)

	if err != nil {
		return "", err
	}
	feedIDint, _ := strconv.Atoi(feedID)
	var createFeed entity.DetailMediaFeed
	createFeed.FeedID = uint(feedIDint)
	createFeed.MediaLink = fileLocation
	err = fmc.repo.Create(&createFeed)
	if err != nil {
		return "", err
	}
	return fileLocation, err
}

func (fmc *detailMediaDetailMediaFeedUsecase) Read(fm *entity.DetailMediaFeed) error {
	return fmc.repo.Read(fm)
}
