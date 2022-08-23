package usecase

import (
	"fmt"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"itdp-group3-backend/repository"
	"strconv"
	"strings"
)

type ProductUseCaseInterface interface {
	CreateProduct(p *dto.ProductRequest) (entity.Product, error)
	GetByAccount(p dto.ProductRequest) ([]dto.ProductResponse, error)
	GetByProduct(p dto.ProductRequest) (dto.ProductResponse, error)
	Delete(id string) error
}

type productUseCase struct {
	repo repository.ProductRepositoryInterface
}

func (pu *productUseCase) Delete(id string) error {
	return pu.repo.Delete(id)
}

func (pu *productUseCase) GetByAccount(p dto.ProductRequest) ([]dto.ProductResponse, error) {
	var products []dto.ProductResponse

	res, err := pu.repo.GetByAccount(p)
	if err != nil {
		return []dto.ProductResponse{}, err
	}

	for _, product := range res {
		products = append(products, dto.ProductResponse{
			ProductID:           fmt.Sprintf("%d", product.ID),
			AccountID:           fmt.Sprintf("%d", product.AccountID),
			ProductName:         product.ProductName,
			Price:               fmt.Sprintf("%.f", product.Price),
			Description:         product.Description,
			DetailMediaProducts: strings.Split(product.DetailMediaProducts, ", "),
		})
	}
	return products, nil
}

func (pu *productUseCase) GetByProduct(p dto.ProductRequest) (dto.ProductResponse, error) {
	var product dto.ProductResponse

	res, err := pu.repo.GetByProduct(p)
	if err != nil {
		return dto.ProductResponse{}, err
	}

	product.ProductID = fmt.Sprintf("%d", res.ID)
	product.AccountID = fmt.Sprintf("%d", res.AccountID)
	product.ProductName = res.ProductName
	product.Price = fmt.Sprintf("%.f", res.Price)
	product.Description = res.Description
	product.DetailMediaProducts = strings.Split(res.DetailMediaProducts, ", ")

	return product, nil
}

func (pu *productUseCase) CreateProduct(p *dto.ProductRequest) (entity.Product, error) {
	var createdProduct entity.Product
	accountId, _ := strconv.Atoi(p.AccountID)
	priceConv, _ := strconv.Atoi(p.Price)

	createdProduct.AccountID = uint(accountId)
	createdProduct.ProductName = p.ProductName
	createdProduct.Price = float64(priceConv)
	createdProduct.Description = p.Description
	createdProduct.DetailMediaProducts = strings.Join(p.DetailMediaProducts, ", ")

	if err := pu.repo.Create(&createdProduct); err != nil {
		return createdProduct, err
	}

	return createdProduct, nil

}

func NewProductUseCase(repo repository.ProductRepositoryInterface) ProductUseCaseInterface {
	return &productUseCase{
		repo: repo,
	}
}
