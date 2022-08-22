package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type FeedUsecase interface {
	Create(f *entity.Feed) error
	Read(f *entity.Feed) error
	ReadByAccountID(id uint, page int, pageLim int) ([]entity.Feed, error)
	ReadByProfileCategory(cat uint, page int, pageLim int) ([]entity.Feed, error)
	ReadByPage(page int, pageLim int) ([]entity.Feed, error)
	Delete(id uint) error
}

type feedUsecase struct {
	repo repository.FeedRepository
}

func NewFeedUsecase(repo repository.FeedRepository) FeedUsecase {
	return &feedUsecase{
		repo: repo,
	}
}

func (fc *feedUsecase) Create(f *entity.Feed) error {
	return fc.repo.Create(f)
}

func (fc *feedUsecase) Read(f *entity.Feed) error {
	return fc.repo.Read(f)
}

func (fc *feedUsecase) ReadByAccountID(id uint, page int, pageLim int) ([]entity.Feed, error) {
	return fc.repo.ReadByAccountID(id, page, pageLim)
}

func (fc *feedUsecase) ReadByProfileCategory(cat uint, page int, pageLim int) ([]entity.Feed, error) {
	return fc.repo.ReadByProfileCategory(cat, page, pageLim)
}

func (fc *feedUsecase) ReadByPage(page int, pageLim int) ([]entity.Feed, error) {
	return fc.repo.ReadByPage(page, pageLim)
}

func (fc *feedUsecase) Delete(id uint) error {
	return fc.repo.Delete(id)
}
