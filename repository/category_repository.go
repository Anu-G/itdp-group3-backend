package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type CategoryRepository interface {
	Create(cat *entity.Category) error
	ReadAll(cat *[]entity.Category) error
	FindById(id uint) (string, error)
}

type categoryRepository struct {
	db *gorm.DB
}

func NewCategoryRepository(db *gorm.DB) CategoryRepository {
	return &categoryRepository{
		db: db,
	}
}

func (catr *categoryRepository) Create(cat *entity.Category) error {
	return catr.db.Create(&cat).Error
}

func (catr *categoryRepository) ReadAll(cat *[]entity.Category) error {
	return catr.db.Find(&cat).Error
}

func (catr *categoryRepository) Delete(cat *entity.Category) error {
	return catr.db.Where("id = ?", cat).Delete(&cat).Error
}

func (catr *categoryRepository) FindById(id uint) (string, error) {
	var category entity.Category
	err := catr.db.First(&category, "id = ?" , id).Error
	return category.CategoryName, err
}

