package usecase

import (
	"fmt"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
	"mime/multipart"
	"strconv"

	"github.com/google/uuid"
	"gorm.io/gorm"
)

type BusinessProfileUseCaseInterface interface {
	CreateBusinessProfile(bp *dto.BusinessProfileRequest) (entity.BusinessProfile, error)
	CreateProfileImage(file multipart.File, fileExt string) (string, error)
	GetBusinessProfile(bp *dto.BusinessProfileRequest) (dto.BusinessProfileResponse, error)
}

type businessProfileUseCase struct {
	repo             repository.BusinessProfileRepositoryInterface
	accountRepo      repository.AccountRepository
	businessHourRepo repository.BusinessHourRepositoryInterface
	businessLinkRepo repository.BusinessLinkRepositoryInterface
	categoryRepo     repository.CategoryRepository
	fileRepo         repository.FileRepository
}

func (b *businessProfileUseCase) GetBusinessProfile(bp *dto.BusinessProfileRequest) (dto.BusinessProfileResponse, error) {
	accountId, _ := strconv.Atoi(bp.AccountID)

	var createdBp entity.BusinessProfile
	var response dto.BusinessProfileResponse
	var account = entity.Account{Model: gorm.Model{ID: uint(accountId)}}

	err := b.accountRepo.FindById(&account)
	if err != nil {
		return dto.BusinessProfileResponse{}, err
	}

	createdBp.AccountID = uint(accountId)
	err = b.repo.GetByIdPreload(&createdBp)
	if err != nil {
		return dto.BusinessProfileResponse{}, err
	}

	categoryName, err := b.categoryRepo.FindById(createdBp.CategoryID)
	if err != nil {
		return dto.BusinessProfileResponse{}, err
	}

	response.BusinessProfile = createdBp
	response.PhoneNumber = account.PhoneNumber
	response.CategoryName = categoryName

	return response, nil
}

func (b *businessProfileUseCase) CreateProfileImage(file multipart.File, fileExt string) (string, error) {
	fileName := fmt.Sprintf("img-bp-%s.%s", uuid.New().String(), fileExt)
	fileLocation, err := b.fileRepo.Save(file, fileName)

	if err != nil {
		return "", err
	}

	return fileLocation, nil
}

func (b *businessProfileUseCase) CreateBusinessProfile(bp *dto.BusinessProfileRequest) (entity.BusinessProfile, error) {
	var createdBusinessProfile entity.BusinessProfile
	var account entity.Account

	accountId, _ := strconv.Atoi(bp.AccountID)
	categoryId, _ := strconv.Atoi(bp.CategoryID)

	createdBusinessProfile.AccountID = uint(accountId)
	createdBusinessProfile.CategoryID = uint(categoryId)
	createdBusinessProfile.Address = bp.Address
	createdBusinessProfile.ProfileImage = bp.ProfileImage
	createdBusinessProfile.ProfileBio = bp.ProfileBio
	createdBusinessProfile.GmapsLink = bp.GmapsLink
	createdBusinessProfile.DisplayName = bp.DisplayName

	for _, businessHour := range bp.BusinessHours {
		convDay, _ := strconv.Atoi(businessHour.Day)
		createdBusinessProfile.BusinessHours = append(createdBusinessProfile.BusinessHours, entity.BusinessHour{
			Day:       convDay,
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

	account.ID = uint(accountId)
	b.accountRepo.FindById(&account)

	if account.Username != "" {
		var businessProfile = entity.BusinessProfile{AccountID: account.ID}
		b.repo.GetByIdPreload(&businessProfile)

		for _, bhour := range businessProfile.BusinessHours {
			if err := b.businessHourRepo.Delete(strconv.FormatUint(uint64(bhour.ID), 10)); err != nil {
				return createdBusinessProfile, err
			}
		}

		for _, blink := range businessProfile.BusinessLinks {
			if err := b.businessLinkRepo.Delete(strconv.FormatUint(uint64(blink.ID), 10)); err != nil {
				return createdBusinessProfile, err
			}
		}

		if err := b.repo.Delete(strconv.FormatUint(uint64(createdBusinessProfile.AccountID), 10)); err != nil {
			return createdBusinessProfile, err
		}
	}

	if err := b.repo.Create(&createdBusinessProfile); err != nil {
		return createdBusinessProfile, err
	}

	return createdBusinessProfile, nil
}

func NewBusinessProfileUseCase(
	repo repository.BusinessProfileRepositoryInterface,
	accountRepo repository.AccountRepository,
	businessHourRepo repository.BusinessHourRepositoryInterface,
	businessLinkRepo repository.BusinessLinkRepositoryInterface,
	categoryRepo     repository.CategoryRepository,
	fileRepo repository.FileRepository,
) BusinessProfileUseCaseInterface {

	return &businessProfileUseCase{
		repo:             repo,
		accountRepo:      accountRepo,
		businessHourRepo: businessHourRepo,
		businessLinkRepo: businessLinkRepo,
		categoryRepo: categoryRepo,
		fileRepo:         fileRepo,
	}

}
