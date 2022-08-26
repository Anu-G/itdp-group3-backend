package repository

import (
	"errors"
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type FAQRepositoryInterface interface {
	Create(p *entity.BusinessFAQ) error
	GetFAQByAccount(id string) ([]entity.BusinessFAQ, error)
	Delete(id string) error
}

type faqRepository struct {
	db *gorm.DB
}

func (pr *faqRepository) Delete(id string) error {
	return pr.db.Where("id = ?", id).Delete(&entity.BusinessFAQ{}).Error
}

func (pr *faqRepository) GetFAQByAccount(id string) ([]entity.BusinessFAQ, error) {
	var faqs []entity.BusinessFAQ
	res := pr.db.Find(&faqs, "m_business_faq.business_profile_id = ?", id)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return faqs, nil
		} else {
			return faqs, err
		}
	}
	return faqs, nil
}

func (pr *faqRepository) Create(p *entity.BusinessFAQ) error {
	return pr.db.Create(&p).Error
}

func NewFAQRepo(db *gorm.DB) FAQRepositoryInterface {
	return &faqRepository{
		db: db,
	}
}
