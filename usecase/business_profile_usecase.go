package usecase

import (
	"errors"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
	"gorm.io/gorm"
)

type BusinessProfileUseCaseInterface interface {
	CreateBusinessProfile(bp *dto.BusinessProfileRequest) (entity.BusinessProfile, error)
	CreateProfileImage(file multipart.File, ctx *gin.Context, folderName string) (string, error)
	GetBusinessProfile(bp *dto.BusinessProfileRequest) (dto.BusinessProfileResponse, error)
	Update(bp *dto.BusinessProfileRequest) (entity.BusinessProfile, error)
}

type businessProfileUseCase struct {
	repo             repository.BusinessProfileRepositoryInterface
	accountRepo      repository.AccountRepository
	businessHourRepo repository.BusinessHourRepositoryInterface
	businessLinkRepo repository.BusinessLinkRepositoryInterface
	categoryRepo     repository.CategoryRepository
	fileRepo         repository.FileRepository
}

func (b *businessProfileUseCase) Update(bp *dto.BusinessProfileRequest) (entity.BusinessProfile, error) {
	var newBusinessHours []entity.BusinessHour
	var newBusinessLinks []entity.BusinessLink
	var with map[string]interface{}
	var oldBusinessProfile entity.BusinessProfile

	accountId, _ := strconv.Atoi(bp.AccountID)
	categoryId, _ := strconv.Atoi(bp.CategoryID)

	// cek apakah sudah terinsert atau belum di tabel m_business_profile
	oldBusinessProfile.AccountID = uint(accountId)
	b.repo.GetByIdPreload(&oldBusinessProfile)

	if oldBusinessProfile.DisplayName == "" {
		return entity.BusinessProfile{}, errors.New("please initialize profile first")
	}

	// arrange new data
	for _, businessHour := range bp.BusinessHours {
		convDay, _ := strconv.Atoi(businessHour.Day)
		newBusinessHours = append(newBusinessHours, entity.BusinessHour{
			Day:               convDay,
			OpenHour:          businessHour.OpenHour,
			CloseHour:         businessHour.CloseHour,
			BusinessProfileID: oldBusinessProfile.ID,
		})
	}
	for _, businessLink := range bp.BusinessLinks {
		newBusinessLinks = append(newBusinessLinks, entity.BusinessLink{
			Label:             businessLink.Label,
			Link:              businessLink.Link,
			BusinessProfileID: oldBusinessProfile.ID,
		})
	}

	// delete old embedded data
	for _, bhour := range oldBusinessProfile.BusinessHours {
		if err := b.businessHourRepo.Delete(strconv.FormatUint(uint64(bhour.ID), 10)); err != nil {
			return entity.BusinessProfile{}, err
		}
	}
	for _, blink := range oldBusinessProfile.BusinessLinks {
		if err := b.businessLinkRepo.Delete(strconv.FormatUint(uint64(blink.ID), 10)); err != nil {
			return entity.BusinessProfile{}, err
		}
	}

	// arrange update
	var update entity.BusinessProfile
	update.ID = oldBusinessProfile.ID

	with = map[string]interface{}{
		"category_id":   uint(categoryId),
		"address":       bp.Address,
		"profile_image": bp.ProfileImage,
		"profile_bio":   bp.ProfileBio,
		"gmaps_link":    bp.GmapsLink,
		"display_name":  bp.DisplayName,
	}

	if err := b.repo.Update(&update, with); err != nil {
		return entity.BusinessProfile{}, err
	}

	// insert new data
	for _, insertBhour := range newBusinessHours {
		if err := b.businessHourRepo.Create(&insertBhour); err != nil {
			return entity.BusinessProfile{}, err
		}
	}

	for _, insertBlink := range newBusinessLinks {
		if err := b.businessLinkRepo.Create(&insertBlink); err != nil {
			return entity.BusinessProfile{}, err
		}
	}

	return entity.BusinessProfile{}, nil
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

func (b *businessProfileUseCase) CreateProfileImage(file multipart.File, ctx *gin.Context, folderName string) (string, error) {
	return b.fileRepo.SaveSingleFile(file, ctx, folderName)
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

	if account.Username == "" {
		return createdBusinessProfile, errors.New("user not found")
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
	categoryRepo repository.CategoryRepository,
	fileRepo repository.FileRepository,
) BusinessProfileUseCaseInterface {

	return &businessProfileUseCase{
		repo:             repo,
		accountRepo:      accountRepo,
		businessHourRepo: businessHourRepo,
		businessLinkRepo: businessLinkRepo,
		categoryRepo:     categoryRepo,
		fileRepo:         fileRepo,
	}

}
