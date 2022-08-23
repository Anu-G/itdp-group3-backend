package usecase

import (
	"fmt"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
	"mime/multipart"
	"strconv"

	"github.com/google/uuid"
)

type NonBusinessProfileUseCaseInterface interface {
	CreateNonBusinessProfile(bp *dto.NonBusinessProfileRequest) (entity.NonBusinessProfile, error)
	CreateProfileImage(file multipart.File, fileExt string) (string, error)
	GetNonBusinessProfile(bp *dto.NonBusinessProfileRequest) (dto.NonBusinessProfileResponse, error)
}

type nonBusinessProfileUseCase struct {
	repo        repository.NonBusinessProfileRepositoryInterface
	accountRepo repository.AccountRepository
	fileRepo    repository.FileRepository
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

func (b *nonBusinessProfileUseCase) CreateProfileImage(file multipart.File, fileExt string) (string, error) {
	fileName := fmt.Sprintf("img-nbp-%s.%s", uuid.New().String(), fileExt)
	fileLocation, err := b.fileRepo.Save(file, fileName)

	if err != nil {
		return "", err
	}

	return fileLocation, nil
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
		if err := b.repo.Delete(strconv.FormatUint(uint64(createdNonBusinessProfile.AccountID), 10)); err != nil {
			return createdNonBusinessProfile, err
		}
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
