package model

import (
	"encoding/json"
	"time"

	"github.com/samber/lo"
	"gorm.io/datatypes"
	"gorm.io/gorm"

	"github.com/sssjuing/media-vault/internal/app/server/types"
	"github.com/sssjuing/media-vault/internal/pkg/config"
)

type Video struct {
	gorm.Model
	SerialNumber string          `gorm:"type:varchar(128);uniqueIndex;not null"`
	CoverPath    string          `gorm:"size:1024;not null" ` // 封面在对象存储中桶内的路径
	Title        *string         `gorm:"size:512;index"`
	ChineseTitle *string         `gorm:"size:512;index"`
	Actresses    []*Actress      `gorm:"many2many:actress_video"`
	ReleaseDate  *time.Time      // 发行日期
	VideoPath    *string         `gorm:"size:1024"` // 视频在对象存储中桶内的路径
	Mosaic       *bool           // 是否打马赛克
	Tags         *datatypes.JSON // 标签，JSON 数组格式存储
	M3U8URL      *string         `gorm:"size:1024;column:m3u8_url"` // M3U8 播放列表 URL
	Synopsis     *string         // 概要
}

var publicUrl string

func init() {
	publicUrl = config.GetMinioPublicUrl()
}

func (v *Video) DTO() *types.VideoDTO {
	if v == nil {
		return nil
	}
	var videoUrl *string
	if v.VideoPath != nil && *v.VideoPath != "" {
		videoUrl = new(string)
		*videoUrl = publicUrl + *v.VideoPath
	}
	return &types.VideoDTO{
		ID:           v.ID,
		SerialNumber: v.SerialNumber,
		CoverUrl:     publicUrl + v.CoverPath,
		Title:        v.Title,
		ChineseTitle: v.ChineseTitle,
		Actresses: lo.Map(v.Actresses, func(a *Actress, _ int) *types.ActressDTO {
			return a.DTO()
		}),
		ReleaseDate: v.ReleaseDate,
		VideoUrl:    videoUrl,
		Mosaic:      v.Mosaic,
		Tags:        (*json.RawMessage)(v.Tags),
		M3U8Url:     v.M3U8URL,
		Synopsis:    v.Synopsis,
		CreatedAt:   v.CreatedAt,
		UpdatedAt:   v.UpdatedAt,
	}
}
