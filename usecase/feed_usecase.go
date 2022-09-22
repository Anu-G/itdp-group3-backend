package usecase

import (
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type FeedUsecase interface {
	Create(f *entity.Feed) error
	Read(f *[]entity.Feed) error
	ReadByID(f *entity.Feed) error
	ReadDetailByID(id uint, page int, pageLim int) (dto.FeedDetailRequest, error)
	ReadForTimeline(page int, pageLim int) ([]dto.FeedDetailRequest, error)
	ReadByAccountID(id uint) ([]dto.FeedDetailRequest, error)
	ReadByProfileCategory(cat uint, page int, pageLim int) ([]dto.FeedDetailRequest, error)
	ReadByPage(page int, pageLim int) ([]entity.Feed, error)
	ReadByFollowerAccountID(id uint, page int, pageLim int) ([]dto.FeedDetailRequest, error)
	Update(f *entity.Feed) error
	Delete(id uint) error
	ReadForDetailTimeline(page int, pageLim int, feedId uint) ([]dto.FeedDetailRequest, error)
}

type feedUsecase struct {
	repo   repository.FeedRepository
	repoAc repository.AccountRepository
}

func NewFeedUsecase(repo repository.FeedRepository, repoAc repository.AccountRepository) FeedUsecase {
	return &feedUsecase{
		repo:   repo,
		repoAc: repoAc,
	}
}

func (fc *feedUsecase) Create(f *entity.Feed) error {
	return fc.repo.Create(f)
}

func (fc *feedUsecase) Read(f *[]entity.Feed) error {
	return fc.repo.Read(f)
}

func (fc *feedUsecase) ReadByID(f *entity.Feed) error {
	return fc.repo.ReadByID(f)
}

func (fc *feedUsecase) ReadDetailByID(id uint, page int, pageLim int) (dto.FeedDetailRequest, error) {
	return fc.repo.ReadDetailByID(id, page, pageLim)
}

func (fc *feedUsecase) ReadForTimeline(page int, pageLim int) ([]dto.FeedDetailRequest, error) {
	return fc.repo.ReadForTimeline(page, pageLim)
}

func (fc *feedUsecase) ReadByAccountID(id uint) ([]dto.FeedDetailRequest, error) {
	return fc.repo.ReadByAccountID(int(id))
}

func (fc *feedUsecase) ReadByProfileCategory(cat uint, page int, pageLim int) ([]dto.FeedDetailRequest, error) {
	return fc.repo.ReadByProfileCategory(cat, page, pageLim)
}

func (fc *feedUsecase) ReadByPage(page int, pageLim int) ([]entity.Feed, error) {
	return fc.repo.ReadByPage(page, pageLim)
}

func (fc *feedUsecase) ReadByFollowerAccountID(id uint, page int, pageLim int) ([]dto.FeedDetailRequest, error) {
	var accountInput entity.Account
	var accountIDList []uint
	accountInput.ID = id
	err := fc.repoAc.FindById(&accountInput)
	if err != nil {
		return nil, err
	}
	for _, follow := range accountInput.Followed {
		accountIDList = append(accountIDList, follow.FollowedAccountID)
	}
	feedList, err := fc.repo.ReadByFollowerAccountID(accountIDList, page, pageLim)
	if err != nil {
		return nil, err
	}
	return feedList, nil
}

func (fc *feedUsecase) Update(f *entity.Feed) error {
	return fc.repo.Update(f)
}

func (fc *feedUsecase) Delete(id uint) error {
	return fc.repo.Delete(id)
}

func (fc *feedUsecase) ReadForDetailTimeline(page int, pageLim int, feedId uint) ([]dto.FeedDetailRequest, error) {
	return fc.repo.ReadForDetailTimeline(page, pageLim, feedId)
}

