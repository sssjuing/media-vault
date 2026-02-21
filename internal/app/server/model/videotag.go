package model

import "gorm.io/gorm"

type VideoTag struct {
	gorm.Model
	Name string  `gorm:"type:varchar(128);uniqueIndex;not null" json:"name"`
	Rank float64 `gorm:"type:float;default:0;not null" json:"rank"`
}
