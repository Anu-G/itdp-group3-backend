package repository

import "gorm.io/gorm"

type BaseRepository interface {
	Paging(db *gorm.DB, page int, pageLim int) *gorm.DB
	Delete(id uint) error
}
