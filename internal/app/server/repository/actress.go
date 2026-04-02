package repository

import (
	"errors"

	"github.com/sssjuing/media-vault/internal/app/server/model"
	"gorm.io/gorm"
	"gorm.io/gorm/clause"
)

type ActressRepository interface {
	FindAll() ([]model.Actress, error)
	FindByID(uint) (*model.Actress, error)
	Create(*model.Actress) error
	Update(*model.Actress) error
	Delete(*model.Actress) error
	FindVideos(uint) ([]*model.Video, error)
}

type ActressRepositoryImpl struct {
	db *gorm.DB
}

func NewActressRepositoryImpl(db *gorm.DB) *ActressRepositoryImpl {
	return &ActressRepositoryImpl{db: db}
}

func (r *ActressRepositoryImpl) FindAll() ([]model.Actress, error) {
	var actresses []model.Actress
	if err := r.db.Preload("NameSets").Preload("NameSets.Names").Order("created_at desc").Find(&actresses).Error; err != nil {
		return nil, err
	}
	return actresses, nil
}

func (r *ActressRepositoryImpl) FindByID(id uint) (*model.Actress, error) {
	var a model.Actress
	if err := r.db.Preload("NameSets").Preload("NameSets.Names").First(&a, id).Error; err != nil {
		if errors.Is(err, gorm.ErrRecordNotFound) {
			return nil, nil
		}
		return nil, err
	}
	return &a, nil
}

func (r *ActressRepositoryImpl) Create(a *model.Actress) error {
	return r.db.Create(a).Error
}

func (r *ActressRepositoryImpl) Update(a *model.Actress) error {
	return r.db.Model(a).Updates(a).Error
}

func (r *ActressRepositoryImpl) Delete(a *model.Actress) error {
	return r.db.Select(clause.Associations).Unscoped().Delete(a).Error
}

func (r *ActressRepositoryImpl) FindVideos(id uint) ([]*model.Video, error) {
	var videos []*model.Video
	if err := r.db.
		Preload("ActressNameSets").
		Preload("ActressNameSets.Actress").
		Joins("JOIN video_actress_name_set ON videos.id = video_actress_name_set.video_id").
		Joins("JOIN actress_name_sets ON video_actress_name_set.actress_name_set_id = actress_name_sets.id").
		Where("actress_name_sets.actress_id = ?", id).
		Order("videos.release_date desc").
		Find(&videos).Error; err != nil {
		return nil, err
	}
	return videos, nil
}
