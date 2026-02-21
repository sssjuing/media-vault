package db

import (
	"github.com/samber/lo"
	"github.com/sssjuing/media-vault/internal/app/server/model"
	"gorm.io/driver/postgres"
	"gorm.io/gorm"
	"gorm.io/plugin/dbresolver"
)

func NewDB(dsn string, replicas []string) *gorm.DB {
	db, err := gorm.Open(postgres.Open(dsn), &gorm.Config{})

	if err != nil {
		panic("failed to connect database")
	}

	// 配置读写分离
	db.Use(dbresolver.Register(dbresolver.Config{
		Sources: []gorm.Dialector{
			postgres.Open(dsn), // 主库
		},
		Replicas: lo.Map(replicas, func(item string, index int) gorm.Dialector {
			return postgres.Open(item)
		}),
		Policy: dbresolver.RandomPolicy{}, // 负载均衡策略
	}))

	return db
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Actress{},
		&model.Video{},
		&model.VideoTag{},
	)
}
