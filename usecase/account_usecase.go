package usecase

import (
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type AccountUsecase interface {
	Update(a *entity.Account) error
	UpdateByID(id uint) (entity.Account, error)
	ReadForPostTimeline(a *[]entity.Account) error
	ReadForProductDetail(a *entity.Account) error
	ReadForFeedDetail(a *entity.Account) error
	FollowList(rf dto.FollowListRequest) ([]dto.FollowListResponse, error)
}

type accountUsecase struct {
	repo repository.AccountRepository
}

func NewAccountUsecse(repo repository.AccountRepository) AccountUsecase {
	return &accountUsecase{
		repo: repo,
	}
}

func (ac *accountUsecase) Update(a *entity.Account) error {
	return ac.repo.Update(a)
}

func (ac *accountUsecase) UpdateByID(id uint) (entity.Account, error) {
	var newAccount entity.Account
	newAccount.ID = id
	err := ac.repo.FindById(&newAccount)
	if err != nil {
		return newAccount, err
	}
	newAccount.RoleID = 2
	return newAccount, ac.repo.Update(&newAccount)
}

func (ac *accountUsecase) ReadForPostTimeline(a *[]entity.Account) error {
	return ac.repo.ReadForPostTimeline(a)
}

func (ac *accountUsecase) ReadForProductDetail(a *entity.Account) error {
	return ac.repo.ReadForProductDetail(a)
}

func (ac *accountUsecase) ReadForFeedDetail(a *entity.Account) error {
	return ac.repo.ReadForFeedDetail(a)
}

func (ac *accountUsecase) FollowList(rf dto.FollowListRequest) ([]dto.FollowListResponse, error) {
	var accountInput entity.Account
	var response []dto.FollowListResponse
	accountInput.ID = rf.AccountID
	err := ac.repo.FindById(&accountInput)
	if err != nil {
		return nil, err
	}
	var accountIDList []uint
	if rf.FollowStatus {
		for _, follow := range accountInput.Follower {
			accountIDList = append(accountIDList, follow.FollowerAccountID)
		}
	} else {
		for _, follow := range accountInput.Followed {
			accountIDList = append(accountIDList, follow.FollowedAccountID)
		}
	}
	accountList, err := ac.repo.FindListById(accountIDList)
	for _, account := range accountList {
		response = append(response, dto.FollowListResponse{
			AccountID: account.ID,
			Username:  account.Username,
		})
	}
	return response, err
}
