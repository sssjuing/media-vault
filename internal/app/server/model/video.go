package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Video struct {
	gorm.Model
	SerialNumber string          `gorm:"size:128;uniqueIndex;not null"`
	CoverPath    string          `gorm:"size:1024;not null"` // 封面在对象存储中桶内的路径
	Title        *string         `gorm:"size:512;index"`
	ChineseTitle *string         `gorm:"size:512;index"`
	ActressNames []*ActressName  `gorm:"many2many:actress_name_video"` // 通过演员名字关联
	ReleaseDate  *time.Time      // 发行日期
	VideoPath    *string         `gorm:"size:1024"` // 视频在对象存储中桶内的路径
	Mosaic       *bool           // 是否打马赛克
	Tags         *datatypes.JSON // 标签，JSON 数组格式存储
	Synopsis     *string         // 概要
	M3u8URL      *string         `gorm:"size:1024"` // M3U8 播放列表 URL
	M3u8Referer  *string         `gorm:"size:1024"` // M3U8 引用来源
}
