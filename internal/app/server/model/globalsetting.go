package model

import (
	"time"

	"gorm.io/datatypes"
	"gorm.io/gorm"
)

type GlobalSetting struct {
	ID        bool                        `gorm:"primaryKey;default:true;check:id"` // 固定为 true，确保只有一行
	CreatedAt *time.Time                  `json:"created_at"`
	UpdatedAt *time.Time                  `json:"updated_at"`
	DeletedAt gorm.DeletedAt              `gorm:"index" json:"-"`
	VideoTags datatypes.JSONSlice[string] `gorm:"type:jsonb;not null;default:'[]'" json:"video_tags"`
}
