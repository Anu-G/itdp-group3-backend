package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type DetailCommentUsecase interface {
	Create(cm *entity.DetailComment) error
	Read(cm *entity.DetailComment) error
	Delete(cm *entity.DetailComment) error
}

type detailCommentUsecase struct {
	repo repository.DetailCommentRepository
}

func NewDetailCommentUsecase(repo repository.DetailCommentRepository) DetailCommentUsecase {
	return &detailCommentUsecase{
		repo: repo,
	}
}

func (cmc *detailCommentUsecase) Create(cm *entity.DetailComment) error {
	return cmc.repo.Create(cm)
}

func (cmc *detailCommentUsecase) Read(cm *entity.DetailComment) error {
	return cmc.repo.Read(cm)
}

func (cmc *detailCommentUsecase) Delete(cm *entity.DetailComment) error {
	return cmc.repo.Delete(cm)
}
