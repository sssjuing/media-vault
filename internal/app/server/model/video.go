//go:generate ../../../../bin/gen-requests -output ../handler/video_requests_gen.go video.go

package model

import (
	"time"

	"gorm.io/datatypes"
)

// Video 视频模型
type Video struct {
	BaseModel
	SerialNumber    string            `gorm:"size:128;uniqueIndex;not null" json:"serial_number"`        // 视频序列号，唯一标识
	CoverPath       string            `gorm:"size:1024;not null" json:"cover_path"`                      // 封面在对象存储中桶内的路径
	Title           *string           `gorm:"size:512;index" json:"title"`                               // 视频原标题
	ChineseTitle    *string           `gorm:"size:512;index" json:"chinese_title"`                       // 视频中文标题
	ActressNameSets []*ActressNameSet `gorm:"many2many:video_actress_name_set" json:"actress_name_sets"` // 视频关联的演员名字集（同一个演员在不同的作品中可能使用不同的名字）
	ReleaseDate     *time.Time        `json:"release_date"`                                              // 发行日期
	VideoPath       *string           `gorm:"size:1024" json:"video_path"`                               // 视频在对象存储中桶内的路径
	Mosaic          *bool             `json:"mosaic"`                                                    // 是否打马赛克
	Tags            *datatypes.JSON   `json:"tags"`                                                      // 标签，JSON 数组格式存储
	Synopsis        *string           `json:"synopsis"`                                                  // 概要
	M3u8URL         *string           `gorm:"size:1024" json:"m3u8_url"`                                 // M3U8 播放列表 URL
	M3u8Referer     *string           `gorm:"size:1024" json:"m3u8_referer"`                             // M3U8 引用来源
}
