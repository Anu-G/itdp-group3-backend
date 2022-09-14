package usecase

import (
	"errors"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type DetailLikeUsecase interface {
	Like(dl *dto.LikeRequest) error
	Unlike(dl *dto.LikeRequest) error
}

type detailLikeUsecase struct {
	repo repository.DetailLikeRepository
}

func NewDetailLikeUsecase(repo repository.DetailLikeRepository) DetailLikeUsecase {
	return &detailLikeUsecase{
		repo: repo,
	}
}

func (dlu *detailLikeUsecase) Like(dl *dto.LikeRequest) error {
	var dlReq entity.DetailLike
	res, _ := dlu.repo.Find(dl)
	if res.FeedID != 0 {
		return errors.New("account already liked")
	}
	dlReq.AccountID = dl.AccountID
	dlReq.FeedID = dl.FeedID
	return dlu.repo.Like(&dlReq)
}

func (dlu *detailLikeUsecase) Unlike(dl *dto.LikeRequest) error {
	var dlReq entity.DetailLike
	dlReq.AccountID = dl.AccountID
	dlReq.FeedID = dl.FeedID
	return dlu.repo.Unlike(&dlReq)
}
