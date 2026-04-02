package model

// NameType 名字类型，用于标识不同语言或类型的名字
type NameType string

const (
	NameTypeJA NameType = "ja" // Japanese
	NameTypeZH NameType = "zh" // Chinese
	NameTypeEN NameType = "en" // English
)

// ActressName 演员单个名字，属于某个名字集
type ActressName struct {
	BaseModel
	NameSetID uint            `gorm:"not null;index:idx_name_set_id" json:"-"`               // 所属的名字集ID
	NameSet   *ActressNameSet `gorm:"foreignKey:NameSetID" json:"-"`                         // 所属的名字集
	Name      string          `gorm:"size:128;not null;index:idx_actress_name" json:"name"`  // 名字内容
	NameType  NameType        `gorm:"size:32;not null;index:idx_name_type" json:"name_type"` // 名字类型（如：日文名、中文名、英文名等）
}

// ActressNameSet 演员名字集，一套完整的名字（包含多种语言版本），一个演员可以有多套名字
type ActressNameSet struct {
	BaseModel
	ActressID uint           `gorm:"not null;index:idx_name_set_actress" json:"-"`   // 关联的演员ID
	Actress   *Actress       `gorm:"foreignKey:ActressID" json:"actress"`            // 关联的演员
	Names     []*ActressName `gorm:"foreignKey:NameSetID" json:"names"`              // 该套名字包含的所有语言版本的名字
	Videos    []*Video       `gorm:"many2many:video_actress_name_set" json:"videos"` // 使用了这套名字的视频
}
