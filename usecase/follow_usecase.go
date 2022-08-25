package usecase

import (
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
)

type FollowUsecase interface {
	Create(fl *dto.FollowRequest) error
	Delete(fl *dto.UnfollowRequest) error
	FindForVerif(fl *dto.FollowRequest) (entity.Follower, error)
}

type followUsecase struct {
	repoFr repository.FollowerRepository
	repoFd repository.FollowedRepository
}

func NewFollowUsecase(repoFr repository.FollowerRepository, repoFd repository.FollowedRepository) FollowUsecase {
	return &followUsecase{
		repoFr: repoFr,
		repoFd: repoFd,
	}
}

func (flc *followUsecase) Create(fl *dto.FollowRequest) error {
	var followerControl entity.Follower
	var followedControl entity.Followed
	followerControl.AccountID = fl.FollowedAccountID
	followerControl.FollowerAccountID = fl.FollowerAccounID
	followedControl.AccountID = fl.FollowerAccounID
	followedControl.FollowedAccountID = fl.FollowedAccountID
	err := flc.repoFd.Create(&followedControl)
	if err != nil {
		return err
	}
	err = flc.repoFr.Create(&followerControl)
	if err != nil {
		return err
	}
	return nil
}

func (flc *followUsecase) Delete(fl *dto.UnfollowRequest) error {
	var followerControl entity.Follower
	var followedControl entity.Followed
	followerControl.AccountID = fl.FollowedAccountID
	followerControl.FollowerAccountID = fl.FollowerAccounID
	followedControl.AccountID = fl.FollowerAccounID
	followedControl.FollowedAccountID = fl.FollowedAccountID
	err := flc.repoFd.Delete(&followedControl)
	if err != nil {
		return err
	}
	err = flc.repoFr.Delete(&followerControl)
	if err != nil {
		return err
	}
	return nil
}

func (flc *followUsecase) FindForVerif(fl *dto.FollowRequest) (entity.Follower, error) {
	return flc.repoFr.FindForVerif(fl)
}
