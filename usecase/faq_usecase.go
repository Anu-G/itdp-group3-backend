package usecase

import (
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
	"strconv"
)

type FAQUseCaseInterface interface {
	CreateFAQ(p *dto.FAQRequest) (entity.BusinessFAQ, error)
	GetFAQByAccount(id string) ([]entity.BusinessFAQ, error)
	Delete(id string) error
}

type faqUseCase struct {
	repo   repository.FAQRepositoryInterface
	bpRepo repository.BusinessProfileRepositoryInterface
}

func (pu *faqUseCase) Delete(id string) error {
	return pu.repo.Delete(id)
}

func (pu *faqUseCase) GetFAQByAccount(id string) ([]entity.BusinessFAQ, error) {
	return pu.repo.GetFAQByAccount(id)
}

func (pu *faqUseCase) CreateFAQ(p *dto.FAQRequest) (entity.BusinessFAQ, error) {
	var createdFAQ entity.BusinessFAQ
	var bp entity.BusinessProfile
	accountId, _ := strconv.Atoi(p.AccountID)

	bp.AccountID = uint(accountId)

	err := pu.bpRepo.GetByIdPreload(&bp)
	if err != nil {
		return entity.BusinessFAQ{}, err
	}

	createdFAQ.BusinessProfileID = bp.ID
	createdFAQ.Answer = p.Answer
	createdFAQ.Question = p.Question

	if err := pu.repo.Create(&createdFAQ); err != nil {
		return createdFAQ, err
	}

	return createdFAQ, nil

}

func NewFAQUseCase(repo repository.FAQRepositoryInterface, bpRepo repository.BusinessProfileRepositoryInterface) FAQUseCaseInterface {
	return &faqUseCase{
		repo:   repo,
		bpRepo: bpRepo,
	}
}
