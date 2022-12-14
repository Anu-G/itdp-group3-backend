package repository

import (
	"fmt"
	"itdp-group3-backend/model/dto"
	"itdp-group3-backend/model/entity"

	"gorm.io/gorm"
)

type FeedRepository interface {
	Create(f *entity.Feed) error
	Read(f *[]entity.Feed) error
	ReadByID(f *entity.Feed) error
	ReadForTimeline(page int, pageLim int) ([]dto.FeedDetailRequest, error)
	ReadByAccountID(id uint, page int, pageLim int) ([]entity.Feed, error)
	ReadByProfileCategory(cat uint, page int, pageLim int) ([]dto.FeedDetailRequest, error)
	ReadByPage(page int, pageLim int) ([]entity.Feed, error)
	ReadByFollowerAccountID(ids []uint, page int, pageLim int) ([]dto.FeedDetailRequest, error)
	Update(f *entity.Feed) error
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

func (fr *feedRepository) Read(f *[]entity.Feed) error {
	return fr.db.Preload("DetailComments").Find(&f).Error
}

func (fr *feedRepository) ReadByID(f *entity.Feed) error {
	return fr.db.Preload("DetailComments").Find(&f, "id = ?", f.ID).Error
}

func (fr *feedRepository) ReadForTimeline(page int, pageLim int) ([]dto.FeedDetailRequest, error) {
	var feed *entity.Feed
	var feedRequest *[]dto.FeedDetailRequest
	selectQuery := fmt.Sprintln(`
	m_feed.id as post_id, m_feed.account_id, BP.profile_image as profile_image,BP.display_name as display_name, m_feed.caption_post as caption_post, m_feed.created_at as created_at, m_feed.detail_media_feeds as detail_media_feeds`)
	joinQuery := fmt.Sprintln(`
	JOIN m_account as A on A.id = m_feed.account_id 
	JOIN m_business_profile as BP on BP.account_id = m_feed.account_id`)
	read := fr.db.Model(&feed).Select(selectQuery).Joins(joinQuery).Preload("DetailComments")
	res := fr.Paging(read, page, pageLim).Find(&feedRequest).Order("m_feed.created_at")
	return *feedRequest, res.Error
}

func (fr *feedRepository) ReadByAccountID(id uint, page int, pageLim int) ([]entity.Feed, error) {
	var f entity.Feed
	var feedRes []entity.Feed
	read := fr.db.Model(&f).Where("account_id = ?", id).Preload("DetailComments")
	res := fr.Paging(read, page, pageLim).Find(&feedRes).Order("id").Error
	return feedRes, res
}

func (fr *feedRepository) ReadByFollowerAccountID(ids []uint, page int, pageLim int) ([]dto.FeedDetailRequest, error) {
	var f entity.Feed
	var feedRes []dto.FeedDetailRequest
	selectQuery := fmt.Sprintln(`
	m_feed.id as post_id, BP.profile_image as profile_image,BP.display_name as display_name, m_feed.caption_post as caption_post, m_feed.created_at as created_at, m_feed.detail_media_feeds as detail_media_feeds`)
	joinQuery := fmt.Sprintln(`
	JOIN m_account as A on A.id = m_feed.account_id 
	JOIN m_business_profile as BP on BP.account_id = m_feed.account_id`)
	if len(ids) == 0 {
		return nil, nil
	}
	read := fr.db.Model(&f).Select(selectQuery).Joins(joinQuery).Where("m_feed.account_id = ?", ids[0])
	for i := 1; i < len(ids); i++ {
		read = read.Or("account_id = ?", ids[i])
	}
	res := fr.Paging(read, page, pageLim).Find(&feedRes).Order("created_at").Error
	return feedRes, res
}

func (fr *feedRepository) ReadByProfileCategory(cat uint, page int, pageLim int) ([]dto.FeedDetailRequest, error) {
	var f entity.Feed
	var feedRes []dto.FeedDetailRequest
	selectQuery := fmt.Sprintln(`
	m_feed.id as post_id, BP.profile_image as profile_image,BP.display_name as display_name, m_feed.caption_post as caption_post, m_feed.created_at as created_at, m_feed.detail_media_feeds as detail_media_feeds`)
	joinQuery := fmt.Sprintln(`
	JOIN m_account as A on A.id = m_feed.account_id 
	JOIN m_business_profile as BP on BP.account_id = m_feed.account_id`)
	read := fr.db.Model(&f).Where("bp.category_id = ?", cat).Select(selectQuery).Joins(joinQuery)
	res := fr.Paging(read, page, pageLim).Find(&feedRes).Order("id").Error
	return feedRes, res
}

func (fr *feedRepository) ReadByPage(page int, pageLim int) ([]entity.Feed, error) {
	var feedRes []entity.Feed
	read := fr.db.Preload("DetailComments")
	res := fr.Paging(read, page, pageLim).Find(&feedRes).Order("id").Error
	return feedRes, res
}

func (fr *feedRepository) Update(f *entity.Feed) error {
	return fr.db.Save(&f).Error
}

func (fr *feedRepository) Paging(db *gorm.DB, page int, pageLim int) *gorm.DB {
	lim := pageLim
	offset := (page - 1) * lim
	res := db.Offset(offset).Limit(lim)
	return res
}

func (fr *feedRepository) Delete(id uint) error {
	var feedRes entity.Feed
	return fr.db.Where("id = ?", id).Delete(&feedRes).Error
}
