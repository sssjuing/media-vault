package db

import (
	"fmt"

	"github.com/sssjuing/media-vault/internal/app/server/model"
	"gorm.io/driver/postgres"
	"gorm.io/driver/sqlite"
	"gorm.io/gorm"
)

type DatabaseType string

const (
	DatabaseTypeSQLite   DatabaseType = "sqlite"
	DatabaseTypePostgres DatabaseType = "postgres"
)

func NewDB(dbType DatabaseType, dsn string) (*gorm.DB, error) {
	var dialector gorm.Dialector

	switch dbType {
	case DatabaseTypeSQLite:
		dialector = sqlite.Open(dsn)
	case DatabaseTypePostgres:
		dialector = postgres.Open(dsn)
	default:
		return nil, fmt.Errorf("unsupported database type: %s", dbType)
	}

	db, err := gorm.Open(dialector, &gorm.Config{})
	if err != nil {
		return nil, err
	}

	return db, nil
}

func AutoMigrate(db *gorm.DB) error {
	return db.AutoMigrate(
		&model.Actress{},
		&model.ActressNameSet{},
		&model.ActressName{},
		&model.Video{},
		&model.GlobalSetting{},
	)
}
