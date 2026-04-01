package model

import (
	"gorm.io/gorm"
)

// NameType 名字类型，用于标识不同语言或类型的名字
type NameType string

const (
	NameTypeJA NameType = "ja" // Japanese
	NameTypeZH NameType = "zh" // Chinese
	NameTypeEN NameType = "en" // English
)

// ActressName 演员单个名字，属于某个名字集
type ActressName struct {
	gorm.Model
	NameSetID uint            `gorm:"not null;index:idx_name_set_id"`           // 所属的名字集ID
	NameSet   *ActressNameSet `gorm:"foreignKey:NameSetID"`                     // 所属的名字集
	Name      string          `gorm:"size:128;not null;index:idx_actress_name"` // 名字内容
	NameType  NameType        `gorm:"size:32;not null;index:idx_name_type"`     // 名字类型（如：日文名、中文名、英文名等）
}

// ActressNameSet 演员名字集，一套完整的名字（包含多种语言版本），一个演员可以有多套名字
type ActressNameSet struct {
	gorm.Model
	ActressID uint           `gorm:"not null;index:idx_name_set_actress"` // 关联的演员ID
	Actress   *Actress       `gorm:"foreignKey:ActressID"`                // 关联的演员
	Names     []*ActressName `gorm:"foreignKey:NameSetID"`                // 该套名字包含的所有语言版本的名字
	Videos    []*Video       `gorm:"many2many:video_actress_name_set"`    // 使用了这套名字的视频
}
