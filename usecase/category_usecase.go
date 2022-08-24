package usecase

import (
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type CategoryUsecase interface {
	Create(cat *entity.Category) error
	ReadAll(cat *[]entity.Category) error
}

type categoryUsecase struct {
	repo repository.CategoryRepository
}

func NewCategoryUsecase(repo repository.CategoryRepository) CategoryUsecase {
	return &categoryUsecase{
		repo: repo,
	}
}

func (catc *categoryUsecase) Create(cat *entity.Category) error {
	return catc.repo.Create(cat)
}

func (catc *categoryUsecase) ReadAll(cat *[]entity.Category) error {
	return catc.repo.ReadAll(cat)
}
