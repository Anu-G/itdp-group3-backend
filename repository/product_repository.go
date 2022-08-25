package repository

import (
	"errors"
	"fmt"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"
	"strings"

	"gorm.io/gorm"
)

type ProductRepositoryInterface interface {
	Create(p *entity.Product) error
	GetByAccount(p dto.ProductRequest) ([]entity.Product, error)
	GetByProduct(p dto.ProductRequest) (entity.Product, error)
	SearchProduct(keyword string) ([]entity.Product, error)
	Delete(id string) error
}

type productRepository struct {
	db *gorm.DB
}

func (pr *productRepository) SearchProduct(keyword string) ([]entity.Product, error) {
	var products []entity.Product
	var keywordLike = strings.Split(keyword, " ")
	var newKeywordLike string
	newKeyword := strings.ReplaceAll(keyword, " ", " | ")

	for i := 0; i < len(keywordLike); i++ {
		var key = keywordLike[i]
		newKeywordLike += fmt.Sprintf(" P.product_name LIKE '%%%s%%'", key)
		if i != len(keywordLike)-1 {
			newKeywordLike += "AND "
		}
	}

	query := fmt.Sprintf(`
SELECT *
FROM m_product as P
WHERE P.product_name @@ to_tsquery('` + newKeyword + `' )
OR regexp_replace(REPLACE((regexp_replace(P.description, '(^|\s)[^#]+(\s|$)', '', 'g')),'#',''), E'[\\n\\r]+', ' ', 'g') @@ to_tsquery('` + newKeyword + `' )
OR %v
ORDER BY ts_rank_cd(
	to_tsvector('indonesian',P.product_name), 
	to_tsquery('` + newKeyword + `')) +
	ts_rank_cd(
	to_tsvector('indonesian',regexp_replace(REPLACE((regexp_replace(P.description, '(^|\s)[^#]+(\s|$)', '', 'g')),'#',''), E'[\\n\\r]+', ' ', 'g')), 
	to_tsquery('` + newKeyword + `'), 32 /* rank/(rank+1) */) DESC;
	`, newKeywordLike)

	res := pr.db.Raw(query).Find(&products)
	if err := res.Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return products, nil
		} else {
			return products, err
		}
	}
	return products, nil
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
