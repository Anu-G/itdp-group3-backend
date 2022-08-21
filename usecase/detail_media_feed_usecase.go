package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type DetailMediaFeedUsecase interface {
	Create(fm *entity.DetailMediaFeed) error
	Read(fm *entity.DetailMediaFeed) error
}

type detailMediaDetailMediaFeedUsecase struct {
	repo repository.DetailMediaFeedRepository
}

func NewDetailMediaFeedUsecase(repo repository.DetailMediaFeedRepository) DetailMediaFeedUsecase {
	return &detailMediaDetailMediaFeedUsecase{
		repo: repo,
	}
}

func (fmc *detailMediaDetailMediaFeedUsecase) Create(fm *entity.DetailMediaFeed) error {
	return fmc.repo.Create(fm)
}

func (fmc *detailMediaDetailMediaFeedUsecase) Read(fm *entity.DetailMediaFeed) error {
	return fmc.repo.Read(fm)
}
