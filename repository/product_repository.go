package repository

import (
	"errors"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	Create(p *entity.Product) error
	GetByAccount(p dto.ProductRequest) ([]entity.Product, error)
	GetByProduct(p dto.ProductRequest) (entity.Product, error)
	Delete(id string) error
}

type productRepository struct {
	db *gorm.DB
}

func (pr *productRepository) Delete(id string) error {
	return pr.db.Where("id = ?", id).Delete(&entity.Product{}).Error
}

func (pr *productRepository) GetByAccount(p dto.ProductRequest) ([]entity.Product, error) {
	var products []entity.Product
	res := pr.db.Find(&products, "m_product.account_id = ?", p.AccountID)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return products, nil
		} else {
			return products, err
		}
	}
	return products, nil
}

func (pr *productRepository) GetByProduct(p dto.ProductRequest) (entity.Product, error) {
	var product entity.Product
	res := pr.db.Find(&product, "m_product.account_id = ? AND m_product.id = ?", p.AccountID, p.ProductID)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return product, nil
		} else {
			return product, err
		}
	}
	return product, nil
}

func (pr *productRepository) Create(p *entity.Product) error {
	return pr.db.Create(&p).Error
}

func NewProductRepo(db *gorm.DB) ProductRepositoryInterface {
	return &productRepository{
		db: db,
	}
}
