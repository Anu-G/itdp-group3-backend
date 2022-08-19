package usecase

import (
	"fmt"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
	"mime/multipart"
)

type BusinessProfileUseCaseInterface interface {
	CreateBusinessProfile(bp *dto.BusinessProfileRequest) (entity.BusinessProfile, error)
	CreateProfileImage(bp *entity.BusinessProfile , file multipart.File, fileExt string) (entity.BusinessProfile, error)
}

type businessProfileUseCase struct {
	repo         repository.BusinessProfileRepositoryInterface
	fileRepo     repository.FileRepository
}

func (b *businessProfileUseCase) CreateProfileImage(bp *entity.BusinessProfile , file multipart.File, fileExt string) (entity.BusinessProfile, error)  {
	var createdBusinessProfile entity.BusinessProfile

	fileName := fmt.Sprintf("img-%d.%s", bp.AccountID, fileExt)
	fileLocation, err := b.fileRepo.Save(file, fileName)

	if err != nil {
		return entity.BusinessProfile{}, err
	}

	createdBusinessProfile.ProfileImage = fileLocation

	if err := b.repo.Create(&createdBusinessProfile); err != nil {
		return createdBusinessProfile, err
	}

	return entity.BusinessProfile{}, nil
}

func (b *businessProfileUseCase) CreateBusinessProfile(bp *dto.BusinessProfileRequest) (entity.BusinessProfile, error) {
	var createdBusinessProfile entity.BusinessProfile

	createdBusinessProfile.AccountID = bp.AccountID
	createdBusinessProfile.CategoryID = bp.CategoryID
	createdBusinessProfile.Address = bp.Address
	createdBusinessProfile.ProfileBio = bp.ProfileBio
	createdBusinessProfile.GmapsLink = bp.GmapsLink

	for _, businessHour := range bp.BusinessHours {
		createdBusinessProfile.BusinessHours = append(createdBusinessProfile.BusinessHours, entity.BusinessHour{
			Day:       businessHour.Day,
			OpenHour:  businessHour.OpenHour,
			CloseHour: businessHour.CloseHour,
		})
	}

	for _, businessLink := range bp.BusinessLinks {
		createdBusinessProfile.BusinessLinks = append(createdBusinessProfile.BusinessLinks, entity.BusinessLink{
			Label: businessLink.Label,
			Link:  businessLink.Link,
		})
	}

	if err := b.repo.Create(&createdBusinessProfile); err != nil {
		return createdBusinessProfile, err
	}

	return createdBusinessProfile, nil
}

func NewBusinessProfileUseCase(repo repository.BusinessProfileRepositoryInterface, fileRepo repository.FileRepository) BusinessProfileUseCaseInterface {
	return &businessProfileUseCase{
		repo:         repo,
		fileRepo:     fileRepo,
	}
}
