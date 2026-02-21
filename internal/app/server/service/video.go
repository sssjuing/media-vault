package service

import (
	"github.com/samber/lo"
	"github.com/sssjuing/media-vault/internal/app/server/model"
	"github.com/sssjuing/media-vault/internal/app/server/repository"
	"github.com/sssjuing/media-vault/internal/app/server/types"
	"gorm.io/datatypes"
)

type VideoService interface {
	// GetList(repository.VidoesQueryOptions) ([]dto.VideoDTO, int64, error)
	Create(*types.CreateVideoDTO) (*types.VideoDTO, error)
	Update(*model.Video, *types.UpdateVideoDTO) (*types.VideoDTO, error)
}

type VideoServiceImpl struct {
	videoRepo repository.VideoRepository
}

func NewVideoService(videoRepo repository.VideoRepository) *VideoServiceImpl {
	return &VideoServiceImpl{
		videoRepo: videoRepo,
	}
}

// func (svc *VideoServiceImpl) GetList(options repository.VidoesQueryOptions) ([]dto.VideoDTO, int64, error) {
// 	videos, count, err := svc.videoRepo.FindAll(options)
// 	if err != nil {
// 		return nil, 0, err
// 	}
// 	list := utils.Map(videos, func(v model.Video) dto.VideoDTO {
// 		return *v.DTO()
// 	})
// 	return list, count, nil
// }

func (svc *VideoServiceImpl) Create(cvd *types.CreateVideoDTO) (*types.VideoDTO, error) {
	v := model.Video{
		SerialNumber: cvd.SerialNumber,
		CoverPath:    cvd.CoverPath,
		Title:        cvd.Title,
		ChineseTitle: cvd.ChineseTitle,
		Actresses: lo.Map(cvd.Actresses, func(ad *types.ActressDTO, _ int) *model.Actress {
			actress := &model.Actress{}
			actress.ID = ad.ID
			return actress
		}),
		ReleaseDate: cvd.ReleaseDate,
		VideoPath:   cvd.VideoPath,
		Mosaic:      cvd.Mosaic,
		Tags:        (*datatypes.JSON)(cvd.Tags),
		M3U8URL:     cvd.M3U8Url,
		Synopsis:    cvd.Synopsis,
	}
	if err := svc.videoRepo.Create(&v); err != nil {
		return nil, err
	}
	return v.DTO(), nil
}

func (svc *VideoServiceImpl) Update(v *model.Video, uvd *types.UpdateVideoDTO) (*types.VideoDTO, error) {
	if uvd.SerialNumber != "" {
		v.SerialNumber = uvd.SerialNumber
	}
	if uvd.CoverPath != "" {
		v.CoverPath = uvd.CoverPath
	}
	v.Title = uvd.Title
	v.ChineseTitle = uvd.ChineseTitle
	if uvd.Actresses != nil {
		v.Actresses = lo.Map(uvd.Actresses, func(ad *types.ActressDTO, _ int) *model.Actress {
			actress := &model.Actress{}
			actress.ID = ad.ID
			return actress
		})
	}
	v.ReleaseDate = uvd.ReleaseDate
	v.VideoPath = uvd.VideoPath
	v.Mosaic = uvd.Mosaic
	v.Tags = (*datatypes.JSON)(uvd.Tags)
	v.M3U8URL = uvd.M3U8Url
	v.Synopsis = uvd.Synopsis
	if err := svc.videoRepo.Update(v); err != nil {
		return nil, err
	}
	return v.DTO(), nil
}
