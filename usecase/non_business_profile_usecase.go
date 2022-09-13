package usecase

import (
	"errors"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
	"mime/multipart"
	"strconv"

	"github.com/gin-gonic/gin"
)

type NonBusinessProfileUseCaseInterface interface {
	CreateNonBusinessProfile(bp *dto.NonBusinessProfileRequest) (entity.NonBusinessProfile, error)
	CreateProfileImage(file multipart.File, ctx *gin.Context, folderName string) (string, error)
	GetNonBusinessProfile(bp *dto.NonBusinessProfileRequest) (dto.NonBusinessProfileResponse, error)
	UpdateNonBusinessProfile(bp *dto.NonBusinessProfileRequest) (entity.NonBusinessProfile, error)
}

type nonBusinessProfileUseCase struct {
	repo        repository.NonBusinessProfileRepositoryInterface
	accountRepo repository.AccountRepository
	fileRepo    repository.FileRepository
}

func (b *nonBusinessProfileUseCase) UpdateNonBusinessProfile(bp *dto.NonBusinessProfileRequest) (entity.NonBusinessProfile, error) {
	var with map[string]interface{}
	var oldNonBusinessProfile entity.NonBusinessProfile

	accountId, _ := strconv.Atoi(bp.AccountID)

	oldNonBusinessProfile.AccountID = uint(accountId)
	b.repo.GetById(&oldNonBusinessProfile)

	if oldNonBusinessProfile.DisplayName == "" {
		return entity.NonBusinessProfile{}, errors.New("please initialize profile first")
	}

	var update entity.NonBusinessProfile
	update.ID = oldNonBusinessProfile.ID

	with = map[string]interface{}{
		"profile_image": bp.ProfileImage,
		"profile_bio":   bp.ProfileBio,
		"display_name":  bp.DisplayName,
	}

	if err := b.repo.Update(&update, with); err != nil {
		return entity.NonBusinessProfile{}, err
	}

	return entity.NonBusinessProfile{}, nil
}

func (b *nonBusinessProfileUseCase) GetNonBusinessProfile(bp *dto.NonBusinessProfileRequest) (dto.NonBusinessProfileResponse, error) {
	var createdBp entity.NonBusinessProfile
	var response dto.NonBusinessProfileResponse
	var account entity.Account
	accountId, _ := strconv.Atoi(bp.AccountID)

	account.ID = uint(accountId)
	err := b.accountRepo.FindById(&account)
	if err != nil {
		return dto.NonBusinessProfileResponse{}, err
	}

	createdBp.AccountID = uint(accountId)
	err = b.repo.GetById(&createdBp)
	if err != nil {
		return dto.NonBusinessProfileResponse{}, err
	}

	response.NonBusinessProfile = createdBp
	response.PhoneNumber = account.PhoneNumber

	return response, nil
}

func (b *nonBusinessProfileUseCase) CreateProfileImage(file multipart.File, ctx *gin.Context, folderName string) (string, error) {
	return b.fileRepo.SaveSingleFile(file, ctx, folderName)
}

func (b *nonBusinessProfileUseCase) CreateNonBusinessProfile(bp *dto.NonBusinessProfileRequest) (entity.NonBusinessProfile, error) {
	var createdNonBusinessProfile entity.NonBusinessProfile
	var account entity.Account

	accountId, _ := strconv.Atoi(bp.AccountID)

	createdNonBusinessProfile.AccountID = uint(accountId)
	createdNonBusinessProfile.ProfileImage = bp.ProfileImage
	createdNonBusinessProfile.ProfileBio = bp.ProfileBio
	createdNonBusinessProfile.DisplayName = bp.DisplayName

	account.ID = uint(accountId)
	b.accountRepo.FindById(&account)

	if account.Username != "" {
		return createdNonBusinessProfile, errors.New("user not found")
	}

	if err := b.repo.Create(&createdNonBusinessProfile); err != nil {
		return createdNonBusinessProfile, err
	}

	return createdNonBusinessProfile, nil
}

func NewNonBusinessProfileUseCase(repo repository.NonBusinessProfileRepositoryInterface, accountRepo repository.AccountRepository, fileRepo repository.FileRepository) NonBusinessProfileUseCaseInterface {
	return &nonBusinessProfileUseCase{
		repo:        repo,
		accountRepo: accountRepo,
		fileRepo:    fileRepo,
	}
}
