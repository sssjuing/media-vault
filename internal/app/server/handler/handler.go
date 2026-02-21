package handler

import (
	"github.com/sssjuing/media-vault/internal/app/server/repository"
	"github.com/sssjuing/media-vault/internal/app/server/service"
)

type Handler struct {
	actressRepo  repository.ActressRepository
	actressSvc   service.ActressService
	videoRepo    repository.VideoRepository
	videoSvc     service.VideoService
	videoTagRepo repository.VideoTagRepository
}

func NewHandler(ar repository.ActressRepository, vr repository.VideoRepository, vtr repository.VideoTagRepository) *Handler {
	as := service.NewActressService(ar)
	vs := service.NewVideoService(vr)
	return &Handler{
		actressRepo:  ar,
		actressSvc:   as,
		videoRepo:    vr,
		videoSvc:     vs,
		videoTagRepo: vtr,
	}
}
