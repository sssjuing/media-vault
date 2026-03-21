package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type Actress struct {
	gorm.Model
	UniqueName   string          `gorm:"type:varchar(128);uniqueIndex;not null"` // 唯一名称，一般为日文名
	ChineseName  string          `gorm:"type:varchar(128);uniqueIndex;not null"` // 中文名
	EnglishName  *string         `gorm:"type:varchar(128)"`                      // 英文名
	OtherNames   *datatypes.JSON // 其他名称，JSON 数组格式存储
	BirthDate    *time.Time      // 出生日期
	BirthPlace   *string         `gorm:"size:255"` // 出生地
	Height       *int            // 身高
	Weight       *int            // 体重
	Measurements *datatypes.JSON // 测量数据，JSON 格式存储
	Cup          *string         `gorm:"size:2"` // 杯型
	BloodGroup   *string         `gorm:"size:2"` // 血型
	DebutDate    *time.Time      // 出道日期
	Hobbies      *string         `gorm:"size:255"` // 爱好
	Notes        *string         // 备注
	Videos       []*Video        `gorm:"many2many:actress_video"` // 出演的视频
}
