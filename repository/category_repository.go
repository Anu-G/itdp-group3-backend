package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type CategoryRepositoryInterface interface {
	FindById(id uint) (entity.Category, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func (c *categoryRepository) FindById(id uint) (entity.Category, error) {
	var category entity.Category
	err := c.db.First(&category, "category_id=?", id).Error
	if err!=nil {
		return entity.Category{}, err
	}
	return category, nil
}

func NewCategoryRepository(db *gorm.DB) CategoryRepositoryInterface {
	return &categoryRepository{
		db: db,
	}
}
