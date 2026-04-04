//go:generate ../../../../bin/gen-requests -exclude BWH -output ../handler/actress_requests_gen.go actress.go

package model

import "time"

// BWH 三围数据
type BWH struct {
	Bust  *int `json:"bust"`  // 胸围（单位：厘米）
	Waist *int `json:"waist"` // 腰围（单位：厘米）
	Hips  *int `json:"hips"`  // 臀围（单位：厘米）
}

// Actress 演员模型
type Actress struct {
	BaseModel
	UniqueName string            `gorm:"size:128;uniqueIndex;not null" json:"unique_name"` // 演员的唯一名称，用于标识演员
	NameSets   []*ActressNameSet `gorm:"foreignKey:ActressID" json:"name_sets"`            // 演员的所有名字集（一套名字包含多种语言版本，如日文名+中文名+英文名）
	BirthDate  *time.Time        `json:"birth_date"`                                       // 出生日期
	BirthPlace *string           `gorm:"size:255" json:"birth_place"`                      // 出生地
	Height     *int              `json:"height"`                                           // 身高（单位：厘米）
	Weight     *int              `json:"weight"`                                           // 体重（单位：公斤）
	BWH        *BWH              `gorm:"embedded" json:"bwh"`                              // 三围数据
	Cup        *string           `gorm:"size:2" json:"cup"`                                // 杯型
	BloodGroup *string           `gorm:"size:2" json:"blood_group"`                        // 血型
	DebutDate  *time.Time        `json:"debut_date"`                                       // 出道日期
	Hobbies    *string           `gorm:"size:255" json:"hobbies"`                          // 爱好
	Notes      *string           `json:"notes"`                                            // 备注
}
