package repository

import (
	"github.com/sssjuing/media-vault/internal/app/server/model"
	"gorm.io/gorm"
)

type VideoTagRepository interface {
	FindAll() ([]model.VideoTag, error)
	// Create(*model.VideoTag) error
	// Delete(*model.VideoTag) error
}

type VideoTagRepositoryImpl struct {
	db *gorm.DB
}

func NewVideoTagRepositoryImpl(db *gorm.DB) *VideoTagRepositoryImpl {
	return &VideoTagRepositoryImpl{db: db}
}

func (r *VideoTagRepositoryImpl) FindAll() ([]model.VideoTag, error) {
	var videoTags []model.VideoTag

	if err := r.db.Order("rank").Find(&videoTags).Error; err != nil {
		return nil, err
	}
	return videoTags, nil
}
