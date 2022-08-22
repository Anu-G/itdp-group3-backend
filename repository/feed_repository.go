package repository

import (
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type FeedRepository interface {
	Create(f *entity.Feed) error
	Read(f *entity.Feed) error
	ReadByAccountID(id uint, page int, pageLim int) ([]entity.Feed, error)
	ReadByProfileCategory(cat uint, page int, pageLim int) ([]entity.Feed, error)
	ReadByPage(page int, pageLim int) ([]entity.Feed, error)
	BaseRepository
}

type feedRepository struct {
	db *gorm.DB
}

func NewFeedRepository(db *gorm.DB) FeedRepository {
	return &feedRepository{
		db: db,
	}
}

func (fr *feedRepository) Create(f *entity.Feed) error {
	return fr.db.Create(&f).Error
}

func (fr *feedRepository) Read(f *entity.Feed) error {
	return fr.db.Preload("DetailMediaFeeds").Preload("DetailComments").Find(&f).Error
}

func (fr *feedRepository) ReadByAccountID(id uint, page int, pageLim int) ([]entity.Feed, error) {
	var f entity.Feed
	var feedRes []entity.Feed
	read := fr.db.Model(&f).Where("account_id = ?", id).Preload("DetailMediaFeeds").Preload("DetailComments").Find(&feedRes)
	res := fr.Paging(read, page, pageLim).Error
	return feedRes, res
}

func (fr *feedRepository) ReadByProfileCategory(cat uint, page int, pageLim int) ([]entity.Feed, error) {
	var f entity.Feed
	var feedRes []entity.Feed
	read := fr.db.Model(&f).Where("bp.category_id = ?", cat).Joins("join m_business_profile as bp on bp.account_id = m_feed.account_id").Scan(&feedRes)
	res := fr.Paging(read, page, pageLim).Error
	return feedRes, res
}

func (fr *feedRepository) ReadByPage(page int, pageLim int) ([]entity.Feed, error) {
	var feedRes []entity.Feed
	read := fr.db.Preload("DetailMediaFeeds").Preload("DetailComments").Find(&feedRes)
	res := fr.Paging(read, page, pageLim).Error
	return feedRes, res
}

func (fr *feedRepository) Paging(db *gorm.DB, page int, pageLim int) *gorm.DB {
	lim := pageLim
	offset := (page - 1) * lim
	res := db.Offset(offset).Limit(lim).Order("id")
	return res
}

func (fr *feedRepository) Delete(id uint) error {
	var feedRes entity.Feed
	return fr.db.Where("id = ?", id).Delete(&feedRes).Error
}
