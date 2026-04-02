package model

import (
	"time"

	"gorm.io/gorm"
)

// BWH 三围数据
type BWH struct {
	Bust  *int `json:"bust"`  // 胸围（单位：厘米）
	Waist *int `json:"waist"` // 腰围（单位：厘米）
	Hips  *int `json:"hips"`  // 臀围（单位：厘米）
}

// Actress 演员模型
type Actress struct {
	gorm.Model
	UniqueName string            `gorm:"size:128;uniqueIndex;not null"` // 演员的唯一名称，用于标识演员
	NameSets   []*ActressNameSet `gorm:"foreignKey:ActressID"`          // 演员的所有名字集（一套名字包含多种语言版本，如日文名+中文名+英文名）
	BirthDate  *time.Time        ``                                     // 出生日期
	BirthPlace *string           `gorm:"size:255"`                      // 出生地
	Height     *int              ``                                     // 身高（单位：厘米）
	Weight     *int              ``                                     // 体重（单位：公斤）
	BWH        *BWH              `gorm:"embedded"`                      // 三围数据
	Cup        *string           `gorm:"size:2"`                        // 杯型
	BloodGroup *string           `gorm:"size:2"`                        // 血型
	DebutDate  *time.Time        ``                                     // 出道日期
	Hobbies    *string           `gorm:"size:255"`                      // 爱好
	Notes      *string           ``                                     // 备注
}
