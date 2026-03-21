package handler

import (
	"github.com/sssjuing/media-vault/internal/app/server/repository"
	"gorm.io/gorm"
)

type Handler struct {
	db          *gorm.DB
	actressRepo repository.ActressRepository
	videoRepo   repository.VideoRepository
}

func NewHandler(db *gorm.DB, ar repository.ActressRepository, vr repository.VideoRepository) *Handler {
	return &Handler{
		db:          db,
		actressRepo: ar,
		videoRepo:   vr,
	}
}
