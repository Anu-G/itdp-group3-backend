package usecase

import (
	"fmt"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
	"mime/multipart"
	"strconv"
	"time"
)

type BusinessProfileUseCaseInterface interface {
	CreateBusinessProfile(bp *dto.BusinessProfileRequest) (entity.BusinessProfile, error)
	CreateProfileImage(account_id string , file multipart.File, fileExt string) (string, error)
}

type businessProfileUseCase struct {
	repo         repository.BusinessProfileRepositoryInterface
	fileRepo     repository.FileRepository
}

func (b *businessProfileUseCase) CreateProfileImage(account_id string, file multipart.File, fileExt string) (string, error)  {
	fileName := fmt.Sprintf("img-%s.%s", account_id, fileExt)
	fileLocation, err := b.fileRepo.Save(file, fileName)

	if err != nil {
		return "", err
	}

	return fileLocation, nil
}

func (b *businessProfileUseCase) CreateBusinessProfile(bp *dto.BusinessProfileRequest) (entity.BusinessProfile, error) {
	var createdBusinessProfile entity.BusinessProfile
	accountId, _ := strconv.Atoi(bp.AccountID)
	categoryId, _ := strconv.Atoi(bp.CategoryID)

	createdBusinessProfile.AccountID = uint(accountId)
	createdBusinessProfile.CategoryID = uint(categoryId)
	createdBusinessProfile.Address = bp.Address
	createdBusinessProfile.ProfileImage = bp.ProfileImage
	createdBusinessProfile.ProfileBio = bp.ProfileBio
	createdBusinessProfile.GmapsLink = bp.GmapsLink

	for _, businessHour := range bp.BusinessHours {
		layoutFormat := "15:04"
		open, _ := time.Parse(layoutFormat, businessHour.OpenHour)
		close, _ := time.Parse(layoutFormat, businessHour.CloseHour)
		convDay, _ := strconv.Atoi(businessHour.Day)

		createdBusinessProfile.BusinessHours = append(createdBusinessProfile.BusinessHours, entity.BusinessHour{
			Day:       convDay,
			OpenHour:  open,
			CloseHour: close,
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
