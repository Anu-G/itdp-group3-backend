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
	ReadDetailByID(id uint, page int, pageLim int) (dto.FeedDetailRequest, error)
	ReadForTimeline(page int, pageLim int) ([]dto.FeedDetailRequest, error)
	ReadByAccountID(id int) ([]dto.FeedDetailRequest, error)
	ReadByProfileCategory(cat uint, page int, pageLim int) ([]dto.FeedDetailRequest, error)
	ReadByPage(page int, pageLim int) ([]entity.Feed, error)
	ReadByFollowerAccountID(ids []uint, page int, pageLim int) ([]dto.FeedDetailRequest, error)
	Update(f *entity.Feed) error
	ReadForDetailTimeline(page int, pageLim int, feedId uint) ([]dto.FeedDetailRequest, error)
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
	return fr.db.Preload("DetailComments").Preload("DetailLikes").Find(&f).Error
}

func (fr *feedRepository) ReadByID(f *entity.Feed) error {
	return fr.db.Preload("DetailComments").Preload("DetailLikes").Find(&f, "id = ?", f.ID).Error
}

func (fr *feedRepository) ReadDetailByID(id uint, page int, pageLim int) (dto.FeedDetailRequest, error) {
	var feed *entity.Feed
	var feedCL *dto.FeedResponse
	var feedRequest *dto.FeedDetailRequest
	var err error
	selectQuery := fmt.Sprintln(`
	m_feed.id as post_id, m_feed.account_id, BP.profile_image AS profile_image,NBP.profile_image AS profile_image,BP.display_name AS display_name,NBP.display_name AS display_name, 
	m_feed.caption_post as caption_post, m_feed.created_at as created_at, m_feed.detail_media_feeds as detail_media_feeds`)
	joinQuery := fmt.Sprintln(`
	JOIN m_account as A on A.id = m_feed.account_id 
	LEFT OUTER JOIN m_business_profile as BP on BP.account_id = m_feed.account_id
	LEFT OUTER JOIN m_non_business_profile as NBP on NBP.account_id = m_feed.account_id`)
	read := fr.db.Model(&feed).Where("m_feed.id = ?", id).Select(selectQuery).Joins(joinQuery)
	readCL := fr.db.Model(&entity.Feed{}).Where("m_feed.id = ?", id).Preload("DetailComments", func(db *gorm.DB) *gorm.DB {
		return fr.db.Model(&entity.DetailComment{}).Select(`
m_detail_comment.feed_id, m_detail_comment.account_id, m_detail_comment.comment_fill,  
CASE WHEN mbp.account_id IS NOT NULL 
THEN mbp.display_name ELSE mnbp.display_name
END AS "display_name",
CASE WHEN mbp.account_id IS NOT NULL 
THEN mbp.profile_image ELSE mnbp.profile_image
END AS "profile_image"
`).Joins(`
LEFT OUTER JOIN m_business_profile mbp ON mbp.account_id = m_detail_comment.account_id 
LEFT OUTER JOIN m_non_business_profile mnbp ON mnbp.account_id = m_detail_comment.account_id `)
	}).Preload("DetailLikes")
	res := fr.Paging(read, page, pageLim).Order("m_feed.created_at DESC").Find(&feedRequest)
	resCL := fr.Paging(readCL, page, pageLim).Order("m_feed.created_at DESC").Find(&feedCL)
	feedRequest.DetailComment = feedCL.DetailComments
	feedRequest.DetailLike = feedCL.DetailLikes
	if res.Error == nil {
		err = resCL.Error
	} else {
		err = res.Error
	}
	return *feedRequest, err
}

func (fr *feedRepository) ReadForTimeline(page int, pageLim int) ([]dto.FeedDetailRequest, error) {
	var feed *entity.Feed
	var feedCL *[]dto.FeedResponse
	var feedRequest *[]dto.FeedDetailRequest
	var err error
	selectQuery := fmt.Sprintln(`
	m_feed.id as post_id, 
	m_feed.account_id, 
	CASE 
		WHEN BP.account_id IS NOT NULL THEN BP.profile_image 
		ELSE NBP.profile_image
	END AS "profile_image", 
	CASE 
		WHEN BP.account_id IS NOT NULL THEN BP.display_name 
		ELSE NBP.display_name
	END AS "display_name",
		m_feed.caption_post as caption_post, m_feed.created_at as created_at, m_feed.detail_media_feeds as detail_media_feeds
	`)
	joinQuery := fmt.Sprintln(`
	JOIN m_account as A on A.id = m_feed.account_id 
	LEFT OUTER JOIN m_business_profile as BP on BP.account_id = m_feed.account_id
	LEFT OUTER JOIN m_non_business_profile as NBP on NBP.account_id = m_feed.account_id`)
	read := fr.db.Model(&feed).Select(selectQuery).Joins(joinQuery)
	readCL := fr.db.Model(&entity.Feed{}).Preload("DetailComments", func(db *gorm.DB) *gorm.DB {
		return fr.db.Model(&entity.DetailComment{}).Select(`
m_detail_comment.feed_id, m_detail_comment.account_id, m_detail_comment.comment_fill,  
CASE WHEN mbp.account_id IS NOT NULL 
THEN mbp.display_name ELSE mnbp.display_name
END AS "display_name",
CASE WHEN mbp.account_id IS NOT NULL 
THEN mbp.profile_image ELSE mnbp.profile_image
END AS "profile_image"
`).Joins(`
LEFT OUTER JOIN m_business_profile mbp ON mbp.account_id = m_detail_comment.account_id 
LEFT OUTER JOIN m_non_business_profile mnbp ON mnbp.account_id = m_detail_comment.account_id `)
	}).Preload("DetailLikes")
	res := fr.Paging(read, page, pageLim).Order("m_feed.created_at DESC").Find(&feedRequest)
	resCL := fr.Paging(readCL, page, pageLim).Order("m_feed.created_at DESC").Find(&feedCL)
	for i, feed := range *feedCL {
		(*feedRequest)[i].DetailComment = feed.DetailComments
		(*feedRequest)[i].DetailLike = feed.DetailLikes
	}
	if res.Error == nil {
		err = resCL.Error
	} else {
		err = res.Error
	}
	return *feedRequest, err
}

func (fr *feedRepository) ReadByAccountID(id int) ([]dto.FeedDetailRequest, error) {
	var feed *entity.Feed
	var feedCL *[]dto.FeedResponse
	var feedRequest *[]dto.FeedDetailRequest
	var err error
	selectQuery := fmt.Sprintln(`
	m_feed.id as post_id, m_feed.account_id, BP.profile_image AS profile_image,NBP.profile_image AS profile_image,BP.display_name AS display_name,NBP.display_name AS display_name, 
	m_feed.caption_post as caption_post, m_feed.created_at as created_at, m_feed.detail_media_feeds as detail_media_feeds`)
	joinQuery := fmt.Sprintln(`
	JOIN m_account as A on A.id = m_feed.account_id 
	LEFT OUTER JOIN m_business_profile as BP on BP.account_id = m_feed.account_id
	LEFT OUTER JOIN m_non_business_profile as NBP on NBP.account_id = m_feed.account_id`)
	res := fr.db.Model(&feed).Where("m_feed.account_id = ?", id).Select(selectQuery).Joins(joinQuery).Order("m_feed.created_at DESC").Find(&feedRequest)
	resCL := fr.db.Model(&entity.Feed{}).Where("account_id = ?", id).Preload("DetailComments", func(db *gorm.DB) *gorm.DB {
		return fr.db.Model(&entity.DetailComment{}).Select(`
m_detail_comment.feed_id, m_detail_comment.account_id, m_detail_comment.comment_fill,  
CASE WHEN mbp.account_id IS NOT NULL 
THEN mbp.display_name ELSE mnbp.display_name
END AS "display_name",
CASE WHEN mbp.account_id IS NOT NULL 
THEN mbp.profile_image ELSE mnbp.profile_image
END AS "profile_image"
`).Joins(`
LEFT OUTER JOIN m_business_profile mbp ON mbp.account_id = m_detail_comment.account_id 
LEFT OUTER JOIN m_non_business_profile mnbp ON mnbp.account_id = m_detail_comment.account_id `)
	}).Preload("DetailLikes").Order("m_feed.created_at DESC").Find(&feedCL)
	fmt.Println(len(*feedRequest))
	fmt.Println("comment\n", len(*feedCL))
	for i, feed := range *feedCL {
		(*feedRequest)[i].DetailComment = feed.DetailComments
		(*feedRequest)[i].DetailLike = feed.DetailLikes
	}
	if res.Error == nil {
		err = resCL.Error
	} else {
		err = res.Error
	}
	return *feedRequest, err
}

func (fr *feedRepository) ReadByFollowerAccountID(ids []uint, page int, pageLim int) ([]dto.FeedDetailRequest, error) {
	var f entity.Feed
	var feedCL []dto.FeedResponse
	var feedRes []dto.FeedDetailRequest
	var err error
	selectQuery := fmt.Sprintln(`
	m_feed.id as post_id, BP.profile_image as profile_image,BP.display_name as display_name, m_feed.caption_post as caption_post, m_feed.created_at as created_at, m_feed.detail_media_feeds as detail_media_feeds`)
	joinQuery := fmt.Sprintln(`
	JOIN m_account as A on A.id = m_feed.account_id 
	LEFT OUTER JOIN m_business_profile as BP on BP.account_id = m_feed.account_id
	LEFT OUTER JOIN m_non_business_profile as NBP on NBP.account_id = m_feed.account_id`)
	if len(ids) == 0 {
		return nil, nil
	}
	read := fr.db.Model(&f).Select(selectQuery).Joins(joinQuery).Where("m_feed.account_id = ?", ids[0])
	readCL := fr.db.Model(&entity.Feed{}).Preload("DetailComments", func(db *gorm.DB) *gorm.DB {
		return fr.db.Model(&entity.DetailComment{}).Select(`
m_detail_comment.feed_id, m_detail_comment.account_id, m_detail_comment.comment_fill,  
CASE WHEN mbp.account_id IS NOT NULL 
THEN mbp.display_name ELSE mnbp.display_name
END AS "display_name",
CASE WHEN mbp.account_id IS NOT NULL 
THEN mbp.profile_image ELSE mnbp.profile_image
END AS "profile_image"
`).Joins(`
LEFT OUTER JOIN m_business_profile mbp ON mbp.account_id = m_detail_comment.account_id 
LEFT OUTER JOIN m_non_business_profile mnbp ON mnbp.account_id = m_detail_comment.account_id `)
	}).Preload("DetailLikes")
	for i := 1; i < len(ids); i++ {
		read = read.Or("account_id = ?", ids[i])
	}
	res := fr.Paging(read, page, pageLim).Order("created_at DESC").Find(&feedRes)
	resCL := fr.Paging(readCL, page, pageLim).Order("m_feed.created_at DESC").Find(&feedCL)
	for i, feed := range feedCL {
		feedRes[i].DetailComment = feed.DetailComments
		feedRes[i].DetailLike = feed.DetailLikes
	}
	if res.Error == nil {
		err = resCL.Error
	} else {
		err = res.Error
	}
	return feedRes, err
}

func (fr *feedRepository) ReadByProfileCategory(cat uint, page int, pageLim int) ([]dto.FeedDetailRequest, error) {
	var f entity.Feed
	var feedCL *[]dto.FeedResponse
	var feedRes []dto.FeedDetailRequest
	var err error
	selectQuery := fmt.Sprintln(`
	m_feed.id as post_id, m_feed.account_id, BP.profile_image AS profile_image,NBP.profile_image AS profile_image,BP.display_name AS display_name,NBP.display_name AS display_name, 
	m_feed.caption_post as caption_post, m_feed.created_at as created_at, m_feed.detail_media_feeds as detail_media_feeds`)
	joinQuery := fmt.Sprintln(`
	JOIN m_account as A on A.id = m_feed.account_id 
	LEFT OUTER JOIN m_business_profile as BP on BP.account_id = m_feed.account_id
	LEFT OUTER JOIN m_non_business_profile as NBP on NBP.account_id = m_feed.account_id`)
	read := fr.db.Model(&f).Where("bp.category_id = ?", cat).Select(selectQuery).Joins(joinQuery)
	readCL := fr.db.Model(&entity.Feed{}).Preload("DetailComments", func(db *gorm.DB) *gorm.DB {
		return fr.db.Model(&entity.DetailComment{}).Select(`
m_detail_comment.feed_id, m_detail_comment.account_id, m_detail_comment.comment_fill,  
CASE WHEN mbp.account_id IS NOT NULL 
THEN mbp.display_name ELSE mnbp.display_name
END AS "display_name",
CASE WHEN mbp.account_id IS NOT NULL 
THEN mbp.profile_image ELSE mnbp.profile_image
END AS "profile_image"
`).Joins(`
LEFT OUTER JOIN m_business_profile mbp ON mbp.account_id = m_detail_comment.account_id 
LEFT OUTER JOIN m_non_business_profile mnbp ON mnbp.account_id = m_detail_comment.account_id `)
	}).Preload("DetailLikes")
	res := fr.Paging(read, page, pageLim).Order("m_feed.created_at DESC").Find(&feedRes)
	resCL := fr.Paging(readCL, page, pageLim).Order("m_feed.created_at DESC").Find(&feedCL)
	for i, feed := range *feedCL {
		if i < len(feedRes) {
			feedRes[i].DetailComment = feed.DetailComments
			feedRes[i].DetailLike = feed.DetailLikes
		}
	}
	if res.Error == nil {
		err = resCL.Error
	} else {
		err = res.Error
	}
	return feedRes, err
}

func (fr *feedRepository) ReadByPage(page int, pageLim int) ([]entity.Feed, error) {
	var feedRes []entity.Feed
	read := fr.db.Preload("DetailComments").Preload("DetailLikes")
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

func (fr *feedRepository) ReadForDetailTimeline(page int, pageLim int, feedId uint) ([]dto.FeedDetailRequest, error) {
	var feed *entity.Feed
	var feedCL *[]dto.FeedResponse
	var feedRequest *[]dto.FeedDetailRequest
	var err error
	selectQuery := fmt.Sprintln(`
	m_feed.id as post_id, m_feed.account_id, BP.profile_image AS profile_image,NBP.profile_image AS profile_image,BP.display_name AS display_name,NBP.display_name AS display_name, 
	m_feed.caption_post as caption_post, m_feed.created_at as created_at, m_feed.detail_media_feeds as detail_media_feeds`)
	joinQuery := fmt.Sprintln(`
	JOIN m_account as A on A.id = m_feed.account_id 
	LEFT OUTER JOIN m_business_profile as BP on BP.account_id = m_feed.account_id
	LEFT OUTER JOIN m_non_business_profile as NBP on NBP.account_id = m_feed.account_id`)
	read := fr.db.Model(&feed).Select(selectQuery).Joins(joinQuery).Where(`"m_feed"."id" = ?`, feedId)
	readCL := fr.db.Model(&entity.Feed{}).Where(`"m_feed"."id" = ?`, feedId).Preload("DetailComments", func(db *gorm.DB) *gorm.DB {
		return fr.db.Model(&entity.DetailComment{}).Select("m_detail_comment.feed_id, m_detail_comment.account_id, m_detail_comment.comment_fill, bp.display_name, bp.profile_image").Joins("JOIN m_business_profile as bp on bp.account_id = m_detail_comment.account_id")
	}).Preload("DetailLikes")
	res := fr.Paging(read, page, pageLim).Order("m_feed.created_at DESC").Find(&feedRequest)
	resCL := fr.Paging(readCL, page, pageLim).Order("m_feed.created_at DESC").Find(&feedCL)
	for i, feed := range *feedCL {
		(*feedRequest)[i].DetailComment = feed.DetailComments
		(*feedRequest)[i].DetailLike = feed.DetailLikes
	}
	if res.Error == nil {
		err = resCL.Error
	} else {
		err = res.Error
	}
	return *feedRequest, err
}
