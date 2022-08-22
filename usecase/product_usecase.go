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

type ProductUseCaseInterface interface {
	CreateProduct(p *dto.ProductRequest) (entity.Product,error)
	CreateProductImage(file multipart.File, fileExt string) (string, error)
}

type productUseCase struct {
	repo     repository.ProductRepositoryInterface
	fileRepo repository.FileRepository
}

func (pu *productUseCase) CreateProductImage(file multipart.File, fileExt string) (string, error) {
	fileName := fmt.Sprintf("img-product-%s.%s", uuid.New().String(), fileExt)
	fileLocation, err := pu.fileRepo.Save(file, fileName)

	if err != nil {
		return "", err
	}

	return fileLocation, nil
}

func (pu *productUseCase) CreateProduct(p *dto.ProductRequest) (entity.Product,error) {
	var createdProduct entity.Product
	accountId, _ := strconv.Atoi(p.AccountID)
	priceConv, _ := strconv.Atoi(p.Price)

	createdProduct.AccountID = uint(accountId)
	createdProduct.ProductName = p.ProductName
	createdProduct.Price = float64(priceConv)
	createdProduct.Description = p.Description

	for _, detailMediaProduct := range p.DetailMediaProducts {
		createdProduct.DetailMediaProducts = append(createdProduct.DetailMediaProducts, entity.DetailMediaProduct{
			MediaLink: detailMediaProduct.MediaLink,
		})
	}

	if err := pu.repo.Create(&createdProduct); err != nil {
		return createdProduct, err
	}

	return createdProduct, nil

}

func NewProductUseCase(repo repository.ProductRepositoryInterface, fileRepo repository.FileRepository) ProductUseCaseInterface {
	return &productUseCase{
		repo:     repo,
		fileRepo: fileRepo,
	}
}
