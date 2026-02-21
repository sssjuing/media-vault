package repository

import (
	"encoding/json"

	"gorm.io/gorm"
)

func paginate(page, pageSize int) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		if page == 0 {
			page = 1
		}
		if pageSize == 0 {
			pageSize = 10
		}
		return db.Limit(pageSize).Offset((page - 1) * pageSize)
	}
}

func videoTagsFilter(tags []string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		jsonTags, _ := json.Marshal(tags)
		return db.Where("tags @> ?", string(jsonTags))
	}
}

func searchVideos(keyword string) func(db *gorm.DB) *gorm.DB {
	return func(db *gorm.DB) *gorm.DB {
		kw := "%" + keyword + "%"
		return db.Where("serial_number LIKE ? OR title LIKE ? OR chinese_title LIKE ?", kw, kw, kw)
	}
}
