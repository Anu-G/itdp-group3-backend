package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type DetailMediaFeedUsecase interface {
	Create(feedID uint, filePath string) error
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

func (fmc *detailMediaDetailMediaFeedUsecase) Create(feedID uint, filePath string) error {
	var createFeed entity.DetailMediaFeed
	createFeed.FeedID = uint(feedID)
	createFeed.MediaLink = filePath
	err := fmc.repo.Create(&createFeed)
	if err != nil {
		return err
	}
	return nil
}

func (fmc *detailMediaDetailMediaFeedUsecase) Read(fm *entity.DetailMediaFeed) error {
	return fmc.repo.Read(fm)
}
