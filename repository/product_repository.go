package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	Create(p *entity.Product) error
	GetAllPreload(p *[]entity.Product) error
	GetByIdPreload(p *entity.Product) error
}

type productRepository struct {
	db *gorm.DB
}

func (pr *productRepository) GetByIdPreload(p *entity.Product) error {
	return pr.db.Preload("DetailMediaProducts").First(&p, "m_product.id = ? AND m_product.account_id = ?", p.ID, p.AccountID).Error
}

func (pr *productRepository) GetAllPreload(p *[]entity.Product) error {
	return pr.db.Preload("DetailMediaProducts").Find(&p, "m_product.account_id = ?", ).Error
}

func (pr *productRepository) Create(p *entity.Product) error {
	return pr.db.Create(&p).Error
}

func NewProductRepo(db *gorm.DB) ProductRepositoryInterface {
	return &productRepository{
		db: db,
	}
}
