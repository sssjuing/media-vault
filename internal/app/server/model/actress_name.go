package model

import (
	"gorm.io/gorm"
)

type ActressName struct {
	gorm.Model
	ActressID   uint     `gorm:"not null;index"`
	Actress     *Actress `gorm:"foreignKey:ActressID"`
	Name        string   `gorm:"size:128;uniqueIndex;not null"`
	ChineseName *string  `gorm:"size:128"`
	EnglishName *string  `gorm:"size:128"` // 英文名
	Videos      []*Video `gorm:"many2many:actress_name_video"`
}
