package model

import (
	"time"

	"gorm.io/gorm"
)

type Measurements struct {
	Bust  *int // 胸围
	Waist *int // 腰围
	Hips  *int // 臀围
}

type Actress struct {
	gorm.Model
	Names        []*ActressName `gorm:"foreignKey:ActressID"` // 演员的所有名字
	BirthDate    *time.Time     // 出生日期
	BirthPlace   *string        `gorm:"size:255"` // 出生地
	Height       *int           // 身高
	Weight       *int           // 体重
	Measurements Measurements   `gorm:"embedded"` // 三围数据
	Cup          *string        `gorm:"size:2"`   // 杯型
	BloodGroup   *string        `gorm:"size:2"`   // 血型
	DebutDate    *time.Time     // 出道日期
	Hobbies      *string        `gorm:"size:255"` // 爱好
	Notes        *string        // 备注
}
